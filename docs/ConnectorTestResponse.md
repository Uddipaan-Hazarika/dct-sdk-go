# ConnectorTestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Connection status, SUCCEEDED or FAILED | 
**Message** | **string** | A message describing the result of the masking connector test. | 

## Methods

### NewConnectorTestResponse

`func NewConnectorTestResponse(status string, message string, ) *ConnectorTestResponse`

NewConnectorTestResponse instantiates a new ConnectorTestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorTestResponseWithDefaults

`func NewConnectorTestResponseWithDefaults() *ConnectorTestResponse`

NewConnectorTestResponseWithDefaults instantiates a new ConnectorTestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ConnectorTestResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConnectorTestResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConnectorTestResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *ConnectorTestResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ConnectorTestResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ConnectorTestResponse) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


