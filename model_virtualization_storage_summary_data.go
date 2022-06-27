/*
Delphix DCT API

Delphix DCT API

API version: 2.0.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// VirtualizationStorageSummaryData struct for VirtualizationStorageSummaryData
type VirtualizationStorageSummaryData struct {
	// A reference to the engine.
	EngineId string `json:"engine_id"`
	// The engine name.
	EngineName string `json:"engine_name"`
	// The engine hostname.
	EngineHostname string `json:"engine_hostname"`
	// The total amount of storage allocated for engine objects and system metadata, in bytes.
	TotalCapacity *int64 `json:"total_capacity,omitempty"`
	// The amount of available storage, in bytes.
	FreeStorage *int64 `json:"free_storage,omitempty"`
	// The amount of storage used by engine objects and system metadata, in bytes.
	UsedStorage *int64 `json:"used_storage,omitempty"`
	// The percentage of storage used.
	UsedPercentage *float32 `json:"used_percentage,omitempty"`
	// The number of dSources on the engine.
	DsourceCount *int64 `json:"dsource_count,omitempty"`
	// The number of VDBs on the engine.
	VdbCount *int64 `json:"vdb_count,omitempty"`
	// The total number of dSources and VDBs on the engine.
	TotalObjectCount *int64 `json:"total_object_count,omitempty"`
}

// NewVirtualizationStorageSummaryData instantiates a new VirtualizationStorageSummaryData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationStorageSummaryData(engineId string, engineName string, engineHostname string) *VirtualizationStorageSummaryData {
	this := VirtualizationStorageSummaryData{}
	this.EngineId = engineId
	this.EngineName = engineName
	this.EngineHostname = engineHostname
	return &this
}

// NewVirtualizationStorageSummaryDataWithDefaults instantiates a new VirtualizationStorageSummaryData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationStorageSummaryDataWithDefaults() *VirtualizationStorageSummaryData {
	this := VirtualizationStorageSummaryData{}
	return &this
}

// GetEngineId returns the EngineId field value
func (o *VirtualizationStorageSummaryData) GetEngineId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EngineId
}

// GetEngineIdOk returns a tuple with the EngineId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationStorageSummaryData) GetEngineIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EngineId, true
}

// SetEngineId sets field value
func (o *VirtualizationStorageSummaryData) SetEngineId(v string) {
	o.EngineId = v
}

// GetEngineName returns the EngineName field value
func (o *VirtualizationStorageSummaryData) GetEngineName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EngineName
}

// GetEngineNameOk returns a tuple with the EngineName field value
// and a boolean to check if the value has been set.
func (o *VirtualizationStorageSummaryData) GetEngineNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EngineName, true
}

// SetEngineName sets field value
func (o *VirtualizationStorageSummaryData) SetEngineName(v string) {
	o.EngineName = v
}

// GetEngineHostname returns the EngineHostname field value
func (o *VirtualizationStorageSummaryData) GetEngineHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EngineHostname
}

// GetEngineHostnameOk returns a tuple with the EngineHostname field value
// and a boolean to check if the value has been set.
func (o *VirtualizationStorageSummaryData) GetEngineHostnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EngineHostname, true
}

// SetEngineHostname sets field value
func (o *VirtualizationStorageSummaryData) SetEngineHostname(v string) {
	o.EngineHostname = v
}

// GetTotalCapacity returns the TotalCapacity field value if set, zero value otherwise.
func (o *VirtualizationStorageSummaryData) GetTotalCapacity() int64 {
	if o == nil || o.TotalCapacity == nil {
		var ret int64
		return ret
	}
	return *o.TotalCapacity
}

// GetTotalCapacityOk returns a tuple with the TotalCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationStorageSummaryData) GetTotalCapacityOk() (*int64, bool) {
	if o == nil || o.TotalCapacity == nil {
		return nil, false
	}
	return o.TotalCapacity, true
}

// HasTotalCapacity returns a boolean if a field has been set.
func (o *VirtualizationStorageSummaryData) HasTotalCapacity() bool {
	if o != nil && o.TotalCapacity != nil {
		return true
	}

	return false
}

// SetTotalCapacity gets a reference to the given int64 and assigns it to the TotalCapacity field.
func (o *VirtualizationStorageSummaryData) SetTotalCapacity(v int64) {
	o.TotalCapacity = &v
}

// GetFreeStorage returns the FreeStorage field value if set, zero value otherwise.
func (o *VirtualizationStorageSummaryData) GetFreeStorage() int64 {
	if o == nil || o.FreeStorage == nil {
		var ret int64
		return ret
	}
	return *o.FreeStorage
}

// GetFreeStorageOk returns a tuple with the FreeStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationStorageSummaryData) GetFreeStorageOk() (*int64, bool) {
	if o == nil || o.FreeStorage == nil {
		return nil, false
	}
	return o.FreeStorage, true
}

// HasFreeStorage returns a boolean if a field has been set.
func (o *VirtualizationStorageSummaryData) HasFreeStorage() bool {
	if o != nil && o.FreeStorage != nil {
		return true
	}

	return false
}

// SetFreeStorage gets a reference to the given int64 and assigns it to the FreeStorage field.
func (o *VirtualizationStorageSummaryData) SetFreeStorage(v int64) {
	o.FreeStorage = &v
}

// GetUsedStorage returns the UsedStorage field value if set, zero value otherwise.
func (o *VirtualizationStorageSummaryData) GetUsedStorage() int64 {
	if o == nil || o.UsedStorage == nil {
		var ret int64
		return ret
	}
	return *o.UsedStorage
}

// GetUsedStorageOk returns a tuple with the UsedStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationStorageSummaryData) GetUsedStorageOk() (*int64, bool) {
	if o == nil || o.UsedStorage == nil {
		return nil, false
	}
	return o.UsedStorage, true
}

// HasUsedStorage returns a boolean if a field has been set.
func (o *VirtualizationStorageSummaryData) HasUsedStorage() bool {
	if o != nil && o.UsedStorage != nil {
		return true
	}

	return false
}

// SetUsedStorage gets a reference to the given int64 and assigns it to the UsedStorage field.
func (o *VirtualizationStorageSummaryData) SetUsedStorage(v int64) {
	o.UsedStorage = &v
}

// GetUsedPercentage returns the UsedPercentage field value if set, zero value otherwise.
func (o *VirtualizationStorageSummaryData) GetUsedPercentage() float32 {
	if o == nil || o.UsedPercentage == nil {
		var ret float32
		return ret
	}
	return *o.UsedPercentage
}

// GetUsedPercentageOk returns a tuple with the UsedPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationStorageSummaryData) GetUsedPercentageOk() (*float32, bool) {
	if o == nil || o.UsedPercentage == nil {
		return nil, false
	}
	return o.UsedPercentage, true
}

// HasUsedPercentage returns a boolean if a field has been set.
func (o *VirtualizationStorageSummaryData) HasUsedPercentage() bool {
	if o != nil && o.UsedPercentage != nil {
		return true
	}

	return false
}

// SetUsedPercentage gets a reference to the given float32 and assigns it to the UsedPercentage field.
func (o *VirtualizationStorageSummaryData) SetUsedPercentage(v float32) {
	o.UsedPercentage = &v
}

// GetDsourceCount returns the DsourceCount field value if set, zero value otherwise.
func (o *VirtualizationStorageSummaryData) GetDsourceCount() int64 {
	if o == nil || o.DsourceCount == nil {
		var ret int64
		return ret
	}
	return *o.DsourceCount
}

// GetDsourceCountOk returns a tuple with the DsourceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationStorageSummaryData) GetDsourceCountOk() (*int64, bool) {
	if o == nil || o.DsourceCount == nil {
		return nil, false
	}
	return o.DsourceCount, true
}

// HasDsourceCount returns a boolean if a field has been set.
func (o *VirtualizationStorageSummaryData) HasDsourceCount() bool {
	if o != nil && o.DsourceCount != nil {
		return true
	}

	return false
}

// SetDsourceCount gets a reference to the given int64 and assigns it to the DsourceCount field.
func (o *VirtualizationStorageSummaryData) SetDsourceCount(v int64) {
	o.DsourceCount = &v
}

// GetVdbCount returns the VdbCount field value if set, zero value otherwise.
func (o *VirtualizationStorageSummaryData) GetVdbCount() int64 {
	if o == nil || o.VdbCount == nil {
		var ret int64
		return ret
	}
	return *o.VdbCount
}

// GetVdbCountOk returns a tuple with the VdbCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationStorageSummaryData) GetVdbCountOk() (*int64, bool) {
	if o == nil || o.VdbCount == nil {
		return nil, false
	}
	return o.VdbCount, true
}

// HasVdbCount returns a boolean if a field has been set.
func (o *VirtualizationStorageSummaryData) HasVdbCount() bool {
	if o != nil && o.VdbCount != nil {
		return true
	}

	return false
}

// SetVdbCount gets a reference to the given int64 and assigns it to the VdbCount field.
func (o *VirtualizationStorageSummaryData) SetVdbCount(v int64) {
	o.VdbCount = &v
}

// GetTotalObjectCount returns the TotalObjectCount field value if set, zero value otherwise.
func (o *VirtualizationStorageSummaryData) GetTotalObjectCount() int64 {
	if o == nil || o.TotalObjectCount == nil {
		var ret int64
		return ret
	}
	return *o.TotalObjectCount
}

// GetTotalObjectCountOk returns a tuple with the TotalObjectCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationStorageSummaryData) GetTotalObjectCountOk() (*int64, bool) {
	if o == nil || o.TotalObjectCount == nil {
		return nil, false
	}
	return o.TotalObjectCount, true
}

// HasTotalObjectCount returns a boolean if a field has been set.
func (o *VirtualizationStorageSummaryData) HasTotalObjectCount() bool {
	if o != nil && o.TotalObjectCount != nil {
		return true
	}

	return false
}

// SetTotalObjectCount gets a reference to the given int64 and assigns it to the TotalObjectCount field.
func (o *VirtualizationStorageSummaryData) SetTotalObjectCount(v int64) {
	o.TotalObjectCount = &v
}

func (o VirtualizationStorageSummaryData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["engine_id"] = o.EngineId
	}
	if true {
		toSerialize["engine_name"] = o.EngineName
	}
	if true {
		toSerialize["engine_hostname"] = o.EngineHostname
	}
	if o.TotalCapacity != nil {
		toSerialize["total_capacity"] = o.TotalCapacity
	}
	if o.FreeStorage != nil {
		toSerialize["free_storage"] = o.FreeStorage
	}
	if o.UsedStorage != nil {
		toSerialize["used_storage"] = o.UsedStorage
	}
	if o.UsedPercentage != nil {
		toSerialize["used_percentage"] = o.UsedPercentage
	}
	if o.DsourceCount != nil {
		toSerialize["dsource_count"] = o.DsourceCount
	}
	if o.VdbCount != nil {
		toSerialize["vdb_count"] = o.VdbCount
	}
	if o.TotalObjectCount != nil {
		toSerialize["total_object_count"] = o.TotalObjectCount
	}
	return json.Marshal(toSerialize)
}

type NullableVirtualizationStorageSummaryData struct {
	value *VirtualizationStorageSummaryData
	isSet bool
}

func (v NullableVirtualizationStorageSummaryData) Get() *VirtualizationStorageSummaryData {
	return v.value
}

func (v *NullableVirtualizationStorageSummaryData) Set(val *VirtualizationStorageSummaryData) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationStorageSummaryData) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationStorageSummaryData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationStorageSummaryData(val *VirtualizationStorageSummaryData) *NullableVirtualizationStorageSummaryData {
	return &NullableVirtualizationStorageSummaryData{value: val, isSet: true}
}

func (v NullableVirtualizationStorageSummaryData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationStorageSummaryData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


