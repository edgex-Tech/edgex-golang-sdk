package user

// CheckUserExist represents user existence response data.
type CheckUserExist struct {
	IsUserExist bool `json:"isUserExist"`
}

// CheckUserExistResponse represents the API response for check user existence.
type CheckUserExistResponse struct {
	Code       string         `json:"code"`
	Data       CheckUserExist `json:"data"`
	ErrorParam interface{}    `json:"errorParam"`
	ErrorMsg   string         `json:"msg"`
}

// GetAccountIdForRegister represents account id lookup response data.
type GetAccountIdForRegister struct {
	AccountId string `json:"accountId"`
}

// GetAccountIdForRegisterResponse represents the API response for account id lookup.
type GetAccountIdForRegisterResponse struct {
	Code       string                  `json:"code"`
	Data       GetAccountIdForRegister `json:"data"`
	ErrorParam interface{}             `json:"errorParam"`
	ErrorMsg   string                  `json:"msg"`
}

// SignerWithPermissions defines an extra signer and permission bitmask.
type SignerWithPermissions struct {
	Signer      string `json:"signer"`
	Permissions string `json:"permissions"`
}

// OnboardSiteV2Params defines the request payload for v2 onboarding.
type OnboardSiteV2Params struct {
	EthAddress            string                  `json:"ethAddress"`
	OnlySignOn            string                  `json:"onlySignOn"`
	Param                 string                  `json:"param,omitempty"`
	Signature             string                  `json:"signature"`
	ClientAccountId       string                  `json:"clientAccountId,omitempty"`
	PrivyIdentityToken    string                  `json:"privyIdentityToken,omitempty"`
	EthSignature          string                  `json:"ethSignature,omitempty"`
	ExtraSigners          []string                `json:"extraSigners,omitempty"`
	SignerWithPermissions []SignerWithPermissions `json:"signerWithPermissions,omitempty"`
	HintAccountId         string                  `json:"hintAccountId,omitempty"`
}

// OnboardSiteResponse represents the API response for onboarding.
type OnboardSiteResponse struct {
	Code       string                 `json:"code"`
	Data       map[string]interface{} `json:"data"`
	ErrorParam interface{}            `json:"errorParam"`
	ErrorMsg   string                 `json:"msg"`
}
