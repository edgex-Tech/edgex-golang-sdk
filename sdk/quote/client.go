package quote

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// Client represents the new quote client without OpenAPI dependencies
type Client struct {
	c clientInterface
}

type clientInterface interface {
	GetBaseURL() string
	HttpRequest(urlStr string, method string, data map[string]interface{}, params map[string]string) (*http.Response, error)
}

// NewClient creates a new quote client
func NewClient(client clientInterface) *Client {
	return &Client{
		c: client,
	}
}

func decodeResponse[T any](resp *http.Response, operation string) (*T, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result T
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &result, nil
}

// Get24HourQuote gets the 24-hour quotes for given contracts
func (c *Client) Get24HourQuote(ctx context.Context, contractId string) (*ResultListTicker, error) {
	url := fmt.Sprintf("%s/api/v2/public/quote/getTicker", c.c.GetBaseURL())
	queryParams := map[string]string{
		"contractId": contractId,
	}

	resp, err := c.c.HttpRequest(url, "GET", nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get 24-hour quotes: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeResponse[ResultListTicker](resp, "get 24-hour quotes")
	if err != nil {
		return nil, err
	}
	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return result, nil
}

// GetKLine gets the K-line data for a contract
func (c *Client) GetKLine(ctx context.Context, params GetKLineParams) (*ResultPageDataKline, error) {
	url := fmt.Sprintf("%s/api/v2/public/quote/getKline", c.c.GetBaseURL())
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

	resp, err := c.c.HttpRequest(url, "GET", nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get k-line data: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeResponse[ResultPageDataKline](resp, "get k-line data")
	if err != nil {
		return nil, err
	}
	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return result, nil
}

// GetOrderBookDepth gets the order book depth for a contract
func (c *Client) GetOrderBookDepth(ctx context.Context, params GetOrderBookDepthParams) (*ResultListDepth, error) {
	url := fmt.Sprintf("%s/api/v2/public/quote/getDepth", c.c.GetBaseURL())
	queryParams := map[string]string{
		"contractId": params.ContractID,
		"level":      strconv.FormatInt(int64(params.Size), 10),
	}

	if params.Precision != nil {
		queryParams["precision"] = *params.Precision
	}

	resp, err := c.c.HttpRequest(url, "GET", nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get order book depth: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeResponse[ResultListDepth](resp, "get order book depth")
	if err != nil {
		return nil, err
	}
	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return result, nil
}

// GetMultiContractKLine gets the K-line data for multiple contracts
func (c *Client) GetMultiContractKLine(ctx context.Context, params GetMultiContractKLineParams) (*ResultListContractKline, error) {
	url := fmt.Sprintf("%s/api/v2/public/quote/getMultiContractKline", c.c.GetBaseURL())
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

	resp, err := c.c.HttpRequest(url, "GET", nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get multi-contract k-line data: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeResponse[ResultListContractKline](resp, "get multi-contract k-line data")
	if err != nil {
		return nil, err
	}
	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return result, nil
}

func (c *Client) GetMarketStatus(ctx context.Context, params GetMarketStatusParams) (*ResultGetMarketStatusModel, error) {
	url := fmt.Sprintf("%s/api/v2/public/quote/getMarketStatus", c.c.GetBaseURL())
	queryParams := map[string]string{}
	if params.ContractID != nil {
		queryParams["contractId"] = strconv.FormatInt(*params.ContractID, 10)
	}

	resp, err := c.c.HttpRequest(url, "GET", nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get market status: %w", err)
	}
	defer resp.Body.Close()

	result, err := decodeResponse[ResultGetMarketStatusModel](resp, "get market status")
	if err != nil {
		return nil, err
	}
	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}
	return result, nil
}
