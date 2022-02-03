# GetConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Oconf** | [**[]OutputConfig**](OutputConfig.md) |  | 
**Pconf** | [**[]ProfileConfig**](ProfileConfig.md) |  | 
**Mconf** | [**[]ModuleConfig**](ModuleConfig.md) |  | 
**Dconf** | [**[]DisplayConfig**](DisplayConfig.md) |  | 
**Iconf** | [**[]InputConfig**](InputConfig.md) |  | 
**Nconf** | [**NetworkConfig**](NetworkConfig.md) |  | 
**Clock** | [**ClockConfig**](ClockConfig.md) |  | 
**Misc** | [**MiscellaneousConfig**](MiscellaneousConfig.md) |  | 
**Season** | [**SeasonConfig**](SeasonConfig.md) |  | 
**Cal** | [**CalibrationConfig**](CalibrationConfig.md) |  | 

## Methods

### NewGetConfigResponse

`func NewGetConfigResponse(oconf []OutputConfig, pconf []ProfileConfig, mconf []ModuleConfig, dconf []DisplayConfig, iconf []InputConfig, nconf NetworkConfig, clock ClockConfig, misc MiscellaneousConfig, season SeasonConfig, cal CalibrationConfig, ) *GetConfigResponse`

NewGetConfigResponse instantiates a new GetConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConfigResponseWithDefaults

`func NewGetConfigResponseWithDefaults() *GetConfigResponse`

NewGetConfigResponseWithDefaults instantiates a new GetConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOconf

`func (o *GetConfigResponse) GetOconf() []OutputConfig`

GetOconf returns the Oconf field if non-nil, zero value otherwise.

### GetOconfOk

`func (o *GetConfigResponse) GetOconfOk() (*[]OutputConfig, bool)`

GetOconfOk returns a tuple with the Oconf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOconf

`func (o *GetConfigResponse) SetOconf(v []OutputConfig)`

SetOconf sets Oconf field to given value.


### GetPconf

`func (o *GetConfigResponse) GetPconf() []ProfileConfig`

GetPconf returns the Pconf field if non-nil, zero value otherwise.

### GetPconfOk

`func (o *GetConfigResponse) GetPconfOk() (*[]ProfileConfig, bool)`

GetPconfOk returns a tuple with the Pconf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPconf

`func (o *GetConfigResponse) SetPconf(v []ProfileConfig)`

SetPconf sets Pconf field to given value.


### GetMconf

`func (o *GetConfigResponse) GetMconf() []ModuleConfig`

GetMconf returns the Mconf field if non-nil, zero value otherwise.

### GetMconfOk

`func (o *GetConfigResponse) GetMconfOk() (*[]ModuleConfig, bool)`

GetMconfOk returns a tuple with the Mconf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMconf

`func (o *GetConfigResponse) SetMconf(v []ModuleConfig)`

SetMconf sets Mconf field to given value.


### GetDconf

`func (o *GetConfigResponse) GetDconf() []DisplayConfig`

GetDconf returns the Dconf field if non-nil, zero value otherwise.

### GetDconfOk

`func (o *GetConfigResponse) GetDconfOk() (*[]DisplayConfig, bool)`

GetDconfOk returns a tuple with the Dconf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDconf

`func (o *GetConfigResponse) SetDconf(v []DisplayConfig)`

SetDconf sets Dconf field to given value.


### GetIconf

`func (o *GetConfigResponse) GetIconf() []InputConfig`

GetIconf returns the Iconf field if non-nil, zero value otherwise.

### GetIconfOk

`func (o *GetConfigResponse) GetIconfOk() (*[]InputConfig, bool)`

GetIconfOk returns a tuple with the Iconf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconf

`func (o *GetConfigResponse) SetIconf(v []InputConfig)`

SetIconf sets Iconf field to given value.


### GetNconf

`func (o *GetConfigResponse) GetNconf() NetworkConfig`

GetNconf returns the Nconf field if non-nil, zero value otherwise.

### GetNconfOk

`func (o *GetConfigResponse) GetNconfOk() (*NetworkConfig, bool)`

GetNconfOk returns a tuple with the Nconf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNconf

`func (o *GetConfigResponse) SetNconf(v NetworkConfig)`

SetNconf sets Nconf field to given value.


### GetClock

`func (o *GetConfigResponse) GetClock() ClockConfig`

GetClock returns the Clock field if non-nil, zero value otherwise.

### GetClockOk

`func (o *GetConfigResponse) GetClockOk() (*ClockConfig, bool)`

GetClockOk returns a tuple with the Clock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClock

`func (o *GetConfigResponse) SetClock(v ClockConfig)`

SetClock sets Clock field to given value.


### GetMisc

`func (o *GetConfigResponse) GetMisc() MiscellaneousConfig`

GetMisc returns the Misc field if non-nil, zero value otherwise.

### GetMiscOk

`func (o *GetConfigResponse) GetMiscOk() (*MiscellaneousConfig, bool)`

GetMiscOk returns a tuple with the Misc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMisc

`func (o *GetConfigResponse) SetMisc(v MiscellaneousConfig)`

SetMisc sets Misc field to given value.


### GetSeason

`func (o *GetConfigResponse) GetSeason() SeasonConfig`

GetSeason returns the Season field if non-nil, zero value otherwise.

### GetSeasonOk

`func (o *GetConfigResponse) GetSeasonOk() (*SeasonConfig, bool)`

GetSeasonOk returns a tuple with the Season field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeason

`func (o *GetConfigResponse) SetSeason(v SeasonConfig)`

SetSeason sets Season field to given value.


### GetCal

`func (o *GetConfigResponse) GetCal() CalibrationConfig`

GetCal returns the Cal field if non-nil, zero value otherwise.

### GetCalOk

`func (o *GetConfigResponse) GetCalOk() (*CalibrationConfig, bool)`

GetCalOk returns a tuple with the Cal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCal

`func (o *GetConfigResponse) SetCal(v CalibrationConfig)`

SetCal sets Cal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


