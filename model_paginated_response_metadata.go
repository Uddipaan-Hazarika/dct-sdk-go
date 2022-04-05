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

// PaginatedResponseMetadata struct for PaginatedResponseMetadata
type PaginatedResponseMetadata struct {
	// Pointer to the previous page of results. Use this value as a cursor query parameter in a subsequent request, along with limit, to navigate through the collection by virtual page.
	PrevCursor *string `json:"prev_cursor,omitempty"`
	// Pointer to the next page of results. Use this value as a cursor query parameter in a subsequent request, along with limit, to navigate through the collection by virtual page.
	NextCursor *string `json:"next_cursor,omitempty"`
	// The total number of results. This value may not be provided.
	Total *int32 `json:"total,omitempty"`
}

// NewPaginatedResponseMetadata instantiates a new PaginatedResponseMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedResponseMetadata() *PaginatedResponseMetadata {
	this := PaginatedResponseMetadata{}
	return &this
}

// NewPaginatedResponseMetadataWithDefaults instantiates a new PaginatedResponseMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedResponseMetadataWithDefaults() *PaginatedResponseMetadata {
	this := PaginatedResponseMetadata{}
	return &this
}

// GetPrevCursor returns the PrevCursor field value if set, zero value otherwise.
func (o *PaginatedResponseMetadata) GetPrevCursor() string {
	if o == nil || o.PrevCursor == nil {
		var ret string
		return ret
	}
	return *o.PrevCursor
}

// GetPrevCursorOk returns a tuple with the PrevCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedResponseMetadata) GetPrevCursorOk() (*string, bool) {
	if o == nil || o.PrevCursor == nil {
		return nil, false
	}
	return o.PrevCursor, true
}

// HasPrevCursor returns a boolean if a field has been set.
func (o *PaginatedResponseMetadata) HasPrevCursor() bool {
	if o != nil && o.PrevCursor != nil {
		return true
	}

	return false
}

// SetPrevCursor gets a reference to the given string and assigns it to the PrevCursor field.
func (o *PaginatedResponseMetadata) SetPrevCursor(v string) {
	o.PrevCursor = &v
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise.
func (o *PaginatedResponseMetadata) GetNextCursor() string {
	if o == nil || o.NextCursor == nil {
		var ret string
		return ret
	}
	return *o.NextCursor
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedResponseMetadata) GetNextCursorOk() (*string, bool) {
	if o == nil || o.NextCursor == nil {
		return nil, false
	}
	return o.NextCursor, true
}

// HasNextCursor returns a boolean if a field has been set.
func (o *PaginatedResponseMetadata) HasNextCursor() bool {
	if o != nil && o.NextCursor != nil {
		return true
	}

	return false
}

// SetNextCursor gets a reference to the given string and assigns it to the NextCursor field.
func (o *PaginatedResponseMetadata) SetNextCursor(v string) {
	o.NextCursor = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *PaginatedResponseMetadata) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedResponseMetadata) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *PaginatedResponseMetadata) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *PaginatedResponseMetadata) SetTotal(v int32) {
	o.Total = &v
}

func (o PaginatedResponseMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PrevCursor != nil {
		toSerialize["prev_cursor"] = o.PrevCursor
	}
	if o.NextCursor != nil {
		toSerialize["next_cursor"] = o.NextCursor
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedResponseMetadata struct {
	value *PaginatedResponseMetadata
	isSet bool
}

func (v NullablePaginatedResponseMetadata) Get() *PaginatedResponseMetadata {
	return v.value
}

func (v *NullablePaginatedResponseMetadata) Set(val *PaginatedResponseMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedResponseMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedResponseMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedResponseMetadata(val *PaginatedResponseMetadata) *NullablePaginatedResponseMetadata {
	return &NullablePaginatedResponseMetadata{value: val, isSet: true}
}

func (v NullablePaginatedResponseMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedResponseMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


