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

// checks if the OracleRacCustomEnvFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OracleRacCustomEnvFile{}

// OracleRacCustomEnvFile struct for OracleRacCustomEnvFile
type OracleRacCustomEnvFile struct {
	NodeId *string `json:"node_id,omitempty"`
	PathParameters *string `json:"path_parameters,omitempty"`
}

// NewOracleRacCustomEnvFile instantiates a new OracleRacCustomEnvFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOracleRacCustomEnvFile() *OracleRacCustomEnvFile {
	this := OracleRacCustomEnvFile{}
	return &this
}

// NewOracleRacCustomEnvFileWithDefaults instantiates a new OracleRacCustomEnvFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOracleRacCustomEnvFileWithDefaults() *OracleRacCustomEnvFile {
	this := OracleRacCustomEnvFile{}
	return &this
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *OracleRacCustomEnvFile) GetNodeId() string {
	if o == nil || IsNil(o.NodeId) {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleRacCustomEnvFile) GetNodeIdOk() (*string, bool) {
	if o == nil || IsNil(o.NodeId) {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *OracleRacCustomEnvFile) HasNodeId() bool {
	if o != nil && !IsNil(o.NodeId) {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *OracleRacCustomEnvFile) SetNodeId(v string) {
	o.NodeId = &v
}

// GetPathParameters returns the PathParameters field value if set, zero value otherwise.
func (o *OracleRacCustomEnvFile) GetPathParameters() string {
	if o == nil || IsNil(o.PathParameters) {
		var ret string
		return ret
	}
	return *o.PathParameters
}

// GetPathParametersOk returns a tuple with the PathParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleRacCustomEnvFile) GetPathParametersOk() (*string, bool) {
	if o == nil || IsNil(o.PathParameters) {
		return nil, false
	}
	return o.PathParameters, true
}

// HasPathParameters returns a boolean if a field has been set.
func (o *OracleRacCustomEnvFile) HasPathParameters() bool {
	if o != nil && !IsNil(o.PathParameters) {
		return true
	}

	return false
}

// SetPathParameters gets a reference to the given string and assigns it to the PathParameters field.
func (o *OracleRacCustomEnvFile) SetPathParameters(v string) {
	o.PathParameters = &v
}

func (o OracleRacCustomEnvFile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OracleRacCustomEnvFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NodeId) {
		toSerialize["node_id"] = o.NodeId
	}
	if !IsNil(o.PathParameters) {
		toSerialize["path_parameters"] = o.PathParameters
	}
	return toSerialize, nil
}

type NullableOracleRacCustomEnvFile struct {
	value *OracleRacCustomEnvFile
	isSet bool
}

func (v NullableOracleRacCustomEnvFile) Get() *OracleRacCustomEnvFile {
	return v.value
}

func (v *NullableOracleRacCustomEnvFile) Set(val *OracleRacCustomEnvFile) {
	v.value = val
	v.isSet = true
}

func (v NullableOracleRacCustomEnvFile) IsSet() bool {
	return v.isSet
}

func (v *NullableOracleRacCustomEnvFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOracleRacCustomEnvFile(val *OracleRacCustomEnvFile) *NullableOracleRacCustomEnvFile {
	return &NullableOracleRacCustomEnvFile{value: val, isSet: true}
}

func (v NullableOracleRacCustomEnvFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOracleRacCustomEnvFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


