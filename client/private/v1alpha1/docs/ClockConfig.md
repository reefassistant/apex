# ClockConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timezone** | **[]string** |  | 
**Date** | **int64** |  | 
**Dst** | **bool** |  | 
**Auto** | **bool** |  | 

## Methods

### NewClockConfig

`func NewClockConfig(timezone []string, date int64, dst bool, auto bool, ) *ClockConfig`

NewClockConfig instantiates a new ClockConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClockConfigWithDefaults

`func NewClockConfigWithDefaults() *ClockConfig`

NewClockConfigWithDefaults instantiates a new ClockConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimezone

`func (o *ClockConfig) GetTimezone() []string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ClockConfig) GetTimezoneOk() (*[]string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ClockConfig) SetTimezone(v []string)`

SetTimezone sets Timezone field to given value.


### GetDate

`func (o *ClockConfig) GetDate() int64`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ClockConfig) GetDateOk() (*int64, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ClockConfig) SetDate(v int64)`

SetDate sets Date field to given value.


### GetDst

`func (o *ClockConfig) GetDst() bool`

GetDst returns the Dst field if non-nil, zero value otherwise.

### GetDstOk

`func (o *ClockConfig) GetDstOk() (*bool, bool)`

GetDstOk returns a tuple with the Dst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDst

`func (o *ClockConfig) SetDst(v bool)`

SetDst sets Dst field to given value.


### GetAuto

`func (o *ClockConfig) GetAuto() bool`

GetAuto returns the Auto field if non-nil, zero value otherwise.

### GetAutoOk

`func (o *ClockConfig) GetAutoOk() (*bool, bool)`

GetAutoOk returns a tuple with the Auto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuto

`func (o *ClockConfig) SetAuto(v bool)`

SetAuto sets Auto field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


