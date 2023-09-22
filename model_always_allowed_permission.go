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

// checks if the AlwaysAllowedPermission type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlwaysAllowedPermission{}

// AlwaysAllowedPermission struct for AlwaysAllowedPermission
type AlwaysAllowedPermission struct {
	ObjectType ObjectTypeEnum `json:"object_type"`
	Permission PermissionEnum `json:"permission"`
}

// NewAlwaysAllowedPermission instantiates a new AlwaysAllowedPermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlwaysAllowedPermission(objectType ObjectTypeEnum, permission PermissionEnum) *AlwaysAllowedPermission {
	this := AlwaysAllowedPermission{}
	this.ObjectType = objectType
	this.Permission = permission
	return &this
}

// NewAlwaysAllowedPermissionWithDefaults instantiates a new AlwaysAllowedPermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlwaysAllowedPermissionWithDefaults() *AlwaysAllowedPermission {
	this := AlwaysAllowedPermission{}
	return &this
}

// GetObjectType returns the ObjectType field value
func (o *AlwaysAllowedPermission) GetObjectType() ObjectTypeEnum {
	if o == nil {
		var ret ObjectTypeEnum
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AlwaysAllowedPermission) GetObjectTypeOk() (*ObjectTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AlwaysAllowedPermission) SetObjectType(v ObjectTypeEnum) {
	o.ObjectType = v
}

// GetPermission returns the Permission field value
func (o *AlwaysAllowedPermission) GetPermission() PermissionEnum {
	if o == nil {
		var ret PermissionEnum
		return ret
	}

	return o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value
// and a boolean to check if the value has been set.
func (o *AlwaysAllowedPermission) GetPermissionOk() (*PermissionEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permission, true
}

// SetPermission sets field value
func (o *AlwaysAllowedPermission) SetPermission(v PermissionEnum) {
	o.Permission = v
}

func (o AlwaysAllowedPermission) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlwaysAllowedPermission) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object_type"] = o.ObjectType
	toSerialize["permission"] = o.Permission
	return toSerialize, nil
}

type NullableAlwaysAllowedPermission struct {
	value *AlwaysAllowedPermission
	isSet bool
}

func (v NullableAlwaysAllowedPermission) Get() *AlwaysAllowedPermission {
	return v.value
}

func (v *NullableAlwaysAllowedPermission) Set(val *AlwaysAllowedPermission) {
	v.value = val
	v.isSet = true
}

func (v NullableAlwaysAllowedPermission) IsSet() bool {
	return v.isSet
}

func (v *NullableAlwaysAllowedPermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlwaysAllowedPermission(val *AlwaysAllowedPermission) *NullableAlwaysAllowedPermission {
	return &NullableAlwaysAllowedPermission{value: val, isSet: true}
}

func (v NullableAlwaysAllowedPermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlwaysAllowedPermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


