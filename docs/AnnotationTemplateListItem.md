# AnnotationTemplateListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | the id of the annotation template | [optional] 
**Name** | Pointer to **string** | name of the annotation template | [optional] 
**CreateAt** | Pointer to **int64** | Unix timestamp in ms | [optional] 
**Type** | Pointer to **string** | the type of the annotation template | [optional] 
**LabelCount** | Pointer to **int32** | the number of labels annotation template has | [optional] 

## Methods

### NewAnnotationTemplateListItem

`func NewAnnotationTemplateListItem() *AnnotationTemplateListItem`

NewAnnotationTemplateListItem instantiates a new AnnotationTemplateListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnnotationTemplateListItemWithDefaults

`func NewAnnotationTemplateListItemWithDefaults() *AnnotationTemplateListItem`

NewAnnotationTemplateListItemWithDefaults instantiates a new AnnotationTemplateListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AnnotationTemplateListItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AnnotationTemplateListItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AnnotationTemplateListItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AnnotationTemplateListItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AnnotationTemplateListItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AnnotationTemplateListItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AnnotationTemplateListItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AnnotationTemplateListItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreateAt

`func (o *AnnotationTemplateListItem) GetCreateAt() int64`

GetCreateAt returns the CreateAt field if non-nil, zero value otherwise.

### GetCreateAtOk

`func (o *AnnotationTemplateListItem) GetCreateAtOk() (*int64, bool)`

GetCreateAtOk returns a tuple with the CreateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateAt

`func (o *AnnotationTemplateListItem) SetCreateAt(v int64)`

SetCreateAt sets CreateAt field to given value.

### HasCreateAt

`func (o *AnnotationTemplateListItem) HasCreateAt() bool`

HasCreateAt returns a boolean if a field has been set.

### GetType

`func (o *AnnotationTemplateListItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AnnotationTemplateListItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AnnotationTemplateListItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AnnotationTemplateListItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLabelCount

`func (o *AnnotationTemplateListItem) GetLabelCount() int32`

GetLabelCount returns the LabelCount field if non-nil, zero value otherwise.

### GetLabelCountOk

`func (o *AnnotationTemplateListItem) GetLabelCountOk() (*int32, bool)`

GetLabelCountOk returns a tuple with the LabelCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelCount

`func (o *AnnotationTemplateListItem) SetLabelCount(v int32)`

SetLabelCount sets LabelCount field to given value.

### HasLabelCount

`func (o *AnnotationTemplateListItem) HasLabelCount() bool`

HasLabelCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


