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

// checks if the DeleteEnvironmentUserResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteEnvironmentUserResponse{}

// DeleteEnvironmentUserResponse struct for DeleteEnvironmentUserResponse
type DeleteEnvironmentUserResponse struct {
	Job *Job `json:"job,omitempty"`
}

// NewDeleteEnvironmentUserResponse instantiates a new DeleteEnvironmentUserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteEnvironmentUserResponse() *DeleteEnvironmentUserResponse {
	this := DeleteEnvironmentUserResponse{}
	return &this
}

// NewDeleteEnvironmentUserResponseWithDefaults instantiates a new DeleteEnvironmentUserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteEnvironmentUserResponseWithDefaults() *DeleteEnvironmentUserResponse {
	this := DeleteEnvironmentUserResponse{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *DeleteEnvironmentUserResponse) GetJob() Job {
	if o == nil || IsNil(o.Job) {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteEnvironmentUserResponse) GetJobOk() (*Job, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *DeleteEnvironmentUserResponse) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *DeleteEnvironmentUserResponse) SetJob(v Job) {
	o.Job = &v
}

func (o DeleteEnvironmentUserResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteEnvironmentUserResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableDeleteEnvironmentUserResponse struct {
	value *DeleteEnvironmentUserResponse
	isSet bool
}

func (v NullableDeleteEnvironmentUserResponse) Get() *DeleteEnvironmentUserResponse {
	return v.value
}

func (v *NullableDeleteEnvironmentUserResponse) Set(val *DeleteEnvironmentUserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteEnvironmentUserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteEnvironmentUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteEnvironmentUserResponse(val *DeleteEnvironmentUserResponse) *NullableDeleteEnvironmentUserResponse {
	return &NullableDeleteEnvironmentUserResponse{value: val, isSet: true}
}

func (v NullableDeleteEnvironmentUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteEnvironmentUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


