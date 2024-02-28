/*
Delphix DCT API

Delphix DCT API

API version: 3.9.0
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
	// The username of the Connector used by the MaskingJob. (Standard Job only).
	ConnectorUsername *string `json:"connector_username,omitempty"`
	// The password of the Connector used by the MaskingJob. (Standard Job only).
	ConnectorPassword *string `json:"connector_password,omitempty"`
	// The username of the source Connector used by the on-the-fly MaskingJob. (Standard Job only).
	OnTheFlySourceConnectorUsername *string `json:"on_the_fly_source_connector_username,omitempty"`
	// The password of the source Connector used by the on-the-fly MaskingJob. (Standard Job only).
	OnTheFlySourceConnectorPassword *string `json:"on_the_fly_source_connector_password,omitempty"`
	// Description of the Job (Hyperscale Job only).
	Description *string `json:"description,omitempty"`
	// Dataset of the Hyperscale Job (Hyperscale Job only).
	DatasetId *string `json:"dataset_id,omitempty"`
	// Defines whether execution data will be stored after execution is complete (Hyperscale Job only).
	RetainExecutionData *string `json:"retain_execution_data,omitempty"`
	// Maximum memory to be allocated for each Masking job (Hyperscale Job only).
	MaxMemory *int32 `json:"max_memory,omitempty"`
	// Minimum memory to be allocated for each Masking job (Hyperscale Job only).
	MinMemory *int32 `json:"min_memory,omitempty"`
	// Feedback Size for each Masking job (Hyperscale Job only).
	FeedbackSize *int32 `json:"feedback_size,omitempty"`
	// Stream Row Limit for each Masking job (Hyperscale Job only).
	StreamRowLimit *int32 `json:"stream_row_limit,omitempty"`
	// Number of input streams to be configured for Masking Job (Hyperscale Job only).
	NumInputStreams *int32 `json:"num_input_streams,omitempty"`
	// Maximum number of parallel connection that the Hyperscale instance can have with the source datasource (Hyperscale Job only).
	MaxConcurrentSourceConnections *int32 `json:"max_concurrent_source_connections,omitempty"`
	// Maximum number of parallel connection that the Hyperscale instance can have with the target datasource (Hyperscale Job only).
	MaxConcurrentTargetConnections *int32 `json:"max_concurrent_target_connections,omitempty"`
	// The degree of parallelism (DOP) per Oracle job to recreate the index in the post-load process (Hyperscale Job only).
	ParallelismDegree *int32 `json:"parallelism_degree,omitempty"`
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

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateMaskingJobParameters) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMaskingJobParameters) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateMaskingJobParameters) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateMaskingJobParameters) SetDescription(v string) {
	o.Description = &v
}

// GetDatasetId returns the DatasetId field value if set, zero value otherwise.
func (o *UpdateMaskingJobParameters) GetDatasetId() string {
	if o == nil || IsNil(o.DatasetId) {
		var ret string
		return ret
	}
	return *o.DatasetId
}

// GetDatasetIdOk returns a tuple with the DatasetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMaskingJobParameters) GetDatasetIdOk() (*string, bool) {
	if o == nil || IsNil(o.DatasetId) {
		return nil, false
	}
	return o.DatasetId, true
}

// HasDatasetId returns a boolean if a field has been set.
func (o *UpdateMaskingJobParameters) HasDatasetId() bool {
	if o != nil && !IsNil(o.DatasetId) {
		return true
	}

	return false
}

// SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.
func (o *UpdateMaskingJobParameters) SetDatasetId(v string) {
	o.DatasetId = &v
}

// GetRetainExecutionData returns the RetainExecutionData field value if set, zero value otherwise.
func (o *UpdateMaskingJobParameters) GetRetainExecutionData() string {
	if o == nil || IsNil(o.RetainExecutionData) {
		var ret string
		return ret
	}
	return *o.RetainExecutionData
}

// GetRetainExecutionDataOk returns a tuple with the RetainExecutionData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMaskingJobParameters) GetRetainExecutionDataOk() (*string, bool) {
	if o == nil || IsNil(o.RetainExecutionData) {
		return nil, false
	}
	return o.RetainExecutionData, true
}

// HasRetainExecutionData returns a boolean if a field has been set.
func (o *UpdateMaskingJobParameters) HasRetainExecutionData() bool {
	if o != nil && !IsNil(o.RetainExecutionData) {
		return true
	}

	return false
}

// SetRetainExecutionData gets a reference to the given string and assigns it to the RetainExecutionData field.
func (o *UpdateMaskingJobParameters) SetRetainExecutionData(v string) {
	o.RetainExecutionData = &v
}

// GetMaxMemory returns the MaxMemory field value if set, zero value otherwise.
func (o *UpdateMaskingJobParameters) GetMaxMemory() int32 {
	if o == nil || IsNil(o.MaxMemory) {
		var ret int32
		return ret
	}
	return *o.MaxMemory
}

// GetMaxMemoryOk returns a tuple with the MaxMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMaskingJobParameters) GetMaxMemoryOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxMemory) {
		return nil, false
	}
	return o.MaxMemory, true
}

// HasMaxMemory returns a boolean if a field has been set.
func (o *UpdateMaskingJobParameters) HasMaxMemory() bool {
	if o != nil && !IsNil(o.MaxMemory) {
		return true
	}

	return false
}

// SetMaxMemory gets a reference to the given int32 and assigns it to the MaxMemory field.
func (o *UpdateMaskingJobParameters) SetMaxMemory(v int32) {
	o.MaxMemory = &v
}

// GetMinMemory returns the MinMemory field value if set, zero value otherwise.
func (o *UpdateMaskingJobParameters) GetMinMemory() int32 {
	if o == nil || IsNil(o.MinMemory) {
		var ret int32
		return ret
	}
	return *o.MinMemory
}

// GetMinMemoryOk returns a tuple with the MinMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMaskingJobParameters) GetMinMemoryOk() (*int32, bool) {
	if o == nil || IsNil(o.MinMemory) {
		return nil, false
	}
	return o.MinMemory, true
}

// HasMinMemory returns a boolean if a field has been set.
func (o *UpdateMaskingJobParameters) HasMinMemory() bool {
	if o != nil && !IsNil(o.MinMemory) {
		return true
	}

	return false
}

// SetMinMemory gets a reference to the given int32 and assigns it to the MinMemory field.
func (o *UpdateMaskingJobParameters) SetMinMemory(v int32) {
	o.MinMemory = &v
}

// GetFeedbackSize returns the FeedbackSize field value if set, zero value otherwise.
func (o *UpdateMaskingJobParameters) GetFeedbackSize() int32 {
	if o == nil || IsNil(o.FeedbackSize) {
		var ret int32
		return ret
	}
	return *o.FeedbackSize
}

// GetFeedbackSizeOk returns a tuple with the FeedbackSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMaskingJobParameters) GetFeedbackSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.FeedbackSize) {
		return nil, false
	}
	return o.FeedbackSize, true
}

// HasFeedbackSize returns a boolean if a field has been set.
func (o *UpdateMaskingJobParameters) HasFeedbackSize() bool {
	if o != nil && !IsNil(o.FeedbackSize) {
		return true
	}

	return false
}

// SetFeedbackSize gets a reference to the given int32 and assigns it to the FeedbackSize field.
func (o *UpdateMaskingJobParameters) SetFeedbackSize(v int32) {
	o.FeedbackSize = &v
}

// GetStreamRowLimit returns the StreamRowLimit field value if set, zero value otherwise.
func (o *UpdateMaskingJobParameters) GetStreamRowLimit() int32 {
	if o == nil || IsNil(o.StreamRowLimit) {
		var ret int32
		return ret
	}
	return *o.StreamRowLimit
}

// GetStreamRowLimitOk returns a tuple with the StreamRowLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMaskingJobParameters) GetStreamRowLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.StreamRowLimit) {
		return nil, false
	}
	return o.StreamRowLimit, true
}

// HasStreamRowLimit returns a boolean if a field has been set.
func (o *UpdateMaskingJobParameters) HasStreamRowLimit() bool {
	if o != nil && !IsNil(o.StreamRowLimit) {
		return true
	}

	return false
}

// SetStreamRowLimit gets a reference to the given int32 and assigns it to the StreamRowLimit field.
func (o *UpdateMaskingJobParameters) SetStreamRowLimit(v int32) {
	o.StreamRowLimit = &v
}

// GetNumInputStreams returns the NumInputStreams field value if set, zero value otherwise.
func (o *UpdateMaskingJobParameters) GetNumInputStreams() int32 {
	if o == nil || IsNil(o.NumInputStreams) {
		var ret int32
		return ret
	}
	return *o.NumInputStreams
}

// GetNumInputStreamsOk returns a tuple with the NumInputStreams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMaskingJobParameters) GetNumInputStreamsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumInputStreams) {
		return nil, false
	}
	return o.NumInputStreams, true
}

// HasNumInputStreams returns a boolean if a field has been set.
func (o *UpdateMaskingJobParameters) HasNumInputStreams() bool {
	if o != nil && !IsNil(o.NumInputStreams) {
		return true
	}

	return false
}

// SetNumInputStreams gets a reference to the given int32 and assigns it to the NumInputStreams field.
func (o *UpdateMaskingJobParameters) SetNumInputStreams(v int32) {
	o.NumInputStreams = &v
}

// GetMaxConcurrentSourceConnections returns the MaxConcurrentSourceConnections field value if set, zero value otherwise.
func (o *UpdateMaskingJobParameters) GetMaxConcurrentSourceConnections() int32 {
	if o == nil || IsNil(o.MaxConcurrentSourceConnections) {
		var ret int32
		return ret
	}
	return *o.MaxConcurrentSourceConnections
}

// GetMaxConcurrentSourceConnectionsOk returns a tuple with the MaxConcurrentSourceConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMaskingJobParameters) GetMaxConcurrentSourceConnectionsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxConcurrentSourceConnections) {
		return nil, false
	}
	return o.MaxConcurrentSourceConnections, true
}

// HasMaxConcurrentSourceConnections returns a boolean if a field has been set.
func (o *UpdateMaskingJobParameters) HasMaxConcurrentSourceConnections() bool {
	if o != nil && !IsNil(o.MaxConcurrentSourceConnections) {
		return true
	}

	return false
}

// SetMaxConcurrentSourceConnections gets a reference to the given int32 and assigns it to the MaxConcurrentSourceConnections field.
func (o *UpdateMaskingJobParameters) SetMaxConcurrentSourceConnections(v int32) {
	o.MaxConcurrentSourceConnections = &v
}

// GetMaxConcurrentTargetConnections returns the MaxConcurrentTargetConnections field value if set, zero value otherwise.
func (o *UpdateMaskingJobParameters) GetMaxConcurrentTargetConnections() int32 {
	if o == nil || IsNil(o.MaxConcurrentTargetConnections) {
		var ret int32
		return ret
	}
	return *o.MaxConcurrentTargetConnections
}

// GetMaxConcurrentTargetConnectionsOk returns a tuple with the MaxConcurrentTargetConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMaskingJobParameters) GetMaxConcurrentTargetConnectionsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxConcurrentTargetConnections) {
		return nil, false
	}
	return o.MaxConcurrentTargetConnections, true
}

// HasMaxConcurrentTargetConnections returns a boolean if a field has been set.
func (o *UpdateMaskingJobParameters) HasMaxConcurrentTargetConnections() bool {
	if o != nil && !IsNil(o.MaxConcurrentTargetConnections) {
		return true
	}

	return false
}

// SetMaxConcurrentTargetConnections gets a reference to the given int32 and assigns it to the MaxConcurrentTargetConnections field.
func (o *UpdateMaskingJobParameters) SetMaxConcurrentTargetConnections(v int32) {
	o.MaxConcurrentTargetConnections = &v
}

// GetParallelismDegree returns the ParallelismDegree field value if set, zero value otherwise.
func (o *UpdateMaskingJobParameters) GetParallelismDegree() int32 {
	if o == nil || IsNil(o.ParallelismDegree) {
		var ret int32
		return ret
	}
	return *o.ParallelismDegree
}

// GetParallelismDegreeOk returns a tuple with the ParallelismDegree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMaskingJobParameters) GetParallelismDegreeOk() (*int32, bool) {
	if o == nil || IsNil(o.ParallelismDegree) {
		return nil, false
	}
	return o.ParallelismDegree, true
}

// HasParallelismDegree returns a boolean if a field has been set.
func (o *UpdateMaskingJobParameters) HasParallelismDegree() bool {
	if o != nil && !IsNil(o.ParallelismDegree) {
		return true
	}

	return false
}

// SetParallelismDegree gets a reference to the given int32 and assigns it to the ParallelismDegree field.
func (o *UpdateMaskingJobParameters) SetParallelismDegree(v int32) {
	o.ParallelismDegree = &v
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
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DatasetId) {
		toSerialize["dataset_id"] = o.DatasetId
	}
	if !IsNil(o.RetainExecutionData) {
		toSerialize["retain_execution_data"] = o.RetainExecutionData
	}
	if !IsNil(o.MaxMemory) {
		toSerialize["max_memory"] = o.MaxMemory
	}
	if !IsNil(o.MinMemory) {
		toSerialize["min_memory"] = o.MinMemory
	}
	if !IsNil(o.FeedbackSize) {
		toSerialize["feedback_size"] = o.FeedbackSize
	}
	if !IsNil(o.StreamRowLimit) {
		toSerialize["stream_row_limit"] = o.StreamRowLimit
	}
	if !IsNil(o.NumInputStreams) {
		toSerialize["num_input_streams"] = o.NumInputStreams
	}
	if !IsNil(o.MaxConcurrentSourceConnections) {
		toSerialize["max_concurrent_source_connections"] = o.MaxConcurrentSourceConnections
	}
	if !IsNil(o.MaxConcurrentTargetConnections) {
		toSerialize["max_concurrent_target_connections"] = o.MaxConcurrentTargetConnections
	}
	if !IsNil(o.ParallelismDegree) {
		toSerialize["parallelism_degree"] = o.ParallelismDegree
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


