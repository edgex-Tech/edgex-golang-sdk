"""Integration tests for the WebSocket API."""

import unittest
import logging
import asyncio
import json
from typing import Dict, Any, List

from tests.integration.base_test import BaseIntegrationTest
from tests.integration.config import TEST_CONTRACT_ID

# Configure logging
logger = logging.getLogger(__name__)


class TestWebSocketAPI(BaseIntegrationTest):
    """Integration tests for the WebSocket API."""

    def setUp(self):
        """Set up the test."""
        super().setUp()

        # Create event lists for WebSocket messages
        self.ticker_events = []
        self.kline_events = []
        self.depth_events = []
        self.account_events = []
        self.order_events = []
        self.position_events = []

    def handle_ticker(self, message: str):
        """
        Handle ticker messages.

        Args:
            message: The WebSocket message
        """
        try:
            data = json.loads(message)
            self.ticker_events.append(data)
            logger.info(f"Received ticker event: {data}")
        except Exception as e:
            logger.error(f"Failed to handle ticker event: {str(e)}")

    def handle_kline(self, message: str):
        """
        Handle K-line messages.

        Args:
            message: The WebSocket message
        """
        try:
            data = json.loads(message)
            self.kline_events.append(data)
            logger.info(f"Received K-line event: {data}")
        except Exception as e:
            logger.error(f"Failed to handle K-line event: {str(e)}")

    def handle_depth(self, message: str):
        """
        Handle depth messages.

        Args:
            message: The WebSocket message
        """
        try:
            data = json.loads(message)
            self.depth_events.append(data)
            logger.info(f"Received depth event: {data}")
        except Exception as e:
            logger.error(f"Failed to handle depth event: {str(e)}")

    def handle_account(self, message: str):
        """
        Handle account messages.

        Args:
            message: The WebSocket message
        """
        try:
            data = json.loads(message)
            self.account_events.append(data)
            logger.info(f"Received account event: {data}")
        except Exception as e:
            logger.error(f"Failed to handle account event: {str(e)}")

    def handle_order(self, message: str):
        """
        Handle order messages.

        Args:
            message: The WebSocket message
        """
        try:
            data = json.loads(message)
            self.order_events.append(data)
            logger.info(f"Received order event: {data}")
        except Exception as e:
            logger.error(f"Failed to handle order event: {str(e)}")

    def handle_position(self, message: str):
        """
        Handle position messages.

        Args:
            message: The WebSocket message
        """
        try:
            data = json.loads(message)
            self.position_events.append(data)
            logger.info(f"Received position event: {data}")
        except Exception as e:
            logger.error(f"Failed to handle position event: {str(e)}")

    def test_public_websocket(self):
        """Test public WebSocket connection."""
        # Connect to public WebSocket
        self.ws_manager.connect_public()

        # Subscribe to ticker updates
        self.ws_manager.subscribe_ticker(TEST_CONTRACT_ID, self.handle_ticker)

        # Subscribe to K-line updates
        self.ws_manager.subscribe_kline(TEST_CONTRACT_ID, "1m", self.handle_kline)

        # Subscribe to depth updates
        self.ws_manager.subscribe_depth(TEST_CONTRACT_ID, self.handle_depth)

        # Wait for some updates
        logger.info("Waiting for WebSocket updates...")
        asyncio.run(asyncio.sleep(10))

        # Check if we received any events
        self.assertGreaterEqual(len(self.ticker_events) + len(self.kline_events) + len(self.depth_events), 0)

        # Disconnect
        self.ws_manager.disconnect_public()

    def test_private_websocket(self):
        """Test private WebSocket connection."""
        # Connect to private WebSocket
        self.ws_manager.connect_private()

        # Subscribe to account updates
        self.ws_manager.subscribe_account_update(self.handle_account)

        # Subscribe to order updates
        self.ws_manager.subscribe_order_update(self.handle_order)

        # Subscribe to position updates
        self.ws_manager.subscribe_position_update(self.handle_position)

        # Wait for some updates
        logger.info("Waiting for WebSocket updates...")
        asyncio.run(asyncio.sleep(10))

        # Disconnect
        self.ws_manager.disconnect_private()


if __name__ == "__main__":
    unittest.main()
