# DatasetZipLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataViewId** | Pointer to **string** |  | [optional] 
**ViewType** | Pointer to [**DataViewType**](DataViewType.md) |  | [optional] 
**DataItems** | Pointer to [**[]S3StorageInner**](S3StorageInner.md) |  | [optional] 

## Methods

### NewDatasetZipLocation

`func NewDatasetZipLocation() *DatasetZipLocation`

NewDatasetZipLocation instantiates a new DatasetZipLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetZipLocationWithDefaults

`func NewDatasetZipLocationWithDefaults() *DatasetZipLocation`

NewDatasetZipLocationWithDefaults instantiates a new DatasetZipLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataViewId

`func (o *DatasetZipLocation) GetDataViewId() string`

GetDataViewId returns the DataViewId field if non-nil, zero value otherwise.

### GetDataViewIdOk

`func (o *DatasetZipLocation) GetDataViewIdOk() (*string, bool)`

GetDataViewIdOk returns a tuple with the DataViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataViewId

`func (o *DatasetZipLocation) SetDataViewId(v string)`

SetDataViewId sets DataViewId field to given value.

### HasDataViewId

`func (o *DatasetZipLocation) HasDataViewId() bool`

HasDataViewId returns a boolean if a field has been set.

### GetViewType

`func (o *DatasetZipLocation) GetViewType() DataViewType`

GetViewType returns the ViewType field if non-nil, zero value otherwise.

### GetViewTypeOk

`func (o *DatasetZipLocation) GetViewTypeOk() (*DataViewType, bool)`

GetViewTypeOk returns a tuple with the ViewType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewType

`func (o *DatasetZipLocation) SetViewType(v DataViewType)`

SetViewType sets ViewType field to given value.

### HasViewType

`func (o *DatasetZipLocation) HasViewType() bool`

HasViewType returns a boolean if a field has been set.

### GetDataItems

`func (o *DatasetZipLocation) GetDataItems() []S3StorageInner`

GetDataItems returns the DataItems field if non-nil, zero value otherwise.

### GetDataItemsOk

`func (o *DatasetZipLocation) GetDataItemsOk() (*[]S3StorageInner, bool)`

GetDataItemsOk returns a tuple with the DataItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataItems

`func (o *DatasetZipLocation) SetDataItems(v []S3StorageInner)`

SetDataItems sets DataItems field to given value.

### HasDataItems

`func (o *DatasetZipLocation) HasDataItems() bool`

HasDataItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


