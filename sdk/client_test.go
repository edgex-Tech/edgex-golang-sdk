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
