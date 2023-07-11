# UpdateDatasetZipRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Progress** | Pointer to **float32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TrainRawDataViewId** | Pointer to **string** |  | [optional] 
**TrainAnnotationViewId** | Pointer to **string** |  | [optional] 
**ValRawDataViewId** | Pointer to **string** |  | [optional] 
**ValAnnotationViewId** | Pointer to **string** |  | [optional] 
**AnnotationTemplateId** | Pointer to **string** |  | [optional] 
**RawDataViewId** | Pointer to **string** |  | [optional] 
**AnnotationViewId** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateDatasetZipRequest

`func NewUpdateDatasetZipRequest() *UpdateDatasetZipRequest`

NewUpdateDatasetZipRequest instantiates a new UpdateDatasetZipRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDatasetZipRequestWithDefaults

`func NewUpdateDatasetZipRequestWithDefaults() *UpdateDatasetZipRequest`

NewUpdateDatasetZipRequestWithDefaults instantiates a new UpdateDatasetZipRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProgress

`func (o *UpdateDatasetZipRequest) GetProgress() float32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *UpdateDatasetZipRequest) GetProgressOk() (*float32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *UpdateDatasetZipRequest) SetProgress(v float32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *UpdateDatasetZipRequest) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateDatasetZipRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateDatasetZipRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateDatasetZipRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateDatasetZipRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTrainRawDataViewId

`func (o *UpdateDatasetZipRequest) GetTrainRawDataViewId() string`

GetTrainRawDataViewId returns the TrainRawDataViewId field if non-nil, zero value otherwise.

### GetTrainRawDataViewIdOk

`func (o *UpdateDatasetZipRequest) GetTrainRawDataViewIdOk() (*string, bool)`

GetTrainRawDataViewIdOk returns a tuple with the TrainRawDataViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainRawDataViewId

`func (o *UpdateDatasetZipRequest) SetTrainRawDataViewId(v string)`

SetTrainRawDataViewId sets TrainRawDataViewId field to given value.

### HasTrainRawDataViewId

`func (o *UpdateDatasetZipRequest) HasTrainRawDataViewId() bool`

HasTrainRawDataViewId returns a boolean if a field has been set.

### GetTrainAnnotationViewId

`func (o *UpdateDatasetZipRequest) GetTrainAnnotationViewId() string`

GetTrainAnnotationViewId returns the TrainAnnotationViewId field if non-nil, zero value otherwise.

### GetTrainAnnotationViewIdOk

`func (o *UpdateDatasetZipRequest) GetTrainAnnotationViewIdOk() (*string, bool)`

GetTrainAnnotationViewIdOk returns a tuple with the TrainAnnotationViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainAnnotationViewId

`func (o *UpdateDatasetZipRequest) SetTrainAnnotationViewId(v string)`

SetTrainAnnotationViewId sets TrainAnnotationViewId field to given value.

### HasTrainAnnotationViewId

`func (o *UpdateDatasetZipRequest) HasTrainAnnotationViewId() bool`

HasTrainAnnotationViewId returns a boolean if a field has been set.

### GetValRawDataViewId

`func (o *UpdateDatasetZipRequest) GetValRawDataViewId() string`

GetValRawDataViewId returns the ValRawDataViewId field if non-nil, zero value otherwise.

### GetValRawDataViewIdOk

`func (o *UpdateDatasetZipRequest) GetValRawDataViewIdOk() (*string, bool)`

GetValRawDataViewIdOk returns a tuple with the ValRawDataViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValRawDataViewId

`func (o *UpdateDatasetZipRequest) SetValRawDataViewId(v string)`

SetValRawDataViewId sets ValRawDataViewId field to given value.

### HasValRawDataViewId

`func (o *UpdateDatasetZipRequest) HasValRawDataViewId() bool`

HasValRawDataViewId returns a boolean if a field has been set.

### GetValAnnotationViewId

`func (o *UpdateDatasetZipRequest) GetValAnnotationViewId() string`

GetValAnnotationViewId returns the ValAnnotationViewId field if non-nil, zero value otherwise.

### GetValAnnotationViewIdOk

`func (o *UpdateDatasetZipRequest) GetValAnnotationViewIdOk() (*string, bool)`

GetValAnnotationViewIdOk returns a tuple with the ValAnnotationViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValAnnotationViewId

`func (o *UpdateDatasetZipRequest) SetValAnnotationViewId(v string)`

SetValAnnotationViewId sets ValAnnotationViewId field to given value.

### HasValAnnotationViewId

`func (o *UpdateDatasetZipRequest) HasValAnnotationViewId() bool`

HasValAnnotationViewId returns a boolean if a field has been set.

### GetAnnotationTemplateId

`func (o *UpdateDatasetZipRequest) GetAnnotationTemplateId() string`

GetAnnotationTemplateId returns the AnnotationTemplateId field if non-nil, zero value otherwise.

### GetAnnotationTemplateIdOk

`func (o *UpdateDatasetZipRequest) GetAnnotationTemplateIdOk() (*string, bool)`

GetAnnotationTemplateIdOk returns a tuple with the AnnotationTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationTemplateId

`func (o *UpdateDatasetZipRequest) SetAnnotationTemplateId(v string)`

SetAnnotationTemplateId sets AnnotationTemplateId field to given value.

### HasAnnotationTemplateId

`func (o *UpdateDatasetZipRequest) HasAnnotationTemplateId() bool`

HasAnnotationTemplateId returns a boolean if a field has been set.

### GetRawDataViewId

`func (o *UpdateDatasetZipRequest) GetRawDataViewId() string`

GetRawDataViewId returns the RawDataViewId field if non-nil, zero value otherwise.

### GetRawDataViewIdOk

`func (o *UpdateDatasetZipRequest) GetRawDataViewIdOk() (*string, bool)`

GetRawDataViewIdOk returns a tuple with the RawDataViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawDataViewId

`func (o *UpdateDatasetZipRequest) SetRawDataViewId(v string)`

SetRawDataViewId sets RawDataViewId field to given value.

### HasRawDataViewId

`func (o *UpdateDatasetZipRequest) HasRawDataViewId() bool`

HasRawDataViewId returns a boolean if a field has been set.

### GetAnnotationViewId

`func (o *UpdateDatasetZipRequest) GetAnnotationViewId() string`

GetAnnotationViewId returns the AnnotationViewId field if non-nil, zero value otherwise.

### GetAnnotationViewIdOk

`func (o *UpdateDatasetZipRequest) GetAnnotationViewIdOk() (*string, bool)`

GetAnnotationViewIdOk returns a tuple with the AnnotationViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationViewId

`func (o *UpdateDatasetZipRequest) SetAnnotationViewId(v string)`

SetAnnotationViewId sets AnnotationViewId field to given value.

### HasAnnotationViewId

`func (o *UpdateDatasetZipRequest) HasAnnotationViewId() bool`

HasAnnotationViewId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


