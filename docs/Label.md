# Label

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | the id of the label | [optional] 
**Name** | **string** | the name of the label | 
**SuperCategoryName** | Pointer to **string** | the super category name | [optional] 
**Color** | **int32** |  | 
**KeyPointDef** | Pointer to **[]string** | the defination of the key points | [optional] 
**KeyPointSkeleton** | Pointer to **[][]int32** |  | [optional] 
**CoverImageUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewLabel

`func NewLabel(name string, color int32, ) *Label`

NewLabel instantiates a new Label object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelWithDefaults

`func NewLabelWithDefaults() *Label`

NewLabelWithDefaults instantiates a new Label object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Label) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Label) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Label) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Label) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Label) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Label) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Label) SetName(v string)`

SetName sets Name field to given value.


### GetSuperCategoryName

`func (o *Label) GetSuperCategoryName() string`

GetSuperCategoryName returns the SuperCategoryName field if non-nil, zero value otherwise.

### GetSuperCategoryNameOk

`func (o *Label) GetSuperCategoryNameOk() (*string, bool)`

GetSuperCategoryNameOk returns a tuple with the SuperCategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperCategoryName

`func (o *Label) SetSuperCategoryName(v string)`

SetSuperCategoryName sets SuperCategoryName field to given value.

### HasSuperCategoryName

`func (o *Label) HasSuperCategoryName() bool`

HasSuperCategoryName returns a boolean if a field has been set.

### GetColor

`func (o *Label) GetColor() int32`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *Label) GetColorOk() (*int32, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *Label) SetColor(v int32)`

SetColor sets Color field to given value.


### GetKeyPointDef

`func (o *Label) GetKeyPointDef() []string`

GetKeyPointDef returns the KeyPointDef field if non-nil, zero value otherwise.

### GetKeyPointDefOk

`func (o *Label) GetKeyPointDefOk() (*[]string, bool)`

GetKeyPointDefOk returns a tuple with the KeyPointDef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPointDef

`func (o *Label) SetKeyPointDef(v []string)`

SetKeyPointDef sets KeyPointDef field to given value.

### HasKeyPointDef

`func (o *Label) HasKeyPointDef() bool`

HasKeyPointDef returns a boolean if a field has been set.

### GetKeyPointSkeleton

`func (o *Label) GetKeyPointSkeleton() [][]int32`

GetKeyPointSkeleton returns the KeyPointSkeleton field if non-nil, zero value otherwise.

### GetKeyPointSkeletonOk

`func (o *Label) GetKeyPointSkeletonOk() (*[][]int32, bool)`

GetKeyPointSkeletonOk returns a tuple with the KeyPointSkeleton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPointSkeleton

`func (o *Label) SetKeyPointSkeleton(v [][]int32)`

SetKeyPointSkeleton sets KeyPointSkeleton field to given value.

### HasKeyPointSkeleton

`func (o *Label) HasKeyPointSkeleton() bool`

HasKeyPointSkeleton returns a boolean if a field has been set.

### GetCoverImageUrl

`func (o *Label) GetCoverImageUrl() string`

GetCoverImageUrl returns the CoverImageUrl field if non-nil, zero value otherwise.

### GetCoverImageUrlOk

`func (o *Label) GetCoverImageUrlOk() (*string, bool)`

GetCoverImageUrlOk returns a tuple with the CoverImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverImageUrl

`func (o *Label) SetCoverImageUrl(v string)`

SetCoverImageUrl sets CoverImageUrl field to given value.

### HasCoverImageUrl

`func (o *Label) HasCoverImageUrl() bool`

HasCoverImageUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


