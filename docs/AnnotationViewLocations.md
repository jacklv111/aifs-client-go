# AnnotationViewLocations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataViewId** | Pointer to **string** |  | [optional] 
**AnnotationTemplateId** | Pointer to **string** |  | [optional] 
**ViewType** | Pointer to [**DataViewType**](DataViewType.md) |  | [optional] 
**DataItems** | Pointer to [**[]S3StorageInner**](S3StorageInner.md) |  | [optional] 

## Methods

### NewAnnotationViewLocations

`func NewAnnotationViewLocations() *AnnotationViewLocations`

NewAnnotationViewLocations instantiates a new AnnotationViewLocations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnnotationViewLocationsWithDefaults

`func NewAnnotationViewLocationsWithDefaults() *AnnotationViewLocations`

NewAnnotationViewLocationsWithDefaults instantiates a new AnnotationViewLocations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataViewId

`func (o *AnnotationViewLocations) GetDataViewId() string`

GetDataViewId returns the DataViewId field if non-nil, zero value otherwise.

### GetDataViewIdOk

`func (o *AnnotationViewLocations) GetDataViewIdOk() (*string, bool)`

GetDataViewIdOk returns a tuple with the DataViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataViewId

`func (o *AnnotationViewLocations) SetDataViewId(v string)`

SetDataViewId sets DataViewId field to given value.

### HasDataViewId

`func (o *AnnotationViewLocations) HasDataViewId() bool`

HasDataViewId returns a boolean if a field has been set.

### GetAnnotationTemplateId

`func (o *AnnotationViewLocations) GetAnnotationTemplateId() string`

GetAnnotationTemplateId returns the AnnotationTemplateId field if non-nil, zero value otherwise.

### GetAnnotationTemplateIdOk

`func (o *AnnotationViewLocations) GetAnnotationTemplateIdOk() (*string, bool)`

GetAnnotationTemplateIdOk returns a tuple with the AnnotationTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationTemplateId

`func (o *AnnotationViewLocations) SetAnnotationTemplateId(v string)`

SetAnnotationTemplateId sets AnnotationTemplateId field to given value.

### HasAnnotationTemplateId

`func (o *AnnotationViewLocations) HasAnnotationTemplateId() bool`

HasAnnotationTemplateId returns a boolean if a field has been set.

### GetViewType

`func (o *AnnotationViewLocations) GetViewType() DataViewType`

GetViewType returns the ViewType field if non-nil, zero value otherwise.

### GetViewTypeOk

`func (o *AnnotationViewLocations) GetViewTypeOk() (*DataViewType, bool)`

GetViewTypeOk returns a tuple with the ViewType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewType

`func (o *AnnotationViewLocations) SetViewType(v DataViewType)`

SetViewType sets ViewType field to given value.

### HasViewType

`func (o *AnnotationViewLocations) HasViewType() bool`

HasViewType returns a boolean if a field has been set.

### GetDataItems

`func (o *AnnotationViewLocations) GetDataItems() []S3StorageInner`

GetDataItems returns the DataItems field if non-nil, zero value otherwise.

### GetDataItemsOk

`func (o *AnnotationViewLocations) GetDataItemsOk() (*[]S3StorageInner, bool)`

GetDataItemsOk returns a tuple with the DataItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataItems

`func (o *AnnotationViewLocations) SetDataItems(v []S3StorageInner)`

SetDataItems sets DataItems field to given value.

### HasDataItems

`func (o *AnnotationViewLocations) HasDataItems() bool`

HasDataItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


