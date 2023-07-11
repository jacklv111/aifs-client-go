# DataViewDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | the id of the data view | [optional] 
**Name** | Pointer to **string** | the name of the data view | [optional] 
**ViewType** | Pointer to [**DataViewType**](DataViewType.md) |  | [optional] 
**Description** | Pointer to **string** | the description of the data view | [optional] 
**RawDataType** | Pointer to [**RawDataType**](RawDataType.md) |  | [optional] 
**AnnotationTemplateType** | Pointer to **string** | if data view is an annotation view, it has annotation template type | [optional] 
**CreateAt** | Pointer to **int64** | Unix timestamp in ms | [optional] 
**Progress** | Pointer to **float32** |  | [optional] 
**CommitId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ZipFormat** | Pointer to [**ZipFormat**](ZipFormat.md) |  | [optional] 
**TrainRawDataViewId** | Pointer to **string** |  | [optional] 
**TrainAnnotationViewId** | Pointer to **string** |  | [optional] 
**ValRawDataViewId** | Pointer to **string** |  | [optional] 
**ValAnnotationViewId** | Pointer to **string** |  | [optional] 
**AnnotationTemplateId** | Pointer to **string** |  | [optional] 
**RawDataViewId** | Pointer to **string** |  | [optional] 
**AnnotationViewId** | Pointer to **string** |  | [optional] 

## Methods

### NewDataViewDetails

`func NewDataViewDetails() *DataViewDetails`

NewDataViewDetails instantiates a new DataViewDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataViewDetailsWithDefaults

`func NewDataViewDetailsWithDefaults() *DataViewDetails`

NewDataViewDetailsWithDefaults instantiates a new DataViewDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataViewDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataViewDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataViewDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DataViewDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DataViewDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataViewDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataViewDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataViewDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetViewType

`func (o *DataViewDetails) GetViewType() DataViewType`

GetViewType returns the ViewType field if non-nil, zero value otherwise.

### GetViewTypeOk

`func (o *DataViewDetails) GetViewTypeOk() (*DataViewType, bool)`

GetViewTypeOk returns a tuple with the ViewType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewType

`func (o *DataViewDetails) SetViewType(v DataViewType)`

SetViewType sets ViewType field to given value.

### HasViewType

`func (o *DataViewDetails) HasViewType() bool`

HasViewType returns a boolean if a field has been set.

### GetDescription

`func (o *DataViewDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DataViewDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DataViewDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DataViewDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRawDataType

`func (o *DataViewDetails) GetRawDataType() RawDataType`

GetRawDataType returns the RawDataType field if non-nil, zero value otherwise.

### GetRawDataTypeOk

`func (o *DataViewDetails) GetRawDataTypeOk() (*RawDataType, bool)`

GetRawDataTypeOk returns a tuple with the RawDataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawDataType

`func (o *DataViewDetails) SetRawDataType(v RawDataType)`

SetRawDataType sets RawDataType field to given value.

### HasRawDataType

`func (o *DataViewDetails) HasRawDataType() bool`

HasRawDataType returns a boolean if a field has been set.

### GetAnnotationTemplateType

`func (o *DataViewDetails) GetAnnotationTemplateType() string`

GetAnnotationTemplateType returns the AnnotationTemplateType field if non-nil, zero value otherwise.

### GetAnnotationTemplateTypeOk

`func (o *DataViewDetails) GetAnnotationTemplateTypeOk() (*string, bool)`

GetAnnotationTemplateTypeOk returns a tuple with the AnnotationTemplateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationTemplateType

`func (o *DataViewDetails) SetAnnotationTemplateType(v string)`

SetAnnotationTemplateType sets AnnotationTemplateType field to given value.

### HasAnnotationTemplateType

`func (o *DataViewDetails) HasAnnotationTemplateType() bool`

HasAnnotationTemplateType returns a boolean if a field has been set.

### GetCreateAt

`func (o *DataViewDetails) GetCreateAt() int64`

GetCreateAt returns the CreateAt field if non-nil, zero value otherwise.

### GetCreateAtOk

`func (o *DataViewDetails) GetCreateAtOk() (*int64, bool)`

GetCreateAtOk returns a tuple with the CreateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateAt

`func (o *DataViewDetails) SetCreateAt(v int64)`

SetCreateAt sets CreateAt field to given value.

### HasCreateAt

`func (o *DataViewDetails) HasCreateAt() bool`

HasCreateAt returns a boolean if a field has been set.

### GetProgress

`func (o *DataViewDetails) GetProgress() float32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *DataViewDetails) GetProgressOk() (*float32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *DataViewDetails) SetProgress(v float32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *DataViewDetails) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetCommitId

`func (o *DataViewDetails) GetCommitId() string`

GetCommitId returns the CommitId field if non-nil, zero value otherwise.

### GetCommitIdOk

`func (o *DataViewDetails) GetCommitIdOk() (*string, bool)`

GetCommitIdOk returns a tuple with the CommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitId

`func (o *DataViewDetails) SetCommitId(v string)`

SetCommitId sets CommitId field to given value.

### HasCommitId

`func (o *DataViewDetails) HasCommitId() bool`

HasCommitId returns a boolean if a field has been set.

### GetStatus

`func (o *DataViewDetails) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DataViewDetails) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DataViewDetails) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DataViewDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetZipFormat

`func (o *DataViewDetails) GetZipFormat() ZipFormat`

GetZipFormat returns the ZipFormat field if non-nil, zero value otherwise.

### GetZipFormatOk

`func (o *DataViewDetails) GetZipFormatOk() (*ZipFormat, bool)`

GetZipFormatOk returns a tuple with the ZipFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipFormat

`func (o *DataViewDetails) SetZipFormat(v ZipFormat)`

SetZipFormat sets ZipFormat field to given value.

### HasZipFormat

`func (o *DataViewDetails) HasZipFormat() bool`

HasZipFormat returns a boolean if a field has been set.

### GetTrainRawDataViewId

`func (o *DataViewDetails) GetTrainRawDataViewId() string`

GetTrainRawDataViewId returns the TrainRawDataViewId field if non-nil, zero value otherwise.

### GetTrainRawDataViewIdOk

`func (o *DataViewDetails) GetTrainRawDataViewIdOk() (*string, bool)`

GetTrainRawDataViewIdOk returns a tuple with the TrainRawDataViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainRawDataViewId

`func (o *DataViewDetails) SetTrainRawDataViewId(v string)`

SetTrainRawDataViewId sets TrainRawDataViewId field to given value.

### HasTrainRawDataViewId

`func (o *DataViewDetails) HasTrainRawDataViewId() bool`

HasTrainRawDataViewId returns a boolean if a field has been set.

### GetTrainAnnotationViewId

`func (o *DataViewDetails) GetTrainAnnotationViewId() string`

GetTrainAnnotationViewId returns the TrainAnnotationViewId field if non-nil, zero value otherwise.

### GetTrainAnnotationViewIdOk

`func (o *DataViewDetails) GetTrainAnnotationViewIdOk() (*string, bool)`

GetTrainAnnotationViewIdOk returns a tuple with the TrainAnnotationViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainAnnotationViewId

`func (o *DataViewDetails) SetTrainAnnotationViewId(v string)`

SetTrainAnnotationViewId sets TrainAnnotationViewId field to given value.

### HasTrainAnnotationViewId

`func (o *DataViewDetails) HasTrainAnnotationViewId() bool`

HasTrainAnnotationViewId returns a boolean if a field has been set.

### GetValRawDataViewId

`func (o *DataViewDetails) GetValRawDataViewId() string`

GetValRawDataViewId returns the ValRawDataViewId field if non-nil, zero value otherwise.

### GetValRawDataViewIdOk

`func (o *DataViewDetails) GetValRawDataViewIdOk() (*string, bool)`

GetValRawDataViewIdOk returns a tuple with the ValRawDataViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValRawDataViewId

`func (o *DataViewDetails) SetValRawDataViewId(v string)`

SetValRawDataViewId sets ValRawDataViewId field to given value.

### HasValRawDataViewId

`func (o *DataViewDetails) HasValRawDataViewId() bool`

HasValRawDataViewId returns a boolean if a field has been set.

### GetValAnnotationViewId

`func (o *DataViewDetails) GetValAnnotationViewId() string`

GetValAnnotationViewId returns the ValAnnotationViewId field if non-nil, zero value otherwise.

### GetValAnnotationViewIdOk

`func (o *DataViewDetails) GetValAnnotationViewIdOk() (*string, bool)`

GetValAnnotationViewIdOk returns a tuple with the ValAnnotationViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValAnnotationViewId

`func (o *DataViewDetails) SetValAnnotationViewId(v string)`

SetValAnnotationViewId sets ValAnnotationViewId field to given value.

### HasValAnnotationViewId

`func (o *DataViewDetails) HasValAnnotationViewId() bool`

HasValAnnotationViewId returns a boolean if a field has been set.

### GetAnnotationTemplateId

`func (o *DataViewDetails) GetAnnotationTemplateId() string`

GetAnnotationTemplateId returns the AnnotationTemplateId field if non-nil, zero value otherwise.

### GetAnnotationTemplateIdOk

`func (o *DataViewDetails) GetAnnotationTemplateIdOk() (*string, bool)`

GetAnnotationTemplateIdOk returns a tuple with the AnnotationTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationTemplateId

`func (o *DataViewDetails) SetAnnotationTemplateId(v string)`

SetAnnotationTemplateId sets AnnotationTemplateId field to given value.

### HasAnnotationTemplateId

`func (o *DataViewDetails) HasAnnotationTemplateId() bool`

HasAnnotationTemplateId returns a boolean if a field has been set.

### GetRawDataViewId

`func (o *DataViewDetails) GetRawDataViewId() string`

GetRawDataViewId returns the RawDataViewId field if non-nil, zero value otherwise.

### GetRawDataViewIdOk

`func (o *DataViewDetails) GetRawDataViewIdOk() (*string, bool)`

GetRawDataViewIdOk returns a tuple with the RawDataViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawDataViewId

`func (o *DataViewDetails) SetRawDataViewId(v string)`

SetRawDataViewId sets RawDataViewId field to given value.

### HasRawDataViewId

`func (o *DataViewDetails) HasRawDataViewId() bool`

HasRawDataViewId returns a boolean if a field has been set.

### GetAnnotationViewId

`func (o *DataViewDetails) GetAnnotationViewId() string`

GetAnnotationViewId returns the AnnotationViewId field if non-nil, zero value otherwise.

### GetAnnotationViewIdOk

`func (o *DataViewDetails) GetAnnotationViewIdOk() (*string, bool)`

GetAnnotationViewIdOk returns a tuple with the AnnotationViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationViewId

`func (o *DataViewDetails) SetAnnotationViewId(v string)`

SetAnnotationViewId sets AnnotationViewId field to given value.

### HasAnnotationViewId

`func (o *DataViewDetails) HasAnnotationViewId() bool`

HasAnnotationViewId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


