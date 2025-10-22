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
func (c *Client) CreateTransferOut(ctx context.Context, params *CreateTransferOutParams, metadata *metadatapkg.MetaData) (*ResultCreateTransferOut, error) {
	// Find coin from metadata
	var coin *metadatapkg.Coin
	if metadata != nil && metadata.CoinList != nil {
		for i := range metadata.CoinList {
			if metadata.CoinList[i].CoinId == params.CoinId {
				coin = &metadata.CoinList[i]
				break
			}
		}
	}
	if coin == nil {
		return nil, fmt.Errorf("coin not found: %s", params.CoinId)
	}
	assetID, err := internal.HexToBigInteger(coin.StarkExAssetId)
	if err != nil {
		return nil, fmt.Errorf("failed to parse asset ID: %w", err)
	}

	// Get collateral coin from metadata
	var collateralCoin *metadatapkg.Coin
	if metadata != nil && metadata.Global != nil {
		collateralCoin = metadata.Global.StarkExCollateralCoin
	}

	if collateralCoin == nil {
		return nil, fmt.Errorf("collateral coin not found in metadata")
	}
	assetIDFee, err := internal.HexToBigInteger(collateralCoin.StarkExAssetId)
	if err != nil {
		return nil, fmt.Errorf("failed to parse asset ID fee: %w", err)
	}

	// Generate client transfer ID if not provided
	clientTransferId := internal.GenerateUUID()

	// Parse decimal amount
	amountDm, err := decimal.NewFromString(params.Amount)
	if err != nil {
		return nil, fmt.Errorf("failed to parse amount: %w", err)
	}

	// Calculate nonce and expiration time
	nonce := internal.CalcNonce(clientTransferId)
	l2ExpireTime := time.Now().Add(14 * 24 * time.Hour).UnixMilli()
	l2ExpireHour := l2ExpireTime / (60 * 60 * 1000)

	// Remove 0x prefix from receiver L2 key if present
	receiverL2Key := strings.TrimPrefix(params.ReceiverL2Key, "0x")

	// Convert receiver L2 key to big.Int
	receiverPublicKey, ok := new(big.Int).SetString(receiverL2Key, 16)
	if !ok {
		return nil, fmt.Errorf("invalid receiver L2 key format: %s", receiverL2Key)
	}

	// Parse receiver account ID
	receiverPositionId, err := strconv.ParseInt(params.ReceiverAccountId, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid receiver account ID: %w", err)
	}

	// Get position IDs
	senderPositionId := c.Client.GetAccountID()
	feePositionId := senderPositionId // Fee position is same as sender

	// limitFee := amountDm.Mul(price).Mul(feeRate).Ceil()
	// maxAmountFee := limitFee.Mul(usdtStarkExResolutionDecimal)
	maxAmountFee := int64(0)

	// Convert amount to protocol format (shift by 6 decimal places)
	shiftFactor := decimal.NewFromInt(1000000)
	amount := amountDm.Mul(shiftFactor).IntPart()

	// Calculate transfer hash and sign it
	msgHash := internal.CalcTransferHash(
		assetID,
		assetIDFee,
		receiverPublicKey,
		senderPositionId,
		receiverPositionId,
		feePositionId,
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
		"amount":            amountDm.String(),
		"receiverAccountId": params.ReceiverAccountId,
		"receiverL2Key":     params.ReceiverL2Key,
		"clientTransferId":  clientTransferId,
		"transferReason":    params.TransferReason,
		"l2Nonce":           strconv.FormatInt(nonce, 10),
		"l2ExpireTime":      strconv.FormatInt(l2ExpireTime, 10),
		"l2Signature":       fmt.Sprintf("%s%s%s", signature.R, signature.S, signature.V),
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
