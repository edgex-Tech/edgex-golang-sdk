package sdk

import (
	"testing"
	"time"

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

func TestNewClient_UsesAssetBaseURL(t *testing.T) {
	client, err := NewClient(&ClientConfig{
		BaseURL:      "https://edgex-prod-v2.edgex.exchange/api/",
		AssetBaseURL: "https://spot.edgex.exchange/api/",
		AccountID:    1,
	})
	require.NoError(t, err)

	assert.Equal(t, "https://edgex-prod-v2.edgex.exchange", client.GetBaseURL())
	assert.Equal(t, "https://spot.edgex.exchange", client.GetAssetBaseURL())
}
