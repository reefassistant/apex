# StatusOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **[]string** | status | 
**Name** | **string** | name | 
**Gid** | **string** | gid | 
**Type** | **string** | type | 
**ID** | **float32** | ID | 
**Did** | **string** | did | 
**Intensity** | Pointer to **float32** | intensity (optional) | [optional] 

## Methods

### NewStatusOutput

`func NewStatusOutput(status []string, name string, gid string, type_ string, iD float32, did string, ) *StatusOutput`

NewStatusOutput instantiates a new StatusOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusOutputWithDefaults

`func NewStatusOutputWithDefaults() *StatusOutput`

NewStatusOutputWithDefaults instantiates a new StatusOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *StatusOutput) GetStatus() []string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StatusOutput) GetStatusOk() (*[]string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StatusOutput) SetStatus(v []string)`

SetStatus sets Status field to given value.


### GetName

`func (o *StatusOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StatusOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StatusOutput) SetName(v string)`

SetName sets Name field to given value.


### GetGid

`func (o *StatusOutput) GetGid() string`

GetGid returns the Gid field if non-nil, zero value otherwise.

### GetGidOk

`func (o *StatusOutput) GetGidOk() (*string, bool)`

GetGidOk returns a tuple with the Gid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGid

`func (o *StatusOutput) SetGid(v string)`

SetGid sets Gid field to given value.


### GetType

`func (o *StatusOutput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StatusOutput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StatusOutput) SetType(v string)`

SetType sets Type field to given value.


### GetID

`func (o *StatusOutput) GetID() float32`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *StatusOutput) GetIDOk() (*float32, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *StatusOutput) SetID(v float32)`

SetID sets ID field to given value.


### GetDid

`func (o *StatusOutput) GetDid() string`

GetDid returns the Did field if non-nil, zero value otherwise.

### GetDidOk

`func (o *StatusOutput) GetDidOk() (*string, bool)`

GetDidOk returns a tuple with the Did field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDid

`func (o *StatusOutput) SetDid(v string)`

SetDid sets Did field to given value.


### GetIntensity

`func (o *StatusOutput) GetIntensity() float32`

GetIntensity returns the Intensity field if non-nil, zero value otherwise.

### GetIntensityOk

`func (o *StatusOutput) GetIntensityOk() (*float32, bool)`

GetIntensityOk returns a tuple with the Intensity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntensity

`func (o *StatusOutput) SetIntensity(v float32)`

SetIntensity sets Intensity field to given value.

### HasIntensity

`func (o *StatusOutput) HasIntensity() bool`

HasIntensity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


