package transfer

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"strconv"
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
	if metadata.Global == nil || metadata.Global.StarkExCollateralCoin == nil {
		return nil, fmt.Errorf("metadata global is nil")
	}
	coin := metadata.Global.StarkExCollateralCoin
	assetID, err := internal.HexToBigInteger(coin.StarkExAssetId)
	if err != nil {
		return nil, fmt.Errorf("failed to parse asset ID: %w", err)
	}
	fmt.Println(coin.CoinId)

	// Generate client transfer ID if not provided
	clientTransferId := internal.GetRandomClientId()

	// Calculate nonce and expiration time
	nonce := internal.CalcNonce(clientTransferId)
	l2ExpireTime := params.ExpireTime.Add(14 * 24 * time.Hour).UnixMilli()
	l2ExpireHour := l2ExpireTime / (60 * 60 * 1000)

	// Convert receiver L2 key to big.Int
	receiverPublicKey, err := internal.HexToBigInteger(params.ReceiverL2Key)
	if err != nil {
		return nil, fmt.Errorf("invalid receiver L2 key format: %s", params.ReceiverL2Key)
	}

	// Parse receiver account ID
	receiverPositionId, err := strconv.ParseInt(params.ReceiverAccountId, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid receiver account ID: %w", err)
	}

	// Convert amount to protocol format (shift by 6 decimal places)
	amountDm, err := decimal.NewFromString(params.Amount)
	if err != nil {
		return nil, fmt.Errorf("failed to parse amount: %w", err)
	}
	amount := amountDm.Shift(6).IntPart()
	maxAmountFee := int64(0)

	// Calculate transfer hash and sign it
	msgHash := internal.CalcTransferHash(
		assetID,
		big.NewInt(0),
		receiverPublicKey,
		c.Client.GetAccountID(),
		receiverPositionId,
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

	// Build request body
	body := map[string]interface{}{
		"accountId":         strconv.FormatInt(c.Client.GetAccountID(), 10),
		"coinId":            params.CoinId,
		"amount":            params.Amount,
		"receiverAccountId": params.ReceiverAccountId,
		"receiverL2Key":     params.ReceiverL2Key,
		"clientTransferId":  clientTransferId,
		"transferReason":    params.TransferReason,
		"l2Nonce":           strconv.FormatInt(nonce, 10),
		"l2ExpireTime":      strconv.FormatInt(l2ExpireTime, 10),
		"l2Signature":       fmt.Sprintf("%s%s", signature.R, signature.S),
	}
	if params.ExtraType != nil {
		body["extraType"] = *params.ExtraType
	}
	if params.ExtraDataJson != nil {
		body["extraDataJson"] = *params.ExtraDataJson
	}

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
	fmt.Println(string(respBody))

	var result ResultCreateTransferOut
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}
