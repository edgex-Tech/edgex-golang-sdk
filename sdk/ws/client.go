package ws

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Client represents a WebSocket client
type Client struct {
	conn              *websocket.Conn
	url               string
	mu                sync.RWMutex
	handlers          map[string]MessageHandler
	done              chan struct{}
	pingTicker        *time.Ticker
	isPrivate         bool
	subscriptions     map[string]struct{}
	onConnectHooks    []func()
	onMessageHooks    []func([]byte)
	onDisconnectHooks []func(error)
	accountID         int64
	apiKey            string
	apiPassphrase     string
	apiSecret         string
	requestURI        string
	nowFn             func() time.Time
}

// PrivateAuthConfig defines private websocket auth configuration.
type PrivateAuthConfig struct {
	APIKey        string
	APIPassphrase string
	APISecret     string
	RequestURI    string
}

// MessageHandler is a function type for handling WebSocket messages
type MessageHandler func(message []byte)

// Message represents a WebSocket message
type Message struct {
	Type string          `json:"type"`
	Time string          `json:"time,omitempty"`
	Data json.RawMessage `json:"data,omitempty"`
}

// NewClient creates a new WebSocket client with HMAC authentication
func NewClient(wsURL string, isPrivate bool, accountID int64, apiKey, apiPassphrase, apiSecret string) *Client {
	return NewClientWithConfig(wsURL, isPrivate, accountID, &PrivateAuthConfig{
		APIKey:        apiKey,
		APIPassphrase: apiPassphrase,
		APISecret:     apiSecret,
	})
}

// NewClientWithConfig creates a new websocket client with configurable private auth.
func NewClientWithConfig(wsURL string, isPrivate bool, accountID int64, cfg *PrivateAuthConfig) *Client {
	if cfg == nil {
		cfg = &PrivateAuthConfig{}
	}

	requestURI := strings.TrimSpace(cfg.RequestURI)
	if requestURI == "" {
		if parsed, err := url.Parse(wsURL); err == nil {
			requestURI = parsed.Path
		}
	}
	if requestURI == "" {
		requestURI = "/api/v1/private/ws" // WebSocket still uses v1 per API docs
	}
	if !strings.HasPrefix(requestURI, "/") {
		requestURI = "/" + requestURI
	}

	return &Client{
		url:           wsURL,
		handlers:      make(map[string]MessageHandler),
		done:          make(chan struct{}),
		isPrivate:     isPrivate,
		subscriptions: make(map[string]struct{}),
		accountID:     accountID,
		apiKey:        strings.TrimSpace(cfg.APIKey),
		apiPassphrase: strings.TrimSpace(cfg.APIPassphrase),
		apiSecret:     cfg.APISecret,
		requestURI:    requestURI,
		nowFn:         time.Now,
	}
}

// Connect establishes a WebSocket connection
func (c *Client) Connect(ctx context.Context) error {
	dialer := websocket.Dialer{}
	headers := http.Header{}
	dialURL := c.url

	if c.isPrivate {
		var err error
		dialURL, err = c.configureHMACPrivateConnection(&dialer, headers)
		if err != nil {
			fmt.Printf("Connect err:%v", err)
			return err
		}
	}

	conn, resp, err := dialer.DialContext(ctx, dialURL, headers)
	if err != nil {
		if resp != nil {
			return fmt.Errorf("failed to connect to WebSocket: %w (status: %d)", err, resp.StatusCode)
		}
		return fmt.Errorf("failed to connect to WebSocket: %w", err)
	}

	c.mu.Lock()
	c.conn = conn
	c.mu.Unlock()

	// Start ping ticker
	c.pingTicker = time.NewTicker(30 * time.Second)

	// Start message handling
	go c.handleMessages()
	go c.handlePing()

	// Call connect hooks
	for _, hook := range c.onConnectHooks {
		hook()
	}

	return nil
}

func (c *Client) configureHMACPrivateConnection(dialer *websocket.Dialer, headers http.Header) (string, error) {
	if c.apiKey == "" || c.apiPassphrase == "" || c.apiSecret == "" {
		return "", fmt.Errorf("hmac credentials are required for private websocket")
	}

	timestamp := fmt.Sprintf("%d", c.nowFn().UnixMilli())

	// Build URL with query parameters (server expects accountId in query)
	parsedURL, err := url.Parse(c.url)
	if err != nil {
		return "", fmt.Errorf("failed to parse websocket url: %w", err)
	}

	query := parsedURL.Query()
	query.Set("accountId", fmt.Sprintf("%d", c.accountID))
	query.Set("timestamp", timestamp)

	// Convert url.Values to map[string]string for signature calculation
	params := make(map[string]string)
	for k, v := range query {
		if len(v) > 0 {
			params[k] = v[0]
		}
	}

	// Build requestBody from sorted query parameters (same as HTTP GET)
	requestBody := buildSortedQueryString(params)

	// Build signature
	signature := buildHMACSignature(c.apiSecret, timestamp, http.MethodGet, c.requestURI, requestBody)

	// Set HMAC authentication headers
	headers.Set("X-edgeX-Api-Key", c.apiKey)
	headers.Set("X-edgeX-Passphrase", c.apiPassphrase)
	headers.Set("X-edgeX-Signature", signature)
	headers.Set("X-edgeX-Timestamp", timestamp)

	parsedURL.RawQuery = query.Encode()
	return parsedURL.String(), nil
}

// buildSortedQueryString builds sorted query string (same as HTTP client)
func buildSortedQueryString(params map[string]string) string {
	if len(params) == 0 {
		return ""
	}
	keys := make([]string, 0, len(params))
	for k := range params {
		if params[k] == "" {
			continue // Skip empty values
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)
	pairs := make([]string, 0, len(keys))
	for _, k := range keys {
		pairs = append(pairs, fmt.Sprintf("%s=%s", k, params[k]))
	}
	return strings.Join(pairs, "&")
}

// buildHMACSignature builds HMAC signature (same as HTTP requests)
func buildHMACSignature(secret string, timestamp string, method string, requestURI string, requestBody string) string {
	message := timestamp + method + requestURI + requestBody
	// Same as HTTP: base64(secret)
	base64Key := base64.StdEncoding.EncodeToString([]byte(secret))

	mac := hmac.New(sha256.New, []byte(base64Key))
	mac.Write([]byte(message))
	return fmt.Sprintf("%x", mac.Sum(nil))
}

// Close closes the WebSocket connection
func (c *Client) Close() error {
	close(c.done)
	if c.pingTicker != nil {
		c.pingTicker.Stop()
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

// QuoteEvent represents a quote event message
type QuoteEvent struct {
	Type    string `json:"type"`
	Channel string `json:"channel"`
	Content struct {
		Channel  string          `json:"channel"`
		DataType string          `json:"dataType"`
		Data     json.RawMessage `json:"data"`
	} `json:"content"`
}

// handleMessages processes incoming WebSocket messages
func (c *Client) handleMessages() {
	for {
		select {
		case <-c.done:
			return
		default:
			c.mu.RLock()
			conn := c.conn
			c.mu.RUnlock()

			if conn == nil {
				return
			}

			_, message, err := conn.ReadMessage()
			if err != nil {
				for _, hook := range c.onDisconnectHooks {
					hook(err)
				}
				return
			}

			// Call message hooks
			for _, hook := range c.onMessageHooks {
				hook(message)
			}

			var msg Message
			if err := json.Unmarshal(message, &msg); err != nil {
				continue
			}

			// Handle ping messages
			if msg.Type == "ping" {
				c.handlePong(msg.Time)
				continue
			}

			// Handle quote events
			if msg.Type == "quote-event" {
				var quoteEvent QuoteEvent
				if err := json.Unmarshal(message, &quoteEvent); err != nil {
					continue
				}

				// Extract channel type from channel string (e.g., "ticker" from "ticker.10000001")
				channelType := strings.Split(quoteEvent.Channel, ".")[0]
				if handler, ok := c.handlers[channelType]; ok {
					handler(message)
				}
				continue
			}

			// Call registered handlers for other message types
			if handler, ok := c.handlers[msg.Type]; ok {
				handler(message)
			}
		}
	}
}

// handlePing sends periodic ping messages
func (c *Client) handlePing() {
	for {
		select {
		case <-c.done:
			return
		case <-c.pingTicker.C:
			c.mu.RLock()
			conn := c.conn
			c.mu.RUnlock()

			if conn == nil {
				return
			}

			pingMsg := Message{
				Type: "ping",
				Time: fmt.Sprintf("%d", time.Now().UnixMilli()),
			}

			if err := c.sendMessage(pingMsg); err != nil {
				return
			}
		}
	}
}

// handlePong sends pong response to server ping
func (c *Client) handlePong(timestamp string) {
	pongMsg := Message{
		Type: "pong",
		Time: timestamp,
	}

	_ = c.sendMessage(pongMsg)
}

// Subscribe subscribes to a topic (for public WebSocket)
func (c *Client) Subscribe(topic string, params map[string]interface{}) error {
	if c.isPrivate {
		return fmt.Errorf("cannot subscribe on private WebSocket connection")
	}

	subMsg := map[string]interface{}{
		"type":    "subscribe",
		"channel": topic,
	}

	if err := c.sendMessage(subMsg); err != nil {
		return err
	}

	c.mu.Lock()
	c.subscriptions[topic] = struct{}{}
	c.mu.Unlock()

	return nil
}

// Unsubscribe unsubscribes from a topic (for public WebSocket)
func (c *Client) Unsubscribe(topic string) error {
	if c.isPrivate {
		return fmt.Errorf("cannot unsubscribe on private WebSocket connection")
	}

	unsubMsg := map[string]interface{}{
		"type":    "unsubscribe",
		"channel": topic,
	}

	if err := c.sendMessage(unsubMsg); err != nil {
		return err
	}

	c.mu.Lock()
	delete(c.subscriptions, topic)
	c.mu.Unlock()

	return nil
}

// OnMessage registers a handler for a specific message type
func (c *Client) OnMessage(msgType string, handler MessageHandler) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.handlers[msgType] = handler
}

// OnMessageHook registers a hook that will be called for all messages
func (c *Client) OnMessageHook(hook MessageHandler) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.onMessageHooks = append(c.onMessageHooks, hook)
}

// OnConnect registers a hook that will be called when connection is established
func (c *Client) OnConnect(hook func()) {
	c.onConnectHooks = append(c.onConnectHooks, hook)
}

// OnDisconnect registers a hook that will be called when connection is closed
func (c *Client) OnDisconnect(hook func(error)) {
	c.onDisconnectHooks = append(c.onDisconnectHooks, hook)
}

// sendMessage sends a message through the WebSocket connection
func (c *Client) sendMessage(msg interface{}) error {
	c.mu.RLock()
	conn := c.conn
	c.mu.RUnlock()

	if conn == nil {
		return fmt.Errorf("WebSocket connection is not established")
	}

	return conn.WriteJSON(msg)
}
