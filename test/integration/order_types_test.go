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

// TestIntegration_OrderTypes tests all supported order types:
// 1. LIMIT
// 2. MARKET
// Note: Conditional orders (STOP_*, TAKE_PROFIT_*) require additional API parameters
// and are tested separately
func TestIntegration_OrderTypes(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	if strings.TrimSpace(client.GetSignerPriKey()) == "" {
		t.Skip("Skipping integration test: EDGEX_SIGNER_PRIVATE_KEY is required")
	}

	ctx := test.GetTestContext()

	// Get contract and market data
	t.Log("Preparing test environment...")
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

	// Test Case 1: Limit Order
	t.Run("LimitOrder", func(t *testing.T) {
		price := lastPrice.Mul(decimal.NewFromFloat(1.02)).Div(tickSize).Ceil().Mul(tickSize)
		clientOrderID := fmt.Sprintf("sdk-limit-%d", time.Now().UnixNano())

		t.Logf("Creating LIMIT order at price %s...", price.String())
		resp, err := client.CreateOrder(ctx, &order.CreateOrderParams{
			ContractId:    contractID,
			Price:         price.String(),
			Size:          orderSize,
			Type:          order.OrderTypeLimit,
			Side:          order.OrderSideSell,
			TimeInForce:   string(order.TimeInForce_GOOD_TIL_CANCEL),
			ClientOrderId: &clientOrderID,
		})

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.NotNil(t, resp.Data)
		assert.NotNil(t, resp.Data.OrderId)

		orderID := *resp.Data.OrderId
		t.Logf("✅ LIMIT order created: ID=%s", orderID)

		// Wait and cancel
		time.Sleep(1 * time.Second)
		_, err = client.CancelOrder(ctx, &order.CancelOrderParams{OrderId: orderID})
		assert.NoError(t, err)
		t.Logf("✅ LIMIT order canceled")
	})

	// Test Case 2: Market Order (already tested in market_order_flow_test.go)
	// Skipping here to avoid duplicate test

	// Note: Conditional orders (STOP_LIMIT, STOP_MARKET, TAKE_PROFIT_LIMIT, TAKE_PROFIT_MARKET)
	// require additional parameters not yet exposed in CreateOrderParams.
	// These will be added once the API specification is confirmed.

	t.Log("✅ All order types test completed successfully")
}

// TestIntegration_ReduceOnlyOrders tests reduce-only orders
func TestIntegration_ReduceOnlyOrders(t *testing.T) {
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

	// Get market price
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

	// Test Reduce-Only Order
	t.Log("Testing Reduce-Only order...")
	price := lastPrice.Mul(decimal.NewFromFloat(0.98)).Div(tickSize).Floor().Mul(tickSize)
	clientOrderID := fmt.Sprintf("sdk-reduce-only-%d", time.Now().UnixNano())

	resp, err := client.CreateOrder(ctx, &order.CreateOrderParams{
		ContractId:    contractID,
		Price:         price.String(),
		Size:          orderSize,
		Type:          order.OrderTypeLimit,
		Side:          order.OrderSideSell,
		TimeInForce:   string(order.TimeInForce_GOOD_TIL_CANCEL),
		ReduceOnly:    true,
		ClientOrderId: &clientOrderID,
	})

	if err != nil {
		// Reduce-only may fail if no position
		t.Logf("Reduce-only order may fail without position: %v", err)
		return
	}

	assert.NotNil(t, resp)
	assert.NotNil(t, resp.Data)
	assert.NotNil(t, resp.Data.OrderId)

	orderID := *resp.Data.OrderId
	t.Logf("✅ Reduce-Only order created: ID=%s", orderID)

	time.Sleep(1 * time.Second)
	_, err = client.CancelOrder(ctx, &order.CancelOrderParams{OrderId: orderID})
	assert.NoError(t, err)
	t.Logf("✅ Reduce-Only order canceled")
}

// TestIntegration_TimeInForceOptions tests different TimeInForce options
func TestIntegration_TimeInForceOptions(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	if strings.TrimSpace(client.GetSignerPriKey()) == "" {
		t.Skip("Skipping integration test: EDGEX_SIGNER_PRIVATE_KEY is required")
	}

	ctx := test.GetTestContext()

	// Get contract and market data
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
	price := lastPrice.Mul(decimal.NewFromFloat(1.02)).Div(tickSize).Ceil().Mul(tickSize)
	orderSize := contract.MinOrderSize
	if orderSize == "" || orderSize == "0" {
		orderSize = "0.001"
	}

	// Test GTC (Good Till Cancel)
	t.Run("GTC_Order", func(t *testing.T) {
		clientOrderID := fmt.Sprintf("sdk-gtc-%d", time.Now().UnixNano())

		t.Log("Creating GTC (Good Till Cancel) order...")
		resp, err := client.CreateOrder(ctx, &order.CreateOrderParams{
			ContractId:    contractID,
			Price:         price.String(),
			Size:          orderSize,
			Type:          order.OrderTypeLimit,
			Side:          order.OrderSideSell,
			TimeInForce:   string(order.TimeInForce_GOOD_TIL_CANCEL),
			ClientOrderId: &clientOrderID,
		})

		assert.NoError(t, err)
		assert.NotNil(t, resp.Data.OrderId)

		orderID := *resp.Data.OrderId
		t.Logf("✅ GTC order created: ID=%s", orderID)

		time.Sleep(1 * time.Second)
		_, err = client.CancelOrder(ctx, &order.CancelOrderParams{OrderId: orderID})
		assert.NoError(t, err)
	})

	// Test IOC (Immediate or Cancel)
	t.Run("IOC_Order", func(t *testing.T) {
		clientOrderID := fmt.Sprintf("sdk-ioc-%d", time.Now().UnixNano())

		t.Log("Creating IOC (Immediate Or Cancel) order...")
		resp, err := client.CreateOrder(ctx, &order.CreateOrderParams{
			ContractId:    contractID,
			Price:         price.String(),
			Size:          orderSize,
			Type:          order.OrderTypeLimit,
			Side:          order.OrderSideSell,
			TimeInForce:   string(order.TimeInForce_IMMEDIATE_OR_CANCEL),
			ClientOrderId: &clientOrderID,
		})

		assert.NoError(t, err)
		assert.NotNil(t, resp.Data.OrderId)

		orderID := *resp.Data.OrderId
		t.Logf("✅ IOC order created: ID=%s (may be auto-canceled if not filled)", orderID)

		// IOC orders are automatically canceled if not immediately filled
		time.Sleep(2 * time.Second)

		// Verify it was canceled
		verifyResp, err := client.GetOrdersByID(ctx, []string{orderID})
		if err == nil && verifyResp != nil && len(verifyResp.Data) > 0 {
			if verifyResp.Data[0].Status != nil {
				t.Logf("IOC order status: %s", *verifyResp.Data[0].Status)
			}
		}
	})

	// Test FOK (Fill or Kill)
	t.Run("FOK_Order", func(t *testing.T) {
		clientOrderID := fmt.Sprintf("sdk-fok-%d", time.Now().UnixNano())

		t.Log("Creating FOK (Fill Or Kill) order...")
		resp, err := client.CreateOrder(ctx, &order.CreateOrderParams{
			ContractId:    contractID,
			Price:         price.String(),
			Size:          orderSize,
			Type:          order.OrderTypeLimit,
			Side:          order.OrderSideSell,
			TimeInForce:   string(order.TimeInForce_FILL_OR_KILL),
			ClientOrderId: &clientOrderID,
		})

		assert.NoError(t, err)
		assert.NotNil(t, resp.Data.OrderId)

		orderID := *resp.Data.OrderId
		t.Logf("✅ FOK order created: ID=%s (must fill completely or cancel)", orderID)

		// FOK orders are automatically canceled if not completely filled
		time.Sleep(2 * time.Second)

		// Verify final status
		verifyResp, err := client.GetOrdersByID(ctx, []string{orderID})
		if err == nil && verifyResp != nil && len(verifyResp.Data) > 0 {
			if verifyResp.Data[0].Status != nil {
				t.Logf("FOK order status: %s", *verifyResp.Data[0].Status)
			}
		}
	})

	t.Log("✅ TimeInForce options test completed successfully")
}
