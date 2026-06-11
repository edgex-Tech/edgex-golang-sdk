package integration

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/edgex-Tech/edgex-golang-sdk/v2/sdk/ws"
	"github.com/edgex-Tech/edgex-golang-sdk/v2/test"
	"github.com/stretchr/testify/assert"
)

// TestIntegration_WebSocketPublic tests public WebSocket functionality:
// 1. Connect to public WebSocket
// 2. Subscribe to ticker channel
// 3. Verify data reception
// 4. Close connection
func TestIntegration_WebSocketPublic(t *testing.T) {
	wsBaseURL := os.Getenv("EDGEX_WS_BASE_URL")
	if wsBaseURL == "" {
		t.Skip("Skipping WebSocket test: EDGEX_WS_BASE_URL not set")
	}

	ctx := test.GetTestContext()

	// Step 1: Create WebSocket manager
	t.Log("Step 1: Creating WebSocket manager...")
	manager := ws.NewManager(wsBaseURL, 0, "", "", "")
	assert.NotNil(t, manager)
	defer manager.Close()

	// Step 2: Connect to public WebSocket
	t.Log("Step 2: Connecting to public WebSocket...")
	err := manager.ConnectPublic(ctx)
	assert.NoError(t, err)
	t.Logf("Connected to public WebSocket: %s", wsBaseURL)
	t.Logf("Public WS connect request: baseURL=%s path=%s", wsBaseURL, "/api/v1/public/ws")

	// Wait for connection to stabilize
	time.Sleep(1 * time.Second)

	// Step 3: Get contract ID for subscription
	t.Log("Step 3: Getting contract ID...")
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	contractID, err := test.ResolveTestContractID(ctx, client)
	assert.NoError(t, err)
	t.Logf("Using contract: %s", contractID)

	// Step 4: Subscribe to ticker channel with handler
	t.Log("Step 4: Subscribing to market ticker...")
	receivedData := false
	handler := func(msg []byte) {
		receivedData = true
		t.Logf("Received ticker data: %s", string(msg))
	}

	t.Logf("Public WS subscribe request: channel=%s", fmt.Sprintf("ticker.%s", contractID))
	err = manager.SubscribeMarketTicker(contractID, handler)
	assert.NoError(t, err)
	t.Log("Subscribed to market ticker")

	// Step 5: Wait for ticker data
	t.Log("Step 5: Waiting for ticker data...")
	timeout := time.After(10 * time.Second)
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if receivedData {
				t.Log("Received ticker data successfully")
				goto cleanup
			}
		case <-timeout:
			t.Fatal("No ticker data received within 10 seconds")
		}
	}

cleanup:
	// Step 6: Close connection
	t.Log("Step 6: Closing WebSocket connection...")
	manager.Close()
	t.Log("Closed WebSocket connection")

	t.Log("✅ Public WebSocket test completed successfully")
}

// TestIntegration_WebSocketPrivate tests private WebSocket functionality:
// 1. Connect to private WebSocket
// 2. Register message handlers
// 3. Monitor for updates
// 4. Clean up
func TestIntegration_WebSocketPrivate(t *testing.T) {
	wsBaseURL := os.Getenv("EDGEX_WS_BASE_URL")
	if wsBaseURL == "" {
		t.Skip("Skipping WebSocket test: EDGEX_WS_BASE_URL not set")
	}

	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	if strings.TrimSpace(client.GetSignerPriKey()) == "" {
		t.Skip("Skipping private WebSocket test: EDGEX_SIGNER_PRIVATE_KEY is required")
	}

	ctx := test.GetTestContext()

	// Step 1: Get account ID and credentials
	t.Log("Step 1: Getting account credentials...")
	accountID := client.GetAccountID()
	apiKey := os.Getenv("EDGEX_API_KEY")
	apiSecret := os.Getenv("EDGEX_API_SECRET")
	apiPassphrase := os.Getenv("EDGEX_API_PASSPHRASE")

	// Step 2: Create WebSocket manager
	t.Log("Step 2: Creating private WebSocket manager...")
	t.Logf("WebSocket URL: %s", wsBaseURL)
	t.Logf("Account ID: %d", accountID)
	t.Logf("API Key: %s", apiKey)
	t.Logf("API Passphrase: %s", apiPassphrase)
	t.Logf("API Secret length: %d", len(apiSecret))
	t.Logf("Private WS connect request: baseURL=%s path=%s accountId=%d", wsBaseURL, "/api/v1/private/ws", accountID)

	manager := ws.NewManager(wsBaseURL, accountID, apiKey, apiPassphrase, apiSecret)
	assert.NotNil(t, manager)
	defer manager.Close()

	// Step 3: Connect to private WebSocket
	t.Log("Step 3: Connecting to private WebSocket...")
	t.Log("Expected URL: wss://.../api/v1/private/ws")
	t.Log("Authentication: HMAC via WebSocket Subprotocol")

	err = manager.ConnectPrivate(ctx)
	if err != nil {
		t.Logf("Connection error details: %v", err)
		t.Logf("Error type: %T", err)
		t.Log("Possible issues:")
		t.Log("  1. API credentials incorrect")
		t.Log("  2. Server doesn't support WebSocket Subprotocol authentication")
		t.Log("  3. Signature algorithm mismatch")
		t.Log("  4. WebSocket URL path incorrect")
	}
	assert.NoError(t, err)
	t.Logf("Connected to private WebSocket for account: %d", accountID)

	// Wait for connection to stabilize
	time.Sleep(2 * time.Second)

	// Step 4: Register handler for private messages
	t.Log("Step 4: Registering private message handlers...")
	updateCount := 0
	done := make(chan struct{}, 1)
	handler := func(msg []byte) {
		updateCount++
		t.Logf("Received private message #%d: %s", updateCount, string(msg))
		select {
		case done <- struct{}{}:
		default:
		}
	}

	// Register handlers for different message types
	t.Logf("Private WS handler registration: group=%s", "account")
	err = manager.OnPrivateMessage("account", handler)
	if err != nil {
		t.Logf("Warning: Failed to register account handler: %v", err)
	}

	t.Logf("Private WS handler registration: group=%s", "order")
	err = manager.OnPrivateMessage("order", handler)
	if err != nil {
		t.Logf("Warning: Failed to register order handler: %v", err)
	}

	// Step 5: Monitor for updates (for a short period)
	t.Log("Step 5: Monitoring for private messages...")
	select {
	case <-done:
		t.Log("Received private data successfully")
	case <-time.After(5 * time.Second):
		t.Fatal("No private WebSocket data received within 5 seconds")
	}

	// Step 6: Clean up
	t.Log("Step 6: Cleaning up...")
	manager.Close()
	t.Log("Closed private WebSocket connection")

	assert.Greater(t, updateCount, 0, "private WebSocket should receive at least one message")
	t.Logf("Private WebSocket test completed (received %d messages)", updateCount)
}

// TestIntegration_WebSocketReconnection tests WebSocket reconnection behavior
func TestIntegration_WebSocketReconnection(t *testing.T) {
	wsBaseURL := os.Getenv("EDGEX_WS_BASE_URL")
	if wsBaseURL == "" {
		t.Skip("Skipping WebSocket test: EDGEX_WS_BASE_URL not set")
	}

	ctx := test.GetTestContext()

	t.Log("Testing WebSocket reconnection behavior...")

	// Connect
	t.Log("Initial connection...")
	manager := ws.NewManager(wsBaseURL, 0, "", "", "")
	err := manager.ConnectPublic(ctx)
	assert.NoError(t, err)

	time.Sleep(1 * time.Second)

	// Close
	t.Log("Closing connection...")
	manager.Close()

	time.Sleep(500 * time.Millisecond)

	// Reconnect
	t.Log("Reconnecting...")
	manager2 := ws.NewManager(wsBaseURL, 0, "", "", "")
	err = manager2.ConnectPublic(ctx)
	assert.NoError(t, err)

	time.Sleep(1 * time.Second)

	// Final cleanup
	manager2.Close()

	t.Log("WebSocket reconnection test completed successfully")
}
