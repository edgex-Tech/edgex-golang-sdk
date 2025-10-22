package account

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/edgex-Tech/edgex-golang-sdk/sdk/internal"
)

// Client represents the account client
type Client struct {
	*internal.Client
}

// NewClient creates a new account client
func NewClient(client *internal.Client) *Client {
	return &Client{
		Client: client,
	}
}

// GetAccountAsset gets the account asset information
func (c *Client) GetAccountAsset(ctx context.Context) (*GetAccountAssetResponse, error) {
	url := fmt.Sprintf("%s/api/v1/private/account/getAccountAsset", c.Client.GetBaseURL())
	params := map[string]string{
		"accountId": strconv.FormatInt(c.Client.GetAccountID(), 10),
	}

	resp, err := c.Client.HttpRequest(url, "GET", nil, params)
	if err != nil {
		return nil, fmt.Errorf("failed to get account asset: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result GetAccountAssetResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s, errorParam: %v", result.Code, result.ErrorParam)
	}

	return &result, nil
}

// GetAccountPositions gets the account positions
func (c *Client) GetAccountPositions(ctx context.Context) (*ListPositionResponse, error) {
	url := fmt.Sprintf("%s/api/v1/private/account/getAccountAsset", c.Client.GetBaseURL())
	params := map[string]string{
		"accountId": strconv.FormatInt(c.Client.GetAccountID(), 10),
	}

	resp, err := c.Client.HttpRequest(url, "GET", nil, params)
	if err != nil {
		return nil, fmt.Errorf("failed to get account positions: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var assetResp GetAccountAssetResponse
	if err := json.Unmarshal(body, &assetResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if assetResp.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", assetResp.Code)
	}

	result := &ListPositionResponse{
		Code: assetResp.Code,
		Data: assetResp.Data.PositionList,
	}

	return result, nil
}


// GetPositionTransactionPage gets the position transactions with pagination
func (c *Client) GetPositionTransactionPage(ctx context.Context, params GetPositionTransactionPageParams) (*PageDataPositionTransactionResponse, error) {
	url := fmt.Sprintf("%s/api/v1/private/account/getPositionTransactionPage", c.Client.GetBaseURL())
	queryParams := map[string]string{
		"accountId": strconv.FormatInt(c.Client.GetAccountID(), 10),
		"size":      strconv.FormatInt(int64(params.Size), 10),
	}

	if params.OffsetData != "" {
		queryParams["offsetData"] = params.OffsetData
	}
	if len(params.FilterCoinIDList) > 0 {
		queryParams["filterCoinIdList"] = internal.JoinStrings(params.FilterCoinIDList)
	}
	if len(params.FilterContractIDList) > 0 {
		queryParams["filterContractIdList"] = internal.JoinStrings(params.FilterContractIDList)
	}
	if len(params.FilterTypeList) > 0 {
		queryParams["filterTypeList"] = internal.JoinStrings(params.FilterTypeList)
	}
	if params.FilterStartCreatedTime > 0 {
		queryParams["filterStartCreatedTimeInclusive"] = strconv.FormatInt(params.FilterStartCreatedTime, 10)
	}
	if params.FilterEndCreatedTime > 0 {
		queryParams["filterEndCreatedTimeExclusive"] = strconv.FormatInt(params.FilterEndCreatedTime, 10)
	}
	if params.FilterCloseOnly != nil {
		queryParams["filterCloseOnly"] = fmt.Sprintf("%v", *params.FilterCloseOnly)
	}
	if params.FilterOpenOnly != nil {
		queryParams["filterOpenOnly"] = fmt.Sprintf("%v", *params.FilterOpenOnly)
	}

	resp, err := c.Client.HttpRequest(url, "GET", nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get position transaction page: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result PageDataPositionTransactionResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}

// GetCollateralTransactionPage gets the collateral transactions with pagination
func (c *Client) GetCollateralTransactionPage(ctx context.Context, params GetCollateralTransactionPageParams) (*PageDataCollateralTransactionResponse, error) {
	url := fmt.Sprintf("%s/api/v1/private/account/getCollateralTransactionPage", c.Client.GetBaseURL())
	queryParams := map[string]string{
		"accountId": strconv.FormatInt(c.Client.GetAccountID(), 10),
		"size":      strconv.FormatInt(int64(params.Size), 10),
	}

	if params.OffsetData != "" {
		queryParams["offsetData"] = params.OffsetData
	}
	if len(params.FilterCoinIDList) > 0 {
		queryParams["filterCoinIdList"] = internal.JoinStrings(params.FilterCoinIDList)
	}
	if len(params.FilterTypeList) > 0 {
		queryParams["filterTypeList"] = internal.JoinStrings(params.FilterTypeList)
	}
	if params.FilterStartCreatedTime > 0 {
		queryParams["filterStartCreatedTimeInclusive"] = strconv.FormatInt(params.FilterStartCreatedTime, 10)
	}
	if params.FilterEndCreatedTime > 0 {
		queryParams["filterEndCreatedTimeExclusive"] = strconv.FormatInt(params.FilterEndCreatedTime, 10)
	}

	resp, err := c.Client.HttpRequest(url, "GET", nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get collateral transaction page: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result PageDataCollateralTransactionResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}

// GetPositionByContractID gets position information for specific contracts
func (c *Client) GetPositionByContractID(ctx context.Context, contractIDs []string) (*ListPositionResponse, error) {
	if len(contractIDs) == 0 {
		return nil, fmt.Errorf("at least one contractId is required")
	}

	url := fmt.Sprintf("%s/api/v1/private/account/getPositionByContractId", c.Client.GetBaseURL())
	queryParams := map[string]string{
		"accountId":      strconv.FormatInt(c.Client.GetAccountID(), 10),
		"contractIdList": internal.JoinStrings(contractIDs),
	}

	resp, err := c.Client.HttpRequest(url, "GET", nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get position by contract ID: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result ListPositionResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}

// GetPositionTermPage gets position terms with pagination
func (c *Client) GetPositionTermPage(ctx context.Context, params GetPositionTermPageParams) (*PageDataPositionTermResponse, error) {
	url := fmt.Sprintf("%s/api/v1/private/account/getPositionTermPage", c.Client.GetBaseURL())
	queryParams := map[string]string{
		"accountId": strconv.FormatInt(c.Client.GetAccountID(), 10),
		"size":      strconv.FormatInt(int64(params.Size), 10),
	}

	if params.OffsetData != "" {
		queryParams["offsetData"] = params.OffsetData
	}
	if len(params.FilterCoinIDList) > 0 {
		queryParams["filterCoinIdList"] = internal.JoinStrings(params.FilterCoinIDList)
	}
	if len(params.FilterContractIDList) > 0 {
		queryParams["filterContractIdList"] = internal.JoinStrings(params.FilterContractIDList)
	}
	if params.FilterIsLongPosition != nil {
		queryParams["filterIsLongPosition"] = fmt.Sprintf("%v", *params.FilterIsLongPosition)
	}
	if params.FilterStartCreatedTime > 0 {
		queryParams["filterStartCreatedTimeInclusive"] = strconv.FormatInt(params.FilterStartCreatedTime, 10)
	}
	if params.FilterEndCreatedTime > 0 {
		queryParams["filterEndCreatedTimeExclusive"] = strconv.FormatInt(params.FilterEndCreatedTime, 10)
	}

	resp, err := c.Client.HttpRequest(url, "GET", nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get position term page: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result PageDataPositionTermResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}

// GetCollateralByCoinID gets collateral information for specific coins
func (c *Client) GetCollateralByCoinID(ctx context.Context, coinIDs []string) (*ListCollateralResponse, error) {
	url := fmt.Sprintf("%s/api/v1/private/account/getCollateralByCoinId", c.Client.GetBaseURL())
	queryParams := map[string]string{
		"accountId": strconv.FormatInt(c.Client.GetAccountID(), 10),
	}

	if len(coinIDs) > 0 {
		queryParams["coinIdList"] = internal.JoinStrings(coinIDs)
	}

	resp, err := c.Client.HttpRequest(url, "GET", nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get collateral by coin ID: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result ListCollateralResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}

// GetAccountByID gets account information by ID
func (c *Client) GetAccountByID(ctx context.Context) (*AccountResponse, error) {
	url := fmt.Sprintf("%s/api/v1/private/account/getAccountById", c.Client.GetBaseURL())
	queryParams := map[string]string{
		"accountId": strconv.FormatInt(c.Client.GetAccountID(), 10),
	}

	resp, err := c.Client.HttpRequest(url, "GET", nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get account by ID: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result AccountResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}

// GetAccountAssetSnapshotPage gets account asset snapshots with pagination
func (c *Client) GetAccountAssetSnapshotPage(ctx context.Context, params GetAccountAssetSnapshotPageParams) (*PageDataAccountAssetSnapshotResponse, error) {
	if params.CoinID == "" {
		return nil, fmt.Errorf("coinId is required")
	}

	url := fmt.Sprintf("%s/api/v1/private/account/getAccountAssetSnapshotPage", c.Client.GetBaseURL())
	queryParams := map[string]string{
		"accountId": strconv.FormatInt(c.Client.GetAccountID(), 10),
		"size":      strconv.FormatInt(int64(params.Size), 10),
		"coinId":    params.CoinID,
	}

	if params.OffsetData != "" {
		queryParams["offsetData"] = params.OffsetData
	}
	if params.FilterTimeTag != nil {
		queryParams["filterTimeTag"] = strconv.FormatInt(int64(*params.FilterTimeTag), 10)
	}
	if params.FilterStartTime > 0 {
		queryParams["filterStartTimeInclusive"] = strconv.FormatInt(params.FilterStartTime, 10)
	}
	if params.FilterEndTime > 0 {
		queryParams["filterEndTimeExclusive"] = strconv.FormatInt(params.FilterEndTime, 10)
	}

	resp, err := c.Client.HttpRequest(url, "GET", nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get account asset snapshot page: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result PageDataAccountAssetSnapshotResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}

// GetPositionTransactionByID gets specific position transactions by IDs
func (c *Client) GetPositionTransactionByID(ctx context.Context, transactionIDs []string) (*ListPositionTransactionResponse, error) {
	if len(transactionIDs) == 0 {
		return nil, fmt.Errorf("at least one transactionId is required")
	}

	url := fmt.Sprintf("%s/api/v1/private/account/getPositionTransactionById", c.Client.GetBaseURL())
	queryParams := map[string]string{
		"accountId":                 strconv.FormatInt(c.Client.GetAccountID(), 10),
		"positionTransactionIdList": internal.JoinStrings(transactionIDs),
	}

	resp, err := c.Client.HttpRequest(url, "GET", nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get position transaction by ID: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result ListPositionTransactionResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}

// GetCollateralTransactionByID gets specific collateral transactions by IDs
func (c *Client) GetCollateralTransactionByID(ctx context.Context, transactionIDs []string) (*ListCollateralTransactionResponse, error) {
	if len(transactionIDs) == 0 {
		return nil, fmt.Errorf("at least one transactionId is required")
	}

	url := fmt.Sprintf("%s/api/v1/private/account/getCollateralTransactionById", c.Client.GetBaseURL())
	queryParams := map[string]string{
		"accountId":                   strconv.FormatInt(c.Client.GetAccountID(), 10),
		"collateralTransactionIdList": internal.JoinStrings(transactionIDs),
	}

	resp, err := c.Client.HttpRequest(url, "GET", nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get collateral transaction by ID: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result ListCollateralTransactionResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}

// GetAccountDeleverageLight gets account deleverage light information
func (c *Client) GetAccountDeleverageLight(ctx context.Context) (*GetAccountDeleverageLightResponse, error) {
	url := fmt.Sprintf("%s/api/v1/private/account/getAccountDeleverageLight", c.Client.GetBaseURL())
	queryParams := map[string]string{
		"accountId": strconv.FormatInt(c.Client.GetAccountID(), 10),
	}

	resp, err := c.Client.HttpRequest(url, "GET", nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get account deleverage light: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result GetAccountDeleverageLightResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}

// UpdateLeverageSetting updates the account leverage settings
func (c *Client) UpdateLeverageSetting(ctx context.Context, contractID string, leverage string) error {
	url := fmt.Sprintf("%s/api/v1/private/account/updateLeverageSetting", c.Client.GetBaseURL())
	data := map[string]interface{}{
		"accountId":  c.Client.GetAccountID(),
		"contractId": contractID,
		"leverage":   leverage,
	}

	resp, err := c.Client.HttpRequest(url, "POST", data, nil)
	if err != nil {
		return fmt.Errorf("failed to update leverage setting: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	var result UpdateLeverageSettingResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return fmt.Errorf("request failed with code: %s", result.Code)
	}

	return nil
}
