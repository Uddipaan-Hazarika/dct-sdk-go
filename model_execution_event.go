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
	"time"
)

// checks if the ExecutionEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExecutionEvent{}

// ExecutionEvent Events, such as warnings or errors, associated with job executions.
type ExecutionEvent struct {
	// The ExecutionEvent entity ID.
	Id *string `json:"id,omitempty"`
	// The ID of the execution.
	ExecutionId *string `json:"execution_id,omitempty"`
	// The type of execution event.
	EventType *string `json:"event_type,omitempty"`
	// The severity of the execution event.
	Severity *string `json:"severity,omitempty"`
	// The cause of the execution event.
	Cause *string `json:"cause,omitempty"`
	// The number of times the execution event occurred.
	Count *int64 `json:"count,omitempty"`
	// The date and time that this execution event first occurred.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// The name of the column, field, or other object being masked when this event occurred, if applicable.
	MaskedObjectName *string `json:"masked_object_name,omitempty"`
	// The name of the masking algorithm running when this event occurred, if applicable.
	AlgorithmName *string `json:"algorithm_name,omitempty"`
	// The Java class of the exception that triggered this event, if applicable.
	ExceptionType *string `json:"exception_type,omitempty"`
	// The details associated with the Java exception that triggered this event, if applicable.
	ExceptionDetail *string `json:"exception_detail,omitempty"`
}

// NewExecutionEvent instantiates a new ExecutionEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecutionEvent() *ExecutionEvent {
	this := ExecutionEvent{}
	return &this
}

// NewExecutionEventWithDefaults instantiates a new ExecutionEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecutionEventWithDefaults() *ExecutionEvent {
	this := ExecutionEvent{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ExecutionEvent) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionEvent) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ExecutionEvent) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ExecutionEvent) SetId(v string) {
	o.Id = &v
}

// GetExecutionId returns the ExecutionId field value if set, zero value otherwise.
func (o *ExecutionEvent) GetExecutionId() string {
	if o == nil || IsNil(o.ExecutionId) {
		var ret string
		return ret
	}
	return *o.ExecutionId
}

// GetExecutionIdOk returns a tuple with the ExecutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionEvent) GetExecutionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExecutionId) {
		return nil, false
	}
	return o.ExecutionId, true
}

// HasExecutionId returns a boolean if a field has been set.
func (o *ExecutionEvent) HasExecutionId() bool {
	if o != nil && !IsNil(o.ExecutionId) {
		return true
	}

	return false
}

// SetExecutionId gets a reference to the given string and assigns it to the ExecutionId field.
func (o *ExecutionEvent) SetExecutionId(v string) {
	o.ExecutionId = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *ExecutionEvent) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionEvent) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *ExecutionEvent) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *ExecutionEvent) SetEventType(v string) {
	o.EventType = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *ExecutionEvent) GetSeverity() string {
	if o == nil || IsNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionEvent) GetSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *ExecutionEvent) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *ExecutionEvent) SetSeverity(v string) {
	o.Severity = &v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *ExecutionEvent) GetCause() string {
	if o == nil || IsNil(o.Cause) {
		var ret string
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionEvent) GetCauseOk() (*string, bool) {
	if o == nil || IsNil(o.Cause) {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *ExecutionEvent) HasCause() bool {
	if o != nil && !IsNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given string and assigns it to the Cause field.
func (o *ExecutionEvent) SetCause(v string) {
	o.Cause = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ExecutionEvent) GetCount() int64 {
	if o == nil || IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionEvent) GetCountOk() (*int64, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ExecutionEvent) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *ExecutionEvent) SetCount(v int64) {
	o.Count = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ExecutionEvent) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionEvent) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ExecutionEvent) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *ExecutionEvent) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetMaskedObjectName returns the MaskedObjectName field value if set, zero value otherwise.
func (o *ExecutionEvent) GetMaskedObjectName() string {
	if o == nil || IsNil(o.MaskedObjectName) {
		var ret string
		return ret
	}
	return *o.MaskedObjectName
}

// GetMaskedObjectNameOk returns a tuple with the MaskedObjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionEvent) GetMaskedObjectNameOk() (*string, bool) {
	if o == nil || IsNil(o.MaskedObjectName) {
		return nil, false
	}
	return o.MaskedObjectName, true
}

// HasMaskedObjectName returns a boolean if a field has been set.
func (o *ExecutionEvent) HasMaskedObjectName() bool {
	if o != nil && !IsNil(o.MaskedObjectName) {
		return true
	}

	return false
}

// SetMaskedObjectName gets a reference to the given string and assigns it to the MaskedObjectName field.
func (o *ExecutionEvent) SetMaskedObjectName(v string) {
	o.MaskedObjectName = &v
}

// GetAlgorithmName returns the AlgorithmName field value if set, zero value otherwise.
func (o *ExecutionEvent) GetAlgorithmName() string {
	if o == nil || IsNil(o.AlgorithmName) {
		var ret string
		return ret
	}
	return *o.AlgorithmName
}

// GetAlgorithmNameOk returns a tuple with the AlgorithmName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionEvent) GetAlgorithmNameOk() (*string, bool) {
	if o == nil || IsNil(o.AlgorithmName) {
		return nil, false
	}
	return o.AlgorithmName, true
}

// HasAlgorithmName returns a boolean if a field has been set.
func (o *ExecutionEvent) HasAlgorithmName() bool {
	if o != nil && !IsNil(o.AlgorithmName) {
		return true
	}

	return false
}

// SetAlgorithmName gets a reference to the given string and assigns it to the AlgorithmName field.
func (o *ExecutionEvent) SetAlgorithmName(v string) {
	o.AlgorithmName = &v
}

// GetExceptionType returns the ExceptionType field value if set, zero value otherwise.
func (o *ExecutionEvent) GetExceptionType() string {
	if o == nil || IsNil(o.ExceptionType) {
		var ret string
		return ret
	}
	return *o.ExceptionType
}

// GetExceptionTypeOk returns a tuple with the ExceptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionEvent) GetExceptionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ExceptionType) {
		return nil, false
	}
	return o.ExceptionType, true
}

// HasExceptionType returns a boolean if a field has been set.
func (o *ExecutionEvent) HasExceptionType() bool {
	if o != nil && !IsNil(o.ExceptionType) {
		return true
	}

	return false
}

// SetExceptionType gets a reference to the given string and assigns it to the ExceptionType field.
func (o *ExecutionEvent) SetExceptionType(v string) {
	o.ExceptionType = &v
}

// GetExceptionDetail returns the ExceptionDetail field value if set, zero value otherwise.
func (o *ExecutionEvent) GetExceptionDetail() string {
	if o == nil || IsNil(o.ExceptionDetail) {
		var ret string
		return ret
	}
	return *o.ExceptionDetail
}

// GetExceptionDetailOk returns a tuple with the ExceptionDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionEvent) GetExceptionDetailOk() (*string, bool) {
	if o == nil || IsNil(o.ExceptionDetail) {
		return nil, false
	}
	return o.ExceptionDetail, true
}

// HasExceptionDetail returns a boolean if a field has been set.
func (o *ExecutionEvent) HasExceptionDetail() bool {
	if o != nil && !IsNil(o.ExceptionDetail) {
		return true
	}

	return false
}

// SetExceptionDetail gets a reference to the given string and assigns it to the ExceptionDetail field.
func (o *ExecutionEvent) SetExceptionDetail(v string) {
	o.ExceptionDetail = &v
}

func (o ExecutionEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExecutionEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ExecutionId) {
		toSerialize["execution_id"] = o.ExecutionId
	}
	if !IsNil(o.EventType) {
		toSerialize["event_type"] = o.EventType
	}
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !IsNil(o.Cause) {
		toSerialize["cause"] = o.Cause
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.MaskedObjectName) {
		toSerialize["masked_object_name"] = o.MaskedObjectName
	}
	if !IsNil(o.AlgorithmName) {
		toSerialize["algorithm_name"] = o.AlgorithmName
	}
	if !IsNil(o.ExceptionType) {
		toSerialize["exception_type"] = o.ExceptionType
	}
	if !IsNil(o.ExceptionDetail) {
		toSerialize["exception_detail"] = o.ExceptionDetail
	}
	return toSerialize, nil
}

type NullableExecutionEvent struct {
	value *ExecutionEvent
	isSet bool
}

func (v NullableExecutionEvent) Get() *ExecutionEvent {
	return v.value
}

func (v *NullableExecutionEvent) Set(val *ExecutionEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableExecutionEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableExecutionEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecutionEvent(val *ExecutionEvent) *NullableExecutionEvent {
	return &NullableExecutionEvent{value: val, isSet: true}
}

func (v NullableExecutionEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecutionEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


