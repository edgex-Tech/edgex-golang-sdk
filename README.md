# edgeX Golang SDK

A Go SDK for interacting with the edgeX Exchange API.

## Deprecation Notice

This `v1` branch contains the legacy EdgeX V1 SDK and is deprecated. New integrations should use the `main` branch for the current EdgeX V2 SDK.

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
    client, err := sdk.NewClient(
        sdk.WithBaseURL("https://testnet.edgex.exchange"),
        sdk.WithAccountID(12345),
        sdk.WithStarkPrivateKey("your-stark-private-key"),
    )
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
- `TEST_STARK_PRIVATE_KEY`: Your stark private key

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin feature/my-new-feature`)
5. Create a new Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
