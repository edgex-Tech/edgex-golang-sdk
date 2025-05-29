import json
import time
from decimal import Decimal
from typing import Dict, Any, Optional, List

import requests

from ..internal.client import Client as InternalClient
from .types import (
    CreateOrderParams,
    CancelOrderParams,
    GetActiveOrderParams,
    OrderFillTransactionParams,
    TimeInForce,
    OrderType,
    ResponseCode
)


class Client:
    """Client for order-related API endpoints."""

    def __init__(self, internal_client: InternalClient, session: requests.Session):
        """
        Initialize the order client.

        Args:
            internal_client: The internal client for common functionality
            session: The HTTP session for making requests
        """
        self.internal_client = internal_client
        self.session = session
        self.base_url = internal_client.base_url

    async def create_order(self, params: CreateOrderParams, metadata: Dict[str, Any]) -> Dict[str, Any]:
        """
        Create a new order with the given parameters.

        Args:
            params: Order parameters
            metadata: Exchange metadata

        Returns:
            Dict[str, Any]: The created order

        Raises:
            ValueError: If required parameters are missing or invalid
        """
        # Set default TimeInForce based on order type if not specified
        if not params.time_in_force:
            if params.type == OrderType.MARKET:
                params.time_in_force = TimeInForce.IMMEDIATE_OR_CANCEL
            elif params.type == OrderType.LIMIT:
                params.time_in_force = TimeInForce.GOOD_TIL_CANCEL

        # Find the contract from metadata
        contract = None
        contract_list = metadata.get("contractList", [])
        for c in contract_list:
            if c.get("contractId") == params.contract_id:
                contract = c
                break

        if not contract:
            raise ValueError(f"contract not found: {params.contract_id}")

        # Get collateral coin from metadata
        global_data = metadata.get("global", {})
        collateral_coin = global_data.get("starkExCollateralCoin", {})

        # Parse decimal values
        try:
            size = Decimal(params.size)
            price = Decimal(params.price)
        except (ValueError, TypeError):
            raise ValueError("failed to parse size or price")

        # Convert hex resolution to decimal
        hex_resolution = contract.get("starkExResolution", "0x0")
        # Remove "0x" prefix if present
        hex_resolution = hex_resolution.replace("0x", "")
        # Parse hex string to int
        try:
            resolution_int = int(hex_resolution, 16)
            resolution = Decimal(resolution_int)
        except (ValueError, TypeError):
            raise ValueError("failed to parse hex resolution")

        client_order_id = params.client_order_id or self.internal_client.generate_uuid()

        # Calculate values
        value_dm = price * size
        amount_synthetic = int(size * resolution)
        amount_collateral = int(value_dm * Decimal("1000000"))  # Shift 6 decimal places

        # Calculate fee based on order type (maker/taker)
        try:
            fee_rate = Decimal(contract.get("defaultTakerFeeRate", "0"))
        except (ValueError, TypeError):
            raise ValueError("failed to parse fee rate")

        # Calculate fee amount in decimal with 6 decimal places
        amount_fee_dm = (value_dm * fee_rate).quantize(Decimal("0.000001"))
        amount_fee_str = str(amount_fee_dm)

        # Convert to the required integer format for the protocol
        amount_fee = int(amount_fee_dm * Decimal("1000000"))  # Shift 6 decimal places

        nonce = self.internal_client.calc_nonce(client_order_id)
        l2_expire_time = int(time.time() * 1000) + (14 * 24 * 60 * 60 * 1000)  # 14 days

        # Calculate signature using asset IDs from metadata
        expire_time_unix = l2_expire_time // (60 * 60 * 1000)
        sig_hash = self.internal_client.calc_limit_order_hash(
            contract.get("starkExSyntheticAssetId", ""),
            collateral_coin.get("starkExAssetId", ""),
            collateral_coin.get("starkExAssetId", ""),
            params.side == "BUY",
            amount_synthetic,
            amount_collateral,
            amount_fee,
            nonce,
            self.internal_client.get_account_id(),
            expire_time_unix
        )

        # Sign the order
        sig = self.internal_client.sign(sig_hash)

        # Convert signature to string
        sig_str = f"{sig.r}{sig.s}{sig.v}"

        # Create order request
        account_id = str(self.internal_client.get_account_id())
        nonce_str = str(nonce)
        l2_expire_time_str = str(l2_expire_time)
        expire_time_str = str(l2_expire_time - 864000000)  # 10 days earlier
        value_str = str(value_dm)

        price_str = params.price if params.type == OrderType.LIMIT else "0"

        # Prepare request data
        request_data = {
            "accountId": account_id,
            "contractId": params.contract_id,
            "price": price_str,
            "size": params.size,
            "type": str(params.type),
            "timeInForce": params.time_in_force,
            "side": params.side,
            "l2Signature": sig_str,
            "l2Nonce": nonce_str,
            "l2ExpireTime": l2_expire_time_str,
            "l2Value": value_str,
            "l2Size": params.size,
            "l2LimitFee": amount_fee_str,
            "clientOrderId": client_order_id,
            "expireTime": expire_time_str,
            "reduceOnly": params.reduce_only
        }

        # Execute request
        url = f"{self.base_url}/api/v1/private/order/createOrder"
        response = self.session.post(url, json=request_data)

        if response.status_code != 200:
            raise ValueError(f"request failed with status code: {response.status_code}")

        resp_data = response.json()

        if resp_data.get("code") != ResponseCode.SUCCESS:
            error_param = resp_data.get("errorParam")
            if error_param:
                raise ValueError(f"request failed with error params: {error_param}")
            raise ValueError(f"request failed with code: {resp_data.get('code')}")

        return resp_data

    async def cancel_order(self, params: CancelOrderParams) -> Dict[str, Any]:
        """
        Cancel a specific order.

        Args:
            params: Cancel order parameters

        Returns:
            Dict[str, Any]: The cancellation result

        Raises:
            ValueError: If required parameters are missing or invalid
        """
        account_id = str(self.internal_client.get_account_id())

        if params.order_id:
            url = f"{self.base_url}/api/v1/private/order/cancelOrderById"
            request_data = {
                "accountId": account_id,
                "orderIdList": [params.order_id]
            }
        elif params.client_id:
            url = f"{self.base_url}/api/v1/private/order/cancelOrderByClientOrderId"
            request_data = {
                "accountId": account_id,
                "clientOrderIdList": [params.client_id]
            }
        elif params.contract_id:
            url = f"{self.base_url}/api/v1/private/order/cancelAllOrder"
            request_data = {
                "accountId": account_id,
                "filterContractIdList": [params.contract_id]
            }
        else:
            raise ValueError("must provide either order_id, client_id, or contract_id")

        # Execute request
        response = self.session.post(url, json=request_data)

        if response.status_code != 200:
            raise ValueError(f"request failed with status code: {response.status_code}")

        resp_data = response.json()

        if resp_data.get("code") != ResponseCode.SUCCESS:
            error_param = resp_data.get("errorParam")
            if error_param:
                raise ValueError(f"request failed with error params: {error_param}")
            raise ValueError(f"request failed with code: {resp_data.get('code')}")

        return resp_data

    async def get_active_orders(self, params: GetActiveOrderParams) -> Dict[str, Any]:
        """
        Get active orders with pagination and filters.

        Args:
            params: Active order query parameters

        Returns:
            Dict[str, Any]: The active orders

        Raises:
            ValueError: If the request fails
        """
        # Build query parameters
        query_params = {
            "accountId": str(self.internal_client.get_account_id())
        }

        # Add pagination parameters
        if params.size:
            query_params["size"] = params.size
        if params.offset_data:
            query_params["offsetData"] = params.offset_data

        # Add filter parameters
        if params.filter_coin_id_list:
            query_params["filterCoinIdList"] = ",".join(params.filter_coin_id_list)
        if params.filter_contract_id_list:
            query_params["filterContractIdList"] = ",".join(params.filter_contract_id_list)
        if params.filter_type_list:
            query_params["filterTypeList"] = ",".join(params.filter_type_list)
        if params.filter_status_list:
            query_params["filterStatusList"] = ",".join(params.filter_status_list)

        # Add boolean filters
        if params.filter_is_liquidate is not None:
            query_params["filterIsLiquidateList"] = str(params.filter_is_liquidate).lower()
        if params.filter_is_deleverage is not None:
            query_params["filterIsDeleverageList"] = str(params.filter_is_deleverage).lower()
        if params.filter_is_position_tpsl is not None:
            query_params["filterIsPositionTpslList"] = str(params.filter_is_position_tpsl).lower()

        # Add time filters
        if params.filter_start_created_time_inclusive > 0:
            query_params["filterStartCreatedTimeInclusive"] = str(params.filter_start_created_time_inclusive)
        if params.filter_end_created_time_exclusive > 0:
            query_params["filterEndCreatedTimeExclusive"] = str(params.filter_end_created_time_exclusive)

        # Execute request
        url = f"{self.base_url}/api/v1/private/order/getActiveOrderPage"
        response = self.session.get(url, params=query_params)

        if response.status_code != 200:
            raise ValueError(f"request failed with status code: {response.status_code}")

        resp_data = response.json()

        if resp_data.get("code") != ResponseCode.SUCCESS:
            error_param = resp_data.get("errorParam")
            if error_param:
                raise ValueError(f"request failed with error params: {error_param}")
            raise ValueError(f"request failed with code: {resp_data.get('code')}")

        return resp_data

    async def get_order_fill_transactions(self, params: OrderFillTransactionParams) -> Dict[str, Any]:
        """
        Get order fill transactions with pagination and filters.

        Args:
            params: Order fill transaction query parameters

        Returns:
            Dict[str, Any]: The order fill transactions

        Raises:
            ValueError: If the request fails
        """
        # Build query parameters
        query_params = {
            "accountId": str(self.internal_client.get_account_id())
        }

        # Add pagination parameters
        if params.size:
            query_params["size"] = params.size
        if params.offset_data:
            query_params["offsetData"] = params.offset_data

        # Add filter parameters
        if params.filter_coin_id_list:
            query_params["filterCoinIdList"] = ",".join(params.filter_coin_id_list)
        if params.filter_contract_id_list:
            query_params["filterContractIdList"] = ",".join(params.filter_contract_id_list)
        if params.filter_order_id_list:
            query_params["filterOrderIdList"] = ",".join(params.filter_order_id_list)

        # Add boolean filters
        if params.filter_is_liquidate is not None:
            query_params["filterIsLiquidateList"] = str(params.filter_is_liquidate).lower()
        if params.filter_is_deleverage is not None:
            query_params["filterIsDeleverageList"] = str(params.filter_is_deleverage).lower()
        if params.filter_is_position_tpsl is not None:
            query_params["filterIsPositionTpslList"] = str(params.filter_is_position_tpsl).lower()

        # Add time filters
        if params.filter_start_created_time_inclusive > 0:
            query_params["filterStartCreatedTimeInclusive"] = str(params.filter_start_created_time_inclusive)
        if params.filter_end_created_time_exclusive > 0:
            query_params["filterEndCreatedTimeExclusive"] = str(params.filter_end_created_time_exclusive)

        # Execute request
        url = f"{self.base_url}/api/v1/private/order/getHistoryOrderFillTransactionPage"
        response = self.session.get(url, params=query_params)

        if response.status_code != 200:
            raise ValueError(f"request failed with status code: {response.status_code}")

        resp_data = response.json()

        if resp_data.get("code") != ResponseCode.SUCCESS:
            error_param = resp_data.get("errorParam")
            if error_param:
                raise ValueError(f"request failed with error params: {error_param}")
            raise ValueError(f"request failed with code: {resp_data.get('code')}")

        return resp_data

    async def get_max_order_size(self, contract_id: str, price: float) -> Dict[str, Any]:
        """
        Get the maximum order size for a given contract and price.

        Args:
            contract_id: The contract ID
            price: The price

        Returns:
            Dict[str, Any]: The maximum order size information

        Raises:
            ValueError: If the request fails
        """
        # Build request body (API expects POST with JSON body)
        data = {
            "accountId": str(self.internal_client.get_account_id()),
            "contractId": contract_id,
            "price": str(price)
        }

        # Execute request
        url = f"{self.base_url}/api/v1/private/order/getMaxCreateOrderSize"
        response = self.session.post(url, json=data)

        if response.status_code != 200:
            raise ValueError(f"request failed with status code: {response.status_code}")

        resp_data = response.json()

        if resp_data.get("code") != ResponseCode.SUCCESS:
            error_param = resp_data.get("errorParam")
            if error_param:
                raise ValueError(f"request failed with error params: {error_param}")
            raise ValueError(f"request failed with code: {resp_data.get('code')}")

        return resp_data
