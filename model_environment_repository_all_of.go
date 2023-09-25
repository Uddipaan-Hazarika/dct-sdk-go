/*
Delphix DCT API

Delphix DCT API

API version: 3.5.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// checks if the EnvironmentRepositoryAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentRepositoryAllOf{}

// EnvironmentRepositoryAllOf struct for EnvironmentRepositoryAllOf
type EnvironmentRepositoryAllOf struct {
	// The environment ID.
	EnvironmentId *string `json:"environment_id,omitempty"`
}

// NewEnvironmentRepositoryAllOf instantiates a new EnvironmentRepositoryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentRepositoryAllOf() *EnvironmentRepositoryAllOf {
	this := EnvironmentRepositoryAllOf{}
	return &this
}

// NewEnvironmentRepositoryAllOfWithDefaults instantiates a new EnvironmentRepositoryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentRepositoryAllOfWithDefaults() *EnvironmentRepositoryAllOf {
	this := EnvironmentRepositoryAllOf{}
	return &this
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *EnvironmentRepositoryAllOf) GetEnvironmentId() string {
	if o == nil || IsNil(o.EnvironmentId) {
		var ret string
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentRepositoryAllOf) GetEnvironmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnvironmentId) {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *EnvironmentRepositoryAllOf) HasEnvironmentId() bool {
	if o != nil && !IsNil(o.EnvironmentId) {
		return true
	}

	return false
}

// SetEnvironmentId gets a reference to the given string and assigns it to the EnvironmentId field.
func (o *EnvironmentRepositoryAllOf) SetEnvironmentId(v string) {
	o.EnvironmentId = &v
}

func (o EnvironmentRepositoryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentRepositoryAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnvironmentId) {
		toSerialize["environment_id"] = o.EnvironmentId
	}
	return toSerialize, nil
}

type NullableEnvironmentRepositoryAllOf struct {
	value *EnvironmentRepositoryAllOf
	isSet bool
}

func (v NullableEnvironmentRepositoryAllOf) Get() *EnvironmentRepositoryAllOf {
	return v.value
}

func (v *NullableEnvironmentRepositoryAllOf) Set(val *EnvironmentRepositoryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentRepositoryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentRepositoryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentRepositoryAllOf(val *EnvironmentRepositoryAllOf) *NullableEnvironmentRepositoryAllOf {
	return &NullableEnvironmentRepositoryAllOf{value: val, isSet: true}
}

func (v NullableEnvironmentRepositoryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentRepositoryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

