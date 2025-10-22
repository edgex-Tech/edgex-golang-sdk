package account

import (
	"encoding/json"
	"testing"

	"github.com/edgex-Tech/edgex-golang-sdk/sdk/account"
	"github.com/edgex-Tech/edgex-golang-sdk/test"
	"github.com/stretchr/testify/assert"
)

func TestGetAccountAsset(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	asset, err := client.GetAccountAsset(ctx)
	jsonData, _ := json.MarshalIndent(asset, "", "  ")
	t.Logf("Account Asset: %s", string(jsonData))
	assert.NoError(t, err)
	assert.NotNil(t, asset)
	assert.Equal(t, "SUCCESS", asset.Code)

	data := asset.Data
	assert.NotNil(t, data)
	assert.NotEmpty(t, data.CollateralList)
	assert.NotEmpty(t, data.PositionList)
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
		Size: 10,
	}

	transactions, err := client.GetPositionTransactionPage(ctx, params)
	jsonData, _ := json.MarshalIndent(transactions, "", "  ")
	t.Logf("Position Transaction Page: %s", string(jsonData))
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
		Size: 10,
	}

	transactions, err := client.GetCollateralTransactionPage(ctx, params)
	jsonData, _ := json.MarshalIndent(transactions, "", "  ")
	t.Logf("Collateral Transaction Page: %s", string(jsonData))
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
		Size: 10,
	}

	terms, err := client.GetPositionTermPage(ctx, params)
	jsonData, _ := json.MarshalIndent(terms, "", "  ")
	t.Logf("Position Term Page: %s", string(jsonData))
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

	params := account.GetAccountAssetSnapshotPageParams{
		Size:   10,
		CoinID: "1000", // Example coin ID
	}

	snapshots, err := client.GetAccountAssetSnapshotPage(ctx, params)
	jsonData, _ := json.MarshalIndent(snapshots, "", "  ")
	t.Logf("Account Asset Snapshot Page: %s", string(jsonData))
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

	// Test updating leverage setting within allowed range on an active contract
	err = client.UpdateLeverageSetting(ctx, "20000018", "5")
	assert.NoError(t, err)
}
