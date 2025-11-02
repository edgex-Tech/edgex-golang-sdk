package sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/edgex-Tech/edgex-golang-sdk/sdk/account"
	"github.com/edgex-Tech/edgex-golang-sdk/sdk/asset"
	"github.com/edgex-Tech/edgex-golang-sdk/sdk/funding"
	"github.com/edgex-Tech/edgex-golang-sdk/sdk/internal"
	"github.com/edgex-Tech/edgex-golang-sdk/sdk/metadata"
	"github.com/edgex-Tech/edgex-golang-sdk/sdk/order"
	"github.com/edgex-Tech/edgex-golang-sdk/sdk/quote"
	"github.com/edgex-Tech/edgex-golang-sdk/sdk/transfer"
	"github.com/shopspring/decimal"
	"golang.org/x/crypto/sha3"
)

// Client represents an EdgeX SDK client
type Client struct {
	*internal.Client
	metadataCache     *metadata.ResultMetaData
	metadataCacheTime time.Time
	metadataCacheTTL  *time.Duration
	Order             *order.Client
	Metadata          *metadata.Client
	Account           *account.Client
	Quote             *quote.Client
	Funding           *funding.Client
	Transfer          *transfer.Client
	Asset             *asset.Client
}

// ClientConfig holds the configuration for creating a new Client
type ClientConfig struct {
	BaseURL          string
	AccountID        int64
	StarkPriKey      string
	MetaDataCacheTTL *time.Duration
}

// NewClient creates a new EdgeX SDK client
func NewClient(cfg *ClientConfig) (*Client, error) {
	internalClient, err := internal.NewClient(&internal.ClientConfig{
		BaseURL:     cfg.BaseURL,
		AccountID:   cfg.AccountID,
		StarkPriKey: cfg.StarkPriKey,
	})
	if err != nil {
		return nil, err
	}

	return &Client{
		Client:           internalClient,
		metadataCacheTTL: cfg.MetaDataCacheTTL,
		Order:            order.NewClient(internalClient),
		Metadata:         metadata.NewClient(internalClient),
		Account:          account.NewClient(internalClient),
		Quote:            quote.NewClient(internalClient),
		Funding:          funding.NewClient(internalClient),
		Transfer:         transfer.NewClient(internalClient),
		Asset:            asset.NewClient(internalClient),
	}, nil
}

// requestInterceptor implements http.RoundTripper to intercept requests
type requestInterceptor struct {
	transport      http.RoundTripper
	internalClient *internal.Client
	baseURL        string
}

// RoundTrip implements http.RoundTripper
func (i *requestInterceptor) RoundTrip(req *http.Request) (*http.Response, error) {
	// Add timestamp header
	timestamp := time.Now().UnixMilli()
	req.Header.Set("X-edgeX-Api-Timestamp", fmt.Sprintf("%d", timestamp))

	// Generate signature content
	path := strings.TrimPrefix(req.URL.Path, i.baseURL)
	var signContent string
	if req.Body != nil {
		// Read and restore body since it will be consumed
		bodyBytes, err := io.ReadAll(req.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read request body: %w", err)
		}
		req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		// Convert body to sorted string format
		var bodyMap map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &bodyMap); err != nil {
			return nil, fmt.Errorf("failed to unmarshal body: %w", err)
		}

		bodyStr := internal.GetValue(bodyMap)
		signContent = fmt.Sprintf("%d%s%s%s", timestamp, req.Method, path, bodyStr)
	} else {
		// For requests without body, use query parameters if present
		if req.URL.RawQuery != "" {
			// Sort query parameters
			params := strings.Split(req.URL.RawQuery, "&")
			sort.Strings(params)
			signContent = fmt.Sprintf("%d%s%s%s", timestamp, req.Method, path, strings.Join(params, "&"))
		} else {
			signContent = fmt.Sprintf("%d%s%s", timestamp, req.Method, path)
		}
	}

	// Sign the content using stark private key
	hash := sha3.NewLegacyKeccak256()
	hash.Write([]byte(signContent))
	contentHash := hash.Sum(nil)

	sig, err := i.internalClient.Sign(contentHash)
	if err != nil {
		return nil, fmt.Errorf("failed to sign request: %w", err)
	}

	// Combine r and s into a single signature
	sigStr := fmt.Sprintf("%s%s", sig.R, sig.S)
	req.Header.Set("X-edgeX-Api-Signature", sigStr)

	// Forward the request to the underlying transport
	return i.transport.RoundTrip(req)
}

// GetMetaData gets the exchange metadata
func (c *Client) GetMetaData(ctx context.Context) (*metadata.ResultMetaData, error) {
	if c.metadataCacheTTL != nil {
		// Check if metadata is cached and not expired
		if c.metadataCache != nil && time.Since(c.metadataCacheTime) < *c.metadataCacheTTL {
			return c.metadataCache, nil
		}
		c.metadataCacheTime = time.Now()
	}
	metadataCache, err := c.Metadata.GetMetaData(ctx)
	if err != nil {
		return nil, err
	}
	c.metadataCache = metadataCache
	return c.metadataCache, nil
}

// GetServerTime gets the current server time
func (c *Client) GetServerTime(ctx context.Context) (*metadata.ResultGetServerTime, error) {
	return c.Metadata.GetServerTime(ctx)
}

// CreateOrder creates a new order with the given parameters
func (c *Client) CreateOrder(ctx context.Context, params *order.CreateOrderParams) (*order.ResultCreateOrder, error) {
	// Get metadata first
	metadataResp, err := c.GetMetaData(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get metadata: %w", err)
	}
	l2Price := params.Price
	if params.Type == order.OrderTypeMarket {
		price, err := c.getMarketOrderPrice(ctx, params.ContractId, params.Side)
		if err != nil {
			return nil, fmt.Errorf("failed to get market order price: %w", err)
		}
		l2Price = *price
	}

	// Convert l2Price string to decimal.Decimal
	l2PriceDecimal, err := decimal.NewFromString(l2Price)
	if err != nil {
		return nil, fmt.Errorf("invalid price format: %w", err)
	}

	return c.Order.CreateOrder(ctx, params, metadataResp.Data, l2PriceDecimal)
}

func (c *Client) CreateNormalWithdraw(ctx context.Context, params *asset.CreateNormalWithdrawParams) (*asset.ResultCreateNormalWithdraw, error) {
	// Get metadata first
	metadataResp, err := c.GetMetaData(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get metadata: %w", err)
	}

	return c.Asset.CreateNormalWithdraw(ctx, params, metadataResp.Data)
}

// func (c *Client) GetCrossWithdrawSignInfo(ctx context.Context, params *asset.CreateCrossWithdrawParams) (*asset.ResultCreateCrossWithdraw, error) {
// 	// Get metadata first
// }

// GetMaxOrderSize gets the maximum order size for a given contract and price
func (c *Client) GetMaxOrderSize(ctx context.Context, contractID string, price decimal.Decimal) (*order.ResultGetMaxCreateOrderSize, error) {
	return c.Order.GetMaxOrderSize(ctx, contractID, price)
}

// CancelOrder cancels a specific order
func (c *Client) CancelOrder(ctx context.Context, params *order.CancelOrderParams) (interface{}, error) {
	return c.Order.CancelOrder(ctx, params)
}

// GetActiveOrders gets active orders with pagination and filters
func (c *Client) GetActiveOrders(ctx context.Context, params *order.GetActiveOrderParams) (*order.ResultPageDataOrder, error) {
	return c.Order.GetActiveOrders(ctx, params)
}

// GetOrdersByID retrieves orders using exchange order IDs.
func (c *Client) GetOrdersByID(ctx context.Context, orderIDs []string) (*order.ResultListOrder, error) {
	return c.Order.GetOrdersByID(ctx, orderIDs)
}

// GetOrdersByClientOrderID retrieves orders using client-provided order IDs.
func (c *Client) GetOrdersByClientOrderID(ctx context.Context, clientOrderIDs []string) (*order.ResultListOrder, error) {
	return c.Order.GetOrdersByClientOrderID(ctx, clientOrderIDs)
}

// GetOrderFillTransactions gets order fill transactions with pagination and filters
func (c *Client) GetOrderFillTransactions(ctx context.Context, params *order.OrderFillTransactionParams) (*order.ResultPageDataOrderFillTransaction, error) {
	return c.Order.GetOrderFillTransactions(ctx, params)
}

// GetAccountAsset gets the account asset information
func (c *Client) GetAccountAsset(ctx context.Context) (*account.GetAccountAssetResponse, error) {
	return c.Account.GetAccountAsset(ctx)
}

// GetAccountPositions gets the account positions
func (c *Client) GetAccountPositions(ctx context.Context) (*account.ListPositionResponse, error) {
	return c.Account.GetAccountPositions(ctx)
}

// GetPositionTransactionPage gets the position transactions with pagination
func (c *Client) GetPositionTransactionPage(ctx context.Context, params account.GetPositionTransactionPageParams) (*account.PageDataPositionTransactionResponse, error) {
	return c.Account.GetPositionTransactionPage(ctx, params)
}

// GetCollateralTransactionPage gets the collateral transactions with pagination
func (c *Client) GetCollateralTransactionPage(ctx context.Context, params account.GetCollateralTransactionPageParams) (*account.PageDataCollateralTransactionResponse, error) {
	return c.Account.GetCollateralTransactionPage(ctx, params)
}

// GetPositionTermPage gets the position terms with pagination
func (c *Client) GetPositionTermPage(ctx context.Context, params account.GetPositionTermPageParams) (*account.PageDataPositionTermResponse, error) {
	return c.Account.GetPositionTermPage(ctx, params)
}

// GetAccountByID gets account information by ID
func (c *Client) GetAccountByID(ctx context.Context) (*account.AccountResponse, error) {
	return c.Account.GetAccountByID(ctx)
}

// GetAccountDeleverageLight gets account deleverage light information
func (c *Client) GetAccountDeleverageLight(ctx context.Context) (*account.GetAccountDeleverageLightResponse, error) {
	return c.Account.GetAccountDeleverageLight(ctx)
}

// GetAccountAssetSnapshotPage gets account asset snapshots with pagination
func (c *Client) GetAccountAssetSnapshotPage(ctx context.Context, params account.GetAccountAssetSnapshotPageParams) (*account.PageDataAccountAssetSnapshotResponse, error) {
	return c.Account.GetAccountAssetSnapshotPage(ctx, params)
}

// GetPositionTransactionByID gets position transactions by IDs
func (c *Client) GetPositionTransactionByID(ctx context.Context, transactionIDs []string) (*account.ListPositionTransactionResponse, error) {
	return c.Account.GetPositionTransactionByID(ctx, transactionIDs)
}

// GetCollateralTransactionByID gets collateral transactions by IDs
func (c *Client) GetCollateralTransactionByID(ctx context.Context, transactionIDs []string) (*account.ListCollateralTransactionResponse, error) {
	return c.Account.GetCollateralTransactionByID(ctx, transactionIDs)
}

// GetQuoteSummary gets the quote summary for a given contract
func (c *Client) GetQuoteSummary(ctx context.Context, contractID string) (*quote.ResultGetTickerSummaryModel, error) {
	return c.Quote.GetQuoteSummary(ctx, contractID)
}

// Get24HourQuotes gets the 24-hour quotes for given contracts
func (c *Client) Get24HourQuote(ctx context.Context, contractId string) (*quote.ResultListTicker, error) {
	return c.Quote.Get24HourQuote(ctx, contractId)
}

// GetKLine gets the K-line data for a contract
func (c *Client) GetKLine(ctx context.Context, params quote.GetKLineParams) (*quote.ResultPageDataKline, error) {
	return c.Quote.GetKLine(ctx, params)
}

// GetOrderBookDepth gets the order book depth for a contract
func (c *Client) GetOrderBookDepth(ctx context.Context, params quote.GetOrderBookDepthParams) (*quote.ResultListDepth, error) {
	return c.Quote.GetOrderBookDepth(ctx, params)
}

// GetMultiContractKLine gets the K-line data for multiple contracts
func (c *Client) GetMultiContractKLine(ctx context.Context, params quote.GetMultiContractKLineParams) (*quote.ResultListContractKline, error) {
	return c.Quote.GetMultiContractKLine(ctx, params)
}

// GetTransferOutById gets a transfer out record by ID
func (c *Client) GetTransferOutById(ctx context.Context, params transfer.GetTransferOutByIdParams) (*transfer.ResultListTransferOut, error) {
	return c.Transfer.GetTransferOutById(ctx, params)
}

// GetTransferInById gets a transfer in record by ID
func (c *Client) GetTransferInById(ctx context.Context, params transfer.GetTransferInByIdParams) (*transfer.ResultListTransferIn, error) {
	return c.Transfer.GetTransferInById(ctx, params)
}

// GetWithdrawAvailableAmount gets the available withdrawal amount
func (c *Client) GetWithdrawAvailableAmount(ctx context.Context, params transfer.GetWithdrawAvailableAmountParams) (*transfer.ResultGetTransferOutAvailableAmount, error) {
	return c.Transfer.GetWithdrawAvailableAmount(ctx, params)
}

// CreateTransferOut creates a new transfer out order
func (c *Client) CreateTransferOut(ctx context.Context, params *transfer.CreateTransferOutParams) (*transfer.ResultCreateTransferOut, error) {
	// Get metadata first
	metadataResp, err := c.GetMetaData(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get metadata: %w", err)
	}

	return c.Transfer.CreateTransferOut(ctx, params, metadataResp.Data)
}

// UpdateLeverageSetting updates the account leverage settings
func (c *Client) UpdateLeverageSetting(ctx context.Context, contractID string, leverage string) error {
	return c.Account.UpdateLeverageSetting(ctx, contractID, leverage)
}

// CreateMarketOrder creates a new market order with the given parameters
func (c *Client) getMarketOrderPrice(ctx context.Context, contractId, side string) (*string, error) {
	// Get metadata for contract info
	metadataResp, err := c.GetMetaData(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get metadata: %w", err)
	}

	// Find the contract
	var contract *metadata.Contract
	contractList := metadataResp.Data.ContractList
	for i, ct := range contractList {
		if ct.ContractId == contractId {
			contract = &contractList[i]
			break
		}
	}
	if contract == nil {
		return nil, fmt.Errorf("contract not found: %s", contractId)
	}

	// Calculate price based on side
	var price string
	if side == order.OrderSideBuy {
		// For buy orders: oracle_price * 10, rounded to price precision
		quote, err := c.Get24HourQuote(ctx, contractId)
		if err != nil {
			return nil, fmt.Errorf("failed to get 24-hour quotes: %w", err)
		}

		data := quote.Data
		if len(data) == 0 {
			return nil, fmt.Errorf("no quote data available for contract: %s", contractId)
		}

		// Extract oracle price from Ticker
		tickerData := data[0]
		oraclePriceStr := ""
		if tickerData.OraclePrice != nil {
			oraclePriceStr = *tickerData.OraclePrice
		}
		if oraclePriceStr == "" {
			return nil, fmt.Errorf("oracle price not found or invalid format")
		}

		oraclePrice, err := decimal.NewFromString(oraclePriceStr)
		if err != nil {
			return nil, fmt.Errorf("invalid oracle price: %s", oraclePriceStr)
		}
		multiplier := decimal.NewFromInt(10)
		tickSize, err := decimal.NewFromString(contract.TickSize)
		if err != nil {
			return nil, fmt.Errorf("invalid tick size: %s", contract.TickSize)
		}
		precision := int32(tickSize.Exponent())
		price = oraclePrice.Mul(multiplier).Round(precision).String()
	} else {
		// For sell orders: use tick size
		price = contract.TickSize
	}
	return &price, nil
}
