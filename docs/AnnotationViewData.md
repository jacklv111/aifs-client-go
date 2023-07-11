# AnnotationViewData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataViewId** | Pointer to **string** |  | [optional] 
**AnnotationTemplateId** | Pointer to **string** |  | [optional] 
**ViewType** | Pointer to [**DataViewType**](DataViewType.md) |  | [optional] 
**DataItems** | Pointer to [**[]AnnotationDataInner**](AnnotationDataInner.md) |  | [optional] 

## Methods

### NewAnnotationViewData

`func NewAnnotationViewData() *AnnotationViewData`

NewAnnotationViewData instantiates a new AnnotationViewData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnnotationViewDataWithDefaults

`func NewAnnotationViewDataWithDefaults() *AnnotationViewData`

NewAnnotationViewDataWithDefaults instantiates a new AnnotationViewData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataViewId

`func (o *AnnotationViewData) GetDataViewId() string`

GetDataViewId returns the DataViewId field if non-nil, zero value otherwise.

### GetDataViewIdOk

`func (o *AnnotationViewData) GetDataViewIdOk() (*string, bool)`

GetDataViewIdOk returns a tuple with the DataViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataViewId

`func (o *AnnotationViewData) SetDataViewId(v string)`

SetDataViewId sets DataViewId field to given value.

### HasDataViewId

`func (o *AnnotationViewData) HasDataViewId() bool`

HasDataViewId returns a boolean if a field has been set.

### GetAnnotationTemplateId

`func (o *AnnotationViewData) GetAnnotationTemplateId() string`

GetAnnotationTemplateId returns the AnnotationTemplateId field if non-nil, zero value otherwise.

### GetAnnotationTemplateIdOk

`func (o *AnnotationViewData) GetAnnotationTemplateIdOk() (*string, bool)`

GetAnnotationTemplateIdOk returns a tuple with the AnnotationTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationTemplateId

`func (o *AnnotationViewData) SetAnnotationTemplateId(v string)`

SetAnnotationTemplateId sets AnnotationTemplateId field to given value.

### HasAnnotationTemplateId

`func (o *AnnotationViewData) HasAnnotationTemplateId() bool`

HasAnnotationTemplateId returns a boolean if a field has been set.

### GetViewType

`func (o *AnnotationViewData) GetViewType() DataViewType`

GetViewType returns the ViewType field if non-nil, zero value otherwise.

### GetViewTypeOk

`func (o *AnnotationViewData) GetViewTypeOk() (*DataViewType, bool)`

GetViewTypeOk returns a tuple with the ViewType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewType

`func (o *AnnotationViewData) SetViewType(v DataViewType)`

SetViewType sets ViewType field to given value.

### HasViewType

`func (o *AnnotationViewData) HasViewType() bool`

HasViewType returns a boolean if a field has been set.

### GetDataItems

`func (o *AnnotationViewData) GetDataItems() []AnnotationDataInner`

GetDataItems returns the DataItems field if non-nil, zero value otherwise.

### GetDataItemsOk

`func (o *AnnotationViewData) GetDataItemsOk() (*[]AnnotationDataInner, bool)`

GetDataItemsOk returns a tuple with the DataItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataItems

`func (o *AnnotationViewData) SetDataItems(v []AnnotationDataInner)`

SetDataItems sets DataItems field to given value.

### HasDataItems

`func (o *AnnotationViewData) HasDataItems() bool`

HasDataItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


