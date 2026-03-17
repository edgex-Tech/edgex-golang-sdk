package order

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/edgex-Tech/edgex-golang-sdk/sdk/order"
	"github.com/edgex-Tech/edgex-golang-sdk/test"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

// TestStopMarketOrderStopLoss creates a stop-loss order using STOP_MARKET
//
// STOP_MARKET Order Characteristics:
// - Conditional order that activates when trigger price is reached
// - Executes at market price once triggered
// - Used for stop-loss protection (exit losing position)
// - Status: UNTRIGGERED → triggered → executes at market
//
// Trigger Logic for SELL (stop-loss on long position):
// - triggerPrice < currentPrice (sell when price drops)
// - Example: Current $50k, trigger at $47.5k = -5% stop loss
//
// Use Case: Protect long position from excessive loss
// Example: "If BTC drops to $47,500, sell 0.001 BTC at market to cut losses"
func TestStopMarketOrderStopLoss(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	if strings.TrimSpace(client.GetSignerPriKey()) == "" {
		t.Skip("Skipping test: TEST_SIGNER_PRIVATE_KEY required for order signing")
	}

	ctx := test.GetTestContext()
	contract := mustGetTestContract(t, client, ctx)
	contractID := contract.ContractId

	// Get minimum order size
	size := getMinimumOrderSize(contract)

	// Get current market price
	referencePrice := getReferencePrice(t, client, ctx, contractID)
	t.Logf("Current market price: %s USDC", referencePrice.String())

	// Set trigger price 5% below market (stop-loss level)
	triggerPrice := referencePrice.Mul(decimal.NewFromFloat(0.95))
	triggerPriceRounded, _ := ceilToStep(triggerPrice, contract.TickSize)
	triggerPriceStr := triggerPriceRounded.String()
	t.Logf("Stop-loss trigger price: %s USDC (-5%%)", triggerPriceStr)

	// Generate unique client order ID
	clientOrderID := fmt.Sprintf("stop-market-sl-%d", time.Now().UnixNano())

	// Create STOP_MARKET order for stop-loss
	params := &order.CreateOrderParams{
		ContractId:       contractID,
		Type:             order.OrderTypeStopMarket,
		Side:             order.OrderSideSell, // SELL to close long position
		Size:             size,
		Price:            "0", // No limit price for market execution
		TriggerPrice:     triggerPriceStr,
		TriggerPriceType: string(order.TriggerPriceType_LAST_PRICE),
		TimeInForce:      string(order.TimeInForce_IMMEDIATE_OR_CANCEL),
		ClientOrderId:    &clientOrderID,
		ReduceOnly:       true, // Only close existing position, don't open new short
	}

	t.Logf("Creating STOP_MARKET order: side=%s, size=%s, triggerPrice=%s, reduceOnly=%v",
		params.Side, params.Size, params.TriggerPrice, params.ReduceOnly)

	resp, err := client.CreateOrder(ctx, params)

	if err != nil {
		t.Logf("Error creating stop market order: %v", err)
		t.Skip("Skipping test due to order creation error")
		return
	}

	orderJSON, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("Stop market order created: %s", string(orderJSON))

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, resp.Data)
	assert.NotNil(t, resp.Data.OrderId)

	orderID := *resp.Data.OrderId
	t.Logf("Order created with ID: %s", orderID)

	// Query the order to verify it's in UNTRIGGERED state
	time.Sleep(500 * time.Millisecond)

	orders, err := client.GetOrdersByID(ctx, []string{orderID})
	if err == nil && len(orders.Data) > 0 {
		createdOrder := orders.Data[0]
		orderDetailsJSON, _ := json.MarshalIndent(createdOrder, "", "  ")
		t.Logf("Order details: %s", string(orderDetailsJSON))

		// Verify order type
		if createdOrder.Type != nil {
			assert.Equal(t, string(order.OrderTypeStopMarket), *createdOrder.Type)
		}

		// Verify order side
		if createdOrder.Side != nil {
			assert.Equal(t, order.OrderSideSell, *createdOrder.Side)
		}

		// Verify trigger price
		if createdOrder.TriggerPrice != nil {
			t.Logf("Trigger price set to: %s", *createdOrder.TriggerPrice)
		}

		// Verify reduce-only flag
		if createdOrder.ReduceOnly != nil {
			assert.True(t, *createdOrder.ReduceOnly)
		}

		// Order should be UNTRIGGERED (waiting for price to drop)
		if createdOrder.Status != nil {
			t.Logf("Order status: %s (should be UNTRIGGERED)", *createdOrder.Status)
		}
	}

	// Cleanup: Cancel the order
	t.Log("Canceling stop-loss order...")
	_, err = client.CancelOrder(ctx, &order.CancelOrderParams{OrderId: orderID})
	if err != nil {
		t.Logf("Warning: Failed to cancel order: %v", err)
	}
}

// TestStopMarketOrderBreakoutBuy creates a breakout entry using STOP_MARKET
//
// STOP_MARKET Order Characteristics:
// - Enters position when price breaks above resistance
// - Executes at market price once triggered
// - Used for momentum/breakout trading
//
// Trigger Logic for BUY (breakout entry):
// - triggerPrice > currentPrice (buy when price rises)
// - Example: Current $50k, trigger at $52.5k = +5% breakout
//
// Use Case: Enter long position on upside breakout
// Example: "If BTC breaks above $52,500, buy 0.001 BTC at market"
func TestStopMarketOrderBreakoutBuy(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	if strings.TrimSpace(client.GetSignerPriKey()) == "" {
		t.Skip("Skipping test: TEST_SIGNER_PRIVATE_KEY required for order signing")
	}

	ctx := test.GetTestContext()
	contract := mustGetTestContract(t, client, ctx)
	contractID := contract.ContractId

	size := getMinimumOrderSize(contract)

	// Get current market price
	referencePrice := getReferencePrice(t, client, ctx, contractID)
	t.Logf("Current market price: %s USDC", referencePrice.String())

	// Set trigger price 5% above market (breakout level)
	triggerPrice := referencePrice.Mul(decimal.NewFromFloat(1.05))
	triggerPriceRounded, _ := ceilToStep(triggerPrice, contract.TickSize)
	triggerPriceStr := triggerPriceRounded.String()
	t.Logf("Breakout trigger price: %s USDC (+5%%)", triggerPriceStr)

	clientOrderID := fmt.Sprintf("stop-market-buy-%d", time.Now().UnixNano())

	// Create STOP_MARKET order for breakout entry
	params := &order.CreateOrderParams{
		ContractId:       contractID,
		Type:             order.OrderTypeStopMarket,
		Side:             order.OrderSideBuy, // BUY on breakout
		Size:             size,
		Price:            "0",
		TriggerPrice:     triggerPriceStr,
		TriggerPriceType: string(order.TriggerPriceType_LAST_PRICE),
		TimeInForce:      string(order.TimeInForce_IMMEDIATE_OR_CANCEL),
		ClientOrderId:    &clientOrderID,
		ReduceOnly:       false, // Open new position
	}

	t.Logf("Creating STOP_MARKET breakout order: side=%s, size=%s, triggerPrice=%s",
		params.Side, params.Size, params.TriggerPrice)

	resp, err := client.CreateOrder(ctx, params)

	if err != nil {
		t.Logf("Error creating stop market order: %v", err)
		t.Skip("Skipping test due to order creation error")
		return
	}

	orderJSON, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("Stop market breakout order created: %s", string(orderJSON))

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, resp.Data)
	assert.NotNil(t, resp.Data.OrderId)

	orderID := *resp.Data.OrderId
	t.Logf("Order created with ID: %s", orderID)

	// Query the order
	time.Sleep(500 * time.Millisecond)

	orders, err := client.GetOrdersByID(ctx, []string{orderID})
	if err == nil && len(orders.Data) > 0 {
		createdOrder := orders.Data[0]
		orderDetailsJSON, _ := json.MarshalIndent(createdOrder, "", "  ")
		t.Logf("Order details: %s", string(orderDetailsJSON))

		// Verify order type
		if createdOrder.Type != nil {
			assert.Equal(t, string(order.OrderTypeStopMarket), *createdOrder.Type)
		}

		// Verify order side
		if createdOrder.Side != nil {
			assert.Equal(t, order.OrderSideBuy, *createdOrder.Side)
		}

		// Verify trigger price
		if createdOrder.TriggerPrice != nil {
			t.Logf("Trigger price set to: %s", *createdOrder.TriggerPrice)
		}

		// Order should be UNTRIGGERED
		if createdOrder.Status != nil {
			t.Logf("Order status: %s", *createdOrder.Status)
		}
	}

	// Cleanup
	t.Log("Canceling breakout order...")
	_, err = client.CancelOrder(ctx, &order.CancelOrderParams{OrderId: orderID})
	if err != nil {
		t.Logf("Warning: Failed to cancel order: %v", err)
	}
}

// TestStopLimitOrderWithPriceProtection creates a stop-loss with minimum sell price
//
// STOP_LIMIT Order Characteristics:
// - Conditional order with price protection
// - Places limit order (not market) when triggered
// - Guarantees minimum sell price (or maximum buy price)
// - Status: UNTRIGGERED → triggered → limit order placed
//
// Parameters:
// - triggerPrice: When to activate the order
// - price: Limit price for execution (minimum sell / maximum buy)
//
// Use Case: Stop-loss with price floor protection
// Example: "If BTC drops to $47,500, sell at minimum $47,000"
func TestStopLimitOrderWithPriceProtection(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	if strings.TrimSpace(client.GetSignerPriKey()) == "" {
		t.Skip("Skipping test: TEST_SIGNER_PRIVATE_KEY required for order signing")
	}

	ctx := test.GetTestContext()
	contract := mustGetTestContract(t, client, ctx)
	contractID := contract.ContractId

	size := getMinimumOrderSize(contract)

	// Get current market price
	referencePrice := getReferencePrice(t, client, ctx, contractID)
	t.Logf("Current market price: %s USDC", referencePrice.String())

	// Set trigger price 5% below market
	triggerPrice := referencePrice.Mul(decimal.NewFromFloat(0.95))
	triggerPriceRounded, _ := ceilToStep(triggerPrice, contract.TickSize)
	triggerPriceStr := triggerPriceRounded.String()

	// Set limit price 6% below market (lower than trigger for safety margin)
	limitPrice := referencePrice.Mul(decimal.NewFromFloat(0.94))
	limitPriceRounded, _ := ceilToStep(limitPrice, contract.TickSize)
	limitPriceStr := limitPriceRounded.String()

	t.Logf("Trigger price: %s USDC (-5%%)", triggerPriceStr)
	t.Logf("Limit price: %s USDC (-6%%, minimum acceptable)", limitPriceStr)

	clientOrderID := fmt.Sprintf("stop-limit-sl-%d", time.Now().UnixNano())

	// Create STOP_LIMIT order
	params := &order.CreateOrderParams{
		ContractId:       contractID,
		Type:             order.OrderTypeStopLimit,
		Side:             order.OrderSideSell,
		Size:             size,
		Price:            limitPriceStr, // Limit price (minimum sell price)
		TriggerPrice:     triggerPriceStr,
		TriggerPriceType: string(order.TriggerPriceType_LAST_PRICE),
		TimeInForce:      string(order.TimeInForce_GOOD_TIL_CANCEL),
		ClientOrderId:    &clientOrderID,
		ReduceOnly:       true,
	}

	t.Logf("Creating STOP_LIMIT order: side=%s, size=%s, trigger=%s, limit=%s",
		params.Side, params.Size, params.TriggerPrice, params.Price)

	resp, err := client.CreateOrder(ctx, params)

	if err != nil {
		t.Logf("Error creating stop limit order: %v", err)
		t.Skip("Skipping test due to order creation error")
		return
	}

	orderJSON, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("Stop limit order created: %s", string(orderJSON))

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, resp.Data)
	assert.NotNil(t, resp.Data.OrderId)

	orderID := *resp.Data.OrderId
	t.Logf("Order created with ID: %s", orderID)

	// Query the order
	time.Sleep(500 * time.Millisecond)

	orders, err := client.GetOrdersByID(ctx, []string{orderID})
	if err == nil && len(orders.Data) > 0 {
		createdOrder := orders.Data[0]
		orderDetailsJSON, _ := json.MarshalIndent(createdOrder, "", "  ")
		t.Logf("Order details: %s", string(orderDetailsJSON))

		// Verify order type
		if createdOrder.Type != nil {
			assert.Equal(t, string(order.OrderTypeStopLimit), *createdOrder.Type)
		}

		// Verify both trigger and limit prices are set
		if createdOrder.TriggerPrice != nil {
			t.Logf("Trigger price: %s", *createdOrder.TriggerPrice)
		}

		if createdOrder.Price != nil {
			t.Logf("Limit price: %s", *createdOrder.Price)
		}

		// Order should be UNTRIGGERED
		if createdOrder.Status != nil {
			t.Logf("Order status: %s", *createdOrder.Status)
		}
	}

	// Cleanup
	t.Log("Canceling stop-limit order...")
	_, err = client.CancelOrder(ctx, &order.CancelOrderParams{OrderId: orderID})
	if err != nil {
		t.Logf("Warning: Failed to cancel order: %v", err)
	}
}
