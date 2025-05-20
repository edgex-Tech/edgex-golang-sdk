"""
Main module for running integration tests.
"""

import unittest
import sys
import os
import logging

# Configure logging
logging.basicConfig(
    level=logging.INFO,
    format='%(asctime)s - %(name)s - %(levelname)s - %(message)s'
)
logger = logging.getLogger(__name__)

# Set the mock signing adapter in the environment
os.environ["EDGEX_SIGNING_ADAPTER"] = "mock"

# Log information about the mock tests
logger.info("Running integration tests with the mock signing adapter.")
logger.info("This means that cryptographic operations are not performed using the actual Stark curve.")

# Discover and run tests
test_loader = unittest.TestLoader()
test_suite = test_loader.discover(os.path.dirname(__file__), pattern='test_*.py')

# Run the tests
test_runner = unittest.TextTestRunner(verbosity=2)
result = test_runner.run(test_suite)

# Exit with the number of failures and errors
sys.exit(len(result.failures) + len(result.errors))
