# UpdateAnnotationTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | the id of the annotation template | [optional] 
**Name** | **string** | the name of the annotation template | 
**Type** | **string** | the type of the annotation template | 
**Description** | Pointer to **string** | the description of the annotation template | [optional] 
**Labels** | Pointer to [**[]Label**](Label.md) |  | [optional] 
**WordList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUpdateAnnotationTemplateRequest

`func NewUpdateAnnotationTemplateRequest(name string, type_ string, ) *UpdateAnnotationTemplateRequest`

NewUpdateAnnotationTemplateRequest instantiates a new UpdateAnnotationTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAnnotationTemplateRequestWithDefaults

`func NewUpdateAnnotationTemplateRequestWithDefaults() *UpdateAnnotationTemplateRequest`

NewUpdateAnnotationTemplateRequestWithDefaults instantiates a new UpdateAnnotationTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateAnnotationTemplateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateAnnotationTemplateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateAnnotationTemplateRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateAnnotationTemplateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateAnnotationTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAnnotationTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAnnotationTemplateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *UpdateAnnotationTemplateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateAnnotationTemplateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateAnnotationTemplateRequest) SetType(v string)`

SetType sets Type field to given value.


### GetDescription

`func (o *UpdateAnnotationTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateAnnotationTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateAnnotationTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateAnnotationTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *UpdateAnnotationTemplateRequest) GetLabels() []Label`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateAnnotationTemplateRequest) GetLabelsOk() (*[]Label, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateAnnotationTemplateRequest) SetLabels(v []Label)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateAnnotationTemplateRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetWordList

`func (o *UpdateAnnotationTemplateRequest) GetWordList() []string`

GetWordList returns the WordList field if non-nil, zero value otherwise.

### GetWordListOk

`func (o *UpdateAnnotationTemplateRequest) GetWordListOk() (*[]string, bool)`

GetWordListOk returns a tuple with the WordList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWordList

`func (o *UpdateAnnotationTemplateRequest) SetWordList(v []string)`

SetWordList sets WordList field to given value.

### HasWordList

`func (o *UpdateAnnotationTemplateRequest) HasWordList() bool`

HasWordList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


