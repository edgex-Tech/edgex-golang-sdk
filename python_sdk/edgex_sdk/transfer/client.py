from typing import Dict, Any, List, Optional

import requests

from ..internal.client import Client as InternalClient
from ..order.types import ResponseCode


class GetTransferOutByIdParams:
    """Parameters for getting transfer out records by ID."""

    def __init__(self, transfer_id_list: List[str]):
        self.transfer_id_list = transfer_id_list


class GetTransferInByIdParams:
    """Parameters for getting transfer in records by ID."""

    def __init__(self, transfer_id_list: List[str]):
        self.transfer_id_list = transfer_id_list


class GetWithdrawAvailableAmountParams:
    """Parameters for getting available withdrawal amount."""

    def __init__(self, coin_id: str):
        self.coin_id = coin_id


class CreateTransferOutParams:
    """Parameters for creating a transfer out order."""

    def __init__(
        self,
        coin_id: str,
        amount: str,
        address: str,
        network: str,
        memo: str = "",
        client_order_id: str = None
    ):
        self.coin_id = coin_id
        self.amount = amount
        self.address = address
        self.network = network
        self.memo = memo
        self.client_order_id = client_order_id


class Client:
    """Client for transfer-related API endpoints."""

    def __init__(self, internal_client: InternalClient, session: requests.Session):
        """
        Initialize the transfer client.

        Args:
            internal_client: The internal client for common functionality
            session: The HTTP session for making requests
        """
        self.internal_client = internal_client
        self.session = session
        self.base_url = internal_client.base_url

    async def get_transfer_out_by_id(self, params: GetTransferOutByIdParams) -> Dict[str, Any]:
        """
        Get transfer out records by ID.

        Args:
            params: Transfer out query parameters

        Returns:
            Dict[str, Any]: The transfer out records

        Raises:
            ValueError: If the request fails
        """
        url = f"{self.base_url}/api/v1/private/transfer/getTransferOutById"
        query_params = {
            "accountId": str(self.internal_client.get_account_id()),
            "transferIdList": ",".join(params.transfer_id_list)
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

    async def get_transfer_in_by_id(self, params: GetTransferInByIdParams) -> Dict[str, Any]:
        """
        Get transfer in records by ID.

        Args:
            params: Transfer in query parameters

        Returns:
            Dict[str, Any]: The transfer in records

        Raises:
            ValueError: If the request fails
        """
        url = f"{self.base_url}/api/v1/private/transfer/getTransferInById"
        query_params = {
            "accountId": str(self.internal_client.get_account_id()),
            "transferIdList": ",".join(params.transfer_id_list)
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

    async def get_withdraw_available_amount(self, params: GetWithdrawAvailableAmountParams) -> Dict[str, Any]:
        """
        Get the available withdrawal amount.

        Args:
            params: Withdrawal available amount query parameters

        Returns:
            Dict[str, Any]: The available withdrawal amount

        Raises:
            ValueError: If the request fails
        """
        url = f"{self.base_url}/api/v1/private/transfer/getTransferOutAvailableAmount"
        query_params = {
            "accountId": str(self.internal_client.get_account_id()),
            "coinId": params.coin_id
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

    async def create_transfer_out(self, params: CreateTransferOutParams, metadata: Dict[str, Any]) -> Dict[str, Any]:
        """
        Create a new transfer out order.

        Args:
            params: Transfer out parameters
            metadata: Exchange metadata

        Returns:
            Dict[str, Any]: The created transfer out order

        Raises:
            ValueError: If the request fails
        """
        url = f"{self.base_url}/api/v1/private/transfer/createTransferOut"

        client_order_id = params.client_order_id or self.internal_client.generate_uuid()

        data = {
            "accountId": str(self.internal_client.get_account_id()),
            "coinId": params.coin_id,
            "amount": params.amount,
            "address": params.address,
            "network": params.network,
            "clientOrderId": client_order_id
        }

        if params.memo:
            data["memo"] = params.memo

        # TODO: Implement signature calculation for transfer out
        # This would require additional logic similar to the order creation

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

    async def get_transfer_out_page(
        self,
        size: str = "",
        offset_data: str = "",
        filter_coin_id_list: List[str] = None,
        filter_status_list: List[str] = None,
        filter_start_created_time_inclusive: int = 0,
        filter_end_created_time_exclusive: int = 0
    ) -> Dict[str, Any]:
        """
        Get transfer out records with pagination.

        Args:
            size: Size of the page
            offset_data: Offset data for pagination
            filter_coin_id_list: Filter by coin IDs
            filter_status_list: Filter by status
            filter_start_created_time_inclusive: Filter start time (inclusive)
            filter_end_created_time_exclusive: Filter end time (exclusive)

        Returns:
            Dict[str, Any]: The transfer out records

        Raises:
            ValueError: If the request fails
        """
        url = f"{self.base_url}/api/v1/private/transfer/getActiveTransferOut"
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

    async def get_transfer_in_page(
        self,
        size: str = "",
        offset_data: str = "",
        filter_coin_id_list: List[str] = None,
        filter_status_list: List[str] = None,
        filter_start_created_time_inclusive: int = 0,
        filter_end_created_time_exclusive: int = 0
    ) -> Dict[str, Any]:
        """
        Get transfer in records with pagination.

        Args:
            size: Size of the page
            offset_data: Offset data for pagination
            filter_coin_id_list: Filter by coin IDs
            filter_status_list: Filter by status
            filter_start_created_time_inclusive: Filter start time (inclusive)
            filter_end_created_time_exclusive: Filter end time (exclusive)

        Returns:
            Dict[str, Any]: The transfer in records

        Raises:
            ValueError: If the request fails
        """
        url = f"{self.base_url}/api/v1/private/transfer/getActiveTransferIn"
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
