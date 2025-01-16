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

// checks if the AssetOrder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetOrder{}

// AssetOrder Asset订单信息
type AssetOrder struct {
	// 订单id
	OrderId *string `json:"orderId,omitempty"`
	// 订单创建时间
	Time *string `json:"time,omitempty"`
	// 订单类型
	Type *string `json:"type,omitempty"`
	// 订单状态
	Status *int32 `json:"status,omitempty"`
	// 订单金额
	Amount *string `json:"amount,omitempty"`
	// 订单手续费
	Fee *string `json:"fee,omitempty"`
	// 链上tx_id
	TxId *string `json:"txId,omitempty"`
	// 链
	Chain *string `json:"chain,omitempty"`
	// 地址
	Address *string `json:"address,omitempty"`
	// 币种
	Coin *string `json:"coin,omitempty"`
	// 链id
	ChainId *string `json:"chainId,omitempty"`
	// transfer转出账号id
	TransferSenderAccountId *string `json:"transferSenderAccountId,omitempty"`
	// transfer转入账号id
	TransferReceiverAccountId *string `json:"transferReceiverAccountId,omitempty"`
}

// NewAssetOrder instantiates a new AssetOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetOrder() *AssetOrder {
	this := AssetOrder{}
	return &this
}

// NewAssetOrderWithDefaults instantiates a new AssetOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetOrderWithDefaults() *AssetOrder {
	this := AssetOrder{}
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *AssetOrder) GetOrderId() string {
	if o == nil || IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetOrder) GetOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *AssetOrder) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *AssetOrder) SetOrderId(v string) {
	o.OrderId = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *AssetOrder) GetTime() string {
	if o == nil || IsNil(o.Time) {
		var ret string
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetOrder) GetTimeOk() (*string, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *AssetOrder) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given string and assigns it to the Time field.
func (o *AssetOrder) SetTime(v string) {
	o.Time = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AssetOrder) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetOrder) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AssetOrder) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AssetOrder) SetType(v string) {
	o.Type = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AssetOrder) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetOrder) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AssetOrder) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *AssetOrder) SetStatus(v int32) {
	o.Status = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *AssetOrder) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetOrder) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *AssetOrder) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *AssetOrder) SetAmount(v string) {
	o.Amount = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *AssetOrder) GetFee() string {
	if o == nil || IsNil(o.Fee) {
		var ret string
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetOrder) GetFeeOk() (*string, bool) {
	if o == nil || IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *AssetOrder) HasFee() bool {
	if o != nil && !IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given string and assigns it to the Fee field.
func (o *AssetOrder) SetFee(v string) {
	o.Fee = &v
}

// GetTxId returns the TxId field value if set, zero value otherwise.
func (o *AssetOrder) GetTxId() string {
	if o == nil || IsNil(o.TxId) {
		var ret string
		return ret
	}
	return *o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetOrder) GetTxIdOk() (*string, bool) {
	if o == nil || IsNil(o.TxId) {
		return nil, false
	}
	return o.TxId, true
}

// HasTxId returns a boolean if a field has been set.
func (o *AssetOrder) HasTxId() bool {
	if o != nil && !IsNil(o.TxId) {
		return true
	}

	return false
}

// SetTxId gets a reference to the given string and assigns it to the TxId field.
func (o *AssetOrder) SetTxId(v string) {
	o.TxId = &v
}

// GetChain returns the Chain field value if set, zero value otherwise.
func (o *AssetOrder) GetChain() string {
	if o == nil || IsNil(o.Chain) {
		var ret string
		return ret
	}
	return *o.Chain
}

// GetChainOk returns a tuple with the Chain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetOrder) GetChainOk() (*string, bool) {
	if o == nil || IsNil(o.Chain) {
		return nil, false
	}
	return o.Chain, true
}

// HasChain returns a boolean if a field has been set.
func (o *AssetOrder) HasChain() bool {
	if o != nil && !IsNil(o.Chain) {
		return true
	}

	return false
}

// SetChain gets a reference to the given string and assigns it to the Chain field.
func (o *AssetOrder) SetChain(v string) {
	o.Chain = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *AssetOrder) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetOrder) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *AssetOrder) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *AssetOrder) SetAddress(v string) {
	o.Address = &v
}

// GetCoin returns the Coin field value if set, zero value otherwise.
func (o *AssetOrder) GetCoin() string {
	if o == nil || IsNil(o.Coin) {
		var ret string
		return ret
	}
	return *o.Coin
}

// GetCoinOk returns a tuple with the Coin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetOrder) GetCoinOk() (*string, bool) {
	if o == nil || IsNil(o.Coin) {
		return nil, false
	}
	return o.Coin, true
}

// HasCoin returns a boolean if a field has been set.
func (o *AssetOrder) HasCoin() bool {
	if o != nil && !IsNil(o.Coin) {
		return true
	}

	return false
}

// SetCoin gets a reference to the given string and assigns it to the Coin field.
func (o *AssetOrder) SetCoin(v string) {
	o.Coin = &v
}

// GetChainId returns the ChainId field value if set, zero value otherwise.
func (o *AssetOrder) GetChainId() string {
	if o == nil || IsNil(o.ChainId) {
		var ret string
		return ret
	}
	return *o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetOrder) GetChainIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChainId) {
		return nil, false
	}
	return o.ChainId, true
}

// HasChainId returns a boolean if a field has been set.
func (o *AssetOrder) HasChainId() bool {
	if o != nil && !IsNil(o.ChainId) {
		return true
	}

	return false
}

// SetChainId gets a reference to the given string and assigns it to the ChainId field.
func (o *AssetOrder) SetChainId(v string) {
	o.ChainId = &v
}

// GetTransferSenderAccountId returns the TransferSenderAccountId field value if set, zero value otherwise.
func (o *AssetOrder) GetTransferSenderAccountId() string {
	if o == nil || IsNil(o.TransferSenderAccountId) {
		var ret string
		return ret
	}
	return *o.TransferSenderAccountId
}

// GetTransferSenderAccountIdOk returns a tuple with the TransferSenderAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetOrder) GetTransferSenderAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransferSenderAccountId) {
		return nil, false
	}
	return o.TransferSenderAccountId, true
}

// HasTransferSenderAccountId returns a boolean if a field has been set.
func (o *AssetOrder) HasTransferSenderAccountId() bool {
	if o != nil && !IsNil(o.TransferSenderAccountId) {
		return true
	}

	return false
}

// SetTransferSenderAccountId gets a reference to the given string and assigns it to the TransferSenderAccountId field.
func (o *AssetOrder) SetTransferSenderAccountId(v string) {
	o.TransferSenderAccountId = &v
}

// GetTransferReceiverAccountId returns the TransferReceiverAccountId field value if set, zero value otherwise.
func (o *AssetOrder) GetTransferReceiverAccountId() string {
	if o == nil || IsNil(o.TransferReceiverAccountId) {
		var ret string
		return ret
	}
	return *o.TransferReceiverAccountId
}

// GetTransferReceiverAccountIdOk returns a tuple with the TransferReceiverAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetOrder) GetTransferReceiverAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransferReceiverAccountId) {
		return nil, false
	}
	return o.TransferReceiverAccountId, true
}

// HasTransferReceiverAccountId returns a boolean if a field has been set.
func (o *AssetOrder) HasTransferReceiverAccountId() bool {
	if o != nil && !IsNil(o.TransferReceiverAccountId) {
		return true
	}

	return false
}

// SetTransferReceiverAccountId gets a reference to the given string and assigns it to the TransferReceiverAccountId field.
func (o *AssetOrder) SetTransferReceiverAccountId(v string) {
	o.TransferReceiverAccountId = &v
}

func (o AssetOrder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetOrder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	if !IsNil(o.TxId) {
		toSerialize["txId"] = o.TxId
	}
	if !IsNil(o.Chain) {
		toSerialize["chain"] = o.Chain
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Coin) {
		toSerialize["coin"] = o.Coin
	}
	if !IsNil(o.ChainId) {
		toSerialize["chainId"] = o.ChainId
	}
	if !IsNil(o.TransferSenderAccountId) {
		toSerialize["transferSenderAccountId"] = o.TransferSenderAccountId
	}
	if !IsNil(o.TransferReceiverAccountId) {
		toSerialize["transferReceiverAccountId"] = o.TransferReceiverAccountId
	}
	return toSerialize, nil
}

type NullableAssetOrder struct {
	value *AssetOrder
	isSet bool
}

func (v NullableAssetOrder) Get() *AssetOrder {
	return v.value
}

func (v *NullableAssetOrder) Set(val *AssetOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetOrder(val *AssetOrder) *NullableAssetOrder {
	return &NullableAssetOrder{value: val, isSet: true}
}

func (v NullableAssetOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


