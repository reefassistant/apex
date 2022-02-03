# SystemStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | **string** |  | 
**Software** | **string** |  | 
**Hardware** | **string** |  | 
**Serial** | **string** |  | 
**Type** | **string** |  | 
**Extra** | **map[string]map[string]interface{}** |  | 
**Timezone** | **string** |  | 
**Date** | **int64** |  | 

## Methods

### NewSystemStatus

`func NewSystemStatus(hostname string, software string, hardware string, serial string, type_ string, extra map[string]map[string]interface{}, timezone string, date int64, ) *SystemStatus`

NewSystemStatus instantiates a new SystemStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemStatusWithDefaults

`func NewSystemStatusWithDefaults() *SystemStatus`

NewSystemStatusWithDefaults instantiates a new SystemStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *SystemStatus) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *SystemStatus) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *SystemStatus) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetSoftware

`func (o *SystemStatus) GetSoftware() string`

GetSoftware returns the Software field if non-nil, zero value otherwise.

### GetSoftwareOk

`func (o *SystemStatus) GetSoftwareOk() (*string, bool)`

GetSoftwareOk returns a tuple with the Software field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftware

`func (o *SystemStatus) SetSoftware(v string)`

SetSoftware sets Software field to given value.


### GetHardware

`func (o *SystemStatus) GetHardware() string`

GetHardware returns the Hardware field if non-nil, zero value otherwise.

### GetHardwareOk

`func (o *SystemStatus) GetHardwareOk() (*string, bool)`

GetHardwareOk returns a tuple with the Hardware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardware

`func (o *SystemStatus) SetHardware(v string)`

SetHardware sets Hardware field to given value.


### GetSerial

`func (o *SystemStatus) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *SystemStatus) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *SystemStatus) SetSerial(v string)`

SetSerial sets Serial field to given value.


### GetType

`func (o *SystemStatus) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SystemStatus) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SystemStatus) SetType(v string)`

SetType sets Type field to given value.


### GetExtra

`func (o *SystemStatus) GetExtra() map[string]map[string]interface{}`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *SystemStatus) GetExtraOk() (*map[string]map[string]interface{}, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *SystemStatus) SetExtra(v map[string]map[string]interface{})`

SetExtra sets Extra field to given value.


### GetTimezone

`func (o *SystemStatus) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *SystemStatus) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *SystemStatus) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetDate

`func (o *SystemStatus) GetDate() int64`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *SystemStatus) GetDateOk() (*int64, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *SystemStatus) SetDate(v int64)`

SetDate sets Date field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


