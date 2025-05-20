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
        # The response might contain 'serverTime' or 'timeMillis'
        self.assertTrue("serverTime" in data or "timeMillis" in data, "Neither 'serverTime' nor 'timeMillis' found in response")

        # Get the time value
        time_value = data.get("serverTime") or data.get("timeMillis")
        # The time value might be an int or a string
        self.assertTrue(isinstance(time_value, int) or isinstance(time_value, str), "Time value is not an int or string")

        # Log server time
        time_value = data.get("serverTime") or data.get("timeMillis")
        logger.info(f"Server time: {time_value}")


if __name__ == "__main__":
    unittest.main()
