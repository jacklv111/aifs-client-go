/*
Aifs api

aifs api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"os"
)

// checks if the ImageData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageData{}

// ImageData struct for ImageData
type ImageData struct {
	Format *string `json:"format,omitempty"`
	Data **os.File `json:"data,omitempty"`
}

// NewImageData instantiates a new ImageData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageData() *ImageData {
	this := ImageData{}
	return &this
}

// NewImageDataWithDefaults instantiates a new ImageData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageDataWithDefaults() *ImageData {
	this := ImageData{}
	return &this
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *ImageData) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageData) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *ImageData) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *ImageData) SetFormat(v string) {
	o.Format = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ImageData) GetData() *os.File {
	if o == nil || IsNil(o.Data) {
		var ret *os.File
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageData) GetDataOk() (**os.File, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ImageData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given *os.File and assigns it to the Data field.
func (o *ImageData) SetData(v *os.File) {
	o.Data = &v
}

func (o ImageData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableImageData struct {
	value *ImageData
	isSet bool
}

func (v NullableImageData) Get() *ImageData {
	return v.value
}

func (v *NullableImageData) Set(val *ImageData) {
	v.value = val
	v.isSet = true
}

func (v NullableImageData) IsSet() bool {
	return v.isSet
}

func (v *NullableImageData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageData(val *ImageData) *NullableImageData {
	return &NullableImageData{value: val, isSet: true}
}

func (v NullableImageData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

