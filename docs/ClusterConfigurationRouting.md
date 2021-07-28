# ClusterConfigurationRouting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | Pointer to **string** | The routing profile for which the travel times and distances should be calculated. Other profiles are listed [here](#section/Map-Data-and-Routing-Profiles/OpenStreetMap) | [optional] 
**CostPerSecond** | Pointer to **float64** | Cost per second (travel time) | [optional] 
**CostPerMeter** | Pointer to **float64** | Cost per meter (travel distance) | [optional] 

## Methods

### NewClusterConfigurationRouting

`func NewClusterConfigurationRouting() *ClusterConfigurationRouting`

NewClusterConfigurationRouting instantiates a new ClusterConfigurationRouting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterConfigurationRoutingWithDefaults

`func NewClusterConfigurationRoutingWithDefaults() *ClusterConfigurationRouting`

NewClusterConfigurationRoutingWithDefaults instantiates a new ClusterConfigurationRouting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfile

`func (o *ClusterConfigurationRouting) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ClusterConfigurationRouting) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ClusterConfigurationRouting) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *ClusterConfigurationRouting) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetCostPerSecond

`func (o *ClusterConfigurationRouting) GetCostPerSecond() float64`

GetCostPerSecond returns the CostPerSecond field if non-nil, zero value otherwise.

### GetCostPerSecondOk

`func (o *ClusterConfigurationRouting) GetCostPerSecondOk() (*float64, bool)`

GetCostPerSecondOk returns a tuple with the CostPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerSecond

`func (o *ClusterConfigurationRouting) SetCostPerSecond(v float64)`

SetCostPerSecond sets CostPerSecond field to given value.

### HasCostPerSecond

`func (o *ClusterConfigurationRouting) HasCostPerSecond() bool`

HasCostPerSecond returns a boolean if a field has been set.

### GetCostPerMeter

`func (o *ClusterConfigurationRouting) GetCostPerMeter() float64`

GetCostPerMeter returns the CostPerMeter field if non-nil, zero value otherwise.

### GetCostPerMeterOk

`func (o *ClusterConfigurationRouting) GetCostPerMeterOk() (*float64, bool)`

GetCostPerMeterOk returns a tuple with the CostPerMeter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerMeter

`func (o *ClusterConfigurationRouting) SetCostPerMeter(v float64)`

SetCostPerMeter sets CostPerMeter field to given value.

### HasCostPerMeter

`func (o *ClusterConfigurationRouting) HasCostPerMeter() bool`

HasCostPerMeter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


