# MergeDataViewsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | name of the merged data view | [optional] 
**Description** | Pointer to **string** | description of the merged data view | [optional] 
**DataViewIdList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewMergeDataViewsRequest

`func NewMergeDataViewsRequest() *MergeDataViewsRequest`

NewMergeDataViewsRequest instantiates a new MergeDataViewsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMergeDataViewsRequestWithDefaults

`func NewMergeDataViewsRequestWithDefaults() *MergeDataViewsRequest`

NewMergeDataViewsRequestWithDefaults instantiates a new MergeDataViewsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MergeDataViewsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MergeDataViewsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MergeDataViewsRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MergeDataViewsRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *MergeDataViewsRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MergeDataViewsRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MergeDataViewsRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MergeDataViewsRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDataViewIdList

`func (o *MergeDataViewsRequest) GetDataViewIdList() []string`

GetDataViewIdList returns the DataViewIdList field if non-nil, zero value otherwise.

### GetDataViewIdListOk

`func (o *MergeDataViewsRequest) GetDataViewIdListOk() (*[]string, bool)`

GetDataViewIdListOk returns a tuple with the DataViewIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataViewIdList

`func (o *MergeDataViewsRequest) SetDataViewIdList(v []string)`

SetDataViewIdList sets DataViewIdList field to given value.

### HasDataViewIdList

`func (o *MergeDataViewsRequest) HasDataViewIdList() bool`

HasDataViewIdList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


