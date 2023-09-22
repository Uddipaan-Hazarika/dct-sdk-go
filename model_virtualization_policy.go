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

// checks if the VirtualizationPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualizationPolicy{}

// VirtualizationPolicy struct for VirtualizationPolicy
type VirtualizationPolicy struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
	// The namespace id of this virtualization policy.
	NamespaceId *string `json:"namespace_id,omitempty"`
	// The namespace name of this virtualization policy..
	NamespaceName *string `json:"namespace_name,omitempty"`
	// Is this a replicated object.
	IsReplica *bool `json:"is_replica,omitempty"`
	EngineId *string `json:"engine_id,omitempty"`
	PolicyType *string `json:"policy_type,omitempty"`
	TimezoneId *string `json:"timezone_id,omitempty"`
	Schedules []VirtualizationSchedule `json:"schedules,omitempty"`
}

// NewVirtualizationPolicy instantiates a new VirtualizationPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationPolicy() *VirtualizationPolicy {
	this := VirtualizationPolicy{}
	return &this
}

// NewVirtualizationPolicyWithDefaults instantiates a new VirtualizationPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationPolicyWithDefaults() *VirtualizationPolicy {
	this := VirtualizationPolicy{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VirtualizationPolicy) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationPolicy) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VirtualizationPolicy) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VirtualizationPolicy) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualizationPolicy) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationPolicy) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualizationPolicy) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualizationPolicy) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *VirtualizationPolicy) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationPolicy) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *VirtualizationPolicy) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *VirtualizationPolicy) SetNamespace(v string) {
	o.Namespace = &v
}

// GetNamespaceId returns the NamespaceId field value if set, zero value otherwise.
func (o *VirtualizationPolicy) GetNamespaceId() string {
	if o == nil || IsNil(o.NamespaceId) {
		var ret string
		return ret
	}
	return *o.NamespaceId
}

// GetNamespaceIdOk returns a tuple with the NamespaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationPolicy) GetNamespaceIdOk() (*string, bool) {
	if o == nil || IsNil(o.NamespaceId) {
		return nil, false
	}
	return o.NamespaceId, true
}

// HasNamespaceId returns a boolean if a field has been set.
func (o *VirtualizationPolicy) HasNamespaceId() bool {
	if o != nil && !IsNil(o.NamespaceId) {
		return true
	}

	return false
}

// SetNamespaceId gets a reference to the given string and assigns it to the NamespaceId field.
func (o *VirtualizationPolicy) SetNamespaceId(v string) {
	o.NamespaceId = &v
}

// GetNamespaceName returns the NamespaceName field value if set, zero value otherwise.
func (o *VirtualizationPolicy) GetNamespaceName() string {
	if o == nil || IsNil(o.NamespaceName) {
		var ret string
		return ret
	}
	return *o.NamespaceName
}

// GetNamespaceNameOk returns a tuple with the NamespaceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationPolicy) GetNamespaceNameOk() (*string, bool) {
	if o == nil || IsNil(o.NamespaceName) {
		return nil, false
	}
	return o.NamespaceName, true
}

// HasNamespaceName returns a boolean if a field has been set.
func (o *VirtualizationPolicy) HasNamespaceName() bool {
	if o != nil && !IsNil(o.NamespaceName) {
		return true
	}

	return false
}

// SetNamespaceName gets a reference to the given string and assigns it to the NamespaceName field.
func (o *VirtualizationPolicy) SetNamespaceName(v string) {
	o.NamespaceName = &v
}

// GetIsReplica returns the IsReplica field value if set, zero value otherwise.
func (o *VirtualizationPolicy) GetIsReplica() bool {
	if o == nil || IsNil(o.IsReplica) {
		var ret bool
		return ret
	}
	return *o.IsReplica
}

// GetIsReplicaOk returns a tuple with the IsReplica field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationPolicy) GetIsReplicaOk() (*bool, bool) {
	if o == nil || IsNil(o.IsReplica) {
		return nil, false
	}
	return o.IsReplica, true
}

// HasIsReplica returns a boolean if a field has been set.
func (o *VirtualizationPolicy) HasIsReplica() bool {
	if o != nil && !IsNil(o.IsReplica) {
		return true
	}

	return false
}

// SetIsReplica gets a reference to the given bool and assigns it to the IsReplica field.
func (o *VirtualizationPolicy) SetIsReplica(v bool) {
	o.IsReplica = &v
}

// GetEngineId returns the EngineId field value if set, zero value otherwise.
func (o *VirtualizationPolicy) GetEngineId() string {
	if o == nil || IsNil(o.EngineId) {
		var ret string
		return ret
	}
	return *o.EngineId
}

// GetEngineIdOk returns a tuple with the EngineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationPolicy) GetEngineIdOk() (*string, bool) {
	if o == nil || IsNil(o.EngineId) {
		return nil, false
	}
	return o.EngineId, true
}

// HasEngineId returns a boolean if a field has been set.
func (o *VirtualizationPolicy) HasEngineId() bool {
	if o != nil && !IsNil(o.EngineId) {
		return true
	}

	return false
}

// SetEngineId gets a reference to the given string and assigns it to the EngineId field.
func (o *VirtualizationPolicy) SetEngineId(v string) {
	o.EngineId = &v
}

// GetPolicyType returns the PolicyType field value if set, zero value otherwise.
func (o *VirtualizationPolicy) GetPolicyType() string {
	if o == nil || IsNil(o.PolicyType) {
		var ret string
		return ret
	}
	return *o.PolicyType
}

// GetPolicyTypeOk returns a tuple with the PolicyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationPolicy) GetPolicyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyType) {
		return nil, false
	}
	return o.PolicyType, true
}

// HasPolicyType returns a boolean if a field has been set.
func (o *VirtualizationPolicy) HasPolicyType() bool {
	if o != nil && !IsNil(o.PolicyType) {
		return true
	}

	return false
}

// SetPolicyType gets a reference to the given string and assigns it to the PolicyType field.
func (o *VirtualizationPolicy) SetPolicyType(v string) {
	o.PolicyType = &v
}

// GetTimezoneId returns the TimezoneId field value if set, zero value otherwise.
func (o *VirtualizationPolicy) GetTimezoneId() string {
	if o == nil || IsNil(o.TimezoneId) {
		var ret string
		return ret
	}
	return *o.TimezoneId
}

// GetTimezoneIdOk returns a tuple with the TimezoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationPolicy) GetTimezoneIdOk() (*string, bool) {
	if o == nil || IsNil(o.TimezoneId) {
		return nil, false
	}
	return o.TimezoneId, true
}

// HasTimezoneId returns a boolean if a field has been set.
func (o *VirtualizationPolicy) HasTimezoneId() bool {
	if o != nil && !IsNil(o.TimezoneId) {
		return true
	}

	return false
}

// SetTimezoneId gets a reference to the given string and assigns it to the TimezoneId field.
func (o *VirtualizationPolicy) SetTimezoneId(v string) {
	o.TimezoneId = &v
}

// GetSchedules returns the Schedules field value if set, zero value otherwise.
func (o *VirtualizationPolicy) GetSchedules() []VirtualizationSchedule {
	if o == nil || IsNil(o.Schedules) {
		var ret []VirtualizationSchedule
		return ret
	}
	return o.Schedules
}

// GetSchedulesOk returns a tuple with the Schedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationPolicy) GetSchedulesOk() ([]VirtualizationSchedule, bool) {
	if o == nil || IsNil(o.Schedules) {
		return nil, false
	}
	return o.Schedules, true
}

// HasSchedules returns a boolean if a field has been set.
func (o *VirtualizationPolicy) HasSchedules() bool {
	if o != nil && !IsNil(o.Schedules) {
		return true
	}

	return false
}

// SetSchedules gets a reference to the given []VirtualizationSchedule and assigns it to the Schedules field.
func (o *VirtualizationPolicy) SetSchedules(v []VirtualizationSchedule) {
	o.Schedules = v
}

func (o VirtualizationPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualizationPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	if !IsNil(o.NamespaceId) {
		toSerialize["namespace_id"] = o.NamespaceId
	}
	if !IsNil(o.NamespaceName) {
		toSerialize["namespace_name"] = o.NamespaceName
	}
	if !IsNil(o.IsReplica) {
		toSerialize["is_replica"] = o.IsReplica
	}
	if !IsNil(o.EngineId) {
		toSerialize["engine_id"] = o.EngineId
	}
	if !IsNil(o.PolicyType) {
		toSerialize["policy_type"] = o.PolicyType
	}
	if !IsNil(o.TimezoneId) {
		toSerialize["timezone_id"] = o.TimezoneId
	}
	if !IsNil(o.Schedules) {
		toSerialize["schedules"] = o.Schedules
	}
	return toSerialize, nil
}

type NullableVirtualizationPolicy struct {
	value *VirtualizationPolicy
	isSet bool
}

func (v NullableVirtualizationPolicy) Get() *VirtualizationPolicy {
	return v.value
}

func (v *NullableVirtualizationPolicy) Set(val *VirtualizationPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationPolicy(val *VirtualizationPolicy) *NullableVirtualizationPolicy {
	return &NullableVirtualizationPolicy{value: val, isSet: true}
}

func (v NullableVirtualizationPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


