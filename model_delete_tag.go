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

// checks if the DeleteTag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteTag{}

// DeleteTag struct for DeleteTag
type DeleteTag struct {
	// Key of the tag
	Key *string `json:"key,omitempty"`
	// Value of the tag
	Value *string `json:"value,omitempty"`
	// List of tags to be deleted
	Tags []Tag `json:"tags,omitempty"`
}

// NewDeleteTag instantiates a new DeleteTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteTag() *DeleteTag {
	this := DeleteTag{}
	return &this
}

// NewDeleteTagWithDefaults instantiates a new DeleteTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteTagWithDefaults() *DeleteTag {
	this := DeleteTag{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *DeleteTag) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteTag) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *DeleteTag) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *DeleteTag) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DeleteTag) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteTag) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DeleteTag) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *DeleteTag) SetValue(v string) {
	o.Value = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DeleteTag) GetTags() []Tag {
	if o == nil || IsNil(o.Tags) {
		var ret []Tag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteTag) GetTagsOk() ([]Tag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DeleteTag) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []Tag and assigns it to the Tags field.
func (o *DeleteTag) SetTags(v []Tag) {
	o.Tags = v
}

func (o DeleteTag) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteTag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableDeleteTag struct {
	value *DeleteTag
	isSet bool
}

func (v NullableDeleteTag) Get() *DeleteTag {
	return v.value
}

func (v *NullableDeleteTag) Set(val *DeleteTag) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteTag) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteTag(val *DeleteTag) *NullableDeleteTag {
	return &NullableDeleteTag{value: val, isSet: true}
}

func (v NullableDeleteTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


