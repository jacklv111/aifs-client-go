/*
Aifs api

aifs api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aifsclientgo

import (
	"encoding/json"
)

// checks if the CreateDataViewSuccessResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDataViewSuccessResp{}

// CreateDataViewSuccessResp struct for CreateDataViewSuccessResp
type CreateDataViewSuccessResp struct {
	DataViewId *string `json:"dataViewId,omitempty"`
}

// NewCreateDataViewSuccessResp instantiates a new CreateDataViewSuccessResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDataViewSuccessResp() *CreateDataViewSuccessResp {
	this := CreateDataViewSuccessResp{}
	return &this
}

// NewCreateDataViewSuccessRespWithDefaults instantiates a new CreateDataViewSuccessResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDataViewSuccessRespWithDefaults() *CreateDataViewSuccessResp {
	this := CreateDataViewSuccessResp{}
	return &this
}

// GetDataViewId returns the DataViewId field value if set, zero value otherwise.
func (o *CreateDataViewSuccessResp) GetDataViewId() string {
	if o == nil || IsNil(o.DataViewId) {
		var ret string
		return ret
	}
	return *o.DataViewId
}

// GetDataViewIdOk returns a tuple with the DataViewId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDataViewSuccessResp) GetDataViewIdOk() (*string, bool) {
	if o == nil || IsNil(o.DataViewId) {
		return nil, false
	}
	return o.DataViewId, true
}

// HasDataViewId returns a boolean if a field has been set.
func (o *CreateDataViewSuccessResp) HasDataViewId() bool {
	if o != nil && !IsNil(o.DataViewId) {
		return true
	}

	return false
}

// SetDataViewId gets a reference to the given string and assigns it to the DataViewId field.
func (o *CreateDataViewSuccessResp) SetDataViewId(v string) {
	o.DataViewId = &v
}

func (o CreateDataViewSuccessResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDataViewSuccessResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DataViewId) {
		toSerialize["dataViewId"] = o.DataViewId
	}
	return toSerialize, nil
}

type NullableCreateDataViewSuccessResp struct {
	value *CreateDataViewSuccessResp
	isSet bool
}

func (v NullableCreateDataViewSuccessResp) Get() *CreateDataViewSuccessResp {
	return v.value
}

func (v *NullableCreateDataViewSuccessResp) Set(val *CreateDataViewSuccessResp) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDataViewSuccessResp) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDataViewSuccessResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDataViewSuccessResp(val *CreateDataViewSuccessResp) *NullableCreateDataViewSuccessResp {
	return &NullableCreateDataViewSuccessResp{value: val, isSet: true}
}

func (v NullableCreateDataViewSuccessResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDataViewSuccessResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


