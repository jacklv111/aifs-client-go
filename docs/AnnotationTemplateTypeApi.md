# \AnnotationTemplateTypeApi

All URIs are relative to *https://www.example.com/api/open/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAnnoTemplateTypeList**](AnnotationTemplateTypeApi.md#GetAnnoTemplateTypeList) | **Get** /annotation-template-types | Get annotation template type list



## GetAnnoTemplateTypeList

> []string GetAnnoTemplateTypeList(ctx).Offset(offset).Limit(limit).Execute()

Get annotation template type list



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnnotationTemplateTypeApi.GetAnnoTemplateTypeList(context.Background()).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnnotationTemplateTypeApi.GetAnnoTemplateTypeList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnnoTemplateTypeList`: []string
    fmt.Fprintf(os.Stdout, "Response from `AnnotationTemplateTypeApi.GetAnnoTemplateTypeList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAnnoTemplateTypeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | The number of items to skip before starting to collect the result set | [default to 0]
 **limit** | **int32** | The numbers of items to return | [default to 10]

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

