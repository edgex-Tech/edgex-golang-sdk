package ws

import (
	"context"
	"fmt"
	"strings"
	"sync"
)

// Manager handles WebSocket connections
type Manager struct {
	publicClient  *Client
	privateClient *Client
	baseURL       string
	accountID     int64
	privateWSPath string
	apiKey        string
	apiPassphrase string
	apiSecret     string
	mu            sync.RWMutex
}

// ManagerConfig defines private websocket auth configuration.
type ManagerConfig struct {
	APIKey        string
	APIPassphrase string
	APISecret     string
	PrivateWSPath string
}

// NewManager creates a new WebSocket manager with HMAC authentication
func NewManager(baseURL string, accountID int64, apiKey, apiPassphrase, apiSecret string) *Manager {
	return NewManagerWithConfig(baseURL, accountID, &ManagerConfig{
		APIKey:        apiKey,
		APIPassphrase: apiPassphrase,
		APISecret:     apiSecret,
	})
}

// NewManagerWithConfig creates a new websocket manager with configurable private auth.
func NewManagerWithConfig(baseURL string, accountID int64, cfg *ManagerConfig) *Manager {
	if cfg == nil {
		cfg = &ManagerConfig{}
	}

	privateWSPath := strings.TrimSpace(cfg.PrivateWSPath)
	if privateWSPath == "" {
		privateWSPath = "/api/v1/private/ws" // WebSocket still uses v1 per API docs
	}
	if !strings.HasPrefix(privateWSPath, "/") {
		privateWSPath = "/" + privateWSPath
	}

	return &Manager{
		baseURL:       strings.TrimRight(strings.TrimSpace(baseURL), "/"),
		accountID:     accountID,
		privateWSPath: privateWSPath,
		apiKey:        strings.TrimSpace(cfg.APIKey),
		apiPassphrase: strings.TrimSpace(cfg.APIPassphrase),
		apiSecret:     cfg.APISecret,
	}
}

// ConnectPublic connects to the public WebSocket endpoint
func (m *Manager) ConnectPublic(ctx context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.publicClient != nil {
		return nil
	}

	url := fmt.Sprintf("%s/api/v1/public/ws", m.baseURL) // WebSocket still uses v1 per API docs
	client := NewClient(url, false, 0, "", "", "")       // No auth needed for public
	if err := client.Connect(ctx); err != nil {
		return err
	}

	m.publicClient = client
	return nil
}

// ConnectPrivate connects to the private WebSocket endpoint
func (m *Manager) ConnectPrivate(ctx context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.privateClient != nil {
		return nil
	}

	url := fmt.Sprintf("%s%s", m.baseURL, m.privateWSPath)
	client := NewClientWithConfig(url, true, m.accountID, &PrivateAuthConfig{
		APIKey:        m.apiKey,
		APIPassphrase: m.apiPassphrase,
		APISecret:     m.apiSecret,
		RequestURI:    m.privateWSPath,
	})
	if err := client.Connect(ctx); err != nil {
		return err
	}

	m.privateClient = client
	return nil
}

// SubscribeMarketTicker subscribes to 24-hour market ticker updates
func (m *Manager) SubscribeMarketTicker(contractID string, handler MessageHandler) error {
	m.mu.RLock()
	client := m.publicClient
	m.mu.RUnlock()

	if client == nil {
		return fmt.Errorf("public WebSocket connection not established")
	}

	client.OnMessage("ticker", handler)
	return client.Subscribe(fmt.Sprintf("ticker.%s", contractID), nil)
}

// SubscribeKLine subscribes to K-line (candlestick) data
func (m *Manager) SubscribeKLine(contractID string, interval string, handler MessageHandler) error {
	m.mu.RLock()
	client := m.publicClient
	m.mu.RUnlock()

	if client == nil {
		return fmt.Errorf("public WebSocket connection not established")
	}

	client.OnMessage("kline", handler)
	return client.Subscribe(fmt.Sprintf("kline.LAST_PRICE.%s.%s", contractID, interval), nil)
}

// SubscribeDepth subscribes to market depth updates
func (m *Manager) SubscribeDepth(contractID string, handler MessageHandler) error {
	m.mu.RLock()
	client := m.publicClient
	m.mu.RUnlock()

	if client == nil {
		return fmt.Errorf("public WebSocket connection not established")
	}

	client.OnMessage("depth", handler)
	return client.Subscribe(fmt.Sprintf("depth.%s.15", contractID), nil)
}

// SubscribeTrades subscribes to latest trades
func (m *Manager) SubscribeTrades(contractID string, handler MessageHandler) error {
	m.mu.RLock()
	client := m.publicClient
	m.mu.RUnlock()

	if client == nil {
		return fmt.Errorf("public WebSocket connection not established")
	}

	client.OnMessage("trades", handler)
	return client.Subscribe(fmt.Sprintf("trades.%s", contractID), nil)
}

// OnPrivateMessage registers a handler for private WebSocket messages
func (m *Manager) OnPrivateMessage(msgType string, handler MessageHandler) error {
	m.mu.RLock()
	client := m.privateClient
	m.mu.RUnlock()

	if client == nil {
		return fmt.Errorf("private WebSocket connection not established")
	}

	client.OnMessage(msgType, handler)
	return nil
}

// OnPublicMessage registers a handler for all public WebSocket messages
func (m *Manager) OnPublicMessage(handler MessageHandler) error {
	m.mu.RLock()
	client := m.publicClient
	m.mu.RUnlock()

	if client == nil {
		return fmt.Errorf("public WebSocket connection not established")
	}

	client.OnMessageHook(handler)
	return nil
}

// Close closes all WebSocket connections
func (m *Manager) Close() {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.publicClient != nil {
		m.publicClient.Close()
		m.publicClient = nil
	}

	if m.privateClient != nil {
		m.privateClient.Close()
		m.privateClient = nil
	}
}
