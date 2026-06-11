package account

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/edgex-Tech/edgex-golang-sdk/v2/sdk/internal"
	metadatapkg "github.com/edgex-Tech/edgex-golang-sdk/v2/sdk/metadata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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

func TestSetMarginModeUsesTradingSignature(t *testing.T) {
	var gotPath string
	var gotBody map[string]interface{}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		defer r.Body.Close()
		_ = json.NewDecoder(r.Body).Decode(&gotBody)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code":"SUCCESS","data":{"updated":true}}`))
	}))
	defer srv.Close()

	mockCli := newMockClient(42, "0x59c6995e998f97a5a0044966f0945380f7d7f4fbcbe8f85f8f19853f51e7f7b4", srv.URL, func(urlStr, method string, data map[string]interface{}, params map[string]string) (*http.Response, error) {
		payload, _ := json.Marshal(data)
		req, _ := http.NewRequest(method, urlStr, bytes.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")
		return srv.Client().Do(req)
	})

	client := NewClient(mockCli)
	clientOrderID := "margin-fixed-id"

	result, err := client.SetMarginMode(context.Background(), &SetMarginModeParams{
		ContractID:    "10000001",
		MarginMode:    "1",
		ClientOrderID: clientOrderID,
	}, &metadatapkg.MetaData{
		Global: &metadatapkg.Global{
			NativeChainId:   "42161",
			ContractAddress: "0x0000000000000000000000000000000000000001",
		},
	})
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "/api/v2/private/account/setMarginMode", gotPath)
	assert.Equal(t, "42", gotBody["accountId"])
	assert.Equal(t, "10000001", gotBody["contractId"])
	assert.Equal(t, "1", gotBody["marginMode"])
	assert.NotContains(t, gotBody, "clientOrderId")
	assert.Equal(t, strconv.FormatInt(internal.CalcNonce(clientOrderID), 10), gotBody["l2Nonce"])
	assert.NotEmpty(t, gotBody["l2ExpireTime"])

	signer, ok := gotBody["signer"].(string)
	require.True(t, ok)
	assert.True(t, strings.HasPrefix(signer, "0x"))
	assert.Len(t, signer, 42)

	signature, ok := gotBody["l2Signature"].(string)
	require.True(t, ok)
	assert.True(t, strings.HasPrefix(signature, "0x"))
	assert.Len(t, signature, 132)
}

func TestSetMarginModeRequiresTradingPrivateKey(t *testing.T) {
	mockCli := newMockClient(42, "", "http://127.0.0.1", func(urlStr string, method string, data map[string]interface{}, params map[string]string) (*http.Response, error) {
		return nil, assert.AnError
	})

	client := NewClient(mockCli)
	_, err := client.SetMarginMode(context.Background(), &SetMarginModeParams{
		ContractID: "10000001",
		MarginMode: "0",
	}, &metadatapkg.MetaData{
		Global: &metadatapkg.Global{
			NativeChainId:   "42161",
			ContractAddress: "0x0000000000000000000000000000000000000001",
		},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "trading private key is required")
}

func TestGetAccountAssetUnmarshalExpandedFields(t *testing.T) {
	responseBody := `{
		"code":"SUCCESS",
		"data":{
			"account":{
				"id":"1",
				"userId":"2",
				"status":"NORMAL",
				"signers":["0xabc"],
				"contractIdToMarginMode":{"30000051":"0"}
			},
			"version":"9",
			"positionList":[
				{
					"contractId":"30000051",
					"openSize":"1.5",
					"marginMode":"CROSS",
					"longTermStat":{"cumOpenSize":"2"}
				}
			],
			"collateralList":[
				{"coinId":"1000","amount":"12.3","cumTransferOutAmount":"1"}
			],
			"positionAssetList":[
				{"contractId":"30000051","positionValue":"100"}
			],
			"collateralAssetModelList":[
				{"coinId":"1000","walletBalance":"20"}
			],
			"oraclePriceList":[
				{"contractId":"30000051","priceType":"ORACLE_PRICE","priceValue":"2500"}
			],
			"markPriceList":[
				{"contractId":"30000051","priceType":"MARK_PRICE","priceValue":"2499"}
			]
		}
	}`

	mockCli := newMockClient(42, "", "https://example.com", func(urlStr string, method string, data map[string]interface{}, params map[string]string) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(responseBody)),
			Header:     make(http.Header),
		}, nil
	})

	client := NewClient(mockCli)
	resp, err := client.GetAccountAsset(context.Background())
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, resp.Data)
	require.NotNil(t, resp.Data.Account)
	assert.Equal(t, "NORMAL", resp.Data.Account.Status)
	assert.Equal(t, "0xabc", resp.Data.Account.Signers[0])
	assert.Equal(t, "0", resp.Data.Account.ContractIDToMarginMode["30000051"])
	assert.Equal(t, "9", resp.Data.Version)
	assert.Equal(t, "1.5", resp.Data.PositionList[0].OpenSize)
	require.NotNil(t, resp.Data.PositionList[0].LongTermStat)
	assert.Equal(t, "2", resp.Data.PositionList[0].LongTermStat.CumOpenSize)
	assert.Equal(t, "1", resp.Data.CollateralList[0].CumTransferOutAmount)
	assert.Equal(t, "100", resp.Data.PositionAssetList[0].PositionValue)
	assert.Equal(t, "20", resp.Data.CollateralAssetModelList[0].WalletBalance)
	assert.Equal(t, "2500", resp.Data.OraclePriceList[0].PriceValue)
	assert.Equal(t, "2499", resp.Data.MarkPriceList[0].PriceValue)
}
