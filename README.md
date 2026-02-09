# edgeX Golang SDK

A Go SDK for interacting with the edgeX Exchange API.

## Installation

```bash
go get github.com/edgex-Tech/edgex-golang-sdk
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/edgex-Tech/edgex-golang-sdk/sdk"
)

func main() {
    // Create a new client
    client, err := sdk.NewClient(&sdk.ClientConfig{
        BaseURL:     "https://testnet.edgex.exchange",
        AccountID:   12345,
        StarkPriKey: "your-stark-private-key",
    })
    if err != nil {
        log.Fatal(err)
    }

    // Create context
    ctx := context.Background()

    // Get account assets
    assets, err := client.Asset.GetAccountAsset(ctx)
    if err != nil {
        log.Fatal(err)
    }

    // Print asset information
    fmt.Printf("Account Assets: %+v\n", assets)
}
```

## V2 Migration Support

The SDK now supports V2 API routing and HMAC request signing (aligned with `edgex-web`).

Use these `sdk.ClientConfig` fields:

- `APIVersion`: `sdk.APIVersionV1` (default) or `sdk.APIVersionV2`
- `SigningMethod`: `sdk.SigningMethodStark` (default for v1) or `sdk.SigningMethodHMAC` (default for v2)
- `APIKey`, `APIPassphrase`, `APISecret`: required for private requests in HMAC mode
- `AuthHeaderKey`: app header prefix, default `edgeX` (e.g. `X-edgeX-Api-Key`)
- `TradingPriKey` / `TradingAddr`: EIP-712 trading signer key/address (used for V2 order signatures)
- `WalletPriKey` / `WalletAddr`: EIP-712 wallet signer key/address (used for V2 transfer signatures)

Example:

```go
client, err := sdk.NewClient(&sdk.ClientConfig{
    BaseURL:       "https://edgex-testnet-internal-v2.edgex.exchange",
    APIVersion:    sdk.APIVersionV2,
    SigningMethod: sdk.SigningMethodHMAC,
    APIKey:        "your-api-key",
    APIPassphrase: "your-passphrase",
    APISecret:     "your-secret",
    AccountID:     12345,
    TradingPriKey: "0x...trading-secp256k1-private-key...",
    WalletPriKey:  "0x...wallet-secp256k1-private-key...",
})
```

New V2 onboarding/register helpers:

- `client.CheckUserExist(...)`
- `client.GetAccountIdForRegister(...)`
- `client.OnboardSiteV2(...)`
- `client.RegisterAccountV2(...)`

EIP-712 flows ported in SDK:

- V2 order create: `client.CreateOrder(...)` when `APIVersion = sdk.APIVersionV2`
- V2 transfer create: `client.CreateTransferOut(...)` when `APIVersion = sdk.APIVersionV2`
- V2 register account: `client.RegisterAccountV2(...)` can auto-sign when `EthSignature` is empty and `WalletPriKey` is configured (optional overrides: `Owner`, `ChainID`, `VerifyingContract`)

V2 private WebSocket auth (perp) is supported via `ws.NewManagerWithConfig(...)`:

```go
manager := ws.NewManagerWithConfig("wss://edgex-quote-testnet-v2.edgex.exchange", accountID, &ws.ManagerConfig{
    APIVersion:    sdk.APIVersionV2,
    SigningMethod: sdk.SigningMethodHMAC,
    APIKey:        "your-api-key",
    APIPassphrase: "your-passphrase",
    APISecret:     "your-secret",
    // optional, default: "edgeX"
    AuthHeaderKey: "EDGEX",
})
```

## Available APIs

The SDK currently supports the following API modules:

- **Account API**: Manage account positions, retrieve position transactions, and handle collateral transactions
  - Get account positions
  - Get position by contract ID
  - Get position transaction history
  - Get collateral transaction details

- **Asset API**: Handle asset management and withdrawals
  - Get asset orders with pagination
  - Get coin rates
  - Manage withdrawals (normal, cross-chain, and fast)
  - Get withdrawal records and sign information
  - Check withdrawable amounts

- **Funding API**: Manage funding operations and account balance
  - Handle funding transactions
  - Manage funding accounts

- **Metadata API**: Access exchange system information
  - Get server time
  - Get exchange metadata (trading pairs, contracts, etc.)

- **Order API**: Comprehensive order management
  - Create and cancel orders
  - Get active orders
  - Get order fill transactions
  - Calculate maximum order sizes
  - Manage order history

- **Quote API**: Access market data and pricing
  - Get multi-contract K-line data
  - Get order book depth
  - Access real-time market quotes

- **Transfer API**: Handle asset transfers
  - Create transfer out orders
  - Get transfer records (in/out)
  - Check available withdrawal amounts
  - Manage transfer history

For detailed examples of each API endpoint, please refer to the test files in the `test` directory.

## Environment Variables

For testing, the following environment variables need to be set:

- `TEST_BASE_URL`: Base URL for HTTP API endpoints (e.g., "https://api-testnet.edgex.exchange")
- `TEST_WS_BASE_URL`: Base URL for WebSocket endpoints (e.g., "wss://quote-testnet.edgex.exchange")
- `TEST_ACCOUNT_ID`: Your account ID
- `TEST_API_VERSION`: `v1` or `v2` (optional; auto-detected from URL/credentials)
- `TEST_SIGNING_METHOD`: `stark` or `hmac` (optional; auto-detected from API version/credentials)
- `TEST_STARK_PRIVATE_KEY`: Stark private key (required for v1 Stark signing flows)
- `TEST_TRADING_PRIVATE_KEY`: required for v2 EIP-712 order signing tests
- `TEST_WALLET_PRIVATE_KEY`: required for v2 EIP-712 transfer signing tests
- `TEST_WS_SIGNING_METHOD`: `stark` or `hmac` (optional; auto-detected from credentials)
- `TEST_API_KEY`, `TEST_API_PASSPHRASE`, `TEST_API_SECRET`: required for v2 HMAC private WS tests
- `TEST_AUTH_HEADER_KEY`: optional HMAC header prefix override (defaults to `edgeX`)
- `TEST_TRANSFER_RECEIVER_ACCOUNT_ID`, `TEST_TRANSFER_RECEIVER_L2_KEY`: optional; required only for `TestCreateTransferOut`
- `TEST_ENABLE_MUTATION_TESTS`: set to `true` to enable withdraw/create mutation tests

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin feature/my-new-feature`)
5. Create a new Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
