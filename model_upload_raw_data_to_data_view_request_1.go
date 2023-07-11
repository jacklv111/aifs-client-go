/*
Aifs api

aifs api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UploadRawDataToDataViewRequest1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UploadRawDataToDataViewRequest1{}

// UploadRawDataToDataViewRequest1 struct for UploadRawDataToDataViewRequest1
type UploadRawDataToDataViewRequest1 struct {
	// the folder path of the resource
	ResourcePath string `json:"resourcePath"`
}

// NewUploadRawDataToDataViewRequest1 instantiates a new UploadRawDataToDataViewRequest1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadRawDataToDataViewRequest1(resourcePath string) *UploadRawDataToDataViewRequest1 {
	this := UploadRawDataToDataViewRequest1{}
	this.ResourcePath = resourcePath
	return &this
}

// NewUploadRawDataToDataViewRequest1WithDefaults instantiates a new UploadRawDataToDataViewRequest1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadRawDataToDataViewRequest1WithDefaults() *UploadRawDataToDataViewRequest1 {
	this := UploadRawDataToDataViewRequest1{}
	return &this
}

// GetResourcePath returns the ResourcePath field value
func (o *UploadRawDataToDataViewRequest1) GetResourcePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourcePath
}

// GetResourcePathOk returns a tuple with the ResourcePath field value
// and a boolean to check if the value has been set.
func (o *UploadRawDataToDataViewRequest1) GetResourcePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourcePath, true
}

// SetResourcePath sets field value
func (o *UploadRawDataToDataViewRequest1) SetResourcePath(v string) {
	o.ResourcePath = v
}

func (o UploadRawDataToDataViewRequest1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UploadRawDataToDataViewRequest1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resourcePath"] = o.ResourcePath
	return toSerialize, nil
}

type NullableUploadRawDataToDataViewRequest1 struct {
	value *UploadRawDataToDataViewRequest1
	isSet bool
}

func (v NullableUploadRawDataToDataViewRequest1) Get() *UploadRawDataToDataViewRequest1 {
	return v.value
}

func (v *NullableUploadRawDataToDataViewRequest1) Set(val *UploadRawDataToDataViewRequest1) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadRawDataToDataViewRequest1) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadRawDataToDataViewRequest1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadRawDataToDataViewRequest1(val *UploadRawDataToDataViewRequest1) *NullableUploadRawDataToDataViewRequest1 {
	return &NullableUploadRawDataToDataViewRequest1{value: val, isSet: true}
}

func (v NullableUploadRawDataToDataViewRequest1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadRawDataToDataViewRequest1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

