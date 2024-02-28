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

// checks if the UpdateAlgorithmRevisionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAlgorithmRevisionResponse{}

// UpdateAlgorithmRevisionResponse struct for UpdateAlgorithmRevisionResponse
type UpdateAlgorithmRevisionResponse struct {
	Job *Job `json:"job,omitempty"`
}

// NewUpdateAlgorithmRevisionResponse instantiates a new UpdateAlgorithmRevisionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAlgorithmRevisionResponse() *UpdateAlgorithmRevisionResponse {
	this := UpdateAlgorithmRevisionResponse{}
	return &this
}

// NewUpdateAlgorithmRevisionResponseWithDefaults instantiates a new UpdateAlgorithmRevisionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAlgorithmRevisionResponseWithDefaults() *UpdateAlgorithmRevisionResponse {
	this := UpdateAlgorithmRevisionResponse{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *UpdateAlgorithmRevisionResponse) GetJob() Job {
	if o == nil || IsNil(o.Job) {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAlgorithmRevisionResponse) GetJobOk() (*Job, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *UpdateAlgorithmRevisionResponse) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *UpdateAlgorithmRevisionResponse) SetJob(v Job) {
	o.Job = &v
}

func (o UpdateAlgorithmRevisionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAlgorithmRevisionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableUpdateAlgorithmRevisionResponse struct {
	value *UpdateAlgorithmRevisionResponse
	isSet bool
}

func (v NullableUpdateAlgorithmRevisionResponse) Get() *UpdateAlgorithmRevisionResponse {
	return v.value
}

func (v *NullableUpdateAlgorithmRevisionResponse) Set(val *UpdateAlgorithmRevisionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAlgorithmRevisionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAlgorithmRevisionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAlgorithmRevisionResponse(val *UpdateAlgorithmRevisionResponse) *NullableUpdateAlgorithmRevisionResponse {
	return &NullableUpdateAlgorithmRevisionResponse{value: val, isSet: true}
}

func (v NullableUpdateAlgorithmRevisionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAlgorithmRevisionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


