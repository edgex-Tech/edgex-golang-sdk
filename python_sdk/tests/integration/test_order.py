"""Integration tests for the order API."""

import unittest
import logging
from typing import Dict, Any
from decimal import Decimal

from edgex_sdk import (
    OrderSide,
    OrderType,
    TimeInForce,
    CreateOrderParams,
    CancelOrderParams,
    GetActiveOrderParams,
    OrderFillTransactionParams
)
from tests.integration.base_test import BaseIntegrationTest
from tests.integration.config import TEST_CONTRACT_ID, TEST_ORDER_SIZE, TEST_ORDER_PRICE

# Configure logging
logger = logging.getLogger(__name__)


class TestOrderAPI(BaseIntegrationTest):
    """Integration tests for the order API."""

    def test_get_max_order_size(self):
        """Test get_max_order_size method."""
        # Get max order size
        max_size = self.run_async(self.client.get_max_order_size(TEST_CONTRACT_ID, Decimal(TEST_ORDER_PRICE)))

        # Check response
        self.assertResponseSuccess(max_size)

        # Check data
        data = max_size.get("data", {})
        self.assertIn("maxBuySize", data)
        self.assertIn("maxSellSize", data)

        # Log max order size
        logger.info(f"Max buy size: {data.get('maxBuySize')}, Max sell size: {data.get('maxSellSize')}")

    def test_get_active_orders(self):
        """Test get_active_orders method."""
        # Create parameters
        params = GetActiveOrderParams(
            size="10"
        )

        # Get active orders
        orders = self.run_async(self.client.get_active_orders(params))

        # Check response
        self.assertResponseSuccess(orders)

        # Check data
        data = orders.get("data", {})
        self.assertIn("dataList", data)
        self.assertIsInstance(data["dataList"], list)

        # Store active orders for other tests
        self.__class__.test_data["active_orders"] = data.get("dataList", [])

        # Log order count
        logger.info(f"Found {len(data.get('dataList', []))} active orders")

    def test_get_order_fill_transactions(self):
        """Test get_order_fill_transactions method."""
        # Create parameters
        params = OrderFillTransactionParams(
            size="10"
        )

        # Get order fill transactions
        transactions = self.run_async(self.client.get_order_fill_transactions(params))

        # Check response
        self.assertResponseSuccess(transactions)

        # Check data
        data = transactions.get("data", {})
        self.assertIn("dataList", data)
        self.assertIsInstance(data["dataList"], list)

        # Log transaction count
        logger.info(f"Found {len(data.get('dataList', []))} order fill transactions")

    @unittest.skip("Skipping order creation to avoid actual orders")
    def test_create_and_cancel_order(self):
        """Test create_order and cancel_order methods."""
        # Create order parameters
        params = CreateOrderParams(
            contract_id=TEST_CONTRACT_ID,
            size=TEST_ORDER_SIZE,
            price=TEST_ORDER_PRICE,
            type=OrderType.LIMIT,
            side=OrderSide.BUY,
            time_in_force=TimeInForce.GOOD_TIL_CANCEL
        )

        # Create order
        order = self.run_async(self.client.create_order(params))

        # Check response
        self.assertResponseSuccess(order)

        # Check data
        data = order.get("data", {})
        self.assertIn("orderId", data)

        # Store order ID
        order_id = data["orderId"]

        # Log order details
        logger.info(f"Created order: {order_id}")

        # Cancel order
        cancel_params = CancelOrderParams(
            order_id=order_id
        )

        # Cancel order
        cancel = self.run_async(self.client.cancel_order(cancel_params))

        # Check response
        self.assertResponseSuccess(cancel)

        # Log cancellation details
        logger.info(f"Cancelled order: {order_id}")


if __name__ == "__main__":
    unittest.main()
