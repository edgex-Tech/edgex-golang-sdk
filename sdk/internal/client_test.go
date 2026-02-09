package internal

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClientDefaultsToHMACForV2(t *testing.T) {
	client, err := NewClient(&ClientConfig{
		BaseURL:    "https://example.com",
		APIVersion: APIVersionV2,
	})
	require.NoError(t, err)
	assert.Equal(t, SigningMethodHMAC, client.GetSigningMethod())
}

func TestNewClientNormalizesBaseURL(t *testing.T) {
	client, err := NewClient(&ClientConfig{
		BaseURL: "https://example.com/api/",
	})
	require.NoError(t, err)
	assert.Equal(t, "https://example.com", client.GetBaseURL())
}

func TestHttpRequestV2HMACPrivate(t *testing.T) {
	const (
		headerKey  = "edgeX"
		apiKey     = "test-api-key"
		passphrase = "test-passphrase"
		secret     = "test-secret"
	)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v2/private/account/getAccountAsset", r.URL.Path)
		assert.Equal(t, "1", r.URL.Query().Get("accountId"))

		timestamp := r.Header.Get(fmt.Sprintf("X-%s-Timestamp", headerKey))
		require.NotEmpty(t, timestamp)
		assert.Equal(t, apiKey, r.Header.Get(fmt.Sprintf("X-%s-Api-Key", headerKey)))
		assert.Equal(t, passphrase, r.Header.Get(fmt.Sprintf("X-%s-Passphrase", headerKey)))

		expectedSig := expectedHMACSignature(secret, timestamp, "GET", "/api/v2/private/account/getAccountAsset", "accountId=1")
		assert.Equal(t, expectedSig, r.Header.Get(fmt.Sprintf("X-%s-Signature", headerKey)))
		assert.Empty(t, r.Header.Get("X-edgeX-Api-Signature"))

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code":"SUCCESS","data":{}}`))
	}))
	defer server.Close()

	client, err := NewClient(&ClientConfig{
		BaseURL:       server.URL,
		APIVersion:    APIVersionV2,
		SigningMethod: SigningMethodHMAC,
		APIKey:        apiKey,
		APIPassphrase: passphrase,
		APISecret:     secret,
		AuthHeaderKey: headerKey,
	})
	require.NoError(t, err)

	url := fmt.Sprintf("%s/api/v1/private/account/getAccountAsset", client.GetBaseURL())
	resp, err := client.HttpRequest(url, "GET", nil, map[string]string{"accountId": "1"})
	require.NoError(t, err)
	resp.Body.Close()
}

func TestHttpRequestV2PublicSkipsAuthHeaders(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v2/public/meta/getMetaData", r.URL.Path)
		assert.Empty(t, r.Header.Get("X-edgeX-Api-Key"))
		assert.Empty(t, r.Header.Get("X-edgeX-Passphrase"))
		assert.Empty(t, r.Header.Get("X-edgeX-Signature"))
		assert.Empty(t, r.Header.Get("X-edgeX-Timestamp"))
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code":"SUCCESS","data":{}}`))
	}))
	defer server.Close()

	client, err := NewClient(&ClientConfig{
		BaseURL:       server.URL,
		APIVersion:    APIVersionV2,
		SigningMethod: SigningMethodHMAC,
	})
	require.NoError(t, err)

	url := fmt.Sprintf("%s/api/v1/public/meta/getMetaData", client.GetBaseURL())
	resp, err := client.HttpRequest(url, "GET", nil, nil)
	require.NoError(t, err)
	resp.Body.Close()
}

func TestHttpRequestV2PrivateRequiresCredentials(t *testing.T) {
	client, err := NewClient(&ClientConfig{
		BaseURL:       "https://example.com",
		APIVersion:    APIVersionV2,
		SigningMethod: SigningMethodHMAC,
	})
	require.NoError(t, err)

	_, err = client.HttpRequest("https://example.com/api/v1/private/account/getAccountAsset", "GET", nil, map[string]string{"accountId": "1"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "hmac credentials are required")
}

func expectedHMACSignature(secret, timestamp, method, requestURI, requestBody string) string {
	base64Key := base64.StdEncoding.EncodeToString([]byte(secret))
	mac := hmac.New(sha256.New, []byte(base64Key))
	mac.Write([]byte(timestamp + method + requestURI + requestBody))
	return hex.EncodeToString(mac.Sum(nil))
}
