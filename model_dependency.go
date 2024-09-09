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

// checks if the Dependency type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dependency{}

// Dependency A dependency relationship.
type Dependency struct {
	// The ID of the child entity.
	ChildId *string `json:"child_id,omitempty"`
	// The name of the child entity.
	ChildName *string `json:"child_name,omitempty"`
	// The type of the child entity.
	ChildType *string `json:"child_type,omitempty"`
}

// NewDependency instantiates a new Dependency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDependency() *Dependency {
	this := Dependency{}
	return &this
}

// NewDependencyWithDefaults instantiates a new Dependency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDependencyWithDefaults() *Dependency {
	this := Dependency{}
	return &this
}

// GetChildId returns the ChildId field value if set, zero value otherwise.
func (o *Dependency) GetChildId() string {
	if o == nil || IsNil(o.ChildId) {
		var ret string
		return ret
	}
	return *o.ChildId
}

// GetChildIdOk returns a tuple with the ChildId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dependency) GetChildIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChildId) {
		return nil, false
	}
	return o.ChildId, true
}

// HasChildId returns a boolean if a field has been set.
func (o *Dependency) HasChildId() bool {
	if o != nil && !IsNil(o.ChildId) {
		return true
	}

	return false
}

// SetChildId gets a reference to the given string and assigns it to the ChildId field.
func (o *Dependency) SetChildId(v string) {
	o.ChildId = &v
}

// GetChildName returns the ChildName field value if set, zero value otherwise.
func (o *Dependency) GetChildName() string {
	if o == nil || IsNil(o.ChildName) {
		var ret string
		return ret
	}
	return *o.ChildName
}

// GetChildNameOk returns a tuple with the ChildName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dependency) GetChildNameOk() (*string, bool) {
	if o == nil || IsNil(o.ChildName) {
		return nil, false
	}
	return o.ChildName, true
}

// HasChildName returns a boolean if a field has been set.
func (o *Dependency) HasChildName() bool {
	if o != nil && !IsNil(o.ChildName) {
		return true
	}

	return false
}

// SetChildName gets a reference to the given string and assigns it to the ChildName field.
func (o *Dependency) SetChildName(v string) {
	o.ChildName = &v
}

// GetChildType returns the ChildType field value if set, zero value otherwise.
func (o *Dependency) GetChildType() string {
	if o == nil || IsNil(o.ChildType) {
		var ret string
		return ret
	}
	return *o.ChildType
}

// GetChildTypeOk returns a tuple with the ChildType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dependency) GetChildTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ChildType) {
		return nil, false
	}
	return o.ChildType, true
}

// HasChildType returns a boolean if a field has been set.
func (o *Dependency) HasChildType() bool {
	if o != nil && !IsNil(o.ChildType) {
		return true
	}

	return false
}

// SetChildType gets a reference to the given string and assigns it to the ChildType field.
func (o *Dependency) SetChildType(v string) {
	o.ChildType = &v
}

func (o Dependency) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dependency) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChildId) {
		toSerialize["child_id"] = o.ChildId
	}
	if !IsNil(o.ChildName) {
		toSerialize["child_name"] = o.ChildName
	}
	if !IsNil(o.ChildType) {
		toSerialize["child_type"] = o.ChildType
	}
	return toSerialize, nil
}

type NullableDependency struct {
	value *Dependency
	isSet bool
}

func (v NullableDependency) Get() *Dependency {
	return v.value
}

func (v *NullableDependency) Set(val *Dependency) {
	v.value = val
	v.isSet = true
}

func (v NullableDependency) IsSet() bool {
	return v.isSet
}

func (v *NullableDependency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDependency(val *Dependency) *NullableDependency {
	return &NullableDependency{value: val, isSet: true}
}

func (v NullableDependency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDependency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

