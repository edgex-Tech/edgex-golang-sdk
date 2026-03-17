package funding

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// Client represents the new funding client without OpenAPI dependencies
type Client struct {
	c clientInterface
}

type clientInterface interface {
	GetBaseURL() string
	HttpRequest(urlStr string, method string, data map[string]interface{}, params map[string]string) (*http.Response, error)
}

// NewClient creates a new funding client
func NewClient(client clientInterface) *Client {
	return &Client{
		c: client,
	}
}

// GetFundingRate gets the funding rate for a contract
func (c *Client) GetFundingRate(ctx context.Context, params GetFundingRateParams) (*ResultPageDataFundingRate, error) {
	url := fmt.Sprintf("%s/api/v2/public/funding/getFundingRatePage", c.c.GetBaseURL())
	queryParams := map[string]string{
		"contractId":                   params.ContractID,
		"filterSettlementFundingRate": "true",
	}

	if params.Size != nil {
		queryParams["size"] = strconv.FormatInt(int64(*params.Size), 10)
	}
	if params.Offset != nil {
		queryParams["offsetData"] = *params.Offset
	}
	if params.From != nil {
		queryParams["filterBeginTimeInclusive"] = strconv.FormatInt(*params.From, 10)
	}
	if params.To != nil {
		queryParams["filterEndTimeExclusive"] = strconv.FormatInt(*params.To, 10)
	}

	resp, err := c.c.HttpRequest(url, "GET", nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get funding rate: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result ResultPageDataFundingRate
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}

// GetLatestFundingRate gets the latest funding rate for a contract
func (c *Client) GetLatestFundingRate(ctx context.Context, params GetLatestFundingRateParams) (*ResultListFundingRate, error) {
	url := fmt.Sprintf("%s/api/v2/public/funding/getLatestFundingRate", c.c.GetBaseURL())
	queryParams := map[string]string{
		"contractId": params.ContractID,
	}

	resp, err := c.c.HttpRequest(url, "GET", nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest funding rate: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result ResultListFundingRate
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}

