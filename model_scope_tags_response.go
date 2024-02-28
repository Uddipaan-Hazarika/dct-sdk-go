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

// checks if the ScopeTagsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScopeTagsResponse{}

// ScopeTagsResponse struct for ScopeTagsResponse
type ScopeTagsResponse struct {
	// Array of tags with key value pairs along with optional object_type and permissions
	Tags []ScopeTag `json:"tags,omitempty"`
}

// NewScopeTagsResponse instantiates a new ScopeTagsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScopeTagsResponse() *ScopeTagsResponse {
	this := ScopeTagsResponse{}
	return &this
}

// NewScopeTagsResponseWithDefaults instantiates a new ScopeTagsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScopeTagsResponseWithDefaults() *ScopeTagsResponse {
	this := ScopeTagsResponse{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ScopeTagsResponse) GetTags() []ScopeTag {
	if o == nil || IsNil(o.Tags) {
		var ret []ScopeTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScopeTagsResponse) GetTagsOk() ([]ScopeTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ScopeTagsResponse) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ScopeTag and assigns it to the Tags field.
func (o *ScopeTagsResponse) SetTags(v []ScopeTag) {
	o.Tags = v
}

func (o ScopeTagsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScopeTagsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableScopeTagsResponse struct {
	value *ScopeTagsResponse
	isSet bool
}

func (v NullableScopeTagsResponse) Get() *ScopeTagsResponse {
	return v.value
}

func (v *NullableScopeTagsResponse) Set(val *ScopeTagsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableScopeTagsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableScopeTagsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScopeTagsResponse(val *ScopeTagsResponse) *NullableScopeTagsResponse {
	return &NullableScopeTagsResponse{value: val, isSet: true}
}

func (v NullableScopeTagsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScopeTagsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


