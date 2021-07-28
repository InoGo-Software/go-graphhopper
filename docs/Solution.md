# Solution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Costs** | Pointer to **int32** |  | [optional] 
**Distance** | Pointer to **int32** | Overall distance travelled in meter, i.e. the sum of each route&#39;s transport distance | [optional] 
**Time** | Pointer to **int64** | Use &#x60;transport_time&#x60; instead. | [optional] 
**TransportTime** | Pointer to **int64** | Overall time travelled in seconds, i.e. the sum of each route&#39;s transport time. | [optional] 
**MaxOperationTime** | Pointer to **int64** | Operation time of longest route in seconds. | [optional] 
**WaitingTime** | Pointer to **int64** | Overall waiting time in seconds. | [optional] 
**ServiceDuration** | Pointer to **int64** | Overall service time in seconds. | [optional] 
**PreparationTime** | Pointer to **int64** | Overall preparation time in seconds. | [optional] 
**CompletionTime** | Pointer to **int64** | Overall completion time in seconds, i.e. the sum of each routes/drivers operation time. | [optional] 
**NoVehicles** | Pointer to **int32** | Number of employed vehicles. | [optional] 
**NoUnassigned** | Pointer to **int32** | Number of jobs that could not be assigned to final solution. | [optional] 
**Routes** | Pointer to [**[]Route**](Route.md) | An array of routes | [optional] 
**Unassigned** | Pointer to [**SolutionUnassigned**](SolutionUnassigned.md) |  | [optional] 

## Methods

### NewSolution

`func NewSolution() *Solution`

NewSolution instantiates a new Solution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSolutionWithDefaults

`func NewSolutionWithDefaults() *Solution`

NewSolutionWithDefaults instantiates a new Solution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCosts

`func (o *Solution) GetCosts() int32`

GetCosts returns the Costs field if non-nil, zero value otherwise.

### GetCostsOk

`func (o *Solution) GetCostsOk() (*int32, bool)`

GetCostsOk returns a tuple with the Costs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosts

`func (o *Solution) SetCosts(v int32)`

SetCosts sets Costs field to given value.

### HasCosts

`func (o *Solution) HasCosts() bool`

HasCosts returns a boolean if a field has been set.

### GetDistance

`func (o *Solution) GetDistance() int32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *Solution) GetDistanceOk() (*int32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *Solution) SetDistance(v int32)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *Solution) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### GetTime

`func (o *Solution) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *Solution) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *Solution) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *Solution) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTransportTime

`func (o *Solution) GetTransportTime() int64`

GetTransportTime returns the TransportTime field if non-nil, zero value otherwise.

### GetTransportTimeOk

`func (o *Solution) GetTransportTimeOk() (*int64, bool)`

GetTransportTimeOk returns a tuple with the TransportTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportTime

`func (o *Solution) SetTransportTime(v int64)`

SetTransportTime sets TransportTime field to given value.

### HasTransportTime

`func (o *Solution) HasTransportTime() bool`

HasTransportTime returns a boolean if a field has been set.

### GetMaxOperationTime

`func (o *Solution) GetMaxOperationTime() int64`

GetMaxOperationTime returns the MaxOperationTime field if non-nil, zero value otherwise.

### GetMaxOperationTimeOk

`func (o *Solution) GetMaxOperationTimeOk() (*int64, bool)`

GetMaxOperationTimeOk returns a tuple with the MaxOperationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOperationTime

`func (o *Solution) SetMaxOperationTime(v int64)`

SetMaxOperationTime sets MaxOperationTime field to given value.

### HasMaxOperationTime

`func (o *Solution) HasMaxOperationTime() bool`

HasMaxOperationTime returns a boolean if a field has been set.

### GetWaitingTime

`func (o *Solution) GetWaitingTime() int64`

GetWaitingTime returns the WaitingTime field if non-nil, zero value otherwise.

### GetWaitingTimeOk

`func (o *Solution) GetWaitingTimeOk() (*int64, bool)`

GetWaitingTimeOk returns a tuple with the WaitingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitingTime

`func (o *Solution) SetWaitingTime(v int64)`

SetWaitingTime sets WaitingTime field to given value.

### HasWaitingTime

`func (o *Solution) HasWaitingTime() bool`

HasWaitingTime returns a boolean if a field has been set.

### GetServiceDuration

`func (o *Solution) GetServiceDuration() int64`

GetServiceDuration returns the ServiceDuration field if non-nil, zero value otherwise.

### GetServiceDurationOk

`func (o *Solution) GetServiceDurationOk() (*int64, bool)`

GetServiceDurationOk returns a tuple with the ServiceDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDuration

`func (o *Solution) SetServiceDuration(v int64)`

SetServiceDuration sets ServiceDuration field to given value.

### HasServiceDuration

`func (o *Solution) HasServiceDuration() bool`

HasServiceDuration returns a boolean if a field has been set.

### GetPreparationTime

`func (o *Solution) GetPreparationTime() int64`

GetPreparationTime returns the PreparationTime field if non-nil, zero value otherwise.

### GetPreparationTimeOk

`func (o *Solution) GetPreparationTimeOk() (*int64, bool)`

GetPreparationTimeOk returns a tuple with the PreparationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparationTime

`func (o *Solution) SetPreparationTime(v int64)`

SetPreparationTime sets PreparationTime field to given value.

### HasPreparationTime

`func (o *Solution) HasPreparationTime() bool`

HasPreparationTime returns a boolean if a field has been set.

### GetCompletionTime

`func (o *Solution) GetCompletionTime() int64`

GetCompletionTime returns the CompletionTime field if non-nil, zero value otherwise.

### GetCompletionTimeOk

`func (o *Solution) GetCompletionTimeOk() (*int64, bool)`

GetCompletionTimeOk returns a tuple with the CompletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTime

`func (o *Solution) SetCompletionTime(v int64)`

SetCompletionTime sets CompletionTime field to given value.

### HasCompletionTime

`func (o *Solution) HasCompletionTime() bool`

HasCompletionTime returns a boolean if a field has been set.

### GetNoVehicles

`func (o *Solution) GetNoVehicles() int32`

GetNoVehicles returns the NoVehicles field if non-nil, zero value otherwise.

### GetNoVehiclesOk

`func (o *Solution) GetNoVehiclesOk() (*int32, bool)`

GetNoVehiclesOk returns a tuple with the NoVehicles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoVehicles

`func (o *Solution) SetNoVehicles(v int32)`

SetNoVehicles sets NoVehicles field to given value.

### HasNoVehicles

`func (o *Solution) HasNoVehicles() bool`

HasNoVehicles returns a boolean if a field has been set.

### GetNoUnassigned

`func (o *Solution) GetNoUnassigned() int32`

GetNoUnassigned returns the NoUnassigned field if non-nil, zero value otherwise.

### GetNoUnassignedOk

`func (o *Solution) GetNoUnassignedOk() (*int32, bool)`

GetNoUnassignedOk returns a tuple with the NoUnassigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoUnassigned

`func (o *Solution) SetNoUnassigned(v int32)`

SetNoUnassigned sets NoUnassigned field to given value.

### HasNoUnassigned

`func (o *Solution) HasNoUnassigned() bool`

HasNoUnassigned returns a boolean if a field has been set.

### GetRoutes

`func (o *Solution) GetRoutes() []Route`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *Solution) GetRoutesOk() (*[]Route, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *Solution) SetRoutes(v []Route)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *Solution) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetUnassigned

`func (o *Solution) GetUnassigned() SolutionUnassigned`

GetUnassigned returns the Unassigned field if non-nil, zero value otherwise.

### GetUnassignedOk

`func (o *Solution) GetUnassignedOk() (*SolutionUnassigned, bool)`

GetUnassignedOk returns a tuple with the Unassigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnassigned

`func (o *Solution) SetUnassigned(v SolutionUnassigned)`

SetUnassigned sets Unassigned field to given value.

### HasUnassigned

`func (o *Solution) HasUnassigned() bool`

HasUnassigned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


