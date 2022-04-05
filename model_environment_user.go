/*
Delphix DCT API

Delphix DCT API

API version: 1.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// EnvironmentUser struct for EnvironmentUser
type EnvironmentUser struct {
	// Environment user reference
	UserRef *string `json:"user_ref,omitempty"`
	// Username of environment user
	Username *string `json:"username,omitempty"`
	// This indicates if this user is primary or not
	PrimaryUser *bool `json:"primary_user,omitempty"`
}

// NewEnvironmentUser instantiates a new EnvironmentUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentUser() *EnvironmentUser {
	this := EnvironmentUser{}
	return &this
}

// NewEnvironmentUserWithDefaults instantiates a new EnvironmentUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentUserWithDefaults() *EnvironmentUser {
	this := EnvironmentUser{}
	return &this
}

// GetUserRef returns the UserRef field value if set, zero value otherwise.
func (o *EnvironmentUser) GetUserRef() string {
	if o == nil || o.UserRef == nil {
		var ret string
		return ret
	}
	return *o.UserRef
}

// GetUserRefOk returns a tuple with the UserRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUser) GetUserRefOk() (*string, bool) {
	if o == nil || o.UserRef == nil {
		return nil, false
	}
	return o.UserRef, true
}

// HasUserRef returns a boolean if a field has been set.
func (o *EnvironmentUser) HasUserRef() bool {
	if o != nil && o.UserRef != nil {
		return true
	}

	return false
}

// SetUserRef gets a reference to the given string and assigns it to the UserRef field.
func (o *EnvironmentUser) SetUserRef(v string) {
	o.UserRef = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *EnvironmentUser) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUser) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *EnvironmentUser) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *EnvironmentUser) SetUsername(v string) {
	o.Username = &v
}

// GetPrimaryUser returns the PrimaryUser field value if set, zero value otherwise.
func (o *EnvironmentUser) GetPrimaryUser() bool {
	if o == nil || o.PrimaryUser == nil {
		var ret bool
		return ret
	}
	return *o.PrimaryUser
}

// GetPrimaryUserOk returns a tuple with the PrimaryUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUser) GetPrimaryUserOk() (*bool, bool) {
	if o == nil || o.PrimaryUser == nil {
		return nil, false
	}
	return o.PrimaryUser, true
}

// HasPrimaryUser returns a boolean if a field has been set.
func (o *EnvironmentUser) HasPrimaryUser() bool {
	if o != nil && o.PrimaryUser != nil {
		return true
	}

	return false
}

// SetPrimaryUser gets a reference to the given bool and assigns it to the PrimaryUser field.
func (o *EnvironmentUser) SetPrimaryUser(v bool) {
	o.PrimaryUser = &v
}

func (o EnvironmentUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UserRef != nil {
		toSerialize["user_ref"] = o.UserRef
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.PrimaryUser != nil {
		toSerialize["primary_user"] = o.PrimaryUser
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentUser struct {
	value *EnvironmentUser
	isSet bool
}

func (v NullableEnvironmentUser) Get() *EnvironmentUser {
	return v.value
}

func (v *NullableEnvironmentUser) Set(val *EnvironmentUser) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentUser) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentUser(val *EnvironmentUser) *NullableEnvironmentUser {
	return &NullableEnvironmentUser{value: val, isSet: true}
}

func (v NullableEnvironmentUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


