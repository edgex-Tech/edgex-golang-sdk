package test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/edgex-Tech/edgex-golang-sdk/sdk"
	"github.com/edgex-Tech/edgex-golang-sdk/sdk/metadata"
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
	baseURL := os.Getenv("TEST_BASE_URL")
	if baseURL == "" {
		//return nil, fmt.Errorf("TEST_BASE_URL environment variable is not set")
		baseURL = "https://testnet.edgex.exchange"
	}

	accountIDStr := os.Getenv("TEST_ACCOUNT_ID")
	if accountIDStr == "" {
		// return nil, fmt.Errorf("TEST_ACCOUNT_ID environment variable is not set")
		accountIDStr = "665403845421039873"
	}

	accountID, err := strconv.ParseInt(accountIDStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid TEST_ACCOUNT_ID: %w", err)
	}

	apiVersion := strings.ToLower(strings.TrimSpace(os.Getenv("TEST_API_VERSION")))
	if apiVersion == "" {
		baseURLLower := strings.ToLower(baseURL)
		if strings.Contains(baseURLLower, "-v2.") ||
			strings.Contains(baseURLLower, "/v2") ||
			strings.TrimSpace(os.Getenv("TEST_API_KEY")) != "" ||
			strings.TrimSpace(os.Getenv("TEST_API_PASSPHRASE")) != "" ||
			strings.TrimSpace(os.Getenv("TEST_API_SECRET")) != "" {
			apiVersion = sdk.APIVersionV2
		} else {
			apiVersion = sdk.APIVersionV1
		}
	}

	signingMethod := strings.ToLower(strings.TrimSpace(os.Getenv("TEST_SIGNING_METHOD")))
	if signingMethod == "" {
		signingMethod = strings.ToLower(strings.TrimSpace(os.Getenv("TEST_WS_SIGNING_METHOD")))
	}
	if signingMethod == "" {
		if apiVersion == sdk.APIVersionV2 {
			signingMethod = sdk.SigningMethodHMAC
		} else {
			signingMethod = sdk.SigningMethodStark
		}
	}

	starkPrivateKey := strings.TrimSpace(os.Getenv("TEST_STARK_PRIVATE_KEY"))
	if apiVersion == sdk.APIVersionV1 && starkPrivateKey == "" {
		// Keep v1 fallback for backwards compatibility.
		starkPrivateKey = "04a266bc1e005725a278034bc4ab0f3075a7110a47d390b0b1b7841cabac0c4d"
	}

	return sdk.NewClient(&sdk.ClientConfig{
		BaseURL:          baseURL,
		AccountID:        accountID,
		StarkPriKey:      starkPrivateKey,
		TradingPriKey:    strings.TrimSpace(os.Getenv("TEST_TRADING_PRIVATE_KEY")),
		WalletPriKey:     strings.TrimSpace(os.Getenv("TEST_WALLET_PRIVATE_KEY")),
		TradingAddr:      strings.TrimSpace(os.Getenv("TEST_TRADING_ADDRESS")),
		WalletAddr:       strings.TrimSpace(os.Getenv("TEST_WALLET_ADDRESS")),
		APIVersion:       apiVersion,
		SigningMethod:    signingMethod,
		APIKey:           strings.TrimSpace(os.Getenv("TEST_API_KEY")),
		APIPassphrase:    strings.TrimSpace(os.Getenv("TEST_API_PASSPHRASE")),
		APISecret:        strings.TrimSpace(os.Getenv("TEST_API_SECRET")),
		AuthHeaderKey:    strings.TrimSpace(os.Getenv("TEST_AUTH_HEADER_KEY")),
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
