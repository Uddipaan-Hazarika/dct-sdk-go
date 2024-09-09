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
	"time"
)

// checks if the BundleUploadEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BundleUploadEvent{}

// BundleUploadEvent Bundle Upload Event
type BundleUploadEvent struct {
	// The ID of the bundle generation event.
	Id *string `json:"id,omitempty"`
	// The time when the bundle generation started.
	GenerationStartTime *time.Time `json:"generationStartTime,omitempty"`
	// The time when the bundle generation ended.
	GenerationEndTime *time.Time `json:"generationEndTime,omitempty"`
	// True if the bundle was fully built, False if it was trimmed short due to constraints
	AllGapsFilled *bool `json:"allGapsFilled,omitempty"`
	// The list of dates for which the data is included in the bundle.
	DataDates []string `json:"dataDates,omitempty"`
}

// NewBundleUploadEvent instantiates a new BundleUploadEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleUploadEvent() *BundleUploadEvent {
	this := BundleUploadEvent{}
	return &this
}

// NewBundleUploadEventWithDefaults instantiates a new BundleUploadEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleUploadEventWithDefaults() *BundleUploadEvent {
	this := BundleUploadEvent{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BundleUploadEvent) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleUploadEvent) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BundleUploadEvent) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BundleUploadEvent) SetId(v string) {
	o.Id = &v
}

// GetGenerationStartTime returns the GenerationStartTime field value if set, zero value otherwise.
func (o *BundleUploadEvent) GetGenerationStartTime() time.Time {
	if o == nil || IsNil(o.GenerationStartTime) {
		var ret time.Time
		return ret
	}
	return *o.GenerationStartTime
}

// GetGenerationStartTimeOk returns a tuple with the GenerationStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleUploadEvent) GetGenerationStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.GenerationStartTime) {
		return nil, false
	}
	return o.GenerationStartTime, true
}

// HasGenerationStartTime returns a boolean if a field has been set.
func (o *BundleUploadEvent) HasGenerationStartTime() bool {
	if o != nil && !IsNil(o.GenerationStartTime) {
		return true
	}

	return false
}

// SetGenerationStartTime gets a reference to the given time.Time and assigns it to the GenerationStartTime field.
func (o *BundleUploadEvent) SetGenerationStartTime(v time.Time) {
	o.GenerationStartTime = &v
}

// GetGenerationEndTime returns the GenerationEndTime field value if set, zero value otherwise.
func (o *BundleUploadEvent) GetGenerationEndTime() time.Time {
	if o == nil || IsNil(o.GenerationEndTime) {
		var ret time.Time
		return ret
	}
	return *o.GenerationEndTime
}

// GetGenerationEndTimeOk returns a tuple with the GenerationEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleUploadEvent) GetGenerationEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.GenerationEndTime) {
		return nil, false
	}
	return o.GenerationEndTime, true
}

// HasGenerationEndTime returns a boolean if a field has been set.
func (o *BundleUploadEvent) HasGenerationEndTime() bool {
	if o != nil && !IsNil(o.GenerationEndTime) {
		return true
	}

	return false
}

// SetGenerationEndTime gets a reference to the given time.Time and assigns it to the GenerationEndTime field.
func (o *BundleUploadEvent) SetGenerationEndTime(v time.Time) {
	o.GenerationEndTime = &v
}

// GetAllGapsFilled returns the AllGapsFilled field value if set, zero value otherwise.
func (o *BundleUploadEvent) GetAllGapsFilled() bool {
	if o == nil || IsNil(o.AllGapsFilled) {
		var ret bool
		return ret
	}
	return *o.AllGapsFilled
}

// GetAllGapsFilledOk returns a tuple with the AllGapsFilled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleUploadEvent) GetAllGapsFilledOk() (*bool, bool) {
	if o == nil || IsNil(o.AllGapsFilled) {
		return nil, false
	}
	return o.AllGapsFilled, true
}

// HasAllGapsFilled returns a boolean if a field has been set.
func (o *BundleUploadEvent) HasAllGapsFilled() bool {
	if o != nil && !IsNil(o.AllGapsFilled) {
		return true
	}

	return false
}

// SetAllGapsFilled gets a reference to the given bool and assigns it to the AllGapsFilled field.
func (o *BundleUploadEvent) SetAllGapsFilled(v bool) {
	o.AllGapsFilled = &v
}

// GetDataDates returns the DataDates field value if set, zero value otherwise.
func (o *BundleUploadEvent) GetDataDates() []string {
	if o == nil || IsNil(o.DataDates) {
		var ret []string
		return ret
	}
	return o.DataDates
}

// GetDataDatesOk returns a tuple with the DataDates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleUploadEvent) GetDataDatesOk() ([]string, bool) {
	if o == nil || IsNil(o.DataDates) {
		return nil, false
	}
	return o.DataDates, true
}

// HasDataDates returns a boolean if a field has been set.
func (o *BundleUploadEvent) HasDataDates() bool {
	if o != nil && !IsNil(o.DataDates) {
		return true
	}

	return false
}

// SetDataDates gets a reference to the given []string and assigns it to the DataDates field.
func (o *BundleUploadEvent) SetDataDates(v []string) {
	o.DataDates = v
}

func (o BundleUploadEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BundleUploadEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.GenerationStartTime) {
		toSerialize["generationStartTime"] = o.GenerationStartTime
	}
	if !IsNil(o.GenerationEndTime) {
		toSerialize["generationEndTime"] = o.GenerationEndTime
	}
	if !IsNil(o.AllGapsFilled) {
		toSerialize["allGapsFilled"] = o.AllGapsFilled
	}
	if !IsNil(o.DataDates) {
		toSerialize["dataDates"] = o.DataDates
	}
	return toSerialize, nil
}

type NullableBundleUploadEvent struct {
	value *BundleUploadEvent
	isSet bool
}

func (v NullableBundleUploadEvent) Get() *BundleUploadEvent {
	return v.value
}

func (v *NullableBundleUploadEvent) Set(val *BundleUploadEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleUploadEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleUploadEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleUploadEvent(val *BundleUploadEvent) *NullableBundleUploadEvent {
	return &NullableBundleUploadEvent{value: val, isSet: true}
}

func (v NullableBundleUploadEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleUploadEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

