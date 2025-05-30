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

// checks if the ReadSiteMessageParam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReadSiteMessageParam{}

// ReadSiteMessageParam 标识阅读站内信消息-请求
type ReadSiteMessageParam struct {
	// 消息Id. 可传多个。不传全部已读
	MessageIdList []string `json:"messageIdList,omitempty"`
}

// NewReadSiteMessageParam instantiates a new ReadSiteMessageParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadSiteMessageParam() *ReadSiteMessageParam {
	this := ReadSiteMessageParam{}
	return &this
}

// NewReadSiteMessageParamWithDefaults instantiates a new ReadSiteMessageParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadSiteMessageParamWithDefaults() *ReadSiteMessageParam {
	this := ReadSiteMessageParam{}
	return &this
}

// GetMessageIdList returns the MessageIdList field value if set, zero value otherwise.
func (o *ReadSiteMessageParam) GetMessageIdList() []string {
	if o == nil || IsNil(o.MessageIdList) {
		var ret []string
		return ret
	}
	return o.MessageIdList
}

// GetMessageIdListOk returns a tuple with the MessageIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadSiteMessageParam) GetMessageIdListOk() ([]string, bool) {
	if o == nil || IsNil(o.MessageIdList) {
		return nil, false
	}
	return o.MessageIdList, true
}

// HasMessageIdList returns a boolean if a field has been set.
func (o *ReadSiteMessageParam) HasMessageIdList() bool {
	if o != nil && !IsNil(o.MessageIdList) {
		return true
	}

	return false
}

// SetMessageIdList gets a reference to the given []string and assigns it to the MessageIdList field.
func (o *ReadSiteMessageParam) SetMessageIdList(v []string) {
	o.MessageIdList = v
}

func (o ReadSiteMessageParam) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadSiteMessageParam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MessageIdList) {
		toSerialize["messageIdList"] = o.MessageIdList
	}
	return toSerialize, nil
}

type NullableReadSiteMessageParam struct {
	value *ReadSiteMessageParam
	isSet bool
}

func (v NullableReadSiteMessageParam) Get() *ReadSiteMessageParam {
	return v.value
}

func (v *NullableReadSiteMessageParam) Set(val *ReadSiteMessageParam) {
	v.value = val
	v.isSet = true
}

func (v NullableReadSiteMessageParam) IsSet() bool {
	return v.isSet
}

func (v *NullableReadSiteMessageParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadSiteMessageParam(val *ReadSiteMessageParam) *NullableReadSiteMessageParam {
	return &NullableReadSiteMessageParam{value: val, isSet: true}
}

func (v NullableReadSiteMessageParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadSiteMessageParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


