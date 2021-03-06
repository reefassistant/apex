/*
Apex Private API

Private (authenticated) API to integrate with your local Apex aquarium controller (AOS 5 compatible).

API version: 1.0.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1alpha1

import (
	"encoding/json"
)

// GetUserResponse struct for GetUserResponse
type GetUserResponse struct {
	Username string `json:"username"`
}

// NewGetUserResponse instantiates a new GetUserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserResponse(username string) *GetUserResponse {
	this := GetUserResponse{}
	this.Username = username
	return &this
}

// NewGetUserResponseWithDefaults instantiates a new GetUserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserResponseWithDefaults() *GetUserResponse {
	this := GetUserResponse{}
	return &this
}

// GetUsername returns the Username field value
func (o *GetUserResponse) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *GetUserResponse) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *GetUserResponse) SetUsername(v string) {
	o.Username = v
}

func (o GetUserResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableGetUserResponse struct {
	value *GetUserResponse
	isSet bool
}

func (v NullableGetUserResponse) Get() *GetUserResponse {
	return v.value
}

func (v *NullableGetUserResponse) Set(val *GetUserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserResponse(val *GetUserResponse) *NullableGetUserResponse {
	return &NullableGetUserResponse{value: val, isSet: true}
}

func (v NullableGetUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
