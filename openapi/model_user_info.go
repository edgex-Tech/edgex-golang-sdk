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

// checks if the UserInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserInfo{}

// UserInfo 用户信息和用户设置信息
type UserInfo struct {
	User *User `json:"user,omitempty"`
	UserPreference *UserPreference `json:"userPreference,omitempty"`
}

// NewUserInfo instantiates a new UserInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserInfo() *UserInfo {
	this := UserInfo{}
	return &this
}

// NewUserInfoWithDefaults instantiates a new UserInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserInfoWithDefaults() *UserInfo {
	this := UserInfo{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *UserInfo) GetUser() User {
	if o == nil || IsNil(o.User) {
		var ret User
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInfo) GetUserOk() (*User, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *UserInfo) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given User and assigns it to the User field.
func (o *UserInfo) SetUser(v User) {
	o.User = &v
}

// GetUserPreference returns the UserPreference field value if set, zero value otherwise.
func (o *UserInfo) GetUserPreference() UserPreference {
	if o == nil || IsNil(o.UserPreference) {
		var ret UserPreference
		return ret
	}
	return *o.UserPreference
}

// GetUserPreferenceOk returns a tuple with the UserPreference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInfo) GetUserPreferenceOk() (*UserPreference, bool) {
	if o == nil || IsNil(o.UserPreference) {
		return nil, false
	}
	return o.UserPreference, true
}

// HasUserPreference returns a boolean if a field has been set.
func (o *UserInfo) HasUserPreference() bool {
	if o != nil && !IsNil(o.UserPreference) {
		return true
	}

	return false
}

// SetUserPreference gets a reference to the given UserPreference and assigns it to the UserPreference field.
func (o *UserInfo) SetUserPreference(v UserPreference) {
	o.UserPreference = &v
}

func (o UserInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.UserPreference) {
		toSerialize["userPreference"] = o.UserPreference
	}
	return toSerialize, nil
}

type NullableUserInfo struct {
	value *UserInfo
	isSet bool
}

func (v NullableUserInfo) Get() *UserInfo {
	return v.value
}

func (v *NullableUserInfo) Set(val *UserInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUserInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUserInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserInfo(val *UserInfo) *NullableUserInfo {
	return &NullableUserInfo{value: val, isSet: true}
}

func (v NullableUserInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


