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

// checks if the SearchHashicorpVaultsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchHashicorpVaultsResponse{}

// SearchHashicorpVaultsResponse struct for SearchHashicorpVaultsResponse
type SearchHashicorpVaultsResponse struct {
	Items []HashicorpVault `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewSearchHashicorpVaultsResponse instantiates a new SearchHashicorpVaultsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchHashicorpVaultsResponse() *SearchHashicorpVaultsResponse {
	this := SearchHashicorpVaultsResponse{}
	return &this
}

// NewSearchHashicorpVaultsResponseWithDefaults instantiates a new SearchHashicorpVaultsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchHashicorpVaultsResponseWithDefaults() *SearchHashicorpVaultsResponse {
	this := SearchHashicorpVaultsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SearchHashicorpVaultsResponse) GetItems() []HashicorpVault {
	if o == nil || IsNil(o.Items) {
		var ret []HashicorpVault
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchHashicorpVaultsResponse) GetItemsOk() ([]HashicorpVault, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SearchHashicorpVaultsResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []HashicorpVault and assigns it to the Items field.
func (o *SearchHashicorpVaultsResponse) SetItems(v []HashicorpVault) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *SearchHashicorpVaultsResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchHashicorpVaultsResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *SearchHashicorpVaultsResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *SearchHashicorpVaultsResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o SearchHashicorpVaultsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchHashicorpVaultsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return toSerialize, nil
}

type NullableSearchHashicorpVaultsResponse struct {
	value *SearchHashicorpVaultsResponse
	isSet bool
}

func (v NullableSearchHashicorpVaultsResponse) Get() *SearchHashicorpVaultsResponse {
	return v.value
}

func (v *NullableSearchHashicorpVaultsResponse) Set(val *SearchHashicorpVaultsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchHashicorpVaultsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchHashicorpVaultsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchHashicorpVaultsResponse(val *SearchHashicorpVaultsResponse) *NullableSearchHashicorpVaultsResponse {
	return &NullableSearchHashicorpVaultsResponse{value: val, isSet: true}
}

func (v NullableSearchHashicorpVaultsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchHashicorpVaultsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


