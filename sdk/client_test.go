package sdk

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/edgex-Tech/edgex-golang-sdk/v2/sdk/metadata"
	"github.com/edgex-Tech/edgex-golang-sdk/v2/sdk/order"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClient_DefaultMetadataCacheTTL(t *testing.T) {
	client, err := NewClient(&ClientConfig{
		BaseURL:   "https://example.com",
		AccountID: 1,
	})
	require.NoError(t, err)

	require.NotNil(t, client.metadataCacheTTL)
	assert.Equal(t, defaultMetaDataCacheTTL, *client.metadataCacheTTL)
}

func TestNewClient_UsesProvidedMetadataCacheTTL(t *testing.T) {
	customTTL := 5 * time.Minute
	client, err := NewClient(&ClientConfig{
		BaseURL:          "https://example.com",
		AccountID:        1,
		MetaDataCacheTTL: &customTTL,
	})
	require.NoError(t, err)

	require.NotNil(t, client.metadataCacheTTL)
	assert.Equal(t, customTTL, *client.metadataCacheTTL)
}

func TestNewClient_InitializesUnifiedAssetNamespace(t *testing.T) {
	client, err := NewClient(&ClientConfig{
		BaseURL:   "https://example.com",
		AccountID: 1,
	})
	require.NoError(t, err)

	require.NotNil(t, client.UnifiedAsset)
}

func TestCreateConditionalMarketOrderUsesMarketProtectionPrice(t *testing.T) {
	tests := []struct {
		name            string
		orderType       order.OrderType
		side            string
		expectedL2Value string
	}{
		{name: "stop market sell", orderType: order.OrderTypeStopMarket, side: order.OrderSideSell, expectedL2Value: "0.079"},
		{name: "take profit market sell", orderType: order.OrderTypeTakeProfitMarket, side: order.OrderSideSell, expectedL2Value: "0.079"},
		{name: "stop market buy", orderType: order.OrderTypeStopMarket, side: order.OrderSideBuy, expectedL2Value: "1917.33"},
		{name: "take profit market buy", orderType: order.OrderTypeTakeProfitMarket, side: order.OrderSideBuy, expectedL2Value: "1917.33"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotBody map[string]interface{}
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				switch r.URL.Path {
				case "/api/v2/public/quote/getTicker":
					_, _ = w.Write([]byte(`{"code":"SUCCESS","data":[{"oraclePrice":"2.427"}]}`))
				case "/api/v2/private/order/createOrder":
					if err := json.NewDecoder(r.Body).Decode(&gotBody); err != nil {
						t.Errorf("decode order request: %v", err)
						http.Error(w, err.Error(), http.StatusBadRequest)
						return
					}
					_, _ = w.Write([]byte(`{"code":"SUCCESS","data":{"orderId":"123"}}`))
				default:
					http.NotFound(w, r)
				}
			}))
			defer srv.Close()

			client, err := NewClient(&ClientConfig{
				BaseURL:       srv.URL,
				AccountID:     42,
				SignerPriKey:  "59c6995e998f97a5a0044966f0945380f7d7f4fbcbe8f85f8f19853f51e7f7b4",
				APIKey:        "test-api-key",
				APIPassphrase: "test-passphrase",
				APISecret:     "test-secret",
			})
			require.NoError(t, err)

			client.metadataCache = &metadata.ResultMetaData{Data: &metadata.MetaData{
				Global: &metadata.Global{
					NativeChainId:   "42161",
					ContractAddress: "0x0000000000000000000000000000000000000001",
				},
				CoinList: []metadata.Coin{{CoinId: "2", Resolution: "1000000"}},
				ContractList: []metadata.Contract{{
					ContractId:          "30000043",
					QuoteCoinId:         "2",
					TickSize:            "0.001",
					Resolution:          "100000000",
					DefaultTakerFeeRate: "0.00043",
					DefaultMakerFeeRate: "0.00038",
				}},
			}}
			client.metadataCacheTime = time.Now()

			_, err = client.CreateOrder(context.Background(), &order.CreateOrderParams{
				ContractId:       "30000043",
				Price:            "0",
				Size:             "79",
				Type:             tt.orderType,
				Side:             tt.side,
				TriggerPrice:     "2.481",
				TriggerPriceType: string(order.TriggerPriceType_LAST_PRICE),
				ReduceOnly:       true,
			})
			require.NoError(t, err)

			assert.Equal(t, "0", gotBody["price"])
			assert.Equal(t, "2.481", gotBody["triggerPrice"])
			assert.Equal(t, string(order.TriggerPriceType_LAST_PRICE), gotBody["triggerPriceType"])
			assert.Equal(t, string(order.TimeInForce_IMMEDIATE_OR_CANCEL), gotBody["timeInForce"])
			assert.Equal(t, tt.expectedL2Value, gotBody["l2Value"])
			assert.Equal(t, "79", gotBody["l2Size"])
		})
	}
}
