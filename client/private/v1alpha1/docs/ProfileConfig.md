# ProfileConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ID** | **float32** |  | 
**Type** | **string** |  | 
**Data** | **map[string]map[string]interface{}** |  | 

## Methods

### NewProfileConfig

`func NewProfileConfig(name string, iD float32, type_ string, data map[string]map[string]interface{}, ) *ProfileConfig`

NewProfileConfig instantiates a new ProfileConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileConfigWithDefaults

`func NewProfileConfigWithDefaults() *ProfileConfig`

NewProfileConfigWithDefaults instantiates a new ProfileConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProfileConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProfileConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProfileConfig) SetName(v string)`

SetName sets Name field to given value.


### GetID

`func (o *ProfileConfig) GetID() float32`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *ProfileConfig) GetIDOk() (*float32, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *ProfileConfig) SetID(v float32)`

SetID sets ID field to given value.


### GetType

`func (o *ProfileConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProfileConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProfileConfig) SetType(v string)`

SetType sets Type field to given value.


### GetData

`func (o *ProfileConfig) GetData() map[string]map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ProfileConfig) GetDataOk() (*map[string]map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ProfileConfig) SetData(v map[string]map[string]interface{})`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


