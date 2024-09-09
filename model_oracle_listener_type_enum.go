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
	"fmt"
)

// OracleListenerTypeEnum the model 'OracleListenerTypeEnum'
type OracleListenerTypeEnum string

// List of OracleListenerTypeEnum
const (
	ORACLELISTENERTYPEENUM_NODE OracleListenerTypeEnum = "NODE"
	ORACLELISTENERTYPEENUM_SCAN OracleListenerTypeEnum = "SCAN"
)

// All allowed values of OracleListenerTypeEnum enum
var AllowedOracleListenerTypeEnumEnumValues = []OracleListenerTypeEnum{
	"NODE",
	"SCAN",
}

func (v *OracleListenerTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OracleListenerTypeEnum(value)
	for _, existing := range AllowedOracleListenerTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OracleListenerTypeEnum", value)
}

// NewOracleListenerTypeEnumFromValue returns a pointer to a valid OracleListenerTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOracleListenerTypeEnumFromValue(v string) (*OracleListenerTypeEnum, error) {
	ev := OracleListenerTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OracleListenerTypeEnum: valid values are %v", v, AllowedOracleListenerTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OracleListenerTypeEnum) IsValid() bool {
	for _, existing := range AllowedOracleListenerTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OracleListenerTypeEnum value
func (v OracleListenerTypeEnum) Ptr() *OracleListenerTypeEnum {
	return &v
}

type NullableOracleListenerTypeEnum struct {
	value *OracleListenerTypeEnum
	isSet bool
}

func (v NullableOracleListenerTypeEnum) Get() *OracleListenerTypeEnum {
	return v.value
}

func (v *NullableOracleListenerTypeEnum) Set(val *OracleListenerTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableOracleListenerTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableOracleListenerTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOracleListenerTypeEnum(val *OracleListenerTypeEnum) *NullableOracleListenerTypeEnum {
	return &NullableOracleListenerTypeEnum{value: val, isSet: true}
}

func (v NullableOracleListenerTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOracleListenerTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
