package internal

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"golang.org/x/crypto/sha3"

	"github.com/edgex-Tech/edgex-golang-sdk/starkcurve"
)

// Client represents the base client with common functionality
type Client struct {
	httpClient    *http.Client
	baseURL       string
	accountID     int64
	starkPriKey   string
	tradingPriKey string
	walletPriKey  string
	tradingAddr   string
	walletAddr    string
	apiVersion    string
	signingMethod string
	apiKey        string
	apiPassphrase string
	apiSecret     string
	authHeaderKey string
}

// ClientConfig holds the configuration for creating a new Client
type ClientConfig struct {
	BaseURL       string
	AccountID     int64
	StarkPriKey   string
	TradingPriKey string
	WalletPriKey  string
	TradingAddr   string
	WalletAddr    string
	APIVersion    string
	SigningMethod string
	APIKey        string
	APIPassphrase string
	APISecret     string
	AuthHeaderKey string
}

// NewClient creates a new base client
func NewClient(cfg *ClientConfig) (*Client, error) {
	apiVersion := strings.ToLower(strings.TrimSpace(cfg.APIVersion))
	if apiVersion == "" {
		apiVersion = APIVersionV1
	}
	if apiVersion != APIVersionV1 && apiVersion != APIVersionV2 {
		return nil, fmt.Errorf("unsupported api version: %s", cfg.APIVersion)
	}

	signingMethod := strings.ToLower(strings.TrimSpace(cfg.SigningMethod))
	if signingMethod == "" {
		if apiVersion == APIVersionV2 {
			signingMethod = SigningMethodHMAC
		} else {
			signingMethod = SigningMethodStark
		}
	}
	if signingMethod != SigningMethodStark && signingMethod != SigningMethodHMAC {
		return nil, fmt.Errorf("unsupported signing method: %s", cfg.SigningMethod)
	}

	authHeaderKey := strings.TrimSpace(cfg.AuthHeaderKey)
	if authHeaderKey == "" {
		authHeaderKey = DefaultHeaderKey
	}

	return &Client{
		httpClient:    &http.Client{Timeout: 30 * time.Second},
		baseURL:       normalizeBaseURL(cfg.BaseURL),
		accountID:     cfg.AccountID,
		starkPriKey:   strings.TrimPrefix(cfg.StarkPriKey, "0x"),
		tradingPriKey: strings.TrimPrefix(cfg.TradingPriKey, "0x"),
		walletPriKey:  strings.TrimPrefix(cfg.WalletPriKey, "0x"),
		tradingAddr:   strings.TrimSpace(cfg.TradingAddr),
		walletAddr:    strings.TrimSpace(cfg.WalletAddr),
		apiVersion:    apiVersion,
		signingMethod: signingMethod,
		apiKey:        cfg.APIKey,
		apiPassphrase: cfg.APIPassphrase,
		apiSecret:     cfg.APISecret,
		authHeaderKey: authHeaderKey,
	}, nil
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

// GetStarkPriKey returns the stark private key
func (c *Client) GetStarkPriKey() string {
	return c.starkPriKey
}

// GetTradingPriKey returns the EIP-712 trading private key.
func (c *Client) GetTradingPriKey() string {
	return c.tradingPriKey
}

// GetWalletPriKey returns the EIP-712 wallet private key.
func (c *Client) GetWalletPriKey() string {
	return c.walletPriKey
}

// GetTradingAddr returns the configured trading signer address.
func (c *Client) GetTradingAddr() string {
	return c.tradingAddr
}

// GetWalletAddr returns the configured wallet signer address.
func (c *Client) GetWalletAddr() string {
	return c.walletAddr
}

// GetBaseURL returns the base URL
func (c *Client) GetBaseURL() string {
	return c.baseURL
}

// GetAPIVersion returns the configured API version
func (c *Client) GetAPIVersion() string {
	return c.apiVersion
}

// GetSigningMethod returns the configured request signing method
func (c *Client) GetSigningMethod() string {
	return c.signingMethod
}

// GetAuthHeaderKey returns the configured auth header key prefix
func (c *Client) GetAuthHeaderKey() string {
	return c.authHeaderKey
}

// HttpRequest makes an authenticated HTTP request
func (c *Client) HttpRequest(urlStr string, method string, data map[string]interface{}, params map[string]string) (*http.Response, error) {
	method = strings.ToUpper(method)

	urlStr, err := c.rewriteURLByVersion(urlStr)
	if err != nil {
		return nil, fmt.Errorf("failed to rewrite URL by version: %w", err)
	}

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

	// Add authentication headers
	if c.signingMethod == SigningMethodHMAC {
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
	} else {
		signContent := c.buildSignatureContent(timestamp, method, signPath, data, params)

		// Calculate Keccak256 hash
		hash := sha3.NewLegacyKeccak256()
		hash.Write([]byte(signContent))
		contentHash := hash.Sum(nil)

		sig, err := c.Sign(contentHash)
		if err != nil {
			return nil, fmt.Errorf("failed to sign request: %w", err)
		}

		req.Header.Set("X-edgeX-Api-Timestamp", timestampStr)
		req.Header.Set("X-edgeX-Api-Signature", fmt.Sprintf("%s%s", sig.R, sig.S))
	}

	req.Header.Set("Accept", "application/json")

	// Execute request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
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

func (c *Client) rewriteURLByVersion(rawURL string) (string, error) {
	if c.apiVersion != APIVersionV2 {
		return rawURL, nil
	}

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	parsedURL.Path = rewritePathToV2(parsedURL.Path)
	return parsedURL.String(), nil
}

func rewritePathToV2(path string) string {
	switch {
	case strings.Contains(path, "/api/api/v1/"):
		return strings.Replace(path, "/api/api/v1/", "/api/v2/", 1)
	case strings.HasSuffix(path, "/api/api/v1"):
		return strings.TrimSuffix(path, "/api/api/v1") + "/api/v2"
	case strings.Contains(path, "/api/v1/"):
		return strings.Replace(path, "/api/v1/", "/api/v2/", 1)
	case strings.HasSuffix(path, "/api/v1"):
		return strings.TrimSuffix(path, "/api/v1") + "/api/v2"
	case strings.Contains(path, "/v1/"):
		return strings.Replace(path, "/v1/", "/v2/", 1)
	case strings.HasSuffix(path, "/v1"):
		return strings.TrimSuffix(path, "/v1") + "/v2"
	default:
		return path
	}
}

func (c *Client) shouldSignWithHMAC(path string) bool {
	return strings.Contains(path, "private") ||
		strings.Contains(path, "/opt/") ||
		strings.Contains(path, "invite") ||
		strings.Contains(path, "points") ||
		strings.Contains(path, "/public/codeInfo") ||
		strings.HasPrefix(path, "/ranking") ||
		strings.Contains(path, "nft") ||
		strings.Contains(path, "vault") ||
		strings.HasPrefix(path, "/activity") ||
		strings.Contains(path, "public/info") ||
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
	// Match web implementation: base64(encodeURI(secret)).
	base64Key := base64.StdEncoding.EncodeToString([]byte(c.apiSecret))

	mac := hmac.New(sha256.New, []byte(base64Key))
	mac.Write([]byte(message))
	return hex.EncodeToString(mac.Sum(nil))
}

// buildSignatureContent builds the content string for signature generation
func (c *Client) buildSignatureContent(timestamp int64, method string, path string, data map[string]interface{}, params map[string]string) string {
	if len(data) > 0 {
		// Convert body to sorted string format
		bodyStr := c.getValue(data)
		return fmt.Sprintf("%d%s%s%s", timestamp, method, path, bodyStr)
	}

	// For requests without body, use query parameters if present
	if len(params) > 0 {
		// Sort query parameters
		keys := make([]string, 0, len(params))
		for k := range params {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		var paramPairs []string
		for _, k := range keys {
			paramPairs = append(paramPairs, fmt.Sprintf("%s=%s", k, params[k]))
		}
		queryString := strings.Join(paramPairs, "&")
		return fmt.Sprintf("%d%s%s%s", timestamp, method, path, queryString)
	}

	return fmt.Sprintf("%d%s%s", timestamp, method, path)
}

// getValue converts a value to a string representation for signing
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
		// Convert all values to strings and sort by keys
		sortedMap := make(map[string]string)
		for key, val := range v {
			sortedMap[key] = c.getValue(val)
		}

		// Get sorted keys
		keys := make([]string, 0, len(sortedMap))
		for k := range sortedMap {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		// Build key=value pairs
		var pairs []string
		for _, k := range keys {
			pairs = append(pairs, fmt.Sprintf("%s=%s", k, sortedMap[k]))
		}
		return strings.Join(pairs, "&")
	default:
		return fmt.Sprintf("%v", v)
	}
}

// Sign signs a message hash using the client's Stark private key
func (c *Client) Sign(messageHash []byte) (*L2Signature, error) {
	privateKey := c.GetStarkPriKey()
	if privateKey == "" {
		return nil, fmt.Errorf("stark private key not set")
	}

	privKeyBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to decode private key: %w", err)
	}

	starkPrivKey := big.NewInt(0).SetBytes(privKeyBytes)

	msgHashInt := big.NewInt(0).SetBytes(messageHash)
	msgHashInt = msgHashInt.Mod(msgHashInt, starkcurve.NewStarkCurve().N)

	r, s, err := starkcurve.Sign(starkPrivKey.Bytes(), msgHashInt.Bytes())
	if err != nil {
		return nil, err
	}

	rBytes := append(bytes.Repeat([]byte{0}, 32-len(r.Bytes())), r.Bytes()...)
	sBytes := append(bytes.Repeat([]byte{0}, 32-len(s.Bytes())), s.Bytes()...)

	// Convert r, s and y to hex strings
	signature := &L2Signature{
		R: hex.EncodeToString(rBytes),
		S: hex.EncodeToString(sBytes),
		V: "",
	}

	return signature, nil
}

// ResolveTradingSignerAddress returns the configured trading signer address or derives it from private key.
func (c *Client) ResolveTradingSignerAddress() (string, error) {
	if c.tradingAddr != "" {
		return c.tradingAddr, nil
	}
	if c.tradingPriKey == "" {
		return "", fmt.Errorf("trading signer address/private key not set")
	}
	return DeriveAddressFromPrivateKey(c.tradingPriKey)
}

// ResolveWalletSignerAddress returns the configured wallet signer address or derives it from private key.
func (c *Client) ResolveWalletSignerAddress() (string, error) {
	if c.walletAddr != "" {
		return c.walletAddr, nil
	}
	if c.walletPriKey == "" {
		return "", fmt.Errorf("wallet signer address/private key not set")
	}
	return DeriveAddressFromPrivateKey(c.walletPriKey)
}

// SignTypedDataWithTradingKey signs EIP-712 typed data using the configured trading key.
func (c *Client) SignTypedDataWithTradingKey(typedData TypedData) (string, error) {
	if c.tradingPriKey == "" {
		return "", fmt.Errorf("trading private key not set")
	}
	return SignTypedDataWithPrivateKey(c.tradingPriKey, typedData)
}

// SignTypedDataWithWalletKey signs EIP-712 typed data using the configured wallet key.
func (c *Client) SignTypedDataWithWalletKey(typedData TypedData) (string, error) {
	if c.walletPriKey == "" {
		return "", fmt.Errorf("wallet private key not set")
	}
	return SignTypedDataWithPrivateKey(c.walletPriKey, typedData)
}

// func L2Sign(msgHashStr string) (*L2Signature, error) {
// 	msgHashBig, _ := L2SignUtils.HexToBigInteger(msgHashStr)
// 	privateKeyBig, _ := L2SignUtils.HexToBigInteger(privateKeyStr)
// 	privateKey := ecdsa.Create(privateKeyBig)
// 	return ecdsa.Sign(msgHashBig, privateKey)
// }
