package asset

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/edgex-Tech/edgex-golang-sdk/sdk/internal"
)

// Client represents the new asset client without OpenAPI dependencies
type Client struct {
	*internal.Client
}

// NewClient creates a new asset client
func NewClient(client *internal.Client) *Client {
	return &Client{
		Client: client,
	}
}

// GetAllOrdersPage gets all asset orders with pagination
func (c *Client) GetAllOrdersPage(ctx context.Context, params GetAllOrdersPageParams) (*ResultPageDataAssetOrder, error) {
	url := fmt.Sprintf("%s/api/v1/private/asset/getAllOrdersPage", c.Client.GetBaseURL())
	queryParams := map[string]string{
		"accountId": strconv.FormatInt(c.Client.GetAccountID(), 10),
	}

	if params.StartTime != "" {
		queryParams["startTime"] = params.StartTime
	}
	if params.EndTime != "" {
		queryParams["endTime"] = params.EndTime
	}
	if params.ChainId != "" {
		queryParams["chainId"] = params.ChainId
	}
	if params.TypeList != "" {
		queryParams["typeList"] = params.TypeList
	}
	if params.Size != "" {
		queryParams["size"] = params.Size
	}
	if params.OffsetData != "" {
		queryParams["offsetData"] = params.OffsetData
	}

	resp, err := c.Client.HttpRequest(url, "GET", nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get asset orders: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result ResultPageDataAssetOrder
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}

// GetCoinRate gets the coin rate
func (c *Client) GetCoinRate(ctx context.Context, params GetCoinRateParams) (*ResultGetCoinRate, error) {
	url := fmt.Sprintf("%s/api/v1/private/asset/getCoinRate", c.Client.GetBaseURL())
	queryParams := map[string]string{}

	if params.ChainId != "" {
		queryParams["chainId"] = params.ChainId
	}
	if params.Coin != "" {
		queryParams["coin"] = params.Coin
	}

	resp, err := c.Client.HttpRequest(url, "GET", nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get coin rate: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result ResultGetCoinRate
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}

// GetCrossWithdrawById gets cross withdraw records by ID
func (c *Client) GetCrossWithdrawById(ctx context.Context, params GetCrossWithdrawByIdParams) (*ResultListCrossWithdraw, error) {
	url := fmt.Sprintf("%s/api/v1/private/asset/getCrossWithdrawById", c.Client.GetBaseURL())
	queryParams := map[string]string{
		"accountId": strconv.FormatInt(c.Client.GetAccountID(), 10),
	}

	if params.CrossWithdrawIdList != "" {
		queryParams["crossWithdrawIdList"] = params.CrossWithdrawIdList
	}

	resp, err := c.Client.HttpRequest(url, "GET", nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get cross withdraw by id: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result ResultListCrossWithdraw
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}

// GetCrossWithdrawSignInfo gets cross withdraw sign info
func (c *Client) GetCrossWithdrawSignInfo(ctx context.Context, params GetCrossWithdrawSignInfoParams) (*ResultGetCrossWithdrawSignInfo, error) {
	url := fmt.Sprintf("%s/api/v1/private/asset/getCrossWithdrawSignInfo", c.Client.GetBaseURL())
	queryParams := map[string]string{}

	if params.ChainId != "" {
		queryParams["chainId"] = params.ChainId
	}
	if params.Amount != "" {
		queryParams["amount"] = params.Amount
	}

	resp, err := c.Client.HttpRequest(url, "GET", nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get cross withdraw sign info: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result ResultGetCrossWithdrawSignInfo
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}

// GetFastWithdrawById gets fast withdraw records by ID
func (c *Client) GetFastWithdrawById(ctx context.Context, params GetFastWithdrawByIdParams) (*ResultListFastWithdraw, error) {
	url := fmt.Sprintf("%s/api/v1/private/asset/getFastWithdrawById", c.Client.GetBaseURL())
	queryParams := map[string]string{
		"accountId": strconv.FormatInt(c.Client.GetAccountID(), 10),
	}

	if params.FastWithdrawIdList != "" {
		queryParams["fastWithdrawIdList"] = params.FastWithdrawIdList
	}

	resp, err := c.Client.HttpRequest(url, "GET", nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get fast withdraw by id: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result ResultListFastWithdraw
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}

// GetFastWithdrawSignInfo gets fast withdraw sign info
func (c *Client) GetFastWithdrawSignInfo(ctx context.Context, params GetFastWithdrawSignInfoParams) (*ResultGetFastWithdrawSignInfo, error) {
	url := fmt.Sprintf("%s/api/v1/private/asset/getFastWithdrawSignInfo", c.Client.GetBaseURL())
	queryParams := map[string]string{}

	if params.ChainId != "" {
		queryParams["chainId"] = params.ChainId
	}
	if params.Amount != "" {
		queryParams["amount"] = params.Amount
	}

	resp, err := c.Client.HttpRequest(url, "GET", nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get fast withdraw sign info: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result ResultGetFastWithdrawSignInfo
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}

// GetNormalWithdrawById gets normal withdraw records by ID
func (c *Client) GetNormalWithdrawById(ctx context.Context, params GetNormalWithdrawByIdParams) (*ResultListNormalWithdraw, error) {
	url := fmt.Sprintf("%s/api/v1/private/asset/getNormalWithdrawById", c.Client.GetBaseURL())
	queryParams := map[string]string{
		"accountId": strconv.FormatInt(c.Client.GetAccountID(), 10),
	}

	if params.NormalWithdrawIdList != "" {
		queryParams["normalWithdrawIdList"] = params.NormalWithdrawIdList
	}

	resp, err := c.Client.HttpRequest(url, "GET", nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get normal withdraw by id: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result ResultListNormalWithdraw
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}

// GetNormalWithdrawableAmount gets normal withdrawable amount
func (c *Client) GetNormalWithdrawableAmount(ctx context.Context, params GetNormalWithdrawableAmountParams) (*ResultGetNormalWithdrawableAmount, error) {
	url := fmt.Sprintf("%s/api/v1/private/asset/getNormalWithdrawableAmount", c.Client.GetBaseURL())
	queryParams := map[string]string{}

	if params.Address != "" {
		queryParams["address"] = params.Address
	}

	resp, err := c.Client.HttpRequest(url, "GET", nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get normal withdrawable amount: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result ResultGetNormalWithdrawableAmount
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}

// CreateNormalWithdraw creates a normal withdrawal order
func (c *Client) CreateNormalWithdraw(ctx context.Context, params CreateNormalWithdrawParams) (*ResultCreateNormalWithdraw, error) {
	url := fmt.Sprintf("%s/api/v1/private/asset/createNormalWithdraw", c.Client.GetBaseURL())

	body := map[string]interface{}{
		"accountId":        strconv.FormatInt(c.Client.GetAccountID(), 10),
		"coinId":           params.CoinId,
		"amount":           params.Amount,
		"ethAddress":       params.EthAddress,
		"clientWithdrawId": params.ClientWithdrawId,
		"expireTime":       params.ExpireTime,
		"l2Signature":      params.L2Signature,
	}

	resp, err := c.Client.HttpRequest(url, "POST", body, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create normal withdraw: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result ResultCreateNormalWithdraw
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}

// CreateCrossWithdraw creates a cross-chain withdrawal order
func (c *Client) CreateCrossWithdraw(ctx context.Context, params CreateCrossWithdrawParams) (*ResultCreateCrossWithdraw, error) {
	url := fmt.Sprintf("%s/api/v1/private/asset/createCrossWithdraw", c.Client.GetBaseURL())

	body := map[string]interface{}{
		"accountId":             strconv.FormatInt(c.Client.GetAccountID(), 10),
		"coinId":                params.CoinId,
		"amount":                params.Amount,
		"ethAddress":            params.EthAddress,
		"erc20Address":          params.Erc20Address,
		"lpAccountId":           params.LpAccountId,
		"clientCrossWithdrawId": params.ClientCrossWithdrawId,
		"expireTime":            params.ExpireTime,
		"l2Signature":           params.L2Signature,
		"fee":                   params.Fee,
		"chainId":               params.ChainId,
		"mpcAddress":            params.MpcAddress,
		"mpcSignature":          params.MpcSignature,
		"mpcSignTime":           params.MpcSignTime,
	}

	resp, err := c.Client.HttpRequest(url, "POST", body, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create cross withdraw: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result ResultCreateCrossWithdraw
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}
