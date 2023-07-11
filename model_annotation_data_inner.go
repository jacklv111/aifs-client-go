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

// checks if the AnnotationDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnnotationDataInner{}

// AnnotationDataInner struct for AnnotationDataInner
type AnnotationDataInner struct {
	DataItemId *string `json:"dataItemId,omitempty"`
	Labels []string `json:"labels,omitempty"`
	// text format annotation data
	TextData *string `json:"textData,omitempty"`
}

// NewAnnotationDataInner instantiates a new AnnotationDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnnotationDataInner() *AnnotationDataInner {
	this := AnnotationDataInner{}
	return &this
}

// NewAnnotationDataInnerWithDefaults instantiates a new AnnotationDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnnotationDataInnerWithDefaults() *AnnotationDataInner {
	this := AnnotationDataInner{}
	return &this
}

// GetDataItemId returns the DataItemId field value if set, zero value otherwise.
func (o *AnnotationDataInner) GetDataItemId() string {
	if o == nil || IsNil(o.DataItemId) {
		var ret string
		return ret
	}
	return *o.DataItemId
}

// GetDataItemIdOk returns a tuple with the DataItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotationDataInner) GetDataItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.DataItemId) {
		return nil, false
	}
	return o.DataItemId, true
}

// HasDataItemId returns a boolean if a field has been set.
func (o *AnnotationDataInner) HasDataItemId() bool {
	if o != nil && !IsNil(o.DataItemId) {
		return true
	}

	return false
}

// SetDataItemId gets a reference to the given string and assigns it to the DataItemId field.
func (o *AnnotationDataInner) SetDataItemId(v string) {
	o.DataItemId = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *AnnotationDataInner) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotationDataInner) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *AnnotationDataInner) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *AnnotationDataInner) SetLabels(v []string) {
	o.Labels = v
}

// GetTextData returns the TextData field value if set, zero value otherwise.
func (o *AnnotationDataInner) GetTextData() string {
	if o == nil || IsNil(o.TextData) {
		var ret string
		return ret
	}
	return *o.TextData
}

// GetTextDataOk returns a tuple with the TextData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotationDataInner) GetTextDataOk() (*string, bool) {
	if o == nil || IsNil(o.TextData) {
		return nil, false
	}
	return o.TextData, true
}

// HasTextData returns a boolean if a field has been set.
func (o *AnnotationDataInner) HasTextData() bool {
	if o != nil && !IsNil(o.TextData) {
		return true
	}

	return false
}

// SetTextData gets a reference to the given string and assigns it to the TextData field.
func (o *AnnotationDataInner) SetTextData(v string) {
	o.TextData = &v
}

func (o AnnotationDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnnotationDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DataItemId) {
		toSerialize["dataItemId"] = o.DataItemId
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.TextData) {
		toSerialize["textData"] = o.TextData
	}
	return toSerialize, nil
}

type NullableAnnotationDataInner struct {
	value *AnnotationDataInner
	isSet bool
}

func (v NullableAnnotationDataInner) Get() *AnnotationDataInner {
	return v.value
}

func (v *NullableAnnotationDataInner) Set(val *AnnotationDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAnnotationDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAnnotationDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnnotationDataInner(val *AnnotationDataInner) *NullableAnnotationDataInner {
	return &NullableAnnotationDataInner{value: val, isSet: true}
}

func (v NullableAnnotationDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnnotationDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


