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

// checks if the OracleClusterNodeInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OracleClusterNodeInstance{}

// OracleClusterNodeInstance An oracle cluster node instance.
type OracleClusterNodeInstance struct {
	// The name of this instance.
	InstanceName *string `json:"instance_name,omitempty"`
	// The number of this instance.
	InstanceNumber *int32 `json:"instance_number,omitempty"`
}

// NewOracleClusterNodeInstance instantiates a new OracleClusterNodeInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOracleClusterNodeInstance() *OracleClusterNodeInstance {
	this := OracleClusterNodeInstance{}
	return &this
}

// NewOracleClusterNodeInstanceWithDefaults instantiates a new OracleClusterNodeInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOracleClusterNodeInstanceWithDefaults() *OracleClusterNodeInstance {
	this := OracleClusterNodeInstance{}
	return &this
}

// GetInstanceName returns the InstanceName field value if set, zero value otherwise.
func (o *OracleClusterNodeInstance) GetInstanceName() string {
	if o == nil || IsNil(o.InstanceName) {
		var ret string
		return ret
	}
	return *o.InstanceName
}

// GetInstanceNameOk returns a tuple with the InstanceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleClusterNodeInstance) GetInstanceNameOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceName) {
		return nil, false
	}
	return o.InstanceName, true
}

// HasInstanceName returns a boolean if a field has been set.
func (o *OracleClusterNodeInstance) HasInstanceName() bool {
	if o != nil && !IsNil(o.InstanceName) {
		return true
	}

	return false
}

// SetInstanceName gets a reference to the given string and assigns it to the InstanceName field.
func (o *OracleClusterNodeInstance) SetInstanceName(v string) {
	o.InstanceName = &v
}

// GetInstanceNumber returns the InstanceNumber field value if set, zero value otherwise.
func (o *OracleClusterNodeInstance) GetInstanceNumber() int32 {
	if o == nil || IsNil(o.InstanceNumber) {
		var ret int32
		return ret
	}
	return *o.InstanceNumber
}

// GetInstanceNumberOk returns a tuple with the InstanceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleClusterNodeInstance) GetInstanceNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.InstanceNumber) {
		return nil, false
	}
	return o.InstanceNumber, true
}

// HasInstanceNumber returns a boolean if a field has been set.
func (o *OracleClusterNodeInstance) HasInstanceNumber() bool {
	if o != nil && !IsNil(o.InstanceNumber) {
		return true
	}

	return false
}

// SetInstanceNumber gets a reference to the given int32 and assigns it to the InstanceNumber field.
func (o *OracleClusterNodeInstance) SetInstanceNumber(v int32) {
	o.InstanceNumber = &v
}

func (o OracleClusterNodeInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OracleClusterNodeInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InstanceName) {
		toSerialize["instance_name"] = o.InstanceName
	}
	if !IsNil(o.InstanceNumber) {
		toSerialize["instance_number"] = o.InstanceNumber
	}
	return toSerialize, nil
}

type NullableOracleClusterNodeInstance struct {
	value *OracleClusterNodeInstance
	isSet bool
}

func (v NullableOracleClusterNodeInstance) Get() *OracleClusterNodeInstance {
	return v.value
}

func (v *NullableOracleClusterNodeInstance) Set(val *OracleClusterNodeInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableOracleClusterNodeInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableOracleClusterNodeInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOracleClusterNodeInstance(val *OracleClusterNodeInstance) *NullableOracleClusterNodeInstance {
	return &NullableOracleClusterNodeInstance{value: val, isSet: true}
}

func (v NullableOracleClusterNodeInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOracleClusterNodeInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


