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

// checks if the SearchBookmarksByVDBGroupsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchBookmarksByVDBGroupsResponse{}

// SearchBookmarksByVDBGroupsResponse struct for SearchBookmarksByVDBGroupsResponse
type SearchBookmarksByVDBGroupsResponse struct {
	Items []Bookmark `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewSearchBookmarksByVDBGroupsResponse instantiates a new SearchBookmarksByVDBGroupsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchBookmarksByVDBGroupsResponse() *SearchBookmarksByVDBGroupsResponse {
	this := SearchBookmarksByVDBGroupsResponse{}
	return &this
}

// NewSearchBookmarksByVDBGroupsResponseWithDefaults instantiates a new SearchBookmarksByVDBGroupsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchBookmarksByVDBGroupsResponseWithDefaults() *SearchBookmarksByVDBGroupsResponse {
	this := SearchBookmarksByVDBGroupsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SearchBookmarksByVDBGroupsResponse) GetItems() []Bookmark {
	if o == nil || IsNil(o.Items) {
		var ret []Bookmark
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBookmarksByVDBGroupsResponse) GetItemsOk() ([]Bookmark, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SearchBookmarksByVDBGroupsResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Bookmark and assigns it to the Items field.
func (o *SearchBookmarksByVDBGroupsResponse) SetItems(v []Bookmark) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *SearchBookmarksByVDBGroupsResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBookmarksByVDBGroupsResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *SearchBookmarksByVDBGroupsResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *SearchBookmarksByVDBGroupsResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o SearchBookmarksByVDBGroupsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchBookmarksByVDBGroupsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return toSerialize, nil
}

type NullableSearchBookmarksByVDBGroupsResponse struct {
	value *SearchBookmarksByVDBGroupsResponse
	isSet bool
}

func (v NullableSearchBookmarksByVDBGroupsResponse) Get() *SearchBookmarksByVDBGroupsResponse {
	return v.value
}

func (v *NullableSearchBookmarksByVDBGroupsResponse) Set(val *SearchBookmarksByVDBGroupsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchBookmarksByVDBGroupsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchBookmarksByVDBGroupsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchBookmarksByVDBGroupsResponse(val *SearchBookmarksByVDBGroupsResponse) *NullableSearchBookmarksByVDBGroupsResponse {
	return &NullableSearchBookmarksByVDBGroupsResponse{value: val, isSet: true}
}

func (v NullableSearchBookmarksByVDBGroupsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchBookmarksByVDBGroupsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


