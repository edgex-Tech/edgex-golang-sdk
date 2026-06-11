package transfer

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/edgex-Tech/edgex-golang-sdk/v2/sdk/internal"
	metadatapkg "github.com/edgex-Tech/edgex-golang-sdk/v2/sdk/metadata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateTransferOutV2UsesEIP712Signature(t *testing.T) {
	var gotPath string
	var gotBody map[string]interface{}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		defer r.Body.Close()
		_ = json.NewDecoder(r.Body).Decode(&gotBody)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code":"SUCCESS","data":{"id":"t-1"}}`))
	}))
	defer srv.Close()

	mockCli := newMockClient(42, "0x59c6995e998f97a5a0044966f0945380f7d7f4fbcbe8f85f8f19853f51e7f7b4", srv.URL, func(urlStr, method string, data map[string]interface{}, params map[string]string) (*http.Response, error) {
		payload, _ := json.Marshal(data)
		req, _ := http.NewRequest(method, urlStr, bytes.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")
		return srv.Client().Do(req)
	})

	client := NewClient(mockCli)
	expire := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli()
	params := &CreateTransferOutParams{
		ClientTransferId:  "transfer-v2-fixed-id",
		CoinId:            "2",
		Amount:            "1.234567",
		ReceiverAccountId: "99",
		TransferReason:    USER_TRANSFER.String(),
		L2ExpireTime:      strconv.FormatInt(expire, 10),
	}

	res, err := client.CreateTransferOut(context.Background(), params, &metadatapkg.MetaData{
		Global: &metadatapkg.Global{
			ChainId:         "42161",
			ContractAddress: "0x0000000000000000000000000000000000000001",
		},
		CoinList: []metadatapkg.Coin{
			{CoinId: "2", StepSize: "0.000001"},
		},
	})
	require.NoError(t, err)
	require.NotNil(t, res)
	assert.Equal(t, "/api/v2/private/transfer/createTransferOut", gotPath)

	assert.Equal(t, "transfer-v2-fixed-id", gotBody["clientTransferId"])
	assert.Equal(t, strconv.FormatInt(internal.CalcNonce("transfer-v2-fixed-id"), 10), gotBody["l2Nonce"])
	assert.Equal(t, strconv.FormatInt(expire, 10), gotBody["l2ExpireTime"])
	assert.Equal(t, "42", gotBody["accountId"])
	assert.Equal(t, "99", gotBody["receiverAccountId"])
	assert.NotContains(t, gotBody, "receiverL2Key")

	l2Signature, ok := gotBody["l2Signature"].(string)
	require.True(t, ok)
	assert.True(t, strings.HasPrefix(l2Signature, "0x"))
	assert.Len(t, l2Signature, 132)
}

func TestCreateTransferOutV2RequiresWalletKeyWhenSignatureMissing(t *testing.T) {
	mockCli := newMockClient(42, "", "http://127.0.0.1", nil)

	client := NewClient(mockCli)
	_, err := client.CreateTransferOut(context.Background(), &CreateTransferOutParams{
		ClientTransferId:  "transfer-v2-fixed-id",
		CoinId:            "2",
		Amount:            "1.234567",
		ReceiverAccountId: "99",
		TransferReason:    USER_TRANSFER.String(),
		L2ExpireTime:      strconv.FormatInt(time.Now().Add(24*time.Hour).UnixMilli(), 10),
	}, &metadatapkg.MetaData{
		Global: &metadatapkg.Global{
			ChainId:         "42161",
			ContractAddress: "0x0000000000000000000000000000000000000001",
		},
		CoinList: []metadatapkg.Coin{
			{CoinId: "2", StepSize: "0.000001"},
		},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "wallet")
}

func TestCreateTransferOutV2AllowsProvidedSignatureWithoutWalletKey(t *testing.T) {
	var gotBody map[string]interface{}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		_ = json.NewDecoder(r.Body).Decode(&gotBody)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code":"SUCCESS","data":{"id":"t-1"}}`))
	}))
	defer srv.Close()

	mockCli := newMockClient(42, "", srv.URL, func(urlStr, method string, data map[string]interface{}, params map[string]string) (*http.Response, error) {
		payload, _ := json.Marshal(data)
		req, _ := http.NewRequest(method, urlStr, bytes.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")
		return srv.Client().Do(req)
	})

	client := NewClient(mockCli)
	expire := strconv.FormatInt(time.Date(2026, 1, 2, 0, 0, 0, 0, time.UTC).UnixMilli(), 10)
	res, err := client.CreateTransferOut(context.Background(), &CreateTransferOutParams{
		ClientTransferId:  "transfer-v2-fixed-id",
		CoinId:            "2",
		Amount:            "1.234567",
		ReceiverAccountId: "99",
		TransferReason:    USER_TRANSFER.String(),
		L2Nonce:           "888",
		L2ExpireTime:      expire,
		L2Signature:       "0xdeadbeef",
	}, &metadatapkg.MetaData{
		Global: &metadatapkg.Global{
			ChainId:         "42161",
			ContractAddress: "0x0000000000000000000000000000000000000001",
		},
		CoinList: []metadatapkg.Coin{
			{CoinId: "2", StepSize: "0.000001"},
		},
	})
	require.NoError(t, err)
	require.NotNil(t, res)
	assert.Equal(t, "0xdeadbeef", gotBody["l2Signature"])
	assert.Equal(t, "888", gotBody["l2Nonce"])
	assert.Equal(t, expire, gotBody["l2ExpireTime"])
}
