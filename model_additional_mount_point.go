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

// checks if the AdditionalMountPoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdditionalMountPoint{}

// AdditionalMountPoint Specifies an additional location on which to mount a subdirectory of an AppData container.
type AdditionalMountPoint struct {
	// Relative path within the container of the directory that should be mounted.
	SharedPath *string `json:"shared_path,omitempty"`
	// Absolute path on the target environment were the filesystem should be mounted
	MountPath *string `json:"mount_path,omitempty"`
	// The entity ID of the environment on which the file system will be mounted.
	EnvironmentId *string `json:"environment_id,omitempty"`
}

// NewAdditionalMountPoint instantiates a new AdditionalMountPoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalMountPoint() *AdditionalMountPoint {
	this := AdditionalMountPoint{}
	return &this
}

// NewAdditionalMountPointWithDefaults instantiates a new AdditionalMountPoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalMountPointWithDefaults() *AdditionalMountPoint {
	this := AdditionalMountPoint{}
	return &this
}

// GetSharedPath returns the SharedPath field value if set, zero value otherwise.
func (o *AdditionalMountPoint) GetSharedPath() string {
	if o == nil || IsNil(o.SharedPath) {
		var ret string
		return ret
	}
	return *o.SharedPath
}

// GetSharedPathOk returns a tuple with the SharedPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalMountPoint) GetSharedPathOk() (*string, bool) {
	if o == nil || IsNil(o.SharedPath) {
		return nil, false
	}
	return o.SharedPath, true
}

// HasSharedPath returns a boolean if a field has been set.
func (o *AdditionalMountPoint) HasSharedPath() bool {
	if o != nil && !IsNil(o.SharedPath) {
		return true
	}

	return false
}

// SetSharedPath gets a reference to the given string and assigns it to the SharedPath field.
func (o *AdditionalMountPoint) SetSharedPath(v string) {
	o.SharedPath = &v
}

// GetMountPath returns the MountPath field value if set, zero value otherwise.
func (o *AdditionalMountPoint) GetMountPath() string {
	if o == nil || IsNil(o.MountPath) {
		var ret string
		return ret
	}
	return *o.MountPath
}

// GetMountPathOk returns a tuple with the MountPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalMountPoint) GetMountPathOk() (*string, bool) {
	if o == nil || IsNil(o.MountPath) {
		return nil, false
	}
	return o.MountPath, true
}

// HasMountPath returns a boolean if a field has been set.
func (o *AdditionalMountPoint) HasMountPath() bool {
	if o != nil && !IsNil(o.MountPath) {
		return true
	}

	return false
}

// SetMountPath gets a reference to the given string and assigns it to the MountPath field.
func (o *AdditionalMountPoint) SetMountPath(v string) {
	o.MountPath = &v
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *AdditionalMountPoint) GetEnvironmentId() string {
	if o == nil || IsNil(o.EnvironmentId) {
		var ret string
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalMountPoint) GetEnvironmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnvironmentId) {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *AdditionalMountPoint) HasEnvironmentId() bool {
	if o != nil && !IsNil(o.EnvironmentId) {
		return true
	}

	return false
}

// SetEnvironmentId gets a reference to the given string and assigns it to the EnvironmentId field.
func (o *AdditionalMountPoint) SetEnvironmentId(v string) {
	o.EnvironmentId = &v
}

func (o AdditionalMountPoint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionalMountPoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SharedPath) {
		toSerialize["shared_path"] = o.SharedPath
	}
	if !IsNil(o.MountPath) {
		toSerialize["mount_path"] = o.MountPath
	}
	if !IsNil(o.EnvironmentId) {
		toSerialize["environment_id"] = o.EnvironmentId
	}
	return toSerialize, nil
}

type NullableAdditionalMountPoint struct {
	value *AdditionalMountPoint
	isSet bool
}

func (v NullableAdditionalMountPoint) Get() *AdditionalMountPoint {
	return v.value
}

func (v *NullableAdditionalMountPoint) Set(val *AdditionalMountPoint) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalMountPoint) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalMountPoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalMountPoint(val *AdditionalMountPoint) *NullableAdditionalMountPoint {
	return &NullableAdditionalMountPoint{value: val, isSet: true}
}

func (v NullableAdditionalMountPoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalMountPoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


