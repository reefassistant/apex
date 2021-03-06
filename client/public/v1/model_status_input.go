/*
Apex Public API

Public (unauthenticated) API to integrate with your local Apex aquarium controller (AOS 5 compatible).

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// StatusInput struct for StatusInput
type StatusInput struct {
	// did
	Did string `json:"did"`
	// Name
	Name string `json:"name"`
	// Type
	Type string `json:"type"`
	// Value
	Value float32 `json:"value"`
}

// NewStatusInput instantiates a new StatusInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusInput(did string, name string, type_ string, value float32) *StatusInput {
	this := StatusInput{}
	this.Did = did
	this.Name = name
	this.Type = type_
	this.Value = value
	return &this
}

// NewStatusInputWithDefaults instantiates a new StatusInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusInputWithDefaults() *StatusInput {
	this := StatusInput{}
	return &this
}

// GetDid returns the Did field value
func (o *StatusInput) GetDid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Did
}

// GetDidOk returns a tuple with the Did field value
// and a boolean to check if the value has been set.
func (o *StatusInput) GetDidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Did, true
}

// SetDid sets field value
func (o *StatusInput) SetDid(v string) {
	o.Did = v
}

// GetName returns the Name field value
func (o *StatusInput) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StatusInput) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StatusInput) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *StatusInput) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StatusInput) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *StatusInput) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *StatusInput) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *StatusInput) GetValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *StatusInput) SetValue(v float32) {
	o.Value = v
}

func (o StatusInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["did"] = o.Did
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableStatusInput struct {
	value *StatusInput
	isSet bool
}

func (v NullableStatusInput) Get() *StatusInput {
	return v.value
}

func (v *NullableStatusInput) Set(val *StatusInput) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusInput) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusInput(val *StatusInput) *NullableStatusInput {
	return &NullableStatusInput{value: val, isSet: true}
}

func (v NullableStatusInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
