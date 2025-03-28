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

// checks if the CheckUserExist type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckUserExist{}

// CheckUserExist 判断用户是否存在-响应
type CheckUserExist struct {
	// 用户是否存在
	IsUserExist *bool `json:"isUserExist,omitempty"`
}

// NewCheckUserExist instantiates a new CheckUserExist object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckUserExist() *CheckUserExist {
	this := CheckUserExist{}
	return &this
}

// NewCheckUserExistWithDefaults instantiates a new CheckUserExist object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckUserExistWithDefaults() *CheckUserExist {
	this := CheckUserExist{}
	return &this
}

// GetIsUserExist returns the IsUserExist field value if set, zero value otherwise.
func (o *CheckUserExist) GetIsUserExist() bool {
	if o == nil || IsNil(o.IsUserExist) {
		var ret bool
		return ret
	}
	return *o.IsUserExist
}

// GetIsUserExistOk returns a tuple with the IsUserExist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckUserExist) GetIsUserExistOk() (*bool, bool) {
	if o == nil || IsNil(o.IsUserExist) {
		return nil, false
	}
	return o.IsUserExist, true
}

// HasIsUserExist returns a boolean if a field has been set.
func (o *CheckUserExist) HasIsUserExist() bool {
	if o != nil && !IsNil(o.IsUserExist) {
		return true
	}

	return false
}

// SetIsUserExist gets a reference to the given bool and assigns it to the IsUserExist field.
func (o *CheckUserExist) SetIsUserExist(v bool) {
	o.IsUserExist = &v
}

func (o CheckUserExist) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckUserExist) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsUserExist) {
		toSerialize["isUserExist"] = o.IsUserExist
	}
	return toSerialize, nil
}

type NullableCheckUserExist struct {
	value *CheckUserExist
	isSet bool
}

func (v NullableCheckUserExist) Get() *CheckUserExist {
	return v.value
}

func (v *NullableCheckUserExist) Set(val *CheckUserExist) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckUserExist) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckUserExist) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckUserExist(val *CheckUserExist) *NullableCheckUserExist {
	return &NullableCheckUserExist{value: val, isSet: true}
}

func (v NullableCheckUserExist) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckUserExist) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


