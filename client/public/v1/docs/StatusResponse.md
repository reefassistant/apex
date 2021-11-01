# StatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Istat** | [**StatusContainer**](StatusContainer.md) |  | 

## Methods

### NewStatusResponse

`func NewStatusResponse(istat StatusContainer, ) *StatusResponse`

NewStatusResponse instantiates a new StatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusResponseWithDefaults

`func NewStatusResponseWithDefaults() *StatusResponse`

NewStatusResponseWithDefaults instantiates a new StatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIstat

`func (o *StatusResponse) GetIstat() StatusContainer`

GetIstat returns the Istat field if non-nil, zero value otherwise.

### GetIstatOk

`func (o *StatusResponse) GetIstatOk() (*StatusContainer, bool)`

GetIstatOk returns a tuple with the Istat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIstat

`func (o *StatusResponse) SetIstat(v StatusContainer)`

SetIstat sets Istat field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


