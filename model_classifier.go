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

// checks if the Classifier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Classifier{}

// Classifier A classifier.
type Classifier struct {
	// The classifier ID.
	Id *string `json:"id,omitempty"`
	// The name of this classifier.
	Name *string `json:"name,omitempty"`
	// The framework of this classifier.
	Framework *string `json:"framework,omitempty"`
	// A description of this classifier.
	Description NullableString `json:"description,omitempty"`
	// The id of the data class associated with this classifier.
	DataClassId *string `json:"data_class_id,omitempty"`
	// The name of the data class associated with this classifier.
	DataClassName *string `json:"data_class_name,omitempty"`
	// The configuration of this algorithm.
	Config map[string]interface{} `json:"config,omitempty"`
	// The ID of the engine this classifier originated from.
	EngineId NullableString `json:"engine_id,omitempty"`
	// The name of the engine this classifier originated from.
	EngineName NullableString `json:"engine_name,omitempty"`
	// The tags of this classifier.
	Tags []Tag `json:"tags,omitempty"`
}

// NewClassifier instantiates a new Classifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClassifier() *Classifier {
	this := Classifier{}
	return &this
}

// NewClassifierWithDefaults instantiates a new Classifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClassifierWithDefaults() *Classifier {
	this := Classifier{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Classifier) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Classifier) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Classifier) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Classifier) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Classifier) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Classifier) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Classifier) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Classifier) SetName(v string) {
	o.Name = &v
}

// GetFramework returns the Framework field value if set, zero value otherwise.
func (o *Classifier) GetFramework() string {
	if o == nil || IsNil(o.Framework) {
		var ret string
		return ret
	}
	return *o.Framework
}

// GetFrameworkOk returns a tuple with the Framework field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Classifier) GetFrameworkOk() (*string, bool) {
	if o == nil || IsNil(o.Framework) {
		return nil, false
	}
	return o.Framework, true
}

// HasFramework returns a boolean if a field has been set.
func (o *Classifier) HasFramework() bool {
	if o != nil && !IsNil(o.Framework) {
		return true
	}

	return false
}

// SetFramework gets a reference to the given string and assigns it to the Framework field.
func (o *Classifier) SetFramework(v string) {
	o.Framework = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Classifier) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Classifier) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Classifier) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Classifier) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Classifier) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Classifier) UnsetDescription() {
	o.Description.Unset()
}

// GetDataClassId returns the DataClassId field value if set, zero value otherwise.
func (o *Classifier) GetDataClassId() string {
	if o == nil || IsNil(o.DataClassId) {
		var ret string
		return ret
	}
	return *o.DataClassId
}

// GetDataClassIdOk returns a tuple with the DataClassId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Classifier) GetDataClassIdOk() (*string, bool) {
	if o == nil || IsNil(o.DataClassId) {
		return nil, false
	}
	return o.DataClassId, true
}

// HasDataClassId returns a boolean if a field has been set.
func (o *Classifier) HasDataClassId() bool {
	if o != nil && !IsNil(o.DataClassId) {
		return true
	}

	return false
}

// SetDataClassId gets a reference to the given string and assigns it to the DataClassId field.
func (o *Classifier) SetDataClassId(v string) {
	o.DataClassId = &v
}

// GetDataClassName returns the DataClassName field value if set, zero value otherwise.
func (o *Classifier) GetDataClassName() string {
	if o == nil || IsNil(o.DataClassName) {
		var ret string
		return ret
	}
	return *o.DataClassName
}

// GetDataClassNameOk returns a tuple with the DataClassName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Classifier) GetDataClassNameOk() (*string, bool) {
	if o == nil || IsNil(o.DataClassName) {
		return nil, false
	}
	return o.DataClassName, true
}

// HasDataClassName returns a boolean if a field has been set.
func (o *Classifier) HasDataClassName() bool {
	if o != nil && !IsNil(o.DataClassName) {
		return true
	}

	return false
}

// SetDataClassName gets a reference to the given string and assigns it to the DataClassName field.
func (o *Classifier) SetDataClassName(v string) {
	o.DataClassName = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *Classifier) GetConfig() map[string]interface{} {
	if o == nil || IsNil(o.Config) {
		var ret map[string]interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Classifier) GetConfigOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Config) {
		return map[string]interface{}{}, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *Classifier) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *Classifier) SetConfig(v map[string]interface{}) {
	o.Config = v
}

// GetEngineId returns the EngineId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Classifier) GetEngineId() string {
	if o == nil || IsNil(o.EngineId.Get()) {
		var ret string
		return ret
	}
	return *o.EngineId.Get()
}

// GetEngineIdOk returns a tuple with the EngineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Classifier) GetEngineIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EngineId.Get(), o.EngineId.IsSet()
}

// HasEngineId returns a boolean if a field has been set.
func (o *Classifier) HasEngineId() bool {
	if o != nil && o.EngineId.IsSet() {
		return true
	}

	return false
}

// SetEngineId gets a reference to the given NullableString and assigns it to the EngineId field.
func (o *Classifier) SetEngineId(v string) {
	o.EngineId.Set(&v)
}
// SetEngineIdNil sets the value for EngineId to be an explicit nil
func (o *Classifier) SetEngineIdNil() {
	o.EngineId.Set(nil)
}

// UnsetEngineId ensures that no value is present for EngineId, not even an explicit nil
func (o *Classifier) UnsetEngineId() {
	o.EngineId.Unset()
}

// GetEngineName returns the EngineName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Classifier) GetEngineName() string {
	if o == nil || IsNil(o.EngineName.Get()) {
		var ret string
		return ret
	}
	return *o.EngineName.Get()
}

// GetEngineNameOk returns a tuple with the EngineName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Classifier) GetEngineNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EngineName.Get(), o.EngineName.IsSet()
}

// HasEngineName returns a boolean if a field has been set.
func (o *Classifier) HasEngineName() bool {
	if o != nil && o.EngineName.IsSet() {
		return true
	}

	return false
}

// SetEngineName gets a reference to the given NullableString and assigns it to the EngineName field.
func (o *Classifier) SetEngineName(v string) {
	o.EngineName.Set(&v)
}
// SetEngineNameNil sets the value for EngineName to be an explicit nil
func (o *Classifier) SetEngineNameNil() {
	o.EngineName.Set(nil)
}

// UnsetEngineName ensures that no value is present for EngineName, not even an explicit nil
func (o *Classifier) UnsetEngineName() {
	o.EngineName.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Classifier) GetTags() []Tag {
	if o == nil || IsNil(o.Tags) {
		var ret []Tag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Classifier) GetTagsOk() ([]Tag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Classifier) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []Tag and assigns it to the Tags field.
func (o *Classifier) SetTags(v []Tag) {
	o.Tags = v
}

func (o Classifier) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Classifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Framework) {
		toSerialize["framework"] = o.Framework
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.DataClassId) {
		toSerialize["data_class_id"] = o.DataClassId
	}
	if !IsNil(o.DataClassName) {
		toSerialize["data_class_name"] = o.DataClassName
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if o.EngineId.IsSet() {
		toSerialize["engine_id"] = o.EngineId.Get()
	}
	if o.EngineName.IsSet() {
		toSerialize["engine_name"] = o.EngineName.Get()
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableClassifier struct {
	value *Classifier
	isSet bool
}

func (v NullableClassifier) Get() *Classifier {
	return v.value
}

func (v *NullableClassifier) Set(val *Classifier) {
	v.value = val
	v.isSet = true
}

func (v NullableClassifier) IsSet() bool {
	return v.isSet
}

func (v *NullableClassifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClassifier(val *Classifier) *NullableClassifier {
	return &NullableClassifier{value: val, isSet: true}
}

func (v NullableClassifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClassifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

