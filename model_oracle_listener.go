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

// checks if the OracleListener type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OracleListener{}

// OracleListener struct for OracleListener
type OracleListener struct {
	// Id of this listener.
	Id *string `json:"id,omitempty"`
	// Name of this listener.
	Name *string `json:"name,omitempty"`
	// Type of this listener.
	Type *string `json:"type,omitempty"`
	// The list of protocol addresses for this listener.
	ProtocolAddresses []string `json:"protocol_addresses,omitempty"`
	// The list of client endpoints for this listener.
	ClientEndpoints []string `json:"client_endpoints,omitempty"`
	// Whether this listener was automatically discovered or not.
	IsDiscovered *bool `json:"is_discovered,omitempty"`
	// Id to the host this listener is associated with.
	HostId *string `json:"host_id,omitempty"`
}

// NewOracleListener instantiates a new OracleListener object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOracleListener() *OracleListener {
	this := OracleListener{}
	return &this
}

// NewOracleListenerWithDefaults instantiates a new OracleListener object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOracleListenerWithDefaults() *OracleListener {
	this := OracleListener{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OracleListener) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleListener) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OracleListener) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OracleListener) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OracleListener) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleListener) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OracleListener) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OracleListener) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OracleListener) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleListener) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OracleListener) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OracleListener) SetType(v string) {
	o.Type = &v
}

// GetProtocolAddresses returns the ProtocolAddresses field value if set, zero value otherwise.
func (o *OracleListener) GetProtocolAddresses() []string {
	if o == nil || IsNil(o.ProtocolAddresses) {
		var ret []string
		return ret
	}
	return o.ProtocolAddresses
}

// GetProtocolAddressesOk returns a tuple with the ProtocolAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleListener) GetProtocolAddressesOk() ([]string, bool) {
	if o == nil || IsNil(o.ProtocolAddresses) {
		return nil, false
	}
	return o.ProtocolAddresses, true
}

// HasProtocolAddresses returns a boolean if a field has been set.
func (o *OracleListener) HasProtocolAddresses() bool {
	if o != nil && !IsNil(o.ProtocolAddresses) {
		return true
	}

	return false
}

// SetProtocolAddresses gets a reference to the given []string and assigns it to the ProtocolAddresses field.
func (o *OracleListener) SetProtocolAddresses(v []string) {
	o.ProtocolAddresses = v
}

// GetClientEndpoints returns the ClientEndpoints field value if set, zero value otherwise.
func (o *OracleListener) GetClientEndpoints() []string {
	if o == nil || IsNil(o.ClientEndpoints) {
		var ret []string
		return ret
	}
	return o.ClientEndpoints
}

// GetClientEndpointsOk returns a tuple with the ClientEndpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleListener) GetClientEndpointsOk() ([]string, bool) {
	if o == nil || IsNil(o.ClientEndpoints) {
		return nil, false
	}
	return o.ClientEndpoints, true
}

// HasClientEndpoints returns a boolean if a field has been set.
func (o *OracleListener) HasClientEndpoints() bool {
	if o != nil && !IsNil(o.ClientEndpoints) {
		return true
	}

	return false
}

// SetClientEndpoints gets a reference to the given []string and assigns it to the ClientEndpoints field.
func (o *OracleListener) SetClientEndpoints(v []string) {
	o.ClientEndpoints = v
}

// GetIsDiscovered returns the IsDiscovered field value if set, zero value otherwise.
func (o *OracleListener) GetIsDiscovered() bool {
	if o == nil || IsNil(o.IsDiscovered) {
		var ret bool
		return ret
	}
	return *o.IsDiscovered
}

// GetIsDiscoveredOk returns a tuple with the IsDiscovered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleListener) GetIsDiscoveredOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDiscovered) {
		return nil, false
	}
	return o.IsDiscovered, true
}

// HasIsDiscovered returns a boolean if a field has been set.
func (o *OracleListener) HasIsDiscovered() bool {
	if o != nil && !IsNil(o.IsDiscovered) {
		return true
	}

	return false
}

// SetIsDiscovered gets a reference to the given bool and assigns it to the IsDiscovered field.
func (o *OracleListener) SetIsDiscovered(v bool) {
	o.IsDiscovered = &v
}

// GetHostId returns the HostId field value if set, zero value otherwise.
func (o *OracleListener) GetHostId() string {
	if o == nil || IsNil(o.HostId) {
		var ret string
		return ret
	}
	return *o.HostId
}

// GetHostIdOk returns a tuple with the HostId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleListener) GetHostIdOk() (*string, bool) {
	if o == nil || IsNil(o.HostId) {
		return nil, false
	}
	return o.HostId, true
}

// HasHostId returns a boolean if a field has been set.
func (o *OracleListener) HasHostId() bool {
	if o != nil && !IsNil(o.HostId) {
		return true
	}

	return false
}

// SetHostId gets a reference to the given string and assigns it to the HostId field.
func (o *OracleListener) SetHostId(v string) {
	o.HostId = &v
}

func (o OracleListener) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OracleListener) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.ProtocolAddresses) {
		toSerialize["protocol_addresses"] = o.ProtocolAddresses
	}
	if !IsNil(o.ClientEndpoints) {
		toSerialize["client_endpoints"] = o.ClientEndpoints
	}
	if !IsNil(o.IsDiscovered) {
		toSerialize["is_discovered"] = o.IsDiscovered
	}
	if !IsNil(o.HostId) {
		toSerialize["host_id"] = o.HostId
	}
	return toSerialize, nil
}

type NullableOracleListener struct {
	value *OracleListener
	isSet bool
}

func (v NullableOracleListener) Get() *OracleListener {
	return v.value
}

func (v *NullableOracleListener) Set(val *OracleListener) {
	v.value = val
	v.isSet = true
}

func (v NullableOracleListener) IsSet() bool {
	return v.isSet
}

func (v *NullableOracleListener) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOracleListener(val *OracleListener) *NullableOracleListener {
	return &NullableOracleListener{value: val, isSet: true}
}

func (v NullableOracleListener) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOracleListener) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


