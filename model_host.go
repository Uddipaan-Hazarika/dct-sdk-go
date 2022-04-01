/*
Delphix API Gateway

Delphix API Gateway API

API version: 1.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// Host A physical/virtual server.
type Host struct {
	// The hostname or IP address of this host.
	Hostname *string `json:"hostname,omitempty"`
	// The name of the OS on this host.
	OsName *string `json:"os_name,omitempty"`
	// The version of the OS on this host.
	OsVersion *string `json:"os_version,omitempty"`
	// The total amount of memory on this host in bytes.
	MemorySize *int64 `json:"memory_size,omitempty"`
}

// NewHost instantiates a new Host object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHost() *Host {
	this := Host{}
	return &this
}

// NewHostWithDefaults instantiates a new Host object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostWithDefaults() *Host {
	this := Host{}
	return &this
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *Host) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Host) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *Host) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *Host) SetHostname(v string) {
	o.Hostname = &v
}

// GetOsName returns the OsName field value if set, zero value otherwise.
func (o *Host) GetOsName() string {
	if o == nil || o.OsName == nil {
		var ret string
		return ret
	}
	return *o.OsName
}

// GetOsNameOk returns a tuple with the OsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Host) GetOsNameOk() (*string, bool) {
	if o == nil || o.OsName == nil {
		return nil, false
	}
	return o.OsName, true
}

// HasOsName returns a boolean if a field has been set.
func (o *Host) HasOsName() bool {
	if o != nil && o.OsName != nil {
		return true
	}

	return false
}

// SetOsName gets a reference to the given string and assigns it to the OsName field.
func (o *Host) SetOsName(v string) {
	o.OsName = &v
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *Host) GetOsVersion() string {
	if o == nil || o.OsVersion == nil {
		var ret string
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Host) GetOsVersionOk() (*string, bool) {
	if o == nil || o.OsVersion == nil {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *Host) HasOsVersion() bool {
	if o != nil && o.OsVersion != nil {
		return true
	}

	return false
}

// SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.
func (o *Host) SetOsVersion(v string) {
	o.OsVersion = &v
}

// GetMemorySize returns the MemorySize field value if set, zero value otherwise.
func (o *Host) GetMemorySize() int64 {
	if o == nil || o.MemorySize == nil {
		var ret int64
		return ret
	}
	return *o.MemorySize
}

// GetMemorySizeOk returns a tuple with the MemorySize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Host) GetMemorySizeOk() (*int64, bool) {
	if o == nil || o.MemorySize == nil {
		return nil, false
	}
	return o.MemorySize, true
}

// HasMemorySize returns a boolean if a field has been set.
func (o *Host) HasMemorySize() bool {
	if o != nil && o.MemorySize != nil {
		return true
	}

	return false
}

// SetMemorySize gets a reference to the given int64 and assigns it to the MemorySize field.
func (o *Host) SetMemorySize(v int64) {
	o.MemorySize = &v
}

func (o Host) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	if o.OsName != nil {
		toSerialize["os_name"] = o.OsName
	}
	if o.OsVersion != nil {
		toSerialize["os_version"] = o.OsVersion
	}
	if o.MemorySize != nil {
		toSerialize["memory_size"] = o.MemorySize
	}
	return json.Marshal(toSerialize)
}

type NullableHost struct {
	value *Host
	isSet bool
}

func (v NullableHost) Get() *Host {
	return v.value
}

func (v *NullableHost) Set(val *Host) {
	v.value = val
	v.isSet = true
}

func (v NullableHost) IsSet() bool {
	return v.isSet
}

func (v *NullableHost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHost(val *Host) *NullableHost {
	return &NullableHost{value: val, isSet: true}
}

func (v NullableHost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


