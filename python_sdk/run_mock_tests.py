#!/usr/bin/env python3
"""
Mock test runner for the EdgeX Python SDK.

This script runs the integration tests with dummy values for the environment variables.
It's useful for testing the SDK without actual API credentials.

Note: This will not make actual API calls, as the tests will fail when trying to connect
to the EdgeX API with invalid credentials. However, it's useful for testing the SDK
structure and mock signing adapter.
"""

import os
import sys
import subprocess
import logging

# Configure logging
logging.basicConfig(
    level=logging.INFO,
    format='%(asctime)s - %(name)s - %(levelname)s - %(message)s'
)
logger = logging.getLogger(__name__)

# Set dummy environment variables
os.environ["EDGEX_BASE_URL"] = "https://testnet.edgex.exchange"
os.environ["EDGEX_WS_URL"] = "wss://testnet.edgex.exchange"
os.environ["EDGEX_ACCOUNT_ID"] = "12345"
os.environ["EDGEX_STARK_PRIVATE_KEY"] = "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"
os.environ["EDGEX_SIGNING_ADAPTER"] = "mock"

# Log information about the mock tests
logger.info("Running integration tests with dummy credentials and a mock signing adapter.")
logger.info("This means that API calls will fail, but the SDK structure and mock signing adapter can be tested.")
logger.info("For actual API testing, use the run_integration_tests.py script with valid credentials.")

# Run the integration tests as a module
result = subprocess.run([sys.executable, "-m", "tests.integration"])

# Exit with the same exit code
sys.exit(result.returncode)
