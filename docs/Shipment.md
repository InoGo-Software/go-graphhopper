# Shipment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Specifies the id of the shipment. Ids need to be unique so there must not be two services/shipments with the same id. | 
**Name** | Pointer to **string** | Meaningful name for shipment, e.g. \&quot;pickup and deliver pizza to Peter\&quot;. | [optional] 
**Priority** | Pointer to **int32** | Specifies the priority. Can be 1 &#x3D; high priority to 10 &#x3D; low priority. Often there are more services/shipments than the available vehicle fleet can handle. Then you can set priorities to differentiate high priority tasks from those that could be left unassigned. I.e. the lower the priority the earlier these tasks are omitted in the solution. | [optional] [default to 2]
**Pickup** | [**Stop**](Stop.md) |  | 
**Delivery** | [**Stop**](Stop.md) |  | 
**Size** | Pointer to **[]int32** | Size can have multiple dimensions and should be in line with the capacity dimension array of the vehicle type. For example, if the item that needs to be delivered has two size dimension, volume and weight, then specify it as follow [ 20, 5 ] assuming a volume of 20 and a weight of 5. | [optional] [default to [0]]
**RequiredSkills** | Pointer to **[]string** | Specifies an array of required skills, i.e. array of string (not case sensitive). For example, if this shipment needs to be conducted by a technician having a &#x60;drilling_machine&#x60; and a &#x60;screw_driver&#x60; then specify the array as follows: &#x60;[\&quot;drilling_machine\&quot;,\&quot;screw_driver\&quot;]&#x60;. This means that the service can only be done by a vehicle (technician) that has the skills &#x60;drilling_machine&#x60; AND &#x60;screw_driver&#x60; in its skill array. Otherwise it remains unassigned. | [optional] 
**AllowedVehicles** | Pointer to **[]string** | Specifies an array of allowed vehicles, i.e. array of vehicle ids. For example, if this shipment can only be conducted EITHER by \&quot;technician_peter\&quot; OR \&quot;technician_stefan\&quot; specify this as follows: [\&quot;technician_peter\&quot;,\&quot;technician_stefan\&quot;]. | [optional] 
**DisallowedVehicles** | Pointer to **[]string** | Specifies an array of disallowed vehicles, i.e. array of vehicle ids. | [optional] 
**MaxTimeInVehicle** | Pointer to **int64** | Specifies the maximum time in seconds a shipment can stay in the vehicle. | [optional] 

## Methods

### NewShipment

`func NewShipment(id string, pickup Stop, delivery Stop, ) *Shipment`

NewShipment instantiates a new Shipment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentWithDefaults

`func NewShipmentWithDefaults() *Shipment`

NewShipmentWithDefaults instantiates a new Shipment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Shipment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Shipment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Shipment) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Shipment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Shipment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Shipment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Shipment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriority

`func (o *Shipment) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Shipment) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Shipment) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *Shipment) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetPickup

`func (o *Shipment) GetPickup() Stop`

GetPickup returns the Pickup field if non-nil, zero value otherwise.

### GetPickupOk

`func (o *Shipment) GetPickupOk() (*Stop, bool)`

GetPickupOk returns a tuple with the Pickup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickup

`func (o *Shipment) SetPickup(v Stop)`

SetPickup sets Pickup field to given value.


### GetDelivery

`func (o *Shipment) GetDelivery() Stop`

GetDelivery returns the Delivery field if non-nil, zero value otherwise.

### GetDeliveryOk

`func (o *Shipment) GetDeliveryOk() (*Stop, bool)`

GetDeliveryOk returns a tuple with the Delivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivery

`func (o *Shipment) SetDelivery(v Stop)`

SetDelivery sets Delivery field to given value.


### GetSize

`func (o *Shipment) GetSize() []int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Shipment) GetSizeOk() (*[]int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Shipment) SetSize(v []int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *Shipment) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetRequiredSkills

`func (o *Shipment) GetRequiredSkills() []string`

GetRequiredSkills returns the RequiredSkills field if non-nil, zero value otherwise.

### GetRequiredSkillsOk

`func (o *Shipment) GetRequiredSkillsOk() (*[]string, bool)`

GetRequiredSkillsOk returns a tuple with the RequiredSkills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredSkills

`func (o *Shipment) SetRequiredSkills(v []string)`

SetRequiredSkills sets RequiredSkills field to given value.

### HasRequiredSkills

`func (o *Shipment) HasRequiredSkills() bool`

HasRequiredSkills returns a boolean if a field has been set.

### GetAllowedVehicles

`func (o *Shipment) GetAllowedVehicles() []string`

GetAllowedVehicles returns the AllowedVehicles field if non-nil, zero value otherwise.

### GetAllowedVehiclesOk

`func (o *Shipment) GetAllowedVehiclesOk() (*[]string, bool)`

GetAllowedVehiclesOk returns a tuple with the AllowedVehicles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedVehicles

`func (o *Shipment) SetAllowedVehicles(v []string)`

SetAllowedVehicles sets AllowedVehicles field to given value.

### HasAllowedVehicles

`func (o *Shipment) HasAllowedVehicles() bool`

HasAllowedVehicles returns a boolean if a field has been set.

### GetDisallowedVehicles

`func (o *Shipment) GetDisallowedVehicles() []string`

GetDisallowedVehicles returns the DisallowedVehicles field if non-nil, zero value otherwise.

### GetDisallowedVehiclesOk

`func (o *Shipment) GetDisallowedVehiclesOk() (*[]string, bool)`

GetDisallowedVehiclesOk returns a tuple with the DisallowedVehicles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisallowedVehicles

`func (o *Shipment) SetDisallowedVehicles(v []string)`

SetDisallowedVehicles sets DisallowedVehicles field to given value.

### HasDisallowedVehicles

`func (o *Shipment) HasDisallowedVehicles() bool`

HasDisallowedVehicles returns a boolean if a field has been set.

### GetMaxTimeInVehicle

`func (o *Shipment) GetMaxTimeInVehicle() int64`

GetMaxTimeInVehicle returns the MaxTimeInVehicle field if non-nil, zero value otherwise.

### GetMaxTimeInVehicleOk

`func (o *Shipment) GetMaxTimeInVehicleOk() (*int64, bool)`

GetMaxTimeInVehicleOk returns a tuple with the MaxTimeInVehicle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTimeInVehicle

`func (o *Shipment) SetMaxTimeInVehicle(v int64)`

SetMaxTimeInVehicle sets MaxTimeInVehicle field to given value.

### HasMaxTimeInVehicle

`func (o *Shipment) HasMaxTimeInVehicle() bool`

HasMaxTimeInVehicle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


