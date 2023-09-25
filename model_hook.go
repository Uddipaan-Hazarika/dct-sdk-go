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

// checks if the Hook type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Hook{}

// Hook struct for Hook
type Hook struct {
	Name *string `json:"name,omitempty"`
	Command string `json:"command"`
	Shell *string `json:"shell,omitempty"`
	ElementId *string `json:"element_id,omitempty"`
	HasCredentials *bool `json:"has_credentials,omitempty"`
}

// NewHook instantiates a new Hook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHook(command string) *Hook {
	this := Hook{}
	this.Command = command
	return &this
}

// NewHookWithDefaults instantiates a new Hook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHookWithDefaults() *Hook {
	this := Hook{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Hook) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hook) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Hook) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Hook) SetName(v string) {
	o.Name = &v
}

// GetCommand returns the Command field value
func (o *Hook) GetCommand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Command
}

// GetCommandOk returns a tuple with the Command field value
// and a boolean to check if the value has been set.
func (o *Hook) GetCommandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Command, true
}

// SetCommand sets field value
func (o *Hook) SetCommand(v string) {
	o.Command = v
}

// GetShell returns the Shell field value if set, zero value otherwise.
func (o *Hook) GetShell() string {
	if o == nil || IsNil(o.Shell) {
		var ret string
		return ret
	}
	return *o.Shell
}

// GetShellOk returns a tuple with the Shell field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hook) GetShellOk() (*string, bool) {
	if o == nil || IsNil(o.Shell) {
		return nil, false
	}
	return o.Shell, true
}

// HasShell returns a boolean if a field has been set.
func (o *Hook) HasShell() bool {
	if o != nil && !IsNil(o.Shell) {
		return true
	}

	return false
}

// SetShell gets a reference to the given string and assigns it to the Shell field.
func (o *Hook) SetShell(v string) {
	o.Shell = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *Hook) GetElementId() string {
	if o == nil || IsNil(o.ElementId) {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hook) GetElementIdOk() (*string, bool) {
	if o == nil || IsNil(o.ElementId) {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *Hook) HasElementId() bool {
	if o != nil && !IsNil(o.ElementId) {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *Hook) SetElementId(v string) {
	o.ElementId = &v
}

// GetHasCredentials returns the HasCredentials field value if set, zero value otherwise.
func (o *Hook) GetHasCredentials() bool {
	if o == nil || IsNil(o.HasCredentials) {
		var ret bool
		return ret
	}
	return *o.HasCredentials
}

// GetHasCredentialsOk returns a tuple with the HasCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hook) GetHasCredentialsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasCredentials) {
		return nil, false
	}
	return o.HasCredentials, true
}

// HasHasCredentials returns a boolean if a field has been set.
func (o *Hook) HasHasCredentials() bool {
	if o != nil && !IsNil(o.HasCredentials) {
		return true
	}

	return false
}

// SetHasCredentials gets a reference to the given bool and assigns it to the HasCredentials field.
func (o *Hook) SetHasCredentials(v bool) {
	o.HasCredentials = &v
}

func (o Hook) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Hook) ToMap() (map[string]interface{}, error) {
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

type NullableHook struct {
	value *Hook
	isSet bool
}

func (v NullableHook) Get() *Hook {
	return v.value
}

func (v *NullableHook) Set(val *Hook) {
	v.value = val
	v.isSet = true
}

func (v NullableHook) IsSet() bool {
	return v.isSet
}

func (v *NullableHook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHook(val *Hook) *NullableHook {
	return &NullableHook{value: val, isSet: true}
}

func (v NullableHook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


