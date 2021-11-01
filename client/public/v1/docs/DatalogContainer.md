# DatalogContainer

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
**Record** | [**[]DatalogRecord**](DatalogRecord.md) |  | 

## Methods

### NewDatalogContainer

`func NewDatalogContainer(hostname string, software string, hardware string, serial string, type_ string, timezone string, record []DatalogRecord, ) *DatalogContainer`

NewDatalogContainer instantiates a new DatalogContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatalogContainerWithDefaults

`func NewDatalogContainerWithDefaults() *DatalogContainer`

NewDatalogContainerWithDefaults instantiates a new DatalogContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *DatalogContainer) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *DatalogContainer) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *DatalogContainer) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetSoftware

`func (o *DatalogContainer) GetSoftware() string`

GetSoftware returns the Software field if non-nil, zero value otherwise.

### GetSoftwareOk

`func (o *DatalogContainer) GetSoftwareOk() (*string, bool)`

GetSoftwareOk returns a tuple with the Software field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftware

`func (o *DatalogContainer) SetSoftware(v string)`

SetSoftware sets Software field to given value.


### GetHardware

`func (o *DatalogContainer) GetHardware() string`

GetHardware returns the Hardware field if non-nil, zero value otherwise.

### GetHardwareOk

`func (o *DatalogContainer) GetHardwareOk() (*string, bool)`

GetHardwareOk returns a tuple with the Hardware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardware

`func (o *DatalogContainer) SetHardware(v string)`

SetHardware sets Hardware field to given value.


### GetSerial

`func (o *DatalogContainer) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *DatalogContainer) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *DatalogContainer) SetSerial(v string)`

SetSerial sets Serial field to given value.


### GetType

`func (o *DatalogContainer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DatalogContainer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DatalogContainer) SetType(v string)`

SetType sets Type field to given value.


### GetExtra

`func (o *DatalogContainer) GetExtra() map[string]string`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *DatalogContainer) GetExtraOk() (*map[string]string, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *DatalogContainer) SetExtra(v map[string]string)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *DatalogContainer) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetTimezone

`func (o *DatalogContainer) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *DatalogContainer) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *DatalogContainer) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetRecord

`func (o *DatalogContainer) GetRecord() []DatalogRecord`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *DatalogContainer) GetRecordOk() (*[]DatalogRecord, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *DatalogContainer) SetRecord(v []DatalogRecord)`

SetRecord sets Record field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


