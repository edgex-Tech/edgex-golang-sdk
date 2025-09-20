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

	openapi "github.com/edgex-Tech/edgex-golang-sdk/openapi"
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
	Order    *order.Client
	Metadata *metadata.Client
	Account  *account.Client
	Quote    *quote.Client
	Funding  *funding.Client
	Transfer *transfer.Client
	Asset    *asset.Client
}

// ClientConfig holds the configuration for creating a new Client
type ClientConfig struct {
	BaseURL     string
	AccountID   int64
	StarkPriKey string
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

	// Create OpenAPI client configuration
	openapiConfig := openapi.NewConfiguration()
	openapiConfig.Servers = []openapi.ServerConfiguration{
		{
			URL: cfg.BaseURL,
		},
	}

	// Create transport for request interception
	transport := http.DefaultTransport
	openapiConfig.HTTPClient = &http.Client{
		Transport: &requestInterceptor{
			transport:      transport,
			internalClient: internalClient,
			baseURL:        cfg.BaseURL,
		},
	}

	openapiClient := openapi.NewAPIClient(openapiConfig)

	return &Client{
		Client:   internalClient,
		Order:    order.NewClient(internalClient, openapiClient),
		Metadata: metadata.NewClient(internalClient, openapiClient),
		Account:  account.NewClient(internalClient, openapiClient),
		Quote:    quote.NewClient(internalClient, openapiClient),
		Funding:  funding.NewClient(internalClient, openapiClient),
		Transfer: transfer.NewClient(internalClient, openapiClient),
		Asset:    asset.NewClient(internalClient, openapiClient),
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
func (c *Client) GetMetaData(ctx context.Context) (*openapi.ResultMetaData, error) {
	return c.Metadata.GetMetaData(ctx)
}

// GetServerTime gets the current server time
func (c *Client) GetServerTime(ctx context.Context) (*openapi.ResultGetServerTime, error) {
	return c.Metadata.GetServerTime(ctx)
}

// CreateOrder creates a new order with the given parameters
func (c *Client) CreateOrder(ctx context.Context, params *order.CreateOrderParams) (*openapi.ResultCreateOrder, error) {
	// Get metadata first
	metadata, err := c.GetMetaData(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get metadata: %w", err)
	}

	return c.Order.CreateOrder(ctx, params, metadata.GetData())
}

// GetMaxOrderSize gets the maximum order size for a given contract and price
func (c *Client) GetMaxOrderSize(ctx context.Context, contractID string, price decimal.Decimal) (*openapi.ResultGetMaxCreateOrderSize, error) {
	priceFloat, _ := price.Float64()
	return c.Order.GetMaxOrderSize(ctx, contractID, priceFloat)
}

// CancelOrder cancels a specific order
func (c *Client) CancelOrder(ctx context.Context, params *order.CancelOrderParams) (interface{}, error) {
	return c.Order.CancelOrder(ctx, params)
}

// GetActiveOrders gets active orders with pagination and filters
func (c *Client) GetActiveOrders(ctx context.Context, params *order.GetActiveOrderParams) (*openapi.ResultPageDataOrder, error) {
	return c.Order.GetActiveOrders(ctx, params)
}

// GetOrdersByID retrieves orders using exchange order IDs.
func (c *Client) GetOrdersByID(ctx context.Context, orderIDs []string) (*openapi.ResultListOrder, error) {
	return c.Order.GetOrdersByID(ctx, orderIDs)
}

// GetOrdersByClientOrderID retrieves orders using client-provided order IDs.
func (c *Client) GetOrdersByClientOrderID(ctx context.Context, clientOrderIDs []string) (*openapi.ResultListOrder, error) {
	return c.Order.GetOrdersByClientOrderID(ctx, clientOrderIDs)
}

// GetOrderFillTransactions gets order fill transactions with pagination and filters
func (c *Client) GetOrderFillTransactions(ctx context.Context, params *order.OrderFillTransactionParams) (*openapi.ResultPageDataOrderFillTransaction, error) {
	return c.Order.GetOrderFillTransactions(ctx, params)
}

// GetAccountAsset gets the account asset information
func (c *Client) GetAccountAsset(ctx context.Context) (*openapi.ResultGetAccountAsset, error) {
	return c.Account.GetAccountAsset(ctx)
}

// GetAccountPositions gets the account positions
func (c *Client) GetAccountPositions(ctx context.Context) (*openapi.ResultListPosition, error) {
	return c.Account.GetAccountPositions(ctx)
}

// GetPositionTransactionPage gets the position transactions with pagination
func (c *Client) GetPositionTransactionPage(ctx context.Context, params account.GetPositionTransactionPageParams) (*openapi.ResultPageDataPositionTransaction, error) {
	return c.Account.GetPositionTransactionPage(ctx, params)
}

// GetCollateralTransactionPage gets the collateral transactions with pagination
func (c *Client) GetCollateralTransactionPage(ctx context.Context, params account.GetCollateralTransactionPageParams) (*openapi.ResultPageDataCollateralTransaction, error) {
	return c.Account.GetCollateralTransactionPage(ctx, params)
}

// GetPositionTermPage gets the position terms with pagination
func (c *Client) GetPositionTermPage(ctx context.Context, params account.GetPositionTermPageParams) (*openapi.ResultPageDataPositionTerm, error) {
	return c.Account.GetPositionTermPage(ctx, params)
}

// GetAccountByID gets account information by ID
func (c *Client) GetAccountByID(ctx context.Context) (*openapi.ResultAccount, error) {
	return c.Account.GetAccountByID(ctx)
}

// GetAccountDeleverageLight gets account deleverage light information
func (c *Client) GetAccountDeleverageLight(ctx context.Context) (*openapi.ResultGetAccountDeleverageLight, error) {
	return c.Account.GetAccountDeleverageLight(ctx)
}

// GetAccountAssetSnapshotPage gets account asset snapshots with pagination
func (c *Client) GetAccountAssetSnapshotPage(ctx context.Context, params account.GetAccountAssetSnapshotPageParams) (*openapi.ResultPageDataAccountAssetSnapshot, error) {
	return c.Account.GetAccountAssetSnapshotPage(ctx, params)
}

// GetPositionTransactionByID gets position transactions by IDs
func (c *Client) GetPositionTransactionByID(ctx context.Context, transactionIDs []string) (*openapi.ResultListPositionTransaction, error) {
	return c.Account.GetPositionTransactionByID(ctx, transactionIDs)
}

// GetCollateralTransactionByID gets collateral transactions by IDs
func (c *Client) GetCollateralTransactionByID(ctx context.Context, transactionIDs []string) (*openapi.ResultListCollateralTransaction, error) {
	return c.Account.GetCollateralTransactionByID(ctx, transactionIDs)
}

// GetQuoteSummary gets the quote summary for a given contract
func (c *Client) GetQuoteSummary(ctx context.Context, contractID string) (*openapi.ResultGetTickerSummaryModel, error) {
	return c.Quote.GetQuoteSummary(ctx, contractID)
}

// Get24HourQuotes gets the 24-hour quotes for given contracts
func (c *Client) Get24HourQuote(ctx context.Context, contractId string) (*openapi.ResultListTicker, error) {
	return c.Quote.Get24HourQuote(ctx, contractId)
}

// GetKLine gets the K-line data for a contract
func (c *Client) GetKLine(ctx context.Context, params quote.GetKLineParams) (*openapi.ResultPageDataKline, error) {
	return c.Quote.GetKLine(ctx, params)
}

// GetOrderBookDepth gets the order book depth for a contract
func (c *Client) GetOrderBookDepth(ctx context.Context, params quote.GetOrderBookDepthParams) (*openapi.ResultListDepth, error) {
	return c.Quote.GetOrderBookDepth(ctx, params)
}

// GetMultiContractKLine gets the K-line data for multiple contracts
func (c *Client) GetMultiContractKLine(ctx context.Context, params quote.GetMultiContractKLineParams) (*openapi.ResultListContractKline, error) {
	return c.Quote.GetMultiContractKLine(ctx, params)
}

// GetTransferOutById gets a transfer out record by ID
func (c *Client) GetTransferOutById(ctx context.Context, params transfer.GetTransferOutByIdParams) (*openapi.ResultListTransferOut, error) {
	return c.Transfer.GetTransferOutById(ctx, params)
}

// GetTransferInById gets a transfer in record by ID
func (c *Client) GetTransferInById(ctx context.Context, params transfer.GetTransferInByIdParams) (*openapi.ResultListTransferIn, error) {
	return c.Transfer.GetTransferInById(ctx, params)
}

// GetWithdrawAvailableAmount gets the available withdrawal amount
func (c *Client) GetWithdrawAvailableAmount(ctx context.Context, params transfer.GetWithdrawAvailableAmountParams) (*openapi.ResultGetTransferOutAvailableAmount, error) {
	return c.Transfer.GetWithdrawAvailableAmount(ctx, params)
}

// CreateTransferOut creates a new transfer out order
func (c *Client) CreateTransferOut(ctx context.Context, params transfer.CreateTransferOutParams) (*openapi.ResultCreateTransferOut, error) {
	// Get metadata first
	metadata, err := c.GetMetaData(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get metadata: %w", err)
	}

	return c.Transfer.CreateTransferOut(ctx, params, metadata.GetData())
}

// UpdateLeverageSetting updates the account leverage settings
func (c *Client) UpdateLeverageSetting(ctx context.Context, contractID string, leverage string) error {
	return c.Account.UpdateLeverageSetting(ctx, contractID, leverage)
}

// CreateLimitOrder creates a new limit order with the given parameters
func (c *Client) CreateLimitOrder(ctx context.Context, contractId, size, price, side string, clientOrderId *string) (*openapi.ResultCreateOrder, error) {
	params := &order.CreateOrderParams{
		ContractId:    contractId,
		Size:          size,
		Price:         price,
		Side:          side,
		Type:          order.OrderTypeLimit,
		ClientOrderId: clientOrderId,
	}
	return c.CreateOrder(ctx, params)
}

// CreateMarketOrder creates a new market order with the given parameters
func (c *Client) CreateMarketOrder(ctx context.Context, contractId, size, side string, clientOrderId *string) (*openapi.ResultCreateOrder, error) {
	// Get metadata for contract info
	metadata, err := c.GetMetaData(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get metadata: %w", err)
	}

	// Find the contract
	var contract *openapi.Contract
	contractList := metadata.Data.ContractList
	for _, c := range contractList {
		if *c.ContractId == contractId {
			contract = &c
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
		oraclePrice, err := decimal.NewFromString(quote.GetData()[0].GetOraclePrice())
		if err != nil {
			return nil, fmt.Errorf("invalid oracle price: %s", quote.GetData()[0].GetOraclePrice())
		}
		multiplier := decimal.NewFromInt(10)
		tickSize, err := decimal.NewFromString(*contract.TickSize)
		if err != nil {
			return nil, fmt.Errorf("invalid tick size: %s", *contract.TickSize)
		}
		precision := int32(tickSize.Exponent())
		price = oraclePrice.Mul(multiplier).Round(precision).String()
	} else {
		// For sell orders: use tick size
		price = *contract.TickSize
	}

	params := &order.CreateOrderParams{
		ContractId:    contractId,
		Size:          size,
		Price:         price,
		Side:          side,
		Type:          order.OrderTypeMarket,
		ClientOrderId: clientOrderId,
	}
	return c.CreateOrder(ctx, params)
}
