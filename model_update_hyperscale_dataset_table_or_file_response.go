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

// checks if the UpdateHyperscaleDatasetTableOrFileResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateHyperscaleDatasetTableOrFileResponse{}

// UpdateHyperscaleDatasetTableOrFileResponse struct for UpdateHyperscaleDatasetTableOrFileResponse
type UpdateHyperscaleDatasetTableOrFileResponse struct {
	Job *Job `json:"job,omitempty"`
}

// NewUpdateHyperscaleDatasetTableOrFileResponse instantiates a new UpdateHyperscaleDatasetTableOrFileResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateHyperscaleDatasetTableOrFileResponse() *UpdateHyperscaleDatasetTableOrFileResponse {
	this := UpdateHyperscaleDatasetTableOrFileResponse{}
	return &this
}

// NewUpdateHyperscaleDatasetTableOrFileResponseWithDefaults instantiates a new UpdateHyperscaleDatasetTableOrFileResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateHyperscaleDatasetTableOrFileResponseWithDefaults() *UpdateHyperscaleDatasetTableOrFileResponse {
	this := UpdateHyperscaleDatasetTableOrFileResponse{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *UpdateHyperscaleDatasetTableOrFileResponse) GetJob() Job {
	if o == nil || IsNil(o.Job) {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateHyperscaleDatasetTableOrFileResponse) GetJobOk() (*Job, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *UpdateHyperscaleDatasetTableOrFileResponse) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *UpdateHyperscaleDatasetTableOrFileResponse) SetJob(v Job) {
	o.Job = &v
}

func (o UpdateHyperscaleDatasetTableOrFileResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateHyperscaleDatasetTableOrFileResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableUpdateHyperscaleDatasetTableOrFileResponse struct {
	value *UpdateHyperscaleDatasetTableOrFileResponse
	isSet bool
}

func (v NullableUpdateHyperscaleDatasetTableOrFileResponse) Get() *UpdateHyperscaleDatasetTableOrFileResponse {
	return v.value
}

func (v *NullableUpdateHyperscaleDatasetTableOrFileResponse) Set(val *UpdateHyperscaleDatasetTableOrFileResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateHyperscaleDatasetTableOrFileResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateHyperscaleDatasetTableOrFileResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateHyperscaleDatasetTableOrFileResponse(val *UpdateHyperscaleDatasetTableOrFileResponse) *NullableUpdateHyperscaleDatasetTableOrFileResponse {
	return &NullableUpdateHyperscaleDatasetTableOrFileResponse{value: val, isSet: true}
}

func (v NullableUpdateHyperscaleDatasetTableOrFileResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateHyperscaleDatasetTableOrFileResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

