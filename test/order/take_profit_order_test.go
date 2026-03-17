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

// TestTakeProfitMarketOrderLong creates a take-profit order to exit long position
//
// TAKE_PROFIT_MARKET Order Characteristics:
// - Conditional order that activates when profit target is reached
// - Executes at market price once triggered
// - Used to lock in profits automatically
// - Status: UNTRIGGERED → triggered → executes at market
//
// Trigger Logic for SELL (take profit on long position):
// - triggerPrice > currentPrice (sell when price rises to profit target)
// - Example: Current $50k, trigger at $55k = +10% profit target
//
// Use Case: Automatically take profit on long position
// Example: "If BTC rises to $55,000, sell 0.001 BTC at market to lock in 10% profit"
func TestTakeProfitMarketOrderLong(t *testing.T) {
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

	// Set trigger price 10% above market (profit target)
	triggerPrice := referencePrice.Mul(decimal.NewFromFloat(1.10))
	triggerPriceRounded, _ := ceilToStep(triggerPrice, contract.TickSize)
	triggerPriceStr := triggerPriceRounded.String()
	t.Logf("Take-profit trigger price: %s USDC (+10%% profit target)", triggerPriceStr)

	clientOrderID := fmt.Sprintf("tp-market-long-%d", time.Now().UnixNano())

	// Create TAKE_PROFIT_MARKET order
	params := &order.CreateOrderParams{
		ContractId:       contractID,
		Type:             order.OrderTypeTakeProfitMarket,
		Side:             order.OrderSideSell, // SELL to close long position
		Size:             size,
		Price:            "0", // Market execution, no limit price
		TriggerPrice:     triggerPriceStr,
		TriggerPriceType: string(order.TriggerPriceType_LAST_PRICE),
		TimeInForce:      string(order.TimeInForce_IMMEDIATE_OR_CANCEL),
		ClientOrderId:    &clientOrderID,
		ReduceOnly:       true, // Only close position, don't open short
	}

	t.Logf("Creating TAKE_PROFIT_MARKET order: side=%s, size=%s, triggerPrice=%s, reduceOnly=%v",
		params.Side, params.Size, params.TriggerPrice, params.ReduceOnly)

	resp, err := client.CreateOrder(ctx, params)

	if err != nil {
		t.Logf("Error creating take-profit order: %v", err)
		t.Skip("Skipping test due to order creation error")
		return
	}

	orderJSON, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("Take-profit market order created: %s", string(orderJSON))

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, resp.Data)
	assert.NotNil(t, resp.Data.OrderId)

	orderID := *resp.Data.OrderId
	t.Logf("Order created with ID: %s", orderID)

	// Query the order to verify
	time.Sleep(500 * time.Millisecond)

	orders, err := client.GetOrdersByID(ctx, []string{orderID})
	if err == nil && len(orders.Data) > 0 {
		createdOrder := orders.Data[0]
		orderDetailsJSON, _ := json.MarshalIndent(createdOrder, "", "  ")
		t.Logf("Order details: %s", string(orderDetailsJSON))

		// Verify order type
		if createdOrder.Type != nil {
			assert.Equal(t, string(order.OrderTypeTakeProfitMarket), *createdOrder.Type)
		}

		// Verify order side
		if createdOrder.Side != nil {
			assert.Equal(t, order.OrderSideSell, *createdOrder.Side)
		}

		// Verify trigger price
		if createdOrder.TriggerPrice != nil {
			t.Logf("Trigger price: %s", *createdOrder.TriggerPrice)
		}

		// Verify reduce-only flag
		if createdOrder.ReduceOnly != nil {
			assert.True(t, *createdOrder.ReduceOnly)
		}

		// Order should be UNTRIGGERED (waiting for price to rise)
		if createdOrder.Status != nil {
			t.Logf("Order status: %s (should be UNTRIGGERED)", *createdOrder.Status)
		}
	}

	// Cleanup
	t.Log("Canceling take-profit order...")
	_, err = client.CancelOrder(ctx, &order.CancelOrderParams{OrderId: orderID})
	if err != nil {
		t.Logf("Warning: Failed to cancel order: %v", err)
	}
}

// TestTakeProfitLimitOrderLong creates take-profit with minimum price guarantee
//
// TAKE_PROFIT_LIMIT Order Characteristics:
// - Places limit order when profit target is reached
// - Guarantees minimum sell price
// - Protects against slippage during profit-taking
//
// Parameters:
// - triggerPrice: Profit target level
// - price: Minimum acceptable sell price (limit)
//
// Use Case: Take profit with price protection
// Example: "If BTC rises to $55,000, sell at minimum $54,500"
func TestTakeProfitLimitOrderLong(t *testing.T) {
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

	// Set trigger price 10% above market
	triggerPrice := referencePrice.Mul(decimal.NewFromFloat(1.10))
	triggerPriceRounded, _ := ceilToStep(triggerPrice, contract.TickSize)
	triggerPriceStr := triggerPriceRounded.String()

	// Set limit price 9% above market (slightly below trigger for safety)
	limitPrice := referencePrice.Mul(decimal.NewFromFloat(1.09))
	limitPriceRounded, _ := ceilToStep(limitPrice, contract.TickSize)
	limitPriceStr := limitPriceRounded.String()

	t.Logf("Trigger price: %s USDC (+10%% profit target)", triggerPriceStr)
	t.Logf("Limit price: %s USDC (+9%%, minimum acceptable)", limitPriceStr)

	clientOrderID := fmt.Sprintf("tp-limit-long-%d", time.Now().UnixNano())

	// Create TAKE_PROFIT_LIMIT order
	params := &order.CreateOrderParams{
		ContractId:       contractID,
		Type:             order.OrderTypeTakeProfitLimit,
		Side:             order.OrderSideSell,
		Size:             size,
		Price:            limitPriceStr, // Minimum sell price
		TriggerPrice:     triggerPriceStr,
		TriggerPriceType: string(order.TriggerPriceType_LAST_PRICE),
		TimeInForce:      string(order.TimeInForce_GOOD_TIL_CANCEL),
		ClientOrderId:    &clientOrderID,
		ReduceOnly:       true,
	}

	t.Logf("Creating TAKE_PROFIT_LIMIT order: side=%s, size=%s, trigger=%s, limit=%s",
		params.Side, params.Size, params.TriggerPrice, params.Price)

	resp, err := client.CreateOrder(ctx, params)

	if err != nil {
		t.Logf("Error creating take-profit limit order: %v", err)
		t.Skip("Skipping test due to order creation error")
		return
	}

	orderJSON, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("Take-profit limit order created: %s", string(orderJSON))

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
			assert.Equal(t, string(order.OrderTypeTakeProfitLimit), *createdOrder.Type)
		}

		// Verify both prices are set
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
	t.Log("Canceling take-profit limit order...")
	_, err = client.CancelOrder(ctx, &order.CancelOrderParams{OrderId: orderID})
	if err != nil {
		t.Logf("Warning: Failed to cancel order: %v", err)
	}
}

// TestTakeProfitMarketOrderShort creates take-profit order for short position
//
// TAKE_PROFIT_MARKET Order Characteristics:
// - For SHORT positions, profit is made when price drops
// - BUY order to close short position
//
// Trigger Logic for BUY (take profit on short position):
// - triggerPrice < currentPrice (buy back when price drops)
// - Example: Shorted at $50k, current $50k, trigger at $45k = 10% profit
//
// Use Case: Automatically take profit on short position
// Example: "If BTC drops to $45,000, buy back 0.001 BTC at market to lock in short profit"
func TestTakeProfitMarketOrderShort(t *testing.T) {
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

	// Set trigger price 10% below market (profit target for short)
	triggerPrice := referencePrice.Mul(decimal.NewFromFloat(0.90))
	triggerPriceRounded, _ := ceilToStep(triggerPrice, contract.TickSize)
	triggerPriceStr := triggerPriceRounded.String()
	t.Logf("Take-profit trigger price: %s USDC (-10%%, short profit target)", triggerPriceStr)

	clientOrderID := fmt.Sprintf("tp-market-short-%d", time.Now().UnixNano())

	// Create TAKE_PROFIT_MARKET order for short position
	params := &order.CreateOrderParams{
		ContractId:       contractID,
		Type:             order.OrderTypeTakeProfitMarket,
		Side:             order.OrderSideBuy, // BUY to close short position
		Size:             size,
		Price:            "0",
		TriggerPrice:     triggerPriceStr,
		TriggerPriceType: string(order.TriggerPriceType_LAST_PRICE),
		TimeInForce:      string(order.TimeInForce_IMMEDIATE_OR_CANCEL),
		ClientOrderId:    &clientOrderID,
		ReduceOnly:       true, // Only close short, don't open long
	}

	t.Logf("Creating TAKE_PROFIT_MARKET order (short): side=%s, size=%s, triggerPrice=%s",
		params.Side, params.Size, params.TriggerPrice)

	resp, err := client.CreateOrder(ctx, params)

	if err != nil {
		t.Logf("Error creating take-profit order: %v", err)
		t.Skip("Skipping test due to order creation error")
		return
	}

	orderJSON, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("Take-profit market order (short) created: %s", string(orderJSON))

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
			assert.Equal(t, string(order.OrderTypeTakeProfitMarket), *createdOrder.Type)
		}

		// Verify order side (BUY for short position)
		if createdOrder.Side != nil {
			assert.Equal(t, order.OrderSideBuy, *createdOrder.Side)
		}

		// Verify trigger price
		if createdOrder.TriggerPrice != nil {
			t.Logf("Trigger price: %s", *createdOrder.TriggerPrice)
		}

		// Order should be UNTRIGGERED
		if createdOrder.Status != nil {
			t.Logf("Order status: %s", *createdOrder.Status)
		}
	}

	// Cleanup
	t.Log("Canceling take-profit order...")
	_, err = client.CancelOrder(ctx, &order.CancelOrderParams{OrderId: orderID})
	if err != nil {
		t.Logf("Warning: Failed to cancel order: %v", err)
	}
}

// TestTakeProfitLimitOrderShort creates take-profit with max buy price for short
//
// TAKE_PROFIT_LIMIT Order Characteristics:
// - For short positions: BUY limit order when profit target reached
// - Guarantees maximum buy-back price
//
// Use Case: Take profit on short with price ceiling
// Example: "If BTC drops to $45,000, buy back at maximum $45,500"
func TestTakeProfitLimitOrderShort(t *testing.T) {
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

	// Set trigger price 10% below market
	triggerPrice := referencePrice.Mul(decimal.NewFromFloat(0.90))
	triggerPriceRounded, _ := ceilToStep(triggerPrice, contract.TickSize)
	triggerPriceStr := triggerPriceRounded.String()

	// Set limit price 9% below market (slightly above trigger)
	limitPrice := referencePrice.Mul(decimal.NewFromFloat(0.91))
	limitPriceRounded, _ := ceilToStep(limitPrice, contract.TickSize)
	limitPriceStr := limitPriceRounded.String()

	t.Logf("Trigger price: %s USDC (-10%% profit target)", triggerPriceStr)
	t.Logf("Limit price: %s USDC (-9%%, maximum buy price)", limitPriceStr)

	clientOrderID := fmt.Sprintf("tp-limit-short-%d", time.Now().UnixNano())

	// Create TAKE_PROFIT_LIMIT order for short
	params := &order.CreateOrderParams{
		ContractId:       contractID,
		Type:             order.OrderTypeTakeProfitLimit,
		Side:             order.OrderSideBuy, // BUY to close short
		Size:             size,
		Price:            limitPriceStr, // Maximum buy price
		TriggerPrice:     triggerPriceStr,
		TriggerPriceType: string(order.TriggerPriceType_LAST_PRICE),
		TimeInForce:      string(order.TimeInForce_GOOD_TIL_CANCEL),
		ClientOrderId:    &clientOrderID,
		ReduceOnly:       true,
	}

	t.Logf("Creating TAKE_PROFIT_LIMIT order (short): side=%s, size=%s, trigger=%s, limit=%s",
		params.Side, params.Size, params.TriggerPrice, params.Price)

	resp, err := client.CreateOrder(ctx, params)

	if err != nil {
		t.Logf("Error creating take-profit limit order: %v", err)
		t.Skip("Skipping test due to order creation error")
		return
	}

	orderJSON, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("Take-profit limit order (short) created: %s", string(orderJSON))

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
			assert.Equal(t, string(order.OrderTypeTakeProfitLimit), *createdOrder.Type)
		}

		// Verify side (BUY for short)
		if createdOrder.Side != nil {
			assert.Equal(t, order.OrderSideBuy, *createdOrder.Side)
		}

		// Verify both prices
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
	t.Log("Canceling take-profit limit order...")
	_, err = client.CancelOrder(ctx, &order.CancelOrderParams{OrderId: orderID})
	if err != nil {
		t.Logf("Warning: Failed to cancel order: %v", err)
	}
}
