# StatusLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkState** | **int32** |  | 
**LinkKey** | **string** |  | 
**Link** | **bool** |  | 

## Methods

### NewStatusLink

`func NewStatusLink(linkState int32, linkKey string, link bool, ) *StatusLink`

NewStatusLink instantiates a new StatusLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusLinkWithDefaults

`func NewStatusLinkWithDefaults() *StatusLink`

NewStatusLinkWithDefaults instantiates a new StatusLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkState

`func (o *StatusLink) GetLinkState() int32`

GetLinkState returns the LinkState field if non-nil, zero value otherwise.

### GetLinkStateOk

`func (o *StatusLink) GetLinkStateOk() (*int32, bool)`

GetLinkStateOk returns a tuple with the LinkState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkState

`func (o *StatusLink) SetLinkState(v int32)`

SetLinkState sets LinkState field to given value.


### GetLinkKey

`func (o *StatusLink) GetLinkKey() string`

GetLinkKey returns the LinkKey field if non-nil, zero value otherwise.

### GetLinkKeyOk

`func (o *StatusLink) GetLinkKeyOk() (*string, bool)`

GetLinkKeyOk returns a tuple with the LinkKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkKey

`func (o *StatusLink) SetLinkKey(v string)`

SetLinkKey sets LinkKey field to given value.


### GetLink

`func (o *StatusLink) GetLink() bool`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *StatusLink) GetLinkOk() (*bool, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *StatusLink) SetLink(v bool)`

SetLink sets Link field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


