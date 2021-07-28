# Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Specifies the id of the service. Ids need to be unique so there must not be two services/shipments with the same id. | 
**Type** | Pointer to **string** | Specifies type of service. This makes a difference if items are loaded or unloaded, i.e. if one of the size dimensions &gt; 0. If it is specified as &#x60;service&#x60; or &#x60;pickup&#x60;, items are loaded and will stay in the vehicle for the rest of the route (and thus consumes capacity for the rest of the route). If it is a &#x60;delivery&#x60;, items are implicitly loaded at the beginning of the route and will stay in the route until delivery (and thus releases capacity for the rest of the route). | [optional] [default to "service"]
**Priority** | Pointer to **int32** | Specifies the priority. Can be 1 &#x3D; high priority to 10 &#x3D; low priority. Often there are more services/shipments than the available vehicle fleet can handle. Then you can set priorities to differentiate high priority tasks from those that could be left unassigned. I.e. the lower the priority the earlier these tasks are omitted in the solution. | [optional] [default to 2]
**Name** | Pointer to **string** | Meaningful name for service, e.g. &#x60;\&quot;deliver pizza\&quot;&#x60;. | [optional] 
**Address** | Pointer to [**Address**](Address.md) |  | [optional] 
**Duration** | Pointer to **int64** | Specifies the duration of the service in seconds, i.e. how long it takes at the customer site. | [optional] [default to 0]
**PreparationTime** | Pointer to **int64** | Specifies the preparation time in seconds. It can be used to model parking lot search time since if you have 3 identical locations in a row, it only falls due once. | [optional] [default to 0]
**TimeWindows** | Pointer to [**[]TimeWindow**](TimeWindow.md) | Specifies an array of time window objects (see time_window object below). Specify the time either with the recommended Unix time stamp (the number of seconds since 1970-01-01) or you can also count the seconds relative to Monday morning 00:00 and define the whole week in seconds. For example, Monday 9am is then represented by 9hour * 3600sec/hour &#x3D; 32400. In turn, Wednesday 1pm corresponds to 2day * 24hour/day * 3600sec/hour + 1day * 13hour/day * 3600sec/hour &#x3D; 219600. See this tutorial for more information. | [optional] 
**Size** | Pointer to **[]int32** | Size can have multiple dimensions and should be in line with the capacity dimension array of the vehicle type. For example, if the item that needs to be delivered has two size dimension, volume and weight, then specify it as follow [ 20, 5 ] assuming a volume of 20 and a weight of 5. | [optional] [default to [0]]
**RequiredSkills** | Pointer to **[]string** | Specifies an array of required skills, i.e. array of string (not case sensitive). For example, if this service needs to be conducted by a technician having a &#x60;drilling_machine&#x60; and a &#x60;screw_driver&#x60; then specify the array as follows: &#x60;[\&quot;drilling_machine\&quot;,\&quot;screw_driver\&quot;]&#x60;. This means that the service can only be done by a vehicle (technician) that has the skills &#x60;drilling_machine&#x60; AND &#x60;screw_driver&#x60; in its skill array. Otherwise it remains unassigned. | [optional] 
**AllowedVehicles** | Pointer to **[]string** | Specifies an array of allowed vehicles, i.e. array of vehicle ids. For example, if this service can only be conducted EITHER by &#x60;technician_peter&#x60; OR &#x60;technician_stefan&#x60; specify this as follows: &#x60;[\&quot;technician_peter\&quot;,\&quot;technician_stefan\&quot;]&#x60;. | [optional] 
**DisallowedVehicles** | Pointer to **[]string** | Specifies an array of disallowed vehicles, i.e. array of vehicle ids. | [optional] 
**MaxTimeInVehicle** | Pointer to **int64** | Specifies the maximum time in seconds a delivery can stay in the vehicle. Currently, it only works with services of \&quot;type\&quot;:\&quot;delivery\&quot;. | [optional] 
**Group** | Pointer to **string** | Group this service belongs to. See the group relation and [this post](https://discuss.graphhopper.com/t/4040) on how to utilize this. | [optional] 

## Methods

### NewService

`func NewService(id string, ) *Service`

NewService instantiates a new Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWithDefaults

`func NewServiceWithDefaults() *Service`

NewServiceWithDefaults instantiates a new Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Service) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Service) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Service) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *Service) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Service) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Service) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Service) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPriority

`func (o *Service) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Service) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Service) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *Service) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetName

`func (o *Service) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Service) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Service) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Service) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddress

`func (o *Service) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Service) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Service) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Service) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetDuration

`func (o *Service) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Service) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Service) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *Service) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetPreparationTime

`func (o *Service) GetPreparationTime() int64`

GetPreparationTime returns the PreparationTime field if non-nil, zero value otherwise.

### GetPreparationTimeOk

`func (o *Service) GetPreparationTimeOk() (*int64, bool)`

GetPreparationTimeOk returns a tuple with the PreparationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparationTime

`func (o *Service) SetPreparationTime(v int64)`

SetPreparationTime sets PreparationTime field to given value.

### HasPreparationTime

`func (o *Service) HasPreparationTime() bool`

HasPreparationTime returns a boolean if a field has been set.

### GetTimeWindows

`func (o *Service) GetTimeWindows() []TimeWindow`

GetTimeWindows returns the TimeWindows field if non-nil, zero value otherwise.

### GetTimeWindowsOk

`func (o *Service) GetTimeWindowsOk() (*[]TimeWindow, bool)`

GetTimeWindowsOk returns a tuple with the TimeWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindows

`func (o *Service) SetTimeWindows(v []TimeWindow)`

SetTimeWindows sets TimeWindows field to given value.

### HasTimeWindows

`func (o *Service) HasTimeWindows() bool`

HasTimeWindows returns a boolean if a field has been set.

### GetSize

`func (o *Service) GetSize() []int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Service) GetSizeOk() (*[]int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Service) SetSize(v []int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *Service) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetRequiredSkills

`func (o *Service) GetRequiredSkills() []string`

GetRequiredSkills returns the RequiredSkills field if non-nil, zero value otherwise.

### GetRequiredSkillsOk

`func (o *Service) GetRequiredSkillsOk() (*[]string, bool)`

GetRequiredSkillsOk returns a tuple with the RequiredSkills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredSkills

`func (o *Service) SetRequiredSkills(v []string)`

SetRequiredSkills sets RequiredSkills field to given value.

### HasRequiredSkills

`func (o *Service) HasRequiredSkills() bool`

HasRequiredSkills returns a boolean if a field has been set.

### GetAllowedVehicles

`func (o *Service) GetAllowedVehicles() []string`

GetAllowedVehicles returns the AllowedVehicles field if non-nil, zero value otherwise.

### GetAllowedVehiclesOk

`func (o *Service) GetAllowedVehiclesOk() (*[]string, bool)`

GetAllowedVehiclesOk returns a tuple with the AllowedVehicles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedVehicles

`func (o *Service) SetAllowedVehicles(v []string)`

SetAllowedVehicles sets AllowedVehicles field to given value.

### HasAllowedVehicles

`func (o *Service) HasAllowedVehicles() bool`

HasAllowedVehicles returns a boolean if a field has been set.

### GetDisallowedVehicles

`func (o *Service) GetDisallowedVehicles() []string`

GetDisallowedVehicles returns the DisallowedVehicles field if non-nil, zero value otherwise.

### GetDisallowedVehiclesOk

`func (o *Service) GetDisallowedVehiclesOk() (*[]string, bool)`

GetDisallowedVehiclesOk returns a tuple with the DisallowedVehicles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisallowedVehicles

`func (o *Service) SetDisallowedVehicles(v []string)`

SetDisallowedVehicles sets DisallowedVehicles field to given value.

### HasDisallowedVehicles

`func (o *Service) HasDisallowedVehicles() bool`

HasDisallowedVehicles returns a boolean if a field has been set.

### GetMaxTimeInVehicle

`func (o *Service) GetMaxTimeInVehicle() int64`

GetMaxTimeInVehicle returns the MaxTimeInVehicle field if non-nil, zero value otherwise.

### GetMaxTimeInVehicleOk

`func (o *Service) GetMaxTimeInVehicleOk() (*int64, bool)`

GetMaxTimeInVehicleOk returns a tuple with the MaxTimeInVehicle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTimeInVehicle

`func (o *Service) SetMaxTimeInVehicle(v int64)`

SetMaxTimeInVehicle sets MaxTimeInVehicle field to given value.

### HasMaxTimeInVehicle

`func (o *Service) HasMaxTimeInVehicle() bool`

HasMaxTimeInVehicle returns a boolean if a field has been set.

### GetGroup

`func (o *Service) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *Service) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *Service) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *Service) HasGroup() bool`

HasGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


