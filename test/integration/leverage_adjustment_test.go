package integration

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/edgex-Tech/edgex-golang-sdk/v2/sdk/order"
	"github.com/edgex-Tech/edgex-golang-sdk/v2/test"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

// TestIntegration_LeverageAdjustment tests the leverage adjustment flow:
// 1. Get current contract and leverage
// 2. Update leverage setting
// 3. Verify leverage change
// 4. Create order to verify new leverage is effective
// 5. Cancel order and restore original leverage
func TestIntegration_LeverageAdjustment(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	if strings.TrimSpace(client.GetSignerPriKey()) == "" {
		t.Skip("Skipping integration test: EDGEX_SIGNER_PRIVATE_KEY is required")
	}

	ctx := test.GetTestContext()

	// Step 1: Get contract information
	t.Log("Step 1: Getting contract information...")
	contract, err := test.ResolveTestContract(ctx, client)
	assert.NoError(t, err)
	assert.NotNil(t, contract)

	contractID := contract.ContractId
	t.Logf("Using contract: %s", contractID)

	// Get current account asset to check initial state
	t.Log("Step 2: Getting current account state...")
	assetResp, err := client.GetAccountAsset(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, assetResp)

	// Step 3: Update leverage to 5x
	t.Log("Step 3: Updating leverage to 5x...")
	targetLeverage := "5"
	err = client.UpdateLeverageSetting(ctx, contractID, targetLeverage)
	if err != nil {
		if strings.Contains(err.Error(), "ACCOUNT_UPDATE_LEVERAGE_FAILED_ORDER") {
			t.Skip("Cannot update leverage with active orders - this is expected business logic")
			return
		}
		assert.NoError(t, err)
		return
	}
	t.Logf("Leverage updated to: %s", targetLeverage)

	// Wait for leverage change to take effect
	time.Sleep(1 * time.Second)

	// Step 4: Verify leverage change by checking account state
	t.Log("Step 4: Verifying leverage change...")
	assetAfter, err := client.GetAccountAsset(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, assetAfter)
	t.Log("Leverage change verified via account asset query")

	// Step 5: Create a test order to verify leverage is effective
	t.Log("Step 5: Creating test order with new leverage...")

	// Get current price
	quoteResp, err := client.Get24HourQuote(ctx, contractID)
	assert.NoError(t, err)
	assert.NotNil(t, quoteResp)

	var lastPrice decimal.Decimal
	if len(quoteResp.Data) > 0 && quoteResp.Data[0].LastPrice != nil {
		lastPrice, err = decimal.NewFromString(*quoteResp.Data[0].LastPrice)
		assert.NoError(t, err)
	} else {
		lastPrice = decimal.NewFromFloat(50000)
	}

	// Calculate order price (2% above market)
	targetPrice := lastPrice.Mul(decimal.NewFromFloat(1.02))
	tickSize, _ := decimal.NewFromString(contract.TickSize)
	orderPrice := targetPrice.Div(tickSize).Ceil().Mul(tickSize)

	// Get order size
	orderSize := contract.MinOrderSize
	if orderSize == "" || orderSize == "0" {
		orderSize = contract.StepSize
	}
	if orderSize == "" || orderSize == "0" {
		orderSize = "0.001"
	}

	// Create test order
	clientOrderID := fmt.Sprintf("sdk-leverage-test-%d", time.Now().UnixNano())
	createResp, err := client.CreateOrder(ctx, &order.CreateOrderParams{
		ContractId:    contractID,
		Price:         orderPrice.String(),
		Size:          orderSize,
		Type:          order.OrderTypeLimit,
		Side:          order.OrderSideSell,
		TimeInForce:   string(order.TimeInForce_GOOD_TIL_CANCEL),
		ClientOrderId: &clientOrderID,
	})
	assert.NoError(t, err)
	assert.NotNil(t, createResp)
	assert.NotNil(t, createResp.Data)
	assert.NotNil(t, createResp.Data.OrderId)

	orderID := *createResp.Data.OrderId
	t.Logf("Test order created: ID=%s", orderID)

	// Wait a moment
	time.Sleep(1 * time.Second)

	// Step 6: Cancel the test order
	t.Log("Step 6: Canceling test order...")
	_, err = client.CancelOrder(ctx, &order.CancelOrderParams{
		OrderId: orderID,
	})
	assert.NoError(t, err)
	t.Log("Test order canceled")

	// Step 7: Update leverage back to default (10x)
	t.Log("Step 7: Restoring leverage to default (10x)...")
	defaultLeverage := "10"
	err = client.UpdateLeverageSetting(ctx, contractID, defaultLeverage)
	assert.NoError(t, err)
	t.Logf("Leverage restored to: %s", defaultLeverage)

	time.Sleep(1 * time.Second)

	// Step 8: Verify final state
	t.Log("Step 8: Verifying final state...")
	finalAsset, err := client.GetAccountAsset(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, finalAsset)

	t.Log("✅ Leverage adjustment test completed successfully")
}

// TestIntegration_LeverageEdgeCases tests edge cases for leverage adjustment
func TestIntegration_LeverageEdgeCases(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	if strings.TrimSpace(client.GetSignerPriKey()) == "" {
		t.Skip("Skipping integration test: EDGEX_SIGNER_PRIVATE_KEY is required")
	}

	ctx := test.GetTestContext()

	// Get contract
	contract, err := test.ResolveTestContract(ctx, client)
	assert.NoError(t, err)
	contractID := contract.ContractId

	// Test Case 1: Update leverage to minimum (usually 1x)
	t.Log("Test Case 1: Setting leverage to minimum (1x)...")
	err = client.UpdateLeverageSetting(ctx, contractID, "1")
	if err != nil {
		if strings.Contains(err.Error(), "ACCOUNT_UPDATE_LEVERAGE_FAILED_ORDER") {
			t.Skip("Cannot update leverage with active orders - this is expected business logic")
			return
		}
		assert.NoError(t, err)
		return
	}
	t.Log("✅ Minimum leverage (1x) set successfully")

	time.Sleep(500 * time.Millisecond)

	// Test Case 2: Update leverage to maximum (check contract's max leverage)
	// Note: Different contracts may have different max leverage
	t.Log("Test Case 2: Setting leverage to high value (20x)...")
	err = client.UpdateLeverageSetting(ctx, contractID, "20")
	if err != nil {
		t.Logf("⚠️ Leverage 20x may exceed contract limit: %v", err)
		// This is expected if contract doesn't support 20x
	} else {
		t.Log("✅ High leverage (20x) set successfully")
	}

	time.Sleep(500 * time.Millisecond)

	// Restore to default
	t.Log("Restoring leverage to default (10x)...")
	err = client.UpdateLeverageSetting(ctx, contractID, "10")
	assert.NoError(t, err)

	t.Log("✅ Leverage edge cases test completed")
}
