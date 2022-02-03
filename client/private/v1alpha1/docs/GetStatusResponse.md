# GetStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**System** | [**SystemStatus**](SystemStatus.md) |  | 
**Modules** | **[]map[string]interface{}** |  | 
**Nstat** | **map[string]interface{}** |  | 
**Feed** | **map[string]interface{}** |  | 
**Power** | **map[string]interface{}** |  | 
**Outputs** | **[]map[string]interface{}** |  | 
**Inputs** | **[]map[string]interface{}** |  | 
**Link** | **map[string]interface{}** |  | 

## Methods

### NewGetStatusResponse

`func NewGetStatusResponse(system SystemStatus, modules []map[string]interface{}, nstat map[string]interface{}, feed map[string]interface{}, power map[string]interface{}, outputs []map[string]interface{}, inputs []map[string]interface{}, link map[string]interface{}, ) *GetStatusResponse`

NewGetStatusResponse instantiates a new GetStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStatusResponseWithDefaults

`func NewGetStatusResponseWithDefaults() *GetStatusResponse`

NewGetStatusResponseWithDefaults instantiates a new GetStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSystem

`func (o *GetStatusResponse) GetSystem() SystemStatus`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *GetStatusResponse) GetSystemOk() (*SystemStatus, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *GetStatusResponse) SetSystem(v SystemStatus)`

SetSystem sets System field to given value.


### GetModules

`func (o *GetStatusResponse) GetModules() []map[string]interface{}`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *GetStatusResponse) GetModulesOk() (*[]map[string]interface{}, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *GetStatusResponse) SetModules(v []map[string]interface{})`

SetModules sets Modules field to given value.


### GetNstat

`func (o *GetStatusResponse) GetNstat() map[string]interface{}`

GetNstat returns the Nstat field if non-nil, zero value otherwise.

### GetNstatOk

`func (o *GetStatusResponse) GetNstatOk() (*map[string]interface{}, bool)`

GetNstatOk returns a tuple with the Nstat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNstat

`func (o *GetStatusResponse) SetNstat(v map[string]interface{})`

SetNstat sets Nstat field to given value.


### GetFeed

`func (o *GetStatusResponse) GetFeed() map[string]interface{}`

GetFeed returns the Feed field if non-nil, zero value otherwise.

### GetFeedOk

`func (o *GetStatusResponse) GetFeedOk() (*map[string]interface{}, bool)`

GetFeedOk returns a tuple with the Feed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeed

`func (o *GetStatusResponse) SetFeed(v map[string]interface{})`

SetFeed sets Feed field to given value.


### GetPower

`func (o *GetStatusResponse) GetPower() map[string]interface{}`

GetPower returns the Power field if non-nil, zero value otherwise.

### GetPowerOk

`func (o *GetStatusResponse) GetPowerOk() (*map[string]interface{}, bool)`

GetPowerOk returns a tuple with the Power field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPower

`func (o *GetStatusResponse) SetPower(v map[string]interface{})`

SetPower sets Power field to given value.


### GetOutputs

`func (o *GetStatusResponse) GetOutputs() []map[string]interface{}`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *GetStatusResponse) GetOutputsOk() (*[]map[string]interface{}, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *GetStatusResponse) SetOutputs(v []map[string]interface{})`

SetOutputs sets Outputs field to given value.


### GetInputs

`func (o *GetStatusResponse) GetInputs() []map[string]interface{}`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *GetStatusResponse) GetInputsOk() (*[]map[string]interface{}, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *GetStatusResponse) SetInputs(v []map[string]interface{})`

SetInputs sets Inputs field to given value.


### GetLink

`func (o *GetStatusResponse) GetLink() map[string]interface{}`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *GetStatusResponse) GetLinkOk() (*map[string]interface{}, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *GetStatusResponse) SetLink(v map[string]interface{})`

SetLink sets Link field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


