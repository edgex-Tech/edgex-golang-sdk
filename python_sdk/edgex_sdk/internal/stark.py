"""
Stark curve implementation for signing messages.

This module provides functionality for signing messages using the Stark curve,
which is required for authentication with the EdgeX API.
"""

import binascii
import hashlib
from typing import Tuple, Optional

import sha3


class StarkCurve:
    """Implementation of the Stark curve for signing."""
    
    # Stark curve parameters
    FIELD_PRIME = 0x800000000000011000000000000000000000000000000000000000000000001
    ALPHA = 1
    BETA = 0x6f21413efbe40de150e596d72f7a8c5609ad26c15c915c1f4cdfcb99cee9e89
    N = 0x800000000000010ffffffffffffffffb781126dcae7b2321e66a241adc64d2f
    
    def __init__(self):
        """Initialize the Stark curve."""
        pass
    
    @staticmethod
    def is_on_curve(x: int, y: int) -> bool:
        """
        Check if a point is on the Stark curve.
        
        Args:
            x: The x-coordinate
            y: The y-coordinate
            
        Returns:
            bool: Whether the point is on the curve
        """
        # Check if y^2 = x^3 + alpha*x + beta (mod p)
        left = (y * y) % StarkCurve.FIELD_PRIME
        right = (x * x * x + StarkCurve.ALPHA * x + StarkCurve.BETA) % StarkCurve.FIELD_PRIME
        return left == right
    
    @staticmethod
    def add_points(p1: Optional[Tuple[int, int]], p2: Optional[Tuple[int, int]]) -> Optional[Tuple[int, int]]:
        """
        Add two points on the Stark curve.
        
        Args:
            p1: The first point (x, y) or None for the point at infinity
            p2: The second point (x, y) or None for the point at infinity
            
        Returns:
            Optional[Tuple[int, int]]: The resulting point or None for the point at infinity
        """
        if p1 is None:
            return p2
        if p2 is None:
            return p1
        
        x1, y1 = p1
        x2, y2 = p2
        
        # Handle point doubling
        if x1 == x2 and y1 == y2:
            # Calculate slope for point doubling
            # slope = (3 * x1^2 + alpha) / (2 * y1)
            numerator = (3 * x1 * x1 + StarkCurve.ALPHA) % StarkCurve.FIELD_PRIME
            denominator = (2 * y1) % StarkCurve.FIELD_PRIME
            
            # Calculate modular inverse of denominator
            denominator_inv = pow(denominator, StarkCurve.FIELD_PRIME - 2, StarkCurve.FIELD_PRIME)
            slope = (numerator * denominator_inv) % StarkCurve.FIELD_PRIME
        elif x1 == x2:
            # If x1 == x2 but y1 != y2, the result is the point at infinity
            return None
        else:
            # Calculate slope for point addition
            # slope = (y2 - y1) / (x2 - x1)
            numerator = (y2 - y1) % StarkCurve.FIELD_PRIME
            denominator = (x2 - x1) % StarkCurve.FIELD_PRIME
            
            # Calculate modular inverse of denominator
            denominator_inv = pow(denominator, StarkCurve.FIELD_PRIME - 2, StarkCurve.FIELD_PRIME)
            slope = (numerator * denominator_inv) % StarkCurve.FIELD_PRIME
        
        # Calculate new point
        x3 = (slope * slope - x1 - x2) % StarkCurve.FIELD_PRIME
        y3 = (slope * (x1 - x3) - y1) % StarkCurve.FIELD_PRIME
        
        return (x3, y3)
    
    @staticmethod
    def scalar_multiply(k: int, point: Tuple[int, int]) -> Optional[Tuple[int, int]]:
        """
        Multiply a point by a scalar.
        
        Args:
            k: The scalar
            point: The point (x, y)
            
        Returns:
            Optional[Tuple[int, int]]: The resulting point or None for the point at infinity
        """
        if k == 0:
            return None
        
        result = None
        addend = point
        
        while k:
            if k & 1:
                result = StarkCurve.add_points(result, addend)
            addend = StarkCurve.add_points(addend, addend)
            k >>= 1
        
        return result
    
    @staticmethod
    def get_public_key(private_key: bytes) -> Tuple[int, int]:
        """
        Get the public key from a private key.
        
        Args:
            private_key: The private key as bytes
            
        Returns:
            Tuple[int, int]: The public key as (x, y) coordinates
        """
        # Convert private key to integer
        priv_key_int = int.from_bytes(private_key, byteorder='big')
        
        # Ensure the private key is in the valid range
        priv_key_int = priv_key_int % StarkCurve.N
        
        # Generator point for Stark curve
        # Note: This is a placeholder. The actual generator point should be used.
        generator_x = 0x1ef15c18599971b7beced415a40f0c7deacfd9b0d1819e03d723d8bc943cfca
        generator_y = 0x5668060aa49730b7be4801df46ec62de53ecd11abe43a32873000c36e8dc1f
        generator = (generator_x, generator_y)
        
        # Perform scalar multiplication
        public_key = StarkCurve.scalar_multiply(priv_key_int, generator)
        
        if public_key is None:
            raise ValueError("Invalid private key")
        
        return public_key
    
    @staticmethod
    def sign(private_key: bytes, message_hash: bytes) -> Tuple[int, int]:
        """
        Sign a message hash using the private key.
        
        Args:
            private_key: The private key as bytes
            message_hash: The hash of the message to sign
            
        Returns:
            Tuple[int, int]: The signature as (r, s) values
            
        Raises:
            ValueError: If the private key is invalid
        """
        # Convert private key and message hash to integers
        priv_key_int = int.from_bytes(private_key, byteorder='big')
        msg_hash_int = int.from_bytes(message_hash, byteorder='big')
        
        # Ensure the private key is in the valid range
        priv_key_int = priv_key_int % StarkCurve.N
        
        # Ensure the message hash is in the valid range
        msg_hash_int = msg_hash_int % StarkCurve.N
        
        # Generate a deterministic k value
        # This is a simplified implementation; a more secure method should be used in production
        k_bytes = hashlib.sha256(private_key + message_hash).digest()
        k = int.from_bytes(k_bytes, byteorder='big') % StarkCurve.N
        
        # Generator point for Stark curve
        generator_x = 0x1ef15c18599971b7beced415a40f0c7deacfd9b0d1819e03d723d8bc943cfca
        generator_y = 0x5668060aa49730b7be4801df46ec62de53ecd11abe43a32873000c36e8dc1f
        generator = (generator_x, generator_y)
        
        # Calculate r = x-coordinate of k*G
        point_r = StarkCurve.scalar_multiply(k, generator)
        if point_r is None:
            # If k*G is the point at infinity, try again with a different k
            return StarkCurve.sign(private_key, message_hash)
        
        r = point_r[0]
        
        # Calculate s = (message_hash + r * private_key) / k mod n
        s = ((msg_hash_int + r * priv_key_int) * pow(k, StarkCurve.N - 2, StarkCurve.N)) % StarkCurve.N
        
        return (r, s)


def sign_message(private_key_hex: str, message_hash: bytes) -> Tuple[str, str]:
    """
    Sign a message hash using a private key in hex format.
    
    Args:
        private_key_hex: The private key as a hex string
        message_hash: The hash of the message to sign
        
    Returns:
        Tuple[str, str]: The signature as (r, s) hex strings
        
    Raises:
        ValueError: If the private key is invalid
    """
    try:
        private_key = binascii.unhexlify(private_key_hex)
    except binascii.Error:
        raise ValueError("Invalid private key hex string")
    
    r, s = StarkCurve().sign(private_key, message_hash)
    
    # Convert r and s to 32-byte hex strings
    r_hex = format(r, '064x')
    s_hex = format(s, '064x')
    
    return (r_hex, s_hex)


def get_public_key(private_key_hex: str) -> Tuple[str, str]:
    """
    Get the public key from a private key in hex format.
    
    Args:
        private_key_hex: The private key as a hex string
        
    Returns:
        Tuple[str, str]: The public key as (x, y) hex strings
        
    Raises:
        ValueError: If the private key is invalid
    """
    try:
        private_key = binascii.unhexlify(private_key_hex)
    except binascii.Error:
        raise ValueError("Invalid private key hex string")
    
    x, y = StarkCurve().get_public_key(private_key)
    
    # Convert x and y to 32-byte hex strings
    x_hex = format(x, '064x')
    y_hex = format(y, '064x')
    
    return (x_hex, y_hex)


def verify_signature(public_key: Tuple[int, int], message_hash: bytes, signature: Tuple[int, int]) -> bool:
    """
    Verify a signature using a public key.
    
    Args:
        public_key: The public key as (x, y) coordinates
        message_hash: The hash of the message
        signature: The signature as (r, s) values
        
    Returns:
        bool: Whether the signature is valid
    """
    # Convert message hash to integer
    msg_hash_int = int.from_bytes(message_hash, byteorder='big')
    
    # Ensure the message hash is in the valid range
    msg_hash_int = msg_hash_int % StarkCurve.N
    
    # Extract r and s from signature
    r, s = signature
    
    # Check if r and s are in the valid range
    if r <= 0 or r >= StarkCurve.N or s <= 0 or s >= StarkCurve.N:
        return False
    
    # Calculate s_inv = s^-1 mod n
    s_inv = pow(s, StarkCurve.N - 2, StarkCurve.N)
    
    # Calculate u1 = message_hash * s_inv mod n
    u1 = (msg_hash_int * s_inv) % StarkCurve.N
    
    # Calculate u2 = r * s_inv mod n
    u2 = (r * s_inv) % StarkCurve.N
    
    # Generator point for Stark curve
    generator_x = 0x1ef15c18599971b7beced415a40f0c7deacfd9b0d1819e03d723d8bc943cfca
    generator_y = 0x5668060aa49730b7be4801df46ec62de53ecd11abe43a32873000c36e8dc1f
    generator = (generator_x, generator_y)
    
    # Calculate u1*G + u2*Q
    point1 = StarkCurve.scalar_multiply(u1, generator)
    point2 = StarkCurve.scalar_multiply(u2, public_key)
    point = StarkCurve.add_points(point1, point2)
    
    if point is None:
        return False
    
    # The signature is valid if the x-coordinate of the resulting point equals r
    return point[0] == r
