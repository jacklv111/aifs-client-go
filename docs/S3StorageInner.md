# S3StorageInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataItemId** | Pointer to **string** |  | [optional] 
**DataName** | Pointer to **string** |  | [optional] 
**ObjectKey** | Pointer to **string** | the object key of the data in s3 storage | [optional] 

## Methods

### NewS3StorageInner

`func NewS3StorageInner() *S3StorageInner`

NewS3StorageInner instantiates a new S3StorageInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3StorageInnerWithDefaults

`func NewS3StorageInnerWithDefaults() *S3StorageInner`

NewS3StorageInnerWithDefaults instantiates a new S3StorageInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataItemId

`func (o *S3StorageInner) GetDataItemId() string`

GetDataItemId returns the DataItemId field if non-nil, zero value otherwise.

### GetDataItemIdOk

`func (o *S3StorageInner) GetDataItemIdOk() (*string, bool)`

GetDataItemIdOk returns a tuple with the DataItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataItemId

`func (o *S3StorageInner) SetDataItemId(v string)`

SetDataItemId sets DataItemId field to given value.

### HasDataItemId

`func (o *S3StorageInner) HasDataItemId() bool`

HasDataItemId returns a boolean if a field has been set.

### GetDataName

`func (o *S3StorageInner) GetDataName() string`

GetDataName returns the DataName field if non-nil, zero value otherwise.

### GetDataNameOk

`func (o *S3StorageInner) GetDataNameOk() (*string, bool)`

GetDataNameOk returns a tuple with the DataName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataName

`func (o *S3StorageInner) SetDataName(v string)`

SetDataName sets DataName field to given value.

### HasDataName

`func (o *S3StorageInner) HasDataName() bool`

HasDataName returns a boolean if a field has been set.

### GetObjectKey

`func (o *S3StorageInner) GetObjectKey() string`

GetObjectKey returns the ObjectKey field if non-nil, zero value otherwise.

### GetObjectKeyOk

`func (o *S3StorageInner) GetObjectKeyOk() (*string, bool)`

GetObjectKeyOk returns a tuple with the ObjectKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectKey

`func (o *S3StorageInner) SetObjectKey(v string)`

SetObjectKey sets ObjectKey field to given value.

### HasObjectKey

`func (o *S3StorageInner) HasObjectKey() bool`

HasObjectKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


