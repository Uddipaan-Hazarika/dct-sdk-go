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

// checks if the VirtualizationSchedule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualizationSchedule{}

// VirtualizationSchedule struct for VirtualizationSchedule
type VirtualizationSchedule struct {
	CronString string `json:"cron_string"`
	CutoffTime int64 `json:"cutoff_time"`
}

// NewVirtualizationSchedule instantiates a new VirtualizationSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationSchedule(cronString string, cutoffTime int64) *VirtualizationSchedule {
	this := VirtualizationSchedule{}
	this.CronString = cronString
	this.CutoffTime = cutoffTime
	return &this
}

// NewVirtualizationScheduleWithDefaults instantiates a new VirtualizationSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationScheduleWithDefaults() *VirtualizationSchedule {
	this := VirtualizationSchedule{}
	return &this
}

// GetCronString returns the CronString field value
func (o *VirtualizationSchedule) GetCronString() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CronString
}

// GetCronStringOk returns a tuple with the CronString field value
// and a boolean to check if the value has been set.
func (o *VirtualizationSchedule) GetCronStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CronString, true
}

// SetCronString sets field value
func (o *VirtualizationSchedule) SetCronString(v string) {
	o.CronString = v
}

// GetCutoffTime returns the CutoffTime field value
func (o *VirtualizationSchedule) GetCutoffTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CutoffTime
}

// GetCutoffTimeOk returns a tuple with the CutoffTime field value
// and a boolean to check if the value has been set.
func (o *VirtualizationSchedule) GetCutoffTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CutoffTime, true
}

// SetCutoffTime sets field value
func (o *VirtualizationSchedule) SetCutoffTime(v int64) {
	o.CutoffTime = v
}

func (o VirtualizationSchedule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualizationSchedule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cron_string"] = o.CronString
	toSerialize["cutoff_time"] = o.CutoffTime
	return toSerialize, nil
}

type NullableVirtualizationSchedule struct {
	value *VirtualizationSchedule
	isSet bool
}

func (v NullableVirtualizationSchedule) Get() *VirtualizationSchedule {
	return v.value
}

func (v *NullableVirtualizationSchedule) Set(val *VirtualizationSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationSchedule(val *VirtualizationSchedule) *NullableVirtualizationSchedule {
	return &NullableVirtualizationSchedule{value: val, isSet: true}
}

func (v NullableVirtualizationSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


