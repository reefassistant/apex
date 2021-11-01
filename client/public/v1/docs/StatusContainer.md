# StatusContainer

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
**Date** | **int64** |  | 
**Feed** | Pointer to [**StatusFeed**](StatusFeed.md) |  | [optional] 
**Power** | Pointer to [**StatusPower**](StatusPower.md) |  | [optional] 
**Inputs** | [**[]StatusInput**](StatusInput.md) |  | 
**Outputs** | [**[]StatusOutput**](StatusOutput.md) |  | 
**Link** | Pointer to [**StatusLink**](StatusLink.md) |  | [optional] 

## Methods

### NewStatusContainer

`func NewStatusContainer(hostname string, software string, hardware string, serial string, type_ string, timezone string, date int64, inputs []StatusInput, outputs []StatusOutput, ) *StatusContainer`

NewStatusContainer instantiates a new StatusContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusContainerWithDefaults

`func NewStatusContainerWithDefaults() *StatusContainer`

NewStatusContainerWithDefaults instantiates a new StatusContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *StatusContainer) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *StatusContainer) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *StatusContainer) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetSoftware

`func (o *StatusContainer) GetSoftware() string`

GetSoftware returns the Software field if non-nil, zero value otherwise.

### GetSoftwareOk

`func (o *StatusContainer) GetSoftwareOk() (*string, bool)`

GetSoftwareOk returns a tuple with the Software field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftware

`func (o *StatusContainer) SetSoftware(v string)`

SetSoftware sets Software field to given value.


### GetHardware

`func (o *StatusContainer) GetHardware() string`

GetHardware returns the Hardware field if non-nil, zero value otherwise.

### GetHardwareOk

`func (o *StatusContainer) GetHardwareOk() (*string, bool)`

GetHardwareOk returns a tuple with the Hardware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardware

`func (o *StatusContainer) SetHardware(v string)`

SetHardware sets Hardware field to given value.


### GetSerial

`func (o *StatusContainer) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *StatusContainer) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *StatusContainer) SetSerial(v string)`

SetSerial sets Serial field to given value.


### GetType

`func (o *StatusContainer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StatusContainer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StatusContainer) SetType(v string)`

SetType sets Type field to given value.


### GetExtra

`func (o *StatusContainer) GetExtra() map[string]string`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *StatusContainer) GetExtraOk() (*map[string]string, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *StatusContainer) SetExtra(v map[string]string)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *StatusContainer) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetTimezone

`func (o *StatusContainer) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *StatusContainer) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *StatusContainer) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetDate

`func (o *StatusContainer) GetDate() int64`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *StatusContainer) GetDateOk() (*int64, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *StatusContainer) SetDate(v int64)`

SetDate sets Date field to given value.


### GetFeed

`func (o *StatusContainer) GetFeed() StatusFeed`

GetFeed returns the Feed field if non-nil, zero value otherwise.

### GetFeedOk

`func (o *StatusContainer) GetFeedOk() (*StatusFeed, bool)`

GetFeedOk returns a tuple with the Feed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeed

`func (o *StatusContainer) SetFeed(v StatusFeed)`

SetFeed sets Feed field to given value.

### HasFeed

`func (o *StatusContainer) HasFeed() bool`

HasFeed returns a boolean if a field has been set.

### GetPower

`func (o *StatusContainer) GetPower() StatusPower`

GetPower returns the Power field if non-nil, zero value otherwise.

### GetPowerOk

`func (o *StatusContainer) GetPowerOk() (*StatusPower, bool)`

GetPowerOk returns a tuple with the Power field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPower

`func (o *StatusContainer) SetPower(v StatusPower)`

SetPower sets Power field to given value.

### HasPower

`func (o *StatusContainer) HasPower() bool`

HasPower returns a boolean if a field has been set.

### GetInputs

`func (o *StatusContainer) GetInputs() []StatusInput`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *StatusContainer) GetInputsOk() (*[]StatusInput, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *StatusContainer) SetInputs(v []StatusInput)`

SetInputs sets Inputs field to given value.


### GetOutputs

`func (o *StatusContainer) GetOutputs() []StatusOutput`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *StatusContainer) GetOutputsOk() (*[]StatusOutput, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *StatusContainer) SetOutputs(v []StatusOutput)`

SetOutputs sets Outputs field to given value.


### GetLink

`func (o *StatusContainer) GetLink() StatusLink`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *StatusContainer) GetLinkOk() (*StatusLink, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *StatusContainer) SetLink(v StatusLink)`

SetLink sets Link field to given value.

### HasLink

`func (o *StatusContainer) HasLink() bool`

HasLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


