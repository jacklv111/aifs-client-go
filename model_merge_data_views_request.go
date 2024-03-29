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

// checks if the MergeDataViewsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MergeDataViewsRequest{}

// MergeDataViewsRequest struct for MergeDataViewsRequest
type MergeDataViewsRequest struct {
	// name of the merged data view
	Name *string `json:"name,omitempty"`
	// description of the merged data view
	Description *string `json:"description,omitempty"`
	DataViewIdList []string `json:"dataViewIdList,omitempty"`
}

// NewMergeDataViewsRequest instantiates a new MergeDataViewsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMergeDataViewsRequest() *MergeDataViewsRequest {
	this := MergeDataViewsRequest{}
	return &this
}

// NewMergeDataViewsRequestWithDefaults instantiates a new MergeDataViewsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMergeDataViewsRequestWithDefaults() *MergeDataViewsRequest {
	this := MergeDataViewsRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MergeDataViewsRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeDataViewsRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MergeDataViewsRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MergeDataViewsRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MergeDataViewsRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeDataViewsRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MergeDataViewsRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MergeDataViewsRequest) SetDescription(v string) {
	o.Description = &v
}

// GetDataViewIdList returns the DataViewIdList field value if set, zero value otherwise.
func (o *MergeDataViewsRequest) GetDataViewIdList() []string {
	if o == nil || IsNil(o.DataViewIdList) {
		var ret []string
		return ret
	}
	return o.DataViewIdList
}

// GetDataViewIdListOk returns a tuple with the DataViewIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeDataViewsRequest) GetDataViewIdListOk() ([]string, bool) {
	if o == nil || IsNil(o.DataViewIdList) {
		return nil, false
	}
	return o.DataViewIdList, true
}

// HasDataViewIdList returns a boolean if a field has been set.
func (o *MergeDataViewsRequest) HasDataViewIdList() bool {
	if o != nil && !IsNil(o.DataViewIdList) {
		return true
	}

	return false
}

// SetDataViewIdList gets a reference to the given []string and assigns it to the DataViewIdList field.
func (o *MergeDataViewsRequest) SetDataViewIdList(v []string) {
	o.DataViewIdList = v
}

func (o MergeDataViewsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MergeDataViewsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DataViewIdList) {
		toSerialize["dataViewIdList"] = o.DataViewIdList
	}
	return toSerialize, nil
}

type NullableMergeDataViewsRequest struct {
	value *MergeDataViewsRequest
	isSet bool
}

func (v NullableMergeDataViewsRequest) Get() *MergeDataViewsRequest {
	return v.value
}

func (v *NullableMergeDataViewsRequest) Set(val *MergeDataViewsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMergeDataViewsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMergeDataViewsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMergeDataViewsRequest(val *MergeDataViewsRequest) *NullableMergeDataViewsRequest {
	return &NullableMergeDataViewsRequest{value: val, isSet: true}
}

func (v NullableMergeDataViewsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMergeDataViewsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


