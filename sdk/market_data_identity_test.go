package sdk

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/edgex-Tech/edgex-golang-sdk/v2/sdk/funding"
	"github.com/edgex-Tech/edgex-golang-sdk/v2/sdk/quote"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Recorded with an independent HMAC-SHA256 (key = base64("test-secret")) over
// "1700000000000" + "GET" + path + sorted query.
const (
	identityTimestamp = "1700000000000"
	// GET /api/v2/public/quote/getDepth, accountId=42&contractId=10000001&level=15
	depthIdentitySignature = "66a9037e65c02d8b5cc2c65006ada384a8509d0854d29e6b0ab8a2ec640511b5"
	// GET /api/v2/public/quote/getMarketStatus, accountId=42
	marketStatusIdentitySignature = "e847caa79cab8d46e23e5fa921441e08ee4a5055e762f86184962dcbea8b6780"
)

var identityHeaders = []string{
	"X-edgeX-Api-Key", "X-edgeX-Passphrase", "X-edgeX-Signature", "X-edgeX-Timestamp",
}

func newIdentityTestClient(t *testing.T, cfg ClientConfig) (*Client, *[]*http.Request) {
	t.Helper()
	var requests []*http.Request
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests = append(requests, r)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code":"SUCCESS","data":[]}`))
	}))
	t.Cleanup(srv.Close)

	cfg.BaseURL = srv.URL
	client, err := NewClient(&cfg)
	require.NoError(t, err)
	client.nowFn = func() time.Time { return time.UnixMilli(1700000000000) }
	return client, &requests
}

func fullCredentials() ClientConfig {
	return ClientConfig{
		AccountID:     42,
		APIKey:        "test-api-key",
		APIPassphrase: "test-passphrase",
		APISecret:     "test-secret",
	}
}

func TestMarketDataIdentitySignsDepthLikeAPrivateRequest(t *testing.T) {
	client, requests := newIdentityTestClient(t, fullCredentials())

	_, err := client.Quote.GetOrderBookDepth(context.Background(), quote.GetOrderBookDepthParams{
		ContractID: "10000001",
		Size:       15,
	})
	require.NoError(t, err)

	require.Len(t, *requests, 1)
	got := (*requests)[0]
	assert.Equal(t, "/api/v2/public/quote/getDepth", got.URL.Path)
	assert.Equal(t, "accountId=42&contractId=10000001&level=15", got.URL.RawQuery)
	assert.Equal(t, "test-api-key", got.Header.Get("X-edgeX-Api-Key"))
	assert.Equal(t, "test-passphrase", got.Header.Get("X-edgeX-Passphrase"))
	assert.Equal(t, identityTimestamp, got.Header.Get("X-edgeX-Timestamp"))
	assert.Equal(t, depthIdentitySignature, got.Header.Get("X-edgeX-Signature"))
}

func TestMarketDataIdentityWithNoOtherQueryParameter(t *testing.T) {
	client, requests := newIdentityTestClient(t, fullCredentials())

	resp, err := client.HttpRequest(client.GetBaseURL()+"/api/v2/public/quote/getMarketStatus", "GET", nil, nil)
	require.NoError(t, err)
	resp.Body.Close()

	got := (*requests)[0]
	assert.Equal(t, "accountId=42", got.URL.RawQuery)
	assert.Equal(t, marketStatusIdentitySignature, got.Header.Get("X-edgeX-Signature"))
}

func TestMarketDataIdentityCoversEveryMappedRouteAndNothingElse(t *testing.T) {
	client, requests := newIdentityTestClient(t, fullCredentials())

	mapped := []string{
		"/api/v2/public/quote/getTicker",
		"/api/v2/public/quote/getDepth",
		"/api/v2/public/quote/getKline",
		"/api/v2/public/quote/getMultiContractKline",
		"/api/v2/public/quote/getMultiDayKline",
		"/api/v2/public/quote/getAccurateOpenInterest",
		"/api/v2/public/quote/getMarketStatus",
		"/api/v2/public/funding/getLatestFundingRate",
		"/api/v2/public/funding/getFundingRatePage",
		"/api/v2/public/index/getIndexPriceConfig",
		"/api/v2/public/index/getMarketTradingSchedule",
	}
	assert.Len(t, marketDataIdentityPaths, len(mapped))
	for _, path := range mapped {
		resp, err := client.HttpRequest(client.GetBaseURL()+path, "GET", nil, map[string]string{"contractId": "10000001"})
		require.NoError(t, err)
		resp.Body.Close()
		got := (*requests)[len(*requests)-1]
		assert.Equal(t, "42", got.URL.Query().Get("accountId"), path)
		assert.NotEmpty(t, got.Header.Get("X-edgeX-Signature"), path)
	}

	anonymous := []string{
		"/api/v2/public/meta/getMetaData",
		"/api/v2/public/meta/getServerTime",
		"/api/v2/public/quote/getExchangeLongShortRatio",
		"/api/v2/public/quote/getTicketSummary",
	}
	for _, path := range anonymous {
		resp, err := client.HttpRequest(client.GetBaseURL()+path, "GET", nil, map[string]string{"contractId": "10000001"})
		require.NoError(t, err)
		resp.Body.Close()
		got := (*requests)[len(*requests)-1]
		assert.Equal(t, "contractId=10000001", got.URL.RawQuery, path)
		for _, name := range identityHeaders {
			assert.Empty(t, got.Header.Get(name), path)
		}
	}
}

func TestMarketDataIdentityReplacesACallerSuppliedAccount(t *testing.T) {
	client, requests := newIdentityTestClient(t, fullCredentials())

	resp, err := client.HttpRequest(
		client.GetBaseURL()+"/api/v2/public/quote/getDepth?accountId=7&level=15",
		"GET", nil, map[string]string{"contractId": "10000001"})
	require.NoError(t, err)
	resp.Body.Close()

	got := (*requests)[0]
	assert.Equal(t, []string{"42"}, got.URL.Query()["accountId"])
	assert.Equal(t, depthIdentitySignature, got.Header.Get("X-edgeX-Signature"))
}

func TestMarketDataStaysAnonymousWithoutFullIdentity(t *testing.T) {
	disabled := fullCredentials()
	disabled.DisableMarketDataIdentity = true
	noAccount := fullCredentials()
	noAccount.AccountID = 0
	noSecret := fullCredentials()
	noSecret.APISecret = ""

	cases := map[string]ClientConfig{
		"no credentials":    {},
		"account only":      {AccountID: 42},
		"no account":        noAccount,
		"incomplete key":    noSecret,
		"identity disabled": disabled,
	}
	for name, cfg := range cases {
		t.Run(name, func(t *testing.T) {
			client, requests := newIdentityTestClient(t, cfg)

			_, err := client.Quote.GetOrderBookDepth(context.Background(), quote.GetOrderBookDepthParams{
				ContractID: "10000001",
				Size:       15,
			})
			require.NoError(t, err)
			_, err = client.Funding.GetLatestFundingRate(context.Background(), funding.GetLatestFundingRateParams{
				ContractID: "10000001",
			})
			require.NoError(t, err)

			require.Len(t, *requests, 2)
			assert.Equal(t, "contractId=10000001&level=15", (*requests)[0].URL.RawQuery)
			assert.Equal(t, "contractId=10000001", (*requests)[1].URL.RawQuery)
			for _, got := range *requests {
				for _, header := range identityHeaders {
					assert.Empty(t, got.Header.Get(header))
				}
			}
		})
	}
}
