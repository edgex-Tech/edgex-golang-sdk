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

// checks if the PositionAsset type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PositionAsset{}

// PositionAsset 仓位相关的资产信息
type PositionAsset struct {
	// 所属用户id
	UserId *string `json:"userId,omitempty"`
	// 所属账户id
	AccountId *string `json:"accountId,omitempty"`
	// 所属抵押品币种id
	CoinId *string `json:"coinId,omitempty"`
	// 所属合约id
	ContractId *string `json:"contractId,omitempty"`
	// 仓位价值，(正数为多仓，负数为空仓)
	PositionValue *string `json:"positionValue,omitempty"`
	// 当前合约仓位最大可开杠杆
	MaxLeverage *string `json:"maxLeverage,omitempty"`
	// 仓位初始保证金
	InitialMarginRequirement *string `json:"initialMarginRequirement,omitempty"`
	// 结合风险档位计算出的starkEx风险率。本质类似于维持保证金率，只是精度不同
	StarkExRiskRate *string `json:"starkExRiskRate,omitempty"`
	// starkEx 风险金额，本质类似于维持保证金，只是精度不同。
	StarkExRiskValue *string `json:"starkExRiskValue,omitempty"`
	// 平均入场价格
	AvgEntryPrice *string `json:"avgEntryPrice,omitempty"`
	// 清算价(强平价)。即当预言机价格到达此价格，必定触发清算
	LiquidatePrice *string `json:"liquidatePrice,omitempty"`
	// 破产价。即当预言机价格到达此价格，账户总价值小于0
	BankruptPrice *string `json:"bankruptPrice,omitempty"`
	// 最差平仓价。即平仓成交价格不能劣于这个价格
	WorstClosePrice *string `json:"worstClosePrice,omitempty"`
	// 仓位未实现盈亏
	UnrealizePnl *string `json:"unrealizePnl,omitempty"`
	// 仓位term已实现盈亏
	TermRealizePnl *string `json:"termRealizePnl,omitempty"`
	// 仓位total已实现盈亏
	TotalRealizePnl *string `json:"totalRealizePnl,omitempty"`
}

// NewPositionAsset instantiates a new PositionAsset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPositionAsset() *PositionAsset {
	this := PositionAsset{}
	return &this
}

// NewPositionAssetWithDefaults instantiates a new PositionAsset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPositionAssetWithDefaults() *PositionAsset {
	this := PositionAsset{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *PositionAsset) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionAsset) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *PositionAsset) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *PositionAsset) SetUserId(v string) {
	o.UserId = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *PositionAsset) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionAsset) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *PositionAsset) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *PositionAsset) SetAccountId(v string) {
	o.AccountId = &v
}

// GetCoinId returns the CoinId field value if set, zero value otherwise.
func (o *PositionAsset) GetCoinId() string {
	if o == nil || IsNil(o.CoinId) {
		var ret string
		return ret
	}
	return *o.CoinId
}

// GetCoinIdOk returns a tuple with the CoinId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionAsset) GetCoinIdOk() (*string, bool) {
	if o == nil || IsNil(o.CoinId) {
		return nil, false
	}
	return o.CoinId, true
}

// HasCoinId returns a boolean if a field has been set.
func (o *PositionAsset) HasCoinId() bool {
	if o != nil && !IsNil(o.CoinId) {
		return true
	}

	return false
}

// SetCoinId gets a reference to the given string and assigns it to the CoinId field.
func (o *PositionAsset) SetCoinId(v string) {
	o.CoinId = &v
}

// GetContractId returns the ContractId field value if set, zero value otherwise.
func (o *PositionAsset) GetContractId() string {
	if o == nil || IsNil(o.ContractId) {
		var ret string
		return ret
	}
	return *o.ContractId
}

// GetContractIdOk returns a tuple with the ContractId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionAsset) GetContractIdOk() (*string, bool) {
	if o == nil || IsNil(o.ContractId) {
		return nil, false
	}
	return o.ContractId, true
}

// HasContractId returns a boolean if a field has been set.
func (o *PositionAsset) HasContractId() bool {
	if o != nil && !IsNil(o.ContractId) {
		return true
	}

	return false
}

// SetContractId gets a reference to the given string and assigns it to the ContractId field.
func (o *PositionAsset) SetContractId(v string) {
	o.ContractId = &v
}

// GetPositionValue returns the PositionValue field value if set, zero value otherwise.
func (o *PositionAsset) GetPositionValue() string {
	if o == nil || IsNil(o.PositionValue) {
		var ret string
		return ret
	}
	return *o.PositionValue
}

// GetPositionValueOk returns a tuple with the PositionValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionAsset) GetPositionValueOk() (*string, bool) {
	if o == nil || IsNil(o.PositionValue) {
		return nil, false
	}
	return o.PositionValue, true
}

// HasPositionValue returns a boolean if a field has been set.
func (o *PositionAsset) HasPositionValue() bool {
	if o != nil && !IsNil(o.PositionValue) {
		return true
	}

	return false
}

// SetPositionValue gets a reference to the given string and assigns it to the PositionValue field.
func (o *PositionAsset) SetPositionValue(v string) {
	o.PositionValue = &v
}

// GetMaxLeverage returns the MaxLeverage field value if set, zero value otherwise.
func (o *PositionAsset) GetMaxLeverage() string {
	if o == nil || IsNil(o.MaxLeverage) {
		var ret string
		return ret
	}
	return *o.MaxLeverage
}

// GetMaxLeverageOk returns a tuple with the MaxLeverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionAsset) GetMaxLeverageOk() (*string, bool) {
	if o == nil || IsNil(o.MaxLeverage) {
		return nil, false
	}
	return o.MaxLeverage, true
}

// HasMaxLeverage returns a boolean if a field has been set.
func (o *PositionAsset) HasMaxLeverage() bool {
	if o != nil && !IsNil(o.MaxLeverage) {
		return true
	}

	return false
}

// SetMaxLeverage gets a reference to the given string and assigns it to the MaxLeverage field.
func (o *PositionAsset) SetMaxLeverage(v string) {
	o.MaxLeverage = &v
}

// GetInitialMarginRequirement returns the InitialMarginRequirement field value if set, zero value otherwise.
func (o *PositionAsset) GetInitialMarginRequirement() string {
	if o == nil || IsNil(o.InitialMarginRequirement) {
		var ret string
		return ret
	}
	return *o.InitialMarginRequirement
}

// GetInitialMarginRequirementOk returns a tuple with the InitialMarginRequirement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionAsset) GetInitialMarginRequirementOk() (*string, bool) {
	if o == nil || IsNil(o.InitialMarginRequirement) {
		return nil, false
	}
	return o.InitialMarginRequirement, true
}

// HasInitialMarginRequirement returns a boolean if a field has been set.
func (o *PositionAsset) HasInitialMarginRequirement() bool {
	if o != nil && !IsNil(o.InitialMarginRequirement) {
		return true
	}

	return false
}

// SetInitialMarginRequirement gets a reference to the given string and assigns it to the InitialMarginRequirement field.
func (o *PositionAsset) SetInitialMarginRequirement(v string) {
	o.InitialMarginRequirement = &v
}

// GetStarkExRiskRate returns the StarkExRiskRate field value if set, zero value otherwise.
func (o *PositionAsset) GetStarkExRiskRate() string {
	if o == nil || IsNil(o.StarkExRiskRate) {
		var ret string
		return ret
	}
	return *o.StarkExRiskRate
}

// GetStarkExRiskRateOk returns a tuple with the StarkExRiskRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionAsset) GetStarkExRiskRateOk() (*string, bool) {
	if o == nil || IsNil(o.StarkExRiskRate) {
		return nil, false
	}
	return o.StarkExRiskRate, true
}

// HasStarkExRiskRate returns a boolean if a field has been set.
func (o *PositionAsset) HasStarkExRiskRate() bool {
	if o != nil && !IsNil(o.StarkExRiskRate) {
		return true
	}

	return false
}

// SetStarkExRiskRate gets a reference to the given string and assigns it to the StarkExRiskRate field.
func (o *PositionAsset) SetStarkExRiskRate(v string) {
	o.StarkExRiskRate = &v
}

// GetStarkExRiskValue returns the StarkExRiskValue field value if set, zero value otherwise.
func (o *PositionAsset) GetStarkExRiskValue() string {
	if o == nil || IsNil(o.StarkExRiskValue) {
		var ret string
		return ret
	}
	return *o.StarkExRiskValue
}

// GetStarkExRiskValueOk returns a tuple with the StarkExRiskValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionAsset) GetStarkExRiskValueOk() (*string, bool) {
	if o == nil || IsNil(o.StarkExRiskValue) {
		return nil, false
	}
	return o.StarkExRiskValue, true
}

// HasStarkExRiskValue returns a boolean if a field has been set.
func (o *PositionAsset) HasStarkExRiskValue() bool {
	if o != nil && !IsNil(o.StarkExRiskValue) {
		return true
	}

	return false
}

// SetStarkExRiskValue gets a reference to the given string and assigns it to the StarkExRiskValue field.
func (o *PositionAsset) SetStarkExRiskValue(v string) {
	o.StarkExRiskValue = &v
}

// GetAvgEntryPrice returns the AvgEntryPrice field value if set, zero value otherwise.
func (o *PositionAsset) GetAvgEntryPrice() string {
	if o == nil || IsNil(o.AvgEntryPrice) {
		var ret string
		return ret
	}
	return *o.AvgEntryPrice
}

// GetAvgEntryPriceOk returns a tuple with the AvgEntryPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionAsset) GetAvgEntryPriceOk() (*string, bool) {
	if o == nil || IsNil(o.AvgEntryPrice) {
		return nil, false
	}
	return o.AvgEntryPrice, true
}

// HasAvgEntryPrice returns a boolean if a field has been set.
func (o *PositionAsset) HasAvgEntryPrice() bool {
	if o != nil && !IsNil(o.AvgEntryPrice) {
		return true
	}

	return false
}

// SetAvgEntryPrice gets a reference to the given string and assigns it to the AvgEntryPrice field.
func (o *PositionAsset) SetAvgEntryPrice(v string) {
	o.AvgEntryPrice = &v
}

// GetLiquidatePrice returns the LiquidatePrice field value if set, zero value otherwise.
func (o *PositionAsset) GetLiquidatePrice() string {
	if o == nil || IsNil(o.LiquidatePrice) {
		var ret string
		return ret
	}
	return *o.LiquidatePrice
}

// GetLiquidatePriceOk returns a tuple with the LiquidatePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionAsset) GetLiquidatePriceOk() (*string, bool) {
	if o == nil || IsNil(o.LiquidatePrice) {
		return nil, false
	}
	return o.LiquidatePrice, true
}

// HasLiquidatePrice returns a boolean if a field has been set.
func (o *PositionAsset) HasLiquidatePrice() bool {
	if o != nil && !IsNil(o.LiquidatePrice) {
		return true
	}

	return false
}

// SetLiquidatePrice gets a reference to the given string and assigns it to the LiquidatePrice field.
func (o *PositionAsset) SetLiquidatePrice(v string) {
	o.LiquidatePrice = &v
}

// GetBankruptPrice returns the BankruptPrice field value if set, zero value otherwise.
func (o *PositionAsset) GetBankruptPrice() string {
	if o == nil || IsNil(o.BankruptPrice) {
		var ret string
		return ret
	}
	return *o.BankruptPrice
}

// GetBankruptPriceOk returns a tuple with the BankruptPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionAsset) GetBankruptPriceOk() (*string, bool) {
	if o == nil || IsNil(o.BankruptPrice) {
		return nil, false
	}
	return o.BankruptPrice, true
}

// HasBankruptPrice returns a boolean if a field has been set.
func (o *PositionAsset) HasBankruptPrice() bool {
	if o != nil && !IsNil(o.BankruptPrice) {
		return true
	}

	return false
}

// SetBankruptPrice gets a reference to the given string and assigns it to the BankruptPrice field.
func (o *PositionAsset) SetBankruptPrice(v string) {
	o.BankruptPrice = &v
}

// GetWorstClosePrice returns the WorstClosePrice field value if set, zero value otherwise.
func (o *PositionAsset) GetWorstClosePrice() string {
	if o == nil || IsNil(o.WorstClosePrice) {
		var ret string
		return ret
	}
	return *o.WorstClosePrice
}

// GetWorstClosePriceOk returns a tuple with the WorstClosePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionAsset) GetWorstClosePriceOk() (*string, bool) {
	if o == nil || IsNil(o.WorstClosePrice) {
		return nil, false
	}
	return o.WorstClosePrice, true
}

// HasWorstClosePrice returns a boolean if a field has been set.
func (o *PositionAsset) HasWorstClosePrice() bool {
	if o != nil && !IsNil(o.WorstClosePrice) {
		return true
	}

	return false
}

// SetWorstClosePrice gets a reference to the given string and assigns it to the WorstClosePrice field.
func (o *PositionAsset) SetWorstClosePrice(v string) {
	o.WorstClosePrice = &v
}

// GetUnrealizePnl returns the UnrealizePnl field value if set, zero value otherwise.
func (o *PositionAsset) GetUnrealizePnl() string {
	if o == nil || IsNil(o.UnrealizePnl) {
		var ret string
		return ret
	}
	return *o.UnrealizePnl
}

// GetUnrealizePnlOk returns a tuple with the UnrealizePnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionAsset) GetUnrealizePnlOk() (*string, bool) {
	if o == nil || IsNil(o.UnrealizePnl) {
		return nil, false
	}
	return o.UnrealizePnl, true
}

// HasUnrealizePnl returns a boolean if a field has been set.
func (o *PositionAsset) HasUnrealizePnl() bool {
	if o != nil && !IsNil(o.UnrealizePnl) {
		return true
	}

	return false
}

// SetUnrealizePnl gets a reference to the given string and assigns it to the UnrealizePnl field.
func (o *PositionAsset) SetUnrealizePnl(v string) {
	o.UnrealizePnl = &v
}

// GetTermRealizePnl returns the TermRealizePnl field value if set, zero value otherwise.
func (o *PositionAsset) GetTermRealizePnl() string {
	if o == nil || IsNil(o.TermRealizePnl) {
		var ret string
		return ret
	}
	return *o.TermRealizePnl
}

// GetTermRealizePnlOk returns a tuple with the TermRealizePnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionAsset) GetTermRealizePnlOk() (*string, bool) {
	if o == nil || IsNil(o.TermRealizePnl) {
		return nil, false
	}
	return o.TermRealizePnl, true
}

// HasTermRealizePnl returns a boolean if a field has been set.
func (o *PositionAsset) HasTermRealizePnl() bool {
	if o != nil && !IsNil(o.TermRealizePnl) {
		return true
	}

	return false
}

// SetTermRealizePnl gets a reference to the given string and assigns it to the TermRealizePnl field.
func (o *PositionAsset) SetTermRealizePnl(v string) {
	o.TermRealizePnl = &v
}

// GetTotalRealizePnl returns the TotalRealizePnl field value if set, zero value otherwise.
func (o *PositionAsset) GetTotalRealizePnl() string {
	if o == nil || IsNil(o.TotalRealizePnl) {
		var ret string
		return ret
	}
	return *o.TotalRealizePnl
}

// GetTotalRealizePnlOk returns a tuple with the TotalRealizePnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositionAsset) GetTotalRealizePnlOk() (*string, bool) {
	if o == nil || IsNil(o.TotalRealizePnl) {
		return nil, false
	}
	return o.TotalRealizePnl, true
}

// HasTotalRealizePnl returns a boolean if a field has been set.
func (o *PositionAsset) HasTotalRealizePnl() bool {
	if o != nil && !IsNil(o.TotalRealizePnl) {
		return true
	}

	return false
}

// SetTotalRealizePnl gets a reference to the given string and assigns it to the TotalRealizePnl field.
func (o *PositionAsset) SetTotalRealizePnl(v string) {
	o.TotalRealizePnl = &v
}

func (o PositionAsset) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PositionAsset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	if !IsNil(o.CoinId) {
		toSerialize["coinId"] = o.CoinId
	}
	if !IsNil(o.ContractId) {
		toSerialize["contractId"] = o.ContractId
	}
	if !IsNil(o.PositionValue) {
		toSerialize["positionValue"] = o.PositionValue
	}
	if !IsNil(o.MaxLeverage) {
		toSerialize["maxLeverage"] = o.MaxLeverage
	}
	if !IsNil(o.InitialMarginRequirement) {
		toSerialize["initialMarginRequirement"] = o.InitialMarginRequirement
	}
	if !IsNil(o.StarkExRiskRate) {
		toSerialize["starkExRiskRate"] = o.StarkExRiskRate
	}
	if !IsNil(o.StarkExRiskValue) {
		toSerialize["starkExRiskValue"] = o.StarkExRiskValue
	}
	if !IsNil(o.AvgEntryPrice) {
		toSerialize["avgEntryPrice"] = o.AvgEntryPrice
	}
	if !IsNil(o.LiquidatePrice) {
		toSerialize["liquidatePrice"] = o.LiquidatePrice
	}
	if !IsNil(o.BankruptPrice) {
		toSerialize["bankruptPrice"] = o.BankruptPrice
	}
	if !IsNil(o.WorstClosePrice) {
		toSerialize["worstClosePrice"] = o.WorstClosePrice
	}
	if !IsNil(o.UnrealizePnl) {
		toSerialize["unrealizePnl"] = o.UnrealizePnl
	}
	if !IsNil(o.TermRealizePnl) {
		toSerialize["termRealizePnl"] = o.TermRealizePnl
	}
	if !IsNil(o.TotalRealizePnl) {
		toSerialize["totalRealizePnl"] = o.TotalRealizePnl
	}
	return toSerialize, nil
}

type NullablePositionAsset struct {
	value *PositionAsset
	isSet bool
}

func (v NullablePositionAsset) Get() *PositionAsset {
	return v.value
}

func (v *NullablePositionAsset) Set(val *PositionAsset) {
	v.value = val
	v.isSet = true
}

func (v NullablePositionAsset) IsSet() bool {
	return v.isSet
}

func (v *NullablePositionAsset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePositionAsset(val *PositionAsset) *NullablePositionAsset {
	return &NullablePositionAsset{value: val, isSet: true}
}

func (v NullablePositionAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePositionAsset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


