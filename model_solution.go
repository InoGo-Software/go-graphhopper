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

// Solution Only available if status field indicates `finished`.
type Solution struct {
	Costs *int32 `json:"costs,omitempty"`
	// Overall distance travelled in meter, i.e. the sum of each route's transport distance
	Distance *int32 `json:"distance,omitempty"`
	// Use `transport_time` instead.
	Time *int64 `json:"time,omitempty"`
	// Overall time travelled in seconds, i.e. the sum of each route's transport time.
	TransportTime *int64 `json:"transport_time,omitempty"`
	// Operation time of longest route in seconds.
	MaxOperationTime *int64 `json:"max_operation_time,omitempty"`
	// Overall waiting time in seconds.
	WaitingTime *int64 `json:"waiting_time,omitempty"`
	// Overall service time in seconds.
	ServiceDuration *int64 `json:"service_duration,omitempty"`
	// Overall preparation time in seconds.
	PreparationTime *int64 `json:"preparation_time,omitempty"`
	// Overall completion time in seconds, i.e. the sum of each routes/drivers operation time.
	CompletionTime *int64 `json:"completion_time,omitempty"`
	// Number of employed vehicles.
	NoVehicles *int32 `json:"no_vehicles,omitempty"`
	// Number of jobs that could not be assigned to final solution.
	NoUnassigned *int32 `json:"no_unassigned,omitempty"`
	// An array of routes
	Routes *[]Route `json:"routes,omitempty"`
	Unassigned *SolutionUnassigned `json:"unassigned,omitempty"`
}

// NewSolution instantiates a new Solution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSolution() *Solution {
	this := Solution{}
	return &this
}

// NewSolutionWithDefaults instantiates a new Solution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSolutionWithDefaults() *Solution {
	this := Solution{}
	return &this
}

// GetCosts returns the Costs field value if set, zero value otherwise.
func (o *Solution) GetCosts() int32 {
	if o == nil || o.Costs == nil {
		var ret int32
		return ret
	}
	return *o.Costs
}

// GetCostsOk returns a tuple with the Costs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Solution) GetCostsOk() (*int32, bool) {
	if o == nil || o.Costs == nil {
		return nil, false
	}
	return o.Costs, true
}

// HasCosts returns a boolean if a field has been set.
func (o *Solution) HasCosts() bool {
	if o != nil && o.Costs != nil {
		return true
	}

	return false
}

// SetCosts gets a reference to the given int32 and assigns it to the Costs field.
func (o *Solution) SetCosts(v int32) {
	o.Costs = &v
}

// GetDistance returns the Distance field value if set, zero value otherwise.
func (o *Solution) GetDistance() int32 {
	if o == nil || o.Distance == nil {
		var ret int32
		return ret
	}
	return *o.Distance
}

// GetDistanceOk returns a tuple with the Distance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Solution) GetDistanceOk() (*int32, bool) {
	if o == nil || o.Distance == nil {
		return nil, false
	}
	return o.Distance, true
}

// HasDistance returns a boolean if a field has been set.
func (o *Solution) HasDistance() bool {
	if o != nil && o.Distance != nil {
		return true
	}

	return false
}

// SetDistance gets a reference to the given int32 and assigns it to the Distance field.
func (o *Solution) SetDistance(v int32) {
	o.Distance = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *Solution) GetTime() int64 {
	if o == nil || o.Time == nil {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Solution) GetTimeOk() (*int64, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *Solution) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *Solution) SetTime(v int64) {
	o.Time = &v
}

// GetTransportTime returns the TransportTime field value if set, zero value otherwise.
func (o *Solution) GetTransportTime() int64 {
	if o == nil || o.TransportTime == nil {
		var ret int64
		return ret
	}
	return *o.TransportTime
}

// GetTransportTimeOk returns a tuple with the TransportTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Solution) GetTransportTimeOk() (*int64, bool) {
	if o == nil || o.TransportTime == nil {
		return nil, false
	}
	return o.TransportTime, true
}

// HasTransportTime returns a boolean if a field has been set.
func (o *Solution) HasTransportTime() bool {
	if o != nil && o.TransportTime != nil {
		return true
	}

	return false
}

// SetTransportTime gets a reference to the given int64 and assigns it to the TransportTime field.
func (o *Solution) SetTransportTime(v int64) {
	o.TransportTime = &v
}

// GetMaxOperationTime returns the MaxOperationTime field value if set, zero value otherwise.
func (o *Solution) GetMaxOperationTime() int64 {
	if o == nil || o.MaxOperationTime == nil {
		var ret int64
		return ret
	}
	return *o.MaxOperationTime
}

// GetMaxOperationTimeOk returns a tuple with the MaxOperationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Solution) GetMaxOperationTimeOk() (*int64, bool) {
	if o == nil || o.MaxOperationTime == nil {
		return nil, false
	}
	return o.MaxOperationTime, true
}

// HasMaxOperationTime returns a boolean if a field has been set.
func (o *Solution) HasMaxOperationTime() bool {
	if o != nil && o.MaxOperationTime != nil {
		return true
	}

	return false
}

// SetMaxOperationTime gets a reference to the given int64 and assigns it to the MaxOperationTime field.
func (o *Solution) SetMaxOperationTime(v int64) {
	o.MaxOperationTime = &v
}

// GetWaitingTime returns the WaitingTime field value if set, zero value otherwise.
func (o *Solution) GetWaitingTime() int64 {
	if o == nil || o.WaitingTime == nil {
		var ret int64
		return ret
	}
	return *o.WaitingTime
}

// GetWaitingTimeOk returns a tuple with the WaitingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Solution) GetWaitingTimeOk() (*int64, bool) {
	if o == nil || o.WaitingTime == nil {
		return nil, false
	}
	return o.WaitingTime, true
}

// HasWaitingTime returns a boolean if a field has been set.
func (o *Solution) HasWaitingTime() bool {
	if o != nil && o.WaitingTime != nil {
		return true
	}

	return false
}

// SetWaitingTime gets a reference to the given int64 and assigns it to the WaitingTime field.
func (o *Solution) SetWaitingTime(v int64) {
	o.WaitingTime = &v
}

// GetServiceDuration returns the ServiceDuration field value if set, zero value otherwise.
func (o *Solution) GetServiceDuration() int64 {
	if o == nil || o.ServiceDuration == nil {
		var ret int64
		return ret
	}
	return *o.ServiceDuration
}

// GetServiceDurationOk returns a tuple with the ServiceDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Solution) GetServiceDurationOk() (*int64, bool) {
	if o == nil || o.ServiceDuration == nil {
		return nil, false
	}
	return o.ServiceDuration, true
}

// HasServiceDuration returns a boolean if a field has been set.
func (o *Solution) HasServiceDuration() bool {
	if o != nil && o.ServiceDuration != nil {
		return true
	}

	return false
}

// SetServiceDuration gets a reference to the given int64 and assigns it to the ServiceDuration field.
func (o *Solution) SetServiceDuration(v int64) {
	o.ServiceDuration = &v
}

// GetPreparationTime returns the PreparationTime field value if set, zero value otherwise.
func (o *Solution) GetPreparationTime() int64 {
	if o == nil || o.PreparationTime == nil {
		var ret int64
		return ret
	}
	return *o.PreparationTime
}

// GetPreparationTimeOk returns a tuple with the PreparationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Solution) GetPreparationTimeOk() (*int64, bool) {
	if o == nil || o.PreparationTime == nil {
		return nil, false
	}
	return o.PreparationTime, true
}

// HasPreparationTime returns a boolean if a field has been set.
func (o *Solution) HasPreparationTime() bool {
	if o != nil && o.PreparationTime != nil {
		return true
	}

	return false
}

// SetPreparationTime gets a reference to the given int64 and assigns it to the PreparationTime field.
func (o *Solution) SetPreparationTime(v int64) {
	o.PreparationTime = &v
}

// GetCompletionTime returns the CompletionTime field value if set, zero value otherwise.
func (o *Solution) GetCompletionTime() int64 {
	if o == nil || o.CompletionTime == nil {
		var ret int64
		return ret
	}
	return *o.CompletionTime
}

// GetCompletionTimeOk returns a tuple with the CompletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Solution) GetCompletionTimeOk() (*int64, bool) {
	if o == nil || o.CompletionTime == nil {
		return nil, false
	}
	return o.CompletionTime, true
}

// HasCompletionTime returns a boolean if a field has been set.
func (o *Solution) HasCompletionTime() bool {
	if o != nil && o.CompletionTime != nil {
		return true
	}

	return false
}

// SetCompletionTime gets a reference to the given int64 and assigns it to the CompletionTime field.
func (o *Solution) SetCompletionTime(v int64) {
	o.CompletionTime = &v
}

// GetNoVehicles returns the NoVehicles field value if set, zero value otherwise.
func (o *Solution) GetNoVehicles() int32 {
	if o == nil || o.NoVehicles == nil {
		var ret int32
		return ret
	}
	return *o.NoVehicles
}

// GetNoVehiclesOk returns a tuple with the NoVehicles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Solution) GetNoVehiclesOk() (*int32, bool) {
	if o == nil || o.NoVehicles == nil {
		return nil, false
	}
	return o.NoVehicles, true
}

// HasNoVehicles returns a boolean if a field has been set.
func (o *Solution) HasNoVehicles() bool {
	if o != nil && o.NoVehicles != nil {
		return true
	}

	return false
}

// SetNoVehicles gets a reference to the given int32 and assigns it to the NoVehicles field.
func (o *Solution) SetNoVehicles(v int32) {
	o.NoVehicles = &v
}

// GetNoUnassigned returns the NoUnassigned field value if set, zero value otherwise.
func (o *Solution) GetNoUnassigned() int32 {
	if o == nil || o.NoUnassigned == nil {
		var ret int32
		return ret
	}
	return *o.NoUnassigned
}

// GetNoUnassignedOk returns a tuple with the NoUnassigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Solution) GetNoUnassignedOk() (*int32, bool) {
	if o == nil || o.NoUnassigned == nil {
		return nil, false
	}
	return o.NoUnassigned, true
}

// HasNoUnassigned returns a boolean if a field has been set.
func (o *Solution) HasNoUnassigned() bool {
	if o != nil && o.NoUnassigned != nil {
		return true
	}

	return false
}

// SetNoUnassigned gets a reference to the given int32 and assigns it to the NoUnassigned field.
func (o *Solution) SetNoUnassigned(v int32) {
	o.NoUnassigned = &v
}

// GetRoutes returns the Routes field value if set, zero value otherwise.
func (o *Solution) GetRoutes() []Route {
	if o == nil || o.Routes == nil {
		var ret []Route
		return ret
	}
	return *o.Routes
}

// GetRoutesOk returns a tuple with the Routes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Solution) GetRoutesOk() (*[]Route, bool) {
	if o == nil || o.Routes == nil {
		return nil, false
	}
	return o.Routes, true
}

// HasRoutes returns a boolean if a field has been set.
func (o *Solution) HasRoutes() bool {
	if o != nil && o.Routes != nil {
		return true
	}

	return false
}

// SetRoutes gets a reference to the given []Route and assigns it to the Routes field.
func (o *Solution) SetRoutes(v []Route) {
	o.Routes = &v
}

// GetUnassigned returns the Unassigned field value if set, zero value otherwise.
func (o *Solution) GetUnassigned() SolutionUnassigned {
	if o == nil || o.Unassigned == nil {
		var ret SolutionUnassigned
		return ret
	}
	return *o.Unassigned
}

// GetUnassignedOk returns a tuple with the Unassigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Solution) GetUnassignedOk() (*SolutionUnassigned, bool) {
	if o == nil || o.Unassigned == nil {
		return nil, false
	}
	return o.Unassigned, true
}

// HasUnassigned returns a boolean if a field has been set.
func (o *Solution) HasUnassigned() bool {
	if o != nil && o.Unassigned != nil {
		return true
	}

	return false
}

// SetUnassigned gets a reference to the given SolutionUnassigned and assigns it to the Unassigned field.
func (o *Solution) SetUnassigned(v SolutionUnassigned) {
	o.Unassigned = &v
}

func (o Solution) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Costs != nil {
		toSerialize["costs"] = o.Costs
	}
	if o.Distance != nil {
		toSerialize["distance"] = o.Distance
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	if o.TransportTime != nil {
		toSerialize["transport_time"] = o.TransportTime
	}
	if o.MaxOperationTime != nil {
		toSerialize["max_operation_time"] = o.MaxOperationTime
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
	if o.CompletionTime != nil {
		toSerialize["completion_time"] = o.CompletionTime
	}
	if o.NoVehicles != nil {
		toSerialize["no_vehicles"] = o.NoVehicles
	}
	if o.NoUnassigned != nil {
		toSerialize["no_unassigned"] = o.NoUnassigned
	}
	if o.Routes != nil {
		toSerialize["routes"] = o.Routes
	}
	if o.Unassigned != nil {
		toSerialize["unassigned"] = o.Unassigned
	}
	return json.Marshal(toSerialize)
}

type NullableSolution struct {
	value *Solution
	isSet bool
}

func (v NullableSolution) Get() *Solution {
	return v.value
}

func (v *NullableSolution) Set(val *Solution) {
	v.value = val
	v.isSet = true
}

func (v NullableSolution) IsSet() bool {
	return v.isSet
}

func (v *NullableSolution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSolution(val *Solution) *NullableSolution {
	return &NullableSolution{value: val, isSet: true}
}

func (v NullableSolution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSolution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


