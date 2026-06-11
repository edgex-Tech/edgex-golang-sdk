package ws

import (
	"encoding/json"
	"sync/atomic"
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

func TestTradeEventDispatchesByDataGroup(t *testing.T) {
	client := NewClient("wss://quote.example.com/api/v1/private/ws", true, 1001, "", "", "")

	var accountCalls atomic.Int32
	var orderCalls atomic.Int32
	var eventCalls atomic.Int32

	client.OnMessage("account", func(message []byte) {
		accountCalls.Add(1)
	})
	client.OnMessage("order", func(message []byte) {
		orderCalls.Add(1)
	})
	client.OnMessage("trade-event", func(message []byte) {
		eventCalls.Add(1)
	})

	message := []byte(`{
		"type":"trade-event",
		"content":{
			"event":"ORDER_UPDATE",
			"version":1,
			"time":1773832863190,
			"accountId":0,
			"data":{
				"account":[{"id":"1"}],
				"order":[{"orderId":"2"}]
			}
		}
	}`)

	var msg Message
	if err := json.Unmarshal(message, &msg); err != nil {
		t.Fatalf("failed to unmarshal message: %v", err)
	}

	if msg.Type == "trade-event" {
		if handler, ok := client.handlers[msg.Type]; ok {
			handler(message)
		}
		var tradeEvent TradeEvent
		if err := json.Unmarshal(message, &tradeEvent); err != nil {
			t.Fatalf("failed to unmarshal trade-event: %v", err)
		}
		for group := range tradeEvent.Content.Data {
			if handler, ok := client.handlers[group]; ok {
				handler(message)
			}
		}
	}

	if eventCalls.Load() != 1 {
		t.Fatalf("expected trade-event handler to be called once, got %d", eventCalls.Load())
	}
	if accountCalls.Load() != 1 {
		t.Fatalf("expected account handler to be called once, got %d", accountCalls.Load())
	}
	if orderCalls.Load() != 1 {
		t.Fatalf("expected order handler to be called once, got %d", orderCalls.Load())
	}
}
