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

// mustGetTestContract fetches test contract metadata
func mustGetTestContract(t *testing.T, client *sdk.Client, ctx context.Context) *metadata.Contract {
	t.Helper()
	meta, err := client.GetMetaData(ctx)
	if err != nil {
		t.Fatalf("Failed to get metadata: %v", err)
	}
	if meta == nil || meta.Data == nil || len(meta.Data.ContractList) == 0 {
		t.Fatal("No contracts available in metadata")
	}
	return &meta.Data.ContractList[0]
}

// getReferencePrice gets a reference price for testing
func getReferencePrice(t *testing.T, client *sdk.Client, ctx context.Context, contractID string) decimal.Decimal {
	t.Helper()
	ticker, err := client.Quote.Get24HourQuote(ctx, contractID)
	if err != nil || ticker == nil || len(ticker.Data) == 0 {
		return decimal.NewFromFloat(50000) // Default fallback
	}
	if ticker.Data[0].LastPrice != nil {
		price, err := decimal.NewFromString(*ticker.Data[0].LastPrice)
		if err == nil && !price.IsZero() {
			return price
		}
	}
	return decimal.NewFromFloat(50000)
}

// ceilToStep rounds price to nearest tick size
func ceilToStep(price decimal.Decimal, tickSize string) (decimal.Decimal, error) {
	tick, err := decimal.NewFromString(tickSize)
	if err != nil || tick.IsZero() {
		return price, nil
	}
	return price.Div(tick).Ceil().Mul(tick), nil
}

func TestV2CreateAndCancelOrder(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	if strings.TrimSpace(client.GetSignerPriKey()) == "" {
		t.Skip("Skipping v2 order test: TEST_SIGNER_PRIVATE_KEY is required")
	}

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

	clientOrderID := fmt.Sprintf("sdk-v2-%d", time.Now().UnixNano())
	resp, err := client.CreateOrder(ctx, &order.CreateOrderParams{
		ContractId:    contractID,
		Price:         price.String(),
		Size:          size,
		Type:          order.OrderTypeLimit,
		Side:          order.OrderSideSell,
		TimeInForce:   string(order.TimeInForce_GOOD_TIL_CANCEL),
		ClientOrderId: &clientOrderID,
	})
	orderJSON, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("V2 created order: %s", string(orderJSON))

	assert.NoError(t, err)
	if !assert.NotNil(t, resp) || !assert.NotNil(t, resp.Data) || !assert.NotNil(t, resp.Data.OrderId) {
		return
	}

	orderID := *resp.Data.OrderId
	cancelResp, err := client.CancelOrder(ctx, &order.CancelOrderParams{OrderId: orderID})
	cancelJSON, _ := json.MarshalIndent(cancelResp, "", "  ")
	t.Logf("V2 cancel order: %s", string(cancelJSON))

	assert.NoError(t, err)
	assert.NotNil(t, cancelResp)
}

func TestV2CancelAllOrdersByContract(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()
	contract := mustGetTestContract(t, client, ctx)

	resp, err := client.CancelOrder(ctx, &order.CancelOrderParams{
		ContractId: contract.ContractId,
	})
	respJSON, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("V2 cancel all orders response: %s", string(respJSON))

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestGetOrdersByID(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	// Use dummy order IDs for testing
	orderIDs := []string{"123456789"}

	orders, err := client.GetOrdersByID(ctx, orderIDs)

	if err != nil {
		assert.NoError(t, err)
		return
	}

	if !assert.NotNil(t, orders) {
		return
	}

	assert.Equal(t, "SUCCESS", orders.Code)
	assert.NotNil(t, orders.Data)

	t.Logf("Orders found: %d", len(orders.Data))
}

func TestGetOrdersByClientOrderID(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	// Use dummy client order IDs for testing
	clientOrderIDs := []string{"sdk-test-123"}

	orders, err := client.GetOrdersByClientOrderID(ctx, clientOrderIDs)

	if err != nil {
		assert.NoError(t, err)
		return
	}

	if !assert.NotNil(t, orders) {
		return
	}

	assert.Equal(t, "SUCCESS", orders.Code)
	assert.NotNil(t, orders.Data)

	t.Logf("Orders found: %d", len(orders.Data))
}

func TestGetOrderFillTransactions(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	params := order.OrderFillTransactionParams{
		PaginationParams: order.PaginationParams{
			Size: "10",
		},
	}

	transactions, err := client.GetOrderFillTransactions(ctx, &params)

	if err != nil {
		assert.NoError(t, err)
		return
	}

	if !assert.NotNil(t, transactions) {
		return
	}

	assert.Equal(t, "SUCCESS", transactions.Code)
	assert.NotNil(t, transactions.Data)
	assert.NotNil(t, transactions.Data.DataList)

	t.Logf("Order fill transactions: %d", len(transactions.Data.DataList))
}

func TestGetMaxOrderSize(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()
	contract := mustGetTestContract(t, client, ctx)

	price := decimal.NewFromFloat(50000)

	maxSize, err := client.GetMaxOrderSize(ctx, contract.ContractId, price)

	if err != nil {
		assert.NoError(t, err)
		return
	}

	if !assert.NotNil(t, maxSize) {
		return
	}

	assert.Equal(t, "SUCCESS", maxSize.Code)
	assert.NotNil(t, maxSize.Data)

	if maxSize.Data.MaxBuySize != nil {
		t.Logf("Max buy size: %s", *maxSize.Data.MaxBuySize)
	}
	if maxSize.Data.MaxSellSize != nil {
		t.Logf("Max sell size: %s", *maxSize.Data.MaxSellSize)
	}
}

func TestGetActiveOrders(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	params := order.GetActiveOrderParams{
		PaginationParams: order.PaginationParams{
			Size: "10",
		},
	}

	orders, err := client.GetActiveOrders(ctx, &params)

	if err != nil {
		assert.NoError(t, err)
		return
	}

	if !assert.NotNil(t, orders) {
		return
	}

	assert.Equal(t, "SUCCESS", orders.Code)
	assert.NotNil(t, orders.Data)
	assert.NotNil(t, orders.Data.DataList)

	t.Logf("Active orders: %d", len(orders.Data.DataList))
}
