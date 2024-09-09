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
	"bytes"
	"fmt"
)

// checks if the PhoneHomeBundleProductHistory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PhoneHomeBundleProductHistory{}

// PhoneHomeBundleProductHistory The product history data included in the phonehome bundle's product information.
type PhoneHomeBundleProductHistory struct {
	// The product version.
	Version string `json:"version"`
	// The UTC time when the version was installed (ISO 8601 format)
	InstalledOn string `json:"installed_on"`
}

type _PhoneHomeBundleProductHistory PhoneHomeBundleProductHistory

// NewPhoneHomeBundleProductHistory instantiates a new PhoneHomeBundleProductHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhoneHomeBundleProductHistory(version string, installedOn string) *PhoneHomeBundleProductHistory {
	this := PhoneHomeBundleProductHistory{}
	this.Version = version
	this.InstalledOn = installedOn
	return &this
}

// NewPhoneHomeBundleProductHistoryWithDefaults instantiates a new PhoneHomeBundleProductHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhoneHomeBundleProductHistoryWithDefaults() *PhoneHomeBundleProductHistory {
	this := PhoneHomeBundleProductHistory{}
	return &this
}

// GetVersion returns the Version field value
func (o *PhoneHomeBundleProductHistory) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *PhoneHomeBundleProductHistory) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *PhoneHomeBundleProductHistory) SetVersion(v string) {
	o.Version = v
}

// GetInstalledOn returns the InstalledOn field value
func (o *PhoneHomeBundleProductHistory) GetInstalledOn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstalledOn
}

// GetInstalledOnOk returns a tuple with the InstalledOn field value
// and a boolean to check if the value has been set.
func (o *PhoneHomeBundleProductHistory) GetInstalledOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstalledOn, true
}

// SetInstalledOn sets field value
func (o *PhoneHomeBundleProductHistory) SetInstalledOn(v string) {
	o.InstalledOn = v
}

func (o PhoneHomeBundleProductHistory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PhoneHomeBundleProductHistory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["version"] = o.Version
	toSerialize["installed_on"] = o.InstalledOn
	return toSerialize, nil
}

func (o *PhoneHomeBundleProductHistory) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"version",
		"installed_on",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPhoneHomeBundleProductHistory := _PhoneHomeBundleProductHistory{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPhoneHomeBundleProductHistory)

	if err != nil {
		return err
	}

	*o = PhoneHomeBundleProductHistory(varPhoneHomeBundleProductHistory)

	return err
}

type NullablePhoneHomeBundleProductHistory struct {
	value *PhoneHomeBundleProductHistory
	isSet bool
}

func (v NullablePhoneHomeBundleProductHistory) Get() *PhoneHomeBundleProductHistory {
	return v.value
}

func (v *NullablePhoneHomeBundleProductHistory) Set(val *PhoneHomeBundleProductHistory) {
	v.value = val
	v.isSet = true
}

func (v NullablePhoneHomeBundleProductHistory) IsSet() bool {
	return v.isSet
}

func (v *NullablePhoneHomeBundleProductHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhoneHomeBundleProductHistory(val *PhoneHomeBundleProductHistory) *NullablePhoneHomeBundleProductHistory {
	return &NullablePhoneHomeBundleProductHistory{value: val, isSet: true}
}

func (v NullablePhoneHomeBundleProductHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhoneHomeBundleProductHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

