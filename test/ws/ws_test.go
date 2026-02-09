package ws_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/edgex-Tech/edgex-golang-sdk/sdk/ws"
	"github.com/edgex-Tech/edgex-golang-sdk/test"
)

type metadataContract struct {
	ContractID     string `json:"contractId"`
	IsOpenPosition bool   `json:"isOpenPosition"`
}

type metadataData struct {
	ContractList []metadataContract `json:"contractList"`
}

type metadataResponse struct {
	Code string       `json:"code"`
	Data metadataData `json:"data"`
}

func fetchContractIDFromMetadata(ctx context.Context, baseURL string) (string, error) {
	baseURL = strings.TrimSpace(strings.TrimRight(baseURL, "/"))
	if baseURL == "" {
		return "", fmt.Errorf("base url is empty")
	}

	client := &http.Client{Timeout: 10 * time.Second}
	candidates := []string{
		baseURL + "/api/v2/public/meta/getMetaData",
		baseURL + "/api/v1/public/meta/getMetaData",
	}

	for _, url := range candidates {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			continue
		}

		resp, err := client.Do(req)
		if err != nil {
			continue
		}

		body, readErr := io.ReadAll(resp.Body)
		resp.Body.Close()
		if readErr != nil || resp.StatusCode != http.StatusOK {
			continue
		}

		var metadataResp metadataResponse
		if err := json.Unmarshal(body, &metadataResp); err != nil {
			continue
		}
		if metadataResp.Code != "SUCCESS" {
			continue
		}

		for _, c := range metadataResp.Data.ContractList {
			if c.IsOpenPosition && strings.TrimSpace(c.ContractID) != "" {
				return c.ContractID, nil
			}
		}
		for _, c := range metadataResp.Data.ContractList {
			if strings.TrimSpace(c.ContractID) != "" {
				return c.ContractID, nil
			}
		}
	}

	return "", fmt.Errorf("no valid contract id found from metadata")
}

func TestWebSocket(t *testing.T) {
	ctx := test.GetTestContext()
	baseURL := os.Getenv("TEST_WS_BASE_URL")
	if baseURL == "" {
		t.Skip("Skipping test: TEST_WS_BASE_URL environment variable is not set")
	}

	manager := ws.NewManager(baseURL, 0, "") // No auth needed for public WebSocket

	// Connect to public WebSocket
	err := manager.ConnectPublic(ctx)
	if err != nil {
		t.Fatalf("Failed to connect to public WebSocket: %v", err)
	}

	metadataBaseURL := os.Getenv("TEST_BASE_URL")
	contractID, err := fetchContractIDFromMetadata(ctx, metadataBaseURL)
	if err != nil {
		t.Fatalf("Failed to resolve contract ID from metadata: %v", err)
	}
	t.Logf("Using contract ID from metadata: %s", contractID)

	// Create channels to track message receipt
	tickerMsgCh := make(chan struct{})
	klineMsgCh := make(chan struct{})
	depthMsgCh := make(chan struct{})
	tradesMsgCh := make(chan struct{})

	// Add a debug hook to log all messages
	manager.OnPublicMessage(func(message []byte) {
		t.Logf("Raw message received: %s", string(message))
	})

	// Test cases for different subscription types
	var tickerReceived, klineReceived, depthReceived, tradesReceived bool
	testCases := []struct {
		name    string
		subFunc func() error
		msgCh   chan struct{}
	}{
		{
			name: "Market Ticker",
			subFunc: func() error {
				return manager.SubscribeMarketTicker(contractID, func(message []byte) {
					t.Logf("Ticker message received: %s", string(message))
					if !tickerReceived {
						close(tickerMsgCh)
						tickerReceived = true
					}
				})
			},
			msgCh: tickerMsgCh,
		},
		{
			name: "KLine",
			subFunc: func() error {
				return manager.SubscribeKLine(contractID, "DAY_1", func(message []byte) {
					t.Logf("KLine message received: %s", string(message))
					if !klineReceived {
						close(klineMsgCh)
						klineReceived = true
					}
				})
			},
			msgCh: klineMsgCh,
		},
		{
			name: "Depth",
			subFunc: func() error {
				return manager.SubscribeDepth(contractID, func(message []byte) {
					t.Logf("Depth message received: %s", string(message))
					if !depthReceived {
						close(depthMsgCh)
						depthReceived = true
					}
				})
			},
			msgCh: depthMsgCh,
		},
		{
			name: "Trades",
			subFunc: func() error {
				return manager.SubscribeTrades(contractID, func(message []byte) {
					t.Logf("Trades message received: %s", string(message))
					if !tradesReceived {
						close(tradesMsgCh)
						tradesReceived = true
					}
				})
			},
			msgCh: tradesMsgCh,
		},
	}

	// Run each test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Subscribe to the channel
			err := tc.subFunc()
			if err != nil {
				t.Fatalf("Failed to subscribe to %s: %v", tc.name, err)
			}

			// Wait for message or timeout
			select {
			case <-tc.msgCh:
				t.Logf("%s message received successfully", tc.name)
			case <-time.After(5 * time.Second):
				t.Errorf("Timeout waiting for %s message", tc.name)
			}
		})
	}

	// Clean up
	manager.Close()
}

func TestPrivateWebSocket(t *testing.T) {
	// Get test credentials from environment variables
	baseURL := os.Getenv("TEST_WS_BASE_URL")
	if baseURL == "" {
		t.Skip("Skipping test: TEST_WS_BASE_URL environment variable is not set")
	}

	accountIDStr := os.Getenv("TEST_ACCOUNT_ID")
	if accountIDStr == "" {
		t.Skip("Skipping test: TEST_ACCOUNT_ID environment variable is not set")
	}

	accountID, err := strconv.ParseInt(accountIDStr, 10, 64)
	if err != nil {
		t.Fatalf("Invalid TEST_ACCOUNT_ID: %v", err)
	}

	signingMethod := strings.ToLower(strings.TrimSpace(os.Getenv("TEST_WS_SIGNING_METHOD")))
	if signingMethod == "" {
		if os.Getenv("TEST_API_KEY") != "" &&
			os.Getenv("TEST_API_PASSPHRASE") != "" &&
			os.Getenv("TEST_API_SECRET") != "" {
			signingMethod = "hmac"
		} else {
			signingMethod = "stark"
		}
	}

	ctx := test.GetTestContext()
	var manager *ws.Manager
	if signingMethod == "hmac" {
		apiKey := strings.TrimSpace(os.Getenv("TEST_API_KEY"))
		apiPassphrase := strings.TrimSpace(os.Getenv("TEST_API_PASSPHRASE"))
		apiSecret := os.Getenv("TEST_API_SECRET")
		if apiKey == "" || apiPassphrase == "" || apiSecret == "" {
			t.Skip("Skipping v2 private ws test: TEST_API_KEY/TEST_API_PASSPHRASE/TEST_API_SECRET are required")
		}

		manager = ws.NewManagerWithConfig(baseURL, accountID, &ws.ManagerConfig{
			APIVersion:    "v2",
			SigningMethod: "hmac",
			APIKey:        apiKey,
			APIPassphrase: apiPassphrase,
			APISecret:     apiSecret,
			AuthHeaderKey: strings.TrimSpace(os.Getenv("TEST_AUTH_HEADER_KEY")),
		})
	} else {
		starkPrivateKey := strings.TrimSpace(os.Getenv("TEST_STARK_PRIVATE_KEY"))
		if starkPrivateKey == "" {
			t.Skip("Skipping v1 private ws test: TEST_STARK_PRIVATE_KEY environment variable is not set")
		}

		manager = ws.NewManager(baseURL, accountID, starkPrivateKey)
	}

	// Connect to private WebSocket
	err = manager.ConnectPrivate(ctx)
	if err != nil {
		t.Logf("Failed to connect to private WebSocket: %v", err)
	}

	// Listen for account updates
	done := make(chan struct{})
	err = manager.OnPrivateMessage("ACCOUNT_UPDATE", func(message []byte) {
		var msg map[string]interface{}
		err := json.Unmarshal(message, &msg)
		if err != nil {
			t.Logf("Failed to unmarshal message: %v", err)
			return
		}
		t.Logf("Received account update: %v", msg)
		close(done)
	})
	if err != nil {
		t.Fatalf("Failed to register account update handler: %v", err)
	}

	// Wait for message or timeout
	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Log("Timeout waiting for account update")
	}

	// Clean up
	manager.Close()
}
