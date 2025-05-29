from typing import Dict, Any, List, Optional

import requests

from ..internal.client import Client as InternalClient
from ..order.types import ResponseCode


class GetAssetOrdersParams:
    """Parameters for getting asset orders."""

    def __init__(self, size: str = "10", offset_data: str = "", filter_coin_id_list: List[str] = None,
                 filter_start_created_time_inclusive: int = 0, filter_end_created_time_exclusive: int = 0):
        self.size = size
        self.offset_data = offset_data
        self.filter_coin_id_list = filter_coin_id_list or []
        self.filter_start_created_time_inclusive = filter_start_created_time_inclusive
        self.filter_end_created_time_exclusive = filter_end_created_time_exclusive


class CreateWithdrawalParams:
    """Parameters for creating a withdrawal."""

    def __init__(self, coin_id: str, amount: str, address: str, tag: str = ""):
        self.coin_id = coin_id
        self.amount = amount
        self.address = address
        self.tag = tag


class GetWithdrawalRecordsParams:
    """Parameters for getting withdrawal records."""

    def __init__(self, size: str = "10", offset_data: str = "", filter_coin_id_list: List[str] = None,
                 filter_status_list: List[str] = None, filter_start_created_time_inclusive: int = 0,
                 filter_end_created_time_exclusive: int = 0):
        self.size = size
        self.offset_data = offset_data
        self.filter_coin_id_list = filter_coin_id_list or []
        self.filter_status_list = filter_status_list or []
        self.filter_start_created_time_inclusive = filter_start_created_time_inclusive
        self.filter_end_created_time_exclusive = filter_end_created_time_exclusive


class Client:
    """Client for asset-related API endpoints."""

    def __init__(self, internal_client: InternalClient, session: requests.Session):
        """
        Initialize the asset client.

        Args:
            internal_client: The internal client for common functionality
            session: The HTTP session for making requests
        """
        self.internal_client = internal_client
        self.session = session
        self.base_url = internal_client.base_url

    async def get_account_asset(self) -> Dict[str, Any]:
        """
        Get the account asset information.
        Note: This method delegates to the account client since it's an account endpoint.

        Returns:
            Dict[str, Any]: The account asset information

        Raises:
            ValueError: If the request fails
        """
        # This is actually an account endpoint, not an asset endpoint
        # We should delegate to the account client
        raise NotImplementedError("This method should be called from the account client: client.account.get_account_asset()")

    async def get_asset_orders(
        self,
        params: GetAssetOrdersParams
    ) -> Dict[str, Any]:
        """
        Get asset orders with pagination.

        Args:
            params: Parameters for the request

        Returns:
            Dict[str, Any]: The asset orders

        Raises:
            ValueError: If the request fails
        """
        url = f"{self.base_url}/api/v1/private/assets/getAllOrdersPage"
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

        # Add time filters
        if params.filter_start_created_time_inclusive > 0:
            query_params["filterStartCreatedTimeInclusive"] = str(params.filter_start_created_time_inclusive)
        if params.filter_end_created_time_exclusive > 0:
            query_params["filterEndCreatedTimeExclusive"] = str(params.filter_end_created_time_exclusive)

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

    async def get_coin_rates(self, chain_id: str = "1", coin: str = "0xdac17f958d2ee523a2206206994597c13d831ec7") -> Dict[str, Any]:
        """
        Get coin rates.

        Args:
            chain_id: Chain ID (default: "1" for Ethereum mainnet)
            coin: Coin contract address (default: USDT)

        Returns:
            Dict[str, Any]: The coin rates

        Raises:
            ValueError: If the request fails
        """
        url = f"{self.base_url}/api/v1/private/assets/getCoinRate"
        params = {
            "chainId": chain_id,
            "coin": coin
        }

        response = self.session.get(url, params=params)

        if response.status_code != 200:
            raise ValueError(f"request failed with status code: {response.status_code}")

        resp_data = response.json()

        if resp_data.get("code") != ResponseCode.SUCCESS:
            error_param = resp_data.get("errorParam")
            if error_param:
                raise ValueError(f"request failed with error params: {error_param}")
            raise ValueError(f"request failed with code: {resp_data.get('code')}")

        return resp_data

    async def create_withdrawal(
        self,
        coin_id: str,
        amount: str,
        address: str,
        network: str,
        memo: str = "",
        client_order_id: str = None
    ) -> Dict[str, Any]:
        """
        Create a withdrawal request.

        Args:
            coin_id: The coin ID
            amount: The withdrawal amount
            address: The withdrawal address
            network: The network
            memo: Optional memo
            client_order_id: Optional client order ID

        Returns:
            Dict[str, Any]: The withdrawal result

        Raises:
            ValueError: If the request fails
        """
        url = f"{self.base_url}/api/v1/private/assets/createNormalWithdraw"

        data = {
            "accountId": str(self.internal_client.get_account_id()),
            "coinId": coin_id,
            "amount": amount,
            "address": address,
            "network": network
        }

        if memo:
            data["memo"] = memo

        if client_order_id:
            data["clientOrderId"] = client_order_id
        else:
            data["clientOrderId"] = self.internal_client.generate_uuid()

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

    async def get_withdrawal_records(
        self,
        size: str = "",
        offset_data: str = "",
        filter_coin_id_list: List[str] = None,
        filter_status_list: List[str] = None,
        filter_start_created_time_inclusive: int = 0,
        filter_end_created_time_exclusive: int = 0
    ) -> Dict[str, Any]:
        """
        Get withdrawal records with pagination.

        Args:
            size: Size of the page
            offset_data: Offset data for pagination
            filter_coin_id_list: Filter by coin IDs
            filter_status_list: Filter by status
            filter_start_created_time_inclusive: Filter start time (inclusive)
            filter_end_created_time_exclusive: Filter end time (exclusive)

        Returns:
            Dict[str, Any]: The withdrawal records

        Raises:
            ValueError: If the request fails
        """
        url = f"{self.base_url}/api/v1/private/assets/getNormalWithdrawById"
        query_params = {
            "accountId": str(self.internal_client.get_account_id())
        }

        # Add pagination parameters
        if size:
            query_params["size"] = size
        if offset_data:
            query_params["offsetData"] = offset_data

        # Add filter parameters
        if filter_coin_id_list:
            query_params["filterCoinIdList"] = ",".join(filter_coin_id_list)
        if filter_status_list:
            query_params["filterStatusList"] = ",".join(filter_status_list)

        # Add time filters
        if filter_start_created_time_inclusive > 0:
            query_params["filterStartCreatedTimeInclusive"] = str(filter_start_created_time_inclusive)
        if filter_end_created_time_exclusive > 0:
            query_params["filterEndCreatedTimeExclusive"] = str(filter_end_created_time_exclusive)

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

    async def get_withdrawable_amount(self, address: str) -> Dict[str, Any]:
        """
        Get the withdrawable amount for a coin.

        Args:
            address: The coin contract address

        Returns:
            Dict[str, Any]: The withdrawable amount information

        Raises:
            ValueError: If the request fails
        """
        url = f"{self.base_url}/api/v1/private/assets/getNormalWithdrawableAmount"
        query_params = {
            "address": address
        }

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
