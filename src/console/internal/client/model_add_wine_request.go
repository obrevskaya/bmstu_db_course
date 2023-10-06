/*
API for ppo project

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the AddWineRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddWineRequest{}

// AddWineRequest struct for AddWineRequest
type AddWineRequest struct {
	Name     string `json:"name"`
	Count    string `json:"count"`
	Year     int32  `json:"year"`
	Strength int32  `json:"strength"`
	Price    string `json:"price"`
	Type     string `json:"type"`
}

// NewAddWineRequest instantiates a new AddWineRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddWineRequest(name string, count string, year int32, strength int32, price string, type_ string) *AddWineRequest {
	this := AddWineRequest{}
	this.Name = name
	this.Count = count
	this.Year = year
	this.Strength = strength
	this.Price = price
	this.Type = type_
	return &this
}

// NewAddWineRequestWithDefaults instantiates a new AddWineRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddWineRequestWithDefaults() *AddWineRequest {
	this := AddWineRequest{}
	return &this
}

// GetName returns the Name field value
func (o *AddWineRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AddWineRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AddWineRequest) SetName(v string) {
	o.Name = v
}

// GetCount returns the Count field value
func (o *AddWineRequest) GetCount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *AddWineRequest) GetCountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *AddWineRequest) SetCount(v string) {
	o.Count = v
}

// GetYear returns the Year field value
func (o *AddWineRequest) GetYear() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Year
}

// GetYearOk returns a tuple with the Year field value
// and a boolean to check if the value has been set.
func (o *AddWineRequest) GetYearOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Year, true
}

// SetYear sets field value
func (o *AddWineRequest) SetYear(v int32) {
	o.Year = v
}

// GetStrength returns the Strength field value
func (o *AddWineRequest) GetStrength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Strength
}

// GetStrengthOk returns a tuple with the Strength field value
// and a boolean to check if the value has been set.
func (o *AddWineRequest) GetStrengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Strength, true
}

// SetStrength sets field value
func (o *AddWineRequest) SetStrength(v int32) {
	o.Strength = v
}

// GetPrice returns the Price field value
func (o *AddWineRequest) GetPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *AddWineRequest) GetPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *AddWineRequest) SetPrice(v string) {
	o.Price = v
}

// GetType returns the Type field value
func (o *AddWineRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AddWineRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AddWineRequest) SetType(v string) {
	o.Type = v
}

func (o AddWineRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddWineRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["count"] = o.Count
	toSerialize["year"] = o.Year
	toSerialize["strength"] = o.Strength
	toSerialize["price"] = o.Price
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableAddWineRequest struct {
	value *AddWineRequest
	isSet bool
}

func (v NullableAddWineRequest) Get() *AddWineRequest {
	return v.value
}

func (v *NullableAddWineRequest) Set(val *AddWineRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddWineRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddWineRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddWineRequest(val *AddWineRequest) *NullableAddWineRequest {
	return &NullableAddWineRequest{value: val, isSet: true}
}

func (v NullableAddWineRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddWineRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}