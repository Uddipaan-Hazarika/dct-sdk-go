/*
Delphix DCT API

Delphix DCT API

API version: 1.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// ListSourcesResponse struct for ListSourcesResponse
type ListSourcesResponse struct {
	Items []Source `json:"items,omitempty"`
	// Sadly, sometimes requests to the API are not successful. Failures can occur for a wide range of reasons. The Errors object contains information about full or partial failures which might have occurred during the request.
	Errors []Error `json:"errors,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewListSourcesResponse instantiates a new ListSourcesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSourcesResponse() *ListSourcesResponse {
	this := ListSourcesResponse{}
	return &this
}

// NewListSourcesResponseWithDefaults instantiates a new ListSourcesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSourcesResponseWithDefaults() *ListSourcesResponse {
	this := ListSourcesResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListSourcesResponse) GetItems() []Source {
	if o == nil || o.Items == nil {
		var ret []Source
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSourcesResponse) GetItemsOk() ([]Source, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListSourcesResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Source and assigns it to the Items field.
func (o *ListSourcesResponse) SetItems(v []Source) {
	o.Items = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ListSourcesResponse) GetErrors() []Error {
	if o == nil || o.Errors == nil {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSourcesResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ListSourcesResponse) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *ListSourcesResponse) SetErrors(v []Error) {
	o.Errors = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *ListSourcesResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || o.ResponseMetadata == nil {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSourcesResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || o.ResponseMetadata == nil {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *ListSourcesResponse) HasResponseMetadata() bool {
	if o != nil && o.ResponseMetadata != nil {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *ListSourcesResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o ListSourcesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	if o.ResponseMetadata != nil {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return json.Marshal(toSerialize)
}

type NullableListSourcesResponse struct {
	value *ListSourcesResponse
	isSet bool
}

func (v NullableListSourcesResponse) Get() *ListSourcesResponse {
	return v.value
}

func (v *NullableListSourcesResponse) Set(val *ListSourcesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListSourcesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListSourcesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSourcesResponse(val *ListSourcesResponse) *NullableListSourcesResponse {
	return &NullableListSourcesResponse{value: val, isSet: true}
}

func (v NullableListSourcesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSourcesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


