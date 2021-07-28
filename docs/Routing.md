# Routing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CalcPoints** | Pointer to **bool** | It lets you specify whether the API should provide you with route geometries for vehicle routes or not. Thus, you do not need to do extra routing to get the polyline for each route. | [optional] [default to false]
**ConsiderTraffic** | Pointer to **bool** | indicates whether historical traffic information should be considered | [optional] [default to false]
**NetworkDataProvider** | Pointer to **string** | specifies the data provider, read more about it [here](#section/Map-Data-and-Routing-Profiles). | [optional] [default to "openstreetmap"]
**CurbsideStrictness** | Pointer to **string** | In some cases curbside constraints cannot be fulfilled. For example in one-way streets you cannot arrive at a building that is on the left side of the street such that the building is to the right of you (unless you drove the one-way street the wrong/illegal way). You can set the &#x60;curbside_strictness&#x60; to &#x60;soft&#x60; to ignore the curbside constraint in such cases or set it to &#x60;strict&#x60; to get an error response instead. You can also set it to &#x60;ignore&#x60; to ignore all curbside constraints (this is useful to compare the results with and without constraints without modifying every single address). | [optional] [default to "soft"]
**FailFast** | Pointer to **bool** | indicates whether matrix calculation should fail fast when points cannot be connected | [optional] [default to true]
**ReturnSnappedWaypoints** | Pointer to **bool** | Indicates whether a solution includes snapped waypoints. In contrary to the address coordinate a snapped waypoint is the access point to the (road) network. | [optional] [default to false]
**SnapPreventions** | Pointer to **[]string** | Prevents snapping locations to road links of specified road types, e.g. to motorway. | [optional] 

## Methods

### NewRouting

`func NewRouting() *Routing`

NewRouting instantiates a new Routing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingWithDefaults

`func NewRoutingWithDefaults() *Routing`

NewRoutingWithDefaults instantiates a new Routing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalcPoints

`func (o *Routing) GetCalcPoints() bool`

GetCalcPoints returns the CalcPoints field if non-nil, zero value otherwise.

### GetCalcPointsOk

`func (o *Routing) GetCalcPointsOk() (*bool, bool)`

GetCalcPointsOk returns a tuple with the CalcPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalcPoints

`func (o *Routing) SetCalcPoints(v bool)`

SetCalcPoints sets CalcPoints field to given value.

### HasCalcPoints

`func (o *Routing) HasCalcPoints() bool`

HasCalcPoints returns a boolean if a field has been set.

### GetConsiderTraffic

`func (o *Routing) GetConsiderTraffic() bool`

GetConsiderTraffic returns the ConsiderTraffic field if non-nil, zero value otherwise.

### GetConsiderTrafficOk

`func (o *Routing) GetConsiderTrafficOk() (*bool, bool)`

GetConsiderTrafficOk returns a tuple with the ConsiderTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsiderTraffic

`func (o *Routing) SetConsiderTraffic(v bool)`

SetConsiderTraffic sets ConsiderTraffic field to given value.

### HasConsiderTraffic

`func (o *Routing) HasConsiderTraffic() bool`

HasConsiderTraffic returns a boolean if a field has been set.

### GetNetworkDataProvider

`func (o *Routing) GetNetworkDataProvider() string`

GetNetworkDataProvider returns the NetworkDataProvider field if non-nil, zero value otherwise.

### GetNetworkDataProviderOk

`func (o *Routing) GetNetworkDataProviderOk() (*string, bool)`

GetNetworkDataProviderOk returns a tuple with the NetworkDataProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDataProvider

`func (o *Routing) SetNetworkDataProvider(v string)`

SetNetworkDataProvider sets NetworkDataProvider field to given value.

### HasNetworkDataProvider

`func (o *Routing) HasNetworkDataProvider() bool`

HasNetworkDataProvider returns a boolean if a field has been set.

### GetCurbsideStrictness

`func (o *Routing) GetCurbsideStrictness() string`

GetCurbsideStrictness returns the CurbsideStrictness field if non-nil, zero value otherwise.

### GetCurbsideStrictnessOk

`func (o *Routing) GetCurbsideStrictnessOk() (*string, bool)`

GetCurbsideStrictnessOk returns a tuple with the CurbsideStrictness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurbsideStrictness

`func (o *Routing) SetCurbsideStrictness(v string)`

SetCurbsideStrictness sets CurbsideStrictness field to given value.

### HasCurbsideStrictness

`func (o *Routing) HasCurbsideStrictness() bool`

HasCurbsideStrictness returns a boolean if a field has been set.

### GetFailFast

`func (o *Routing) GetFailFast() bool`

GetFailFast returns the FailFast field if non-nil, zero value otherwise.

### GetFailFastOk

`func (o *Routing) GetFailFastOk() (*bool, bool)`

GetFailFastOk returns a tuple with the FailFast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailFast

`func (o *Routing) SetFailFast(v bool)`

SetFailFast sets FailFast field to given value.

### HasFailFast

`func (o *Routing) HasFailFast() bool`

HasFailFast returns a boolean if a field has been set.

### GetReturnSnappedWaypoints

`func (o *Routing) GetReturnSnappedWaypoints() bool`

GetReturnSnappedWaypoints returns the ReturnSnappedWaypoints field if non-nil, zero value otherwise.

### GetReturnSnappedWaypointsOk

`func (o *Routing) GetReturnSnappedWaypointsOk() (*bool, bool)`

GetReturnSnappedWaypointsOk returns a tuple with the ReturnSnappedWaypoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnSnappedWaypoints

`func (o *Routing) SetReturnSnappedWaypoints(v bool)`

SetReturnSnappedWaypoints sets ReturnSnappedWaypoints field to given value.

### HasReturnSnappedWaypoints

`func (o *Routing) HasReturnSnappedWaypoints() bool`

HasReturnSnappedWaypoints returns a boolean if a field has been set.

### GetSnapPreventions

`func (o *Routing) GetSnapPreventions() []string`

GetSnapPreventions returns the SnapPreventions field if non-nil, zero value otherwise.

### GetSnapPreventionsOk

`func (o *Routing) GetSnapPreventionsOk() (*[]string, bool)`

GetSnapPreventionsOk returns a tuple with the SnapPreventions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapPreventions

`func (o *Routing) SetSnapPreventions(v []string)`

SetSnapPreventions sets SnapPreventions field to given value.

### HasSnapPreventions

`func (o *Routing) HasSnapPreventions() bool`

HasSnapPreventions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


