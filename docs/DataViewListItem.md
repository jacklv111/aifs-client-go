# DataViewListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | the id of the data view | [optional] 
**Name** | Pointer to **string** | the name of the data view | [optional] 
**ViewType** | Pointer to [**DataViewType**](DataViewType.md) |  | [optional] 
**RawDataType** | Pointer to [**RawDataType**](RawDataType.md) |  | [optional] 
**AnnotationTemplateId** | Pointer to **string** | if data view is an annotation view, it has annotation template id | [optional] 
**AnnotationTemplateType** | Pointer to **string** | if data view is an annotation view, it has annotation template type | [optional] 
**CreateAt** | Pointer to **int64** | Unix timestamp in ms | [optional] 

## Methods

### NewDataViewListItem

`func NewDataViewListItem() *DataViewListItem`

NewDataViewListItem instantiates a new DataViewListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataViewListItemWithDefaults

`func NewDataViewListItemWithDefaults() *DataViewListItem`

NewDataViewListItemWithDefaults instantiates a new DataViewListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataViewListItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataViewListItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataViewListItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DataViewListItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DataViewListItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataViewListItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataViewListItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataViewListItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetViewType

`func (o *DataViewListItem) GetViewType() DataViewType`

GetViewType returns the ViewType field if non-nil, zero value otherwise.

### GetViewTypeOk

`func (o *DataViewListItem) GetViewTypeOk() (*DataViewType, bool)`

GetViewTypeOk returns a tuple with the ViewType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewType

`func (o *DataViewListItem) SetViewType(v DataViewType)`

SetViewType sets ViewType field to given value.

### HasViewType

`func (o *DataViewListItem) HasViewType() bool`

HasViewType returns a boolean if a field has been set.

### GetRawDataType

`func (o *DataViewListItem) GetRawDataType() RawDataType`

GetRawDataType returns the RawDataType field if non-nil, zero value otherwise.

### GetRawDataTypeOk

`func (o *DataViewListItem) GetRawDataTypeOk() (*RawDataType, bool)`

GetRawDataTypeOk returns a tuple with the RawDataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawDataType

`func (o *DataViewListItem) SetRawDataType(v RawDataType)`

SetRawDataType sets RawDataType field to given value.

### HasRawDataType

`func (o *DataViewListItem) HasRawDataType() bool`

HasRawDataType returns a boolean if a field has been set.

### GetAnnotationTemplateId

`func (o *DataViewListItem) GetAnnotationTemplateId() string`

GetAnnotationTemplateId returns the AnnotationTemplateId field if non-nil, zero value otherwise.

### GetAnnotationTemplateIdOk

`func (o *DataViewListItem) GetAnnotationTemplateIdOk() (*string, bool)`

GetAnnotationTemplateIdOk returns a tuple with the AnnotationTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationTemplateId

`func (o *DataViewListItem) SetAnnotationTemplateId(v string)`

SetAnnotationTemplateId sets AnnotationTemplateId field to given value.

### HasAnnotationTemplateId

`func (o *DataViewListItem) HasAnnotationTemplateId() bool`

HasAnnotationTemplateId returns a boolean if a field has been set.

### GetAnnotationTemplateType

`func (o *DataViewListItem) GetAnnotationTemplateType() string`

GetAnnotationTemplateType returns the AnnotationTemplateType field if non-nil, zero value otherwise.

### GetAnnotationTemplateTypeOk

`func (o *DataViewListItem) GetAnnotationTemplateTypeOk() (*string, bool)`

GetAnnotationTemplateTypeOk returns a tuple with the AnnotationTemplateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationTemplateType

`func (o *DataViewListItem) SetAnnotationTemplateType(v string)`

SetAnnotationTemplateType sets AnnotationTemplateType field to given value.

### HasAnnotationTemplateType

`func (o *DataViewListItem) HasAnnotationTemplateType() bool`

HasAnnotationTemplateType returns a boolean if a field has been set.

### GetCreateAt

`func (o *DataViewListItem) GetCreateAt() int64`

GetCreateAt returns the CreateAt field if non-nil, zero value otherwise.

### GetCreateAtOk

`func (o *DataViewListItem) GetCreateAtOk() (*int64, bool)`

GetCreateAtOk returns a tuple with the CreateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateAt

`func (o *DataViewListItem) SetCreateAt(v int64)`

SetCreateAt sets CreateAt field to given value.

### HasCreateAt

`func (o *DataViewListItem) HasCreateAt() bool`

HasCreateAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


