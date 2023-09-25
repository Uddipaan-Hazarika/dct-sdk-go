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

// checks if the SearchHyperscaleDatasetsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchHyperscaleDatasetsResponse{}

// SearchHyperscaleDatasetsResponse struct for SearchHyperscaleDatasetsResponse
type SearchHyperscaleDatasetsResponse struct {
	Items []HyperscaleDataset `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewSearchHyperscaleDatasetsResponse instantiates a new SearchHyperscaleDatasetsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchHyperscaleDatasetsResponse() *SearchHyperscaleDatasetsResponse {
	this := SearchHyperscaleDatasetsResponse{}
	return &this
}

// NewSearchHyperscaleDatasetsResponseWithDefaults instantiates a new SearchHyperscaleDatasetsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchHyperscaleDatasetsResponseWithDefaults() *SearchHyperscaleDatasetsResponse {
	this := SearchHyperscaleDatasetsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SearchHyperscaleDatasetsResponse) GetItems() []HyperscaleDataset {
	if o == nil || IsNil(o.Items) {
		var ret []HyperscaleDataset
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchHyperscaleDatasetsResponse) GetItemsOk() ([]HyperscaleDataset, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SearchHyperscaleDatasetsResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []HyperscaleDataset and assigns it to the Items field.
func (o *SearchHyperscaleDatasetsResponse) SetItems(v []HyperscaleDataset) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *SearchHyperscaleDatasetsResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchHyperscaleDatasetsResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *SearchHyperscaleDatasetsResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *SearchHyperscaleDatasetsResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o SearchHyperscaleDatasetsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchHyperscaleDatasetsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return toSerialize, nil
}

type NullableSearchHyperscaleDatasetsResponse struct {
	value *SearchHyperscaleDatasetsResponse
	isSet bool
}

func (v NullableSearchHyperscaleDatasetsResponse) Get() *SearchHyperscaleDatasetsResponse {
	return v.value
}

func (v *NullableSearchHyperscaleDatasetsResponse) Set(val *SearchHyperscaleDatasetsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchHyperscaleDatasetsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchHyperscaleDatasetsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchHyperscaleDatasetsResponse(val *SearchHyperscaleDatasetsResponse) *NullableSearchHyperscaleDatasetsResponse {
	return &NullableSearchHyperscaleDatasetsResponse{value: val, isSet: true}
}

func (v NullableSearchHyperscaleDatasetsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchHyperscaleDatasetsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

