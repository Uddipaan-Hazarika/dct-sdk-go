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

// checks if the SMTPConfigValidate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SMTPConfigValidate{}

// SMTPConfigValidate Parameters to validate SMTP Config
type SMTPConfigValidate struct {
	ToAddress string `json:"to_address"`
}

// NewSMTPConfigValidate instantiates a new SMTPConfigValidate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSMTPConfigValidate(toAddress string) *SMTPConfigValidate {
	this := SMTPConfigValidate{}
	this.ToAddress = toAddress
	return &this
}

// NewSMTPConfigValidateWithDefaults instantiates a new SMTPConfigValidate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSMTPConfigValidateWithDefaults() *SMTPConfigValidate {
	this := SMTPConfigValidate{}
	return &this
}

// GetToAddress returns the ToAddress field value
func (o *SMTPConfigValidate) GetToAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToAddress
}

// GetToAddressOk returns a tuple with the ToAddress field value
// and a boolean to check if the value has been set.
func (o *SMTPConfigValidate) GetToAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToAddress, true
}

// SetToAddress sets field value
func (o *SMTPConfigValidate) SetToAddress(v string) {
	o.ToAddress = v
}

func (o SMTPConfigValidate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SMTPConfigValidate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["to_address"] = o.ToAddress
	return toSerialize, nil
}

type NullableSMTPConfigValidate struct {
	value *SMTPConfigValidate
	isSet bool
}

func (v NullableSMTPConfigValidate) Get() *SMTPConfigValidate {
	return v.value
}

func (v *NullableSMTPConfigValidate) Set(val *SMTPConfigValidate) {
	v.value = val
	v.isSet = true
}

func (v NullableSMTPConfigValidate) IsSet() bool {
	return v.isSet
}

func (v *NullableSMTPConfigValidate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSMTPConfigValidate(val *SMTPConfigValidate) *NullableSMTPConfigValidate {
	return &NullableSMTPConfigValidate{value: val, isSet: true}
}

func (v NullableSMTPConfigValidate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSMTPConfigValidate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


