# CreateDataViewRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataViewName** | Pointer to **string** | the name of the data view | [optional] 
**Description** | Pointer to **string** | the description of the data view | [optional] 
**ViewType** | Pointer to [**DataViewType**](DataViewType.md) |  | [optional] 
**RawDataType** | Pointer to [**RawDataType**](RawDataType.md) |  | [optional] 
**ZipFormat** | Pointer to [**ZipFormat**](ZipFormat.md) |  | [optional] 
**RelatedDataViewId** | Pointer to **string** | If it is an annotation type data view, it must have a related raw-data data view | [optional] 
**AnnotationTemplateId** | Pointer to **string** | If it is an annotation type data view, it must have a related annotation template id. If it is a dataset-zip data view, it can have an annotation template id to indicate the annotation template of the annotation data. | [optional] 
**RawDataViewId** | Pointer to **string** | If it is a dataset-zip type data view, it can have a raw data view id to upload raw data to the data view | [optional] 
**AnnotationViewId** | Pointer to **string** | If it is a dataset-zip type data view, it can have a annotation view id to upload annotation data to the data view | [optional] 

## Methods

### NewCreateDataViewRequest

`func NewCreateDataViewRequest() *CreateDataViewRequest`

NewCreateDataViewRequest instantiates a new CreateDataViewRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDataViewRequestWithDefaults

`func NewCreateDataViewRequestWithDefaults() *CreateDataViewRequest`

NewCreateDataViewRequestWithDefaults instantiates a new CreateDataViewRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataViewName

`func (o *CreateDataViewRequest) GetDataViewName() string`

GetDataViewName returns the DataViewName field if non-nil, zero value otherwise.

### GetDataViewNameOk

`func (o *CreateDataViewRequest) GetDataViewNameOk() (*string, bool)`

GetDataViewNameOk returns a tuple with the DataViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataViewName

`func (o *CreateDataViewRequest) SetDataViewName(v string)`

SetDataViewName sets DataViewName field to given value.

### HasDataViewName

`func (o *CreateDataViewRequest) HasDataViewName() bool`

HasDataViewName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateDataViewRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateDataViewRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateDataViewRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateDataViewRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetViewType

`func (o *CreateDataViewRequest) GetViewType() DataViewType`

GetViewType returns the ViewType field if non-nil, zero value otherwise.

### GetViewTypeOk

`func (o *CreateDataViewRequest) GetViewTypeOk() (*DataViewType, bool)`

GetViewTypeOk returns a tuple with the ViewType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewType

`func (o *CreateDataViewRequest) SetViewType(v DataViewType)`

SetViewType sets ViewType field to given value.

### HasViewType

`func (o *CreateDataViewRequest) HasViewType() bool`

HasViewType returns a boolean if a field has been set.

### GetRawDataType

`func (o *CreateDataViewRequest) GetRawDataType() RawDataType`

GetRawDataType returns the RawDataType field if non-nil, zero value otherwise.

### GetRawDataTypeOk

`func (o *CreateDataViewRequest) GetRawDataTypeOk() (*RawDataType, bool)`

GetRawDataTypeOk returns a tuple with the RawDataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawDataType

`func (o *CreateDataViewRequest) SetRawDataType(v RawDataType)`

SetRawDataType sets RawDataType field to given value.

### HasRawDataType

`func (o *CreateDataViewRequest) HasRawDataType() bool`

HasRawDataType returns a boolean if a field has been set.

### GetZipFormat

`func (o *CreateDataViewRequest) GetZipFormat() ZipFormat`

GetZipFormat returns the ZipFormat field if non-nil, zero value otherwise.

### GetZipFormatOk

`func (o *CreateDataViewRequest) GetZipFormatOk() (*ZipFormat, bool)`

GetZipFormatOk returns a tuple with the ZipFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipFormat

`func (o *CreateDataViewRequest) SetZipFormat(v ZipFormat)`

SetZipFormat sets ZipFormat field to given value.

### HasZipFormat

`func (o *CreateDataViewRequest) HasZipFormat() bool`

HasZipFormat returns a boolean if a field has been set.

### GetRelatedDataViewId

`func (o *CreateDataViewRequest) GetRelatedDataViewId() string`

GetRelatedDataViewId returns the RelatedDataViewId field if non-nil, zero value otherwise.

### GetRelatedDataViewIdOk

`func (o *CreateDataViewRequest) GetRelatedDataViewIdOk() (*string, bool)`

GetRelatedDataViewIdOk returns a tuple with the RelatedDataViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedDataViewId

`func (o *CreateDataViewRequest) SetRelatedDataViewId(v string)`

SetRelatedDataViewId sets RelatedDataViewId field to given value.

### HasRelatedDataViewId

`func (o *CreateDataViewRequest) HasRelatedDataViewId() bool`

HasRelatedDataViewId returns a boolean if a field has been set.

### GetAnnotationTemplateId

`func (o *CreateDataViewRequest) GetAnnotationTemplateId() string`

GetAnnotationTemplateId returns the AnnotationTemplateId field if non-nil, zero value otherwise.

### GetAnnotationTemplateIdOk

`func (o *CreateDataViewRequest) GetAnnotationTemplateIdOk() (*string, bool)`

GetAnnotationTemplateIdOk returns a tuple with the AnnotationTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationTemplateId

`func (o *CreateDataViewRequest) SetAnnotationTemplateId(v string)`

SetAnnotationTemplateId sets AnnotationTemplateId field to given value.

### HasAnnotationTemplateId

`func (o *CreateDataViewRequest) HasAnnotationTemplateId() bool`

HasAnnotationTemplateId returns a boolean if a field has been set.

### GetRawDataViewId

`func (o *CreateDataViewRequest) GetRawDataViewId() string`

GetRawDataViewId returns the RawDataViewId field if non-nil, zero value otherwise.

### GetRawDataViewIdOk

`func (o *CreateDataViewRequest) GetRawDataViewIdOk() (*string, bool)`

GetRawDataViewIdOk returns a tuple with the RawDataViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawDataViewId

`func (o *CreateDataViewRequest) SetRawDataViewId(v string)`

SetRawDataViewId sets RawDataViewId field to given value.

### HasRawDataViewId

`func (o *CreateDataViewRequest) HasRawDataViewId() bool`

HasRawDataViewId returns a boolean if a field has been set.

### GetAnnotationViewId

`func (o *CreateDataViewRequest) GetAnnotationViewId() string`

GetAnnotationViewId returns the AnnotationViewId field if non-nil, zero value otherwise.

### GetAnnotationViewIdOk

`func (o *CreateDataViewRequest) GetAnnotationViewIdOk() (*string, bool)`

GetAnnotationViewIdOk returns a tuple with the AnnotationViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationViewId

`func (o *CreateDataViewRequest) SetAnnotationViewId(v string)`

SetAnnotationViewId sets AnnotationViewId field to given value.

### HasAnnotationViewId

`func (o *CreateDataViewRequest) HasAnnotationViewId() bool`

HasAnnotationViewId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


