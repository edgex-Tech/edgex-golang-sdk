#!/usr/bin/env python3
"""
Direct test script for the Pedersen hash implementation.

This script tests the new Pedersen hash implementation by importing modules directly.
"""

import sys
import os

# Add the current directory to the path
current_dir = os.path.dirname(__file__)
sys.path.insert(0, current_dir)

# Import the modules directly without going through __init__.py
import importlib.util

def load_module_from_path(module_name, file_path):
    """Load a module from a file path."""
    spec = importlib.util.spec_from_file_location(module_name, file_path)
    module = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(module)
    return module

# Add the crypto directory to the path so modules can import each other
crypto_dir = os.path.join(current_dir, 'edgex_sdk', 'crypto')
sys.path.insert(0, crypto_dir)

# Load the constants module
constants_path = os.path.join(crypto_dir, 'constants.py')
constants = load_module_from_path('constants', constants_path)

# Load the pedersen_hash module
pedersen_path = os.path.join(crypto_dir, 'pedersen_hash.py')
pedersen = load_module_from_path('pedersen_hash', pedersen_path)


def test_constants():
    """Test that constants are loaded correctly."""
    print("Testing constants...")
    
    try:
        print(f"FIELD_PRIME = {hex(constants.FIELD_PRIME)}")
        print(f"SHIFT_POINT = [{hex(constants.SHIFT_POINT[0])}, {hex(constants.SHIFT_POINT[1])}]")
        print(f"Number of CONSTANT_POINTS = {len(constants.CONSTANT_POINTS)}")
        
        # Verify shift point is the first constant point
        if constants.SHIFT_POINT == constants.CONSTANT_POINTS[0]:
            print("✓ SHIFT_POINT matches first CONSTANT_POINT")
        else:
            print("✗ SHIFT_POINT doesn't match first CONSTANT_POINT")
            return False
        
        print("✓ Constants test passed")
        return True
        
    except Exception as e:
        print(f"✗ Constants test failed: {e}")
        import traceback
        traceback.print_exc()
        return False


def test_basic_pedersen_hash():
    """Test basic Pedersen hash functionality."""
    print("\nTesting basic Pedersen hash...")
    
    try:
        # Test with simple values
        result = pedersen.pedersen_hash(123, 456)
        print(f"pedersen_hash(123, 456) = {hex(result)}")
        
        # Test with single value
        result_single = pedersen.pedersen_hash(789)
        print(f"pedersen_hash(789) = {hex(result_single)}")
        
        # Test as point
        point = pedersen.pedersen_hash_as_point(123, 456)
        print(f"pedersen_hash_as_point(123, 456) = ({hex(point[0])}, {hex(point[1])})")
        
        # Test as bytes
        result_bytes = pedersen.pedersen_hash_bytes(123, 456)
        print(f"pedersen_hash_bytes(123, 456) = {result_bytes.hex()}")
        
        print("✓ Basic Pedersen hash tests passed")
        return True
        
    except Exception as e:
        print(f"✗ Basic Pedersen hash test failed: {e}")
        import traceback
        traceback.print_exc()
        return False


def test_consistency():
    """Test consistency of different hash functions."""
    print("\nTesting consistency...")
    
    try:
        # Test that pedersen_hash returns the x-coordinate of pedersen_hash_as_point
        elements = [123, 456]
        hash_result = pedersen.pedersen_hash(*elements)
        point_result = pedersen.pedersen_hash_as_point(*elements)
        
        if hash_result == point_result[0]:
            print("✓ pedersen_hash and pedersen_hash_as_point are consistent")
        else:
            print("✗ pedersen_hash and pedersen_hash_as_point are inconsistent")
            print(f"  hash_result = {hex(hash_result)}")
            print(f"  point_result[0] = {hex(point_result[0])}")
            return False
        
        # Test that pedersen_hash_bytes returns the same as pedersen_hash
        bytes_result = pedersen.pedersen_hash_bytes(*elements)
        expected_bytes = hash_result.to_bytes(32, byteorder='big')
        
        if bytes_result == expected_bytes:
            print("✓ pedersen_hash_bytes and pedersen_hash are consistent")
            return True
        else:
            print("✗ pedersen_hash_bytes and pedersen_hash are inconsistent")
            print(f"  bytes_result = {bytes_result.hex()}")
            print(f"  expected_bytes = {expected_bytes.hex()}")
            return False
            
    except Exception as e:
        print(f"✗ Consistency test failed: {e}")
        import traceback
        traceback.print_exc()
        return False


def test_edge_cases():
    """Test edge cases."""
    print("\nTesting edge cases...")
    
    try:
        # Test with zero
        result_zero = pedersen.pedersen_hash(0)
        print(f"pedersen_hash(0) = {hex(result_zero)}")
        
        # Test with small number
        result_small = pedersen.pedersen_hash(1)
        print(f"pedersen_hash(1) = {hex(result_small)}")
        
        # Test with bytes input
        result_bytes_input = pedersen.pedersen_hash_bytes(b'\x01\x02\x03', 456)
        print(f"pedersen_hash_bytes(b'\\x01\\x02\\x03', 456) = {result_bytes_input.hex()}")
        
        print("✓ Edge case tests passed")
        return True
        
    except Exception as e:
        print(f"✗ Edge case test failed: {e}")
        import traceback
        traceback.print_exc()
        return False


def main():
    """Run all tests."""
    print("Running Pedersen hash implementation tests...\n")
    
    tests = [
        test_constants,
        test_basic_pedersen_hash,
        test_consistency,
        test_edge_cases,
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
