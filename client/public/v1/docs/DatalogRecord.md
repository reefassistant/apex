# DatalogRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **int64** |  | 
**Data** | [**[]DatalogData**](DatalogData.md) |  | 

## Methods

### NewDatalogRecord

`func NewDatalogRecord(date int64, data []DatalogData, ) *DatalogRecord`

NewDatalogRecord instantiates a new DatalogRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatalogRecordWithDefaults

`func NewDatalogRecordWithDefaults() *DatalogRecord`

NewDatalogRecordWithDefaults instantiates a new DatalogRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *DatalogRecord) GetDate() int64`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *DatalogRecord) GetDateOk() (*int64, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *DatalogRecord) SetDate(v int64)`

SetDate sets Date field to given value.


### GetData

`func (o *DatalogRecord) GetData() []DatalogData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DatalogRecord) GetDataOk() (*[]DatalogData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DatalogRecord) SetData(v []DatalogData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


