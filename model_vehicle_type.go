/*
 * GraphHopper Directions API
 *
 *  With the [GraphHopper Directions API](https://www.graphhopper.com/products/) you can integrate A-to-B route planning, turn-by-turn navigation, route optimization, isochrone calculations and other tools in your application.  The GraphHopper Directions API consists of the following RESTful web services:   * [Routing API](#tag/Routing-API),  * [Route Optimization API](#tag/Route-Optimization-API),  * [Isochrone API](#tag/Isochrone-API),  * [Map Matching API](#tag/Map-Matching-API),  * [Matrix API](#tag/Matrix-API),  * [Geocoding API](#tag/Geocoding-API) and  * [Cluster API](#tag/Cluster-API).  # Explore our APIs  ## Get started  1. [Sign up for GraphHopper](https://support.graphhopper.com/a/solutions/articles/44001976025) 2. [Create an API key](https://support.graphhopper.com/a/solutions/articles/44001976027)  Each API part has its own documentation. Jump to the desired API part and learn about the API through the given examples and tutorials.  In addition, for each API there are specific sample requests that you can send via Insomnia or Postman to see what the requests and responses look like.  ## Insomnia  To explore our APIs with [Insomnia](https://insomnia.rest/), follow these steps:  1. Open Insomnia and Import [our workspace](https://raw.githubusercontent.com/graphhopper/directions-api-doc/master/web/restclients/GraphHopper-Direction-API-Insomnia.json). 2. Specify [your API key](https://graphhopper.com/dashboard/#/register) in your workspace: Manage Environments -> Base Environment -> `\"api_key\": your API key` 3. Start exploring  ![Insomnia](./img/insomnia.png)  ## Postman  To explore our APIs with [Postman](https://www.getpostman.com/), follow these steps:  1. Import our [request collections](https://raw.githubusercontent.com/graphhopper/directions-api-doc/master/web/restclients/graphhopper_directions_api.postman_collection.json) as well as our [environment file](https://raw.githubusercontent.com/graphhopper/directions-api-doc/master/web/restclients/graphhopper_directions_api.postman_environment.json). 2. Specify [your API key](https://graphhopper.com/dashboard/#/register) in your environment: `\"api_key\": your API key` 3. Start exploring  ![Postman](./img/postman.png)  ## API Client Libraries  To speed up development and make coding easier, we offer the following client libraries:   * [JavaScript client](https://github.com/graphhopper/directions-api-js-client) - try the [live examples](https://graphhopper.com/api/1/examples/)  * [Others](https://github.com/graphhopper/directions-api-clients) like C#, Ruby, PHP, Python, ... automatically created for the Route Optimization API  ### Bandwidth reduction  If you create your own client, make sure it supports http/2 and gzipped responses for best speed.  If you use the Matrix, the Route Optimization API or the Cluster API and want to solve large problems, we recommend you to reduce bandwidth by [compressing your POST request](https://gist.github.com/karussell/82851e303ea7b3459b2dea01f18949f4) and specifying the header as follows: `Content-Encoding: gzip`. This will also avoid the HTTP 413 error \"Request Entity Too Large\".  ## Contact Us  If you have problems or questions, please read the following information:  - [FAQ](https://graphhopper.com/api/1/docs/FAQ/) - [Public forum](https://discuss.graphhopper.com/c/directions-api) - [Contact us](https://www.graphhopper.com/contact-form/) - [GraphHopper Status Page](https://status.graphhopper.com/)  To stay informed about the latest developments, you can  - follow us on [twitter](https://twitter.com/graphhopper/), - read [our blog](https://graphhopper.com/blog/), - watch [our documentation repository](https://github.com/graphhopper/directions-api-doc), - sign up for our newsletter or - [our forum](https://discuss.graphhopper.com/c/directions-api).  Select the channel you like the most.   # Map Data and Routing Profiles  Currently, our main data source is [OpenStreetMap](https://www.openstreetmap.org). We also integrated other network data providers. This chapter gives an overview about the options you have.  ## OpenStreetMap  #### Geographical Coverage  [OpenStreetMap](https://www.openstreetmap.org) covers the whole world. If you want to see for yourself if we can provide data suitable for your region, please visit [GraphHopper Maps](https://graphhopper.com/maps/). You can edit and modify OpenStreetMap data if you find that important information is missing, e.g. a weight limit for a bridge. [Here](https://wiki.openstreetmap.org/wiki/Beginners%27_guide) is a beginner's guide that shows how to add data. If you have edited data, we usually consider your data after 1 week at the latest.  #### Supported Vehicle Profiles  The Routing, Matrix and Route Optimization APIs support the following vehicle profiles:  Name       | Description           | Restrictions              | Icon -----------|:----------------------|:--------------------------|:--------------------------------------------------------- car        | Car mode              | car access                | ![car image](https://graphhopper.com/maps/img/car.png) small_truck| Small truck like a Mercedes Sprinter, Ford Transit or Iveco Daily | height=2.7m, width=2+0.4m, length=5.5m, weight=2080+1400 kg | ![small truck image](https://graphhopper.com/maps/img/small_truck.png) truck      | Truck like a MAN or Mercedes-Benz Actros | height=3.7m, width=2.6+0.5m, length=12m, weight=13000 + 13000 kg, hgv=yes, 3 Axes | ![truck image](https://graphhopper.com/maps/img/truck.png) scooter    | Moped mode | Fast inner city, often used for food delivery, is able to ignore certain bollards, maximum speed of roughly 50km/h | ![scooter image](https://graphhopper.com/maps/img/scooter.png) foot       | Pedestrian or walking without dangerous [SAC-scales](https://wiki.openstreetmap.org/wiki/Key:sac_scale) | foot access         | ![foot image](https://graphhopper.com/maps/img/foot.png) hike       | Pedestrian or walking with priority for more beautiful hiking tours and potentially a bit longer than `foot`. Walking duration is influenced by elevation differences.  | foot access         | ![hike image](https://graphhopper.com/maps/img/hike.png) bike       | Trekking bike avoiding hills | bike access  | ![bike image](https://graphhopper.com/maps/img/bike.png) mtb        | Mountainbike          | bike access         | ![Mountainbike image](https://graphhopper.com/maps/img/mtb.png) racingbike| Bike preferring roads | bike access         | ![racingbike image](https://graphhopper.com/maps/img/racingbike.png)  Please note:   * all motor vehicles (`car`, `small_truck`, `truck` and `scooter`) support turn restrictions via `turn_costs=true`  * the free package supports only the vehicle profiles `car`, `bike` or `foot`  * up to 2 different vehicle profiles can be used in a single optimization request. The number of vehicles is unaffected and depends on your subscription.  * we offer custom vehicle profiles with different properties, different speed profiles or different access options. To find out more about custom profiles, please [contact us](https://www.graphhopper.com/contact-form/).  * a sophisticated `motorcycle` profile is available up on request. It is powered by the [Kurviger](https://kurviger.de/en) Routing API and favors curves and slopes while avoiding cities and highways.   ## TomTom  If you want to include traffic, you can purchase the TomTom Add-on. This Add-on only uses TomTom's road network and historical traffic information. Live traffic is not yet considered. If you are interested to learn how we consider traffic information, we recommend that you read [this article](https://www.graphhopper.com/blog/2017/11/06/time-dependent-optimization/).  Please note the following:   * Currently we only offer this for our [Route Optimization API](#tag/Route-Optimization-API).  * In addition to our terms, you need to accept TomTom's [End User License Aggreement](https://www.graphhopper.com/tomtom-end-user-license-agreement/).  * We do *not* use TomTom's web services. We only use their data with our software.   [Contact us](https://www.graphhopper.com/contact-form/) for more details.  #### Geographical Coverage  We offer  - Europe including Russia - North, Central and South America - Saudi Arabia - United Arab Emirates - South Africa - Australia  #### Supported Vehicle Profiles  Name       | Description           | Restrictions              | Icon -----------|:----------------------|:--------------------------|:--------------------------------------------------------- car        | Car mode              | car access                | ![car image](https://graphhopper.com/maps/img/car.png) small_truck| Small truck like a Mercedes Sprinter, Ford Transit or Iveco Daily | height=2.7m, width=2+0.4m, length=5.5m, weight=2080+1400 kg | ![small truck image](https://graphhopper.com/maps/img/small_truck.png) 
 *
 * API version: 1.0.0
 * Contact: support@graphhopper.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// VehicleType struct for VehicleType
type VehicleType struct {
	// Specifies the id of the vehicle type. If a vehicle needs to be of this type, it should refer to this with its type_id attribute.
	TypeId string `json:"type_id"`
	Profile *interface{} `json:"profile,omitempty"`
	// Specifies an array of capacity dimension values which need to be int values. For example, if there are two dimensions such as volume and weight then it needs to be defined as [ 1000, 300 ] assuming a maximum volume of 1000 and a maximum weight of 300.
	Capacity *[]int32 `json:"capacity,omitempty"`
	// Specifies a speed factor for this vehicle type. If the vehicle that uses this type needs to be only half as fast as what is actually calculated with our routing engine then set the speed factor to 0.5.
	SpeedFactor *float64 `json:"speed_factor,omitempty"`
	// Specifies a service time factor for this vehicle type. If the vehicle/driver that uses this type is able to conduct the service as double as fast as it is determined in the corresponding service or shipment then set it to 0.5.
	ServiceTimeFactor *float64 `json:"service_time_factor,omitempty"`
	// **_BETA feature_**! Cost parameter per distance unit, here meter is used
	CostPerMeter *float64 `json:"cost_per_meter,omitempty"`
	// **_BETA feature_**! Cost parameter per time unit, here second is used
	CostPerSecond *float64 `json:"cost_per_second,omitempty"`
	// **_BETA feature_**! Cost parameter vehicle activation, i.e. fixed costs per vehicle
	CostPerActivation *float64 `json:"cost_per_activation,omitempty"`
	// Specifies whether traffic should be considered. if \"tomtom\" is used and this is false, free flow travel times from \"tomtom\" are calculated. If this is true, historical traffic info are used. We do not yet have traffic data for \"openstreetmap\", thus, setting this true has no effect at all.
	ConsiderTraffic *bool `json:"consider_traffic,omitempty"`
	// Specifies the network data provider. Either use [`openstreetmap`](#section/Map-Data-and-Routing-Profiles/OpenStreetMap) (default) or [`tomtom`](#section/Map-Data-and-Routing-Profiles/TomTom) (add-on required).
	NetworkDataProvider *string `json:"network_data_provider,omitempty"`
}

// NewVehicleType instantiates a new VehicleType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVehicleType(typeId string) *VehicleType {
	this := VehicleType{}
	this.TypeId = typeId
	var speedFactor float64 = 1
	this.SpeedFactor = &speedFactor
	var serviceTimeFactor float64 = 1
	this.ServiceTimeFactor = &serviceTimeFactor
	var considerTraffic bool = false
	this.ConsiderTraffic = &considerTraffic
	var networkDataProvider string = "openstreetmap"
	this.NetworkDataProvider = &networkDataProvider
	return &this
}

// NewVehicleTypeWithDefaults instantiates a new VehicleType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVehicleTypeWithDefaults() *VehicleType {
	this := VehicleType{}
	var speedFactor float64 = 1
	this.SpeedFactor = &speedFactor
	var serviceTimeFactor float64 = 1
	this.ServiceTimeFactor = &serviceTimeFactor
	var considerTraffic bool = false
	this.ConsiderTraffic = &considerTraffic
	var networkDataProvider string = "openstreetmap"
	this.NetworkDataProvider = &networkDataProvider
	return &this
}

// GetTypeId returns the TypeId field value
func (o *VehicleType) GetTypeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TypeId
}

// GetTypeIdOk returns a tuple with the TypeId field value
// and a boolean to check if the value has been set.
func (o *VehicleType) GetTypeIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TypeId, true
}

// SetTypeId sets field value
func (o *VehicleType) SetTypeId(v string) {
	o.TypeId = v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *VehicleType) GetProfile() interface{} {
	if o == nil || o.Profile == nil {
		var ret interface{}
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleType) GetProfileOk() (*interface{}, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *VehicleType) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given interface{} and assigns it to the Profile field.
func (o *VehicleType) SetProfile(v interface{}) {
	o.Profile = &v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *VehicleType) GetCapacity() []int32 {
	if o == nil || o.Capacity == nil {
		var ret []int32
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleType) GetCapacityOk() (*[]int32, bool) {
	if o == nil || o.Capacity == nil {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *VehicleType) HasCapacity() bool {
	if o != nil && o.Capacity != nil {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given []int32 and assigns it to the Capacity field.
func (o *VehicleType) SetCapacity(v []int32) {
	o.Capacity = &v
}

// GetSpeedFactor returns the SpeedFactor field value if set, zero value otherwise.
func (o *VehicleType) GetSpeedFactor() float64 {
	if o == nil || o.SpeedFactor == nil {
		var ret float64
		return ret
	}
	return *o.SpeedFactor
}

// GetSpeedFactorOk returns a tuple with the SpeedFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleType) GetSpeedFactorOk() (*float64, bool) {
	if o == nil || o.SpeedFactor == nil {
		return nil, false
	}
	return o.SpeedFactor, true
}

// HasSpeedFactor returns a boolean if a field has been set.
func (o *VehicleType) HasSpeedFactor() bool {
	if o != nil && o.SpeedFactor != nil {
		return true
	}

	return false
}

// SetSpeedFactor gets a reference to the given float64 and assigns it to the SpeedFactor field.
func (o *VehicleType) SetSpeedFactor(v float64) {
	o.SpeedFactor = &v
}

// GetServiceTimeFactor returns the ServiceTimeFactor field value if set, zero value otherwise.
func (o *VehicleType) GetServiceTimeFactor() float64 {
	if o == nil || o.ServiceTimeFactor == nil {
		var ret float64
		return ret
	}
	return *o.ServiceTimeFactor
}

// GetServiceTimeFactorOk returns a tuple with the ServiceTimeFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleType) GetServiceTimeFactorOk() (*float64, bool) {
	if o == nil || o.ServiceTimeFactor == nil {
		return nil, false
	}
	return o.ServiceTimeFactor, true
}

// HasServiceTimeFactor returns a boolean if a field has been set.
func (o *VehicleType) HasServiceTimeFactor() bool {
	if o != nil && o.ServiceTimeFactor != nil {
		return true
	}

	return false
}

// SetServiceTimeFactor gets a reference to the given float64 and assigns it to the ServiceTimeFactor field.
func (o *VehicleType) SetServiceTimeFactor(v float64) {
	o.ServiceTimeFactor = &v
}

// GetCostPerMeter returns the CostPerMeter field value if set, zero value otherwise.
func (o *VehicleType) GetCostPerMeter() float64 {
	if o == nil || o.CostPerMeter == nil {
		var ret float64
		return ret
	}
	return *o.CostPerMeter
}

// GetCostPerMeterOk returns a tuple with the CostPerMeter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleType) GetCostPerMeterOk() (*float64, bool) {
	if o == nil || o.CostPerMeter == nil {
		return nil, false
	}
	return o.CostPerMeter, true
}

// HasCostPerMeter returns a boolean if a field has been set.
func (o *VehicleType) HasCostPerMeter() bool {
	if o != nil && o.CostPerMeter != nil {
		return true
	}

	return false
}

// SetCostPerMeter gets a reference to the given float64 and assigns it to the CostPerMeter field.
func (o *VehicleType) SetCostPerMeter(v float64) {
	o.CostPerMeter = &v
}

// GetCostPerSecond returns the CostPerSecond field value if set, zero value otherwise.
func (o *VehicleType) GetCostPerSecond() float64 {
	if o == nil || o.CostPerSecond == nil {
		var ret float64
		return ret
	}
	return *o.CostPerSecond
}

// GetCostPerSecondOk returns a tuple with the CostPerSecond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleType) GetCostPerSecondOk() (*float64, bool) {
	if o == nil || o.CostPerSecond == nil {
		return nil, false
	}
	return o.CostPerSecond, true
}

// HasCostPerSecond returns a boolean if a field has been set.
func (o *VehicleType) HasCostPerSecond() bool {
	if o != nil && o.CostPerSecond != nil {
		return true
	}

	return false
}

// SetCostPerSecond gets a reference to the given float64 and assigns it to the CostPerSecond field.
func (o *VehicleType) SetCostPerSecond(v float64) {
	o.CostPerSecond = &v
}

// GetCostPerActivation returns the CostPerActivation field value if set, zero value otherwise.
func (o *VehicleType) GetCostPerActivation() float64 {
	if o == nil || o.CostPerActivation == nil {
		var ret float64
		return ret
	}
	return *o.CostPerActivation
}

// GetCostPerActivationOk returns a tuple with the CostPerActivation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleType) GetCostPerActivationOk() (*float64, bool) {
	if o == nil || o.CostPerActivation == nil {
		return nil, false
	}
	return o.CostPerActivation, true
}

// HasCostPerActivation returns a boolean if a field has been set.
func (o *VehicleType) HasCostPerActivation() bool {
	if o != nil && o.CostPerActivation != nil {
		return true
	}

	return false
}

// SetCostPerActivation gets a reference to the given float64 and assigns it to the CostPerActivation field.
func (o *VehicleType) SetCostPerActivation(v float64) {
	o.CostPerActivation = &v
}

// GetConsiderTraffic returns the ConsiderTraffic field value if set, zero value otherwise.
func (o *VehicleType) GetConsiderTraffic() bool {
	if o == nil || o.ConsiderTraffic == nil {
		var ret bool
		return ret
	}
	return *o.ConsiderTraffic
}

// GetConsiderTrafficOk returns a tuple with the ConsiderTraffic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleType) GetConsiderTrafficOk() (*bool, bool) {
	if o == nil || o.ConsiderTraffic == nil {
		return nil, false
	}
	return o.ConsiderTraffic, true
}

// HasConsiderTraffic returns a boolean if a field has been set.
func (o *VehicleType) HasConsiderTraffic() bool {
	if o != nil && o.ConsiderTraffic != nil {
		return true
	}

	return false
}

// SetConsiderTraffic gets a reference to the given bool and assigns it to the ConsiderTraffic field.
func (o *VehicleType) SetConsiderTraffic(v bool) {
	o.ConsiderTraffic = &v
}

// GetNetworkDataProvider returns the NetworkDataProvider field value if set, zero value otherwise.
func (o *VehicleType) GetNetworkDataProvider() string {
	if o == nil || o.NetworkDataProvider == nil {
		var ret string
		return ret
	}
	return *o.NetworkDataProvider
}

// GetNetworkDataProviderOk returns a tuple with the NetworkDataProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleType) GetNetworkDataProviderOk() (*string, bool) {
	if o == nil || o.NetworkDataProvider == nil {
		return nil, false
	}
	return o.NetworkDataProvider, true
}

// HasNetworkDataProvider returns a boolean if a field has been set.
func (o *VehicleType) HasNetworkDataProvider() bool {
	if o != nil && o.NetworkDataProvider != nil {
		return true
	}

	return false
}

// SetNetworkDataProvider gets a reference to the given string and assigns it to the NetworkDataProvider field.
func (o *VehicleType) SetNetworkDataProvider(v string) {
	o.NetworkDataProvider = &v
}

func (o VehicleType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type_id"] = o.TypeId
	}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}
	if o.Capacity != nil {
		toSerialize["capacity"] = o.Capacity
	}
	if o.SpeedFactor != nil {
		toSerialize["speed_factor"] = o.SpeedFactor
	}
	if o.ServiceTimeFactor != nil {
		toSerialize["service_time_factor"] = o.ServiceTimeFactor
	}
	if o.CostPerMeter != nil {
		toSerialize["cost_per_meter"] = o.CostPerMeter
	}
	if o.CostPerSecond != nil {
		toSerialize["cost_per_second"] = o.CostPerSecond
	}
	if o.CostPerActivation != nil {
		toSerialize["cost_per_activation"] = o.CostPerActivation
	}
	if o.ConsiderTraffic != nil {
		toSerialize["consider_traffic"] = o.ConsiderTraffic
	}
	if o.NetworkDataProvider != nil {
		toSerialize["network_data_provider"] = o.NetworkDataProvider
	}
	return json.Marshal(toSerialize)
}

type NullableVehicleType struct {
	value *VehicleType
	isSet bool
}

func (v NullableVehicleType) Get() *VehicleType {
	return v.value
}

func (v *NullableVehicleType) Set(val *VehicleType) {
	v.value = val
	v.isSet = true
}

func (v NullableVehicleType) IsSet() bool {
	return v.isSet
}

func (v *NullableVehicleType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVehicleType(val *VehicleType) *NullableVehicleType {
	return &NullableVehicleType{value: val, isSet: true}
}

func (v NullableVehicleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVehicleType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


