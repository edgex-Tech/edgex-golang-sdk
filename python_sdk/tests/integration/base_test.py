"""Base test class for integration tests."""

import unittest
import asyncio
import logging
import os
from typing import Dict, Any, Optional

from edgex_sdk import Client, WebSocketManager
from .config import BASE_URL, WS_URL, ACCOUNT_ID, STARK_PRIVATE_KEY, STARKEX_SIGNING_ADAPTER, check_env_vars

# Configure logging
logging.basicConfig(
    level=logging.INFO,
    format='%(asctime)s - %(name)s - %(levelname)s - %(message)s'
)
logger = logging.getLogger(__name__)


class BaseIntegrationTest(unittest.TestCase):
    """Base class for integration tests."""

    @classmethod
    def setUpClass(cls):
        """Set up the test class."""
        # Check if required environment variables are set
        env_status = check_env_vars()
        if not env_status["all_set"]:
            missing_vars = env_status["missing_vars"]
            raise unittest.SkipTest(
                f"Skipping integration tests because the following environment variables are not set: {', '.join(missing_vars)}"
            )

        # Create client (uses StarkEx signing adapter by default)
        cls.client = Client(
            base_url=BASE_URL,
            account_id=ACCOUNT_ID,
            stark_private_key=STARK_PRIVATE_KEY
        )

        # Log which signing adapter is being used
        logger.info("Using StarkEx signing adapter (default)")

        # Create WebSocket manager (uses StarkEx signing adapter by default)
        cls.ws_manager = WebSocketManager(
            base_url=WS_URL,
            account_id=ACCOUNT_ID,
            stark_pri_key=STARK_PRIVATE_KEY
        )

        # Store test data
        cls.test_data = {}

    @classmethod
    def tearDownClass(cls):
        """Tear down the test class."""
        # Close WebSocket connections
        if hasattr(cls, "ws_manager"):
            cls.ws_manager.disconnect_all()

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

    def assertResponseSuccess(self, response: Dict[str, Any], msg: Optional[str] = None):
        """
        Assert that a response is successful.

        Args:
            response: The response to check
            msg: Optional message to display on failure
        """
        self.assertIn("code", response, msg=msg)
        self.assertEqual(response["code"], "SUCCESS", msg=msg)
        self.assertIn("data", response, msg=msg)
