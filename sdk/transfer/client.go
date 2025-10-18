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
	url := fmt.Sprintf("%s/api/v1/private/transfer/getWithdrawAvailableAmount", c.Client.GetBaseURL())
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
func (c *Client) CreateTransferOut(ctx context.Context, params CreateTransferOutParams, metadata interface{}) (*ResultCreateTransferOut, error) {
	url := fmt.Sprintf("%s/api/v1/private/transfer/createTransferOut", c.Client.GetBaseURL())

	// Generate client transfer ID if not provided
	if params.ClientTransferId == "" {
		params.ClientTransferId = internal.GenerateUUID()
	}

	l2ExpireTime := strconv.FormatInt(time.Now().Add(14*24*time.Hour).UnixMilli(), 10)

	// Convert parameters to appropriate types for hash calculation
	amountDm, _ := decimal.NewFromString(params.Amount)
	amount := amountDm.Shift(6).IntPart()
	nonce := internal.CalcNonce(params.ClientTransferId)
	expireTime, _ := strconv.ParseInt(l2ExpireTime, 10, 64)
	expireTimeUnix := expireTime / (60 * 60 * 1000) // Convert to hours

	// Remove 0x prefix from receiver L2 key if present
	receiverL2Key := strings.TrimPrefix(params.ReceiverL2Key, "0x")

	// Convert receiver L2 key to big.Int
	receiverPublicKey, ok := new(big.Int).SetString(receiverL2Key, 16)
	if !ok {
		return nil, fmt.Errorf("invalid receiver L2 key format: %s", receiverL2Key)
	}

	// Get position IDs (same as account IDs)
	senderPositionId := c.Client.GetAccountID()
	receiverPositionId, err := strconv.ParseInt(params.ReceiverAccountId, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid receiver account ID: %w", err)
	}
	feePositionId := senderPositionId // Fee position is same as sender for now
	maxAmountFee := int64(0)

	// For now, use a placeholder asset ID (0)
	// In production, this should come from metadata
	assetID := big.NewInt(0)

	// Calculate transfer hash and sign it
	msgHash := internal.CalcTransferHash(
		assetID,
		big.NewInt(0),
		receiverPublicKey,
		senderPositionId,
		receiverPositionId,
		feePositionId,
		nonce,
		amount,
		maxAmountFee,
		expireTimeUnix,
	)
	signature, err := c.Client.Sign(msgHash)
	if err != nil {
		return nil, fmt.Errorf("failed to sign transfer hash: %w", err)
	}

	// Build request body
	body := map[string]interface{}{
		"accountId":        strconv.FormatInt(c.Client.GetAccountID(), 10),
		"coinId":           params.CoinId,
		"amount":           amountDm.String(),
		"receiverAccountId": params.ReceiverAccountId,
		"receiverL2Key":    params.ReceiverL2Key,
		"clientTransferId": params.ClientTransferId,
		"transferReason":   params.TransferReason,
		"l2Nonce":          strconv.FormatInt(nonce, 10),
		"l2ExpireTime":     l2ExpireTime,
		"l2Signature":      fmt.Sprintf("%s%s%s", signature.R, signature.S, signature.V),
	}

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

