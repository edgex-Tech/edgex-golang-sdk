package ws

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

// Recorded with an independent HMAC-SHA256 (key = base64("test-secret")) over
// "1700000000000GET/api/v1/public/wsaccountId=42&timestamp=1700000000000".
const publicIdentitySignature = "fefae46abb70c5e69968f56b9b5ac83155d12f049a28beb826dd53f6267d2615"

// handshakeRecorder upgrades every connection and keeps the handshake request.
func handshakeRecorder(t *testing.T) (string, *[]*http.Request) {
	t.Helper()
	var handshakes []*http.Request
	upgrader := websocket.Upgrader{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handshakes = append(handshakes, r)
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer conn.Close()
		for {
			if _, _, err := conn.ReadMessage(); err != nil {
				return
			}
		}
	}))
	t.Cleanup(srv.Close)
	return "ws" + strings.TrimPrefix(srv.URL, "http"), &handshakes
}

func TestPublicIdentityHandshakeCarriesTheSignedAccount(t *testing.T) {
	baseURL, handshakes := handshakeRecorder(t)

	client := NewClientWithConfig(baseURL+publicWSPath, false, 42, &PrivateAuthConfig{
		APIKey:         "test-api-key",
		APIPassphrase:  "test-passphrase",
		APISecret:      "test-secret",
		RequestURI:     publicWSPath,
		PublicIdentity: true,
	})
	client.nowFn = func() time.Time { return time.UnixMilli(1700000000000) }
	if err := client.Connect(context.Background()); err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer client.Close()

	got := (*handshakes)[0]
	if got.URL.Path != "/api/v1/public/ws" {
		t.Fatalf("unexpected path: %s", got.URL.Path)
	}
	if got.URL.RawQuery != "accountId=42&timestamp=1700000000000" {
		t.Fatalf("unexpected query: %s", got.URL.RawQuery)
	}
	want := map[string]string{
		"X-edgeX-Api-Key":    "test-api-key",
		"X-edgeX-Passphrase": "test-passphrase",
		"X-edgeX-Signature":  publicIdentitySignature,
		"X-edgeX-Timestamp":  "1700000000000",
	}
	for name, value := range want {
		if got.Header.Get(name) != value {
			t.Fatalf("header %s: got %q, want %q", name, got.Header.Get(name), value)
		}
	}

	// An identified public connection is still a public one.
	if err := client.Subscribe("depth.10000001.15", nil); err != nil {
		t.Fatalf("subscribe on identified public connection: %v", err)
	}
}

func TestManagerPublicConnectionIdentity(t *testing.T) {
	full := &ManagerConfig{APIKey: "test-api-key", APIPassphrase: "test-passphrase", APISecret: "test-secret"}
	disabled := *full
	disabled.DisableMarketDataIdentity = true

	cases := []struct {
		name       string
		accountID  int64
		cfg        *ManagerConfig
		identified bool
	}{
		{"account and full credentials", 42, full, true},
		{"identity disabled", 42, &disabled, false},
		{"no account", 0, full, false},
		{"no credentials", 42, &ManagerConfig{}, false},
		{"incomplete credentials", 42, &ManagerConfig{APIKey: "test-api-key", APIPassphrase: "test-passphrase"}, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			baseURL, handshakes := handshakeRecorder(t)
			manager := NewManagerWithConfig(baseURL, tc.accountID, tc.cfg)
			if err := manager.ConnectPublic(context.Background()); err != nil {
				t.Fatalf("connect public: %v", err)
			}
			defer manager.Close()

			got := (*handshakes)[0]
			if got.URL.Path != "/api/v1/public/ws" {
				t.Fatalf("unexpected path: %s", got.URL.Path)
			}
			if tc.identified {
				if got.URL.Query().Get("accountId") != "42" || got.URL.Query().Get("timestamp") == "" {
					t.Fatalf("identity query missing: %s", got.URL.RawQuery)
				}
				if got.Header.Get("X-edgeX-Api-Key") != "test-api-key" || got.Header.Get("X-edgeX-Signature") == "" {
					t.Fatalf("identity headers missing")
				}
				return
			}
			if got.URL.RawQuery != "" {
				t.Fatalf("anonymous handshake has a query: %s", got.URL.RawQuery)
			}
			for _, name := range []string{"X-edgeX-Api-Key", "X-edgeX-Passphrase", "X-edgeX-Signature", "X-edgeX-Timestamp"} {
				if got.Header.Get(name) != "" {
					t.Fatalf("anonymous handshake carries %s", name)
				}
			}
		})
	}
}
