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

// checks if the AnnotationViewData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnnotationViewData{}

// AnnotationViewData struct for AnnotationViewData
type AnnotationViewData struct {
	DataViewId *string `json:"dataViewId,omitempty"`
	AnnotationTemplateId *string `json:"annotationTemplateId,omitempty"`
	ViewType *DataViewType `json:"viewType,omitempty"`
	DataItems []AnnotationDataInner `json:"dataItems,omitempty"`
}

// NewAnnotationViewData instantiates a new AnnotationViewData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnnotationViewData() *AnnotationViewData {
	this := AnnotationViewData{}
	return &this
}

// NewAnnotationViewDataWithDefaults instantiates a new AnnotationViewData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnnotationViewDataWithDefaults() *AnnotationViewData {
	this := AnnotationViewData{}
	return &this
}

// GetDataViewId returns the DataViewId field value if set, zero value otherwise.
func (o *AnnotationViewData) GetDataViewId() string {
	if o == nil || IsNil(o.DataViewId) {
		var ret string
		return ret
	}
	return *o.DataViewId
}

// GetDataViewIdOk returns a tuple with the DataViewId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotationViewData) GetDataViewIdOk() (*string, bool) {
	if o == nil || IsNil(o.DataViewId) {
		return nil, false
	}
	return o.DataViewId, true
}

// HasDataViewId returns a boolean if a field has been set.
func (o *AnnotationViewData) HasDataViewId() bool {
	if o != nil && !IsNil(o.DataViewId) {
		return true
	}

	return false
}

// SetDataViewId gets a reference to the given string and assigns it to the DataViewId field.
func (o *AnnotationViewData) SetDataViewId(v string) {
	o.DataViewId = &v
}

// GetAnnotationTemplateId returns the AnnotationTemplateId field value if set, zero value otherwise.
func (o *AnnotationViewData) GetAnnotationTemplateId() string {
	if o == nil || IsNil(o.AnnotationTemplateId) {
		var ret string
		return ret
	}
	return *o.AnnotationTemplateId
}

// GetAnnotationTemplateIdOk returns a tuple with the AnnotationTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotationViewData) GetAnnotationTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.AnnotationTemplateId) {
		return nil, false
	}
	return o.AnnotationTemplateId, true
}

// HasAnnotationTemplateId returns a boolean if a field has been set.
func (o *AnnotationViewData) HasAnnotationTemplateId() bool {
	if o != nil && !IsNil(o.AnnotationTemplateId) {
		return true
	}

	return false
}

// SetAnnotationTemplateId gets a reference to the given string and assigns it to the AnnotationTemplateId field.
func (o *AnnotationViewData) SetAnnotationTemplateId(v string) {
	o.AnnotationTemplateId = &v
}

// GetViewType returns the ViewType field value if set, zero value otherwise.
func (o *AnnotationViewData) GetViewType() DataViewType {
	if o == nil || IsNil(o.ViewType) {
		var ret DataViewType
		return ret
	}
	return *o.ViewType
}

// GetViewTypeOk returns a tuple with the ViewType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotationViewData) GetViewTypeOk() (*DataViewType, bool) {
	if o == nil || IsNil(o.ViewType) {
		return nil, false
	}
	return o.ViewType, true
}

// HasViewType returns a boolean if a field has been set.
func (o *AnnotationViewData) HasViewType() bool {
	if o != nil && !IsNil(o.ViewType) {
		return true
	}

	return false
}

// SetViewType gets a reference to the given DataViewType and assigns it to the ViewType field.
func (o *AnnotationViewData) SetViewType(v DataViewType) {
	o.ViewType = &v
}

// GetDataItems returns the DataItems field value if set, zero value otherwise.
func (o *AnnotationViewData) GetDataItems() []AnnotationDataInner {
	if o == nil || IsNil(o.DataItems) {
		var ret []AnnotationDataInner
		return ret
	}
	return o.DataItems
}

// GetDataItemsOk returns a tuple with the DataItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotationViewData) GetDataItemsOk() ([]AnnotationDataInner, bool) {
	if o == nil || IsNil(o.DataItems) {
		return nil, false
	}
	return o.DataItems, true
}

// HasDataItems returns a boolean if a field has been set.
func (o *AnnotationViewData) HasDataItems() bool {
	if o != nil && !IsNil(o.DataItems) {
		return true
	}

	return false
}

// SetDataItems gets a reference to the given []AnnotationDataInner and assigns it to the DataItems field.
func (o *AnnotationViewData) SetDataItems(v []AnnotationDataInner) {
	o.DataItems = v
}

func (o AnnotationViewData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnnotationViewData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DataViewId) {
		toSerialize["dataViewId"] = o.DataViewId
	}
	if !IsNil(o.AnnotationTemplateId) {
		toSerialize["annotationTemplateId"] = o.AnnotationTemplateId
	}
	if !IsNil(o.ViewType) {
		toSerialize["viewType"] = o.ViewType
	}
	if !IsNil(o.DataItems) {
		toSerialize["dataItems"] = o.DataItems
	}
	return toSerialize, nil
}

type NullableAnnotationViewData struct {
	value *AnnotationViewData
	isSet bool
}

func (v NullableAnnotationViewData) Get() *AnnotationViewData {
	return v.value
}

func (v *NullableAnnotationViewData) Set(val *AnnotationViewData) {
	v.value = val
	v.isSet = true
}

func (v NullableAnnotationViewData) IsSet() bool {
	return v.isSet
}

func (v *NullableAnnotationViewData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnnotationViewData(val *AnnotationViewData) *NullableAnnotationViewData {
	return &NullableAnnotationViewData{value: val, isSet: true}
}

func (v NullableAnnotationViewData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnnotationViewData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

