package unified_asset

import (
	"bytes"
	"context"
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/edgex-Tech/edgex-golang-sdk/sdk/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
		walletPriKey: walletPriKey,
		baseURL:      baseURL,
		httpFunc:     httpFunc,
	}
}

func (m *mockClient) GetAccountID() int64 {
	return m.accountID
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

func TestBuildWithdrawAttemptUsesRawAmountAndZeroPrivy(t *testing.T) {
	attempt := buildWithdrawAttempt(
		"0xFCAd0B19bB29D4674531d6f115237E16AfCE377c",
		"spot",
		"12345",
		"0x98d2919b9A214E6Fa5384AC81E6864bA686Ad74c",
		"1000",
		3343,
		123456,
		"849849126827855872",
		ZERO_ADDRESS,
	)

	assert.Equal(t, "spot", attempt["source"])
	assert.Equal(t, "12345", attempt["sourceAccount"])
	assert.Equal(t, "1000", attempt["amount"])
	assert.Equal(t, "0", attempt["fee"])
	assert.Equal(t, "chain-3343", attempt["destination"])
	assert.Equal(t, ZERO_ADDRESS, attempt["privyAddress"])
}

func TestCreateWithdrawUsesUnifiedAssetV1Paths(t *testing.T) {
	var paths []string
	var submitBody map[string]interface{}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paths = append(paths, r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/api/v1/private/unified-asset/getFeeByAssetFlow":
			_, _ = w.Write([]byte(`{"code":"SUCCESS","data":{"fee":"10"}}`))
		case "/api/v1/private/unified-asset/getEIP712Data":
			_, _ = w.Write([]byte(`{"code":"SUCCESS","data":{"types":{"AssetFlowAttempt":{"fields":[{"name":"amount","type":"uint256"}]}},"primaryType":"AssetFlowAttempt","domain":{"name":"EdgeX","version":"1"},"messageJson":"{\"amount\":\"990\"}"}}`))
		case "/api/v1/private/unified-asset/submitAssetFlow":
			defer r.Body.Close()
			_ = json.NewDecoder(r.Body).Decode(&submitBody)
			_, _ = w.Write([]byte(`{"code":"SUCCESS","data":{"flowId":"1"}}`))
		default:
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
	}))
	defer srv.Close()

	mockCli := newMockClient(12345, "59c6995e998f97a5a0044966f0945380f7d7f4fbcbe8f85f8f19853f51e7f7b4", srv.URL, func(urlStr, method string, data map[string]interface{}, params map[string]string) (*http.Response, error) {
		payload, _ := json.Marshal(data)
		req, _ := http.NewRequest(method, urlStr, bytes.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")
		return srv.Client().Do(req)
	})

	client := NewClient(mockCli)
	result, err := client.CreateWithdraw(context.Background(), CreateWithdrawParams{
		AmountRaw:        "1000",
		UserAddress:      "0xFCAd0B19bB29D4674531d6f115237E16AfCE377c",
		TokenAddress:     "0x98d2919b9A214E6Fa5384AC81E6864bA686Ad74c",
		ChainID:          3343,
		ClientWithdrawID: "849849126827855872",
		ExpireTime:       123456,
	})
	require.NoError(t, err)
	assert.Equal(t, []string{
		"/api/v1/private/unified-asset/getFeeByAssetFlow",
		"/api/v1/private/unified-asset/getEIP712Data",
		"/api/v1/private/unified-asset/submitAssetFlow",
	}, paths)
	require.NotNil(t, result)
	assert.Equal(t, "1", result["flowId"])
	assert.Equal(t, "10", submitBody["attempt"].(map[string]interface{})["fee"])
	assert.Equal(t, "990", submitBody["attempt"].(map[string]interface{})["amount"])
	assert.Contains(t, submitBody["userSignature"].(string), "0x")
}

func TestBuildTypedDataFromServerResponseNormalizesFields(t *testing.T) {
	typedData, err := buildTypedDataFromServerResponse(map[string]interface{}{
		"types": map[string]interface{}{
			"AssetFlowAttempt": []interface{}{
				map[string]interface{}{"name": "amount", "type": "uint256"},
				map[string]interface{}{"name": "deadline", "type": "uint256"},
			},
		},
		"primaryType": "AssetFlowAttempt",
		"domain": map[string]interface{}{
			"name":              "EdgeX",
			"version":           "1",
			"chainId":           "0xd0f",
			"verifyingContract": "0x0000000000000000000000000000000000000001",
		},
		"messageJson": "{\"amount\":\"990\",\"deadline\":\"123456\"}",
	})
	require.NoError(t, err)

	require.Contains(t, typedData.Types, "EIP712Domain")
	assert.Equal(t, "AssetFlowAttempt", typedData.PrimaryType)
	assert.Equal(t, "990", typedData.Message["amount"])
	assert.Equal(t, "123456", typedData.Message["deadline"])
	assert.Equal(t, "EdgeX", typedData.Domain.Name)
	assert.Equal(t, "1", typedData.Domain.Version)
	require.NotNil(t, typedData.Domain.ChainId)
	assert.Equal(t, "3343", (*big.Int)(typedData.Domain.ChainId).String())
	assert.Equal(t, "0x0000000000000000000000000000000000000001", typedData.Domain.VerifyingContract)
}

func TestGetSpotDepositDataUsesUnifiedAssetPath(t *testing.T) {
	var gotPath string

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code":"SUCCESS","data":{"chainId":33431}}`))
	}))
	defer srv.Close()

	mockCli := newMockClient(12345, "59c6995e998f97a5a0044966f0945380f7d7f4fbcbe8f85f8f19853f51e7f7b4", srv.URL, func(urlStr, method string, data map[string]interface{}, params map[string]string) (*http.Response, error) {
		payload, _ := json.Marshal(data)
		req, _ := http.NewRequest(method, urlStr, bytes.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")
		return srv.Client().Do(req)
	})

	client := NewClient(mockCli)
	result, err := client.GetSpotDepositData(context.Background(), CreateSpotDepositParams{
		AmountRaw:     "1000",
		UserAddress:   "0xFCAd0B19bB29D4674531d6f115237E16AfCE377c",
		TokenAddress:  "0x98d2919b9A214E6Fa5384AC81E6864bA686Ad74c",
		ChainID:       33431,
		SpotAccountID: "12345",
	})
	require.NoError(t, err)
	assert.Equal(t, "/api/v1/private/unified-asset/getDepositData", gotPath)
	assert.Equal(t, float64(33431), result["chainId"])
}
