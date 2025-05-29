# EdgeX Python SDK Testing Documentation

This document provides an overview of the testing strategy for the EdgeX Python SDK, including which components are covered by tests and which are not.

## Table of Contents

- [Test Structure](#test-structure)
- [Unit Tests](#unit-tests)
  - [Signing Adapter Tests](#signing-adapter-tests)
  - [Client Tests](#client-tests)
  - [Internal Client Tests](#internal-client-tests)
- [Integration Tests](#integration-tests)
  - [Account API Tests](#account-api-tests)
  - [Metadata API Tests](#metadata-api-tests)
  - [Order API Tests](#order-api-tests)
  - [Quote API Tests](#quote-api-tests)
  - [WebSocket API Tests](#websocket-api-tests)
- [Test Coverage](#test-coverage)
  - [Well-Covered Components](#well-covered-components)
  - [Partially Covered Components](#partially-covered-components)
  - [Components Lacking Coverage](#components-lacking-coverage)
- [Running Tests](#running-tests)
  - [Running Unit Tests](#running-unit-tests)
  - [Running Integration Tests](#running-integration-tests)
  - [Running Mock Tests](#running-mock-tests)
- [Test Environment](#test-environment)
- [Future Test Improvements](#future-test-improvements)

## Test Structure

The EdgeX Python SDK test suite is organized into two main categories:

1. **Unit Tests**: Located in `python_sdk/tests/`, these tests focus on individual components and functions without requiring external API connections.

2. **Integration Tests**: Located in `python_sdk/tests/integration/`, these tests verify the SDK's interaction with the EdgeX API endpoints.

## Unit Tests

### Signing Adapter Tests

File: `tests/test_starkex_signing_adapter.py`

These tests verify the functionality of the StarkEx signing adapter, which is responsible for cryptographic operations using the Stark curve.

**Coverage:**

- ✅ Sign and verify operations
- ✅ Public key derivation
- ✅ Error handling for invalid inputs
- ✅ Edge cases (zero values, large values)
- ✅ Signature verification with different messages and keys

**Test Cases:**
- `test_sign_and_verify`: Tests basic signing and verification
- `test_sign_different_messages`: Ensures different messages produce different signatures
- `test_sign_different_keys`: Ensures different keys produce different signatures
- `test_verify_invalid_signature`: Tests rejection of invalid signatures
- `test_verify_wrong_message`: Tests rejection when verifying with wrong message
- `test_verify_wrong_public_key`: Tests rejection when verifying with wrong public key
- `test_invalid_private_key`: Tests handling of invalid private keys
- `test_invalid_message_hash`: Tests handling of invalid message hashes

### Client Tests

File: `tests/test_client.py`

These tests verify the functionality of the main client class, which serves as the entry point for the SDK.

**Coverage:**

- ✅ Client initialization
- ✅ Parameter validation
- ✅ Method delegation to specialized clients

**Test Cases:**
- `test_init`: Tests client initialization with valid parameters
- `test_init_invalid_params`: Tests client initialization with invalid parameters

### Internal Client Tests

File: `tests/test_internal_client.py`

These tests verify the functionality of the internal client, which handles low-level API communication and signing.

**Coverage:**

- ✅ Request signing
- ✅ Timestamp generation
- ✅ Header construction
- ✅ Parameter validation

**Test Cases:**
- `test_init`: Tests internal client initialization
- `test_get_timestamp`: Tests timestamp generation
- `test_sign`: Tests message signing
- `test_get_stark_pri_key`: Tests private key retrieval

## Integration Tests

Integration tests verify the SDK's interaction with the EdgeX API endpoints. Most of these tests require valid API credentials to pass, except for tests of public endpoints.

> **Note**: For information about public endpoints that can be tested without credentials, see [PUBLIC_ENDPOINTS.md](PUBLIC_ENDPOINTS.md).

### Account API Tests

File: `tests/integration/test_account.py`

**Coverage:**

- ✅ Account retrieval
- ✅ Asset information retrieval
- ✅ Position information retrieval
- ✅ Transaction history retrieval

**Test Cases:**
- `test_get_account_by_id`: Tests retrieving account information
- `test_get_account_asset`: Tests retrieving account asset information
- `test_get_account_positions`: Tests retrieving account positions
- `test_get_position_transaction_page`: Tests retrieving position transaction history
- `test_get_collateral_transaction_page`: Tests retrieving collateral transaction history

### Metadata API Tests

File: `tests/integration/test_metadata.py`

**Coverage:**

- ✅ Contract metadata retrieval
- ✅ Server time retrieval

**Test Cases:**
- `test_get_metadata`: Tests retrieving contract metadata
- `test_get_server_time`: Tests retrieving server time
- `test_contract_exists`: Tests checking if a contract exists

### Order API Tests

File: `tests/integration/test_order.py`

**Coverage:**

- ✅ Order creation
- ✅ Order cancellation
- ✅ Active order retrieval
- ✅ Order fill transaction retrieval
- ✅ Maximum order size retrieval

**Test Cases:**
- `test_create_and_cancel_order`: Tests creating and canceling an order
- `test_get_active_orders`: Tests retrieving active orders
- `test_get_order_fill_transactions`: Tests retrieving order fill transactions
- `test_get_max_order_size`: Tests retrieving maximum order size

### Quote API Tests

File: `tests/integration/test_quote.py`

**Coverage:**

- ✅ 24-hour quote retrieval
- ✅ K-line data retrieval
- ✅ Order book depth retrieval
- ✅ Multi-contract K-line data retrieval

**Test Cases:**
- `test_get_24_hour_quote`: Tests retrieving 24-hour quotes
- `test_get_k_line`: Tests retrieving K-line data
- `test_get_order_book_depth`: Tests retrieving order book depth
- `test_get_multi_contract_k_line`: Tests retrieving multi-contract K-line data

### WebSocket API Tests

File: `tests/integration/test_websocket.py`

**Coverage:**

- ✅ Public WebSocket connection
- ✅ Private WebSocket connection
- ✅ Subscription to various channels
- ✅ Message handling

**Test Cases:**
- `test_public_websocket`: Tests connecting to the public WebSocket
- `test_private_websocket`: Tests connecting to the private WebSocket

## Test Coverage

### Well-Covered Components

The following components have good test coverage:

1. **StarkEx Signing Adapter**: Comprehensive unit tests for all cryptographic operations
2. **Client Initialization**: Thorough testing of parameter validation and initialization
3. **API Endpoint Structure**: Integration tests cover all major API endpoints

### Partially Covered Components

The following components have partial test coverage:

1. **Error Handling**: Some error conditions are tested, but not all possible error scenarios
2. **WebSocket Message Handling**: Basic connection is tested, but not all message types
3. **Response Parsing**: Basic parsing is tested, but not all edge cases

### Components Lacking Coverage

The following components could benefit from additional test coverage:

1. **Reconnection Logic**: WebSocket reconnection logic is not thoroughly tested
2. **Rate Limiting**: Handling of API rate limits is not tested
3. **Long-Running Operations**: Tests for long-running operations or pagination
4. **Concurrency**: Tests for concurrent API calls
5. **Transfer and Asset Clients**: Missing dedicated integration tests

## Running Tests

### Running Unit Tests

To run all unit tests:

```bash
cd python_sdk
python -m unittest discover tests
```

To run a specific test file:

```bash
cd python_sdk
python -m unittest tests/test_starkex_signing_adapter.py
```

### Running Integration Tests

To run integration tests with real API credentials:

```bash
cd python_sdk
python run_integration_tests.py
```

**Note**: This requires setting the following environment variables:
- `EDGEX_BASE_URL`: The base URL for the EdgeX API
- `EDGEX_WS_URL`: The WebSocket URL for the EdgeX API
- `EDGEX_ACCOUNT_ID`: Your EdgeX account ID
- `EDGEX_STARK_PRIVATE_KEY`: Your Stark private key

Example:
```bash
export EDGEX_ACCOUNT_ID="your_account_id_here"
export EDGEX_STARK_PRIVATE_KEY="your_stark_private_key_here"
```

**⚠️ SECURITY WARNING**: Never commit these credentials to version control. Always use environment variables or secure credential management systems.

### Running Public Endpoint Tests

To run only the tests for public endpoints (which don't require authentication):

```bash
cd python_sdk
python run_public_tests.py
```

This script uses dummy values for the required environment variables and only tests endpoints that don't require authentication.

**Note**: Some tests, particularly the WebSocket test, may be skipped due to connection issues or API limitations. This is expected behavior and doesn't indicate a problem with the SDK.

### Running Mock Tests

To run integration tests with mock values (for testing the SDK structure without real API calls):

```bash
cd python_sdk
python run_mock_tests.py
```

## Test Environment

The tests are designed to run in the following environments:

- **Unit Tests**: Can run in any environment with Python 3.7+
- **Integration Tests**: Require access to the EdgeX API
- **Mock Tests**: Can run in any environment with Python 3.7+

## Future Test Improvements

The following improvements could be made to the test suite:

1. **Separate Public and Private Endpoint Tests**: Fully separate tests for public endpoints from those requiring authentication

1. **Increased Unit Test Coverage**: Add tests for edge cases and error conditions
2. **Mocking External Dependencies**: Use mocks for API responses to test more scenarios
3. **Property-Based Testing**: Implement property-based testing for cryptographic operations
4. **Performance Testing**: Add tests for performance and resource usage
5. **Dedicated Transfer and Asset Tests**: Add dedicated integration tests for these clients
6. **WebSocket Message Testing**: Add more tests for different WebSocket message types
7. **Concurrency Testing**: Add tests for concurrent API calls
8. **Long-Running Operation Testing**: Add tests for pagination and long-running operations
