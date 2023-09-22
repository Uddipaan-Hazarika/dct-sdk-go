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

// checks if the AccessGroupAccountIdsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessGroupAccountIdsRequest{}

// AccessGroupAccountIdsRequest struct for AccessGroupAccountIdsRequest
type AccessGroupAccountIdsRequest struct {
	AccountIds []int64 `json:"account_ids"`
}

// NewAccessGroupAccountIdsRequest instantiates a new AccessGroupAccountIdsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessGroupAccountIdsRequest(accountIds []int64) *AccessGroupAccountIdsRequest {
	this := AccessGroupAccountIdsRequest{}
	this.AccountIds = accountIds
	return &this
}

// NewAccessGroupAccountIdsRequestWithDefaults instantiates a new AccessGroupAccountIdsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessGroupAccountIdsRequestWithDefaults() *AccessGroupAccountIdsRequest {
	this := AccessGroupAccountIdsRequest{}
	return &this
}

// GetAccountIds returns the AccountIds field value
func (o *AccessGroupAccountIdsRequest) GetAccountIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.AccountIds
}

// GetAccountIdsOk returns a tuple with the AccountIds field value
// and a boolean to check if the value has been set.
func (o *AccessGroupAccountIdsRequest) GetAccountIdsOk() ([]int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountIds, true
}

// SetAccountIds sets field value
func (o *AccessGroupAccountIdsRequest) SetAccountIds(v []int64) {
	o.AccountIds = v
}

func (o AccessGroupAccountIdsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessGroupAccountIdsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_ids"] = o.AccountIds
	return toSerialize, nil
}

type NullableAccessGroupAccountIdsRequest struct {
	value *AccessGroupAccountIdsRequest
	isSet bool
}

func (v NullableAccessGroupAccountIdsRequest) Get() *AccessGroupAccountIdsRequest {
	return v.value
}

func (v *NullableAccessGroupAccountIdsRequest) Set(val *AccessGroupAccountIdsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessGroupAccountIdsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessGroupAccountIdsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessGroupAccountIdsRequest(val *AccessGroupAccountIdsRequest) *NullableAccessGroupAccountIdsRequest {
	return &NullableAccessGroupAccountIdsRequest{value: val, isSet: true}
}

func (v NullableAccessGroupAccountIdsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessGroupAccountIdsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


