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

// checks if the PageDataWithdraw type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PageDataWithdraw{}

// PageDataWithdraw 通用翻页返回
type PageDataWithdraw struct {
	// 数据列表
	DataList []Withdraw `json:"dataList,omitempty"`
	// 获取下一页偏移。如果没有下一页数据，则为空串
	NextPageOffsetData *string `json:"nextPageOffsetData,omitempty"`
}

// NewPageDataWithdraw instantiates a new PageDataWithdraw object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageDataWithdraw() *PageDataWithdraw {
	this := PageDataWithdraw{}
	return &this
}

// NewPageDataWithdrawWithDefaults instantiates a new PageDataWithdraw object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageDataWithdrawWithDefaults() *PageDataWithdraw {
	this := PageDataWithdraw{}
	return &this
}

// GetDataList returns the DataList field value if set, zero value otherwise.
func (o *PageDataWithdraw) GetDataList() []Withdraw {
	if o == nil || IsNil(o.DataList) {
		var ret []Withdraw
		return ret
	}
	return o.DataList
}

// GetDataListOk returns a tuple with the DataList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageDataWithdraw) GetDataListOk() ([]Withdraw, bool) {
	if o == nil || IsNil(o.DataList) {
		return nil, false
	}
	return o.DataList, true
}

// HasDataList returns a boolean if a field has been set.
func (o *PageDataWithdraw) HasDataList() bool {
	if o != nil && !IsNil(o.DataList) {
		return true
	}

	return false
}

// SetDataList gets a reference to the given []Withdraw and assigns it to the DataList field.
func (o *PageDataWithdraw) SetDataList(v []Withdraw) {
	o.DataList = v
}

// GetNextPageOffsetData returns the NextPageOffsetData field value if set, zero value otherwise.
func (o *PageDataWithdraw) GetNextPageOffsetData() string {
	if o == nil || IsNil(o.NextPageOffsetData) {
		var ret string
		return ret
	}
	return *o.NextPageOffsetData
}

// GetNextPageOffsetDataOk returns a tuple with the NextPageOffsetData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageDataWithdraw) GetNextPageOffsetDataOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageOffsetData) {
		return nil, false
	}
	return o.NextPageOffsetData, true
}

// HasNextPageOffsetData returns a boolean if a field has been set.
func (o *PageDataWithdraw) HasNextPageOffsetData() bool {
	if o != nil && !IsNil(o.NextPageOffsetData) {
		return true
	}

	return false
}

// SetNextPageOffsetData gets a reference to the given string and assigns it to the NextPageOffsetData field.
func (o *PageDataWithdraw) SetNextPageOffsetData(v string) {
	o.NextPageOffsetData = &v
}

func (o PageDataWithdraw) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PageDataWithdraw) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DataList) {
		toSerialize["dataList"] = o.DataList
	}
	if !IsNil(o.NextPageOffsetData) {
		toSerialize["nextPageOffsetData"] = o.NextPageOffsetData
	}
	return toSerialize, nil
}

type NullablePageDataWithdraw struct {
	value *PageDataWithdraw
	isSet bool
}

func (v NullablePageDataWithdraw) Get() *PageDataWithdraw {
	return v.value
}

func (v *NullablePageDataWithdraw) Set(val *PageDataWithdraw) {
	v.value = val
	v.isSet = true
}

func (v NullablePageDataWithdraw) IsSet() bool {
	return v.isSet
}

func (v *NullablePageDataWithdraw) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageDataWithdraw(val *PageDataWithdraw) *NullablePageDataWithdraw {
	return &NullablePageDataWithdraw{value: val, isSet: true}
}

func (v NullablePageDataWithdraw) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageDataWithdraw) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


