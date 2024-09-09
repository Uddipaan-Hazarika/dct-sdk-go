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

// checks if the ClassifiersSearchResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClassifiersSearchResponse{}

// ClassifiersSearchResponse struct for ClassifiersSearchResponse
type ClassifiersSearchResponse struct {
	Items []Classifier `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewClassifiersSearchResponse instantiates a new ClassifiersSearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClassifiersSearchResponse() *ClassifiersSearchResponse {
	this := ClassifiersSearchResponse{}
	return &this
}

// NewClassifiersSearchResponseWithDefaults instantiates a new ClassifiersSearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClassifiersSearchResponseWithDefaults() *ClassifiersSearchResponse {
	this := ClassifiersSearchResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ClassifiersSearchResponse) GetItems() []Classifier {
	if o == nil || IsNil(o.Items) {
		var ret []Classifier
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassifiersSearchResponse) GetItemsOk() ([]Classifier, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ClassifiersSearchResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Classifier and assigns it to the Items field.
func (o *ClassifiersSearchResponse) SetItems(v []Classifier) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *ClassifiersSearchResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassifiersSearchResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *ClassifiersSearchResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *ClassifiersSearchResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o ClassifiersSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClassifiersSearchResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return toSerialize, nil
}

type NullableClassifiersSearchResponse struct {
	value *ClassifiersSearchResponse
	isSet bool
}

func (v NullableClassifiersSearchResponse) Get() *ClassifiersSearchResponse {
	return v.value
}

func (v *NullableClassifiersSearchResponse) Set(val *ClassifiersSearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableClassifiersSearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableClassifiersSearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClassifiersSearchResponse(val *ClassifiersSearchResponse) *NullableClassifiersSearchResponse {
	return &NullableClassifiersSearchResponse{value: val, isSet: true}
}

func (v NullableClassifiersSearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClassifiersSearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

