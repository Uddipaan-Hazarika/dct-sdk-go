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

// ListBookmarksResponse struct for ListBookmarksResponse
type ListBookmarksResponse struct {
	Items []Bookmark `json:"items,omitempty"`
	// Sadly, sometimes requests to the API are not successful. Failures can occur for a wide range of reasons. The Errors object contains information about full or partial failures which might have occurred during the request.
	Errors []Error `json:"errors,omitempty"`
}

// NewListBookmarksResponse instantiates a new ListBookmarksResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListBookmarksResponse() *ListBookmarksResponse {
	this := ListBookmarksResponse{}
	return &this
}

// NewListBookmarksResponseWithDefaults instantiates a new ListBookmarksResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListBookmarksResponseWithDefaults() *ListBookmarksResponse {
	this := ListBookmarksResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListBookmarksResponse) GetItems() []Bookmark {
	if o == nil || o.Items == nil {
		var ret []Bookmark
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListBookmarksResponse) GetItemsOk() ([]Bookmark, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListBookmarksResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Bookmark and assigns it to the Items field.
func (o *ListBookmarksResponse) SetItems(v []Bookmark) {
	o.Items = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ListBookmarksResponse) GetErrors() []Error {
	if o == nil || o.Errors == nil {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListBookmarksResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ListBookmarksResponse) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *ListBookmarksResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o ListBookmarksResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableListBookmarksResponse struct {
	value *ListBookmarksResponse
	isSet bool
}

func (v NullableListBookmarksResponse) Get() *ListBookmarksResponse {
	return v.value
}

func (v *NullableListBookmarksResponse) Set(val *ListBookmarksResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListBookmarksResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListBookmarksResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListBookmarksResponse(val *ListBookmarksResponse) *NullableListBookmarksResponse {
	return &NullableListBookmarksResponse{value: val, isSet: true}
}

func (v NullableListBookmarksResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListBookmarksResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


