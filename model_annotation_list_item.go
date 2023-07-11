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

// checks if the AnnotationListItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnnotationListItem{}

// AnnotationListItem the result of get annotation list
type AnnotationListItem struct {
	// the id of the raw data
	RawDataId *string `json:"rawDataId,omitempty"`
	// the id of the annotation data item
	DataItemId *string `json:"dataItemId,omitempty"`
	Url *string `json:"url,omitempty"`
	Labels []string `json:"labels,omitempty"`
}

// NewAnnotationListItem instantiates a new AnnotationListItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnnotationListItem() *AnnotationListItem {
	this := AnnotationListItem{}
	return &this
}

// NewAnnotationListItemWithDefaults instantiates a new AnnotationListItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnnotationListItemWithDefaults() *AnnotationListItem {
	this := AnnotationListItem{}
	return &this
}

// GetRawDataId returns the RawDataId field value if set, zero value otherwise.
func (o *AnnotationListItem) GetRawDataId() string {
	if o == nil || IsNil(o.RawDataId) {
		var ret string
		return ret
	}
	return *o.RawDataId
}

// GetRawDataIdOk returns a tuple with the RawDataId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotationListItem) GetRawDataIdOk() (*string, bool) {
	if o == nil || IsNil(o.RawDataId) {
		return nil, false
	}
	return o.RawDataId, true
}

// HasRawDataId returns a boolean if a field has been set.
func (o *AnnotationListItem) HasRawDataId() bool {
	if o != nil && !IsNil(o.RawDataId) {
		return true
	}

	return false
}

// SetRawDataId gets a reference to the given string and assigns it to the RawDataId field.
func (o *AnnotationListItem) SetRawDataId(v string) {
	o.RawDataId = &v
}

// GetDataItemId returns the DataItemId field value if set, zero value otherwise.
func (o *AnnotationListItem) GetDataItemId() string {
	if o == nil || IsNil(o.DataItemId) {
		var ret string
		return ret
	}
	return *o.DataItemId
}

// GetDataItemIdOk returns a tuple with the DataItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotationListItem) GetDataItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.DataItemId) {
		return nil, false
	}
	return o.DataItemId, true
}

// HasDataItemId returns a boolean if a field has been set.
func (o *AnnotationListItem) HasDataItemId() bool {
	if o != nil && !IsNil(o.DataItemId) {
		return true
	}

	return false
}

// SetDataItemId gets a reference to the given string and assigns it to the DataItemId field.
func (o *AnnotationListItem) SetDataItemId(v string) {
	o.DataItemId = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AnnotationListItem) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotationListItem) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AnnotationListItem) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AnnotationListItem) SetUrl(v string) {
	o.Url = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *AnnotationListItem) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotationListItem) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *AnnotationListItem) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *AnnotationListItem) SetLabels(v []string) {
	o.Labels = v
}

func (o AnnotationListItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnnotationListItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RawDataId) {
		toSerialize["rawDataId"] = o.RawDataId
	}
	if !IsNil(o.DataItemId) {
		toSerialize["dataItemId"] = o.DataItemId
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	return toSerialize, nil
}

type NullableAnnotationListItem struct {
	value *AnnotationListItem
	isSet bool
}

func (v NullableAnnotationListItem) Get() *AnnotationListItem {
	return v.value
}

func (v *NullableAnnotationListItem) Set(val *AnnotationListItem) {
	v.value = val
	v.isSet = true
}

func (v NullableAnnotationListItem) IsSet() bool {
	return v.isSet
}

func (v *NullableAnnotationListItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnnotationListItem(val *AnnotationListItem) *NullableAnnotationListItem {
	return &NullableAnnotationListItem{value: val, isSet: true}
}

func (v NullableAnnotationListItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnnotationListItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

