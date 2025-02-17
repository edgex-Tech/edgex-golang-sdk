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

// checks if the Get1inchQuote type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Get1inchQuote{}

// Get1inchQuote 获取1inch币对兑换汇率-响应
type Get1inchQuote struct {
	// 返回的汇率，带精度
	ToAmount *string `json:"toAmount,omitempty"`
	// gas
	Gas *string `json:"gas,omitempty"`
}

// NewGet1inchQuote instantiates a new Get1inchQuote object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGet1inchQuote() *Get1inchQuote {
	this := Get1inchQuote{}
	return &this
}

// NewGet1inchQuoteWithDefaults instantiates a new Get1inchQuote object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGet1inchQuoteWithDefaults() *Get1inchQuote {
	this := Get1inchQuote{}
	return &this
}

// GetToAmount returns the ToAmount field value if set, zero value otherwise.
func (o *Get1inchQuote) GetToAmount() string {
	if o == nil || IsNil(o.ToAmount) {
		var ret string
		return ret
	}
	return *o.ToAmount
}

// GetToAmountOk returns a tuple with the ToAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Get1inchQuote) GetToAmountOk() (*string, bool) {
	if o == nil || IsNil(o.ToAmount) {
		return nil, false
	}
	return o.ToAmount, true
}

// HasToAmount returns a boolean if a field has been set.
func (o *Get1inchQuote) HasToAmount() bool {
	if o != nil && !IsNil(o.ToAmount) {
		return true
	}

	return false
}

// SetToAmount gets a reference to the given string and assigns it to the ToAmount field.
func (o *Get1inchQuote) SetToAmount(v string) {
	o.ToAmount = &v
}

// GetGas returns the Gas field value if set, zero value otherwise.
func (o *Get1inchQuote) GetGas() string {
	if o == nil || IsNil(o.Gas) {
		var ret string
		return ret
	}
	return *o.Gas
}

// GetGasOk returns a tuple with the Gas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Get1inchQuote) GetGasOk() (*string, bool) {
	if o == nil || IsNil(o.Gas) {
		return nil, false
	}
	return o.Gas, true
}

// HasGas returns a boolean if a field has been set.
func (o *Get1inchQuote) HasGas() bool {
	if o != nil && !IsNil(o.Gas) {
		return true
	}

	return false
}

// SetGas gets a reference to the given string and assigns it to the Gas field.
func (o *Get1inchQuote) SetGas(v string) {
	o.Gas = &v
}

func (o Get1inchQuote) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Get1inchQuote) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ToAmount) {
		toSerialize["toAmount"] = o.ToAmount
	}
	if !IsNil(o.Gas) {
		toSerialize["gas"] = o.Gas
	}
	return toSerialize, nil
}

type NullableGet1inchQuote struct {
	value *Get1inchQuote
	isSet bool
}

func (v NullableGet1inchQuote) Get() *Get1inchQuote {
	return v.value
}

func (v *NullableGet1inchQuote) Set(val *Get1inchQuote) {
	v.value = val
	v.isSet = true
}

func (v NullableGet1inchQuote) IsSet() bool {
	return v.isSet
}

func (v *NullableGet1inchQuote) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGet1inchQuote(val *Get1inchQuote) *NullableGet1inchQuote {
	return &NullableGet1inchQuote{value: val, isSet: true}
}

func (v NullableGet1inchQuote) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGet1inchQuote) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


