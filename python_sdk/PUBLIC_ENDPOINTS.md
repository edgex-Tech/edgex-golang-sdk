# EdgeX Python SDK Public Endpoints

This document provides information about the public endpoints in the EdgeX Python SDK that can be tested without authentication credentials.

## Public Endpoints Overview

The EdgeX API provides several public endpoints that don't require authentication. These endpoints can be tested without providing account credentials, making them ideal for basic integration testing and verification of the SDK's functionality.

## Available Public Endpoints

The following endpoints are public and can be accessed without authentication:

### Metadata Endpoints

1. **Get Metadata** (`/api/v1/public/meta/getMetaData`)
   - Returns information about available contracts and exchange settings
   - Implemented in `metadata.client.get_metadata()`

2. **Get Server Time** (`/api/v1/public/meta/getServerTime`)
   - Returns the current server time
   - Implemented in `metadata.client.get_server_time()`

### Quote Endpoints

1. **Get Quote Summary** (`/api/v1/public/quote/getTicketSummary`)
   - Returns a summary of quotes for a specific contract
   - Implemented in `quote.client.get_quote_summary()`

2. **Get 24-Hour Ticker** (`/api/v1/public/quote/getTicker`)
   - Returns 24-hour ticker information for a specific contract
   - Implemented in `quote.client.get_24_hour_quote()`

3. **Get K-Line Data** (`/api/v1/public/quote/getKline`)
   - Returns K-line (candlestick) data for a specific contract
   - Implemented in `quote.client.get_k_line()`

4. **Get Order Book Depth** (`/api/v1/public/quote/getDepth`)
   - Returns the order book depth for a specific contract
   - Implemented in `quote.client.get_order_book_depth()`

5. **Get Multi-Contract K-Line Data** (`/api/v1/public/quote/getMultiContractKline`)
   - Returns K-line data for multiple contracts
   - Implemented in `quote.client.get_multi_contract_k_line()`

### Funding Endpoints

1. **Get Funding Rate** (`/api/v1/public/funding/getFundingRatePage`)
   - Returns funding rate information
   - Implemented in `funding.client.get_funding_rate_page()`

2. **Get Latest Funding Rate** (`/api/v1/public/funding/getLatestFundingRate`)
   - Returns the latest funding rate
   - Implemented in `funding.client.get_latest_funding_rate()`

### WebSocket Endpoints

1. **Public WebSocket** (`/ws/public`)
   - Provides real-time market data
   - Implemented in `ws.manager.connect_public()`
   - Note: Testing WebSocket connections may be challenging without proper credentials or in certain environments. The test is designed to skip rather than fail if it encounters connection issues.

## Testing Public Endpoints

You can test these public endpoints without providing authentication credentials. Here's how to set up and run tests for public endpoints only:

### Creating a Public Endpoints Test Runner

Create a new file `run_public_tests.py` in the `python_sdk` directory:

```python
#!/usr/bin/env python3
"""
Public endpoints test runner for the EdgeX Python SDK.

This script runs tests for public endpoints that don't require authentication.
"""

import os
import sys
import unittest
import logging

# Configure logging
logging.basicConfig(
    level=logging.INFO,
    format='%(asctime)s - %(name)s - %(levelname)s - %(message)s'
)
logger = logging.getLogger(__name__)

# Set dummy values for required environment variables
# These won't be used for authentication but are needed for client initialization
os.environ["EDGEX_BASE_URL"] = "https://testnet.edgex.exchange"
os.environ["EDGEX_WS_URL"] = "wss://testnet.edgex.exchange"
os.environ["EDGEX_ACCOUNT_ID"] = "0"  # Dummy value
os.environ["EDGEX_STARK_PRIVATE_KEY"] = "0" * 64  # Dummy value
os.environ["EDGEX_SIGNING_ADAPTER"] = "mock"  # Use mock adapter
os.environ["EDGEX_PUBLIC_ONLY"] = "true"  # Flag to indicate public endpoints only

# Log information
logger.info("Running tests for public endpoints only")
logger.info("These tests don't require authentication credentials")

# Discover and run tests
test_loader = unittest.TestLoader()
test_suite = test_loader.discover('tests/integration/public', pattern='test_*.py')

# Run the tests
test_runner = unittest.TextTestRunner(verbosity=2)
result = test_runner.run(test_suite)

# Exit with the number of failures and errors
sys.exit(len(result.failures) + len(result.errors))
```

### Creating a Base Test Class for Public Endpoints

Create a new file `tests/integration/public/base_test.py`:

```python
"""Base test class for public endpoint tests."""

import unittest
import asyncio
import logging
import os
from typing import Dict, Any

from edgex_sdk import Client
from edgex_sdk.internal.mock_signing_adapter import MockSigningAdapter
from tests.integration.config import BASE_URL

# Configure logging
logging.basicConfig(
    level=logging.INFO,
    format='%(asctime)s - %(name)s - %(levelname)s - %(message)s'
)
logger = logging.getLogger(__name__)


class BasePublicEndpointTest(unittest.TestCase):
    """Base class for public endpoint tests."""

    @classmethod
    def setUpClass(cls):
        """Set up the test class."""
        # Create a mock signing adapter
        signing_adapter = MockSigningAdapter()

        # Create client with dummy values
        # The account_id and stark_private_key won't be used for public endpoints
        cls.client = Client(
            base_url=BASE_URL,
            account_id=0,  # Dummy value
            stark_private_key="0" * 64,  # Dummy value
            signing_adapter=signing_adapter
        )

        # Store test data
        cls.test_data = {}

    def run_async(self, coro):
        """
        Run an async coroutine in the current event loop.

        Args:
            coro: The coroutine to run

        Returns:
            Any: The result of the coroutine
        """
        return asyncio.get_event_loop().run_until_complete(coro)

    def setUp(self):
        """Set up the test."""
        # Create a new event loop for each test
        self.loop = asyncio.new_event_loop()
        asyncio.set_event_loop(self.loop)

    def tearDown(self):
        """Tear down the test."""
        # Close the event loop
        self.loop.close()

    def assertResponseSuccess(self, response: Dict[str, Any]):
        """
        Assert that a response is successful.

        Args:
            response: The response to check
        """
        self.assertIn("code", response)
        self.assertEqual(response["code"], "SUCCESS")
        self.assertIn("data", response)
```

### Creating Test Files for Public Endpoints

Create test files for each public endpoint category:

1. `tests/integration/public/test_metadata.py`:

```python
"""Tests for public metadata endpoints."""

import unittest
import logging

from tests.integration.public.base_test import BasePublicEndpointTest

# Configure logging
logger = logging.getLogger(__name__)


class TestPublicMetadataAPI(BasePublicEndpointTest):
    """Tests for public metadata endpoints."""

    def test_get_metadata(self):
        """Test get_metadata method."""
        # Get metadata
        metadata = self.run_async(self.client.get_metadata())

        # Check response
        self.assertResponseSuccess(metadata)

        # Check data
        data = metadata.get("data", {})
        self.assertIn("contractList", data)
        self.assertIsInstance(data["contractList"], list)

        # Log contract count
        logger.info(f"Found {len(data['contractList'])} contracts")

    def test_get_server_time(self):
        """Test get_server_time method."""
        # Get server time
        server_time = self.run_async(self.client.get_server_time())

        # Check response
        self.assertResponseSuccess(server_time)

        # Check data
        data = server_time.get("data", {})
        self.assertIn("serverTime", data)
        self.assertIsInstance(data["serverTime"], int)

        # Log server time
        logger.info(f"Server time: {data['serverTime']}")


if __name__ == "__main__":
    unittest.main()
```

2. `tests/integration/public/test_quote.py`:

```python
"""Tests for public quote endpoints."""

import unittest
import logging

from edgex_sdk import GetKLineParams, GetOrderBookDepthParams, GetMultiContractKLineParams
from tests.integration.public.base_test import BasePublicEndpointTest

# Configure logging
logger = logging.getLogger(__name__)

# Test contract ID
TEST_CONTRACT_ID = "BTC-USDT"  # Use a common contract that should exist


class TestPublicQuoteAPI(BasePublicEndpointTest):
    """Tests for public quote endpoints."""

    def test_get_24_hour_quote(self):
        """Test get_24_hour_quote method."""
        # Get 24-hour quote
        quote = self.run_async(self.client.quote.get_24_hour_quote(TEST_CONTRACT_ID))

        # Check response
        self.assertResponseSuccess(quote)

        # Check data
        data = quote.get("data", [])
        self.assertIsInstance(data, list)

        # Log quote details
        if data:
            first_quote = data[0]
            logger.info(f"24-hour quote for {TEST_CONTRACT_ID}: {first_quote.get('lastPrice')}")
        else:
            logger.info(f"No 24-hour quote data for {TEST_CONTRACT_ID}")

    def test_get_k_line(self):
        """Test get_k_line method."""
        # Create parameters
        params = GetKLineParams(
            contract_id=TEST_CONTRACT_ID,
            interval="1m",
            size="10"
        )

        # Get K-line data
        klines = self.run_async(self.client.quote.get_k_line(params))

        # Check response
        self.assertResponseSuccess(klines)

        # Check data
        data = klines.get("data", {})

        # Log K-line details
        if "list" in data and data["list"]:
            first_kline = data["list"][0]
            logger.info(f"First K-line for {TEST_CONTRACT_ID}: {first_kline}")
        else:
            logger.info(f"No K-line data for {TEST_CONTRACT_ID}")

    def test_get_order_book_depth(self):
        """Test get_order_book_depth method."""
        # Create parameters
        params = GetOrderBookDepthParams(
            contract_id=TEST_CONTRACT_ID,
            limit=10
        )

        # Get order book depth
        depth = self.run_async(self.client.quote.get_order_book_depth(params))

        # Check response
        self.assertResponseSuccess(depth)

        # Check data
        data = depth.get("data", {})
        self.assertIn("asks", data)
        self.assertIn("bids", data)
        self.assertIsInstance(data["asks"], list)
        self.assertIsInstance(data["bids"], list)

        # Log depth details
        asks = data["asks"]
        bids = data["bids"]
        logger.info(f"Order book depth for {TEST_CONTRACT_ID}: {len(asks)} asks, {len(bids)} bids")


if __name__ == "__main__":
    unittest.main()
```

## Modifying the Client Implementation

To properly support public endpoints without authentication, you may need to modify the internal client implementation to skip authentication for public endpoints. Here's a suggested approach:

1. Modify the `internal/client.py` file to check if the URL contains `/public/` and skip adding authentication headers in that case.

2. Alternatively, create a separate public client class that doesn't require authentication credentials.

## Benefits of Testing Public Endpoints

Testing public endpoints provides several benefits:

1. **Easier CI/CD Integration**: Tests can run in CI/CD pipelines without requiring secure storage of API credentials.

2. **Basic SDK Verification**: Verifies that the SDK can correctly communicate with the API and parse responses.

3. **Documentation Validation**: Ensures that the API documentation for public endpoints is accurate.

4. **Dependency Verification**: Confirms that all required dependencies are correctly installed and functioning.

## Conclusion

By separating tests for public endpoints, you can ensure that basic SDK functionality can be verified without requiring authentication credentials. This makes it easier to run tests in various environments and provides a good starting point for users who want to try out the SDK without setting up full authentication.
