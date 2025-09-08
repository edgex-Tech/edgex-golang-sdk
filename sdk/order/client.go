package order

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	openapi "github.com/edgex-Tech/edgex-golang-sdk/openapi"
	"github.com/edgex-Tech/edgex-golang-sdk/sdk/internal"
	"github.com/shopspring/decimal"
)

// Client represents the order client
type Client struct {
	*internal.Client
	openapiClient *openapi.APIClient
}

// NewClient creates a new order client
func NewClient(client *internal.Client, openapiClient *openapi.APIClient) *Client {
	return &Client{
		Client:        client,
		openapiClient: openapiClient,
	}
}

// CreateOrder creates a new order with the given parameters
func (c *Client) CreateOrder(ctx context.Context, params *CreateOrderParams, metadata openapi.MetaData) (*openapi.ResultCreateOrder, error) {
	// Set default TimeInForce based on order type if not specified
	if params.TimeInForce == "" {
		switch params.Type {
		case OrderTypeMarket:
			params.TimeInForce = string(TimeInForce_IMMEDIATE_OR_CANCEL)
		case OrderTypeLimit:
			params.TimeInForce = string(TimeInForce_GOOD_TIL_CANCEL)
		}
	}

	// Find the contract from metadata
	var contract *openapi.Contract
	contractList := metadata.GetContractList()
	for i := range contractList {
		if contractList[i].GetContractId() == params.ContractId {
			contract = &contractList[i]
			break
		}
	}
	if contract == nil {
		return nil, fmt.Errorf("contract not found: %s", params.ContractId)
	}

	// Get collateral coin from metadata
	global := metadata.GetGlobal()
	collateralCoin := global.GetStarkExCollateralCoin()

	// Parse decimal values
	size, err := decimal.NewFromString(params.Size)
	if err != nil {
		return nil, fmt.Errorf("failed to parse size: %w", err)
	}

	price, err := decimal.NewFromString(params.Price)
	if err != nil {
		return nil, fmt.Errorf("failed to parse price: %w", err)
	}

	// Convert hex resolution to decimal
	hexResolution := contract.GetStarkExResolution()
	// Remove "0x" prefix if present
	hexResolution = strings.TrimPrefix(hexResolution, "0x")
	// Parse hex string to int64
	resolutionInt, err := strconv.ParseInt(hexResolution, 16, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse hex resolution: %w", err)
	}
	resolution := decimal.NewFromInt(resolutionInt)

	clientOrderId := internal.GenerateUUID()
	if params.ClientOrderId != nil {
		clientOrderId = *params.ClientOrderId
	}

	// Calculate values
	valueDm := price.Mul(size)
	amountSynthetic := size.Mul(resolution).IntPart()
	amountCollateral := valueDm.Shift(6).IntPart()

	// Calculate fee based on order type (maker/taker)
	feeRate, err := decimal.NewFromString(contract.GetDefaultTakerFeeRate())
	if err != nil {
		return nil, fmt.Errorf("failed to parse fee rate: %w", err)
	}

	// Calculate fee amount in decimal with ceiling (向上取整)
	amountFeeDm := valueDm.Mul(feeRate).Ceil()
	amountFeeStr := amountFeeDm.String()

	// Convert to the required integer format for the protocol
	amountFee := amountFeeDm.Shift(6).IntPart()

	nonce := internal.CalcNonce(clientOrderId)
	l2ExpireTime := time.Now().Add(14 * 24 * time.Hour).UnixMilli()

	// Calculate signature using asset IDs from metadata
	expireTimeUnix := l2ExpireTime / (60 * 60 * 1000)
	sigHash := internal.CalcLimitOrderHash(
		contract.GetStarkExSyntheticAssetId(),
		collateralCoin.GetStarkExAssetId(),
		collateralCoin.GetStarkExAssetId(),
		params.Side == OrderSideBuy,
		amountSynthetic,
		amountCollateral,
		amountFee,
		nonce,
		c.Client.GetAccountID(),
		expireTimeUnix,
	)

	// Sign the order
	sig, err := c.Client.Sign(sigHash)
	if err != nil {
		return nil, fmt.Errorf("failed to sign order: %w", err)
	}

	// Convert signature to string
	sigStr := fmt.Sprintf("%s%s%s", sig.R, sig.S, sig.V)

	// Create order request
	accountID := strconv.FormatInt(c.Client.GetAccountID(), 10)
	nonceStr := strconv.FormatInt(nonce, 10)
	l2ExpireTimeStr := strconv.FormatInt(l2ExpireTime, 10)
	expireTimeStr := strconv.FormatInt(l2ExpireTime-864000000, 10)
	valueStr := valueDm.String()

	var price_ string
	if string(params.Type) == string(OrderTypeLimit) {
		price_ = params.Price
	} else {
		price_ = "0"
	}

	req := c.openapiClient.Class04OrderPrivateApiAPI.CreateOrder(ctx).
		CreateOrderParam(openapi.CreateOrderParam{
			AccountId:     &accountID,
			ContractId:    &params.ContractId,
			Price:         &price_,
			Size:          &params.Size,
			Type:          (*string)(&params.Type),
			TimeInForce:   &params.TimeInForce,
			Side:          &params.Side,
			L2Signature:   &sigStr,
			L2Nonce:       &nonceStr,
			L2ExpireTime:  &l2ExpireTimeStr,
			L2Value:       &valueStr,
			L2Size:        &params.Size,
			L2LimitFee:    &amountFeeStr,
			ClientOrderId: &clientOrderId,
			ExpireTime:    &expireTimeStr,
			ReduceOnly:    &params.ReduceOnly,
		})

	// Execute request
	resp, _, err := req.Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to create order: %w", err)
	}

	if resp.GetCode() != ResponseCodeSuccess {
		if errorParam := resp.GetErrorParam(); errorParam != nil {
			return nil, fmt.Errorf("request failed with error params: %v", errorParam)
		}
		return nil, fmt.Errorf("request failed with code: %s", resp.GetCode())
	}

	return resp, nil
}

// CancelOrder cancels a specific order
func (c *Client) CancelOrder(ctx context.Context, params *CancelOrderParams) (interface{}, error) {
	if params.OrderId != "" {
		req := c.openapiClient.Class04OrderPrivateApiAPI.CancelOrderById(ctx)
		accountID := strconv.FormatInt(c.GetAccountID(), 10)
		cancelParam := openapi.CancelOrderByIdParam{
			AccountId:   &accountID,
			OrderIdList: []string{params.OrderId},
		}
		req = req.CancelOrderByIdParam(cancelParam)
		resp, _, err := req.Execute()
		if err != nil {
			return nil, err
		}
		if resp.GetCode() != ResponseCodeSuccess {
			if errorParam := resp.GetErrorParam(); errorParam != nil {
				return nil, fmt.Errorf("request failed with error params: %v", errorParam)
			}
			return nil, fmt.Errorf("request failed with code: %s", resp.GetCode())
		}
		return resp, nil
	} else if params.ClientId != "" {
		req := c.openapiClient.Class04OrderPrivateApiAPI.CancelOrderByClientOrderId(ctx)
		accountID := strconv.FormatInt(c.GetAccountID(), 10)
		cancelParam := openapi.CancelOrderByClientOrderIdParam{
			AccountId:         &accountID,
			ClientOrderIdList: []string{params.ClientId},
		}
		req = req.CancelOrderByClientOrderIdParam(cancelParam)
		resp, _, err := req.Execute()
		if err != nil {
			return nil, err
		}
		if resp.GetCode() != ResponseCodeSuccess {
			if errorParam := resp.GetErrorParam(); errorParam != nil {
				return nil, fmt.Errorf("request failed with error params: %v", errorParam)
			}
			return nil, fmt.Errorf("request failed with code: %s", resp.GetCode())
		}
		return resp, nil
	} else if params.ContractId != "" {
		req := c.openapiClient.Class04OrderPrivateApiAPI.CancelAllOrder(ctx)
		accountID := strconv.FormatInt(c.GetAccountID(), 10)
		cancelParam := openapi.CancelAllOrderParam{
			AccountId:            &accountID,
			FilterContractIdList: []string{params.ContractId},
		}
		req = req.CancelAllOrderParam(cancelParam)
		resp, _, err := req.Execute()
		if err != nil {
			return nil, err
		}
		if resp.GetCode() != ResponseCodeSuccess {
			if errorParam := resp.GetErrorParam(); errorParam != nil {
				return nil, fmt.Errorf("request failed with error params: %v", errorParam)
			}
			return nil, fmt.Errorf("request failed with code: %s", resp.GetCode())
		}
		return resp, nil
	}
	return nil, fmt.Errorf("must provide either OrderId, ClientId, or ContractId")
}

// GetActiveOrders gets active orders with pagination and filters
func (c *Client) GetActiveOrders(ctx context.Context, params *GetActiveOrderParams) (*openapi.ResultPageDataOrder, error) {
	req := c.openapiClient.Class04OrderPrivateApiAPI.GetActiveOrderPage(ctx)

	// Set account ID and pagination
	req = req.AccountId(strconv.FormatInt(c.GetAccountID(), 10))
	if params.Size != "" {
		req = req.Size(params.Size)
	}
	if params.OffsetData != "" {
		req = req.OffsetData(params.OffsetData)
	}

	// Set filters
	if len(params.FilterCoinIdList) > 0 {
		req = req.FilterCoinIdList(strings.Join(params.FilterCoinIdList, ","))
	}
	if len(params.FilterContractIdList) > 0 {
		req = req.FilterContractIdList(strings.Join(params.FilterContractIdList, ","))
	}
	if len(params.FilterTypeList) > 0 {
		req = req.FilterTypeList(strings.Join(params.FilterTypeList, ","))
	}
	if len(params.FilterStatusList) > 0 {
		req = req.FilterStatusList(strings.Join(params.FilterStatusList, ","))
	}

	// Set boolean filters
	if params.FilterIsLiquidate != nil {
		req = req.FilterIsLiquidateList(strconv.FormatBool(*params.FilterIsLiquidate))
	}
	if params.FilterIsDeleverage != nil {
		req = req.FilterIsDeleverageList(strconv.FormatBool(*params.FilterIsDeleverage))
	}
	if params.FilterIsPositionTpsl != nil {
		req = req.FilterIsPositionTpslList(strconv.FormatBool(*params.FilterIsPositionTpsl))
	}

	// Set time filters
	if params.FilterStartCreatedTimeInclusive > 0 {
		req = req.FilterStartCreatedTimeInclusive(strconv.FormatUint(params.FilterStartCreatedTimeInclusive, 10))
	}
	if params.FilterEndCreatedTimeExclusive > 0 {
		req = req.FilterEndCreatedTimeExclusive(strconv.FormatUint(params.FilterEndCreatedTimeExclusive, 10))
	}

	resp, _, err := req.Execute()
	if err != nil {
		return nil, err
	}
	if resp.GetCode() != ResponseCodeSuccess {
		if errorParam := resp.GetErrorParam(); errorParam != nil {
			return nil, fmt.Errorf("request failed with error params: %v", errorParam)
		}
		return nil, fmt.Errorf("request failed with code: %s", resp.GetCode())
	}
	return resp, nil
}

// GetOrderFillTransactions gets order fill transactions with pagination and filters
func (c *Client) GetOrderFillTransactions(ctx context.Context, params *OrderFillTransactionParams) (*openapi.ResultPageDataOrderFillTransaction, error) {
	req := c.openapiClient.Class04OrderPrivateApiAPI.GetHistoryOrderFillTransactionPage(ctx)

	// Set account ID and pagination
	req = req.AccountId(strconv.FormatInt(c.GetAccountID(), 10))
	if params.Size != "" {
		req = req.Size(params.Size)
	}
	if params.OffsetData != "" {
		req = req.OffsetData(params.OffsetData)
	}

	// Set filters
	if len(params.FilterCoinIdList) > 0 {
		req = req.FilterCoinIdList(strings.Join(params.FilterCoinIdList, ","))
	}
	if len(params.FilterContractIdList) > 0 {
		req = req.FilterContractIdList(strings.Join(params.FilterContractIdList, ","))
	}
	if len(params.FilterOrderIdList) > 0 {
		req = req.FilterOrderIdList(strings.Join(params.FilterOrderIdList, ","))
	}

	// Set boolean filters
	if params.FilterIsLiquidate != nil {
		req = req.FilterIsLiquidateList(strconv.FormatBool(*params.FilterIsLiquidate))
	}
	if params.FilterIsDeleverage != nil {
		req = req.FilterIsDeleverageList(strconv.FormatBool(*params.FilterIsDeleverage))
	}
	if params.FilterIsPositionTpsl != nil {
		req = req.FilterIsPositionTpslList(strconv.FormatBool(*params.FilterIsPositionTpsl))
	}

	// Set time filters
	if params.FilterStartCreatedTimeInclusive > 0 {
		req = req.FilterStartCreatedTimeInclusive(strconv.FormatUint(params.FilterStartCreatedTimeInclusive, 10))
	}
	if params.FilterEndCreatedTimeExclusive > 0 {
		req = req.FilterEndCreatedTimeExclusive(strconv.FormatUint(params.FilterEndCreatedTimeExclusive, 10))
	}

	resp, _, err := req.Execute()
	if err != nil {
		return nil, err
	}
	if resp.GetCode() != ResponseCodeSuccess {
		if errorParam := resp.GetErrorParam(); errorParam != nil {
			return nil, fmt.Errorf("request failed with error params: %v", errorParam)
		}
		return nil, fmt.Errorf("request failed with code: %s", resp.GetCode())
	}
	return resp, nil
}

// GetMaxOrderSize gets the maximum order size for a given contract and price
func (c *Client) GetMaxOrderSize(ctx context.Context, contractID string, price float64) (*openapi.ResultGetMaxCreateOrderSize, error) {
	req := c.openapiClient.Class04OrderPrivateApiAPI.GetMaxCreateOrderSize(ctx)
	accountID := strconv.FormatInt(c.GetAccountID(), 10)
	priceStr := fmt.Sprintf("%f", price)
	param := openapi.GetMaxCreateOrderSizeParam{
		AccountId:  &accountID,
		ContractId: &contractID,
		Price:      &priceStr,
	}
	resp, _, err := req.GetMaxCreateOrderSizeParam(param).Execute()
	if err != nil {
		return nil, err
	}
	if resp.GetCode() != ResponseCodeSuccess {
		if errorParam := resp.GetErrorParam(); errorParam != nil {
			return nil, fmt.Errorf("request failed with error params: %v", errorParam)
		}
		return nil, fmt.Errorf("request failed with code: %s", resp.GetCode())
	}
	return resp, nil
}
