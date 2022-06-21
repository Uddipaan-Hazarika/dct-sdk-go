/*
Delphix DCT API

Delphix DCT API

API version: 2.0.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// ReportingSchedule struct for ReportingSchedule
type ReportingSchedule struct {
	ReportId *int32 `json:"report_id,omitempty"`
	ReportType string `json:"report_type"`
	// Standard cron expressions are supported e.g. 0 15 10 L * ?  - Schedule at 10:15 AM on the last day of every month, 0 0 2 ? * Mon-Fri - Schedule at 2:00 AM every Monday, Tuesday, Wednesday, Thursday and Friday. For more details kindly refer- \"http://www.quartz-scheduler.org/documentation/\"
	CronExpression string `json:"cron_expression"`
	// Timezones are specified according to the Olson tzinfo database - \"https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\".
	TimeZone *string `json:"time_zone,omitempty"`
	Message *string `json:"message,omitempty"`
	FileFormat string `json:"file_format"`
	Enabled bool `json:"enabled"`
	Recipients []string `json:"recipients"`
	SortColumn *string `json:"sort_column,omitempty"`
	RowCount *int32 `json:"row_count,omitempty"`
}

// NewReportingSchedule instantiates a new ReportingSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportingSchedule(reportType string, cronExpression string, fileFormat string, enabled bool, recipients []string) *ReportingSchedule {
	this := ReportingSchedule{}
	this.ReportType = reportType
	this.CronExpression = cronExpression
	this.FileFormat = fileFormat
	this.Enabled = enabled
	this.Recipients = recipients
	return &this
}

// NewReportingScheduleWithDefaults instantiates a new ReportingSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportingScheduleWithDefaults() *ReportingSchedule {
	this := ReportingSchedule{}
	var enabled bool = true
	this.Enabled = enabled
	return &this
}

// GetReportId returns the ReportId field value if set, zero value otherwise.
func (o *ReportingSchedule) GetReportId() int32 {
	if o == nil || o.ReportId == nil {
		var ret int32
		return ret
	}
	return *o.ReportId
}

// GetReportIdOk returns a tuple with the ReportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingSchedule) GetReportIdOk() (*int32, bool) {
	if o == nil || o.ReportId == nil {
		return nil, false
	}
	return o.ReportId, true
}

// HasReportId returns a boolean if a field has been set.
func (o *ReportingSchedule) HasReportId() bool {
	if o != nil && o.ReportId != nil {
		return true
	}

	return false
}

// SetReportId gets a reference to the given int32 and assigns it to the ReportId field.
func (o *ReportingSchedule) SetReportId(v int32) {
	o.ReportId = &v
}

// GetReportType returns the ReportType field value
func (o *ReportingSchedule) GetReportType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportType
}

// GetReportTypeOk returns a tuple with the ReportType field value
// and a boolean to check if the value has been set.
func (o *ReportingSchedule) GetReportTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReportType, true
}

// SetReportType sets field value
func (o *ReportingSchedule) SetReportType(v string) {
	o.ReportType = v
}

// GetCronExpression returns the CronExpression field value
func (o *ReportingSchedule) GetCronExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CronExpression
}

// GetCronExpressionOk returns a tuple with the CronExpression field value
// and a boolean to check if the value has been set.
func (o *ReportingSchedule) GetCronExpressionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CronExpression, true
}

// SetCronExpression sets field value
func (o *ReportingSchedule) SetCronExpression(v string) {
	o.CronExpression = v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *ReportingSchedule) GetTimeZone() string {
	if o == nil || o.TimeZone == nil {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingSchedule) GetTimeZoneOk() (*string, bool) {
	if o == nil || o.TimeZone == nil {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *ReportingSchedule) HasTimeZone() bool {
	if o != nil && o.TimeZone != nil {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *ReportingSchedule) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ReportingSchedule) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingSchedule) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ReportingSchedule) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ReportingSchedule) SetMessage(v string) {
	o.Message = &v
}

// GetFileFormat returns the FileFormat field value
func (o *ReportingSchedule) GetFileFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileFormat
}

// GetFileFormatOk returns a tuple with the FileFormat field value
// and a boolean to check if the value has been set.
func (o *ReportingSchedule) GetFileFormatOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FileFormat, true
}

// SetFileFormat sets field value
func (o *ReportingSchedule) SetFileFormat(v string) {
	o.FileFormat = v
}

// GetEnabled returns the Enabled field value
func (o *ReportingSchedule) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ReportingSchedule) GetEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ReportingSchedule) SetEnabled(v bool) {
	o.Enabled = v
}

// GetRecipients returns the Recipients field value
func (o *ReportingSchedule) GetRecipients() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *ReportingSchedule) GetRecipientsOk() ([]string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Recipients, true
}

// SetRecipients sets field value
func (o *ReportingSchedule) SetRecipients(v []string) {
	o.Recipients = v
}

// GetSortColumn returns the SortColumn field value if set, zero value otherwise.
func (o *ReportingSchedule) GetSortColumn() string {
	if o == nil || o.SortColumn == nil {
		var ret string
		return ret
	}
	return *o.SortColumn
}

// GetSortColumnOk returns a tuple with the SortColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingSchedule) GetSortColumnOk() (*string, bool) {
	if o == nil || o.SortColumn == nil {
		return nil, false
	}
	return o.SortColumn, true
}

// HasSortColumn returns a boolean if a field has been set.
func (o *ReportingSchedule) HasSortColumn() bool {
	if o != nil && o.SortColumn != nil {
		return true
	}

	return false
}

// SetSortColumn gets a reference to the given string and assigns it to the SortColumn field.
func (o *ReportingSchedule) SetSortColumn(v string) {
	o.SortColumn = &v
}

// GetRowCount returns the RowCount field value if set, zero value otherwise.
func (o *ReportingSchedule) GetRowCount() int32 {
	if o == nil || o.RowCount == nil {
		var ret int32
		return ret
	}
	return *o.RowCount
}

// GetRowCountOk returns a tuple with the RowCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingSchedule) GetRowCountOk() (*int32, bool) {
	if o == nil || o.RowCount == nil {
		return nil, false
	}
	return o.RowCount, true
}

// HasRowCount returns a boolean if a field has been set.
func (o *ReportingSchedule) HasRowCount() bool {
	if o != nil && o.RowCount != nil {
		return true
	}

	return false
}

// SetRowCount gets a reference to the given int32 and assigns it to the RowCount field.
func (o *ReportingSchedule) SetRowCount(v int32) {
	o.RowCount = &v
}

func (o ReportingSchedule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ReportId != nil {
		toSerialize["report_id"] = o.ReportId
	}
	if true {
		toSerialize["report_type"] = o.ReportType
	}
	if true {
		toSerialize["cron_expression"] = o.CronExpression
	}
	if o.TimeZone != nil {
		toSerialize["time_zone"] = o.TimeZone
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["file_format"] = o.FileFormat
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["recipients"] = o.Recipients
	}
	if o.SortColumn != nil {
		toSerialize["sort_column"] = o.SortColumn
	}
	if o.RowCount != nil {
		toSerialize["row_count"] = o.RowCount
	}
	return json.Marshal(toSerialize)
}

type NullableReportingSchedule struct {
	value *ReportingSchedule
	isSet bool
}

func (v NullableReportingSchedule) Get() *ReportingSchedule {
	return v.value
}

func (v *NullableReportingSchedule) Set(val *ReportingSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableReportingSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableReportingSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportingSchedule(val *ReportingSchedule) *NullableReportingSchedule {
	return &NullableReportingSchedule{value: val, isSet: true}
}

func (v NullableReportingSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportingSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


