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

// RouteRequest struct for RouteRequest
type RouteRequest struct {
	// The points for the route in an array of `[longitude,latitude]`. For instance, if you want to calculate a route from point A to B to C then you specify `points: [ [A_longitude, A_latitude], [B_longitude, B_latitude], [C_longitude, C_latitude]] 
	Points *[][]float64 `json:"points,omitempty"`
	// Optional parameter. Specifies a hint for each point in the `points` array to prefer a certain street for the closest location lookup. E.g. if there is an address or house with two or more neighboring streets you can control for which street the closest location is looked up.
	PointHints *[]string `json:"point_hints,omitempty"`
	// Optional parameter to avoid snapping to a certain road class or road environment. Current supported values `motorway`, `trunk`, `ferry`, `tunnel`, `bridge` and `ford`
	SnapPreventions *[]string `json:"snap_preventions,omitempty"`
	// Optional parameter. It specifies on which side a point should be relative to the driver when she leaves/arrives at a start/target/via point. You need to specify this parameter for either none or all points. Only supported for motor vehicles and OpenStreetMap.
	Curbsides *[]string `json:"curbsides,omitempty"`
	Vehicle *VehicleProfileId `json:"vehicle,omitempty"`
	// The locale of the resulting turn instructions. E.g. `pt_PT` for Portuguese or `de` for German. 
	Locale *string `json:"locale,omitempty"`
	// If `true`, a third coordinate, the altitude, is included with all positions in the response. This changes the format of the `points` and `snapped_waypoints` fields of the response, in both their encodings. Unless you switch off the `points_encoded` parameter, you need special code on the client side that can handle three-dimensional coordinates. A request can fail if the vehicle profile does not support elevation. See the features object for every vehicle profile. 
	Elevation *bool `json:"elevation,omitempty"`
	// Optional parameter to retrieve path details. You can request additional details for the route: `street_name`, `time`, `distance`, `max_speed`, `toll`, `road_class`, `road_class_link`, `road_access`, `road_environment`, `lanes`, and `surface`. Read more about the usage of path details [here](https://discuss.graphhopper.com/t/2539). 
	Details *[]string `json:"details,omitempty"`
	// Normally, the calculated route will visit the points in the order you specified them. If you have more than two points, you can set this parameter to `\"true\"` and the points may be re-ordered to minimize the total travel time. Keep in mind that the limits on the number of locations of the Route Optimization API applies, and the request costs more credits. 
	Optimize *string `json:"optimize,omitempty"`
	// If instructions should be calculated and returned 
	Instructions *bool `json:"instructions,omitempty"`
	// If the points for the route should be calculated at all. 
	CalcPoints *bool `json:"calc_points,omitempty"`
	// If `true`, the output will be formatted. 
	Debug *bool `json:"debug,omitempty"`
	// Allows changing the encoding of location data in the response. The default is polyline encoding, which is compact but requires special client code to unpack. (We provide it in our JavaScript client library!) Set this parameter to `false` to switch the encoding to simple coordinate pairs like `[lon,lat]`, or `[lon,lat,elevation]`. See the description of the response format for more information. 
	PointsEncoded *bool `json:"points_encoded,omitempty"`
	// Use this parameter in combination with one or more parameters from below. 
	ChDisable *bool `json:"ch.disable,omitempty"`
	// Determines the way the ''best'' route is calculated. Default is `fastest`. Other options are `shortest` (e.g. for `vehicle=foot` or `bike`) and `short_fastest` which finds a reasonable balance between `shortest` and `fastest`. Requires `ch.disable=true`. 
	Weighting *string `json:"weighting,omitempty"`
	// Favour a heading direction for a certain point. Specify either one heading for the start point or as many as there are points. In this case headings are associated by their order to the specific points. Headings are given as north based clockwise angle between 0 and 360 degree. This parameter also influences the tour generated with `algorithm=round_trip` and forces the initial direction.  Requires `ch.disable=true`. 
	Headings *[]int32 `json:"headings,omitempty"`
	// Time penalty in seconds for not obeying a specified heading. Requires `ch.disable=true`. 
	HeadingPenalty *int32 `json:"heading_penalty,omitempty"`
	// If `true`, u-turns are avoided at via-points with regard to the `heading_penalty`. Requires `ch.disable=true`. 
	PassThrough *bool `json:"pass_through,omitempty"`
	// Block road access via a point with the format `latitude,longitude` or an area defined by a circle `lat,lon,radius` or a rectangle `lat1,lon1,lat2,lon2`. Separate several values with `;`. Requires `ch.disable=true`. 
	BlockArea *string `json:"block_area,omitempty"`
	// Specify which road classes and environments you would like to avoid. Possible values are `motorway`, `steps`, `track`, `toll`, `ferry`, `tunnel` and `bridge`. Separate several values with `;`. Obviously not all the values make sense for all vehicle profiles e.g. `bike` is already forbidden on a `motorway`. Requires `ch.disable=true`. 
	Avoid *string `json:"avoid,omitempty"`
	// Rather than looking for the shortest or fastest path, this lets you solve two different problems related to routing: With `round_trip`, the route will get you back to where you started. This is meant for fun (think of a bike trip), so we will add some randomness. This requires `ch.disable=true`. With `alternative_route`, we give you not one but several routes that are close to optimal, but not too similar to each other. You can control both of these features with additional parameters, see below. 
	Algorithm *string `json:"algorithm,omitempty"`
	// If `algorithm=round_trip`, this parameter configures approximative length of the resulting round trip. Requires `ch.disable=true`. 
	RoundTripDistance *int32 `json:"round_trip.distance,omitempty"`
	// If `algorithm=round_trip`, this sets the random seed. Change this to get a different tour for each value. 
	RoundTripSeed *int64 `json:"round_trip.seed,omitempty"`
	// If `algorithm=alternative_route`, this parameter sets the number of maximum paths which should be calculated. Increasing can lead to worse alternatives. 
	AlternativeRouteMaxPaths *int32 `json:"alternative_route.max_paths,omitempty"`
	// If `algorithm=alternative_route`, this parameter sets the factor by which the alternatives routes can be longer than the optimal route. Increasing can lead to worse alternatives. 
	AlternativeRouteMaxWeightFactor *float32 `json:"alternative_route.max_weight_factor,omitempty"`
	// If `algorithm=alternative_route`, this parameter specifies how similar an alternative route can be to the optimal route. Increasing can lead to worse alternatives. 
	AlternativeRouteMaxShareFactor *float32 `json:"alternative_route.max_share_factor,omitempty"`
}

// NewRouteRequest instantiates a new RouteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRouteRequest() *RouteRequest {
	this := RouteRequest{}
	var locale string = "en"
	this.Locale = &locale
	var elevation bool = false
	this.Elevation = &elevation
	var optimize string = "false"
	this.Optimize = &optimize
	var instructions bool = true
	this.Instructions = &instructions
	var calcPoints bool = true
	this.CalcPoints = &calcPoints
	var debug bool = false
	this.Debug = &debug
	var pointsEncoded bool = true
	this.PointsEncoded = &pointsEncoded
	var chDisable bool = false
	this.ChDisable = &chDisable
	var weighting string = "fastest"
	this.Weighting = &weighting
	var headingPenalty int32 = 120
	this.HeadingPenalty = &headingPenalty
	var passThrough bool = false
	this.PassThrough = &passThrough
	var roundTripDistance int32 = 10000
	this.RoundTripDistance = &roundTripDistance
	var alternativeRouteMaxPaths int32 = 2
	this.AlternativeRouteMaxPaths = &alternativeRouteMaxPaths
	var alternativeRouteMaxWeightFactor float32 = 1.4
	this.AlternativeRouteMaxWeightFactor = &alternativeRouteMaxWeightFactor
	var alternativeRouteMaxShareFactor float32 = 0.6
	this.AlternativeRouteMaxShareFactor = &alternativeRouteMaxShareFactor
	return &this
}

// NewRouteRequestWithDefaults instantiates a new RouteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteRequestWithDefaults() *RouteRequest {
	this := RouteRequest{}
	var locale string = "en"
	this.Locale = &locale
	var elevation bool = false
	this.Elevation = &elevation
	var optimize string = "false"
	this.Optimize = &optimize
	var instructions bool = true
	this.Instructions = &instructions
	var calcPoints bool = true
	this.CalcPoints = &calcPoints
	var debug bool = false
	this.Debug = &debug
	var pointsEncoded bool = true
	this.PointsEncoded = &pointsEncoded
	var chDisable bool = false
	this.ChDisable = &chDisable
	var weighting string = "fastest"
	this.Weighting = &weighting
	var headingPenalty int32 = 120
	this.HeadingPenalty = &headingPenalty
	var passThrough bool = false
	this.PassThrough = &passThrough
	var roundTripDistance int32 = 10000
	this.RoundTripDistance = &roundTripDistance
	var alternativeRouteMaxPaths int32 = 2
	this.AlternativeRouteMaxPaths = &alternativeRouteMaxPaths
	var alternativeRouteMaxWeightFactor float32 = 1.4
	this.AlternativeRouteMaxWeightFactor = &alternativeRouteMaxWeightFactor
	var alternativeRouteMaxShareFactor float32 = 0.6
	this.AlternativeRouteMaxShareFactor = &alternativeRouteMaxShareFactor
	return &this
}

// GetPoints returns the Points field value if set, zero value otherwise.
func (o *RouteRequest) GetPoints() [][]float64 {
	if o == nil || o.Points == nil {
		var ret [][]float64
		return ret
	}
	return *o.Points
}

// GetPointsOk returns a tuple with the Points field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteRequest) GetPointsOk() (*[][]float64, bool) {
	if o == nil || o.Points == nil {
		return nil, false
	}
	return o.Points, true
}

// HasPoints returns a boolean if a field has been set.
func (o *RouteRequest) HasPoints() bool {
	if o != nil && o.Points != nil {
		return true
	}

	return false
}

// SetPoints gets a reference to the given [][]float64 and assigns it to the Points field.
func (o *RouteRequest) SetPoints(v [][]float64) {
	o.Points = &v
}

// GetPointHints returns the PointHints field value if set, zero value otherwise.
func (o *RouteRequest) GetPointHints() []string {
	if o == nil || o.PointHints == nil {
		var ret []string
		return ret
	}
	return *o.PointHints
}

// GetPointHintsOk returns a tuple with the PointHints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteRequest) GetPointHintsOk() (*[]string, bool) {
	if o == nil || o.PointHints == nil {
		return nil, false
	}
	return o.PointHints, true
}

// HasPointHints returns a boolean if a field has been set.
func (o *RouteRequest) HasPointHints() bool {
	if o != nil && o.PointHints != nil {
		return true
	}

	return false
}

// SetPointHints gets a reference to the given []string and assigns it to the PointHints field.
func (o *RouteRequest) SetPointHints(v []string) {
	o.PointHints = &v
}

// GetSnapPreventions returns the SnapPreventions field value if set, zero value otherwise.
func (o *RouteRequest) GetSnapPreventions() []string {
	if o == nil || o.SnapPreventions == nil {
		var ret []string
		return ret
	}
	return *o.SnapPreventions
}

// GetSnapPreventionsOk returns a tuple with the SnapPreventions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteRequest) GetSnapPreventionsOk() (*[]string, bool) {
	if o == nil || o.SnapPreventions == nil {
		return nil, false
	}
	return o.SnapPreventions, true
}

// HasSnapPreventions returns a boolean if a field has been set.
func (o *RouteRequest) HasSnapPreventions() bool {
	if o != nil && o.SnapPreventions != nil {
		return true
	}

	return false
}

// SetSnapPreventions gets a reference to the given []string and assigns it to the SnapPreventions field.
func (o *RouteRequest) SetSnapPreventions(v []string) {
	o.SnapPreventions = &v
}

// GetCurbsides returns the Curbsides field value if set, zero value otherwise.
func (o *RouteRequest) GetCurbsides() []string {
	if o == nil || o.Curbsides == nil {
		var ret []string
		return ret
	}
	return *o.Curbsides
}

// GetCurbsidesOk returns a tuple with the Curbsides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteRequest) GetCurbsidesOk() (*[]string, bool) {
	if o == nil || o.Curbsides == nil {
		return nil, false
	}
	return o.Curbsides, true
}

// HasCurbsides returns a boolean if a field has been set.
func (o *RouteRequest) HasCurbsides() bool {
	if o != nil && o.Curbsides != nil {
		return true
	}

	return false
}

// SetCurbsides gets a reference to the given []string and assigns it to the Curbsides field.
func (o *RouteRequest) SetCurbsides(v []string) {
	o.Curbsides = &v
}

// GetVehicle returns the Vehicle field value if set, zero value otherwise.
func (o *RouteRequest) GetVehicle() VehicleProfileId {
	if o == nil || o.Vehicle == nil {
		var ret VehicleProfileId
		return ret
	}
	return *o.Vehicle
}

// GetVehicleOk returns a tuple with the Vehicle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteRequest) GetVehicleOk() (*VehicleProfileId, bool) {
	if o == nil || o.Vehicle == nil {
		return nil, false
	}
	return o.Vehicle, true
}

// HasVehicle returns a boolean if a field has been set.
func (o *RouteRequest) HasVehicle() bool {
	if o != nil && o.Vehicle != nil {
		return true
	}

	return false
}

// SetVehicle gets a reference to the given VehicleProfileId and assigns it to the Vehicle field.
func (o *RouteRequest) SetVehicle(v VehicleProfileId) {
	o.Vehicle = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *RouteRequest) GetLocale() string {
	if o == nil || o.Locale == nil {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteRequest) GetLocaleOk() (*string, bool) {
	if o == nil || o.Locale == nil {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *RouteRequest) HasLocale() bool {
	if o != nil && o.Locale != nil {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *RouteRequest) SetLocale(v string) {
	o.Locale = &v
}

// GetElevation returns the Elevation field value if set, zero value otherwise.
func (o *RouteRequest) GetElevation() bool {
	if o == nil || o.Elevation == nil {
		var ret bool
		return ret
	}
	return *o.Elevation
}

// GetElevationOk returns a tuple with the Elevation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteRequest) GetElevationOk() (*bool, bool) {
	if o == nil || o.Elevation == nil {
		return nil, false
	}
	return o.Elevation, true
}

// HasElevation returns a boolean if a field has been set.
func (o *RouteRequest) HasElevation() bool {
	if o != nil && o.Elevation != nil {
		return true
	}

	return false
}

// SetElevation gets a reference to the given bool and assigns it to the Elevation field.
func (o *RouteRequest) SetElevation(v bool) {
	o.Elevation = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *RouteRequest) GetDetails() []string {
	if o == nil || o.Details == nil {
		var ret []string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteRequest) GetDetailsOk() (*[]string, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *RouteRequest) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []string and assigns it to the Details field.
func (o *RouteRequest) SetDetails(v []string) {
	o.Details = &v
}

// GetOptimize returns the Optimize field value if set, zero value otherwise.
func (o *RouteRequest) GetOptimize() string {
	if o == nil || o.Optimize == nil {
		var ret string
		return ret
	}
	return *o.Optimize
}

// GetOptimizeOk returns a tuple with the Optimize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteRequest) GetOptimizeOk() (*string, bool) {
	if o == nil || o.Optimize == nil {
		return nil, false
	}
	return o.Optimize, true
}

// HasOptimize returns a boolean if a field has been set.
func (o *RouteRequest) HasOptimize() bool {
	if o != nil && o.Optimize != nil {
		return true
	}

	return false
}

// SetOptimize gets a reference to the given string and assigns it to the Optimize field.
func (o *RouteRequest) SetOptimize(v string) {
	o.Optimize = &v
}

// GetInstructions returns the Instructions field value if set, zero value otherwise.
func (o *RouteRequest) GetInstructions() bool {
	if o == nil || o.Instructions == nil {
		var ret bool
		return ret
	}
	return *o.Instructions
}

// GetInstructionsOk returns a tuple with the Instructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteRequest) GetInstructionsOk() (*bool, bool) {
	if o == nil || o.Instructions == nil {
		return nil, false
	}
	return o.Instructions, true
}

// HasInstructions returns a boolean if a field has been set.
func (o *RouteRequest) HasInstructions() bool {
	if o != nil && o.Instructions != nil {
		return true
	}

	return false
}

// SetInstructions gets a reference to the given bool and assigns it to the Instructions field.
func (o *RouteRequest) SetInstructions(v bool) {
	o.Instructions = &v
}

// GetCalcPoints returns the CalcPoints field value if set, zero value otherwise.
func (o *RouteRequest) GetCalcPoints() bool {
	if o == nil || o.CalcPoints == nil {
		var ret bool
		return ret
	}
	return *o.CalcPoints
}

// GetCalcPointsOk returns a tuple with the CalcPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteRequest) GetCalcPointsOk() (*bool, bool) {
	if o == nil || o.CalcPoints == nil {
		return nil, false
	}
	return o.CalcPoints, true
}

// HasCalcPoints returns a boolean if a field has been set.
func (o *RouteRequest) HasCalcPoints() bool {
	if o != nil && o.CalcPoints != nil {
		return true
	}

	return false
}

// SetCalcPoints gets a reference to the given bool and assigns it to the CalcPoints field.
func (o *RouteRequest) SetCalcPoints(v bool) {
	o.CalcPoints = &v
}

// GetDebug returns the Debug field value if set, zero value otherwise.
func (o *RouteRequest) GetDebug() bool {
	if o == nil || o.Debug == nil {
		var ret bool
		return ret
	}
	return *o.Debug
}

// GetDebugOk returns a tuple with the Debug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteRequest) GetDebugOk() (*bool, bool) {
	if o == nil || o.Debug == nil {
		return nil, false
	}
	return o.Debug, true
}

// HasDebug returns a boolean if a field has been set.
func (o *RouteRequest) HasDebug() bool {
	if o != nil && o.Debug != nil {
		return true
	}

	return false
}

// SetDebug gets a reference to the given bool and assigns it to the Debug field.
func (o *RouteRequest) SetDebug(v bool) {
	o.Debug = &v
}

// GetPointsEncoded returns the PointsEncoded field value if set, zero value otherwise.
func (o *RouteRequest) GetPointsEncoded() bool {
	if o == nil || o.PointsEncoded == nil {
		var ret bool
		return ret
	}
	return *o.PointsEncoded
}

// GetPointsEncodedOk returns a tuple with the PointsEncoded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteRequest) GetPointsEncodedOk() (*bool, bool) {
	if o == nil || o.PointsEncoded == nil {
		return nil, false
	}
	return o.PointsEncoded, true
}

// HasPointsEncoded returns a boolean if a field has been set.
func (o *RouteRequest) HasPointsEncoded() bool {
	if o != nil && o.PointsEncoded != nil {
		return true
	}

	return false
}

// SetPointsEncoded gets a reference to the given bool and assigns it to the PointsEncoded field.
func (o *RouteRequest) SetPointsEncoded(v bool) {
	o.PointsEncoded = &v
}

// GetChDisable returns the ChDisable field value if set, zero value otherwise.
func (o *RouteRequest) GetChDisable() bool {
	if o == nil || o.ChDisable == nil {
		var ret bool
		return ret
	}
	return *o.ChDisable
}

// GetChDisableOk returns a tuple with the ChDisable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteRequest) GetChDisableOk() (*bool, bool) {
	if o == nil || o.ChDisable == nil {
		return nil, false
	}
	return o.ChDisable, true
}

// HasChDisable returns a boolean if a field has been set.
func (o *RouteRequest) HasChDisable() bool {
	if o != nil && o.ChDisable != nil {
		return true
	}

	return false
}

// SetChDisable gets a reference to the given bool and assigns it to the ChDisable field.
func (o *RouteRequest) SetChDisable(v bool) {
	o.ChDisable = &v
}

// GetWeighting returns the Weighting field value if set, zero value otherwise.
func (o *RouteRequest) GetWeighting() string {
	if o == nil || o.Weighting == nil {
		var ret string
		return ret
	}
	return *o.Weighting
}

// GetWeightingOk returns a tuple with the Weighting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteRequest) GetWeightingOk() (*string, bool) {
	if o == nil || o.Weighting == nil {
		return nil, false
	}
	return o.Weighting, true
}

// HasWeighting returns a boolean if a field has been set.
func (o *RouteRequest) HasWeighting() bool {
	if o != nil && o.Weighting != nil {
		return true
	}

	return false
}

// SetWeighting gets a reference to the given string and assigns it to the Weighting field.
func (o *RouteRequest) SetWeighting(v string) {
	o.Weighting = &v
}

// GetHeadings returns the Headings field value if set, zero value otherwise.
func (o *RouteRequest) GetHeadings() []int32 {
	if o == nil || o.Headings == nil {
		var ret []int32
		return ret
	}
	return *o.Headings
}

// GetHeadingsOk returns a tuple with the Headings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteRequest) GetHeadingsOk() (*[]int32, bool) {
	if o == nil || o.Headings == nil {
		return nil, false
	}
	return o.Headings, true
}

// HasHeadings returns a boolean if a field has been set.
func (o *RouteRequest) HasHeadings() bool {
	if o != nil && o.Headings != nil {
		return true
	}

	return false
}

// SetHeadings gets a reference to the given []int32 and assigns it to the Headings field.
func (o *RouteRequest) SetHeadings(v []int32) {
	o.Headings = &v
}

// GetHeadingPenalty returns the HeadingPenalty field value if set, zero value otherwise.
func (o *RouteRequest) GetHeadingPenalty() int32 {
	if o == nil || o.HeadingPenalty == nil {
		var ret int32
		return ret
	}
	return *o.HeadingPenalty
}

// GetHeadingPenaltyOk returns a tuple with the HeadingPenalty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteRequest) GetHeadingPenaltyOk() (*int32, bool) {
	if o == nil || o.HeadingPenalty == nil {
		return nil, false
	}
	return o.HeadingPenalty, true
}

// HasHeadingPenalty returns a boolean if a field has been set.
func (o *RouteRequest) HasHeadingPenalty() bool {
	if o != nil && o.HeadingPenalty != nil {
		return true
	}

	return false
}

// SetHeadingPenalty gets a reference to the given int32 and assigns it to the HeadingPenalty field.
func (o *RouteRequest) SetHeadingPenalty(v int32) {
	o.HeadingPenalty = &v
}

// GetPassThrough returns the PassThrough field value if set, zero value otherwise.
func (o *RouteRequest) GetPassThrough() bool {
	if o == nil || o.PassThrough == nil {
		var ret bool
		return ret
	}
	return *o.PassThrough
}

// GetPassThroughOk returns a tuple with the PassThrough field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteRequest) GetPassThroughOk() (*bool, bool) {
	if o == nil || o.PassThrough == nil {
		return nil, false
	}
	return o.PassThrough, true
}

// HasPassThrough returns a boolean if a field has been set.
func (o *RouteRequest) HasPassThrough() bool {
	if o != nil && o.PassThrough != nil {
		return true
	}

	return false
}

// SetPassThrough gets a reference to the given bool and assigns it to the PassThrough field.
func (o *RouteRequest) SetPassThrough(v bool) {
	o.PassThrough = &v
}

// GetBlockArea returns the BlockArea field value if set, zero value otherwise.
func (o *RouteRequest) GetBlockArea() string {
	if o == nil || o.BlockArea == nil {
		var ret string
		return ret
	}
	return *o.BlockArea
}

// GetBlockAreaOk returns a tuple with the BlockArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteRequest) GetBlockAreaOk() (*string, bool) {
	if o == nil || o.BlockArea == nil {
		return nil, false
	}
	return o.BlockArea, true
}

// HasBlockArea returns a boolean if a field has been set.
func (o *RouteRequest) HasBlockArea() bool {
	if o != nil && o.BlockArea != nil {
		return true
	}

	return false
}

// SetBlockArea gets a reference to the given string and assigns it to the BlockArea field.
func (o *RouteRequest) SetBlockArea(v string) {
	o.BlockArea = &v
}

// GetAvoid returns the Avoid field value if set, zero value otherwise.
func (o *RouteRequest) GetAvoid() string {
	if o == nil || o.Avoid == nil {
		var ret string
		return ret
	}
	return *o.Avoid
}

// GetAvoidOk returns a tuple with the Avoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteRequest) GetAvoidOk() (*string, bool) {
	if o == nil || o.Avoid == nil {
		return nil, false
	}
	return o.Avoid, true
}

// HasAvoid returns a boolean if a field has been set.
func (o *RouteRequest) HasAvoid() bool {
	if o != nil && o.Avoid != nil {
		return true
	}

	return false
}

// SetAvoid gets a reference to the given string and assigns it to the Avoid field.
func (o *RouteRequest) SetAvoid(v string) {
	o.Avoid = &v
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise.
func (o *RouteRequest) GetAlgorithm() string {
	if o == nil || o.Algorithm == nil {
		var ret string
		return ret
	}
	return *o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteRequest) GetAlgorithmOk() (*string, bool) {
	if o == nil || o.Algorithm == nil {
		return nil, false
	}
	return o.Algorithm, true
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *RouteRequest) HasAlgorithm() bool {
	if o != nil && o.Algorithm != nil {
		return true
	}

	return false
}

// SetAlgorithm gets a reference to the given string and assigns it to the Algorithm field.
func (o *RouteRequest) SetAlgorithm(v string) {
	o.Algorithm = &v
}

// GetRoundTripDistance returns the RoundTripDistance field value if set, zero value otherwise.
func (o *RouteRequest) GetRoundTripDistance() int32 {
	if o == nil || o.RoundTripDistance == nil {
		var ret int32
		return ret
	}
	return *o.RoundTripDistance
}

// GetRoundTripDistanceOk returns a tuple with the RoundTripDistance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteRequest) GetRoundTripDistanceOk() (*int32, bool) {
	if o == nil || o.RoundTripDistance == nil {
		return nil, false
	}
	return o.RoundTripDistance, true
}

// HasRoundTripDistance returns a boolean if a field has been set.
func (o *RouteRequest) HasRoundTripDistance() bool {
	if o != nil && o.RoundTripDistance != nil {
		return true
	}

	return false
}

// SetRoundTripDistance gets a reference to the given int32 and assigns it to the RoundTripDistance field.
func (o *RouteRequest) SetRoundTripDistance(v int32) {
	o.RoundTripDistance = &v
}

// GetRoundTripSeed returns the RoundTripSeed field value if set, zero value otherwise.
func (o *RouteRequest) GetRoundTripSeed() int64 {
	if o == nil || o.RoundTripSeed == nil {
		var ret int64
		return ret
	}
	return *o.RoundTripSeed
}

// GetRoundTripSeedOk returns a tuple with the RoundTripSeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteRequest) GetRoundTripSeedOk() (*int64, bool) {
	if o == nil || o.RoundTripSeed == nil {
		return nil, false
	}
	return o.RoundTripSeed, true
}

// HasRoundTripSeed returns a boolean if a field has been set.
func (o *RouteRequest) HasRoundTripSeed() bool {
	if o != nil && o.RoundTripSeed != nil {
		return true
	}

	return false
}

// SetRoundTripSeed gets a reference to the given int64 and assigns it to the RoundTripSeed field.
func (o *RouteRequest) SetRoundTripSeed(v int64) {
	o.RoundTripSeed = &v
}

// GetAlternativeRouteMaxPaths returns the AlternativeRouteMaxPaths field value if set, zero value otherwise.
func (o *RouteRequest) GetAlternativeRouteMaxPaths() int32 {
	if o == nil || o.AlternativeRouteMaxPaths == nil {
		var ret int32
		return ret
	}
	return *o.AlternativeRouteMaxPaths
}

// GetAlternativeRouteMaxPathsOk returns a tuple with the AlternativeRouteMaxPaths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteRequest) GetAlternativeRouteMaxPathsOk() (*int32, bool) {
	if o == nil || o.AlternativeRouteMaxPaths == nil {
		return nil, false
	}
	return o.AlternativeRouteMaxPaths, true
}

// HasAlternativeRouteMaxPaths returns a boolean if a field has been set.
func (o *RouteRequest) HasAlternativeRouteMaxPaths() bool {
	if o != nil && o.AlternativeRouteMaxPaths != nil {
		return true
	}

	return false
}

// SetAlternativeRouteMaxPaths gets a reference to the given int32 and assigns it to the AlternativeRouteMaxPaths field.
func (o *RouteRequest) SetAlternativeRouteMaxPaths(v int32) {
	o.AlternativeRouteMaxPaths = &v
}

// GetAlternativeRouteMaxWeightFactor returns the AlternativeRouteMaxWeightFactor field value if set, zero value otherwise.
func (o *RouteRequest) GetAlternativeRouteMaxWeightFactor() float32 {
	if o == nil || o.AlternativeRouteMaxWeightFactor == nil {
		var ret float32
		return ret
	}
	return *o.AlternativeRouteMaxWeightFactor
}

// GetAlternativeRouteMaxWeightFactorOk returns a tuple with the AlternativeRouteMaxWeightFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteRequest) GetAlternativeRouteMaxWeightFactorOk() (*float32, bool) {
	if o == nil || o.AlternativeRouteMaxWeightFactor == nil {
		return nil, false
	}
	return o.AlternativeRouteMaxWeightFactor, true
}

// HasAlternativeRouteMaxWeightFactor returns a boolean if a field has been set.
func (o *RouteRequest) HasAlternativeRouteMaxWeightFactor() bool {
	if o != nil && o.AlternativeRouteMaxWeightFactor != nil {
		return true
	}

	return false
}

// SetAlternativeRouteMaxWeightFactor gets a reference to the given float32 and assigns it to the AlternativeRouteMaxWeightFactor field.
func (o *RouteRequest) SetAlternativeRouteMaxWeightFactor(v float32) {
	o.AlternativeRouteMaxWeightFactor = &v
}

// GetAlternativeRouteMaxShareFactor returns the AlternativeRouteMaxShareFactor field value if set, zero value otherwise.
func (o *RouteRequest) GetAlternativeRouteMaxShareFactor() float32 {
	if o == nil || o.AlternativeRouteMaxShareFactor == nil {
		var ret float32
		return ret
	}
	return *o.AlternativeRouteMaxShareFactor
}

// GetAlternativeRouteMaxShareFactorOk returns a tuple with the AlternativeRouteMaxShareFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteRequest) GetAlternativeRouteMaxShareFactorOk() (*float32, bool) {
	if o == nil || o.AlternativeRouteMaxShareFactor == nil {
		return nil, false
	}
	return o.AlternativeRouteMaxShareFactor, true
}

// HasAlternativeRouteMaxShareFactor returns a boolean if a field has been set.
func (o *RouteRequest) HasAlternativeRouteMaxShareFactor() bool {
	if o != nil && o.AlternativeRouteMaxShareFactor != nil {
		return true
	}

	return false
}

// SetAlternativeRouteMaxShareFactor gets a reference to the given float32 and assigns it to the AlternativeRouteMaxShareFactor field.
func (o *RouteRequest) SetAlternativeRouteMaxShareFactor(v float32) {
	o.AlternativeRouteMaxShareFactor = &v
}

func (o RouteRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Points != nil {
		toSerialize["points"] = o.Points
	}
	if o.PointHints != nil {
		toSerialize["point_hints"] = o.PointHints
	}
	if o.SnapPreventions != nil {
		toSerialize["snap_preventions"] = o.SnapPreventions
	}
	if o.Curbsides != nil {
		toSerialize["curbsides"] = o.Curbsides
	}
	if o.Vehicle != nil {
		toSerialize["vehicle"] = o.Vehicle
	}
	if o.Locale != nil {
		toSerialize["locale"] = o.Locale
	}
	if o.Elevation != nil {
		toSerialize["elevation"] = o.Elevation
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.Optimize != nil {
		toSerialize["optimize"] = o.Optimize
	}
	if o.Instructions != nil {
		toSerialize["instructions"] = o.Instructions
	}
	if o.CalcPoints != nil {
		toSerialize["calc_points"] = o.CalcPoints
	}
	if o.Debug != nil {
		toSerialize["debug"] = o.Debug
	}
	if o.PointsEncoded != nil {
		toSerialize["points_encoded"] = o.PointsEncoded
	}
	if o.ChDisable != nil {
		toSerialize["ch.disable"] = o.ChDisable
	}
	if o.Weighting != nil {
		toSerialize["weighting"] = o.Weighting
	}
	if o.Headings != nil {
		toSerialize["headings"] = o.Headings
	}
	if o.HeadingPenalty != nil {
		toSerialize["heading_penalty"] = o.HeadingPenalty
	}
	if o.PassThrough != nil {
		toSerialize["pass_through"] = o.PassThrough
	}
	if o.BlockArea != nil {
		toSerialize["block_area"] = o.BlockArea
	}
	if o.Avoid != nil {
		toSerialize["avoid"] = o.Avoid
	}
	if o.Algorithm != nil {
		toSerialize["algorithm"] = o.Algorithm
	}
	if o.RoundTripDistance != nil {
		toSerialize["round_trip.distance"] = o.RoundTripDistance
	}
	if o.RoundTripSeed != nil {
		toSerialize["round_trip.seed"] = o.RoundTripSeed
	}
	if o.AlternativeRouteMaxPaths != nil {
		toSerialize["alternative_route.max_paths"] = o.AlternativeRouteMaxPaths
	}
	if o.AlternativeRouteMaxWeightFactor != nil {
		toSerialize["alternative_route.max_weight_factor"] = o.AlternativeRouteMaxWeightFactor
	}
	if o.AlternativeRouteMaxShareFactor != nil {
		toSerialize["alternative_route.max_share_factor"] = o.AlternativeRouteMaxShareFactor
	}
	return json.Marshal(toSerialize)
}

type NullableRouteRequest struct {
	value *RouteRequest
	isSet bool
}

func (v NullableRouteRequest) Get() *RouteRequest {
	return v.value
}

func (v *NullableRouteRequest) Set(val *RouteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteRequest(val *RouteRequest) *NullableRouteRequest {
	return &NullableRouteRequest{value: val, isSet: true}
}

func (v NullableRouteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


