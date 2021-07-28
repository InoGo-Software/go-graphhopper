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

// Shipment struct for Shipment
type Shipment struct {
	// Specifies the id of the shipment. Ids need to be unique so there must not be two services/shipments with the same id.
	Id string `json:"id"`
	// Meaningful name for shipment, e.g. \"pickup and deliver pizza to Peter\".
	Name *string `json:"name,omitempty"`
	// Specifies the priority. Can be 1 = high priority to 10 = low priority. Often there are more services/shipments than the available vehicle fleet can handle. Then you can set priorities to differentiate high priority tasks from those that could be left unassigned. I.e. the lower the priority the earlier these tasks are omitted in the solution.
	Priority *int32 `json:"priority,omitempty"`
	Pickup Stop `json:"pickup"`
	Delivery Stop `json:"delivery"`
	// Size can have multiple dimensions and should be in line with the capacity dimension array of the vehicle type. For example, if the item that needs to be delivered has two size dimension, volume and weight, then specify it as follow [ 20, 5 ] assuming a volume of 20 and a weight of 5.
	Size *[]int32 `json:"size,omitempty"`
	// Specifies an array of required skills, i.e. array of string (not case sensitive). For example, if this shipment needs to be conducted by a technician having a `drilling_machine` and a `screw_driver` then specify the array as follows: `[\"drilling_machine\",\"screw_driver\"]`. This means that the service can only be done by a vehicle (technician) that has the skills `drilling_machine` AND `screw_driver` in its skill array. Otherwise it remains unassigned.
	RequiredSkills *[]string `json:"required_skills,omitempty"`
	// Specifies an array of allowed vehicles, i.e. array of vehicle ids. For example, if this shipment can only be conducted EITHER by \"technician_peter\" OR \"technician_stefan\" specify this as follows: [\"technician_peter\",\"technician_stefan\"].
	AllowedVehicles *[]string `json:"allowed_vehicles,omitempty"`
	// Specifies an array of disallowed vehicles, i.e. array of vehicle ids.
	DisallowedVehicles *[]string `json:"disallowed_vehicles,omitempty"`
	// Specifies the maximum time in seconds a shipment can stay in the vehicle.
	MaxTimeInVehicle *int64 `json:"max_time_in_vehicle,omitempty"`
}

// NewShipment instantiates a new Shipment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipment(id string, pickup Stop, delivery Stop) *Shipment {
	this := Shipment{}
	this.Id = id
	var priority int32 = 2
	this.Priority = &priority
	this.Pickup = pickup
	this.Delivery = delivery
	return &this
}

// NewShipmentWithDefaults instantiates a new Shipment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentWithDefaults() *Shipment {
	this := Shipment{}
	var priority int32 = 2
	this.Priority = &priority
	return &this
}

// GetId returns the Id field value
func (o *Shipment) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Shipment) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Shipment) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Shipment) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Shipment) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Shipment) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Shipment) SetName(v string) {
	o.Name = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *Shipment) GetPriority() int32 {
	if o == nil || o.Priority == nil {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Shipment) GetPriorityOk() (*int32, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *Shipment) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *Shipment) SetPriority(v int32) {
	o.Priority = &v
}

// GetPickup returns the Pickup field value
func (o *Shipment) GetPickup() Stop {
	if o == nil {
		var ret Stop
		return ret
	}

	return o.Pickup
}

// GetPickupOk returns a tuple with the Pickup field value
// and a boolean to check if the value has been set.
func (o *Shipment) GetPickupOk() (*Stop, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Pickup, true
}

// SetPickup sets field value
func (o *Shipment) SetPickup(v Stop) {
	o.Pickup = v
}

// GetDelivery returns the Delivery field value
func (o *Shipment) GetDelivery() Stop {
	if o == nil {
		var ret Stop
		return ret
	}

	return o.Delivery
}

// GetDeliveryOk returns a tuple with the Delivery field value
// and a boolean to check if the value has been set.
func (o *Shipment) GetDeliveryOk() (*Stop, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Delivery, true
}

// SetDelivery sets field value
func (o *Shipment) SetDelivery(v Stop) {
	o.Delivery = v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *Shipment) GetSize() []int32 {
	if o == nil || o.Size == nil {
		var ret []int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Shipment) GetSizeOk() (*[]int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *Shipment) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given []int32 and assigns it to the Size field.
func (o *Shipment) SetSize(v []int32) {
	o.Size = &v
}

// GetRequiredSkills returns the RequiredSkills field value if set, zero value otherwise.
func (o *Shipment) GetRequiredSkills() []string {
	if o == nil || o.RequiredSkills == nil {
		var ret []string
		return ret
	}
	return *o.RequiredSkills
}

// GetRequiredSkillsOk returns a tuple with the RequiredSkills field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Shipment) GetRequiredSkillsOk() (*[]string, bool) {
	if o == nil || o.RequiredSkills == nil {
		return nil, false
	}
	return o.RequiredSkills, true
}

// HasRequiredSkills returns a boolean if a field has been set.
func (o *Shipment) HasRequiredSkills() bool {
	if o != nil && o.RequiredSkills != nil {
		return true
	}

	return false
}

// SetRequiredSkills gets a reference to the given []string and assigns it to the RequiredSkills field.
func (o *Shipment) SetRequiredSkills(v []string) {
	o.RequiredSkills = &v
}

// GetAllowedVehicles returns the AllowedVehicles field value if set, zero value otherwise.
func (o *Shipment) GetAllowedVehicles() []string {
	if o == nil || o.AllowedVehicles == nil {
		var ret []string
		return ret
	}
	return *o.AllowedVehicles
}

// GetAllowedVehiclesOk returns a tuple with the AllowedVehicles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Shipment) GetAllowedVehiclesOk() (*[]string, bool) {
	if o == nil || o.AllowedVehicles == nil {
		return nil, false
	}
	return o.AllowedVehicles, true
}

// HasAllowedVehicles returns a boolean if a field has been set.
func (o *Shipment) HasAllowedVehicles() bool {
	if o != nil && o.AllowedVehicles != nil {
		return true
	}

	return false
}

// SetAllowedVehicles gets a reference to the given []string and assigns it to the AllowedVehicles field.
func (o *Shipment) SetAllowedVehicles(v []string) {
	o.AllowedVehicles = &v
}

// GetDisallowedVehicles returns the DisallowedVehicles field value if set, zero value otherwise.
func (o *Shipment) GetDisallowedVehicles() []string {
	if o == nil || o.DisallowedVehicles == nil {
		var ret []string
		return ret
	}
	return *o.DisallowedVehicles
}

// GetDisallowedVehiclesOk returns a tuple with the DisallowedVehicles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Shipment) GetDisallowedVehiclesOk() (*[]string, bool) {
	if o == nil || o.DisallowedVehicles == nil {
		return nil, false
	}
	return o.DisallowedVehicles, true
}

// HasDisallowedVehicles returns a boolean if a field has been set.
func (o *Shipment) HasDisallowedVehicles() bool {
	if o != nil && o.DisallowedVehicles != nil {
		return true
	}

	return false
}

// SetDisallowedVehicles gets a reference to the given []string and assigns it to the DisallowedVehicles field.
func (o *Shipment) SetDisallowedVehicles(v []string) {
	o.DisallowedVehicles = &v
}

// GetMaxTimeInVehicle returns the MaxTimeInVehicle field value if set, zero value otherwise.
func (o *Shipment) GetMaxTimeInVehicle() int64 {
	if o == nil || o.MaxTimeInVehicle == nil {
		var ret int64
		return ret
	}
	return *o.MaxTimeInVehicle
}

// GetMaxTimeInVehicleOk returns a tuple with the MaxTimeInVehicle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Shipment) GetMaxTimeInVehicleOk() (*int64, bool) {
	if o == nil || o.MaxTimeInVehicle == nil {
		return nil, false
	}
	return o.MaxTimeInVehicle, true
}

// HasMaxTimeInVehicle returns a boolean if a field has been set.
func (o *Shipment) HasMaxTimeInVehicle() bool {
	if o != nil && o.MaxTimeInVehicle != nil {
		return true
	}

	return false
}

// SetMaxTimeInVehicle gets a reference to the given int64 and assigns it to the MaxTimeInVehicle field.
func (o *Shipment) SetMaxTimeInVehicle(v int64) {
	o.MaxTimeInVehicle = &v
}

func (o Shipment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if true {
		toSerialize["pickup"] = o.Pickup
	}
	if true {
		toSerialize["delivery"] = o.Delivery
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.RequiredSkills != nil {
		toSerialize["required_skills"] = o.RequiredSkills
	}
	if o.AllowedVehicles != nil {
		toSerialize["allowed_vehicles"] = o.AllowedVehicles
	}
	if o.DisallowedVehicles != nil {
		toSerialize["disallowed_vehicles"] = o.DisallowedVehicles
	}
	if o.MaxTimeInVehicle != nil {
		toSerialize["max_time_in_vehicle"] = o.MaxTimeInVehicle
	}
	return json.Marshal(toSerialize)
}

type NullableShipment struct {
	value *Shipment
	isSet bool
}

func (v NullableShipment) Get() *Shipment {
	return v.value
}

func (v *NullableShipment) Set(val *Shipment) {
	v.value = val
	v.isSet = true
}

func (v NullableShipment) IsSet() bool {
	return v.isSet
}

func (v *NullableShipment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipment(val *Shipment) *NullableShipment {
	return &NullableShipment{value: val, isSet: true}
}

func (v NullableShipment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


