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

// checks if the CredentialsEnvVariable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialsEnvVariable{}

// CredentialsEnvVariable struct for CredentialsEnvVariable
type CredentialsEnvVariable struct {
	// Base name of the environment variables. Variables are named by appending '_USER', '_PASSWORD', '_PUBKEY' and '_PRIVKEY' to this base name, respectively. Variables whose values are not entered or are not present in the type of credential or vault selected, will not be set.
	BaseVarName string `json:"base_var_name"`
	// Password to assign to the environment variables.
	Password *string `json:"password,omitempty"`
	// The name or reference of the vault to assign to the environment variables.
	Vault *string `json:"vault,omitempty"`
	// Vault engine name where the credential is stored.
	HashicorpVaultEngine *string `json:"hashicorp_vault_engine,omitempty"`
	// Path in the vault engine where the credential is stored.
	HashicorpVaultSecretPath *string `json:"hashicorp_vault_secret_path,omitempty"`
	// Hashicorp vault key for the username in the key-value store.
	HashicorpVaultUsernameKey *string `json:"hashicorp_vault_username_key,omitempty"`
	// Hashicorp vault key for the password in the key-value store.
	HashicorpVaultSecretKey *string `json:"hashicorp_vault_secret_key,omitempty"`
	// Azure key vault name.
	AzureVaultName *string `json:"azure_vault_name,omitempty"`
	// Azure vault key in the key-value store.
	AzureVaultUsernameKey *string `json:"azure_vault_username_key,omitempty"`
	// Azure vault key in the key-value store.
	AzureVaultSecretKey *string `json:"azure_vault_secret_key,omitempty"`
	// Query to find a credential in the CyberArk vault.
	CyberarkVaultQueryString *string `json:"cyberark_vault_query_string,omitempty"`
}

// NewCredentialsEnvVariable instantiates a new CredentialsEnvVariable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialsEnvVariable(baseVarName string) *CredentialsEnvVariable {
	this := CredentialsEnvVariable{}
	this.BaseVarName = baseVarName
	return &this
}

// NewCredentialsEnvVariableWithDefaults instantiates a new CredentialsEnvVariable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialsEnvVariableWithDefaults() *CredentialsEnvVariable {
	this := CredentialsEnvVariable{}
	return &this
}

// GetBaseVarName returns the BaseVarName field value
func (o *CredentialsEnvVariable) GetBaseVarName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseVarName
}

// GetBaseVarNameOk returns a tuple with the BaseVarName field value
// and a boolean to check if the value has been set.
func (o *CredentialsEnvVariable) GetBaseVarNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseVarName, true
}

// SetBaseVarName sets field value
func (o *CredentialsEnvVariable) SetBaseVarName(v string) {
	o.BaseVarName = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *CredentialsEnvVariable) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsEnvVariable) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *CredentialsEnvVariable) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *CredentialsEnvVariable) SetPassword(v string) {
	o.Password = &v
}

// GetVault returns the Vault field value if set, zero value otherwise.
func (o *CredentialsEnvVariable) GetVault() string {
	if o == nil || IsNil(o.Vault) {
		var ret string
		return ret
	}
	return *o.Vault
}

// GetVaultOk returns a tuple with the Vault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsEnvVariable) GetVaultOk() (*string, bool) {
	if o == nil || IsNil(o.Vault) {
		return nil, false
	}
	return o.Vault, true
}

// HasVault returns a boolean if a field has been set.
func (o *CredentialsEnvVariable) HasVault() bool {
	if o != nil && !IsNil(o.Vault) {
		return true
	}

	return false
}

// SetVault gets a reference to the given string and assigns it to the Vault field.
func (o *CredentialsEnvVariable) SetVault(v string) {
	o.Vault = &v
}

// GetHashicorpVaultEngine returns the HashicorpVaultEngine field value if set, zero value otherwise.
func (o *CredentialsEnvVariable) GetHashicorpVaultEngine() string {
	if o == nil || IsNil(o.HashicorpVaultEngine) {
		var ret string
		return ret
	}
	return *o.HashicorpVaultEngine
}

// GetHashicorpVaultEngineOk returns a tuple with the HashicorpVaultEngine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsEnvVariable) GetHashicorpVaultEngineOk() (*string, bool) {
	if o == nil || IsNil(o.HashicorpVaultEngine) {
		return nil, false
	}
	return o.HashicorpVaultEngine, true
}

// HasHashicorpVaultEngine returns a boolean if a field has been set.
func (o *CredentialsEnvVariable) HasHashicorpVaultEngine() bool {
	if o != nil && !IsNil(o.HashicorpVaultEngine) {
		return true
	}

	return false
}

// SetHashicorpVaultEngine gets a reference to the given string and assigns it to the HashicorpVaultEngine field.
func (o *CredentialsEnvVariable) SetHashicorpVaultEngine(v string) {
	o.HashicorpVaultEngine = &v
}

// GetHashicorpVaultSecretPath returns the HashicorpVaultSecretPath field value if set, zero value otherwise.
func (o *CredentialsEnvVariable) GetHashicorpVaultSecretPath() string {
	if o == nil || IsNil(o.HashicorpVaultSecretPath) {
		var ret string
		return ret
	}
	return *o.HashicorpVaultSecretPath
}

// GetHashicorpVaultSecretPathOk returns a tuple with the HashicorpVaultSecretPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsEnvVariable) GetHashicorpVaultSecretPathOk() (*string, bool) {
	if o == nil || IsNil(o.HashicorpVaultSecretPath) {
		return nil, false
	}
	return o.HashicorpVaultSecretPath, true
}

// HasHashicorpVaultSecretPath returns a boolean if a field has been set.
func (o *CredentialsEnvVariable) HasHashicorpVaultSecretPath() bool {
	if o != nil && !IsNil(o.HashicorpVaultSecretPath) {
		return true
	}

	return false
}

// SetHashicorpVaultSecretPath gets a reference to the given string and assigns it to the HashicorpVaultSecretPath field.
func (o *CredentialsEnvVariable) SetHashicorpVaultSecretPath(v string) {
	o.HashicorpVaultSecretPath = &v
}

// GetHashicorpVaultUsernameKey returns the HashicorpVaultUsernameKey field value if set, zero value otherwise.
func (o *CredentialsEnvVariable) GetHashicorpVaultUsernameKey() string {
	if o == nil || IsNil(o.HashicorpVaultUsernameKey) {
		var ret string
		return ret
	}
	return *o.HashicorpVaultUsernameKey
}

// GetHashicorpVaultUsernameKeyOk returns a tuple with the HashicorpVaultUsernameKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsEnvVariable) GetHashicorpVaultUsernameKeyOk() (*string, bool) {
	if o == nil || IsNil(o.HashicorpVaultUsernameKey) {
		return nil, false
	}
	return o.HashicorpVaultUsernameKey, true
}

// HasHashicorpVaultUsernameKey returns a boolean if a field has been set.
func (o *CredentialsEnvVariable) HasHashicorpVaultUsernameKey() bool {
	if o != nil && !IsNil(o.HashicorpVaultUsernameKey) {
		return true
	}

	return false
}

// SetHashicorpVaultUsernameKey gets a reference to the given string and assigns it to the HashicorpVaultUsernameKey field.
func (o *CredentialsEnvVariable) SetHashicorpVaultUsernameKey(v string) {
	o.HashicorpVaultUsernameKey = &v
}

// GetHashicorpVaultSecretKey returns the HashicorpVaultSecretKey field value if set, zero value otherwise.
func (o *CredentialsEnvVariable) GetHashicorpVaultSecretKey() string {
	if o == nil || IsNil(o.HashicorpVaultSecretKey) {
		var ret string
		return ret
	}
	return *o.HashicorpVaultSecretKey
}

// GetHashicorpVaultSecretKeyOk returns a tuple with the HashicorpVaultSecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsEnvVariable) GetHashicorpVaultSecretKeyOk() (*string, bool) {
	if o == nil || IsNil(o.HashicorpVaultSecretKey) {
		return nil, false
	}
	return o.HashicorpVaultSecretKey, true
}

// HasHashicorpVaultSecretKey returns a boolean if a field has been set.
func (o *CredentialsEnvVariable) HasHashicorpVaultSecretKey() bool {
	if o != nil && !IsNil(o.HashicorpVaultSecretKey) {
		return true
	}

	return false
}

// SetHashicorpVaultSecretKey gets a reference to the given string and assigns it to the HashicorpVaultSecretKey field.
func (o *CredentialsEnvVariable) SetHashicorpVaultSecretKey(v string) {
	o.HashicorpVaultSecretKey = &v
}

// GetAzureVaultName returns the AzureVaultName field value if set, zero value otherwise.
func (o *CredentialsEnvVariable) GetAzureVaultName() string {
	if o == nil || IsNil(o.AzureVaultName) {
		var ret string
		return ret
	}
	return *o.AzureVaultName
}

// GetAzureVaultNameOk returns a tuple with the AzureVaultName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsEnvVariable) GetAzureVaultNameOk() (*string, bool) {
	if o == nil || IsNil(o.AzureVaultName) {
		return nil, false
	}
	return o.AzureVaultName, true
}

// HasAzureVaultName returns a boolean if a field has been set.
func (o *CredentialsEnvVariable) HasAzureVaultName() bool {
	if o != nil && !IsNil(o.AzureVaultName) {
		return true
	}

	return false
}

// SetAzureVaultName gets a reference to the given string and assigns it to the AzureVaultName field.
func (o *CredentialsEnvVariable) SetAzureVaultName(v string) {
	o.AzureVaultName = &v
}

// GetAzureVaultUsernameKey returns the AzureVaultUsernameKey field value if set, zero value otherwise.
func (o *CredentialsEnvVariable) GetAzureVaultUsernameKey() string {
	if o == nil || IsNil(o.AzureVaultUsernameKey) {
		var ret string
		return ret
	}
	return *o.AzureVaultUsernameKey
}

// GetAzureVaultUsernameKeyOk returns a tuple with the AzureVaultUsernameKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsEnvVariable) GetAzureVaultUsernameKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AzureVaultUsernameKey) {
		return nil, false
	}
	return o.AzureVaultUsernameKey, true
}

// HasAzureVaultUsernameKey returns a boolean if a field has been set.
func (o *CredentialsEnvVariable) HasAzureVaultUsernameKey() bool {
	if o != nil && !IsNil(o.AzureVaultUsernameKey) {
		return true
	}

	return false
}

// SetAzureVaultUsernameKey gets a reference to the given string and assigns it to the AzureVaultUsernameKey field.
func (o *CredentialsEnvVariable) SetAzureVaultUsernameKey(v string) {
	o.AzureVaultUsernameKey = &v
}

// GetAzureVaultSecretKey returns the AzureVaultSecretKey field value if set, zero value otherwise.
func (o *CredentialsEnvVariable) GetAzureVaultSecretKey() string {
	if o == nil || IsNil(o.AzureVaultSecretKey) {
		var ret string
		return ret
	}
	return *o.AzureVaultSecretKey
}

// GetAzureVaultSecretKeyOk returns a tuple with the AzureVaultSecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsEnvVariable) GetAzureVaultSecretKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AzureVaultSecretKey) {
		return nil, false
	}
	return o.AzureVaultSecretKey, true
}

// HasAzureVaultSecretKey returns a boolean if a field has been set.
func (o *CredentialsEnvVariable) HasAzureVaultSecretKey() bool {
	if o != nil && !IsNil(o.AzureVaultSecretKey) {
		return true
	}

	return false
}

// SetAzureVaultSecretKey gets a reference to the given string and assigns it to the AzureVaultSecretKey field.
func (o *CredentialsEnvVariable) SetAzureVaultSecretKey(v string) {
	o.AzureVaultSecretKey = &v
}

// GetCyberarkVaultQueryString returns the CyberarkVaultQueryString field value if set, zero value otherwise.
func (o *CredentialsEnvVariable) GetCyberarkVaultQueryString() string {
	if o == nil || IsNil(o.CyberarkVaultQueryString) {
		var ret string
		return ret
	}
	return *o.CyberarkVaultQueryString
}

// GetCyberarkVaultQueryStringOk returns a tuple with the CyberarkVaultQueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsEnvVariable) GetCyberarkVaultQueryStringOk() (*string, bool) {
	if o == nil || IsNil(o.CyberarkVaultQueryString) {
		return nil, false
	}
	return o.CyberarkVaultQueryString, true
}

// HasCyberarkVaultQueryString returns a boolean if a field has been set.
func (o *CredentialsEnvVariable) HasCyberarkVaultQueryString() bool {
	if o != nil && !IsNil(o.CyberarkVaultQueryString) {
		return true
	}

	return false
}

// SetCyberarkVaultQueryString gets a reference to the given string and assigns it to the CyberarkVaultQueryString field.
func (o *CredentialsEnvVariable) SetCyberarkVaultQueryString(v string) {
	o.CyberarkVaultQueryString = &v
}

func (o CredentialsEnvVariable) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CredentialsEnvVariable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["base_var_name"] = o.BaseVarName
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Vault) {
		toSerialize["vault"] = o.Vault
	}
	if !IsNil(o.HashicorpVaultEngine) {
		toSerialize["hashicorp_vault_engine"] = o.HashicorpVaultEngine
	}
	if !IsNil(o.HashicorpVaultSecretPath) {
		toSerialize["hashicorp_vault_secret_path"] = o.HashicorpVaultSecretPath
	}
	if !IsNil(o.HashicorpVaultUsernameKey) {
		toSerialize["hashicorp_vault_username_key"] = o.HashicorpVaultUsernameKey
	}
	if !IsNil(o.HashicorpVaultSecretKey) {
		toSerialize["hashicorp_vault_secret_key"] = o.HashicorpVaultSecretKey
	}
	if !IsNil(o.AzureVaultName) {
		toSerialize["azure_vault_name"] = o.AzureVaultName
	}
	if !IsNil(o.AzureVaultUsernameKey) {
		toSerialize["azure_vault_username_key"] = o.AzureVaultUsernameKey
	}
	if !IsNil(o.AzureVaultSecretKey) {
		toSerialize["azure_vault_secret_key"] = o.AzureVaultSecretKey
	}
	if !IsNil(o.CyberarkVaultQueryString) {
		toSerialize["cyberark_vault_query_string"] = o.CyberarkVaultQueryString
	}
	return toSerialize, nil
}

type NullableCredentialsEnvVariable struct {
	value *CredentialsEnvVariable
	isSet bool
}

func (v NullableCredentialsEnvVariable) Get() *CredentialsEnvVariable {
	return v.value
}

func (v *NullableCredentialsEnvVariable) Set(val *CredentialsEnvVariable) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialsEnvVariable) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialsEnvVariable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialsEnvVariable(val *CredentialsEnvVariable) *NullableCredentialsEnvVariable {
	return &NullableCredentialsEnvVariable{value: val, isSet: true}
}

func (v NullableCredentialsEnvVariable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialsEnvVariable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

