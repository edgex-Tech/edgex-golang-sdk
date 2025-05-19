"""Integration tests for the quote API."""

import unittest
import logging
from typing import Dict, Any

from edgex_sdk import GetKLineParams, GetOrderBookDepthParams, GetMultiContractKLineParams
from tests.integration.base_test import BaseIntegrationTest
from tests.integration.config import TEST_CONTRACT_ID

# Configure logging
logger = logging.getLogger(__name__)


class TestQuoteAPI(BaseIntegrationTest):
    """Integration tests for the quote API."""

    def test_get_24_hour_quote(self):
        """Test get_24_hour_quote method."""
        # Get 24-hour quote
        quote = self.run_async(self.client.quote.get_24_hour_quote(TEST_CONTRACT_ID))

        # Check response
        self.assertResponseSuccess(quote)

        # Check data
        data = quote.get("data", [])
        self.assertIsInstance(data, list)
        self.assertGreater(len(data), 0)

        # Check first quote
        first_quote = data[0]
        self.assertIn("contractId", first_quote)
        self.assertEqual(first_quote["contractId"], TEST_CONTRACT_ID)

        # Log quote details
        logger.info(f"24-hour quote for {TEST_CONTRACT_ID}: {first_quote.get('lastPrice')}")

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
        self.assertIn("list", data)
        self.assertIsInstance(data["list"], list)

        # Check K-line count
        kline_list = data["list"]
        self.assertLessEqual(len(kline_list), 10)

        # Check first K-line
        if kline_list:
            first_kline = kline_list[0]
            self.assertIn("open", first_kline)
            self.assertIn("high", first_kline)
            self.assertIn("low", first_kline)
            self.assertIn("close", first_kline)
            self.assertIn("volume", first_kline)

            # Log K-line details
            logger.info(f"First K-line for {TEST_CONTRACT_ID}: {first_kline}")

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

        # Check ask and bid count
        asks = data["asks"]
        bids = data["bids"]
        self.assertLessEqual(len(asks), 10)
        self.assertLessEqual(len(bids), 10)

        # Log depth details
        logger.info(f"Order book depth for {TEST_CONTRACT_ID}: {len(asks)} asks, {len(bids)} bids")

    def test_get_multi_contract_k_line(self):
        """Test get_multi_contract_k_line method."""
        # Create parameters
        params = GetMultiContractKLineParams(
            contract_id_list=[TEST_CONTRACT_ID],
            interval="1m",
            limit=1
        )

        # Get multi-contract K-line data
        klines = self.run_async(self.client.quote.get_multi_contract_k_line(params))

        # Check response
        self.assertResponseSuccess(klines)

        # Check data
        data = klines.get("data", {})
        self.assertIn("list", data)
        self.assertIsInstance(data["list"], list)

        # Check K-line count
        kline_list = data["list"]
        self.assertLessEqual(len(kline_list), 1)

        # Check first K-line
        if kline_list:
            first_kline = kline_list[0]
            self.assertIn("contractId", first_kline)
            self.assertEqual(first_kline["contractId"], TEST_CONTRACT_ID)

            # Log K-line details
            logger.info(f"Multi-contract K-line for {TEST_CONTRACT_ID}: {first_kline}")


if __name__ == "__main__":
    unittest.main()
