# MoveDataViewItemsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SrcDataViewId** | Pointer to **string** | the id of the source data view | [optional] 
**DstDataViewId** | Pointer to **string** | the id of the destination data view | [optional] 

## Methods

### NewMoveDataViewItemsRequest

`func NewMoveDataViewItemsRequest() *MoveDataViewItemsRequest`

NewMoveDataViewItemsRequest instantiates a new MoveDataViewItemsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoveDataViewItemsRequestWithDefaults

`func NewMoveDataViewItemsRequestWithDefaults() *MoveDataViewItemsRequest`

NewMoveDataViewItemsRequestWithDefaults instantiates a new MoveDataViewItemsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSrcDataViewId

`func (o *MoveDataViewItemsRequest) GetSrcDataViewId() string`

GetSrcDataViewId returns the SrcDataViewId field if non-nil, zero value otherwise.

### GetSrcDataViewIdOk

`func (o *MoveDataViewItemsRequest) GetSrcDataViewIdOk() (*string, bool)`

GetSrcDataViewIdOk returns a tuple with the SrcDataViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcDataViewId

`func (o *MoveDataViewItemsRequest) SetSrcDataViewId(v string)`

SetSrcDataViewId sets SrcDataViewId field to given value.

### HasSrcDataViewId

`func (o *MoveDataViewItemsRequest) HasSrcDataViewId() bool`

HasSrcDataViewId returns a boolean if a field has been set.

### GetDstDataViewId

`func (o *MoveDataViewItemsRequest) GetDstDataViewId() string`

GetDstDataViewId returns the DstDataViewId field if non-nil, zero value otherwise.

### GetDstDataViewIdOk

`func (o *MoveDataViewItemsRequest) GetDstDataViewIdOk() (*string, bool)`

GetDstDataViewIdOk returns a tuple with the DstDataViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstDataViewId

`func (o *MoveDataViewItemsRequest) SetDstDataViewId(v string)`

SetDstDataViewId sets DstDataViewId field to given value.

### HasDstDataViewId

`func (o *MoveDataViewItemsRequest) HasDstDataViewId() bool`

HasDstDataViewId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


