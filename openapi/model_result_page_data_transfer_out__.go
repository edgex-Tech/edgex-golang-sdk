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

// checks if the ResultPageDataTransferOut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResultPageDataTransferOut{}

// ResultPageDataTransferOut struct for ResultPageDataTransferOut
type ResultPageDataTransferOut struct {
	// 状态码.成功返回为\"SUCCESS\",其他都为失败
	Code *string `json:"code,omitempty"`
	Data *PageDataTransferOut `json:"data,omitempty"`
	// 错误消息中的参数信息
	ErrorParam *map[string]string `json:"errorParam,omitempty"`
	// 服务器请求接收时间
	RequestTime *string `json:"requestTime,omitempty"`
	// 服务器响应返回时间
	ResponseTime *string `json:"responseTime,omitempty"`
	// 调用traceId
	TraceId *string `json:"traceId,omitempty"`
}

// NewResultPageDataTransferOut instantiates a new ResultPageDataTransferOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResultPageDataTransferOut() *ResultPageDataTransferOut {
	this := ResultPageDataTransferOut{}
	return &this
}

// NewResultPageDataTransferOutWithDefaults instantiates a new ResultPageDataTransferOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResultPageDataTransferOutWithDefaults() *ResultPageDataTransferOut {
	this := ResultPageDataTransferOut{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ResultPageDataTransferOut) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultPageDataTransferOut) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ResultPageDataTransferOut) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ResultPageDataTransferOut) SetCode(v string) {
	o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ResultPageDataTransferOut) GetData() PageDataTransferOut {
	if o == nil || IsNil(o.Data) {
		var ret PageDataTransferOut
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultPageDataTransferOut) GetDataOk() (*PageDataTransferOut, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ResultPageDataTransferOut) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PageDataTransferOut and assigns it to the Data field.
func (o *ResultPageDataTransferOut) SetData(v PageDataTransferOut) {
	o.Data = &v
}

// GetErrorParam returns the ErrorParam field value if set, zero value otherwise.
func (o *ResultPageDataTransferOut) GetErrorParam() map[string]string {
	if o == nil || IsNil(o.ErrorParam) {
		var ret map[string]string
		return ret
	}
	return *o.ErrorParam
}

// GetErrorParamOk returns a tuple with the ErrorParam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultPageDataTransferOut) GetErrorParamOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ErrorParam) {
		return nil, false
	}
	return o.ErrorParam, true
}

// HasErrorParam returns a boolean if a field has been set.
func (o *ResultPageDataTransferOut) HasErrorParam() bool {
	if o != nil && !IsNil(o.ErrorParam) {
		return true
	}

	return false
}

// SetErrorParam gets a reference to the given map[string]string and assigns it to the ErrorParam field.
func (o *ResultPageDataTransferOut) SetErrorParam(v map[string]string) {
	o.ErrorParam = &v
}

// GetRequestTime returns the RequestTime field value if set, zero value otherwise.
func (o *ResultPageDataTransferOut) GetRequestTime() string {
	if o == nil || IsNil(o.RequestTime) {
		var ret string
		return ret
	}
	return *o.RequestTime
}

// GetRequestTimeOk returns a tuple with the RequestTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultPageDataTransferOut) GetRequestTimeOk() (*string, bool) {
	if o == nil || IsNil(o.RequestTime) {
		return nil, false
	}
	return o.RequestTime, true
}

// HasRequestTime returns a boolean if a field has been set.
func (o *ResultPageDataTransferOut) HasRequestTime() bool {
	if o != nil && !IsNil(o.RequestTime) {
		return true
	}

	return false
}

// SetRequestTime gets a reference to the given string and assigns it to the RequestTime field.
func (o *ResultPageDataTransferOut) SetRequestTime(v string) {
	o.RequestTime = &v
}

// GetResponseTime returns the ResponseTime field value if set, zero value otherwise.
func (o *ResultPageDataTransferOut) GetResponseTime() string {
	if o == nil || IsNil(o.ResponseTime) {
		var ret string
		return ret
	}
	return *o.ResponseTime
}

// GetResponseTimeOk returns a tuple with the ResponseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultPageDataTransferOut) GetResponseTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ResponseTime) {
		return nil, false
	}
	return o.ResponseTime, true
}

// HasResponseTime returns a boolean if a field has been set.
func (o *ResultPageDataTransferOut) HasResponseTime() bool {
	if o != nil && !IsNil(o.ResponseTime) {
		return true
	}

	return false
}

// SetResponseTime gets a reference to the given string and assigns it to the ResponseTime field.
func (o *ResultPageDataTransferOut) SetResponseTime(v string) {
	o.ResponseTime = &v
}

// GetTraceId returns the TraceId field value if set, zero value otherwise.
func (o *ResultPageDataTransferOut) GetTraceId() string {
	if o == nil || IsNil(o.TraceId) {
		var ret string
		return ret
	}
	return *o.TraceId
}

// GetTraceIdOk returns a tuple with the TraceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultPageDataTransferOut) GetTraceIdOk() (*string, bool) {
	if o == nil || IsNil(o.TraceId) {
		return nil, false
	}
	return o.TraceId, true
}

// HasTraceId returns a boolean if a field has been set.
func (o *ResultPageDataTransferOut) HasTraceId() bool {
	if o != nil && !IsNil(o.TraceId) {
		return true
	}

	return false
}

// SetTraceId gets a reference to the given string and assigns it to the TraceId field.
func (o *ResultPageDataTransferOut) SetTraceId(v string) {
	o.TraceId = &v
}

func (o ResultPageDataTransferOut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResultPageDataTransferOut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.ErrorParam) {
		toSerialize["errorParam"] = o.ErrorParam
	}
	if !IsNil(o.RequestTime) {
		toSerialize["requestTime"] = o.RequestTime
	}
	if !IsNil(o.ResponseTime) {
		toSerialize["responseTime"] = o.ResponseTime
	}
	if !IsNil(o.TraceId) {
		toSerialize["traceId"] = o.TraceId
	}
	return toSerialize, nil
}

type NullableResultPageDataTransferOut struct {
	value *ResultPageDataTransferOut
	isSet bool
}

func (v NullableResultPageDataTransferOut) Get() *ResultPageDataTransferOut {
	return v.value
}

func (v *NullableResultPageDataTransferOut) Set(val *ResultPageDataTransferOut) {
	v.value = val
	v.isSet = true
}

func (v NullableResultPageDataTransferOut) IsSet() bool {
	return v.isSet
}

func (v *NullableResultPageDataTransferOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultPageDataTransferOut(val *ResultPageDataTransferOut) *NullableResultPageDataTransferOut {
	return &NullableResultPageDataTransferOut{value: val, isSet: true}
}

func (v NullableResultPageDataTransferOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResultPageDataTransferOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


