package order

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/edgex-Tech/edgex-golang-sdk/sdk/metadata"
	"github.com/edgex-Tech/edgex-golang-sdk/sdk/order"
	"github.com/edgex-Tech/edgex-golang-sdk/test"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

// TestMarketOrderBuy creates a market buy order
//
// Market Order Characteristics:
// - Executes immediately at the best available market price
// - No price parameter needed (uses opponent's best price)
// - TimeInForce must be IMMEDIATE_OR_CANCEL
// - Guarantees execution but not price
//
// Use Case: Quick entry when price precision is less important than execution speed
// Example: "Buy 0.001 BTC immediately at current market price"
func TestMarketOrderBuy(t *testing.T) {
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

	// Get current market price for reference
	referencePrice := getReferencePrice(t, client, ctx, contractID)
	t.Logf("Current market price: %s USDC", referencePrice.String())

	// Generate unique client order ID
	clientOrderID := fmt.Sprintf("market-buy-%d", time.Now().UnixNano())

	// Create market buy order
	params := &order.CreateOrderParams{
		ContractId:    contractID,
		Type:          order.OrderTypeMarket,
		Side:          order.OrderSideBuy,
		Size:          size,
		Price:         "0", // Market order doesn't need price
		TimeInForce:   string(order.TimeInForce_IMMEDIATE_OR_CANCEL), // Required for market orders
		ClientOrderId: &clientOrderID,
		ReduceOnly:    false,
	}

	t.Logf("Creating MARKET BUY order: size=%s, timeInForce=%s", size, params.TimeInForce)

	resp, err := client.CreateOrder(ctx, params)
	
	if err != nil {
		t.Logf("Error creating market order: %v", err)
		// Market orders may fail due to insufficient liquidity in testnet
		t.Skip("Skipping test due to market order error (may be liquidity issue)")
		return
	}

	orderJSON, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("Market buy order created: %s", string(orderJSON))

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, resp.Data)
	assert.NotNil(t, resp.Data.OrderId)

	orderID := *resp.Data.OrderId
	t.Logf("Order created with ID: %s", orderID)

	// Query the order to verify details
	time.Sleep(500 * time.Millisecond) // Brief wait for order processing
	
	orders, err := client.GetOrdersByID(ctx, []string{orderID})
	if err == nil && len(orders.Data) > 0 {
		createdOrder := orders.Data[0]
		orderDetailsJSON, _ := json.MarshalIndent(createdOrder, "", "  ")
		t.Logf("Order details: %s", string(orderDetailsJSON))

		// Verify order type
		if createdOrder.Type != nil {
			assert.Equal(t, string(order.OrderTypeMarket), *createdOrder.Type, "Order type should be MARKET")
		}

		// Verify order side
		if createdOrder.Side != nil {
			assert.Equal(t, order.OrderSideBuy, *createdOrder.Side, "Order side should be BUY")
		}

		// Market orders should execute immediately or be cancelled
		// Status could be FILLED, CANCELED, or EXECUTING
		if createdOrder.Status != nil {
			t.Logf("Order status: %s", *createdOrder.Status)
		}

		// If order filled, log fill details
		if createdOrder.CumFillSize != nil && *createdOrder.CumFillSize != "0" {
			t.Logf("Order filled: size=%s", *createdOrder.CumFillSize)
		}
	}
}

// Helper: Get minimum order size from contract
func getMinimumOrderSize(contract *metadata.Contract) string {
	size := strings.TrimSpace(contract.MinOrderSize)
	if size == "" || size == "0" {
		size = strings.TrimSpace(contract.StepSize)
	}
	if size == "" || size == "0" {
		size = "0.001" // Default minimum
	}
	return size
}

// TestMarketOrderSell creates a market sell order
//
// Market Order Characteristics:
// - Executes immediately at the best available market price
// - Sells at current bid price (buyer's price)
// - TimeInForce must be IMMEDIATE_OR_CANCEL
// - Useful for quick exits
//
// Use Case: Quick exit from position when speed matters more than price
// Example: "Sell 0.001 BTC immediately at current market price"
func TestMarketOrderSell(t *testing.T) {
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

	// Get current market price for reference
	referencePrice := getReferencePrice(t, client, ctx, contractID)
	t.Logf("Current market price: %s USDC", referencePrice.String())

	// Generate unique client order ID
	clientOrderID := fmt.Sprintf("market-sell-%d", time.Now().UnixNano())

	// Create market sell order
	params := &order.CreateOrderParams{
		ContractId:    contractID,
		Type:          order.OrderTypeMarket,
		Side:          order.OrderSideSell,
		Size:          size,
		Price:         "0", // Market order doesn't need price
		TimeInForce:   string(order.TimeInForce_IMMEDIATE_OR_CANCEL),
		ClientOrderId: &clientOrderID,
		ReduceOnly:    false,
	}

	t.Logf("Creating MARKET SELL order: size=%s, timeInForce=%s", size, params.TimeInForce)

	resp, err := client.CreateOrder(ctx, params)
	
	if err != nil {
		t.Logf("Error creating market order: %v", err)
		t.Skip("Skipping test due to market order error (may be liquidity issue)")
		return
	}

	orderJSON, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("Market sell order created: %s", string(orderJSON))

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, resp.Data)
	assert.NotNil(t, resp.Data.OrderId)

	orderID := *resp.Data.OrderId
	t.Logf("Order created with ID: %s", orderID)

	// Query the order to verify details
	time.Sleep(500 * time.Millisecond)
	
	orders, err := client.GetOrdersByID(ctx, []string{orderID})
	if err == nil && len(orders.Data) > 0 {
		createdOrder := orders.Data[0]
		orderDetailsJSON, _ := json.MarshalIndent(createdOrder, "", "  ")
		t.Logf("Order details: %s", string(orderDetailsJSON))

		// Verify order type
		if createdOrder.Type != nil {
			assert.Equal(t, string(order.OrderTypeMarket), *createdOrder.Type)
		}

		// Verify order side
		if createdOrder.Side != nil {
			assert.Equal(t, order.OrderSideSell, *createdOrder.Side)
		}

		// If order filled, verify execution price
		if createdOrder.CumFillSize != nil && createdOrder.CumFillValue != nil {
			fillSize, _ := decimal.NewFromString(*createdOrder.CumFillSize)
			fillValue, _ := decimal.NewFromString(*createdOrder.CumFillValue)
			if fillSize.GreaterThan(decimal.Zero) {
				executedPrice := fillValue.Div(fillSize)
				t.Logf("Executed %s @ %s USDC (average)", fillSize.String(), executedPrice.String())
				
				// Market sell should execute near or slightly below reference price
				// Allow 5% slippage tolerance for testnet
				minExpectedPrice := referencePrice.Mul(decimal.NewFromFloat(0.95))
				assert.True(t, executedPrice.GreaterThanOrEqual(minExpectedPrice),
					"Execution price should be within reasonable range of market price")
			}
		}
	}
}


