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

// JobRelation struct for JobRelation
type JobRelation struct {
	// Specifies the type of relation. It must be either of type `in_same_route`, `in_sequence` or `in_direct_sequence`.  `in_same_route`: As the name suggest, it enforces the specified services or shipments to be in the same route. It can be specified as follows:  ```json {    \"type\": \"in_same_route\",    \"ids\": [\"serv_i_id\",\"serv_j_id\"] } ```  This enforces service i to be in the same route as service j no matter which vehicle will be employed. If a specific vehicle (driver) is required to conduct this, just add a `vehicle_id` like this:  ``` {    \"type\": \"in_same_route\",    \"ids\": [\"serv_i_id\",\"serv_j_id\"],    \"vehicle_id\": \"vehicle1\" } ```  This not only enforce service i and j to be in the same route, but also makes sure that both services are in the route of `vehicle1`.  *Tip*: This way initial loads and vehicle routes can be modelled. For example, if your vehicles are already on the road and new orders come in, then vehicles can still be rescheduled subject to the orders that have already been assigned to these vehicles.    `in_sequence`: This relation type enforces n jobs to be in sequence. It can be specified as  ```json {    \"type\": \"in_sequence\",    \"ids\": [\"serv_i_id\",\"serv_j_id\"] } ```  which means that service j need to be in the same route as service i AND it needs to occur somewhere after service i. As described above if a specific vehicle needs to conduct this, just add `vehicle_id`.   `in_direct_sequence`: This enforces n services or shipments to be in direct sequence. It can be specified as  ```json {    \"type\": \"in_direct_sequence\",    \"ids\": [\"serv_i_id\",\"serv_j_id\",\"serv_k_id\"] } ```  yielding service j to occur directly after service i, and service k to occur directly after service j i.e. in strong order. Again, a vehicle can be assigned a priority by adding a `vehicle_id` to the relation.   *Special IDs*: If you look at the previous example and you want service i to be the first in the route, use the special ID `start` as follows:  ```json {    \"type\": \"in_direct_sequence\",    \"ids\": [\"start\",\"serv_i_id\",\"serv_j_id\",\"serv_k_id\"] } ```  Latter enforces the direct sequence of i, j and k at the beginning of the route. If this sequence should be bound to the end of the route, use the special ID `end` like this:  ```json {    \"type\": \"in_direct_sequence\",    \"ids\": [\"serv_i_id\",\"service_j_id\",\"serv_k_id\",\"end\"] } ```  If you deal with services then you need to use the 'id' of your services in the field 'ids'. To also consider sequences of the pickups and deliveries of your shipments, you need to use a special ID, i.e. use the shipment id plus the keyword `_pickup` or `_delivery`. For example, to ensure that the pickup and delivery of the shipment with the id 'my_shipment' are direct neighbors, you need the following specification:  ``` {    \"type\": \"in_direct_sequence\",    \"ids\": [\"my_ship_pickup\",\"my_ship_delivery\"] } ```  
	Type string `json:"type"`
	// Specifies an array of shipment and/or service ids that are in relation. If you deal with services then you need to use the id of your services in ids. To also consider sequences of the pickups and deliveries of your shipments, you need to use a special ID, i.e. use your shipment id plus the keyword `_pickup` or `_delivery`. If you want to place a service or shipment activity at the beginning of your route, use the special ID `start`. In turn, use `end` to place it at the end of the route.
	Ids []string `json:"ids"`
	// Id of pre-assigned vehicle, i.e. the vehicle id that is determined to conduct the services and shipments in this relation.
	VehicleId *string `json:"vehicle_id,omitempty"`
}

// NewJobRelation instantiates a new JobRelation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobRelation(type_ string, ids []string) *JobRelation {
	this := JobRelation{}
	this.Type = type_
	this.Ids = ids
	return &this
}

// NewJobRelationWithDefaults instantiates a new JobRelation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobRelationWithDefaults() *JobRelation {
	this := JobRelation{}
	return &this
}

// GetType returns the Type field value
func (o *JobRelation) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *JobRelation) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *JobRelation) SetType(v string) {
	o.Type = v
}

// GetIds returns the Ids field value
func (o *JobRelation) GetIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value
// and a boolean to check if the value has been set.
func (o *JobRelation) GetIdsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Ids, true
}

// SetIds sets field value
func (o *JobRelation) SetIds(v []string) {
	o.Ids = v
}

// GetVehicleId returns the VehicleId field value if set, zero value otherwise.
func (o *JobRelation) GetVehicleId() string {
	if o == nil || o.VehicleId == nil {
		var ret string
		return ret
	}
	return *o.VehicleId
}

// GetVehicleIdOk returns a tuple with the VehicleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobRelation) GetVehicleIdOk() (*string, bool) {
	if o == nil || o.VehicleId == nil {
		return nil, false
	}
	return o.VehicleId, true
}

// HasVehicleId returns a boolean if a field has been set.
func (o *JobRelation) HasVehicleId() bool {
	if o != nil && o.VehicleId != nil {
		return true
	}

	return false
}

// SetVehicleId gets a reference to the given string and assigns it to the VehicleId field.
func (o *JobRelation) SetVehicleId(v string) {
	o.VehicleId = &v
}

func (o JobRelation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["ids"] = o.Ids
	}
	if o.VehicleId != nil {
		toSerialize["vehicle_id"] = o.VehicleId
	}
	return json.Marshal(toSerialize)
}

type NullableJobRelation struct {
	value *JobRelation
	isSet bool
}

func (v NullableJobRelation) Get() *JobRelation {
	return v.value
}

func (v *NullableJobRelation) Set(val *JobRelation) {
	v.value = val
	v.isSet = true
}

func (v NullableJobRelation) IsSet() bool {
	return v.isSet
}

func (v *NullableJobRelation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobRelation(val *JobRelation) *NullableJobRelation {
	return &NullableJobRelation{value: val, isSet: true}
}

func (v NullableJobRelation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobRelation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


