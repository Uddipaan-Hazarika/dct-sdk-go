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

// checks if the StartVCDBParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StartVCDBParameters{}

// StartVCDBParameters Parameters to start a vCDB.
type StartVCDBParameters struct {
	// List of specific Virtual Container Database Instances to start.
	Instances []int32 `json:"instances,omitempty"`
}

// NewStartVCDBParameters instantiates a new StartVCDBParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartVCDBParameters() *StartVCDBParameters {
	this := StartVCDBParameters{}
	return &this
}

// NewStartVCDBParametersWithDefaults instantiates a new StartVCDBParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartVCDBParametersWithDefaults() *StartVCDBParameters {
	this := StartVCDBParameters{}
	return &this
}

// GetInstances returns the Instances field value if set, zero value otherwise.
func (o *StartVCDBParameters) GetInstances() []int32 {
	if o == nil || IsNil(o.Instances) {
		var ret []int32
		return ret
	}
	return o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartVCDBParameters) GetInstancesOk() ([]int32, bool) {
	if o == nil || IsNil(o.Instances) {
		return nil, false
	}
	return o.Instances, true
}

// HasInstances returns a boolean if a field has been set.
func (o *StartVCDBParameters) HasInstances() bool {
	if o != nil && !IsNil(o.Instances) {
		return true
	}

	return false
}

// SetInstances gets a reference to the given []int32 and assigns it to the Instances field.
func (o *StartVCDBParameters) SetInstances(v []int32) {
	o.Instances = v
}

func (o StartVCDBParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StartVCDBParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Instances) {
		toSerialize["instances"] = o.Instances
	}
	return toSerialize, nil
}

type NullableStartVCDBParameters struct {
	value *StartVCDBParameters
	isSet bool
}

func (v NullableStartVCDBParameters) Get() *StartVCDBParameters {
	return v.value
}

func (v *NullableStartVCDBParameters) Set(val *StartVCDBParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableStartVCDBParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableStartVCDBParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartVCDBParameters(val *StartVCDBParameters) *NullableStartVCDBParameters {
	return &NullableStartVCDBParameters{value: val, isSet: true}
}

func (v NullableStartVCDBParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartVCDBParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

