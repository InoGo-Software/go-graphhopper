# \RoutingAPIApi

All URIs are relative to *https://graphhopper.com/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRoute**](RoutingAPIApi.md#GetRoute) | **Get** /route | GET Route Endpoint
[**PostRoute**](RoutingAPIApi.md#PostRoute) | **Post** /route | POST Route Endpoint
[**RouteInfoGet**](RoutingAPIApi.md#RouteInfoGet) | **Get** /route/info | Coverage information



## GetRoute

> RouteResponse GetRoute(ctx).Point(point).PointHint(pointHint).SnapPrevention(snapPrevention).Vehicle(vehicle).Curbside(curbside).TurnCosts(turnCosts).Locale(locale).Elevation(elevation).Details(details).Optimize(optimize).Instructions(instructions).CalcPoints(calcPoints).Debug(debug).PointsEncoded(pointsEncoded).ChDisable(chDisable).Weighting(weighting).Heading(heading).HeadingPenalty(headingPenalty).PassThrough(passThrough).BlockArea(blockArea).Avoid(avoid).Algorithm(algorithm).RoundTripDistance(roundTripDistance).RoundTripSeed(roundTripSeed).AlternativeRouteMaxPaths(alternativeRouteMaxPaths).AlternativeRouteMaxWeightFactor(alternativeRouteMaxWeightFactor).AlternativeRouteMaxShareFactor(alternativeRouteMaxShareFactor).Execute()

GET Route Endpoint



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
    point := []string{"Inner_example"} // []string | The points for which the route should be calculated. Format: `[latitude,longitude]`. Specify at least an origin and a destination. Via points are possible. The maximum number depends on your plan. 
    pointHint := []string{"Inner_example"} // []string | The `point_hint` is typically a road name to which the associated `point` parameter should be snapped to. Specify no `point_hint` parameter or the same number as you have `point` parameters.  (optional)
    snapPrevention := []string{"Inner_example"} // []string | Optional parameter to avoid snapping to a certain road class or road environment. Currently supported values are `motorway`, `trunk`, `ferry`, `tunnel`, `bridge` and `ford`. Multiple values are specified like `snap_prevention=ferry&snap_prevention=motorway`.  (optional)
    vehicle := openapiclient.VehicleProfileId("car") // VehicleProfileId | The vehicle profile for which the route should be calculated.  (optional) (default to "car")
    curbside := []string{"Curbside_example"} // []string | Optional parameter. It specifies on which side a point should be relative to the driver when she leaves/arrives at a start/target/via point. You need to specify this parameter for either none or all points. Only supported for motor vehicles and OpenStreetMap.  (optional)
    turnCosts := true // bool | Specifies if turn restrictions should be considered. Enabling this option increases the route computation time. Only supported for motor vehicles and OpenStreetMap.  (optional) (default to false)
    locale := "locale_example" // string | The locale of the resulting turn instructions. E.g. `pt_PT` for Portuguese or `de` for German.  (optional) (default to "en")
    elevation := true // bool | If `true`, a third coordinate, the altitude, is included with all positions in the response. This changes the format of the `points` and `snapped_waypoints` fields of the response, in both their encodings. Unless you switch off the `points_encoded` parameter, you need special code on the client side that can handle three-dimensional coordinates. A request can fail if the vehicle profile does not support elevation. See the features object for every vehicle profile.  (optional) (default to false)
    details := []string{"Inner_example"} // []string | Optional parameter to retrieve path details. You can request additional details for the route: `street_name`,  `time`, `distance`, `max_speed`, `toll`, `road_class`, `road_class_link`, `road_access`, `road_environment`, `lanes`, and `surface`. Read more about the usage of path details [here](https://discuss.graphhopper.com/t/2539).  (optional)
    optimize := "optimize_example" // string | Normally, the calculated route will visit the points in the order you specified them. If you have more than two points, you can set this parameter to `\"true\"` and the points may be re-ordered to minimize the total travel time. Keep in mind that the limits on the number of locations of the Route Optimization API applies, and the request costs more credits.  (optional) (default to "false")
    instructions := true // bool | If instructions should be calculated and returned  (optional) (default to true)
    calcPoints := true // bool | If the points for the route should be calculated at all.  (optional) (default to true)
    debug := true // bool | If `true`, the output will be formatted.  (optional) (default to false)
    pointsEncoded := true // bool | Allows changing the encoding of location data in the response. The default is polyline encoding, which is compact but requires special client code to unpack. (We provide it in our JavaScript client library!) Set this parameter to `false` to switch the encoding to simple coordinate pairs like `[lon,lat]`, or `[lon,lat,elevation]`. See the description of the response format for more information.  (optional) (default to true)
    chDisable := true // bool | Use this parameter in combination with one or more parameters from below.  (optional) (default to false)
    weighting := "weighting_example" // string | Determines the way the \"best\" route is calculated. Besides `fastest` you can use `short_fastest` which finds a reasonable balance between the distance influence (`shortest`) and the time (`fastest`). You could also use `shortest` but is deprecated and not recommended for motor vehicles. All except `fastest` require `ch.disable=true`.  (optional) (default to "fastest")
    heading := []int32{int32(123)} // []int32 | Favour a heading direction for a certain point. Specify either one heading for the start point or as many as there are points. In this case headings are associated by their order to the specific points. Headings are given as north based clockwise angle between 0 and 360 degree. This parameter also influences the tour generated with `algorithm=round_trip` and forces the initial direction.  Requires `ch.disable=true`.  (optional)
    headingPenalty := int32(56) // int32 | Time penalty in seconds for not obeying a specified heading. Requires `ch.disable=true`.  (optional) (default to 120)
    passThrough := true // bool | If `true`, u-turns are avoided at via-points with regard to the `heading_penalty`. Requires `ch.disable=true`.  (optional) (default to false)
    blockArea := "blockArea_example" // string | Block road access by specifying a point close to the road segment to be blocked, with the format `lat,lon`. You can also block all road segments crossing a geometric shape. Specify a circle using the format `lat,lon,radius`, or a polygon using the format `lat1,lon1,lat2,lon2,...,latN,lonN`. You can specify several shapes, separating them with `;`. Requires `ch.disable=true`.  (optional)
    avoid := "avoid_example" // string | Specify which road classes and environments you would like to avoid.  Possible values are `motorway`, `steps`, `track`, `toll`, `ferry`, `tunnel` and `bridge`. Separate several values with `;`. Obviously not all the values make sense for all vehicle profiles e.g. `bike` is already forbidden on a `motorway`. Requires `ch.disable=true`.  (optional)
    algorithm := "algorithm_example" // string | Rather than looking for the shortest or fastest path, this parameter lets you solve two different problems related to routing: With `alternative_route`, we give you not one but several routes that are close to optimal, but not too similar to each other.  With `round_trip`, the route will get you back to where you started. This is meant for fun (think of a bike trip), so we will add some randomness. The `round_trip` option requires `ch.disable=true`. You can control both of these features with additional parameters, see below.   (optional)
    roundTripDistance := int32(56) // int32 | If `algorithm=round_trip`, this parameter configures approximative length of the resulting round trip. Requires `ch.disable=true`.  (optional) (default to 10000)
    roundTripSeed := int64(789) // int64 | If `algorithm=round_trip`, this sets the random seed. Change this to get a different tour for each value.  (optional)
    alternativeRouteMaxPaths := int32(56) // int32 | If `algorithm=alternative_route`, this parameter sets the number of maximum paths which should be calculated. Increasing can lead to worse alternatives.  (optional) (default to 2)
    alternativeRouteMaxWeightFactor := float32(8.14) // float32 | If `algorithm=alternative_route`, this parameter sets the factor by which the alternatives routes can be longer than the optimal route. Increasing can lead to worse alternatives.  (optional) (default to 1.4)
    alternativeRouteMaxShareFactor := float32(8.14) // float32 | If `algorithm=alternative_route`, this parameter specifies how similar an alternative route can be to the optimal route. Increasing can lead to worse alternatives.  (optional) (default to 0.6)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoutingAPIApi.GetRoute(context.Background()).Point(point).PointHint(pointHint).SnapPrevention(snapPrevention).Vehicle(vehicle).Curbside(curbside).TurnCosts(turnCosts).Locale(locale).Elevation(elevation).Details(details).Optimize(optimize).Instructions(instructions).CalcPoints(calcPoints).Debug(debug).PointsEncoded(pointsEncoded).ChDisable(chDisable).Weighting(weighting).Heading(heading).HeadingPenalty(headingPenalty).PassThrough(passThrough).BlockArea(blockArea).Avoid(avoid).Algorithm(algorithm).RoundTripDistance(roundTripDistance).RoundTripSeed(roundTripSeed).AlternativeRouteMaxPaths(alternativeRouteMaxPaths).AlternativeRouteMaxWeightFactor(alternativeRouteMaxWeightFactor).AlternativeRouteMaxShareFactor(alternativeRouteMaxShareFactor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPIApi.GetRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoute`: RouteResponse
    fmt.Fprintf(os.Stdout, "Response from `RoutingAPIApi.GetRoute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **point** | **[]string** | The points for which the route should be calculated. Format: &#x60;[latitude,longitude]&#x60;. Specify at least an origin and a destination. Via points are possible. The maximum number depends on your plan.  | 
 **pointHint** | **[]string** | The &#x60;point_hint&#x60; is typically a road name to which the associated &#x60;point&#x60; parameter should be snapped to. Specify no &#x60;point_hint&#x60; parameter or the same number as you have &#x60;point&#x60; parameters.  | 
 **snapPrevention** | **[]string** | Optional parameter to avoid snapping to a certain road class or road environment. Currently supported values are &#x60;motorway&#x60;, &#x60;trunk&#x60;, &#x60;ferry&#x60;, &#x60;tunnel&#x60;, &#x60;bridge&#x60; and &#x60;ford&#x60;. Multiple values are specified like &#x60;snap_prevention&#x3D;ferry&amp;snap_prevention&#x3D;motorway&#x60;.  | 
 **vehicle** | [**VehicleProfileId**](VehicleProfileId.md) | The vehicle profile for which the route should be calculated.  | [default to &quot;car&quot;]
 **curbside** | **[]string** | Optional parameter. It specifies on which side a point should be relative to the driver when she leaves/arrives at a start/target/via point. You need to specify this parameter for either none or all points. Only supported for motor vehicles and OpenStreetMap.  | 
 **turnCosts** | **bool** | Specifies if turn restrictions should be considered. Enabling this option increases the route computation time. Only supported for motor vehicles and OpenStreetMap.  | [default to false]
 **locale** | **string** | The locale of the resulting turn instructions. E.g. &#x60;pt_PT&#x60; for Portuguese or &#x60;de&#x60; for German.  | [default to &quot;en&quot;]
 **elevation** | **bool** | If &#x60;true&#x60;, a third coordinate, the altitude, is included with all positions in the response. This changes the format of the &#x60;points&#x60; and &#x60;snapped_waypoints&#x60; fields of the response, in both their encodings. Unless you switch off the &#x60;points_encoded&#x60; parameter, you need special code on the client side that can handle three-dimensional coordinates. A request can fail if the vehicle profile does not support elevation. See the features object for every vehicle profile.  | [default to false]
 **details** | **[]string** | Optional parameter to retrieve path details. You can request additional details for the route: &#x60;street_name&#x60;,  &#x60;time&#x60;, &#x60;distance&#x60;, &#x60;max_speed&#x60;, &#x60;toll&#x60;, &#x60;road_class&#x60;, &#x60;road_class_link&#x60;, &#x60;road_access&#x60;, &#x60;road_environment&#x60;, &#x60;lanes&#x60;, and &#x60;surface&#x60;. Read more about the usage of path details [here](https://discuss.graphhopper.com/t/2539).  | 
 **optimize** | **string** | Normally, the calculated route will visit the points in the order you specified them. If you have more than two points, you can set this parameter to &#x60;\&quot;true\&quot;&#x60; and the points may be re-ordered to minimize the total travel time. Keep in mind that the limits on the number of locations of the Route Optimization API applies, and the request costs more credits.  | [default to &quot;false&quot;]
 **instructions** | **bool** | If instructions should be calculated and returned  | [default to true]
 **calcPoints** | **bool** | If the points for the route should be calculated at all.  | [default to true]
 **debug** | **bool** | If &#x60;true&#x60;, the output will be formatted.  | [default to false]
 **pointsEncoded** | **bool** | Allows changing the encoding of location data in the response. The default is polyline encoding, which is compact but requires special client code to unpack. (We provide it in our JavaScript client library!) Set this parameter to &#x60;false&#x60; to switch the encoding to simple coordinate pairs like &#x60;[lon,lat]&#x60;, or &#x60;[lon,lat,elevation]&#x60;. See the description of the response format for more information.  | [default to true]
 **chDisable** | **bool** | Use this parameter in combination with one or more parameters from below.  | [default to false]
 **weighting** | **string** | Determines the way the \&quot;best\&quot; route is calculated. Besides &#x60;fastest&#x60; you can use &#x60;short_fastest&#x60; which finds a reasonable balance between the distance influence (&#x60;shortest&#x60;) and the time (&#x60;fastest&#x60;). You could also use &#x60;shortest&#x60; but is deprecated and not recommended for motor vehicles. All except &#x60;fastest&#x60; require &#x60;ch.disable&#x3D;true&#x60;.  | [default to &quot;fastest&quot;]
 **heading** | **[]int32** | Favour a heading direction for a certain point. Specify either one heading for the start point or as many as there are points. In this case headings are associated by their order to the specific points. Headings are given as north based clockwise angle between 0 and 360 degree. This parameter also influences the tour generated with &#x60;algorithm&#x3D;round_trip&#x60; and forces the initial direction.  Requires &#x60;ch.disable&#x3D;true&#x60;.  | 
 **headingPenalty** | **int32** | Time penalty in seconds for not obeying a specified heading. Requires &#x60;ch.disable&#x3D;true&#x60;.  | [default to 120]
 **passThrough** | **bool** | If &#x60;true&#x60;, u-turns are avoided at via-points with regard to the &#x60;heading_penalty&#x60;. Requires &#x60;ch.disable&#x3D;true&#x60;.  | [default to false]
 **blockArea** | **string** | Block road access by specifying a point close to the road segment to be blocked, with the format &#x60;lat,lon&#x60;. You can also block all road segments crossing a geometric shape. Specify a circle using the format &#x60;lat,lon,radius&#x60;, or a polygon using the format &#x60;lat1,lon1,lat2,lon2,...,latN,lonN&#x60;. You can specify several shapes, separating them with &#x60;;&#x60;. Requires &#x60;ch.disable&#x3D;true&#x60;.  | 
 **avoid** | **string** | Specify which road classes and environments you would like to avoid.  Possible values are &#x60;motorway&#x60;, &#x60;steps&#x60;, &#x60;track&#x60;, &#x60;toll&#x60;, &#x60;ferry&#x60;, &#x60;tunnel&#x60; and &#x60;bridge&#x60;. Separate several values with &#x60;;&#x60;. Obviously not all the values make sense for all vehicle profiles e.g. &#x60;bike&#x60; is already forbidden on a &#x60;motorway&#x60;. Requires &#x60;ch.disable&#x3D;true&#x60;.  | 
 **algorithm** | **string** | Rather than looking for the shortest or fastest path, this parameter lets you solve two different problems related to routing: With &#x60;alternative_route&#x60;, we give you not one but several routes that are close to optimal, but not too similar to each other.  With &#x60;round_trip&#x60;, the route will get you back to where you started. This is meant for fun (think of a bike trip), so we will add some randomness. The &#x60;round_trip&#x60; option requires &#x60;ch.disable&#x3D;true&#x60;. You can control both of these features with additional parameters, see below.   | 
 **roundTripDistance** | **int32** | If &#x60;algorithm&#x3D;round_trip&#x60;, this parameter configures approximative length of the resulting round trip. Requires &#x60;ch.disable&#x3D;true&#x60;.  | [default to 10000]
 **roundTripSeed** | **int64** | If &#x60;algorithm&#x3D;round_trip&#x60;, this sets the random seed. Change this to get a different tour for each value.  | 
 **alternativeRouteMaxPaths** | **int32** | If &#x60;algorithm&#x3D;alternative_route&#x60;, this parameter sets the number of maximum paths which should be calculated. Increasing can lead to worse alternatives.  | [default to 2]
 **alternativeRouteMaxWeightFactor** | **float32** | If &#x60;algorithm&#x3D;alternative_route&#x60;, this parameter sets the factor by which the alternatives routes can be longer than the optimal route. Increasing can lead to worse alternatives.  | [default to 1.4]
 **alternativeRouteMaxShareFactor** | **float32** | If &#x60;algorithm&#x3D;alternative_route&#x60;, this parameter specifies how similar an alternative route can be to the optimal route. Increasing can lead to worse alternatives.  | [default to 0.6]

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


## PostRoute

> RouteResponse PostRoute(ctx).RouteRequest(routeRequest).Execute()

POST Route Endpoint



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
    routeRequest := *openapiclient.NewRouteRequest() // RouteRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoutingAPIApi.PostRoute(context.Background()).RouteRequest(routeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPIApi.PostRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostRoute`: RouteResponse
    fmt.Fprintf(os.Stdout, "Response from `RoutingAPIApi.PostRoute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **routeRequest** | [**RouteRequest**](RouteRequest.md) |  | 

### Return type

[**RouteResponse**](RouteResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RouteInfoGet

> InfoResponse RouteInfoGet(ctx).Execute()

Coverage information



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
    resp, r, err := api_client.RoutingAPIApi.RouteInfoGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPIApi.RouteInfoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RouteInfoGet`: InfoResponse
    fmt.Fprintf(os.Stdout, "Response from `RoutingAPIApi.RouteInfoGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRouteInfoGetRequest struct via the builder pattern


### Return type

[**InfoResponse**](InfoResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

