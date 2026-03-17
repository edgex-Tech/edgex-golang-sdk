package sdk

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
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
)

const defaultMetaDataCacheTTL = time.Hour
const defaultHeaderKey = "edgeX"

// Client represents an EdgeX SDK client
type Client struct {
	// HTTP client
	httpClient *http.Client
	baseURL    string
	accountID  int64

	// EIP-712 keys
	signerPriKey string
	walletPriKey string
	signerAddr   string
	walletAddr   string

	// HMAC credentials
	apiKey        string
	apiPassphrase string
	apiSecret     string
	authHeaderKey string

	// SDK config
	sdkConfig *internal.SDKConfig

	// Metadata cache
	metadataCache     *metadata.ResultMetaData
	metadataCacheTime time.Time
	metadataCacheTTL  *time.Duration

	// Sub-clients
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
	BaseURL          string
	AccountID        int64
	SignerPriKey     string // EIP-712 private key for order signing
	WalletPriKey     string // EIP-712 private key for withdrawal signing
	SignerAddr       string // Signer address (auto-derived if not provided)
	WalletAddr       string // Wallet address (auto-derived if not provided)
	APIKey           string // HMAC API Key
	APIPassphrase    string // HMAC API Passphrase
	APISecret        string // HMAC API Secret
	AuthHeaderKey    string // Custom auth header key (optional)
	MetaDataCacheTTL *time.Duration
}

// NewClient creates a new EdgeX SDK client
func NewClient(cfg *ClientConfig) (*Client, error) {
	authHeaderKey := strings.TrimSpace(cfg.AuthHeaderKey)
	if authHeaderKey == "" {
		authHeaderKey = defaultHeaderKey
	}

	// Load SDK configuration from file
	sdkConfig, err := internal.LoadConfig()
	if err != nil {
		fmt.Printf("LoadConfig err:%v", err)
	}

	metadataCacheTTL := cfg.MetaDataCacheTTL
	if metadataCacheTTL == nil {
		defaultTTL := defaultMetaDataCacheTTL
		metadataCacheTTL = &defaultTTL
	}

	client := &Client{
		httpClient:       &http.Client{Timeout: 30 * time.Second},
		baseURL:          normalizeBaseURL(cfg.BaseURL),
		accountID:        cfg.AccountID,
		signerPriKey:     strings.TrimPrefix(cfg.SignerPriKey, "0x"),
		walletPriKey:     strings.TrimPrefix(cfg.WalletPriKey, "0x"),
		signerAddr:       strings.TrimSpace(cfg.SignerAddr),
		walletAddr:       strings.TrimSpace(cfg.WalletAddr),
		apiKey:           cfg.APIKey,
		apiPassphrase:    cfg.APIPassphrase,
		apiSecret:        cfg.APISecret,
		authHeaderKey:    authHeaderKey,
		sdkConfig:        sdkConfig,
		metadataCacheTTL: metadataCacheTTL,
	}

	// Initialize sub-clients
	client.Order = order.NewClient(client)
	client.Metadata = metadata.NewClient(client)
	client.Account = account.NewClient(client)
	client.Quote = quote.NewClient(client)
	client.Funding = funding.NewClient(client)
	client.Transfer = transfer.NewClient(client)
	client.Asset = asset.NewClient(client)

	return client, nil
}

func normalizeBaseURL(baseURL string) string {
	baseURL = strings.TrimSpace(baseURL)
	baseURL = strings.TrimRight(baseURL, "/")
	lowerBaseURL := strings.ToLower(baseURL)
	if strings.HasSuffix(lowerBaseURL, "/api") && len(baseURL) > 4 {
		return baseURL[:len(baseURL)-4]
	}
	return baseURL
}

// GetAccountID returns the account ID
func (c *Client) GetAccountID() int64 {
	return c.accountID
}

// GetSignerPriKey returns the EIP-712 signer private key
func (c *Client) GetSignerPriKey() string {
	return c.signerPriKey
}

// GetWalletPriKey returns the EIP-712 wallet private key
func (c *Client) GetWalletPriKey() string {
	return c.walletPriKey
}

// GetSignerAddr returns the configured signer address
func (c *Client) GetSignerAddr() string {
	return c.signerAddr
}

// GetWalletAddr returns the configured wallet address
func (c *Client) GetWalletAddr() string {
	return c.walletAddr
}

// GetBaseURL returns the base URL
func (c *Client) GetBaseURL() string {
	return c.baseURL
}

// GetAuthHeaderKey returns the configured auth header key prefix
func (c *Client) GetAuthHeaderKey() string {
	return c.authHeaderKey
}

// ResolveSignerAddress returns the configured signer address or derives it from private key
func (c *Client) ResolveSignerAddress() (string, error) {
	if c.signerAddr != "" {
		return c.signerAddr, nil
	}
	if c.signerPriKey == "" {
		return "", fmt.Errorf("signer address/private key not set")
	}
	return internal.DeriveAddressFromPrivateKey(c.signerPriKey)
}

// ResolveWalletAddress returns the configured wallet address or derives it from private key
func (c *Client) ResolveWalletAddress() (string, error) {
	if c.walletAddr != "" {
		return c.walletAddr, nil
	}
	if c.walletPriKey == "" {
		return "", fmt.Errorf("wallet address/private key not set")
	}
	return internal.DeriveAddressFromPrivateKey(c.walletPriKey)
}

// SignTypedDataWithSignerKey signs EIP-712 typed data using the configured signer key
func (c *Client) SignTypedDataWithSignerKey(typedData internal.TypedData) (string, error) {
	if c.signerPriKey == "" {
		return "", fmt.Errorf("signer private key not set")
	}
	return internal.SignTypedDataWithPrivateKey(c.signerPriKey, typedData)
}

// SignTypedDataWithWalletKey signs EIP-712 typed data using the configured wallet key
func (c *Client) SignTypedDataWithWalletKey(typedData internal.TypedData) (string, error) {
	if c.walletPriKey == "" {
		return "", fmt.Errorf("wallet private key not set")
	}
	return internal.SignTypedDataWithPrivateKey(c.walletPriKey, typedData)
}

// HttpRequest makes an authenticated HTTP request
func (c *Client) HttpRequest(urlStr string, method string, data map[string]interface{}, params map[string]string) (*http.Response, error) {
	method = strings.ToUpper(method)

	// Generate timestamp
	timestamp := time.Now().UnixMilli()
	timestampStr := fmt.Sprintf("%d", timestamp)

	// Parse URL to extract path
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}
	signPath := parsedURL.Path
	if parsedURL.RawQuery != "" {
		signPath = signPath + "?" + parsedURL.RawQuery
	}

	// Create request
	var req *http.Request
	if method == http.MethodGet {
		// For GET requests, add params to URL
		if len(params) > 0 {
			q := parsedURL.Query()
			for k, v := range params {
				q.Add(k, v)
			}
			parsedURL.RawQuery = q.Encode()
			urlStr = parsedURL.String()
		}
		req, err = http.NewRequest(method, urlStr, nil)
	} else {
		// For POST/PUT requests, send data as JSON body
		var body io.Reader
		if len(data) > 0 {
			bodyBytes, err := json.Marshal(data)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal request body: %w", err)
			}
			body = bytes.NewReader(bodyBytes)
		}
		req, err = http.NewRequest(method, urlStr, body)
		if body != nil {
			req.Header.Set("Content-Type", "application/json")
		}
	}

	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Add HMAC authentication headers
	if c.shouldSignWithHMAC(signPath) {
		if c.apiKey == "" || c.apiPassphrase == "" || c.apiSecret == "" {
			return nil, fmt.Errorf("hmac credentials are required for private endpoints")
		}

		signParams := mergeQueryParams(parsedURL.Query(), params)
		requestBody := c.buildRequestBodyForSignature(method, data, signParams)
		requestURI := c.buildHMACRequestURI(parsedURL.Path)
		signature := c.buildHMACSignature(timestampStr, method, requestURI, requestBody)

		req.Header.Set(fmt.Sprintf("X-%s-Api-Key", c.authHeaderKey), c.apiKey)
		req.Header.Set(fmt.Sprintf("X-%s-Passphrase", c.authHeaderKey), c.apiPassphrase)
		req.Header.Set(fmt.Sprintf("X-%s-Signature", c.authHeaderKey), signature)
		req.Header.Set(fmt.Sprintf("X-%s-Timestamp", c.authHeaderKey), timestampStr)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-edgeX-Channel", "Golang-SDK")

	// Add custom headers from configuration file
	if c.sdkConfig != nil {
		for key, value := range c.sdkConfig.CustomHeaders {
			req.Header.Set(key, value)
		}
	}

	// Log request if enabled
	if c.sdkConfig != nil && c.sdkConfig.Logging.EnableRequestLog {
		c.logRequest(req, data, params)
	}

	// Execute request
	reqStartTime := time.Now()
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}

	// Log response if enabled
	if c.sdkConfig != nil && c.sdkConfig.Logging.EnableResponseLog {
		c.logResponse(resp, reqStartTime)
	}

	// Log error if status code >= 400
	if resp.StatusCode >= 400 {
		c.logAPIError(resp, reqStartTime, data, params)
	}

	return resp, nil
}

func mergeQueryParams(query url.Values, params map[string]string) map[string]string {
	merged := make(map[string]string, len(params)+len(query))
	for k, values := range query {
		if len(values) == 0 {
			continue
		}
		merged[k] = values[0]
	}
	for k, v := range params {
		merged[k] = v
	}
	return merged
}

func (c *Client) shouldSignWithHMAC(path string) bool {
	// Only sign private endpoints
	// Public endpoints like /public/meta, /public/quote, /public/funding should NOT be signed
	return strings.Contains(path, "private") ||
		strings.Contains(path, "/opt/") ||
		strings.Contains(path, "invite") ||
		strings.Contains(path, "points") ||
		strings.Contains(path, "/public/codeInfo") ||
		strings.Contains(path, "/public/info/") || // Note: specific path with trailing slash
		strings.HasPrefix(path, "/ranking") ||
		strings.Contains(path, "nft") ||
		strings.Contains(path, "vault") ||
		strings.HasPrefix(path, "/activity") ||
		strings.Contains(path, "/download")
}

func (c *Client) buildHMACRequestURI(path string) string {
	if strings.Contains(path, "private") && !strings.HasPrefix(path, "/api") {
		return "/api" + path
	}
	return path
}

func (c *Client) buildRequestBodyForSignature(method string, data map[string]interface{}, params map[string]string) string {
	if method == http.MethodGet {
		return c.buildSortedQueryString(params)
	}
	if len(data) == 0 {
		return ""
	}
	return c.getValue(data)
}

func (c *Client) buildSortedQueryString(params map[string]string) string {
	if len(params) == 0 {
		return ""
	}
	keys := make([]string, 0, len(params))
	for k := range params {
		if params[k] == "" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)
	pairs := make([]string, 0, len(keys))
	for _, k := range keys {
		pairs = append(pairs, fmt.Sprintf("%s=%s", k, params[k]))
	}
	return strings.Join(pairs, "&")
}

func (c *Client) buildHMACSignature(timestamp string, method string, requestURI string, requestBody string) string {
	message := timestamp + method + requestURI + requestBody
	base64Key := base64.StdEncoding.EncodeToString([]byte(c.apiSecret))

	mac := hmac.New(sha256.New, []byte(base64Key))
	mac.Write([]byte(message))
	return hex.EncodeToString(mac.Sum(nil))
}

func (c *Client) getValue(data interface{}) string {
	switch v := data.(type) {
	case nil:
		return ""
	case string:
		return v
	case bool:
		return strings.ToLower(fmt.Sprintf("%v", v))
	case int, int32, int64, float32, float64:
		return fmt.Sprintf("%v", v)
	case []interface{}:
		if len(v) == 0 {
			return ""
		}
		var values []string
		for _, item := range v {
			values = append(values, c.getValue(item))
		}
		return strings.Join(values, "&")
	case []string:
		if len(v) == 0 {
			return ""
		}
		var values []string
		for _, item := range v {
			values = append(values, c.getValue(item))
		}
		return strings.Join(values, "&")
	case map[string]interface{}:
		sortedMap := make(map[string]string)
		for key, val := range v {
			sortedMap[key] = c.getValue(val)
		}

		keys := make([]string, 0, len(sortedMap))
		for k := range sortedMap {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		var pairs []string
		for _, k := range keys {
			pairs = append(pairs, fmt.Sprintf("%s=%s", k, sortedMap[k]))
		}
		return strings.Join(pairs, "&")
	default:
		return fmt.Sprintf("%v", v)
	}
}

func (c *Client) logRequest(req *http.Request, data map[string]interface{}, params map[string]string) {
	logData := map[string]interface{}{
		"time":   time.Now().Format(time.RFC3339),
		"method": req.Method,
		"url":    req.URL.String(),
		"headers": map[string]string{
			"Content-Type": req.Header.Get("Content-Type"),
			"Accept":       req.Header.Get("Accept"),
		},
	}

	if len(data) > 0 {
		logData["body"] = data
	}
	if len(params) > 0 {
		logData["params"] = params
	}

	if jsonBytes, err := json.Marshal(logData); err == nil {
		fmt.Printf("SDK Request: %s\n", string(jsonBytes))
	}
}

func (c *Client) logResponse(resp *http.Response, reqStartTime time.Time) {
	duration := time.Since(reqStartTime)

	// Read body for logging
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	// Restore body for further reading
	resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// Extract traceId from response body (EdgeX API standard)
	traceId := ""
	var respBody map[string]interface{}
	if json.Unmarshal(bodyBytes, &respBody) == nil {
		if traceStr, ok := respBody["traceId"].(string); ok {
			traceId = traceStr
		}
	}

	logData := map[string]interface{}{
		"time":     time.Now().Format(time.RFC3339),
		"status":   resp.StatusCode,
		"duration": duration.String(),
		"traceId":  traceId,
	}

	// Try to parse as JSON for better readability
	var bodyJSON interface{}
	if json.Unmarshal(bodyBytes, &bodyJSON) == nil {
		logData["body"] = bodyJSON
	} else {
		// If not JSON, show first 200 chars
		bodyStr := string(bodyBytes)
		if len(bodyStr) > 200 {
			bodyStr = bodyStr[:200] + "..."
		}
		logData["body"] = bodyStr
	}

	if jsonBytes, err := json.Marshal(logData); err == nil {
		fmt.Printf("SDK Response: %s\n", string(jsonBytes))
	}
}

func (c *Client) logAPIError(resp *http.Response, reqStartTime time.Time, data map[string]interface{}, params map[string]string) {
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	var respBody map[string]interface{}
	msg := string(bodyBytes)
	traceIdFromBody := ""

	if json.Unmarshal(bodyBytes, &respBody) == nil {
		// Extract msg
		if msgStr, ok := respBody["msg"].(string); ok {
			msg = msgStr
		} else if messageStr, ok := respBody["message"].(string); ok {
			msg = messageStr
		}

		// Extract traceId from response body
		if traceStr, ok := respBody["traceId"].(string); ok {
			traceIdFromBody = traceStr
		}
	}

	errorParam := make(map[string]interface{})
	for k, v := range data {
		errorParam[k] = v
	}
	for k, v := range params {
		errorParam[k] = v
	}

	apiErr := map[string]interface{}{
		"code":         resp.StatusCode,
		"traceId":      traceIdFromBody,
		"errorParam":   errorParam,
		"msg":          msg,
		"requestTime":  reqStartTime.Format(time.RFC3339),
		"responseTime": time.Now().Format(time.RFC3339),
	}

	if c.sdkConfig != nil && c.sdkConfig.Logging.ErrorLogFormat == "text" {
		fmt.Printf("API Error: code=%d, traceId=%s, msg=%s, requestTime=%s, responseTime=%s\n",
			apiErr["code"], apiErr["traceId"], apiErr["msg"], apiErr["requestTime"], apiErr["responseTime"])
	} else {
		if jsonBytes, err := json.Marshal(apiErr); err == nil {
			fmt.Printf("API Error: %s\n", string(jsonBytes))
		}
	}
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

func (c *Client) PrepareWithdrawSignInfo(ctx context.Context, params asset.PrepareWithdrawSignInfoParams) (*asset.PreparedWithdrawSignInfo, error) {
	metadataResp, err := c.GetMetaData(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get metadata: %w", err)
	}

	return c.Asset.PrepareWithdrawSignInfo(ctx, metadataResp.Data, params)
}

func (c *Client) CreateCrossWithdrawAuto(ctx context.Context, params *asset.CreateCrossWithdrawAutoParams) (*asset.ResultCreateCrossWithdraw, error) {
	metadataResp, err := c.GetMetaData(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get metadata: %w", err)
	}

	return c.Asset.CreateCrossWithdrawAuto(ctx, params, metadataResp.Data)
}

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

// GetAccountPage gets paginated account list (v2 endpoint).
func (c *Client) GetAccountPage(ctx context.Context, params account.GetAccountPageParams) (*account.PageDataAccountResponse, error) {
	return c.Account.GetAccountPage(ctx, params)
}

// GetAccountAssetSnapshotPage gets account asset snapshots with pagination
func (c *Client) GetAccountAssetSnapshotPage(ctx context.Context, params account.GetAccountAssetSnapshotPageParams) (*account.PageDataAccountAssetSnapshotResponse, error) {
	return c.Account.GetAccountAssetSnapshotPage(ctx, params)
}

// GetPositionByContractID gets positions by contract IDs
func (c *Client) GetPositionByContractID(ctx context.Context, contractIDs []string) (*account.ListPositionResponse, error) {
	return c.Account.GetPositionByContractID(ctx, contractIDs)
}

// GetCollateralByCoinID gets collaterals by coin IDs
func (c *Client) GetCollateralByCoinID(ctx context.Context, coinIDs []string) (*account.ListCollateralResponse, error) {
	return c.Account.GetCollateralByCoinID(ctx, coinIDs)
}

// GetPositionOrders gets orders related to a position
func (c *Client) GetPositionOrders(ctx context.Context, params account.GetPositionOrdersParams) (*account.GetPositionOrdersResponse, error) {
	return c.Account.GetPositionOrders(ctx, &params)
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
