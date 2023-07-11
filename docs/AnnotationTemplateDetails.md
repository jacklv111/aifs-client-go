# AnnotationTemplateDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | the id of the annotation template | [optional] 
**Name** | Pointer to **string** | name of the annotation template | [optional] 
**CreateAt** | Pointer to **int64** | Unix timestamp in ms | [optional] 
**UpdateAt** | Pointer to **int64** | Unix timestamp in ms | [optional] 
**Type** | Pointer to **string** | the type of the annotation template | [optional] 
**Labels** | Pointer to [**[]Label**](Label.md) |  | [optional] 
**WordList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAnnotationTemplateDetails

`func NewAnnotationTemplateDetails() *AnnotationTemplateDetails`

NewAnnotationTemplateDetails instantiates a new AnnotationTemplateDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnnotationTemplateDetailsWithDefaults

`func NewAnnotationTemplateDetailsWithDefaults() *AnnotationTemplateDetails`

NewAnnotationTemplateDetailsWithDefaults instantiates a new AnnotationTemplateDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AnnotationTemplateDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AnnotationTemplateDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AnnotationTemplateDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AnnotationTemplateDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AnnotationTemplateDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AnnotationTemplateDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AnnotationTemplateDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AnnotationTemplateDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreateAt

`func (o *AnnotationTemplateDetails) GetCreateAt() int64`

GetCreateAt returns the CreateAt field if non-nil, zero value otherwise.

### GetCreateAtOk

`func (o *AnnotationTemplateDetails) GetCreateAtOk() (*int64, bool)`

GetCreateAtOk returns a tuple with the CreateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateAt

`func (o *AnnotationTemplateDetails) SetCreateAt(v int64)`

SetCreateAt sets CreateAt field to given value.

### HasCreateAt

`func (o *AnnotationTemplateDetails) HasCreateAt() bool`

HasCreateAt returns a boolean if a field has been set.

### GetUpdateAt

`func (o *AnnotationTemplateDetails) GetUpdateAt() int64`

GetUpdateAt returns the UpdateAt field if non-nil, zero value otherwise.

### GetUpdateAtOk

`func (o *AnnotationTemplateDetails) GetUpdateAtOk() (*int64, bool)`

GetUpdateAtOk returns a tuple with the UpdateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAt

`func (o *AnnotationTemplateDetails) SetUpdateAt(v int64)`

SetUpdateAt sets UpdateAt field to given value.

### HasUpdateAt

`func (o *AnnotationTemplateDetails) HasUpdateAt() bool`

HasUpdateAt returns a boolean if a field has been set.

### GetType

`func (o *AnnotationTemplateDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AnnotationTemplateDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AnnotationTemplateDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AnnotationTemplateDetails) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLabels

`func (o *AnnotationTemplateDetails) GetLabels() []Label`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AnnotationTemplateDetails) GetLabelsOk() (*[]Label, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AnnotationTemplateDetails) SetLabels(v []Label)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AnnotationTemplateDetails) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetWordList

`func (o *AnnotationTemplateDetails) GetWordList() []string`

GetWordList returns the WordList field if non-nil, zero value otherwise.

### GetWordListOk

`func (o *AnnotationTemplateDetails) GetWordListOk() (*[]string, bool)`

GetWordListOk returns a tuple with the WordList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWordList

`func (o *AnnotationTemplateDetails) SetWordList(v []string)`

SetWordList sets WordList field to given value.

### HasWordList

`func (o *AnnotationTemplateDetails) HasWordList() bool`

HasWordList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


