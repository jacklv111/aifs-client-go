# ImageData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to **string** |  | [optional] 
**Data** | Pointer to ***os.File** |  | [optional] 

## Methods

### NewImageData

`func NewImageData() *ImageData`

NewImageData instantiates a new ImageData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageDataWithDefaults

`func NewImageDataWithDefaults() *ImageData`

NewImageDataWithDefaults instantiates a new ImageData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *ImageData) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ImageData) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ImageData) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *ImageData) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetData

`func (o *ImageData) GetData() *os.File`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ImageData) GetDataOk() (**os.File, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ImageData) SetData(v *os.File)`

SetData sets Data field to given value.

### HasData

`func (o *ImageData) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


