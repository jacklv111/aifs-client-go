# DataViewStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemCount** | Pointer to **int32** |  | [optional] 
**LabelCount** | Pointer to **int32** | the number of labels in the annotation data view. | [optional] 
**LabelDistribution** | Pointer to [**[]LabelDistribution**](LabelDistribution.md) | the distribution of labels in the annotation data view. | [optional] 
**TotalDataSize** | Pointer to **int64** | the total size of the data in the data view in bytes. | [optional] 

## Methods

### NewDataViewStatistics

`func NewDataViewStatistics() *DataViewStatistics`

NewDataViewStatistics instantiates a new DataViewStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataViewStatisticsWithDefaults

`func NewDataViewStatisticsWithDefaults() *DataViewStatistics`

NewDataViewStatisticsWithDefaults instantiates a new DataViewStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemCount

`func (o *DataViewStatistics) GetItemCount() int32`

GetItemCount returns the ItemCount field if non-nil, zero value otherwise.

### GetItemCountOk

`func (o *DataViewStatistics) GetItemCountOk() (*int32, bool)`

GetItemCountOk returns a tuple with the ItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCount

`func (o *DataViewStatistics) SetItemCount(v int32)`

SetItemCount sets ItemCount field to given value.

### HasItemCount

`func (o *DataViewStatistics) HasItemCount() bool`

HasItemCount returns a boolean if a field has been set.

### GetLabelCount

`func (o *DataViewStatistics) GetLabelCount() int32`

GetLabelCount returns the LabelCount field if non-nil, zero value otherwise.

### GetLabelCountOk

`func (o *DataViewStatistics) GetLabelCountOk() (*int32, bool)`

GetLabelCountOk returns a tuple with the LabelCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelCount

`func (o *DataViewStatistics) SetLabelCount(v int32)`

SetLabelCount sets LabelCount field to given value.

### HasLabelCount

`func (o *DataViewStatistics) HasLabelCount() bool`

HasLabelCount returns a boolean if a field has been set.

### GetLabelDistribution

`func (o *DataViewStatistics) GetLabelDistribution() []LabelDistribution`

GetLabelDistribution returns the LabelDistribution field if non-nil, zero value otherwise.

### GetLabelDistributionOk

`func (o *DataViewStatistics) GetLabelDistributionOk() (*[]LabelDistribution, bool)`

GetLabelDistributionOk returns a tuple with the LabelDistribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelDistribution

`func (o *DataViewStatistics) SetLabelDistribution(v []LabelDistribution)`

SetLabelDistribution sets LabelDistribution field to given value.

### HasLabelDistribution

`func (o *DataViewStatistics) HasLabelDistribution() bool`

HasLabelDistribution returns a boolean if a field has been set.

### GetTotalDataSize

`func (o *DataViewStatistics) GetTotalDataSize() int64`

GetTotalDataSize returns the TotalDataSize field if non-nil, zero value otherwise.

### GetTotalDataSizeOk

`func (o *DataViewStatistics) GetTotalDataSizeOk() (*int64, bool)`

GetTotalDataSizeOk returns a tuple with the TotalDataSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDataSize

`func (o *DataViewStatistics) SetTotalDataSize(v int64)`

SetTotalDataSize sets TotalDataSize field to given value.

### HasTotalDataSize

`func (o *DataViewStatistics) HasTotalDataSize() bool`

HasTotalDataSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


