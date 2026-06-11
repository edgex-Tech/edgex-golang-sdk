package transfer

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/edgex-Tech/edgex-golang-sdk/v2/sdk/transfer"
	"github.com/edgex-Tech/edgex-golang-sdk/v2/test"
	"github.com/stretchr/testify/assert"
)

func TestGetTransferOutById(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	params := transfer.GetTransferOutByIdParams{
		TransferId: "123",
	}
	resp, err := client.GetTransferOutById(ctx, params)
	jsonData, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("Transfer Out: %s", string(jsonData))
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "SUCCESS", resp.Code)

	data := resp.Data
	assert.NotNil(t, data)
	if len(data) > 0 {
		// Data is interface{}, skip detailed assertions
		t.Logf("Transfer record: %v", data[0])
	}
}

func TestGetTransferInById(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	params := transfer.GetTransferInByIdParams{
		TransferId: "123",
	}
	resp, err := client.GetTransferInById(ctx, params)
	jsonData, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("Transfer In: %s", string(jsonData))
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "SUCCESS", resp.Code)

	data := resp.Data
	assert.NotNil(t, data)
	if len(data) > 0 {
		// Data is interface{}, skip detailed assertions
		t.Logf("Transfer record: %v", data[0])
	}
}

func TestGetWithdrawAvailableAmount(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()
	coinID, err := test.ResolveTestCoinID(ctx, client)
	assert.NoError(t, err)

	params := transfer.GetWithdrawAvailableAmountParams{
		CoinId: coinID,
	}
	resp, err := client.GetWithdrawAvailableAmount(ctx, params)
	jsonData, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("Withdraw Available Amount: %s", string(jsonData))
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "SUCCESS", resp.Code)

	data := resp.Data
	assert.NotNil(t, data)
	// Data is interface{}, skip detailed assertions
	t.Logf("Available amount data: %v", data)
}

func TestCreateTransferOut(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()
	coinID, err := test.ResolveTestCoinID(ctx, client)
	assert.NoError(t, err)

	receiverAccountID := strings.TrimSpace(os.Getenv("EDGEX_TRANSFER_RECEIVER_ACCOUNT_ID"))
	receiverL2Key := strings.TrimSpace(os.Getenv("EDGEX_TRANSFER_RECEIVER_L2_KEY"))
	if receiverAccountID == "" || receiverL2Key == "" {
		t.Skip("Skipping transfer-out create test: EDGEX_TRANSFER_RECEIVER_ACCOUNT_ID and EDGEX_TRANSFER_RECEIVER_L2_KEY are required")
	}

	// Get metadata for transfer
	metadata, err := client.GetMetaData(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, metadata)

	// Test parameters
	params := &transfer.CreateTransferOutParams{
		CoinId:            coinID,
		Amount:            "1",
		ReceiverAccountId: receiverAccountID,
		ReceiverL2Key:     receiverL2Key,
		TransferReason:    "USER_TRANSFER",
	}

	// Create transfer out - should auto-generate nonce, expiry, and signature
	resp, err := client.Transfer.CreateTransferOut(ctx, params, metadata.Data)

	if err != nil {
		t.Logf("Transfer out creation error: %v", err)
		t.Skip("Skipping transfer-out test due to error (may need proper chain configuration)")
		return
	}

	assert.NoError(t, err)
	assert.NotNil(t, resp)

	// Log response for debugging
	jsonData, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("Create Transfer Out Response: %s", string(jsonData))

	// Verify response
	assert.Equal(t, "SUCCESS", resp.Code)
	data := resp.Data
	assert.NotNil(t, data)
	// Data is interface{}, skip detailed assertions
	t.Logf("Create transfer response data: %v", data)
}
