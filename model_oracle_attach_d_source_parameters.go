/*
Delphix DCT API

Delphix DCT API

API version: 3.16.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the OracleAttachDSourceParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OracleAttachDSourceParameters{}

// OracleAttachDSourceParameters struct for OracleAttachDSourceParameters
type OracleAttachDSourceParameters struct {
	// Id of the source to attach.
	SourceId string `json:"source_id"`
	// Boolean value indicates whether LEVEL-based incremental backups can be used on the source database.
	BackupLevelEnabled *bool `json:"backup_level_enabled,omitempty"`
	// Bandwidth limit (MB/s) for SnapSync and LogSync network traffic. A value of 0 means no limit.
	BandwidthLimit *int32 `json:"bandwidth_limit,omitempty"`
	// True if extended block checking should be used for this linked database.
	CheckLogical *bool `json:"check_logical,omitempty"`
	// True if SnapSync data from the source should be compressed over the network. Enabling this feature will reduce network bandwidth consumption and may significantly improve throughput, especially over slow network.
	CompressedLinkingEnabled *bool `json:"compressed_linking_enabled,omitempty"`
	// True if two SnapSyncs should be performed in immediate succession to reduce the number of logs required to provision the snapshot. This may significantly reduce the time necessary to provision from a snapshot.
	DoubleSync *bool `json:"double_sync,omitempty"`
	// True if SnapSync data from the source should be retrieved through an encrypted connection. Enabling this feature can decrease the performance of SnapSync from the source but has no impact on the performance of VDBs created from the retrieved data.
	EncryptedLinkingEnabled *bool `json:"encrypted_linking_enabled,omitempty"`
	// Reference to the user that should be used in the host.
	EnvironmentUser *string `json:"environment_user,omitempty"`
	// External file path.
	ExternalFilePath *string `json:"external_file_path,omitempty"`
	// Number of data files to include in each RMAN backup set.
	FilesPerSet *int32 `json:"files_per_set,omitempty"`
	// If true, attach will succeed even if the resetlogs of the new database does not match the resetlogs information of the original database.
	Force *bool `json:"force,omitempty"`
	// True if initial load should be done immediately.
	LinkNow *bool `json:"link_now,omitempty"`
	// Total number of transport connections to use during SnapSync.
	NumberOfConnections *int32 `json:"number_of_connections,omitempty"`
	// Operations to perform after syncing a created dSource and before running the LogSync.
	Operations []SourceOperation `json:"operations,omitempty"`
	// The database fallback username. Optional if bequeath connections are enabled (to be used in case of bequeath connection failures). Only required for username-password auth.
	OracleFallbackUser *string `json:"oracle_fallback_user,omitempty"`
	// Password for fallback username.
	OracleFallbackCredentials *string `json:"oracle_fallback_credentials,omitempty"`
	// Number of parallel channels to use.
	RmanChannels *int32 `json:"rman_channels,omitempty"`
}

type _OracleAttachDSourceParameters OracleAttachDSourceParameters

// NewOracleAttachDSourceParameters instantiates a new OracleAttachDSourceParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOracleAttachDSourceParameters(sourceId string) *OracleAttachDSourceParameters {
	this := OracleAttachDSourceParameters{}
	this.SourceId = sourceId
	var bandwidthLimit int32 = 0
	this.BandwidthLimit = &bandwidthLimit
	var checkLogical bool = false
	this.CheckLogical = &checkLogical
	var compressedLinkingEnabled bool = true
	this.CompressedLinkingEnabled = &compressedLinkingEnabled
	var doubleSync bool = false
	this.DoubleSync = &doubleSync
	var encryptedLinkingEnabled bool = false
	this.EncryptedLinkingEnabled = &encryptedLinkingEnabled
	var filesPerSet int32 = 5
	this.FilesPerSet = &filesPerSet
	var force bool = false
	this.Force = &force
	var linkNow bool = false
	this.LinkNow = &linkNow
	var numberOfConnections int32 = 1
	this.NumberOfConnections = &numberOfConnections
	var rmanChannels int32 = 2
	this.RmanChannels = &rmanChannels
	return &this
}

// NewOracleAttachDSourceParametersWithDefaults instantiates a new OracleAttachDSourceParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOracleAttachDSourceParametersWithDefaults() *OracleAttachDSourceParameters {
	this := OracleAttachDSourceParameters{}
	var bandwidthLimit int32 = 0
	this.BandwidthLimit = &bandwidthLimit
	var checkLogical bool = false
	this.CheckLogical = &checkLogical
	var compressedLinkingEnabled bool = true
	this.CompressedLinkingEnabled = &compressedLinkingEnabled
	var doubleSync bool = false
	this.DoubleSync = &doubleSync
	var encryptedLinkingEnabled bool = false
	this.EncryptedLinkingEnabled = &encryptedLinkingEnabled
	var filesPerSet int32 = 5
	this.FilesPerSet = &filesPerSet
	var force bool = false
	this.Force = &force
	var linkNow bool = false
	this.LinkNow = &linkNow
	var numberOfConnections int32 = 1
	this.NumberOfConnections = &numberOfConnections
	var rmanChannels int32 = 2
	this.RmanChannels = &rmanChannels
	return &this
}

// GetSourceId returns the SourceId field value
func (o *OracleAttachDSourceParameters) GetSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value
// and a boolean to check if the value has been set.
func (o *OracleAttachDSourceParameters) GetSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceId, true
}

// SetSourceId sets field value
func (o *OracleAttachDSourceParameters) SetSourceId(v string) {
	o.SourceId = v
}

// GetBackupLevelEnabled returns the BackupLevelEnabled field value if set, zero value otherwise.
func (o *OracleAttachDSourceParameters) GetBackupLevelEnabled() bool {
	if o == nil || IsNil(o.BackupLevelEnabled) {
		var ret bool
		return ret
	}
	return *o.BackupLevelEnabled
}

// GetBackupLevelEnabledOk returns a tuple with the BackupLevelEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleAttachDSourceParameters) GetBackupLevelEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.BackupLevelEnabled) {
		return nil, false
	}
	return o.BackupLevelEnabled, true
}

// HasBackupLevelEnabled returns a boolean if a field has been set.
func (o *OracleAttachDSourceParameters) HasBackupLevelEnabled() bool {
	if o != nil && !IsNil(o.BackupLevelEnabled) {
		return true
	}

	return false
}

// SetBackupLevelEnabled gets a reference to the given bool and assigns it to the BackupLevelEnabled field.
func (o *OracleAttachDSourceParameters) SetBackupLevelEnabled(v bool) {
	o.BackupLevelEnabled = &v
}

// GetBandwidthLimit returns the BandwidthLimit field value if set, zero value otherwise.
func (o *OracleAttachDSourceParameters) GetBandwidthLimit() int32 {
	if o == nil || IsNil(o.BandwidthLimit) {
		var ret int32
		return ret
	}
	return *o.BandwidthLimit
}

// GetBandwidthLimitOk returns a tuple with the BandwidthLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleAttachDSourceParameters) GetBandwidthLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.BandwidthLimit) {
		return nil, false
	}
	return o.BandwidthLimit, true
}

// HasBandwidthLimit returns a boolean if a field has been set.
func (o *OracleAttachDSourceParameters) HasBandwidthLimit() bool {
	if o != nil && !IsNil(o.BandwidthLimit) {
		return true
	}

	return false
}

// SetBandwidthLimit gets a reference to the given int32 and assigns it to the BandwidthLimit field.
func (o *OracleAttachDSourceParameters) SetBandwidthLimit(v int32) {
	o.BandwidthLimit = &v
}

// GetCheckLogical returns the CheckLogical field value if set, zero value otherwise.
func (o *OracleAttachDSourceParameters) GetCheckLogical() bool {
	if o == nil || IsNil(o.CheckLogical) {
		var ret bool
		return ret
	}
	return *o.CheckLogical
}

// GetCheckLogicalOk returns a tuple with the CheckLogical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleAttachDSourceParameters) GetCheckLogicalOk() (*bool, bool) {
	if o == nil || IsNil(o.CheckLogical) {
		return nil, false
	}
	return o.CheckLogical, true
}

// HasCheckLogical returns a boolean if a field has been set.
func (o *OracleAttachDSourceParameters) HasCheckLogical() bool {
	if o != nil && !IsNil(o.CheckLogical) {
		return true
	}

	return false
}

// SetCheckLogical gets a reference to the given bool and assigns it to the CheckLogical field.
func (o *OracleAttachDSourceParameters) SetCheckLogical(v bool) {
	o.CheckLogical = &v
}

// GetCompressedLinkingEnabled returns the CompressedLinkingEnabled field value if set, zero value otherwise.
func (o *OracleAttachDSourceParameters) GetCompressedLinkingEnabled() bool {
	if o == nil || IsNil(o.CompressedLinkingEnabled) {
		var ret bool
		return ret
	}
	return *o.CompressedLinkingEnabled
}

// GetCompressedLinkingEnabledOk returns a tuple with the CompressedLinkingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleAttachDSourceParameters) GetCompressedLinkingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.CompressedLinkingEnabled) {
		return nil, false
	}
	return o.CompressedLinkingEnabled, true
}

// HasCompressedLinkingEnabled returns a boolean if a field has been set.
func (o *OracleAttachDSourceParameters) HasCompressedLinkingEnabled() bool {
	if o != nil && !IsNil(o.CompressedLinkingEnabled) {
		return true
	}

	return false
}

// SetCompressedLinkingEnabled gets a reference to the given bool and assigns it to the CompressedLinkingEnabled field.
func (o *OracleAttachDSourceParameters) SetCompressedLinkingEnabled(v bool) {
	o.CompressedLinkingEnabled = &v
}

// GetDoubleSync returns the DoubleSync field value if set, zero value otherwise.
func (o *OracleAttachDSourceParameters) GetDoubleSync() bool {
	if o == nil || IsNil(o.DoubleSync) {
		var ret bool
		return ret
	}
	return *o.DoubleSync
}

// GetDoubleSyncOk returns a tuple with the DoubleSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleAttachDSourceParameters) GetDoubleSyncOk() (*bool, bool) {
	if o == nil || IsNil(o.DoubleSync) {
		return nil, false
	}
	return o.DoubleSync, true
}

// HasDoubleSync returns a boolean if a field has been set.
func (o *OracleAttachDSourceParameters) HasDoubleSync() bool {
	if o != nil && !IsNil(o.DoubleSync) {
		return true
	}

	return false
}

// SetDoubleSync gets a reference to the given bool and assigns it to the DoubleSync field.
func (o *OracleAttachDSourceParameters) SetDoubleSync(v bool) {
	o.DoubleSync = &v
}

// GetEncryptedLinkingEnabled returns the EncryptedLinkingEnabled field value if set, zero value otherwise.
func (o *OracleAttachDSourceParameters) GetEncryptedLinkingEnabled() bool {
	if o == nil || IsNil(o.EncryptedLinkingEnabled) {
		var ret bool
		return ret
	}
	return *o.EncryptedLinkingEnabled
}

// GetEncryptedLinkingEnabledOk returns a tuple with the EncryptedLinkingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleAttachDSourceParameters) GetEncryptedLinkingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.EncryptedLinkingEnabled) {
		return nil, false
	}
	return o.EncryptedLinkingEnabled, true
}

// HasEncryptedLinkingEnabled returns a boolean if a field has been set.
func (o *OracleAttachDSourceParameters) HasEncryptedLinkingEnabled() bool {
	if o != nil && !IsNil(o.EncryptedLinkingEnabled) {
		return true
	}

	return false
}

// SetEncryptedLinkingEnabled gets a reference to the given bool and assigns it to the EncryptedLinkingEnabled field.
func (o *OracleAttachDSourceParameters) SetEncryptedLinkingEnabled(v bool) {
	o.EncryptedLinkingEnabled = &v
}

// GetEnvironmentUser returns the EnvironmentUser field value if set, zero value otherwise.
func (o *OracleAttachDSourceParameters) GetEnvironmentUser() string {
	if o == nil || IsNil(o.EnvironmentUser) {
		var ret string
		return ret
	}
	return *o.EnvironmentUser
}

// GetEnvironmentUserOk returns a tuple with the EnvironmentUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleAttachDSourceParameters) GetEnvironmentUserOk() (*string, bool) {
	if o == nil || IsNil(o.EnvironmentUser) {
		return nil, false
	}
	return o.EnvironmentUser, true
}

// HasEnvironmentUser returns a boolean if a field has been set.
func (o *OracleAttachDSourceParameters) HasEnvironmentUser() bool {
	if o != nil && !IsNil(o.EnvironmentUser) {
		return true
	}

	return false
}

// SetEnvironmentUser gets a reference to the given string and assigns it to the EnvironmentUser field.
func (o *OracleAttachDSourceParameters) SetEnvironmentUser(v string) {
	o.EnvironmentUser = &v
}

// GetExternalFilePath returns the ExternalFilePath field value if set, zero value otherwise.
func (o *OracleAttachDSourceParameters) GetExternalFilePath() string {
	if o == nil || IsNil(o.ExternalFilePath) {
		var ret string
		return ret
	}
	return *o.ExternalFilePath
}

// GetExternalFilePathOk returns a tuple with the ExternalFilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleAttachDSourceParameters) GetExternalFilePathOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalFilePath) {
		return nil, false
	}
	return o.ExternalFilePath, true
}

// HasExternalFilePath returns a boolean if a field has been set.
func (o *OracleAttachDSourceParameters) HasExternalFilePath() bool {
	if o != nil && !IsNil(o.ExternalFilePath) {
		return true
	}

	return false
}

// SetExternalFilePath gets a reference to the given string and assigns it to the ExternalFilePath field.
func (o *OracleAttachDSourceParameters) SetExternalFilePath(v string) {
	o.ExternalFilePath = &v
}

// GetFilesPerSet returns the FilesPerSet field value if set, zero value otherwise.
func (o *OracleAttachDSourceParameters) GetFilesPerSet() int32 {
	if o == nil || IsNil(o.FilesPerSet) {
		var ret int32
		return ret
	}
	return *o.FilesPerSet
}

// GetFilesPerSetOk returns a tuple with the FilesPerSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleAttachDSourceParameters) GetFilesPerSetOk() (*int32, bool) {
	if o == nil || IsNil(o.FilesPerSet) {
		return nil, false
	}
	return o.FilesPerSet, true
}

// HasFilesPerSet returns a boolean if a field has been set.
func (o *OracleAttachDSourceParameters) HasFilesPerSet() bool {
	if o != nil && !IsNil(o.FilesPerSet) {
		return true
	}

	return false
}

// SetFilesPerSet gets a reference to the given int32 and assigns it to the FilesPerSet field.
func (o *OracleAttachDSourceParameters) SetFilesPerSet(v int32) {
	o.FilesPerSet = &v
}

// GetForce returns the Force field value if set, zero value otherwise.
func (o *OracleAttachDSourceParameters) GetForce() bool {
	if o == nil || IsNil(o.Force) {
		var ret bool
		return ret
	}
	return *o.Force
}

// GetForceOk returns a tuple with the Force field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleAttachDSourceParameters) GetForceOk() (*bool, bool) {
	if o == nil || IsNil(o.Force) {
		return nil, false
	}
	return o.Force, true
}

// HasForce returns a boolean if a field has been set.
func (o *OracleAttachDSourceParameters) HasForce() bool {
	if o != nil && !IsNil(o.Force) {
		return true
	}

	return false
}

// SetForce gets a reference to the given bool and assigns it to the Force field.
func (o *OracleAttachDSourceParameters) SetForce(v bool) {
	o.Force = &v
}

// GetLinkNow returns the LinkNow field value if set, zero value otherwise.
func (o *OracleAttachDSourceParameters) GetLinkNow() bool {
	if o == nil || IsNil(o.LinkNow) {
		var ret bool
		return ret
	}
	return *o.LinkNow
}

// GetLinkNowOk returns a tuple with the LinkNow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleAttachDSourceParameters) GetLinkNowOk() (*bool, bool) {
	if o == nil || IsNil(o.LinkNow) {
		return nil, false
	}
	return o.LinkNow, true
}

// HasLinkNow returns a boolean if a field has been set.
func (o *OracleAttachDSourceParameters) HasLinkNow() bool {
	if o != nil && !IsNil(o.LinkNow) {
		return true
	}

	return false
}

// SetLinkNow gets a reference to the given bool and assigns it to the LinkNow field.
func (o *OracleAttachDSourceParameters) SetLinkNow(v bool) {
	o.LinkNow = &v
}

// GetNumberOfConnections returns the NumberOfConnections field value if set, zero value otherwise.
func (o *OracleAttachDSourceParameters) GetNumberOfConnections() int32 {
	if o == nil || IsNil(o.NumberOfConnections) {
		var ret int32
		return ret
	}
	return *o.NumberOfConnections
}

// GetNumberOfConnectionsOk returns a tuple with the NumberOfConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleAttachDSourceParameters) GetNumberOfConnectionsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfConnections) {
		return nil, false
	}
	return o.NumberOfConnections, true
}

// HasNumberOfConnections returns a boolean if a field has been set.
func (o *OracleAttachDSourceParameters) HasNumberOfConnections() bool {
	if o != nil && !IsNil(o.NumberOfConnections) {
		return true
	}

	return false
}

// SetNumberOfConnections gets a reference to the given int32 and assigns it to the NumberOfConnections field.
func (o *OracleAttachDSourceParameters) SetNumberOfConnections(v int32) {
	o.NumberOfConnections = &v
}

// GetOperations returns the Operations field value if set, zero value otherwise.
func (o *OracleAttachDSourceParameters) GetOperations() []SourceOperation {
	if o == nil || IsNil(o.Operations) {
		var ret []SourceOperation
		return ret
	}
	return o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleAttachDSourceParameters) GetOperationsOk() ([]SourceOperation, bool) {
	if o == nil || IsNil(o.Operations) {
		return nil, false
	}
	return o.Operations, true
}

// HasOperations returns a boolean if a field has been set.
func (o *OracleAttachDSourceParameters) HasOperations() bool {
	if o != nil && !IsNil(o.Operations) {
		return true
	}

	return false
}

// SetOperations gets a reference to the given []SourceOperation and assigns it to the Operations field.
func (o *OracleAttachDSourceParameters) SetOperations(v []SourceOperation) {
	o.Operations = v
}

// GetOracleFallbackUser returns the OracleFallbackUser field value if set, zero value otherwise.
func (o *OracleAttachDSourceParameters) GetOracleFallbackUser() string {
	if o == nil || IsNil(o.OracleFallbackUser) {
		var ret string
		return ret
	}
	return *o.OracleFallbackUser
}

// GetOracleFallbackUserOk returns a tuple with the OracleFallbackUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleAttachDSourceParameters) GetOracleFallbackUserOk() (*string, bool) {
	if o == nil || IsNil(o.OracleFallbackUser) {
		return nil, false
	}
	return o.OracleFallbackUser, true
}

// HasOracleFallbackUser returns a boolean if a field has been set.
func (o *OracleAttachDSourceParameters) HasOracleFallbackUser() bool {
	if o != nil && !IsNil(o.OracleFallbackUser) {
		return true
	}

	return false
}

// SetOracleFallbackUser gets a reference to the given string and assigns it to the OracleFallbackUser field.
func (o *OracleAttachDSourceParameters) SetOracleFallbackUser(v string) {
	o.OracleFallbackUser = &v
}

// GetOracleFallbackCredentials returns the OracleFallbackCredentials field value if set, zero value otherwise.
func (o *OracleAttachDSourceParameters) GetOracleFallbackCredentials() string {
	if o == nil || IsNil(o.OracleFallbackCredentials) {
		var ret string
		return ret
	}
	return *o.OracleFallbackCredentials
}

// GetOracleFallbackCredentialsOk returns a tuple with the OracleFallbackCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleAttachDSourceParameters) GetOracleFallbackCredentialsOk() (*string, bool) {
	if o == nil || IsNil(o.OracleFallbackCredentials) {
		return nil, false
	}
	return o.OracleFallbackCredentials, true
}

// HasOracleFallbackCredentials returns a boolean if a field has been set.
func (o *OracleAttachDSourceParameters) HasOracleFallbackCredentials() bool {
	if o != nil && !IsNil(o.OracleFallbackCredentials) {
		return true
	}

	return false
}

// SetOracleFallbackCredentials gets a reference to the given string and assigns it to the OracleFallbackCredentials field.
func (o *OracleAttachDSourceParameters) SetOracleFallbackCredentials(v string) {
	o.OracleFallbackCredentials = &v
}

// GetRmanChannels returns the RmanChannels field value if set, zero value otherwise.
func (o *OracleAttachDSourceParameters) GetRmanChannels() int32 {
	if o == nil || IsNil(o.RmanChannels) {
		var ret int32
		return ret
	}
	return *o.RmanChannels
}

// GetRmanChannelsOk returns a tuple with the RmanChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleAttachDSourceParameters) GetRmanChannelsOk() (*int32, bool) {
	if o == nil || IsNil(o.RmanChannels) {
		return nil, false
	}
	return o.RmanChannels, true
}

// HasRmanChannels returns a boolean if a field has been set.
func (o *OracleAttachDSourceParameters) HasRmanChannels() bool {
	if o != nil && !IsNil(o.RmanChannels) {
		return true
	}

	return false
}

// SetRmanChannels gets a reference to the given int32 and assigns it to the RmanChannels field.
func (o *OracleAttachDSourceParameters) SetRmanChannels(v int32) {
	o.RmanChannels = &v
}

func (o OracleAttachDSourceParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OracleAttachDSourceParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source_id"] = o.SourceId
	if !IsNil(o.BackupLevelEnabled) {
		toSerialize["backup_level_enabled"] = o.BackupLevelEnabled
	}
	if !IsNil(o.BandwidthLimit) {
		toSerialize["bandwidth_limit"] = o.BandwidthLimit
	}
	if !IsNil(o.CheckLogical) {
		toSerialize["check_logical"] = o.CheckLogical
	}
	if !IsNil(o.CompressedLinkingEnabled) {
		toSerialize["compressed_linking_enabled"] = o.CompressedLinkingEnabled
	}
	if !IsNil(o.DoubleSync) {
		toSerialize["double_sync"] = o.DoubleSync
	}
	if !IsNil(o.EncryptedLinkingEnabled) {
		toSerialize["encrypted_linking_enabled"] = o.EncryptedLinkingEnabled
	}
	if !IsNil(o.EnvironmentUser) {
		toSerialize["environment_user"] = o.EnvironmentUser
	}
	if !IsNil(o.ExternalFilePath) {
		toSerialize["external_file_path"] = o.ExternalFilePath
	}
	if !IsNil(o.FilesPerSet) {
		toSerialize["files_per_set"] = o.FilesPerSet
	}
	if !IsNil(o.Force) {
		toSerialize["force"] = o.Force
	}
	if !IsNil(o.LinkNow) {
		toSerialize["link_now"] = o.LinkNow
	}
	if !IsNil(o.NumberOfConnections) {
		toSerialize["number_of_connections"] = o.NumberOfConnections
	}
	if !IsNil(o.Operations) {
		toSerialize["operations"] = o.Operations
	}
	if !IsNil(o.OracleFallbackUser) {
		toSerialize["oracle_fallback_user"] = o.OracleFallbackUser
	}
	if !IsNil(o.OracleFallbackCredentials) {
		toSerialize["oracle_fallback_credentials"] = o.OracleFallbackCredentials
	}
	if !IsNil(o.RmanChannels) {
		toSerialize["rman_channels"] = o.RmanChannels
	}
	return toSerialize, nil
}

func (o *OracleAttachDSourceParameters) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"source_id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOracleAttachDSourceParameters := _OracleAttachDSourceParameters{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOracleAttachDSourceParameters)

	if err != nil {
		return err
	}

	*o = OracleAttachDSourceParameters(varOracleAttachDSourceParameters)

	return err
}

type NullableOracleAttachDSourceParameters struct {
	value *OracleAttachDSourceParameters
	isSet bool
}

func (v NullableOracleAttachDSourceParameters) Get() *OracleAttachDSourceParameters {
	return v.value
}

func (v *NullableOracleAttachDSourceParameters) Set(val *OracleAttachDSourceParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableOracleAttachDSourceParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableOracleAttachDSourceParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOracleAttachDSourceParameters(val *OracleAttachDSourceParameters) *NullableOracleAttachDSourceParameters {
	return &NullableOracleAttachDSourceParameters{value: val, isSet: true}
}

func (v NullableOracleAttachDSourceParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOracleAttachDSourceParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

