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

// checks if the DataViewStatistics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataViewStatistics{}

// DataViewStatistics struct for DataViewStatistics
type DataViewStatistics struct {
	ItemCount *int32 `json:"itemCount,omitempty"`
	// the number of labels in the annotation data view.
	LabelCount *int32 `json:"labelCount,omitempty"`
	// the distribution of labels in the annotation data view.
	LabelDistribution []LabelDistribution `json:"labelDistribution,omitempty"`
	// the total size of the data in the data view in bytes.
	TotalDataSize *int64 `json:"totalDataSize,omitempty"`
}

// NewDataViewStatistics instantiates a new DataViewStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataViewStatistics() *DataViewStatistics {
	this := DataViewStatistics{}
	return &this
}

// NewDataViewStatisticsWithDefaults instantiates a new DataViewStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataViewStatisticsWithDefaults() *DataViewStatistics {
	this := DataViewStatistics{}
	return &this
}

// GetItemCount returns the ItemCount field value if set, zero value otherwise.
func (o *DataViewStatistics) GetItemCount() int32 {
	if o == nil || IsNil(o.ItemCount) {
		var ret int32
		return ret
	}
	return *o.ItemCount
}

// GetItemCountOk returns a tuple with the ItemCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataViewStatistics) GetItemCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ItemCount) {
		return nil, false
	}
	return o.ItemCount, true
}

// HasItemCount returns a boolean if a field has been set.
func (o *DataViewStatistics) HasItemCount() bool {
	if o != nil && !IsNil(o.ItemCount) {
		return true
	}

	return false
}

// SetItemCount gets a reference to the given int32 and assigns it to the ItemCount field.
func (o *DataViewStatistics) SetItemCount(v int32) {
	o.ItemCount = &v
}

// GetLabelCount returns the LabelCount field value if set, zero value otherwise.
func (o *DataViewStatistics) GetLabelCount() int32 {
	if o == nil || IsNil(o.LabelCount) {
		var ret int32
		return ret
	}
	return *o.LabelCount
}

// GetLabelCountOk returns a tuple with the LabelCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataViewStatistics) GetLabelCountOk() (*int32, bool) {
	if o == nil || IsNil(o.LabelCount) {
		return nil, false
	}
	return o.LabelCount, true
}

// HasLabelCount returns a boolean if a field has been set.
func (o *DataViewStatistics) HasLabelCount() bool {
	if o != nil && !IsNil(o.LabelCount) {
		return true
	}

	return false
}

// SetLabelCount gets a reference to the given int32 and assigns it to the LabelCount field.
func (o *DataViewStatistics) SetLabelCount(v int32) {
	o.LabelCount = &v
}

// GetLabelDistribution returns the LabelDistribution field value if set, zero value otherwise.
func (o *DataViewStatistics) GetLabelDistribution() []LabelDistribution {
	if o == nil || IsNil(o.LabelDistribution) {
		var ret []LabelDistribution
		return ret
	}
	return o.LabelDistribution
}

// GetLabelDistributionOk returns a tuple with the LabelDistribution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataViewStatistics) GetLabelDistributionOk() ([]LabelDistribution, bool) {
	if o == nil || IsNil(o.LabelDistribution) {
		return nil, false
	}
	return o.LabelDistribution, true
}

// HasLabelDistribution returns a boolean if a field has been set.
func (o *DataViewStatistics) HasLabelDistribution() bool {
	if o != nil && !IsNil(o.LabelDistribution) {
		return true
	}

	return false
}

// SetLabelDistribution gets a reference to the given []LabelDistribution and assigns it to the LabelDistribution field.
func (o *DataViewStatistics) SetLabelDistribution(v []LabelDistribution) {
	o.LabelDistribution = v
}

// GetTotalDataSize returns the TotalDataSize field value if set, zero value otherwise.
func (o *DataViewStatistics) GetTotalDataSize() int64 {
	if o == nil || IsNil(o.TotalDataSize) {
		var ret int64
		return ret
	}
	return *o.TotalDataSize
}

// GetTotalDataSizeOk returns a tuple with the TotalDataSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataViewStatistics) GetTotalDataSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalDataSize) {
		return nil, false
	}
	return o.TotalDataSize, true
}

// HasTotalDataSize returns a boolean if a field has been set.
func (o *DataViewStatistics) HasTotalDataSize() bool {
	if o != nil && !IsNil(o.TotalDataSize) {
		return true
	}

	return false
}

// SetTotalDataSize gets a reference to the given int64 and assigns it to the TotalDataSize field.
func (o *DataViewStatistics) SetTotalDataSize(v int64) {
	o.TotalDataSize = &v
}

func (o DataViewStatistics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataViewStatistics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ItemCount) {
		toSerialize["itemCount"] = o.ItemCount
	}
	if !IsNil(o.LabelCount) {
		toSerialize["labelCount"] = o.LabelCount
	}
	if !IsNil(o.LabelDistribution) {
		toSerialize["labelDistribution"] = o.LabelDistribution
	}
	if !IsNil(o.TotalDataSize) {
		toSerialize["totalDataSize"] = o.TotalDataSize
	}
	return toSerialize, nil
}

type NullableDataViewStatistics struct {
	value *DataViewStatistics
	isSet bool
}

func (v NullableDataViewStatistics) Get() *DataViewStatistics {
	return v.value
}

func (v *NullableDataViewStatistics) Set(val *DataViewStatistics) {
	v.value = val
	v.isSet = true
}

func (v NullableDataViewStatistics) IsSet() bool {
	return v.isSet
}

func (v *NullableDataViewStatistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataViewStatistics(val *DataViewStatistics) *NullableDataViewStatistics {
	return &NullableDataViewStatistics{value: val, isSet: true}
}

func (v NullableDataViewStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataViewStatistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

