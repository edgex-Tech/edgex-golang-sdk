package quote

import (
	"encoding/json"
	"testing"

	"github.com/edgex-Tech/edgex-golang-sdk/sdk/quote"
	"github.com/edgex-Tech/edgex-golang-sdk/test"
	"github.com/stretchr/testify/assert"
)

func TestGetQuoteSummary(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()
	resp, err := client.GetQuoteSummary(ctx, "LAST_DAY_1")
	jsonData, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("Quote Summary: %s", string(jsonData))
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "SUCCESS", resp.Code)

	data := resp.Data
	assert.NotNil(t, data)
	assert.NotNil(t, data.TickerSummary)
	t.Logf("Ticker summary data: %v", data.TickerSummary)
}

func TestGet24HourQuotes(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()
	contractID, err := test.ResolveTestContractID(ctx, client)
	assert.NoError(t, err)

	resp, err := client.Get24HourQuote(ctx, contractID)
	jsonData, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("24-Hour Quotes: %s", string(jsonData))
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "SUCCESS", resp.Code)

	data := resp.Data
	assert.NotNil(t, data)
	// Data is []interface{}, skip detailed assertions
	t.Logf("24-hour quotes data: %v", data)
}

func TestGetKLine(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()
	contractID, err := test.ResolveTestContractID(ctx, client)
	assert.NoError(t, err)

	params := quote.GetKLineParams{
		ContractID: contractID,
		Interval:   quote.KlineType1Hour,
		Size:       1,
		PriceType:  quote.PriceTypeLastPrice,
	}
	paramsJSON, _ := json.MarshalIndent(params, "", "  ")
	t.Logf("GetKLine params: %s", string(paramsJSON))
	resp, err := client.GetKLine(ctx, params)
	jsonData, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("GetKLine response: %s", string(jsonData))
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "SUCCESS", resp.Code)

	data := resp.Data
	assert.NotNil(t, data)
	assert.NotNil(t, data.DataList)
	// Verify Kline data structure
	for _, kline := range data.DataList {
		assert.NotNil(t, kline)
		assert.NotEmpty(t, kline.ContractId)
	}
}

func TestGetOrderBookDepth(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()
	contractID, err := test.ResolveTestContractID(ctx, client)
	assert.NoError(t, err)

	params := quote.GetOrderBookDepthParams{
		ContractID: contractID,
		Size:       15, // API supports 15 or 200 levels
	}
	resp, err := client.GetOrderBookDepth(ctx, params)
	jsonData, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("Order Book Depth: %s", string(jsonData))
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "SUCCESS", resp.Code)

	data := resp.Data
	assert.NotNil(t, data)
	// Verify Depth data structure
	for _, depth := range data {
		assert.NotNil(t, depth)
		assert.NotEmpty(t, depth.ContractId)
		assert.NotNil(t, depth.Asks)
		assert.NotNil(t, depth.Bids)
	}
}

func TestGetMultiContractKLine(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()
	contractID, err := test.ResolveTestContractID(ctx, client)
	assert.NoError(t, err)

	params := quote.GetMultiContractKLineParams{
		ContractIDs: []string{contractID},
		Interval:    quote.KlineType1Hour,
		Size:        1,
		PriceType:   quote.PriceTypeLastPrice,
	}
	paramsJSON, _ := json.MarshalIndent(params, "", "  ")
	t.Logf("GetMultiContractKLine params: %s", string(paramsJSON))
	resp, err := client.GetMultiContractKLine(ctx, params)
	jsonData, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("GetMultiContractKLine response: %s", string(jsonData))
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "SUCCESS", resp.Code)

	data := resp.Data
	assert.NotNil(t, data)
	// Verify ContractMultiKline data structure
	for _, contractKline := range data {
		assert.NotNil(t, contractKline)
		assert.NotEmpty(t, contractKline.ContractId)
	}
}
