# DisplayConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | **float32** |  | 
**LineEnables** | **[]bool** |  | 
**Probes** | **[]string** |  | 
**Outpus** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDisplayConfig

`func NewDisplayConfig(page float32, lineEnables []bool, probes []string, ) *DisplayConfig`

NewDisplayConfig instantiates a new DisplayConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisplayConfigWithDefaults

`func NewDisplayConfigWithDefaults() *DisplayConfig`

NewDisplayConfigWithDefaults instantiates a new DisplayConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *DisplayConfig) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *DisplayConfig) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *DisplayConfig) SetPage(v float32)`

SetPage sets Page field to given value.


### GetLineEnables

`func (o *DisplayConfig) GetLineEnables() []bool`

GetLineEnables returns the LineEnables field if non-nil, zero value otherwise.

### GetLineEnablesOk

`func (o *DisplayConfig) GetLineEnablesOk() (*[]bool, bool)`

GetLineEnablesOk returns a tuple with the LineEnables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineEnables

`func (o *DisplayConfig) SetLineEnables(v []bool)`

SetLineEnables sets LineEnables field to given value.


### GetProbes

`func (o *DisplayConfig) GetProbes() []string`

GetProbes returns the Probes field if non-nil, zero value otherwise.

### GetProbesOk

`func (o *DisplayConfig) GetProbesOk() (*[]string, bool)`

GetProbesOk returns a tuple with the Probes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbes

`func (o *DisplayConfig) SetProbes(v []string)`

SetProbes sets Probes field to given value.


### GetOutpus

`func (o *DisplayConfig) GetOutpus() []string`

GetOutpus returns the Outpus field if non-nil, zero value otherwise.

### GetOutpusOk

`func (o *DisplayConfig) GetOutpusOk() (*[]string, bool)`

GetOutpusOk returns a tuple with the Outpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutpus

`func (o *DisplayConfig) SetOutpus(v []string)`

SetOutpus sets Outpus field to given value.

### HasOutpus

`func (o *DisplayConfig) HasOutpus() bool`

HasOutpus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


