#!/usr/bin/env python3
"""
Test script for the signing adapter with Pedersen hash integration.

This script tests the signing adapter to ensure it correctly uses the new Pedersen hash implementation.
"""

import sys
import os
import importlib.util

# Add the current directory to the path
current_dir = os.path.dirname(__file__)
sys.path.insert(0, current_dir)

# Add the crypto directory to the path so modules can import each other
crypto_dir = os.path.join(current_dir, 'edgex_sdk', 'crypto')
sys.path.insert(0, crypto_dir)

# Add the internal directory to the path
internal_dir = os.path.join(current_dir, 'edgex_sdk', 'internal')
sys.path.insert(0, internal_dir)

def load_module_from_path(module_name, file_path):
    """Load a module from a file path."""
    spec = importlib.util.spec_from_file_location(module_name, file_path)
    module = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(module)
    return module

# Load the required modules
constants_path = os.path.join(crypto_dir, 'constants.py')
constants = load_module_from_path('constants', constants_path)

pedersen_path = os.path.join(crypto_dir, 'pedersen_hash.py')
pedersen = load_module_from_path('pedersen_hash', pedersen_path)

# Load the signing adapter base class
signing_adapter_path = os.path.join(internal_dir, 'signing_adapter.py')
signing_adapter_base = load_module_from_path('signing_adapter', signing_adapter_path)

# Load the StarkEx signing adapter
starkex_adapter_path = os.path.join(internal_dir, 'starkex_signing_adapter.py')
starkex_adapter = load_module_from_path('starkex_signing_adapter', starkex_adapter_path)


def test_signing_adapter_pedersen_hash():
    """Test the signing adapter's Pedersen hash method."""
    print("Testing signing adapter Pedersen hash integration...")
    
    try:
        # Create an instance of the signing adapter
        adapter = starkex_adapter.StarkExSigningAdapter()
        
        # Test the pedersen_hash method
        elements = [123, 456, 789]
        result = adapter.pedersen_hash(elements)
        print(f"adapter.pedersen_hash({elements}) = {result.hex()}")
        
        # Compare with direct call to our implementation
        direct_result = pedersen.pedersen_hash_bytes(*elements)
        print(f"direct pedersen_hash_bytes({elements}) = {direct_result.hex()}")
        
        if result == direct_result:
            print("✓ Signing adapter Pedersen hash matches direct implementation")
            return True
        else:
            print("✗ Results don't match between adapter and direct call")
            print(f"  Adapter result:  {result.hex()}")
            print(f"  Direct result:   {direct_result.hex()}")
            return False
            
    except Exception as e:
        print(f"✗ Signing adapter Pedersen hash test failed: {e}")
        import traceback
        traceback.print_exc()
        return False


def test_signing_adapter_basic_functionality():
    """Test basic signing adapter functionality."""
    print("\nTesting signing adapter basic functionality...")
    
    try:
        # Create an instance of the signing adapter
        adapter = starkex_adapter.StarkExSigningAdapter()
        
        # Test key generation
        private_key = "0123456789abcdef" * 8  # 64 hex chars = 32 bytes
        public_key = adapter.get_public_key(private_key)
        print(f"Generated public key: {public_key}")
        
        # Test signing
        message_hash = b"Hello, StarkEx!"
        signature = adapter.sign(message_hash, private_key)
        print(f"Signature: r={signature[0]}, s={signature[1]}")
        
        # Test verification
        is_valid = adapter.verify(message_hash, signature, public_key)
        print(f"Signature verification: {is_valid}")
        
        if is_valid:
            print("✓ Basic signing adapter functionality works")
            return True
        else:
            print("✗ Signature verification failed")
            return False
            
    except Exception as e:
        print(f"✗ Basic signing adapter test failed: {e}")
        import traceback
        traceback.print_exc()
        return False


def test_pedersen_hash_with_different_inputs():
    """Test Pedersen hash with various input types."""
    print("\nTesting Pedersen hash with different inputs...")
    
    try:
        adapter = starkex_adapter.StarkExSigningAdapter()
        
        # Test with single element
        result1 = adapter.pedersen_hash([42])
        print(f"pedersen_hash([42]) = {result1.hex()}")
        
        # Test with multiple elements
        result2 = adapter.pedersen_hash([1, 2, 3, 4, 5])
        print(f"pedersen_hash([1, 2, 3, 4, 5]) = {result2.hex()}")
        
        # Test with large numbers
        large_num = 2**200
        result3 = adapter.pedersen_hash([large_num])
        print(f"pedersen_hash([{large_num}]) = {result3.hex()}")
        
        # Test with zero
        result4 = adapter.pedersen_hash([0])
        print(f"pedersen_hash([0]) = {result4.hex()}")
        
        print("✓ Pedersen hash works with different input types")
        return True
        
    except Exception as e:
        print(f"✗ Pedersen hash different inputs test failed: {e}")
        import traceback
        traceback.print_exc()
        return False


def test_error_handling():
    """Test error handling in the signing adapter."""
    print("\nTesting error handling...")
    
    try:
        adapter = starkex_adapter.StarkExSigningAdapter()
        
        # Test with invalid private key
        try:
            adapter.get_public_key("invalid_hex")
            print("✗ Should have failed with invalid private key")
            return False
        except ValueError:
            print("✓ Correctly caught error for invalid private key")
        
        # Test with number too large for Pedersen hash
        try:
            # This should work with our implementation since we handle large numbers
            very_large = constants.FIELD_PRIME + 1
            adapter.pedersen_hash([very_large])
            print("✗ Should have failed with number too large")
            return False
        except ValueError:
            print("✓ Correctly caught error for number too large")
        
        print("✓ Error handling tests passed")
        return True
        
    except Exception as e:
        print(f"✗ Error handling test failed: {e}")
        import traceback
        traceback.print_exc()
        return False


def main():
    """Run all tests."""
    print("Running signing adapter integration tests...\n")
    
    tests = [
        test_signing_adapter_pedersen_hash,
        test_signing_adapter_basic_functionality,
        test_pedersen_hash_with_different_inputs,
        test_error_handling,
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
