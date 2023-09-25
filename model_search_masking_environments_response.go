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

// checks if the SearchMaskingEnvironmentsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchMaskingEnvironmentsResponse{}

// SearchMaskingEnvironmentsResponse struct for SearchMaskingEnvironmentsResponse
type SearchMaskingEnvironmentsResponse struct {
	Items []MaskingEnvironment `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewSearchMaskingEnvironmentsResponse instantiates a new SearchMaskingEnvironmentsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchMaskingEnvironmentsResponse() *SearchMaskingEnvironmentsResponse {
	this := SearchMaskingEnvironmentsResponse{}
	return &this
}

// NewSearchMaskingEnvironmentsResponseWithDefaults instantiates a new SearchMaskingEnvironmentsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchMaskingEnvironmentsResponseWithDefaults() *SearchMaskingEnvironmentsResponse {
	this := SearchMaskingEnvironmentsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SearchMaskingEnvironmentsResponse) GetItems() []MaskingEnvironment {
	if o == nil || IsNil(o.Items) {
		var ret []MaskingEnvironment
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchMaskingEnvironmentsResponse) GetItemsOk() ([]MaskingEnvironment, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SearchMaskingEnvironmentsResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []MaskingEnvironment and assigns it to the Items field.
func (o *SearchMaskingEnvironmentsResponse) SetItems(v []MaskingEnvironment) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *SearchMaskingEnvironmentsResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchMaskingEnvironmentsResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *SearchMaskingEnvironmentsResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *SearchMaskingEnvironmentsResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o SearchMaskingEnvironmentsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchMaskingEnvironmentsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return toSerialize, nil
}

type NullableSearchMaskingEnvironmentsResponse struct {
	value *SearchMaskingEnvironmentsResponse
	isSet bool
}

func (v NullableSearchMaskingEnvironmentsResponse) Get() *SearchMaskingEnvironmentsResponse {
	return v.value
}

func (v *NullableSearchMaskingEnvironmentsResponse) Set(val *SearchMaskingEnvironmentsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchMaskingEnvironmentsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchMaskingEnvironmentsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchMaskingEnvironmentsResponse(val *SearchMaskingEnvironmentsResponse) *NullableSearchMaskingEnvironmentsResponse {
	return &NullableSearchMaskingEnvironmentsResponse{value: val, isSet: true}
}

func (v NullableSearchMaskingEnvironmentsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchMaskingEnvironmentsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

