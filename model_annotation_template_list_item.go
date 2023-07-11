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

// checks if the AnnotationTemplateListItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnnotationTemplateListItem{}

// AnnotationTemplateListItem struct for AnnotationTemplateListItem
type AnnotationTemplateListItem struct {
	// the id of the annotation template
	Id *string `json:"id,omitempty"`
	// name of the annotation template
	Name *string `json:"name,omitempty"`
	// Unix timestamp in ms
	CreateAt *int64 `json:"createAt,omitempty"`
	// the type of the annotation template
	Type *string `json:"type,omitempty"`
	// the number of labels annotation template has
	LabelCount *int32 `json:"labelCount,omitempty"`
}

// NewAnnotationTemplateListItem instantiates a new AnnotationTemplateListItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnnotationTemplateListItem() *AnnotationTemplateListItem {
	this := AnnotationTemplateListItem{}
	return &this
}

// NewAnnotationTemplateListItemWithDefaults instantiates a new AnnotationTemplateListItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnnotationTemplateListItemWithDefaults() *AnnotationTemplateListItem {
	this := AnnotationTemplateListItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AnnotationTemplateListItem) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotationTemplateListItem) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AnnotationTemplateListItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AnnotationTemplateListItem) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AnnotationTemplateListItem) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotationTemplateListItem) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AnnotationTemplateListItem) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AnnotationTemplateListItem) SetName(v string) {
	o.Name = &v
}

// GetCreateAt returns the CreateAt field value if set, zero value otherwise.
func (o *AnnotationTemplateListItem) GetCreateAt() int64 {
	if o == nil || IsNil(o.CreateAt) {
		var ret int64
		return ret
	}
	return *o.CreateAt
}

// GetCreateAtOk returns a tuple with the CreateAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotationTemplateListItem) GetCreateAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateAt) {
		return nil, false
	}
	return o.CreateAt, true
}

// HasCreateAt returns a boolean if a field has been set.
func (o *AnnotationTemplateListItem) HasCreateAt() bool {
	if o != nil && !IsNil(o.CreateAt) {
		return true
	}

	return false
}

// SetCreateAt gets a reference to the given int64 and assigns it to the CreateAt field.
func (o *AnnotationTemplateListItem) SetCreateAt(v int64) {
	o.CreateAt = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AnnotationTemplateListItem) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotationTemplateListItem) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AnnotationTemplateListItem) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AnnotationTemplateListItem) SetType(v string) {
	o.Type = &v
}

// GetLabelCount returns the LabelCount field value if set, zero value otherwise.
func (o *AnnotationTemplateListItem) GetLabelCount() int32 {
	if o == nil || IsNil(o.LabelCount) {
		var ret int32
		return ret
	}
	return *o.LabelCount
}

// GetLabelCountOk returns a tuple with the LabelCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotationTemplateListItem) GetLabelCountOk() (*int32, bool) {
	if o == nil || IsNil(o.LabelCount) {
		return nil, false
	}
	return o.LabelCount, true
}

// HasLabelCount returns a boolean if a field has been set.
func (o *AnnotationTemplateListItem) HasLabelCount() bool {
	if o != nil && !IsNil(o.LabelCount) {
		return true
	}

	return false
}

// SetLabelCount gets a reference to the given int32 and assigns it to the LabelCount field.
func (o *AnnotationTemplateListItem) SetLabelCount(v int32) {
	o.LabelCount = &v
}

func (o AnnotationTemplateListItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnnotationTemplateListItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.CreateAt) {
		toSerialize["createAt"] = o.CreateAt
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.LabelCount) {
		toSerialize["labelCount"] = o.LabelCount
	}
	return toSerialize, nil
}

type NullableAnnotationTemplateListItem struct {
	value *AnnotationTemplateListItem
	isSet bool
}

func (v NullableAnnotationTemplateListItem) Get() *AnnotationTemplateListItem {
	return v.value
}

func (v *NullableAnnotationTemplateListItem) Set(val *AnnotationTemplateListItem) {
	v.value = val
	v.isSet = true
}

func (v NullableAnnotationTemplateListItem) IsSet() bool {
	return v.isSet
}

func (v *NullableAnnotationTemplateListItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnnotationTemplateListItem(val *AnnotationTemplateListItem) *NullableAnnotationTemplateListItem {
	return &NullableAnnotationTemplateListItem{value: val, isSet: true}
}

func (v NullableAnnotationTemplateListItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnnotationTemplateListItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


