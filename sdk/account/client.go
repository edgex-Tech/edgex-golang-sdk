package account

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/edgex-Tech/edgex-golang-sdk/sdk/internal"
	"github.com/edgex-Tech/edgex-golang-sdk/sdk/metadata"
	"github.com/ethereum/go-ethereum/common"
)

// Client represents the account client
type Client struct {
	c clientInterface
}

type clientInterface interface {
	GetAccountID() int64
	GetBaseURL() string
	ResolveSignerAddress() (string, error)
	SignTypedDataWithSignerKey(typedData internal.TypedData) (string, error)
	HttpRequest(urlStr string, method string, data map[string]interface{}, params map[string]string) (*http.Response, error)
}

func calcSetMarginModeL2ExpireTime(now time.Time) string {
	nowMillis := now.UnixMilli()
	nextHourMillis := ((nowMillis + 3600000 - 1) / 3600000) * 3600000
	return strconv.FormatInt(nextHourMillis+14*24*60*60*1000, 10)
}

// NewClient creates a new account client
func NewClient(client clientInterface) *Client {
	return &Client{
		c: client,
	}
}

// GetAccountAsset gets the account asset information
func (c *Client) GetAccountAsset(ctx context.Context) (*GetAccountAssetResponse, error) {
	url := fmt.Sprintf("%s/api/v2/private/account/getAccountAsset", c.c.GetBaseURL())
	params := map[string]string{
		"accountId": strconv.FormatInt(c.c.GetAccountID(), 10),
	}

	resp, err := c.c.HttpRequest(url, "GET", nil, params)
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
	url := fmt.Sprintf("%s/api/v2/private/account/getAccountAsset", c.c.GetBaseURL())
	params := map[string]string{
		"accountId": strconv.FormatInt(c.c.GetAccountID(), 10),
	}

	resp, err := c.c.HttpRequest(url, "GET", nil, params)
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
	url := fmt.Sprintf("%s/api/v2/private/account/getPositionTransactionPage", c.c.GetBaseURL())
	queryParams := map[string]string{
		"accountId": strconv.FormatInt(c.c.GetAccountID(), 10),
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

	resp, err := c.c.HttpRequest(url, "GET", nil, queryParams)
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
	url := fmt.Sprintf("%s/api/v2/private/account/getCollateralTransactionPage", c.c.GetBaseURL())
	queryParams := map[string]string{
		"accountId": strconv.FormatInt(c.c.GetAccountID(), 10),
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

	resp, err := c.c.HttpRequest(url, "GET", nil, queryParams)
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

	url := fmt.Sprintf("%s/api/v2/private/account/getPositionByContractId", c.c.GetBaseURL())
	queryParams := map[string]string{
		"accountId":      strconv.FormatInt(c.c.GetAccountID(), 10),
		"contractIdList": internal.JoinStrings(contractIDs),
	}

	resp, err := c.c.HttpRequest(url, "GET", nil, queryParams)
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
	url := fmt.Sprintf("%s/api/v2/private/account/getPositionTermPage", c.c.GetBaseURL())
	queryParams := map[string]string{
		"accountId": strconv.FormatInt(c.c.GetAccountID(), 10),
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

	resp, err := c.c.HttpRequest(url, "GET", nil, queryParams)
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
	url := fmt.Sprintf("%s/api/v2/private/account/getCollateralByCoinId", c.c.GetBaseURL())
	queryParams := map[string]string{
		"accountId": strconv.FormatInt(c.c.GetAccountID(), 10),
	}

	if len(coinIDs) > 0 {
		queryParams["coinIdList"] = internal.JoinStrings(coinIDs)
	}

	resp, err := c.c.HttpRequest(url, "GET", nil, queryParams)
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
	url := fmt.Sprintf("%s/api/v2/private/account/getAccountById", c.c.GetBaseURL())
	queryParams := map[string]string{
		"accountId": strconv.FormatInt(c.c.GetAccountID(), 10),
	}

	resp, err := c.c.HttpRequest(url, "GET", nil, queryParams)
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

	url := fmt.Sprintf("%s/api/v2/private/account/getAccountAssetSnapshotPage", c.c.GetBaseURL())
	queryParams := map[string]string{
		"accountId": strconv.FormatInt(c.c.GetAccountID(), 10),
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

	resp, err := c.c.HttpRequest(url, "GET", nil, queryParams)
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

	url := fmt.Sprintf("%s/api/v2/private/account/getPositionTransactionById", c.c.GetBaseURL())
	queryParams := map[string]string{
		"accountId":                 strconv.FormatInt(c.c.GetAccountID(), 10),
		"positionTransactionIdList": internal.JoinStrings(transactionIDs),
	}

	resp, err := c.c.HttpRequest(url, "GET", nil, queryParams)
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

	url := fmt.Sprintf("%s/api/v2/private/account/getCollateralTransactionById", c.c.GetBaseURL())
	queryParams := map[string]string{
		"accountId":                   strconv.FormatInt(c.c.GetAccountID(), 10),
		"collateralTransactionIdList": internal.JoinStrings(transactionIDs),
	}

	resp, err := c.c.HttpRequest(url, "GET", nil, queryParams)
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
	url := fmt.Sprintf("%s/api/v2/private/account/getAccountDeleverageLight", c.c.GetBaseURL())
	queryParams := map[string]string{
		"accountId": strconv.FormatInt(c.c.GetAccountID(), 10),
	}

	resp, err := c.c.HttpRequest(url, "GET", nil, queryParams)
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
	url := fmt.Sprintf("%s/api/v2/private/account/updateLeverageSetting", c.c.GetBaseURL())
	data := map[string]interface{}{
		"accountId":  c.c.GetAccountID(),
		"contractId": contractID,
		"leverage":   leverage,
	}

	resp, err := c.c.HttpRequest(url, "POST", data, nil)
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

// SetMarginMode signs the margin-mode update with the trading key and submits the v2 request.
func (c *Client) SetMarginMode(ctx context.Context, params *SetMarginModeParams, md *metadata.MetaData) (*SetMarginModeResponse, error) {
	if params == nil {
		return nil, fmt.Errorf("params is nil")
	}
	if md == nil || md.Global == nil {
		return nil, fmt.Errorf("metadata.global is required")
	}

	chainID := strings.TrimSpace(md.Global.NativeChainId)
	if chainID == "" {
		chainID = strings.TrimSpace(md.Global.ChainId)
	}
	if chainID == "" {
		return nil, fmt.Errorf("metadata.global.nativeChainId/chainId is required")
	}
	verifyingContract := strings.TrimSpace(md.Global.ContractAddress)
	if verifyingContract == "" {
		return nil, fmt.Errorf("metadata.global.contractAddress is required")
	}

	clientOrderID := strings.TrimSpace(params.ClientOrderID)
	if clientOrderID == "" {
		clientOrderID = internal.GetRandomClientId()
	}
	l2Nonce := internal.CalcNonce(clientOrderID)
	l2ExpireTime := calcSetMarginModeL2ExpireTime(time.Now())

	marginMode := strings.TrimSpace(params.MarginMode)
	marginModeUint, err := strconv.ParseUint(marginMode, 10, 8)
	if err != nil {
		return nil, fmt.Errorf("invalid marginMode: %w", err)
	}

	tradingSigner, err := c.c.ResolveSignerAddress()
	if err != nil {
		return nil, fmt.Errorf("trading private key is required for v2 EIP-712 margin mode signing: %w", err)
	}
	if !common.IsHexAddress(tradingSigner) {
		return nil, fmt.Errorf("invalid signer address: %s", tradingSigner)
	}
	tradingSigner = common.HexToAddress(tradingSigner).Hex()

	domain, err := internal.NewTypedDataDomain("EdgeX", "1", chainID, verifyingContract)
	if err != nil {
		return nil, fmt.Errorf("failed to build EIP-712 domain: %w", err)
	}

	typedData := internal.TypedData{
		Types: internal.TypedDataTypes{
			"EIP712Domain": {
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
			"SetMarginPreferenceParams": {
				{Name: "accountId", Type: "uint64"},
				{Name: "assetId", Type: "uint64"},
				{Name: "marginMode", Type: "uint8"},
				{Name: "nonce", Type: "uint256"},
				{Name: "signer", Type: "address"},
			},
		},
		PrimaryType: "SetMarginPreferenceParams",
		Domain:      domain,
		Message: internal.TypedDataMessage{
			"accountId":  strconv.FormatInt(c.c.GetAccountID(), 10),
			"assetId":    strings.TrimSpace(params.ContractID),
			"marginMode": strconv.FormatUint(marginModeUint, 10),
			"nonce":      strconv.FormatInt(l2Nonce, 10),
			"signer":     tradingSigner,
		},
	}

	l2Signature, err := c.c.SignTypedDataWithSignerKey(typedData)
	if err != nil {
		return nil, fmt.Errorf("failed to sign margin mode payload: %w", err)
	}

	url := fmt.Sprintf("%s/api/v2/private/account/setMarginMode", c.c.GetBaseURL())
	data := map[string]interface{}{
		"accountId":    strconv.FormatInt(c.c.GetAccountID(), 10),
		"contractId":   strings.TrimSpace(params.ContractID),
		"marginMode":   marginMode,
		"l2Nonce":      strconv.FormatInt(l2Nonce, 10),
		"l2ExpireTime": l2ExpireTime,
		"signer":       tradingSigner,
		"l2Signature":  l2Signature,
	}

	resp, err := c.c.HttpRequest(url, "POST", data, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to set margin mode: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result SetMarginModeResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}

// GetAccountPage gets paginated account list (mainly used in v2).
func (c *Client) GetAccountPage(ctx context.Context, params GetAccountPageParams) (*PageDataAccountResponse, error) {
	url := fmt.Sprintf("%s/api/v2/private/account/getAccountPage", c.c.GetBaseURL())
	queryParams := map[string]string{
		"accountId": strconv.FormatInt(c.c.GetAccountID(), 10),
	}

	if params.Size > 0 {
		queryParams["size"] = strconv.FormatInt(int64(params.Size), 10)
	}
	if params.OffsetData != "" {
		queryParams["offsetData"] = params.OffsetData
	}

	resp, err := c.c.HttpRequest(url, "GET", nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get account page: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result PageDataAccountResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}
