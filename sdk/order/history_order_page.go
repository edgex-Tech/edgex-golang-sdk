package order

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
)

// GetHistoryOrderPageParams represents parameters for getting history orders with POST method (V2 only)
type GetHistoryOrderPageParams struct {
	Size                            int32    `json:"size"`
	OffsetData                      string   `json:"offsetData,omitempty"`
	FilterCoinIdList                []string `json:"filterCoinIdList,omitempty"`
	FilterContractIdList            []string `json:"filterContractIdList,omitempty"`
	FilterTypeList                  []string `json:"filterTypeList,omitempty"`
	FilterStatusList                []string `json:"filterStatusList,omitempty"`
	FilterStartCreatedTimeInclusive *int64   `json:"filterStartCreatedTimeInclusive,omitempty"`
	FilterEndCreatedTimeExclusive   *int64   `json:"filterEndCreatedTimeExclusive,omitempty"`
	FilterOrderIdList               []string `json:"filterOrderIdList,omitempty"`
	FilterIsLiquidate               *bool    `json:"-"`
	FilterIsDeleverage              *bool    `json:"-"`
	FilterIsPositionTpsl            *bool    `json:"-"`
}

// GetHistoryOrderPageResponse represents the response for history order page query
type GetHistoryOrderPageResponse struct {
	Code      string                `json:"code"`
	Msg       string                `json:"msg"`
	Data      *HistoryOrderPageData `json:"data"`
	RequestID string                `json:"requestId"`
}

// HistoryOrderPageData contains paginated history order data
type HistoryOrderPageData struct {
	DataList   []HistoryOrder `json:"dataList"`
	HasNext    bool           `json:"hasNext"`
	OffsetData string         `json:"offsetData"`
}

// HistoryOrder represents a historical order record
type HistoryOrder struct {
	OrderID            string `json:"orderId"`
	ClientOrderID      string `json:"clientOrderId"`
	AccountID          string `json:"accountId"`
	ContractID         string `json:"contractId"`
	CoinID             string `json:"coinId"`
	Type               string `json:"type"`
	Side               string `json:"side"`
	Price              string `json:"price"`
	Size               string `json:"size"`
	FilledSize         string `json:"filledSize"`
	UnfilledSize       string `json:"unfilledSize"`
	Status             string `json:"status"`
	TimeInForce        string `json:"timeInForce"`
	ReduceOnly         bool   `json:"reduceOnly"`
	IsLiquidate        bool   `json:"isLiquidate"`
	IsDeleverage       bool   `json:"isDeleverage"`
	IsPositionTPSL     bool   `json:"isPositionTpsl"`
	CreatedTime        string `json:"createdTime"`
	UpdatedTime        string `json:"updatedTime"`
	ExpireTime         string `json:"expireTime"`
	AverageFilledPrice string `json:"averageFilledPrice"`
	TriggerPrice       string `json:"triggerPrice"`
	TriggerPriceType   string `json:"triggerPriceType"`
	Remark             string `json:"remark"`
	Fee                string `json:"fee"`
	FeeCoinID          string `json:"feeCoinId"`
	RealizedPnl        string `json:"realizedPnl"`
}

// GetHistoryOrderPage gets historical orders with complex filters using POST method (V2 new feature)
func (c *Client) GetHistoryOrderPage(ctx context.Context, params *GetHistoryOrderPageParams) (*GetHistoryOrderPageResponse, error) {
	url := fmt.Sprintf("%s/api/v2/private/order/getHistoryOrderPage", c.c.GetBaseURL())

	// Build request body
	body := map[string]interface{}{
		"accountId": fmt.Sprintf("%d", c.c.GetAccountID()),
		"size":      params.Size,
	}

	if params.OffsetData != "" {
		body["offsetData"] = params.OffsetData
	}
	if len(params.FilterCoinIdList) > 0 {
		body["filterCoinIdList"] = params.FilterCoinIdList
	}
	if len(params.FilterContractIdList) > 0 {
		body["filterContractIdList"] = params.FilterContractIdList
	}
	if len(params.FilterTypeList) > 0 {
		body["filterTypeList"] = params.FilterTypeList
	}
	if len(params.FilterStatusList) > 0 {
		body["filterStatusList"] = params.FilterStatusList
	}
	if params.FilterIsLiquidate != nil {
		body["filterIsLiquidateList"] = []bool{*params.FilterIsLiquidate}
	}
	if params.FilterIsDeleverage != nil {
		body["filterIsDeleverageList"] = []bool{*params.FilterIsDeleverage}
	}
	if params.FilterIsPositionTpsl != nil {
		body["filterIsPositionTpslList"] = []bool{*params.FilterIsPositionTpsl}
	}
	if params.FilterStartCreatedTimeInclusive != nil {
		body["filterStartCreatedTimeInclusive"] = fmt.Sprintf("%d", *params.FilterStartCreatedTimeInclusive)
	}
	if params.FilterEndCreatedTimeExclusive != nil {
		body["filterEndCreatedTimeExclusive"] = fmt.Sprintf("%d", *params.FilterEndCreatedTimeExclusive)
	}
	if len(params.FilterOrderIdList) > 0 {
		body["filterOrderIdList"] = params.FilterOrderIdList
	}
	resp, err := c.c.HttpRequest(url, "POST", body, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get history order page: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result GetHistoryOrderPageResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s, msg: %s", result.Code, result.Msg)
	}

	return &result, nil
}
