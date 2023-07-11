# AnnotationDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataItemId** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**TextData** | Pointer to **string** | text format annotation data | [optional] 

## Methods

### NewAnnotationDataInner

`func NewAnnotationDataInner() *AnnotationDataInner`

NewAnnotationDataInner instantiates a new AnnotationDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnnotationDataInnerWithDefaults

`func NewAnnotationDataInnerWithDefaults() *AnnotationDataInner`

NewAnnotationDataInnerWithDefaults instantiates a new AnnotationDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataItemId

`func (o *AnnotationDataInner) GetDataItemId() string`

GetDataItemId returns the DataItemId field if non-nil, zero value otherwise.

### GetDataItemIdOk

`func (o *AnnotationDataInner) GetDataItemIdOk() (*string, bool)`

GetDataItemIdOk returns a tuple with the DataItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataItemId

`func (o *AnnotationDataInner) SetDataItemId(v string)`

SetDataItemId sets DataItemId field to given value.

### HasDataItemId

`func (o *AnnotationDataInner) HasDataItemId() bool`

HasDataItemId returns a boolean if a field has been set.

### GetLabels

`func (o *AnnotationDataInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AnnotationDataInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AnnotationDataInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AnnotationDataInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTextData

`func (o *AnnotationDataInner) GetTextData() string`

GetTextData returns the TextData field if non-nil, zero value otherwise.

### GetTextDataOk

`func (o *AnnotationDataInner) GetTextDataOk() (*string, bool)`

GetTextDataOk returns a tuple with the TextData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextData

`func (o *AnnotationDataInner) SetTextData(v string)`

SetTextData sets TextData field to given value.

### HasTextData

`func (o *AnnotationDataInner) HasTextData() bool`

HasTextData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


