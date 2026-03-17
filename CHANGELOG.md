# Changelog

All notable changes to the EdgeX Golang SDK will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [2.0.0] - 2026-03-10

### Major Changes 🚀

This is a major release with **V2 API only** support. V1 API has been removed to focus on the new V2 architecture.

### Added ✨

#### Core Features
- 🔐 **EIP-712 Signature Support** - Ethereum-compatible L2 operation signing
  - Order creation signing (`LimitOrderParams`)
  - Transfer signing (`TransferParams`)
  - Account registration signing (`RegisterAccountParams`)
- 🔑 **HMAC-SHA256 Authentication** - Secure HTTP request authentication
  - Automatic signature generation for private endpoints
  - Configurable auth header prefix (default: `edgeX`)
- 📡 **WebSocket HMAC Authentication** - Private WebSocket with HMAC auth
  - Support for both public and private WebSocket connections
  - HMAC authentication via WebSocket subprotocol

#### SDK Enhancements
- 📝 **Configuration File Support** - `edgex-sdk.yaml` for test environment customization
  - Custom headers support (e.g., Cloudflare Access authentication)
  - Logging configuration (JSON/text format)
  - Completely optional and hidden from production code
- 🏷️ **Fixed Header** - Automatic `X-edgeX-Channel: Golang-SDK` on all requests
- 📊 **Structured Error Logging** - JSON-formatted error logs
  - Includes: `code`, `traceId`, `errorParam`, `msg`, `requestTime`, `responseTime`
  - Configurable format (JSON or text)
  - Helps with debugging and issue tracking
- 🧪 **Comprehensive Test Suite**
  - 10 order scenario tests (market, limit, stop, post-only, etc.)
  - Integration tests for full workflows
  - WebSocket connection tests

### Changed 🔄

#### API Changes
- 🌐 **All API paths updated** - `/api/v1/*` → `/api/v2/*`
- 🔧 **Client configuration simplified** - V2-only configuration structure
- ⚡ **Order creation** - Uses EIP-712 signing instead of Stark signing
- 💸 **Transfer operations** - Uses EIP-712 signing instead of Stark signing
- 📡 **WebSocket** - HMAC authentication for private connections

#### Internal Improvements
- 🏗️ Improved client initialization with configuration validation
- 🔒 Enhanced security with proper key management
- ⚡ Better performance with optimized signing operations
- 📝 Improved error messages with more context

### Fixed 🐛

- 🐛 Improved error handling with proper context propagation
- 🐛 Enhanced metadata caching mechanism
- 🐛 Better WebSocket connection management
- 🐛 Fixed timeout handling in HTTP requests

### Security 🔒

- 🔐 EIP-712 signing follows Ethereum standards
- 🔑 HMAC-SHA256 with proper timestamp validation
- 🛡️ Sensitive configuration file (`.edgex-sdk.yaml`) excluded from git
- 🔒 Private keys never logged or exposed

### Dependencies 📦

- ✅ Added `gopkg.in/yaml.v3` for configuration file parsing
- ✅ Retained `github.com/ethereum/go-ethereum` v1.14.12 for EIP-712
- ✅ Go 1.22+ required

---

## Migration Guide (V1 → V2)

### Breaking Changes ⚠️

1. **API Version** - Only V2 is supported
   ```go
   // V1 (Old - No longer supported)
   client, err := sdk.NewClient(&sdk.ClientConfig{
       BaseURL:     "https://edgex-prod-v2.edgex.exchange",
       AccountID:   12345,
       StarkPriKey: "your-stark-key",
   })
   
   // V2 (New)
   client, err := sdk.NewClient(&sdk.ClientConfig{
       BaseURL:       "https://edgex-prod-v2.edgex.exchange",
       AccountID:     12345,
       APIKey:        "your-api-key",
       APISecret:     "your-api-secret",
       APIPassphrase: "your-passphrase",
       SignerPriKey:  "your-trading-key",
       WalletPriKey:  "your-wallet-key",
   })
   ```

2. **Authentication** - HMAC instead of Stark signing
   - Obtain API credentials from EdgeX dashboard
   - Generate EIP-712 keys (secp256k1) for trading and wallet operations

3. **Order Creation** - Automatic EIP-712 signing
   ```go
   // The SDK automatically handles EIP-712 signing
   order, err := client.Order.CreateOrder(ctx, params, nil, nil)
   ```

4. **Transfer** - Automatic EIP-712 signing
   ```go
   // The SDK automatically handles EIP-712 signing
   transfer, err := client.Transfer.CreateTransferOut(ctx, params, nil)
   ```

### New Features You Can Use 🎁

1. **Update Leverage**
   ```go
   err := client.Account.UpdateLeverageSetting(ctx, contractID, leverage)
   ```

2. **Cancel by Client Order ID**
   ```go
   result, err := client.CancelOrder(ctx, &order.CancelOrderParams{
       ClientOrderId: &clientOrderID,
   })
   ```

3. **Configuration File** (Optional, for test environment)
   ```yaml
   # .edgex-sdk.yaml
   logging:
     error_log_format: "json"
   ```

### Testing Your Migration 🧪

1. Update your configuration to V2
2. Run your test suite
3. Check error logs in JSON format
4. Verify all operations work correctly

---

## Support

For questions or issues:
- 📝 [GitHub Issues](https://github.com/edgex-Tech/edgex-golang-sdk/issues)
- 📧 Email: support@edgex.exchange
- 📚 Docs: https://docs.edgex.exchange
