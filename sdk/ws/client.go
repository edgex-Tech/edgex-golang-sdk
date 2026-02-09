package ws

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/edgex-Tech/edgex-golang-sdk/sdk/internal"
	"github.com/edgex-Tech/edgex-golang-sdk/starkcurve"
	"github.com/gorilla/websocket"
	"golang.org/x/crypto/sha3"
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
	signingMethod     string
	starkPriKey       string
	apiKey            string
	apiPassphrase     string
	apiSecret         string
	authHeaderKey     string
	requestURI        string
	nowFn             func() time.Time
}

// PrivateAuthConfig defines private websocket auth behavior.
type PrivateAuthConfig struct {
	SigningMethod string
	StarkPriKey   string
	APIKey        string
	APIPassphrase string
	APISecret     string
	AuthHeaderKey string
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

// NewClient creates a new WebSocket client
func NewClient(wsURL string, isPrivate bool, accountID int64, starkPriKey string) *Client {
	return NewClientWithConfig(wsURL, isPrivate, accountID, &PrivateAuthConfig{
		SigningMethod: internal.SigningMethodStark,
		StarkPriKey:   starkPriKey,
	})
}

// NewClientWithConfig creates a new websocket client with configurable private auth.
func NewClientWithConfig(wsURL string, isPrivate bool, accountID int64, cfg *PrivateAuthConfig) *Client {
	if cfg == nil {
		cfg = &PrivateAuthConfig{}
	}

	signingMethod := strings.ToLower(strings.TrimSpace(cfg.SigningMethod))
	if signingMethod == "" {
		if strings.TrimSpace(cfg.APIKey) != "" ||
			strings.TrimSpace(cfg.APIPassphrase) != "" ||
			strings.TrimSpace(cfg.APISecret) != "" {
			signingMethod = internal.SigningMethodHMAC
		} else {
			signingMethod = internal.SigningMethodStark
		}
	}

	authHeaderKey := strings.TrimSpace(cfg.AuthHeaderKey)
	if authHeaderKey == "" {
		authHeaderKey = internal.DefaultHeaderKey
	}

	requestURI := strings.TrimSpace(cfg.RequestURI)
	if requestURI == "" {
		if parsed, err := url.Parse(wsURL); err == nil {
			requestURI = parsed.Path
		}
	}
	if requestURI == "" {
		requestURI = "/api/v1/private/ws"
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
		signingMethod: signingMethod,
		starkPriKey:   strings.TrimPrefix(strings.TrimSpace(cfg.StarkPriKey), "0x"),
		apiKey:        strings.TrimSpace(cfg.APIKey),
		apiPassphrase: strings.TrimSpace(cfg.APIPassphrase),
		apiSecret:     cfg.APISecret,
		authHeaderKey: authHeaderKey,
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
		if c.signingMethod == internal.SigningMethodHMAC {
			dialURL, err = c.configureHMACPrivateConnection(&dialer, headers)
		} else {
			err = c.configureStarkPrivateHeaders(headers)
		}
		if err != nil {
			return err
		}
	}

	conn, _, err := dialer.DialContext(ctx, dialURL, headers)
	if err != nil {
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

func (c *Client) configureStarkPrivateHeaders(headers http.Header) error {
	timestamp := c.nowFn().UnixMilli()
	headers.Set("X-edgeX-Api-Timestamp", fmt.Sprintf("%d", timestamp))

	// Keep legacy v1 signing content format for backward compatibility.
	path := fmt.Sprintf("/api/v1/private/wsaccountId=%d", c.accountID)
	signContent := fmt.Sprintf("%d%s%s", timestamp, http.MethodGet, path)

	hash := sha3.NewLegacyKeccak256()
	hash.Write([]byte(signContent))
	messageHash := hash.Sum(nil)

	privKeyBytes, err := hex.DecodeString(c.starkPriKey)
	if err != nil {
		return fmt.Errorf("failed to decode private key: %w", err)
	}

	starkPrivKey := big.NewInt(0).SetBytes(privKeyBytes)
	msgHashInt := big.NewInt(0).SetBytes(messageHash)
	msgHashInt = msgHashInt.Mod(msgHashInt, starkcurve.NewStarkCurve().N)

	r, s, err := starkcurve.Sign(starkPrivKey.Bytes(), msgHashInt.Bytes())
	if err != nil {
		return fmt.Errorf("failed to sign message: %w", err)
	}

	rBytes := append(bytes.Repeat([]byte{0}, 32-len(r.Bytes())), r.Bytes()...)
	sBytes := append(bytes.Repeat([]byte{0}, 32-len(s.Bytes())), s.Bytes()...)

	headers.Set("X-edgeX-Api-Signature", fmt.Sprintf("%s%s", hex.EncodeToString(rBytes), hex.EncodeToString(sBytes)))
	return nil
}

func (c *Client) configureHMACPrivateConnection(dialer *websocket.Dialer, headers http.Header) (string, error) {
	if c.apiKey == "" || c.apiPassphrase == "" || c.apiSecret == "" {
		return "", fmt.Errorf("hmac credentials are required for private websocket")
	}

	timestamp := fmt.Sprintf("%d", c.nowFn().UnixMilli())
	requestBody := fmt.Sprintf("accountId=%d&timestamp=%s", c.accountID, timestamp)
	signature := buildHMACSignature(c.apiSecret, timestamp, http.MethodGet, c.requestURI, requestBody)

	signHeaders := map[string]string{
		fmt.Sprintf("X-%s-Api-Key", c.authHeaderKey):    c.apiKey,
		fmt.Sprintf("X-%s-Passphrase", c.authHeaderKey): c.apiPassphrase,
		fmt.Sprintf("X-%s-Signature", c.authHeaderKey):  signature,
		fmt.Sprintf("X-%s-Timestamp", c.authHeaderKey):  timestamp,
	}

	protocol, err := buildWSAuthProtocol(signHeaders)
	if err != nil {
		return "", err
	}
	dialer.Subprotocols = []string{protocol}

	// Add headers as well for non-browser clients (web uses subprotocol only).
	for key, value := range signHeaders {
		headers.Set(key, value)
	}

	parsedURL, err := url.Parse(c.url)
	if err != nil {
		return "", fmt.Errorf("failed to parse websocket url: %w", err)
	}

	query := parsedURL.Query()
	query.Set("accountId", fmt.Sprintf("%d", c.accountID))
	query.Set("timestamp", timestamp)
	parsedURL.RawQuery = query.Encode()

	return parsedURL.String(), nil
}

func buildWSAuthProtocol(headers map[string]string) (string, error) {
	headerBytes, err := json.Marshal(headers)
	if err != nil {
		return "", fmt.Errorf("failed to marshal websocket auth headers: %w", err)
	}
	return strings.TrimRight(base64.StdEncoding.EncodeToString(headerBytes), "="), nil
}

func buildHMACSignature(secret string, timestamp string, method string, requestURI string, requestBody string) string {
	message := timestamp + method + requestURI + requestBody
	// Match edgex-web behavior: base64(encodeURI(secret)); encodeURI is a no-op for typical API secrets.
	base64Key := base64.StdEncoding.EncodeToString([]byte(secret))

	mac := hmac.New(sha256.New, []byte(base64Key))
	mac.Write([]byte(message))
	return hex.EncodeToString(mac.Sum(nil))
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
