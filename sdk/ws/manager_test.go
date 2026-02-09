package ws

import (
	"testing"

	"github.com/edgex-Tech/edgex-golang-sdk/sdk/internal"
)

func TestNewManagerWithConfigDefaultsV2ToHMAC(t *testing.T) {
	manager := NewManagerWithConfig("wss://quote.example.com/", 1001, &ManagerConfig{
		APIVersion: internal.APIVersionV2,
	})

	if manager.signingMethod != internal.SigningMethodHMAC {
		t.Fatalf("unexpected signing method: got %s want %s", manager.signingMethod, internal.SigningMethodHMAC)
	}
	if manager.privateWSPath != "/api/v1/private/ws" {
		t.Fatalf("unexpected private ws path: got %s", manager.privateWSPath)
	}
	if manager.authHeaderKey != internal.DefaultHeaderKey {
		t.Fatalf("unexpected auth header key: got %s want %s", manager.authHeaderKey, internal.DefaultHeaderKey)
	}
	if manager.baseURL != "wss://quote.example.com" {
		t.Fatalf("unexpected baseURL normalization: got %s", manager.baseURL)
	}
}

func TestNewManagerKeepsBackwardCompatibleStark(t *testing.T) {
	manager := NewManager("wss://quote.example.com", 1001, "0xabc123")

	if manager.signingMethod != internal.SigningMethodStark {
		t.Fatalf("unexpected signing method: got %s want %s", manager.signingMethod, internal.SigningMethodStark)
	}
	if manager.starkPriKey != "abc123" {
		t.Fatalf("stark key should strip 0x prefix, got %s", manager.starkPriKey)
	}
}
