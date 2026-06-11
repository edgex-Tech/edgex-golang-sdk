package integration

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/edgex-Tech/edgex-golang-sdk/v2/sdk/order"
	"github.com/edgex-Tech/edgex-golang-sdk/v2/test"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

// TestIntegration_OrderQuery tests comprehensive order query and filtering:
// 1. Create multiple orders with different types
// 2. Query active orders with filters
// 3. Query by client order ID
// 4. Query order fill transactions
// 5. Query by order ID
// 6. Clean up all test orders
func TestIntegration_OrderQuery(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	if strings.TrimSpace(client.GetSignerPriKey()) == "" {
		t.Skip("Skipping integration test: EDGEX_SIGNER_PRIVATE_KEY is required")
	}

	ctx := test.GetTestContext()

	// Step 1: Prepare test data
	t.Log("Step 1: Getting contract and market data...")
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

	// Step 2: Create multiple test orders
	t.Log("Step 2: Creating multiple test orders...")
	orderIDs := []string{}
	clientOrderIDs := []string{}

	// Create Order 1: Sell at +2%
	price1 := lastPrice.Mul(decimal.NewFromFloat(1.02)).Div(tickSize).Ceil().Mul(tickSize)
	clientOrderID1 := fmt.Sprintf("sdk-query-test-1-%d", time.Now().UnixNano())
	resp1, err := client.CreateOrder(ctx, &order.CreateOrderParams{
		ContractId:    contractID,
		Price:         price1.String(),
		Size:          orderSize,
		Type:          order.OrderTypeLimit,
		Side:          order.OrderSideSell,
		TimeInForce:   string(order.TimeInForce_GOOD_TIL_CANCEL),
		ClientOrderId: &clientOrderID1,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp1.Data.OrderId)
	orderIDs = append(orderIDs, *resp1.Data.OrderId)
	clientOrderIDs = append(clientOrderIDs, clientOrderID1)
	t.Logf("Created order 1: ID=%s, ClientID=%s, Price=%s", *resp1.Data.OrderId, clientOrderID1, price1.String())

	time.Sleep(500 * time.Millisecond)

	// Create Order 2: Sell at +3%
	price2 := lastPrice.Mul(decimal.NewFromFloat(1.03)).Div(tickSize).Ceil().Mul(tickSize)
	clientOrderID2 := fmt.Sprintf("sdk-query-test-2-%d", time.Now().UnixNano())
	resp2, err := client.CreateOrder(ctx, &order.CreateOrderParams{
		ContractId:    contractID,
		Price:         price2.String(),
		Size:          orderSize,
		Type:          order.OrderTypeLimit,
		Side:          order.OrderSideSell,
		TimeInForce:   string(order.TimeInForce_GOOD_TIL_CANCEL),
		ClientOrderId: &clientOrderID2,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp2.Data.OrderId)
	orderIDs = append(orderIDs, *resp2.Data.OrderId)
	clientOrderIDs = append(clientOrderIDs, clientOrderID2)
	t.Logf("Created order 2: ID=%s, ClientID=%s, Price=%s", *resp2.Data.OrderId, clientOrderID2, price2.String())

	time.Sleep(500 * time.Millisecond)

	// Step 3: Query active orders
	t.Log("Step 3: Querying active orders...")
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
	activeParams := order.GetActiveOrderParams{
		PaginationParams: order.PaginationParams{
			Size: "1",
		},
		OrderFilterParams: order.OrderFilterParams{
			FilterContractIdList: []string{contractID},
		},
	}
	activeParamsJSON, _ := json.MarshalIndent(activeParams, "", "  ")
	activeRespJSON, _ := json.MarshalIndent(activeResp, "", "  ")
	t.Logf("GetActiveOrders params: %s", string(activeParamsJSON))
	t.Logf("GetActiveOrders response: %s", string(activeRespJSON))
	t.Logf("Total active orders for contract %s: %d", contractID, len(activeResp.Data.DataList))

	// Verify our orders are in the list
	foundCount := 0
	for _, o := range activeResp.Data.DataList {
		if o.Id != nil {
			for _, testOrderID := range orderIDs {
				if *o.Id == testOrderID {
					foundCount++
					if o.Status != nil {
						t.Logf("Found order %s: Status=%s", *o.Id, *o.Status)
					}
					break
				}
			}
		}
	}
	assert.GreaterOrEqual(t, foundCount, 2, "Should find at least 2 test orders")

	// Step 4: Query by client order ID
	t.Log("Step 4: Querying orders by client order ID...")
	clientOrderResp, err := client.GetOrdersByClientOrderID(ctx, clientOrderIDs)
	assert.NoError(t, err)
	assert.NotNil(t, clientOrderResp)
	assert.GreaterOrEqual(t, len(clientOrderResp.Data), 2, "Should find at least 2 orders by client ID")

	for i, o := range clientOrderResp.Data {
		if o.ClientOrderId != nil && o.Status != nil {
			t.Logf("Order %d by ClientID: ClientOrderID=%s, Status=%s", i+1, *o.ClientOrderId, *o.Status)
		}
	}

	// Step 5: Query by order ID
	t.Log("Step 5: Querying orders by order ID...")
	orderIDResp, err := client.GetOrdersByID(ctx, orderIDs)
	assert.NoError(t, err)
	assert.NotNil(t, orderIDResp)
	assert.Equal(t, len(orderIDs), len(orderIDResp.Data), "Should find all orders by ID")

	for i, o := range orderIDResp.Data {
		if o.Id != nil && o.Price != nil && o.Size != nil {
			t.Logf("Order %d details: ID=%s, Price=%s, Size=%s", i+1, *o.Id, *o.Price, *o.Size)
		}
	}

	// Step 6: Query order fill transactions (may be empty for limit orders that haven't filled)
	t.Log("Step 6: Querying order fill transactions...")
	fillResp, err := client.GetOrderFillTransactions(ctx, &order.OrderFillTransactionParams{
		PaginationParams: order.PaginationParams{
			Size: "1",
		},
		OrderFilterParams: order.OrderFilterParams{
			FilterContractIdList: []string{contractID},
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, fillResp)
	fillParams := order.OrderFillTransactionParams{
		PaginationParams: order.PaginationParams{
			Size: "1",
		},
		OrderFilterParams: order.OrderFilterParams{
			FilterContractIdList: []string{contractID},
		},
	}
	fillParamsJSON, _ := json.MarshalIndent(fillParams, "", "  ")
	fillRespJSON, _ := json.MarshalIndent(fillResp, "", "  ")
	t.Logf("GetOrderFillTransactions params: %s", string(fillParamsJSON))
	t.Logf("GetOrderFillTransactions response: %s", string(fillRespJSON))
	t.Logf("Total fill transactions: %d", len(fillResp.Data.DataList))

	// Step 7: Clean up - cancel all test orders
	t.Log("Step 7: Cleaning up test orders...")
	for i, orderID := range orderIDs {
		_, err := client.CancelOrder(ctx, &order.CancelOrderParams{
			OrderId: orderID,
		})
		if err != nil {
			t.Logf("Warning: Failed to cancel order %d (%s): %v", i+1, orderID, err)
		} else {
			t.Logf("Canceled order %d: %s", i+1, orderID)
		}
		time.Sleep(200 * time.Millisecond)
	}

	// Step 8: Verify orders are canceled
	t.Log("Step 8: Verifying orders are canceled...")
	time.Sleep(1 * time.Second)

	verifyResp, err := client.GetOrdersByID(ctx, orderIDs)
	assert.NoError(t, err)
	if verifyResp != nil {
		for _, o := range verifyResp.Data {
			if o.Status != nil {
				status := strings.ToUpper(*o.Status)
				assert.Contains(t, []string{"CANCELLED", "CANCELING", "CANCELED"}, status,
					"Order should be cancelled")
			}
		}
	}

	t.Log("✅ Order query and filtering test completed successfully")
}

// TestIntegration_OrderFiltering tests advanced filtering capabilities
func TestIntegration_OrderFiltering(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	// Test filtering by different parameters
	t.Log("Testing order filtering capabilities...")

	// Test 1: Filter by contract
	t.Log("Test 1: Filtering active orders by contract...")
	contractID, err := test.ResolveTestContractID(ctx, client)
	assert.NoError(t, err)

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
	activeFilterParamsJSON, _ := json.MarshalIndent(order.GetActiveOrderParams{
		PaginationParams: order.PaginationParams{
			Size: "1",
		},
		OrderFilterParams: order.OrderFilterParams{
			FilterContractIdList: []string{contractID},
		},
	}, "", "  ")
	activeFilterRespJSON, _ := json.MarshalIndent(activeResp, "", "  ")
	t.Logf("Filtered GetActiveOrders params: %s", string(activeFilterParamsJSON))
	t.Logf("Filtered GetActiveOrders response: %s", string(activeFilterRespJSON))
	t.Logf("Active orders for contract %s: %d", contractID, len(activeResp.Data.DataList))

	// Test 2: Pagination
	t.Log("Test 2: Testing pagination...")
	page1, err := client.GetActiveOrders(ctx, &order.GetActiveOrderParams{
		PaginationParams: order.PaginationParams{
			Size: "1",
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, page1)
	page1ParamsJSON, _ := json.MarshalIndent(order.GetActiveOrderParams{
		PaginationParams: order.PaginationParams{
			Size: "1",
		},
	}, "", "  ")
	page1RespJSON, _ := json.MarshalIndent(page1, "", "  ")
	t.Logf("GetActiveOrders page1 params: %s", string(page1ParamsJSON))
	t.Logf("GetActiveOrders page1 response: %s", string(page1RespJSON))
	t.Logf("Page 1 orders: %d", len(page1.Data.DataList))

	if page1.Data.NextPageOffsetData != nil && *page1.Data.NextPageOffsetData != "" {
		page2, err := client.GetActiveOrders(ctx, &order.GetActiveOrderParams{
			PaginationParams: order.PaginationParams{
				Size:       "1",
				OffsetData: *page1.Data.NextPageOffsetData,
			},
		})
		assert.NoError(t, err)
		assert.NotNil(t, page2)
		page2ParamsJSON, _ := json.MarshalIndent(order.GetActiveOrderParams{
			PaginationParams: order.PaginationParams{
				Size:       "1",
				OffsetData: *page1.Data.NextPageOffsetData,
			},
		}, "", "  ")
		page2RespJSON, _ := json.MarshalIndent(page2, "", "  ")
		t.Logf("GetActiveOrders page2 params: %s", string(page2ParamsJSON))
		t.Logf("GetActiveOrders page2 response: %s", string(page2RespJSON))
		t.Logf("Page 2 orders: %d", len(page2.Data.DataList))
	}

	t.Log("✅ Order filtering test completed successfully")
}
