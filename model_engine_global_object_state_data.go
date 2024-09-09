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

// checks if the EngineGlobalObjectStateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EngineGlobalObjectStateData{}

// EngineGlobalObjectStateData struct for EngineGlobalObjectStateData
type EngineGlobalObjectStateData struct {
	// The entity ID of the engine.
	EngineId *string `json:"engine_id,omitempty"`
	// The name of the engine.
	EngineName *string `json:"engine_name,omitempty"`
	// The number of masking jobs present on the engine.
	JobsCount *int32 `json:"jobs_count,omitempty"`
	// The number of connectors present on the engine.
	ConnectorsCount *int32 `json:"connectors_count,omitempty"`
	// The number of rulesets present on the engine.
	RulesetsCount *int32 `json:"rulesets_count,omitempty"`
	// The revisionHash of the GLOBAL_OBJECT.
	GlobalObjectRevisionHash *string `json:"global_object_revision_hash,omitempty"`
}

// NewEngineGlobalObjectStateData instantiates a new EngineGlobalObjectStateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEngineGlobalObjectStateData() *EngineGlobalObjectStateData {
	this := EngineGlobalObjectStateData{}
	return &this
}

// NewEngineGlobalObjectStateDataWithDefaults instantiates a new EngineGlobalObjectStateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEngineGlobalObjectStateDataWithDefaults() *EngineGlobalObjectStateData {
	this := EngineGlobalObjectStateData{}
	return &this
}

// GetEngineId returns the EngineId field value if set, zero value otherwise.
func (o *EngineGlobalObjectStateData) GetEngineId() string {
	if o == nil || IsNil(o.EngineId) {
		var ret string
		return ret
	}
	return *o.EngineId
}

// GetEngineIdOk returns a tuple with the EngineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineGlobalObjectStateData) GetEngineIdOk() (*string, bool) {
	if o == nil || IsNil(o.EngineId) {
		return nil, false
	}
	return o.EngineId, true
}

// HasEngineId returns a boolean if a field has been set.
func (o *EngineGlobalObjectStateData) HasEngineId() bool {
	if o != nil && !IsNil(o.EngineId) {
		return true
	}

	return false
}

// SetEngineId gets a reference to the given string and assigns it to the EngineId field.
func (o *EngineGlobalObjectStateData) SetEngineId(v string) {
	o.EngineId = &v
}

// GetEngineName returns the EngineName field value if set, zero value otherwise.
func (o *EngineGlobalObjectStateData) GetEngineName() string {
	if o == nil || IsNil(o.EngineName) {
		var ret string
		return ret
	}
	return *o.EngineName
}

// GetEngineNameOk returns a tuple with the EngineName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineGlobalObjectStateData) GetEngineNameOk() (*string, bool) {
	if o == nil || IsNil(o.EngineName) {
		return nil, false
	}
	return o.EngineName, true
}

// HasEngineName returns a boolean if a field has been set.
func (o *EngineGlobalObjectStateData) HasEngineName() bool {
	if o != nil && !IsNil(o.EngineName) {
		return true
	}

	return false
}

// SetEngineName gets a reference to the given string and assigns it to the EngineName field.
func (o *EngineGlobalObjectStateData) SetEngineName(v string) {
	o.EngineName = &v
}

// GetJobsCount returns the JobsCount field value if set, zero value otherwise.
func (o *EngineGlobalObjectStateData) GetJobsCount() int32 {
	if o == nil || IsNil(o.JobsCount) {
		var ret int32
		return ret
	}
	return *o.JobsCount
}

// GetJobsCountOk returns a tuple with the JobsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineGlobalObjectStateData) GetJobsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.JobsCount) {
		return nil, false
	}
	return o.JobsCount, true
}

// HasJobsCount returns a boolean if a field has been set.
func (o *EngineGlobalObjectStateData) HasJobsCount() bool {
	if o != nil && !IsNil(o.JobsCount) {
		return true
	}

	return false
}

// SetJobsCount gets a reference to the given int32 and assigns it to the JobsCount field.
func (o *EngineGlobalObjectStateData) SetJobsCount(v int32) {
	o.JobsCount = &v
}

// GetConnectorsCount returns the ConnectorsCount field value if set, zero value otherwise.
func (o *EngineGlobalObjectStateData) GetConnectorsCount() int32 {
	if o == nil || IsNil(o.ConnectorsCount) {
		var ret int32
		return ret
	}
	return *o.ConnectorsCount
}

// GetConnectorsCountOk returns a tuple with the ConnectorsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineGlobalObjectStateData) GetConnectorsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ConnectorsCount) {
		return nil, false
	}
	return o.ConnectorsCount, true
}

// HasConnectorsCount returns a boolean if a field has been set.
func (o *EngineGlobalObjectStateData) HasConnectorsCount() bool {
	if o != nil && !IsNil(o.ConnectorsCount) {
		return true
	}

	return false
}

// SetConnectorsCount gets a reference to the given int32 and assigns it to the ConnectorsCount field.
func (o *EngineGlobalObjectStateData) SetConnectorsCount(v int32) {
	o.ConnectorsCount = &v
}

// GetRulesetsCount returns the RulesetsCount field value if set, zero value otherwise.
func (o *EngineGlobalObjectStateData) GetRulesetsCount() int32 {
	if o == nil || IsNil(o.RulesetsCount) {
		var ret int32
		return ret
	}
	return *o.RulesetsCount
}

// GetRulesetsCountOk returns a tuple with the RulesetsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineGlobalObjectStateData) GetRulesetsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.RulesetsCount) {
		return nil, false
	}
	return o.RulesetsCount, true
}

// HasRulesetsCount returns a boolean if a field has been set.
func (o *EngineGlobalObjectStateData) HasRulesetsCount() bool {
	if o != nil && !IsNil(o.RulesetsCount) {
		return true
	}

	return false
}

// SetRulesetsCount gets a reference to the given int32 and assigns it to the RulesetsCount field.
func (o *EngineGlobalObjectStateData) SetRulesetsCount(v int32) {
	o.RulesetsCount = &v
}

// GetGlobalObjectRevisionHash returns the GlobalObjectRevisionHash field value if set, zero value otherwise.
func (o *EngineGlobalObjectStateData) GetGlobalObjectRevisionHash() string {
	if o == nil || IsNil(o.GlobalObjectRevisionHash) {
		var ret string
		return ret
	}
	return *o.GlobalObjectRevisionHash
}

// GetGlobalObjectRevisionHashOk returns a tuple with the GlobalObjectRevisionHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineGlobalObjectStateData) GetGlobalObjectRevisionHashOk() (*string, bool) {
	if o == nil || IsNil(o.GlobalObjectRevisionHash) {
		return nil, false
	}
	return o.GlobalObjectRevisionHash, true
}

// HasGlobalObjectRevisionHash returns a boolean if a field has been set.
func (o *EngineGlobalObjectStateData) HasGlobalObjectRevisionHash() bool {
	if o != nil && !IsNil(o.GlobalObjectRevisionHash) {
		return true
	}

	return false
}

// SetGlobalObjectRevisionHash gets a reference to the given string and assigns it to the GlobalObjectRevisionHash field.
func (o *EngineGlobalObjectStateData) SetGlobalObjectRevisionHash(v string) {
	o.GlobalObjectRevisionHash = &v
}

func (o EngineGlobalObjectStateData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EngineGlobalObjectStateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EngineId) {
		toSerialize["engine_id"] = o.EngineId
	}
	if !IsNil(o.EngineName) {
		toSerialize["engine_name"] = o.EngineName
	}
	if !IsNil(o.JobsCount) {
		toSerialize["jobs_count"] = o.JobsCount
	}
	if !IsNil(o.ConnectorsCount) {
		toSerialize["connectors_count"] = o.ConnectorsCount
	}
	if !IsNil(o.RulesetsCount) {
		toSerialize["rulesets_count"] = o.RulesetsCount
	}
	if !IsNil(o.GlobalObjectRevisionHash) {
		toSerialize["global_object_revision_hash"] = o.GlobalObjectRevisionHash
	}
	return toSerialize, nil
}

type NullableEngineGlobalObjectStateData struct {
	value *EngineGlobalObjectStateData
	isSet bool
}

func (v NullableEngineGlobalObjectStateData) Get() *EngineGlobalObjectStateData {
	return v.value
}

func (v *NullableEngineGlobalObjectStateData) Set(val *EngineGlobalObjectStateData) {
	v.value = val
	v.isSet = true
}

func (v NullableEngineGlobalObjectStateData) IsSet() bool {
	return v.isSet
}

func (v *NullableEngineGlobalObjectStateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEngineGlobalObjectStateData(val *EngineGlobalObjectStateData) *NullableEngineGlobalObjectStateData {
	return &NullableEngineGlobalObjectStateData{value: val, isSet: true}
}

func (v NullableEngineGlobalObjectStateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEngineGlobalObjectStateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

