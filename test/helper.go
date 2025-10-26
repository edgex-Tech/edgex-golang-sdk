package test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/edgex-Tech/edgex-golang-sdk/sdk"
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

	starkPrivateKey := os.Getenv("TEST_STARK_PRIVATE_KEY")
	if starkPrivateKey == "" {
		//return nil, fmt.Errorf("TEST_STARK_PRIVATE_KEY environment variable is not set")
		starkPrivateKey = "04a266bc1e005725a278034bc4ab0f3075a7110a47d390b0b1b7841cabac0c4d"
	}

	return sdk.NewClient(&sdk.ClientConfig{
		BaseURL:          baseURL,
		AccountID:        accountID,
		StarkPriKey:      starkPrivateKey,
		MetaDataCacheTTL: nil,
	})
}

// GetTestContext returns a context for testing
func GetTestContext() context.Context {
	return context.Background()
}
