package quote

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
)

type mockClient struct {
	baseURL string
	do      func(urlStr string, method string, data map[string]interface{}, params map[string]string) (*http.Response, error)
}

func (m *mockClient) GetBaseURL() string {
	return m.baseURL
}

func (m *mockClient) HttpRequest(urlStr string, method string, data map[string]interface{}, params map[string]string) (*http.Response, error) {
	return m.do(urlStr, method, data, params)
}

func TestGetQuoteSummaryUsesPeriodQuery(t *testing.T) {
	var gotParams map[string]string

	client := NewClient(&mockClient{
		baseURL: "https://example.com",
		do: func(urlStr string, method string, data map[string]interface{}, params map[string]string) (*http.Response, error) {
			gotParams = params
			return &http.Response{
				StatusCode: http.StatusOK,
				Body: io.NopCloser(strings.NewReader(`{
					"code":"SUCCESS",
					"data":{"tickerSummary":{"period":"LAST_DAY_1","trades":"1","value":"2","openInterest":"3"}}
				}`)),
				Header: make(http.Header),
			}, nil
		},
	})

	resp, err := client.GetQuoteSummary(context.Background(), "LAST_DAY_1")
	if err != nil {
		t.Fatalf("GetQuoteSummary returned error: %v", err)
	}
	if gotParams["period"] != "LAST_DAY_1" {
		t.Fatalf("unexpected period query: got=%q", gotParams["period"])
	}
	if resp == nil || resp.Data == nil || resp.Data.TickerSummary == nil {
		t.Fatalf("unexpected response: %+v", resp)
	}
}

func TestGetEstimatedFeeUsesTimeRangeQuery(t *testing.T) {
	var gotParams map[string]string

	client := NewClient(&mockClient{
		baseURL: "https://example.com",
		do: func(urlStr string, method string, data map[string]interface{}, params map[string]string) (*http.Response, error) {
			gotParams = params
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(`{"code":"SUCCESS","data":[]}`)),
				Header:     make(http.Header),
			}, nil
		},
	})

	_, err := client.GetEstimatedFee(context.Background(), GetEstimatedFeeParams{
		FilterBeginKlineTimeInclusive: 100,
		FilterEndKlineTimeExclusive:   200,
	})
	if err != nil {
		t.Fatalf("GetEstimatedFee returned error: %v", err)
	}
	if gotParams["filterBeginKlineTimeInclusive"] != "100" {
		t.Fatalf("unexpected begin query: got=%q", gotParams["filterBeginKlineTimeInclusive"])
	}
	if gotParams["filterEndKlineTimeExclusive"] != "200" {
		t.Fatalf("unexpected end query: got=%q", gotParams["filterEndKlineTimeExclusive"])
	}
}
