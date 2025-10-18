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
func (c *Client) CreateOrder(ctx context.Context, params *CreateOrderParams, metadata interface{}) (*ResultCreateOrder, error) {
	// Set default TimeInForce based on order type if not specified
	if params.TimeInForce == "" {
		switch params.Type {
		case OrderTypeMarket:
			params.TimeInForce = string(TimeInForce_IMMEDIATE_OR_CANCEL)
		case OrderTypeLimit:
			params.TimeInForce = string(TimeInForce_GOOD_TIL_CANCEL)
		}
	}

	// Parse decimal values
	size, err := decimal.NewFromString(params.Size)
	if err != nil {
		return nil, fmt.Errorf("failed to parse size: %w", err)
	}

	price, err := decimal.NewFromString(params.Price)
	if err != nil {
		return nil, fmt.Errorf("failed to parse price: %w", err)
	}

	clientOrderId := internal.GenerateUUID()
	if params.ClientOrderId != nil {
		clientOrderId = *params.ClientOrderId
	}

	// Calculate values
	valueDm := price.Mul(size)
	// amountCollateral := valueDm.Shift(6).IntPart()

	// Calculate fee based on order type (using default taker fee rate)
	feeRate, _ := decimal.NewFromString("0.001") // Default fee rate
	amountFeeDm := valueDm.Mul(feeRate).Ceil()
	amountFeeStr := amountFeeDm.String()
	// amountFee := amountFeeDm.Shift(6).IntPart()

	nonce := internal.CalcNonce(clientOrderId)
	l2ExpireTime := time.Now().Add(14 * 24 * time.Hour).UnixMilli()

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
		"l2Nonce":       strconv.FormatInt(nonce, 10),
		"l2ExpireTime":  strconv.FormatInt(l2ExpireTime, 10),
		"l2Value":       valueDm.String(),
		"l2Size":        params.Size,
		"l2LimitFee":    amountFeeStr,
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
	url := fmt.Sprintf("%s/api/v1/private/order/cancelOrder", c.Client.GetBaseURL())
	accountID := strconv.FormatInt(c.Client.GetAccountID(), 10)

	var body map[string]interface{}

	if params.OrderId != "" {
		body = map[string]interface{}{
			"accountId":   accountID,
			"orderIdList": []string{params.OrderId},
		}
	} else if params.ClientId != "" {
		body = map[string]interface{}{
			"accountId":         accountID,
			"clientOrderIdList": []string{params.ClientId},
		}
	} else if params.ContractId != "" {
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
func (c *Client) GetMaxOrderSize(ctx context.Context, contractID string, price float64) (*ResultGetMaxCreateOrderSize, error) {
	url := fmt.Sprintf("%s/api/v1/private/order/getMaxCreateOrderSize", c.Client.GetBaseURL())
	queryParams := map[string]string{
		"accountId":  strconv.FormatInt(c.Client.GetAccountID(), 10),
		"contractId": contractID,
		"price":      fmt.Sprintf("%f", price),
	}

	resp, err := c.Client.HttpRequest(url, "GET", nil, queryParams)
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
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}
