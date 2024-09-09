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

// checks if the ListConnectorsMaskingJobsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListConnectorsMaskingJobsResponse{}

// ListConnectorsMaskingJobsResponse struct for ListConnectorsMaskingJobsResponse
type ListConnectorsMaskingJobsResponse struct {
	Items []MaskingJobWithConnectorRole `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewListConnectorsMaskingJobsResponse instantiates a new ListConnectorsMaskingJobsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConnectorsMaskingJobsResponse() *ListConnectorsMaskingJobsResponse {
	this := ListConnectorsMaskingJobsResponse{}
	return &this
}

// NewListConnectorsMaskingJobsResponseWithDefaults instantiates a new ListConnectorsMaskingJobsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConnectorsMaskingJobsResponseWithDefaults() *ListConnectorsMaskingJobsResponse {
	this := ListConnectorsMaskingJobsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListConnectorsMaskingJobsResponse) GetItems() []MaskingJobWithConnectorRole {
	if o == nil || IsNil(o.Items) {
		var ret []MaskingJobWithConnectorRole
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectorsMaskingJobsResponse) GetItemsOk() ([]MaskingJobWithConnectorRole, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListConnectorsMaskingJobsResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []MaskingJobWithConnectorRole and assigns it to the Items field.
func (o *ListConnectorsMaskingJobsResponse) SetItems(v []MaskingJobWithConnectorRole) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *ListConnectorsMaskingJobsResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectorsMaskingJobsResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *ListConnectorsMaskingJobsResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *ListConnectorsMaskingJobsResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o ListConnectorsMaskingJobsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListConnectorsMaskingJobsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return toSerialize, nil
}

type NullableListConnectorsMaskingJobsResponse struct {
	value *ListConnectorsMaskingJobsResponse
	isSet bool
}

func (v NullableListConnectorsMaskingJobsResponse) Get() *ListConnectorsMaskingJobsResponse {
	return v.value
}

func (v *NullableListConnectorsMaskingJobsResponse) Set(val *ListConnectorsMaskingJobsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListConnectorsMaskingJobsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListConnectorsMaskingJobsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConnectorsMaskingJobsResponse(val *ListConnectorsMaskingJobsResponse) *NullableListConnectorsMaskingJobsResponse {
	return &NullableListConnectorsMaskingJobsResponse{value: val, isSet: true}
}

func (v NullableListConnectorsMaskingJobsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConnectorsMaskingJobsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

