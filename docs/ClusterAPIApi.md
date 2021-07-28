# \ClusterAPIApi

All URIs are relative to *https://graphhopper.com/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AsyncClusteringProblem**](ClusterAPIApi.md#AsyncClusteringProblem) | **Post** /cluster/calculate | Batch Cluster Endpoint
[**GetClusterSolution**](ClusterAPIApi.md#GetClusterSolution) | **Get** /cluster/solution/{jobId} | GET Batch Solution Endpoint
[**SolveClusteringProblem**](ClusterAPIApi.md#SolveClusteringProblem) | **Post** /cluster | POST Cluster Endpoint



## AsyncClusteringProblem

> JobId AsyncClusteringProblem(ctx).ClusterRequest(clusterRequest).Execute()

Batch Cluster Endpoint



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
    clusterRequest := *openapiclient.NewClusterRequest() // ClusterRequest | Request object that contains the problem to be solved

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterAPIApi.AsyncClusteringProblem(context.Background()).ClusterRequest(clusterRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterAPIApi.AsyncClusteringProblem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AsyncClusteringProblem`: JobId
    fmt.Fprintf(os.Stdout, "Response from `ClusterAPIApi.AsyncClusteringProblem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAsyncClusteringProblemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterRequest** | [**ClusterRequest**](ClusterRequest.md) | Request object that contains the problem to be solved | 

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


## GetClusterSolution

> ClusterResponse GetClusterSolution(ctx, jobId).Execute()

GET Batch Solution Endpoint



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
    resp, r, err := api_client.ClusterAPIApi.GetClusterSolution(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterAPIApi.GetClusterSolution``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterSolution`: ClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `ClusterAPIApi.GetClusterSolution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Request solution with jobId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterSolutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterResponse**](ClusterResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SolveClusteringProblem

> ClusterResponse SolveClusteringProblem(ctx).ClusterRequest(clusterRequest).Execute()

POST Cluster Endpoint



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
    clusterRequest := *openapiclient.NewClusterRequest() // ClusterRequest | Request object that contains the problem to be solved

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterAPIApi.SolveClusteringProblem(context.Background()).ClusterRequest(clusterRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterAPIApi.SolveClusteringProblem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SolveClusteringProblem`: ClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `ClusterAPIApi.SolveClusteringProblem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSolveClusteringProblemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterRequest** | [**ClusterRequest**](ClusterRequest.md) | Request object that contains the problem to be solved | 

### Return type

[**ClusterResponse**](ClusterResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

