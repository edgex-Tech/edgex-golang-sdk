package order

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/edgex-Tech/edgex-golang-sdk/sdk/internal"
	metadatapkg "github.com/edgex-Tech/edgex-golang-sdk/sdk/metadata"
	"github.com/shopspring/decimal"
)

// Client represents the new order client without OpenAPI dependencies
type Client struct {
	*internal.Client
}

// NewClient creates a new order client
func NewClient(client *internal.Client) *Client {
	return &Client{
		Client: client,
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

	// Find contract from metadata
	var contract *metadatapkg.Contract
	if metadata != nil && metadata.ContractList != nil {
		for i := range metadata.ContractList {
			if metadata.ContractList[i].ContractId == params.ContractId {
				contract = &metadata.ContractList[i]
				break
			}
		}
	}

	if contract == nil {
		return nil, fmt.Errorf("contract not found: %s", params.ContractId)
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
		return nil, fmt.Errorf("coin not found: %s", contract.QuoteCoinId)
	}

	syntheticFactorBig, err := internal.HexToBigInteger(contract.StarkExResolution)
	if err != nil {
		return nil, fmt.Errorf("failed to parse synthetic factor: %w", err)
	}
	syntheticFactor := decimal.NewFromBigInt(syntheticFactorBig, 0)

	shiftFactorBig, err := internal.HexToBigInteger(quoteCoin.StarkExResolution)
	if err != nil {
		return nil, fmt.Errorf("failed to parse shift factor: %w", err)
	}
	shiftFactor := decimal.NewFromBigInt(shiftFactorBig, 0)
	// Parse decimal values
	size, err := decimal.NewFromString(params.Size)
	if err != nil {
		return nil, fmt.Errorf("failed to parse size: %w", err)
	}

	// Calculate values
	valueDm := l2Price.Mul(size)

	amountSynthetic := size.Mul(syntheticFactor).IntPart()
	amountCollateral := valueDm.Mul(shiftFactor).IntPart()

	// Get fee rate from contract or use default
	var feeRate decimal.Decimal
	if contract.DefaultTakerFeeRate != "" {
		feeRateVal, err := decimal.NewFromString(contract.DefaultTakerFeeRate)
		if err != nil {
			return nil, fmt.Errorf("failed to parse fee rate: %w", err)
		}
		feeRate = feeRateVal
	} else {
		feeRate, _ = decimal.NewFromString("0.001") // Default fee rate
	}

	// Calculate fee amount in decimal with ceiling to integer
	limitFee := size.Mul(l2Price).Mul(feeRate).Ceil()
	maxAmountFee := limitFee.Mul(shiftFactor)

	clientOrderId := internal.GetRandomClientId()

	nonce := internal.CalcNonce(clientOrderId)
	l2ExpireTime := params.ExpireTime.Add(time.Hour * 9 * 24).UnixMilli()
	l2ExpireHour := l2ExpireTime / (60 * 60 * 1000)

	msgHash := internal.CalcLimitOrderHash(
		contract.StarkExSyntheticAssetId,
		quoteCoin.StarkExAssetId,
		quoteCoin.StarkExAssetId,
		params.Side == "BUY",
		amountSynthetic,
		amountCollateral,
		maxAmountFee.BigInt().Int64(),
		nonce,
		c.Client.GetAccountID(),
		l2ExpireHour,
	)
	signature, err := c.Client.Sign(msgHash)
	if err != nil {
		return nil, fmt.Errorf("failed to sign withdrawal hash: %w", err)
	}
	sig_str := fmt.Sprintf("%s%s%s", signature.R, signature.S, signature.V)

	// Build request body
	body := map[string]interface{}{
		"accountId":     strconv.FormatInt(c.Client.GetAccountID(), 10),
		"contractId":    params.ContractId,
		"price":         params.Price,
		"size":          params.Size,
		"type":          string(params.Type),
		"side":          params.Side,
		"timeInForce":   params.TimeInForce,
		"clientOrderId": clientOrderId,
		"expireTime":    strconv.FormatInt(params.ExpireTime.UnixMilli(), 10),
		"l2Nonce":       strconv.FormatInt(nonce, 10),
		"l2Signature":   sig_str,
		"l2ExpireTime":  strconv.FormatInt(l2ExpireTime, 10),
		"l2Value":       valueDm.String(),
		"l2Size":        params.Size,
		"l2LimitFee":    limitFee.String(),
		"reduceOnly":    params.ReduceOnly,
	}

	url := fmt.Sprintf("%s/api/v1/private/order/createOrder", c.Client.GetBaseURL())
	resp, err := c.Client.HttpRequest(url, "POST", body, nil)
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

// CancelOrder cancels a specific order
func (c *Client) CancelOrder(ctx context.Context, params *CancelOrderParams) (interface{}, error) {
	var url string
	accountID := strconv.FormatInt(c.Client.GetAccountID(), 10)

	var body map[string]interface{}

	if params.OrderId != "" {
		url = fmt.Sprintf("%s/api/v1/private/order/cancelOrderById", c.Client.GetBaseURL())
		body = map[string]interface{}{
			"accountId":   accountID,
			"orderIdList": []string{params.OrderId},
		}
	} else if params.ClientId != "" {
		url = fmt.Sprintf("%s/api/v1/private/order/cancelOrderByClientOrderId", c.Client.GetBaseURL())
		body = map[string]interface{}{
			"accountId":         accountID,
			"clientOrderIdList": []string{params.ClientId},
		}
	} else if params.ContractId != "" {
		url = "/api/v1/private/order/cancelAllOrder"
		body = map[string]interface{}{
			"accountId":            accountID,
			"filterContractIdList": []string{params.ContractId},
		}
	} else {
		return nil, fmt.Errorf("must provide either OrderId, ClientId, or ContractId")
	}

	resp, err := c.Client.HttpRequest(url, "POST", body, nil)
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
	url := fmt.Sprintf("%s/api/v1/private/order/getActiveOrderPage", c.Client.GetBaseURL())
	queryParams := map[string]string{
		"accountId": strconv.FormatInt(c.Client.GetAccountID(), 10),
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
	if params.FilterIsLiquidate != nil {
		queryParams["filterIsLiquidate"] = strconv.FormatBool(*params.FilterIsLiquidate)
	}
	if params.FilterIsDeleverage != nil {
		queryParams["filterIsDeleverage"] = strconv.FormatBool(*params.FilterIsDeleverage)
	}
	if params.FilterIsPositionTpsl != nil {
		queryParams["filterIsPositionTpsl"] = strconv.FormatBool(*params.FilterIsPositionTpsl)
	}
	if params.FilterStartCreatedTimeInclusive > 0 {
		queryParams["filterStartCreatedTimeInclusive"] = strconv.FormatUint(params.FilterStartCreatedTimeInclusive, 10)
	}
	if params.FilterEndCreatedTimeExclusive > 0 {
		queryParams["filterEndCreatedTimeExclusive"] = strconv.FormatUint(params.FilterEndCreatedTimeExclusive, 10)
	}

	resp, err := c.Client.HttpRequest(url, "GET", nil, queryParams)
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

// GetOrderFillTransactions gets order fill transactions with pagination and filters
func (c *Client) GetOrderFillTransactions(ctx context.Context, params *OrderFillTransactionParams) (*ResultPageDataOrderFillTransaction, error) {
	url := fmt.Sprintf("%s/api/v1/private/order/getHistoryOrderFillTransactionPage", c.Client.GetBaseURL())
	queryParams := map[string]string{
		"accountId": strconv.FormatInt(c.Client.GetAccountID(), 10),
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
	if params.FilterIsLiquidate != nil {
		queryParams["filterIsLiquidate"] = strconv.FormatBool(*params.FilterIsLiquidate)
	}
	if params.FilterIsDeleverage != nil {
		queryParams["filterIsDeleverage"] = strconv.FormatBool(*params.FilterIsDeleverage)
	}
	if params.FilterIsPositionTpsl != nil {
		queryParams["filterIsPositionTpsl"] = strconv.FormatBool(*params.FilterIsPositionTpsl)
	}
	if params.FilterStartCreatedTimeInclusive > 0 {
		queryParams["filterStartCreatedTimeInclusive"] = strconv.FormatUint(params.FilterStartCreatedTimeInclusive, 10)
	}
	if params.FilterEndCreatedTimeExclusive > 0 {
		queryParams["filterEndCreatedTimeExclusive"] = strconv.FormatUint(params.FilterEndCreatedTimeExclusive, 10)
	}

	resp, err := c.Client.HttpRequest(url, "GET", nil, queryParams)
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

	url := fmt.Sprintf("%s/api/v1/private/order/getOrderById", c.Client.GetBaseURL())
	queryParams := map[string]string{
		"accountId":   strconv.FormatInt(c.Client.GetAccountID(), 10),
		"orderIdList": strings.Join(orderIDs, ","),
	}

	resp, err := c.Client.HttpRequest(url, "GET", nil, queryParams)
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

	url := fmt.Sprintf("%s/api/v1/private/order/getOrderByClientOrderId", c.Client.GetBaseURL())
	queryParams := map[string]string{
		"accountId":         strconv.FormatInt(c.Client.GetAccountID(), 10),
		"clientOrderIdList": strings.Join(clientOrderIDs, ","),
	}

	resp, err := c.Client.HttpRequest(url, "GET", nil, queryParams)
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
	url := fmt.Sprintf("%s/api/v1/private/order/getMaxCreateOrderSize", c.Client.GetBaseURL())
	queryParams := map[string]interface{}{
		"accountId":  strconv.FormatInt(c.Client.GetAccountID(), 10),
		"contractId": contractID,
		"price":      price.String(),
	}

	resp, err := c.Client.HttpRequest(url, "POST", queryParams, nil)
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
