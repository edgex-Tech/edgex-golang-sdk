"""
Unit tests for the Stark curve implementation.
"""

import unittest
import binascii

from edgex_sdk.internal.stark import StarkCurve, sign_message, get_public_key, verify_signature


class TestStarkCurve(unittest.TestCase):
    """Test cases for the Stark curve implementation."""
    
    def setUp(self):
        """Set up test fixtures."""
        self.stark_curve = StarkCurve()
        
        # Test private key (32 bytes)
        self.private_key_hex = "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"
        self.private_key = binascii.unhexlify(self.private_key_hex)
        
        # Test message hash (32 bytes)
        self.message_hash_hex = "0000000000000000000000000000000000000000000000000000000000000001"
        self.message_hash = binascii.unhexlify(self.message_hash_hex)
    
    def test_is_on_curve(self):
        """Test is_on_curve method."""
        # Generator point for Stark curve
        generator_x = 0x1ef15c18599971b7beced415a40f0c7deacfd9b0d1819e03d723d8bc943cfca
        generator_y = 0x5668060aa49730b7be4801df46ec62de53ecd11abe43a32873000c36e8dc1f
        
        # Check that generator point is on the curve
        self.assertTrue(self.stark_curve.is_on_curve(generator_x, generator_y))
        
        # Check that a random point is not on the curve
        self.assertFalse(self.stark_curve.is_on_curve(123, 456))
    
    def test_add_points(self):
        """Test add_points method."""
        # Generator point for Stark curve
        generator_x = 0x1ef15c18599971b7beced415a40f0c7deacfd9b0d1819e03d723d8bc943cfca
        generator_y = 0x5668060aa49730b7be4801df46ec62de53ecd11abe43a32873000c36e8dc1f
        generator = (generator_x, generator_y)
        
        # Add point to itself
        result = self.stark_curve.add_points(generator, generator)
        
        # Check that result is a tuple of two integers
        self.assertIsInstance(result, tuple)
        self.assertEqual(len(result), 2)
        self.assertIsInstance(result[0], int)
        self.assertIsInstance(result[1], int)
        
        # Check that result is on the curve
        self.assertTrue(self.stark_curve.is_on_curve(result[0], result[1]))
        
        # Add point to None (point at infinity)
        result = self.stark_curve.add_points(generator, None)
        self.assertEqual(result, generator)
        
        # Add None to point
        result = self.stark_curve.add_points(None, generator)
        self.assertEqual(result, generator)
    
    def test_scalar_multiply(self):
        """Test scalar_multiply method."""
        # Generator point for Stark curve
        generator_x = 0x1ef15c18599971b7beced415a40f0c7deacfd9b0d1819e03d723d8bc943cfca
        generator_y = 0x5668060aa49730b7be4801df46ec62de53ecd11abe43a32873000c36e8dc1f
        generator = (generator_x, generator_y)
        
        # Multiply by 0
        result = self.stark_curve.scalar_multiply(0, generator)
        self.assertIsNone(result)
        
        # Multiply by 1
        result = self.stark_curve.scalar_multiply(1, generator)
        self.assertEqual(result, generator)
        
        # Multiply by 2
        result = self.stark_curve.scalar_multiply(2, generator)
        
        # Check that result is a tuple of two integers
        self.assertIsInstance(result, tuple)
        self.assertEqual(len(result), 2)
        self.assertIsInstance(result[0], int)
        self.assertIsInstance(result[1], int)
        
        # Check that result is on the curve
        self.assertTrue(self.stark_curve.is_on_curve(result[0], result[1]))
    
    def test_get_public_key(self):
        """Test get_public_key method."""
        # Get public key
        public_key = self.stark_curve.get_public_key(self.private_key)
        
        # Check that public key is a tuple of two integers
        self.assertIsInstance(public_key, tuple)
        self.assertEqual(len(public_key), 2)
        self.assertIsInstance(public_key[0], int)
        self.assertIsInstance(public_key[1], int)
        
        # Check that public key is on the curve
        self.assertTrue(self.stark_curve.is_on_curve(public_key[0], public_key[1]))
    
    def test_sign(self):
        """Test sign method."""
        # Sign message
        signature = self.stark_curve.sign(self.private_key, self.message_hash)
        
        # Check that signature is a tuple of two integers
        self.assertIsInstance(signature, tuple)
        self.assertEqual(len(signature), 2)
        self.assertIsInstance(signature[0], int)
        self.assertIsInstance(signature[1], int)
        
        # Check that signature components are in the valid range
        self.assertTrue(0 < signature[0] < self.stark_curve.N)
        self.assertTrue(0 < signature[1] < self.stark_curve.N)
    
    def test_sign_message(self):
        """Test sign_message function."""
        # Sign message
        r, s = sign_message(self.private_key_hex, self.message_hash)
        
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
    
    def test_get_public_key_function(self):
        """Test get_public_key function."""
        # Get public key
        x, y = get_public_key(self.private_key_hex)
        
        # Check that x and y are hex strings of length 64
        self.assertIsInstance(x, str)
        self.assertIsInstance(y, str)
        self.assertEqual(len(x), 64)
        self.assertEqual(len(y), 64)
        
        # Check that x and y are valid hex strings
        try:
            x_int = int(x, 16)
            y_int = int(y, 16)
        except ValueError:
            self.fail("x or y is not a valid hex string")
        
        # Check that the point is on the curve
        self.assertTrue(self.stark_curve.is_on_curve(x_int, y_int))
    
    def test_verify_signature(self):
        """Test verify_signature function."""
        # Get public key
        public_key = self.stark_curve.get_public_key(self.private_key)
        
        # Sign message
        signature = self.stark_curve.sign(self.private_key, self.message_hash)
        
        # Verify signature
        result = verify_signature(public_key, self.message_hash, signature)
        
        # Check that signature is valid
        self.assertTrue(result)
        
        # Verify with wrong message
        wrong_message = b"wrong message"
        result = verify_signature(public_key, wrong_message, signature)
        
        # Check that signature is invalid
        self.assertFalse(result)
        
        # Verify with wrong signature
        wrong_signature = (signature[0], signature[1] + 1)
        result = verify_signature(public_key, self.message_hash, wrong_signature)
        
        # Check that signature is invalid
        self.assertFalse(result)


if __name__ == '__main__':
    unittest.main()
