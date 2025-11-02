package internal

import (
	"bytes"
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
	httpClient  *http.Client
	baseURL     string
	accountID   int64
	starkPriKey string
}

// ClientConfig holds the configuration for creating a new Client
type ClientConfig struct {
	BaseURL     string
	AccountID   int64
	StarkPriKey string
}

// NewClient creates a new base client
func NewClient(cfg *ClientConfig) (*Client, error) {
	return &Client{
		httpClient:  &http.Client{Timeout: 30 * time.Second},
		baseURL:     cfg.BaseURL,
		accountID:   cfg.AccountID,
		starkPriKey: cfg.StarkPriKey,
	}, nil
}

// GetAccountID returns the account ID
func (c *Client) GetAccountID() int64 {
	return c.accountID
}

// GetStarkPriKey returns the stark private key
func (c *Client) GetStarkPriKey() string {
	return c.starkPriKey
}

// GetBaseURL returns the base URL
func (c *Client) GetBaseURL() string {
	return c.baseURL
}

// HttpRequest makes an authenticated HTTP request
func (c *Client) HttpRequest(urlStr string, method string, data map[string]interface{}, params map[string]string) (*http.Response, error) {
	// Generate timestamp
	timestamp := time.Now().UnixMilli()

	// Parse URL to extract path
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}
	path := parsedURL.Path
	if parsedURL.RawQuery != "" {
		path = path + "?" + parsedURL.RawQuery
	}

	// Build signature content
	signContent := c.buildSignatureContent(timestamp, method, path, data, params)

	// Calculate Keccak256 hash
	hash := sha3.NewLegacyKeccak256()
	hash.Write([]byte(signContent))
	contentHash := hash.Sum(nil)

	// Sign the hash
	sig, err := c.Sign(contentHash)
	if err != nil {
		return nil, fmt.Errorf("failed to sign request: %w", err)
	}

	// Create request
	var req *http.Request
	if method == "GET" {
		// For GET requests, add params to URL
		if len(params) > 0 {
			q := url.Values{}
			for k, v := range params {
				q.Add(k, v)
			}
			urlStr = urlStr + "?" + q.Encode()
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
	req.Header.Set("X-edgeX-Api-Timestamp", fmt.Sprintf("%d", timestamp))
	req.Header.Set("X-edgeX-Api-Signature", fmt.Sprintf("%s%s", sig.R, sig.S))
	req.Header.Set("Accept", "application/json")

	// Execute request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}

	return resp, nil
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

// func L2Sign(msgHashStr string) (*L2Signature, error) {
// 	msgHashBig, _ := L2SignUtils.HexToBigInteger(msgHashStr)
// 	privateKeyBig, _ := L2SignUtils.HexToBigInteger(privateKeyStr)
// 	privateKey := ecdsa.Create(privateKeyBig)
// 	return ecdsa.Sign(msgHashBig, privateKey)
// }
