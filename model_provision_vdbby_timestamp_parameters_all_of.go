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

// ProvisionVDBByTimestampParametersAllOf struct for ProvisionVDBByTimestampParametersAllOf
type ProvisionVDBByTimestampParametersAllOf struct {
	// The ID of the source object (dSource or VDB) to provision from. All other objects referenced by the parameters must live on the same engine as the source.
	SourceDataId string `json:"source_data_id"`
}

// NewProvisionVDBByTimestampParametersAllOf instantiates a new ProvisionVDBByTimestampParametersAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisionVDBByTimestampParametersAllOf(sourceDataId string) *ProvisionVDBByTimestampParametersAllOf {
	this := ProvisionVDBByTimestampParametersAllOf{}
	this.SourceDataId = sourceDataId
	return &this
}

// NewProvisionVDBByTimestampParametersAllOfWithDefaults instantiates a new ProvisionVDBByTimestampParametersAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisionVDBByTimestampParametersAllOfWithDefaults() *ProvisionVDBByTimestampParametersAllOf {
	this := ProvisionVDBByTimestampParametersAllOf{}
	return &this
}

// GetSourceDataId returns the SourceDataId field value
func (o *ProvisionVDBByTimestampParametersAllOf) GetSourceDataId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceDataId
}

// GetSourceDataIdOk returns a tuple with the SourceDataId field value
// and a boolean to check if the value has been set.
func (o *ProvisionVDBByTimestampParametersAllOf) GetSourceDataIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SourceDataId, true
}

// SetSourceDataId sets field value
func (o *ProvisionVDBByTimestampParametersAllOf) SetSourceDataId(v string) {
	o.SourceDataId = v
}

func (o ProvisionVDBByTimestampParametersAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["source_data_id"] = o.SourceDataId
	}
	return json.Marshal(toSerialize)
}

type NullableProvisionVDBByTimestampParametersAllOf struct {
	value *ProvisionVDBByTimestampParametersAllOf
	isSet bool
}

func (v NullableProvisionVDBByTimestampParametersAllOf) Get() *ProvisionVDBByTimestampParametersAllOf {
	return v.value
}

func (v *NullableProvisionVDBByTimestampParametersAllOf) Set(val *ProvisionVDBByTimestampParametersAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisionVDBByTimestampParametersAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisionVDBByTimestampParametersAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisionVDBByTimestampParametersAllOf(val *ProvisionVDBByTimestampParametersAllOf) *NullableProvisionVDBByTimestampParametersAllOf {
	return &NullableProvisionVDBByTimestampParametersAllOf{value: val, isSet: true}
}

func (v NullableProvisionVDBByTimestampParametersAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisionVDBByTimestampParametersAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


