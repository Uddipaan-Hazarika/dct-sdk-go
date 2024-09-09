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

// checks if the BookmarkVDBs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BookmarkVDBs{}

// BookmarkVDBs VDB id and name associated with bookmark.
type BookmarkVDBs struct {
	// The VDB id.
	VdbId *string `json:"vdb_id,omitempty"`
	// The VDB name.
	VdbName *string `json:"vdb_name,omitempty"`
}

// NewBookmarkVDBs instantiates a new BookmarkVDBs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBookmarkVDBs() *BookmarkVDBs {
	this := BookmarkVDBs{}
	return &this
}

// NewBookmarkVDBsWithDefaults instantiates a new BookmarkVDBs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBookmarkVDBsWithDefaults() *BookmarkVDBs {
	this := BookmarkVDBs{}
	return &this
}

// GetVdbId returns the VdbId field value if set, zero value otherwise.
func (o *BookmarkVDBs) GetVdbId() string {
	if o == nil || IsNil(o.VdbId) {
		var ret string
		return ret
	}
	return *o.VdbId
}

// GetVdbIdOk returns a tuple with the VdbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookmarkVDBs) GetVdbIdOk() (*string, bool) {
	if o == nil || IsNil(o.VdbId) {
		return nil, false
	}
	return o.VdbId, true
}

// HasVdbId returns a boolean if a field has been set.
func (o *BookmarkVDBs) HasVdbId() bool {
	if o != nil && !IsNil(o.VdbId) {
		return true
	}

	return false
}

// SetVdbId gets a reference to the given string and assigns it to the VdbId field.
func (o *BookmarkVDBs) SetVdbId(v string) {
	o.VdbId = &v
}

// GetVdbName returns the VdbName field value if set, zero value otherwise.
func (o *BookmarkVDBs) GetVdbName() string {
	if o == nil || IsNil(o.VdbName) {
		var ret string
		return ret
	}
	return *o.VdbName
}

// GetVdbNameOk returns a tuple with the VdbName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookmarkVDBs) GetVdbNameOk() (*string, bool) {
	if o == nil || IsNil(o.VdbName) {
		return nil, false
	}
	return o.VdbName, true
}

// HasVdbName returns a boolean if a field has been set.
func (o *BookmarkVDBs) HasVdbName() bool {
	if o != nil && !IsNil(o.VdbName) {
		return true
	}

	return false
}

// SetVdbName gets a reference to the given string and assigns it to the VdbName field.
func (o *BookmarkVDBs) SetVdbName(v string) {
	o.VdbName = &v
}

func (o BookmarkVDBs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BookmarkVDBs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VdbId) {
		toSerialize["vdb_id"] = o.VdbId
	}
	if !IsNil(o.VdbName) {
		toSerialize["vdb_name"] = o.VdbName
	}
	return toSerialize, nil
}

type NullableBookmarkVDBs struct {
	value *BookmarkVDBs
	isSet bool
}

func (v NullableBookmarkVDBs) Get() *BookmarkVDBs {
	return v.value
}

func (v *NullableBookmarkVDBs) Set(val *BookmarkVDBs) {
	v.value = val
	v.isSet = true
}

func (v NullableBookmarkVDBs) IsSet() bool {
	return v.isSet
}

func (v *NullableBookmarkVDBs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBookmarkVDBs(val *BookmarkVDBs) *NullableBookmarkVDBs {
	return &NullableBookmarkVDBs{value: val, isSet: true}
}

func (v NullableBookmarkVDBs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBookmarkVDBs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

