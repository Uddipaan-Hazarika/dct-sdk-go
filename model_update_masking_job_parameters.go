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

// checks if the UpdateMaskingJobParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateMaskingJobParameters{}

// UpdateMaskingJobParameters Parameters to update a MaskingJob.
type UpdateMaskingJobParameters struct {
	// The name of the MaskingJob.
	Name *string `json:"name,omitempty"`
	// The username of the Connector used by the MaskingJob.
	ConnectorUsername *string `json:"connector_username,omitempty"`
	// The password of the Connector used by the MaskingJob.
	ConnectorPassword *string `json:"connector_password,omitempty"`
	// The username of the source Connector used by the on-the-fly MaskingJob.
	OnTheFlySourceConnectorUsername *string `json:"on_the_fly_source_connector_username,omitempty"`
	// The password of the source Connector used by the on-the-fly MaskingJob.
	OnTheFlySourceConnectorPassword *string `json:"on_the_fly_source_connector_password,omitempty"`
}

// NewUpdateMaskingJobParameters instantiates a new UpdateMaskingJobParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMaskingJobParameters() *UpdateMaskingJobParameters {
	this := UpdateMaskingJobParameters{}
	return &this
}

// NewUpdateMaskingJobParametersWithDefaults instantiates a new UpdateMaskingJobParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMaskingJobParametersWithDefaults() *UpdateMaskingJobParameters {
	this := UpdateMaskingJobParameters{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateMaskingJobParameters) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMaskingJobParameters) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateMaskingJobParameters) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateMaskingJobParameters) SetName(v string) {
	o.Name = &v
}

// GetConnectorUsername returns the ConnectorUsername field value if set, zero value otherwise.
func (o *UpdateMaskingJobParameters) GetConnectorUsername() string {
	if o == nil || IsNil(o.ConnectorUsername) {
		var ret string
		return ret
	}
	return *o.ConnectorUsername
}

// GetConnectorUsernameOk returns a tuple with the ConnectorUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMaskingJobParameters) GetConnectorUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectorUsername) {
		return nil, false
	}
	return o.ConnectorUsername, true
}

// HasConnectorUsername returns a boolean if a field has been set.
func (o *UpdateMaskingJobParameters) HasConnectorUsername() bool {
	if o != nil && !IsNil(o.ConnectorUsername) {
		return true
	}

	return false
}

// SetConnectorUsername gets a reference to the given string and assigns it to the ConnectorUsername field.
func (o *UpdateMaskingJobParameters) SetConnectorUsername(v string) {
	o.ConnectorUsername = &v
}

// GetConnectorPassword returns the ConnectorPassword field value if set, zero value otherwise.
func (o *UpdateMaskingJobParameters) GetConnectorPassword() string {
	if o == nil || IsNil(o.ConnectorPassword) {
		var ret string
		return ret
	}
	return *o.ConnectorPassword
}

// GetConnectorPasswordOk returns a tuple with the ConnectorPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMaskingJobParameters) GetConnectorPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectorPassword) {
		return nil, false
	}
	return o.ConnectorPassword, true
}

// HasConnectorPassword returns a boolean if a field has been set.
func (o *UpdateMaskingJobParameters) HasConnectorPassword() bool {
	if o != nil && !IsNil(o.ConnectorPassword) {
		return true
	}

	return false
}

// SetConnectorPassword gets a reference to the given string and assigns it to the ConnectorPassword field.
func (o *UpdateMaskingJobParameters) SetConnectorPassword(v string) {
	o.ConnectorPassword = &v
}

// GetOnTheFlySourceConnectorUsername returns the OnTheFlySourceConnectorUsername field value if set, zero value otherwise.
func (o *UpdateMaskingJobParameters) GetOnTheFlySourceConnectorUsername() string {
	if o == nil || IsNil(o.OnTheFlySourceConnectorUsername) {
		var ret string
		return ret
	}
	return *o.OnTheFlySourceConnectorUsername
}

// GetOnTheFlySourceConnectorUsernameOk returns a tuple with the OnTheFlySourceConnectorUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMaskingJobParameters) GetOnTheFlySourceConnectorUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.OnTheFlySourceConnectorUsername) {
		return nil, false
	}
	return o.OnTheFlySourceConnectorUsername, true
}

// HasOnTheFlySourceConnectorUsername returns a boolean if a field has been set.
func (o *UpdateMaskingJobParameters) HasOnTheFlySourceConnectorUsername() bool {
	if o != nil && !IsNil(o.OnTheFlySourceConnectorUsername) {
		return true
	}

	return false
}

// SetOnTheFlySourceConnectorUsername gets a reference to the given string and assigns it to the OnTheFlySourceConnectorUsername field.
func (o *UpdateMaskingJobParameters) SetOnTheFlySourceConnectorUsername(v string) {
	o.OnTheFlySourceConnectorUsername = &v
}

// GetOnTheFlySourceConnectorPassword returns the OnTheFlySourceConnectorPassword field value if set, zero value otherwise.
func (o *UpdateMaskingJobParameters) GetOnTheFlySourceConnectorPassword() string {
	if o == nil || IsNil(o.OnTheFlySourceConnectorPassword) {
		var ret string
		return ret
	}
	return *o.OnTheFlySourceConnectorPassword
}

// GetOnTheFlySourceConnectorPasswordOk returns a tuple with the OnTheFlySourceConnectorPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMaskingJobParameters) GetOnTheFlySourceConnectorPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.OnTheFlySourceConnectorPassword) {
		return nil, false
	}
	return o.OnTheFlySourceConnectorPassword, true
}

// HasOnTheFlySourceConnectorPassword returns a boolean if a field has been set.
func (o *UpdateMaskingJobParameters) HasOnTheFlySourceConnectorPassword() bool {
	if o != nil && !IsNil(o.OnTheFlySourceConnectorPassword) {
		return true
	}

	return false
}

// SetOnTheFlySourceConnectorPassword gets a reference to the given string and assigns it to the OnTheFlySourceConnectorPassword field.
func (o *UpdateMaskingJobParameters) SetOnTheFlySourceConnectorPassword(v string) {
	o.OnTheFlySourceConnectorPassword = &v
}

func (o UpdateMaskingJobParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateMaskingJobParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ConnectorUsername) {
		toSerialize["connector_username"] = o.ConnectorUsername
	}
	if !IsNil(o.ConnectorPassword) {
		toSerialize["connector_password"] = o.ConnectorPassword
	}
	if !IsNil(o.OnTheFlySourceConnectorUsername) {
		toSerialize["on_the_fly_source_connector_username"] = o.OnTheFlySourceConnectorUsername
	}
	if !IsNil(o.OnTheFlySourceConnectorPassword) {
		toSerialize["on_the_fly_source_connector_password"] = o.OnTheFlySourceConnectorPassword
	}
	return toSerialize, nil
}

type NullableUpdateMaskingJobParameters struct {
	value *UpdateMaskingJobParameters
	isSet bool
}

func (v NullableUpdateMaskingJobParameters) Get() *UpdateMaskingJobParameters {
	return v.value
}

func (v *NullableUpdateMaskingJobParameters) Set(val *UpdateMaskingJobParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMaskingJobParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMaskingJobParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMaskingJobParameters(val *UpdateMaskingJobParameters) *NullableUpdateMaskingJobParameters {
	return &NullableUpdateMaskingJobParameters{value: val, isSet: true}
}

func (v NullableUpdateMaskingJobParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMaskingJobParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


