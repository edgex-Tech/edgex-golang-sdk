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

// checks if the NormalWithdraw type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NormalWithdraw{}

// NormalWithdraw 普通提现单
type NormalWithdraw struct {
	// 提现单id
	Id *string `json:"id,omitempty"`
	// 所属用户id
	UserId *string `json:"userId,omitempty"`
	// 所属账户id
	AccountId *string `json:"accountId,omitempty"`
	// 所属抵押品币种id
	CoinId *string `json:"coinId,omitempty"`
	// 充值数量
	Amount *string `json:"amount,omitempty"`
	// 充值的eth地址，可能与账户里的eth地址不一样。
	EthAddress *string `json:"ethAddress,omitempty"`
	// 客户自定义id，用于幂等校验
	ClientWithdrawId *string `json:"clientWithdrawId,omitempty"`
	// l2签名nonce。取sha256(client_withdraw_id) 前32个bit
	L2Nonce *string `json:"l2Nonce,omitempty"`
	// l2签名过期时间。unix时间的小时数，至少要比下单时间晚24个小时
	L2ExpireTime *string `json:"l2ExpireTime,omitempty"`
	L2Signature *L2Signature `json:"l2Signature,omitempty"`
	// 普通提现单状态
	Status *string `json:"status,omitempty"`
	// 对应的交易服务withdraw单id
	TradeWithdrawId *string `json:"tradeWithdrawId,omitempty"`
	RiskSignature *L2Signature `json:"riskSignature,omitempty"`
	L1ConfirmedTx *L1Tx `json:"l1ConfirmedTx,omitempty"`
	// l1交易确认时间
	L1ConfirmedTime *string `json:"l1ConfirmedTime,omitempty"`
	// l1提现完成时间
	L1CompletedTime *string `json:"l1CompletedTime,omitempty"`
	// 创建时间
	CreatedTime *string `json:"createdTime,omitempty"`
	// 更新时间
	UpdatedTime *string `json:"updatedTime,omitempty"`
}

// NewNormalWithdraw instantiates a new NormalWithdraw object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNormalWithdraw() *NormalWithdraw {
	this := NormalWithdraw{}
	return &this
}

// NewNormalWithdrawWithDefaults instantiates a new NormalWithdraw object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNormalWithdrawWithDefaults() *NormalWithdraw {
	this := NormalWithdraw{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NormalWithdraw) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalWithdraw) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NormalWithdraw) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NormalWithdraw) SetId(v string) {
	o.Id = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *NormalWithdraw) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalWithdraw) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *NormalWithdraw) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *NormalWithdraw) SetUserId(v string) {
	o.UserId = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *NormalWithdraw) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalWithdraw) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *NormalWithdraw) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *NormalWithdraw) SetAccountId(v string) {
	o.AccountId = &v
}

// GetCoinId returns the CoinId field value if set, zero value otherwise.
func (o *NormalWithdraw) GetCoinId() string {
	if o == nil || IsNil(o.CoinId) {
		var ret string
		return ret
	}
	return *o.CoinId
}

// GetCoinIdOk returns a tuple with the CoinId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalWithdraw) GetCoinIdOk() (*string, bool) {
	if o == nil || IsNil(o.CoinId) {
		return nil, false
	}
	return o.CoinId, true
}

// HasCoinId returns a boolean if a field has been set.
func (o *NormalWithdraw) HasCoinId() bool {
	if o != nil && !IsNil(o.CoinId) {
		return true
	}

	return false
}

// SetCoinId gets a reference to the given string and assigns it to the CoinId field.
func (o *NormalWithdraw) SetCoinId(v string) {
	o.CoinId = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *NormalWithdraw) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalWithdraw) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *NormalWithdraw) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *NormalWithdraw) SetAmount(v string) {
	o.Amount = &v
}

// GetEthAddress returns the EthAddress field value if set, zero value otherwise.
func (o *NormalWithdraw) GetEthAddress() string {
	if o == nil || IsNil(o.EthAddress) {
		var ret string
		return ret
	}
	return *o.EthAddress
}

// GetEthAddressOk returns a tuple with the EthAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalWithdraw) GetEthAddressOk() (*string, bool) {
	if o == nil || IsNil(o.EthAddress) {
		return nil, false
	}
	return o.EthAddress, true
}

// HasEthAddress returns a boolean if a field has been set.
func (o *NormalWithdraw) HasEthAddress() bool {
	if o != nil && !IsNil(o.EthAddress) {
		return true
	}

	return false
}

// SetEthAddress gets a reference to the given string and assigns it to the EthAddress field.
func (o *NormalWithdraw) SetEthAddress(v string) {
	o.EthAddress = &v
}

// GetClientWithdrawId returns the ClientWithdrawId field value if set, zero value otherwise.
func (o *NormalWithdraw) GetClientWithdrawId() string {
	if o == nil || IsNil(o.ClientWithdrawId) {
		var ret string
		return ret
	}
	return *o.ClientWithdrawId
}

// GetClientWithdrawIdOk returns a tuple with the ClientWithdrawId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalWithdraw) GetClientWithdrawIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientWithdrawId) {
		return nil, false
	}
	return o.ClientWithdrawId, true
}

// HasClientWithdrawId returns a boolean if a field has been set.
func (o *NormalWithdraw) HasClientWithdrawId() bool {
	if o != nil && !IsNil(o.ClientWithdrawId) {
		return true
	}

	return false
}

// SetClientWithdrawId gets a reference to the given string and assigns it to the ClientWithdrawId field.
func (o *NormalWithdraw) SetClientWithdrawId(v string) {
	o.ClientWithdrawId = &v
}

// GetL2Nonce returns the L2Nonce field value if set, zero value otherwise.
func (o *NormalWithdraw) GetL2Nonce() string {
	if o == nil || IsNil(o.L2Nonce) {
		var ret string
		return ret
	}
	return *o.L2Nonce
}

// GetL2NonceOk returns a tuple with the L2Nonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalWithdraw) GetL2NonceOk() (*string, bool) {
	if o == nil || IsNil(o.L2Nonce) {
		return nil, false
	}
	return o.L2Nonce, true
}

// HasL2Nonce returns a boolean if a field has been set.
func (o *NormalWithdraw) HasL2Nonce() bool {
	if o != nil && !IsNil(o.L2Nonce) {
		return true
	}

	return false
}

// SetL2Nonce gets a reference to the given string and assigns it to the L2Nonce field.
func (o *NormalWithdraw) SetL2Nonce(v string) {
	o.L2Nonce = &v
}

// GetL2ExpireTime returns the L2ExpireTime field value if set, zero value otherwise.
func (o *NormalWithdraw) GetL2ExpireTime() string {
	if o == nil || IsNil(o.L2ExpireTime) {
		var ret string
		return ret
	}
	return *o.L2ExpireTime
}

// GetL2ExpireTimeOk returns a tuple with the L2ExpireTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalWithdraw) GetL2ExpireTimeOk() (*string, bool) {
	if o == nil || IsNil(o.L2ExpireTime) {
		return nil, false
	}
	return o.L2ExpireTime, true
}

// HasL2ExpireTime returns a boolean if a field has been set.
func (o *NormalWithdraw) HasL2ExpireTime() bool {
	if o != nil && !IsNil(o.L2ExpireTime) {
		return true
	}

	return false
}

// SetL2ExpireTime gets a reference to the given string and assigns it to the L2ExpireTime field.
func (o *NormalWithdraw) SetL2ExpireTime(v string) {
	o.L2ExpireTime = &v
}

// GetL2Signature returns the L2Signature field value if set, zero value otherwise.
func (o *NormalWithdraw) GetL2Signature() L2Signature {
	if o == nil || IsNil(o.L2Signature) {
		var ret L2Signature
		return ret
	}
	return *o.L2Signature
}

// GetL2SignatureOk returns a tuple with the L2Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalWithdraw) GetL2SignatureOk() (*L2Signature, bool) {
	if o == nil || IsNil(o.L2Signature) {
		return nil, false
	}
	return o.L2Signature, true
}

// HasL2Signature returns a boolean if a field has been set.
func (o *NormalWithdraw) HasL2Signature() bool {
	if o != nil && !IsNil(o.L2Signature) {
		return true
	}

	return false
}

// SetL2Signature gets a reference to the given L2Signature and assigns it to the L2Signature field.
func (o *NormalWithdraw) SetL2Signature(v L2Signature) {
	o.L2Signature = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NormalWithdraw) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalWithdraw) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NormalWithdraw) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *NormalWithdraw) SetStatus(v string) {
	o.Status = &v
}

// GetTradeWithdrawId returns the TradeWithdrawId field value if set, zero value otherwise.
func (o *NormalWithdraw) GetTradeWithdrawId() string {
	if o == nil || IsNil(o.TradeWithdrawId) {
		var ret string
		return ret
	}
	return *o.TradeWithdrawId
}

// GetTradeWithdrawIdOk returns a tuple with the TradeWithdrawId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalWithdraw) GetTradeWithdrawIdOk() (*string, bool) {
	if o == nil || IsNil(o.TradeWithdrawId) {
		return nil, false
	}
	return o.TradeWithdrawId, true
}

// HasTradeWithdrawId returns a boolean if a field has been set.
func (o *NormalWithdraw) HasTradeWithdrawId() bool {
	if o != nil && !IsNil(o.TradeWithdrawId) {
		return true
	}

	return false
}

// SetTradeWithdrawId gets a reference to the given string and assigns it to the TradeWithdrawId field.
func (o *NormalWithdraw) SetTradeWithdrawId(v string) {
	o.TradeWithdrawId = &v
}

// GetRiskSignature returns the RiskSignature field value if set, zero value otherwise.
func (o *NormalWithdraw) GetRiskSignature() L2Signature {
	if o == nil || IsNil(o.RiskSignature) {
		var ret L2Signature
		return ret
	}
	return *o.RiskSignature
}

// GetRiskSignatureOk returns a tuple with the RiskSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalWithdraw) GetRiskSignatureOk() (*L2Signature, bool) {
	if o == nil || IsNil(o.RiskSignature) {
		return nil, false
	}
	return o.RiskSignature, true
}

// HasRiskSignature returns a boolean if a field has been set.
func (o *NormalWithdraw) HasRiskSignature() bool {
	if o != nil && !IsNil(o.RiskSignature) {
		return true
	}

	return false
}

// SetRiskSignature gets a reference to the given L2Signature and assigns it to the RiskSignature field.
func (o *NormalWithdraw) SetRiskSignature(v L2Signature) {
	o.RiskSignature = &v
}

// GetL1ConfirmedTx returns the L1ConfirmedTx field value if set, zero value otherwise.
func (o *NormalWithdraw) GetL1ConfirmedTx() L1Tx {
	if o == nil || IsNil(o.L1ConfirmedTx) {
		var ret L1Tx
		return ret
	}
	return *o.L1ConfirmedTx
}

// GetL1ConfirmedTxOk returns a tuple with the L1ConfirmedTx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalWithdraw) GetL1ConfirmedTxOk() (*L1Tx, bool) {
	if o == nil || IsNil(o.L1ConfirmedTx) {
		return nil, false
	}
	return o.L1ConfirmedTx, true
}

// HasL1ConfirmedTx returns a boolean if a field has been set.
func (o *NormalWithdraw) HasL1ConfirmedTx() bool {
	if o != nil && !IsNil(o.L1ConfirmedTx) {
		return true
	}

	return false
}

// SetL1ConfirmedTx gets a reference to the given L1Tx and assigns it to the L1ConfirmedTx field.
func (o *NormalWithdraw) SetL1ConfirmedTx(v L1Tx) {
	o.L1ConfirmedTx = &v
}

// GetL1ConfirmedTime returns the L1ConfirmedTime field value if set, zero value otherwise.
func (o *NormalWithdraw) GetL1ConfirmedTime() string {
	if o == nil || IsNil(o.L1ConfirmedTime) {
		var ret string
		return ret
	}
	return *o.L1ConfirmedTime
}

// GetL1ConfirmedTimeOk returns a tuple with the L1ConfirmedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalWithdraw) GetL1ConfirmedTimeOk() (*string, bool) {
	if o == nil || IsNil(o.L1ConfirmedTime) {
		return nil, false
	}
	return o.L1ConfirmedTime, true
}

// HasL1ConfirmedTime returns a boolean if a field has been set.
func (o *NormalWithdraw) HasL1ConfirmedTime() bool {
	if o != nil && !IsNil(o.L1ConfirmedTime) {
		return true
	}

	return false
}

// SetL1ConfirmedTime gets a reference to the given string and assigns it to the L1ConfirmedTime field.
func (o *NormalWithdraw) SetL1ConfirmedTime(v string) {
	o.L1ConfirmedTime = &v
}

// GetL1CompletedTime returns the L1CompletedTime field value if set, zero value otherwise.
func (o *NormalWithdraw) GetL1CompletedTime() string {
	if o == nil || IsNil(o.L1CompletedTime) {
		var ret string
		return ret
	}
	return *o.L1CompletedTime
}

// GetL1CompletedTimeOk returns a tuple with the L1CompletedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalWithdraw) GetL1CompletedTimeOk() (*string, bool) {
	if o == nil || IsNil(o.L1CompletedTime) {
		return nil, false
	}
	return o.L1CompletedTime, true
}

// HasL1CompletedTime returns a boolean if a field has been set.
func (o *NormalWithdraw) HasL1CompletedTime() bool {
	if o != nil && !IsNil(o.L1CompletedTime) {
		return true
	}

	return false
}

// SetL1CompletedTime gets a reference to the given string and assigns it to the L1CompletedTime field.
func (o *NormalWithdraw) SetL1CompletedTime(v string) {
	o.L1CompletedTime = &v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *NormalWithdraw) GetCreatedTime() string {
	if o == nil || IsNil(o.CreatedTime) {
		var ret string
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalWithdraw) GetCreatedTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *NormalWithdraw) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given string and assigns it to the CreatedTime field.
func (o *NormalWithdraw) SetCreatedTime(v string) {
	o.CreatedTime = &v
}

// GetUpdatedTime returns the UpdatedTime field value if set, zero value otherwise.
func (o *NormalWithdraw) GetUpdatedTime() string {
	if o == nil || IsNil(o.UpdatedTime) {
		var ret string
		return ret
	}
	return *o.UpdatedTime
}

// GetUpdatedTimeOk returns a tuple with the UpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalWithdraw) GetUpdatedTimeOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedTime) {
		return nil, false
	}
	return o.UpdatedTime, true
}

// HasUpdatedTime returns a boolean if a field has been set.
func (o *NormalWithdraw) HasUpdatedTime() bool {
	if o != nil && !IsNil(o.UpdatedTime) {
		return true
	}

	return false
}

// SetUpdatedTime gets a reference to the given string and assigns it to the UpdatedTime field.
func (o *NormalWithdraw) SetUpdatedTime(v string) {
	o.UpdatedTime = &v
}

func (o NormalWithdraw) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NormalWithdraw) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
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
	if !IsNil(o.ClientWithdrawId) {
		toSerialize["clientWithdrawId"] = o.ClientWithdrawId
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
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TradeWithdrawId) {
		toSerialize["tradeWithdrawId"] = o.TradeWithdrawId
	}
	if !IsNil(o.RiskSignature) {
		toSerialize["riskSignature"] = o.RiskSignature
	}
	if !IsNil(o.L1ConfirmedTx) {
		toSerialize["l1ConfirmedTx"] = o.L1ConfirmedTx
	}
	if !IsNil(o.L1ConfirmedTime) {
		toSerialize["l1ConfirmedTime"] = o.L1ConfirmedTime
	}
	if !IsNil(o.L1CompletedTime) {
		toSerialize["l1CompletedTime"] = o.L1CompletedTime
	}
	if !IsNil(o.CreatedTime) {
		toSerialize["createdTime"] = o.CreatedTime
	}
	if !IsNil(o.UpdatedTime) {
		toSerialize["updatedTime"] = o.UpdatedTime
	}
	return toSerialize, nil
}

type NullableNormalWithdraw struct {
	value *NormalWithdraw
	isSet bool
}

func (v NullableNormalWithdraw) Get() *NormalWithdraw {
	return v.value
}

func (v *NullableNormalWithdraw) Set(val *NormalWithdraw) {
	v.value = val
	v.isSet = true
}

func (v NullableNormalWithdraw) IsSet() bool {
	return v.isSet
}

func (v *NullableNormalWithdraw) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNormalWithdraw(val *NormalWithdraw) *NullableNormalWithdraw {
	return &NullableNormalWithdraw{value: val, isSet: true}
}

func (v NullableNormalWithdraw) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNormalWithdraw) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


