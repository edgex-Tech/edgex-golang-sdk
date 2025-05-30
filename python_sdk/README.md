# EdgeX Python SDK

A Python SDK for interacting with the EdgeX Exchange API. This SDK provides a comprehensive interface to the EdgeX API, allowing you to easily integrate EdgeX functionality into your Python applications.

## Features

- **Complete API Coverage**: Access all EdgeX API endpoints
- **WebSocket Support**: Real-time data streaming
- **Async/Await**: Modern Python async interface
- **Type Hints**: Comprehensive type annotations for better IDE support
- **Error Handling**: Proper error handling and validation
- **Pagination**: Support for paginated API endpoints
- **Authentication**: Automatic request signing

## Installation

### From PyPI

```bash
pip install edgex-python-sdk
```

### From Source

```bash
git clone https://github.com/edgex-Tech/edgex-python-sdk.git
cd edgex-python-sdk/python_sdk
pip install -e .
```

### Using Requirements Files

For production use:
```bash
pip install -r requirements.txt
```

For development (includes testing and linting tools):
```bash
pip install -r requirements-dev.txt
```

### Virtual Environment (Recommended)

It's recommended to use a virtual environment:

```bash
# Create virtual environment
python3 -m venv venv

# Activate virtual environment
source venv/bin/activate  # On Windows: venv\Scripts\activate

# Install dependencies
pip install -r requirements.txt

# Or install in development mode
pip install -e .
```

## Quick Start

```python
import asyncio
from edgex_sdk import Client, OrderSide, OrderType

async def main():
    # Create a new client
    client = Client(
        base_url="https://testnet.edgex.exchange",
        account_id=12345,
        stark_private_key="your-stark-private-key"
    )

    # Get server time
    server_time = await client.get_server_time()
    print(f"Server Time: {server_time}")

    # Get account assets
    assets = await client.get_account_asset()
    print(f"Account Assets: {assets}")

    # Create a limit order
    order = await client.create_limit_order(
        contract_id="BTC-USDT",
        size="0.001",
        price="30000",
        side=OrderSide.BUY
    )
    print(f"Order created: {order}")

# Run the async function
asyncio.run(main())
```

## Architecture

The SDK is organized into modules that correspond to the EdgeX API structure:

```
edgex_sdk/
├── __init__.py
├── client.py           # Main client
├── account/            # Account API
├── asset/              # Asset API
├── funding/            # Funding API
├── internal/           # Internal utilities
├── metadata/           # Metadata API
├── order/              # Order API
├── quote/              # Quote API
├── transfer/           # Transfer API
└── ws/                 # WebSocket API
```

## Available APIs

The SDK currently supports the following API modules:

- **Account API**: Manage account positions, retrieve position transactions, and handle collateral transactions
  - Get account positions
  - Get position by contract ID
  - Get position transaction history
  - Get collateral transaction details
  - Update leverage settings

- **Asset API**: Handle asset management and withdrawals
  - Get asset orders with pagination
  - Get coin rates
  - Manage withdrawals (normal, cross-chain, and fast)
  - Get withdrawal records and sign information
  - Check withdrawable amounts

- **Funding API**: Manage funding operations and account balance
  - Handle funding transactions
  - Manage funding accounts
  - Get funding transaction history

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
  - Get 24-hour ticker data

- **Transfer API**: Handle asset transfers
  - Create transfer out orders
  - Get transfer records (in/out)
  - Check available withdrawal amounts
  - Manage transfer history

- **WebSocket API**: Real-time data streaming
  - Market data (tickers, K-lines, order book, trades)
  - Account updates
  - Order updates
  - Position updates

## WebSocket Support

The SDK provides a WebSocket manager for handling real-time data:

```python
from edgex_sdk import WebSocketManager

# Create a WebSocket manager (uses StarkExSigningAdapter by default)
ws_manager = WebSocketManager(
    base_url="wss://quote-testnet.edgex.exchange",
    account_id=12345,
    stark_pri_key="your-stark-private-key"
)

# Define message handlers
def ticker_handler(message):
    print(f"Ticker Update: {message}")

# Connect to public WebSocket
ws_manager.connect_public()

# Subscribe to ticker updates
ws_manager.subscribe_ticker("BTC-USDT", ticker_handler)
```

## Signing Adapters

The SDK provides a flexible signing mechanism through signing adapters. **StarkExSigningAdapter is used by default**, so you don't need to explicitly create one:

```python
from edgex_sdk import Client

# Create a client (uses StarkExSigningAdapter by default)
client = Client(
    base_url="https://testnet.edgex.exchange",
    account_id=12345,
    stark_private_key="your-stark-private-key"
)
```

If you need to use a custom signing adapter, you can still provide one:

```python
from edgex_sdk import Client, StarkExSigningAdapter

# Create a custom signing adapter (optional)
signing_adapter = StarkExSigningAdapter()

# Create a client with a custom signing adapter
client = Client(
    base_url="https://testnet.edgex.exchange",
    account_id=12345,
    stark_private_key="your-stark-private-key",
    signing_adapter=signing_adapter
)
```

The SDK includes the following signing adapters:

- **StarkExSigningAdapter** (default): Full implementation using StarkWare cryptographic operations
- **MockSigningAdapter**: A mock implementation for testing that doesn't perform actual cryptographic operations

You can also create your own signing adapter by implementing the `SigningAdapter` interface.

## Error Handling

The SDK provides proper error handling for API requests:

```python
try:
    # Create a limit order
    order = await client.create_limit_order(
        contract_id="BTC-USDT",
        size="0.001",
        price="30000",
        side=OrderSide.BUY
    )
    print(f"Order created: {order}")
except ValueError as e:
    print(f"Failed to create order: {str(e)}")
```

## Pagination

Many API endpoints support pagination:

```python
from edgex_sdk import GetActiveOrderParams

# Create pagination parameters
params = GetActiveOrderParams(
    size="10",
    offset_data=""
)

# Get active orders
orders = await client.get_active_orders(params)

# Get next page
if orders.get("data", {}).get("hasNext"):
    params.offset_data = orders.get("data", {}).get("offsetData")
    next_page = await client.get_active_orders(params)
```

## Examples

For detailed examples of each API endpoint, please refer to the [examples](examples) directory.

## Testing

To run the tests:

```bash
python run_tests.py
```

## Environment Variables

For testing, the following environment variables can be set:

- `EDGEX_BASE_URL`: Base URL for HTTP API endpoints (e.g., "https://api-testnet.edgex.exchange")
- `EDGEX_WS_URL`: Base URL for WebSocket endpoints (e.g., "wss://quote-testnet.edgex.exchange")
- `EDGEX_ACCOUNT_ID`: Your account ID
- `EDGEX_STARK_PRIVATE_KEY`: Your stark private key

## Documentation

For detailed API documentation, please refer to the [EdgeX API documentation](https://docs.edgex.exchange).

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin feature/my-new-feature`)
5. Create a new Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.