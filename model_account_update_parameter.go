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

// checks if the AccountUpdateParameter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountUpdateParameter{}

// AccountUpdateParameter struct for AccountUpdateParameter
type AccountUpdateParameter struct {
	// Numeric ID of the Account.
	Id *int64 `json:"id,omitempty"`
	// The unique ID which is used to identify the identity of an API request. The web server (nginx) configuration must be configured so as to include the external ID as the value of the X_CLIENT_ID HTTP request header when requests are proxied. For OAuth2/JWT based authentication, this typically corresponds to a value extracted from the JWT, uniquely identifying the Account.
	ApiClientId *string `json:"api_client_id,omitempty"`
	// An optional first name for the Account.
	FirstName *string `json:"first_name,omitempty"`
	// An optional last name for the Account.
	LastName *string `json:"last_name,omitempty"`
	// An optional email for the Account.
	Email *string `json:"email,omitempty"`
	// The username for username/password authentication. This can also be used to provide an optional logical name for the Account.
	Username *string `json:"username,omitempty"`
	// This value will be used for linking this account to an LDAP user when authenticated with the same LDAP principal. When accounts authenticate with LDAP, an LDAP principal value is calculated based on the username, msad_domain_name, search_base and username_pattern.
	LdapPrincipal *string `json:"ldap_principal,omitempty"`
}

// NewAccountUpdateParameter instantiates a new AccountUpdateParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountUpdateParameter() *AccountUpdateParameter {
	this := AccountUpdateParameter{}
	return &this
}

// NewAccountUpdateParameterWithDefaults instantiates a new AccountUpdateParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountUpdateParameterWithDefaults() *AccountUpdateParameter {
	this := AccountUpdateParameter{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccountUpdateParameter) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdateParameter) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccountUpdateParameter) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *AccountUpdateParameter) SetId(v int64) {
	o.Id = &v
}

// GetApiClientId returns the ApiClientId field value if set, zero value otherwise.
func (o *AccountUpdateParameter) GetApiClientId() string {
	if o == nil || IsNil(o.ApiClientId) {
		var ret string
		return ret
	}
	return *o.ApiClientId
}

// GetApiClientIdOk returns a tuple with the ApiClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdateParameter) GetApiClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApiClientId) {
		return nil, false
	}
	return o.ApiClientId, true
}

// HasApiClientId returns a boolean if a field has been set.
func (o *AccountUpdateParameter) HasApiClientId() bool {
	if o != nil && !IsNil(o.ApiClientId) {
		return true
	}

	return false
}

// SetApiClientId gets a reference to the given string and assigns it to the ApiClientId field.
func (o *AccountUpdateParameter) SetApiClientId(v string) {
	o.ApiClientId = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *AccountUpdateParameter) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdateParameter) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *AccountUpdateParameter) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *AccountUpdateParameter) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *AccountUpdateParameter) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdateParameter) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *AccountUpdateParameter) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *AccountUpdateParameter) SetLastName(v string) {
	o.LastName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *AccountUpdateParameter) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdateParameter) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *AccountUpdateParameter) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *AccountUpdateParameter) SetEmail(v string) {
	o.Email = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *AccountUpdateParameter) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdateParameter) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *AccountUpdateParameter) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *AccountUpdateParameter) SetUsername(v string) {
	o.Username = &v
}

// GetLdapPrincipal returns the LdapPrincipal field value if set, zero value otherwise.
func (o *AccountUpdateParameter) GetLdapPrincipal() string {
	if o == nil || IsNil(o.LdapPrincipal) {
		var ret string
		return ret
	}
	return *o.LdapPrincipal
}

// GetLdapPrincipalOk returns a tuple with the LdapPrincipal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdateParameter) GetLdapPrincipalOk() (*string, bool) {
	if o == nil || IsNil(o.LdapPrincipal) {
		return nil, false
	}
	return o.LdapPrincipal, true
}

// HasLdapPrincipal returns a boolean if a field has been set.
func (o *AccountUpdateParameter) HasLdapPrincipal() bool {
	if o != nil && !IsNil(o.LdapPrincipal) {
		return true
	}

	return false
}

// SetLdapPrincipal gets a reference to the given string and assigns it to the LdapPrincipal field.
func (o *AccountUpdateParameter) SetLdapPrincipal(v string) {
	o.LdapPrincipal = &v
}

func (o AccountUpdateParameter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountUpdateParameter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	if !IsNil(o.ApiClientId) {
		toSerialize["api_client_id"] = o.ApiClientId
	}
	if !IsNil(o.FirstName) {
		toSerialize["first_name"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.LdapPrincipal) {
		toSerialize["ldap_principal"] = o.LdapPrincipal
	}
	return toSerialize, nil
}

type NullableAccountUpdateParameter struct {
	value *AccountUpdateParameter
	isSet bool
}

func (v NullableAccountUpdateParameter) Get() *AccountUpdateParameter {
	return v.value
}

func (v *NullableAccountUpdateParameter) Set(val *AccountUpdateParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountUpdateParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountUpdateParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountUpdateParameter(val *AccountUpdateParameter) *NullableAccountUpdateParameter {
	return &NullableAccountUpdateParameter{value: val, isSet: true}
}

func (v NullableAccountUpdateParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountUpdateParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


