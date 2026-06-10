package cctp

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainnetDefaultsMatchEdgeToEthereum(t *testing.T) {
	assert.Equal(t, "https://iris-api.circle.com", DefaultIrisBaseURL)
	assert.Equal(t, 28, DefaultSourceCCTPDomain)
	assert.Equal(t, 0, DefaultDestCCTPDomain)
	assert.Equal(t, "0x81D40F21F12A8F0E3252Bccb954D722d4c464B64", DefaultETHMessageTransmitter)
}

func TestBytes32AddressLeftPadsEVMAddress(t *testing.T) {
	value, err := Bytes32Address("0xFCAd0B19bB29D4674531d6f115237E16AfCE377c")
	require.NoError(t, err)
	assert.Equal(t, "0x000000000000000000000000fcad0b19bb29d4674531d6f115237e16afce377c", value)
	assert.Len(t, ZeroBytes32, 66)
}

func TestExtractMinimumFeeBpsUsesMatchingFinalityThreshold(t *testing.T) {
	payload := []map[string]interface{}{
		{"finalityThreshold": 1000, "minimumFee": 1.5},
		{"finalityThreshold": 2000, "minimumFee": 0},
	}
	value, err := ExtractMinimumFeeBps(payload, 1000)
	require.NoError(t, err)
	assert.Equal(t, "1.5", value.String())

	value, err = ExtractMinimumFeeBps(payload, 2000)
	require.NoError(t, err)
	assert.Equal(t, "0", value.String())
}

func TestSelectIrisMessageByEventNonce(t *testing.T) {
	payload := map[string]interface{}{
		"messages": []map[string]interface{}{
			{"status": "pending", "eventNonce": "0x1", "attestation": ""},
			{"status": "complete", "eventNonce": "0x2", "attestation": "0xab"},
		},
	}

	message, err := SelectIrisMessage(payload, nil, "0x2")
	require.NoError(t, err)
	assert.Equal(t, "0x2", message["eventNonce"])
	assert.True(t, IsCompleteIrisMessage(map[string]interface{}{
		"status":      "complete",
		"eventNonce":  "0x2",
		"attestation": "0xab",
	}))
}
