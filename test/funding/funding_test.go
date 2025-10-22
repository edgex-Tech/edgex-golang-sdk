package funding

import (
	"encoding/json"
	"testing"

	"github.com/edgex-Tech/edgex-golang-sdk/sdk/funding"
	"github.com/edgex-Tech/edgex-golang-sdk/test"
	"github.com/stretchr/testify/assert"
)

func TestGetFundingRate(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	size := int32(10)
	params := funding.GetFundingRateParams{
		ContractID: "20000018", // liquid contract with active market making
		Size:       &size,
	}
	resp, err := client.Funding.GetFundingRate(ctx, params)
	jsonData, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("Funding Rate: %s", string(jsonData))
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "SUCCESS", resp.Code)

	data := resp.Data
	assert.NotNil(t, data)
	assert.NotEmpty(t, data.DataList)
	for _, rate := range data.DataList {
		// Verify FundingRate data structure
		assert.NotNil(t, rate)
		assert.NotEmpty(t, rate.ContractId)
		assert.NotEmpty(t, rate.FundingRate)
	}
}

func TestGetLatestFundingRate(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()

	params := funding.GetLatestFundingRateParams{
		ContractID: "20000018", // liquid contract with active market making
	}
	resp, err := client.Funding.GetLatestFundingRate(ctx, params)
	jsonData, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("Latest Funding Rate: %s", string(jsonData))
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "SUCCESS", resp.Code)

	data := resp.Data
	assert.NotNil(t, data)
	assert.NotEmpty(t, data)
	for _, rate := range data {
		// Verify FundingRate data structure
		assert.NotNil(t, rate)
		assert.NotEmpty(t, rate.ContractId)
		assert.NotEmpty(t, rate.FundingRate)
	}
}
