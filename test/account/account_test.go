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
	assert.Equal(t, "SUCCESS", asset.GetCode())

	data := asset.GetData()
	assert.NotNil(t, data)
	assert.NotEmpty(t, data.GetCollateralList())
	assert.NotEmpty(t, data.GetPositionList())
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
	assert.Equal(t, "SUCCESS", positions.GetCode())

	data := positions.GetData()
	assert.NotNil(t, data)
	for _, position := range data {
		assert.NotEmpty(t, position.GetContractId())
		assert.NotEmpty(t, position.GetOpenSize())
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
	assert.Equal(t, "SUCCESS", transactions.GetCode())

	data := transactions.GetData()
	assert.NotNil(t, data)
	assert.NotNil(t, data.GetDataList())
	for _, tx := range data.GetDataList() {
		assert.NotEmpty(t, tx.GetId())
		assert.NotEmpty(t, tx.GetContractId())
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
	assert.Equal(t, "SUCCESS", transactions.GetCode())

	data := transactions.GetData()
	assert.NotNil(t, data)
	assert.NotNil(t, data.GetDataList())
	for _, tx := range data.GetDataList() {
		assert.NotEmpty(t, tx.GetId())
		assert.NotEmpty(t, tx.GetDeltaAmount())
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
	assert.Equal(t, "SUCCESS", terms.GetCode())

	data := terms.GetData()
	assert.NotNil(t, data)
	assert.NotNil(t, data.GetDataList())
	for _, term := range data.GetDataList() {
		assert.NotEmpty(t, term.GetAccountId())
		assert.NotEmpty(t, term.GetContractId())
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
	assert.Equal(t, "SUCCESS", account.GetCode())

	data := account.GetData()
	assert.NotNil(t, data)
	assert.NotEmpty(t, data.GetId())
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
	assert.Equal(t, "SUCCESS", deleverage.GetCode())

	data := deleverage.GetData()
	assert.NotNil(t, data)
	assert.NotNil(t, data.GetPositionContractIdToLightNumberMap())
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
	assert.Equal(t, "SUCCESS", snapshots.GetCode())

	data := snapshots.GetData()
	assert.NotNil(t, data)
	assert.NotNil(t, data.GetDataList())
	for _, snapshot := range data.GetDataList() {
		assert.NotEmpty(t, snapshot.GetCoinId())
		assert.NotEmpty(t, snapshot.GetTotalEquity())
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
	assert.Equal(t, "SUCCESS", transactions.GetCode())

	data := transactions.GetData()
	assert.NotNil(t, data)
	for _, tx := range data {
		assert.NotEmpty(t, tx.GetId())
		assert.NotEmpty(t, tx.GetContractId())
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
	assert.Equal(t, "SUCCESS", transactions.GetCode())

	data := transactions.GetData()
	assert.NotNil(t, data)
	for _, tx := range data {
		assert.NotEmpty(t, tx.GetId())
		assert.NotEmpty(t, tx.GetDeltaAmount())
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
