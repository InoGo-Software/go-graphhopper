# RouteResponsePath

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Distance** | Pointer to **float64** | The total distance, in meters. To get this information for one &#39;leg&#39; please read [this blog post](https://www.graphhopper.com/blog/2019/11/28/routing-api-using-path-details/).  | [optional] 
**Time** | Pointer to **int64** | The total travel time, in milliseconds. To get this information for one &#39;leg&#39; please read [this blog post](https://www.graphhopper.com/blog/2019/11/28/routing-api-using-path-details/).  | [optional] 
**Ascend** | Pointer to **float64** | The total ascent, in meters.  | [optional] 
**Descend** | Pointer to **float64** | The total descent, in meters.  | [optional] 
**Points** | Pointer to **interface{}** |  | [optional] 
**SnappedWaypoints** | Pointer to **interface{}** |  | [optional] 
**PointsEncoded** | Pointer to **bool** | Whether the &#x60;points&#x60; and &#x60;snapped_waypoints&#x60; fields are polyline-encoded strings rather than JSON arrays of coordinates. See the field description for more information on the two formats.  | [optional] 
**Bbox** | Pointer to **[]float64** | The bounding box of the route geometry. Format: &#x60;[minLon, minLat, maxLon, maxLat]&#x60;.  | [optional] 
**Instructions** | Pointer to [**[]RouteResponsePathInstructions**](RouteResponsePathInstructions.md) | The instructions for this route. This feature is under active development, and our instructions can sometimes be misleading, so be mindful when using them for navigation.  | [optional] 
**Details** | Pointer to **map[string]interface{}** | Details, as requested with the &#x60;details&#x60; parameter. Consider the value &#x60;{\&quot;street_name\&quot;: [[0,2,\&quot;Frankfurter Straße\&quot;],[2,6,\&quot;Zollweg\&quot;]]}&#x60;. In this example, the route uses two streets: The first, Frankfurter Straße, is used between &#x60;points[0]&#x60; and &#x60;points[2]&#x60;, and the second, Zollweg, between &#x60;points[2]&#x60; and &#x60;points[6]&#x60;. See [here](https://discuss.graphhopper.com/t/2539) for discussion.  | [optional] 
**PointsOrder** | Pointer to **[]int32** | An array of indices (zero-based), specifiying the order in which the input points are visited. Only present if the &#x60;optimize&#x60; parameter was used.  | [optional] 

## Methods

### NewRouteResponsePath

`func NewRouteResponsePath() *RouteResponsePath`

NewRouteResponsePath instantiates a new RouteResponsePath object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteResponsePathWithDefaults

`func NewRouteResponsePathWithDefaults() *RouteResponsePath`

NewRouteResponsePathWithDefaults instantiates a new RouteResponsePath object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistance

`func (o *RouteResponsePath) GetDistance() float64`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *RouteResponsePath) GetDistanceOk() (*float64, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *RouteResponsePath) SetDistance(v float64)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *RouteResponsePath) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### GetTime

`func (o *RouteResponsePath) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *RouteResponsePath) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *RouteResponsePath) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *RouteResponsePath) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetAscend

`func (o *RouteResponsePath) GetAscend() float64`

GetAscend returns the Ascend field if non-nil, zero value otherwise.

### GetAscendOk

`func (o *RouteResponsePath) GetAscendOk() (*float64, bool)`

GetAscendOk returns a tuple with the Ascend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAscend

`func (o *RouteResponsePath) SetAscend(v float64)`

SetAscend sets Ascend field to given value.

### HasAscend

`func (o *RouteResponsePath) HasAscend() bool`

HasAscend returns a boolean if a field has been set.

### GetDescend

`func (o *RouteResponsePath) GetDescend() float64`

GetDescend returns the Descend field if non-nil, zero value otherwise.

### GetDescendOk

`func (o *RouteResponsePath) GetDescendOk() (*float64, bool)`

GetDescendOk returns a tuple with the Descend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescend

`func (o *RouteResponsePath) SetDescend(v float64)`

SetDescend sets Descend field to given value.

### HasDescend

`func (o *RouteResponsePath) HasDescend() bool`

HasDescend returns a boolean if a field has been set.

### GetPoints

`func (o *RouteResponsePath) GetPoints() interface{}`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *RouteResponsePath) GetPointsOk() (*interface{}, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *RouteResponsePath) SetPoints(v interface{})`

SetPoints sets Points field to given value.

### HasPoints

`func (o *RouteResponsePath) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### GetSnappedWaypoints

`func (o *RouteResponsePath) GetSnappedWaypoints() interface{}`

GetSnappedWaypoints returns the SnappedWaypoints field if non-nil, zero value otherwise.

### GetSnappedWaypointsOk

`func (o *RouteResponsePath) GetSnappedWaypointsOk() (*interface{}, bool)`

GetSnappedWaypointsOk returns a tuple with the SnappedWaypoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnappedWaypoints

`func (o *RouteResponsePath) SetSnappedWaypoints(v interface{})`

SetSnappedWaypoints sets SnappedWaypoints field to given value.

### HasSnappedWaypoints

`func (o *RouteResponsePath) HasSnappedWaypoints() bool`

HasSnappedWaypoints returns a boolean if a field has been set.

### GetPointsEncoded

`func (o *RouteResponsePath) GetPointsEncoded() bool`

GetPointsEncoded returns the PointsEncoded field if non-nil, zero value otherwise.

### GetPointsEncodedOk

`func (o *RouteResponsePath) GetPointsEncodedOk() (*bool, bool)`

GetPointsEncodedOk returns a tuple with the PointsEncoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointsEncoded

`func (o *RouteResponsePath) SetPointsEncoded(v bool)`

SetPointsEncoded sets PointsEncoded field to given value.

### HasPointsEncoded

`func (o *RouteResponsePath) HasPointsEncoded() bool`

HasPointsEncoded returns a boolean if a field has been set.

### GetBbox

`func (o *RouteResponsePath) GetBbox() []float64`

GetBbox returns the Bbox field if non-nil, zero value otherwise.

### GetBboxOk

`func (o *RouteResponsePath) GetBboxOk() (*[]float64, bool)`

GetBboxOk returns a tuple with the Bbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBbox

`func (o *RouteResponsePath) SetBbox(v []float64)`

SetBbox sets Bbox field to given value.

### HasBbox

`func (o *RouteResponsePath) HasBbox() bool`

HasBbox returns a boolean if a field has been set.

### GetInstructions

`func (o *RouteResponsePath) GetInstructions() []RouteResponsePathInstructions`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *RouteResponsePath) GetInstructionsOk() (*[]RouteResponsePathInstructions, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *RouteResponsePath) SetInstructions(v []RouteResponsePathInstructions)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *RouteResponsePath) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetDetails

`func (o *RouteResponsePath) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *RouteResponsePath) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *RouteResponsePath) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *RouteResponsePath) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetPointsOrder

`func (o *RouteResponsePath) GetPointsOrder() []int32`

GetPointsOrder returns the PointsOrder field if non-nil, zero value otherwise.

### GetPointsOrderOk

`func (o *RouteResponsePath) GetPointsOrderOk() (*[]int32, bool)`

GetPointsOrderOk returns a tuple with the PointsOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointsOrder

`func (o *RouteResponsePath) SetPointsOrder(v []int32)`

SetPointsOrder sets PointsOrder field to given value.

### HasPointsOrder

`func (o *RouteResponsePath) HasPointsOrder() bool`

HasPointsOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


