/*
Delphix DCT API

Delphix DCT API

API version: 3.16.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// checks if the ComplianceApplicationSetting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComplianceApplicationSetting{}

// ComplianceApplicationSetting An application setting of a compliance engine.
type ComplianceApplicationSetting struct {
	// The ID of the application setting.
	Id *string `json:"id,omitempty"`
	// The group of the application setting.
	Group *string `json:"group,omitempty"`
	// The name of the application setting.
	Name *string `json:"name,omitempty"`
	// The value of the application setting.
	Value *string `json:"value,omitempty"`
}

// NewComplianceApplicationSetting instantiates a new ComplianceApplicationSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComplianceApplicationSetting() *ComplianceApplicationSetting {
	this := ComplianceApplicationSetting{}
	return &this
}

// NewComplianceApplicationSettingWithDefaults instantiates a new ComplianceApplicationSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComplianceApplicationSettingWithDefaults() *ComplianceApplicationSetting {
	this := ComplianceApplicationSetting{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ComplianceApplicationSetting) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplianceApplicationSetting) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ComplianceApplicationSetting) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ComplianceApplicationSetting) SetId(v string) {
	o.Id = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *ComplianceApplicationSetting) GetGroup() string {
	if o == nil || IsNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplianceApplicationSetting) GetGroupOk() (*string, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *ComplianceApplicationSetting) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *ComplianceApplicationSetting) SetGroup(v string) {
	o.Group = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ComplianceApplicationSetting) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplianceApplicationSetting) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ComplianceApplicationSetting) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ComplianceApplicationSetting) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ComplianceApplicationSetting) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplianceApplicationSetting) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ComplianceApplicationSetting) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ComplianceApplicationSetting) SetValue(v string) {
	o.Value = &v
}

func (o ComplianceApplicationSetting) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComplianceApplicationSetting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableComplianceApplicationSetting struct {
	value *ComplianceApplicationSetting
	isSet bool
}

func (v NullableComplianceApplicationSetting) Get() *ComplianceApplicationSetting {
	return v.value
}

func (v *NullableComplianceApplicationSetting) Set(val *ComplianceApplicationSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableComplianceApplicationSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableComplianceApplicationSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComplianceApplicationSetting(val *ComplianceApplicationSetting) *NullableComplianceApplicationSetting {
	return &NullableComplianceApplicationSetting{value: val, isSet: true}
}

func (v NullableComplianceApplicationSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComplianceApplicationSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

