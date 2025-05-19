"""
Unit tests for the StarkEx signing adapter.
"""

import unittest
import binascii
import hashlib

from edgex_sdk.internal.starkex_signing_adapter import StarkExSigningAdapter


class TestStarkExSigningAdapter(unittest.TestCase):
    """Test cases for the StarkEx signing adapter."""

    def setUp(self):
        """Set up test fixtures."""
        self.adapter = StarkExSigningAdapter()

        # Test private key (32 bytes)
        self.private_key_hex = "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"

        # Test message hash (32 bytes)
        self.message_hash_hex = "0000000000000000000000000000000000000000000000000000000000000001"
        self.message_hash = binascii.unhexlify(self.message_hash_hex)

    def test_sign_and_verify(self):
        """Test sign and verify methods."""
        # Sign the message
        r, s = self.adapter.sign(self.message_hash, self.private_key_hex)

        # Check that r and s are hex strings of length 64
        self.assertIsInstance(r, str)
        self.assertIsInstance(s, str)
        self.assertEqual(len(r), 64)
        self.assertEqual(len(s), 64)

        # Check that r and s are valid hex strings
        try:
            int(r, 16)
            int(s, 16)
        except ValueError:
            self.fail("r or s is not a valid hex string")

        # Get the public key
        public_key = self.adapter.get_public_key(self.private_key_hex)

        # Check that the public key is a hex string of length 64
        self.assertIsInstance(public_key, str)
        self.assertEqual(len(public_key), 64)

        # Verify the signature
        result = self.adapter.verify(self.message_hash, (r, s), public_key)

        # Check that the signature is valid
        self.assertTrue(result)

    def test_sign_different_messages(self):
        """Test signing different messages produces different signatures."""
        # Sign the first message
        r1, s1 = self.adapter.sign(self.message_hash, self.private_key_hex)

        # Create a different message hash
        message_hash2 = hashlib.sha256(b"different message").digest()

        # Sign the second message
        r2, s2 = self.adapter.sign(message_hash2, self.private_key_hex)

        # Check that the signatures are different
        self.assertNotEqual((r1, s1), (r2, s2))

    def test_sign_different_keys(self):
        """Test signing the same message with different keys produces different signatures."""
        # Sign with the first key
        r1, s1 = self.adapter.sign(self.message_hash, self.private_key_hex)

        # Create a different private key
        private_key2 = "fedcba9876543210fedcba9876543210fedcba9876543210fedcba9876543210"

        # Sign with the second key
        r2, s2 = self.adapter.sign(self.message_hash, private_key2)

        # Check that the signatures are different
        self.assertNotEqual((r1, s1), (r2, s2))

    def test_verify_invalid_signature(self):
        """Test verifying an invalid signature."""
        # Sign the message
        r, s = self.adapter.sign(self.message_hash, self.private_key_hex)

        # Get the public key
        public_key = self.adapter.get_public_key(self.private_key_hex)

        # Modify the signature
        r_invalid = format(int(r, 16) + 1, '064x')

        # Verify the invalid signature
        result = self.adapter.verify(self.message_hash, (r_invalid, s), public_key)

        # Check that the signature is invalid
        self.assertFalse(result)

    def test_verify_wrong_message(self):
        """Test verifying a signature with the wrong message."""
        # Sign the message
        r, s = self.adapter.sign(self.message_hash, self.private_key_hex)

        # Get the public key
        public_key = self.adapter.get_public_key(self.private_key_hex)

        # Create a different message hash
        message_hash2 = hashlib.sha256(b"different message").digest()

        # Verify the signature with the wrong message
        result = self.adapter.verify(message_hash2, (r, s), public_key)

        # Check that the signature is invalid
        self.assertFalse(result)

    def test_verify_wrong_public_key(self):
        """Test verifying a signature with the wrong public key."""
        # Sign the message
        r, s = self.adapter.sign(self.message_hash, self.private_key_hex)

        # Create a different private key
        private_key2 = "fedcba9876543210fedcba9876543210fedcba9876543210fedcba9876543210"

        # Get the wrong public key
        wrong_public_key = self.adapter.get_public_key(private_key2)

        # Verify the signature with the wrong public key
        result = self.adapter.verify(self.message_hash, (r, s), wrong_public_key)

        # Check that the signature is invalid
        self.assertFalse(result)

    def test_invalid_private_key(self):
        """Test signing with an invalid private key."""
        # Invalid hex string
        with self.assertRaises(ValueError):
            self.adapter.sign(self.message_hash, "not a hex string")

        # Private key with value 0 should be handled
        try:
            self.adapter.sign(self.message_hash, "0" * 64)
        except ValueError:
            self.fail("sign() raised ValueError unexpectedly with private key of all zeros")

    def test_invalid_message_hash(self):
        """Test signing with an invalid message hash."""
        # Message hash out of range should be handled
        huge_hash = b"\xff" * 32
        try:
            self.adapter.sign(huge_hash, self.private_key_hex)
        except ValueError:
            self.fail("sign() raised ValueError unexpectedly with large message hash")


if __name__ == '__main__':
    unittest.main()
