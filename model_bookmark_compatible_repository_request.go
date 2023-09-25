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

// checks if the BookmarkCompatibleRepositoryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BookmarkCompatibleRepositoryRequest{}

// BookmarkCompatibleRepositoryRequest struct for BookmarkCompatibleRepositoryRequest
type BookmarkCompatibleRepositoryRequest struct {
	// The ID of the bookmark from which to execute the operation. The bookmark must contain only one VDB.
	BookmarkId string `json:"bookmark_id"`
	// The ID or name of the target environment.
	EnvironmentId *string `json:"environment_id,omitempty"`
}

// NewBookmarkCompatibleRepositoryRequest instantiates a new BookmarkCompatibleRepositoryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBookmarkCompatibleRepositoryRequest(bookmarkId string) *BookmarkCompatibleRepositoryRequest {
	this := BookmarkCompatibleRepositoryRequest{}
	this.BookmarkId = bookmarkId
	return &this
}

// NewBookmarkCompatibleRepositoryRequestWithDefaults instantiates a new BookmarkCompatibleRepositoryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBookmarkCompatibleRepositoryRequestWithDefaults() *BookmarkCompatibleRepositoryRequest {
	this := BookmarkCompatibleRepositoryRequest{}
	return &this
}

// GetBookmarkId returns the BookmarkId field value
func (o *BookmarkCompatibleRepositoryRequest) GetBookmarkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BookmarkId
}

// GetBookmarkIdOk returns a tuple with the BookmarkId field value
// and a boolean to check if the value has been set.
func (o *BookmarkCompatibleRepositoryRequest) GetBookmarkIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BookmarkId, true
}

// SetBookmarkId sets field value
func (o *BookmarkCompatibleRepositoryRequest) SetBookmarkId(v string) {
	o.BookmarkId = v
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *BookmarkCompatibleRepositoryRequest) GetEnvironmentId() string {
	if o == nil || IsNil(o.EnvironmentId) {
		var ret string
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookmarkCompatibleRepositoryRequest) GetEnvironmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnvironmentId) {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *BookmarkCompatibleRepositoryRequest) HasEnvironmentId() bool {
	if o != nil && !IsNil(o.EnvironmentId) {
		return true
	}

	return false
}

// SetEnvironmentId gets a reference to the given string and assigns it to the EnvironmentId field.
func (o *BookmarkCompatibleRepositoryRequest) SetEnvironmentId(v string) {
	o.EnvironmentId = &v
}

func (o BookmarkCompatibleRepositoryRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BookmarkCompatibleRepositoryRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bookmark_id"] = o.BookmarkId
	if !IsNil(o.EnvironmentId) {
		toSerialize["environment_id"] = o.EnvironmentId
	}
	return toSerialize, nil
}

type NullableBookmarkCompatibleRepositoryRequest struct {
	value *BookmarkCompatibleRepositoryRequest
	isSet bool
}

func (v NullableBookmarkCompatibleRepositoryRequest) Get() *BookmarkCompatibleRepositoryRequest {
	return v.value
}

func (v *NullableBookmarkCompatibleRepositoryRequest) Set(val *BookmarkCompatibleRepositoryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBookmarkCompatibleRepositoryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBookmarkCompatibleRepositoryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBookmarkCompatibleRepositoryRequest(val *BookmarkCompatibleRepositoryRequest) *NullableBookmarkCompatibleRepositoryRequest {
	return &NullableBookmarkCompatibleRepositoryRequest{value: val, isSet: true}
}

func (v NullableBookmarkCompatibleRepositoryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBookmarkCompatibleRepositoryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

