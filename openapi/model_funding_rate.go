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

// checks if the FundingRate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FundingRate{}

// FundingRate 资金费率
type FundingRate struct {
	// 合约ID
	ContractId *string `json:"contractId,omitempty"`
	// 资金费率结算时间 如08:00-09:00这一期的资金费率，是通过上一期07:00-08:00的数据计算而得，在08:00的时候已经确定，在09:00进行结算时使用
	FundingTime *string `json:"fundingTime,omitempty"`
	// 费率计算时间 millisecond
	FundingTimestamp *string `json:"fundingTimestamp,omitempty"`
	// 预言机价格
	OraclePrice *string `json:"oraclePrice,omitempty"`
	// 指数价格
	IndexPrice *string `json:"indexPrice,omitempty"`
	// 资金费率
	FundingRate *string `json:"fundingRate,omitempty"`
	// 资金费率结算标志
	IsSettlement *bool `json:"isSettlement,omitempty"`
	// 预测资金费率
	ForecastFundingRate *string `json:"forecastFundingRate,omitempty"`
	// 上一次的资金费率
	PreviousFundingRate *string `json:"previousFundingRate,omitempty"`
	// 上一次资金费率计算时间 millisecond
	PreviousFundingTimestamp *string `json:"previousFundingTimestamp,omitempty"`
	// 溢价指数
	PremiumIndex *string `json:"premiumIndex,omitempty"`
	// 平均溢价指数
	AvgPremiumIndex *string `json:"avgPremiumIndex,omitempty"`
	// 溢价指数计算时间
	PremiumIndexTimestamp *string `json:"premiumIndexTimestamp,omitempty"`
	// 深度加权买卖价格需计算的数量
	ImpactMarginNotional *string `json:"impactMarginNotional,omitempty"`
	// 深度加权卖价
	ImpactAskPrice *string `json:"impactAskPrice,omitempty"`
	// 深度加权买价
	ImpactBidPrice *string `json:"impactBidPrice,omitempty"`
	// 固定利率
	InterestRate *string `json:"interestRate,omitempty"`
	// 综合利率 interestRate/频率
	PredictedFundingRate *string `json:"predictedFundingRate,omitempty"`
	// 资金费率时间间隔 min
	FundingRateIntervalMin *string `json:"fundingRateIntervalMin,omitempty"`
	// 资金费指数
	StarkExFundingIndex *string `json:"starkExFundingIndex,omitempty"`
}

// NewFundingRate instantiates a new FundingRate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFundingRate() *FundingRate {
	this := FundingRate{}
	return &this
}

// NewFundingRateWithDefaults instantiates a new FundingRate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFundingRateWithDefaults() *FundingRate {
	this := FundingRate{}
	return &this
}

// GetContractId returns the ContractId field value if set, zero value otherwise.
func (o *FundingRate) GetContractId() string {
	if o == nil || IsNil(o.ContractId) {
		var ret string
		return ret
	}
	return *o.ContractId
}

// GetContractIdOk returns a tuple with the ContractId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingRate) GetContractIdOk() (*string, bool) {
	if o == nil || IsNil(o.ContractId) {
		return nil, false
	}
	return o.ContractId, true
}

// HasContractId returns a boolean if a field has been set.
func (o *FundingRate) HasContractId() bool {
	if o != nil && !IsNil(o.ContractId) {
		return true
	}

	return false
}

// SetContractId gets a reference to the given string and assigns it to the ContractId field.
func (o *FundingRate) SetContractId(v string) {
	o.ContractId = &v
}

// GetFundingTime returns the FundingTime field value if set, zero value otherwise.
func (o *FundingRate) GetFundingTime() string {
	if o == nil || IsNil(o.FundingTime) {
		var ret string
		return ret
	}
	return *o.FundingTime
}

// GetFundingTimeOk returns a tuple with the FundingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingRate) GetFundingTimeOk() (*string, bool) {
	if o == nil || IsNil(o.FundingTime) {
		return nil, false
	}
	return o.FundingTime, true
}

// HasFundingTime returns a boolean if a field has been set.
func (o *FundingRate) HasFundingTime() bool {
	if o != nil && !IsNil(o.FundingTime) {
		return true
	}

	return false
}

// SetFundingTime gets a reference to the given string and assigns it to the FundingTime field.
func (o *FundingRate) SetFundingTime(v string) {
	o.FundingTime = &v
}

// GetFundingTimestamp returns the FundingTimestamp field value if set, zero value otherwise.
func (o *FundingRate) GetFundingTimestamp() string {
	if o == nil || IsNil(o.FundingTimestamp) {
		var ret string
		return ret
	}
	return *o.FundingTimestamp
}

// GetFundingTimestampOk returns a tuple with the FundingTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingRate) GetFundingTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.FundingTimestamp) {
		return nil, false
	}
	return o.FundingTimestamp, true
}

// HasFundingTimestamp returns a boolean if a field has been set.
func (o *FundingRate) HasFundingTimestamp() bool {
	if o != nil && !IsNil(o.FundingTimestamp) {
		return true
	}

	return false
}

// SetFundingTimestamp gets a reference to the given string and assigns it to the FundingTimestamp field.
func (o *FundingRate) SetFundingTimestamp(v string) {
	o.FundingTimestamp = &v
}

// GetOraclePrice returns the OraclePrice field value if set, zero value otherwise.
func (o *FundingRate) GetOraclePrice() string {
	if o == nil || IsNil(o.OraclePrice) {
		var ret string
		return ret
	}
	return *o.OraclePrice
}

// GetOraclePriceOk returns a tuple with the OraclePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingRate) GetOraclePriceOk() (*string, bool) {
	if o == nil || IsNil(o.OraclePrice) {
		return nil, false
	}
	return o.OraclePrice, true
}

// HasOraclePrice returns a boolean if a field has been set.
func (o *FundingRate) HasOraclePrice() bool {
	if o != nil && !IsNil(o.OraclePrice) {
		return true
	}

	return false
}

// SetOraclePrice gets a reference to the given string and assigns it to the OraclePrice field.
func (o *FundingRate) SetOraclePrice(v string) {
	o.OraclePrice = &v
}

// GetIndexPrice returns the IndexPrice field value if set, zero value otherwise.
func (o *FundingRate) GetIndexPrice() string {
	if o == nil || IsNil(o.IndexPrice) {
		var ret string
		return ret
	}
	return *o.IndexPrice
}

// GetIndexPriceOk returns a tuple with the IndexPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingRate) GetIndexPriceOk() (*string, bool) {
	if o == nil || IsNil(o.IndexPrice) {
		return nil, false
	}
	return o.IndexPrice, true
}

// HasIndexPrice returns a boolean if a field has been set.
func (o *FundingRate) HasIndexPrice() bool {
	if o != nil && !IsNil(o.IndexPrice) {
		return true
	}

	return false
}

// SetIndexPrice gets a reference to the given string and assigns it to the IndexPrice field.
func (o *FundingRate) SetIndexPrice(v string) {
	o.IndexPrice = &v
}

// GetFundingRate returns the FundingRate field value if set, zero value otherwise.
func (o *FundingRate) GetFundingRate() string {
	if o == nil || IsNil(o.FundingRate) {
		var ret string
		return ret
	}
	return *o.FundingRate
}

// GetFundingRateOk returns a tuple with the FundingRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingRate) GetFundingRateOk() (*string, bool) {
	if o == nil || IsNil(o.FundingRate) {
		return nil, false
	}
	return o.FundingRate, true
}

// HasFundingRate returns a boolean if a field has been set.
func (o *FundingRate) HasFundingRate() bool {
	if o != nil && !IsNil(o.FundingRate) {
		return true
	}

	return false
}

// SetFundingRate gets a reference to the given string and assigns it to the FundingRate field.
func (o *FundingRate) SetFundingRate(v string) {
	o.FundingRate = &v
}

// GetIsSettlement returns the IsSettlement field value if set, zero value otherwise.
func (o *FundingRate) GetIsSettlement() bool {
	if o == nil || IsNil(o.IsSettlement) {
		var ret bool
		return ret
	}
	return *o.IsSettlement
}

// GetIsSettlementOk returns a tuple with the IsSettlement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingRate) GetIsSettlementOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSettlement) {
		return nil, false
	}
	return o.IsSettlement, true
}

// HasIsSettlement returns a boolean if a field has been set.
func (o *FundingRate) HasIsSettlement() bool {
	if o != nil && !IsNil(o.IsSettlement) {
		return true
	}

	return false
}

// SetIsSettlement gets a reference to the given bool and assigns it to the IsSettlement field.
func (o *FundingRate) SetIsSettlement(v bool) {
	o.IsSettlement = &v
}

// GetForecastFundingRate returns the ForecastFundingRate field value if set, zero value otherwise.
func (o *FundingRate) GetForecastFundingRate() string {
	if o == nil || IsNil(o.ForecastFundingRate) {
		var ret string
		return ret
	}
	return *o.ForecastFundingRate
}

// GetForecastFundingRateOk returns a tuple with the ForecastFundingRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingRate) GetForecastFundingRateOk() (*string, bool) {
	if o == nil || IsNil(o.ForecastFundingRate) {
		return nil, false
	}
	return o.ForecastFundingRate, true
}

// HasForecastFundingRate returns a boolean if a field has been set.
func (o *FundingRate) HasForecastFundingRate() bool {
	if o != nil && !IsNil(o.ForecastFundingRate) {
		return true
	}

	return false
}

// SetForecastFundingRate gets a reference to the given string and assigns it to the ForecastFundingRate field.
func (o *FundingRate) SetForecastFundingRate(v string) {
	o.ForecastFundingRate = &v
}

// GetPreviousFundingRate returns the PreviousFundingRate field value if set, zero value otherwise.
func (o *FundingRate) GetPreviousFundingRate() string {
	if o == nil || IsNil(o.PreviousFundingRate) {
		var ret string
		return ret
	}
	return *o.PreviousFundingRate
}

// GetPreviousFundingRateOk returns a tuple with the PreviousFundingRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingRate) GetPreviousFundingRateOk() (*string, bool) {
	if o == nil || IsNil(o.PreviousFundingRate) {
		return nil, false
	}
	return o.PreviousFundingRate, true
}

// HasPreviousFundingRate returns a boolean if a field has been set.
func (o *FundingRate) HasPreviousFundingRate() bool {
	if o != nil && !IsNil(o.PreviousFundingRate) {
		return true
	}

	return false
}

// SetPreviousFundingRate gets a reference to the given string and assigns it to the PreviousFundingRate field.
func (o *FundingRate) SetPreviousFundingRate(v string) {
	o.PreviousFundingRate = &v
}

// GetPreviousFundingTimestamp returns the PreviousFundingTimestamp field value if set, zero value otherwise.
func (o *FundingRate) GetPreviousFundingTimestamp() string {
	if o == nil || IsNil(o.PreviousFundingTimestamp) {
		var ret string
		return ret
	}
	return *o.PreviousFundingTimestamp
}

// GetPreviousFundingTimestampOk returns a tuple with the PreviousFundingTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingRate) GetPreviousFundingTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.PreviousFundingTimestamp) {
		return nil, false
	}
	return o.PreviousFundingTimestamp, true
}

// HasPreviousFundingTimestamp returns a boolean if a field has been set.
func (o *FundingRate) HasPreviousFundingTimestamp() bool {
	if o != nil && !IsNil(o.PreviousFundingTimestamp) {
		return true
	}

	return false
}

// SetPreviousFundingTimestamp gets a reference to the given string and assigns it to the PreviousFundingTimestamp field.
func (o *FundingRate) SetPreviousFundingTimestamp(v string) {
	o.PreviousFundingTimestamp = &v
}

// GetPremiumIndex returns the PremiumIndex field value if set, zero value otherwise.
func (o *FundingRate) GetPremiumIndex() string {
	if o == nil || IsNil(o.PremiumIndex) {
		var ret string
		return ret
	}
	return *o.PremiumIndex
}

// GetPremiumIndexOk returns a tuple with the PremiumIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingRate) GetPremiumIndexOk() (*string, bool) {
	if o == nil || IsNil(o.PremiumIndex) {
		return nil, false
	}
	return o.PremiumIndex, true
}

// HasPremiumIndex returns a boolean if a field has been set.
func (o *FundingRate) HasPremiumIndex() bool {
	if o != nil && !IsNil(o.PremiumIndex) {
		return true
	}

	return false
}

// SetPremiumIndex gets a reference to the given string and assigns it to the PremiumIndex field.
func (o *FundingRate) SetPremiumIndex(v string) {
	o.PremiumIndex = &v
}

// GetAvgPremiumIndex returns the AvgPremiumIndex field value if set, zero value otherwise.
func (o *FundingRate) GetAvgPremiumIndex() string {
	if o == nil || IsNil(o.AvgPremiumIndex) {
		var ret string
		return ret
	}
	return *o.AvgPremiumIndex
}

// GetAvgPremiumIndexOk returns a tuple with the AvgPremiumIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingRate) GetAvgPremiumIndexOk() (*string, bool) {
	if o == nil || IsNil(o.AvgPremiumIndex) {
		return nil, false
	}
	return o.AvgPremiumIndex, true
}

// HasAvgPremiumIndex returns a boolean if a field has been set.
func (o *FundingRate) HasAvgPremiumIndex() bool {
	if o != nil && !IsNil(o.AvgPremiumIndex) {
		return true
	}

	return false
}

// SetAvgPremiumIndex gets a reference to the given string and assigns it to the AvgPremiumIndex field.
func (o *FundingRate) SetAvgPremiumIndex(v string) {
	o.AvgPremiumIndex = &v
}

// GetPremiumIndexTimestamp returns the PremiumIndexTimestamp field value if set, zero value otherwise.
func (o *FundingRate) GetPremiumIndexTimestamp() string {
	if o == nil || IsNil(o.PremiumIndexTimestamp) {
		var ret string
		return ret
	}
	return *o.PremiumIndexTimestamp
}

// GetPremiumIndexTimestampOk returns a tuple with the PremiumIndexTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingRate) GetPremiumIndexTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.PremiumIndexTimestamp) {
		return nil, false
	}
	return o.PremiumIndexTimestamp, true
}

// HasPremiumIndexTimestamp returns a boolean if a field has been set.
func (o *FundingRate) HasPremiumIndexTimestamp() bool {
	if o != nil && !IsNil(o.PremiumIndexTimestamp) {
		return true
	}

	return false
}

// SetPremiumIndexTimestamp gets a reference to the given string and assigns it to the PremiumIndexTimestamp field.
func (o *FundingRate) SetPremiumIndexTimestamp(v string) {
	o.PremiumIndexTimestamp = &v
}

// GetImpactMarginNotional returns the ImpactMarginNotional field value if set, zero value otherwise.
func (o *FundingRate) GetImpactMarginNotional() string {
	if o == nil || IsNil(o.ImpactMarginNotional) {
		var ret string
		return ret
	}
	return *o.ImpactMarginNotional
}

// GetImpactMarginNotionalOk returns a tuple with the ImpactMarginNotional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingRate) GetImpactMarginNotionalOk() (*string, bool) {
	if o == nil || IsNil(o.ImpactMarginNotional) {
		return nil, false
	}
	return o.ImpactMarginNotional, true
}

// HasImpactMarginNotional returns a boolean if a field has been set.
func (o *FundingRate) HasImpactMarginNotional() bool {
	if o != nil && !IsNil(o.ImpactMarginNotional) {
		return true
	}

	return false
}

// SetImpactMarginNotional gets a reference to the given string and assigns it to the ImpactMarginNotional field.
func (o *FundingRate) SetImpactMarginNotional(v string) {
	o.ImpactMarginNotional = &v
}

// GetImpactAskPrice returns the ImpactAskPrice field value if set, zero value otherwise.
func (o *FundingRate) GetImpactAskPrice() string {
	if o == nil || IsNil(o.ImpactAskPrice) {
		var ret string
		return ret
	}
	return *o.ImpactAskPrice
}

// GetImpactAskPriceOk returns a tuple with the ImpactAskPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingRate) GetImpactAskPriceOk() (*string, bool) {
	if o == nil || IsNil(o.ImpactAskPrice) {
		return nil, false
	}
	return o.ImpactAskPrice, true
}

// HasImpactAskPrice returns a boolean if a field has been set.
func (o *FundingRate) HasImpactAskPrice() bool {
	if o != nil && !IsNil(o.ImpactAskPrice) {
		return true
	}

	return false
}

// SetImpactAskPrice gets a reference to the given string and assigns it to the ImpactAskPrice field.
func (o *FundingRate) SetImpactAskPrice(v string) {
	o.ImpactAskPrice = &v
}

// GetImpactBidPrice returns the ImpactBidPrice field value if set, zero value otherwise.
func (o *FundingRate) GetImpactBidPrice() string {
	if o == nil || IsNil(o.ImpactBidPrice) {
		var ret string
		return ret
	}
	return *o.ImpactBidPrice
}

// GetImpactBidPriceOk returns a tuple with the ImpactBidPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingRate) GetImpactBidPriceOk() (*string, bool) {
	if o == nil || IsNil(o.ImpactBidPrice) {
		return nil, false
	}
	return o.ImpactBidPrice, true
}

// HasImpactBidPrice returns a boolean if a field has been set.
func (o *FundingRate) HasImpactBidPrice() bool {
	if o != nil && !IsNil(o.ImpactBidPrice) {
		return true
	}

	return false
}

// SetImpactBidPrice gets a reference to the given string and assigns it to the ImpactBidPrice field.
func (o *FundingRate) SetImpactBidPrice(v string) {
	o.ImpactBidPrice = &v
}

// GetInterestRate returns the InterestRate field value if set, zero value otherwise.
func (o *FundingRate) GetInterestRate() string {
	if o == nil || IsNil(o.InterestRate) {
		var ret string
		return ret
	}
	return *o.InterestRate
}

// GetInterestRateOk returns a tuple with the InterestRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingRate) GetInterestRateOk() (*string, bool) {
	if o == nil || IsNil(o.InterestRate) {
		return nil, false
	}
	return o.InterestRate, true
}

// HasInterestRate returns a boolean if a field has been set.
func (o *FundingRate) HasInterestRate() bool {
	if o != nil && !IsNil(o.InterestRate) {
		return true
	}

	return false
}

// SetInterestRate gets a reference to the given string and assigns it to the InterestRate field.
func (o *FundingRate) SetInterestRate(v string) {
	o.InterestRate = &v
}

// GetPredictedFundingRate returns the PredictedFundingRate field value if set, zero value otherwise.
func (o *FundingRate) GetPredictedFundingRate() string {
	if o == nil || IsNil(o.PredictedFundingRate) {
		var ret string
		return ret
	}
	return *o.PredictedFundingRate
}

// GetPredictedFundingRateOk returns a tuple with the PredictedFundingRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingRate) GetPredictedFundingRateOk() (*string, bool) {
	if o == nil || IsNil(o.PredictedFundingRate) {
		return nil, false
	}
	return o.PredictedFundingRate, true
}

// HasPredictedFundingRate returns a boolean if a field has been set.
func (o *FundingRate) HasPredictedFundingRate() bool {
	if o != nil && !IsNil(o.PredictedFundingRate) {
		return true
	}

	return false
}

// SetPredictedFundingRate gets a reference to the given string and assigns it to the PredictedFundingRate field.
func (o *FundingRate) SetPredictedFundingRate(v string) {
	o.PredictedFundingRate = &v
}

// GetFundingRateIntervalMin returns the FundingRateIntervalMin field value if set, zero value otherwise.
func (o *FundingRate) GetFundingRateIntervalMin() string {
	if o == nil || IsNil(o.FundingRateIntervalMin) {
		var ret string
		return ret
	}
	return *o.FundingRateIntervalMin
}

// GetFundingRateIntervalMinOk returns a tuple with the FundingRateIntervalMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingRate) GetFundingRateIntervalMinOk() (*string, bool) {
	if o == nil || IsNil(o.FundingRateIntervalMin) {
		return nil, false
	}
	return o.FundingRateIntervalMin, true
}

// HasFundingRateIntervalMin returns a boolean if a field has been set.
func (o *FundingRate) HasFundingRateIntervalMin() bool {
	if o != nil && !IsNil(o.FundingRateIntervalMin) {
		return true
	}

	return false
}

// SetFundingRateIntervalMin gets a reference to the given string and assigns it to the FundingRateIntervalMin field.
func (o *FundingRate) SetFundingRateIntervalMin(v string) {
	o.FundingRateIntervalMin = &v
}

// GetStarkExFundingIndex returns the StarkExFundingIndex field value if set, zero value otherwise.
func (o *FundingRate) GetStarkExFundingIndex() string {
	if o == nil || IsNil(o.StarkExFundingIndex) {
		var ret string
		return ret
	}
	return *o.StarkExFundingIndex
}

// GetStarkExFundingIndexOk returns a tuple with the StarkExFundingIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundingRate) GetStarkExFundingIndexOk() (*string, bool) {
	if o == nil || IsNil(o.StarkExFundingIndex) {
		return nil, false
	}
	return o.StarkExFundingIndex, true
}

// HasStarkExFundingIndex returns a boolean if a field has been set.
func (o *FundingRate) HasStarkExFundingIndex() bool {
	if o != nil && !IsNil(o.StarkExFundingIndex) {
		return true
	}

	return false
}

// SetStarkExFundingIndex gets a reference to the given string and assigns it to the StarkExFundingIndex field.
func (o *FundingRate) SetStarkExFundingIndex(v string) {
	o.StarkExFundingIndex = &v
}

func (o FundingRate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FundingRate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContractId) {
		toSerialize["contractId"] = o.ContractId
	}
	if !IsNil(o.FundingTime) {
		toSerialize["fundingTime"] = o.FundingTime
	}
	if !IsNil(o.FundingTimestamp) {
		toSerialize["fundingTimestamp"] = o.FundingTimestamp
	}
	if !IsNil(o.OraclePrice) {
		toSerialize["oraclePrice"] = o.OraclePrice
	}
	if !IsNil(o.IndexPrice) {
		toSerialize["indexPrice"] = o.IndexPrice
	}
	if !IsNil(o.FundingRate) {
		toSerialize["fundingRate"] = o.FundingRate
	}
	if !IsNil(o.IsSettlement) {
		toSerialize["isSettlement"] = o.IsSettlement
	}
	if !IsNil(o.ForecastFundingRate) {
		toSerialize["forecastFundingRate"] = o.ForecastFundingRate
	}
	if !IsNil(o.PreviousFundingRate) {
		toSerialize["previousFundingRate"] = o.PreviousFundingRate
	}
	if !IsNil(o.PreviousFundingTimestamp) {
		toSerialize["previousFundingTimestamp"] = o.PreviousFundingTimestamp
	}
	if !IsNil(o.PremiumIndex) {
		toSerialize["premiumIndex"] = o.PremiumIndex
	}
	if !IsNil(o.AvgPremiumIndex) {
		toSerialize["avgPremiumIndex"] = o.AvgPremiumIndex
	}
	if !IsNil(o.PremiumIndexTimestamp) {
		toSerialize["premiumIndexTimestamp"] = o.PremiumIndexTimestamp
	}
	if !IsNil(o.ImpactMarginNotional) {
		toSerialize["impactMarginNotional"] = o.ImpactMarginNotional
	}
	if !IsNil(o.ImpactAskPrice) {
		toSerialize["impactAskPrice"] = o.ImpactAskPrice
	}
	if !IsNil(o.ImpactBidPrice) {
		toSerialize["impactBidPrice"] = o.ImpactBidPrice
	}
	if !IsNil(o.InterestRate) {
		toSerialize["interestRate"] = o.InterestRate
	}
	if !IsNil(o.PredictedFundingRate) {
		toSerialize["predictedFundingRate"] = o.PredictedFundingRate
	}
	if !IsNil(o.FundingRateIntervalMin) {
		toSerialize["fundingRateIntervalMin"] = o.FundingRateIntervalMin
	}
	if !IsNil(o.StarkExFundingIndex) {
		toSerialize["starkExFundingIndex"] = o.StarkExFundingIndex
	}
	return toSerialize, nil
}

type NullableFundingRate struct {
	value *FundingRate
	isSet bool
}

func (v NullableFundingRate) Get() *FundingRate {
	return v.value
}

func (v *NullableFundingRate) Set(val *FundingRate) {
	v.value = val
	v.isSet = true
}

func (v NullableFundingRate) IsSet() bool {
	return v.isSet
}

func (v *NullableFundingRate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFundingRate(val *FundingRate) *NullableFundingRate {
	return &NullableFundingRate{value: val, isSet: true}
}

func (v NullableFundingRate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFundingRate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


