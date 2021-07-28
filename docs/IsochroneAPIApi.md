# \IsochroneAPIApi

All URIs are relative to *https://graphhopper.com/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIsochrone**](IsochroneAPIApi.md#GetIsochrone) | **Get** /isochrone | Isochrone Endpoint



## GetIsochrone

> IsochroneResponse GetIsochrone(ctx).Point(point).TimeLimit(timeLimit).DistanceLimit(distanceLimit).Vehicle(vehicle).Buckets(buckets).ReverseFlow(reverseFlow).Weighting(weighting).Execute()

Isochrone Endpoint



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
    point := "point_example" // string | Specify the start coordinate
    timeLimit := int32(56) // int32 | Specify which time the vehicle should travel. In seconds. (optional) (default to 600)
    distanceLimit := int32(56) // int32 | Specify which distance the vehicle should travel. In meters. (optional)
    vehicle := openapiclient.VehicleProfileId("car") // VehicleProfileId | The vehicle profile for which the route should be calculated.  (optional) (default to "car")
    buckets := int32(56) // int32 | Number by which to divide the given `time_limit` to create `buckets` nested isochrones of time intervals `time_limit-n*time_limit/buckets`. Applies analogously to `distance_limit`. (optional) (default to 1)
    reverseFlow := true // bool | If `false` the flow goes from point to the polygon, if `true` the flow goes from the polygon \"inside\" to the point. Example use case for `false`&#58; *How many potential customer can be reached within 30min travel time from your store* vs. `true`&#58; *How many customers can reach your store within 30min travel time.*  (optional) (default to false)
    weighting := "weighting_example" // string | Use `\"shortest\"` to get an isodistance line instead of an isochrone. (optional) (default to "fastest")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IsochroneAPIApi.GetIsochrone(context.Background()).Point(point).TimeLimit(timeLimit).DistanceLimit(distanceLimit).Vehicle(vehicle).Buckets(buckets).ReverseFlow(reverseFlow).Weighting(weighting).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IsochroneAPIApi.GetIsochrone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIsochrone`: IsochroneResponse
    fmt.Fprintf(os.Stdout, "Response from `IsochroneAPIApi.GetIsochrone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIsochroneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **point** | **string** | Specify the start coordinate | 
 **timeLimit** | **int32** | Specify which time the vehicle should travel. In seconds. | [default to 600]
 **distanceLimit** | **int32** | Specify which distance the vehicle should travel. In meters. | 
 **vehicle** | [**VehicleProfileId**](VehicleProfileId.md) | The vehicle profile for which the route should be calculated.  | [default to &quot;car&quot;]
 **buckets** | **int32** | Number by which to divide the given &#x60;time_limit&#x60; to create &#x60;buckets&#x60; nested isochrones of time intervals &#x60;time_limit-n*time_limit/buckets&#x60;. Applies analogously to &#x60;distance_limit&#x60;. | [default to 1]
 **reverseFlow** | **bool** | If &#x60;false&#x60; the flow goes from point to the polygon, if &#x60;true&#x60; the flow goes from the polygon \&quot;inside\&quot; to the point. Example use case for &#x60;false&#x60;&amp;#58; *How many potential customer can be reached within 30min travel time from your store* vs. &#x60;true&#x60;&amp;#58; *How many customers can reach your store within 30min travel time.*  | [default to false]
 **weighting** | **string** | Use &#x60;\&quot;shortest\&quot;&#x60; to get an isodistance line instead of an isochrone. | [default to &quot;fastest&quot;]

### Return type

[**IsochroneResponse**](IsochroneResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

