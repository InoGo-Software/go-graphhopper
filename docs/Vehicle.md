# Vehicle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VehicleId** | **string** | Specifies the ID of the vehicle. Ids must be unique, i.e. if there are two vehicles with the same ID, an error is returned. | 
**TypeId** | Pointer to **string** | The type ID assigns a vehicle type to this vehicle. You can specify types in the array of vehicle types. If you omit the type ID, the default type is used. The default type is a &#x60;car&#x60; with a capacity of 0. | [optional] [default to "default-type"]
**StartAddress** | [**Address**](Address.md) |  | 
**EndAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**Break** | Pointer to [**AnyOfTimeWindowBreakDriveTimeBreak**](anyOf&lt;TimeWindowBreak,DriveTimeBreak&gt;.md) |  | [optional] 
**ReturnToDepot** | Pointer to **bool** | If it is false, the algorithm decides where to end the vehicle route. It ends in one of your customers&#39; locations. The end is chosen such that it contributes to the overall objective function, e.g. min transport_time. If it is true, you can either specify a specific end location (which is then regarded as end depot) or you can leave it and the driver returns to its start location. | [optional] [default to true]
**EarliestStart** | Pointer to **int64** | Earliest start of vehicle in seconds. It is recommended to use the unix timestamp. | [optional] [default to 0]
**LatestEnd** | Pointer to **int64** | Latest end of vehicle in seconds, i.e. the time the vehicle needs to be at its end location at latest. | [optional] 
**Skills** | Pointer to **[]string** | Array of skills, i.e. array of string (not case sensitive). | [optional] 
**MaxDistance** | Pointer to **int64** | Specifies the maximum distance (in meters) a vehicle can go. | [optional] 
**MaxDrivingTime** | Pointer to **int64** | Specifies the maximum drive time (in seconds) a vehicle/driver can go, i.e. the maximum time on the road (service and waiting times are not included here) | [optional] 
**MaxJobs** | Pointer to **int32** | Specifies the maximum number of jobs a vehicle can load. | [optional] 
**MinJobs** | Pointer to **int32** | Specifies the minimum number of jobs a vehicle should load. This is a soft constraint, i.e. if it is not possible to fulfill “min_jobs”, we will still try to get as close as possible to this constraint. | [optional] 
**MaxActivities** | Pointer to **int32** | Specifies the maximum number of activities a vehicle can conduct. | [optional] 
**MoveToEndAddress** | Pointer to **bool** | Indicates whether a vehicle should be moved even though it has not been assigned any jobs. | [optional] 

## Methods

### NewVehicle

`func NewVehicle(vehicleId string, startAddress Address, ) *Vehicle`

NewVehicle instantiates a new Vehicle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleWithDefaults

`func NewVehicleWithDefaults() *Vehicle`

NewVehicleWithDefaults instantiates a new Vehicle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVehicleId

`func (o *Vehicle) GetVehicleId() string`

GetVehicleId returns the VehicleId field if non-nil, zero value otherwise.

### GetVehicleIdOk

`func (o *Vehicle) GetVehicleIdOk() (*string, bool)`

GetVehicleIdOk returns a tuple with the VehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleId

`func (o *Vehicle) SetVehicleId(v string)`

SetVehicleId sets VehicleId field to given value.


### GetTypeId

`func (o *Vehicle) GetTypeId() string`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *Vehicle) GetTypeIdOk() (*string, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *Vehicle) SetTypeId(v string)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *Vehicle) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### GetStartAddress

`func (o *Vehicle) GetStartAddress() Address`

GetStartAddress returns the StartAddress field if non-nil, zero value otherwise.

### GetStartAddressOk

`func (o *Vehicle) GetStartAddressOk() (*Address, bool)`

GetStartAddressOk returns a tuple with the StartAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAddress

`func (o *Vehicle) SetStartAddress(v Address)`

SetStartAddress sets StartAddress field to given value.


### GetEndAddress

`func (o *Vehicle) GetEndAddress() Address`

GetEndAddress returns the EndAddress field if non-nil, zero value otherwise.

### GetEndAddressOk

`func (o *Vehicle) GetEndAddressOk() (*Address, bool)`

GetEndAddressOk returns a tuple with the EndAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAddress

`func (o *Vehicle) SetEndAddress(v Address)`

SetEndAddress sets EndAddress field to given value.

### HasEndAddress

`func (o *Vehicle) HasEndAddress() bool`

HasEndAddress returns a boolean if a field has been set.

### GetBreak

`func (o *Vehicle) GetBreak() AnyOfTimeWindowBreakDriveTimeBreak`

GetBreak returns the Break field if non-nil, zero value otherwise.

### GetBreakOk

`func (o *Vehicle) GetBreakOk() (*AnyOfTimeWindowBreakDriveTimeBreak, bool)`

GetBreakOk returns a tuple with the Break field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreak

`func (o *Vehicle) SetBreak(v AnyOfTimeWindowBreakDriveTimeBreak)`

SetBreak sets Break field to given value.

### HasBreak

`func (o *Vehicle) HasBreak() bool`

HasBreak returns a boolean if a field has been set.

### GetReturnToDepot

`func (o *Vehicle) GetReturnToDepot() bool`

GetReturnToDepot returns the ReturnToDepot field if non-nil, zero value otherwise.

### GetReturnToDepotOk

`func (o *Vehicle) GetReturnToDepotOk() (*bool, bool)`

GetReturnToDepotOk returns a tuple with the ReturnToDepot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnToDepot

`func (o *Vehicle) SetReturnToDepot(v bool)`

SetReturnToDepot sets ReturnToDepot field to given value.

### HasReturnToDepot

`func (o *Vehicle) HasReturnToDepot() bool`

HasReturnToDepot returns a boolean if a field has been set.

### GetEarliestStart

`func (o *Vehicle) GetEarliestStart() int64`

GetEarliestStart returns the EarliestStart field if non-nil, zero value otherwise.

### GetEarliestStartOk

`func (o *Vehicle) GetEarliestStartOk() (*int64, bool)`

GetEarliestStartOk returns a tuple with the EarliestStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestStart

`func (o *Vehicle) SetEarliestStart(v int64)`

SetEarliestStart sets EarliestStart field to given value.

### HasEarliestStart

`func (o *Vehicle) HasEarliestStart() bool`

HasEarliestStart returns a boolean if a field has been set.

### GetLatestEnd

`func (o *Vehicle) GetLatestEnd() int64`

GetLatestEnd returns the LatestEnd field if non-nil, zero value otherwise.

### GetLatestEndOk

`func (o *Vehicle) GetLatestEndOk() (*int64, bool)`

GetLatestEndOk returns a tuple with the LatestEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestEnd

`func (o *Vehicle) SetLatestEnd(v int64)`

SetLatestEnd sets LatestEnd field to given value.

### HasLatestEnd

`func (o *Vehicle) HasLatestEnd() bool`

HasLatestEnd returns a boolean if a field has been set.

### GetSkills

`func (o *Vehicle) GetSkills() []string`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *Vehicle) GetSkillsOk() (*[]string, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *Vehicle) SetSkills(v []string)`

SetSkills sets Skills field to given value.

### HasSkills

`func (o *Vehicle) HasSkills() bool`

HasSkills returns a boolean if a field has been set.

### GetMaxDistance

`func (o *Vehicle) GetMaxDistance() int64`

GetMaxDistance returns the MaxDistance field if non-nil, zero value otherwise.

### GetMaxDistanceOk

`func (o *Vehicle) GetMaxDistanceOk() (*int64, bool)`

GetMaxDistanceOk returns a tuple with the MaxDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDistance

`func (o *Vehicle) SetMaxDistance(v int64)`

SetMaxDistance sets MaxDistance field to given value.

### HasMaxDistance

`func (o *Vehicle) HasMaxDistance() bool`

HasMaxDistance returns a boolean if a field has been set.

### GetMaxDrivingTime

`func (o *Vehicle) GetMaxDrivingTime() int64`

GetMaxDrivingTime returns the MaxDrivingTime field if non-nil, zero value otherwise.

### GetMaxDrivingTimeOk

`func (o *Vehicle) GetMaxDrivingTimeOk() (*int64, bool)`

GetMaxDrivingTimeOk returns a tuple with the MaxDrivingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDrivingTime

`func (o *Vehicle) SetMaxDrivingTime(v int64)`

SetMaxDrivingTime sets MaxDrivingTime field to given value.

### HasMaxDrivingTime

`func (o *Vehicle) HasMaxDrivingTime() bool`

HasMaxDrivingTime returns a boolean if a field has been set.

### GetMaxJobs

`func (o *Vehicle) GetMaxJobs() int32`

GetMaxJobs returns the MaxJobs field if non-nil, zero value otherwise.

### GetMaxJobsOk

`func (o *Vehicle) GetMaxJobsOk() (*int32, bool)`

GetMaxJobsOk returns a tuple with the MaxJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxJobs

`func (o *Vehicle) SetMaxJobs(v int32)`

SetMaxJobs sets MaxJobs field to given value.

### HasMaxJobs

`func (o *Vehicle) HasMaxJobs() bool`

HasMaxJobs returns a boolean if a field has been set.

### GetMinJobs

`func (o *Vehicle) GetMinJobs() int32`

GetMinJobs returns the MinJobs field if non-nil, zero value otherwise.

### GetMinJobsOk

`func (o *Vehicle) GetMinJobsOk() (*int32, bool)`

GetMinJobsOk returns a tuple with the MinJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinJobs

`func (o *Vehicle) SetMinJobs(v int32)`

SetMinJobs sets MinJobs field to given value.

### HasMinJobs

`func (o *Vehicle) HasMinJobs() bool`

HasMinJobs returns a boolean if a field has been set.

### GetMaxActivities

`func (o *Vehicle) GetMaxActivities() int32`

GetMaxActivities returns the MaxActivities field if non-nil, zero value otherwise.

### GetMaxActivitiesOk

`func (o *Vehicle) GetMaxActivitiesOk() (*int32, bool)`

GetMaxActivitiesOk returns a tuple with the MaxActivities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxActivities

`func (o *Vehicle) SetMaxActivities(v int32)`

SetMaxActivities sets MaxActivities field to given value.

### HasMaxActivities

`func (o *Vehicle) HasMaxActivities() bool`

HasMaxActivities returns a boolean if a field has been set.

### GetMoveToEndAddress

`func (o *Vehicle) GetMoveToEndAddress() bool`

GetMoveToEndAddress returns the MoveToEndAddress field if non-nil, zero value otherwise.

### GetMoveToEndAddressOk

`func (o *Vehicle) GetMoveToEndAddressOk() (*bool, bool)`

GetMoveToEndAddressOk returns a tuple with the MoveToEndAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveToEndAddress

`func (o *Vehicle) SetMoveToEndAddress(v bool)`

SetMoveToEndAddress sets MoveToEndAddress field to given value.

### HasMoveToEndAddress

`func (o *Vehicle) HasMoveToEndAddress() bool`

HasMoveToEndAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


