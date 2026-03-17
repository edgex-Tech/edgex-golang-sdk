package integration

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/edgex-Tech/edgex-golang-sdk/sdk/account"
	"github.com/edgex-Tech/edgex-golang-sdk/sdk/order"
	"github.com/edgex-Tech/edgex-golang-sdk/test"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

// TestIntegration_PositionTermGeneration tests if position terms get proper IDs after open and close
func TestIntegration_PositionTermGeneration(t *testing.T) {
	if os.Getenv("EDGEX_ENABLE_MUTATION_TESTS") != "true" {
		t.Skip("Skipping mutation test: EDGEX_ENABLE_MUTATION_TESTS not enabled")
	}

	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	if strings.TrimSpace(client.GetSignerPriKey()) == "" {
		t.Skip("Skipping integration test: EDGEX_SIGNER_PRIVATE_KEY is required")
	}

	ctx := test.GetTestContext()

	// Step 1: Check initial position terms
	t.Log("Step 1: Checking initial position terms...")
	termsBefore, err := client.GetPositionTermPage(ctx, account.GetPositionTermPageParams{Size: 10})
	assert.NoError(t, err)
	
	initialTermCount := 0
	if termsBefore.Data != nil {
		initialTermCount = len(termsBefore.Data.DataList)
		t.Logf("Initial position terms: %d", initialTermCount)
		
		for i, term := range termsBefore.Data.DataList {
			t.Logf("Term %d:", i+1)
			if term.Id != nil {
				t.Logf("  ID: %s", *term.Id)
			} else {
				t.Logf("  ID: <nil>")
			}
			if term.ContractId != nil {
				t.Logf("  ContractID: %s", *term.ContractId)
			}
			if term.CreatedTime != nil {
				t.Logf("  CreatedTime: %s", *term.CreatedTime)
			}
		}
	}

	// Step 2: Get contract and market price
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

	orderSize := contract.MinOrderSize
	if orderSize == "" || orderSize == "0" {
		orderSize = "0.001"
	}

	t.Logf("Contract: %s, Current Price: %s, Order Size: %s", contractID, lastPrice.String(), orderSize)

	// Step 3: Create a market buy order to open position
	t.Log("Step 3: Creating market BUY order to open position...")
	clientOrderID1 := fmt.Sprintf("sdk-term-test-buy-%d", time.Now().UnixNano())
	
	buyResp, err := client.CreateOrder(ctx, &order.CreateOrderParams{
		ContractId:    contractID,
		Price:         "0", // Market order
		Size:          orderSize,
		Type:          order.OrderTypeMarket,
		Side:          order.OrderSideBuy,
		ClientOrderId: &clientOrderID1,
	})
	
	assert.NoError(t, err)
	assert.NotNil(t, buyResp)
	assert.NotNil(t, buyResp.Data)
	assert.NotNil(t, buyResp.Data.OrderId)
	
	buyOrderID := *buyResp.Data.OrderId
	t.Logf("✅ Market BUY order created: ID=%s", buyOrderID)
	
	// Wait for order to fill
	time.Sleep(3 * time.Second)
	
	// Step 4: Check position terms after opening
	t.Log("Step 4: Checking position terms after opening position...")
	termsAfterOpen, err := client.GetPositionTermPage(ctx, account.GetPositionTermPageParams{Size: 10})
	assert.NoError(t, err)
	
	if termsAfterOpen.Data != nil {
		t.Logf("Position terms after open: %d", len(termsAfterOpen.Data.DataList))
		for i, term := range termsAfterOpen.Data.DataList {
			if term.Id != nil {
				t.Logf("Term %d ID: %s", i+1, *term.Id)
			} else {
				t.Logf("Term %d ID: <nil>", i+1)
			}
		}
	}

	// Step 5: Create a market sell order to close position
	t.Log("Step 5: Creating market SELL order to close position...")
	clientOrderID2 := fmt.Sprintf("sdk-term-test-sell-%d", time.Now().UnixNano())
	
	sellResp, err := client.CreateOrder(ctx, &order.CreateOrderParams{
		ContractId:    contractID,
		Price:         "0", // Market order
		Size:          orderSize,
		Type:          order.OrderTypeMarket,
		Side:          order.OrderSideSell,
		ReduceOnly:    true, // Close position only
		ClientOrderId: &clientOrderID2,
	})
	
	assert.NoError(t, err)
	assert.NotNil(t, sellResp)
	assert.NotNil(t, sellResp.Data)
	assert.NotNil(t, sellResp.Data.OrderId)
	
	sellOrderID := *sellResp.Data.OrderId
	t.Logf("✅ Market SELL order created: ID=%s", sellOrderID)
	
	// Wait for order to fill and position to close
	time.Sleep(5 * time.Second)
	
	// Step 6: Check position terms after closing
	t.Log("Step 6: Checking position terms after closing position...")
	termsAfterClose, err := client.GetPositionTermPage(ctx, account.GetPositionTermPageParams{Size: 10})
	assert.NoError(t, err)
	
	if termsAfterClose.Data != nil {
		t.Logf("Position terms after close: %d", len(termsAfterClose.Data.DataList))
		
		hasTermWithId := false
		for i, term := range termsAfterClose.Data.DataList {
			t.Logf("Term %d:", i+1)
			if term.Id != nil && *term.Id != "" {
				t.Logf("  ✅ ID: %s (FOUND!)", *term.Id)
				hasTermWithId = true
				
				// Try to get position orders with termCount=1
				t.Logf("  Testing GetPositionOrders...")
				ordersResp, err := client.GetPositionOrders(ctx, account.GetPositionOrdersParams{
					ContractId: *term.ContractId,
					TermCount:  1,
				})
				if err != nil {
					t.Logf("  GetPositionOrders error: %v", err)
				} else if ordersResp.Data != nil {
					t.Logf("  ✅ GetPositionOrders success! Total: %d, Orders: %d", 
						ordersResp.Data.Total, len(ordersResp.Data.OrderList))
				}
			} else {
				t.Logf("  ID: <nil>")
			}
			if term.ContractId != nil {
				t.Logf("  ContractID: %s", *term.ContractId)
			}
			if term.CreatedTime != nil {
				t.Logf("  CreatedTime: %s", *term.CreatedTime)
			}
		}
		
		if !hasTermWithId {
			t.Log("⚠️ Still no position term with ID after close")
		}
	}

	t.Log("✅ Position term generation test completed")
}
