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

// Routing This contains all routing specific configurations.
type Routing struct {
	// It lets you specify whether the API should provide you with route geometries for vehicle routes or not. Thus, you do not need to do extra routing to get the polyline for each route.
	CalcPoints *bool `json:"calc_points,omitempty"`
	// indicates whether historical traffic information should be considered
	ConsiderTraffic *bool `json:"consider_traffic,omitempty"`
	// specifies the data provider, read more about it [here](#section/Map-Data-and-Routing-Profiles).
	NetworkDataProvider *string `json:"network_data_provider,omitempty"`
	// In some cases curbside constraints cannot be fulfilled. For example in one-way streets you cannot arrive at a building that is on the left side of the street such that the building is to the right of you (unless you drove the one-way street the wrong/illegal way). You can set the `curbside_strictness` to `soft` to ignore the curbside constraint in such cases or set it to `strict` to get an error response instead. You can also set it to `ignore` to ignore all curbside constraints (this is useful to compare the results with and without constraints without modifying every single address).
	CurbsideStrictness *string `json:"curbside_strictness,omitempty"`
	// indicates whether matrix calculation should fail fast when points cannot be connected
	FailFast *bool `json:"fail_fast,omitempty"`
	// Indicates whether a solution includes snapped waypoints. In contrary to the address coordinate a snapped waypoint is the access point to the (road) network.
	ReturnSnappedWaypoints *bool `json:"return_snapped_waypoints,omitempty"`
	// Prevents snapping locations to road links of specified road types, e.g. to motorway.
	SnapPreventions *[]string `json:"snap_preventions,omitempty"`
}

// NewRouting instantiates a new Routing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRouting() *Routing {
	this := Routing{}
	var calcPoints bool = false
	this.CalcPoints = &calcPoints
	var considerTraffic bool = false
	this.ConsiderTraffic = &considerTraffic
	var networkDataProvider string = "openstreetmap"
	this.NetworkDataProvider = &networkDataProvider
	var curbsideStrictness string = "soft"
	this.CurbsideStrictness = &curbsideStrictness
	var failFast bool = true
	this.FailFast = &failFast
	var returnSnappedWaypoints bool = false
	this.ReturnSnappedWaypoints = &returnSnappedWaypoints
	return &this
}

// NewRoutingWithDefaults instantiates a new Routing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingWithDefaults() *Routing {
	this := Routing{}
	var calcPoints bool = false
	this.CalcPoints = &calcPoints
	var considerTraffic bool = false
	this.ConsiderTraffic = &considerTraffic
	var networkDataProvider string = "openstreetmap"
	this.NetworkDataProvider = &networkDataProvider
	var curbsideStrictness string = "soft"
	this.CurbsideStrictness = &curbsideStrictness
	var failFast bool = true
	this.FailFast = &failFast
	var returnSnappedWaypoints bool = false
	this.ReturnSnappedWaypoints = &returnSnappedWaypoints
	return &this
}

// GetCalcPoints returns the CalcPoints field value if set, zero value otherwise.
func (o *Routing) GetCalcPoints() bool {
	if o == nil || o.CalcPoints == nil {
		var ret bool
		return ret
	}
	return *o.CalcPoints
}

// GetCalcPointsOk returns a tuple with the CalcPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Routing) GetCalcPointsOk() (*bool, bool) {
	if o == nil || o.CalcPoints == nil {
		return nil, false
	}
	return o.CalcPoints, true
}

// HasCalcPoints returns a boolean if a field has been set.
func (o *Routing) HasCalcPoints() bool {
	if o != nil && o.CalcPoints != nil {
		return true
	}

	return false
}

// SetCalcPoints gets a reference to the given bool and assigns it to the CalcPoints field.
func (o *Routing) SetCalcPoints(v bool) {
	o.CalcPoints = &v
}

// GetConsiderTraffic returns the ConsiderTraffic field value if set, zero value otherwise.
func (o *Routing) GetConsiderTraffic() bool {
	if o == nil || o.ConsiderTraffic == nil {
		var ret bool
		return ret
	}
	return *o.ConsiderTraffic
}

// GetConsiderTrafficOk returns a tuple with the ConsiderTraffic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Routing) GetConsiderTrafficOk() (*bool, bool) {
	if o == nil || o.ConsiderTraffic == nil {
		return nil, false
	}
	return o.ConsiderTraffic, true
}

// HasConsiderTraffic returns a boolean if a field has been set.
func (o *Routing) HasConsiderTraffic() bool {
	if o != nil && o.ConsiderTraffic != nil {
		return true
	}

	return false
}

// SetConsiderTraffic gets a reference to the given bool and assigns it to the ConsiderTraffic field.
func (o *Routing) SetConsiderTraffic(v bool) {
	o.ConsiderTraffic = &v
}

// GetNetworkDataProvider returns the NetworkDataProvider field value if set, zero value otherwise.
func (o *Routing) GetNetworkDataProvider() string {
	if o == nil || o.NetworkDataProvider == nil {
		var ret string
		return ret
	}
	return *o.NetworkDataProvider
}

// GetNetworkDataProviderOk returns a tuple with the NetworkDataProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Routing) GetNetworkDataProviderOk() (*string, bool) {
	if o == nil || o.NetworkDataProvider == nil {
		return nil, false
	}
	return o.NetworkDataProvider, true
}

// HasNetworkDataProvider returns a boolean if a field has been set.
func (o *Routing) HasNetworkDataProvider() bool {
	if o != nil && o.NetworkDataProvider != nil {
		return true
	}

	return false
}

// SetNetworkDataProvider gets a reference to the given string and assigns it to the NetworkDataProvider field.
func (o *Routing) SetNetworkDataProvider(v string) {
	o.NetworkDataProvider = &v
}

// GetCurbsideStrictness returns the CurbsideStrictness field value if set, zero value otherwise.
func (o *Routing) GetCurbsideStrictness() string {
	if o == nil || o.CurbsideStrictness == nil {
		var ret string
		return ret
	}
	return *o.CurbsideStrictness
}

// GetCurbsideStrictnessOk returns a tuple with the CurbsideStrictness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Routing) GetCurbsideStrictnessOk() (*string, bool) {
	if o == nil || o.CurbsideStrictness == nil {
		return nil, false
	}
	return o.CurbsideStrictness, true
}

// HasCurbsideStrictness returns a boolean if a field has been set.
func (o *Routing) HasCurbsideStrictness() bool {
	if o != nil && o.CurbsideStrictness != nil {
		return true
	}

	return false
}

// SetCurbsideStrictness gets a reference to the given string and assigns it to the CurbsideStrictness field.
func (o *Routing) SetCurbsideStrictness(v string) {
	o.CurbsideStrictness = &v
}

// GetFailFast returns the FailFast field value if set, zero value otherwise.
func (o *Routing) GetFailFast() bool {
	if o == nil || o.FailFast == nil {
		var ret bool
		return ret
	}
	return *o.FailFast
}

// GetFailFastOk returns a tuple with the FailFast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Routing) GetFailFastOk() (*bool, bool) {
	if o == nil || o.FailFast == nil {
		return nil, false
	}
	return o.FailFast, true
}

// HasFailFast returns a boolean if a field has been set.
func (o *Routing) HasFailFast() bool {
	if o != nil && o.FailFast != nil {
		return true
	}

	return false
}

// SetFailFast gets a reference to the given bool and assigns it to the FailFast field.
func (o *Routing) SetFailFast(v bool) {
	o.FailFast = &v
}

// GetReturnSnappedWaypoints returns the ReturnSnappedWaypoints field value if set, zero value otherwise.
func (o *Routing) GetReturnSnappedWaypoints() bool {
	if o == nil || o.ReturnSnappedWaypoints == nil {
		var ret bool
		return ret
	}
	return *o.ReturnSnappedWaypoints
}

// GetReturnSnappedWaypointsOk returns a tuple with the ReturnSnappedWaypoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Routing) GetReturnSnappedWaypointsOk() (*bool, bool) {
	if o == nil || o.ReturnSnappedWaypoints == nil {
		return nil, false
	}
	return o.ReturnSnappedWaypoints, true
}

// HasReturnSnappedWaypoints returns a boolean if a field has been set.
func (o *Routing) HasReturnSnappedWaypoints() bool {
	if o != nil && o.ReturnSnappedWaypoints != nil {
		return true
	}

	return false
}

// SetReturnSnappedWaypoints gets a reference to the given bool and assigns it to the ReturnSnappedWaypoints field.
func (o *Routing) SetReturnSnappedWaypoints(v bool) {
	o.ReturnSnappedWaypoints = &v
}

// GetSnapPreventions returns the SnapPreventions field value if set, zero value otherwise.
func (o *Routing) GetSnapPreventions() []string {
	if o == nil || o.SnapPreventions == nil {
		var ret []string
		return ret
	}
	return *o.SnapPreventions
}

// GetSnapPreventionsOk returns a tuple with the SnapPreventions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Routing) GetSnapPreventionsOk() (*[]string, bool) {
	if o == nil || o.SnapPreventions == nil {
		return nil, false
	}
	return o.SnapPreventions, true
}

// HasSnapPreventions returns a boolean if a field has been set.
func (o *Routing) HasSnapPreventions() bool {
	if o != nil && o.SnapPreventions != nil {
		return true
	}

	return false
}

// SetSnapPreventions gets a reference to the given []string and assigns it to the SnapPreventions field.
func (o *Routing) SetSnapPreventions(v []string) {
	o.SnapPreventions = &v
}

func (o Routing) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CalcPoints != nil {
		toSerialize["calc_points"] = o.CalcPoints
	}
	if o.ConsiderTraffic != nil {
		toSerialize["consider_traffic"] = o.ConsiderTraffic
	}
	if o.NetworkDataProvider != nil {
		toSerialize["network_data_provider"] = o.NetworkDataProvider
	}
	if o.CurbsideStrictness != nil {
		toSerialize["curbside_strictness"] = o.CurbsideStrictness
	}
	if o.FailFast != nil {
		toSerialize["fail_fast"] = o.FailFast
	}
	if o.ReturnSnappedWaypoints != nil {
		toSerialize["return_snapped_waypoints"] = o.ReturnSnappedWaypoints
	}
	if o.SnapPreventions != nil {
		toSerialize["snap_preventions"] = o.SnapPreventions
	}
	return json.Marshal(toSerialize)
}

type NullableRouting struct {
	value *Routing
	isSet bool
}

func (v NullableRouting) Get() *Routing {
	return v.value
}

func (v *NullableRouting) Set(val *Routing) {
	v.value = val
	v.isSet = true
}

func (v NullableRouting) IsSet() bool {
	return v.isSet
}

func (v *NullableRouting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouting(val *Routing) *NullableRouting {
	return &NullableRouting{value: val, isSet: true}
}

func (v NullableRouting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


