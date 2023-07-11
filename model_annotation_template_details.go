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

// checks if the AnnotationTemplateDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnnotationTemplateDetails{}

// AnnotationTemplateDetails struct for AnnotationTemplateDetails
type AnnotationTemplateDetails struct {
	// the id of the annotation template
	Id *string `json:"id,omitempty"`
	// name of the annotation template
	Name *string `json:"name,omitempty"`
	// Unix timestamp in ms
	CreateAt *int64 `json:"createAt,omitempty"`
	// Unix timestamp in ms
	UpdateAt *int64 `json:"updateAt,omitempty"`
	// the type of the annotation template
	Type *string `json:"type,omitempty"`
	Labels []Label `json:"labels,omitempty"`
	WordList []string `json:"wordList,omitempty"`
}

// NewAnnotationTemplateDetails instantiates a new AnnotationTemplateDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnnotationTemplateDetails() *AnnotationTemplateDetails {
	this := AnnotationTemplateDetails{}
	return &this
}

// NewAnnotationTemplateDetailsWithDefaults instantiates a new AnnotationTemplateDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnnotationTemplateDetailsWithDefaults() *AnnotationTemplateDetails {
	this := AnnotationTemplateDetails{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AnnotationTemplateDetails) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotationTemplateDetails) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AnnotationTemplateDetails) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AnnotationTemplateDetails) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AnnotationTemplateDetails) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotationTemplateDetails) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AnnotationTemplateDetails) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AnnotationTemplateDetails) SetName(v string) {
	o.Name = &v
}

// GetCreateAt returns the CreateAt field value if set, zero value otherwise.
func (o *AnnotationTemplateDetails) GetCreateAt() int64 {
	if o == nil || IsNil(o.CreateAt) {
		var ret int64
		return ret
	}
	return *o.CreateAt
}

// GetCreateAtOk returns a tuple with the CreateAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotationTemplateDetails) GetCreateAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateAt) {
		return nil, false
	}
	return o.CreateAt, true
}

// HasCreateAt returns a boolean if a field has been set.
func (o *AnnotationTemplateDetails) HasCreateAt() bool {
	if o != nil && !IsNil(o.CreateAt) {
		return true
	}

	return false
}

// SetCreateAt gets a reference to the given int64 and assigns it to the CreateAt field.
func (o *AnnotationTemplateDetails) SetCreateAt(v int64) {
	o.CreateAt = &v
}

// GetUpdateAt returns the UpdateAt field value if set, zero value otherwise.
func (o *AnnotationTemplateDetails) GetUpdateAt() int64 {
	if o == nil || IsNil(o.UpdateAt) {
		var ret int64
		return ret
	}
	return *o.UpdateAt
}

// GetUpdateAtOk returns a tuple with the UpdateAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotationTemplateDetails) GetUpdateAtOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdateAt) {
		return nil, false
	}
	return o.UpdateAt, true
}

// HasUpdateAt returns a boolean if a field has been set.
func (o *AnnotationTemplateDetails) HasUpdateAt() bool {
	if o != nil && !IsNil(o.UpdateAt) {
		return true
	}

	return false
}

// SetUpdateAt gets a reference to the given int64 and assigns it to the UpdateAt field.
func (o *AnnotationTemplateDetails) SetUpdateAt(v int64) {
	o.UpdateAt = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AnnotationTemplateDetails) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotationTemplateDetails) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AnnotationTemplateDetails) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AnnotationTemplateDetails) SetType(v string) {
	o.Type = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *AnnotationTemplateDetails) GetLabels() []Label {
	if o == nil || IsNil(o.Labels) {
		var ret []Label
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotationTemplateDetails) GetLabelsOk() ([]Label, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *AnnotationTemplateDetails) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []Label and assigns it to the Labels field.
func (o *AnnotationTemplateDetails) SetLabels(v []Label) {
	o.Labels = v
}

// GetWordList returns the WordList field value if set, zero value otherwise.
func (o *AnnotationTemplateDetails) GetWordList() []string {
	if o == nil || IsNil(o.WordList) {
		var ret []string
		return ret
	}
	return o.WordList
}

// GetWordListOk returns a tuple with the WordList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotationTemplateDetails) GetWordListOk() ([]string, bool) {
	if o == nil || IsNil(o.WordList) {
		return nil, false
	}
	return o.WordList, true
}

// HasWordList returns a boolean if a field has been set.
func (o *AnnotationTemplateDetails) HasWordList() bool {
	if o != nil && !IsNil(o.WordList) {
		return true
	}

	return false
}

// SetWordList gets a reference to the given []string and assigns it to the WordList field.
func (o *AnnotationTemplateDetails) SetWordList(v []string) {
	o.WordList = v
}

func (o AnnotationTemplateDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnnotationTemplateDetails) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.UpdateAt) {
		toSerialize["updateAt"] = o.UpdateAt
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.WordList) {
		toSerialize["wordList"] = o.WordList
	}
	return toSerialize, nil
}

type NullableAnnotationTemplateDetails struct {
	value *AnnotationTemplateDetails
	isSet bool
}

func (v NullableAnnotationTemplateDetails) Get() *AnnotationTemplateDetails {
	return v.value
}

func (v *NullableAnnotationTemplateDetails) Set(val *AnnotationTemplateDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableAnnotationTemplateDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableAnnotationTemplateDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnnotationTemplateDetails(val *AnnotationTemplateDetails) *NullableAnnotationTemplateDetails {
	return &NullableAnnotationTemplateDetails{value: val, isSet: true}
}

func (v NullableAnnotationTemplateDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnnotationTemplateDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


