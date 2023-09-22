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

// checks if the ExecuteMaskingJobResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExecuteMaskingJobResponse{}

// ExecuteMaskingJobResponse struct for ExecuteMaskingJobResponse
type ExecuteMaskingJobResponse struct {
	Job *Job `json:"job,omitempty"`
}

// NewExecuteMaskingJobResponse instantiates a new ExecuteMaskingJobResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecuteMaskingJobResponse() *ExecuteMaskingJobResponse {
	this := ExecuteMaskingJobResponse{}
	return &this
}

// NewExecuteMaskingJobResponseWithDefaults instantiates a new ExecuteMaskingJobResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecuteMaskingJobResponseWithDefaults() *ExecuteMaskingJobResponse {
	this := ExecuteMaskingJobResponse{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *ExecuteMaskingJobResponse) GetJob() Job {
	if o == nil || IsNil(o.Job) {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecuteMaskingJobResponse) GetJobOk() (*Job, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *ExecuteMaskingJobResponse) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *ExecuteMaskingJobResponse) SetJob(v Job) {
	o.Job = &v
}

func (o ExecuteMaskingJobResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExecuteMaskingJobResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableExecuteMaskingJobResponse struct {
	value *ExecuteMaskingJobResponse
	isSet bool
}

func (v NullableExecuteMaskingJobResponse) Get() *ExecuteMaskingJobResponse {
	return v.value
}

func (v *NullableExecuteMaskingJobResponse) Set(val *ExecuteMaskingJobResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableExecuteMaskingJobResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableExecuteMaskingJobResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecuteMaskingJobResponse(val *ExecuteMaskingJobResponse) *NullableExecuteMaskingJobResponse {
	return &NullableExecuteMaskingJobResponse{value: val, isSet: true}
}

func (v NullableExecuteMaskingJobResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecuteMaskingJobResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


