/*
Http Gateway

Contains interface documents such as accounts, assets, transactions, etc.

API version: v1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CreateWithdrawParam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateWithdrawParam{}

// CreateWithdrawParam 创建提现单-请求
type CreateWithdrawParam struct {
	// 账户id
	AccountId *string `json:"accountId,omitempty"`
	// 币id
	CoinId *string `json:"coinId,omitempty"`
	// 提现数量
	Amount *string `json:"amount,omitempty"`
	// 提现到的地址。
	EthAddress *string `json:"ethAddress,omitempty"`
	// 提现的币种合约地址
	Erc20Address *string `json:"erc20Address,omitempty"`
	// 客户自定义id，用于幂等校验
	ClientWithdrawId *string `json:"clientWithdrawId,omitempty"`
	// 风控签名
	RiskSignature *string `json:"riskSignature,omitempty"`
	// l2签名nonce。取sha256(client_withdraw_id) 前32个bit
	L2Nonce *string `json:"l2Nonce,omitempty"`
	// l2签名过期时间，单位毫秒。参与签名生成/校验时要取小时数，即 l2_expire_time / 3600000
	L2ExpireTime *string `json:"l2ExpireTime,omitempty"`
	// 提交l2的签名
	L2Signature *string `json:"l2Signature,omitempty"`
	// 附加类型，供上层业务使用
	ExtraType *string `json:"extraType,omitempty"`
	// 额外数据，json格式，默认为空串
	ExtraDataJson *string `json:"extraDataJson,omitempty"`
}

// NewCreateWithdrawParam instantiates a new CreateWithdrawParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWithdrawParam() *CreateWithdrawParam {
	this := CreateWithdrawParam{}
	return &this
}

// NewCreateWithdrawParamWithDefaults instantiates a new CreateWithdrawParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWithdrawParamWithDefaults() *CreateWithdrawParam {
	this := CreateWithdrawParam{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *CreateWithdrawParam) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWithdrawParam) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *CreateWithdrawParam) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *CreateWithdrawParam) SetAccountId(v string) {
	o.AccountId = &v
}

// GetCoinId returns the CoinId field value if set, zero value otherwise.
func (o *CreateWithdrawParam) GetCoinId() string {
	if o == nil || IsNil(o.CoinId) {
		var ret string
		return ret
	}
	return *o.CoinId
}

// GetCoinIdOk returns a tuple with the CoinId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWithdrawParam) GetCoinIdOk() (*string, bool) {
	if o == nil || IsNil(o.CoinId) {
		return nil, false
	}
	return o.CoinId, true
}

// HasCoinId returns a boolean if a field has been set.
func (o *CreateWithdrawParam) HasCoinId() bool {
	if o != nil && !IsNil(o.CoinId) {
		return true
	}

	return false
}

// SetCoinId gets a reference to the given string and assigns it to the CoinId field.
func (o *CreateWithdrawParam) SetCoinId(v string) {
	o.CoinId = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *CreateWithdrawParam) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWithdrawParam) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *CreateWithdrawParam) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *CreateWithdrawParam) SetAmount(v string) {
	o.Amount = &v
}

// GetEthAddress returns the EthAddress field value if set, zero value otherwise.
func (o *CreateWithdrawParam) GetEthAddress() string {
	if o == nil || IsNil(o.EthAddress) {
		var ret string
		return ret
	}
	return *o.EthAddress
}

// GetEthAddressOk returns a tuple with the EthAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWithdrawParam) GetEthAddressOk() (*string, bool) {
	if o == nil || IsNil(o.EthAddress) {
		return nil, false
	}
	return o.EthAddress, true
}

// HasEthAddress returns a boolean if a field has been set.
func (o *CreateWithdrawParam) HasEthAddress() bool {
	if o != nil && !IsNil(o.EthAddress) {
		return true
	}

	return false
}

// SetEthAddress gets a reference to the given string and assigns it to the EthAddress field.
func (o *CreateWithdrawParam) SetEthAddress(v string) {
	o.EthAddress = &v
}

// GetErc20Address returns the Erc20Address field value if set, zero value otherwise.
func (o *CreateWithdrawParam) GetErc20Address() string {
	if o == nil || IsNil(o.Erc20Address) {
		var ret string
		return ret
	}
	return *o.Erc20Address
}

// GetErc20AddressOk returns a tuple with the Erc20Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWithdrawParam) GetErc20AddressOk() (*string, bool) {
	if o == nil || IsNil(o.Erc20Address) {
		return nil, false
	}
	return o.Erc20Address, true
}

// HasErc20Address returns a boolean if a field has been set.
func (o *CreateWithdrawParam) HasErc20Address() bool {
	if o != nil && !IsNil(o.Erc20Address) {
		return true
	}

	return false
}

// SetErc20Address gets a reference to the given string and assigns it to the Erc20Address field.
func (o *CreateWithdrawParam) SetErc20Address(v string) {
	o.Erc20Address = &v
}

// GetClientWithdrawId returns the ClientWithdrawId field value if set, zero value otherwise.
func (o *CreateWithdrawParam) GetClientWithdrawId() string {
	if o == nil || IsNil(o.ClientWithdrawId) {
		var ret string
		return ret
	}
	return *o.ClientWithdrawId
}

// GetClientWithdrawIdOk returns a tuple with the ClientWithdrawId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWithdrawParam) GetClientWithdrawIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientWithdrawId) {
		return nil, false
	}
	return o.ClientWithdrawId, true
}

// HasClientWithdrawId returns a boolean if a field has been set.
func (o *CreateWithdrawParam) HasClientWithdrawId() bool {
	if o != nil && !IsNil(o.ClientWithdrawId) {
		return true
	}

	return false
}

// SetClientWithdrawId gets a reference to the given string and assigns it to the ClientWithdrawId field.
func (o *CreateWithdrawParam) SetClientWithdrawId(v string) {
	o.ClientWithdrawId = &v
}

// GetRiskSignature returns the RiskSignature field value if set, zero value otherwise.
func (o *CreateWithdrawParam) GetRiskSignature() string {
	if o == nil || IsNil(o.RiskSignature) {
		var ret string
		return ret
	}
	return *o.RiskSignature
}

// GetRiskSignatureOk returns a tuple with the RiskSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWithdrawParam) GetRiskSignatureOk() (*string, bool) {
	if o == nil || IsNil(o.RiskSignature) {
		return nil, false
	}
	return o.RiskSignature, true
}

// HasRiskSignature returns a boolean if a field has been set.
func (o *CreateWithdrawParam) HasRiskSignature() bool {
	if o != nil && !IsNil(o.RiskSignature) {
		return true
	}

	return false
}

// SetRiskSignature gets a reference to the given string and assigns it to the RiskSignature field.
func (o *CreateWithdrawParam) SetRiskSignature(v string) {
	o.RiskSignature = &v
}

// GetL2Nonce returns the L2Nonce field value if set, zero value otherwise.
func (o *CreateWithdrawParam) GetL2Nonce() string {
	if o == nil || IsNil(o.L2Nonce) {
		var ret string
		return ret
	}
	return *o.L2Nonce
}

// GetL2NonceOk returns a tuple with the L2Nonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWithdrawParam) GetL2NonceOk() (*string, bool) {
	if o == nil || IsNil(o.L2Nonce) {
		return nil, false
	}
	return o.L2Nonce, true
}

// HasL2Nonce returns a boolean if a field has been set.
func (o *CreateWithdrawParam) HasL2Nonce() bool {
	if o != nil && !IsNil(o.L2Nonce) {
		return true
	}

	return false
}

// SetL2Nonce gets a reference to the given string and assigns it to the L2Nonce field.
func (o *CreateWithdrawParam) SetL2Nonce(v string) {
	o.L2Nonce = &v
}

// GetL2ExpireTime returns the L2ExpireTime field value if set, zero value otherwise.
func (o *CreateWithdrawParam) GetL2ExpireTime() string {
	if o == nil || IsNil(o.L2ExpireTime) {
		var ret string
		return ret
	}
	return *o.L2ExpireTime
}

// GetL2ExpireTimeOk returns a tuple with the L2ExpireTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWithdrawParam) GetL2ExpireTimeOk() (*string, bool) {
	if o == nil || IsNil(o.L2ExpireTime) {
		return nil, false
	}
	return o.L2ExpireTime, true
}

// HasL2ExpireTime returns a boolean if a field has been set.
func (o *CreateWithdrawParam) HasL2ExpireTime() bool {
	if o != nil && !IsNil(o.L2ExpireTime) {
		return true
	}

	return false
}

// SetL2ExpireTime gets a reference to the given string and assigns it to the L2ExpireTime field.
func (o *CreateWithdrawParam) SetL2ExpireTime(v string) {
	o.L2ExpireTime = &v
}

// GetL2Signature returns the L2Signature field value if set, zero value otherwise.
func (o *CreateWithdrawParam) GetL2Signature() string {
	if o == nil || IsNil(o.L2Signature) {
		var ret string
		return ret
	}
	return *o.L2Signature
}

// GetL2SignatureOk returns a tuple with the L2Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWithdrawParam) GetL2SignatureOk() (*string, bool) {
	if o == nil || IsNil(o.L2Signature) {
		return nil, false
	}
	return o.L2Signature, true
}

// HasL2Signature returns a boolean if a field has been set.
func (o *CreateWithdrawParam) HasL2Signature() bool {
	if o != nil && !IsNil(o.L2Signature) {
		return true
	}

	return false
}

// SetL2Signature gets a reference to the given string and assigns it to the L2Signature field.
func (o *CreateWithdrawParam) SetL2Signature(v string) {
	o.L2Signature = &v
}

// GetExtraType returns the ExtraType field value if set, zero value otherwise.
func (o *CreateWithdrawParam) GetExtraType() string {
	if o == nil || IsNil(o.ExtraType) {
		var ret string
		return ret
	}
	return *o.ExtraType
}

// GetExtraTypeOk returns a tuple with the ExtraType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWithdrawParam) GetExtraTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ExtraType) {
		return nil, false
	}
	return o.ExtraType, true
}

// HasExtraType returns a boolean if a field has been set.
func (o *CreateWithdrawParam) HasExtraType() bool {
	if o != nil && !IsNil(o.ExtraType) {
		return true
	}

	return false
}

// SetExtraType gets a reference to the given string and assigns it to the ExtraType field.
func (o *CreateWithdrawParam) SetExtraType(v string) {
	o.ExtraType = &v
}

// GetExtraDataJson returns the ExtraDataJson field value if set, zero value otherwise.
func (o *CreateWithdrawParam) GetExtraDataJson() string {
	if o == nil || IsNil(o.ExtraDataJson) {
		var ret string
		return ret
	}
	return *o.ExtraDataJson
}

// GetExtraDataJsonOk returns a tuple with the ExtraDataJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWithdrawParam) GetExtraDataJsonOk() (*string, bool) {
	if o == nil || IsNil(o.ExtraDataJson) {
		return nil, false
	}
	return o.ExtraDataJson, true
}

// HasExtraDataJson returns a boolean if a field has been set.
func (o *CreateWithdrawParam) HasExtraDataJson() bool {
	if o != nil && !IsNil(o.ExtraDataJson) {
		return true
	}

	return false
}

// SetExtraDataJson gets a reference to the given string and assigns it to the ExtraDataJson field.
func (o *CreateWithdrawParam) SetExtraDataJson(v string) {
	o.ExtraDataJson = &v
}

func (o CreateWithdrawParam) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateWithdrawParam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	if !IsNil(o.CoinId) {
		toSerialize["coinId"] = o.CoinId
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.EthAddress) {
		toSerialize["ethAddress"] = o.EthAddress
	}
	if !IsNil(o.Erc20Address) {
		toSerialize["erc20Address"] = o.Erc20Address
	}
	if !IsNil(o.ClientWithdrawId) {
		toSerialize["clientWithdrawId"] = o.ClientWithdrawId
	}
	if !IsNil(o.RiskSignature) {
		toSerialize["riskSignature"] = o.RiskSignature
	}
	if !IsNil(o.L2Nonce) {
		toSerialize["l2Nonce"] = o.L2Nonce
	}
	if !IsNil(o.L2ExpireTime) {
		toSerialize["l2ExpireTime"] = o.L2ExpireTime
	}
	if !IsNil(o.L2Signature) {
		toSerialize["l2Signature"] = o.L2Signature
	}
	if !IsNil(o.ExtraType) {
		toSerialize["extraType"] = o.ExtraType
	}
	if !IsNil(o.ExtraDataJson) {
		toSerialize["extraDataJson"] = o.ExtraDataJson
	}
	return toSerialize, nil
}

type NullableCreateWithdrawParam struct {
	value *CreateWithdrawParam
	isSet bool
}

func (v NullableCreateWithdrawParam) Get() *CreateWithdrawParam {
	return v.value
}

func (v *NullableCreateWithdrawParam) Set(val *CreateWithdrawParam) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWithdrawParam) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWithdrawParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWithdrawParam(val *CreateWithdrawParam) *NullableCreateWithdrawParam {
	return &NullableCreateWithdrawParam{value: val, isSet: true}
}

func (v NullableCreateWithdrawParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWithdrawParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


