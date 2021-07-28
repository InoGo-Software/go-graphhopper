# Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vehicles** | Pointer to [**[]Vehicle**](Vehicle.md) | Specifies the available vehicles. | [optional] 
**VehicleTypes** | Pointer to [**[]VehicleType**](VehicleType.md) | Specifies the available vehicle types. These types can be assigned to vehicles. | [optional] 
**Services** | Pointer to [**[]Service**](Service.md) | Specifies the orders of the type \&quot;service\&quot;. These are, for example, pick-ups, deliveries or other stops that are to be approached by the specified vehicles. Each of these orders contains only one location. | [optional] 
**Shipments** | Pointer to [**[]Shipment**](Shipment.md) | Specifies the available shipments. Each shipment contains a pickup and a delivery stop, which must be processed one after the other. | [optional] 
**Relations** | Pointer to [**[]AnyOfJobRelationGroupRelation**](AnyOfJobRelationGroupRelation.md) | Defines additional relationships between orders. | [optional] 
**Algorithm** | Pointer to [**Algorithm**](Algorithm.md) |  | [optional] 
**Objectives** | Pointer to [**[]Objective**](Objective.md) | Specifies an objective function. The vehicle routing problem is solved in such a way that this objective function is minimized. | [optional] 
**CostMatrices** | Pointer to [**[]CostMatrix**](CostMatrix.md) | Specifies your own tranport time and distance matrices. | [optional] 
**Configuration** | Pointer to [**Configuration**](Configuration.md) |  | [optional] 

## Methods

### NewRequest

`func NewRequest() *Request`

NewRequest instantiates a new Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestWithDefaults

`func NewRequestWithDefaults() *Request`

NewRequestWithDefaults instantiates a new Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVehicles

`func (o *Request) GetVehicles() []Vehicle`

GetVehicles returns the Vehicles field if non-nil, zero value otherwise.

### GetVehiclesOk

`func (o *Request) GetVehiclesOk() (*[]Vehicle, bool)`

GetVehiclesOk returns a tuple with the Vehicles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicles

`func (o *Request) SetVehicles(v []Vehicle)`

SetVehicles sets Vehicles field to given value.

### HasVehicles

`func (o *Request) HasVehicles() bool`

HasVehicles returns a boolean if a field has been set.

### GetVehicleTypes

`func (o *Request) GetVehicleTypes() []VehicleType`

GetVehicleTypes returns the VehicleTypes field if non-nil, zero value otherwise.

### GetVehicleTypesOk

`func (o *Request) GetVehicleTypesOk() (*[]VehicleType, bool)`

GetVehicleTypesOk returns a tuple with the VehicleTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleTypes

`func (o *Request) SetVehicleTypes(v []VehicleType)`

SetVehicleTypes sets VehicleTypes field to given value.

### HasVehicleTypes

`func (o *Request) HasVehicleTypes() bool`

HasVehicleTypes returns a boolean if a field has been set.

### GetServices

`func (o *Request) GetServices() []Service`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *Request) GetServicesOk() (*[]Service, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *Request) SetServices(v []Service)`

SetServices sets Services field to given value.

### HasServices

`func (o *Request) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetShipments

`func (o *Request) GetShipments() []Shipment`

GetShipments returns the Shipments field if non-nil, zero value otherwise.

### GetShipmentsOk

`func (o *Request) GetShipmentsOk() (*[]Shipment, bool)`

GetShipmentsOk returns a tuple with the Shipments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipments

`func (o *Request) SetShipments(v []Shipment)`

SetShipments sets Shipments field to given value.

### HasShipments

`func (o *Request) HasShipments() bool`

HasShipments returns a boolean if a field has been set.

### GetRelations

`func (o *Request) GetRelations() []AnyOfJobRelationGroupRelation`

GetRelations returns the Relations field if non-nil, zero value otherwise.

### GetRelationsOk

`func (o *Request) GetRelationsOk() (*[]AnyOfJobRelationGroupRelation, bool)`

GetRelationsOk returns a tuple with the Relations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelations

`func (o *Request) SetRelations(v []AnyOfJobRelationGroupRelation)`

SetRelations sets Relations field to given value.

### HasRelations

`func (o *Request) HasRelations() bool`

HasRelations returns a boolean if a field has been set.

### GetAlgorithm

`func (o *Request) GetAlgorithm() Algorithm`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *Request) GetAlgorithmOk() (*Algorithm, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *Request) SetAlgorithm(v Algorithm)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *Request) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetObjectives

`func (o *Request) GetObjectives() []Objective`

GetObjectives returns the Objectives field if non-nil, zero value otherwise.

### GetObjectivesOk

`func (o *Request) GetObjectivesOk() (*[]Objective, bool)`

GetObjectivesOk returns a tuple with the Objectives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectives

`func (o *Request) SetObjectives(v []Objective)`

SetObjectives sets Objectives field to given value.

### HasObjectives

`func (o *Request) HasObjectives() bool`

HasObjectives returns a boolean if a field has been set.

### GetCostMatrices

`func (o *Request) GetCostMatrices() []CostMatrix`

GetCostMatrices returns the CostMatrices field if non-nil, zero value otherwise.

### GetCostMatricesOk

`func (o *Request) GetCostMatricesOk() (*[]CostMatrix, bool)`

GetCostMatricesOk returns a tuple with the CostMatrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostMatrices

`func (o *Request) SetCostMatrices(v []CostMatrix)`

SetCostMatrices sets CostMatrices field to given value.

### HasCostMatrices

`func (o *Request) HasCostMatrices() bool`

HasCostMatrices returns a boolean if a field has been set.

### GetConfiguration

`func (o *Request) GetConfiguration() Configuration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *Request) GetConfigurationOk() (*Configuration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *Request) SetConfiguration(v Configuration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *Request) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


