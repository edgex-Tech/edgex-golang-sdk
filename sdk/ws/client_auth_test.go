package ws

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/edgex-Tech/edgex-golang-sdk/sdk/internal"
	"github.com/gorilla/websocket"
)

func TestConfigureHMACPrivateConnection(t *testing.T) {
	client := NewClientWithConfig("wss://quote.example.com/api/v1/private/ws", true, 12345, &PrivateAuthConfig{
		SigningMethod: internal.SigningMethodHMAC,
		APIKey:        "api-key",
		APIPassphrase: "passphrase",
		APISecret:     "secret",
		AuthHeaderKey: "EDGEX",
		RequestURI:    "/api/v1/private/ws",
	})
	client.nowFn = func() time.Time {
		return time.UnixMilli(1700000000123)
	}

	dialer := websocket.Dialer{}
	headers := http.Header{}
	dialURL, err := client.configureHMACPrivateConnection(&dialer, headers)
	if err != nil {
		t.Fatalf("configureHMACPrivateConnection returned error: %v", err)
	}

	parsedURL, err := url.Parse(dialURL)
	if err != nil {
		t.Fatalf("failed to parse returned dial url: %v", err)
	}

	query := parsedURL.Query()
	if got := query.Get("accountId"); got != "12345" {
		t.Fatalf("accountId query mismatch: got %s", got)
	}
	if got := query.Get("timestamp"); got != "1700000000123" {
		t.Fatalf("timestamp query mismatch: got %s", got)
	}

	requestBody := "accountId=12345&timestamp=1700000000123"
	expectedSig := expectedHMACSignatureForTest("secret", "1700000000123", http.MethodGet, "/api/v1/private/ws", requestBody)
	if got := headers.Get("X-EDGEX-Signature"); got != expectedSig {
		t.Fatalf("signature mismatch: got %s want %s", got, expectedSig)
	}
	if got := headers.Get("X-EDGEX-Api-Key"); got != "api-key" {
		t.Fatalf("api key header mismatch: got %s", got)
	}
	if got := headers.Get("X-EDGEX-Passphrase"); got != "passphrase" {
		t.Fatalf("passphrase header mismatch: got %s", got)
	}
	if got := headers.Get("X-EDGEX-Timestamp"); got != "1700000000123" {
		t.Fatalf("timestamp header mismatch: got %s", got)
	}

	if len(dialer.Subprotocols) != 1 {
		t.Fatalf("expected one websocket protocol, got %d", len(dialer.Subprotocols))
	}
	protocol := dialer.Subprotocols[0]
	if strings.Contains(protocol, "=") {
		t.Fatalf("protocol should not contain '=' padding: %s", protocol)
	}

	decodedProtocol, err := decodeBase64NoPadding(protocol)
	if err != nil {
		t.Fatalf("failed to decode protocol: %v", err)
	}

	var protocolHeaders map[string]string
	if err := json.Unmarshal(decodedProtocol, &protocolHeaders); err != nil {
		t.Fatalf("failed to unmarshal protocol header json: %v", err)
	}

	if got := protocolHeaders["X-EDGEX-Signature"]; got != expectedSig {
		t.Fatalf("protocol signature mismatch: got %s want %s", got, expectedSig)
	}
	if got := protocolHeaders["X-EDGEX-Api-Key"]; got != "api-key" {
		t.Fatalf("protocol api key mismatch: got %s", got)
	}
	if got := protocolHeaders["X-EDGEX-Passphrase"]; got != "passphrase" {
		t.Fatalf("protocol passphrase mismatch: got %s", got)
	}
	if got := protocolHeaders["X-EDGEX-Timestamp"]; got != "1700000000123" {
		t.Fatalf("protocol timestamp mismatch: got %s", got)
	}
}

func TestConfigureHMACPrivateConnectionRequiresCredentials(t *testing.T) {
	client := NewClientWithConfig("wss://quote.example.com/api/v1/private/ws", true, 12345, &PrivateAuthConfig{
		SigningMethod: internal.SigningMethodHMAC,
	})

	dialer := websocket.Dialer{}
	headers := http.Header{}
	_, err := client.configureHMACPrivateConnection(&dialer, headers)
	if err == nil {
		t.Fatal("expected missing credential error but got nil")
	}
}

func TestNewClientWithConfigHMACDefaults(t *testing.T) {
	client := NewClientWithConfig("wss://quote.example.com/api/v1/private/ws", true, 1, &PrivateAuthConfig{
		SigningMethod: internal.SigningMethodHMAC,
		APIKey:        "api-key",
		APIPassphrase: "passphrase",
		APISecret:     "secret",
	})

	if client.authHeaderKey != internal.DefaultHeaderKey {
		t.Fatalf("unexpected auth header key: got %s want %s", client.authHeaderKey, internal.DefaultHeaderKey)
	}
	if client.requestURI != "/api/v1/private/ws" {
		t.Fatalf("unexpected requestURI: got %s", client.requestURI)
	}
}

func expectedHMACSignatureForTest(secret string, timestamp string, method string, requestURI string, requestBody string) string {
	message := timestamp + method + requestURI + requestBody
	base64Key := base64.StdEncoding.EncodeToString([]byte(secret))

	mac := hmac.New(sha256.New, []byte(base64Key))
	mac.Write([]byte(message))
	return hex.EncodeToString(mac.Sum(nil))
}

func decodeBase64NoPadding(raw string) ([]byte, error) {
	switch len(raw) % 4 {
	case 2:
		raw += "=="
	case 3:
		raw += "="
	}
	return base64.StdEncoding.DecodeString(raw)
}
