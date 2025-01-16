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

// checks if the PageDataTransferOut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PageDataTransferOut{}

// PageDataTransferOut 通用翻页返回
type PageDataTransferOut struct {
	// 数据列表
	DataList []TransferOut `json:"dataList,omitempty"`
	// 获取下一页偏移。如果没有下一页数据，则为空串
	NextPageOffsetData *string `json:"nextPageOffsetData,omitempty"`
}

// NewPageDataTransferOut instantiates a new PageDataTransferOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageDataTransferOut() *PageDataTransferOut {
	this := PageDataTransferOut{}
	return &this
}

// NewPageDataTransferOutWithDefaults instantiates a new PageDataTransferOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageDataTransferOutWithDefaults() *PageDataTransferOut {
	this := PageDataTransferOut{}
	return &this
}

// GetDataList returns the DataList field value if set, zero value otherwise.
func (o *PageDataTransferOut) GetDataList() []TransferOut {
	if o == nil || IsNil(o.DataList) {
		var ret []TransferOut
		return ret
	}
	return o.DataList
}

// GetDataListOk returns a tuple with the DataList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageDataTransferOut) GetDataListOk() ([]TransferOut, bool) {
	if o == nil || IsNil(o.DataList) {
		return nil, false
	}
	return o.DataList, true
}

// HasDataList returns a boolean if a field has been set.
func (o *PageDataTransferOut) HasDataList() bool {
	if o != nil && !IsNil(o.DataList) {
		return true
	}

	return false
}

// SetDataList gets a reference to the given []TransferOut and assigns it to the DataList field.
func (o *PageDataTransferOut) SetDataList(v []TransferOut) {
	o.DataList = v
}

// GetNextPageOffsetData returns the NextPageOffsetData field value if set, zero value otherwise.
func (o *PageDataTransferOut) GetNextPageOffsetData() string {
	if o == nil || IsNil(o.NextPageOffsetData) {
		var ret string
		return ret
	}
	return *o.NextPageOffsetData
}

// GetNextPageOffsetDataOk returns a tuple with the NextPageOffsetData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageDataTransferOut) GetNextPageOffsetDataOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageOffsetData) {
		return nil, false
	}
	return o.NextPageOffsetData, true
}

// HasNextPageOffsetData returns a boolean if a field has been set.
func (o *PageDataTransferOut) HasNextPageOffsetData() bool {
	if o != nil && !IsNil(o.NextPageOffsetData) {
		return true
	}

	return false
}

// SetNextPageOffsetData gets a reference to the given string and assigns it to the NextPageOffsetData field.
func (o *PageDataTransferOut) SetNextPageOffsetData(v string) {
	o.NextPageOffsetData = &v
}

func (o PageDataTransferOut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PageDataTransferOut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DataList) {
		toSerialize["dataList"] = o.DataList
	}
	if !IsNil(o.NextPageOffsetData) {
		toSerialize["nextPageOffsetData"] = o.NextPageOffsetData
	}
	return toSerialize, nil
}

type NullablePageDataTransferOut struct {
	value *PageDataTransferOut
	isSet bool
}

func (v NullablePageDataTransferOut) Get() *PageDataTransferOut {
	return v.value
}

func (v *NullablePageDataTransferOut) Set(val *PageDataTransferOut) {
	v.value = val
	v.isSet = true
}

func (v NullablePageDataTransferOut) IsSet() bool {
	return v.isSet
}

func (v *NullablePageDataTransferOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageDataTransferOut(val *PageDataTransferOut) *NullablePageDataTransferOut {
	return &NullablePageDataTransferOut{value: val, isSet: true}
}

func (v NullablePageDataTransferOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageDataTransferOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


