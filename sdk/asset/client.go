package asset

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/edgex-Tech/edgex-golang-sdk/sdk/internal"
	"github.com/edgex-Tech/edgex-golang-sdk/sdk/metadata"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

// Client represents the new asset client without OpenAPI dependencies
type Client struct {
	c clientInterface
}

type clientInterface interface {
	GetAccountID() int64
	GetWalletPriKey() string
	GetBaseURL() string
	ResolveWalletAddress() (string, error)
	SignTypedDataWithWalletKey(typedData internal.TypedData) (string, error)
	HttpRequest(urlStr string, method string, data map[string]interface{}, params map[string]string) (*http.Response, error)
}

// NewClient creates a new asset client
func NewClient(client clientInterface) *Client {
	return &Client{
		c: client,
	}
}

// GetAllOrdersPage gets all asset orders with pagination
func (c *Client) GetAllOrdersPage(ctx context.Context, params GetAllOrdersPageParams) (*ResultPageDataAssetOrder, error) {
	url := fmt.Sprintf("%s/api/v2/private/assets/getAllOrdersPage", c.c.GetBaseURL())
	queryParams := map[string]string{
		"accountId": strconv.FormatInt(c.c.GetAccountID(), 10),
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

	resp, err := c.c.HttpRequest(url, "GET", nil, queryParams)
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

func GetNonceFromClientId(clientId string) string {
	hash := sha256.Sum256([]byte(clientId))
	hashHex := hex.EncodeToString(hash[:])
	s := hashHex[:8]

	val, _ := strconv.ParseInt(s, 16, 64)
	return strconv.FormatInt(val, 10)
}

func resolveCoinFromMetadata(md *metadata.MetaData, coinID string) (*metadata.Coin, error) {
	if md == nil {
		return nil, fmt.Errorf("metadata is nil")
	}
	coinID = strings.TrimSpace(coinID)
	if coinID == "" {
		return nil, fmt.Errorf("coin id is empty")
	}
	for i := range md.CoinList {
		if strings.TrimSpace(md.CoinList[i].CoinId) == coinID {
			return &md.CoinList[i], nil
		}
	}
	return nil, fmt.Errorf("coin not found: %s", coinID)
}

func resolveDefaultCoinID(md *metadata.MetaData) (string, error) {
	if md == nil {
		return "", fmt.Errorf("metadata is nil")
	}
	if md.Global != nil && strings.TrimSpace(md.Global.TransferCoinId) != "" {
		return strings.TrimSpace(md.Global.TransferCoinId), nil
	}
	for _, coin := range md.CoinList {
		coinID := strings.TrimSpace(coin.CoinId)
		if coinID != "" {
			return coinID, nil
		}
	}
	return "", fmt.Errorf("no valid coin id in metadata")
}

func resolveCoinResolution(coin *metadata.Coin) int {
	if coin == nil {
		return 6
	}
	if parsed, err := strconv.Atoi(strings.TrimSpace(coin.Resolution)); err == nil && parsed > 0 {
		return parsed
	}
	return 6
}

func normalizeDomainChainID(rawChainID string) (string, error) {
	rawChainID = strings.TrimSpace(rawChainID)
	if rawChainID == "" {
		return "", fmt.Errorf("chain id is empty")
	}
	if strings.HasPrefix(rawChainID, "0x") || strings.HasPrefix(rawChainID, "0X") {
		chainIDBigInt := new(big.Int)
		if _, ok := chainIDBigInt.SetString(rawChainID[2:], 16); !ok {
			return "", fmt.Errorf("invalid hex chain id: %s", rawChainID)
		}
		return chainIDBigInt.String(), nil
	}
	chainIDBigInt := new(big.Int)
	if _, ok := chainIDBigInt.SetString(rawChainID, 10); !ok {
		return "", fmt.Errorf("invalid chain id: %s", rawChainID)
	}
	return chainIDBigInt.String(), nil
}

func pickPreferredTokenAddress(chain *metadata.Chain) string {
	if chain == nil {
		return ""
	}
	for _, token := range chain.TokenList {
		if strings.EqualFold(strings.TrimSpace(token.Token), "USDC") && strings.TrimSpace(token.TokenAddress) != "" {
			return strings.TrimSpace(token.TokenAddress)
		}
	}
	for _, token := range chain.TokenList {
		tokenAddress := strings.TrimSpace(token.TokenAddress)
		if tokenAddress != "" {
			return tokenAddress
		}
	}
	return ""
}

func resolveWithdrawChainAndToken(md *metadata.MetaData, chainID, tokenAddress string) (string, string, error) {
	if md == nil || md.MultiChain == nil || len(md.MultiChain.ChainList) == 0 {
		return "", "", fmt.Errorf("metadata multiChain chainList is empty")
	}

	chainID = strings.TrimSpace(chainID)
	tokenAddress = strings.TrimSpace(tokenAddress)

	if chainID != "" {
		for i := range md.MultiChain.ChainList {
			chain := &md.MultiChain.ChainList[i]
			if strings.TrimSpace(chain.ChainId) != chainID {
				continue
			}
			if tokenAddress == "" {
				tokenAddress = pickPreferredTokenAddress(chain)
				if tokenAddress == "" {
					return "", "", fmt.Errorf("no tokenAddress found in chain: %s", chainID)
				}
				return chainID, tokenAddress, nil
			}
			for _, token := range chain.TokenList {
				if strings.EqualFold(strings.TrimSpace(token.TokenAddress), tokenAddress) {
					return chainID, strings.TrimSpace(token.TokenAddress), nil
				}
			}
			return "", "", fmt.Errorf("tokenAddress %s not found in chain %s", tokenAddress, chainID)
		}
		return "", "", fmt.Errorf("chain id not found: %s", chainID)
	}

	if tokenAddress != "" {
		for i := range md.MultiChain.ChainList {
			chain := &md.MultiChain.ChainList[i]
			for _, token := range chain.TokenList {
				if strings.EqualFold(strings.TrimSpace(token.TokenAddress), tokenAddress) {
					return strings.TrimSpace(chain.ChainId), strings.TrimSpace(token.TokenAddress), nil
				}
			}
		}
		return "", "", fmt.Errorf("tokenAddress not found in metadata: %s", tokenAddress)
	}

	for i := range md.MultiChain.ChainList {
		chain := &md.MultiChain.ChainList[i]
		addr := pickPreferredTokenAddress(chain)
		if strings.TrimSpace(chain.ChainId) != "" && addr != "" {
			return strings.TrimSpace(chain.ChainId), addr, nil
		}
	}

	return "", "", fmt.Errorf("no valid chainId/tokenAddress for withdraw sign info")
}

func resolveWithdrawFee(resp *ResultGetWithdrawSignInfo) string {
	if resp != nil && resp.Data != nil && resp.Data.Fee != nil {
		fee := strings.TrimSpace(*resp.Data.Fee)
		if fee != "" {
			return fee
		}
	}
	return "0"
}

// PrepareWithdrawSignInfo resolves chain/token and queries withdraw sign info with fee.
func (c *Client) PrepareWithdrawSignInfo(ctx context.Context, md *metadata.MetaData, params PrepareWithdrawSignInfoParams) (*PreparedWithdrawSignInfo, error) {
	if md == nil {
		return nil, fmt.Errorf("metadata is nil")
	}

	amount := strings.TrimSpace(params.Amount)
	if amount == "" {
		return nil, fmt.Errorf("amount is required")
	}

	coinID := strings.TrimSpace(params.CoinId)
	if coinID == "" {
		resolvedCoinID, err := resolveDefaultCoinID(md)
		if err != nil {
			return nil, err
		}
		coinID = resolvedCoinID
	}

	chainID, tokenAddress, err := resolveWithdrawChainAndToken(md, params.ChainId, params.TokenAddress)
	if err != nil {
		return nil, err
	}

	signInfo, err := c.GetWithdrawSignInfo(ctx, GetWithdrawSignInfoParams{
		ChainId:      chainID,
		TokenAddress: tokenAddress,
		Amount:       amount,
	})
	if err != nil {
		return nil, err
	}

	return &PreparedWithdrawSignInfo{
		CoinId:       coinID,
		ChainId:      chainID,
		TokenAddress: tokenAddress,
		Amount:       amount,
		Fee:          resolveWithdrawFee(signInfo),
		SignInfo:     signInfo,
	}, nil
}

// BuildWithdrawV2Signature builds EIP-712 withdraw signature payload used by V2 withdraw APIs.
func (c *Client) BuildWithdrawV2Signature(md *metadata.MetaData, params BuildWithdrawV2SignatureParams) (*WithdrawV2SignatureInfo, error) {
	if md == nil || md.Global == nil {
		return nil, fmt.Errorf("metadata global is nil")
	}

	coin, err := resolveCoinFromMetadata(md, params.CoinId)
	if err != nil {
		return nil, err
	}

	amount := strings.TrimSpace(params.Amount)
	if amount == "" {
		return nil, fmt.Errorf("amount is required")
	}

	fee := strings.TrimSpace(params.Fee)
	if fee == "" {
		fee = "0"
	}

	clientWithdrawID := strings.TrimSpace(params.ClientWithdrawId)
	if clientWithdrawID == "" {
		return nil, fmt.Errorf("clientWithdrawId is required")
	}

	resolution := resolveCoinResolution(coin)
	decimals := decimal.NewFromInt(10).Pow(decimal.NewFromInt(int64(resolution)))

	amountDec, err := decimal.NewFromString(amount)
	if err != nil {
		return nil, fmt.Errorf("failed to parse amount: %w", err)
	}
	feeDec, err := decimal.NewFromString(fee)
	if err != nil {
		return nil, fmt.Errorf("failed to parse fee: %w", err)
	}
	typedAmount := amountDec.Mul(decimals).Floor().String()
	typedFee := feeDec.Mul(decimals).Floor().String()

	signerAddress, err := c.c.ResolveWalletAddress()
	if err != nil {
		return nil, fmt.Errorf("failed to resolve signer address: %w", err)
	}
	if !common.IsHexAddress(signerAddress) {
		return nil, fmt.Errorf("invalid signer address: %s", signerAddress)
	}
	signerAddress = common.HexToAddress(signerAddress).Hex()

	toAddress := strings.TrimSpace(params.ToAddress)
	if toAddress == "" {
		toAddress = signerAddress
	}
	if !common.IsHexAddress(toAddress) {
		return nil, fmt.Errorf("invalid toAddress: %s", toAddress)
	}
	toAddress = common.HexToAddress(toAddress).Hex()

	domainChainIDRaw := strings.TrimSpace(md.Global.ChainId)
	if domainChainIDRaw == "" {
		domainChainIDRaw = strings.TrimSpace(md.Global.NativeChainId)
	}
	domainChainID, err := normalizeDomainChainID(domainChainIDRaw)
	if err != nil {
		return nil, err
	}

	domain, err := internal.NewTypedDataDomain(
		"EdgeX",
		"1",
		domainChainID,
		md.Global.ContractAddress,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to build EIP-712 domain: %w", err)
	}

	nonceID := GetNonceFromClientId(clientWithdrawID)
	expireMs := time.Now().UnixMilli() + (14 * 24 * 60 * 60 * 1000)
	expireSec := expireMs / 1000

	typedData := internal.TypedData{
		Types: internal.TypedDataTypes{
			"EIP712Domain": {
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
			"WithdrawalParams": {
				{Name: "from", Type: "uint64"},
				{Name: "toAddress", Type: "address"},
				{Name: "amount", Type: "int256"},
				{Name: "feeAmount", Type: "int256"},
				{Name: "nonce", Type: "uint256"},
				{Name: "expirationTimestamp", Type: "uint256"},
				{Name: "signer", Type: "address"},
			},
		},
		PrimaryType: "WithdrawalParams",
		Domain:      domain,
		Message: internal.TypedDataMessage{
			"from":                strconv.FormatInt(c.c.GetAccountID(), 10),
			"toAddress":           toAddress,
			"amount":              typedAmount,
			"feeAmount":           typedFee,
			"nonce":               nonceID,
			"expirationTimestamp": strconv.FormatInt(expireSec, 10),
			"signer":              signerAddress,
		},
	}

	signature, err := c.c.SignTypedDataWithWalletKey(typedData)
	if err != nil {
		return nil, fmt.Errorf("failed to sign withdrawal: %w", err)
	}

	return &WithdrawV2SignatureInfo{
		Signature:    signature,
		Signer:       signerAddress,
		Nonce:        nonceID,
		L2ExpireTime: strconv.FormatInt(expireMs, 10),
		ToAddress:    toAddress,
	}, nil
}

// CreateNormalWithdraw creates a normal withdrawal order
func (c *Client) CreateNormalWithdraw(ctx context.Context, params *CreateNormalWithdrawParams, md *metadata.MetaData) (*ResultCreateNormalWithdraw, error) {
	url := fmt.Sprintf("%s/api/v2/private/assets/createNormalWithdraw", c.c.GetBaseURL())

	if params == nil {
		return nil, fmt.Errorf("params is nil")
	}

	clientWithdrawID := internal.GetRandomClientId()
	signInfo, err := c.BuildWithdrawV2Signature(md, BuildWithdrawV2SignatureParams{
		CoinId:           params.CoinId,
		Amount:           params.Amount,
		Fee:              params.Fee,
		ClientWithdrawId: clientWithdrawID,
		ToAddress:        params.EthAddress,
	})
	if err != nil {
		return nil, err
	}

	// Convert nonce to integer for normal withdraw (backend expects Long type)
	nonceInt, err := strconv.ParseInt(signInfo.Nonce, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse nonce: %w", err)
	}

	body := map[string]interface{}{
		"accountId":        strconv.FormatInt(c.c.GetAccountID(), 10),
		"coinId":           params.CoinId,
		"amount":           params.Amount,
		"fee":              params.Fee,
		"ethAddress":       signInfo.ToAddress,
		"clientWithdrawId": clientWithdrawID,
		"signature":        signInfo.Signature,
		"signer":           signInfo.Signer,
		"nonce":            nonceInt,
		"l2ExpireTime":     signInfo.L2ExpireTime,
	}

	resp, err := c.c.HttpRequest(url, "POST", body, nil)
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

// GetWithdrawSignInfo queries withdrawal signature information
func (c *Client) GetWithdrawSignInfo(ctx context.Context, params GetWithdrawSignInfoParams) (*ResultGetWithdrawSignInfo, error) {
	url := fmt.Sprintf("%s/api/v2/private/assets/getWithdrawSignInfo", c.c.GetBaseURL())

	queryParams := map[string]string{}
	if params.ChainId != "" {
		queryParams["chainId"] = params.ChainId
	}
	if params.TokenAddress != "" {
		queryParams["tokenAddress"] = params.TokenAddress
	}
	if params.Amount != "" {
		queryParams["amount"] = params.Amount
	}

	resp, err := c.c.HttpRequest(url, "GET", nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get withdraw sign info: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result ResultGetWithdrawSignInfo
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s, msg: %s", result.Code, result.ErrorMsg)
	}

	return &result, nil
}

// CreateCrossWithdraw creates a cross-chain withdrawal order
func (c *Client) CreateCrossWithdraw(ctx context.Context, params CreateCrossWithdrawParams) (*ResultCreateCrossWithdraw, error) {
	url := fmt.Sprintf("%s/api/v2/private/assets/createCrossWithdraw", c.c.GetBaseURL())

	// Parse nonce to integer
	nonceInt, err := strconv.ParseInt(params.Nonce, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse nonce: %w", err)
	}

	body := map[string]interface{}{
		"accountId":        strconv.FormatInt(c.c.GetAccountID(), 10),
		"coinId":           params.CoinId,
		"amount":           params.Amount,
		"ethAddress":       params.EthAddress,    // Signer address (Privy address)
		"targetAddress":    params.TargetAddress, // Final cross-chain withdrawal target address
		"clientWithdrawId": params.ClientWithdrawId,
		"fee":              params.Fee,
		"signature":        params.Signature, // EIP-712 signature with 0x prefix
		"signer":           params.Signer,    // Signer address
		"nonce":            nonceInt,         // Integer type for cross-chain withdraw
		"l2ExpireTime":     params.L2ExpireTime,
		"chainId":          params.ChainId, // Target chain ID
	}

	resp, err := c.c.HttpRequest(url, "POST", body, nil)
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

// CreateCrossWithdrawAuto creates a cross withdraw with auto fee lookup and signature generation.
func (c *Client) CreateCrossWithdrawAuto(ctx context.Context, params *CreateCrossWithdrawAutoParams, md *metadata.MetaData) (*ResultCreateCrossWithdraw, error) {
	if params == nil {
		return nil, fmt.Errorf("params is nil")
	}
	if md == nil {
		return nil, fmt.Errorf("metadata is nil")
	}

	prepared, err := c.PrepareWithdrawSignInfo(ctx, md, PrepareWithdrawSignInfoParams{
		CoinId:       params.CoinId,
		ChainId:      params.ChainId,
		TokenAddress: params.TokenAddress,
		Amount:       params.Amount,
	})
	if err != nil {
		return nil, err
	}

	signerAddress, err := c.c.ResolveWalletAddress()
	if err != nil {
		return nil, fmt.Errorf("failed to resolve signer address: %w", err)
	}
	if !common.IsHexAddress(signerAddress) {
		return nil, fmt.Errorf("invalid signer address: %s", signerAddress)
	}
	signerAddress = common.HexToAddress(signerAddress).Hex()

	targetAddress := strings.TrimSpace(params.TargetAddress)
	if targetAddress == "" {
		targetAddress = signerAddress
	}
	if !common.IsHexAddress(targetAddress) {
		return nil, fmt.Errorf("invalid target address: %s", targetAddress)
	}
	targetAddress = common.HexToAddress(targetAddress).Hex()

	clientWithdrawID := strings.TrimSpace(params.ClientWithdrawId)
	if clientWithdrawID == "" {
		clientWithdrawID = internal.GetRandomClientId()
	}

	fee := strings.TrimSpace(params.Fee)
	if fee == "" {
		fee = prepared.Fee
	}

	signatureInfo, err := c.BuildWithdrawV2Signature(md, BuildWithdrawV2SignatureParams{
		CoinId:           prepared.CoinId,
		Amount:           prepared.Amount,
		Fee:              fee,
		ClientWithdrawId: clientWithdrawID,
		ToAddress:        signerAddress,
	})
	if err != nil {
		return nil, err
	}

	return c.CreateCrossWithdraw(ctx, CreateCrossWithdrawParams{
		CoinId:           prepared.CoinId,
		Amount:           prepared.Amount,
		EthAddress:       signerAddress,
		TargetAddress:    targetAddress,
		ClientWithdrawId: clientWithdrawID,
		Fee:              fee,
		Signature:        signatureInfo.Signature,
		Signer:           signatureInfo.Signer,
		Nonce:            signatureInfo.Nonce,
		L2ExpireTime:     signatureInfo.L2ExpireTime,
		ChainId:          prepared.ChainId,
	})
}
