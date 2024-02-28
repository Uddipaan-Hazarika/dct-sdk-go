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

// checks if the LinkStagingPushDSourceDefaultRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkStagingPushDSourceDefaultRequest{}

// LinkStagingPushDSourceDefaultRequest struct for LinkStagingPushDSourceDefaultRequest
type LinkStagingPushDSourceDefaultRequest struct {
	// The ID of the environment to be linked.
	EnvironmentId string `json:"environment_id"`
}

// NewLinkStagingPushDSourceDefaultRequest instantiates a new LinkStagingPushDSourceDefaultRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkStagingPushDSourceDefaultRequest(environmentId string) *LinkStagingPushDSourceDefaultRequest {
	this := LinkStagingPushDSourceDefaultRequest{}
	this.EnvironmentId = environmentId
	return &this
}

// NewLinkStagingPushDSourceDefaultRequestWithDefaults instantiates a new LinkStagingPushDSourceDefaultRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkStagingPushDSourceDefaultRequestWithDefaults() *LinkStagingPushDSourceDefaultRequest {
	this := LinkStagingPushDSourceDefaultRequest{}
	return &this
}

// GetEnvironmentId returns the EnvironmentId field value
func (o *LinkStagingPushDSourceDefaultRequest) GetEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *LinkStagingPushDSourceDefaultRequest) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value
func (o *LinkStagingPushDSourceDefaultRequest) SetEnvironmentId(v string) {
	o.EnvironmentId = v
}

func (o LinkStagingPushDSourceDefaultRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinkStagingPushDSourceDefaultRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["environment_id"] = o.EnvironmentId
	return toSerialize, nil
}

type NullableLinkStagingPushDSourceDefaultRequest struct {
	value *LinkStagingPushDSourceDefaultRequest
	isSet bool
}

func (v NullableLinkStagingPushDSourceDefaultRequest) Get() *LinkStagingPushDSourceDefaultRequest {
	return v.value
}

func (v *NullableLinkStagingPushDSourceDefaultRequest) Set(val *LinkStagingPushDSourceDefaultRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkStagingPushDSourceDefaultRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkStagingPushDSourceDefaultRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkStagingPushDSourceDefaultRequest(val *LinkStagingPushDSourceDefaultRequest) *NullableLinkStagingPushDSourceDefaultRequest {
	return &NullableLinkStagingPushDSourceDefaultRequest{value: val, isSet: true}
}

func (v NullableLinkStagingPushDSourceDefaultRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkStagingPushDSourceDefaultRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


