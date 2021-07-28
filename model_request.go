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

// Request struct for Request
type Request struct {
	// Specifies the available vehicles.
	Vehicles *[]Vehicle `json:"vehicles,omitempty"`
	// Specifies the available vehicle types. These types can be assigned to vehicles.
	VehicleTypes *[]VehicleType `json:"vehicle_types,omitempty"`
	// Specifies the orders of the type \"service\". These are, for example, pick-ups, deliveries or other stops that are to be approached by the specified vehicles. Each of these orders contains only one location.
	Services *[]Service `json:"services,omitempty"`
	// Specifies the available shipments. Each shipment contains a pickup and a delivery stop, which must be processed one after the other.
	Shipments *[]Shipment `json:"shipments,omitempty"`
	// Defines additional relationships between orders.
	Relations *[]AnyOfJobRelationGroupRelation `json:"relations,omitempty"`
	Algorithm *Algorithm `json:"algorithm,omitempty"`
	// Specifies an objective function. The vehicle routing problem is solved in such a way that this objective function is minimized.
	Objectives *[]Objective `json:"objectives,omitempty"`
	// Specifies your own tranport time and distance matrices.
	CostMatrices *[]CostMatrix `json:"cost_matrices,omitempty"`
	Configuration *Configuration `json:"configuration,omitempty"`
}

// NewRequest instantiates a new Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequest() *Request {
	this := Request{}
	return &this
}

// NewRequestWithDefaults instantiates a new Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestWithDefaults() *Request {
	this := Request{}
	return &this
}

// GetVehicles returns the Vehicles field value if set, zero value otherwise.
func (o *Request) GetVehicles() []Vehicle {
	if o == nil || o.Vehicles == nil {
		var ret []Vehicle
		return ret
	}
	return *o.Vehicles
}

// GetVehiclesOk returns a tuple with the Vehicles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Request) GetVehiclesOk() (*[]Vehicle, bool) {
	if o == nil || o.Vehicles == nil {
		return nil, false
	}
	return o.Vehicles, true
}

// HasVehicles returns a boolean if a field has been set.
func (o *Request) HasVehicles() bool {
	if o != nil && o.Vehicles != nil {
		return true
	}

	return false
}

// SetVehicles gets a reference to the given []Vehicle and assigns it to the Vehicles field.
func (o *Request) SetVehicles(v []Vehicle) {
	o.Vehicles = &v
}

// GetVehicleTypes returns the VehicleTypes field value if set, zero value otherwise.
func (o *Request) GetVehicleTypes() []VehicleType {
	if o == nil || o.VehicleTypes == nil {
		var ret []VehicleType
		return ret
	}
	return *o.VehicleTypes
}

// GetVehicleTypesOk returns a tuple with the VehicleTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Request) GetVehicleTypesOk() (*[]VehicleType, bool) {
	if o == nil || o.VehicleTypes == nil {
		return nil, false
	}
	return o.VehicleTypes, true
}

// HasVehicleTypes returns a boolean if a field has been set.
func (o *Request) HasVehicleTypes() bool {
	if o != nil && o.VehicleTypes != nil {
		return true
	}

	return false
}

// SetVehicleTypes gets a reference to the given []VehicleType and assigns it to the VehicleTypes field.
func (o *Request) SetVehicleTypes(v []VehicleType) {
	o.VehicleTypes = &v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *Request) GetServices() []Service {
	if o == nil || o.Services == nil {
		var ret []Service
		return ret
	}
	return *o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Request) GetServicesOk() (*[]Service, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *Request) HasServices() bool {
	if o != nil && o.Services != nil {
		return true
	}

	return false
}

// SetServices gets a reference to the given []Service and assigns it to the Services field.
func (o *Request) SetServices(v []Service) {
	o.Services = &v
}

// GetShipments returns the Shipments field value if set, zero value otherwise.
func (o *Request) GetShipments() []Shipment {
	if o == nil || o.Shipments == nil {
		var ret []Shipment
		return ret
	}
	return *o.Shipments
}

// GetShipmentsOk returns a tuple with the Shipments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Request) GetShipmentsOk() (*[]Shipment, bool) {
	if o == nil || o.Shipments == nil {
		return nil, false
	}
	return o.Shipments, true
}

// HasShipments returns a boolean if a field has been set.
func (o *Request) HasShipments() bool {
	if o != nil && o.Shipments != nil {
		return true
	}

	return false
}

// SetShipments gets a reference to the given []Shipment and assigns it to the Shipments field.
func (o *Request) SetShipments(v []Shipment) {
	o.Shipments = &v
}

// GetRelations returns the Relations field value if set, zero value otherwise.
func (o *Request) GetRelations() []AnyOfJobRelationGroupRelation {
	if o == nil || o.Relations == nil {
		var ret []AnyOfJobRelationGroupRelation
		return ret
	}
	return *o.Relations
}

// GetRelationsOk returns a tuple with the Relations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Request) GetRelationsOk() (*[]AnyOfJobRelationGroupRelation, bool) {
	if o == nil || o.Relations == nil {
		return nil, false
	}
	return o.Relations, true
}

// HasRelations returns a boolean if a field has been set.
func (o *Request) HasRelations() bool {
	if o != nil && o.Relations != nil {
		return true
	}

	return false
}

// SetRelations gets a reference to the given []AnyOfJobRelationGroupRelation and assigns it to the Relations field.
func (o *Request) SetRelations(v []AnyOfJobRelationGroupRelation) {
	o.Relations = &v
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise.
func (o *Request) GetAlgorithm() Algorithm {
	if o == nil || o.Algorithm == nil {
		var ret Algorithm
		return ret
	}
	return *o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Request) GetAlgorithmOk() (*Algorithm, bool) {
	if o == nil || o.Algorithm == nil {
		return nil, false
	}
	return o.Algorithm, true
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *Request) HasAlgorithm() bool {
	if o != nil && o.Algorithm != nil {
		return true
	}

	return false
}

// SetAlgorithm gets a reference to the given Algorithm and assigns it to the Algorithm field.
func (o *Request) SetAlgorithm(v Algorithm) {
	o.Algorithm = &v
}

// GetObjectives returns the Objectives field value if set, zero value otherwise.
func (o *Request) GetObjectives() []Objective {
	if o == nil || o.Objectives == nil {
		var ret []Objective
		return ret
	}
	return *o.Objectives
}

// GetObjectivesOk returns a tuple with the Objectives field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Request) GetObjectivesOk() (*[]Objective, bool) {
	if o == nil || o.Objectives == nil {
		return nil, false
	}
	return o.Objectives, true
}

// HasObjectives returns a boolean if a field has been set.
func (o *Request) HasObjectives() bool {
	if o != nil && o.Objectives != nil {
		return true
	}

	return false
}

// SetObjectives gets a reference to the given []Objective and assigns it to the Objectives field.
func (o *Request) SetObjectives(v []Objective) {
	o.Objectives = &v
}

// GetCostMatrices returns the CostMatrices field value if set, zero value otherwise.
func (o *Request) GetCostMatrices() []CostMatrix {
	if o == nil || o.CostMatrices == nil {
		var ret []CostMatrix
		return ret
	}
	return *o.CostMatrices
}

// GetCostMatricesOk returns a tuple with the CostMatrices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Request) GetCostMatricesOk() (*[]CostMatrix, bool) {
	if o == nil || o.CostMatrices == nil {
		return nil, false
	}
	return o.CostMatrices, true
}

// HasCostMatrices returns a boolean if a field has been set.
func (o *Request) HasCostMatrices() bool {
	if o != nil && o.CostMatrices != nil {
		return true
	}

	return false
}

// SetCostMatrices gets a reference to the given []CostMatrix and assigns it to the CostMatrices field.
func (o *Request) SetCostMatrices(v []CostMatrix) {
	o.CostMatrices = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *Request) GetConfiguration() Configuration {
	if o == nil || o.Configuration == nil {
		var ret Configuration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Request) GetConfigurationOk() (*Configuration, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *Request) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given Configuration and assigns it to the Configuration field.
func (o *Request) SetConfiguration(v Configuration) {
	o.Configuration = &v
}

func (o Request) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Vehicles != nil {
		toSerialize["vehicles"] = o.Vehicles
	}
	if o.VehicleTypes != nil {
		toSerialize["vehicle_types"] = o.VehicleTypes
	}
	if o.Services != nil {
		toSerialize["services"] = o.Services
	}
	if o.Shipments != nil {
		toSerialize["shipments"] = o.Shipments
	}
	if o.Relations != nil {
		toSerialize["relations"] = o.Relations
	}
	if o.Algorithm != nil {
		toSerialize["algorithm"] = o.Algorithm
	}
	if o.Objectives != nil {
		toSerialize["objectives"] = o.Objectives
	}
	if o.CostMatrices != nil {
		toSerialize["cost_matrices"] = o.CostMatrices
	}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}
	return json.Marshal(toSerialize)
}

type NullableRequest struct {
	value *Request
	isSet bool
}

func (v NullableRequest) Get() *Request {
	return v.value
}

func (v *NullableRequest) Set(val *Request) {
	v.value = val
	v.isSet = true
}

func (v NullableRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequest(val *Request) *NullableRequest {
	return &NullableRequest{value: val, isSet: true}
}

func (v NullableRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


