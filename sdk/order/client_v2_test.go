package order

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/edgex-Tech/edgex-golang-sdk/sdk/internal"
	metadatapkg "github.com/edgex-Tech/edgex-golang-sdk/sdk/metadata"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateOrderV2UsesEIP712Signature(t *testing.T) {
	var gotPath string
	var gotBody map[string]interface{}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		defer r.Body.Close()
		_ = json.NewDecoder(r.Body).Decode(&gotBody)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code":"SUCCESS","data":{"orderId":"123"}}`))
	}))
	defer srv.Close()

	baseClient, err := internal.NewClient(&internal.ClientConfig{
		BaseURL:       srv.URL,
		AccountID:     42,
		APIVersion:    internal.APIVersionV2,
		SigningMethod: internal.SigningMethodHMAC,
		APIKey:        "api-key",
		APIPassphrase: "pass",
		APISecret:     "secret",
		TradingPriKey: "0x59c6995e998f97a5a0044966f0945380f7d7f4fbcbe8f85f8f19853f51e7f7b4",
	})
	require.NoError(t, err)

	client := NewClient(baseClient)

	clientOrderID := "order-v2-fixed-id"
	expireTime := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := client.CreateOrder(context.Background(), &CreateOrderParams{
		ContractId:    "10000001",
		Price:         "2000",
		Size:          "0.1",
		Type:          OrderTypeLimit,
		Side:          OrderSideBuy,
		TimeInForce:   string(TimeInForce_GOOD_TIL_CANCEL),
		ClientOrderId: &clientOrderID,
		ExpireTime:    expireTime,
	}, &metadatapkg.MetaData{
		Global: &metadatapkg.Global{
			NativeChainId:   "42161",
			ContractAddress: "0x0000000000000000000000000000000000000001",
		},
		CoinList: []metadatapkg.Coin{
			{CoinId: "2", Resolution: "1000000"},
		},
		ContractList: []metadatapkg.Contract{
			{
				ContractId:          "10000001",
				QuoteCoinId:         "2",
				Resolution:          "1000000000",
				DefaultTakerFeeRate: "0.0005",
				DefaultMakerFeeRate: "0.0002",
			},
		},
	}, decimal.RequireFromString("2000"))
	require.NoError(t, err)
	require.NotNil(t, res)
	require.NotNil(t, res.Data)
	assert.Equal(t, "/api/v2/private/order/createOrder", gotPath)

	assert.Equal(t, "order-v2-fixed-id", gotBody["clientOrderId"])
	assert.Equal(t, strconv.FormatInt(internal.CalcNonce(clientOrderID), 10), gotBody["l2Nonce"])
	assert.Equal(t, "200", gotBody["l2Value"])
	assert.Equal(t, "1", gotBody["l2LimitFee"])
	assert.Equal(t, strconv.FormatInt(expireTime.UnixMilli(), 10), gotBody["expireTime"])
	assert.Equal(t, strconv.FormatInt(expireTime.UnixMilli()+8*24*60*60*1000, 10), gotBody["l2ExpireTime"])

	l2Signature, ok := gotBody["l2Signature"].(string)
	require.True(t, ok)
	assert.True(t, strings.HasPrefix(l2Signature, "0x"))
	assert.Len(t, l2Signature, 132)
}

func TestCreateOrderV2RequiresTradingPrivateKey(t *testing.T) {
	baseClient, err := internal.NewClient(&internal.ClientConfig{
		BaseURL:       "http://127.0.0.1",
		AccountID:     42,
		APIVersion:    internal.APIVersionV2,
		SigningMethod: internal.SigningMethodHMAC,
		APIKey:        "api-key",
		APIPassphrase: "pass",
		APISecret:     "secret",
	})
	require.NoError(t, err)

	client := NewClient(baseClient)
	_, err = client.CreateOrder(context.Background(), &CreateOrderParams{
		ContractId: "10000001",
		Price:      "2000",
		Size:       "0.1",
		Type:       OrderTypeLimit,
		Side:       OrderSideBuy,
	}, &metadatapkg.MetaData{
		Global: &metadatapkg.Global{
			NativeChainId:   "42161",
			ContractAddress: "0x0000000000000000000000000000000000000001",
		},
		CoinList: []metadatapkg.Coin{
			{CoinId: "2", Resolution: "1000000"},
		},
		ContractList: []metadatapkg.Contract{
			{
				ContractId:          "10000001",
				QuoteCoinId:         "2",
				Resolution:          "1000000000",
				DefaultTakerFeeRate: "0.0005",
				DefaultMakerFeeRate: "0.0002",
			},
		},
	}, decimal.RequireFromString("2000"))
	require.Error(t, err)
	assert.Contains(t, err.Error(), "trading private key is required")
}

func TestParseResolutionDecimalSupportsHexAndFallback(t *testing.T) {
	res, err := parseResolutionDecimal("0xf4240", "")
	require.NoError(t, err)
	assert.Equal(t, "1000000", res.String())

	res, err = parseResolutionDecimal("bad-resolution", "1000000000")
	require.NoError(t, err)
	assert.Equal(t, "1000000000", res.String())
}
