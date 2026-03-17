package internal

import (
	"encoding/hex"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeriveAddressFromPrivateKey(t *testing.T) {
	address, err := DeriveAddressFromPrivateKey("0x59c6995e998f97a5a0044966f0945380f7d7f4fbcbe8f85f8f19853f51e7f7b4")
	require.NoError(t, err)
	assert.True(t, strings.HasPrefix(address, "0x"))
	assert.Len(t, address, 42)
}

func TestSignTypedDataWithPrivateKey(t *testing.T) {
	privateKey := "0x59c6995e998f97a5a0044966f0945380f7d7f4fbcbe8f85f8f19853f51e7f7b4"
	domain, err := NewTypedDataDomain("EdgeX", "1", "1", "0x0000000000000000000000000000000000000001")
	require.NoError(t, err)

	typedData := TypedData{
		Types: TypedDataTypes{
			"EIP712Domain": []TypedDataType{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
			"TransferParams": []TypedDataType{
				{Name: "from", Type: "uint64"},
				{Name: "to", Type: "uint64"},
				{Name: "assetId", Type: "uint64"},
				{Name: "amount", Type: "int256"},
				{Name: "nonce", Type: "uint256"},
				{Name: "deadline", Type: "uint256"},
				{Name: "signer", Type: "address"},
			},
		},
		PrimaryType: "TransferParams",
		Domain:      domain,
		Message: TypedDataMessage{
			"from":     "1",
			"to":       "2",
			"assetId":  "3",
			"amount":   "1000000",
			"nonce":    "999",
			"deadline": "123456789",
			"signer":   "0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC",
		},
	}

	signatureHex, err := SignTypedDataWithPrivateKey(privateKey, typedData)
	require.NoError(t, err)
	assert.True(t, strings.HasPrefix(signatureHex, "0x"))
	assert.Len(t, signatureHex, 132)

	digest, _, err := apitypes.TypedDataAndHash(typedData)
	require.NoError(t, err)

	signatureRaw, err := hex.DecodeString(strings.TrimPrefix(signatureHex, "0x"))
	require.NoError(t, err)
	require.Len(t, signatureRaw, 65)
	assert.Contains(t, []byte{27, 28}, signatureRaw[64])

	recoverySig := make([]byte, len(signatureRaw))
	copy(recoverySig, signatureRaw)
	recoverySig[64] -= 27

	pubkey, err := crypto.SigToPub(digest, recoverySig)
	require.NoError(t, err)
	recovered := crypto.PubkeyToAddress(*pubkey).Hex()

	expected, err := DeriveAddressFromPrivateKey(privateKey)
	require.NoError(t, err)
	assert.Equal(t, expected, recovered)
}
