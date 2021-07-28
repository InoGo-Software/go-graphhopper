# Route

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VehicleId** | Pointer to **string** | Id of vehicle that operates route | [optional] 
**Distance** | Pointer to **int64** | Distance of route in meter | [optional] 
**TransportTime** | Pointer to **int64** | Transport time of route in seconds | [optional] 
**CompletionTime** | Pointer to **int64** | Completion time of route in seconds | [optional] 
**WaitingTime** | Pointer to **int64** | Waiting time of route in seconds | [optional] 
**ServiceDuration** | Pointer to **int64** | Service duration of route in seconds | [optional] 
**PreparationTime** | Pointer to **int64** | Preparation time of route in seconds | [optional] 
**Activities** | Pointer to [**[]Activity**](Activity.md) | Array of activities | [optional] 
**Points** | Pointer to [**[]RoutePoint**](RoutePoint.md) | Array of route planning points | [optional] 

## Methods

### NewRoute

`func NewRoute() *Route`

NewRoute instantiates a new Route object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteWithDefaults

`func NewRouteWithDefaults() *Route`

NewRouteWithDefaults instantiates a new Route object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVehicleId

`func (o *Route) GetVehicleId() string`

GetVehicleId returns the VehicleId field if non-nil, zero value otherwise.

### GetVehicleIdOk

`func (o *Route) GetVehicleIdOk() (*string, bool)`

GetVehicleIdOk returns a tuple with the VehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleId

`func (o *Route) SetVehicleId(v string)`

SetVehicleId sets VehicleId field to given value.

### HasVehicleId

`func (o *Route) HasVehicleId() bool`

HasVehicleId returns a boolean if a field has been set.

### GetDistance

`func (o *Route) GetDistance() int64`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *Route) GetDistanceOk() (*int64, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *Route) SetDistance(v int64)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *Route) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### GetTransportTime

`func (o *Route) GetTransportTime() int64`

GetTransportTime returns the TransportTime field if non-nil, zero value otherwise.

### GetTransportTimeOk

`func (o *Route) GetTransportTimeOk() (*int64, bool)`

GetTransportTimeOk returns a tuple with the TransportTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportTime

`func (o *Route) SetTransportTime(v int64)`

SetTransportTime sets TransportTime field to given value.

### HasTransportTime

`func (o *Route) HasTransportTime() bool`

HasTransportTime returns a boolean if a field has been set.

### GetCompletionTime

`func (o *Route) GetCompletionTime() int64`

GetCompletionTime returns the CompletionTime field if non-nil, zero value otherwise.

### GetCompletionTimeOk

`func (o *Route) GetCompletionTimeOk() (*int64, bool)`

GetCompletionTimeOk returns a tuple with the CompletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTime

`func (o *Route) SetCompletionTime(v int64)`

SetCompletionTime sets CompletionTime field to given value.

### HasCompletionTime

`func (o *Route) HasCompletionTime() bool`

HasCompletionTime returns a boolean if a field has been set.

### GetWaitingTime

`func (o *Route) GetWaitingTime() int64`

GetWaitingTime returns the WaitingTime field if non-nil, zero value otherwise.

### GetWaitingTimeOk

`func (o *Route) GetWaitingTimeOk() (*int64, bool)`

GetWaitingTimeOk returns a tuple with the WaitingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitingTime

`func (o *Route) SetWaitingTime(v int64)`

SetWaitingTime sets WaitingTime field to given value.

### HasWaitingTime

`func (o *Route) HasWaitingTime() bool`

HasWaitingTime returns a boolean if a field has been set.

### GetServiceDuration

`func (o *Route) GetServiceDuration() int64`

GetServiceDuration returns the ServiceDuration field if non-nil, zero value otherwise.

### GetServiceDurationOk

`func (o *Route) GetServiceDurationOk() (*int64, bool)`

GetServiceDurationOk returns a tuple with the ServiceDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDuration

`func (o *Route) SetServiceDuration(v int64)`

SetServiceDuration sets ServiceDuration field to given value.

### HasServiceDuration

`func (o *Route) HasServiceDuration() bool`

HasServiceDuration returns a boolean if a field has been set.

### GetPreparationTime

`func (o *Route) GetPreparationTime() int64`

GetPreparationTime returns the PreparationTime field if non-nil, zero value otherwise.

### GetPreparationTimeOk

`func (o *Route) GetPreparationTimeOk() (*int64, bool)`

GetPreparationTimeOk returns a tuple with the PreparationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparationTime

`func (o *Route) SetPreparationTime(v int64)`

SetPreparationTime sets PreparationTime field to given value.

### HasPreparationTime

`func (o *Route) HasPreparationTime() bool`

HasPreparationTime returns a boolean if a field has been set.

### GetActivities

`func (o *Route) GetActivities() []Activity`

GetActivities returns the Activities field if non-nil, zero value otherwise.

### GetActivitiesOk

`func (o *Route) GetActivitiesOk() (*[]Activity, bool)`

GetActivitiesOk returns a tuple with the Activities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivities

`func (o *Route) SetActivities(v []Activity)`

SetActivities sets Activities field to given value.

### HasActivities

`func (o *Route) HasActivities() bool`

HasActivities returns a boolean if a field has been set.

### GetPoints

`func (o *Route) GetPoints() []RoutePoint`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *Route) GetPointsOk() (*[]RoutePoint, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *Route) SetPoints(v []RoutePoint)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *Route) HasPoints() bool`

HasPoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


