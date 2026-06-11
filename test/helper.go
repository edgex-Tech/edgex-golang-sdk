package test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/edgex-Tech/edgex-golang-sdk/v2/sdk"
	"github.com/edgex-Tech/edgex-golang-sdk/v2/sdk/metadata"
	"github.com/joho/godotenv"
)

func init() {
	// Get the current file's directory
	_, filename, _, _ := runtime.Caller(0)
	// Go up two directories to reach the project root
	projectRoot := filepath.Dir(filepath.Dir(filename))
	envPath := filepath.Join(projectRoot, ".env")

	// Load .env file if it exists
	_ = godotenv.Load(envPath)
}

// CreateTestClient creates a new SDK client for testing
func CreateTestClient() (*sdk.Client, error) {
	baseURL := os.Getenv("EDGEX_BASE_URL")
	if baseURL == "" {
		baseURL = "https://testnet.edgex.exchange"
	}

	accountIDStr := os.Getenv("EDGEX_ACCOUNT_ID")
	if accountIDStr == "" {
		accountIDStr = "665403845421039873"
	}

	accountID, err := strconv.ParseInt(accountIDStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid EDGEX_ACCOUNT_ID: %w", err)
	}

	return sdk.NewClient(&sdk.ClientConfig{
		BaseURL:          baseURL,
		AssetBaseURL:     strings.TrimSpace(os.Getenv("EDGEX_ASSET_BASE_URL")),
		AccountID:        accountID,
		APIKey:           strings.TrimSpace(os.Getenv("EDGEX_API_KEY")),
		APIPassphrase:    strings.TrimSpace(os.Getenv("EDGEX_API_PASSPHRASE")),
		APISecret:        strings.TrimSpace(os.Getenv("EDGEX_API_SECRET")),
		SignerPriKey:     strings.TrimSpace(os.Getenv("EDGEX_SIGNER_PRIVATE_KEY")),
		SignerAddr:       strings.TrimSpace(os.Getenv("EDGEX_SIGNER_ADDRESS")),
		WalletPriKey:     strings.TrimSpace(os.Getenv("EDGEX_WALLET_PRIVATE_KEY")),
		WalletAddr:       strings.TrimSpace(os.Getenv("EDGEX_WALLET_ADDRESS")),
		AuthHeaderKey:    strings.TrimSpace(os.Getenv("EDGEX_AUTH_HEADER_KEY")),
		MetaDataCacheTTL: nil,
	})
}

// GetTestContext returns a context for testing
func GetTestContext() context.Context {
	return context.Background()
}

// ResolveTestContract returns a valid contract from metadata.
func ResolveTestContract(ctx context.Context, client *sdk.Client) (*metadata.Contract, error) {
	resp, err := client.GetMetaData(ctx)
	if err != nil {
		return nil, err
	}
	if resp == nil || resp.Data == nil {
		return nil, fmt.Errorf("metadata response is nil")
	}

	for i := range resp.Data.ContractList {
		contract := &resp.Data.ContractList[i]
		if strings.TrimSpace(contract.ContractId) != "" {
			return contract, nil
		}
	}
	return nil, fmt.Errorf("no valid contract id in metadata")
}

// ResolveTestContractID returns a valid contract id from metadata.
func ResolveTestContractID(ctx context.Context, client *sdk.Client) (string, error) {
	contract, err := ResolveTestContract(ctx, client)
	if err != nil {
		return "", err
	}
	return contract.ContractId, nil
}

// ResolveTestCoinID returns a valid collateral/transfer coin id from metadata.
func ResolveTestCoinID(ctx context.Context, client *sdk.Client) (string, error) {
	resp, err := client.GetMetaData(ctx)
	if err != nil {
		return "", err
	}
	if resp == nil || resp.Data == nil {
		return "", fmt.Errorf("metadata response is nil")
	}
	if resp.Data.Global != nil && strings.TrimSpace(resp.Data.Global.TransferCoinId) != "" {
		return strings.TrimSpace(resp.Data.Global.TransferCoinId), nil
	}
	if len(resp.Data.CoinList) > 0 {
		coinID := strings.TrimSpace(resp.Data.CoinList[0].CoinId)
		if coinID != "" {
			return coinID, nil
		}
	}
	return "", fmt.Errorf("no valid coin id in metadata")
}
