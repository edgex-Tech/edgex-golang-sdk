package order

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/edgex-Tech/edgex-golang-sdk/v2/sdk/internal"
	metadatapkg "github.com/edgex-Tech/edgex-golang-sdk/v2/sdk/metadata"
	"github.com/shopspring/decimal"
)

// Client represents the new order client without OpenAPI dependencies
type Client struct {
	c clientInterface
}

type clientInterface interface {
	GetAccountID() int64
	GetSignerPriKey() string
	GetBaseURL() string
	ResolveSignerAddress() (string, error)
	SignTypedDataWithSignerKey(typedData internal.TypedData) (string, error)
	HttpRequest(urlStr string, method string, data map[string]interface{}, params map[string]string) (*http.Response, error)
}

// NewClient creates a new order client
func NewClient(client clientInterface) *Client {
	return &Client{
		c: client,
	}
}

// CreateOrder creates a new order with the given parameters
func (c *Client) CreateOrder(ctx context.Context, params *CreateOrderParams, metadata *metadatapkg.MetaData, l2Price decimal.Decimal) (*ResultCreateOrder, error) {
	// Set default TimeInForce based on order type if not specified
	if params.TimeInForce == "" {
		switch params.Type {
		case OrderTypeMarket:
			params.TimeInForce = string(TimeInForce_IMMEDIATE_OR_CANCEL)
		case OrderTypeLimit:
			params.TimeInForce = string(TimeInForce_GOOD_TIL_CANCEL)
		}
	}

	contract, quoteCoin, err := resolveOrderContractAndQuoteCoin(metadata, params.ContractId)
	if err != nil {
		return nil, err
	}

	// V2 only - use EIP-712 signature
	return c.createOrderV2(ctx, params, metadata, contract, quoteCoin, l2Price)
}

func resolveOrderContractAndQuoteCoin(metadata *metadatapkg.MetaData, contractID string) (*metadatapkg.Contract, *metadatapkg.Coin, error) {
	var contract *metadatapkg.Contract
	if metadata != nil && metadata.ContractList != nil {
		for i := range metadata.ContractList {
			if metadata.ContractList[i].ContractId == contractID {
				contract = &metadata.ContractList[i]
				break
			}
		}
	}
	if contract == nil {
		return nil, nil, fmt.Errorf("contract not found: %s", contractID)
	}

	var quoteCoin *metadatapkg.Coin
	if metadata != nil && metadata.CoinList != nil {
		for i := range metadata.CoinList {
			if metadata.CoinList[i].CoinId == contract.QuoteCoinId {
				quoteCoin = &metadata.CoinList[i]
				break
			}
		}
	}
	if quoteCoin == nil {
		return nil, nil, fmt.Errorf("coin not found: %s", contract.QuoteCoinId)
	}

	return contract, quoteCoin, nil
}

func parseResolutionDecimal(resolution string, starkExResolution string) (decimal.Decimal, error) {
	candidates := []string{
		strings.TrimSpace(resolution),
		strings.TrimSpace(starkExResolution),
	}

	var parseErrors []string
	for _, candidate := range candidates {
		if candidate == "" {
			continue
		}

		lower := strings.ToLower(candidate)
		if strings.HasPrefix(lower, "0x") {
			resolutionBig, err := internal.HexToBigInteger(candidate)
			if err != nil {
				parseErrors = append(parseErrors, fmt.Sprintf("hex(%s): %v", candidate, err))
				continue
			}
			return decimal.NewFromBigInt(resolutionBig, 0), nil
		}

		res, err := decimal.NewFromString(candidate)
		if err != nil {
			parseErrors = append(parseErrors, fmt.Sprintf("decimal(%s): %v", candidate, err))
			continue
		}
		return res, nil
	}

	if len(parseErrors) == 0 {
		return decimal.Zero, fmt.Errorf("resolution is empty")
	}
	return decimal.Zero, fmt.Errorf("failed to parse resolution: %s", strings.Join(parseErrors, "; "))
}

func parseMaxFeeRate(contract *metadatapkg.Contract) (decimal.Decimal, error) {
	if contract == nil {
		return decimal.Zero, fmt.Errorf("contract is nil")
	}

	defaultRate, _ := decimal.NewFromString("0.001")
	takerRate := defaultRate
	if strings.TrimSpace(contract.DefaultTakerFeeRate) != "" {
		parsed, err := decimal.NewFromString(contract.DefaultTakerFeeRate)
		if err != nil {
			return decimal.Zero, fmt.Errorf("failed to parse default taker fee rate: %w", err)
		}
		takerRate = parsed
	}

	makerRate := takerRate
	if strings.TrimSpace(contract.DefaultMakerFeeRate) != "" {
		parsed, err := decimal.NewFromString(contract.DefaultMakerFeeRate)
		if err != nil {
			return decimal.Zero, fmt.Errorf("failed to parse default maker fee rate: %w", err)
		}
		makerRate = parsed
	}

	if makerRate.GreaterThan(takerRate) {
		return makerRate, nil
	}
	return takerRate, nil
}

func (c *Client) createOrderV2(ctx context.Context, params *CreateOrderParams, metadata *metadatapkg.MetaData, contract *metadatapkg.Contract, quoteCoin *metadatapkg.Coin, l2Price decimal.Decimal) (*ResultCreateOrder, error) {
	if contract == nil || quoteCoin == nil {
		return nil, fmt.Errorf("contract/quote coin is nil")
	}

	if c.c.GetSignerPriKey() == "" {
		return nil, fmt.Errorf("trading private key is required for v2 EIP-712 order signing")
	}

	size, err := decimal.NewFromString(params.Size)
	if err != nil {
		return nil, fmt.Errorf("failed to parse size: %w", err)
	}

	// For conditional orders (STOP_MARKET, TAKE_PROFIT_MARKET) with triggerPrice,
	// use trigger price for l2Value calculation when price is 0
	if l2Price.IsZero() && params.TriggerPrice != "" {
		triggerPriceDec, err := decimal.NewFromString(params.TriggerPrice)
		if err == nil && triggerPriceDec.GreaterThan(decimal.Zero) {
			l2Price = triggerPriceDec
		}
	}

	l2Value := l2Price.Mul(size)

	feeRate, err := parseMaxFeeRate(contract)
	if err != nil {
		return nil, err
	}
	limitFee := l2Value.Mul(feeRate).Ceil()

	contractResolution, err := parseResolutionDecimal(contract.Resolution, contract.StarkExResolution)
	if err != nil {
		return nil, fmt.Errorf("failed to parse contract resolution: %w", err)
	}
	quoteResolution, err := parseResolutionDecimal(quoteCoin.Resolution, quoteCoin.StarkExResolution)
	if err != nil {
		return nil, fmt.Errorf("failed to parse quote coin resolution: %w", err)
	}

	amountSynthetic := size.Mul(contractResolution).BigInt()
	amountCollateral := l2Value.Mul(quoteResolution).BigInt()
	amountFee := limitFee.Mul(quoteResolution).BigInt()

	clientOrderID := internal.GetRandomClientId()
	if params.ClientOrderId != nil && strings.TrimSpace(*params.ClientOrderId) != "" {
		clientOrderID = *params.ClientOrderId
	}
	l2Nonce := internal.CalcNonce(clientOrderID)

	nowMillis := time.Now().UnixMilli()
	const orderExpireWindowMillis int64 = 30 * 24 * 60 * 60 * 1000
	const l1OffsetMillis int64 = 8 * 24 * 60 * 60 * 1000

	l2ExpireTime := nowMillis + orderExpireWindowMillis
	expireTime := l2ExpireTime - l1OffsetMillis
	if !params.ExpireTime.IsZero() {
		expireTime = params.ExpireTime.UnixMilli()
		l2ExpireTime = expireTime + l1OffsetMillis
	}

	tradingSigner, err := c.c.ResolveSignerAddress()
	if err != nil {
		return nil, fmt.Errorf("failed to resolve trading signer address: %w", err)
	}

	chainID := ""
	verifyingContract := ""
	if metadata != nil && metadata.Global != nil {
		chainID = strings.TrimSpace(metadata.Global.NativeChainId)
		if chainID == "" {
			chainID = strings.TrimSpace(metadata.Global.ChainId)
		}
		verifyingContract = strings.TrimSpace(metadata.Global.ContractAddress)
	}

	if chainID == "" {
		return nil, fmt.Errorf("metadata.global.nativeChainId/chainId is required for v2 order signing")
	}
	if verifyingContract == "" {
		return nil, fmt.Errorf("metadata.global.contractAddress is required for v2 order signing")
	}

	typedDomain, err := internal.NewTypedDataDomain("EdgeX", "1", chainID, verifyingContract)
	if err != nil {
		return nil, fmt.Errorf("failed to build typed data domain: %w", err)
	}

	typedData := internal.TypedData{
		Types: internal.TypedDataTypes{
			"EIP712Domain": []internal.TypedDataType{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
			"OrderBase": []internal.TypedDataType{
				{Name: "nonce", Type: "uint256"},
				{Name: "signer", Type: "address"},
				{Name: "accountId", Type: "uint64"},
				{Name: "expirationTimestamp", Type: "uint256"},
			},
			"LimitOrderParams": []internal.TypedDataType{
				{Name: "base", Type: "OrderBase"},
				{Name: "amountSynthetic", Type: "int256"},
				{Name: "amountCollateral", Type: "int256"},
				{Name: "amountFee", Type: "uint256"},
				{Name: "assetIdSynthetic", Type: "uint64"},
				{Name: "assetIdCollateral", Type: "uint64"},
				{Name: "isBuyingSynthetic", Type: "bool"},
				{Name: "parentOrderHash", Type: "bytes32"},
				{Name: "subOrderIndex", Type: "uint256"},
			},
		},
		PrimaryType: "LimitOrderParams",
		Domain:      typedDomain,
		Message: internal.TypedDataMessage{
			"base": map[string]interface{}{
				"nonce":               strconv.FormatInt(l2Nonce, 10),
				"signer":              tradingSigner,
				"accountId":           strconv.FormatInt(c.c.GetAccountID(), 10),
				"expirationTimestamp": strconv.FormatInt(expireTime/1000, 10),
			},
			"amountSynthetic":  amountSynthetic.String(),
			"amountCollateral": amountCollateral.String(),
			"amountFee":        amountFee.String(),
			"assetIdSynthetic": contract.ContractId,
			"assetIdCollateral": func() string {
				if quoteCoin.CoinId != "" {
					return quoteCoin.CoinId
				}
				return contract.QuoteCoinId
			}(),
			"isBuyingSynthetic": params.Side == OrderSideBuy,
			"parentOrderHash":   "0x0000000000000000000000000000000000000000000000000000000000000000",
			"subOrderIndex":     "0",
		},
	}

	l2Signature, err := c.c.SignTypedDataWithSignerKey(typedData)
	if err != nil {
		return nil, fmt.Errorf("failed to sign v2 order typed data: %w", err)
	}

	price := params.Price
	if params.Type == OrderTypeMarket || params.Type == OrderTypeStopMarket || params.Type == OrderTypeTakeProfitMarket {
		price = "0"
	}

	body := map[string]interface{}{
		"accountId":     strconv.FormatInt(c.c.GetAccountID(), 10),
		"contractId":    params.ContractId,
		"price":         price,
		"size":          params.Size,
		"type":          string(params.Type),
		"side":          params.Side,
		"timeInForce":   params.TimeInForce,
		"clientOrderId": clientOrderID,
		"expireTime":    strconv.FormatInt(expireTime, 10),
		"l2Nonce":       strconv.FormatInt(l2Nonce, 10),
		"l2Signature":   l2Signature,
		"l2ExpireTime":  strconv.FormatInt(l2ExpireTime, 10),
		"l2Value":       l2Value.String(),
		"l2Size":        params.Size,
		"l2LimitFee":    limitFee.String(),
		"reduceOnly":    params.ReduceOnly,
	}

	// Add conditional order fields (for STOP_LIMIT, STOP_MARKET, TAKE_PROFIT_LIMIT, TAKE_PROFIT_MARKET)
	if params.TriggerPrice != "" {
		body["triggerPrice"] = params.TriggerPrice
	}
	if params.TriggerPriceType != "" {
		body["triggerPriceType"] = params.TriggerPriceType
	}

	// Add position TP/SL fields
	if params.IsPositionTpsl {
		body["isPositionTpsl"] = params.IsPositionTpsl
	}
	if params.OpenTpslParentOrderId != "" {
		body["openTpslParentOrderId"] = params.OpenTpslParentOrderId
	}

	// Add open TP/SL on new orders
	if params.IsSetOpenTp {
		body["isSetOpenTp"] = params.IsSetOpenTp
		if params.OpenTp != nil {
			body["openTp"] = params.OpenTp
		}
	}
	if params.IsSetOpenSl {
		body["isSetOpenSl"] = params.IsSetOpenSl
		if params.OpenSl != nil {
			body["openSl"] = params.OpenSl
		}
	}

	// Add additional fields
	if params.SourceKey != "" {
		body["sourceKey"] = params.SourceKey
	}
	if params.ExtraType != "" {
		body["extraType"] = params.ExtraType
	}
	if params.ExtraDataJson != "" {
		body["extraDataJson"] = params.ExtraDataJson
	}

	return c.postCreateOrder(ctx, body)
}

func (c *Client) postCreateOrder(ctx context.Context, body map[string]interface{}) (*ResultCreateOrder, error) {
	_ = ctx
	url := fmt.Sprintf("%s/api/v2/private/order/createOrder", c.c.GetBaseURL())
	resp, err := c.c.HttpRequest(url, "POST", body, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create order: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result ResultCreateOrder
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}

func joinOptionalBoolFilter(value *bool) string {
	if value == nil {
		return ""
	}
	return strconv.FormatBool(*value)
}

// CancelOrder cancels a specific order
func (c *Client) CancelOrder(ctx context.Context, params *CancelOrderParams) (interface{}, error) {
	var url string
	accountID := strconv.FormatInt(c.c.GetAccountID(), 10)

	var body map[string]interface{}

	if params.OrderId != "" {
		url = fmt.Sprintf("%s/api/v2/private/order/cancelOrderById", c.c.GetBaseURL())
		body = map[string]interface{}{
			"accountId":   accountID,
			"orderIdList": []string{params.OrderId},
		}
	} else if params.ClientId != "" {
		url = fmt.Sprintf("%s/api/v2/private/order/cancelOrderByClientOrderId", c.c.GetBaseURL())
		body = map[string]interface{}{
			"accountId":         accountID,
			"clientOrderIdList": []string{params.ClientId},
		}
	} else if params.ContractId != "" {
		url = fmt.Sprintf("%s/api/v2/private/order/cancelAllOrder", c.c.GetBaseURL())
		body = map[string]interface{}{
			"accountId":            accountID,
			"filterContractIdList": []string{params.ContractId},
		}
	} else {
		return nil, fmt.Errorf("must provide either OrderId, ClientId, or ContractId")
	}

	resp, err := c.c.HttpRequest(url, "POST", body, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to cancel order: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if code, ok := result["code"].(string); ok && code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", code)
	}

	return result, nil
}

// GetActiveOrders gets active orders with pagination and filters
func (c *Client) GetActiveOrders(ctx context.Context, params *GetActiveOrderParams) (*ResultPageDataOrder, error) {
	url := fmt.Sprintf("%s/api/v2/private/order/getActiveOrderPage", c.c.GetBaseURL())
	queryParams := map[string]string{
		"accountId": strconv.FormatInt(c.c.GetAccountID(), 10),
	}

	if params.Size != "" {
		queryParams["size"] = params.Size
	}
	if params.OffsetData != "" {
		queryParams["offsetData"] = params.OffsetData
	}

	if len(params.FilterCoinIdList) > 0 {
		queryParams["filterCoinIdList"] = strings.Join(params.FilterCoinIdList, ",")
	}
	if len(params.FilterContractIdList) > 0 {
		queryParams["filterContractIdList"] = strings.Join(params.FilterContractIdList, ",")
	}
	if len(params.FilterTypeList) > 0 {
		queryParams["filterTypeList"] = strings.Join(params.FilterTypeList, ",")
	}
	if len(params.FilterStatusList) > 0 {
		queryParams["filterStatusList"] = strings.Join(params.FilterStatusList, ",")
	}
	if value := joinOptionalBoolFilter(params.FilterIsLiquidate); value != "" {
		queryParams["filterIsLiquidateList"] = value
	}
	if value := joinOptionalBoolFilter(params.FilterIsDeleverage); value != "" {
		queryParams["filterIsDeleverageList"] = value
	}
	if value := joinOptionalBoolFilter(params.FilterIsPositionTpsl); value != "" {
		queryParams["filterIsPositionTpslList"] = value
	}
	if params.FilterStartCreatedTimeInclusive > 0 {
		queryParams["filterStartCreatedTimeInclusive"] = strconv.FormatUint(params.FilterStartCreatedTimeInclusive, 10)
	}
	if params.FilterEndCreatedTimeExclusive > 0 {
		queryParams["filterEndCreatedTimeExclusive"] = strconv.FormatUint(params.FilterEndCreatedTimeExclusive, 10)
	}

	resp, err := c.c.HttpRequest(url, "GET", nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get active orders: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result ResultPageDataOrder
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}

// GetHistoryOrderPage gets historical orders with pagination and filters.
func (c *Client) GetHistoryOrderPage(ctx context.Context, params *GetHistoryOrderParams) (*ResultPageDataOrder, error) {
	url := fmt.Sprintf("%s/api/v2/private/order/getHistoryOrderPage", c.c.GetBaseURL())
	queryParams := map[string]string{
		"accountId": strconv.FormatInt(c.c.GetAccountID(), 10),
	}

	if params != nil {
		if params.Size != "" {
			queryParams["size"] = params.Size
		}
		if params.OffsetData != "" {
			queryParams["offsetData"] = params.OffsetData
		}

		if len(params.FilterCoinIdList) > 0 {
			queryParams["filterCoinIdList"] = strings.Join(params.FilterCoinIdList, ",")
		}
		if len(params.FilterContractIdList) > 0 {
			queryParams["filterContractIdList"] = strings.Join(params.FilterContractIdList, ",")
		}
		if len(params.FilterTypeList) > 0 {
			queryParams["filterTypeList"] = strings.Join(params.FilterTypeList, ",")
		}
		if len(params.FilterStatusList) > 0 {
			queryParams["filterStatusList"] = strings.Join(params.FilterStatusList, ",")
		}
		if value := joinOptionalBoolFilter(params.FilterIsLiquidate); value != "" {
			queryParams["filterIsLiquidateList"] = value
		}
		if value := joinOptionalBoolFilter(params.FilterIsDeleverage); value != "" {
			queryParams["filterIsDeleverageList"] = value
		}
		if value := joinOptionalBoolFilter(params.FilterIsPositionTpsl); value != "" {
			queryParams["filterIsPositionTpslList"] = value
		}
		if params.FilterStartCreatedTimeInclusive > 0 {
			queryParams["filterStartCreatedTimeInclusive"] = strconv.FormatUint(params.FilterStartCreatedTimeInclusive, 10)
		}
		if params.FilterEndCreatedTimeExclusive > 0 {
			queryParams["filterEndCreatedTimeExclusive"] = strconv.FormatUint(params.FilterEndCreatedTimeExclusive, 10)
		}
	}

	resp, err := c.c.HttpRequest(url, "GET", nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get history orders: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result ResultPageDataOrder
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}

// GetOrderFillTransactions gets order fill transactions with pagination and filters
func (c *Client) GetOrderFillTransactions(ctx context.Context, params *OrderFillTransactionParams) (*ResultPageDataOrderFillTransaction, error) {
	url := fmt.Sprintf("%s/api/v2/private/order/getHistoryOrderFillTransactionPage", c.c.GetBaseURL())
	queryParams := map[string]string{
		"accountId": strconv.FormatInt(c.c.GetAccountID(), 10),
	}

	if params.Size != "" {
		queryParams["size"] = params.Size
	}
	if params.OffsetData != "" {
		queryParams["offsetData"] = params.OffsetData
	}

	if len(params.FilterCoinIdList) > 0 {
		queryParams["filterCoinIdList"] = strings.Join(params.FilterCoinIdList, ",")
	}
	if len(params.FilterContractIdList) > 0 {
		queryParams["filterContractIdList"] = strings.Join(params.FilterContractIdList, ",")
	}
	if len(params.FilterOrderIdList) > 0 {
		queryParams["filterOrderIdList"] = strings.Join(params.FilterOrderIdList, ",")
	}
	if value := joinOptionalBoolFilter(params.FilterIsLiquidate); value != "" {
		queryParams["filterIsLiquidateList"] = value
	}
	if value := joinOptionalBoolFilter(params.FilterIsDeleverage); value != "" {
		queryParams["filterIsDeleverageList"] = value
	}
	if value := joinOptionalBoolFilter(params.FilterIsPositionTpsl); value != "" {
		queryParams["filterIsPositionTpslList"] = value
	}
	if params.FilterStartCreatedTimeInclusive > 0 {
		queryParams["filterStartCreatedTimeInclusive"] = strconv.FormatUint(params.FilterStartCreatedTimeInclusive, 10)
	}
	if params.FilterEndCreatedTimeExclusive > 0 {
		queryParams["filterEndCreatedTimeExclusive"] = strconv.FormatUint(params.FilterEndCreatedTimeExclusive, 10)
	}

	resp, err := c.c.HttpRequest(url, "GET", nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get order fill transactions: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result ResultPageDataOrderFillTransaction
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}

// GetOrdersByID retrieves orders by their order IDs
func (c *Client) GetOrdersByID(ctx context.Context, orderIDs []string) (*ResultListOrder, error) {
	if len(orderIDs) == 0 {
		return nil, fmt.Errorf("order IDs must not be empty")
	}

	url := fmt.Sprintf("%s/api/v2/private/order/getOrderById", c.c.GetBaseURL())
	queryParams := map[string]string{
		"accountId":   strconv.FormatInt(c.c.GetAccountID(), 10),
		"orderIdList": strings.Join(orderIDs, ","),
	}

	resp, err := c.c.HttpRequest(url, "GET", nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get orders by id: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result ResultListOrder
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}

// GetOrdersByClientOrderID retrieves orders by their client order IDs
func (c *Client) GetOrdersByClientOrderID(ctx context.Context, clientOrderIDs []string) (*ResultListOrder, error) {
	if len(clientOrderIDs) == 0 {
		return nil, fmt.Errorf("client order IDs must not be empty")
	}

	url := fmt.Sprintf("%s/api/v2/private/order/getOrderByClientOrderId", c.c.GetBaseURL())
	queryParams := map[string]string{
		"accountId":         strconv.FormatInt(c.c.GetAccountID(), 10),
		"clientOrderIdList": strings.Join(clientOrderIDs, ","),
	}

	resp, err := c.c.HttpRequest(url, "GET", nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get orders by client order id: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result ResultListOrder
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}

// GetMaxOrderSize gets the maximum order size for a given contract and price
func (c *Client) GetMaxOrderSize(ctx context.Context, contractID string, price decimal.Decimal) (*ResultGetMaxCreateOrderSize, error) {
	url := fmt.Sprintf("%s/api/v2/private/order/getMaxCreateOrderSize", c.c.GetBaseURL())
	queryParams := map[string]interface{}{
		"accountId":  strconv.FormatInt(c.c.GetAccountID(), 10),
		"contractId": contractID,
		"price":      price.String(),
	}

	resp, err := c.c.HttpRequest(url, "POST", queryParams, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get max order size: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result ResultGetMaxCreateOrderSize
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %v", result)
	}

	return &result, nil
}
