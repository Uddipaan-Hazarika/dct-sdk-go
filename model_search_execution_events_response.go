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

// checks if the SearchExecutionEventsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchExecutionEventsResponse{}

// SearchExecutionEventsResponse struct for SearchExecutionEventsResponse
type SearchExecutionEventsResponse struct {
	Items []ExecutionEvent `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewSearchExecutionEventsResponse instantiates a new SearchExecutionEventsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchExecutionEventsResponse() *SearchExecutionEventsResponse {
	this := SearchExecutionEventsResponse{}
	return &this
}

// NewSearchExecutionEventsResponseWithDefaults instantiates a new SearchExecutionEventsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchExecutionEventsResponseWithDefaults() *SearchExecutionEventsResponse {
	this := SearchExecutionEventsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SearchExecutionEventsResponse) GetItems() []ExecutionEvent {
	if o == nil || IsNil(o.Items) {
		var ret []ExecutionEvent
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchExecutionEventsResponse) GetItemsOk() ([]ExecutionEvent, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SearchExecutionEventsResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ExecutionEvent and assigns it to the Items field.
func (o *SearchExecutionEventsResponse) SetItems(v []ExecutionEvent) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *SearchExecutionEventsResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchExecutionEventsResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *SearchExecutionEventsResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *SearchExecutionEventsResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o SearchExecutionEventsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchExecutionEventsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return toSerialize, nil
}

type NullableSearchExecutionEventsResponse struct {
	value *SearchExecutionEventsResponse
	isSet bool
}

func (v NullableSearchExecutionEventsResponse) Get() *SearchExecutionEventsResponse {
	return v.value
}

func (v *NullableSearchExecutionEventsResponse) Set(val *SearchExecutionEventsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchExecutionEventsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchExecutionEventsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchExecutionEventsResponse(val *SearchExecutionEventsResponse) *NullableSearchExecutionEventsResponse {
	return &NullableSearchExecutionEventsResponse{value: val, isSet: true}
}

func (v NullableSearchExecutionEventsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchExecutionEventsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


