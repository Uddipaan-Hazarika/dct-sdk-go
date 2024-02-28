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

// checks if the UpdateBookmarkResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateBookmarkResponse{}

// UpdateBookmarkResponse struct for UpdateBookmarkResponse
type UpdateBookmarkResponse struct {
	Job *Job `json:"job,omitempty"`
}

// NewUpdateBookmarkResponse instantiates a new UpdateBookmarkResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateBookmarkResponse() *UpdateBookmarkResponse {
	this := UpdateBookmarkResponse{}
	return &this
}

// NewUpdateBookmarkResponseWithDefaults instantiates a new UpdateBookmarkResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateBookmarkResponseWithDefaults() *UpdateBookmarkResponse {
	this := UpdateBookmarkResponse{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *UpdateBookmarkResponse) GetJob() Job {
	if o == nil || IsNil(o.Job) {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBookmarkResponse) GetJobOk() (*Job, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *UpdateBookmarkResponse) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *UpdateBookmarkResponse) SetJob(v Job) {
	o.Job = &v
}

func (o UpdateBookmarkResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateBookmarkResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableUpdateBookmarkResponse struct {
	value *UpdateBookmarkResponse
	isSet bool
}

func (v NullableUpdateBookmarkResponse) Get() *UpdateBookmarkResponse {
	return v.value
}

func (v *NullableUpdateBookmarkResponse) Set(val *UpdateBookmarkResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateBookmarkResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateBookmarkResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateBookmarkResponse(val *UpdateBookmarkResponse) *NullableUpdateBookmarkResponse {
	return &NullableUpdateBookmarkResponse{value: val, isSet: true}
}

func (v NullableUpdateBookmarkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateBookmarkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


