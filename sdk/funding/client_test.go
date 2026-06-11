package funding

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

func TestGetFundingRateOmitsSettlementFilterByDefault(t *testing.T) {
	var gotParams map[string]string

	client := NewClient(&mockClient{
		baseURL: "https://example.com",
		do: func(urlStr string, method string, data map[string]interface{}, params map[string]string) (*http.Response, error) {
			gotParams = params
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(`{"code":"SUCCESS","data":{"dataList":[]}}`)),
				Header:     make(http.Header),
			}, nil
		},
	})

	_, err := client.GetFundingRate(context.Background(), GetFundingRateParams{ContractID: "10000001"})
	if err != nil {
		t.Fatalf("GetFundingRate returned error: %v", err)
	}
	if _, ok := gotParams["filterSettlementFundingRate"]; ok {
		t.Fatalf("unexpected settlement filter in query: %+v", gotParams)
	}
}

func TestGetFundingRateIncludesSettlementFilterWhenSpecified(t *testing.T) {
	var gotParams map[string]string
	settlementOnly := true

	client := NewClient(&mockClient{
		baseURL: "https://example.com",
		do: func(urlStr string, method string, data map[string]interface{}, params map[string]string) (*http.Response, error) {
			gotParams = params
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(`{"code":"SUCCESS","data":{"dataList":[]}}`)),
				Header:     make(http.Header),
			}, nil
		},
	})

	_, err := client.GetFundingRate(context.Background(), GetFundingRateParams{
		ContractID:                  "10000001",
		FilterSettlementFundingRate: &settlementOnly,
	})
	if err != nil {
		t.Fatalf("GetFundingRate returned error: %v", err)
	}
	if gotParams["filterSettlementFundingRate"] != "true" {
		t.Fatalf("unexpected settlement filter query: %+v", gotParams)
	}
}
