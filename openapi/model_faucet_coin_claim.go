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

// checks if the FaucetCoinClaim type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FaucetCoinClaim{}

// FaucetCoinClaim 水龙头领取记录
type FaucetCoinClaim struct {
	// 领取记录id
	Id *string `json:"id,omitempty"`
	// 领取用户id
	UserId *string `json:"userId,omitempty"`
	// 领取用户eth地址
	EthAddress *string `json:"ethAddress,omitempty"`
	// 水龙头发放地址
	FaucetAddress *string `json:"faucetAddress,omitempty"`
	// 领取状态
	Status *string `json:"status,omitempty"`
	// 领取时间
	CreatedTime *string `json:"createdTime,omitempty"`
	// 更新时间
	UpdatedTime *string `json:"updatedTime,omitempty"`
}

// NewFaucetCoinClaim instantiates a new FaucetCoinClaim object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFaucetCoinClaim() *FaucetCoinClaim {
	this := FaucetCoinClaim{}
	return &this
}

// NewFaucetCoinClaimWithDefaults instantiates a new FaucetCoinClaim object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFaucetCoinClaimWithDefaults() *FaucetCoinClaim {
	this := FaucetCoinClaim{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FaucetCoinClaim) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaucetCoinClaim) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FaucetCoinClaim) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FaucetCoinClaim) SetId(v string) {
	o.Id = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *FaucetCoinClaim) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaucetCoinClaim) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *FaucetCoinClaim) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *FaucetCoinClaim) SetUserId(v string) {
	o.UserId = &v
}

// GetEthAddress returns the EthAddress field value if set, zero value otherwise.
func (o *FaucetCoinClaim) GetEthAddress() string {
	if o == nil || IsNil(o.EthAddress) {
		var ret string
		return ret
	}
	return *o.EthAddress
}

// GetEthAddressOk returns a tuple with the EthAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaucetCoinClaim) GetEthAddressOk() (*string, bool) {
	if o == nil || IsNil(o.EthAddress) {
		return nil, false
	}
	return o.EthAddress, true
}

// HasEthAddress returns a boolean if a field has been set.
func (o *FaucetCoinClaim) HasEthAddress() bool {
	if o != nil && !IsNil(o.EthAddress) {
		return true
	}

	return false
}

// SetEthAddress gets a reference to the given string and assigns it to the EthAddress field.
func (o *FaucetCoinClaim) SetEthAddress(v string) {
	o.EthAddress = &v
}

// GetFaucetAddress returns the FaucetAddress field value if set, zero value otherwise.
func (o *FaucetCoinClaim) GetFaucetAddress() string {
	if o == nil || IsNil(o.FaucetAddress) {
		var ret string
		return ret
	}
	return *o.FaucetAddress
}

// GetFaucetAddressOk returns a tuple with the FaucetAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaucetCoinClaim) GetFaucetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.FaucetAddress) {
		return nil, false
	}
	return o.FaucetAddress, true
}

// HasFaucetAddress returns a boolean if a field has been set.
func (o *FaucetCoinClaim) HasFaucetAddress() bool {
	if o != nil && !IsNil(o.FaucetAddress) {
		return true
	}

	return false
}

// SetFaucetAddress gets a reference to the given string and assigns it to the FaucetAddress field.
func (o *FaucetCoinClaim) SetFaucetAddress(v string) {
	o.FaucetAddress = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FaucetCoinClaim) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaucetCoinClaim) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FaucetCoinClaim) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *FaucetCoinClaim) SetStatus(v string) {
	o.Status = &v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *FaucetCoinClaim) GetCreatedTime() string {
	if o == nil || IsNil(o.CreatedTime) {
		var ret string
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaucetCoinClaim) GetCreatedTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *FaucetCoinClaim) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given string and assigns it to the CreatedTime field.
func (o *FaucetCoinClaim) SetCreatedTime(v string) {
	o.CreatedTime = &v
}

// GetUpdatedTime returns the UpdatedTime field value if set, zero value otherwise.
func (o *FaucetCoinClaim) GetUpdatedTime() string {
	if o == nil || IsNil(o.UpdatedTime) {
		var ret string
		return ret
	}
	return *o.UpdatedTime
}

// GetUpdatedTimeOk returns a tuple with the UpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaucetCoinClaim) GetUpdatedTimeOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedTime) {
		return nil, false
	}
	return o.UpdatedTime, true
}

// HasUpdatedTime returns a boolean if a field has been set.
func (o *FaucetCoinClaim) HasUpdatedTime() bool {
	if o != nil && !IsNil(o.UpdatedTime) {
		return true
	}

	return false
}

// SetUpdatedTime gets a reference to the given string and assigns it to the UpdatedTime field.
func (o *FaucetCoinClaim) SetUpdatedTime(v string) {
	o.UpdatedTime = &v
}

func (o FaucetCoinClaim) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FaucetCoinClaim) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.EthAddress) {
		toSerialize["ethAddress"] = o.EthAddress
	}
	if !IsNil(o.FaucetAddress) {
		toSerialize["faucetAddress"] = o.FaucetAddress
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.CreatedTime) {
		toSerialize["createdTime"] = o.CreatedTime
	}
	if !IsNil(o.UpdatedTime) {
		toSerialize["updatedTime"] = o.UpdatedTime
	}
	return toSerialize, nil
}

type NullableFaucetCoinClaim struct {
	value *FaucetCoinClaim
	isSet bool
}

func (v NullableFaucetCoinClaim) Get() *FaucetCoinClaim {
	return v.value
}

func (v *NullableFaucetCoinClaim) Set(val *FaucetCoinClaim) {
	v.value = val
	v.isSet = true
}

func (v NullableFaucetCoinClaim) IsSet() bool {
	return v.isSet
}

func (v *NullableFaucetCoinClaim) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFaucetCoinClaim(val *FaucetCoinClaim) *NullableFaucetCoinClaim {
	return &NullableFaucetCoinClaim{value: val, isSet: true}
}

func (v NullableFaucetCoinClaim) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFaucetCoinClaim) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


