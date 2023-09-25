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

// checks if the LinkDSourceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkDSourceResponse{}

// LinkDSourceResponse struct for LinkDSourceResponse
type LinkDSourceResponse struct {
	Job *Job `json:"job,omitempty"`
	// The ID of the dSource.
	DsourceId *string `json:"dsource_id,omitempty"`
}

// NewLinkDSourceResponse instantiates a new LinkDSourceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkDSourceResponse() *LinkDSourceResponse {
	this := LinkDSourceResponse{}
	return &this
}

// NewLinkDSourceResponseWithDefaults instantiates a new LinkDSourceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkDSourceResponseWithDefaults() *LinkDSourceResponse {
	this := LinkDSourceResponse{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *LinkDSourceResponse) GetJob() Job {
	if o == nil || IsNil(o.Job) {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkDSourceResponse) GetJobOk() (*Job, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *LinkDSourceResponse) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *LinkDSourceResponse) SetJob(v Job) {
	o.Job = &v
}

// GetDsourceId returns the DsourceId field value if set, zero value otherwise.
func (o *LinkDSourceResponse) GetDsourceId() string {
	if o == nil || IsNil(o.DsourceId) {
		var ret string
		return ret
	}
	return *o.DsourceId
}

// GetDsourceIdOk returns a tuple with the DsourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkDSourceResponse) GetDsourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DsourceId) {
		return nil, false
	}
	return o.DsourceId, true
}

// HasDsourceId returns a boolean if a field has been set.
func (o *LinkDSourceResponse) HasDsourceId() bool {
	if o != nil && !IsNil(o.DsourceId) {
		return true
	}

	return false
}

// SetDsourceId gets a reference to the given string and assigns it to the DsourceId field.
func (o *LinkDSourceResponse) SetDsourceId(v string) {
	o.DsourceId = &v
}

func (o LinkDSourceResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinkDSourceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	if !IsNil(o.DsourceId) {
		toSerialize["dsource_id"] = o.DsourceId
	}
	return toSerialize, nil
}

type NullableLinkDSourceResponse struct {
	value *LinkDSourceResponse
	isSet bool
}

func (v NullableLinkDSourceResponse) Get() *LinkDSourceResponse {
	return v.value
}

func (v *NullableLinkDSourceResponse) Set(val *LinkDSourceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkDSourceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkDSourceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkDSourceResponse(val *LinkDSourceResponse) *NullableLinkDSourceResponse {
	return &NullableLinkDSourceResponse{value: val, isSet: true}
}

func (v NullableLinkDSourceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkDSourceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

