package user

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/edgex-Tech/edgex-golang-sdk/sdk/internal"
)

// Client represents the user/onboarding API client.
type Client struct {
	*internal.Client
}

// NewClient creates a new user client.
func NewClient(client *internal.Client) *Client {
	return &Client{Client: client}
}

// CheckUserExist checks whether a wallet user already exists.
func (c *Client) CheckUserExist(ctx context.Context, ethAddress string) (*CheckUserExistResponse, error) {
	_ = ctx
	url := fmt.Sprintf("%s/api/v1/public/user/checkUserExist", c.Client.GetBaseURL())
	params := map[string]string{
		"ethAddress": ethAddress,
	}

	resp, err := c.Client.HttpRequest(url, "GET", nil, params)
	if err != nil {
		return nil, fmt.Errorf("failed to check user exist: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result CheckUserExistResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s, errorParam: %v", result.Code, result.ErrorParam)
	}

	return &result, nil
}

// GetAccountIdForRegister returns server-generated hint account id for v2 registration.
func (c *Client) GetAccountIdForRegister(ctx context.Context, ethAddress string, clientAccountId string) (*GetAccountIdForRegisterResponse, error) {
	_ = ctx
	url := fmt.Sprintf("%s/api/v1/public/account/getAccountIdForRegister", c.Client.GetBaseURL())
	params := map[string]string{
		"ethAddress":      ethAddress,
		"clientAccountId": clientAccountId,
	}

	resp, err := c.Client.HttpRequest(url, "GET", nil, params)
	if err != nil {
		return nil, fmt.Errorf("failed to get account id for register: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result GetAccountIdForRegisterResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s, errorParam: %v", result.Code, result.ErrorParam)
	}

	return &result, nil
}

// OnboardSiteV2 performs v2 onboarding/login with optional registration proof fields.
func (c *Client) OnboardSiteV2(ctx context.Context, params *OnboardSiteV2Params) (*OnboardSiteResponse, error) {
	_ = ctx
	if params == nil {
		return nil, fmt.Errorf("params is required")
	}
	if params.EthAddress == "" {
		return nil, fmt.Errorf("ethAddress is required")
	}
	if params.Signature == "" {
		return nil, fmt.Errorf("signature is required")
	}

	payload := map[string]interface{}{
		"ethAddress": params.EthAddress,
		"onlySignOn": params.OnlySignOn,
		"signature":  params.Signature,
	}
	if params.Param != "" {
		payload["param"] = params.Param
	}
	if params.ClientAccountId != "" {
		payload["clientAccountId"] = params.ClientAccountId
	}
	if params.PrivyIdentityToken != "" {
		payload["privyIdentityToken"] = params.PrivyIdentityToken
	}
	if params.EthSignature != "" {
		payload["ethSignature"] = params.EthSignature
	}
	if params.HintAccountId != "" {
		payload["hintAccountId"] = params.HintAccountId
	}
	if len(params.ExtraSigners) > 0 {
		payload["extraSigners"] = normalizeAndSortAddresses(params.ExtraSigners)
	}
	if len(params.SignerWithPermissions) > 0 {
		payload["signerWithPermissions"] = normalizeAndSortSignerPermissions(params.SignerWithPermissions)
	}

	url := fmt.Sprintf("%s/api/v1/public/user/onboardSite", c.Client.GetBaseURL())
	resp, err := c.Client.HttpRequest(url, "POST", payload, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to onboard site: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result OnboardSiteResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	if result.Code != "SUCCESS" {
		return nil, fmt.Errorf("request failed with code: %s, errorParam: %v", result.Code, result.ErrorParam)
	}

	return &result, nil
}

func normalizeAndSortAddresses(input []string) []string {
	out := make([]string, 0, len(input))
	seen := make(map[string]struct{}, len(input))
	for _, addr := range input {
		addr = strings.TrimSpace(addr)
		if addr == "" {
			continue
		}
		key := strings.ToLower(addr)
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		out = append(out, addr)
	}
	sort.Slice(out, func(i, j int) bool {
		return strings.ToLower(out[i]) < strings.ToLower(out[j])
	})
	return out
}

func normalizeAndSortSignerPermissions(input []SignerWithPermissions) []SignerWithPermissions {
	out := make([]SignerWithPermissions, 0, len(input))
	seen := make(map[string]struct{}, len(input))
	for _, item := range input {
		addr := strings.TrimSpace(item.Signer)
		if addr == "" {
			continue
		}
		key := strings.ToLower(addr) + ":" + strings.TrimSpace(item.Permissions)
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		out = append(out, SignerWithPermissions{
			Signer:      addr,
			Permissions: strings.TrimSpace(item.Permissions),
		})
	}
	sort.Slice(out, func(i, j int) bool {
		return strings.ToLower(out[i].Signer) < strings.ToLower(out[j].Signer)
	})
	return out
}
