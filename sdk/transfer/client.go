package transfer

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/edgex-Tech/edgex-golang-sdk/sdk/internal"
	metadatapkg "github.com/edgex-Tech/edgex-golang-sdk/sdk/metadata"
	"github.com/shopspring/decimal"
)

// Client represents the new transfer client without OpenAPI dependencies
type Client struct {
	*internal.Client
}

// NewClient creates a new transfer client
func NewClient(client *internal.Client) *Client {
	return &Client{
		Client: client,
	}
}

// GetTransferOutById gets a transfer out record by ID
func (c *Client) GetTransferOutById(ctx context.Context, params GetTransferOutByIdParams) (*ResultListTransferOut, error) {
	url := fmt.Sprintf("%s/api/v1/private/transfer/getTransferOutById", c.Client.GetBaseURL())
	queryParams := map[string]string{
		"accountId": strconv.FormatInt(c.Client.GetAccountID(), 10),
	}

	if params.TransferId != "" {
		queryParams["transferOutIdList"] = params.TransferId
	}

	resp, err := c.Client.HttpRequest(url, "GET", nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get transfer out by id: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result ResultListTransferOut
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}

// GetTransferInById gets a transfer in record by ID
func (c *Client) GetTransferInById(ctx context.Context, params GetTransferInByIdParams) (*ResultListTransferIn, error) {
	url := fmt.Sprintf("%s/api/v1/private/transfer/getTransferInById", c.Client.GetBaseURL())
	queryParams := map[string]string{
		"accountId": strconv.FormatInt(c.Client.GetAccountID(), 10),
	}

	if params.TransferId != "" {
		queryParams["transferInIdList"] = params.TransferId
	}

	resp, err := c.Client.HttpRequest(url, "GET", nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get transfer in by id: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result ResultListTransferIn
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}

// GetWithdrawAvailableAmount gets the available withdrawal amount
func (c *Client) GetWithdrawAvailableAmount(ctx context.Context, params GetWithdrawAvailableAmountParams) (*ResultGetTransferOutAvailableAmount, error) {
	url := fmt.Sprintf("%s/api/v1/private/transfer/getTransferOutAvailableAmount", c.Client.GetBaseURL())
	queryParams := map[string]string{
		"accountId": strconv.FormatInt(c.Client.GetAccountID(), 10),
	}

	if params.CoinId != "" {
		queryParams["coinId"] = params.CoinId
	}

	resp, err := c.Client.HttpRequest(url, "GET", nil, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get available withdrawal amount: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result ResultGetTransferOutAvailableAmount
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}

// CreateTransferOut creates a new transfer out order
func (c *Client) CreateTransferOut(ctx context.Context, params *CreateTransferOutParams, metadata *metadatapkg.MetaData) (*ResultCreateTransferOut, error) {
	if c.Client.GetAPIVersion() == internal.APIVersionV2 {
		return c.createTransferOutV2(ctx, params, metadata)
	}
	return c.createTransferOutV1(ctx, params, metadata)
}

func (c *Client) createTransferOutV1(ctx context.Context, params *CreateTransferOutParams, metadata *metadatapkg.MetaData) (*ResultCreateTransferOut, error) {
	if metadata == nil || metadata.Global == nil || metadata.Global.StarkExCollateralCoin == nil {
		return nil, fmt.Errorf("metadata global.starkExCollateralCoin is nil")
	}

	coin := metadata.Global.StarkExCollateralCoin
	assetID, err := internal.HexToBigInteger(coin.StarkExAssetId)
	if err != nil {
		return nil, fmt.Errorf("failed to parse asset ID: %w", err)
	}

	clientTransferID := strings.TrimSpace(params.ClientTransferId)
	if clientTransferID == "" {
		clientTransferID = internal.GetRandomClientId()
	}

	nonce := internal.CalcNonce(clientTransferID)

	expireTime := params.ExpireTime
	if expireTime.IsZero() {
		expireTime = time.Now()
	}
	l2ExpireTime := expireTime.Add(14 * 24 * time.Hour).UnixMilli()
	l2ExpireHour := l2ExpireTime / (60 * 60 * 1000)

	receiverPublicKey, err := internal.HexToBigInteger(params.ReceiverL2Key)
	if err != nil {
		return nil, fmt.Errorf("invalid receiver L2 key format: %s", params.ReceiverL2Key)
	}

	receiverPositionID, err := strconv.ParseInt(params.ReceiverAccountId, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid receiver account ID: %w", err)
	}

	amountDm, err := decimal.NewFromString(params.Amount)
	if err != nil {
		return nil, fmt.Errorf("failed to parse amount: %w", err)
	}
	amount := amountDm.Shift(6).IntPart()
	maxAmountFee := int64(0)

	msgHash := internal.CalcTransferHash(
		assetID,
		big.NewInt(0),
		receiverPublicKey,
		c.Client.GetAccountID(),
		receiverPositionID,
		c.Client.GetAccountID(),
		nonce,
		amount,
		maxAmountFee,
		l2ExpireHour,
	)
	signature, err := c.Client.Sign(msgHash)
	if err != nil {
		return nil, fmt.Errorf("failed to sign transfer hash: %w", err)
	}

	body := map[string]interface{}{
		"accountId":         strconv.FormatInt(c.Client.GetAccountID(), 10),
		"coinId":            params.CoinId,
		"amount":            params.Amount,
		"receiverAccountId": params.ReceiverAccountId,
		"receiverL2Key":     params.ReceiverL2Key,
		"clientTransferId":  clientTransferID,
		"transferReason": func() string {
			if strings.TrimSpace(params.TransferReason) != "" {
				return params.TransferReason
			}
			return USER_TRANSFER.String()
		}(),
		"l2Nonce":      strconv.FormatInt(nonce, 10),
		"l2ExpireTime": strconv.FormatInt(l2ExpireTime, 10),
		"l2Signature":  fmt.Sprintf("%s%s", signature.R, signature.S),
	}
	if params.ExtraType != nil {
		body["extraType"] = *params.ExtraType
	}
	if params.ExtraDataJson != nil {
		body["extraDataJson"] = *params.ExtraDataJson
	}

	return c.postCreateTransferOut(ctx, body)
}

func (c *Client) createTransferOutV2(ctx context.Context, params *CreateTransferOutParams, metadata *metadatapkg.MetaData) (*ResultCreateTransferOut, error) {
	if metadata == nil || metadata.Global == nil {
		return nil, fmt.Errorf("metadata.global is nil")
	}

	clientTransferID := strings.TrimSpace(params.ClientTransferId)
	if clientTransferID == "" {
		clientTransferID = internal.GetRandomClientId()
	}

	l2Nonce := strings.TrimSpace(params.L2Nonce)
	if l2Nonce == "" {
		l2Nonce = strconv.FormatInt(internal.CalcNonce(clientTransferID), 10)
	}

	l2ExpireTime := strings.TrimSpace(params.L2ExpireTime)
	if l2ExpireTime == "" {
		l2ExpireTime = strconv.FormatInt(time.Now().Add(10*24*time.Hour).UnixMilli(), 10)
	}
	l2ExpireTimeMillis, err := strconv.ParseInt(l2ExpireTime, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid l2ExpireTime: %w", err)
	}
	deadline := strconv.FormatInt(l2ExpireTimeMillis/1000, 10)

	stepSize, err := resolveTransferCoinStepSize(metadata, params.CoinId)
	if err != nil {
		return nil, err
	}
	amountDm, err := decimal.NewFromString(params.Amount)
	if err != nil {
		return nil, fmt.Errorf("failed to parse amount: %w", err)
	}
	stepDm, err := decimal.NewFromString(stepSize)
	if err != nil {
		return nil, fmt.Errorf("failed to parse step size: %w", err)
	}
	if stepDm.IsZero() {
		return nil, fmt.Errorf("coin step size cannot be zero")
	}
	amountInt := amountDm.Div(stepDm).Floor().BigInt().String()

	chainID := strings.TrimSpace(metadata.Global.ChainId)
	if chainID == "" {
		chainID = strings.TrimSpace(metadata.Global.NativeChainId)
	}
	if chainID == "" {
		return nil, fmt.Errorf("metadata.global.chainId/nativeChainId is required for v2 transfer signing")
	}
	verifyingContract := strings.TrimSpace(metadata.Global.ContractAddress)
	if verifyingContract == "" {
		return nil, fmt.Errorf("metadata.global.contractAddress is required for v2 transfer signing")
	}

	l2Signature := strings.TrimSpace(params.L2Signature)
	if l2Signature == "" {
		if c.Client.GetWalletPriKey() == "" {
			return nil, fmt.Errorf("wallet private key is required for v2 EIP-712 transfer signing")
		}
		signer := strings.TrimSpace(params.Signer)
		if signer == "" {
			signer, err = c.Client.ResolveWalletSignerAddress()
			if err != nil {
				return nil, fmt.Errorf("failed to resolve wallet signer address: %w", err)
			}
		}

		domain, err := internal.NewTypedDataDomain("EdgeX", "1", chainID, verifyingContract)
		if err != nil {
			return nil, fmt.Errorf("failed to build typed data domain: %w", err)
		}

		typedData := internal.TypedData{
			Types: internal.TypedDataTypes{
				"EIP712Domain": []internal.TypedDataType{
					{Name: "name", Type: "string"},
					{Name: "version", Type: "string"},
					{Name: "chainId", Type: "uint256"},
					{Name: "verifyingContract", Type: "address"},
				},
				"TransferParams": []internal.TypedDataType{
					{Name: "from", Type: "uint64"},
					{Name: "to", Type: "uint64"},
					{Name: "assetId", Type: "uint64"},
					{Name: "amount", Type: "int256"},
					{Name: "nonce", Type: "uint256"},
					{Name: "deadline", Type: "uint256"},
					{Name: "signer", Type: "address"},
				},
			},
			PrimaryType: "TransferParams",
			Domain:      domain,
			Message: internal.TypedDataMessage{
				"from":     strconv.FormatInt(c.Client.GetAccountID(), 10),
				"to":       params.ReceiverAccountId,
				"assetId":  params.CoinId,
				"amount":   amountInt,
				"nonce":    l2Nonce,
				"deadline": deadline,
				"signer":   signer,
			},
		}

		l2Signature, err = c.Client.SignTypedDataWithWalletKey(typedData)
		if err != nil {
			return nil, fmt.Errorf("failed to sign v2 transfer typed data: %w", err)
		}
	}

	body := map[string]interface{}{
		"accountId":         strconv.FormatInt(c.Client.GetAccountID(), 10),
		"coinId":            params.CoinId,
		"amount":            params.Amount,
		"receiverAccountId": params.ReceiverAccountId,
		"clientTransferId":  clientTransferID,
		"transferReason": func() string {
			if strings.TrimSpace(params.TransferReason) != "" {
				return params.TransferReason
			}
			return USER_TRANSFER.String()
		}(),
		"l2Nonce":      l2Nonce,
		"l2ExpireTime": l2ExpireTime,
		"l2Signature":  l2Signature,
	}
	if params.ExtraType != nil {
		body["extraType"] = *params.ExtraType
	}
	if params.ExtraDataJson != nil {
		body["extraDataJson"] = *params.ExtraDataJson
	}

	return c.postCreateTransferOut(ctx, body)
}

func resolveTransferCoinStepSize(metadata *metadatapkg.MetaData, coinID string) (string, error) {
	coinID = strings.TrimSpace(coinID)
	if metadata != nil {
		for i := range metadata.CoinList {
			if metadata.CoinList[i].CoinId == coinID {
				if strings.TrimSpace(metadata.CoinList[i].StepSize) == "" {
					return "", fmt.Errorf("step size not found for coin: %s", coinID)
				}
				return metadata.CoinList[i].StepSize, nil
			}
		}
	}
	return "", fmt.Errorf("coin not found in metadata: %s", coinID)
}

func (c *Client) postCreateTransferOut(ctx context.Context, body map[string]interface{}) (*ResultCreateTransferOut, error) {
	_ = ctx
	url := fmt.Sprintf("%s/api/v1/private/transfer/createTransferOut", c.Client.GetBaseURL())
	resp, err := c.Client.HttpRequest(url, "POST", body, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create transfer out: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result ResultCreateTransferOut
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}
