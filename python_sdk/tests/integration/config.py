"""Configuration for integration tests."""

import os
from typing import Dict, Any

from edgex_sdk.internal.mock_signing_adapter import MockSigningAdapter
from edgex_sdk.internal.starkex_signing_adapter import StarkExSigningAdapter

# Load environment variables
BASE_URL = os.getenv("EDGEX_BASE_URL", "https://testnet.edgex.exchange")
WS_URL = os.getenv("EDGEX_WS_URL", "wss://quote-testnet.edgex.exchange")
ACCOUNT_ID = int(os.getenv("EDGEX_ACCOUNT_ID", "0"))
STARK_PRIVATE_KEY = os.getenv("EDGEX_STARK_PRIVATE_KEY", "")

# Test data
TEST_CONTRACT_ID = "10000004"  # Contract ID provided
TEST_ORDER_SIZE = "0.001"
TEST_ORDER_PRICE = "30000"

# Create signing adapters for testing
MOCK_SIGNING_ADAPTER = MockSigningAdapter()
STARKEX_SIGNING_ADAPTER = StarkExSigningAdapter()

# Check if required environment variables are set
def check_env_vars() -> Dict[str, Any]:
    """
    Check if required environment variables are set.

    Returns:
        Dict[str, Any]: A dictionary with the status of each environment variable
    """
    env_vars = {
        "EDGEX_BASE_URL": BASE_URL,
        "EDGEX_WS_URL": WS_URL,
        "EDGEX_ACCOUNT_ID": ACCOUNT_ID,
        "EDGEX_STARK_PRIVATE_KEY": STARK_PRIVATE_KEY
    }

    missing_vars = []
    for var_name, var_value in env_vars.items():
        if not var_value or (var_name == "EDGEX_ACCOUNT_ID" and var_value == 0):
            missing_vars.append(var_name)

    return {
        "all_set": len(missing_vars) == 0,
        "missing_vars": missing_vars,
        "env_vars": env_vars
    }
