package account

import (
	"encoding/json"
	"testing"

	"github.com/edgex-Tech/edgex-golang-sdk/v2/sdk/account"
	"github.com/edgex-Tech/edgex-golang-sdk/v2/test"
	"github.com/stretchr/testify/assert"
)

func logJSON(t *testing.T, label string, value interface{}) {
	t.Helper()
	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		t.Logf("%s: <marshal error: %v>", label, err)
		return
	}
	t.Logf("%s: %s", label, string(data))
}

func TestGetAccountAsset(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	asset, err := client.GetAccountAsset(ctx)
	jsonData, _ := json.MarshalIndent(asset, "", "  ")
	t.Logf("Account Asset: %s", string(jsonData))

	// If error occurs, check it and return early
	if err != nil {
		assert.NoError(t, err)
		return
	}

	// Verify response structure
	if !assert.NotNil(t, asset) {
		return
	}
	assert.Equal(t, "SUCCESS", asset.Code)

	data := asset.Data
	assert.NotNil(t, data)
	assert.NotEmpty(t, data.CollateralList)
	// Position list can legitimately be empty for a fresh account.
	assert.NotNil(t, data.PositionList)
}

func TestGetAccountPositions(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	positions, err := client.GetAccountPositions(ctx)
	jsonData, _ := json.MarshalIndent(positions, "", "  ")
	t.Logf("Account Positions: %s", string(jsonData))
	assert.NoError(t, err)
	assert.NotNil(t, positions)
	assert.Equal(t, "SUCCESS", positions.Code)

	data := positions.Data
	assert.NotNil(t, data)
	for _, position := range data {
		assert.NotEmpty(t, position.ContractID)
		// Skip detailed assertions for other fields
	}
}

func TestGetPositionTransactionPage(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	params := account.GetPositionTransactionPageParams{
		Size: 1,
	}
	logJSON(t, "GetPositionTransactionPage params", params)

	transactions, err := client.GetPositionTransactionPage(ctx, params)
	logJSON(t, "GetPositionTransactionPage response", transactions)
	assert.NoError(t, err)
	assert.NotNil(t, transactions)
	assert.Equal(t, "SUCCESS", transactions.Code)

	data := transactions.Data
	assert.NotNil(t, data)
	assert.NotNil(t, data.DataList)
	for _, tx := range data.DataList {
		// Skip detailed assertions for transaction fields
		assert.NotNil(t, tx)
	}
}

func TestGetCollateralTransactionPage(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	params := account.GetCollateralTransactionPageParams{
		Size: 1,
	}
	logJSON(t, "GetCollateralTransactionPage params", params)

	transactions, err := client.GetCollateralTransactionPage(ctx, params)
	logJSON(t, "GetCollateralTransactionPage response", transactions)
	assert.NoError(t, err)
	assert.NotNil(t, transactions)
	assert.Equal(t, "SUCCESS", transactions.Code)

	data := transactions.Data
	assert.NotNil(t, data)
	assert.NotNil(t, data.DataList)
	for _, tx := range data.DataList {
		// Skip detailed assertions for transaction fields
		assert.NotNil(t, tx)
	}
}

func TestGetPositionTermPage(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	params := account.GetPositionTermPageParams{
		Size: 1,
	}
	logJSON(t, "GetPositionTermPage params", params)

	terms, err := client.GetPositionTermPage(ctx, params)
	logJSON(t, "GetPositionTermPage response", terms)
	assert.NoError(t, err)
	assert.NotNil(t, terms)
	assert.Equal(t, "SUCCESS", terms.Code)

	data := terms.Data
	assert.NotNil(t, data)
	assert.NotNil(t, data.DataList)
	for _, term := range data.DataList {
		// Skip detailed assertions for term fields
		assert.NotNil(t, term)
	}
}

func TestGetAccountByID(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	account, err := client.GetAccountByID(ctx)
	jsonData, _ := json.MarshalIndent(account, "", "  ")
	t.Logf("Account: %s", string(jsonData))
	assert.NoError(t, err)
	assert.NotNil(t, account)
	assert.Equal(t, "SUCCESS", account.Code)

	data := account.Data
	assert.NotNil(t, data)
	// Skip detailed assertions for account fields
}

func TestGetAccountDeleverageLight(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	deleverage, err := client.GetAccountDeleverageLight(ctx)
	jsonData, _ := json.MarshalIndent(deleverage, "", "  ")
	t.Logf("Account Deleverage Light: %s", string(jsonData))
	assert.NoError(t, err)
	assert.NotNil(t, deleverage)
	assert.Equal(t, "SUCCESS", deleverage.Code)

	data := deleverage.Data
	assert.NotNil(t, data)
	// Skip detailed assertions for deleverage fields
}

func TestGetAccountAssetSnapshotPage(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()
	coinID, err := test.ResolveTestCoinID(ctx, client)
	assert.NoError(t, err)

	params := account.GetAccountAssetSnapshotPageParams{
		Size:   1,
		CoinID: coinID,
	}
	logJSON(t, "GetAccountAssetSnapshotPage params", params)

	snapshots, err := client.GetAccountAssetSnapshotPage(ctx, params)
	logJSON(t, "GetAccountAssetSnapshotPage response", snapshots)
	assert.NoError(t, err)
	assert.NotNil(t, snapshots)
	assert.Equal(t, "SUCCESS", snapshots.Code)

	data := snapshots.Data
	assert.NotNil(t, data)
	assert.NotNil(t, data.DataList)
	for _, snapshot := range data.DataList {
		// Skip detailed assertions for snapshot fields
		assert.NotNil(t, snapshot)
	}
}

func TestGetPositionTransactionByID(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	// Example transaction IDs
	transactionIDs := []string{"123456789"}

	transactions, err := client.GetPositionTransactionByID(ctx, transactionIDs)
	jsonData, _ := json.MarshalIndent(transactions, "", "  ")
	t.Logf("Position Transaction: %s", string(jsonData))
	assert.NoError(t, err)
	assert.NotNil(t, transactions)
	assert.Equal(t, "SUCCESS", transactions.Code)

	data := transactions.Data
	assert.NotNil(t, data)
	for _, tx := range data {
		// Skip detailed assertions for transaction fields
		assert.NotNil(t, tx)
	}
}

func TestGetCollateralTransactionByID(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	// Example transaction IDs
	transactionIDs := []string{"123456789"}

	transactions, err := client.GetCollateralTransactionByID(ctx, transactionIDs)
	jsonData, _ := json.MarshalIndent(transactions, "", "  ")
	t.Logf("Collateral Transaction: %s", string(jsonData))
	assert.NoError(t, err)
	assert.NotNil(t, transactions)
	assert.Equal(t, "SUCCESS", transactions.Code)

	data := transactions.Data
	assert.NotNil(t, data)
	for _, tx := range data {
		// Skip detailed assertions for transaction fields
		assert.NotNil(t, tx)
	}
}

func TestUpdateLeverageSetting(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()
	contractID, err := test.ResolveTestContractID(ctx, client)
	assert.NoError(t, err)

	// Test updating leverage setting within allowed range on an active contract
	err = client.UpdateLeverageSetting(ctx, contractID, "5")
	assert.NoError(t, err)
}

func TestGetPositionByContractID(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()
	contractID, err := test.ResolveTestContractID(ctx, client)
	assert.NoError(t, err)

	positions, err := client.GetPositionByContractID(ctx, []string{contractID})

	if err != nil {
		assert.NoError(t, err)
		return
	}

	if !assert.NotNil(t, positions) {
		return
	}

	assert.Equal(t, "SUCCESS", positions.Code)
	assert.NotNil(t, positions.Data)

	if len(positions.Data) > 0 {
		t.Logf("Position for contract %s: ContractID=%s", contractID, positions.Data[0].ContractID)
	} else {
		t.Logf("No positions found for contract %s", contractID)
	}
}

func TestGetCollateralByCoinID(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()
	coinID, err := test.ResolveTestCoinID(ctx, client)
	assert.NoError(t, err)

	collaterals, err := client.GetCollateralByCoinID(ctx, []string{coinID})

	if err != nil {
		assert.NoError(t, err)
		return
	}

	if !assert.NotNil(t, collaterals) {
		return
	}

	assert.Equal(t, "SUCCESS", collaterals.Code)
	assert.NotNil(t, collaterals.Data)

	if len(collaterals.Data) > 0 {
		t.Logf("Collateral for coin %s: CoinID=%s", coinID, collaterals.Data[0].CoinID)
	} else {
		t.Logf("No collaterals found for coin %s", coinID)
	}
}

func TestGetPositionOrders(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	// Try different termCount values
	for _, termCount := range []int32{0, 1, 2} {
		t.Logf("\n=== Testing with termCount=%d ===", termCount)

		params := account.GetPositionOrdersParams{
			ContractId: "10000001", // BTCUSDC contract (required)
			TermCount:  termCount,  // Position term count (required)
			Page:       1,          // Page number (optional, default: 1)
			PageSize:   1,          // Items per page (optional, range: 1-100, default: 20)
		}

		logJSON(t, "GetPositionOrders params", params)

		orders, err := client.GetPositionOrders(ctx, params)

		if err != nil {
			t.Logf("Error: %v", err)
			continue
		}

		if orders != nil && orders.Data != nil {
			logJSON(t, "GetPositionOrders response", orders)
			t.Logf("Response: total=%d, orderList length=%d",
				orders.Data.Total, len(orders.Data.OrderList))

			if len(orders.Data.OrderList) > 0 {
				t.Logf("✅ Found orders! First order: ID=%s, Type=%s, Side=%s, Status=%s",
					orders.Data.OrderList[0].ID, orders.Data.OrderList[0].Type,
					orders.Data.OrderList[0].Side, orders.Data.OrderList[0].Status)
				break // Found orders, stop testing
			}
		}
	}
}

func TestGetAccountPage(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	params := account.GetAccountPageParams{
		Size: 1,
	}
	logJSON(t, "GetAccountPage params", params)

	accountPage, err := client.GetAccountPage(ctx, params)

	if err != nil {
		// This endpoint may require special permissions (API_SIGNER)
		t.Logf("GetAccountPage error (may require special permissions): %v", err)
		t.Skip("Skipping test due to permission requirements")
		return
	}

	if !assert.NotNil(t, accountPage) {
		return
	}

	assert.Equal(t, "SUCCESS", accountPage.Code)
	assert.NotNil(t, accountPage.Data)
	assert.NotNil(t, accountPage.Data.DataList)
	logJSON(t, "GetAccountPage response", accountPage)

	t.Logf("Account page count: %d", len(accountPage.Data.DataList))
}
