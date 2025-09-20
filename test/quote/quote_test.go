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

	resp, err := client.GetQuoteSummary(ctx, "20000018")
	jsonData, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("Quote Summary: %s", string(jsonData))
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "SUCCESS", resp.GetCode())

	data := resp.GetData()
	assert.NotNil(t, data)
	assert.NotNil(t, data.GetTickerSummary())
}

func TestGet24HourQuotes(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	resp, err := client.Get24HourQuote(ctx, "20000018")
	jsonData, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("24-Hour Quotes: %s", string(jsonData))
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "SUCCESS", resp.GetCode())

	data := resp.GetData()
	assert.NotNil(t, data)
	for _, ticker := range data {
		assert.NotEmpty(t, ticker.GetContractId())
		assert.NotEmpty(t, ticker.GetLastPrice())
	}
}

func TestGetKLine(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	params := quote.GetKLineParams{
		ContractID: "20000018",
		Interval:   "HOUR_1",
		Size:       100,
		PriceType:  "LAST_PRICE",
	}
	resp, err := client.GetKLine(ctx, params)
	jsonData, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("K-Line Data: %s", string(jsonData))
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "SUCCESS", resp.GetCode())

	data := resp.GetData()
	assert.NotNil(t, data)
	assert.NotEmpty(t, data.GetDataList())
}

func TestGetOrderBookDepth(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	params := quote.GetOrderBookDepthParams{
		ContractID: "20000018",
		Size:       15, // API supports 15 or 200 levels
	}
	resp, err := client.GetOrderBookDepth(ctx, params)
	jsonData, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("Order Book Depth: %s", string(jsonData))
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "SUCCESS", resp.GetCode())

	data := resp.GetData()
	assert.NotNil(t, data)
	for _, depth := range data {
		assert.NotNil(t, depth.GetAsks())
		assert.NotNil(t, depth.GetBids())
	}
}

func TestGetMultiContractKLine(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	params := quote.GetMultiContractKLineParams{
		ContractIDs: []string{"20000018"},
		Interval:    "HOUR_1",
		Size:        100,
		PriceType:   "LAST_PRICE",
	}
	resp, err := client.GetMultiContractKLine(ctx, params)
	jsonData, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("Multi-Contract K-Line Data: %s", string(jsonData))
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "SUCCESS", resp.GetCode())

	data := resp.GetData()
	assert.NotNil(t, data)
	for _, kline := range data {
		assert.NotEmpty(t, kline.GetContractId())
		assert.NotEmpty(t, kline.GetKlineList())
	}
}
