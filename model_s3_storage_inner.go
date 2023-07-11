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

// checks if the S3StorageInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &S3StorageInner{}

// S3StorageInner struct for S3StorageInner
type S3StorageInner struct {
	DataItemId *string `json:"dataItemId,omitempty"`
	DataName *string `json:"dataName,omitempty"`
	// the object key of the data in s3 storage
	ObjectKey *string `json:"objectKey,omitempty"`
}

// NewS3StorageInner instantiates a new S3StorageInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewS3StorageInner() *S3StorageInner {
	this := S3StorageInner{}
	return &this
}

// NewS3StorageInnerWithDefaults instantiates a new S3StorageInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewS3StorageInnerWithDefaults() *S3StorageInner {
	this := S3StorageInner{}
	return &this
}

// GetDataItemId returns the DataItemId field value if set, zero value otherwise.
func (o *S3StorageInner) GetDataItemId() string {
	if o == nil || IsNil(o.DataItemId) {
		var ret string
		return ret
	}
	return *o.DataItemId
}

// GetDataItemIdOk returns a tuple with the DataItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3StorageInner) GetDataItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.DataItemId) {
		return nil, false
	}
	return o.DataItemId, true
}

// HasDataItemId returns a boolean if a field has been set.
func (o *S3StorageInner) HasDataItemId() bool {
	if o != nil && !IsNil(o.DataItemId) {
		return true
	}

	return false
}

// SetDataItemId gets a reference to the given string and assigns it to the DataItemId field.
func (o *S3StorageInner) SetDataItemId(v string) {
	o.DataItemId = &v
}

// GetDataName returns the DataName field value if set, zero value otherwise.
func (o *S3StorageInner) GetDataName() string {
	if o == nil || IsNil(o.DataName) {
		var ret string
		return ret
	}
	return *o.DataName
}

// GetDataNameOk returns a tuple with the DataName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3StorageInner) GetDataNameOk() (*string, bool) {
	if o == nil || IsNil(o.DataName) {
		return nil, false
	}
	return o.DataName, true
}

// HasDataName returns a boolean if a field has been set.
func (o *S3StorageInner) HasDataName() bool {
	if o != nil && !IsNil(o.DataName) {
		return true
	}

	return false
}

// SetDataName gets a reference to the given string and assigns it to the DataName field.
func (o *S3StorageInner) SetDataName(v string) {
	o.DataName = &v
}

// GetObjectKey returns the ObjectKey field value if set, zero value otherwise.
func (o *S3StorageInner) GetObjectKey() string {
	if o == nil || IsNil(o.ObjectKey) {
		var ret string
		return ret
	}
	return *o.ObjectKey
}

// GetObjectKeyOk returns a tuple with the ObjectKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3StorageInner) GetObjectKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectKey) {
		return nil, false
	}
	return o.ObjectKey, true
}

// HasObjectKey returns a boolean if a field has been set.
func (o *S3StorageInner) HasObjectKey() bool {
	if o != nil && !IsNil(o.ObjectKey) {
		return true
	}

	return false
}

// SetObjectKey gets a reference to the given string and assigns it to the ObjectKey field.
func (o *S3StorageInner) SetObjectKey(v string) {
	o.ObjectKey = &v
}

func (o S3StorageInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o S3StorageInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DataItemId) {
		toSerialize["dataItemId"] = o.DataItemId
	}
	if !IsNil(o.DataName) {
		toSerialize["dataName"] = o.DataName
	}
	if !IsNil(o.ObjectKey) {
		toSerialize["objectKey"] = o.ObjectKey
	}
	return toSerialize, nil
}

type NullableS3StorageInner struct {
	value *S3StorageInner
	isSet bool
}

func (v NullableS3StorageInner) Get() *S3StorageInner {
	return v.value
}

func (v *NullableS3StorageInner) Set(val *S3StorageInner) {
	v.value = val
	v.isSet = true
}

func (v NullableS3StorageInner) IsSet() bool {
	return v.isSet
}

func (v *NullableS3StorageInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableS3StorageInner(val *S3StorageInner) *NullableS3StorageInner {
	return &NullableS3StorageInner{value: val, isSet: true}
}

func (v NullableS3StorageInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableS3StorageInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

