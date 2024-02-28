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

// checks if the ListToolkitResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListToolkitResponse{}

// ListToolkitResponse struct for ListToolkitResponse
type ListToolkitResponse struct {
	Items []Toolkit `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewListToolkitResponse instantiates a new ListToolkitResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListToolkitResponse() *ListToolkitResponse {
	this := ListToolkitResponse{}
	return &this
}

// NewListToolkitResponseWithDefaults instantiates a new ListToolkitResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListToolkitResponseWithDefaults() *ListToolkitResponse {
	this := ListToolkitResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListToolkitResponse) GetItems() []Toolkit {
	if o == nil || IsNil(o.Items) {
		var ret []Toolkit
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListToolkitResponse) GetItemsOk() ([]Toolkit, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListToolkitResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Toolkit and assigns it to the Items field.
func (o *ListToolkitResponse) SetItems(v []Toolkit) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *ListToolkitResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListToolkitResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *ListToolkitResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *ListToolkitResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o ListToolkitResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListToolkitResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return toSerialize, nil
}

type NullableListToolkitResponse struct {
	value *ListToolkitResponse
	isSet bool
}

func (v NullableListToolkitResponse) Get() *ListToolkitResponse {
	return v.value
}

func (v *NullableListToolkitResponse) Set(val *ListToolkitResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListToolkitResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListToolkitResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListToolkitResponse(val *ListToolkitResponse) *NullableListToolkitResponse {
	return &NullableListToolkitResponse{value: val, isSet: true}
}

func (v NullableListToolkitResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListToolkitResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


