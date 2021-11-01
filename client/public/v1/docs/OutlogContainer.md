# OutlogContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | **string** | The (host) name of the APEX controller. | 
**Software** | **string** | Software version. | 
**Hardware** | **string** | Hardware version. | 
**Serial** | **string** | Serial number. | 
**Type** | **string** | Controller type. | 
**Extra** | Pointer to **map[string]string** | Key/value pairs. | [optional] 
**Timezone** | **string** | Time zone. | 
**Record** | [**[]OutlogRecord**](OutlogRecord.md) |  | 

## Methods

### NewOutlogContainer

`func NewOutlogContainer(hostname string, software string, hardware string, serial string, type_ string, timezone string, record []OutlogRecord, ) *OutlogContainer`

NewOutlogContainer instantiates a new OutlogContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutlogContainerWithDefaults

`func NewOutlogContainerWithDefaults() *OutlogContainer`

NewOutlogContainerWithDefaults instantiates a new OutlogContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *OutlogContainer) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *OutlogContainer) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *OutlogContainer) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetSoftware

`func (o *OutlogContainer) GetSoftware() string`

GetSoftware returns the Software field if non-nil, zero value otherwise.

### GetSoftwareOk

`func (o *OutlogContainer) GetSoftwareOk() (*string, bool)`

GetSoftwareOk returns a tuple with the Software field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftware

`func (o *OutlogContainer) SetSoftware(v string)`

SetSoftware sets Software field to given value.


### GetHardware

`func (o *OutlogContainer) GetHardware() string`

GetHardware returns the Hardware field if non-nil, zero value otherwise.

### GetHardwareOk

`func (o *OutlogContainer) GetHardwareOk() (*string, bool)`

GetHardwareOk returns a tuple with the Hardware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardware

`func (o *OutlogContainer) SetHardware(v string)`

SetHardware sets Hardware field to given value.


### GetSerial

`func (o *OutlogContainer) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *OutlogContainer) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *OutlogContainer) SetSerial(v string)`

SetSerial sets Serial field to given value.


### GetType

`func (o *OutlogContainer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OutlogContainer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OutlogContainer) SetType(v string)`

SetType sets Type field to given value.


### GetExtra

`func (o *OutlogContainer) GetExtra() map[string]string`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *OutlogContainer) GetExtraOk() (*map[string]string, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *OutlogContainer) SetExtra(v map[string]string)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *OutlogContainer) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetTimezone

`func (o *OutlogContainer) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *OutlogContainer) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *OutlogContainer) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetRecord

`func (o *OutlogContainer) GetRecord() []OutlogRecord`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *OutlogContainer) GetRecordOk() (*[]OutlogRecord, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *OutlogContainer) SetRecord(v []OutlogRecord)`

SetRecord sets Record field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


