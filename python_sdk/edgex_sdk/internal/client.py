import binascii
import hashlib
import time
import uuid
from typing import Dict, Any, Optional, Tuple

import requests
import sha3

from .mock_signing_adapter import MockSigningAdapter
from .signing_adapter import SigningAdapter


class L2Signature:
    """Represents a signature for L2 operations."""

    def __init__(self, r: str, s: str, v: str = ""):
        self.r = r
        self.s = s
        self.v = v


class Client:
    """Base client with common functionality."""

    def __init__(self, base_url: str, account_id: int, stark_pri_key: str, signing_adapter: Optional[SigningAdapter] = None):
        """
        Initialize the internal client.

        Args:
            base_url: Base URL for API endpoints
            account_id: Account ID for authentication
            stark_pri_key: Stark private key for signing
            signing_adapter: Optional signing adapter to use for cryptographic operations
        """
        self.http_client = requests.Session()
        self.http_client.headers.update({
            "Content-Type": "application/json",
            "Accept": "application/json"
        })
        self.base_url = base_url
        self.account_id = account_id
        self.stark_pri_key = stark_pri_key

        # Use the provided signing adapter or create a mock one
        self.signing_adapter = signing_adapter or MockSigningAdapter()

    def get_account_id(self) -> int:
        """Get the account ID."""
        return self.account_id

    def get_stark_pri_key(self) -> str:
        """Get the stark private key."""
        return self.stark_pri_key

    def sign(self, message_hash: bytes) -> L2Signature:
        """
        Sign a message hash using the client's Stark private key.

        Args:
            message_hash: The hash of the message to sign

        Returns:
            L2Signature: The signature components

        Raises:
            ValueError: If the stark private key is not set or invalid
        """
        private_key = self.get_stark_pri_key()
        if not private_key:
            raise ValueError("stark private key not set")

        # Sign the message using the signing adapter
        try:
            r, s = self.signing_adapter.sign(message_hash, private_key)
            return L2Signature(r=r, s=s, v="")
        except Exception as e:
            raise ValueError(f"failed to sign message: {str(e)}")

    def generate_uuid(self) -> str:
        """Generate a UUID for client order IDs."""
        return str(uuid.uuid4())

    def calc_nonce(self, client_order_id: str) -> int:
        """
        Calculate a nonce from a client order ID.

        Args:
            client_order_id: The client order ID

        Returns:
            int: The calculated nonce
        """
        # Use the first 8 characters of the hash of the client order ID
        keccak = sha3.keccak_256()
        keccak.update(client_order_id.encode())
        hash_hex = keccak.hexdigest()
        return int(hash_hex[:8], 16)

    def calc_limit_order_hash(
        self,
        synthetic_asset_id: str,
        collateral_asset_id: str,
        fee_asset_id: str,
        is_buy: bool,
        amount_synthetic: int,
        amount_collateral: int,
        amount_fee: int,
        nonce: int,
        account_id: int,
        expire_time: int
    ) -> bytes:
        """
        Calculate the hash for a limit order.

        Args:
            synthetic_asset_id: The synthetic asset ID
            collateral_asset_id: The collateral asset ID
            fee_asset_id: The fee asset ID
            is_buy: Whether the order is a buy order
            amount_synthetic: The synthetic amount
            amount_collateral: The collateral amount
            amount_fee: The fee amount
            nonce: The nonce
            account_id: The account ID
            expire_time: The expiration time

        Returns:
            bytes: The calculated hash
        """
        # TODO: Implement the actual hash calculation
        # This is a placeholder for the actual hash calculation

        # For now, we'll create a mock hash
        keccak = sha3.keccak_256()
        keccak.update(f"{synthetic_asset_id}{collateral_asset_id}{fee_asset_id}{is_buy}{amount_synthetic}{amount_collateral}{amount_fee}{nonce}{account_id}{expire_time}".encode())
        return keccak.digest()

    def get_value(self, data: Dict[str, Any]) -> str:
        """
        Convert a dictionary to a sorted string format for signing.

        Args:
            data: The dictionary to convert

        Returns:
            str: The sorted string representation
        """
        # TODO: Implement the actual conversion
        # This is a placeholder for the actual conversion

        # For now, we'll return a simple string representation
        return str(data)
