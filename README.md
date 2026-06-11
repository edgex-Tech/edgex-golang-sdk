# EdgeX Golang SDK (V2)

Official Golang SDK for EdgeX V2 API - A high-performance, production-ready SDK for perpetual trading.

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

## 🚀 Features

- ✅ **Current Code-Based V2 Coverage** - account/order/transfer plus unified-asset and CCTP helpers
- 🔐 **EIP-712 Signature** - Ethereum-compatible L2 operation signing
- 🔑 **HMAC-SHA256 Authentication** - Secure HTTP request authentication
- 📡 **Real-time WebSocket** - Public market data + Private account updates
- ⚡ **High Performance** - Metadata caching, connection pooling
- 🛡️ **Production Ready** - Comprehensive error handling and logging
- 📝 **Well Documented** - Full API coverage with examples

## 📦 Installation

```bash
go get github.com/edgex-Tech/edgex-golang-sdk/v2@v2.0.0
```

**Requirements:** Go 1.22 or higher

Recommended for macOS 26 developers:

- Prefer Go 1.24 or newer when running tests locally on macOS 26.
- With Go 1.22.7 on macOS 26, some test binaries may fail at startup with:
  - `dyld: missing LC_UUID load command`

## 📌 Documentation Status

This repository contains historical markdown files in the root directory. For current behavior:

- Code under `sdk/` is the source of truth.
- Local knowledge bases and agent skills are intentionally kept outside this repo by default.
- This README is a summary, not the canonical interface inventory.

## 🎯 Quick Start

### Basic Usage

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/edgex-Tech/edgex-golang-sdk/v2/sdk"
	"github.com/edgex-Tech/edgex-golang-sdk/v2/sdk/account"
	"github.com/edgex-Tech/edgex-golang-sdk/v2/sdk/order"
)

func main() {
	client, err := sdk.NewClient(&sdk.ClientConfig{
		BaseURL:       "https://edgex-prod-v2.edgex.exchange",
		AssetBaseURL: "https://spot.edgex.exchange",
		AccountID:     12345,
		APIKey:        "your-api-key",
		APISecret:     "your-api-secret",
		APIPassphrase: "your-passphrase",
		SignerPriKey:  "your-signer-private-key", // For order signing
		WalletPriKey:  "your-wallet-private-key", // For withdraw / transfer signing
	})
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	accountAsset, err := client.GetAccountAsset(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Account Asset: %+v\n", accountAsset)

	marginModeResp, err := client.SetMarginMode(ctx, &account.SetMarginModeParams{
		ContractID: "1001",
		MarginMode: "1",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Margin Mode Updated: %+v\n", marginModeResp)

	orderResp, err := client.CreateOrder(ctx, &order.CreateOrderParams{
		ContractId: "1001",
		Side:       "BUY",
		Type:       order.OrderTypeMarket,
		Size:       "0.01",
		Price:      "0",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Order Created: %+v\n", orderResp)
}
```

## 📚 Configuration

### Client Configuration

```go
type ClientConfig struct {
    BaseURL          string
    AssetBaseURL     string // Optional unified-asset base URL, defaults to BaseURL
    AccountID        int64
    SignerPriKey     string // EIP-712 signer private key for orders
    WalletPriKey     string // EIP-712 wallet private key for withdraw / transfer
    SignerAddr       string // Optional, auto-derived if empty
    WalletAddr       string // Optional, auto-derived if empty
    APIKey           string // HTTP HMAC API key
    APIPassphrase    string // HTTP HMAC passphrase
    APISecret        string // HTTP HMAC secret
    AuthHeaderKey    string // Optional HTTP auth header prefix, default "edgeX"
    MetaDataCacheTTL *time.Duration
}
```

### Configuration File (Optional)

Create `.edgex-sdk.yaml` in your project root for optional local customization:

```yaml
# Logging configuration
logging:
  error_log_format: "json"  # "json" or "text"
  enable_request_log: false
  enable_response_log: false
```

**Note:** This file is optional. Use it only when you need local custom headers or verbose logging.

**Mutation test switch:** `EDGEX_ENABLE_MUTATION_TESTS` only enables real order/withdrawal tests when its value is explicitly `true`. `false`, empty, or unset all mean disabled.

## 🔌 API Coverage

Current code-based counts:

| Area | Count |
| --- | ---: |
| Public HTTP | 9 |
| Private HTTP | 30 |
| WebSocket Paths | 2 |
| Total HTTP | 39 |

Module breakdown:

| Module | Count | Notes |
| --- | ---: | --- |
| Metadata | 2 | `GetServerTime`, `GetMetaData` |
| Quote | 5 | Includes `GetMultiContractKLine` |
| Funding | 2 | Accessed via `client.Funding` |
| Account | 15 | Includes `GetPositionOrders`, `UpdateLeverageSetting`, `SetMarginMode` |
| UnifiedAsset | 4 | Unified withdraw, deposit-data, and asset-flow helpers |
| Order | 7 | Root client exposes 7 order operations |
| Transfer | 4 | Includes `CreateTransferOut` |
| WebSocket | 2 | Still `/api/v1/.../ws` |

This README keeps the stable summary. Fast-changing details are better maintained in your local workspace knowledge base or in dedicated customer docs.

## 💡 Usage Examples

For current examples, prefer:

- `test/order/`
- `test/integration/`
- `test/ws/`

Minimal WebSocket example:

```go
wsManager := ws.NewManagerWithConfig("wss://edgex-quote-prod-v2.edgex.exchange", accountID, &ws.ManagerConfig{
    APIKey:        apiKey,
    APIPassphrase: passphrase,
    APISecret:     secret,
})

if err := wsManager.ConnectPublic(ctx); err != nil {
    log.Fatal(err)
}

if err := wsManager.SubscribeMarketTicker("1001", func(message []byte) {
    fmt.Printf("ticker: %s\n", string(message))
}); err != nil {
    log.Fatal(err)
}

if err := wsManager.ConnectPrivate(ctx); err != nil {
    log.Fatal(err)
}

if err := wsManager.OnPrivateMessage("ORDER_UPDATE", func(message []byte) {
    fmt.Printf("order update: %s\n", string(message))
}); err != nil {
    log.Fatal(err)
}
```

## 🧪 Testing

### Environment Variables

This SDK only supports **V2 API with HMAC authentication**.

```bash
# Required: API Configuration
export EDGEX_BASE_URL="https://edgex-prod-v2.edgex.exchange"
export EDGEX_ASSET_BASE_URL="https://spot.edgex.exchange"
export EDGEX_ACCOUNT_ID=12345

# Required: HMAC Authentication
export EDGEX_API_KEY="your-api-key"
export EDGEX_API_SECRET="your-api-secret"
export EDGEX_API_PASSPHRASE="your-passphrase"

# Required: EIP-712 Signing Keys
export EDGEX_SIGNER_PRIVATE_KEY="0x..."  # For order creation
export EDGEX_WALLET_PRIVATE_KEY="0x..."  # For withdrawal creation

# Optional: Explicit addresses (auto-derived if omitted)
export EDGEX_SIGNER_ADDRESS="0x..."
export EDGEX_WALLET_ADDRESS="0x..."

# Optional: WebSocket Configuration
export EDGEX_WS_BASE_URL="wss://edgex-quote-prod-v2.edgex.exchange"

# Optional: Enable Mutation Tests (creates real orders/withdrawals)
export EDGEX_ENABLE_MUTATION_TESTS="false"

# Optional: Custom HTTP auth header prefix
export EDGEX_AUTH_HEADER_KEY="edgeX"
```

### Run Tests

```bash
# Run all tests
go test ./...

# Run specific test suite
go test ./test/order -v

# Run with coverage
go test ./... -cover

# Run only unit tests (skip integration tests)
go test ./... -short
```

macOS 26 note:

```bash
# If Go 1.22.x test binaries fail with:
# dyld: missing LC_UUID load command
go test -ldflags=-linkmode=external ./sdk/...
go test -ldflags=-linkmode=external ./...
```

### Local Replace For SDK Testing

When testing a local SDK checkout together with `official-golang-sdk-test`, keep the SDK module path as:

```go
module github.com/edgex-Tech/edgex-golang-sdk/v2
```

and use a local `replace` in `official-golang-sdk-test/go.mod`:

```go
require github.com/edgex-Tech/edgex-golang-sdk/v2 v2.0.0
replace github.com/edgex-Tech/edgex-golang-sdk/v2 => ../edgex-golang-sdk-pedro
```

### Test Scenarios

The repository includes:
- module-level tests under `test/`
- integration scenarios under `test/integration/`
- websocket tests under `test/ws/`
- request signing and helper tests under `sdk/**/_test.go`

Use your local workspace test guide if you maintain one; this README keeps only the summary.

### Authentication Enhancements

- **EIP-712 Signing**: Ethereum-compatible L2 operation signatures
- **HMAC Authentication**: Secure HTTP request authentication
- **WebSocket HMAC**: Private WebSocket with HMAC authentication

### SDK Enhancements

- **Fixed Header**: Automatic `X-edgeX-Channel: Golang-SDK` header
- **Structured Logging**: JSON-formatted error logs with trace IDs
- **Better Error Handling**: Comprehensive error messages with context

## 🔧 Error Handling

The SDK automatically logs errors in structured JSON format:

```json
{
  "code": 400,
  "traceId": "abc-123-def-456",
  "errorParam": {"contractId": "1001", "price": "50000"},
  "msg": "Invalid parameter: price out of range",
  "requestTime": "2026-03-10T10:00:00Z",
  "responseTime": "2026-03-10T10:00:01Z"
}
```

All API errors include:
- **code**: HTTP status code
- **traceId**: Request trace ID for debugging
- **errorParam**: Request parameters that caused the error
- **msg**: Error message from server
- **requestTime**: When the request was sent
- **responseTime**: When the response was received

## 📖 Documentation

- [API Reference](https://docs.edgex.exchange/api-v2/)
- [WebSocket API](https://docs.edgex.exchange/api-v2/websocket-api/)
- [Authentication Guide](https://docs.edgex.exchange/api-v2/authentication/)
- [EIP-712 Signing](https://docs.edgex.exchange/api-v2/sign/)

## 🤝 Contributing

We welcome contributions! Please follow these steps:

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development Guidelines

- Follow Go best practices and conventions
- Write tests for new features
- Update documentation as needed
- Ensure all tests pass before submitting PR

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🆘 Support

- **Issues**: [GitHub Issues](https://github.com/edgex-Tech/edgex-golang-sdk/issues)
- **Documentation**: [docs.edgex.exchange](https://docs.edgex.exchange)
- **Email**: support@edgex.exchange

## 🔗 Links

- [EdgeX Exchange](https://edgex.exchange)
- [API Documentation](https://docs.edgex.exchange)
- [TypeScript SDK](https://github.com/edgex-Tech/edgex-typescript-sdk)

---

**Made with ❤️ by the EdgeX Team**
