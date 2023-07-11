# \AnnotationTemplateApi

All URIs are relative to *https://www.example.com/api/open/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CopyAnnotationTemplate**](AnnotationTemplateApi.md#CopyAnnotationTemplate) | **Post** /annotation-templates/{annotationTemplateId}/copy | Copy an annotation template
[**CreateAnnotationTemplate**](AnnotationTemplateApi.md#CreateAnnotationTemplate) | **Post** /annotation-templates | Create an annotation template
[**DeleteAnnotationTemplate**](AnnotationTemplateApi.md#DeleteAnnotationTemplate) | **Delete** /annotation-templates/{annotationTemplateId} | Delete an annotation template
[**GetAnnoTemplateDetails**](AnnotationTemplateApi.md#GetAnnoTemplateDetails) | **Get** /annotation-templates/{annotationTemplateId}/details | Get annotation template details
[**GetAnnoTemplateList**](AnnotationTemplateApi.md#GetAnnoTemplateList) | **Get** /annotation-templates | Get annotation template list
[**UpdateAnnotationTemplate**](AnnotationTemplateApi.md#UpdateAnnotationTemplate) | **Put** /annotation-templates | Update an annotation template



## CopyAnnotationTemplate

> CopyAnnotationTemplate200Response CopyAnnotationTemplate(ctx, annotationTemplateId).Execute()

Copy an annotation template



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
    annotationTemplateId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of an annotation template

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnnotationTemplateApi.CopyAnnotationTemplate(context.Background(), annotationTemplateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnnotationTemplateApi.CopyAnnotationTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CopyAnnotationTemplate`: CopyAnnotationTemplate200Response
    fmt.Fprintf(os.Stdout, "Response from `AnnotationTemplateApi.CopyAnnotationTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**annotationTemplateId** | **string** | The id of an annotation template | 

### Other Parameters

Other parameters are passed through a pointer to a apiCopyAnnotationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CopyAnnotationTemplate200Response**](CopyAnnotationTemplate200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAnnotationTemplate

> CreateAnnoTemplateSuccessResp CreateAnnotationTemplate(ctx).CreateAnnotationTemplateRequest(createAnnotationTemplateRequest).Execute()

Create an annotation template



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
    createAnnotationTemplateRequest := *openapiclient.NewCreateAnnotationTemplateRequest("Name_example", "Type_example") // CreateAnnotationTemplateRequest | Create an new annotation template

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnnotationTemplateApi.CreateAnnotationTemplate(context.Background()).CreateAnnotationTemplateRequest(createAnnotationTemplateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnnotationTemplateApi.CreateAnnotationTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAnnotationTemplate`: CreateAnnoTemplateSuccessResp
    fmt.Fprintf(os.Stdout, "Response from `AnnotationTemplateApi.CreateAnnotationTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAnnotationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAnnotationTemplateRequest** | [**CreateAnnotationTemplateRequest**](CreateAnnotationTemplateRequest.md) | Create an new annotation template | 

### Return type

[**CreateAnnoTemplateSuccessResp**](CreateAnnoTemplateSuccessResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAnnotationTemplate

> DeleteAnnotationTemplate(ctx, annotationTemplateId).Execute()

Delete an annotation template



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
    annotationTemplateId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of an annotation template

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AnnotationTemplateApi.DeleteAnnotationTemplate(context.Background(), annotationTemplateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnnotationTemplateApi.DeleteAnnotationTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**annotationTemplateId** | **string** | The id of an annotation template | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAnnotationTemplateRequest struct via the builder pattern


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


## GetAnnoTemplateDetails

> AnnotationTemplateDetails GetAnnoTemplateDetails(ctx, annotationTemplateId).Execute()

Get annotation template details



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
    annotationTemplateId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of an annotation template

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnnotationTemplateApi.GetAnnoTemplateDetails(context.Background(), annotationTemplateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnnotationTemplateApi.GetAnnoTemplateDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnnoTemplateDetails`: AnnotationTemplateDetails
    fmt.Fprintf(os.Stdout, "Response from `AnnotationTemplateApi.GetAnnoTemplateDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**annotationTemplateId** | **string** | The id of an annotation template | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnnoTemplateDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AnnotationTemplateDetails**](AnnotationTemplateDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnnoTemplateList

> []AnnotationTemplateListItem GetAnnoTemplateList(ctx).Offset(offset).Limit(limit).AnnotationTemplateIdList(annotationTemplateIdList).Execute()

Get annotation template list



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
    annotationTemplateIdList := "aaa,bbb,ccc" // string | annotation template with id in annotation template id list will be got (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnnotationTemplateApi.GetAnnoTemplateList(context.Background()).Offset(offset).Limit(limit).AnnotationTemplateIdList(annotationTemplateIdList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnnotationTemplateApi.GetAnnoTemplateList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnnoTemplateList`: []AnnotationTemplateListItem
    fmt.Fprintf(os.Stdout, "Response from `AnnotationTemplateApi.GetAnnoTemplateList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAnnoTemplateListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | The number of items to skip before starting to collect the result set | [default to 0]
 **limit** | **int32** | The numbers of items to return | [default to 10]
 **annotationTemplateIdList** | **string** | annotation template with id in annotation template id list will be got | 

### Return type

[**[]AnnotationTemplateListItem**](AnnotationTemplateListItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAnnotationTemplate

> UpdateAnnotationTemplate(ctx).UpdateAnnotationTemplateRequest(updateAnnotationTemplateRequest).Execute()

Update an annotation template



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
    updateAnnotationTemplateRequest := *openapiclient.NewUpdateAnnotationTemplateRequest("Name_example", "Type_example") // UpdateAnnotationTemplateRequest | Update an existed annotation template

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AnnotationTemplateApi.UpdateAnnotationTemplate(context.Background()).UpdateAnnotationTemplateRequest(updateAnnotationTemplateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnnotationTemplateApi.UpdateAnnotationTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAnnotationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateAnnotationTemplateRequest** | [**UpdateAnnotationTemplateRequest**](UpdateAnnotationTemplateRequest.md) | Update an existed annotation template | 

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

