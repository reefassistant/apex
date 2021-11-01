# OutlogRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **int64** |  | 
**Did** | **string** | did | 
**Name** | **string** | name | 
**Value** | **string** | value | 

## Methods

### NewOutlogRecord

`func NewOutlogRecord(date int64, did string, name string, value string, ) *OutlogRecord`

NewOutlogRecord instantiates a new OutlogRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutlogRecordWithDefaults

`func NewOutlogRecordWithDefaults() *OutlogRecord`

NewOutlogRecordWithDefaults instantiates a new OutlogRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *OutlogRecord) GetDate() int64`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *OutlogRecord) GetDateOk() (*int64, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *OutlogRecord) SetDate(v int64)`

SetDate sets Date field to given value.


### GetDid

`func (o *OutlogRecord) GetDid() string`

GetDid returns the Did field if non-nil, zero value otherwise.

### GetDidOk

`func (o *OutlogRecord) GetDidOk() (*string, bool)`

GetDidOk returns a tuple with the Did field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDid

`func (o *OutlogRecord) SetDid(v string)`

SetDid sets Did field to given value.


### GetName

`func (o *OutlogRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OutlogRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OutlogRecord) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *OutlogRecord) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *OutlogRecord) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *OutlogRecord) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


