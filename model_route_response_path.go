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

// RouteResponsePath struct for RouteResponsePath
type RouteResponsePath struct {
	// The total distance, in meters. To get this information for one 'leg' please read [this blog post](https://www.graphhopper.com/blog/2019/11/28/routing-api-using-path-details/). 
	Distance *float64 `json:"distance,omitempty"`
	// The total travel time, in milliseconds. To get this information for one 'leg' please read [this blog post](https://www.graphhopper.com/blog/2019/11/28/routing-api-using-path-details/). 
	Time *int64 `json:"time,omitempty"`
	// The total ascent, in meters. 
	Ascend *float64 `json:"ascend,omitempty"`
	// The total descent, in meters. 
	Descend *float64 `json:"descend,omitempty"`
	Points *interface{} `json:"points,omitempty"`
	SnappedWaypoints *interface{} `json:"snapped_waypoints,omitempty"`
	// Whether the `points` and `snapped_waypoints` fields are polyline-encoded strings rather than JSON arrays of coordinates. See the field description for more information on the two formats. 
	PointsEncoded *bool `json:"points_encoded,omitempty"`
	// The bounding box of the route geometry. Format: `[minLon, minLat, maxLon, maxLat]`. 
	Bbox *[]float64 `json:"bbox,omitempty"`
	// The instructions for this route. This feature is under active development, and our instructions can sometimes be misleading, so be mindful when using them for navigation. 
	Instructions *[]RouteResponsePathInstructions `json:"instructions,omitempty"`
	// Details, as requested with the `details` parameter. Consider the value `{\"street_name\": [[0,2,\"Frankfurter Straße\"],[2,6,\"Zollweg\"]]}`. In this example, the route uses two streets: The first, Frankfurter Straße, is used between `points[0]` and `points[2]`, and the second, Zollweg, between `points[2]` and `points[6]`. See [here](https://discuss.graphhopper.com/t/2539) for discussion. 
	Details *map[string]interface{} `json:"details,omitempty"`
	// An array of indices (zero-based), specifiying the order in which the input points are visited. Only present if the `optimize` parameter was used. 
	PointsOrder *[]int32 `json:"points_order,omitempty"`
}

// NewRouteResponsePath instantiates a new RouteResponsePath object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRouteResponsePath() *RouteResponsePath {
	this := RouteResponsePath{}
	return &this
}

// NewRouteResponsePathWithDefaults instantiates a new RouteResponsePath object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteResponsePathWithDefaults() *RouteResponsePath {
	this := RouteResponsePath{}
	return &this
}

// GetDistance returns the Distance field value if set, zero value otherwise.
func (o *RouteResponsePath) GetDistance() float64 {
	if o == nil || o.Distance == nil {
		var ret float64
		return ret
	}
	return *o.Distance
}

// GetDistanceOk returns a tuple with the Distance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteResponsePath) GetDistanceOk() (*float64, bool) {
	if o == nil || o.Distance == nil {
		return nil, false
	}
	return o.Distance, true
}

// HasDistance returns a boolean if a field has been set.
func (o *RouteResponsePath) HasDistance() bool {
	if o != nil && o.Distance != nil {
		return true
	}

	return false
}

// SetDistance gets a reference to the given float64 and assigns it to the Distance field.
func (o *RouteResponsePath) SetDistance(v float64) {
	o.Distance = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *RouteResponsePath) GetTime() int64 {
	if o == nil || o.Time == nil {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteResponsePath) GetTimeOk() (*int64, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *RouteResponsePath) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *RouteResponsePath) SetTime(v int64) {
	o.Time = &v
}

// GetAscend returns the Ascend field value if set, zero value otherwise.
func (o *RouteResponsePath) GetAscend() float64 {
	if o == nil || o.Ascend == nil {
		var ret float64
		return ret
	}
	return *o.Ascend
}

// GetAscendOk returns a tuple with the Ascend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteResponsePath) GetAscendOk() (*float64, bool) {
	if o == nil || o.Ascend == nil {
		return nil, false
	}
	return o.Ascend, true
}

// HasAscend returns a boolean if a field has been set.
func (o *RouteResponsePath) HasAscend() bool {
	if o != nil && o.Ascend != nil {
		return true
	}

	return false
}

// SetAscend gets a reference to the given float64 and assigns it to the Ascend field.
func (o *RouteResponsePath) SetAscend(v float64) {
	o.Ascend = &v
}

// GetDescend returns the Descend field value if set, zero value otherwise.
func (o *RouteResponsePath) GetDescend() float64 {
	if o == nil || o.Descend == nil {
		var ret float64
		return ret
	}
	return *o.Descend
}

// GetDescendOk returns a tuple with the Descend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteResponsePath) GetDescendOk() (*float64, bool) {
	if o == nil || o.Descend == nil {
		return nil, false
	}
	return o.Descend, true
}

// HasDescend returns a boolean if a field has been set.
func (o *RouteResponsePath) HasDescend() bool {
	if o != nil && o.Descend != nil {
		return true
	}

	return false
}

// SetDescend gets a reference to the given float64 and assigns it to the Descend field.
func (o *RouteResponsePath) SetDescend(v float64) {
	o.Descend = &v
}

// GetPoints returns the Points field value if set, zero value otherwise.
func (o *RouteResponsePath) GetPoints() interface{} {
	if o == nil || o.Points == nil {
		var ret interface{}
		return ret
	}
	return *o.Points
}

// GetPointsOk returns a tuple with the Points field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteResponsePath) GetPointsOk() (*interface{}, bool) {
	if o == nil || o.Points == nil {
		return nil, false
	}
	return o.Points, true
}

// HasPoints returns a boolean if a field has been set.
func (o *RouteResponsePath) HasPoints() bool {
	if o != nil && o.Points != nil {
		return true
	}

	return false
}

// SetPoints gets a reference to the given interface{} and assigns it to the Points field.
func (o *RouteResponsePath) SetPoints(v interface{}) {
	o.Points = &v
}

// GetSnappedWaypoints returns the SnappedWaypoints field value if set, zero value otherwise.
func (o *RouteResponsePath) GetSnappedWaypoints() interface{} {
	if o == nil || o.SnappedWaypoints == nil {
		var ret interface{}
		return ret
	}
	return *o.SnappedWaypoints
}

// GetSnappedWaypointsOk returns a tuple with the SnappedWaypoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteResponsePath) GetSnappedWaypointsOk() (*interface{}, bool) {
	if o == nil || o.SnappedWaypoints == nil {
		return nil, false
	}
	return o.SnappedWaypoints, true
}

// HasSnappedWaypoints returns a boolean if a field has been set.
func (o *RouteResponsePath) HasSnappedWaypoints() bool {
	if o != nil && o.SnappedWaypoints != nil {
		return true
	}

	return false
}

// SetSnappedWaypoints gets a reference to the given interface{} and assigns it to the SnappedWaypoints field.
func (o *RouteResponsePath) SetSnappedWaypoints(v interface{}) {
	o.SnappedWaypoints = &v
}

// GetPointsEncoded returns the PointsEncoded field value if set, zero value otherwise.
func (o *RouteResponsePath) GetPointsEncoded() bool {
	if o == nil || o.PointsEncoded == nil {
		var ret bool
		return ret
	}
	return *o.PointsEncoded
}

// GetPointsEncodedOk returns a tuple with the PointsEncoded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteResponsePath) GetPointsEncodedOk() (*bool, bool) {
	if o == nil || o.PointsEncoded == nil {
		return nil, false
	}
	return o.PointsEncoded, true
}

// HasPointsEncoded returns a boolean if a field has been set.
func (o *RouteResponsePath) HasPointsEncoded() bool {
	if o != nil && o.PointsEncoded != nil {
		return true
	}

	return false
}

// SetPointsEncoded gets a reference to the given bool and assigns it to the PointsEncoded field.
func (o *RouteResponsePath) SetPointsEncoded(v bool) {
	o.PointsEncoded = &v
}

// GetBbox returns the Bbox field value if set, zero value otherwise.
func (o *RouteResponsePath) GetBbox() []float64 {
	if o == nil || o.Bbox == nil {
		var ret []float64
		return ret
	}
	return *o.Bbox
}

// GetBboxOk returns a tuple with the Bbox field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteResponsePath) GetBboxOk() (*[]float64, bool) {
	if o == nil || o.Bbox == nil {
		return nil, false
	}
	return o.Bbox, true
}

// HasBbox returns a boolean if a field has been set.
func (o *RouteResponsePath) HasBbox() bool {
	if o != nil && o.Bbox != nil {
		return true
	}

	return false
}

// SetBbox gets a reference to the given []float64 and assigns it to the Bbox field.
func (o *RouteResponsePath) SetBbox(v []float64) {
	o.Bbox = &v
}

// GetInstructions returns the Instructions field value if set, zero value otherwise.
func (o *RouteResponsePath) GetInstructions() []RouteResponsePathInstructions {
	if o == nil || o.Instructions == nil {
		var ret []RouteResponsePathInstructions
		return ret
	}
	return *o.Instructions
}

// GetInstructionsOk returns a tuple with the Instructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteResponsePath) GetInstructionsOk() (*[]RouteResponsePathInstructions, bool) {
	if o == nil || o.Instructions == nil {
		return nil, false
	}
	return o.Instructions, true
}

// HasInstructions returns a boolean if a field has been set.
func (o *RouteResponsePath) HasInstructions() bool {
	if o != nil && o.Instructions != nil {
		return true
	}

	return false
}

// SetInstructions gets a reference to the given []RouteResponsePathInstructions and assigns it to the Instructions field.
func (o *RouteResponsePath) SetInstructions(v []RouteResponsePathInstructions) {
	o.Instructions = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *RouteResponsePath) GetDetails() map[string]interface{} {
	if o == nil || o.Details == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteResponsePath) GetDetailsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *RouteResponsePath) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given map[string]interface{} and assigns it to the Details field.
func (o *RouteResponsePath) SetDetails(v map[string]interface{}) {
	o.Details = &v
}

// GetPointsOrder returns the PointsOrder field value if set, zero value otherwise.
func (o *RouteResponsePath) GetPointsOrder() []int32 {
	if o == nil || o.PointsOrder == nil {
		var ret []int32
		return ret
	}
	return *o.PointsOrder
}

// GetPointsOrderOk returns a tuple with the PointsOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteResponsePath) GetPointsOrderOk() (*[]int32, bool) {
	if o == nil || o.PointsOrder == nil {
		return nil, false
	}
	return o.PointsOrder, true
}

// HasPointsOrder returns a boolean if a field has been set.
func (o *RouteResponsePath) HasPointsOrder() bool {
	if o != nil && o.PointsOrder != nil {
		return true
	}

	return false
}

// SetPointsOrder gets a reference to the given []int32 and assigns it to the PointsOrder field.
func (o *RouteResponsePath) SetPointsOrder(v []int32) {
	o.PointsOrder = &v
}

func (o RouteResponsePath) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Distance != nil {
		toSerialize["distance"] = o.Distance
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	if o.Ascend != nil {
		toSerialize["ascend"] = o.Ascend
	}
	if o.Descend != nil {
		toSerialize["descend"] = o.Descend
	}
	if o.Points != nil {
		toSerialize["points"] = o.Points
	}
	if o.SnappedWaypoints != nil {
		toSerialize["snapped_waypoints"] = o.SnappedWaypoints
	}
	if o.PointsEncoded != nil {
		toSerialize["points_encoded"] = o.PointsEncoded
	}
	if o.Bbox != nil {
		toSerialize["bbox"] = o.Bbox
	}
	if o.Instructions != nil {
		toSerialize["instructions"] = o.Instructions
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.PointsOrder != nil {
		toSerialize["points_order"] = o.PointsOrder
	}
	return json.Marshal(toSerialize)
}

type NullableRouteResponsePath struct {
	value *RouteResponsePath
	isSet bool
}

func (v NullableRouteResponsePath) Get() *RouteResponsePath {
	return v.value
}

func (v *NullableRouteResponsePath) Set(val *RouteResponsePath) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteResponsePath) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteResponsePath) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteResponsePath(val *RouteResponsePath) *NullableRouteResponsePath {
	return &NullableRouteResponsePath{value: val, isSet: true}
}

func (v NullableRouteResponsePath) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteResponsePath) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


