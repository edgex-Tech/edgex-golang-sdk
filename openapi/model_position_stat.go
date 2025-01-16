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

// checks if the PositionStat type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PositionStat{}

// PositionStat 仓位累计统计信息
type PositionStat struct {
	// 累计开仓数量
	CumOpenSize *string `json:"cumOpenSize,omitempty"`
	// 累计开仓价值
	CumOpenValue *string `json:"cumOpenValue,omitempty"`
	// 累计开仓费用
	CumOpenFee *string `json:"cumOpenFee,omitempty"`
	// 累计平仓数量
	CumCloseSize *string `json:"cumCloseSize,omitempty"`
	// 累计平仓价值
	CumCloseValue *string `json:"cumCloseValue,omitempty"`
	// 累计平仓费用
	CumCloseFee *string `json:"cumCloseFee,omitempty"`
	// 累计已结算的资金费用
	CumFundingFee *string `json:"cumFundingFee,omitempty"`
	// 累计清算费用
	CumLiquidateFee *string `json:"cumLiquidateFee,omitempty"`
}

// NewPositionStat instantiates a new PositionStat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPositionStat() *PositionStat {
	this := PositionStat{}
	return &this
}

// NewPositionStatWithDefaults instantiates a new PositionStat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPositionStatWithDefaults() *PositionStat {
	this := PositionStat{}
	return &this
}

// GetCumOpenSize returns the CumOpenSize field value if set, zero value otherwise.
func (o *PositionStat) GetCumOpenSize() string {
	if o == nil || IsNil(o.CumOpenSize) {
		var ret string
		return ret
	}
	return *o.CumOpenSize
}

// GetCumOpenSizeOk returns a tuple with the CumOpenSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionStat) GetCumOpenSizeOk() (*string, bool) {
	if o == nil || IsNil(o.CumOpenSize) {
		return nil, false
	}
	return o.CumOpenSize, true
}

// HasCumOpenSize returns a boolean if a field has been set.
func (o *PositionStat) HasCumOpenSize() bool {
	if o != nil && !IsNil(o.CumOpenSize) {
		return true
	}

	return false
}

// SetCumOpenSize gets a reference to the given string and assigns it to the CumOpenSize field.
func (o *PositionStat) SetCumOpenSize(v string) {
	o.CumOpenSize = &v
}

// GetCumOpenValue returns the CumOpenValue field value if set, zero value otherwise.
func (o *PositionStat) GetCumOpenValue() string {
	if o == nil || IsNil(o.CumOpenValue) {
		var ret string
		return ret
	}
	return *o.CumOpenValue
}

// GetCumOpenValueOk returns a tuple with the CumOpenValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionStat) GetCumOpenValueOk() (*string, bool) {
	if o == nil || IsNil(o.CumOpenValue) {
		return nil, false
	}
	return o.CumOpenValue, true
}

// HasCumOpenValue returns a boolean if a field has been set.
func (o *PositionStat) HasCumOpenValue() bool {
	if o != nil && !IsNil(o.CumOpenValue) {
		return true
	}

	return false
}

// SetCumOpenValue gets a reference to the given string and assigns it to the CumOpenValue field.
func (o *PositionStat) SetCumOpenValue(v string) {
	o.CumOpenValue = &v
}

// GetCumOpenFee returns the CumOpenFee field value if set, zero value otherwise.
func (o *PositionStat) GetCumOpenFee() string {
	if o == nil || IsNil(o.CumOpenFee) {
		var ret string
		return ret
	}
	return *o.CumOpenFee
}

// GetCumOpenFeeOk returns a tuple with the CumOpenFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionStat) GetCumOpenFeeOk() (*string, bool) {
	if o == nil || IsNil(o.CumOpenFee) {
		return nil, false
	}
	return o.CumOpenFee, true
}

// HasCumOpenFee returns a boolean if a field has been set.
func (o *PositionStat) HasCumOpenFee() bool {
	if o != nil && !IsNil(o.CumOpenFee) {
		return true
	}

	return false
}

// SetCumOpenFee gets a reference to the given string and assigns it to the CumOpenFee field.
func (o *PositionStat) SetCumOpenFee(v string) {
	o.CumOpenFee = &v
}

// GetCumCloseSize returns the CumCloseSize field value if set, zero value otherwise.
func (o *PositionStat) GetCumCloseSize() string {
	if o == nil || IsNil(o.CumCloseSize) {
		var ret string
		return ret
	}
	return *o.CumCloseSize
}

// GetCumCloseSizeOk returns a tuple with the CumCloseSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionStat) GetCumCloseSizeOk() (*string, bool) {
	if o == nil || IsNil(o.CumCloseSize) {
		return nil, false
	}
	return o.CumCloseSize, true
}

// HasCumCloseSize returns a boolean if a field has been set.
func (o *PositionStat) HasCumCloseSize() bool {
	if o != nil && !IsNil(o.CumCloseSize) {
		return true
	}

	return false
}

// SetCumCloseSize gets a reference to the given string and assigns it to the CumCloseSize field.
func (o *PositionStat) SetCumCloseSize(v string) {
	o.CumCloseSize = &v
}

// GetCumCloseValue returns the CumCloseValue field value if set, zero value otherwise.
func (o *PositionStat) GetCumCloseValue() string {
	if o == nil || IsNil(o.CumCloseValue) {
		var ret string
		return ret
	}
	return *o.CumCloseValue
}

// GetCumCloseValueOk returns a tuple with the CumCloseValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionStat) GetCumCloseValueOk() (*string, bool) {
	if o == nil || IsNil(o.CumCloseValue) {
		return nil, false
	}
	return o.CumCloseValue, true
}

// HasCumCloseValue returns a boolean if a field has been set.
func (o *PositionStat) HasCumCloseValue() bool {
	if o != nil && !IsNil(o.CumCloseValue) {
		return true
	}

	return false
}

// SetCumCloseValue gets a reference to the given string and assigns it to the CumCloseValue field.
func (o *PositionStat) SetCumCloseValue(v string) {
	o.CumCloseValue = &v
}

// GetCumCloseFee returns the CumCloseFee field value if set, zero value otherwise.
func (o *PositionStat) GetCumCloseFee() string {
	if o == nil || IsNil(o.CumCloseFee) {
		var ret string
		return ret
	}
	return *o.CumCloseFee
}

// GetCumCloseFeeOk returns a tuple with the CumCloseFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionStat) GetCumCloseFeeOk() (*string, bool) {
	if o == nil || IsNil(o.CumCloseFee) {
		return nil, false
	}
	return o.CumCloseFee, true
}

// HasCumCloseFee returns a boolean if a field has been set.
func (o *PositionStat) HasCumCloseFee() bool {
	if o != nil && !IsNil(o.CumCloseFee) {
		return true
	}

	return false
}

// SetCumCloseFee gets a reference to the given string and assigns it to the CumCloseFee field.
func (o *PositionStat) SetCumCloseFee(v string) {
	o.CumCloseFee = &v
}

// GetCumFundingFee returns the CumFundingFee field value if set, zero value otherwise.
func (o *PositionStat) GetCumFundingFee() string {
	if o == nil || IsNil(o.CumFundingFee) {
		var ret string
		return ret
	}
	return *o.CumFundingFee
}

// GetCumFundingFeeOk returns a tuple with the CumFundingFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionStat) GetCumFundingFeeOk() (*string, bool) {
	if o == nil || IsNil(o.CumFundingFee) {
		return nil, false
	}
	return o.CumFundingFee, true
}

// HasCumFundingFee returns a boolean if a field has been set.
func (o *PositionStat) HasCumFundingFee() bool {
	if o != nil && !IsNil(o.CumFundingFee) {
		return true
	}

	return false
}

// SetCumFundingFee gets a reference to the given string and assigns it to the CumFundingFee field.
func (o *PositionStat) SetCumFundingFee(v string) {
	o.CumFundingFee = &v
}

// GetCumLiquidateFee returns the CumLiquidateFee field value if set, zero value otherwise.
func (o *PositionStat) GetCumLiquidateFee() string {
	if o == nil || IsNil(o.CumLiquidateFee) {
		var ret string
		return ret
	}
	return *o.CumLiquidateFee
}

// GetCumLiquidateFeeOk returns a tuple with the CumLiquidateFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionStat) GetCumLiquidateFeeOk() (*string, bool) {
	if o == nil || IsNil(o.CumLiquidateFee) {
		return nil, false
	}
	return o.CumLiquidateFee, true
}

// HasCumLiquidateFee returns a boolean if a field has been set.
func (o *PositionStat) HasCumLiquidateFee() bool {
	if o != nil && !IsNil(o.CumLiquidateFee) {
		return true
	}

	return false
}

// SetCumLiquidateFee gets a reference to the given string and assigns it to the CumLiquidateFee field.
func (o *PositionStat) SetCumLiquidateFee(v string) {
	o.CumLiquidateFee = &v
}

func (o PositionStat) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PositionStat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CumOpenSize) {
		toSerialize["cumOpenSize"] = o.CumOpenSize
	}
	if !IsNil(o.CumOpenValue) {
		toSerialize["cumOpenValue"] = o.CumOpenValue
	}
	if !IsNil(o.CumOpenFee) {
		toSerialize["cumOpenFee"] = o.CumOpenFee
	}
	if !IsNil(o.CumCloseSize) {
		toSerialize["cumCloseSize"] = o.CumCloseSize
	}
	if !IsNil(o.CumCloseValue) {
		toSerialize["cumCloseValue"] = o.CumCloseValue
	}
	if !IsNil(o.CumCloseFee) {
		toSerialize["cumCloseFee"] = o.CumCloseFee
	}
	if !IsNil(o.CumFundingFee) {
		toSerialize["cumFundingFee"] = o.CumFundingFee
	}
	if !IsNil(o.CumLiquidateFee) {
		toSerialize["cumLiquidateFee"] = o.CumLiquidateFee
	}
	return toSerialize, nil
}

type NullablePositionStat struct {
	value *PositionStat
	isSet bool
}

func (v NullablePositionStat) Get() *PositionStat {
	return v.value
}

func (v *NullablePositionStat) Set(val *PositionStat) {
	v.value = val
	v.isSet = true
}

func (v NullablePositionStat) IsSet() bool {
	return v.isSet
}

func (v *NullablePositionStat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePositionStat(val *PositionStat) *NullablePositionStat {
	return &NullablePositionStat{value: val, isSet: true}
}

func (v NullablePositionStat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePositionStat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


