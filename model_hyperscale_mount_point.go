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

// checks if the HyperscaleMountPoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperscaleMountPoint{}

// HyperscaleMountPoint A Mount Point for the Hyperscale instance to write to a staging area and engines to read from.
type HyperscaleMountPoint struct {
	// The ID of the Hyperscale Mount Point.
	Id *string `json:"id,omitempty"`
	// The ID of the Hyperscale instance of this Mount Point.
	HyperscaleInstanceId *string `json:"hyperscale_instance_id,omitempty"`
	// Name of the mount, unique for a hyperscale instance. This name will be used as a directory name by the Hyperscale instance.
	Name *string `json:"name,omitempty"`
	// The host name of the server.
	Hostname *string `json:"hostname,omitempty"`
	// The path to the directory on the filesystem to mount.
	MountPath *string `json:"mount_path,omitempty"`
	// The type of filesystem. Enum having values- CIFS, NFS3, NFS4.
	MountType *string `json:"mount_type,omitempty"`
	// The options for mount. The endpoint will return all default options and user specified options.
	Options *string `json:"options,omitempty"`
}

// NewHyperscaleMountPoint instantiates a new HyperscaleMountPoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperscaleMountPoint() *HyperscaleMountPoint {
	this := HyperscaleMountPoint{}
	return &this
}

// NewHyperscaleMountPointWithDefaults instantiates a new HyperscaleMountPoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperscaleMountPointWithDefaults() *HyperscaleMountPoint {
	this := HyperscaleMountPoint{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HyperscaleMountPoint) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleMountPoint) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HyperscaleMountPoint) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *HyperscaleMountPoint) SetId(v string) {
	o.Id = &v
}

// GetHyperscaleInstanceId returns the HyperscaleInstanceId field value if set, zero value otherwise.
func (o *HyperscaleMountPoint) GetHyperscaleInstanceId() string {
	if o == nil || IsNil(o.HyperscaleInstanceId) {
		var ret string
		return ret
	}
	return *o.HyperscaleInstanceId
}

// GetHyperscaleInstanceIdOk returns a tuple with the HyperscaleInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleMountPoint) GetHyperscaleInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.HyperscaleInstanceId) {
		return nil, false
	}
	return o.HyperscaleInstanceId, true
}

// HasHyperscaleInstanceId returns a boolean if a field has been set.
func (o *HyperscaleMountPoint) HasHyperscaleInstanceId() bool {
	if o != nil && !IsNil(o.HyperscaleInstanceId) {
		return true
	}

	return false
}

// SetHyperscaleInstanceId gets a reference to the given string and assigns it to the HyperscaleInstanceId field.
func (o *HyperscaleMountPoint) SetHyperscaleInstanceId(v string) {
	o.HyperscaleInstanceId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HyperscaleMountPoint) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleMountPoint) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HyperscaleMountPoint) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HyperscaleMountPoint) SetName(v string) {
	o.Name = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *HyperscaleMountPoint) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleMountPoint) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *HyperscaleMountPoint) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *HyperscaleMountPoint) SetHostname(v string) {
	o.Hostname = &v
}

// GetMountPath returns the MountPath field value if set, zero value otherwise.
func (o *HyperscaleMountPoint) GetMountPath() string {
	if o == nil || IsNil(o.MountPath) {
		var ret string
		return ret
	}
	return *o.MountPath
}

// GetMountPathOk returns a tuple with the MountPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleMountPoint) GetMountPathOk() (*string, bool) {
	if o == nil || IsNil(o.MountPath) {
		return nil, false
	}
	return o.MountPath, true
}

// HasMountPath returns a boolean if a field has been set.
func (o *HyperscaleMountPoint) HasMountPath() bool {
	if o != nil && !IsNil(o.MountPath) {
		return true
	}

	return false
}

// SetMountPath gets a reference to the given string and assigns it to the MountPath field.
func (o *HyperscaleMountPoint) SetMountPath(v string) {
	o.MountPath = &v
}

// GetMountType returns the MountType field value if set, zero value otherwise.
func (o *HyperscaleMountPoint) GetMountType() string {
	if o == nil || IsNil(o.MountType) {
		var ret string
		return ret
	}
	return *o.MountType
}

// GetMountTypeOk returns a tuple with the MountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleMountPoint) GetMountTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MountType) {
		return nil, false
	}
	return o.MountType, true
}

// HasMountType returns a boolean if a field has been set.
func (o *HyperscaleMountPoint) HasMountType() bool {
	if o != nil && !IsNil(o.MountType) {
		return true
	}

	return false
}

// SetMountType gets a reference to the given string and assigns it to the MountType field.
func (o *HyperscaleMountPoint) SetMountType(v string) {
	o.MountType = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *HyperscaleMountPoint) GetOptions() string {
	if o == nil || IsNil(o.Options) {
		var ret string
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleMountPoint) GetOptionsOk() (*string, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *HyperscaleMountPoint) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given string and assigns it to the Options field.
func (o *HyperscaleMountPoint) SetOptions(v string) {
	o.Options = &v
}

func (o HyperscaleMountPoint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperscaleMountPoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.HyperscaleInstanceId) {
		toSerialize["hyperscale_instance_id"] = o.HyperscaleInstanceId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.MountPath) {
		toSerialize["mount_path"] = o.MountPath
	}
	if !IsNil(o.MountType) {
		toSerialize["mount_type"] = o.MountType
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return toSerialize, nil
}

type NullableHyperscaleMountPoint struct {
	value *HyperscaleMountPoint
	isSet bool
}

func (v NullableHyperscaleMountPoint) Get() *HyperscaleMountPoint {
	return v.value
}

func (v *NullableHyperscaleMountPoint) Set(val *HyperscaleMountPoint) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperscaleMountPoint) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperscaleMountPoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperscaleMountPoint(val *HyperscaleMountPoint) *NullableHyperscaleMountPoint {
	return &NullableHyperscaleMountPoint{value: val, isSet: true}
}

func (v NullableHyperscaleMountPoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperscaleMountPoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


