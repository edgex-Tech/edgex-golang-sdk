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

// checks if the OpenInterest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenInterest{}

// OpenInterest 持仓量
type OpenInterest struct {
	// 合约ID
	ContractId *string `json:"contractId,omitempty"`
	// 统计时间戳
	Timestamp *string `json:"timestamp,omitempty"`
	// 持仓量
	Size *string `json:"size,omitempty"`
}

// NewOpenInterest instantiates a new OpenInterest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenInterest() *OpenInterest {
	this := OpenInterest{}
	return &this
}

// NewOpenInterestWithDefaults instantiates a new OpenInterest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenInterestWithDefaults() *OpenInterest {
	this := OpenInterest{}
	return &this
}

// GetContractId returns the ContractId field value if set, zero value otherwise.
func (o *OpenInterest) GetContractId() string {
	if o == nil || IsNil(o.ContractId) {
		var ret string
		return ret
	}
	return *o.ContractId
}

// GetContractIdOk returns a tuple with the ContractId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenInterest) GetContractIdOk() (*string, bool) {
	if o == nil || IsNil(o.ContractId) {
		return nil, false
	}
	return o.ContractId, true
}

// HasContractId returns a boolean if a field has been set.
func (o *OpenInterest) HasContractId() bool {
	if o != nil && !IsNil(o.ContractId) {
		return true
	}

	return false
}

// SetContractId gets a reference to the given string and assigns it to the ContractId field.
func (o *OpenInterest) SetContractId(v string) {
	o.ContractId = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *OpenInterest) GetTimestamp() string {
	if o == nil || IsNil(o.Timestamp) {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenInterest) GetTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *OpenInterest) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *OpenInterest) SetTimestamp(v string) {
	o.Timestamp = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *OpenInterest) GetSize() string {
	if o == nil || IsNil(o.Size) {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenInterest) GetSizeOk() (*string, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *OpenInterest) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *OpenInterest) SetSize(v string) {
	o.Size = &v
}

func (o OpenInterest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenInterest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContractId) {
		toSerialize["contractId"] = o.ContractId
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	return toSerialize, nil
}

type NullableOpenInterest struct {
	value *OpenInterest
	isSet bool
}

func (v NullableOpenInterest) Get() *OpenInterest {
	return v.value
}

func (v *NullableOpenInterest) Set(val *OpenInterest) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenInterest) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenInterest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenInterest(val *OpenInterest) *NullableOpenInterest {
	return &NullableOpenInterest{value: val, isSet: true}
}

func (v NullableOpenInterest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenInterest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


