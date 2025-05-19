"""
Basic usage example for the EdgeX Python SDK.
"""

import asyncio
import os
from decimal import Decimal

from edgex_sdk import (
    Client,
    OrderSide,
    GetKLineParams,
    GetOrderBookDepthParams,
    WebSocketManager
)


async def main():
    # Load configuration from environment variables
    base_url = os.getenv("EDGEX_BASE_URL", "https://testnet.edgex.exchange")
    account_id = int(os.getenv("EDGEX_ACCOUNT_ID", "12345"))
    stark_private_key = os.getenv("EDGEX_STARK_PRIVATE_KEY", "your-stark-private-key")
    
    # Create a new client
    client = Client(
        base_url=base_url,
        account_id=account_id,
        stark_private_key=stark_private_key
    )
    
    # Get server time
    server_time = await client.get_server_time()
    print(f"Server Time: {server_time}")
    
    # Get metadata
    metadata = await client.get_metadata()
    print(f"Metadata: {metadata.get('data', {}).get('global', {})}")
    
    # Get account assets
    assets = await client.get_account_asset()
    print(f"Account Assets: {assets}")
    
    # Get account positions
    positions = await client.get_account_positions()
    print(f"Account Positions: {positions}")
    
    # Get K-line data
    kline_params = GetKLineParams(
        contract_id="BTC-USDT",
        interval="1m",
        size="10"
    )
    klines = await client.quote.get_k_line(kline_params)
    print(f"K-lines: {klines}")
    
    # Get order book depth
    depth_params = GetOrderBookDepthParams(
        contract_id="BTC-USDT",
        limit=10
    )
    depth = await client.quote.get_order_book_depth(depth_params)
    print(f"Order Book Depth: {depth}")
    
    # Create a limit order (commented out to avoid actual order creation)
    """
    limit_order = await client.create_limit_order(
        contract_id="BTC-USDT",
        size="0.001",
        price="30000",
        side=OrderSide.BUY
    )
    print(f"Limit Order: {limit_order}")
    """
    
    # WebSocket example
    ws_url = os.getenv("EDGEX_WS_URL", "wss://testnet.edgex.exchange")
    ws_manager = WebSocketManager(
        base_url=ws_url,
        account_id=account_id,
        stark_pri_key=stark_private_key
    )
    
    # Define message handlers
    def ticker_handler(message):
        print(f"Ticker Update: {message}")
    
    def kline_handler(message):
        print(f"K-line Update: {message}")
    
    # Connect to public WebSocket
    ws_manager.connect_public()
    
    # Subscribe to ticker and kline updates
    ws_manager.subscribe_ticker("BTC-USDT", ticker_handler)
    ws_manager.subscribe_kline("BTC-USDT", "1m", kline_handler)
    
    # Wait for some updates
    await asyncio.sleep(30)
    
    # Disconnect
    ws_manager.disconnect_all()


if __name__ == "__main__":
    asyncio.run(main())
