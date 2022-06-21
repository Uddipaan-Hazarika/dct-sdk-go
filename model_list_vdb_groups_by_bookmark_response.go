/*
Delphix DCT API

Delphix DCT API

API version: 2.0.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// ListVDBGroupsByBookmarkResponse struct for ListVDBGroupsByBookmarkResponse
type ListVDBGroupsByBookmarkResponse struct {
	Items []VDBGroup `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewListVDBGroupsByBookmarkResponse instantiates a new ListVDBGroupsByBookmarkResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListVDBGroupsByBookmarkResponse() *ListVDBGroupsByBookmarkResponse {
	this := ListVDBGroupsByBookmarkResponse{}
	return &this
}

// NewListVDBGroupsByBookmarkResponseWithDefaults instantiates a new ListVDBGroupsByBookmarkResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListVDBGroupsByBookmarkResponseWithDefaults() *ListVDBGroupsByBookmarkResponse {
	this := ListVDBGroupsByBookmarkResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListVDBGroupsByBookmarkResponse) GetItems() []VDBGroup {
	if o == nil || o.Items == nil {
		var ret []VDBGroup
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListVDBGroupsByBookmarkResponse) GetItemsOk() ([]VDBGroup, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListVDBGroupsByBookmarkResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []VDBGroup and assigns it to the Items field.
func (o *ListVDBGroupsByBookmarkResponse) SetItems(v []VDBGroup) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *ListVDBGroupsByBookmarkResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || o.ResponseMetadata == nil {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListVDBGroupsByBookmarkResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || o.ResponseMetadata == nil {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *ListVDBGroupsByBookmarkResponse) HasResponseMetadata() bool {
	if o != nil && o.ResponseMetadata != nil {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *ListVDBGroupsByBookmarkResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o ListVDBGroupsByBookmarkResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.ResponseMetadata != nil {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return json.Marshal(toSerialize)
}

type NullableListVDBGroupsByBookmarkResponse struct {
	value *ListVDBGroupsByBookmarkResponse
	isSet bool
}

func (v NullableListVDBGroupsByBookmarkResponse) Get() *ListVDBGroupsByBookmarkResponse {
	return v.value
}

func (v *NullableListVDBGroupsByBookmarkResponse) Set(val *ListVDBGroupsByBookmarkResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListVDBGroupsByBookmarkResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListVDBGroupsByBookmarkResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListVDBGroupsByBookmarkResponse(val *ListVDBGroupsByBookmarkResponse) *NullableListVDBGroupsByBookmarkResponse {
	return &NullableListVDBGroupsByBookmarkResponse{value: val, isSet: true}
}

func (v NullableListVDBGroupsByBookmarkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListVDBGroupsByBookmarkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


