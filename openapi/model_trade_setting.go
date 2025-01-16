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

// checks if the TradeSetting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TradeSetting{}

// TradeSetting 交易设置. 交易设置计算优先级：账户合约交易设置 -> 账户默认交易设置 -> 合约配置交易设置. 注意：is_set_fee_rate 和 is_set_fee_discount 只能有一个为 true
type TradeSetting struct {
	// 是否设置费率(具体值)
	IsSetFeeRate *bool `json:"isSetFeeRate,omitempty"`
	// taker费率，取值范围 [0, 1)，仅当 is_set_fee_rate=true 时有效
	TakerFeeRate *string `json:"takerFeeRate,omitempty"`
	// maker费率，取值范围 [0, 1)，仅当 is_set_fee_rate=true 时有效
	MakerFeeRate *string `json:"makerFeeRate,omitempty"`
	// 是否设置费率折扣
	IsSetFeeDiscount *bool `json:"isSetFeeDiscount,omitempty"`
	// taker费率折扣，取值范围 [0, 1)，仅当 is_set_fee_discount=true 时有效
	TakerFeeDiscount *string `json:"takerFeeDiscount,omitempty"`
	// maker费率折扣，取值范围 [0, 1)，仅当 is_set_fee_discount=true 时有效
	MakerFeeDiscount *string `json:"makerFeeDiscount,omitempty"`
	// 是否设置最大交易杠杆
	IsSetMaxLeverage *bool `json:"isSetMaxLeverage,omitempty"`
	// 最大交易杠杆
	MaxLeverage *string `json:"maxLeverage,omitempty"`
}

// NewTradeSetting instantiates a new TradeSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTradeSetting() *TradeSetting {
	this := TradeSetting{}
	return &this
}

// NewTradeSettingWithDefaults instantiates a new TradeSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTradeSettingWithDefaults() *TradeSetting {
	this := TradeSetting{}
	return &this
}

// GetIsSetFeeRate returns the IsSetFeeRate field value if set, zero value otherwise.
func (o *TradeSetting) GetIsSetFeeRate() bool {
	if o == nil || IsNil(o.IsSetFeeRate) {
		var ret bool
		return ret
	}
	return *o.IsSetFeeRate
}

// GetIsSetFeeRateOk returns a tuple with the IsSetFeeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeSetting) GetIsSetFeeRateOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSetFeeRate) {
		return nil, false
	}
	return o.IsSetFeeRate, true
}

// HasIsSetFeeRate returns a boolean if a field has been set.
func (o *TradeSetting) HasIsSetFeeRate() bool {
	if o != nil && !IsNil(o.IsSetFeeRate) {
		return true
	}

	return false
}

// SetIsSetFeeRate gets a reference to the given bool and assigns it to the IsSetFeeRate field.
func (o *TradeSetting) SetIsSetFeeRate(v bool) {
	o.IsSetFeeRate = &v
}

// GetTakerFeeRate returns the TakerFeeRate field value if set, zero value otherwise.
func (o *TradeSetting) GetTakerFeeRate() string {
	if o == nil || IsNil(o.TakerFeeRate) {
		var ret string
		return ret
	}
	return *o.TakerFeeRate
}

// GetTakerFeeRateOk returns a tuple with the TakerFeeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeSetting) GetTakerFeeRateOk() (*string, bool) {
	if o == nil || IsNil(o.TakerFeeRate) {
		return nil, false
	}
	return o.TakerFeeRate, true
}

// HasTakerFeeRate returns a boolean if a field has been set.
func (o *TradeSetting) HasTakerFeeRate() bool {
	if o != nil && !IsNil(o.TakerFeeRate) {
		return true
	}

	return false
}

// SetTakerFeeRate gets a reference to the given string and assigns it to the TakerFeeRate field.
func (o *TradeSetting) SetTakerFeeRate(v string) {
	o.TakerFeeRate = &v
}

// GetMakerFeeRate returns the MakerFeeRate field value if set, zero value otherwise.
func (o *TradeSetting) GetMakerFeeRate() string {
	if o == nil || IsNil(o.MakerFeeRate) {
		var ret string
		return ret
	}
	return *o.MakerFeeRate
}

// GetMakerFeeRateOk returns a tuple with the MakerFeeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeSetting) GetMakerFeeRateOk() (*string, bool) {
	if o == nil || IsNil(o.MakerFeeRate) {
		return nil, false
	}
	return o.MakerFeeRate, true
}

// HasMakerFeeRate returns a boolean if a field has been set.
func (o *TradeSetting) HasMakerFeeRate() bool {
	if o != nil && !IsNil(o.MakerFeeRate) {
		return true
	}

	return false
}

// SetMakerFeeRate gets a reference to the given string and assigns it to the MakerFeeRate field.
func (o *TradeSetting) SetMakerFeeRate(v string) {
	o.MakerFeeRate = &v
}

// GetIsSetFeeDiscount returns the IsSetFeeDiscount field value if set, zero value otherwise.
func (o *TradeSetting) GetIsSetFeeDiscount() bool {
	if o == nil || IsNil(o.IsSetFeeDiscount) {
		var ret bool
		return ret
	}
	return *o.IsSetFeeDiscount
}

// GetIsSetFeeDiscountOk returns a tuple with the IsSetFeeDiscount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeSetting) GetIsSetFeeDiscountOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSetFeeDiscount) {
		return nil, false
	}
	return o.IsSetFeeDiscount, true
}

// HasIsSetFeeDiscount returns a boolean if a field has been set.
func (o *TradeSetting) HasIsSetFeeDiscount() bool {
	if o != nil && !IsNil(o.IsSetFeeDiscount) {
		return true
	}

	return false
}

// SetIsSetFeeDiscount gets a reference to the given bool and assigns it to the IsSetFeeDiscount field.
func (o *TradeSetting) SetIsSetFeeDiscount(v bool) {
	o.IsSetFeeDiscount = &v
}

// GetTakerFeeDiscount returns the TakerFeeDiscount field value if set, zero value otherwise.
func (o *TradeSetting) GetTakerFeeDiscount() string {
	if o == nil || IsNil(o.TakerFeeDiscount) {
		var ret string
		return ret
	}
	return *o.TakerFeeDiscount
}

// GetTakerFeeDiscountOk returns a tuple with the TakerFeeDiscount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeSetting) GetTakerFeeDiscountOk() (*string, bool) {
	if o == nil || IsNil(o.TakerFeeDiscount) {
		return nil, false
	}
	return o.TakerFeeDiscount, true
}

// HasTakerFeeDiscount returns a boolean if a field has been set.
func (o *TradeSetting) HasTakerFeeDiscount() bool {
	if o != nil && !IsNil(o.TakerFeeDiscount) {
		return true
	}

	return false
}

// SetTakerFeeDiscount gets a reference to the given string and assigns it to the TakerFeeDiscount field.
func (o *TradeSetting) SetTakerFeeDiscount(v string) {
	o.TakerFeeDiscount = &v
}

// GetMakerFeeDiscount returns the MakerFeeDiscount field value if set, zero value otherwise.
func (o *TradeSetting) GetMakerFeeDiscount() string {
	if o == nil || IsNil(o.MakerFeeDiscount) {
		var ret string
		return ret
	}
	return *o.MakerFeeDiscount
}

// GetMakerFeeDiscountOk returns a tuple with the MakerFeeDiscount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeSetting) GetMakerFeeDiscountOk() (*string, bool) {
	if o == nil || IsNil(o.MakerFeeDiscount) {
		return nil, false
	}
	return o.MakerFeeDiscount, true
}

// HasMakerFeeDiscount returns a boolean if a field has been set.
func (o *TradeSetting) HasMakerFeeDiscount() bool {
	if o != nil && !IsNil(o.MakerFeeDiscount) {
		return true
	}

	return false
}

// SetMakerFeeDiscount gets a reference to the given string and assigns it to the MakerFeeDiscount field.
func (o *TradeSetting) SetMakerFeeDiscount(v string) {
	o.MakerFeeDiscount = &v
}

// GetIsSetMaxLeverage returns the IsSetMaxLeverage field value if set, zero value otherwise.
func (o *TradeSetting) GetIsSetMaxLeverage() bool {
	if o == nil || IsNil(o.IsSetMaxLeverage) {
		var ret bool
		return ret
	}
	return *o.IsSetMaxLeverage
}

// GetIsSetMaxLeverageOk returns a tuple with the IsSetMaxLeverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeSetting) GetIsSetMaxLeverageOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSetMaxLeverage) {
		return nil, false
	}
	return o.IsSetMaxLeverage, true
}

// HasIsSetMaxLeverage returns a boolean if a field has been set.
func (o *TradeSetting) HasIsSetMaxLeverage() bool {
	if o != nil && !IsNil(o.IsSetMaxLeverage) {
		return true
	}

	return false
}

// SetIsSetMaxLeverage gets a reference to the given bool and assigns it to the IsSetMaxLeverage field.
func (o *TradeSetting) SetIsSetMaxLeverage(v bool) {
	o.IsSetMaxLeverage = &v
}

// GetMaxLeverage returns the MaxLeverage field value if set, zero value otherwise.
func (o *TradeSetting) GetMaxLeverage() string {
	if o == nil || IsNil(o.MaxLeverage) {
		var ret string
		return ret
	}
	return *o.MaxLeverage
}

// GetMaxLeverageOk returns a tuple with the MaxLeverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeSetting) GetMaxLeverageOk() (*string, bool) {
	if o == nil || IsNil(o.MaxLeverage) {
		return nil, false
	}
	return o.MaxLeverage, true
}

// HasMaxLeverage returns a boolean if a field has been set.
func (o *TradeSetting) HasMaxLeverage() bool {
	if o != nil && !IsNil(o.MaxLeverage) {
		return true
	}

	return false
}

// SetMaxLeverage gets a reference to the given string and assigns it to the MaxLeverage field.
func (o *TradeSetting) SetMaxLeverage(v string) {
	o.MaxLeverage = &v
}

func (o TradeSetting) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TradeSetting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsSetFeeRate) {
		toSerialize["isSetFeeRate"] = o.IsSetFeeRate
	}
	if !IsNil(o.TakerFeeRate) {
		toSerialize["takerFeeRate"] = o.TakerFeeRate
	}
	if !IsNil(o.MakerFeeRate) {
		toSerialize["makerFeeRate"] = o.MakerFeeRate
	}
	if !IsNil(o.IsSetFeeDiscount) {
		toSerialize["isSetFeeDiscount"] = o.IsSetFeeDiscount
	}
	if !IsNil(o.TakerFeeDiscount) {
		toSerialize["takerFeeDiscount"] = o.TakerFeeDiscount
	}
	if !IsNil(o.MakerFeeDiscount) {
		toSerialize["makerFeeDiscount"] = o.MakerFeeDiscount
	}
	if !IsNil(o.IsSetMaxLeverage) {
		toSerialize["isSetMaxLeverage"] = o.IsSetMaxLeverage
	}
	if !IsNil(o.MaxLeverage) {
		toSerialize["maxLeverage"] = o.MaxLeverage
	}
	return toSerialize, nil
}

type NullableTradeSetting struct {
	value *TradeSetting
	isSet bool
}

func (v NullableTradeSetting) Get() *TradeSetting {
	return v.value
}

func (v *NullableTradeSetting) Set(val *TradeSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableTradeSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableTradeSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTradeSetting(val *TradeSetting) *NullableTradeSetting {
	return &NullableTradeSetting{value: val, isSet: true}
}

func (v NullableTradeSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTradeSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


