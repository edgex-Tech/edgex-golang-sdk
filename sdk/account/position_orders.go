package account

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
)

// GetPositionOrdersParams represents the parameters for GetPositionOrders
type GetPositionOrdersParams struct {
	// Required parameters
	ContractId string // Contract ID
	TermCount  int32  // Position term count
	
	// Pagination parameters
	Page     int32 // Page number (>=1)
	PageSize int32 // Page size (1-100)
}

// GetPositionOrdersResponse represents the response for GetPositionOrders
type GetPositionOrdersResponse struct {
	Code       string                  `json:"code"`
	Msg        string                  `json:"msg"`
	Data       *PositionOrdersPageData `json:"data"`
	ErrorParam interface{}             `json:"errorParam"`
}

// PositionOrdersPageData represents the paginated data
type PositionOrdersPageData struct {
	Page      int32               `json:"page"`
	PageSize  int32               `json:"pageSize"`
	Total     int64               `json:"total"`
	OrderList []PositionOrderData `json:"orderList"`
}

// PositionOrderData represents a single order in position history
type PositionOrderData struct {
	ID                 string `json:"id"`
	UserID             string `json:"userId"`
	AccountID          string `json:"accountId"`
	CoinID             string `json:"coinId"`
	ContractID         string `json:"contractId"`
	Type               string `json:"type"`
	Side               string `json:"side"`
	Price              string `json:"price"`
	Size               string `json:"size"`
	UnfillSize         string `json:"unfillSize"`
	ExecutedSize       string `json:"executedSize"`
	ExecutedValue      string `json:"executedValue"`
	ExecutedFee        string `json:"executedFee"`
	ExecutedAvgPrice   string `json:"executedAvgPrice"`
	Status             string `json:"status"`
	TimeInForce        string `json:"timeInForce"`
	ExpireTime         string `json:"expireTime"`
	ClientOrderID      string `json:"clientOrderId"`
	ReduceOnly         bool   `json:"reduceOnly"`
	PostOnly           bool   `json:"postOnly"`
	Hidden             bool   `json:"hidden"`
	CreatedTime        string `json:"createdTime"`
	UpdatedTime        string `json:"updatedTime"`
}

// GetPositionOrders gets historical orders by position term (V2 only)
func (c *Client) GetPositionOrders(ctx context.Context, params *GetPositionOrdersParams) (*GetPositionOrdersResponse, error) {
	url := fmt.Sprintf("%s/api/v2/private/account/getPositionOrders", c.c.GetBaseURL())
	
	queryParams := map[string]string{
		"accountId":  fmt.Sprintf("%d", c.c.GetAccountID()),
		"contractId": params.ContractId,
		"termCount":  fmt.Sprintf("%d", params.TermCount),
	}
	
	// Add optional pagination parameters
	if params.Page > 0 {
		queryParams["page"] = fmt.Sprintf("%d", params.Page)
	}
	if params.PageSize > 0 {
		queryParams["pageSize"] = fmt.Sprintf("%d", params.PageSize)
	}

	resp, err := c.c.HttpRequest(url, "GET", nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get position orders: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result GetPositionOrdersResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s, msg: %s", result.Code, result.Msg)
	}

	return &result, nil
}
