"""
Unit tests for the internal client.
"""

import unittest
import binascii
from unittest.mock import patch, MagicMock

from edgex_sdk.internal.client import Client, L2Signature
from edgex_sdk.internal.mock_signing_adapter import MockSigningAdapter


class TestInternalClient(unittest.TestCase):
    """Test cases for the internal client."""

    def setUp(self):
        """Set up test fixtures."""
        self.base_url = "https://testnet.edgex.exchange"
        self.account_id = 12345
        self.stark_pri_key = "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"
        self.client = Client(
            base_url=self.base_url,
            account_id=self.account_id,
            stark_pri_key=self.stark_pri_key
        )

    def test_init(self):
        """Test client initialization."""
        self.assertEqual(self.client.base_url, self.base_url)
        self.assertEqual(self.client.get_account_id(), self.account_id)
        self.assertEqual(self.client.get_stark_pri_key(), self.stark_pri_key)

    def test_get_account_id(self):
        """Test get_account_id method."""
        self.assertEqual(self.client.get_account_id(), self.account_id)

    def test_get_stark_pri_key(self):
        """Test get_stark_pri_key method."""
        self.assertEqual(self.client.get_stark_pri_key(), self.stark_pri_key)

    def test_sign(self):
        """Test sign method."""
        # Create a mock signing adapter
        mock_adapter = MockSigningAdapter()
        mock_adapter.sign = MagicMock(return_value=("r_value", "s_value"))

        # Set the mock adapter
        self.client.signing_adapter = mock_adapter

        # Create a message hash
        message_hash = b"test message hash"

        # Sign the message
        signature = self.client.sign(message_hash)

        # Check that sign was called with the correct arguments
        mock_adapter.sign.assert_called_once_with(message_hash, self.stark_pri_key)

        # Check the signature
        self.assertEqual(signature.r, "r_value")
        self.assertEqual(signature.s, "s_value")
        self.assertEqual(signature.v, "")

    def test_sign_no_key(self):
        """Test sign method with no private key."""
        # Create a client with no private key
        client = Client(
            base_url=self.base_url,
            account_id=self.account_id,
            stark_pri_key=""
        )

        # Try to sign a message
        with self.assertRaises(ValueError):
            client.sign(b"test message hash")

    def test_generate_uuid(self):
        """Test generate_uuid method."""
        uuid1 = self.client.generate_uuid()
        uuid2 = self.client.generate_uuid()

        # Check that UUIDs are strings
        self.assertIsInstance(uuid1, str)
        self.assertIsInstance(uuid2, str)

        # Check that UUIDs are different
        self.assertNotEqual(uuid1, uuid2)

    def test_calc_nonce(self):
        """Test calc_nonce method."""
        client_order_id = "test-order-id"
        nonce = self.client.calc_nonce(client_order_id)

        # Check that nonce is an integer
        self.assertIsInstance(nonce, int)

        # Check that the same client order ID produces the same nonce
        nonce2 = self.client.calc_nonce(client_order_id)
        self.assertEqual(nonce, nonce2)

        # Check that different client order IDs produce different nonces
        nonce3 = self.client.calc_nonce("different-order-id")
        self.assertNotEqual(nonce, nonce3)

    def test_calc_limit_order_hash(self):
        """Test calc_limit_order_hash method."""
        # Test parameters
        synthetic_asset_id = "synthetic_asset_id"
        collateral_asset_id = "collateral_asset_id"
        fee_asset_id = "fee_asset_id"
        is_buy = True
        amount_synthetic = 1000
        amount_collateral = 2000
        amount_fee = 10
        nonce = 12345
        account_id = 67890
        expire_time = 1234567890

        # Calculate hash
        hash_bytes = self.client.calc_limit_order_hash(
            synthetic_asset_id,
            collateral_asset_id,
            fee_asset_id,
            is_buy,
            amount_synthetic,
            amount_collateral,
            amount_fee,
            nonce,
            account_id,
            expire_time
        )

        # Check that hash is bytes
        self.assertIsInstance(hash_bytes, bytes)

        # Check that hash is not empty
        self.assertTrue(len(hash_bytes) > 0)


if __name__ == '__main__':
    unittest.main()
