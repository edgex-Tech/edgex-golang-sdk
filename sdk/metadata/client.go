package metadata

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/edgex-Tech/edgex-golang-sdk/sdk/internal"
)

// Client represents the metadata client
type Client struct {
	*internal.Client
}

// NewClient creates a new metadata client
func NewClient(client *internal.Client) *Client {
	return &Client{
		Client: client,
	}
}

// GetServerTime gets the current server time
func (c *Client) GetServerTime(ctx context.Context) (*ResultGetServerTime, error) {
	url := fmt.Sprintf("%s/api/v1/public/meta/getServerTime", c.Client.GetBaseURL())

	resp, err := c.Client.HttpRequest(url, "GET", nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get server time: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result ResultGetServerTime
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}

// GetMetaData gets the exchange metadata
func (c *Client) GetMetaData(ctx context.Context) (*ResultMetaData, error) {
	url := fmt.Sprintf("%s/api/v1/public/meta/getMetaData", c.Client.GetBaseURL())

	resp, err := c.Client.HttpRequest(url, "GET", nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get metadata: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result ResultMetaData
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s", result.Code)
	}

	return &result, nil
}
