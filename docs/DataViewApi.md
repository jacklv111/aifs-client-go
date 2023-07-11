# \DataViewApi

All URIs are relative to *https://www.example.com/api/open/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDataView**](DataViewApi.md#CreateDataView) | **Post** /data-views | Create a data view
[**DeleteDataItemInDataView**](DataViewApi.md#DeleteDataItemInDataView) | **Delete** /data-views/{dataViewId}/data-items | Delete data item in a data view
[**DeleteDataView**](DataViewApi.md#DeleteDataView) | **Delete** /data-views/{dataViewId} | Delete a data view
[**DivideDataView**](DataViewApi.md#DivideDataView) | **Post** /data-views/{dataViewId}/raw-data-divide | Divide data view
[**FilterAnnotationsInDataView**](DataViewApi.md#FilterAnnotationsInDataView) | **Post** /data-views/{dataViewId}/annotation-filter | Filter annotations in a data view
[**GetAllAnnotationDataInDataView**](DataViewApi.md#GetAllAnnotationDataInDataView) | **Get** /data-views/{dataViewId}/annotation-data | Get all annotation data in a data view
[**GetAllAnnotationLocationsInDataView**](DataViewApi.md#GetAllAnnotationLocationsInDataView) | **Get** /data-views/{dataViewId}/annotation-locations | Get all annotation locations in a data view
[**GetAllRawDataLocationsInDataView**](DataViewApi.md#GetAllRawDataLocationsInDataView) | **Get** /data-views/{dataViewId}/raw-data-locations | Get all raw data locations in a data view
[**GetAnnotationsInDataView**](DataViewApi.md#GetAnnotationsInDataView) | **Get** /data-views/{dataViewId}/annotations | Get data view annotations
[**GetArtifactLocationsInDataView**](DataViewApi.md#GetArtifactLocationsInDataView) | **Get** /data-views/{dataViewId}/artifact-locations | Get files&#39; locations in artifact data view
[**GetDataViewDetails**](DataViewApi.md#GetDataViewDetails) | **Get** /data-views/{dataViewId}/details | Get data view details
[**GetDataViewList**](DataViewApi.md#GetDataViewList) | **Get** /data-views | Get data view list
[**GetDataViewStatistics**](DataViewApi.md#GetDataViewStatistics) | **Get** /data-views/{dataViewId}/statistics | Get data view statistics
[**GetDatasetZipLocationInDataView**](DataViewApi.md#GetDatasetZipLocationInDataView) | **Get** /data-views/{dataViewId}/dataset-zip-location | Get dataset zip&#39;s location in a data view
[**GetModelDataLocationsInDataView**](DataViewApi.md#GetModelDataLocationsInDataView) | **Get** /data-views/{dataViewId}/model-data-locations | Get all model data locations in a data view
[**GetRawDataHashListInDataView**](DataViewApi.md#GetRawDataHashListInDataView) | **Get** /data-views/{dataViewId}/raw-data-hash-list | Get data view raw data hash list
[**GetRawDataInDataView**](DataViewApi.md#GetRawDataInDataView) | **Get** /data-views/{dataViewId}/raw-data | Get data view raw data
[**HardDeleteDataView**](DataViewApi.md#HardDeleteDataView) | **Delete** /data-views/{dataViewId}/hard-delete | Hard delete a data view
[**MergeDataViews**](DataViewApi.md#MergeDataViews) | **Post** /data-views/merge | Merge data views
[**MergeDataViewsToCrurrent**](DataViewApi.md#MergeDataViewsToCrurrent) | **Post** /data-views/{dataViewId}/merge | Merge data views
[**MoveDataViewItems**](DataViewApi.md#MoveDataViewItems) | **Post** /data-views/move | Move data items between data views
[**UpdateDatasetZipView**](DataViewApi.md#UpdateDatasetZipView) | **Put** /data-views/{dataViewId}/dataset-zip | Update a dataset-zip view meta



## CreateDataView

> CreateDataViewSuccessResp CreateDataView(ctx).CreateDataViewRequest(createDataViewRequest).Execute()

Create a data view



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createDataViewRequest := *openapiclient.NewCreateDataViewRequest() // CreateDataViewRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataViewApi.CreateDataView(context.Background()).CreateDataViewRequest(createDataViewRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataViewApi.CreateDataView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDataView`: CreateDataViewSuccessResp
    fmt.Fprintf(os.Stdout, "Response from `DataViewApi.CreateDataView`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDataViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDataViewRequest** | [**CreateDataViewRequest**](CreateDataViewRequest.md) |  | 

### Return type

[**CreateDataViewSuccessResp**](CreateDataViewSuccessResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDataItemInDataView

> DeleteDataItemInDataView(ctx, dataViewId).DataViewItemIdList(dataViewItemIdList).Execute()

Delete data item in a data view



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    dataViewId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of a data view
    dataViewItemIdList := []string{"Inner_example"} // []string | ids of raw data which are expected to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DataViewApi.DeleteDataItemInDataView(context.Background(), dataViewId).DataViewItemIdList(dataViewItemIdList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataViewApi.DeleteDataItemInDataView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataViewId** | **string** | The id of a data view | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDataItemInDataViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataViewItemIdList** | **[]string** | ids of raw data which are expected to be deleted | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDataView

> DeleteDataView(ctx, dataViewId).Execute()

Delete a data view



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    dataViewId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of a data view

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DataViewApi.DeleteDataView(context.Background(), dataViewId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataViewApi.DeleteDataView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataViewId** | **string** | The id of a data view | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDataViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DivideDataView

> []DivideRawDataDataViewResponseInner DivideDataView(ctx, dataViewId).DivideRawDataDataViewRequestInner(divideRawDataDataViewRequestInner).Execute()

Divide data view



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    dataViewId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of a data view
    divideRawDataDataViewRequestInner := []openapiclient.DivideRawDataDataViewRequestInner{*openapiclient.NewDivideRawDataDataViewRequestInner()} // []DivideRawDataDataViewRequestInner | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataViewApi.DivideDataView(context.Background(), dataViewId).DivideRawDataDataViewRequestInner(divideRawDataDataViewRequestInner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataViewApi.DivideDataView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DivideDataView`: []DivideRawDataDataViewResponseInner
    fmt.Fprintf(os.Stdout, "Response from `DataViewApi.DivideDataView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataViewId** | **string** | The id of a data view | 

### Other Parameters

Other parameters are passed through a pointer to a apiDivideDataViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **divideRawDataDataViewRequestInner** | [**[]DivideRawDataDataViewRequestInner**](DivideRawDataDataViewRequestInner.md) |  | 

### Return type

[**[]DivideRawDataDataViewResponseInner**](DivideRawDataDataViewResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilterAnnotationsInDataView

> FilterAnnotationsInDataViewResponse FilterAnnotationsInDataView(ctx, dataViewId).FilterAnnotationsInDataViewRequest(filterAnnotationsInDataViewRequest).Execute()

Filter annotations in a data view



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    dataViewId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of a data view
    filterAnnotationsInDataViewRequest := *openapiclient.NewFilterAnnotationsInDataViewRequest() // FilterAnnotationsInDataViewRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataViewApi.FilterAnnotationsInDataView(context.Background(), dataViewId).FilterAnnotationsInDataViewRequest(filterAnnotationsInDataViewRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataViewApi.FilterAnnotationsInDataView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilterAnnotationsInDataView`: FilterAnnotationsInDataViewResponse
    fmt.Fprintf(os.Stdout, "Response from `DataViewApi.FilterAnnotationsInDataView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataViewId** | **string** | The id of a data view | 

### Other Parameters

Other parameters are passed through a pointer to a apiFilterAnnotationsInDataViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterAnnotationsInDataViewRequest** | [**FilterAnnotationsInDataViewRequest**](FilterAnnotationsInDataViewRequest.md) |  | 

### Return type

[**FilterAnnotationsInDataViewResponse**](FilterAnnotationsInDataViewResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllAnnotationDataInDataView

> AnnotationViewData GetAllAnnotationDataInDataView(ctx, dataViewId).Execute()

Get all annotation data in a data view



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    dataViewId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of a data view

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataViewApi.GetAllAnnotationDataInDataView(context.Background(), dataViewId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataViewApi.GetAllAnnotationDataInDataView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllAnnotationDataInDataView`: AnnotationViewData
    fmt.Fprintf(os.Stdout, "Response from `DataViewApi.GetAllAnnotationDataInDataView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataViewId** | **string** | The id of a data view | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllAnnotationDataInDataViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AnnotationViewData**](AnnotationViewData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllAnnotationLocationsInDataView

> AnnotationViewLocations GetAllAnnotationLocationsInDataView(ctx, dataViewId).Execute()

Get all annotation locations in a data view



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    dataViewId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of a data view

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataViewApi.GetAllAnnotationLocationsInDataView(context.Background(), dataViewId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataViewApi.GetAllAnnotationLocationsInDataView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllAnnotationLocationsInDataView`: AnnotationViewLocations
    fmt.Fprintf(os.Stdout, "Response from `DataViewApi.GetAllAnnotationLocationsInDataView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataViewId** | **string** | The id of a data view | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllAnnotationLocationsInDataViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AnnotationViewLocations**](AnnotationViewLocations.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllRawDataLocationsInDataView

> RawDataViewLocations GetAllRawDataLocationsInDataView(ctx, dataViewId).Execute()

Get all raw data locations in a data view



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    dataViewId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of a data view

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataViewApi.GetAllRawDataLocationsInDataView(context.Background(), dataViewId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataViewApi.GetAllRawDataLocationsInDataView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllRawDataLocationsInDataView`: RawDataViewLocations
    fmt.Fprintf(os.Stdout, "Response from `DataViewApi.GetAllRawDataLocationsInDataView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataViewId** | **string** | The id of a data view | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllRawDataLocationsInDataViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RawDataViewLocations**](RawDataViewLocations.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnnotationsInDataView

> GetAnnotationsInDataView200Response GetAnnotationsInDataView(ctx, dataViewId).Offset(offset).Limit(limit).RawDataIdList(rawDataIdList).LabelId(labelId).Execute()

Get data view annotations



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    dataViewId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of a data view
    offset := int32(56) // int32 | The number of items to skip before starting to collect the result set (optional) (default to 0)
    limit := int32(56) // int32 | The numbers of items to return (optional) (default to 10)
    rawDataIdList := []string{"Inner_example"} // []string | ids of raw data (optional)
    labelId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataViewApi.GetAnnotationsInDataView(context.Background(), dataViewId).Offset(offset).Limit(limit).RawDataIdList(rawDataIdList).LabelId(labelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataViewApi.GetAnnotationsInDataView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnnotationsInDataView`: GetAnnotationsInDataView200Response
    fmt.Fprintf(os.Stdout, "Response from `DataViewApi.GetAnnotationsInDataView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataViewId** | **string** | The id of a data view | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnnotationsInDataViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | The number of items to skip before starting to collect the result set | [default to 0]
 **limit** | **int32** | The numbers of items to return | [default to 10]
 **rawDataIdList** | **[]string** | ids of raw data | 
 **labelId** | **string** |  | 

### Return type

[**GetAnnotationsInDataView200Response**](GetAnnotationsInDataView200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArtifactLocationsInDataView

> ArtifactLocations GetArtifactLocationsInDataView(ctx, dataViewId).Execute()

Get files' locations in artifact data view



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    dataViewId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of a data view

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataViewApi.GetArtifactLocationsInDataView(context.Background(), dataViewId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataViewApi.GetArtifactLocationsInDataView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetArtifactLocationsInDataView`: ArtifactLocations
    fmt.Fprintf(os.Stdout, "Response from `DataViewApi.GetArtifactLocationsInDataView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataViewId** | **string** | The id of a data view | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArtifactLocationsInDataViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ArtifactLocations**](ArtifactLocations.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataViewDetails

> DataViewDetails GetDataViewDetails(ctx, dataViewId).Execute()

Get data view details



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    dataViewId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of a data view

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataViewApi.GetDataViewDetails(context.Background(), dataViewId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataViewApi.GetDataViewDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataViewDetails`: DataViewDetails
    fmt.Fprintf(os.Stdout, "Response from `DataViewApi.GetDataViewDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataViewId** | **string** | The id of a data view | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataViewDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataViewDetails**](DataViewDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataViewList

> []DataViewListItem GetDataViewList(ctx).Offset(offset).Limit(limit).DataViewIdList(dataViewIdList).DataViewName(dataViewName).Execute()

Get data view list



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    offset := int32(56) // int32 | The number of items to skip before starting to collect the result set (optional) (default to 0)
    limit := int32(56) // int32 | The numbers of items to return (optional) (default to 10)
    dataViewIdList := "aaa,bbb,ccc" // string | data view with id in data view id list will be got (optional)
    dataViewName := "hand" // string | data view name filter, support fuzzy query (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataViewApi.GetDataViewList(context.Background()).Offset(offset).Limit(limit).DataViewIdList(dataViewIdList).DataViewName(dataViewName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataViewApi.GetDataViewList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataViewList`: []DataViewListItem
    fmt.Fprintf(os.Stdout, "Response from `DataViewApi.GetDataViewList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDataViewListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | The number of items to skip before starting to collect the result set | [default to 0]
 **limit** | **int32** | The numbers of items to return | [default to 10]
 **dataViewIdList** | **string** | data view with id in data view id list will be got | 
 **dataViewName** | **string** | data view name filter, support fuzzy query | 

### Return type

[**[]DataViewListItem**](DataViewListItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataViewStatistics

> DataViewStatistics GetDataViewStatistics(ctx, dataViewId).Execute()

Get data view statistics



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    dataViewId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of a data view

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataViewApi.GetDataViewStatistics(context.Background(), dataViewId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataViewApi.GetDataViewStatistics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataViewStatistics`: DataViewStatistics
    fmt.Fprintf(os.Stdout, "Response from `DataViewApi.GetDataViewStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataViewId** | **string** | The id of a data view | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataViewStatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataViewStatistics**](DataViewStatistics.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatasetZipLocationInDataView

> DatasetZipLocation GetDatasetZipLocationInDataView(ctx, dataViewId).Execute()

Get dataset zip's location in a data view



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    dataViewId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of a data view

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataViewApi.GetDatasetZipLocationInDataView(context.Background(), dataViewId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataViewApi.GetDatasetZipLocationInDataView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatasetZipLocationInDataView`: DatasetZipLocation
    fmt.Fprintf(os.Stdout, "Response from `DataViewApi.GetDatasetZipLocationInDataView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataViewId** | **string** | The id of a data view | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatasetZipLocationInDataViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DatasetZipLocation**](DatasetZipLocation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetModelDataLocationsInDataView

> ModelDataViewLocations GetModelDataLocationsInDataView(ctx, dataViewId).Execute()

Get all model data locations in a data view



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    dataViewId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of a data view

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataViewApi.GetModelDataLocationsInDataView(context.Background(), dataViewId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataViewApi.GetModelDataLocationsInDataView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetModelDataLocationsInDataView`: ModelDataViewLocations
    fmt.Fprintf(os.Stdout, "Response from `DataViewApi.GetModelDataLocationsInDataView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataViewId** | **string** | The id of a data view | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetModelDataLocationsInDataViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelDataViewLocations**](ModelDataViewLocations.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRawDataHashListInDataView

> []RawDataHashListInner GetRawDataHashListInDataView(ctx, dataViewId).Offset(offset).Limit(limit).Execute()

Get data view raw data hash list



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    dataViewId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of a data view
    offset := int32(56) // int32 | The number of items to skip before starting to collect the result set (optional) (default to 0)
    limit := int32(56) // int32 | The numbers of items to return (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataViewApi.GetRawDataHashListInDataView(context.Background(), dataViewId).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataViewApi.GetRawDataHashListInDataView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRawDataHashListInDataView`: []RawDataHashListInner
    fmt.Fprintf(os.Stdout, "Response from `DataViewApi.GetRawDataHashListInDataView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataViewId** | **string** | The id of a data view | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRawDataHashListInDataViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | The number of items to skip before starting to collect the result set | [default to 0]
 **limit** | **int32** | The numbers of items to return | [default to 10]

### Return type

[**[]RawDataHashListInner**](RawDataHashListInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRawDataInDataView

> GetRawDataInDataView200Response GetRawDataInDataView(ctx, dataViewId).Offset(offset).Limit(limit).RawDataIdList(rawDataIdList).ExcludedAnnotationViewId(excludedAnnotationViewId).IncludedAnnotationViewId(includedAnnotationViewId).Execute()

Get data view raw data



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    dataViewId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of a data view
    offset := int32(56) // int32 | The number of items to skip before starting to collect the result set (optional) (default to 0)
    limit := int32(56) // int32 | The numbers of items to return (optional) (default to 10)
    rawDataIdList := []string{"Inner_example"} // []string | ids of raw data (optional)
    excludedAnnotationViewId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | excluded annotation view with id. Return the raw data items which have no annotation in the annotation view. (optional)
    includedAnnotationViewId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | included annotation view with id. Return the raw data items which have annotation in the annotation view. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataViewApi.GetRawDataInDataView(context.Background(), dataViewId).Offset(offset).Limit(limit).RawDataIdList(rawDataIdList).ExcludedAnnotationViewId(excludedAnnotationViewId).IncludedAnnotationViewId(includedAnnotationViewId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataViewApi.GetRawDataInDataView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRawDataInDataView`: GetRawDataInDataView200Response
    fmt.Fprintf(os.Stdout, "Response from `DataViewApi.GetRawDataInDataView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataViewId** | **string** | The id of a data view | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRawDataInDataViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | The number of items to skip before starting to collect the result set | [default to 0]
 **limit** | **int32** | The numbers of items to return | [default to 10]
 **rawDataIdList** | **[]string** | ids of raw data | 
 **excludedAnnotationViewId** | **string** | excluded annotation view with id. Return the raw data items which have no annotation in the annotation view. | 
 **includedAnnotationViewId** | **string** | included annotation view with id. Return the raw data items which have annotation in the annotation view. | 

### Return type

[**GetRawDataInDataView200Response**](GetRawDataInDataView200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HardDeleteDataView

> HardDeleteDataView(ctx, dataViewId).Execute()

Hard delete a data view



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    dataViewId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of a data view

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DataViewApi.HardDeleteDataView(context.Background(), dataViewId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataViewApi.HardDeleteDataView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataViewId** | **string** | The id of a data view | 

### Other Parameters

Other parameters are passed through a pointer to a apiHardDeleteDataViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MergeDataViews

> MergeDataViewsSuccessResp MergeDataViews(ctx).MergeDataViewsRequest(mergeDataViewsRequest).Execute()

Merge data views



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    mergeDataViewsRequest := *openapiclient.NewMergeDataViewsRequest() // MergeDataViewsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataViewApi.MergeDataViews(context.Background()).MergeDataViewsRequest(mergeDataViewsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataViewApi.MergeDataViews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MergeDataViews`: MergeDataViewsSuccessResp
    fmt.Fprintf(os.Stdout, "Response from `DataViewApi.MergeDataViews`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMergeDataViewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mergeDataViewsRequest** | [**MergeDataViewsRequest**](MergeDataViewsRequest.md) |  | 

### Return type

[**MergeDataViewsSuccessResp**](MergeDataViewsSuccessResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MergeDataViewsToCrurrent

> MergeDataViewsToCrurrent(ctx, dataViewId).MergeDataViewsRequest(mergeDataViewsRequest).Execute()

Merge data views



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    dataViewId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of a data view
    mergeDataViewsRequest := *openapiclient.NewMergeDataViewsRequest() // MergeDataViewsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DataViewApi.MergeDataViewsToCrurrent(context.Background(), dataViewId).MergeDataViewsRequest(mergeDataViewsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataViewApi.MergeDataViewsToCrurrent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataViewId** | **string** | The id of a data view | 

### Other Parameters

Other parameters are passed through a pointer to a apiMergeDataViewsToCrurrentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mergeDataViewsRequest** | [**MergeDataViewsRequest**](MergeDataViewsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveDataViewItems

> MoveDataViewItems(ctx).MoveDataViewItemsRequest(moveDataViewItemsRequest).Execute()

Move data items between data views



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    moveDataViewItemsRequest := *openapiclient.NewMoveDataViewItemsRequest() // MoveDataViewItemsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DataViewApi.MoveDataViewItems(context.Background()).MoveDataViewItemsRequest(moveDataViewItemsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataViewApi.MoveDataViewItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMoveDataViewItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **moveDataViewItemsRequest** | [**MoveDataViewItemsRequest**](MoveDataViewItemsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDatasetZipView

> UpdateDatasetZipView(ctx, dataViewId).UpdateDatasetZipRequest(updateDatasetZipRequest).Execute()

Update a dataset-zip view meta



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    dataViewId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of a data view
    updateDatasetZipRequest := *openapiclient.NewUpdateDatasetZipRequest() // UpdateDatasetZipRequest | Update an existed dataset-zip view's meta

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DataViewApi.UpdateDatasetZipView(context.Background(), dataViewId).UpdateDatasetZipRequest(updateDatasetZipRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataViewApi.UpdateDatasetZipView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataViewId** | **string** | The id of a data view | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDatasetZipViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDatasetZipRequest** | [**UpdateDatasetZipRequest**](UpdateDatasetZipRequest.md) | Update an existed dataset-zip view&#39;s meta | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

