package order

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/edgex-Tech/edgex-golang-sdk/sdk"
	"github.com/edgex-Tech/edgex-golang-sdk/sdk/order"
	"github.com/edgex-Tech/edgex-golang-sdk/test"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestV2CreateAndCancelOrder(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	if client.GetAPIVersion() != sdk.APIVersionV2 {
		t.Skip("Skipping v2 order test: API version is not v2")
	}
	if strings.TrimSpace(client.GetTradingPriKey()) == "" {
		t.Skip("Skipping v2 order test: TEST_TRADING_PRIVATE_KEY is required")
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

	if client.GetAPIVersion() != sdk.APIVersionV2 {
		t.Skip("Skipping v2 cancel-all test: API version is not v2")
	}

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
