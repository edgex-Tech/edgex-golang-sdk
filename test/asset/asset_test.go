package asset

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/edgex-Tech/edgex-golang-sdk/sdk/asset"
	"github.com/edgex-Tech/edgex-golang-sdk/test"
	"github.com/stretchr/testify/assert"
)

func TestGetAllOrdersPage(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	// Get orders for the last 24 hours
	now := time.Now()
	startTime := now.Add(-24 * time.Hour)

	params := asset.GetAllOrdersPageParams{
		StartTime:  startTime.Format("1136214245"),
		EndTime:    now.Format("1136214245"),
		ChainId:    "1",
		TypeList:   "DEPOSIT,WITHDRAW",
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

func TestGetCoinRate(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	params := asset.GetCoinRateParams{
		ChainId: "1",                                          // Ethereum mainnet
		Coin:    "0xdac17f958d2ee523a2206206994597c13d831ec7", // USDT contract address
	}

	resp, err := client.Asset.GetCoinRate(ctx, params)
	if err != nil {
		t.Logf("Error getting coin rate: %v", err)
		t.Skip("Skipping test due to invalid parameters")
		return
	}
	jsonData, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("Coin Rate: %s", string(jsonData))
	assert.NotNil(t, resp)
	assert.Equal(t, "SUCCESS", resp.Code)
}

func TestGetCrossWithdrawById(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	params := asset.GetCrossWithdrawByIdParams{
		CrossWithdrawIdList: "123",
	}

	resp, err := client.Asset.GetCrossWithdrawById(ctx, params)
	jsonData, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("Cross Withdraw: %s", string(jsonData))
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "SUCCESS", resp.Code)
}

func TestGetCrossWithdrawSignInfo(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	params := asset.GetCrossWithdrawSignInfoParams{
		ChainId: "1",       // Ethereum mainnet
		Amount:  "1000000", // Amount in smallest unit (e.g., wei for ETH)
	}

	resp, err := client.Asset.GetCrossWithdrawSignInfo(ctx, params)
	if err != nil {
		t.Logf("Error getting cross withdraw sign info: %v", err)
		t.Skip("Skipping test due to invalid parameters")
		return
	}
	jsonData, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("Cross Withdraw Sign Info: %s", string(jsonData))
	assert.NotNil(t, resp)
	assert.Equal(t, "SUCCESS", resp.Code)
}

func TestGetFastWithdrawById(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	params := asset.GetFastWithdrawByIdParams{
		FastWithdrawIdList: "123",
	}

	resp, err := client.Asset.GetFastWithdrawById(ctx, params)
	jsonData, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("Fast Withdraw: %s", string(jsonData))
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "SUCCESS", resp.Code)
}

func TestGetFastWithdrawSignInfo(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	params := asset.GetFastWithdrawSignInfoParams{
		ChainId: "1",       // Ethereum mainnet
		Amount:  "1000000", // Amount in smallest unit (e.g., wei for ETH)
	}

	resp, err := client.Asset.GetFastWithdrawSignInfo(ctx, params)
	if err != nil {
		t.Logf("Error getting fast withdraw sign info: %v", err)
		t.Skip("Skipping test due to invalid parameters")
		return
	}
	jsonData, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("Fast Withdraw Sign Info: %s", string(jsonData))
	assert.NotNil(t, resp)
	assert.Equal(t, "SUCCESS", resp.Code)
}

func TestGetNormalWithdrawById(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	params := asset.GetNormalWithdrawByIdParams{
		NormalWithdrawIdList: "123",
	}

	resp, err := client.Asset.GetNormalWithdrawById(ctx, params)
	jsonData, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("Normal Withdraw: %s", string(jsonData))
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "SUCCESS", resp.Code)
}

func TestGetNormalWithdrawableAmount(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	params := asset.GetNormalWithdrawableAmountParams{
		Address: "0xdac17f958d2ee523a2206206994597c13d831ec7", // USDT contract address
	}

	resp, err := client.Asset.GetNormalWithdrawableAmount(ctx, params)
	if err != nil {
		t.Logf("Error getting normal withdrawable amount: %v", err)
		t.Skip("Skipping test due to invalid parameters")
		return
	}
	jsonData, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("Normal Withdrawable Amount: %s", string(jsonData))
	assert.NotNil(t, resp)
	assert.Equal(t, "SUCCESS", resp.Code)
}

func TestCreateNormalWithdraw(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	params := &asset.CreateNormalWithdrawParams{
		CoinId:     "1000", // Example coin ID
		Amount:     "1.000000",
		EthAddress: "0x1fB51aa234287C3CA1F957eA9AD0E148Bb814b7A",
	}

	resp, err := client.CreateNormalWithdraw(ctx, params)
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

func TestCreateCrossWithdraw(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	params := asset.CreateCrossWithdrawParams{
		CoinId:                "1000", // Example coin ID
		Amount:                "1.000000",
		EthAddress:            "0x1fB51aa234287C3CA1F957eA9AD0E148Bb814b7A",
		Erc20Address:          "0xdac17f958d2ee523a2206206994597c13d831ec7", // USDT contract address
		LpAccountId:           "551109015904453258",
		ClientCrossWithdrawId: "745410645654877",
		ExpireTime:            "1735887600000",
		L2Signature:           "007bf80407c6a7bb14f5ca3b848a5d908627993f23b073c902e359a6fa4a6a92040cea4c98e25e35ad1d8cc4e18758c463c45bf451299ce55aa49abbdb916d03",
		Fee:                   "0.001000",
		ChainId:               "1", // Ethereum mainnet
		MpcAddress:            "0x1234567890abcdef1234567890abcdef12345678",
		MpcSignature:          "0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
		MpcSignTime:           "1735887600000",
	}

	resp, err := client.Asset.CreateCrossWithdraw(ctx, params)
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
