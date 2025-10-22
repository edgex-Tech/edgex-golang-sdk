package quote

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/edgex-Tech/edgex-golang-sdk/sdk/internal"
)

// Client represents the new quote client without OpenAPI dependencies
type Client struct {
	*internal.Client
}

// NewClient creates a new quote client
func NewClient(client *internal.Client) *Client {
	return &Client{
		Client: client,
	}
}

// GetQuoteSummary gets the quote summary for a given contract
func (c *Client) GetQuoteSummary(ctx context.Context, contractID string) (*ResultGetTickerSummaryModel, error) {
	url := fmt.Sprintf("%s/api/v1/public/quote/getTicketSummary", c.Client.GetBaseURL())
	queryParams := map[string]string{}

	resp, err := c.Client.HttpRequest(url, "GET", nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get quote summary: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result ResultGetTickerSummaryModel
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}

// Get24HourQuote gets the 24-hour quotes for given contracts
func (c *Client) Get24HourQuote(ctx context.Context, contractId string) (*ResultListTicker, error) {
	url := fmt.Sprintf("%s/api/v1/public/quote/getTicker", c.Client.GetBaseURL())
	queryParams := map[string]string{
		"contractId": contractId,
	}

	resp, err := c.Client.HttpRequest(url, "GET", nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get 24-hour quotes: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result ResultListTicker
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}

// GetKLine gets the K-line data for a contract
func (c *Client) GetKLine(ctx context.Context, params GetKLineParams) (*ResultPageDataKline, error) {
	url := fmt.Sprintf("%s/api/v1/public/quote/getKline", c.Client.GetBaseURL())
	queryParams := map[string]string{
		"contractId": params.ContractID,
		"klineType":  string(params.Interval),
		"priceType":  string(params.PriceType),
		"size":       strconv.FormatInt(int64(params.Size), 10),
	}

	// Add optional parameters
	if params.OffsetData != "" {
		queryParams["offsetData"] = params.OffsetData
	}
	if params.From != nil {
		queryParams["filterBeginKlineTimeInclusive"] = strconv.FormatInt(*params.From, 10)
	}
	if params.To != nil {
		queryParams["filterEndKlineTimeExclusive"] = strconv.FormatInt(*params.To, 10)
	}
	fmt.Println(queryParams)

	resp, err := c.Client.HttpRequest(url, "GET", nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get k-line data: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result ResultPageDataKline
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}

// GetOrderBookDepth gets the order book depth for a contract
func (c *Client) GetOrderBookDepth(ctx context.Context, params GetOrderBookDepthParams) (*ResultListDepth, error) {
	url := fmt.Sprintf("%s/api/v1/public/quote/getDepth", c.Client.GetBaseURL())
	queryParams := map[string]string{
		"contractId": params.ContractID,
		"level":      strconv.FormatInt(int64(params.Size), 10),
	}

	if params.Precision != nil {
		queryParams["precision"] = *params.Precision
	}

	resp, err := c.Client.HttpRequest(url, "GET", nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get order book depth: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result ResultListDepth
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}

// GetMultiContractKLine gets the K-line data for multiple contracts
func (c *Client) GetMultiContractKLine(ctx context.Context, params GetMultiContractKLineParams) (*ResultListContractKline, error) {
	url := fmt.Sprintf("%s/api/v1/public/quote/getMultiContractKline", c.Client.GetBaseURL())
	queryParams := map[string]string{
		"contractIdList": strings.Join(params.ContractIDs, ","),
		"klineType":      string(params.Interval),
		"size":           strconv.FormatInt(int64(params.Size), 10),
	}

	if params.PriceType != "" {
		queryParams["priceType"] = string(params.PriceType)
	}
	if params.From != nil {
		queryParams["filterBeginKlineTimeInclusive"] = strconv.FormatInt(*params.From, 10)
	}
	if params.To != nil {
		queryParams["filterEndKlineTimeExclusive"] = strconv.FormatInt(*params.To, 10)
	}

	resp, err := c.Client.HttpRequest(url, "GET", nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get multi-contract k-line data: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result ResultListContractKline
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}
