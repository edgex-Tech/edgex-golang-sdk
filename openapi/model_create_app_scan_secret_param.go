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

// checks if the CreateAppScanSecretParam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAppScanSecretParam{}

// CreateAppScanSecretParam 创建AppScanSecret请求
type CreateAppScanSecretParam struct {
	// 过期时间戳，单位毫秒
	ExpireTime *string `json:"expireTime,omitempty"`
}

// NewCreateAppScanSecretParam instantiates a new CreateAppScanSecretParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAppScanSecretParam() *CreateAppScanSecretParam {
	this := CreateAppScanSecretParam{}
	return &this
}

// NewCreateAppScanSecretParamWithDefaults instantiates a new CreateAppScanSecretParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAppScanSecretParamWithDefaults() *CreateAppScanSecretParam {
	this := CreateAppScanSecretParam{}
	return &this
}

// GetExpireTime returns the ExpireTime field value if set, zero value otherwise.
func (o *CreateAppScanSecretParam) GetExpireTime() string {
	if o == nil || IsNil(o.ExpireTime) {
		var ret string
		return ret
	}
	return *o.ExpireTime
}

// GetExpireTimeOk returns a tuple with the ExpireTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppScanSecretParam) GetExpireTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ExpireTime) {
		return nil, false
	}
	return o.ExpireTime, true
}

// HasExpireTime returns a boolean if a field has been set.
func (o *CreateAppScanSecretParam) HasExpireTime() bool {
	if o != nil && !IsNil(o.ExpireTime) {
		return true
	}

	return false
}

// SetExpireTime gets a reference to the given string and assigns it to the ExpireTime field.
func (o *CreateAppScanSecretParam) SetExpireTime(v string) {
	o.ExpireTime = &v
}

func (o CreateAppScanSecretParam) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAppScanSecretParam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExpireTime) {
		toSerialize["expireTime"] = o.ExpireTime
	}
	return toSerialize, nil
}

type NullableCreateAppScanSecretParam struct {
	value *CreateAppScanSecretParam
	isSet bool
}

func (v NullableCreateAppScanSecretParam) Get() *CreateAppScanSecretParam {
	return v.value
}

func (v *NullableCreateAppScanSecretParam) Set(val *CreateAppScanSecretParam) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAppScanSecretParam) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAppScanSecretParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAppScanSecretParam(val *CreateAppScanSecretParam) *NullableCreateAppScanSecretParam {
	return &NullableCreateAppScanSecretParam{value: val, isSet: true}
}

func (v NullableCreateAppScanSecretParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAppScanSecretParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


