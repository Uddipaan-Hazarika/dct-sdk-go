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

// checks if the HyperscaleInstanceRegistrationParameter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperscaleInstanceRegistrationParameter{}

// HyperscaleInstanceRegistrationParameter Parameters to register and authenticate a Hyperscale instance
type HyperscaleInstanceRegistrationParameter struct {
	// Name in DCT of the Hyperscale instance.
	Name string `json:"name"`
	// Hostname of the Hyperscale instance.
	Hostname string `json:"hostname"`
	// API key to connect to the Hyperscale instance.
	ApiKey NullableString `json:"api_key"`
	// Allow connections to the hyperscale instance over HTTPs without validating the TLS certificate. Even though the connection to the hyperscale instance might be performed over HTTPs, setting this property eliminates the protection against a man-in-the-middle attach for connections to this engine. Instead, consider creating a truststore with a Certificate Authority to validate the hyperscale instance's certificate, and set the truststore_filename property. 
	InsecureSsl *bool `json:"insecure_ssl,omitempty"`
	// Ignore validation of the name associated to the TLS certificate when connecting to the hyperscale instance over HTTPs. Setting this value must only be done if the TLS certificate of the hyperscale instance does not match the hostname, and the TLS configuration of the hyperscale instance cannot be fixed. Setting this property reduces the protection against a man-in-the-middle attack for connections to this engine. This is ignored if insecure_ssl is set. 
	UnsafeSslHostnameCheck *bool `json:"unsafe_ssl_hostname_check,omitempty"`
	// File name of a truststore which can be used to validate the TLS certificate of the hyperscale instance. The truststore must be available at /etc/config/certs/<truststore_filename> 
	TruststoreFilename NullableString `json:"truststore_filename,omitempty"`
	// Password to read the truststore. 
	TruststorePassword NullableString `json:"truststore_password,omitempty"`
	// The tags to be created for this engine.
	Tags []Tag `json:"tags,omitempty"`
}

// NewHyperscaleInstanceRegistrationParameter instantiates a new HyperscaleInstanceRegistrationParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperscaleInstanceRegistrationParameter(name string, hostname string, apiKey NullableString) *HyperscaleInstanceRegistrationParameter {
	this := HyperscaleInstanceRegistrationParameter{}
	this.Name = name
	this.Hostname = hostname
	this.ApiKey = apiKey
	var insecureSsl bool = false
	this.InsecureSsl = &insecureSsl
	var unsafeSslHostnameCheck bool = false
	this.UnsafeSslHostnameCheck = &unsafeSslHostnameCheck
	return &this
}

// NewHyperscaleInstanceRegistrationParameterWithDefaults instantiates a new HyperscaleInstanceRegistrationParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperscaleInstanceRegistrationParameterWithDefaults() *HyperscaleInstanceRegistrationParameter {
	this := HyperscaleInstanceRegistrationParameter{}
	var insecureSsl bool = false
	this.InsecureSsl = &insecureSsl
	var unsafeSslHostnameCheck bool = false
	this.UnsafeSslHostnameCheck = &unsafeSslHostnameCheck
	return &this
}

// GetName returns the Name field value
func (o *HyperscaleInstanceRegistrationParameter) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *HyperscaleInstanceRegistrationParameter) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *HyperscaleInstanceRegistrationParameter) SetName(v string) {
	o.Name = v
}

// GetHostname returns the Hostname field value
func (o *HyperscaleInstanceRegistrationParameter) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *HyperscaleInstanceRegistrationParameter) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *HyperscaleInstanceRegistrationParameter) SetHostname(v string) {
	o.Hostname = v
}

// GetApiKey returns the ApiKey field value
// If the value is explicit nil, the zero value for string will be returned
func (o *HyperscaleInstanceRegistrationParameter) GetApiKey() string {
	if o == nil || o.ApiKey.Get() == nil {
		var ret string
		return ret
	}

	return *o.ApiKey.Get()
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperscaleInstanceRegistrationParameter) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApiKey.Get(), o.ApiKey.IsSet()
}

// SetApiKey sets field value
func (o *HyperscaleInstanceRegistrationParameter) SetApiKey(v string) {
	o.ApiKey.Set(&v)
}

// GetInsecureSsl returns the InsecureSsl field value if set, zero value otherwise.
func (o *HyperscaleInstanceRegistrationParameter) GetInsecureSsl() bool {
	if o == nil || IsNil(o.InsecureSsl) {
		var ret bool
		return ret
	}
	return *o.InsecureSsl
}

// GetInsecureSslOk returns a tuple with the InsecureSsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleInstanceRegistrationParameter) GetInsecureSslOk() (*bool, bool) {
	if o == nil || IsNil(o.InsecureSsl) {
		return nil, false
	}
	return o.InsecureSsl, true
}

// HasInsecureSsl returns a boolean if a field has been set.
func (o *HyperscaleInstanceRegistrationParameter) HasInsecureSsl() bool {
	if o != nil && !IsNil(o.InsecureSsl) {
		return true
	}

	return false
}

// SetInsecureSsl gets a reference to the given bool and assigns it to the InsecureSsl field.
func (o *HyperscaleInstanceRegistrationParameter) SetInsecureSsl(v bool) {
	o.InsecureSsl = &v
}

// GetUnsafeSslHostnameCheck returns the UnsafeSslHostnameCheck field value if set, zero value otherwise.
func (o *HyperscaleInstanceRegistrationParameter) GetUnsafeSslHostnameCheck() bool {
	if o == nil || IsNil(o.UnsafeSslHostnameCheck) {
		var ret bool
		return ret
	}
	return *o.UnsafeSslHostnameCheck
}

// GetUnsafeSslHostnameCheckOk returns a tuple with the UnsafeSslHostnameCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleInstanceRegistrationParameter) GetUnsafeSslHostnameCheckOk() (*bool, bool) {
	if o == nil || IsNil(o.UnsafeSslHostnameCheck) {
		return nil, false
	}
	return o.UnsafeSslHostnameCheck, true
}

// HasUnsafeSslHostnameCheck returns a boolean if a field has been set.
func (o *HyperscaleInstanceRegistrationParameter) HasUnsafeSslHostnameCheck() bool {
	if o != nil && !IsNil(o.UnsafeSslHostnameCheck) {
		return true
	}

	return false
}

// SetUnsafeSslHostnameCheck gets a reference to the given bool and assigns it to the UnsafeSslHostnameCheck field.
func (o *HyperscaleInstanceRegistrationParameter) SetUnsafeSslHostnameCheck(v bool) {
	o.UnsafeSslHostnameCheck = &v
}

// GetTruststoreFilename returns the TruststoreFilename field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperscaleInstanceRegistrationParameter) GetTruststoreFilename() string {
	if o == nil || IsNil(o.TruststoreFilename.Get()) {
		var ret string
		return ret
	}
	return *o.TruststoreFilename.Get()
}

// GetTruststoreFilenameOk returns a tuple with the TruststoreFilename field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperscaleInstanceRegistrationParameter) GetTruststoreFilenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TruststoreFilename.Get(), o.TruststoreFilename.IsSet()
}

// HasTruststoreFilename returns a boolean if a field has been set.
func (o *HyperscaleInstanceRegistrationParameter) HasTruststoreFilename() bool {
	if o != nil && o.TruststoreFilename.IsSet() {
		return true
	}

	return false
}

// SetTruststoreFilename gets a reference to the given NullableString and assigns it to the TruststoreFilename field.
func (o *HyperscaleInstanceRegistrationParameter) SetTruststoreFilename(v string) {
	o.TruststoreFilename.Set(&v)
}
// SetTruststoreFilenameNil sets the value for TruststoreFilename to be an explicit nil
func (o *HyperscaleInstanceRegistrationParameter) SetTruststoreFilenameNil() {
	o.TruststoreFilename.Set(nil)
}

// UnsetTruststoreFilename ensures that no value is present for TruststoreFilename, not even an explicit nil
func (o *HyperscaleInstanceRegistrationParameter) UnsetTruststoreFilename() {
	o.TruststoreFilename.Unset()
}

// GetTruststorePassword returns the TruststorePassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperscaleInstanceRegistrationParameter) GetTruststorePassword() string {
	if o == nil || IsNil(o.TruststorePassword.Get()) {
		var ret string
		return ret
	}
	return *o.TruststorePassword.Get()
}

// GetTruststorePasswordOk returns a tuple with the TruststorePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperscaleInstanceRegistrationParameter) GetTruststorePasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TruststorePassword.Get(), o.TruststorePassword.IsSet()
}

// HasTruststorePassword returns a boolean if a field has been set.
func (o *HyperscaleInstanceRegistrationParameter) HasTruststorePassword() bool {
	if o != nil && o.TruststorePassword.IsSet() {
		return true
	}

	return false
}

// SetTruststorePassword gets a reference to the given NullableString and assigns it to the TruststorePassword field.
func (o *HyperscaleInstanceRegistrationParameter) SetTruststorePassword(v string) {
	o.TruststorePassword.Set(&v)
}
// SetTruststorePasswordNil sets the value for TruststorePassword to be an explicit nil
func (o *HyperscaleInstanceRegistrationParameter) SetTruststorePasswordNil() {
	o.TruststorePassword.Set(nil)
}

// UnsetTruststorePassword ensures that no value is present for TruststorePassword, not even an explicit nil
func (o *HyperscaleInstanceRegistrationParameter) UnsetTruststorePassword() {
	o.TruststorePassword.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *HyperscaleInstanceRegistrationParameter) GetTags() []Tag {
	if o == nil || IsNil(o.Tags) {
		var ret []Tag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleInstanceRegistrationParameter) GetTagsOk() ([]Tag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *HyperscaleInstanceRegistrationParameter) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []Tag and assigns it to the Tags field.
func (o *HyperscaleInstanceRegistrationParameter) SetTags(v []Tag) {
	o.Tags = v
}

func (o HyperscaleInstanceRegistrationParameter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperscaleInstanceRegistrationParameter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["hostname"] = o.Hostname
	toSerialize["api_key"] = o.ApiKey.Get()
	if !IsNil(o.InsecureSsl) {
		toSerialize["insecure_ssl"] = o.InsecureSsl
	}
	if !IsNil(o.UnsafeSslHostnameCheck) {
		toSerialize["unsafe_ssl_hostname_check"] = o.UnsafeSslHostnameCheck
	}
	if o.TruststoreFilename.IsSet() {
		toSerialize["truststore_filename"] = o.TruststoreFilename.Get()
	}
	if o.TruststorePassword.IsSet() {
		toSerialize["truststore_password"] = o.TruststorePassword.Get()
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableHyperscaleInstanceRegistrationParameter struct {
	value *HyperscaleInstanceRegistrationParameter
	isSet bool
}

func (v NullableHyperscaleInstanceRegistrationParameter) Get() *HyperscaleInstanceRegistrationParameter {
	return v.value
}

func (v *NullableHyperscaleInstanceRegistrationParameter) Set(val *HyperscaleInstanceRegistrationParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperscaleInstanceRegistrationParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperscaleInstanceRegistrationParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperscaleInstanceRegistrationParameter(val *HyperscaleInstanceRegistrationParameter) *NullableHyperscaleInstanceRegistrationParameter {
	return &NullableHyperscaleInstanceRegistrationParameter{value: val, isSet: true}
}

func (v NullableHyperscaleInstanceRegistrationParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperscaleInstanceRegistrationParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

