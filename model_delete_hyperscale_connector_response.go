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

// checks if the DeleteHyperscaleConnectorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteHyperscaleConnectorResponse{}

// DeleteHyperscaleConnectorResponse struct for DeleteHyperscaleConnectorResponse
type DeleteHyperscaleConnectorResponse struct {
	Job *Job `json:"job,omitempty"`
}

// NewDeleteHyperscaleConnectorResponse instantiates a new DeleteHyperscaleConnectorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteHyperscaleConnectorResponse() *DeleteHyperscaleConnectorResponse {
	this := DeleteHyperscaleConnectorResponse{}
	return &this
}

// NewDeleteHyperscaleConnectorResponseWithDefaults instantiates a new DeleteHyperscaleConnectorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteHyperscaleConnectorResponseWithDefaults() *DeleteHyperscaleConnectorResponse {
	this := DeleteHyperscaleConnectorResponse{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *DeleteHyperscaleConnectorResponse) GetJob() Job {
	if o == nil || IsNil(o.Job) {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteHyperscaleConnectorResponse) GetJobOk() (*Job, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *DeleteHyperscaleConnectorResponse) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *DeleteHyperscaleConnectorResponse) SetJob(v Job) {
	o.Job = &v
}

func (o DeleteHyperscaleConnectorResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteHyperscaleConnectorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableDeleteHyperscaleConnectorResponse struct {
	value *DeleteHyperscaleConnectorResponse
	isSet bool
}

func (v NullableDeleteHyperscaleConnectorResponse) Get() *DeleteHyperscaleConnectorResponse {
	return v.value
}

func (v *NullableDeleteHyperscaleConnectorResponse) Set(val *DeleteHyperscaleConnectorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteHyperscaleConnectorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteHyperscaleConnectorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteHyperscaleConnectorResponse(val *DeleteHyperscaleConnectorResponse) *NullableDeleteHyperscaleConnectorResponse {
	return &NullableDeleteHyperscaleConnectorResponse{value: val, isSet: true}
}

func (v NullableDeleteHyperscaleConnectorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteHyperscaleConnectorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


