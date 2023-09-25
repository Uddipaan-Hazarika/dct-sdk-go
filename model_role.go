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

// checks if the Role type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Role{}

// Role A role is a named collection of permissions on DCT objects.
type Role struct {
	// The Role name.
	Name string `json:"name"`
	// Role description.
	Description *string `json:"description,omitempty"`
	// The list of permissions granted by this role.
	PermissionObjects []PermissionObject `json:"permission_objects"`
	Tags []Tag `json:"tags,omitempty"`
	// The Role ID.
	Id *string `json:"id,omitempty"`
	// System role are pre defined roles. System roles cannot be modified.
	SystemRole *bool `json:"system_role,omitempty"`
}

// NewRole instantiates a new Role object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRole(name string, permissionObjects []PermissionObject) *Role {
	this := Role{}
	this.Name = name
	this.PermissionObjects = permissionObjects
	return &this
}

// NewRoleWithDefaults instantiates a new Role object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleWithDefaults() *Role {
	this := Role{}
	return &this
}

// GetName returns the Name field value
func (o *Role) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Role) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Role) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Role) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Role) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Role) SetDescription(v string) {
	o.Description = &v
}

// GetPermissionObjects returns the PermissionObjects field value
func (o *Role) GetPermissionObjects() []PermissionObject {
	if o == nil {
		var ret []PermissionObject
		return ret
	}

	return o.PermissionObjects
}

// GetPermissionObjectsOk returns a tuple with the PermissionObjects field value
// and a boolean to check if the value has been set.
func (o *Role) GetPermissionObjectsOk() ([]PermissionObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.PermissionObjects, true
}

// SetPermissionObjects sets field value
func (o *Role) SetPermissionObjects(v []PermissionObject) {
	o.PermissionObjects = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Role) GetTags() []Tag {
	if o == nil || IsNil(o.Tags) {
		var ret []Tag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetTagsOk() ([]Tag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Role) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []Tag and assigns it to the Tags field.
func (o *Role) SetTags(v []Tag) {
	o.Tags = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Role) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Role) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Role) SetId(v string) {
	o.Id = &v
}

// GetSystemRole returns the SystemRole field value if set, zero value otherwise.
func (o *Role) GetSystemRole() bool {
	if o == nil || IsNil(o.SystemRole) {
		var ret bool
		return ret
	}
	return *o.SystemRole
}

// GetSystemRoleOk returns a tuple with the SystemRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetSystemRoleOk() (*bool, bool) {
	if o == nil || IsNil(o.SystemRole) {
		return nil, false
	}
	return o.SystemRole, true
}

// HasSystemRole returns a boolean if a field has been set.
func (o *Role) HasSystemRole() bool {
	if o != nil && !IsNil(o.SystemRole) {
		return true
	}

	return false
}

// SetSystemRole gets a reference to the given bool and assigns it to the SystemRole field.
func (o *Role) SetSystemRole(v bool) {
	o.SystemRole = &v
}

func (o Role) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Role) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["permission_objects"] = o.PermissionObjects
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.SystemRole) {
		toSerialize["system_role"] = o.SystemRole
	}
	return toSerialize, nil
}

type NullableRole struct {
	value *Role
	isSet bool
}

func (v NullableRole) Get() *Role {
	return v.value
}

func (v *NullableRole) Set(val *Role) {
	v.value = val
	v.isSet = true
}

func (v NullableRole) IsSet() bool {
	return v.isSet
}

func (v *NullableRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRole(val *Role) *NullableRole {
	return &NullableRole{value: val, isSet: true}
}

func (v NullableRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

