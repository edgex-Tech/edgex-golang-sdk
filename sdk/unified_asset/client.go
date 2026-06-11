package unified_asset

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/edgex-Tech/edgex-golang-sdk/sdk/internal"
)

const (
	UNIFIED_ASSET_BASE_PATH = "/api/v1/private/unified-asset"
	ZERO_ADDRESS            = "0x0000000000000000000000000000000000000000"
	NATIVE_TOKEN_ADDRESS    = "0x0000000000000000000000000000000000000000"
	EDGE_MAINNET_USDC       = "0x98d2919b9A214E6Fa5384AC81E6864bA686Ad74c"
	EDGE_TESTNET_USDC       = "0x5a4f218dcb4e257a4586159e9ee925b0fa1df610"
)

var (
	withdrawProfiles = map[string]map[string]interface{}{
		"mainnet-usdc": {
			"asset":         "usdc",
			"network":       "mainnet",
			"source":        "spot",
			"chain_id":      3343,
			"token_address": EDGE_MAINNET_USDC,
		},
		"testnet-usdc": {
			"asset":         "usdc",
			"network":       "testnet",
			"source":        "spot",
			"chain_id":      33431,
			"token_address": EDGE_TESTNET_USDC,
		},
		"mainnet-eth-native": {
			"asset":         "eth-native",
			"network":       "mainnet",
			"source":        "spot",
			"chain_id":      42161,
			"token_address": NATIVE_TOKEN_ADDRESS,
		},
		"testnet-eth-native": {
			"asset":         "eth-native",
			"network":       "testnet",
			"source":        "spot",
			"chain_id":      421614,
			"token_address": NATIVE_TOKEN_ADDRESS,
		},
	}
	assetProfileMap = map[string]string{
		"usdc/mainnet":       "mainnet-usdc",
		"usdc/testnet":       "testnet-usdc",
		"eth-native/mainnet": "mainnet-eth-native",
		"eth-native/testnet": "testnet-eth-native",
	}
	sfLock   sync.Mutex
	sfLastTS int64
	sfSeq    int64
	sfEpoch  int64 = 1577836800000
)

type clientInterface interface {
	GetAccountID() int64
	GetBaseURL() string
	ResolveWalletAddress() (string, error)
	SignTypedDataWithWalletKey(typedData internal.TypedData) (string, error)
	HttpRequest(urlStr string, method string, data map[string]interface{}, params map[string]string) (*http.Response, error)
}

type Client struct {
	c clientInterface
}

type CreateWithdrawParams struct {
	AmountRaw          string
	UserAddress        string
	TokenAddress       string
	ChainID            int
	Source             string
	SourceAccount      string
	Profile            string
	Asset              string
	Network            string
	PrivyAddress       string
	PrivyIdentityToken string
	ClientWithdrawID   string
	ExpireSeconds      int
	ExpireTime         int64
	ExtraData          string
}

type CreateSpotDepositParams struct {
	AmountRaw     string
	UserAddress   string
	TokenAddress  string
	ChainID       int
	SpotAccountID string
	SourceAccount string
	PrivyAddress  string
	ExtraData     string
}

type gatewayEnvelope struct {
	Code    string      `json:"code"`
	Data    interface{} `json:"data"`
	Msg     string      `json:"msg"`
	TraceID string      `json:"traceId"`
}

func NewClient(client clientInterface) *Client {
	return &Client{c: client}
}

func buildWithdrawAttempt(userAddress, source, sourceAccount, tokenAddress, amountRaw string, chainID int, expireTime int64, clientWithdrawID, privyAddress string) map[string]interface{} {
	if strings.TrimSpace(privyAddress) == "" {
		privyAddress = ZERO_ADDRESS
	}
	return map[string]interface{}{
		"userAddress":        userAddress,
		"privyAddress":       privyAddress,
		"source":             source,
		"sourceAccount":      sourceAccount,
		"tokenAddress":       tokenAddress,
		"amount":             amountRaw,
		"fee":                "0",
		"destination":        fmt.Sprintf("chain-%d", chainID),
		"destinationAccount": userAddress,
		"clientWithdrawId":   clientWithdrawID,
		"expireTime":         expireTime,
	}
}

func buildSpotDepositAttempt(userAddress, privyAddress, sourceAccount, tokenAddress, amountRaw string, chainID int, spotAccountID string) map[string]interface{} {
	if strings.TrimSpace(privyAddress) == "" {
		privyAddress = ZERO_ADDRESS
	}
	return map[string]interface{}{
		"userAddress":        userAddress,
		"privyAddress":       privyAddress,
		"source":             fmt.Sprintf("chain-%d", chainID),
		"sourceAccount":      sourceAccount,
		"tokenAddress":       tokenAddress,
		"amount":             amountRaw,
		"fee":                "0",
		"destination":        "spot",
		"destinationAccount": spotAccountID,
	}
}

func applyFeeToAttempt(attempt map[string]interface{}, fee string) error {
	grossAmount := new(big.Int)
	if _, ok := grossAmount.SetString(strings.TrimSpace(fmt.Sprint(attempt["amount"])), 10); !ok {
		return fmt.Errorf("invalid attempt amount: %s", fmt.Sprint(attempt["amount"]))
	}
	feeAmount := new(big.Int)
	if _, ok := feeAmount.SetString(strings.TrimSpace(fee), 10); !ok {
		return fmt.Errorf("invalid fee: %s", fee)
	}
	if feeAmount.Sign() < 0 {
		return fmt.Errorf("fee must be non-negative")
	}
	if feeAmount.Cmp(grossAmount) >= 0 {
		return fmt.Errorf("fee must be less than gross amount: fee=%s, gross=%s", feeAmount.String(), grossAmount.String())
	}
	netAmount := new(big.Int).Sub(grossAmount, feeAmount)
	attempt["fee"] = feeAmount.String()
	attempt["amount"] = netAmount.String()
	return nil
}

func profileNameForAsset(asset, network string) (string, error) {
	key := strings.ToLower(strings.TrimSpace(asset)) + "/" + strings.ToLower(strings.TrimSpace(network))
	profile, ok := assetProfileMap[key]
	if !ok {
		known := make([]string, 0, len(assetProfileMap))
		for item := range assetProfileMap {
			known = append(known, item)
		}
		sort.Strings(known)
		return "", fmt.Errorf("unknown asset/network %q/%q; known: %s", asset, network, strings.Join(known, ", "))
	}
	return profile, nil
}

func resolveWithdrawProfile(profileName string) (map[string]interface{}, error) {
	profile, ok := withdrawProfiles[profileName]
	if !ok {
		known := make([]string, 0, len(withdrawProfiles))
		for item := range withdrawProfiles {
			known = append(known, item)
		}
		sort.Strings(known)
		return nil, fmt.Errorf("unknown withdraw profile %q; known: %s", profileName, strings.Join(known, ", "))
	}
	clone := make(map[string]interface{}, len(profile))
	for k, v := range profile {
		clone[k] = v
	}
	return clone, nil
}

func nextSnowflakeID() int64 {
	sfLock.Lock()
	defer sfLock.Unlock()

	now := time.Now().UnixMilli()
	if now == sfLastTS {
		sfSeq = (sfSeq + 1) & 0xFFF
		if sfSeq == 0 {
			for now <= sfLastTS {
				now = time.Now().UnixMilli()
			}
		}
	} else {
		sfSeq = 0
	}
	sfLastTS = now
	nodeID := int64(1)
	return ((now - sfEpoch) << 22) | ((nodeID & 0x3FF) << 12) | sfSeq
}

func unwrapResponse(body []byte) (map[string]interface{}, error) {
	var response gatewayEnvelope
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	if response.Code != "" && response.Code != "SUCCESS" {
		return nil, fmt.Errorf("[%s] %s (trace=%s)", response.Code, response.Msg, response.TraceID)
	}
	if data, ok := response.Data.(map[string]interface{}); ok {
		return data, nil
	}
	return map[string]interface{}{}, nil
}

func buildTypedDataFromServerResponse(data map[string]interface{}) (internal.TypedData, error) {
	if _, ok := data["primaryType"]; !ok {
		if wrapped, ok := data["data"].(map[string]interface{}); ok {
			data = wrapped
		}
	}

	rawTypes, ok := data["types"].(map[string]interface{})
	if !ok {
		return internal.TypedData{}, fmt.Errorf("types missing from server response")
	}

	types := internal.TypedDataTypes{}
	for typeName, rawDefinition := range rawTypes {
		var rawFields []interface{}
		if definitionMap, ok := rawDefinition.(map[string]interface{}); ok {
			if fields, ok := definitionMap["fields"].([]interface{}); ok {
				rawFields = fields
			}
		} else if fields, ok := rawDefinition.([]interface{}); ok {
			rawFields = fields
		}

		if rawFields == nil {
			continue
		}

		fields := make([]internal.TypedDataType, 0, len(rawFields))
		for _, rawField := range rawFields {
			fieldMap, ok := rawField.(map[string]interface{})
			if !ok {
				continue
			}
			fields = append(fields, internal.TypedDataType{
				Name: fmt.Sprint(fieldMap["name"]),
				Type: fmt.Sprint(fieldMap["type"]),
			})
		}
		types[typeName] = fields
	}

	domainMap, _ := data["domain"].(map[string]interface{})
	if domainMap == nil {
		domainMap = map[string]interface{}{}
	}

	stringValue := func(v interface{}) string {
		if v == nil {
			return ""
		}
		return fmt.Sprint(v)
	}

	if _, ok := types["EIP712Domain"]; !ok {
		domainFields := make([]internal.TypedDataType, 0, 4)
		if stringValue(domainMap["name"]) != "" {
			domainFields = append(domainFields, internal.TypedDataType{Name: "name", Type: "string"})
		}
		if stringValue(domainMap["version"]) != "" {
			domainFields = append(domainFields, internal.TypedDataType{Name: "version", Type: "string"})
		}
		if stringValue(domainMap["chainId"]) != "" {
			domainFields = append(domainFields, internal.TypedDataType{Name: "chainId", Type: "uint256"})
		}
		if stringValue(domainMap["verifyingContract"]) != "" {
			domainFields = append(domainFields, internal.TypedDataType{Name: "verifyingContract", Type: "address"})
		}
		types["EIP712Domain"] = domainFields
	}

	domain, err := internal.NewTypedDataDomain(
		stringValue(domainMap["name"]),
		stringValue(domainMap["version"]),
		stringValue(domainMap["chainId"]),
		stringValue(domainMap["verifyingContract"]),
	)
	if err != nil {
		return internal.TypedData{}, fmt.Errorf("failed to normalize typed data domain: %w", err)
	}

	message := internal.TypedDataMessage{}
	if rawMessageJSON, ok := data["messageJson"]; ok && fmt.Sprint(rawMessageJSON) != "" {
		if err := json.Unmarshal([]byte(fmt.Sprint(rawMessageJSON)), &message); err != nil {
			return internal.TypedData{}, fmt.Errorf("failed to decode messageJson: %w", err)
		}
	} else if rawMessage, ok := data["message"].(map[string]interface{}); ok {
		for k, v := range rawMessage {
			message[k] = v
		}
	}

	return internal.TypedData{
		Types:       types,
		PrimaryType: fmt.Sprint(data["primaryType"]),
		Domain:      domain,
		Message:     message,
	}, nil
}

func (c *Client) request(ctx context.Context, method, path string, data map[string]interface{}, params map[string]string) (map[string]interface{}, error) {
	_ = ctx
	resp, err := c.c.HttpRequest(c.c.GetBaseURL()+path, method, data, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	return unwrapResponse(body)
}

func (c *Client) GetFeeByAssetFlow(ctx context.Context, attempt map[string]interface{}) (map[string]interface{}, error) {
	return c.request(ctx, "POST", UNIFIED_ASSET_BASE_PATH+"/getFeeByAssetFlow", map[string]interface{}{"attempt": attempt}, nil)
}

func (c *Client) GetEIP712Data(ctx context.Context, attempt map[string]interface{}) (map[string]interface{}, error) {
	return c.request(ctx, "POST", UNIFIED_ASSET_BASE_PATH+"/getEIP712Data", map[string]interface{}{"attempt": attempt}, nil)
}

func (c *Client) GetDepositData(ctx context.Context, attempt map[string]interface{}, extraData string) (map[string]interface{}, error) {
	return c.request(ctx, "POST", UNIFIED_ASSET_BASE_PATH+"/getDepositData", map[string]interface{}{"attempt": attempt, "extraData": extraData}, nil)
}

func (c *Client) SubmitAssetFlow(ctx context.Context, attempt map[string]interface{}, userSignature, privyIdentityToken, extraData string) (map[string]interface{}, error) {
	return c.request(ctx, "POST", UNIFIED_ASSET_BASE_PATH+"/submitAssetFlow", map[string]interface{}{
		"attempt":            attempt,
		"userSignature":      userSignature,
		"extraData":          extraData,
		"privyIdentityToken": privyIdentityToken,
	}, nil)
}

func (c *Client) QueryAssetFlows(ctx context.Context, params map[string]string) (map[string]interface{}, error) {
	return c.request(ctx, "GET", UNIFIED_ASSET_BASE_PATH+"/queryAssetFlows", nil, params)
}

func (c *Client) CreateWithdraw(ctx context.Context, params CreateWithdrawParams) (map[string]interface{}, error) {
	profileName := strings.TrimSpace(params.Profile)
	if profileName == "" {
		assetName := strings.TrimSpace(params.Asset)
		if assetName == "" {
			assetName = "usdc"
		}
		networkName := strings.TrimSpace(params.Network)
		if networkName == "" {
			networkName = "mainnet"
		}
		var err error
		profileName, err = profileNameForAsset(assetName, networkName)
		if err != nil {
			return nil, err
		}
	}
	profile, err := resolveWithdrawProfile(profileName)
	if err != nil {
		return nil, err
	}

	source := strings.TrimSpace(params.Source)
	if source == "" {
		source = fmt.Sprint(profile["source"])
	}
	sourceAccount := strings.TrimSpace(params.SourceAccount)
	if sourceAccount == "" {
		sourceAccount = strconv.FormatInt(c.c.GetAccountID(), 10)
	}
	tokenAddress := strings.TrimSpace(params.TokenAddress)
	if tokenAddress == "" {
		tokenAddress = fmt.Sprint(profile["token_address"])
	}
	chainID := params.ChainID
	if chainID == 0 {
		chainID = profile["chain_id"].(int)
	}
	if tokenAddress == "" {
		return nil, fmt.Errorf("token_address is required")
	}
	if chainID <= 0 {
		return nil, fmt.Errorf("chain_id is required")
	}

	clientWithdrawID := strings.TrimSpace(params.ClientWithdrawID)
	if clientWithdrawID == "" {
		clientWithdrawID = strconv.FormatInt(nextSnowflakeID(), 10)
	}
	expireTime := params.ExpireTime
	if expireTime == 0 {
		expireSeconds := params.ExpireSeconds
		if expireSeconds == 0 {
			expireSeconds = 300
		}
		expireTime = time.Now().Unix() + int64(expireSeconds)
	}

	attempt := buildWithdrawAttempt(
		params.UserAddress,
		source,
		sourceAccount,
		tokenAddress,
		params.AmountRaw,
		chainID,
		expireTime,
		clientWithdrawID,
		params.PrivyAddress,
	)

	feeData, err := c.GetFeeByAssetFlow(ctx, attempt)
	if err != nil {
		return nil, err
	}
	if err := applyFeeToAttempt(attempt, fmt.Sprint(feeData["fee"])); err != nil {
		return nil, err
	}

	eip712Data, err := c.GetEIP712Data(ctx, attempt)
	if err != nil {
		return nil, err
	}
	typedData, err := buildTypedDataFromServerResponse(eip712Data)
	if err != nil {
		return nil, err
	}
	signature, err := c.c.SignTypedDataWithWalletKey(typedData)
	if err != nil {
		return nil, fmt.Errorf("failed to sign withdraw payload: %w", err)
	}

	return c.SubmitAssetFlow(ctx, attempt, signature, params.PrivyIdentityToken, params.ExtraData)
}

func (c *Client) GetSpotDepositData(ctx context.Context, params CreateSpotDepositParams) (map[string]interface{}, error) {
	if strings.TrimSpace(params.TokenAddress) == "" {
		return nil, fmt.Errorf("token_address is required")
	}
	if params.ChainID <= 0 {
		return nil, fmt.Errorf("chain_id is required")
	}

	spotAccountID := strings.TrimSpace(params.SpotAccountID)
	if spotAccountID == "" {
		spotAccountID = strconv.FormatInt(c.c.GetAccountID(), 10)
	}
	sourceAccount := strings.TrimSpace(params.SourceAccount)
	if sourceAccount == "" {
		sourceAccount = params.UserAddress
	}
	attempt := buildSpotDepositAttempt(
		params.UserAddress,
		params.PrivyAddress,
		sourceAccount,
		params.TokenAddress,
		params.AmountRaw,
		params.ChainID,
		spotAccountID,
	)
	return c.GetDepositData(ctx, attempt, params.ExtraData)
}
