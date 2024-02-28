/*
Delphix DCT API

Delphix DCT API

API version: 3.9.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// checks if the InstallMaskingPluginParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstallMaskingPluginParameters{}

// InstallMaskingPluginParameters Parameters to install a custom masking plugin.
type InstallMaskingPluginParameters struct {
	// File reference for the plugin JAR.
	PluginFileReference string `json:"plugin_file_reference"`
	// The tags of this plugin.
	Tags []Tag `json:"tags,omitempty"`
}

// NewInstallMaskingPluginParameters instantiates a new InstallMaskingPluginParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstallMaskingPluginParameters(pluginFileReference string) *InstallMaskingPluginParameters {
	this := InstallMaskingPluginParameters{}
	this.PluginFileReference = pluginFileReference
	return &this
}

// NewInstallMaskingPluginParametersWithDefaults instantiates a new InstallMaskingPluginParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstallMaskingPluginParametersWithDefaults() *InstallMaskingPluginParameters {
	this := InstallMaskingPluginParameters{}
	return &this
}

// GetPluginFileReference returns the PluginFileReference field value
func (o *InstallMaskingPluginParameters) GetPluginFileReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PluginFileReference
}

// GetPluginFileReferenceOk returns a tuple with the PluginFileReference field value
// and a boolean to check if the value has been set.
func (o *InstallMaskingPluginParameters) GetPluginFileReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PluginFileReference, true
}

// SetPluginFileReference sets field value
func (o *InstallMaskingPluginParameters) SetPluginFileReference(v string) {
	o.PluginFileReference = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InstallMaskingPluginParameters) GetTags() []Tag {
	if o == nil || IsNil(o.Tags) {
		var ret []Tag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallMaskingPluginParameters) GetTagsOk() ([]Tag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InstallMaskingPluginParameters) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []Tag and assigns it to the Tags field.
func (o *InstallMaskingPluginParameters) SetTags(v []Tag) {
	o.Tags = v
}

func (o InstallMaskingPluginParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstallMaskingPluginParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["plugin_file_reference"] = o.PluginFileReference
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableInstallMaskingPluginParameters struct {
	value *InstallMaskingPluginParameters
	isSet bool
}

func (v NullableInstallMaskingPluginParameters) Get() *InstallMaskingPluginParameters {
	return v.value
}

func (v *NullableInstallMaskingPluginParameters) Set(val *InstallMaskingPluginParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableInstallMaskingPluginParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableInstallMaskingPluginParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstallMaskingPluginParameters(val *InstallMaskingPluginParameters) *NullableInstallMaskingPluginParameters {
	return &NullableInstallMaskingPluginParameters{value: val, isSet: true}
}

func (v NullableInstallMaskingPluginParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstallMaskingPluginParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


