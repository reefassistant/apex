# StatusFeed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **int32** | Feed cycle (A&#x3D;1, B&#x3D;2, ...) | 
**Active** | **int32** | seconds left for active feed cycle | 

## Methods

### NewStatusFeed

`func NewStatusFeed(name int32, active int32, ) *StatusFeed`

NewStatusFeed instantiates a new StatusFeed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusFeedWithDefaults

`func NewStatusFeedWithDefaults() *StatusFeed`

NewStatusFeedWithDefaults instantiates a new StatusFeed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StatusFeed) GetName() int32`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StatusFeed) GetNameOk() (*int32, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StatusFeed) SetName(v int32)`

SetName sets Name field to given value.


### GetActive

`func (o *StatusFeed) GetActive() int32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *StatusFeed) GetActiveOk() (*int32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *StatusFeed) SetActive(v int32)`

SetActive sets Active field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


