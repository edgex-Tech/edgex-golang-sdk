import json
import time
from typing import Dict, Any, Optional, List, Union
from decimal import Decimal

import requests
from Crypto.Hash import keccak

from .internal.client import Client as InternalClient
from .internal.signing_adapter import SigningAdapter
from .internal.starkex_signing_adapter import StarkExSigningAdapter
from .account.client import Client as AccountClient
from .asset.client import Client as AssetClient
from .funding.client import Client as FundingClient
from .metadata.client import Client as MetadataClient
from .order.client import Client as OrderClient
from .quote.client import Client as QuoteClient
from .transfer.client import Client as TransferClient
from .order.types import CreateOrderParams, CancelOrderParams, GetActiveOrderParams, OrderFillTransactionParams


class RequestInterceptor:
    """Intercepts requests to add authentication headers."""

    def __init__(self, internal_client: InternalClient, base_url: str):
        self.internal_client = internal_client
        self.base_url = base_url

    def __call__(self, request):
        """Add authentication headers to the request."""
        # Add timestamp header
        timestamp = int(time.time() * 1000)
        request.headers["X-edgeX-Api-Timestamp"] = str(timestamp)

        # Generate signature content
        full_path = request.url.replace(self.base_url, "")

        # Split path and query
        if '?' in full_path:
            path, query = full_path.split('?', 1)
        else:
            path = full_path
            query = ""

        if request.body:
            # Read body
            body_bytes = request.body
            if isinstance(body_bytes, bytes):
                body_str = body_bytes.decode('utf-8')
            else:
                body_str = body_bytes

            # Convert body to sorted string format
            body_map = json.loads(body_str)
            body_str = self.internal_client.get_value(body_map)
            sign_content = f"{timestamp}{request.method}{path}{body_str}"
        else:
            # For requests without body, use query parameters if present
            if query:
                # Sort query parameters as strings (matching Go SDK exactly)
                params = sorted(query.split('&'))
                sign_content = f"{timestamp}{request.method}{path}{'&'.join(params)}"
            else:
                sign_content = f"{timestamp}{request.method}{path}"

        # Sign the content
        keccak_hash = keccak.new(digest_bits=256)
        keccak_hash.update(sign_content.encode())
        content_hash = keccak_hash.digest()

        sig = self.internal_client.sign(content_hash)

        # Combine r and s into a single signature
        sig_str = f"{sig.r}{sig.s}"
        request.headers["X-edgeX-Api-Signature"] = sig_str



        return request


class Client:
    """Main EdgeX SDK client."""

    def __init__(self, base_url: str, account_id: int, stark_private_key: str, signing_adapter: Optional[SigningAdapter] = None):
        """
        Initialize the EdgeX SDK client.

        Args:
            base_url: Base URL for API endpoints
            account_id: Account ID for authentication
            stark_private_key: Stark private key for signing
            signing_adapter: Optional signing adapter (defaults to StarkExSigningAdapter)
        """
        # Use StarkExSigningAdapter as default if none provided
        if signing_adapter is None:
            signing_adapter = StarkExSigningAdapter()

        self.internal_client = InternalClient(
            base_url=base_url,
            account_id=account_id,
            stark_pri_key=stark_private_key,
            signing_adapter=signing_adapter
        )

        # Create session with interceptor
        session = requests.Session()

        # Create a custom adapter that adds authentication
        class AuthHTTPAdapter(requests.adapters.HTTPAdapter):
            def __init__(self, internal_client, base_url):
                super().__init__()
                self.internal_client = internal_client
                self.base_url = base_url

            def send(self, request, **kwargs):
                # Add authentication headers
                interceptor = RequestInterceptor(self.internal_client, self.base_url)
                request = interceptor(request)
                return super().send(request, **kwargs)

        session.mount(base_url, AuthHTTPAdapter(self.internal_client, base_url))

        # Initialize API clients
        self.order = OrderClient(self.internal_client, session)
        self.metadata = MetadataClient(self.internal_client, session)
        self.account = AccountClient(self.internal_client, session)
        self.quote = QuoteClient(self.internal_client, session)
        self.funding = FundingClient(self.internal_client, session)
        self.transfer = TransferClient(self.internal_client, session)
        self.asset = AssetClient(self.internal_client, session)

    async def get_metadata(self) -> Dict[str, Any]:
        """Get the exchange metadata."""
        return await self.metadata.get_metadata()

    async def get_server_time(self) -> Dict[str, Any]:
        """Get the current server time."""
        return await self.metadata.get_server_time()

    async def create_order(self, params: CreateOrderParams) -> Dict[str, Any]:
        """
        Create a new order with the given parameters.

        Args:
            params: Order parameters

        Returns:
            Dict[str, Any]: The created order
        """
        # Get metadata first
        metadata = await self.get_metadata()
        if not metadata:
            raise ValueError("failed to get metadata")

        return await self.order.create_order(params, metadata.get("data", {}))

    async def get_max_order_size(self, contract_id: str, price: Decimal) -> Dict[str, Any]:
        """
        Get the maximum order size for a given contract and price.

        Args:
            contract_id: The contract ID
            price: The price

        Returns:
            Dict[str, Any]: The maximum order size information
        """
        return await self.order.get_max_order_size(contract_id, float(price))

    async def cancel_order(self, params: CancelOrderParams) -> Dict[str, Any]:
        """
        Cancel a specific order.

        Args:
            params: Cancel order parameters

        Returns:
            Dict[str, Any]: The cancellation result
        """
        return await self.order.cancel_order(params)

    async def get_active_orders(self, params: GetActiveOrderParams) -> Dict[str, Any]:
        """
        Get active orders with pagination and filters.

        Args:
            params: Active order query parameters

        Returns:
            Dict[str, Any]: The active orders
        """
        return await self.order.get_active_orders(params)

    async def get_order_fill_transactions(self, params: OrderFillTransactionParams) -> Dict[str, Any]:
        """
        Get order fill transactions with pagination and filters.

        Args:
            params: Order fill transaction query parameters

        Returns:
            Dict[str, Any]: The order fill transactions
        """
        return await self.order.get_order_fill_transactions(params)

    async def get_account_asset(self) -> Dict[str, Any]:
        """Get the account asset information."""
        return await self.account.get_account_asset()

    async def get_account_positions(self) -> Dict[str, Any]:
        """Get the account positions."""
        return await self.account.get_account_positions()

    async def create_limit_order(
        self,
        contract_id: str,
        size: str,
        price: str,
        side: str,
        client_order_id: Optional[str] = None
    ) -> Dict[str, Any]:
        """
        Create a new limit order with the given parameters.

        Args:
            contract_id: The contract ID
            size: The order size
            price: The order price
            side: The order side (BUY or SELL)
            client_order_id: Optional client order ID

        Returns:
            Dict[str, Any]: The created order
        """
        from .order.types import OrderType

        params = CreateOrderParams(
            contract_id=contract_id,
            size=size,
            price=price,
            side=side,
            type=OrderType.LIMIT,
            client_order_id=client_order_id
        )

        return await self.create_order(params)

    async def create_market_order(
        self,
        contract_id: str,
        size: str,
        side: str,
        client_order_id: Optional[str] = None
    ) -> Dict[str, Any]:
        """
        Create a new market order with the given parameters.

        Args:
            contract_id: The contract ID
            size: The order size
            side: The order side (BUY or SELL)
            client_order_id: Optional client order ID

        Returns:
            Dict[str, Any]: The created order
        """
        # Get metadata for contract info
        metadata = await self.get_metadata()
        if not metadata:
            raise ValueError("failed to get metadata")

        # Find the contract
        contract = None
        contract_list = metadata.get("data", {}).get("contractList", [])
        for c in contract_list:
            if c.get("contractId") == contract_id:
                contract = c
                break

        if not contract:
            raise ValueError(f"contract not found: {contract_id}")

        # Calculate price based on side
        from .order.types import OrderSide, OrderType

        if side == OrderSide.BUY:
            # For buy orders: oracle_price * 10, rounded to price precision
            quote = await self.get_24_hour_quote(contract_id)
            if not quote:
                raise ValueError("failed to get 24-hour quotes")

            oracle_price = Decimal(quote.get("data", [])[0].get("oraclePrice", "0"))
            multiplier = Decimal("10")
            tick_size = Decimal(contract.get("tickSize", "0"))
            precision = abs(tick_size.as_tuple().exponent)
            price = str(round(oracle_price * multiplier, precision))
        else:
            # For sell orders: use tick size
            price = contract.get("tickSize", "0")

        params = CreateOrderParams(
            contract_id=contract_id,
            size=size,
            price=price,
            side=side,
            type=OrderType.MARKET,
            client_order_id=client_order_id
        )

        return await self.create_order(params)

    async def get_24_hour_quote(self, contract_id: str) -> Dict[str, Any]:
        """
        Get the 24-hour quotes for a given contract.

        Args:
            contract_id: The contract ID

        Returns:
            Dict[str, Any]: The 24-hour quotes
        """
        return await self.quote.get_24_hour_quote(contract_id)
