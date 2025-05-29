#!/usr/bin/env python3
"""
Test script for the Pedersen hash implementation.

This script tests the new Pedersen hash implementation to ensure it works correctly.
"""

import sys
import os

# Add the edgex_sdk to the path
sys.path.insert(0, os.path.join(os.path.dirname(__file__)))

# Import directly from the crypto module
from edgex_sdk.crypto.pedersen_hash import pedersen_hash, pedersen_hash_as_point, pedersen_hash_bytes


def test_basic_pedersen_hash():
    """Test basic Pedersen hash functionality."""
    print("Testing basic Pedersen hash...")
    
    # Test with simple values
    try:
        result = pedersen_hash(123, 456)
        print(f"pedersen_hash(123, 456) = {hex(result)}")
        
        # Test with single value
        result_single = pedersen_hash(789)
        print(f"pedersen_hash(789) = {hex(result_single)}")
        
        # Test as point
        point = pedersen_hash_as_point(123, 456)
        print(f"pedersen_hash_as_point(123, 456) = ({hex(point[0])}, {hex(point[1])})")
        
        # Test as bytes
        result_bytes = pedersen_hash_bytes(123, 456)
        print(f"pedersen_hash_bytes(123, 456) = {result_bytes.hex()}")
        
        print("✓ Basic Pedersen hash tests passed")
        return True
        
    except Exception as e:
        print(f"✗ Basic Pedersen hash test failed: {e}")
        return False


def test_consistency():
    """Test consistency of different hash functions."""
    print("\nTesting consistency...")

    try:
        # Test that pedersen_hash returns the x-coordinate of pedersen_hash_as_point
        elements = [123, 456]
        hash_result = pedersen_hash(*elements)
        point_result = pedersen_hash_as_point(*elements)

        if hash_result == point_result[0]:
            print("✓ pedersen_hash and pedersen_hash_as_point are consistent")
        else:
            print("✗ pedersen_hash and pedersen_hash_as_point are inconsistent")
            return False

        # Test that pedersen_hash_bytes returns the same as pedersen_hash
        bytes_result = pedersen_hash_bytes(*elements)
        expected_bytes = hash_result.to_bytes(32, byteorder='big')

        if bytes_result == expected_bytes:
            print("✓ pedersen_hash_bytes and pedersen_hash are consistent")
            return True
        else:
            print("✗ pedersen_hash_bytes and pedersen_hash are inconsistent")
            return False

    except Exception as e:
        print(f"✗ Consistency test failed: {e}")
        return False


def test_edge_cases():
    """Test edge cases and error conditions."""
    print("\nTesting edge cases...")
    
    try:
        # Test with zero
        result_zero = pedersen_hash(0)
        print(f"pedersen_hash(0) = {hex(result_zero)}")
        
        # Test with large number (but within field)
        large_num = 2**250  # Should be within FIELD_PRIME
        result_large = pedersen_hash(large_num)
        print(f"pedersen_hash({large_num}) = {hex(result_large)}")
        
        # Test with bytes input
        result_bytes_input = pedersen_hash_bytes(b'\x01\x02\x03', 456)
        print(f"pedersen_hash_bytes(b'\\x01\\x02\\x03', 456) = {result_bytes_input.hex()}")
        
        print("✓ Edge case tests passed")
        return True
        
    except Exception as e:
        print(f"✗ Edge case test failed: {e}")
        return False


def test_error_conditions():
    """Test error conditions."""
    print("\nTesting error conditions...")
    
    try:
        # Test with number too large
        from edgex_sdk.crypto.constants import FIELD_PRIME
        
        try:
            pedersen_hash(FIELD_PRIME + 1)
            print("✗ Should have failed with number too large")
            return False
        except ValueError as e:
            print(f"✓ Correctly caught error for large number: {e}")
        
        # Test with negative number
        try:
            pedersen_hash(-1)
            print("✗ Should have failed with negative number")
            return False
        except ValueError as e:
            print(f"✓ Correctly caught error for negative number: {e}")
        
        # Test with invalid bytes input
        try:
            pedersen_hash_bytes(b'x' * 33)  # Too long
            print("✗ Should have failed with bytes too long")
            return False
        except ValueError as e:
            print(f"✓ Correctly caught error for long bytes: {e}")
        
        print("✓ Error condition tests passed")
        return True
        
    except Exception as e:
        print(f"✗ Error condition test failed: {e}")
        return False


def main():
    """Run all tests."""
    print("Running Pedersen hash implementation tests...\n")
    
    tests = [
        test_basic_pedersen_hash,
        test_consistency,
        test_edge_cases,
        test_error_conditions,
    ]
    
    passed = 0
    total = len(tests)
    
    for test in tests:
        if test():
            passed += 1
    
    print(f"\n{'='*50}")
    print(f"Test Results: {passed}/{total} tests passed")
    
    if passed == total:
        print("🎉 All tests passed!")
        return 0
    else:
        print("❌ Some tests failed!")
        return 1


if __name__ == "__main__":
    sys.exit(main())
