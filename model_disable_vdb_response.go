/*
Delphix API Gateway

Delphix API Gateway API

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
	// The initiated job id.
	JobId *string `json:"job_id,omitempty"`
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

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *DisableVDBResponse) GetJobId() string {
	if o == nil || o.JobId == nil {
		var ret string
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisableVDBResponse) GetJobIdOk() (*string, bool) {
	if o == nil || o.JobId == nil {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *DisableVDBResponse) HasJobId() bool {
	if o != nil && o.JobId != nil {
		return true
	}

	return false
}

// SetJobId gets a reference to the given string and assigns it to the JobId field.
func (o *DisableVDBResponse) SetJobId(v string) {
	o.JobId = &v
}

func (o DisableVDBResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.JobId != nil {
		toSerialize["job_id"] = o.JobId
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


