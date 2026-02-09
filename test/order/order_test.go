package order

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/edgex-Tech/edgex-golang-sdk/sdk"
	"github.com/edgex-Tech/edgex-golang-sdk/sdk/metadata"
	"github.com/edgex-Tech/edgex-golang-sdk/sdk/order"
	"github.com/edgex-Tech/edgex-golang-sdk/test"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func mustGetTestContract(t *testing.T, client *sdk.Client, ctx context.Context) *metadata.Contract {
	t.Helper()

	metadataResp, err := client.GetMetaData(ctx)
	assert.NoError(t, err)
	if metadataResp == nil || metadataResp.Data == nil {
		t.Fatalf("metadata response is nil")
	}
	if len(metadataResp.Data.ContractList) == 0 {
		t.Fatalf("metadata contract list is empty")
	}

	for i := range metadataResp.Data.ContractList {
		contract := &metadataResp.Data.ContractList[i]
		if strings.TrimSpace(contract.ContractId) != "" {
			return contract
		}
	}

	t.Fatalf("no valid contractId found in metadata")
	return nil
}

func ceilToStep(value decimal.Decimal, step string) (decimal.Decimal, error) {
	stepDec, err := decimal.NewFromString(strings.TrimSpace(step))
	if err != nil {
		return decimal.Zero, fmt.Errorf("invalid step %q: %w", step, err)
	}
	if !stepDec.GreaterThan(decimal.Zero) {
		return decimal.Zero, fmt.Errorf("step must be > 0, got %q", step)
	}
	return value.Div(stepDec).Ceil().Mul(stepDec), nil
}

func getReferencePrice(t *testing.T, client *sdk.Client, ctx context.Context, contractID string) decimal.Decimal {
	t.Helper()

	tickerResp, err := client.Get24HourQuote(ctx, contractID)
	assert.NoError(t, err)
	if tickerResp == nil || len(tickerResp.Data) == 0 {
		t.Fatalf("ticker response is empty for contract %s", contractID)
	}

	ticker := tickerResp.Data[0]
	candidates := []*string{ticker.OraclePrice, ticker.IndexPrice, ticker.LastPrice, ticker.Close}
	for _, candidate := range candidates {
		if candidate == nil || strings.TrimSpace(*candidate) == "" {
			continue
		}
		price, err := decimal.NewFromString(*candidate)
		if err == nil && price.GreaterThan(decimal.Zero) {
			return price
		}
	}

	t.Fatalf("no valid reference price found for contract %s", contractID)
	return decimal.Zero
}

func TestSign(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)
	if strings.TrimSpace(client.GetStarkPriKey()) == "" {
		t.Skip("stark private key is not configured; skipping stark signature test")
	}

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
	contract := mustGetTestContract(t, client, ctx)
	contractID := contract.ContractId

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
	contract := mustGetTestContract(t, client, ctx)
	contractID := contract.ContractId

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
	contract := mustGetTestContract(t, client, ctx)
	contractID := contract.ContractId
	size := strings.TrimSpace(contract.MinOrderSize)
	if size == "" || size == "0" {
		size = strings.TrimSpace(contract.StepSize)
	}
	if size == "" || size == "0" {
		size = "0.01"
	}

	referencePrice := getReferencePrice(t, client, ctx, contractID)
	targetPrice := referencePrice.Mul(decimal.NewFromFloat(1.02))
	price, err := ceilToStep(targetPrice, contract.TickSize)
	assert.NoError(t, err)
	if !price.GreaterThan(decimal.Zero) {
		price = referencePrice
	}

	clientOrderID := fmt.Sprintf("sdk-test-%d", time.Now().UnixNano())

	// Create order
	orderParams := &order.CreateOrderParams{
		ContractId:    contractID,
		Price:         price.String(),
		Size:          size,
		Type:          "LIMIT",
		Side:          "SELL",
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
	contract := mustGetTestContract(t, client, ctx)
	contractID := contract.ContractId
	size := strings.TrimSpace(contract.MinOrderSize)
	if size == "" || size == "0" {
		size = strings.TrimSpace(contract.StepSize)
	}
	if size == "" || size == "0" {
		size = "0.01"
	}

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
