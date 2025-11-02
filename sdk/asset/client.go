package asset

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/edgex-Tech/edgex-golang-sdk/sdk/internal"
	"github.com/edgex-Tech/edgex-golang-sdk/sdk/metadata"
	"github.com/shopspring/decimal"
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
	url := fmt.Sprintf("%s/api/v1/private/assets/getAllOrdersPage", c.Client.GetBaseURL())
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
	url := fmt.Sprintf("%s/api/v1/private/assets/getCoinRate", c.Client.GetBaseURL())
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
	url := fmt.Sprintf("%s/api/v1/private/assets/getCrossWithdrawById", c.Client.GetBaseURL())
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
	url := fmt.Sprintf("%s/api/v1/private/assets/getCrossWithdrawSignInfo", c.Client.GetBaseURL())
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
	url := fmt.Sprintf("%s/api/v1/private/assets/getFastWithdrawById", c.Client.GetBaseURL())
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
	url := fmt.Sprintf("%s/api/v1/private/assets/getFastWithdrawSignInfo", c.Client.GetBaseURL())
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
	url := fmt.Sprintf("%s/api/v1/private/assets/getNormalWithdrawById", c.Client.GetBaseURL())
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
	url := fmt.Sprintf("%s/api/v1/private/assets/getNormalWithdrawableAmount", c.Client.GetBaseURL())
	queryParams := map[string]string{}

	accountID := strconv.FormatInt(c.Client.GetAccountID(), 10)
	if params.Address != "" {
		queryParams["address"] = params.Address
		queryParams["accountId"] = accountID
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

func GetNonceFromClientId(clientId string) string {
	hash := sha256.Sum256([]byte(clientId))
	hashHex := hex.EncodeToString(hash[:])
	s := hashHex[:8]

	val, _ := strconv.ParseInt(s, 16, 64)
	return strconv.FormatInt(val, 10)
}

// CreateNormalWithdraw creates a normal withdrawal order
func (c *Client) CreateNormalWithdraw(ctx context.Context, params *CreateNormalWithdrawParams, md *metadata.MetaData) (*ResultCreateNormalWithdraw, error) {
	url := fmt.Sprintf("%s/api/v1/private/assets/createNormalWithdraw", c.Client.GetBaseURL())

	var coin *metadata.Coin
	if md != nil && md.CoinList != nil {
		for i := range md.CoinList {
			if md.CoinList[i].CoinId == params.CoinId {
				coin = &md.CoinList[i]
				break
			}
		}
	}

	if coin == nil {
		return nil, fmt.Errorf("coin not found: %s", params.CoinId)
	}

	accountID := strconv.FormatInt(c.Client.GetAccountID(), 10)
	clientRandomId := internal.GetRandomClientId()
	nonceId := GetNonceFromClientId(clientRandomId)

	l2ExpireTime := time.Now().UnixMilli() + (14 * 24 * 60 * 60 * 1000) // 14 days
	l2ExpireHour := l2ExpireTime / (60 * 60 * 1000)
	expireTime := strconv.FormatInt(l2ExpireTime, 10)

	ammount, err := decimal.NewFromString(params.Amount)
	if err != nil {
		return nil, fmt.Errorf("failed to parse amount: %w", err)
	}
	normalizedAmount := ammount.Mul(decimal.NewFromInt(1000000)).Floor().String()

	// Calculate withdraw hash and sign it
	msgHash := internal.CalcWithdrawalHash(
		coin.StarkExAssetId,
		params.EthAddress,
		accountID,
		nonceId,
		normalizedAmount,
		strconv.FormatInt(l2ExpireHour, 10),
	)

	signature, err := c.Client.Sign(msgHash)
	if err != nil {
		return nil, fmt.Errorf("failed to sign withdrawal hash: %w", err)
	}
	sig_str := fmt.Sprintf("%s%s", signature.R, signature.S)

	body := map[string]interface{}{
		"accountId":        accountID,
		"coinId":           params.CoinId,
		"amount":           params.Amount,
		"ethAddress":       params.EthAddress,
		"clientWithdrawId": clientRandomId,
		"expireTime":       expireTime,
		"l2Signature":      sig_str,
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
	fmt.Println(string(respBody))

	var result ResultCreateNormalWithdraw
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %v", result)
	}

	return &result, nil
}

// CreateCrossWithdraw creates a cross-chain withdrawal order
func (c *Client) CreateCrossWithdraw(ctx context.Context, params CreateCrossWithdrawParams) (*ResultCreateCrossWithdraw, error) {
	url := fmt.Sprintf("%s/api/v1/private/assets/createCrossWithdraw", c.Client.GetBaseURL())

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
