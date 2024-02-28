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

// checks if the RefreshVDBByLocationParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RefreshVDBByLocationParameters{}

// RefreshVDBByLocationParameters Parameters to refresh by a database-specific identifier (SCN, LSN, etc).
type RefreshVDBByLocationParameters struct {
	// The database specific identifier for tracking transactions (SCN, LSN, etc).
	Location *string `json:"location,omitempty"`
	// ID of the dataset to refresh to, mutually exclusive with timeflow_id.
	DatasetId *string `json:"dataset_id,omitempty"`
	// ID of the timeflow to refresh to, mutually exclusive with dataset_id.
	TimeflowId *string `json:"timeflow_id,omitempty"`
}

// NewRefreshVDBByLocationParameters instantiates a new RefreshVDBByLocationParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefreshVDBByLocationParameters() *RefreshVDBByLocationParameters {
	this := RefreshVDBByLocationParameters{}
	return &this
}

// NewRefreshVDBByLocationParametersWithDefaults instantiates a new RefreshVDBByLocationParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefreshVDBByLocationParametersWithDefaults() *RefreshVDBByLocationParameters {
	this := RefreshVDBByLocationParameters{}
	return &this
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *RefreshVDBByLocationParameters) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshVDBByLocationParameters) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *RefreshVDBByLocationParameters) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *RefreshVDBByLocationParameters) SetLocation(v string) {
	o.Location = &v
}

// GetDatasetId returns the DatasetId field value if set, zero value otherwise.
func (o *RefreshVDBByLocationParameters) GetDatasetId() string {
	if o == nil || IsNil(o.DatasetId) {
		var ret string
		return ret
	}
	return *o.DatasetId
}

// GetDatasetIdOk returns a tuple with the DatasetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshVDBByLocationParameters) GetDatasetIdOk() (*string, bool) {
	if o == nil || IsNil(o.DatasetId) {
		return nil, false
	}
	return o.DatasetId, true
}

// HasDatasetId returns a boolean if a field has been set.
func (o *RefreshVDBByLocationParameters) HasDatasetId() bool {
	if o != nil && !IsNil(o.DatasetId) {
		return true
	}

	return false
}

// SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.
func (o *RefreshVDBByLocationParameters) SetDatasetId(v string) {
	o.DatasetId = &v
}

// GetTimeflowId returns the TimeflowId field value if set, zero value otherwise.
func (o *RefreshVDBByLocationParameters) GetTimeflowId() string {
	if o == nil || IsNil(o.TimeflowId) {
		var ret string
		return ret
	}
	return *o.TimeflowId
}

// GetTimeflowIdOk returns a tuple with the TimeflowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshVDBByLocationParameters) GetTimeflowIdOk() (*string, bool) {
	if o == nil || IsNil(o.TimeflowId) {
		return nil, false
	}
	return o.TimeflowId, true
}

// HasTimeflowId returns a boolean if a field has been set.
func (o *RefreshVDBByLocationParameters) HasTimeflowId() bool {
	if o != nil && !IsNil(o.TimeflowId) {
		return true
	}

	return false
}

// SetTimeflowId gets a reference to the given string and assigns it to the TimeflowId field.
func (o *RefreshVDBByLocationParameters) SetTimeflowId(v string) {
	o.TimeflowId = &v
}

func (o RefreshVDBByLocationParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefreshVDBByLocationParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.DatasetId) {
		toSerialize["dataset_id"] = o.DatasetId
	}
	if !IsNil(o.TimeflowId) {
		toSerialize["timeflow_id"] = o.TimeflowId
	}
	return toSerialize, nil
}

type NullableRefreshVDBByLocationParameters struct {
	value *RefreshVDBByLocationParameters
	isSet bool
}

func (v NullableRefreshVDBByLocationParameters) Get() *RefreshVDBByLocationParameters {
	return v.value
}

func (v *NullableRefreshVDBByLocationParameters) Set(val *RefreshVDBByLocationParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableRefreshVDBByLocationParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableRefreshVDBByLocationParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefreshVDBByLocationParameters(val *RefreshVDBByLocationParameters) *NullableRefreshVDBByLocationParameters {
	return &NullableRefreshVDBByLocationParameters{value: val, isSet: true}
}

func (v NullableRefreshVDBByLocationParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefreshVDBByLocationParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


