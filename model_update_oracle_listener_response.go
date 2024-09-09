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

// checks if the UpdateOracleListenerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOracleListenerResponse{}

// UpdateOracleListenerResponse struct for UpdateOracleListenerResponse
type UpdateOracleListenerResponse struct {
	Job *Job `json:"job,omitempty"`
}

// NewUpdateOracleListenerResponse instantiates a new UpdateOracleListenerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOracleListenerResponse() *UpdateOracleListenerResponse {
	this := UpdateOracleListenerResponse{}
	return &this
}

// NewUpdateOracleListenerResponseWithDefaults instantiates a new UpdateOracleListenerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOracleListenerResponseWithDefaults() *UpdateOracleListenerResponse {
	this := UpdateOracleListenerResponse{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *UpdateOracleListenerResponse) GetJob() Job {
	if o == nil || IsNil(o.Job) {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOracleListenerResponse) GetJobOk() (*Job, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *UpdateOracleListenerResponse) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *UpdateOracleListenerResponse) SetJob(v Job) {
	o.Job = &v
}

func (o UpdateOracleListenerResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOracleListenerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableUpdateOracleListenerResponse struct {
	value *UpdateOracleListenerResponse
	isSet bool
}

func (v NullableUpdateOracleListenerResponse) Get() *UpdateOracleListenerResponse {
	return v.value
}

func (v *NullableUpdateOracleListenerResponse) Set(val *UpdateOracleListenerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOracleListenerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOracleListenerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOracleListenerResponse(val *UpdateOracleListenerResponse) *NullableUpdateOracleListenerResponse {
	return &NullableUpdateOracleListenerResponse{value: val, isSet: true}
}

func (v NullableUpdateOracleListenerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOracleListenerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

