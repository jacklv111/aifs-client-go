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

// checks if the CreateAnnoTemplateSuccessResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAnnoTemplateSuccessResp{}

// CreateAnnoTemplateSuccessResp struct for CreateAnnoTemplateSuccessResp
type CreateAnnoTemplateSuccessResp struct {
	AnnotationTemplateId *string `json:"annotationTemplateId,omitempty"`
}

// NewCreateAnnoTemplateSuccessResp instantiates a new CreateAnnoTemplateSuccessResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAnnoTemplateSuccessResp() *CreateAnnoTemplateSuccessResp {
	this := CreateAnnoTemplateSuccessResp{}
	return &this
}

// NewCreateAnnoTemplateSuccessRespWithDefaults instantiates a new CreateAnnoTemplateSuccessResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAnnoTemplateSuccessRespWithDefaults() *CreateAnnoTemplateSuccessResp {
	this := CreateAnnoTemplateSuccessResp{}
	return &this
}

// GetAnnotationTemplateId returns the AnnotationTemplateId field value if set, zero value otherwise.
func (o *CreateAnnoTemplateSuccessResp) GetAnnotationTemplateId() string {
	if o == nil || IsNil(o.AnnotationTemplateId) {
		var ret string
		return ret
	}
	return *o.AnnotationTemplateId
}

// GetAnnotationTemplateIdOk returns a tuple with the AnnotationTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAnnoTemplateSuccessResp) GetAnnotationTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.AnnotationTemplateId) {
		return nil, false
	}
	return o.AnnotationTemplateId, true
}

// HasAnnotationTemplateId returns a boolean if a field has been set.
func (o *CreateAnnoTemplateSuccessResp) HasAnnotationTemplateId() bool {
	if o != nil && !IsNil(o.AnnotationTemplateId) {
		return true
	}

	return false
}

// SetAnnotationTemplateId gets a reference to the given string and assigns it to the AnnotationTemplateId field.
func (o *CreateAnnoTemplateSuccessResp) SetAnnotationTemplateId(v string) {
	o.AnnotationTemplateId = &v
}

func (o CreateAnnoTemplateSuccessResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAnnoTemplateSuccessResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AnnotationTemplateId) {
		toSerialize["annotationTemplateId"] = o.AnnotationTemplateId
	}
	return toSerialize, nil
}

type NullableCreateAnnoTemplateSuccessResp struct {
	value *CreateAnnoTemplateSuccessResp
	isSet bool
}

func (v NullableCreateAnnoTemplateSuccessResp) Get() *CreateAnnoTemplateSuccessResp {
	return v.value
}

func (v *NullableCreateAnnoTemplateSuccessResp) Set(val *CreateAnnoTemplateSuccessResp) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAnnoTemplateSuccessResp) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAnnoTemplateSuccessResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAnnoTemplateSuccessResp(val *CreateAnnoTemplateSuccessResp) *NullableCreateAnnoTemplateSuccessResp {
	return &NullableCreateAnnoTemplateSuccessResp{value: val, isSet: true}
}

func (v NullableCreateAnnoTemplateSuccessResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAnnoTemplateSuccessResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


