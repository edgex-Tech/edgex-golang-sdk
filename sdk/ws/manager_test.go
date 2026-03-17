package ws

import (
	"testing"
)

func TestNewManagerWithConfigDefaults(t *testing.T) {
	manager := NewManagerWithConfig("wss://quote.example.com/", 1001, &ManagerConfig{})

	if manager.privateWSPath != "/api/v1/private/ws" {
		t.Fatalf("unexpected private ws path: got %s", manager.privateWSPath)
	}
	if manager.baseURL != "wss://quote.example.com" {
		t.Fatalf("unexpected baseURL normalization: got %s", manager.baseURL)
	}
}

func TestNewManagerDefaults(t *testing.T) {
	manager := NewManager("wss://quote.example.com", 1001, "test-key", "test-pass", "test-secret")

	if manager.apiKey != "test-key" {
		t.Fatalf("unexpected apiKey: got %s", manager.apiKey)
	}
	if manager.apiPassphrase != "test-pass" {
		t.Fatalf("unexpected apiPassphrase: got %s", manager.apiPassphrase)
	}
}
