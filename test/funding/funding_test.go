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
	assert.Equal(t, "SUCCESS", resp.GetCode())

	data := resp.GetData()
	assert.NotNil(t, data)
	assert.NotEmpty(t, data.GetDataList())
	for _, rate := range data.GetDataList() {
		assert.NotEmpty(t, rate.GetContractId())
		assert.NotEmpty(t, rate.GetFundingRate())
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
	assert.Equal(t, "SUCCESS", resp.GetCode())

	data := resp.GetData()
	assert.NotNil(t, data)
	assert.NotEmpty(t, data)
	for _, rate := range data {
		assert.NotEmpty(t, rate.GetContractId())
		assert.NotEmpty(t, rate.GetFundingRate())
	}
}
