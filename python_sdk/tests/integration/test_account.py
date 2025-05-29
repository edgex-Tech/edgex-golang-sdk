"""Integration tests for the account API."""

import unittest
import logging
from typing import Dict, Any

from edgex_sdk import GetPositionTransactionPageParams, GetCollateralTransactionPageParams
from tests.integration.base_test import BaseIntegrationTest
from tests.integration.config import TEST_CONTRACT_ID

# Configure logging
logger = logging.getLogger(__name__)


class TestAccountAPI(BaseIntegrationTest):
    """Integration tests for the account API."""

    def test_get_account_asset(self):
        """Test get_account_asset method."""
        # Get account asset
        assets = self.run_async(self.client.get_account_asset())

        # Check response
        self.assertResponseSuccess(assets)

        # Check data
        data = assets.get("data", {})
        self.assertIsInstance(data, dict)

        # Store assets for other tests
        self.__class__.test_data["assets"] = data

        # Log asset details
        logger.info(f"Account assets: {data}")

    def test_get_account_positions(self):
        """Test get_account_positions method."""
        # Get account positions
        positions = self.run_async(self.client.get_account_positions())

        # Check response
        self.assertResponseSuccess(positions)

        # Check data (positions API returns same format as account asset)
        data = positions.get("data", {})
        self.assertIsInstance(data, dict)

        # Check position asset list
        position_assets = data.get("positionAssetList", [])
        self.assertIsInstance(position_assets, list)

        # Store positions for other tests
        self.__class__.test_data["positions"] = position_assets

        # Log position count
        logger.info(f"Found {len(position_assets)} position assets")

        # Log position details
        for position in position_assets:
            logger.info(f"Position: {position.get('contractId')} - {position.get('positionValue')}")

    def test_get_position_transaction_page(self):
        """Test get_position_transaction_page method."""
        # Create parameters
        params = GetPositionTransactionPageParams(
            size="10"
        )

        # Get position transactions
        transactions = self.run_async(self.client.account.get_position_transaction_page(params))

        # Check response
        self.assertResponseSuccess(transactions)

        # Check data
        data = transactions.get("data", {})
        self.assertIn("dataList", data)
        self.assertIsInstance(data["dataList"], list)

        # Log transaction count
        logger.info(f"Found {len(data.get('dataList', []))} position transactions")

    def test_get_collateral_transaction_page(self):
        """Test get_collateral_transaction_page method."""
        # Create parameters
        params = GetCollateralTransactionPageParams(
            size="10"
        )

        # Get collateral transactions
        transactions = self.run_async(self.client.account.get_collateral_transaction_page(params))

        # Check response
        self.assertResponseSuccess(transactions)

        # Check data
        data = transactions.get("data", {})
        self.assertIn("dataList", data)
        self.assertIsInstance(data["dataList"], list)

        # Log transaction count
        logger.info(f"Found {len(data.get('dataList', []))} collateral transactions")

    def test_get_account_by_id(self):
        """Test get_account_by_id method."""
        # Get account
        account = self.run_async(self.client.account.get_account_by_id())

        # Check response
        self.assertResponseSuccess(account)

        # Check data
        data = account.get("data", {})
        self.assertIsInstance(data, dict)

        # Check account ID (field is called "id" in response)
        self.assertIn("id", data)
        self.assertEqual(data["id"], str(self.client.internal_client.get_account_id()))

        # Log account details
        logger.info(f"Account details: {data}")


if __name__ == "__main__":
    unittest.main()
