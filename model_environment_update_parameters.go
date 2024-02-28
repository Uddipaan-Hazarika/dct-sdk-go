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

// checks if the EnvironmentUpdateParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentUpdateParameters{}

// EnvironmentUpdateParameters struct for EnvironmentUpdateParameters
type EnvironmentUpdateParameters struct {
	// The name of the environment.
	Name *string `json:"name,omitempty"`
	// Id of the connector environment which is used to connect to this source environment.
	StagingEnvironment *string `json:"staging_environment,omitempty"`
	// Address of the cluster. This property can be modified for Windows cluster only.
	ClusterAddress *string `json:"cluster_address,omitempty"`
	// Absolute path to cluster home directory. This parameter is for UNIX cluster environments.
	ClusterHome *string `json:"cluster_home,omitempty"`
	// username of the SAP ASE database.
	AseDbUsername *string `json:"ase_db_username,omitempty"`
	// password of the SAP ASE database.
	AseDbPassword *string `json:"ase_db_password,omitempty"`
	// The name or reference of the vault from which to read the ASE database credentials.
	AseDbVault *string `json:"ase_db_vault,omitempty"`
	// Delphix display name for the vault user
	AseDbVaultUsername *string `json:"ase_db_vault_username,omitempty"`
	// Vault engine name where the credential is stored.
	AseDbHashicorpVaultEngine *string `json:"ase_db_hashicorp_vault_engine,omitempty"`
	// Path in the vault engine where the credential is stored.
	AseDbHashicorpVaultSecretPath *string `json:"ase_db_hashicorp_vault_secret_path,omitempty"`
	// Key for the username in the key-value store.
	AseDbHashicorpVaultUsernameKey *string `json:"ase_db_hashicorp_vault_username_key,omitempty"`
	// Key for the password in the key-value store.
	AseDbHashicorpVaultSecretKey *string `json:"ase_db_hashicorp_vault_secret_key,omitempty"`
	// Query to find a credential in the CyberArk vault.
	AseDbCyberarkVaultQueryString *string `json:"ase_db_cyberark_vault_query_string,omitempty"`
	// Whether to use kerberos authentication for ASE DB discovery.
	AseDbUseKerberosAuthentication *bool `json:"ase_db_use_kerberos_authentication,omitempty"`
	// The environment description.
	Description *string `json:"description,omitempty"`
}

// NewEnvironmentUpdateParameters instantiates a new EnvironmentUpdateParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentUpdateParameters() *EnvironmentUpdateParameters {
	this := EnvironmentUpdateParameters{}
	return &this
}

// NewEnvironmentUpdateParametersWithDefaults instantiates a new EnvironmentUpdateParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentUpdateParametersWithDefaults() *EnvironmentUpdateParameters {
	this := EnvironmentUpdateParameters{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EnvironmentUpdateParameters) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUpdateParameters) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EnvironmentUpdateParameters) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EnvironmentUpdateParameters) SetName(v string) {
	o.Name = &v
}

// GetStagingEnvironment returns the StagingEnvironment field value if set, zero value otherwise.
func (o *EnvironmentUpdateParameters) GetStagingEnvironment() string {
	if o == nil || IsNil(o.StagingEnvironment) {
		var ret string
		return ret
	}
	return *o.StagingEnvironment
}

// GetStagingEnvironmentOk returns a tuple with the StagingEnvironment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUpdateParameters) GetStagingEnvironmentOk() (*string, bool) {
	if o == nil || IsNil(o.StagingEnvironment) {
		return nil, false
	}
	return o.StagingEnvironment, true
}

// HasStagingEnvironment returns a boolean if a field has been set.
func (o *EnvironmentUpdateParameters) HasStagingEnvironment() bool {
	if o != nil && !IsNil(o.StagingEnvironment) {
		return true
	}

	return false
}

// SetStagingEnvironment gets a reference to the given string and assigns it to the StagingEnvironment field.
func (o *EnvironmentUpdateParameters) SetStagingEnvironment(v string) {
	o.StagingEnvironment = &v
}

// GetClusterAddress returns the ClusterAddress field value if set, zero value otherwise.
func (o *EnvironmentUpdateParameters) GetClusterAddress() string {
	if o == nil || IsNil(o.ClusterAddress) {
		var ret string
		return ret
	}
	return *o.ClusterAddress
}

// GetClusterAddressOk returns a tuple with the ClusterAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUpdateParameters) GetClusterAddressOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterAddress) {
		return nil, false
	}
	return o.ClusterAddress, true
}

// HasClusterAddress returns a boolean if a field has been set.
func (o *EnvironmentUpdateParameters) HasClusterAddress() bool {
	if o != nil && !IsNil(o.ClusterAddress) {
		return true
	}

	return false
}

// SetClusterAddress gets a reference to the given string and assigns it to the ClusterAddress field.
func (o *EnvironmentUpdateParameters) SetClusterAddress(v string) {
	o.ClusterAddress = &v
}

// GetClusterHome returns the ClusterHome field value if set, zero value otherwise.
func (o *EnvironmentUpdateParameters) GetClusterHome() string {
	if o == nil || IsNil(o.ClusterHome) {
		var ret string
		return ret
	}
	return *o.ClusterHome
}

// GetClusterHomeOk returns a tuple with the ClusterHome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUpdateParameters) GetClusterHomeOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterHome) {
		return nil, false
	}
	return o.ClusterHome, true
}

// HasClusterHome returns a boolean if a field has been set.
func (o *EnvironmentUpdateParameters) HasClusterHome() bool {
	if o != nil && !IsNil(o.ClusterHome) {
		return true
	}

	return false
}

// SetClusterHome gets a reference to the given string and assigns it to the ClusterHome field.
func (o *EnvironmentUpdateParameters) SetClusterHome(v string) {
	o.ClusterHome = &v
}

// GetAseDbUsername returns the AseDbUsername field value if set, zero value otherwise.
func (o *EnvironmentUpdateParameters) GetAseDbUsername() string {
	if o == nil || IsNil(o.AseDbUsername) {
		var ret string
		return ret
	}
	return *o.AseDbUsername
}

// GetAseDbUsernameOk returns a tuple with the AseDbUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUpdateParameters) GetAseDbUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.AseDbUsername) {
		return nil, false
	}
	return o.AseDbUsername, true
}

// HasAseDbUsername returns a boolean if a field has been set.
func (o *EnvironmentUpdateParameters) HasAseDbUsername() bool {
	if o != nil && !IsNil(o.AseDbUsername) {
		return true
	}

	return false
}

// SetAseDbUsername gets a reference to the given string and assigns it to the AseDbUsername field.
func (o *EnvironmentUpdateParameters) SetAseDbUsername(v string) {
	o.AseDbUsername = &v
}

// GetAseDbPassword returns the AseDbPassword field value if set, zero value otherwise.
func (o *EnvironmentUpdateParameters) GetAseDbPassword() string {
	if o == nil || IsNil(o.AseDbPassword) {
		var ret string
		return ret
	}
	return *o.AseDbPassword
}

// GetAseDbPasswordOk returns a tuple with the AseDbPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUpdateParameters) GetAseDbPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.AseDbPassword) {
		return nil, false
	}
	return o.AseDbPassword, true
}

// HasAseDbPassword returns a boolean if a field has been set.
func (o *EnvironmentUpdateParameters) HasAseDbPassword() bool {
	if o != nil && !IsNil(o.AseDbPassword) {
		return true
	}

	return false
}

// SetAseDbPassword gets a reference to the given string and assigns it to the AseDbPassword field.
func (o *EnvironmentUpdateParameters) SetAseDbPassword(v string) {
	o.AseDbPassword = &v
}

// GetAseDbVault returns the AseDbVault field value if set, zero value otherwise.
func (o *EnvironmentUpdateParameters) GetAseDbVault() string {
	if o == nil || IsNil(o.AseDbVault) {
		var ret string
		return ret
	}
	return *o.AseDbVault
}

// GetAseDbVaultOk returns a tuple with the AseDbVault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUpdateParameters) GetAseDbVaultOk() (*string, bool) {
	if o == nil || IsNil(o.AseDbVault) {
		return nil, false
	}
	return o.AseDbVault, true
}

// HasAseDbVault returns a boolean if a field has been set.
func (o *EnvironmentUpdateParameters) HasAseDbVault() bool {
	if o != nil && !IsNil(o.AseDbVault) {
		return true
	}

	return false
}

// SetAseDbVault gets a reference to the given string and assigns it to the AseDbVault field.
func (o *EnvironmentUpdateParameters) SetAseDbVault(v string) {
	o.AseDbVault = &v
}

// GetAseDbVaultUsername returns the AseDbVaultUsername field value if set, zero value otherwise.
func (o *EnvironmentUpdateParameters) GetAseDbVaultUsername() string {
	if o == nil || IsNil(o.AseDbVaultUsername) {
		var ret string
		return ret
	}
	return *o.AseDbVaultUsername
}

// GetAseDbVaultUsernameOk returns a tuple with the AseDbVaultUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUpdateParameters) GetAseDbVaultUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.AseDbVaultUsername) {
		return nil, false
	}
	return o.AseDbVaultUsername, true
}

// HasAseDbVaultUsername returns a boolean if a field has been set.
func (o *EnvironmentUpdateParameters) HasAseDbVaultUsername() bool {
	if o != nil && !IsNil(o.AseDbVaultUsername) {
		return true
	}

	return false
}

// SetAseDbVaultUsername gets a reference to the given string and assigns it to the AseDbVaultUsername field.
func (o *EnvironmentUpdateParameters) SetAseDbVaultUsername(v string) {
	o.AseDbVaultUsername = &v
}

// GetAseDbHashicorpVaultEngine returns the AseDbHashicorpVaultEngine field value if set, zero value otherwise.
func (o *EnvironmentUpdateParameters) GetAseDbHashicorpVaultEngine() string {
	if o == nil || IsNil(o.AseDbHashicorpVaultEngine) {
		var ret string
		return ret
	}
	return *o.AseDbHashicorpVaultEngine
}

// GetAseDbHashicorpVaultEngineOk returns a tuple with the AseDbHashicorpVaultEngine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUpdateParameters) GetAseDbHashicorpVaultEngineOk() (*string, bool) {
	if o == nil || IsNil(o.AseDbHashicorpVaultEngine) {
		return nil, false
	}
	return o.AseDbHashicorpVaultEngine, true
}

// HasAseDbHashicorpVaultEngine returns a boolean if a field has been set.
func (o *EnvironmentUpdateParameters) HasAseDbHashicorpVaultEngine() bool {
	if o != nil && !IsNil(o.AseDbHashicorpVaultEngine) {
		return true
	}

	return false
}

// SetAseDbHashicorpVaultEngine gets a reference to the given string and assigns it to the AseDbHashicorpVaultEngine field.
func (o *EnvironmentUpdateParameters) SetAseDbHashicorpVaultEngine(v string) {
	o.AseDbHashicorpVaultEngine = &v
}

// GetAseDbHashicorpVaultSecretPath returns the AseDbHashicorpVaultSecretPath field value if set, zero value otherwise.
func (o *EnvironmentUpdateParameters) GetAseDbHashicorpVaultSecretPath() string {
	if o == nil || IsNil(o.AseDbHashicorpVaultSecretPath) {
		var ret string
		return ret
	}
	return *o.AseDbHashicorpVaultSecretPath
}

// GetAseDbHashicorpVaultSecretPathOk returns a tuple with the AseDbHashicorpVaultSecretPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUpdateParameters) GetAseDbHashicorpVaultSecretPathOk() (*string, bool) {
	if o == nil || IsNil(o.AseDbHashicorpVaultSecretPath) {
		return nil, false
	}
	return o.AseDbHashicorpVaultSecretPath, true
}

// HasAseDbHashicorpVaultSecretPath returns a boolean if a field has been set.
func (o *EnvironmentUpdateParameters) HasAseDbHashicorpVaultSecretPath() bool {
	if o != nil && !IsNil(o.AseDbHashicorpVaultSecretPath) {
		return true
	}

	return false
}

// SetAseDbHashicorpVaultSecretPath gets a reference to the given string and assigns it to the AseDbHashicorpVaultSecretPath field.
func (o *EnvironmentUpdateParameters) SetAseDbHashicorpVaultSecretPath(v string) {
	o.AseDbHashicorpVaultSecretPath = &v
}

// GetAseDbHashicorpVaultUsernameKey returns the AseDbHashicorpVaultUsernameKey field value if set, zero value otherwise.
func (o *EnvironmentUpdateParameters) GetAseDbHashicorpVaultUsernameKey() string {
	if o == nil || IsNil(o.AseDbHashicorpVaultUsernameKey) {
		var ret string
		return ret
	}
	return *o.AseDbHashicorpVaultUsernameKey
}

// GetAseDbHashicorpVaultUsernameKeyOk returns a tuple with the AseDbHashicorpVaultUsernameKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUpdateParameters) GetAseDbHashicorpVaultUsernameKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AseDbHashicorpVaultUsernameKey) {
		return nil, false
	}
	return o.AseDbHashicorpVaultUsernameKey, true
}

// HasAseDbHashicorpVaultUsernameKey returns a boolean if a field has been set.
func (o *EnvironmentUpdateParameters) HasAseDbHashicorpVaultUsernameKey() bool {
	if o != nil && !IsNil(o.AseDbHashicorpVaultUsernameKey) {
		return true
	}

	return false
}

// SetAseDbHashicorpVaultUsernameKey gets a reference to the given string and assigns it to the AseDbHashicorpVaultUsernameKey field.
func (o *EnvironmentUpdateParameters) SetAseDbHashicorpVaultUsernameKey(v string) {
	o.AseDbHashicorpVaultUsernameKey = &v
}

// GetAseDbHashicorpVaultSecretKey returns the AseDbHashicorpVaultSecretKey field value if set, zero value otherwise.
func (o *EnvironmentUpdateParameters) GetAseDbHashicorpVaultSecretKey() string {
	if o == nil || IsNil(o.AseDbHashicorpVaultSecretKey) {
		var ret string
		return ret
	}
	return *o.AseDbHashicorpVaultSecretKey
}

// GetAseDbHashicorpVaultSecretKeyOk returns a tuple with the AseDbHashicorpVaultSecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUpdateParameters) GetAseDbHashicorpVaultSecretKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AseDbHashicorpVaultSecretKey) {
		return nil, false
	}
	return o.AseDbHashicorpVaultSecretKey, true
}

// HasAseDbHashicorpVaultSecretKey returns a boolean if a field has been set.
func (o *EnvironmentUpdateParameters) HasAseDbHashicorpVaultSecretKey() bool {
	if o != nil && !IsNil(o.AseDbHashicorpVaultSecretKey) {
		return true
	}

	return false
}

// SetAseDbHashicorpVaultSecretKey gets a reference to the given string and assigns it to the AseDbHashicorpVaultSecretKey field.
func (o *EnvironmentUpdateParameters) SetAseDbHashicorpVaultSecretKey(v string) {
	o.AseDbHashicorpVaultSecretKey = &v
}

// GetAseDbCyberarkVaultQueryString returns the AseDbCyberarkVaultQueryString field value if set, zero value otherwise.
func (o *EnvironmentUpdateParameters) GetAseDbCyberarkVaultQueryString() string {
	if o == nil || IsNil(o.AseDbCyberarkVaultQueryString) {
		var ret string
		return ret
	}
	return *o.AseDbCyberarkVaultQueryString
}

// GetAseDbCyberarkVaultQueryStringOk returns a tuple with the AseDbCyberarkVaultQueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUpdateParameters) GetAseDbCyberarkVaultQueryStringOk() (*string, bool) {
	if o == nil || IsNil(o.AseDbCyberarkVaultQueryString) {
		return nil, false
	}
	return o.AseDbCyberarkVaultQueryString, true
}

// HasAseDbCyberarkVaultQueryString returns a boolean if a field has been set.
func (o *EnvironmentUpdateParameters) HasAseDbCyberarkVaultQueryString() bool {
	if o != nil && !IsNil(o.AseDbCyberarkVaultQueryString) {
		return true
	}

	return false
}

// SetAseDbCyberarkVaultQueryString gets a reference to the given string and assigns it to the AseDbCyberarkVaultQueryString field.
func (o *EnvironmentUpdateParameters) SetAseDbCyberarkVaultQueryString(v string) {
	o.AseDbCyberarkVaultQueryString = &v
}

// GetAseDbUseKerberosAuthentication returns the AseDbUseKerberosAuthentication field value if set, zero value otherwise.
func (o *EnvironmentUpdateParameters) GetAseDbUseKerberosAuthentication() bool {
	if o == nil || IsNil(o.AseDbUseKerberosAuthentication) {
		var ret bool
		return ret
	}
	return *o.AseDbUseKerberosAuthentication
}

// GetAseDbUseKerberosAuthenticationOk returns a tuple with the AseDbUseKerberosAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUpdateParameters) GetAseDbUseKerberosAuthenticationOk() (*bool, bool) {
	if o == nil || IsNil(o.AseDbUseKerberosAuthentication) {
		return nil, false
	}
	return o.AseDbUseKerberosAuthentication, true
}

// HasAseDbUseKerberosAuthentication returns a boolean if a field has been set.
func (o *EnvironmentUpdateParameters) HasAseDbUseKerberosAuthentication() bool {
	if o != nil && !IsNil(o.AseDbUseKerberosAuthentication) {
		return true
	}

	return false
}

// SetAseDbUseKerberosAuthentication gets a reference to the given bool and assigns it to the AseDbUseKerberosAuthentication field.
func (o *EnvironmentUpdateParameters) SetAseDbUseKerberosAuthentication(v bool) {
	o.AseDbUseKerberosAuthentication = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EnvironmentUpdateParameters) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUpdateParameters) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EnvironmentUpdateParameters) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EnvironmentUpdateParameters) SetDescription(v string) {
	o.Description = &v
}

func (o EnvironmentUpdateParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentUpdateParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.StagingEnvironment) {
		toSerialize["staging_environment"] = o.StagingEnvironment
	}
	if !IsNil(o.ClusterAddress) {
		toSerialize["cluster_address"] = o.ClusterAddress
	}
	if !IsNil(o.ClusterHome) {
		toSerialize["cluster_home"] = o.ClusterHome
	}
	if !IsNil(o.AseDbUsername) {
		toSerialize["ase_db_username"] = o.AseDbUsername
	}
	if !IsNil(o.AseDbPassword) {
		toSerialize["ase_db_password"] = o.AseDbPassword
	}
	if !IsNil(o.AseDbVault) {
		toSerialize["ase_db_vault"] = o.AseDbVault
	}
	if !IsNil(o.AseDbVaultUsername) {
		toSerialize["ase_db_vault_username"] = o.AseDbVaultUsername
	}
	if !IsNil(o.AseDbHashicorpVaultEngine) {
		toSerialize["ase_db_hashicorp_vault_engine"] = o.AseDbHashicorpVaultEngine
	}
	if !IsNil(o.AseDbHashicorpVaultSecretPath) {
		toSerialize["ase_db_hashicorp_vault_secret_path"] = o.AseDbHashicorpVaultSecretPath
	}
	if !IsNil(o.AseDbHashicorpVaultUsernameKey) {
		toSerialize["ase_db_hashicorp_vault_username_key"] = o.AseDbHashicorpVaultUsernameKey
	}
	if !IsNil(o.AseDbHashicorpVaultSecretKey) {
		toSerialize["ase_db_hashicorp_vault_secret_key"] = o.AseDbHashicorpVaultSecretKey
	}
	if !IsNil(o.AseDbCyberarkVaultQueryString) {
		toSerialize["ase_db_cyberark_vault_query_string"] = o.AseDbCyberarkVaultQueryString
	}
	if !IsNil(o.AseDbUseKerberosAuthentication) {
		toSerialize["ase_db_use_kerberos_authentication"] = o.AseDbUseKerberosAuthentication
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableEnvironmentUpdateParameters struct {
	value *EnvironmentUpdateParameters
	isSet bool
}

func (v NullableEnvironmentUpdateParameters) Get() *EnvironmentUpdateParameters {
	return v.value
}

func (v *NullableEnvironmentUpdateParameters) Set(val *EnvironmentUpdateParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentUpdateParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentUpdateParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentUpdateParameters(val *EnvironmentUpdateParameters) *NullableEnvironmentUpdateParameters {
	return &NullableEnvironmentUpdateParameters{value: val, isSet: true}
}

func (v NullableEnvironmentUpdateParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentUpdateParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


