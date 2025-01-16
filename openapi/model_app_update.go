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

// checks if the AppUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppUpdate{}

// AppUpdate APP升级信息
type AppUpdate struct {
	LatestVersion *string `json:"latestVersion,omitempty"`
	DownloadUrl *string `json:"downloadUrl,omitempty"`
	MarketUrl *string `json:"marketUrl,omitempty"`
	Description *string `json:"description,omitempty"`
	BeForceUpdate *bool `json:"beForceUpdate,omitempty"`
	BeUpdate *bool `json:"beUpdate,omitempty"`
}

// NewAppUpdate instantiates a new AppUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppUpdate() *AppUpdate {
	this := AppUpdate{}
	return &this
}

// NewAppUpdateWithDefaults instantiates a new AppUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppUpdateWithDefaults() *AppUpdate {
	this := AppUpdate{}
	return &this
}

// GetLatestVersion returns the LatestVersion field value if set, zero value otherwise.
func (o *AppUpdate) GetLatestVersion() string {
	if o == nil || IsNil(o.LatestVersion) {
		var ret string
		return ret
	}
	return *o.LatestVersion
}

// GetLatestVersionOk returns a tuple with the LatestVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUpdate) GetLatestVersionOk() (*string, bool) {
	if o == nil || IsNil(o.LatestVersion) {
		return nil, false
	}
	return o.LatestVersion, true
}

// HasLatestVersion returns a boolean if a field has been set.
func (o *AppUpdate) HasLatestVersion() bool {
	if o != nil && !IsNil(o.LatestVersion) {
		return true
	}

	return false
}

// SetLatestVersion gets a reference to the given string and assigns it to the LatestVersion field.
func (o *AppUpdate) SetLatestVersion(v string) {
	o.LatestVersion = &v
}

// GetDownloadUrl returns the DownloadUrl field value if set, zero value otherwise.
func (o *AppUpdate) GetDownloadUrl() string {
	if o == nil || IsNil(o.DownloadUrl) {
		var ret string
		return ret
	}
	return *o.DownloadUrl
}

// GetDownloadUrlOk returns a tuple with the DownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUpdate) GetDownloadUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DownloadUrl) {
		return nil, false
	}
	return o.DownloadUrl, true
}

// HasDownloadUrl returns a boolean if a field has been set.
func (o *AppUpdate) HasDownloadUrl() bool {
	if o != nil && !IsNil(o.DownloadUrl) {
		return true
	}

	return false
}

// SetDownloadUrl gets a reference to the given string and assigns it to the DownloadUrl field.
func (o *AppUpdate) SetDownloadUrl(v string) {
	o.DownloadUrl = &v
}

// GetMarketUrl returns the MarketUrl field value if set, zero value otherwise.
func (o *AppUpdate) GetMarketUrl() string {
	if o == nil || IsNil(o.MarketUrl) {
		var ret string
		return ret
	}
	return *o.MarketUrl
}

// GetMarketUrlOk returns a tuple with the MarketUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUpdate) GetMarketUrlOk() (*string, bool) {
	if o == nil || IsNil(o.MarketUrl) {
		return nil, false
	}
	return o.MarketUrl, true
}

// HasMarketUrl returns a boolean if a field has been set.
func (o *AppUpdate) HasMarketUrl() bool {
	if o != nil && !IsNil(o.MarketUrl) {
		return true
	}

	return false
}

// SetMarketUrl gets a reference to the given string and assigns it to the MarketUrl field.
func (o *AppUpdate) SetMarketUrl(v string) {
	o.MarketUrl = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AppUpdate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AppUpdate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AppUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetBeForceUpdate returns the BeForceUpdate field value if set, zero value otherwise.
func (o *AppUpdate) GetBeForceUpdate() bool {
	if o == nil || IsNil(o.BeForceUpdate) {
		var ret bool
		return ret
	}
	return *o.BeForceUpdate
}

// GetBeForceUpdateOk returns a tuple with the BeForceUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUpdate) GetBeForceUpdateOk() (*bool, bool) {
	if o == nil || IsNil(o.BeForceUpdate) {
		return nil, false
	}
	return o.BeForceUpdate, true
}

// HasBeForceUpdate returns a boolean if a field has been set.
func (o *AppUpdate) HasBeForceUpdate() bool {
	if o != nil && !IsNil(o.BeForceUpdate) {
		return true
	}

	return false
}

// SetBeForceUpdate gets a reference to the given bool and assigns it to the BeForceUpdate field.
func (o *AppUpdate) SetBeForceUpdate(v bool) {
	o.BeForceUpdate = &v
}

// GetBeUpdate returns the BeUpdate field value if set, zero value otherwise.
func (o *AppUpdate) GetBeUpdate() bool {
	if o == nil || IsNil(o.BeUpdate) {
		var ret bool
		return ret
	}
	return *o.BeUpdate
}

// GetBeUpdateOk returns a tuple with the BeUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUpdate) GetBeUpdateOk() (*bool, bool) {
	if o == nil || IsNil(o.BeUpdate) {
		return nil, false
	}
	return o.BeUpdate, true
}

// HasBeUpdate returns a boolean if a field has been set.
func (o *AppUpdate) HasBeUpdate() bool {
	if o != nil && !IsNil(o.BeUpdate) {
		return true
	}

	return false
}

// SetBeUpdate gets a reference to the given bool and assigns it to the BeUpdate field.
func (o *AppUpdate) SetBeUpdate(v bool) {
	o.BeUpdate = &v
}

func (o AppUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LatestVersion) {
		toSerialize["latestVersion"] = o.LatestVersion
	}
	if !IsNil(o.DownloadUrl) {
		toSerialize["downloadUrl"] = o.DownloadUrl
	}
	if !IsNil(o.MarketUrl) {
		toSerialize["marketUrl"] = o.MarketUrl
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.BeForceUpdate) {
		toSerialize["beForceUpdate"] = o.BeForceUpdate
	}
	if !IsNil(o.BeUpdate) {
		toSerialize["beUpdate"] = o.BeUpdate
	}
	return toSerialize, nil
}

type NullableAppUpdate struct {
	value *AppUpdate
	isSet bool
}

func (v NullableAppUpdate) Get() *AppUpdate {
	return v.value
}

func (v *NullableAppUpdate) Set(val *AppUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableAppUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableAppUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppUpdate(val *AppUpdate) *NullableAppUpdate {
	return &NullableAppUpdate{value: val, isSet: true}
}

func (v NullableAppUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


