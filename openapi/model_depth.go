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

// checks if the Depth type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Depth{}

// Depth 深度
type Depth struct {
	// 开始订单簿版本号
	StartVersion *string `json:"startVersion,omitempty"`
	// 结束订单簿版本号
	EndVersion *string `json:"endVersion,omitempty"`
	// 深度档位
	Level *int32 `json:"level,omitempty"`
	// 合约id
	ContractId *string `json:"contractId,omitempty"`
	// 合约名称
	ContractName *string `json:"contractName,omitempty"`
	// ask列表
	Asks []BookOrder `json:"asks,omitempty"`
	// bid列表
	Bids []BookOrder `json:"bids,omitempty"`
	// 盘口深度类型
	DepthType *string `json:"depthType,omitempty"`
}

// NewDepth instantiates a new Depth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepth() *Depth {
	this := Depth{}
	return &this
}

// NewDepthWithDefaults instantiates a new Depth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepthWithDefaults() *Depth {
	this := Depth{}
	return &this
}

// GetStartVersion returns the StartVersion field value if set, zero value otherwise.
func (o *Depth) GetStartVersion() string {
	if o == nil || IsNil(o.StartVersion) {
		var ret string
		return ret
	}
	return *o.StartVersion
}

// GetStartVersionOk returns a tuple with the StartVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Depth) GetStartVersionOk() (*string, bool) {
	if o == nil || IsNil(o.StartVersion) {
		return nil, false
	}
	return o.StartVersion, true
}

// HasStartVersion returns a boolean if a field has been set.
func (o *Depth) HasStartVersion() bool {
	if o != nil && !IsNil(o.StartVersion) {
		return true
	}

	return false
}

// SetStartVersion gets a reference to the given string and assigns it to the StartVersion field.
func (o *Depth) SetStartVersion(v string) {
	o.StartVersion = &v
}

// GetEndVersion returns the EndVersion field value if set, zero value otherwise.
func (o *Depth) GetEndVersion() string {
	if o == nil || IsNil(o.EndVersion) {
		var ret string
		return ret
	}
	return *o.EndVersion
}

// GetEndVersionOk returns a tuple with the EndVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Depth) GetEndVersionOk() (*string, bool) {
	if o == nil || IsNil(o.EndVersion) {
		return nil, false
	}
	return o.EndVersion, true
}

// HasEndVersion returns a boolean if a field has been set.
func (o *Depth) HasEndVersion() bool {
	if o != nil && !IsNil(o.EndVersion) {
		return true
	}

	return false
}

// SetEndVersion gets a reference to the given string and assigns it to the EndVersion field.
func (o *Depth) SetEndVersion(v string) {
	o.EndVersion = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *Depth) GetLevel() int32 {
	if o == nil || IsNil(o.Level) {
		var ret int32
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Depth) GetLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *Depth) HasLevel() bool {
	if o != nil && !IsNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given int32 and assigns it to the Level field.
func (o *Depth) SetLevel(v int32) {
	o.Level = &v
}

// GetContractId returns the ContractId field value if set, zero value otherwise.
func (o *Depth) GetContractId() string {
	if o == nil || IsNil(o.ContractId) {
		var ret string
		return ret
	}
	return *o.ContractId
}

// GetContractIdOk returns a tuple with the ContractId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Depth) GetContractIdOk() (*string, bool) {
	if o == nil || IsNil(o.ContractId) {
		return nil, false
	}
	return o.ContractId, true
}

// HasContractId returns a boolean if a field has been set.
func (o *Depth) HasContractId() bool {
	if o != nil && !IsNil(o.ContractId) {
		return true
	}

	return false
}

// SetContractId gets a reference to the given string and assigns it to the ContractId field.
func (o *Depth) SetContractId(v string) {
	o.ContractId = &v
}

// GetContractName returns the ContractName field value if set, zero value otherwise.
func (o *Depth) GetContractName() string {
	if o == nil || IsNil(o.ContractName) {
		var ret string
		return ret
	}
	return *o.ContractName
}

// GetContractNameOk returns a tuple with the ContractName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Depth) GetContractNameOk() (*string, bool) {
	if o == nil || IsNil(o.ContractName) {
		return nil, false
	}
	return o.ContractName, true
}

// HasContractName returns a boolean if a field has been set.
func (o *Depth) HasContractName() bool {
	if o != nil && !IsNil(o.ContractName) {
		return true
	}

	return false
}

// SetContractName gets a reference to the given string and assigns it to the ContractName field.
func (o *Depth) SetContractName(v string) {
	o.ContractName = &v
}

// GetAsks returns the Asks field value if set, zero value otherwise.
func (o *Depth) GetAsks() []BookOrder {
	if o == nil || IsNil(o.Asks) {
		var ret []BookOrder
		return ret
	}
	return o.Asks
}

// GetAsksOk returns a tuple with the Asks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Depth) GetAsksOk() ([]BookOrder, bool) {
	if o == nil || IsNil(o.Asks) {
		return nil, false
	}
	return o.Asks, true
}

// HasAsks returns a boolean if a field has been set.
func (o *Depth) HasAsks() bool {
	if o != nil && !IsNil(o.Asks) {
		return true
	}

	return false
}

// SetAsks gets a reference to the given []BookOrder and assigns it to the Asks field.
func (o *Depth) SetAsks(v []BookOrder) {
	o.Asks = v
}

// GetBids returns the Bids field value if set, zero value otherwise.
func (o *Depth) GetBids() []BookOrder {
	if o == nil || IsNil(o.Bids) {
		var ret []BookOrder
		return ret
	}
	return o.Bids
}

// GetBidsOk returns a tuple with the Bids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Depth) GetBidsOk() ([]BookOrder, bool) {
	if o == nil || IsNil(o.Bids) {
		return nil, false
	}
	return o.Bids, true
}

// HasBids returns a boolean if a field has been set.
func (o *Depth) HasBids() bool {
	if o != nil && !IsNil(o.Bids) {
		return true
	}

	return false
}

// SetBids gets a reference to the given []BookOrder and assigns it to the Bids field.
func (o *Depth) SetBids(v []BookOrder) {
	o.Bids = v
}

// GetDepthType returns the DepthType field value if set, zero value otherwise.
func (o *Depth) GetDepthType() string {
	if o == nil || IsNil(o.DepthType) {
		var ret string
		return ret
	}
	return *o.DepthType
}

// GetDepthTypeOk returns a tuple with the DepthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Depth) GetDepthTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DepthType) {
		return nil, false
	}
	return o.DepthType, true
}

// HasDepthType returns a boolean if a field has been set.
func (o *Depth) HasDepthType() bool {
	if o != nil && !IsNil(o.DepthType) {
		return true
	}

	return false
}

// SetDepthType gets a reference to the given string and assigns it to the DepthType field.
func (o *Depth) SetDepthType(v string) {
	o.DepthType = &v
}

func (o Depth) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Depth) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartVersion) {
		toSerialize["startVersion"] = o.StartVersion
	}
	if !IsNil(o.EndVersion) {
		toSerialize["endVersion"] = o.EndVersion
	}
	if !IsNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	if !IsNil(o.ContractId) {
		toSerialize["contractId"] = o.ContractId
	}
	if !IsNil(o.ContractName) {
		toSerialize["contractName"] = o.ContractName
	}
	if !IsNil(o.Asks) {
		toSerialize["asks"] = o.Asks
	}
	if !IsNil(o.Bids) {
		toSerialize["bids"] = o.Bids
	}
	if !IsNil(o.DepthType) {
		toSerialize["depthType"] = o.DepthType
	}
	return toSerialize, nil
}

type NullableDepth struct {
	value *Depth
	isSet bool
}

func (v NullableDepth) Get() *Depth {
	return v.value
}

func (v *NullableDepth) Set(val *Depth) {
	v.value = val
	v.isSet = true
}

func (v NullableDepth) IsSet() bool {
	return v.isSet
}

func (v *NullableDepth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepth(val *Depth) *NullableDepth {
	return &NullableDepth{value: val, isSet: true}
}

func (v NullableDepth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


