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

// checks if the DSourceHooksOpsPreSyncInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DSourceHooksOpsPreSyncInner{}

// DSourceHooksOpsPreSyncInner struct for DSourceHooksOpsPreSyncInner
type DSourceHooksOpsPreSyncInner struct {
	Name *string `json:"name,omitempty"`
	Command string `json:"command"`
	Shell *string `json:"shell,omitempty"`
	ElementId *string `json:"element_id,omitempty"`
	HasCredentials *bool `json:"has_credentials,omitempty"`
}

// NewDSourceHooksOpsPreSyncInner instantiates a new DSourceHooksOpsPreSyncInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDSourceHooksOpsPreSyncInner(command string) *DSourceHooksOpsPreSyncInner {
	this := DSourceHooksOpsPreSyncInner{}
	this.Command = command
	return &this
}

// NewDSourceHooksOpsPreSyncInnerWithDefaults instantiates a new DSourceHooksOpsPreSyncInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDSourceHooksOpsPreSyncInnerWithDefaults() *DSourceHooksOpsPreSyncInner {
	this := DSourceHooksOpsPreSyncInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DSourceHooksOpsPreSyncInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DSourceHooksOpsPreSyncInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DSourceHooksOpsPreSyncInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DSourceHooksOpsPreSyncInner) SetName(v string) {
	o.Name = &v
}

// GetCommand returns the Command field value
func (o *DSourceHooksOpsPreSyncInner) GetCommand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Command
}

// GetCommandOk returns a tuple with the Command field value
// and a boolean to check if the value has been set.
func (o *DSourceHooksOpsPreSyncInner) GetCommandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Command, true
}

// SetCommand sets field value
func (o *DSourceHooksOpsPreSyncInner) SetCommand(v string) {
	o.Command = v
}

// GetShell returns the Shell field value if set, zero value otherwise.
func (o *DSourceHooksOpsPreSyncInner) GetShell() string {
	if o == nil || IsNil(o.Shell) {
		var ret string
		return ret
	}
	return *o.Shell
}

// GetShellOk returns a tuple with the Shell field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DSourceHooksOpsPreSyncInner) GetShellOk() (*string, bool) {
	if o == nil || IsNil(o.Shell) {
		return nil, false
	}
	return o.Shell, true
}

// HasShell returns a boolean if a field has been set.
func (o *DSourceHooksOpsPreSyncInner) HasShell() bool {
	if o != nil && !IsNil(o.Shell) {
		return true
	}

	return false
}

// SetShell gets a reference to the given string and assigns it to the Shell field.
func (o *DSourceHooksOpsPreSyncInner) SetShell(v string) {
	o.Shell = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *DSourceHooksOpsPreSyncInner) GetElementId() string {
	if o == nil || IsNil(o.ElementId) {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DSourceHooksOpsPreSyncInner) GetElementIdOk() (*string, bool) {
	if o == nil || IsNil(o.ElementId) {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *DSourceHooksOpsPreSyncInner) HasElementId() bool {
	if o != nil && !IsNil(o.ElementId) {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *DSourceHooksOpsPreSyncInner) SetElementId(v string) {
	o.ElementId = &v
}

// GetHasCredentials returns the HasCredentials field value if set, zero value otherwise.
func (o *DSourceHooksOpsPreSyncInner) GetHasCredentials() bool {
	if o == nil || IsNil(o.HasCredentials) {
		var ret bool
		return ret
	}
	return *o.HasCredentials
}

// GetHasCredentialsOk returns a tuple with the HasCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DSourceHooksOpsPreSyncInner) GetHasCredentialsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasCredentials) {
		return nil, false
	}
	return o.HasCredentials, true
}

// HasHasCredentials returns a boolean if a field has been set.
func (o *DSourceHooksOpsPreSyncInner) HasHasCredentials() bool {
	if o != nil && !IsNil(o.HasCredentials) {
		return true
	}

	return false
}

// SetHasCredentials gets a reference to the given bool and assigns it to the HasCredentials field.
func (o *DSourceHooksOpsPreSyncInner) SetHasCredentials(v bool) {
	o.HasCredentials = &v
}

func (o DSourceHooksOpsPreSyncInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DSourceHooksOpsPreSyncInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["command"] = o.Command
	if !IsNil(o.Shell) {
		toSerialize["shell"] = o.Shell
	}
	if !IsNil(o.ElementId) {
		toSerialize["element_id"] = o.ElementId
	}
	if !IsNil(o.HasCredentials) {
		toSerialize["has_credentials"] = o.HasCredentials
	}
	return toSerialize, nil
}

type NullableDSourceHooksOpsPreSyncInner struct {
	value *DSourceHooksOpsPreSyncInner
	isSet bool
}

func (v NullableDSourceHooksOpsPreSyncInner) Get() *DSourceHooksOpsPreSyncInner {
	return v.value
}

func (v *NullableDSourceHooksOpsPreSyncInner) Set(val *DSourceHooksOpsPreSyncInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDSourceHooksOpsPreSyncInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDSourceHooksOpsPreSyncInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDSourceHooksOpsPreSyncInner(val *DSourceHooksOpsPreSyncInner) *NullableDSourceHooksOpsPreSyncInner {
	return &NullableDSourceHooksOpsPreSyncInner{value: val, isSet: true}
}

func (v NullableDSourceHooksOpsPreSyncInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDSourceHooksOpsPreSyncInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


