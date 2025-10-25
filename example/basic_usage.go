package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/edgex-Tech/edgex-golang-sdk/sdk"
	"github.com/edgex-Tech/edgex-golang-sdk/sdk/asset"
	"github.com/edgex-Tech/edgex-golang-sdk/sdk/order"
	"github.com/edgex-Tech/edgex-golang-sdk/sdk/quote"
	"github.com/edgex-Tech/edgex-golang-sdk/sdk/transfer"
)

// printJSON prints the object as formatted JSON
func printJSON(label string, data interface{}) {
	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("%s: Error marshaling JSON: %v\n", label, err)
		return
	}
	fmt.Printf("%s:\n%s\n\n", label, string(jsonBytes))
}

func main() {
	// Load configuration from environment variables
	baseURL := os.Getenv("EDGEX_BASE_URL")
	if baseURL == "" {
		baseURL = "https://testnet.edgex.exchange"
	}

	accountIDStr := os.Getenv("EDGEX_ACCOUNT_ID")
	if accountIDStr == "" {
		accountIDStr = "675511695258419841"
	}
	accountID, err := strconv.ParseInt(accountIDStr, 10, 64)
	if err != nil {
		log.Fatalf("Invalid account ID: %v", err)
	}

	starkPrivateKey := os.Getenv("EDGEX_STARK_PRIVATE_KEY")
	if starkPrivateKey == "" {
		starkPrivateKey = "05495e45ffcda8224eeac73e279b804308c6dd93ea950f142baa7165968d90cb"
	}

	ctx := context.Background()
	// Create a new client
	client, err := sdk.NewClient(&sdk.ClientConfig{
		BaseURL:     baseURL,
		AccountID:   accountID,
		StarkPriKey: starkPrivateKey,
	})
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Get account assets
	result, err := client.Account.GetAccountAsset(ctx)
	if err != nil {
		log.Fatalf("Failed to get account asset: %v", err)
	}
	printJSON("GetAccountAsset Response", result)

	deleverage, err := client.GetAccountDeleverageLight(ctx)
	if err != nil {
		log.Fatalf("Failed to get account deleverage light: %v", err)
	}
	printJSON("GetAccountDeleverageLight Response", deleverage)

	// Get exchange metadata
	metadata, err := client.GetMetaData(ctx)
	if err != nil {
		log.Fatalf("Failed to get metadata: %v", err)
	}
	printJSON("Available contracts", metadata.Data)

	// Get account positions
	positions, err := client.Account.GetAccountPositions(ctx)
	if err != nil {
		log.Fatalf("Failed to get account positions: %v", err)
	}
	printJSON("Account Positions:", positions)

	// Get 24-hour market data for BNBUSDT (contract ID: 10000004)
	quoteData, err := client.Get24HourQuote(ctx, "10000004")
	if err != nil {
		log.Fatalf("Failed to get 24-hour quote: %v", err)
	}
	printJSON("BNBUSDT Price:", quoteData)

	// Get K-line data for BTCUSDT (contract ID: 10000001)
	klineParams := quote.GetKLineParams{
		ContractID: "10000001", // BTCUSDT
		Interval:   quote.KlineType1Hour,
		PriceType:  quote.PriceTypeLastPrice,
		Size:       10,
	}
	klines, err := client.GetKLine(ctx, klineParams)
	if err != nil {
		log.Fatalf("Failed to get K-lines: %v", err)
	}
	printJSON("K-lines:", klines)

	// Get order book depth for ETHUSDT (contract ID: 10000002)
	depthParams := quote.GetOrderBookDepthParams{
		ContractID: "10000002", // ETHUSDT
		Size:       15,         // Valid values are 15 or 200
	}
	depth, err := client.Quote.GetOrderBookDepth(ctx, depthParams)
	if err != nil {
		log.Fatalf("Failed to get order book depth: %v", err)
	}
	printJSON("Order Book Depth:", depth)

	transferResult, err := client.CreateTransferOut(ctx, &transfer.CreateTransferOutParams{
		CoinId:            "1000",
		Amount:            "10",
		ReceiverAccountId: "675524849547870757",
		ReceiverL2Key:     "0x0711bcc79aecf8533e94d9041d02159d45d239fa78f6bc2b1f2efede31e321b9",
		TransferReason:    transfer.USER_TRANSFER.String(),
		ExpireTime:        time.Now().Add(time.Hour * 24),
	})
	if err != nil {
		log.Fatalf("Failed to create transfer out: %v", err)
	}
	printJSON("Transfer out result:", transferResult)

	// Create a limit order
	orderParams := &order.CreateOrderParams{
		ContractId: "10000004", // BNBUSDT
		Size:       "0.01",
		Price:      "600.00",
		Side:       order.OrderSideBuy,
		Type:       order.OrderTypeLimit,
		ExpireTime: time.Now().Add(time.Hour * 24),
	}
	orderResult, err := client.CreateOrder(ctx, orderParams)
	if err != nil {
		log.Fatalf("Failed to create order: %v", err)
	}
	printJSON("Order created:", orderResult)

	// Create a market order
	marketOrderParams := &order.CreateOrderParams{
		ContractId: "10000004", // BNBUSDT
		Size:       "0.01",
		Price:      "0",
		Side:       order.OrderSideBuy,
		Type:       order.OrderTypeMarket,
		ExpireTime: time.Now().Add(time.Hour * 24),
	}
	orderResult, err = client.CreateOrder(ctx, marketOrderParams)
	if err != nil {
		log.Fatalf("Failed to create order: %v", err)
	}
	printJSON("Order created:", orderResult)

	// Get active orders
	activeOrders, err := client.GetActiveOrders(ctx, &order.GetActiveOrderParams{
		PaginationParams: order.PaginationParams{
			Size: "10",
		},
		OrderFilterParams: order.OrderFilterParams{
			FilterContractIdList: []string{"10000004"},
		},
	})
	if err != nil {
		log.Fatalf("Failed to get active orders: %v", err)
	}
	printJSON("Active Orders:", activeOrders)
	// Cancel the order
	cancelParams := &order.CancelOrderParams{
		OrderId: *orderResult.Data.OrderId,
	}
	cancelResult, err := client.CancelOrder(ctx, cancelParams)
	if err != nil {
		log.Fatalf("Failed to cancel order: %v", err)
	}
	printJSON("Order canceled:", cancelResult)

	// Create a normal withdrawal
	createWithdrawResult, err := client.CreateNormalWithdraw(ctx, &asset.CreateNormalWithdrawParams{
		CoinId:     "1000",
		Amount:     "10",
		EthAddress: "0x94C2bF0F2254eD91a5fBbc8c9F3f3433f18480D8",
	})
	if err != nil {
		log.Fatalf("Failed to create withdrawal: %v", err)
	}
	printJSON("create withdraw result:", createWithdrawResult)

	/*
		// WebSocket example
		wsURL := os.Getenv("EDGEX_WS_URL")
		if wsURL == "" {
			wsURL = "wss://quote-testnet.edgex.exchange"
		}

		// WebSocket functionality would be implemented here
		// This is a placeholder for WebSocket support
		printJSON("WebSocket URL: %s\n", wsURL)
	*/
}
