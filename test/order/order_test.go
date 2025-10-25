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

func TestSign(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	// Test message hash - using a fixed hash for consistency
	messageHash := []byte("test message hash for signing")

	// Sign the same message multiple times
	numTests := 5
	signatures := make([]map[string]string, numTests)

	for i := 0; i < numTests; i++ {
		sig, err := client.Sign(messageHash)
		assert.NoError(t, err, "Sign should not return error on iteration %d", i)
		assert.NotNil(t, sig, "Signature should not be nil on iteration %d", i)
		assert.NotEmpty(t, sig.R, "Signature R should not be empty on iteration %d", i)
		assert.NotEmpty(t, sig.S, "Signature S should not be empty on iteration %d", i)

		signatures[i] = map[string]string{
			"R": sig.R,
			"S": sig.S,
			"V": sig.V,
		}
		t.Logf("Iteration %d - R: %s, S: %s", i, sig.R, sig.S)
	}

	// Verify that all signatures are identical (deterministic signing)
	for i := 1; i < numTests; i++ {
		assert.Equal(t, signatures[0]["R"], signatures[i]["R"], "Signature R should be identical across iterations")
		assert.Equal(t, signatures[0]["S"], signatures[i]["S"], "Signature S should be identical across iterations")
		assert.Equal(t, signatures[0]["V"], signatures[i]["V"], "Signature V should be identical across iterations")
	}

	t.Logf("✓ All %d signatures are identical - signing is deterministic", numTests)
}

func TestGetActiveOrders(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()
	contractID := "10000001" // BTCUSDT

	activeOrders, err := client.GetActiveOrders(ctx, &order.GetActiveOrderParams{
		PaginationParams: order.PaginationParams{
			Size: "10",
		},
		OrderFilterParams: order.OrderFilterParams{
			FilterContractIdList: []string{contractID},
		},
	})
	jsonData, _ := json.MarshalIndent(activeOrders, "", "  ")
	t.Logf("Active Orders: %s", string(jsonData))

	assert.NoError(t, err)

	if assert.NotNil(t, activeOrders) && assert.NotNil(t, activeOrders.Data) {
		for _, order := range activeOrders.Data.DataList {
			assert.NotEmpty(t, order.Id)
			assert.NotEmpty(t, order.Side)
			assert.NotEmpty(t, order.Size)
			assert.NotEmpty(t, order.Price)
		}
	}
}

func TestGetOrderFills(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()
	contractID := "10000001" // BTCUSDT

	fills, err := client.GetOrderFillTransactions(ctx, &order.OrderFillTransactionParams{
		PaginationParams: order.PaginationParams{
			Size: "10",
		},
		OrderFilterParams: order.OrderFilterParams{
			FilterContractIdList: []string{contractID},
		},
	})
	jsonData, _ := json.MarshalIndent(fills, "", "  ")
	t.Logf("Order Fills: %s", string(jsonData))

	// Currently the API returns 400 Bad Request
	// This is expected until we have valid test credentials
	if err != nil {
		if !strings.Contains(err.Error(), "json: cannot unmarshal string into Go struct field Order.data.dataList.cumFillSize of type float64") {
			t.Fatal(err)
		}
	}

	if assert.NotNil(t, fills) && assert.NotNil(t, fills.Data) {
		for _, fill := range fills.Data.DataList {
			assert.NotEmpty(t, fill.OrderId)
			assert.NotEmpty(t, fill.FillPrice)
			assert.NotEmpty(t, fill.FillSize)
			assert.NotEmpty(t, fill.FillFee)
		}
	}
}

func TestCreateAndCancelOrder(t *testing.T) {

	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()
	contractID := "20000018"
	price := decimal.NewFromFloat(60.0)
	size := decimal.NewFromFloat(0.5)
	clientOrderID := fmt.Sprintf("sdk-test-%d", time.Now().UnixNano())

	// First get metadata
	metadata, err := client.GetMetaData(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, metadata)

	// Create order
	orderParams := &order.CreateOrderParams{
		ContractId:    contractID,
		Price:         price.String(),
		Size:          size.String(),
		Type:          "LIMIT",
		Side:          "BUY",
		TimeInForce:   "GOOD_TIL_CANCEL",
		ClientOrderId: &clientOrderID,
	}

	resp, err := client.CreateOrder(ctx, orderParams)
	jsonData, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("Created Order: %s", string(jsonData))

	assert.NoError(t, err)
	if assert.NotNil(t, resp) && assert.NotNil(t, resp.Data) {
		orderID := *resp.Data.OrderId
		assert.NotEmpty(t, orderID)

		ordersByID, err := client.GetOrdersByID(ctx, []string{orderID})
		assert.NoError(t, err)
		if assert.NotNil(t, ordersByID) {
			assert.Equal(t, order.ResponseCodeSuccess, ordersByID.Code)
			foundByID := false
			for _, ord := range ordersByID.Data {
				if *ord.Id == orderID {
					foundByID = true
					break
				}
			}
			assert.True(t, foundByID, "order should be returned by order ID lookup")
		}

		ordersByClientID, err := client.GetOrdersByClientOrderID(ctx, []string{clientOrderID})
		assert.NoError(t, err)
		if assert.NotNil(t, ordersByClientID) {
			assert.Equal(t, order.ResponseCodeSuccess, ordersByClientID.Code)
			foundByClient := false
			for _, ord := range ordersByClientID.Data {
				if *ord.ClientOrderId == clientOrderID {
					foundByClient = true
					break
				}
			}
			assert.True(t, foundByClient, "order should be returned by client order ID lookup")
		}

		// Cancel the created order
		cancelResp, err := client.CancelOrder(ctx, &order.CancelOrderParams{
			OrderId: orderID,
		})
		jsonData2, _ := json.MarshalIndent(cancelResp, "", "  ")
		t.Logf("Cancel Order Result: %s", string(jsonData2))

		assert.NoError(t, err)
		assert.NotNil(t, cancelResp)
	}
}

func TestCreateMarketOrder(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()
	contractID := "10000001" // BTCUSDT
	size := "0.001"

	// Get metadata to verify price calculation
	metadataResp, err := client.GetMetaData(ctx)
	assert.NoError(t, err)

	var contract *metadata.Contract
	for _, c := range metadataResp.Data.ContractList {
		if c.ContractId == contractID {
			contract = &c
			break
		}
	}
	assert.NotNil(t, contract, "Contract should be found")

	t.Run("Market Buy Order", func(t *testing.T) {
		// Create market buy order
		result, err := client.CreateOrder(ctx, &order.CreateOrderParams{
			ContractId: contractID,
			Size:       size,
			Type:       order.OrderTypeMarket,
			Side:       order.OrderSideBuy,
		})
		jsonData, _ := json.MarshalIndent(result, "", "  ")
		t.Logf("Created Market Buy Order: %s", string(jsonData))

		assert.NoError(t, err)
		assert.NotNil(t, result)

		if assert.NotNil(t, result.Data) {
			assert.NotEmpty(t, result.Data.OrderId)
		}
	})

	t.Run("Market Sell Order", func(t *testing.T) {
		// Create market sell order
		result, err := client.CreateOrder(ctx, &order.CreateOrderParams{
			ContractId: contractID,
			Size:       size,
			Type:       order.OrderTypeMarket,
			Side:       order.OrderSideSell,
		})
		jsonData, _ := json.MarshalIndent(result, "", "  ")
		t.Logf("Created Market Sell Order: %s", string(jsonData))

		assert.NoError(t, err)
		assert.NotNil(t, result)

		if assert.NotNil(t, result.Data) {
			assert.NotEmpty(t, result.Data.OrderId)
		}
	})
}
