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

// checks if the OraclePriceSignature type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OraclePriceSignature{}

// OraclePriceSignature 预言机价格签名信息
type OraclePriceSignature struct {
	// 合约id
	ContractId *string `json:"contractId,omitempty"`
	// 签名者标识
	Signer *string `json:"signer,omitempty"`
	// 签名价格 (按照stark ex 精度处理后的价格)
	Price *string `json:"price,omitempty"`
	// Concatenation of the asset name and the oracle name (both in hex encoding).
	ExternalAssetId *string `json:"externalAssetId,omitempty"`
	Signature *L2Signature `json:"signature,omitempty"`
	// The time the signature was created.
	Timestamp *string `json:"timestamp,omitempty"`
}

// NewOraclePriceSignature instantiates a new OraclePriceSignature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOraclePriceSignature() *OraclePriceSignature {
	this := OraclePriceSignature{}
	return &this
}

// NewOraclePriceSignatureWithDefaults instantiates a new OraclePriceSignature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOraclePriceSignatureWithDefaults() *OraclePriceSignature {
	this := OraclePriceSignature{}
	return &this
}

// GetContractId returns the ContractId field value if set, zero value otherwise.
func (o *OraclePriceSignature) GetContractId() string {
	if o == nil || IsNil(o.ContractId) {
		var ret string
		return ret
	}
	return *o.ContractId
}

// GetContractIdOk returns a tuple with the ContractId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OraclePriceSignature) GetContractIdOk() (*string, bool) {
	if o == nil || IsNil(o.ContractId) {
		return nil, false
	}
	return o.ContractId, true
}

// HasContractId returns a boolean if a field has been set.
func (o *OraclePriceSignature) HasContractId() bool {
	if o != nil && !IsNil(o.ContractId) {
		return true
	}

	return false
}

// SetContractId gets a reference to the given string and assigns it to the ContractId field.
func (o *OraclePriceSignature) SetContractId(v string) {
	o.ContractId = &v
}

// GetSigner returns the Signer field value if set, zero value otherwise.
func (o *OraclePriceSignature) GetSigner() string {
	if o == nil || IsNil(o.Signer) {
		var ret string
		return ret
	}
	return *o.Signer
}

// GetSignerOk returns a tuple with the Signer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OraclePriceSignature) GetSignerOk() (*string, bool) {
	if o == nil || IsNil(o.Signer) {
		return nil, false
	}
	return o.Signer, true
}

// HasSigner returns a boolean if a field has been set.
func (o *OraclePriceSignature) HasSigner() bool {
	if o != nil && !IsNil(o.Signer) {
		return true
	}

	return false
}

// SetSigner gets a reference to the given string and assigns it to the Signer field.
func (o *OraclePriceSignature) SetSigner(v string) {
	o.Signer = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *OraclePriceSignature) GetPrice() string {
	if o == nil || IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OraclePriceSignature) GetPriceOk() (*string, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *OraclePriceSignature) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *OraclePriceSignature) SetPrice(v string) {
	o.Price = &v
}

// GetExternalAssetId returns the ExternalAssetId field value if set, zero value otherwise.
func (o *OraclePriceSignature) GetExternalAssetId() string {
	if o == nil || IsNil(o.ExternalAssetId) {
		var ret string
		return ret
	}
	return *o.ExternalAssetId
}

// GetExternalAssetIdOk returns a tuple with the ExternalAssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OraclePriceSignature) GetExternalAssetIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalAssetId) {
		return nil, false
	}
	return o.ExternalAssetId, true
}

// HasExternalAssetId returns a boolean if a field has been set.
func (o *OraclePriceSignature) HasExternalAssetId() bool {
	if o != nil && !IsNil(o.ExternalAssetId) {
		return true
	}

	return false
}

// SetExternalAssetId gets a reference to the given string and assigns it to the ExternalAssetId field.
func (o *OraclePriceSignature) SetExternalAssetId(v string) {
	o.ExternalAssetId = &v
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *OraclePriceSignature) GetSignature() L2Signature {
	if o == nil || IsNil(o.Signature) {
		var ret L2Signature
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OraclePriceSignature) GetSignatureOk() (*L2Signature, bool) {
	if o == nil || IsNil(o.Signature) {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *OraclePriceSignature) HasSignature() bool {
	if o != nil && !IsNil(o.Signature) {
		return true
	}

	return false
}

// SetSignature gets a reference to the given L2Signature and assigns it to the Signature field.
func (o *OraclePriceSignature) SetSignature(v L2Signature) {
	o.Signature = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *OraclePriceSignature) GetTimestamp() string {
	if o == nil || IsNil(o.Timestamp) {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OraclePriceSignature) GetTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *OraclePriceSignature) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *OraclePriceSignature) SetTimestamp(v string) {
	o.Timestamp = &v
}

func (o OraclePriceSignature) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OraclePriceSignature) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContractId) {
		toSerialize["contractId"] = o.ContractId
	}
	if !IsNil(o.Signer) {
		toSerialize["signer"] = o.Signer
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.ExternalAssetId) {
		toSerialize["externalAssetId"] = o.ExternalAssetId
	}
	if !IsNil(o.Signature) {
		toSerialize["signature"] = o.Signature
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	return toSerialize, nil
}

type NullableOraclePriceSignature struct {
	value *OraclePriceSignature
	isSet bool
}

func (v NullableOraclePriceSignature) Get() *OraclePriceSignature {
	return v.value
}

func (v *NullableOraclePriceSignature) Set(val *OraclePriceSignature) {
	v.value = val
	v.isSet = true
}

func (v NullableOraclePriceSignature) IsSet() bool {
	return v.isSet
}

func (v *NullableOraclePriceSignature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOraclePriceSignature(val *OraclePriceSignature) *NullableOraclePriceSignature {
	return &NullableOraclePriceSignature{value: val, isSet: true}
}

func (v NullableOraclePriceSignature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOraclePriceSignature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


