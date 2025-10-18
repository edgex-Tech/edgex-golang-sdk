package funding

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/edgex-Tech/edgex-golang-sdk/sdk/internal"
)

// Client represents the new funding client without OpenAPI dependencies
type Client struct {
	*internal.Client
}

// NewClient creates a new funding client
func NewClient(client *internal.Client) *Client {
	return &Client{
		Client: client,
	}
}

// GetFundingRate gets the funding rate for a contract
func (c *Client) GetFundingRate(ctx context.Context, params GetFundingRateParams) (*ResultPageDataFundingRate, error) {
	url := fmt.Sprintf("%s/api/v1/public/funding/getFundingRatePage", c.Client.GetBaseURL())
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

	resp, err := c.Client.HttpRequest(url, "GET", nil, queryParams)
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
	url := fmt.Sprintf("%s/api/v1/public/funding/getLatestFundingRate", c.Client.GetBaseURL())
	queryParams := map[string]string{
		"contractId": params.ContractID,
	}

	resp, err := c.Client.HttpRequest(url, "GET", nil, queryParams)
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

