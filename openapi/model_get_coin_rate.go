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

// checks if the GetCoinRate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCoinRate{}

// GetCoinRate 查询币种对USDC/USDT的汇率-响应
type GetCoinRate struct {
	Rate *string `json:"rate,omitempty"`
}

// NewGetCoinRate instantiates a new GetCoinRate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCoinRate() *GetCoinRate {
	this := GetCoinRate{}
	return &this
}

// NewGetCoinRateWithDefaults instantiates a new GetCoinRate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCoinRateWithDefaults() *GetCoinRate {
	this := GetCoinRate{}
	return &this
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *GetCoinRate) GetRate() string {
	if o == nil || IsNil(o.Rate) {
		var ret string
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCoinRate) GetRateOk() (*string, bool) {
	if o == nil || IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *GetCoinRate) HasRate() bool {
	if o != nil && !IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given string and assigns it to the Rate field.
func (o *GetCoinRate) SetRate(v string) {
	o.Rate = &v
}

func (o GetCoinRate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCoinRate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rate) {
		toSerialize["rate"] = o.Rate
	}
	return toSerialize, nil
}

type NullableGetCoinRate struct {
	value *GetCoinRate
	isSet bool
}

func (v NullableGetCoinRate) Get() *GetCoinRate {
	return v.value
}

func (v *NullableGetCoinRate) Set(val *GetCoinRate) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCoinRate) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCoinRate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCoinRate(val *GetCoinRate) *NullableGetCoinRate {
	return &NullableGetCoinRate{value: val, isSet: true}
}

func (v NullableGetCoinRate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCoinRate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


