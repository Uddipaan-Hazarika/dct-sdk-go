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

// checks if the PermissionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionsRequest{}

// PermissionsRequest struct for PermissionsRequest
type PermissionsRequest struct {
	// Array of permissions with object type and their permission.
	PermissionObjects []PermissionObject `json:"permission_objects"`
}

// NewPermissionsRequest instantiates a new PermissionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionsRequest(permissionObjects []PermissionObject) *PermissionsRequest {
	this := PermissionsRequest{}
	this.PermissionObjects = permissionObjects
	return &this
}

// NewPermissionsRequestWithDefaults instantiates a new PermissionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionsRequestWithDefaults() *PermissionsRequest {
	this := PermissionsRequest{}
	return &this
}

// GetPermissionObjects returns the PermissionObjects field value
func (o *PermissionsRequest) GetPermissionObjects() []PermissionObject {
	if o == nil {
		var ret []PermissionObject
		return ret
	}

	return o.PermissionObjects
}

// GetPermissionObjectsOk returns a tuple with the PermissionObjects field value
// and a boolean to check if the value has been set.
func (o *PermissionsRequest) GetPermissionObjectsOk() ([]PermissionObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.PermissionObjects, true
}

// SetPermissionObjects sets field value
func (o *PermissionsRequest) SetPermissionObjects(v []PermissionObject) {
	o.PermissionObjects = v
}

func (o PermissionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["permission_objects"] = o.PermissionObjects
	return toSerialize, nil
}

type NullablePermissionsRequest struct {
	value *PermissionsRequest
	isSet bool
}

func (v NullablePermissionsRequest) Get() *PermissionsRequest {
	return v.value
}

func (v *NullablePermissionsRequest) Set(val *PermissionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionsRequest(val *PermissionsRequest) *NullablePermissionsRequest {
	return &NullablePermissionsRequest{value: val, isSet: true}
}

func (v NullablePermissionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


