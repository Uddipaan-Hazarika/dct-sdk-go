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

// checks if the DisableEnvironmentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DisableEnvironmentResponse{}

// DisableEnvironmentResponse struct for DisableEnvironmentResponse
type DisableEnvironmentResponse struct {
	Job *Job `json:"job,omitempty"`
}

// NewDisableEnvironmentResponse instantiates a new DisableEnvironmentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisableEnvironmentResponse() *DisableEnvironmentResponse {
	this := DisableEnvironmentResponse{}
	return &this
}

// NewDisableEnvironmentResponseWithDefaults instantiates a new DisableEnvironmentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDisableEnvironmentResponseWithDefaults() *DisableEnvironmentResponse {
	this := DisableEnvironmentResponse{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *DisableEnvironmentResponse) GetJob() Job {
	if o == nil || IsNil(o.Job) {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisableEnvironmentResponse) GetJobOk() (*Job, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *DisableEnvironmentResponse) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *DisableEnvironmentResponse) SetJob(v Job) {
	o.Job = &v
}

func (o DisableEnvironmentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DisableEnvironmentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableDisableEnvironmentResponse struct {
	value *DisableEnvironmentResponse
	isSet bool
}

func (v NullableDisableEnvironmentResponse) Get() *DisableEnvironmentResponse {
	return v.value
}

func (v *NullableDisableEnvironmentResponse) Set(val *DisableEnvironmentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDisableEnvironmentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDisableEnvironmentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisableEnvironmentResponse(val *DisableEnvironmentResponse) *NullableDisableEnvironmentResponse {
	return &NullableDisableEnvironmentResponse{value: val, isSet: true}
}

func (v NullableDisableEnvironmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisableEnvironmentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


