package integration

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/edgex-Tech/edgex-golang-sdk/sdk/order"
	"github.com/edgex-Tech/edgex-golang-sdk/test"
	"github.com/stretchr/testify/assert"
)

// TestIntegration_MarketOrderFlow tests the market order flow:
// 1. Get metadata
// 2. Get account asset
// 3. Create market order (should fill immediately)
// 4. Verify immediate execution
// 5. Query fill transactions
// 6. Verify position changes
func TestIntegration_MarketOrderFlow(t *testing.T) {
	// Market orders execute immediately and affect real positions
	// Only run if explicitly enabled
	if strings.ToLower(strings.TrimSpace(os.Getenv("EDGEX_ENABLE_MUTATION_TESTS"))) != "true" {
		t.Skip("Skipping market order test: set EDGEX_ENABLE_MUTATION_TESTS=true to enable")
	}

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

	contract := &metaResp.Data.ContractList[0]
	contractID := contract.ContractId
	t.Logf("Using contract: %s", contractID)

	// Step 2: Get account asset before order
	t.Log("Step 2: Getting account asset before order...")
	assetBefore, err := client.GetAccountAsset(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, assetBefore)

	// Get positions before
	positionsBefore, err := client.GetAccountPositions(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, positionsBefore)
	initialPositionCount := len(positionsBefore.Data)
	t.Logf("Initial positions: %d", initialPositionCount)

	// Determine order size (use minimum)
	orderSize := contract.MinOrderSize
	if orderSize == "" || orderSize == "0" {
		orderSize = contract.StepSize
	}
	if orderSize == "" || orderSize == "0" {
		orderSize = "0.001" // Very small size for market order
	}
	t.Logf("Order size: %s", orderSize)

	// Step 3: Create market buy order
	t.Log("Step 3: Creating market buy order...")
	clientOrderID := fmt.Sprintf("sdk-market-%d", time.Now().UnixNano())
	createResp, err := client.CreateOrder(ctx, &order.CreateOrderParams{
		ContractId:    contractID,
		Price:         "", // Market orders don't specify price
		Size:          orderSize,
		Type:          order.OrderTypeMarket,
		Side:          order.OrderSideBuy,
		ClientOrderId: &clientOrderID,
	})
	assert.NoError(t, err)
	assert.NotNil(t, createResp)
	assert.NotNil(t, createResp.Data)

	if createResp.Data.OrderId != nil {
		orderID := *createResp.Data.OrderId
		t.Logf("Market order created: ID=%s", orderID)

		// Step 4: Wait a moment and verify order status
		time.Sleep(2 * time.Second)

		t.Log("Step 4: Verifying order execution...")
		orderResp, err := client.GetOrdersByID(ctx, []string{orderID})
		assert.NoError(t, err)
		assert.NotNil(t, orderResp)

		if len(orderResp.Data) > 0 {
			orderStatus := orderResp.Data[0].Status
			if orderStatus != nil {
				t.Logf("Order status: %s", *orderStatus)
				// Market order should be FILLED or PARTIALLY_FILLED
				statusUpper := strings.ToUpper(*orderStatus)
				assert.True(t,
					strings.Contains(statusUpper, "FILLED") || strings.Contains(statusUpper, "COMPLETE"),
					"Market order should be filled")
			}
		}

		// Step 5: Query fill transactions
		t.Log("Step 5: Querying fill transactions...")
		fillResp, err := client.GetOrderFillTransactions(ctx, &order.OrderFillTransactionParams{
			PaginationParams: order.PaginationParams{
				Size: "1",
			},
			OrderFilterParams: order.OrderFilterParams{},
			FilterOrderIdList: []string{orderID},
		})
		assert.NoError(t, err)
		assert.NotNil(t, fillResp)
		fillParamsJSON, _ := json.MarshalIndent(order.OrderFillTransactionParams{
			PaginationParams: order.PaginationParams{
				Size: "1",
			},
			OrderFilterParams: order.OrderFilterParams{},
			FilterOrderIdList: []string{orderID},
		}, "", "  ")
		fillRespJSON, _ := json.MarshalIndent(fillResp, "", "  ")
		t.Logf("GetOrderFillTransactions params: %s", string(fillParamsJSON))
		t.Logf("GetOrderFillTransactions response: %s", string(fillRespJSON))
		t.Logf("Fill transactions found: %d", len(fillResp.Data.DataList))

		// Step 6: Verify position changes
		t.Log("Step 6: Verifying position changes...")
		positionsAfter, err := client.GetAccountPositions(ctx)
		assert.NoError(t, err)
		assert.NotNil(t, positionsAfter)

		finalPositionCount := len(positionsAfter.Data)
		t.Logf("Final positions: %d", finalPositionCount)

		// Should have at least one position after market buy
		assert.GreaterOrEqual(t, finalPositionCount, initialPositionCount,
			"Position count should increase or stay same after market buy")

		t.Log("✅ Market order flow test completed successfully")
	}
}
