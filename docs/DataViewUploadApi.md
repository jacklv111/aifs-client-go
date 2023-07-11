# \DataViewUploadApi

All URIs are relative to *https://www.example.com/api/open/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UploadAnnotationToDataView**](DataViewUploadApi.md#UploadAnnotationToDataView) | **Post** /data-views/{dataViewId}/annotations | Upload annotations to data view
[**UploadDatasetZipToDataView**](DataViewUploadApi.md#UploadDatasetZipToDataView) | **Post** /data-views/{dataViewId}/dataset-zip | Upload dataset zip
[**UploadFileToDataView**](DataViewUploadApi.md#UploadFileToDataView) | **Post** /data-views/{dataViewId}/artifact | Upload file to data view
[**UploadModelDataToDataView**](DataViewUploadApi.md#UploadModelDataToDataView) | **Post** /data-views/{dataViewId}/model | Upload model data to data view
[**UploadRawDataToDataView**](DataViewUploadApi.md#UploadRawDataToDataView) | **Post** /data-views/{dataViewId}/raw-data | Upload raw data to data view



## UploadAnnotationToDataView

> UploadAnnotationToDataView(ctx, dataViewId).FileMeta(fileMeta).Files(files).Execute()

Upload annotations to data view



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
    fileMeta := os.NewFile(1234, "some_file") // *os.File |  (optional)
    files := []*os.File{"TODO"} // []*os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DataViewUploadApi.UploadAnnotationToDataView(context.Background(), dataViewId).FileMeta(fileMeta).Files(files).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataViewUploadApi.UploadAnnotationToDataView``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUploadAnnotationToDataViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileMeta** | ***os.File** |  | 
 **files** | **[]*os.File** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data, application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadDatasetZipToDataView

> UploadDatasetZipToDataView(ctx, dataViewId).XFileName(xFileName).Body(body).Execute()

Upload dataset zip



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
    xFileName := "xFileName_example" // string | 
    body := os.NewFile(1234, "some_file") // *os.File | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DataViewUploadApi.UploadDatasetZipToDataView(context.Background(), dataViewId).XFileName(xFileName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataViewUploadApi.UploadDatasetZipToDataView``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUploadDatasetZipToDataViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFileName** | **string** |  | 
 **body** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadFileToDataView

> UploadFileToDataView(ctx, dataViewId).XFileName(xFileName).Body(body).Execute()

Upload file to data view



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
    xFileName := "xFileName_example" // string | 
    body := os.NewFile(1234, "some_file") // *os.File | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DataViewUploadApi.UploadFileToDataView(context.Background(), dataViewId).XFileName(xFileName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataViewUploadApi.UploadFileToDataView``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUploadFileToDataViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFileName** | **string** |  | 
 **body** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadModelDataToDataView

> UploadModelDataToDataView(ctx, dataViewId).CommitId(commitId).Progress(progress).ModelJit(modelJit).Onnx(onnx).DynamicOnnx(dynamicOnnx).ConfigPy(configPy).BestPth(bestPth).LastPth(lastPth).OutputTemplate(outputTemplate).Logs(logs).Execute()

Upload model data to data view



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
    commitId := "commitId_example" // string |  (optional)
    progress := "progress_example" // string |  (optional)
    modelJit := os.NewFile(1234, "some_file") // *os.File |  (optional)
    onnx := os.NewFile(1234, "some_file") // *os.File |  (optional)
    dynamicOnnx := os.NewFile(1234, "some_file") // *os.File |  (optional)
    configPy := os.NewFile(1234, "some_file") // *os.File |  (optional)
    bestPth := os.NewFile(1234, "some_file") // *os.File |  (optional)
    lastPth := os.NewFile(1234, "some_file") // *os.File |  (optional)
    outputTemplate := os.NewFile(1234, "some_file") // *os.File |  (optional)
    logs := []*os.File{"TODO"} // []*os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DataViewUploadApi.UploadModelDataToDataView(context.Background(), dataViewId).CommitId(commitId).Progress(progress).ModelJit(modelJit).Onnx(onnx).DynamicOnnx(dynamicOnnx).ConfigPy(configPy).BestPth(bestPth).LastPth(lastPth).OutputTemplate(outputTemplate).Logs(logs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataViewUploadApi.UploadModelDataToDataView``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUploadModelDataToDataViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commitId** | **string** |  | 
 **progress** | **string** |  | 
 **modelJit** | ***os.File** |  | 
 **onnx** | ***os.File** |  | 
 **dynamicOnnx** | ***os.File** |  | 
 **configPy** | ***os.File** |  | 
 **bestPth** | ***os.File** |  | 
 **lastPth** | ***os.File** |  | 
 **outputTemplate** | ***os.File** |  | 
 **logs** | **[]*os.File** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadRawDataToDataView

> UploadRawDataToDataView(ctx, dataViewId).Files(files).Execute()

Upload raw data to data view



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
    files := []*os.File{"TODO"} // []*os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DataViewUploadApi.UploadRawDataToDataView(context.Background(), dataViewId).Files(files).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataViewUploadApi.UploadRawDataToDataView``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUploadRawDataToDataViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **files** | **[]*os.File** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data, application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

