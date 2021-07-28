# VehicleType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeId** | **string** | Specifies the id of the vehicle type. If a vehicle needs to be of this type, it should refer to this with its type_id attribute. | 
**Profile** | Pointer to **interface{}** |  | [optional] 
**Capacity** | Pointer to **[]int32** | Specifies an array of capacity dimension values which need to be int values. For example, if there are two dimensions such as volume and weight then it needs to be defined as [ 1000, 300 ] assuming a maximum volume of 1000 and a maximum weight of 300. | [optional] [default to [0]]
**SpeedFactor** | Pointer to **float64** | Specifies a speed factor for this vehicle type. If the vehicle that uses this type needs to be only half as fast as what is actually calculated with our routing engine then set the speed factor to 0.5. | [optional] [default to 1]
**ServiceTimeFactor** | Pointer to **float64** | Specifies a service time factor for this vehicle type. If the vehicle/driver that uses this type is able to conduct the service as double as fast as it is determined in the corresponding service or shipment then set it to 0.5. | [optional] [default to 1]
**CostPerMeter** | Pointer to **float64** | **_BETA feature_**! Cost parameter per distance unit, here meter is used | [optional] 
**CostPerSecond** | Pointer to **float64** | **_BETA feature_**! Cost parameter per time unit, here second is used | [optional] 
**CostPerActivation** | Pointer to **float64** | **_BETA feature_**! Cost parameter vehicle activation, i.e. fixed costs per vehicle | [optional] 
**ConsiderTraffic** | Pointer to **bool** | Specifies whether traffic should be considered. if \&quot;tomtom\&quot; is used and this is false, free flow travel times from \&quot;tomtom\&quot; are calculated. If this is true, historical traffic info are used. We do not yet have traffic data for \&quot;openstreetmap\&quot;, thus, setting this true has no effect at all. | [optional] [default to false]
**NetworkDataProvider** | Pointer to **string** | Specifies the network data provider. Either use [&#x60;openstreetmap&#x60;](#section/Map-Data-and-Routing-Profiles/OpenStreetMap) (default) or [&#x60;tomtom&#x60;](#section/Map-Data-and-Routing-Profiles/TomTom) (add-on required). | [optional] [default to "openstreetmap"]

## Methods

### NewVehicleType

`func NewVehicleType(typeId string, ) *VehicleType`

NewVehicleType instantiates a new VehicleType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleTypeWithDefaults

`func NewVehicleTypeWithDefaults() *VehicleType`

NewVehicleTypeWithDefaults instantiates a new VehicleType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeId

`func (o *VehicleType) GetTypeId() string`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *VehicleType) GetTypeIdOk() (*string, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *VehicleType) SetTypeId(v string)`

SetTypeId sets TypeId field to given value.


### GetProfile

`func (o *VehicleType) GetProfile() interface{}`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *VehicleType) GetProfileOk() (*interface{}, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *VehicleType) SetProfile(v interface{})`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *VehicleType) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetCapacity

`func (o *VehicleType) GetCapacity() []int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *VehicleType) GetCapacityOk() (*[]int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *VehicleType) SetCapacity(v []int32)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *VehicleType) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetSpeedFactor

`func (o *VehicleType) GetSpeedFactor() float64`

GetSpeedFactor returns the SpeedFactor field if non-nil, zero value otherwise.

### GetSpeedFactorOk

`func (o *VehicleType) GetSpeedFactorOk() (*float64, bool)`

GetSpeedFactorOk returns a tuple with the SpeedFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedFactor

`func (o *VehicleType) SetSpeedFactor(v float64)`

SetSpeedFactor sets SpeedFactor field to given value.

### HasSpeedFactor

`func (o *VehicleType) HasSpeedFactor() bool`

HasSpeedFactor returns a boolean if a field has been set.

### GetServiceTimeFactor

`func (o *VehicleType) GetServiceTimeFactor() float64`

GetServiceTimeFactor returns the ServiceTimeFactor field if non-nil, zero value otherwise.

### GetServiceTimeFactorOk

`func (o *VehicleType) GetServiceTimeFactorOk() (*float64, bool)`

GetServiceTimeFactorOk returns a tuple with the ServiceTimeFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTimeFactor

`func (o *VehicleType) SetServiceTimeFactor(v float64)`

SetServiceTimeFactor sets ServiceTimeFactor field to given value.

### HasServiceTimeFactor

`func (o *VehicleType) HasServiceTimeFactor() bool`

HasServiceTimeFactor returns a boolean if a field has been set.

### GetCostPerMeter

`func (o *VehicleType) GetCostPerMeter() float64`

GetCostPerMeter returns the CostPerMeter field if non-nil, zero value otherwise.

### GetCostPerMeterOk

`func (o *VehicleType) GetCostPerMeterOk() (*float64, bool)`

GetCostPerMeterOk returns a tuple with the CostPerMeter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerMeter

`func (o *VehicleType) SetCostPerMeter(v float64)`

SetCostPerMeter sets CostPerMeter field to given value.

### HasCostPerMeter

`func (o *VehicleType) HasCostPerMeter() bool`

HasCostPerMeter returns a boolean if a field has been set.

### GetCostPerSecond

`func (o *VehicleType) GetCostPerSecond() float64`

GetCostPerSecond returns the CostPerSecond field if non-nil, zero value otherwise.

### GetCostPerSecondOk

`func (o *VehicleType) GetCostPerSecondOk() (*float64, bool)`

GetCostPerSecondOk returns a tuple with the CostPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerSecond

`func (o *VehicleType) SetCostPerSecond(v float64)`

SetCostPerSecond sets CostPerSecond field to given value.

### HasCostPerSecond

`func (o *VehicleType) HasCostPerSecond() bool`

HasCostPerSecond returns a boolean if a field has been set.

### GetCostPerActivation

`func (o *VehicleType) GetCostPerActivation() float64`

GetCostPerActivation returns the CostPerActivation field if non-nil, zero value otherwise.

### GetCostPerActivationOk

`func (o *VehicleType) GetCostPerActivationOk() (*float64, bool)`

GetCostPerActivationOk returns a tuple with the CostPerActivation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerActivation

`func (o *VehicleType) SetCostPerActivation(v float64)`

SetCostPerActivation sets CostPerActivation field to given value.

### HasCostPerActivation

`func (o *VehicleType) HasCostPerActivation() bool`

HasCostPerActivation returns a boolean if a field has been set.

### GetConsiderTraffic

`func (o *VehicleType) GetConsiderTraffic() bool`

GetConsiderTraffic returns the ConsiderTraffic field if non-nil, zero value otherwise.

### GetConsiderTrafficOk

`func (o *VehicleType) GetConsiderTrafficOk() (*bool, bool)`

GetConsiderTrafficOk returns a tuple with the ConsiderTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsiderTraffic

`func (o *VehicleType) SetConsiderTraffic(v bool)`

SetConsiderTraffic sets ConsiderTraffic field to given value.

### HasConsiderTraffic

`func (o *VehicleType) HasConsiderTraffic() bool`

HasConsiderTraffic returns a boolean if a field has been set.

### GetNetworkDataProvider

`func (o *VehicleType) GetNetworkDataProvider() string`

GetNetworkDataProvider returns the NetworkDataProvider field if non-nil, zero value otherwise.

### GetNetworkDataProviderOk

`func (o *VehicleType) GetNetworkDataProviderOk() (*string, bool)`

GetNetworkDataProviderOk returns a tuple with the NetworkDataProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDataProvider

`func (o *VehicleType) SetNetworkDataProvider(v string)`

SetNetworkDataProvider sets NetworkDataProvider field to given value.

### HasNetworkDataProvider

`func (o *VehicleType) HasNetworkDataProvider() bool`

HasNetworkDataProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


