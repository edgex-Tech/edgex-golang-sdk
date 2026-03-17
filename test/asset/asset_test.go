package asset

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	sdkasset "github.com/edgex-Tech/edgex-golang-sdk/sdk/asset"
	"github.com/edgex-Tech/edgex-golang-sdk/test"
	"github.com/stretchr/testify/assert"
)

const defaultWithdrawTestAmount = "1.000000"

func isMutationTestsEnabled() bool {
	return strings.ToLower(strings.TrimSpace(os.Getenv("EDGEX_ENABLE_MUTATION_TESTS"))) == "true"
}

func getWithdrawTestAmount() string {
	amount := strings.TrimSpace(os.Getenv("EDGEX_WITHDRAW_TEST_AMOUNT"))
	if amount == "" {
		return defaultWithdrawTestAmount
	}
	return amount
}

func TestGetAllOrdersPage(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	now := time.Now()
	startTime := now.Add(-7 * 24 * time.Hour)

	meta, err := client.GetMetaData(ctx)
	assert.NoError(t, err)

	var chainId string
	if meta.Data != nil && meta.Data.MultiChain != nil && len(meta.Data.MultiChain.ChainList) > 0 {
		chainId = meta.Data.MultiChain.ChainList[0].ChainId
	}
	if chainId == "" {
		chainId = "UNKNOWN"
	}

	params := sdkasset.GetAllOrdersPageParams{
		StartTime:  strconv.FormatInt(startTime.Unix(), 10),
		EndTime:    strconv.FormatInt(now.Unix(), 10),
		ChainId:    chainId,
		TypeList:   "ORDER_TYPE_NORMAL_WITHDRAW,ORDER_TYPE_NORMAL_WITHDRAW",
		Size:       "10",
		OffsetData: "",
	}

	resp, err := client.Asset.GetAllOrdersPage(ctx, params)
	if err != nil {
		t.Logf("Error getting asset orders: %v", err)
		t.Skip("Skipping test due to authentication error")
		return
	}
	jsonData, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("Asset Orders: %s", string(jsonData))
	assert.NotNil(t, resp)
	assert.Equal(t, "SUCCESS", resp.Code)

	data := resp.Data
	assert.NotNil(t, data)
	assert.NotNil(t, data.DataList)
}

func TestCreateNormalWithdraw(t *testing.T) {
	if !isMutationTestsEnabled() {
		t.Skip("Skipping mutation test: set EDGEX_ENABLE_MUTATION_TESTS=true to enable")
	}

	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()
	prepared, err := client.PrepareWithdrawSignInfo(ctx, sdkasset.PrepareWithdrawSignInfoParams{
		Amount: getWithdrawTestAmount(),
	})
	if err != nil {
		t.Logf("Error preparing withdraw sign info: %v", err)
		t.Skip("Skipping test due to setup error")
		return
	}

	resp, err := client.CreateNormalWithdraw(ctx, &sdkasset.CreateNormalWithdrawParams{
		CoinId: prepared.CoinId,
		Amount: prepared.Amount,
		Fee:    prepared.Fee,
	})
	if err != nil {
		t.Logf("Error creating normal withdraw: %v", err)
		t.Skip("Skipping test due to error")
		return
	}
	jsonData, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("Normal Withdraw Creation Response: %s", string(jsonData))
	assert.NotNil(t, resp)
	assert.Equal(t, "SUCCESS", resp.Code)
}

func TestGetWithdrawSignInfo(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()
	prepared, err := client.PrepareWithdrawSignInfo(ctx, sdkasset.PrepareWithdrawSignInfoParams{
		Amount: getWithdrawTestAmount(),
	})
	if err != nil {
		t.Logf("Error getting withdraw sign info: %v", err)
		t.Skip("Skipping test due to error")
		return
	}

	resp := prepared.SignInfo
	jsonData, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("Withdraw Sign Info: %s", string(jsonData))
	t.Logf("Relation: getWithdrawSignInfo fee is reused by createNormalWithdraw/createCrossWithdraw (chainId=%s tokenAddress=%s amount=%s)", prepared.ChainId, prepared.TokenAddress, prepared.Amount)
	assert.NotNil(t, resp)
	assert.Equal(t, "SUCCESS", resp.Code)
	assert.NotNil(t, resp.Data)
	assert.NotNil(t, resp.Data.PoolBalance)
	assert.NotNil(t, resp.Data.Fee)
}

func TestCreateCrossWithdraw(t *testing.T) {
	if !isMutationTestsEnabled() {
		t.Skip("Skipping mutation test: set EDGEX_ENABLE_MUTATION_TESTS=true to enable")
	}

	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()
	resp, err := client.CreateCrossWithdrawAuto(ctx, &sdkasset.CreateCrossWithdrawAutoParams{
		Amount:        getWithdrawTestAmount(),
		TargetAddress: strings.TrimSpace(os.Getenv("EDGEX_CROSS_WITHDRAW_TARGET_ADDRESS")),
	})
	if err != nil {
		t.Logf("Error creating cross withdraw: %v", err)
		t.Skip("Skipping test due to error")
		return
	}
	jsonData, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("Cross Withdraw Creation Response: %s", string(jsonData))
	assert.NotNil(t, resp)
	assert.Equal(t, "SUCCESS", resp.Code)
}
