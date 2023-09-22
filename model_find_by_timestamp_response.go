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

// checks if the FindByTimestampResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FindByTimestampResponse{}

// FindByTimestampResponse struct for FindByTimestampResponse
type FindByTimestampResponse struct {
	Items []Snapshot `json:"items,omitempty"`
}

// NewFindByTimestampResponse instantiates a new FindByTimestampResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindByTimestampResponse() *FindByTimestampResponse {
	this := FindByTimestampResponse{}
	return &this
}

// NewFindByTimestampResponseWithDefaults instantiates a new FindByTimestampResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindByTimestampResponseWithDefaults() *FindByTimestampResponse {
	this := FindByTimestampResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *FindByTimestampResponse) GetItems() []Snapshot {
	if o == nil || IsNil(o.Items) {
		var ret []Snapshot
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindByTimestampResponse) GetItemsOk() ([]Snapshot, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *FindByTimestampResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Snapshot and assigns it to the Items field.
func (o *FindByTimestampResponse) SetItems(v []Snapshot) {
	o.Items = v
}

func (o FindByTimestampResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FindByTimestampResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableFindByTimestampResponse struct {
	value *FindByTimestampResponse
	isSet bool
}

func (v NullableFindByTimestampResponse) Get() *FindByTimestampResponse {
	return v.value
}

func (v *NullableFindByTimestampResponse) Set(val *FindByTimestampResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFindByTimestampResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFindByTimestampResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindByTimestampResponse(val *FindByTimestampResponse) *NullableFindByTimestampResponse {
	return &NullableFindByTimestampResponse{value: val, isSet: true}
}

func (v NullableFindByTimestampResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindByTimestampResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


