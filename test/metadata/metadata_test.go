package metadata

import (
	"testing"

	"github.com/edgex-Tech/edgex-golang-sdk/test"
	"github.com/stretchr/testify/assert"
)

func TestGetMetadata(t *testing.T) {
	client, err := test.CreateTestClient()
	assert.NoError(t, err)

	ctx := test.GetTestContext()
	result, err := client.GetMetaData(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "SUCCESS", result.Code)

	data := result.Data
	assert.NotNil(t, data)

	// Test some fields to ensure we got valid data
	assert.NotEmpty(t, data.CoinList)
	assert.NotEmpty(t, data.ContractList)
	assert.NotNil(t, data.Global)
}
