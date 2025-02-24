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

// checks if the GetTickerSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTickerSummary{}

// GetTickerSummary 获取行情汇总响应
type GetTickerSummary struct {
	TickerSummary *TickerSummary `json:"tickerSummary,omitempty"`
}

// NewGetTickerSummary instantiates a new GetTickerSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTickerSummary() *GetTickerSummary {
	this := GetTickerSummary{}
	return &this
}

// NewGetTickerSummaryWithDefaults instantiates a new GetTickerSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTickerSummaryWithDefaults() *GetTickerSummary {
	this := GetTickerSummary{}
	return &this
}

// GetTickerSummary returns the TickerSummary field value if set, zero value otherwise.
func (o *GetTickerSummary) GetTickerSummary() TickerSummary {
	if o == nil || IsNil(o.TickerSummary) {
		var ret TickerSummary
		return ret
	}
	return *o.TickerSummary
}

// GetTickerSummaryOk returns a tuple with the TickerSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTickerSummary) GetTickerSummaryOk() (*TickerSummary, bool) {
	if o == nil || IsNil(o.TickerSummary) {
		return nil, false
	}
	return o.TickerSummary, true
}

// HasTickerSummary returns a boolean if a field has been set.
func (o *GetTickerSummary) HasTickerSummary() bool {
	if o != nil && !IsNil(o.TickerSummary) {
		return true
	}

	return false
}

// SetTickerSummary gets a reference to the given TickerSummary and assigns it to the TickerSummary field.
func (o *GetTickerSummary) SetTickerSummary(v TickerSummary) {
	o.TickerSummary = &v
}

func (o GetTickerSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTickerSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TickerSummary) {
		toSerialize["tickerSummary"] = o.TickerSummary
	}
	return toSerialize, nil
}

type NullableGetTickerSummary struct {
	value *GetTickerSummary
	isSet bool
}

func (v NullableGetTickerSummary) Get() *GetTickerSummary {
	return v.value
}

func (v *NullableGetTickerSummary) Set(val *GetTickerSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTickerSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTickerSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTickerSummary(val *GetTickerSummary) *NullableGetTickerSummary {
	return &NullableGetTickerSummary{value: val, isSet: true}
}

func (v NullableGetTickerSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTickerSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


