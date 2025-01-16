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

// checks if the RegisterAccountParam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegisterAccountParam{}

// RegisterAccountParam 注册账户-请求
type RegisterAccountParam struct {
	// L2上的账户key，保证全局唯一。对应starkEx中的starkKey。bigint for hex str
	L2Key *string `json:"l2Key,omitempty"`
	// 只用于验证 l2Signature 是否ok。不返回给C端用户。bigint for hex str
	L2KeyYCoordinate *string `json:"l2KeyYCoordinate,omitempty"`
	// 客户端账户id, 用于幂等性校验
	ClientAccountId *string `json:"clientAccountId,omitempty"`
}

// NewRegisterAccountParam instantiates a new RegisterAccountParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegisterAccountParam() *RegisterAccountParam {
	this := RegisterAccountParam{}
	return &this
}

// NewRegisterAccountParamWithDefaults instantiates a new RegisterAccountParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegisterAccountParamWithDefaults() *RegisterAccountParam {
	this := RegisterAccountParam{}
	return &this
}

// GetL2Key returns the L2Key field value if set, zero value otherwise.
func (o *RegisterAccountParam) GetL2Key() string {
	if o == nil || IsNil(o.L2Key) {
		var ret string
		return ret
	}
	return *o.L2Key
}

// GetL2KeyOk returns a tuple with the L2Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterAccountParam) GetL2KeyOk() (*string, bool) {
	if o == nil || IsNil(o.L2Key) {
		return nil, false
	}
	return o.L2Key, true
}

// HasL2Key returns a boolean if a field has been set.
func (o *RegisterAccountParam) HasL2Key() bool {
	if o != nil && !IsNil(o.L2Key) {
		return true
	}

	return false
}

// SetL2Key gets a reference to the given string and assigns it to the L2Key field.
func (o *RegisterAccountParam) SetL2Key(v string) {
	o.L2Key = &v
}

// GetL2KeyYCoordinate returns the L2KeyYCoordinate field value if set, zero value otherwise.
func (o *RegisterAccountParam) GetL2KeyYCoordinate() string {
	if o == nil || IsNil(o.L2KeyYCoordinate) {
		var ret string
		return ret
	}
	return *o.L2KeyYCoordinate
}

// GetL2KeyYCoordinateOk returns a tuple with the L2KeyYCoordinate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterAccountParam) GetL2KeyYCoordinateOk() (*string, bool) {
	if o == nil || IsNil(o.L2KeyYCoordinate) {
		return nil, false
	}
	return o.L2KeyYCoordinate, true
}

// HasL2KeyYCoordinate returns a boolean if a field has been set.
func (o *RegisterAccountParam) HasL2KeyYCoordinate() bool {
	if o != nil && !IsNil(o.L2KeyYCoordinate) {
		return true
	}

	return false
}

// SetL2KeyYCoordinate gets a reference to the given string and assigns it to the L2KeyYCoordinate field.
func (o *RegisterAccountParam) SetL2KeyYCoordinate(v string) {
	o.L2KeyYCoordinate = &v
}

// GetClientAccountId returns the ClientAccountId field value if set, zero value otherwise.
func (o *RegisterAccountParam) GetClientAccountId() string {
	if o == nil || IsNil(o.ClientAccountId) {
		var ret string
		return ret
	}
	return *o.ClientAccountId
}

// GetClientAccountIdOk returns a tuple with the ClientAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterAccountParam) GetClientAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientAccountId) {
		return nil, false
	}
	return o.ClientAccountId, true
}

// HasClientAccountId returns a boolean if a field has been set.
func (o *RegisterAccountParam) HasClientAccountId() bool {
	if o != nil && !IsNil(o.ClientAccountId) {
		return true
	}

	return false
}

// SetClientAccountId gets a reference to the given string and assigns it to the ClientAccountId field.
func (o *RegisterAccountParam) SetClientAccountId(v string) {
	o.ClientAccountId = &v
}

func (o RegisterAccountParam) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegisterAccountParam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.L2Key) {
		toSerialize["l2Key"] = o.L2Key
	}
	if !IsNil(o.L2KeyYCoordinate) {
		toSerialize["l2KeyYCoordinate"] = o.L2KeyYCoordinate
	}
	if !IsNil(o.ClientAccountId) {
		toSerialize["clientAccountId"] = o.ClientAccountId
	}
	return toSerialize, nil
}

type NullableRegisterAccountParam struct {
	value *RegisterAccountParam
	isSet bool
}

func (v NullableRegisterAccountParam) Get() *RegisterAccountParam {
	return v.value
}

func (v *NullableRegisterAccountParam) Set(val *RegisterAccountParam) {
	v.value = val
	v.isSet = true
}

func (v NullableRegisterAccountParam) IsSet() bool {
	return v.isSet
}

func (v *NullableRegisterAccountParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegisterAccountParam(val *RegisterAccountParam) *NullableRegisterAccountParam {
	return &NullableRegisterAccountParam{value: val, isSet: true}
}

func (v NullableRegisterAccountParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegisterAccountParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


