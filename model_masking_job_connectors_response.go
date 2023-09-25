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

// checks if the MaskingJobConnectorsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MaskingJobConnectorsResponse{}

// MaskingJobConnectorsResponse Connector(s) for a masking job.
type MaskingJobConnectorsResponse struct {
	Connector *Connector `json:"connector,omitempty"`
	OnTheFlyConnector *Connector `json:"on_the_fly_connector,omitempty"`
}

// NewMaskingJobConnectorsResponse instantiates a new MaskingJobConnectorsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaskingJobConnectorsResponse() *MaskingJobConnectorsResponse {
	this := MaskingJobConnectorsResponse{}
	return &this
}

// NewMaskingJobConnectorsResponseWithDefaults instantiates a new MaskingJobConnectorsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaskingJobConnectorsResponseWithDefaults() *MaskingJobConnectorsResponse {
	this := MaskingJobConnectorsResponse{}
	return &this
}

// GetConnector returns the Connector field value if set, zero value otherwise.
func (o *MaskingJobConnectorsResponse) GetConnector() Connector {
	if o == nil || IsNil(o.Connector) {
		var ret Connector
		return ret
	}
	return *o.Connector
}

// GetConnectorOk returns a tuple with the Connector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingJobConnectorsResponse) GetConnectorOk() (*Connector, bool) {
	if o == nil || IsNil(o.Connector) {
		return nil, false
	}
	return o.Connector, true
}

// HasConnector returns a boolean if a field has been set.
func (o *MaskingJobConnectorsResponse) HasConnector() bool {
	if o != nil && !IsNil(o.Connector) {
		return true
	}

	return false
}

// SetConnector gets a reference to the given Connector and assigns it to the Connector field.
func (o *MaskingJobConnectorsResponse) SetConnector(v Connector) {
	o.Connector = &v
}

// GetOnTheFlyConnector returns the OnTheFlyConnector field value if set, zero value otherwise.
func (o *MaskingJobConnectorsResponse) GetOnTheFlyConnector() Connector {
	if o == nil || IsNil(o.OnTheFlyConnector) {
		var ret Connector
		return ret
	}
	return *o.OnTheFlyConnector
}

// GetOnTheFlyConnectorOk returns a tuple with the OnTheFlyConnector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingJobConnectorsResponse) GetOnTheFlyConnectorOk() (*Connector, bool) {
	if o == nil || IsNil(o.OnTheFlyConnector) {
		return nil, false
	}
	return o.OnTheFlyConnector, true
}

// HasOnTheFlyConnector returns a boolean if a field has been set.
func (o *MaskingJobConnectorsResponse) HasOnTheFlyConnector() bool {
	if o != nil && !IsNil(o.OnTheFlyConnector) {
		return true
	}

	return false
}

// SetOnTheFlyConnector gets a reference to the given Connector and assigns it to the OnTheFlyConnector field.
func (o *MaskingJobConnectorsResponse) SetOnTheFlyConnector(v Connector) {
	o.OnTheFlyConnector = &v
}

func (o MaskingJobConnectorsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MaskingJobConnectorsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Connector) {
		toSerialize["connector"] = o.Connector
	}
	if !IsNil(o.OnTheFlyConnector) {
		toSerialize["on_the_fly_connector"] = o.OnTheFlyConnector
	}
	return toSerialize, nil
}

type NullableMaskingJobConnectorsResponse struct {
	value *MaskingJobConnectorsResponse
	isSet bool
}

func (v NullableMaskingJobConnectorsResponse) Get() *MaskingJobConnectorsResponse {
	return v.value
}

func (v *NullableMaskingJobConnectorsResponse) Set(val *MaskingJobConnectorsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMaskingJobConnectorsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMaskingJobConnectorsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaskingJobConnectorsResponse(val *MaskingJobConnectorsResponse) *NullableMaskingJobConnectorsResponse {
	return &NullableMaskingJobConnectorsResponse{value: val, isSet: true}
}

func (v NullableMaskingJobConnectorsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaskingJobConnectorsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


