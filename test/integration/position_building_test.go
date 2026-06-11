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

// TestIntegration_BuildPosition tests building a position without immediately canceling:
// 1. Check current position
// 2. Create multiple limit orders at different prices
// 3. Keep orders open for potential fills
// 4. Verify orders are created successfully
func TestIntegration_BuildPosition(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	if strings.TrimSpace(client.GetSignerPriKey()) == "" {
		t.Skip("Skipping integration test: EDGEX_SIGNER_PRIVATE_KEY is required")
	}

	ctx := test.GetTestContext()

	// Step 1: Check current positions
	t.Log("Step 1: Checking current positions...")
	positionsResp, err := client.GetAccountPositions(ctx)
	assert.NoError(t, err)

	initialPositionCount := 0
	if positionsResp != nil && positionsResp.Data != nil {
		initialPositionCount = len(positionsResp.Data)
	}
	t.Logf("Initial positions: %d", initialPositionCount)

	// Step 2: Get contract and market data
	t.Log("Step 2: Getting contract and market data...")
	contract, err := test.ResolveTestContract(ctx, client)
	assert.NoError(t, err)
	contractID := contract.ContractId

	quoteResp, err := client.Get24HourQuote(ctx, contractID)
	assert.NoError(t, err)

	var lastPrice decimal.Decimal
	if len(quoteResp.Data) > 0 && quoteResp.Data[0].LastPrice != nil {
		lastPrice, _ = decimal.NewFromString(*quoteResp.Data[0].LastPrice)
	} else {
		lastPrice = decimal.NewFromFloat(50000)
	}

	tickSize, _ := decimal.NewFromString(contract.TickSize)
	orderSize := contract.MinOrderSize
	if orderSize == "" || orderSize == "0" {
		orderSize = "0.001"
	}

	t.Logf("Contract: %s, Current Price: %s, Order Size: %s", contractID, lastPrice.String(), orderSize)

	// Step 3: Create buy limit orders at different price levels
	// These orders will stay open unless filled
	t.Log("Step 3: Creating buy limit orders at different levels...")

	orderIDs := []string{}

	// Buy order 1: -1% below market (more likely to fill)
	price1 := lastPrice.Mul(decimal.NewFromFloat(0.99)).Div(tickSize).Floor().Mul(tickSize)
	clientOrderID1 := fmt.Sprintf("sdk-build-buy1-%d", time.Now().UnixNano())

	resp1, err := client.CreateOrder(ctx, &order.CreateOrderParams{
		ContractId:    contractID,
		Price:         price1.String(),
		Size:          orderSize,
		Type:          order.OrderTypeLimit,
		Side:          order.OrderSideBuy,
		TimeInForce:   string(order.TimeInForce_GOOD_TIL_CANCEL),
		ClientOrderId: &clientOrderID1,
	})

	if err != nil {
		t.Logf("Warning: Failed to create buy order 1: %v", err)
	} else if resp1 != nil && resp1.Data != nil && resp1.Data.OrderId != nil {
		orderID1 := *resp1.Data.OrderId
		orderIDs = append(orderIDs, orderID1)
		t.Logf("✅ Buy order 1 created: ID=%s, Price=%s (-1%%)", orderID1, price1.String())
	}

	time.Sleep(500 * time.Millisecond)

	// Buy order 2: -2% below market (less likely to fill)
	price2 := lastPrice.Mul(decimal.NewFromFloat(0.98)).Div(tickSize).Floor().Mul(tickSize)
	clientOrderID2 := fmt.Sprintf("sdk-build-buy2-%d", time.Now().UnixNano())

	resp2, err := client.CreateOrder(ctx, &order.CreateOrderParams{
		ContractId:    contractID,
		Price:         price2.String(),
		Size:          orderSize,
		Type:          order.OrderTypeLimit,
		Side:          order.OrderSideBuy,
		TimeInForce:   string(order.TimeInForce_GOOD_TIL_CANCEL),
		ClientOrderId: &clientOrderID2,
	})

	if err != nil {
		t.Logf("Warning: Failed to create buy order 2: %v", err)
	} else if resp2 != nil && resp2.Data != nil && resp2.Data.OrderId != nil {
		orderID2 := *resp2.Data.OrderId
		orderIDs = append(orderIDs, orderID2)
		t.Logf("✅ Buy order 2 created: ID=%s, Price=%s (-2%%)", orderID2, price2.String())
	}

	// Step 4: Verify active orders
	t.Log("Step 4: Verifying active orders...")
	time.Sleep(1 * time.Second)

	activeResp, err := client.GetActiveOrders(ctx, &order.GetActiveOrderParams{
		PaginationParams: order.PaginationParams{
			Size: "1",
		},
		OrderFilterParams: order.OrderFilterParams{
			FilterContractIdList: []string{contractID},
		},
	})

	assert.NoError(t, err)
	assert.NotNil(t, activeResp)
	activeParamsJSON, _ := json.MarshalIndent(order.GetActiveOrderParams{
		PaginationParams: order.PaginationParams{
			Size: "1",
		},
		OrderFilterParams: order.OrderFilterParams{
			FilterContractIdList: []string{contractID},
		},
	}, "", "  ")
	activeRespJSON, _ := json.MarshalIndent(activeResp, "", "  ")
	t.Logf("GetActiveOrders params: %s", string(activeParamsJSON))
	t.Logf("GetActiveOrders response: %s", string(activeRespJSON))
	t.Logf("Total active orders for contract: %d", len(activeResp.Data.DataList))

	// Step 5: Check if any orders filled
	t.Log("Step 5: Checking order fill status...")
	time.Sleep(2 * time.Second)

	for _, orderID := range orderIDs {
		orderResp, err := client.GetOrdersByID(ctx, []string{orderID})
		if err == nil && orderResp != nil && len(orderResp.Data) > 0 {
			order := orderResp.Data[0]
			if order.Status != nil {
				t.Logf("Order %s status: %s", orderID, *order.Status)
				if strings.Contains(strings.ToUpper(*order.Status), "FILLED") {
					t.Logf("🎉 Order %s was FILLED!", orderID)
				}
			}
		}
	}

	// Step 6: Check final positions
	t.Log("Step 6: Checking final positions...")
	finalPositionsResp, err := client.GetAccountPositions(ctx)
	assert.NoError(t, err)

	finalPositionCount := 0
	if finalPositionsResp != nil && finalPositionsResp.Data != nil {
		finalPositionCount = len(finalPositionsResp.Data)
	}
	t.Logf("Final positions: %d", finalPositionCount)

	if finalPositionCount > initialPositionCount {
		t.Logf("🎉 New position created! Total: %d (was: %d)", finalPositionCount, initialPositionCount)
	}

	t.Logf("✅ Position building test completed")
	t.Logf("📝 Note: Created %d limit buy orders that will stay open until filled or manually canceled", len(orderIDs))
	t.Logf("📝 Order IDs: %v", orderIDs)
}

// TestIntegration_CheckCurrentPosition checks the current position status
func TestIntegration_CheckCurrentPosition(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	// Get positions
	t.Log("Checking current positions...")
	positionsResp, err := client.GetAccountPositions(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, positionsResp)

	if positionsResp.Data == nil || len(positionsResp.Data) == 0 {
		t.Log("No current positions")
		return
	}

	t.Logf("Total positions: %d", len(positionsResp.Data))

	for i, pos := range positionsResp.Data {
		t.Logf("Position %d:", i+1)
		t.Logf("  ContractID: %s", pos.ContractID)
		t.Logf("  Size: %s", pos.Size)
		t.Logf("  Price: %s", pos.Price)
	}

	// Get active orders
	t.Log("Checking active orders...")
	activeResp, err := client.GetActiveOrders(ctx, &order.GetActiveOrderParams{
		PaginationParams: order.PaginationParams{
			Size: "50",
		},
	})

	if err == nil && activeResp != nil {
		t.Logf("Total active orders: %d", len(activeResp.Data.DataList))

		for i, ord := range activeResp.Data.DataList {
			if ord.Id != nil && ord.Status != nil && ord.Price != nil {
				t.Logf("Order %d: ID=%s, Status=%s, Price=%s, Side=%v",
					i+1, *ord.Id, *ord.Status, *ord.Price, ord.Side)
			}
		}
	}
}
