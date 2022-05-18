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

// DisableVDBResponse struct for DisableVDBResponse
type DisableVDBResponse struct {
	Job *Job `json:"job,omitempty"`
}

// NewDisableVDBResponse instantiates a new DisableVDBResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisableVDBResponse() *DisableVDBResponse {
	this := DisableVDBResponse{}
	return &this
}

// NewDisableVDBResponseWithDefaults instantiates a new DisableVDBResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDisableVDBResponseWithDefaults() *DisableVDBResponse {
	this := DisableVDBResponse{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *DisableVDBResponse) GetJob() Job {
	if o == nil || o.Job == nil {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisableVDBResponse) GetJobOk() (*Job, bool) {
	if o == nil || o.Job == nil {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *DisableVDBResponse) HasJob() bool {
	if o != nil && o.Job != nil {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *DisableVDBResponse) SetJob(v Job) {
	o.Job = &v
}

func (o DisableVDBResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Job != nil {
		toSerialize["job"] = o.Job
	}
	return json.Marshal(toSerialize)
}

type NullableDisableVDBResponse struct {
	value *DisableVDBResponse
	isSet bool
}

func (v NullableDisableVDBResponse) Get() *DisableVDBResponse {
	return v.value
}

func (v *NullableDisableVDBResponse) Set(val *DisableVDBResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDisableVDBResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDisableVDBResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisableVDBResponse(val *DisableVDBResponse) *NullableDisableVDBResponse {
	return &NullableDisableVDBResponse{value: val, isSet: true}
}

func (v NullableDisableVDBResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisableVDBResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


