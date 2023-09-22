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

// checks if the EnableScaleTestingRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnableScaleTestingRequest{}

// EnableScaleTestingRequest struct for EnableScaleTestingRequest
type EnableScaleTestingRequest struct {
	// no. of times same engine needs to be registered
	EnginesCount int32 `json:"engines_count"`
	// list of engine hostnames to be registered engines_count times
	EnginesList []string `json:"engines_list"`
	// no. of times to duplicate sources, containers, and timeflows
	VirtObjectsCount int32 `json:"virt_objects_count"`
	// no. of times to duplicate snapshots
	SnapshotsCount int32 `json:"snapshots_count"`
}

// NewEnableScaleTestingRequest instantiates a new EnableScaleTestingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnableScaleTestingRequest(enginesCount int32, enginesList []string, virtObjectsCount int32, snapshotsCount int32) *EnableScaleTestingRequest {
	this := EnableScaleTestingRequest{}
	this.EnginesCount = enginesCount
	this.EnginesList = enginesList
	this.VirtObjectsCount = virtObjectsCount
	this.SnapshotsCount = snapshotsCount
	return &this
}

// NewEnableScaleTestingRequestWithDefaults instantiates a new EnableScaleTestingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnableScaleTestingRequestWithDefaults() *EnableScaleTestingRequest {
	this := EnableScaleTestingRequest{}
	return &this
}

// GetEnginesCount returns the EnginesCount field value
func (o *EnableScaleTestingRequest) GetEnginesCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EnginesCount
}

// GetEnginesCountOk returns a tuple with the EnginesCount field value
// and a boolean to check if the value has been set.
func (o *EnableScaleTestingRequest) GetEnginesCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnginesCount, true
}

// SetEnginesCount sets field value
func (o *EnableScaleTestingRequest) SetEnginesCount(v int32) {
	o.EnginesCount = v
}

// GetEnginesList returns the EnginesList field value
func (o *EnableScaleTestingRequest) GetEnginesList() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.EnginesList
}

// GetEnginesListOk returns a tuple with the EnginesList field value
// and a boolean to check if the value has been set.
func (o *EnableScaleTestingRequest) GetEnginesListOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnginesList, true
}

// SetEnginesList sets field value
func (o *EnableScaleTestingRequest) SetEnginesList(v []string) {
	o.EnginesList = v
}

// GetVirtObjectsCount returns the VirtObjectsCount field value
func (o *EnableScaleTestingRequest) GetVirtObjectsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VirtObjectsCount
}

// GetVirtObjectsCountOk returns a tuple with the VirtObjectsCount field value
// and a boolean to check if the value has been set.
func (o *EnableScaleTestingRequest) GetVirtObjectsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VirtObjectsCount, true
}

// SetVirtObjectsCount sets field value
func (o *EnableScaleTestingRequest) SetVirtObjectsCount(v int32) {
	o.VirtObjectsCount = v
}

// GetSnapshotsCount returns the SnapshotsCount field value
func (o *EnableScaleTestingRequest) GetSnapshotsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SnapshotsCount
}

// GetSnapshotsCountOk returns a tuple with the SnapshotsCount field value
// and a boolean to check if the value has been set.
func (o *EnableScaleTestingRequest) GetSnapshotsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SnapshotsCount, true
}

// SetSnapshotsCount sets field value
func (o *EnableScaleTestingRequest) SetSnapshotsCount(v int32) {
	o.SnapshotsCount = v
}

func (o EnableScaleTestingRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnableScaleTestingRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["engines_count"] = o.EnginesCount
	toSerialize["engines_list"] = o.EnginesList
	toSerialize["virt_objects_count"] = o.VirtObjectsCount
	toSerialize["snapshots_count"] = o.SnapshotsCount
	return toSerialize, nil
}

type NullableEnableScaleTestingRequest struct {
	value *EnableScaleTestingRequest
	isSet bool
}

func (v NullableEnableScaleTestingRequest) Get() *EnableScaleTestingRequest {
	return v.value
}

func (v *NullableEnableScaleTestingRequest) Set(val *EnableScaleTestingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEnableScaleTestingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEnableScaleTestingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnableScaleTestingRequest(val *EnableScaleTestingRequest) *NullableEnableScaleTestingRequest {
	return &NullableEnableScaleTestingRequest{value: val, isSet: true}
}

func (v NullableEnableScaleTestingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnableScaleTestingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


