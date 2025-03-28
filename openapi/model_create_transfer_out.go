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

// checks if the CreateTransferOut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTransferOut{}

// CreateTransferOut 创建转出单-响应
type CreateTransferOut struct {
	// 转出单id
	TransferOutId *string `json:"transferOutId,omitempty"`
}

// NewCreateTransferOut instantiates a new CreateTransferOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTransferOut() *CreateTransferOut {
	this := CreateTransferOut{}
	return &this
}

// NewCreateTransferOutWithDefaults instantiates a new CreateTransferOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTransferOutWithDefaults() *CreateTransferOut {
	this := CreateTransferOut{}
	return &this
}

// GetTransferOutId returns the TransferOutId field value if set, zero value otherwise.
func (o *CreateTransferOut) GetTransferOutId() string {
	if o == nil || IsNil(o.TransferOutId) {
		var ret string
		return ret
	}
	return *o.TransferOutId
}

// GetTransferOutIdOk returns a tuple with the TransferOutId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTransferOut) GetTransferOutIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransferOutId) {
		return nil, false
	}
	return o.TransferOutId, true
}

// HasTransferOutId returns a boolean if a field has been set.
func (o *CreateTransferOut) HasTransferOutId() bool {
	if o != nil && !IsNil(o.TransferOutId) {
		return true
	}

	return false
}

// SetTransferOutId gets a reference to the given string and assigns it to the TransferOutId field.
func (o *CreateTransferOut) SetTransferOutId(v string) {
	o.TransferOutId = &v
}

func (o CreateTransferOut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTransferOut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TransferOutId) {
		toSerialize["transferOutId"] = o.TransferOutId
	}
	return toSerialize, nil
}

type NullableCreateTransferOut struct {
	value *CreateTransferOut
	isSet bool
}

func (v NullableCreateTransferOut) Get() *CreateTransferOut {
	return v.value
}

func (v *NullableCreateTransferOut) Set(val *CreateTransferOut) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTransferOut) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTransferOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTransferOut(val *CreateTransferOut) *NullableCreateTransferOut {
	return &NullableCreateTransferOut{value: val, isSet: true}
}

func (v NullableCreateTransferOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTransferOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


