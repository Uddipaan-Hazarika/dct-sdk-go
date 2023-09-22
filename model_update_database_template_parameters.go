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

// checks if the UpdateDatabaseTemplateParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDatabaseTemplateParameters{}

// UpdateDatabaseTemplateParameters Parameters to update a Database Template.
type UpdateDatabaseTemplateParameters struct {
	// The DatabaseTemplate name.
	Name *string `json:"name,omitempty"`
	// User provided description for this template.
	Description *string `json:"description,omitempty"`
	// The type of the source associated with the template.
	SourceType *string `json:"source_type,omitempty"`
	// A name/value map of string configuration parameters.
	Parameters *map[string]string `json:"parameters,omitempty"`
}

// NewUpdateDatabaseTemplateParameters instantiates a new UpdateDatabaseTemplateParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDatabaseTemplateParameters() *UpdateDatabaseTemplateParameters {
	this := UpdateDatabaseTemplateParameters{}
	return &this
}

// NewUpdateDatabaseTemplateParametersWithDefaults instantiates a new UpdateDatabaseTemplateParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDatabaseTemplateParametersWithDefaults() *UpdateDatabaseTemplateParameters {
	this := UpdateDatabaseTemplateParameters{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateDatabaseTemplateParameters) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDatabaseTemplateParameters) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateDatabaseTemplateParameters) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateDatabaseTemplateParameters) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateDatabaseTemplateParameters) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDatabaseTemplateParameters) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateDatabaseTemplateParameters) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateDatabaseTemplateParameters) SetDescription(v string) {
	o.Description = &v
}

// GetSourceType returns the SourceType field value if set, zero value otherwise.
func (o *UpdateDatabaseTemplateParameters) GetSourceType() string {
	if o == nil || IsNil(o.SourceType) {
		var ret string
		return ret
	}
	return *o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDatabaseTemplateParameters) GetSourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SourceType) {
		return nil, false
	}
	return o.SourceType, true
}

// HasSourceType returns a boolean if a field has been set.
func (o *UpdateDatabaseTemplateParameters) HasSourceType() bool {
	if o != nil && !IsNil(o.SourceType) {
		return true
	}

	return false
}

// SetSourceType gets a reference to the given string and assigns it to the SourceType field.
func (o *UpdateDatabaseTemplateParameters) SetSourceType(v string) {
	o.SourceType = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *UpdateDatabaseTemplateParameters) GetParameters() map[string]string {
	if o == nil || IsNil(o.Parameters) {
		var ret map[string]string
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDatabaseTemplateParameters) GetParametersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *UpdateDatabaseTemplateParameters) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]string and assigns it to the Parameters field.
func (o *UpdateDatabaseTemplateParameters) SetParameters(v map[string]string) {
	o.Parameters = &v
}

func (o UpdateDatabaseTemplateParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDatabaseTemplateParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.SourceType) {
		toSerialize["source_type"] = o.SourceType
	}
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	return toSerialize, nil
}

type NullableUpdateDatabaseTemplateParameters struct {
	value *UpdateDatabaseTemplateParameters
	isSet bool
}

func (v NullableUpdateDatabaseTemplateParameters) Get() *UpdateDatabaseTemplateParameters {
	return v.value
}

func (v *NullableUpdateDatabaseTemplateParameters) Set(val *UpdateDatabaseTemplateParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDatabaseTemplateParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDatabaseTemplateParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDatabaseTemplateParameters(val *UpdateDatabaseTemplateParameters) *NullableUpdateDatabaseTemplateParameters {
	return &NullableUpdateDatabaseTemplateParameters{value: val, isSet: true}
}

func (v NullableUpdateDatabaseTemplateParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDatabaseTemplateParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


