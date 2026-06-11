package transfer

import (
	"net/http"
	"strings"

	"github.com/edgex-Tech/edgex-golang-sdk/v2/sdk/internal"
)

// mockClient implements clientInterface for testing
type mockClient struct {
	accountID    int64
	walletPriKey string
	baseURL      string
	walletAddr   string
	httpFunc     func(urlStr string, method string, data map[string]interface{}, params map[string]string) (*http.Response, error)
}

func newMockClient(accountID int64, walletPriKey, baseURL string, httpFunc func(string, string, map[string]interface{}, map[string]string) (*http.Response, error)) *mockClient {
	return &mockClient{
		accountID:    accountID,
		walletPriKey: strings.TrimPrefix(walletPriKey, "0x"),
		baseURL:      baseURL,
		httpFunc:     httpFunc,
	}
}

func (m *mockClient) GetAccountID() int64 {
	return m.accountID
}

func (m *mockClient) GetWalletPriKey() string {
	return m.walletPriKey
}

func (m *mockClient) GetBaseURL() string {
	return m.baseURL
}

func (m *mockClient) ResolveWalletAddress() (string, error) {
	if m.walletAddr != "" {
		return m.walletAddr, nil
	}
	return internal.DeriveAddressFromPrivateKey(m.walletPriKey)
}

func (m *mockClient) SignTypedDataWithWalletKey(typedData internal.TypedData) (string, error) {
	return internal.SignTypedDataWithPrivateKey(m.walletPriKey, typedData)
}

func (m *mockClient) HttpRequest(urlStr string, method string, data map[string]interface{}, params map[string]string) (*http.Response, error) {
	return m.httpFunc(urlStr, method, data, params)
}
