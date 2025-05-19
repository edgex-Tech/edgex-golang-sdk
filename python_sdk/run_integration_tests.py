#!/usr/bin/env python3
"""
Integration test runner for the EdgeX Python SDK.
"""

import os
import sys
import subprocess
import logging

from tests.integration.config import check_env_vars

# Configure logging
logging.basicConfig(
    level=logging.INFO,
    format='%(asctime)s - %(name)s - %(levelname)s - %(message)s'
)
logger = logging.getLogger(__name__)

# Check if required environment variables are set
env_status = check_env_vars()
if not env_status["all_set"]:
    missing_vars = env_status["missing_vars"]
    logger.error(f"Cannot run integration tests because the following environment variables are not set: {', '.join(missing_vars)}")
    logger.error("Please set these environment variables and try again.")
    sys.exit(1)

# Set the StarkEx signing adapter in the environment
os.environ["EDGEX_SIGNING_ADAPTER"] = "starkex"

# Log information about the signing adapter
logger.info("Running integration tests with the StarkEx signing adapter.")
logger.info("This means that cryptographic operations are performed using the actual Stark curve implementation.")

# Run the integration tests as a module
result = subprocess.run([sys.executable, "-m", "tests.integration"])

# Exit with the same exit code
sys.exit(result.returncode)
