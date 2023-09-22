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

// checks if the ConfigSettingsStg type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigSettingsStg{}

// ConfigSettingsStg Custom Database-Level config settings.
type ConfigSettingsStg struct {
	// Name of the property.
	PropertyName *string `json:"property_name,omitempty"`
	// Value of the property.
	Value *string `json:"value,omitempty"`
	// Select this option to comment out the provided property name in the configuration file.
	CommentProperty *bool `json:"comment_property,omitempty"`
}

// NewConfigSettingsStg instantiates a new ConfigSettingsStg object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigSettingsStg() *ConfigSettingsStg {
	this := ConfigSettingsStg{}
	return &this
}

// NewConfigSettingsStgWithDefaults instantiates a new ConfigSettingsStg object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigSettingsStgWithDefaults() *ConfigSettingsStg {
	this := ConfigSettingsStg{}
	return &this
}

// GetPropertyName returns the PropertyName field value if set, zero value otherwise.
func (o *ConfigSettingsStg) GetPropertyName() string {
	if o == nil || IsNil(o.PropertyName) {
		var ret string
		return ret
	}
	return *o.PropertyName
}

// GetPropertyNameOk returns a tuple with the PropertyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigSettingsStg) GetPropertyNameOk() (*string, bool) {
	if o == nil || IsNil(o.PropertyName) {
		return nil, false
	}
	return o.PropertyName, true
}

// HasPropertyName returns a boolean if a field has been set.
func (o *ConfigSettingsStg) HasPropertyName() bool {
	if o != nil && !IsNil(o.PropertyName) {
		return true
	}

	return false
}

// SetPropertyName gets a reference to the given string and assigns it to the PropertyName field.
func (o *ConfigSettingsStg) SetPropertyName(v string) {
	o.PropertyName = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ConfigSettingsStg) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigSettingsStg) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ConfigSettingsStg) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ConfigSettingsStg) SetValue(v string) {
	o.Value = &v
}

// GetCommentProperty returns the CommentProperty field value if set, zero value otherwise.
func (o *ConfigSettingsStg) GetCommentProperty() bool {
	if o == nil || IsNil(o.CommentProperty) {
		var ret bool
		return ret
	}
	return *o.CommentProperty
}

// GetCommentPropertyOk returns a tuple with the CommentProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigSettingsStg) GetCommentPropertyOk() (*bool, bool) {
	if o == nil || IsNil(o.CommentProperty) {
		return nil, false
	}
	return o.CommentProperty, true
}

// HasCommentProperty returns a boolean if a field has been set.
func (o *ConfigSettingsStg) HasCommentProperty() bool {
	if o != nil && !IsNil(o.CommentProperty) {
		return true
	}

	return false
}

// SetCommentProperty gets a reference to the given bool and assigns it to the CommentProperty field.
func (o *ConfigSettingsStg) SetCommentProperty(v bool) {
	o.CommentProperty = &v
}

func (o ConfigSettingsStg) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigSettingsStg) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PropertyName) {
		toSerialize["property_name"] = o.PropertyName
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.CommentProperty) {
		toSerialize["comment_property"] = o.CommentProperty
	}
	return toSerialize, nil
}

type NullableConfigSettingsStg struct {
	value *ConfigSettingsStg
	isSet bool
}

func (v NullableConfigSettingsStg) Get() *ConfigSettingsStg {
	return v.value
}

func (v *NullableConfigSettingsStg) Set(val *ConfigSettingsStg) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigSettingsStg) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigSettingsStg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigSettingsStg(val *ConfigSettingsStg) *NullableConfigSettingsStg {
	return &NullableConfigSettingsStg{value: val, isSet: true}
}

func (v NullableConfigSettingsStg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigSettingsStg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


