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

// checks if the AnnotationViewLocations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnnotationViewLocations{}

// AnnotationViewLocations struct for AnnotationViewLocations
type AnnotationViewLocations struct {
	DataViewId *string `json:"dataViewId,omitempty"`
	AnnotationTemplateId *string `json:"annotationTemplateId,omitempty"`
	ViewType *DataViewType `json:"viewType,omitempty"`
	DataItems []S3StorageInner `json:"dataItems,omitempty"`
}

// NewAnnotationViewLocations instantiates a new AnnotationViewLocations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnnotationViewLocations() *AnnotationViewLocations {
	this := AnnotationViewLocations{}
	return &this
}

// NewAnnotationViewLocationsWithDefaults instantiates a new AnnotationViewLocations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnnotationViewLocationsWithDefaults() *AnnotationViewLocations {
	this := AnnotationViewLocations{}
	return &this
}

// GetDataViewId returns the DataViewId field value if set, zero value otherwise.
func (o *AnnotationViewLocations) GetDataViewId() string {
	if o == nil || IsNil(o.DataViewId) {
		var ret string
		return ret
	}
	return *o.DataViewId
}

// GetDataViewIdOk returns a tuple with the DataViewId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotationViewLocations) GetDataViewIdOk() (*string, bool) {
	if o == nil || IsNil(o.DataViewId) {
		return nil, false
	}
	return o.DataViewId, true
}

// HasDataViewId returns a boolean if a field has been set.
func (o *AnnotationViewLocations) HasDataViewId() bool {
	if o != nil && !IsNil(o.DataViewId) {
		return true
	}

	return false
}

// SetDataViewId gets a reference to the given string and assigns it to the DataViewId field.
func (o *AnnotationViewLocations) SetDataViewId(v string) {
	o.DataViewId = &v
}

// GetAnnotationTemplateId returns the AnnotationTemplateId field value if set, zero value otherwise.
func (o *AnnotationViewLocations) GetAnnotationTemplateId() string {
	if o == nil || IsNil(o.AnnotationTemplateId) {
		var ret string
		return ret
	}
	return *o.AnnotationTemplateId
}

// GetAnnotationTemplateIdOk returns a tuple with the AnnotationTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotationViewLocations) GetAnnotationTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.AnnotationTemplateId) {
		return nil, false
	}
	return o.AnnotationTemplateId, true
}

// HasAnnotationTemplateId returns a boolean if a field has been set.
func (o *AnnotationViewLocations) HasAnnotationTemplateId() bool {
	if o != nil && !IsNil(o.AnnotationTemplateId) {
		return true
	}

	return false
}

// SetAnnotationTemplateId gets a reference to the given string and assigns it to the AnnotationTemplateId field.
func (o *AnnotationViewLocations) SetAnnotationTemplateId(v string) {
	o.AnnotationTemplateId = &v
}

// GetViewType returns the ViewType field value if set, zero value otherwise.
func (o *AnnotationViewLocations) GetViewType() DataViewType {
	if o == nil || IsNil(o.ViewType) {
		var ret DataViewType
		return ret
	}
	return *o.ViewType
}

// GetViewTypeOk returns a tuple with the ViewType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotationViewLocations) GetViewTypeOk() (*DataViewType, bool) {
	if o == nil || IsNil(o.ViewType) {
		return nil, false
	}
	return o.ViewType, true
}

// HasViewType returns a boolean if a field has been set.
func (o *AnnotationViewLocations) HasViewType() bool {
	if o != nil && !IsNil(o.ViewType) {
		return true
	}

	return false
}

// SetViewType gets a reference to the given DataViewType and assigns it to the ViewType field.
func (o *AnnotationViewLocations) SetViewType(v DataViewType) {
	o.ViewType = &v
}

// GetDataItems returns the DataItems field value if set, zero value otherwise.
func (o *AnnotationViewLocations) GetDataItems() []S3StorageInner {
	if o == nil || IsNil(o.DataItems) {
		var ret []S3StorageInner
		return ret
	}
	return o.DataItems
}

// GetDataItemsOk returns a tuple with the DataItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotationViewLocations) GetDataItemsOk() ([]S3StorageInner, bool) {
	if o == nil || IsNil(o.DataItems) {
		return nil, false
	}
	return o.DataItems, true
}

// HasDataItems returns a boolean if a field has been set.
func (o *AnnotationViewLocations) HasDataItems() bool {
	if o != nil && !IsNil(o.DataItems) {
		return true
	}

	return false
}

// SetDataItems gets a reference to the given []S3StorageInner and assigns it to the DataItems field.
func (o *AnnotationViewLocations) SetDataItems(v []S3StorageInner) {
	o.DataItems = v
}

func (o AnnotationViewLocations) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnnotationViewLocations) ToMap() (map[string]interface{}, error) {
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

type NullableAnnotationViewLocations struct {
	value *AnnotationViewLocations
	isSet bool
}

func (v NullableAnnotationViewLocations) Get() *AnnotationViewLocations {
	return v.value
}

func (v *NullableAnnotationViewLocations) Set(val *AnnotationViewLocations) {
	v.value = val
	v.isSet = true
}

func (v NullableAnnotationViewLocations) IsSet() bool {
	return v.isSet
}

func (v *NullableAnnotationViewLocations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnnotationViewLocations(val *AnnotationViewLocations) *NullableAnnotationViewLocations {
	return &NullableAnnotationViewLocations{value: val, isSet: true}
}

func (v NullableAnnotationViewLocations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnnotationViewLocations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


