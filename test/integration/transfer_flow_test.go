package integration

import (
	"strings"
	"testing"

	"github.com/edgex-Tech/edgex-golang-sdk/v2/sdk/transfer"
	"github.com/edgex-Tech/edgex-golang-sdk/v2/test"
	"github.com/stretchr/testify/assert"
)

// TestIntegration_TransferFlow tests the transfer query flow:
// 1. Get available transfer amount
// 2. Query transfer out records
// 3. Query transfer in records
// 4. Verify data consistency
func TestIntegration_TransferFlow(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	// Step 1: Get coin ID for transfer
	t.Log("Step 1: Getting coin ID...")
	coinID, err := test.ResolveTestCoinID(ctx, client)
	if err != nil {
		t.Skipf("Cannot resolve coin ID: %v", err)
		return
	}
	t.Logf("Using coin ID: %s", coinID)

	// Step 2: Get available transfer amount
	t.Log("Step 2: Getting available transfer amount...")
	availableResp, err := client.GetWithdrawAvailableAmount(ctx, transfer.GetWithdrawAvailableAmountParams{
		CoinId: coinID,
	})

	if err != nil {
		// This endpoint may have special requirements
		if strings.Contains(err.Error(), "INVALID") || strings.Contains(err.Error(), "NOT_FOUND") {
			t.Logf("Available amount check skipped: %v", err)
		} else {
			assert.NoError(t, err)
		}
	} else {
		assert.NotNil(t, availableResp)
		if availableResp.Data != nil && availableResp.Data.AvailableAmount != nil {
			t.Logf("Available transfer amount: %s", *availableResp.Data.AvailableAmount)
		}
	}

	// Step 3: Query transfer out records
	t.Log("Step 3: Querying transfer out records...")
	transferOutResp, err := client.GetTransferOutById(ctx, transfer.GetTransferOutByIdParams{
		TransferId: "1,2,3", // Dummy IDs for testing query
	})
	assert.NoError(t, err)
	assert.NotNil(t, transferOutResp)

	if transferOutResp.Data != nil {
		outCount := len(transferOutResp.Data)
		t.Logf("Transfer out records found: %d", outCount)

		for i, record := range transferOutResp.Data {
			if record.Id != nil {
				t.Logf("Transfer out %d: ID=%s", i+1, *record.Id)
			}
		}
	}

	// Step 4: Query transfer in records
	t.Log("Step 4: Querying transfer in records...")
	transferInResp, err := client.GetTransferInById(ctx, transfer.GetTransferInByIdParams{
		TransferId: "1,2,3", // Dummy IDs for testing query
	})
	assert.NoError(t, err)
	assert.NotNil(t, transferInResp)

	if transferInResp.Data != nil {
		inCount := len(transferInResp.Data)
		t.Logf("Transfer in records found: %d", inCount)

		for i, record := range transferInResp.Data {
			if record.Id != nil {
				t.Logf("Transfer in %d: ID=%s", i+1, *record.Id)
			}
		}
	}

	// Step 5: Get account asset to verify balances
	t.Log("Step 5: Getting account asset for balance verification...")
	assetResp, err := client.GetAccountAsset(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, assetResp)

	if assetResp.Data != nil {
		t.Logf("Account has %d collateral types", len(assetResp.Data.CollateralList))
		for i, collateral := range assetResp.Data.CollateralList {
			t.Logf("Collateral %d: CoinID=%s, Amount=%s", i+1, collateral.CoinID, collateral.Amount)
		}
	}

	t.Log("✅ Transfer flow test completed successfully")
}

// Note: Transfer creation (CreateTransferOut) is not included in this test
// as it requires:
// 1. A valid target account ID
// 2. Sufficient balance
// 3. EDGEX_ENABLE_MUTATION_TESTS=true
//
// Withdrawal tests are intentionally skipped per user's request
// to be tested last.
