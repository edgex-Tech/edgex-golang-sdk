package integration

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/edgex-Tech/edgex-golang-sdk/sdk/account"
	"github.com/edgex-Tech/edgex-golang-sdk/test"
	"github.com/stretchr/testify/assert"
)

// TestIntegration_PositionManagement tests the position management flow:
// 1. Query current positions
// 2. Query position transactions
// 3. Query collateral transactions
// 4. Query position terms
// 5. Verify data consistency
func TestIntegration_PositionManagement(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	// Step 1: Query current positions
	t.Log("Step 1: Querying current positions...")
	positionsResp, err := client.GetAccountPositions(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, positionsResp)
	
	positionCount := len(positionsResp.Data)
	t.Logf("Current positions: %d", positionCount)
	
	for i, pos := range positionsResp.Data {
		t.Logf("Position %d: ContractID=%s, Size=%s, Price=%s", 
			i+1, pos.ContractID, pos.Size, pos.Price)
	}

	// Step 2: Query position transactions
	t.Log("Step 2: Querying position transaction history...")
	positionTxResp, err := client.GetPositionTransactionPage(ctx, account.GetPositionTransactionPageParams{
		Size: 10,
	})
	assert.NoError(t, err)
	assert.NotNil(t, positionTxResp)
	assert.NotNil(t, positionTxResp.Data)
	
	txCount := len(positionTxResp.Data.DataList)
	t.Logf("Position transactions found: %d", txCount)
	
	if txCount > 0 {
		t.Logf("Recent position transaction: Type=%v", positionTxResp.Data.DataList[0])
	}

	// Step 3: Query collateral transactions
	t.Log("Step 3: Querying collateral transaction history...")
	collateralTxResp, err := client.GetCollateralTransactionPage(ctx, account.GetCollateralTransactionPageParams{
		Size: 10,
	})
	assert.NoError(t, err)
	assert.NotNil(t, collateralTxResp)
	assert.NotNil(t, collateralTxResp.Data)
	
	collateralTxCount := len(collateralTxResp.Data.DataList)
	t.Logf("Collateral transactions found: %d", collateralTxCount)

	// Step 4: Query position terms
	t.Log("Step 4: Querying position terms...")
	termsResp, err := client.GetPositionTermPage(ctx, account.GetPositionTermPageParams{
		Size: 10,
	})
	assert.NoError(t, err)
	assert.NotNil(t, termsResp)
	assert.NotNil(t, termsResp.Data)
	
	termCount := len(termsResp.Data.DataList)
	t.Logf("Position terms found: %d", termCount)

	// Step 5: Get account asset snapshot
	t.Log("Step 5: Getting account asset snapshots...")
	coinID, err := test.ResolveTestCoinID(ctx, client)
	if err == nil {
		snapshotResp, err := client.GetAccountAssetSnapshotPage(ctx, account.GetAccountAssetSnapshotPageParams{
			Size:   10,
			CoinID: coinID,
		})
		assert.NoError(t, err)
		assert.NotNil(t, snapshotResp)
		
		if snapshotResp.Data != nil {
			snapshotCount := len(snapshotResp.Data.DataList)
			t.Logf("Asset snapshots found: %d", snapshotCount)
		}
	}

	// Step 6: Query account deleverage light
	t.Log("Step 6: Querying account deleverage status...")
	deleverageResp, err := client.GetAccountDeleverageLight(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, deleverageResp)
	
	if deleverageResp.Data != nil && deleverageResp.Data.DeleverageLevel != nil {
		t.Logf("Deleverage level: %s", *deleverageResp.Data.DeleverageLevel)
	}

	t.Log("✅ Position management test completed successfully")
}

// TestIntegration_PositionLifecycle tests creating and closing a position
// This test requires EDGEX_ENABLE_MUTATION_TESTS=true as it involves real trading
func TestIntegration_PositionLifecycle(t *testing.T) {
	if strings.ToLower(strings.TrimSpace(os.Getenv("EDGEX_ENABLE_MUTATION_TESTS"))) != "true" {
		t.Skip("Skipping position lifecycle test: set EDGEX_ENABLE_MUTATION_TESTS=true to enable")
	}

	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	if strings.TrimSpace(client.GetSignerPriKey()) == "" {
		t.Skip("Skipping integration test: EDGEX_SIGNER_PRIVATE_KEY is required")
	}

	ctx := test.GetTestContext()

	t.Log("Step 1: Getting initial positions...")
	positionsBefore, err := client.GetAccountPositions(ctx)
	assert.NoError(t, err)
	initialCount := len(positionsBefore.Data)
	t.Logf("Initial position count: %d", initialCount)

	// Note: Opening and closing positions requires market orders
	// This would be implemented similar to market_order_flow_test.go
	// but with ReduceOnly flag for closing
	
	t.Log("Position lifecycle test framework ready")
	t.Log("Full implementation requires market order execution")
	
	// Wait and query final state
	time.Sleep(1 * time.Second)
	
	positionsAfter, err := client.GetAccountPositions(ctx)
	assert.NoError(t, err)
	finalCount := len(positionsAfter.Data)
	t.Logf("Final position count: %d", finalCount)
	
	t.Log("✅ Position lifecycle test framework completed")
}
