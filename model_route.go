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

// Route struct for Route
type Route struct {
	// Id of vehicle that operates route
	VehicleId *string `json:"vehicle_id,omitempty"`
	// Distance of route in meter
	Distance *int64 `json:"distance,omitempty"`
	// Transport time of route in seconds
	TransportTime *int64 `json:"transport_time,omitempty"`
	// Completion time of route in seconds
	CompletionTime *int64 `json:"completion_time,omitempty"`
	// Waiting time of route in seconds
	WaitingTime *int64 `json:"waiting_time,omitempty"`
	// Service duration of route in seconds
	ServiceDuration *int64 `json:"service_duration,omitempty"`
	// Preparation time of route in seconds
	PreparationTime *int64 `json:"preparation_time,omitempty"`
	// Array of activities
	Activities *[]Activity `json:"activities,omitempty"`
	// Array of route planning points
	Points *[]RoutePoint `json:"points,omitempty"`
}

// NewRoute instantiates a new Route object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoute() *Route {
	this := Route{}
	return &this
}

// NewRouteWithDefaults instantiates a new Route object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteWithDefaults() *Route {
	this := Route{}
	return &this
}

// GetVehicleId returns the VehicleId field value if set, zero value otherwise.
func (o *Route) GetVehicleId() string {
	if o == nil || o.VehicleId == nil {
		var ret string
		return ret
	}
	return *o.VehicleId
}

// GetVehicleIdOk returns a tuple with the VehicleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetVehicleIdOk() (*string, bool) {
	if o == nil || o.VehicleId == nil {
		return nil, false
	}
	return o.VehicleId, true
}

// HasVehicleId returns a boolean if a field has been set.
func (o *Route) HasVehicleId() bool {
	if o != nil && o.VehicleId != nil {
		return true
	}

	return false
}

// SetVehicleId gets a reference to the given string and assigns it to the VehicleId field.
func (o *Route) SetVehicleId(v string) {
	o.VehicleId = &v
}

// GetDistance returns the Distance field value if set, zero value otherwise.
func (o *Route) GetDistance() int64 {
	if o == nil || o.Distance == nil {
		var ret int64
		return ret
	}
	return *o.Distance
}

// GetDistanceOk returns a tuple with the Distance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetDistanceOk() (*int64, bool) {
	if o == nil || o.Distance == nil {
		return nil, false
	}
	return o.Distance, true
}

// HasDistance returns a boolean if a field has been set.
func (o *Route) HasDistance() bool {
	if o != nil && o.Distance != nil {
		return true
	}

	return false
}

// SetDistance gets a reference to the given int64 and assigns it to the Distance field.
func (o *Route) SetDistance(v int64) {
	o.Distance = &v
}

// GetTransportTime returns the TransportTime field value if set, zero value otherwise.
func (o *Route) GetTransportTime() int64 {
	if o == nil || o.TransportTime == nil {
		var ret int64
		return ret
	}
	return *o.TransportTime
}

// GetTransportTimeOk returns a tuple with the TransportTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetTransportTimeOk() (*int64, bool) {
	if o == nil || o.TransportTime == nil {
		return nil, false
	}
	return o.TransportTime, true
}

// HasTransportTime returns a boolean if a field has been set.
func (o *Route) HasTransportTime() bool {
	if o != nil && o.TransportTime != nil {
		return true
	}

	return false
}

// SetTransportTime gets a reference to the given int64 and assigns it to the TransportTime field.
func (o *Route) SetTransportTime(v int64) {
	o.TransportTime = &v
}

// GetCompletionTime returns the CompletionTime field value if set, zero value otherwise.
func (o *Route) GetCompletionTime() int64 {
	if o == nil || o.CompletionTime == nil {
		var ret int64
		return ret
	}
	return *o.CompletionTime
}

// GetCompletionTimeOk returns a tuple with the CompletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetCompletionTimeOk() (*int64, bool) {
	if o == nil || o.CompletionTime == nil {
		return nil, false
	}
	return o.CompletionTime, true
}

// HasCompletionTime returns a boolean if a field has been set.
func (o *Route) HasCompletionTime() bool {
	if o != nil && o.CompletionTime != nil {
		return true
	}

	return false
}

// SetCompletionTime gets a reference to the given int64 and assigns it to the CompletionTime field.
func (o *Route) SetCompletionTime(v int64) {
	o.CompletionTime = &v
}

// GetWaitingTime returns the WaitingTime field value if set, zero value otherwise.
func (o *Route) GetWaitingTime() int64 {
	if o == nil || o.WaitingTime == nil {
		var ret int64
		return ret
	}
	return *o.WaitingTime
}

// GetWaitingTimeOk returns a tuple with the WaitingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetWaitingTimeOk() (*int64, bool) {
	if o == nil || o.WaitingTime == nil {
		return nil, false
	}
	return o.WaitingTime, true
}

// HasWaitingTime returns a boolean if a field has been set.
func (o *Route) HasWaitingTime() bool {
	if o != nil && o.WaitingTime != nil {
		return true
	}

	return false
}

// SetWaitingTime gets a reference to the given int64 and assigns it to the WaitingTime field.
func (o *Route) SetWaitingTime(v int64) {
	o.WaitingTime = &v
}

// GetServiceDuration returns the ServiceDuration field value if set, zero value otherwise.
func (o *Route) GetServiceDuration() int64 {
	if o == nil || o.ServiceDuration == nil {
		var ret int64
		return ret
	}
	return *o.ServiceDuration
}

// GetServiceDurationOk returns a tuple with the ServiceDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetServiceDurationOk() (*int64, bool) {
	if o == nil || o.ServiceDuration == nil {
		return nil, false
	}
	return o.ServiceDuration, true
}

// HasServiceDuration returns a boolean if a field has been set.
func (o *Route) HasServiceDuration() bool {
	if o != nil && o.ServiceDuration != nil {
		return true
	}

	return false
}

// SetServiceDuration gets a reference to the given int64 and assigns it to the ServiceDuration field.
func (o *Route) SetServiceDuration(v int64) {
	o.ServiceDuration = &v
}

// GetPreparationTime returns the PreparationTime field value if set, zero value otherwise.
func (o *Route) GetPreparationTime() int64 {
	if o == nil || o.PreparationTime == nil {
		var ret int64
		return ret
	}
	return *o.PreparationTime
}

// GetPreparationTimeOk returns a tuple with the PreparationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetPreparationTimeOk() (*int64, bool) {
	if o == nil || o.PreparationTime == nil {
		return nil, false
	}
	return o.PreparationTime, true
}

// HasPreparationTime returns a boolean if a field has been set.
func (o *Route) HasPreparationTime() bool {
	if o != nil && o.PreparationTime != nil {
		return true
	}

	return false
}

// SetPreparationTime gets a reference to the given int64 and assigns it to the PreparationTime field.
func (o *Route) SetPreparationTime(v int64) {
	o.PreparationTime = &v
}

// GetActivities returns the Activities field value if set, zero value otherwise.
func (o *Route) GetActivities() []Activity {
	if o == nil || o.Activities == nil {
		var ret []Activity
		return ret
	}
	return *o.Activities
}

// GetActivitiesOk returns a tuple with the Activities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetActivitiesOk() (*[]Activity, bool) {
	if o == nil || o.Activities == nil {
		return nil, false
	}
	return o.Activities, true
}

// HasActivities returns a boolean if a field has been set.
func (o *Route) HasActivities() bool {
	if o != nil && o.Activities != nil {
		return true
	}

	return false
}

// SetActivities gets a reference to the given []Activity and assigns it to the Activities field.
func (o *Route) SetActivities(v []Activity) {
	o.Activities = &v
}

// GetPoints returns the Points field value if set, zero value otherwise.
func (o *Route) GetPoints() []RoutePoint {
	if o == nil || o.Points == nil {
		var ret []RoutePoint
		return ret
	}
	return *o.Points
}

// GetPointsOk returns a tuple with the Points field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetPointsOk() (*[]RoutePoint, bool) {
	if o == nil || o.Points == nil {
		return nil, false
	}
	return o.Points, true
}

// HasPoints returns a boolean if a field has been set.
func (o *Route) HasPoints() bool {
	if o != nil && o.Points != nil {
		return true
	}

	return false
}

// SetPoints gets a reference to the given []RoutePoint and assigns it to the Points field.
func (o *Route) SetPoints(v []RoutePoint) {
	o.Points = &v
}

func (o Route) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.VehicleId != nil {
		toSerialize["vehicle_id"] = o.VehicleId
	}
	if o.Distance != nil {
		toSerialize["distance"] = o.Distance
	}
	if o.TransportTime != nil {
		toSerialize["transport_time"] = o.TransportTime
	}
	if o.CompletionTime != nil {
		toSerialize["completion_time"] = o.CompletionTime
	}
	if o.WaitingTime != nil {
		toSerialize["waiting_time"] = o.WaitingTime
	}
	if o.ServiceDuration != nil {
		toSerialize["service_duration"] = o.ServiceDuration
	}
	if o.PreparationTime != nil {
		toSerialize["preparation_time"] = o.PreparationTime
	}
	if o.Activities != nil {
		toSerialize["activities"] = o.Activities
	}
	if o.Points != nil {
		toSerialize["points"] = o.Points
	}
	return json.Marshal(toSerialize)
}

type NullableRoute struct {
	value *Route
	isSet bool
}

func (v NullableRoute) Get() *Route {
	return v.value
}

func (v *NullableRoute) Set(val *Route) {
	v.value = val
	v.isSet = true
}

func (v NullableRoute) IsSet() bool {
	return v.isSet
}

func (v *NullableRoute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoute(val *Route) *NullableRoute {
	return &NullableRoute{value: val, isSet: true}
}

func (v NullableRoute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


