#!/usr/bin/env python3
"""
Debug script to test signature generation for API authentication.
"""

import json
import time
import sha3
from edgex_sdk.internal.client import Client as InternalClient
from edgex_sdk.internal.starkex_signing_adapter import StarkExSigningAdapter


def debug_signature_generation():
    """Debug the signature generation process."""
    print("Debugging API signature generation...\n")
    
    # Test credentials
    account_id = 542069824235241782
    stark_private_key = "05fa53d4fba9f28d9d4206025bac96ab07080a748881ec710407c21fd4acafed"
    base_url = "https://testnet.edgex.exchange"
    
    # Create internal client with StarkEx signing adapter
    signing_adapter = StarkExSigningAdapter()
    client = InternalClient(
        base_url=base_url,
        account_id=account_id,
        stark_pri_key=stark_private_key,
        signing_adapter=signing_adapter
    )
    
    # Simulate the request that's failing
    timestamp = int(time.time() * 1000)
    method = "POST"
    path = "/api/v1/private/order/createOrder"
    
    # Sample request body (simplified)
    body_data = {
        "accountId": str(account_id),
        "contractId": "10000004",
        "price": "999999",
        "size": "0.001",
        "type": "LIMIT",
        "timeInForce": "GOOD_TIL_CANCEL",
        "side": "BUY",
        "l2Signature": "test_signature",
        "l2Nonce": "12345",
        "l2ExpireTime": str(timestamp + 14 * 24 * 60 * 60 * 1000),
        "l2Value": "999.999",
        "l2Size": "0.001",
        "l2LimitFee": "1.0",
        "clientOrderId": "test-order-123",
        "expireTime": str(timestamp + 10 * 24 * 60 * 60 * 1000),
        "reduceOnly": False
    }
    
    # Convert body to sorted string format (like the client does)
    body_str = client.get_value(body_data)
    
    # Create sign content
    sign_content = f"{timestamp}{method}{path}{body_str}"
    
    print(f"Timestamp: {timestamp}")
    print(f"Method: {method}")
    print(f"Path: {path}")
    print(f"Body string: {body_str}")
    print(f"Sign content: {sign_content}")
    print()
    
    # Hash the content
    keccak = sha3.keccak_256()
    keccak.update(sign_content.encode())
    content_hash = keccak.digest()
    
    print(f"Content hash (hex): {content_hash.hex()}")
    print()
    
    # Sign the hash
    signature = client.sign(content_hash)
    sig_str = f"{signature.r}{signature.s}"
    
    print(f"Signature r: {signature.r}")
    print(f"Signature s: {signature.s}")
    print(f"Combined signature: {sig_str}")
    print(f"Signature length: {len(sig_str)}")
    print()
    
    # Test with a simpler body
    simple_body = {"accountId": str(account_id)}
    simple_body_str = client.get_value(simple_body)
    simple_sign_content = f"{timestamp}{method}{path}{simple_body_str}"
    
    print("=== Testing with simple body ===")
    print(f"Simple body string: {simple_body_str}")
    print(f"Simple sign content: {simple_sign_content}")
    
    keccak2 = sha3.keccak_256()
    keccak2.update(simple_sign_content.encode())
    simple_hash = keccak2.digest()
    
    simple_signature = client.sign(simple_hash)
    simple_sig_str = f"{simple_signature.r}{simple_signature.s}"
    
    print(f"Simple signature: {simple_sig_str}")
    print()
    
    # Test the public key derivation
    public_key = signing_adapter.get_public_key(stark_private_key)
    print(f"Public key: {public_key}")
    print()
    
    # Test signature verification
    try:
        is_valid = signing_adapter.verify(content_hash, (signature.r, signature.s), public_key)
        print(f"Signature verification: {is_valid}")
    except Exception as e:
        print(f"Signature verification failed: {e}")
    
    print("\n=== Debugging complete ===")


if __name__ == "__main__":
    debug_signature_generation()
