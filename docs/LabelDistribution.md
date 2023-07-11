# LabelDistribution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LabelId** | Pointer to **string** | the label id | [optional] 
**Count** | Pointer to **int32** | the number of the label in the data view. | [optional] 
**Ratio** | Pointer to **float32** | the ratio of the label in the data view. | [optional] 

## Methods

### NewLabelDistribution

`func NewLabelDistribution() *LabelDistribution`

NewLabelDistribution instantiates a new LabelDistribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelDistributionWithDefaults

`func NewLabelDistributionWithDefaults() *LabelDistribution`

NewLabelDistributionWithDefaults instantiates a new LabelDistribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabelId

`func (o *LabelDistribution) GetLabelId() string`

GetLabelId returns the LabelId field if non-nil, zero value otherwise.

### GetLabelIdOk

`func (o *LabelDistribution) GetLabelIdOk() (*string, bool)`

GetLabelIdOk returns a tuple with the LabelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelId

`func (o *LabelDistribution) SetLabelId(v string)`

SetLabelId sets LabelId field to given value.

### HasLabelId

`func (o *LabelDistribution) HasLabelId() bool`

HasLabelId returns a boolean if a field has been set.

### GetCount

`func (o *LabelDistribution) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *LabelDistribution) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *LabelDistribution) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *LabelDistribution) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetRatio

`func (o *LabelDistribution) GetRatio() float32`

GetRatio returns the Ratio field if non-nil, zero value otherwise.

### GetRatioOk

`func (o *LabelDistribution) GetRatioOk() (*float32, bool)`

GetRatioOk returns a tuple with the Ratio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatio

`func (o *LabelDistribution) SetRatio(v float32)`

SetRatio sets Ratio field to given value.

### HasRatio

`func (o *LabelDistribution) HasRatio() bool`

HasRatio returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


