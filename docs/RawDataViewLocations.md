# RawDataViewLocations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataViewId** | Pointer to **string** |  | [optional] 
**ViewType** | Pointer to [**DataViewType**](DataViewType.md) |  | [optional] 
**RawDataType** | Pointer to [**RawDataType**](RawDataType.md) |  | [optional] 
**DataItems** | Pointer to [**[]S3StorageInner**](S3StorageInner.md) |  | [optional] 

## Methods

### NewRawDataViewLocations

`func NewRawDataViewLocations() *RawDataViewLocations`

NewRawDataViewLocations instantiates a new RawDataViewLocations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRawDataViewLocationsWithDefaults

`func NewRawDataViewLocationsWithDefaults() *RawDataViewLocations`

NewRawDataViewLocationsWithDefaults instantiates a new RawDataViewLocations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataViewId

`func (o *RawDataViewLocations) GetDataViewId() string`

GetDataViewId returns the DataViewId field if non-nil, zero value otherwise.

### GetDataViewIdOk

`func (o *RawDataViewLocations) GetDataViewIdOk() (*string, bool)`

GetDataViewIdOk returns a tuple with the DataViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataViewId

`func (o *RawDataViewLocations) SetDataViewId(v string)`

SetDataViewId sets DataViewId field to given value.

### HasDataViewId

`func (o *RawDataViewLocations) HasDataViewId() bool`

HasDataViewId returns a boolean if a field has been set.

### GetViewType

`func (o *RawDataViewLocations) GetViewType() DataViewType`

GetViewType returns the ViewType field if non-nil, zero value otherwise.

### GetViewTypeOk

`func (o *RawDataViewLocations) GetViewTypeOk() (*DataViewType, bool)`

GetViewTypeOk returns a tuple with the ViewType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewType

`func (o *RawDataViewLocations) SetViewType(v DataViewType)`

SetViewType sets ViewType field to given value.

### HasViewType

`func (o *RawDataViewLocations) HasViewType() bool`

HasViewType returns a boolean if a field has been set.

### GetRawDataType

`func (o *RawDataViewLocations) GetRawDataType() RawDataType`

GetRawDataType returns the RawDataType field if non-nil, zero value otherwise.

### GetRawDataTypeOk

`func (o *RawDataViewLocations) GetRawDataTypeOk() (*RawDataType, bool)`

GetRawDataTypeOk returns a tuple with the RawDataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawDataType

`func (o *RawDataViewLocations) SetRawDataType(v RawDataType)`

SetRawDataType sets RawDataType field to given value.

### HasRawDataType

`func (o *RawDataViewLocations) HasRawDataType() bool`

HasRawDataType returns a boolean if a field has been set.

### GetDataItems

`func (o *RawDataViewLocations) GetDataItems() []S3StorageInner`

GetDataItems returns the DataItems field if non-nil, zero value otherwise.

### GetDataItemsOk

`func (o *RawDataViewLocations) GetDataItemsOk() (*[]S3StorageInner, bool)`

GetDataItemsOk returns a tuple with the DataItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataItems

`func (o *RawDataViewLocations) SetDataItems(v []S3StorageInner)`

SetDataItems sets DataItems field to given value.

### HasDataItems

`func (o *RawDataViewLocations) HasDataItems() bool`

HasDataItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


