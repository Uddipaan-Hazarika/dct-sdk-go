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

// EngineConnectivityCheckResponse Response received for connectivity status check.
type EngineConnectivityCheckResponse struct {
	Message string `json:"message"`
}

// NewEngineConnectivityCheckResponse instantiates a new EngineConnectivityCheckResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEngineConnectivityCheckResponse(message string) *EngineConnectivityCheckResponse {
	this := EngineConnectivityCheckResponse{}
	this.Message = message
	return &this
}

// NewEngineConnectivityCheckResponseWithDefaults instantiates a new EngineConnectivityCheckResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEngineConnectivityCheckResponseWithDefaults() *EngineConnectivityCheckResponse {
	this := EngineConnectivityCheckResponse{}
	return &this
}

// GetMessage returns the Message field value
func (o *EngineConnectivityCheckResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *EngineConnectivityCheckResponse) GetMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *EngineConnectivityCheckResponse) SetMessage(v string) {
	o.Message = v
}

func (o EngineConnectivityCheckResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableEngineConnectivityCheckResponse struct {
	value *EngineConnectivityCheckResponse
	isSet bool
}

func (v NullableEngineConnectivityCheckResponse) Get() *EngineConnectivityCheckResponse {
	return v.value
}

func (v *NullableEngineConnectivityCheckResponse) Set(val *EngineConnectivityCheckResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEngineConnectivityCheckResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEngineConnectivityCheckResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEngineConnectivityCheckResponse(val *EngineConnectivityCheckResponse) *NullableEngineConnectivityCheckResponse {
	return &NullableEngineConnectivityCheckResponse{value: val, isSet: true}
}

func (v NullableEngineConnectivityCheckResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEngineConnectivityCheckResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


