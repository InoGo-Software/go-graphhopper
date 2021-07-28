# \RouteOptimizationAPIApi

All URIs are relative to *https://graphhopper.com/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AsyncVRP**](RouteOptimizationAPIApi.md#AsyncVRP) | **Post** /vrp/optimize | POST route optimization problem (batch mode)
[**GetSolution**](RouteOptimizationAPIApi.md#GetSolution) | **Get** /vrp/solution/{jobId} | GET the solution (batch mode)
[**SolveVRP**](RouteOptimizationAPIApi.md#SolveVRP) | **Post** /vrp | POST route optimization problem



## AsyncVRP

> JobId AsyncVRP(ctx).Request(request).Execute()

POST route optimization problem (batch mode)



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
    request := *openapiclient.NewRequest() // Request | The request that contains the problem to be solved.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RouteOptimizationAPIApi.AsyncVRP(context.Background()).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RouteOptimizationAPIApi.AsyncVRP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AsyncVRP`: JobId
    fmt.Fprintf(os.Stdout, "Response from `RouteOptimizationAPIApi.AsyncVRP`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAsyncVRPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**Request**](Request.md) | The request that contains the problem to be solved. | 

### Return type

[**JobId**](JobId.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSolution

> Response GetSolution(ctx, jobId).Execute()

GET the solution (batch mode)



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
    jobId := "jobId_example" // string | Request solution with jobId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RouteOptimizationAPIApi.GetSolution(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RouteOptimizationAPIApi.GetSolution``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSolution`: Response
    fmt.Fprintf(os.Stdout, "Response from `RouteOptimizationAPIApi.GetSolution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Request solution with jobId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSolutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Response**](Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SolveVRP

> Response SolveVRP(ctx).Request(request).Execute()

POST route optimization problem



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
    request := *openapiclient.NewRequest() // Request | The request that contains the vehicle routing problem to be solved.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RouteOptimizationAPIApi.SolveVRP(context.Background()).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RouteOptimizationAPIApi.SolveVRP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SolveVRP`: Response
    fmt.Fprintf(os.Stdout, "Response from `RouteOptimizationAPIApi.SolveVRP`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSolveVRPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**Request**](Request.md) | The request that contains the vehicle routing problem to be solved. | 

### Return type

[**Response**](Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

