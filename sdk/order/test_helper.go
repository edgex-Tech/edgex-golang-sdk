package order

import (
	"net/http"
	"strings"

	"github.com/edgex-Tech/edgex-golang-sdk/sdk/internal"
)

// mockClient implements clientInterface for testing
type mockClient struct {
	accountID    int64
	signerPriKey string
	baseURL      string
	signerAddr   string
	httpFunc     func(urlStr string, method string, data map[string]interface{}, params map[string]string) (*http.Response, error)
}

func newMockClient(accountID int64, signerPriKey, baseURL string, httpFunc func(string, string, map[string]interface{}, map[string]string) (*http.Response, error)) *mockClient {
	return &mockClient{
		accountID:    accountID,
		signerPriKey: strings.TrimPrefix(signerPriKey, "0x"),
		baseURL:      baseURL,
		httpFunc:     httpFunc,
	}
}

func (m *mockClient) GetAccountID() int64 {
	return m.accountID
}

func (m *mockClient) GetSignerPriKey() string {
	return m.signerPriKey
}

func (m *mockClient) GetBaseURL() string {
	return m.baseURL
}

func (m *mockClient) ResolveSignerAddress() (string, error) {
	if m.signerAddr != "" {
		return m.signerAddr, nil
	}
	return internal.DeriveAddressFromPrivateKey(m.signerPriKey)
}

func (m *mockClient) SignTypedDataWithSignerKey(typedData internal.TypedData) (string, error) {
	return internal.SignTypedDataWithPrivateKey(m.signerPriKey, typedData)
}

func (m *mockClient) HttpRequest(urlStr string, method string, data map[string]interface{}, params map[string]string) (*http.Response, error) {
	return m.httpFunc(urlStr, method, data, params)
}
