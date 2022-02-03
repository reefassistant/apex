# InputConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Did** | **string** |  | 
**Enable** | **bool** |  | 
**Type** | **string** |  | 
**Name** | **string** |  | 
**Alarm** | **map[string]float32** | Key/float pairs. | 
**Extra** | **map[string]map[string]interface{}** |  | 

## Methods

### NewInputConfig

`func NewInputConfig(did string, enable bool, type_ string, name string, alarm map[string]float32, extra map[string]map[string]interface{}, ) *InputConfig`

NewInputConfig instantiates a new InputConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputConfigWithDefaults

`func NewInputConfigWithDefaults() *InputConfig`

NewInputConfigWithDefaults instantiates a new InputConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDid

`func (o *InputConfig) GetDid() string`

GetDid returns the Did field if non-nil, zero value otherwise.

### GetDidOk

`func (o *InputConfig) GetDidOk() (*string, bool)`

GetDidOk returns a tuple with the Did field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDid

`func (o *InputConfig) SetDid(v string)`

SetDid sets Did field to given value.


### GetEnable

`func (o *InputConfig) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *InputConfig) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *InputConfig) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetType

`func (o *InputConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputConfig) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *InputConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InputConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InputConfig) SetName(v string)`

SetName sets Name field to given value.


### GetAlarm

`func (o *InputConfig) GetAlarm() map[string]float32`

GetAlarm returns the Alarm field if non-nil, zero value otherwise.

### GetAlarmOk

`func (o *InputConfig) GetAlarmOk() (*map[string]float32, bool)`

GetAlarmOk returns a tuple with the Alarm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarm

`func (o *InputConfig) SetAlarm(v map[string]float32)`

SetAlarm sets Alarm field to given value.


### GetExtra

`func (o *InputConfig) GetExtra() map[string]map[string]interface{}`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *InputConfig) GetExtraOk() (*map[string]map[string]interface{}, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *InputConfig) SetExtra(v map[string]map[string]interface{})`

SetExtra sets Extra field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


