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

// TestIntegration_LimitOrderFlow tests the complete limit order flow:
// 1. Get metadata
// 2. Get account asset
// 3. Get market quote
// 4. Calculate reasonable price
// 5. Create limit order (buy/sell)
// 6. Query active orders
// 7. Query order by ID
// 8. Cancel order
// 9. Verify order status
func TestIntegration_LimitOrderFlow(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	if strings.TrimSpace(client.GetSignerPriKey()) == "" {
		t.Skip("Skipping integration test: EDGEX_SIGNER_PRIVATE_KEY is required")
	}

	ctx := test.GetTestContext()

	// Step 1: Get metadata
	t.Log("Step 1: Getting metadata...")
	metaResp, err := client.GetMetaData(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, metaResp)
	assert.NotEmpty(t, metaResp.Data.ContractList)

	contract := &metaResp.Data.ContractList[0]
	contractID := contract.ContractId
	t.Logf("Using contract: %s", contractID)

	// Step 2: Get account asset
	t.Log("Step 2: Getting account asset...")
	assetResp, err := client.GetAccountAsset(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, assetResp)
	t.Logf("Account has %d collaterals", len(assetResp.Data.CollateralList))

	// Step 3: Get market quote
	t.Log("Step 3: Getting market quote...")
	quoteResp, err := client.Get24HourQuote(ctx, contractID)
	assert.NoError(t, err)
	assert.NotNil(t, quoteResp)
	assert.NotEmpty(t, quoteResp.Data)

	var lastPrice decimal.Decimal
	if quoteResp.Data[0].LastPrice != nil {
		lastPrice, err = decimal.NewFromString(*quoteResp.Data[0].LastPrice)
		assert.NoError(t, err)
	} else {
		lastPrice = decimal.NewFromFloat(50000)
	}
	t.Logf("Current price: %s", lastPrice.String())

	// Step 4: Calculate reasonable price (2% above market for sell order)
	t.Log("Step 4: Calculating order price...")
	targetPrice := lastPrice.Mul(decimal.NewFromFloat(1.02))
	tickSize, err := decimal.NewFromString(contract.TickSize)
	assert.NoError(t, err)
	orderPrice := targetPrice.Div(tickSize).Ceil().Mul(tickSize)
	t.Logf("Order price: %s (tick size: %s)", orderPrice.String(), contract.TickSize)

	// Determine order size
	orderSize := contract.MinOrderSize
	if orderSize == "" || orderSize == "0" {
		orderSize = contract.StepSize
	}
	if orderSize == "" || orderSize == "0" {
		orderSize = "0.01"
	}
	t.Logf("Order size: %s", orderSize)

	// Step 5: Create limit order (sell at higher price - won't fill immediately)
	t.Log("Step 5: Creating limit sell order...")
	clientOrderID := fmt.Sprintf("sdk-integration-%d", time.Now().UnixNano())
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
	t.Logf("Order created: ID=%s, ClientOrderID=%s", orderID, clientOrderID)

	// Wait a moment for order to be processed
	time.Sleep(1 * time.Second)

	// Step 6: Query active orders
	t.Log("Step 6: Querying active orders...")
	activeResp, err := client.GetActiveOrders(ctx, &order.GetActiveOrderParams{
		PaginationParams: order.PaginationParams{
			Size: "1",
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, activeResp)
	activeParamsJSON, _ := json.MarshalIndent(order.GetActiveOrderParams{
		PaginationParams: order.PaginationParams{
			Size: "1",
		},
	}, "", "  ")
	activeRespJSON, _ := json.MarshalIndent(activeResp, "", "  ")
	t.Logf("GetActiveOrders params: %s", string(activeParamsJSON))
	t.Logf("GetActiveOrders response: %s", string(activeRespJSON))
	t.Logf("Active orders count: %d", len(activeResp.Data.DataList))

	// Verify our order is in active list
	orderFound := false
	for _, o := range activeResp.Data.DataList {
		if o.Id != nil && *o.Id == orderID {
			orderFound = true
			if o.Status != nil {
				t.Logf("Found order in active list: Status=%s", *o.Status)
			}
			break
		}
	}
	assert.True(t, orderFound, "Order should be in active orders list")

	// Step 7: Query order by ID
	t.Log("Step 7: Querying order by ID...")
	orderByIDResp, err := client.GetOrdersByID(ctx, []string{orderID})
	assert.NoError(t, err)
	assert.NotNil(t, orderByIDResp)
	assert.NotEmpty(t, orderByIDResp.Data)

	foundOrder := orderByIDResp.Data[0]
	if foundOrder.Id != nil {
		assert.Equal(t, orderID, *foundOrder.Id)
	}
	if foundOrder.ContractId != nil {
		assert.Equal(t, contractID, *foundOrder.ContractId)
	}
	if foundOrder.Status != nil && foundOrder.Price != nil && foundOrder.Size != nil {
		t.Logf("Order details: Status=%s, Price=%s, Size=%s", *foundOrder.Status, *foundOrder.Price, *foundOrder.Size)
	}

	// Step 8: Cancel order
	t.Log("Step 8: Canceling order...")
	cancelResp, err := client.CancelOrder(ctx, &order.CancelOrderParams{
		OrderId: orderID,
	})
	assert.NoError(t, err)
	assert.NotNil(t, cancelResp)
	t.Log("Order canceled successfully")

	// Wait for cancellation to be processed
	time.Sleep(1 * time.Second)

	// Step 9: Verify order status
	t.Log("Step 9: Verifying order status after cancellation...")
	verifyResp, err := client.GetOrdersByID(ctx, []string{orderID})
	assert.NoError(t, err)
	assert.NotNil(t, verifyResp)

	if len(verifyResp.Data) > 0 {
		finalOrder := verifyResp.Data[0]
		if finalOrder.Status != nil {
			t.Logf("Final order status: %s", *finalOrder.Status)
			// Status should be CANCELLED or CANCELLING
			assert.Contains(t, []string{"CANCELLED", "CANCELLING", "CANCELED"}, strings.ToUpper(*finalOrder.Status))
		}
	}

	t.Log("✅ Limit order flow test completed successfully")
}
