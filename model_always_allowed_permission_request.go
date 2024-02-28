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

// checks if the AlwaysAllowedPermissionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlwaysAllowedPermissionRequest{}

// AlwaysAllowedPermissionRequest struct for AlwaysAllowedPermissionRequest
type AlwaysAllowedPermissionRequest struct {
	// An array of always allowed permissions
	AlwaysAllowedPermissions []AlwaysAllowedPermission `json:"always_allowed_permissions"`
}

// NewAlwaysAllowedPermissionRequest instantiates a new AlwaysAllowedPermissionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlwaysAllowedPermissionRequest(alwaysAllowedPermissions []AlwaysAllowedPermission) *AlwaysAllowedPermissionRequest {
	this := AlwaysAllowedPermissionRequest{}
	this.AlwaysAllowedPermissions = alwaysAllowedPermissions
	return &this
}

// NewAlwaysAllowedPermissionRequestWithDefaults instantiates a new AlwaysAllowedPermissionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlwaysAllowedPermissionRequestWithDefaults() *AlwaysAllowedPermissionRequest {
	this := AlwaysAllowedPermissionRequest{}
	return &this
}

// GetAlwaysAllowedPermissions returns the AlwaysAllowedPermissions field value
func (o *AlwaysAllowedPermissionRequest) GetAlwaysAllowedPermissions() []AlwaysAllowedPermission {
	if o == nil {
		var ret []AlwaysAllowedPermission
		return ret
	}

	return o.AlwaysAllowedPermissions
}

// GetAlwaysAllowedPermissionsOk returns a tuple with the AlwaysAllowedPermissions field value
// and a boolean to check if the value has been set.
func (o *AlwaysAllowedPermissionRequest) GetAlwaysAllowedPermissionsOk() ([]AlwaysAllowedPermission, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlwaysAllowedPermissions, true
}

// SetAlwaysAllowedPermissions sets field value
func (o *AlwaysAllowedPermissionRequest) SetAlwaysAllowedPermissions(v []AlwaysAllowedPermission) {
	o.AlwaysAllowedPermissions = v
}

func (o AlwaysAllowedPermissionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlwaysAllowedPermissionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["always_allowed_permissions"] = o.AlwaysAllowedPermissions
	return toSerialize, nil
}

type NullableAlwaysAllowedPermissionRequest struct {
	value *AlwaysAllowedPermissionRequest
	isSet bool
}

func (v NullableAlwaysAllowedPermissionRequest) Get() *AlwaysAllowedPermissionRequest {
	return v.value
}

func (v *NullableAlwaysAllowedPermissionRequest) Set(val *AlwaysAllowedPermissionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAlwaysAllowedPermissionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAlwaysAllowedPermissionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlwaysAllowedPermissionRequest(val *AlwaysAllowedPermissionRequest) *NullableAlwaysAllowedPermissionRequest {
	return &NullableAlwaysAllowedPermissionRequest{value: val, isSet: true}
}

func (v NullableAlwaysAllowedPermissionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlwaysAllowedPermissionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


