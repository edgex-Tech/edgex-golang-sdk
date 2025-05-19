from typing import Dict, Any

import requests

from ..internal.client import Client as InternalClient
from ..order.types import ResponseCode


class Client:
    """Client for metadata-related API endpoints."""
    
    def __init__(self, internal_client: InternalClient, session: requests.Session):
        """
        Initialize the metadata client.
        
        Args:
            internal_client: The internal client for common functionality
            session: The HTTP session for making requests
        """
        self.internal_client = internal_client
        self.session = session
        self.base_url = internal_client.base_url
    
    async def get_metadata(self) -> Dict[str, Any]:
        """
        Get the exchange metadata.
        
        Returns:
            Dict[str, Any]: The exchange metadata
            
        Raises:
            ValueError: If the request fails
        """
        url = f"{self.base_url}/api/v1/metadata"
        response = self.session.get(url)
        
        if response.status_code != 200:
            raise ValueError(f"request failed with status code: {response.status_code}")
        
        resp_data = response.json()
        
        if resp_data.get("code") != ResponseCode.SUCCESS:
            error_param = resp_data.get("errorParam")
            if error_param:
                raise ValueError(f"request failed with error params: {error_param}")
            raise ValueError(f"request failed with code: {resp_data.get('code')}")
        
        return resp_data
    
    async def get_server_time(self) -> Dict[str, Any]:
        """
        Get the current server time.
        
        Returns:
            Dict[str, Any]: The server time information
            
        Raises:
            ValueError: If the request fails
        """
        url = f"{self.base_url}/api/v1/metadata/time"
        response = self.session.get(url)
        
        if response.status_code != 200:
            raise ValueError(f"request failed with status code: {response.status_code}")
        
        resp_data = response.json()
        
        if resp_data.get("code") != ResponseCode.SUCCESS:
            error_param = resp_data.get("errorParam")
            if error_param:
                raise ValueError(f"request failed with error params: {error_param}")
            raise ValueError(f"request failed with code: {resp_data.get('code')}")
        
        return resp_data
