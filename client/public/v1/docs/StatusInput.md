# StatusInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Did** | **string** | did | 
**Name** | **string** | Name | 
**Type** | **string** | Type | 
**Value** | **float32** | Value | 

## Methods

### NewStatusInput

`func NewStatusInput(did string, name string, type_ string, value float32, ) *StatusInput`

NewStatusInput instantiates a new StatusInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusInputWithDefaults

`func NewStatusInputWithDefaults() *StatusInput`

NewStatusInputWithDefaults instantiates a new StatusInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDid

`func (o *StatusInput) GetDid() string`

GetDid returns the Did field if non-nil, zero value otherwise.

### GetDidOk

`func (o *StatusInput) GetDidOk() (*string, bool)`

GetDidOk returns a tuple with the Did field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDid

`func (o *StatusInput) SetDid(v string)`

SetDid sets Did field to given value.


### GetName

`func (o *StatusInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StatusInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StatusInput) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *StatusInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StatusInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StatusInput) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *StatusInput) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *StatusInput) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *StatusInput) SetValue(v float32)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


