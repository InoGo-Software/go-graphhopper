# \MapMatchingAPIApi

All URIs are relative to *https://graphhopper.com/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostGPX**](MapMatchingAPIApi.md#PostGPX) | **Post** /match | Map-match a GPX file



## PostGPX

> RouteResponse PostGPX(ctx).GpsAccuracy(gpsAccuracy).Vehicle(vehicle).Execute()

Map-match a GPX file



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
    gpsAccuracy := int32(56) // int32 | Specify the precision of a point, in meter (optional)
    vehicle := "vehicle_example" // string | Specify the vehicle profile like car (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MapMatchingAPIApi.PostGPX(context.Background()).GpsAccuracy(gpsAccuracy).Vehicle(vehicle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MapMatchingAPIApi.PostGPX``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostGPX`: RouteResponse
    fmt.Fprintf(os.Stdout, "Response from `MapMatchingAPIApi.PostGPX`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostGPXRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gpsAccuracy** | **int32** | Specify the precision of a point, in meter | 
 **vehicle** | **string** | Specify the vehicle profile like car | 

### Return type

[**RouteResponse**](RouteResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

