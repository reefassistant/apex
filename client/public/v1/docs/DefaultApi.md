# \DefaultApi

All URIs are relative to *http://apex.local*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDatalog**](DefaultApi.md#GetDatalog) | **Get** /cgi-bin/datalog.json | Fetch input log records.
[**GetOutlog**](DefaultApi.md#GetOutlog) | **Get** /cgi-bin/outlog.json | Fetch output log records.
[**GetStatus**](DefaultApi.md#GetStatus) | **Get** /cgi-bin/status.json | Fetch system info and current input/output states.



## GetDatalog

> DatalogResponse GetDatalog(ctx).Sdate(sdate).Days(days).Execute()

Fetch input log records.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sdate := "sdate_example" // string | Start date in YYMMDD format. (optional)
    days := int32(56) // int32 | Number of days from start date. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetDatalog(context.Background()).Sdate(sdate).Days(days).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetDatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatalog`: DatalogResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetDatalog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sdate** | **string** | Start date in YYMMDD format. | 
 **days** | **int32** | Number of days from start date. | 

### Return type

[**DatalogResponse**](DatalogResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOutlog

> OutlogResponse GetOutlog(ctx).Sdate(sdate).Days(days).Execute()

Fetch output log records.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sdate := "sdate_example" // string | Start date in YYMMDD format. (optional)
    days := int32(56) // int32 | Number of days from start date. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetOutlog(context.Background()).Sdate(sdate).Days(days).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetOutlog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOutlog`: OutlogResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetOutlog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOutlogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sdate** | **string** | Start date in YYMMDD format. | 
 **days** | **int32** | Number of days from start date. | 

### Return type

[**OutlogResponse**](OutlogResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatus

> StatusResponse GetStatus(ctx).Execute()

Fetch system info and current input/output states.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatus`: StatusResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatusRequest struct via the builder pattern


### Return type

[**StatusResponse**](StatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

