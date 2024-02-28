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

// checks if the SearchHyperscaleDatasetTablesOrFilesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchHyperscaleDatasetTablesOrFilesResponse{}

// SearchHyperscaleDatasetTablesOrFilesResponse struct for SearchHyperscaleDatasetTablesOrFilesResponse
type SearchHyperscaleDatasetTablesOrFilesResponse struct {
	Items []HyperscaleDatasetTableOrFile `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewSearchHyperscaleDatasetTablesOrFilesResponse instantiates a new SearchHyperscaleDatasetTablesOrFilesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchHyperscaleDatasetTablesOrFilesResponse() *SearchHyperscaleDatasetTablesOrFilesResponse {
	this := SearchHyperscaleDatasetTablesOrFilesResponse{}
	return &this
}

// NewSearchHyperscaleDatasetTablesOrFilesResponseWithDefaults instantiates a new SearchHyperscaleDatasetTablesOrFilesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchHyperscaleDatasetTablesOrFilesResponseWithDefaults() *SearchHyperscaleDatasetTablesOrFilesResponse {
	this := SearchHyperscaleDatasetTablesOrFilesResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SearchHyperscaleDatasetTablesOrFilesResponse) GetItems() []HyperscaleDatasetTableOrFile {
	if o == nil || IsNil(o.Items) {
		var ret []HyperscaleDatasetTableOrFile
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchHyperscaleDatasetTablesOrFilesResponse) GetItemsOk() ([]HyperscaleDatasetTableOrFile, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SearchHyperscaleDatasetTablesOrFilesResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []HyperscaleDatasetTableOrFile and assigns it to the Items field.
func (o *SearchHyperscaleDatasetTablesOrFilesResponse) SetItems(v []HyperscaleDatasetTableOrFile) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *SearchHyperscaleDatasetTablesOrFilesResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchHyperscaleDatasetTablesOrFilesResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *SearchHyperscaleDatasetTablesOrFilesResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *SearchHyperscaleDatasetTablesOrFilesResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o SearchHyperscaleDatasetTablesOrFilesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchHyperscaleDatasetTablesOrFilesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return toSerialize, nil
}

type NullableSearchHyperscaleDatasetTablesOrFilesResponse struct {
	value *SearchHyperscaleDatasetTablesOrFilesResponse
	isSet bool
}

func (v NullableSearchHyperscaleDatasetTablesOrFilesResponse) Get() *SearchHyperscaleDatasetTablesOrFilesResponse {
	return v.value
}

func (v *NullableSearchHyperscaleDatasetTablesOrFilesResponse) Set(val *SearchHyperscaleDatasetTablesOrFilesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchHyperscaleDatasetTablesOrFilesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchHyperscaleDatasetTablesOrFilesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchHyperscaleDatasetTablesOrFilesResponse(val *SearchHyperscaleDatasetTablesOrFilesResponse) *NullableSearchHyperscaleDatasetTablesOrFilesResponse {
	return &NullableSearchHyperscaleDatasetTablesOrFilesResponse{value: val, isSet: true}
}

func (v NullableSearchHyperscaleDatasetTablesOrFilesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchHyperscaleDatasetTablesOrFilesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


