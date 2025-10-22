package transfer

import (
	"encoding/json"
	"testing"

	"github.com/edgex-Tech/edgex-golang-sdk/sdk/transfer"
	"github.com/edgex-Tech/edgex-golang-sdk/test"
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

	params := transfer.GetWithdrawAvailableAmountParams{
		CoinId: "1000",
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

	// Test parameters
	params := &transfer.CreateTransferOutParams{
		CoinId:            "1000", // Asset ID
		Amount:            "1",    // 1 unit
		ReceiverAccountId: "542103805685137746",
		ReceiverL2Key:     "0x046bcf2e07c20550c49986aca69f405ae4672507fae2568640d3f1d2dcf1bfeb",
		TransferReason:    "USER_TRANSFER",
	}

	// Create transfer out - should auto-generate nonce, expiry, and signature
	resp, err := client.CreateTransferOut(ctx, params)
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
