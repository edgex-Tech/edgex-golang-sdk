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

	l2ExpireTime := time.Now().UnixMilli() + (37 * 24 * 60 * 60 * 1000) // 14 days
	l2ExpireHour := l2ExpireTime / (60 * 60 * 1000)
	expireTime := strconv.FormatInt(l2ExpireHour, 10)

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
		expireTime,
	)
	// fmt.Printf("assetId: %v,\nethAddress: %v,\naccountId: %v,\nclientRandomId: %v,\nnonceId: %v,\namount: %v,\nexpireTime: %v\n", coin.StarkExAssetId, params.EthAddress, accountID, clientRandomId, nonceId, normalizedAmount, expireTime)
	// msgHash -> hex
	fmt.Printf("msgHash: %v\n", hex.EncodeToString(msgHash))

	signature, err := c.Client.Sign(msgHash)
	if err != nil {
		return nil, fmt.Errorf("failed to sign withdrawal hash: %w", err)
	}
	sig_str := fmt.Sprintf("%s%s%s", signature.R, signature.S, signature.V)

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

// curl ^"https://testnet-internal.edgex.exchange/api/v1/private/assets/createCrossWithdraw^" ^
//
//	-H ^"accept: application/json, text/plain, */*^" ^
//	-H ^"accept-language: zh-CN,zh;q=0.9^" ^
//	-H ^"content-type: application/json^" ^
//	-b ^"_ga=GA1.1.994137743.1759840461; privy-session=privy.edgex.exchange; _vid_t=lZlxYLYRC0F+hIVka0v6hXk3zIjLPbIl6sS9MSVuKL7fV8fI6Uj+lq4KLnDoCiXk/sGnnx24KrwWW2GjIbCxbTa6uKFJxsM=; AMP_745eb4d7a4=JTdCJTIyZGV2aWNlSWQlMjIlM0ElMjI2MTM4NjNkNS1iZTA2LTRlZDgtYmQ3Yi0yN2YxN2FiNGRlMjklMjIlMkMlMjJ1c2VySWQlMjIlM0ElMjI2NzA0MjA3NDQ0NjI2NjM5MzglMjIlMkMlMjJzZXNzaW9uSWQlMjIlM0ExNzYwNzkzMjY0MzYwJTJDJTIyb3B0T3V0JTIyJTNBZmFsc2UlMkMlMjJsYXN0RXZlbnRUaW1lJTIyJTNBMTc2MDc5MzI2NTM4MiUyQyUyMmxhc3RFdmVudElkJTIyJTNBMjQlMkMlMjJwYWdlQ291bnRlciUyMiUzQTAlN0Q=; _cfuvid=jEdzYKmgxSpzsWvYPNwKI1YJrYuoMnPsRMAQqf2uQ_U-1760793266158-0.0.1.1-604800000; _ga_CC15H4GB9Q=GS2.1.s1760793267^$o8^$g0^$t1760793327^$j60^$l0^$h0; CF_AppSession=n1826f624d6d8eb3e; CF_Authorization=eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6IjFjODBmOGRkMGM3NmY5Zjg4NWM2ZmFjZmUzZTdmODdlMjgwYjA1NzM0MmM1MDAxZTk2N2U0NjVmNWFjMjRhNTkifQ.eyJ0eXBlIjoiYXBwIiwiaWF0IjoxNzYxMTQwMjc3LCJleHAiOjE3NjEyMjY2NzgsImlzcyI6Imh0dHBzOi8vZWRnZXgtdGVhbS5jbG91ZGZsYXJlYWNjZXNzLmNvbSIsInN1YiI6ImI0MTVjNDNlLTdjYTQtNWNhZi1iNjM1LTE5MTdhMTlkZTM1ZCIsImF1ZCI6IjE1YWFlODY0YjRkNjc5ZWI2ZDM1NTFlMDlmYWM1MzVjYTg4NDNhNGM0N2RmMjM1NmZjZjc4YTg1ZmNhZGQwNjkiLCJkZXZpY2VfaWQiOiJlMjQ1ZDBkMi05YmJjLTExZjAtOGUwOS1iYTkwMjJmYmU0YjciLCJlbWFpbCI6InJhcGhhZWxAZWRnZXguZXhjaGFuZ2UiLCJ3YXJwX2FzX2F1dGgiOnRydWV9.hGLLSNz6-jVR4xP9E_IN89_mniMVHHmMbhtPbSLjoTB5d2eOxtb69BfEAINPaRpD92ADFO7VrCGK5MdtPzCVefMgBfi9r0JlHOhD41Qm8m92QIt3bokJjRaADL6A_MXQ9L5fJewHsZpaX8OaPrHUlyWYBW3kaRuykZw-bo0O90wU2LVeOBiy9yl-DDEGRZZuR6CEm7_9i7WE-TIhTehCOG7i00t6fYxHqCEDM4w-nYxXvgkXGOSUApCO9PF-b3dLVtkkTLI3TFmDKTp_HT1VPxD7hZrDqSh2QJk7mtN_QfHVOjUk4Bgm3nk_bwtYNXC7BCXaZKwwbXvsNGhTgrYXag; _iidt=xvtNHx8xVZ9DhSaEkoxiPtmI8EmgEimZ/yXvMjLJe1wIHMJBEU0RikalBYVGWCHK6NGtBJBn6hhnS0sOn0tnoziQMBtVVL/582oc5Tk=; edgex_fp_data_t=shDNYICoZE17ZhL3xGAdXs9xTfrwPAGkXVuPVQ7BRWyhBoMUeQHZi4BOAcI0TX33AW61/eLbS98wL7AF74CnZSYXum4fF8Wfq++/VFM=; __cf_bm=cEzuzZcSPRWz2zMOEb8maZhOBj2DtRc7YF6IZvxT55Q-1761141470-1.0.1.1-FgnP0daTqyXhIxjdvgCt05J51ootHi_0T5XtA_9wIdzCBsjZ.d1E_qjQC1NLkdj1k7y63X0mNyYl0VlqnsgFm.kDnIV4p2APMD2Mxl9dtA0; privy-token=eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6InJSazNvekc1T2JKVklhN3JvQnp3dU1HR19xSTJQelRBX1ozZFYyR25rV1kifQ.eyJzaWQiOiJjbWdnanBiM3owMWIwam0wY3QyeTJjd3QwIiwiaXNzIjoicHJpdnkuaW8iLCJpYXQiOjE3NjExNDE0NzAsImF1ZCI6ImNseGNvcHFoNDA0NGpqb2g0aTB5MnV4MWgiLCJzdWIiOiJkaWQ6cHJpdnk6Y21nZ2pwYjVlMDFiMmptMGM4bjN6dXQ0cyIsImV4cCI6MTc2MTE0NTA3MH0.nfilTcTCeG5YBkoNNipw47H1ah2g-x8W4vRDEoumvUkiEzzbnMcmnBkEhBI7wVaHnoYtEzCxlu5iQtMA9wp7DA; privy-id-token=eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6InJSazNvekc1T2JKVklhN3JvQnp3dU1HR19xSTJQelRBX1ozZFYyR25rV1kifQ.eyJjciI6IjE3NTk4NDA2NzkiLCJsaW5rZWRfYWNjb3VudHMiOiJbe1widHlwZVwiOlwiZW1haWxcIixcImFkZHJlc3NcIjpcImNkb25nMDEwMUBnbWFpbC5jb21cIixcImx2XCI6MTc1OTg0MDY4MH0se1widHlwZVwiOlwid2FsbGV0XCIsXCJhZGRyZXNzXCI6XCIweDYxRERGZDQ0RTAwNDA0MWZlZWY4ODM3MjQzMGQ3M0U3MUYxRjQ2QTBcIixcImNoYWluX3R5cGVcIjpcImV0aGVyZXVtXCIsXCJ3YWxsZXRfY2xpZW50X3R5cGVcIjpcInByaXZ5XCIsXCJsdlwiOjE3NTk4NDA2ODF9LHtcInR5cGVcIjpcIndhbGxldFwiLFwiYWRkcmVzc1wiOlwiOURhRzYxUEJQN1I1QjhLOXhQRXp4SnNSU2lROHRicTJRbjNuWFpRNkJLRmVcIixcImNoYWluX3R5cGVcIjpcInNvbGFuYVwiLFwid2FsbGV0X2NsaWVudF90eXBlXCI6XCJwcml2eVwiLFwibHZcIjoxNzU5ODQwNjgxfSx7XCJ0eXBlXCI6XCJzbWFydF93YWxsZXRcIixcInNtYXJ0X3dhbGxldF90eXBlXCI6XCJrZXJuZWxcIixcImx2XCI6MTc1OTg0MDY4NixcImFkZHJlc3NcIjpcIjB4NjcwMURjN0I2NEFCRkMyNjFEMjFDRDVlMWY4ZWU3MWI2Yjc0YjMwNlwifV0iLCJpc3MiOiJwcml2eS5pbyIsImlhdCI6MTc2MTE0MTQ3MCwiYXVkIjoiY2x4Y29wcWg0MDQ0ampvaDRpMHkydXgxaCIsInN1YiI6ImRpZDpwcml2eTpjbWdnanBiNWUwMWIyam0wYzhuM3p1dDRzIiwiZXhwIjoxNzYxMTQ1MDcwfQ.6_X2qz6W5v_Q27QSHDmL9JfCpigx8fEv8vFtVwh_IEJXXhqp0tOl18jA5WdyA_HhLXapHsm8dILgKsKAD9XnfQ; AMP_5496e6fe0f=JTdCJTIyZGV2aWNlSWQlMjIlM0ElMjI3YzEwZTA3Zi1mZDhhLTRiZjItOTJmMS1jYzU4YjY1YzFhZDMlMjIlMkMlMjJ1c2VySWQlMjIlM0ElMjI2NzU1MTE2OTUxOTU1MDQ4OTclMjIlMkMlMjJzZXNzaW9uSWQlMjIlM0ExNzYxMTQwNDE4MDkzJTJDJTIyb3B0T3V0JTIyJTNBZmFsc2UlMkMlMjJsYXN0RXZlbnRUaW1lJTIyJTNBMTc2MTE0MTc5MjgzMiUyQyUyMmxhc3RFdmVudElkJTIyJTNBMTMyJTJDJTIycGFnZUNvdW50ZXIlMjIlM0EwJTdE; _ga_B6HWQM2XZ1=GS2.1.s1761140262^$o16^$g1^$t1761141861^$j60^$l0^$h0^" ^
//	-H ^"origin: https://testnet-internal.edgex.exchange^" ^
//	-H ^"priority: u=1, i^" ^
//	-H ^"referer: https://testnet-internal.edgex.exchange/dashboard/profile^" ^
//	-H ^"sec-ch-ua: ^\^"Google Chrome^\^";v=^\^"141^\^", ^\^"Not?A_Brand^\^";v=^\^"8^\^", ^\^"Chromium^\^";v=^\^"141^\^"^" ^
//	-H ^"sec-ch-ua-mobile: ?0^" ^
//	-H ^"sec-ch-ua-platform: ^\^"Windows^\^"^" ^
//	-H ^"sec-fetch-dest: empty^" ^
//	-H ^"sec-fetch-mode: cors^" ^
//	-H ^"sec-fetch-site: same-origin^" ^
//	-H ^"user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36^" ^
//	-H ^"x-edgex-api-key: eaa90200-42f1-796b-8605-b90986a7a4dc^" ^
//	-H ^"x-edgex-passphrase: ZyC-mNtgjH94Tvm3pgnaJA^" ^
//	-H ^"x-edgex-signature: ad08365af040c1bdbd94a6d331e858206a92e411243353faeabbe1c05aa929aa^" ^
//	-H ^"x-edgex-timestamp: 1761141862027^" ^
//	--data-raw ^"^{^\^"accountId^\^":^\^"675511695258419841^\^",^\^"coinId^\^":^\^"1000^\^",^\^"amount^\^":^\^"14^\^",^\^"ethAddress^\^":^\^"0x94C2bF0F2254eD91a5fBbc8c9F3f3433f18480D8^\^",^\^"erc20Address^\^":^\^"0x5E2522c505A543fA2714c617E3Cd133a6Daa9627^\^",^\^"lpAccountId^\^":^\^"642774866310726088^\^",^\^"clientCrossWithdrawId^\^":^\^"7436117733244607^\^",^\^"expireTime^\^":^\^"1762354800000^\^",^\^"l2Signature^\^":^\^"066ee29d0dbc051e2a726bcd32254f6a946ea72d063fb4a3a794db27202d9809062740bea0ed369833f79da2afc0548bb84cd3467dd4c4873322121508a0c7c1^\^",^\^"fee^\^":^\^"1^\^",^\^"chainId^\^":421614^}^"
//
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
