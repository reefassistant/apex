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

// ProfileConfig struct for ProfileConfig
type ProfileConfig struct {
	Name string                            `json:"name"`
	ID   float32                           `json:"ID"`
	Type string                            `json:"type"`
	Data map[string]map[string]interface{} `json:"data"`
}

// NewProfileConfig instantiates a new ProfileConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileConfig(name string, iD float32, type_ string, data map[string]map[string]interface{}) *ProfileConfig {
	this := ProfileConfig{}
	this.Name = name
	this.ID = iD
	this.Type = type_
	this.Data = data
	return &this
}

// NewProfileConfigWithDefaults instantiates a new ProfileConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileConfigWithDefaults() *ProfileConfig {
	this := ProfileConfig{}
	return &this
}

// GetName returns the Name field value
func (o *ProfileConfig) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProfileConfig) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProfileConfig) SetName(v string) {
	o.Name = v
}

// GetID returns the ID field value
func (o *ProfileConfig) GetID() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ID
}

// GetIDOk returns a tuple with the ID field value
// and a boolean to check if the value has been set.
func (o *ProfileConfig) GetIDOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ID, true
}

// SetID sets field value
func (o *ProfileConfig) SetID(v float32) {
	o.ID = v
}

// GetType returns the Type field value
func (o *ProfileConfig) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ProfileConfig) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ProfileConfig) SetType(v string) {
	o.Type = v
}

// GetData returns the Data field value
func (o *ProfileConfig) GetData() map[string]map[string]interface{} {
	if o == nil {
		var ret map[string]map[string]interface{}
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ProfileConfig) GetDataOk() (*map[string]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ProfileConfig) SetData(v map[string]map[string]interface{}) {
	o.Data = v
}

func (o ProfileConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["ID"] = o.ID
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableProfileConfig struct {
	value *ProfileConfig
	isSet bool
}

func (v NullableProfileConfig) Get() *ProfileConfig {
	return v.value
}

func (v *NullableProfileConfig) Set(val *ProfileConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileConfig(val *ProfileConfig) *NullableProfileConfig {
	return &NullableProfileConfig{value: val, isSet: true}
}

func (v NullableProfileConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
