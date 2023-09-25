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

// checks if the CreateDatabaseTemplateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDatabaseTemplateResponse{}

// CreateDatabaseTemplateResponse struct for CreateDatabaseTemplateResponse
type CreateDatabaseTemplateResponse struct {
	DatabaseTemplate *DatabaseTemplate `json:"database_template,omitempty"`
	Job *Job `json:"job,omitempty"`
}

// NewCreateDatabaseTemplateResponse instantiates a new CreateDatabaseTemplateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDatabaseTemplateResponse() *CreateDatabaseTemplateResponse {
	this := CreateDatabaseTemplateResponse{}
	return &this
}

// NewCreateDatabaseTemplateResponseWithDefaults instantiates a new CreateDatabaseTemplateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDatabaseTemplateResponseWithDefaults() *CreateDatabaseTemplateResponse {
	this := CreateDatabaseTemplateResponse{}
	return &this
}

// GetDatabaseTemplate returns the DatabaseTemplate field value if set, zero value otherwise.
func (o *CreateDatabaseTemplateResponse) GetDatabaseTemplate() DatabaseTemplate {
	if o == nil || IsNil(o.DatabaseTemplate) {
		var ret DatabaseTemplate
		return ret
	}
	return *o.DatabaseTemplate
}

// GetDatabaseTemplateOk returns a tuple with the DatabaseTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatabaseTemplateResponse) GetDatabaseTemplateOk() (*DatabaseTemplate, bool) {
	if o == nil || IsNil(o.DatabaseTemplate) {
		return nil, false
	}
	return o.DatabaseTemplate, true
}

// HasDatabaseTemplate returns a boolean if a field has been set.
func (o *CreateDatabaseTemplateResponse) HasDatabaseTemplate() bool {
	if o != nil && !IsNil(o.DatabaseTemplate) {
		return true
	}

	return false
}

// SetDatabaseTemplate gets a reference to the given DatabaseTemplate and assigns it to the DatabaseTemplate field.
func (o *CreateDatabaseTemplateResponse) SetDatabaseTemplate(v DatabaseTemplate) {
	o.DatabaseTemplate = &v
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *CreateDatabaseTemplateResponse) GetJob() Job {
	if o == nil || IsNil(o.Job) {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatabaseTemplateResponse) GetJobOk() (*Job, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *CreateDatabaseTemplateResponse) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *CreateDatabaseTemplateResponse) SetJob(v Job) {
	o.Job = &v
}

func (o CreateDatabaseTemplateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDatabaseTemplateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DatabaseTemplate) {
		toSerialize["database_template"] = o.DatabaseTemplate
	}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableCreateDatabaseTemplateResponse struct {
	value *CreateDatabaseTemplateResponse
	isSet bool
}

func (v NullableCreateDatabaseTemplateResponse) Get() *CreateDatabaseTemplateResponse {
	return v.value
}

func (v *NullableCreateDatabaseTemplateResponse) Set(val *CreateDatabaseTemplateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDatabaseTemplateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDatabaseTemplateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDatabaseTemplateResponse(val *CreateDatabaseTemplateResponse) *NullableCreateDatabaseTemplateResponse {
	return &NullableCreateDatabaseTemplateResponse{value: val, isSet: true}
}

func (v NullableCreateDatabaseTemplateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDatabaseTemplateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


