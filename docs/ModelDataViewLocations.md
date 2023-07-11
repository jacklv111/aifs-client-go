# ModelDataViewLocations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataViewId** | Pointer to **string** |  | [optional] 
**ViewType** | Pointer to [**DataViewType**](DataViewType.md) |  | [optional] 
**DataItems** | Pointer to [**[]S3StorageInner**](S3StorageInner.md) |  | [optional] 

## Methods

### NewModelDataViewLocations

`func NewModelDataViewLocations() *ModelDataViewLocations`

NewModelDataViewLocations instantiates a new ModelDataViewLocations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelDataViewLocationsWithDefaults

`func NewModelDataViewLocationsWithDefaults() *ModelDataViewLocations`

NewModelDataViewLocationsWithDefaults instantiates a new ModelDataViewLocations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataViewId

`func (o *ModelDataViewLocations) GetDataViewId() string`

GetDataViewId returns the DataViewId field if non-nil, zero value otherwise.

### GetDataViewIdOk

`func (o *ModelDataViewLocations) GetDataViewIdOk() (*string, bool)`

GetDataViewIdOk returns a tuple with the DataViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataViewId

`func (o *ModelDataViewLocations) SetDataViewId(v string)`

SetDataViewId sets DataViewId field to given value.

### HasDataViewId

`func (o *ModelDataViewLocations) HasDataViewId() bool`

HasDataViewId returns a boolean if a field has been set.

### GetViewType

`func (o *ModelDataViewLocations) GetViewType() DataViewType`

GetViewType returns the ViewType field if non-nil, zero value otherwise.

### GetViewTypeOk

`func (o *ModelDataViewLocations) GetViewTypeOk() (*DataViewType, bool)`

GetViewTypeOk returns a tuple with the ViewType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewType

`func (o *ModelDataViewLocations) SetViewType(v DataViewType)`

SetViewType sets ViewType field to given value.

### HasViewType

`func (o *ModelDataViewLocations) HasViewType() bool`

HasViewType returns a boolean if a field has been set.

### GetDataItems

`func (o *ModelDataViewLocations) GetDataItems() []S3StorageInner`

GetDataItems returns the DataItems field if non-nil, zero value otherwise.

### GetDataItemsOk

`func (o *ModelDataViewLocations) GetDataItemsOk() (*[]S3StorageInner, bool)`

GetDataItemsOk returns a tuple with the DataItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataItems

`func (o *ModelDataViewLocations) SetDataItems(v []S3StorageInner)`

SetDataItems sets DataItems field to given value.

### HasDataItems

`func (o *ModelDataViewLocations) HasDataItems() bool`

HasDataItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


