# LoginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Login** | **string** |  | 
**Password** | **string** |  | 
**RememberMe** | **bool** |  | 

## Methods

### NewLoginRequest

`func NewLoginRequest(login string, password string, rememberMe bool, ) *LoginRequest`

NewLoginRequest instantiates a new LoginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginRequestWithDefaults

`func NewLoginRequestWithDefaults() *LoginRequest`

NewLoginRequestWithDefaults instantiates a new LoginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogin

`func (o *LoginRequest) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *LoginRequest) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *LoginRequest) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetPassword

`func (o *LoginRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *LoginRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *LoginRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetRememberMe

`func (o *LoginRequest) GetRememberMe() bool`

GetRememberMe returns the RememberMe field if non-nil, zero value otherwise.

### GetRememberMeOk

`func (o *LoginRequest) GetRememberMeOk() (*bool, bool)`

GetRememberMeOk returns a tuple with the RememberMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberMe

`func (o *LoginRequest) SetRememberMe(v bool)`

SetRememberMe sets RememberMe field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


