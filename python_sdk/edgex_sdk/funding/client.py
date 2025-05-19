from typing import Dict, Any, List, Optional

import requests

from ..internal.client import Client as InternalClient
from ..order.types import ResponseCode


class Client:
    """Client for funding-related API endpoints."""

    def __init__(self, internal_client: InternalClient, session: requests.Session):
        """
        Initialize the funding client.

        Args:
            internal_client: The internal client for common functionality
            session: The HTTP session for making requests
        """
        self.internal_client = internal_client
        self.session = session
        self.base_url = internal_client.base_url

    async def get_funding_transactions(
        self,
        size: str = "",
        offset_data: str = "",
        filter_coin_id_list: List[str] = None,
        filter_type_list: List[str] = None,
        filter_start_created_time_inclusive: int = 0,
        filter_end_created_time_exclusive: int = 0
    ) -> Dict[str, Any]:
        """
        Get funding transactions with pagination.

        Args:
            size: Size of the page
            offset_data: Offset data for pagination
            filter_coin_id_list: Filter by coin IDs
            filter_type_list: Filter by transaction types
            filter_start_created_time_inclusive: Filter start time (inclusive)
            filter_end_created_time_exclusive: Filter end time (exclusive)

        Returns:
            Dict[str, Any]: The funding transactions

        Raises:
            ValueError: If the request fails
        """
        url = f"{self.base_url}/api/v1/public/funding/getFundingRatePage"
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
        if filter_type_list:
            query_params["filterTypeList"] = ",".join(filter_type_list)

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

    async def get_funding_account(self) -> Dict[str, Any]:
        """
        Get funding account information.

        Returns:
            Dict[str, Any]: The funding account information

        Raises:
            ValueError: If the request fails
        """
        url = f"{self.base_url}/api/v1/private/account/getAccountAsset"
        params = {
            "accountId": str(self.internal_client.get_account_id())
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

    async def get_funding_transaction_by_id(self, transaction_ids: List[str]) -> Dict[str, Any]:
        """
        Get funding transactions by IDs.

        Args:
            transaction_ids: List of transaction IDs

        Returns:
            Dict[str, Any]: The funding transactions

        Raises:
            ValueError: If the request fails
        """
        url = f"{self.base_url}/api/v1/public/funding/getLatestFundingRate"
        query_params = {
            "accountId": str(self.internal_client.get_account_id()),
            "transactionIdList": ",".join(transaction_ids)
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
