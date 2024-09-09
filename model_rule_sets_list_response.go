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
)

// checks if the RuleSetsListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleSetsListResponse{}

// RuleSetsListResponse struct for RuleSetsListResponse
type RuleSetsListResponse struct {
	Items []RuleSet `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewRuleSetsListResponse instantiates a new RuleSetsListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleSetsListResponse() *RuleSetsListResponse {
	this := RuleSetsListResponse{}
	return &this
}

// NewRuleSetsListResponseWithDefaults instantiates a new RuleSetsListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleSetsListResponseWithDefaults() *RuleSetsListResponse {
	this := RuleSetsListResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *RuleSetsListResponse) GetItems() []RuleSet {
	if o == nil || IsNil(o.Items) {
		var ret []RuleSet
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleSetsListResponse) GetItemsOk() ([]RuleSet, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *RuleSetsListResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []RuleSet and assigns it to the Items field.
func (o *RuleSetsListResponse) SetItems(v []RuleSet) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *RuleSetsListResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleSetsListResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *RuleSetsListResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *RuleSetsListResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o RuleSetsListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuleSetsListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return toSerialize, nil
}

type NullableRuleSetsListResponse struct {
	value *RuleSetsListResponse
	isSet bool
}

func (v NullableRuleSetsListResponse) Get() *RuleSetsListResponse {
	return v.value
}

func (v *NullableRuleSetsListResponse) Set(val *RuleSetsListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleSetsListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleSetsListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleSetsListResponse(val *RuleSetsListResponse) *NullableRuleSetsListResponse {
	return &NullableRuleSetsListResponse{value: val, isSet: true}
}

func (v NullableRuleSetsListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleSetsListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

