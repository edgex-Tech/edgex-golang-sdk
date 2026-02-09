package account

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/edgex-Tech/edgex-golang-sdk/sdk/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegisterAccountV2SortsSigners(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v2/private/account/registerAccount", r.URL.Path)

		var payload map[string]interface{}
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))

		extraSignersAny, ok := payload["extraSigners"].([]interface{})
		require.True(t, ok)
		require.Len(t, extraSignersAny, 2)
		assert.Equal(t, "0xAaaa", extraSignersAny[0])
		assert.Equal(t, "0xBbbb", extraSignersAny[1])

		signersAny, ok := payload["signerWithPermissions"].([]interface{})
		require.True(t, ok)
		require.Len(t, signersAny, 2)

		first := signersAny[0].(map[string]interface{})
		second := signersAny[1].(map[string]interface{})
		assert.Equal(t, "0x1111", first["signer"])
		assert.Equal(t, "0x2222", second["signer"])

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code":"SUCCESS","data":{"accountId":"123","createdTime":"1","isNewAccount":true}}`))
	}))
	defer server.Close()

	internalClient, err := internal.NewClient(&internal.ClientConfig{
		BaseURL:       server.URL,
		APIVersion:    internal.APIVersionV2,
		SigningMethod: internal.SigningMethodHMAC,
		APIKey:        "api-key",
		APIPassphrase: "passphrase",
		APISecret:     "secret",
		AuthHeaderKey: "edgeX",
	})
	require.NoError(t, err)

	client := NewClient(internalClient)
	res, err := client.RegisterAccountV2(context.Background(), &RegisterAccountV2Params{
		AccountName:     "sub",
		IsSystemAccount: false,
		ExtraSigners:    []string{"0xBbbb", "0xAaaa"},
		SignerWithPermissions: []SignerWithPermissions{
			{Signer: "0x2222", Permissions: "0x20"},
			{Signer: "0x1111", Permissions: "0xffff"},
		},
		EthSignature:  "0xabc",
		HintAccountId: "123",
	})
	require.NoError(t, err)
	require.NotNil(t, res)
	assert.Equal(t, "SUCCESS", res.Code)
	require.NotNil(t, res.Data)
	assert.Equal(t, "123", res.Data.AccountID)
	assert.True(t, res.Data.IsNewAccount)
}

func TestRegisterAccountV2AutoSignsWhenEthSignatureEmpty(t *testing.T) {
	var registerPayload map[string]interface{}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"code":"SUCCESS","data":{"global":{"nativeChainId":"42161","contractAddress":"0x0000000000000000000000000000000000000001"}}}`))
		case strings.Contains(r.URL.Path, "/public/info"):
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"code":"SUCCESS","data":{"global":{"nativeChainId":"42161","contractAddress":"0x0000000000000000000000000000000000000001"}}}`))
		case strings.Contains(r.URL.Path, "/private/account/registerAccount"):
			defer r.Body.Close()
			require.NoError(t, json.NewDecoder(r.Body).Decode(&registerPayload))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"code":"SUCCESS","data":{"accountId":"123","createdTime":"1","isNewAccount":true}}`))
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	internalClient, err := internal.NewClient(&internal.ClientConfig{
		BaseURL:       server.URL,
		APIVersion:    internal.APIVersionV2,
		SigningMethod: internal.SigningMethodHMAC,
		APIKey:        "api-key",
		APIPassphrase: "passphrase",
		APISecret:     "secret",
		AuthHeaderKey: "edgeX",
		WalletPriKey:  "0x59c6995e998f97a5a0044966f0945380f7d7f4fbcbe8f85f8f19853f51e7f7b4",
	})
	require.NoError(t, err)

	client := NewClient(internalClient)
	res, err := client.RegisterAccountV2(context.Background(), &RegisterAccountV2Params{
		AccountName:     "sub",
		IsSystemAccount: false,
		ExtraSigners:    []string{"0xBbbb", "0xAaaa"},
		SignerWithPermissions: []SignerWithPermissions{
			{Signer: "0x2222222222222222222222222222222222222222", Permissions: "0x20"},
			{Signer: "0x1111111111111111111111111111111111111111", Permissions: "0xffff"},
		},
		HintAccountId: "123",
	})
	require.NoError(t, err)
	require.NotNil(t, res)
	require.NotNil(t, registerPayload)

	sig, ok := registerPayload["ethSignature"].(string)
	require.True(t, ok)
	assert.True(t, strings.HasPrefix(sig, "0x"))
	assert.Len(t, sig, 132)
}
