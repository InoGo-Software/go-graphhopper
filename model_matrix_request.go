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

// MatrixRequest struct for MatrixRequest
type MatrixRequest struct {
	// The starting points for the routes in an array of `[longitude,latitude]`. For instance, if you want to calculate three routes from point A such as A->1, A->2, A->3 then you have one `from_point` parameter and three `to_point` parameters.
	FromPoints *[][]float64 `json:"from_points,omitempty"`
	// The destination points for the routes in an array of `[longitude,latitude]`.
	ToPoints *[][]float64 `json:"to_points,omitempty"`
	// See `point_hints`of symmetrical matrix
	FromPointHints *[]string `json:"from_point_hints,omitempty"`
	// See `point_hints`of symmetrical matrix
	ToPointHints *[]string `json:"to_point_hints,omitempty"`
	// See `snap_preventions` of symmetrical matrix
	SnapPreventions *[]string `json:"snap_preventions,omitempty"`
	// See `curbsides`of symmetrical matrix
	FromCurbsides *[]string `json:"from_curbsides,omitempty"`
	// See `curbsides`of symmetrical matrix
	ToCurbsides *[]string `json:"to_curbsides,omitempty"`
	// Specifies which matrices should be included in the response. Specify one or more of the following options `weights`, `times`, `distances`. The units of the entries of `distances` are meters, of `times` are seconds and of `weights` is arbitrary and it can differ for different vehicles or versions of this API.
	OutArrays *[]string `json:"out_arrays,omitempty"`
	Vehicle *VehicleProfileId `json:"vehicle,omitempty"`
	// Specifies whether or not the matrix calculation should return with an error as soon as possible in case some points cannot be found or some points are not connected. If set to `false` the time/weight/distance matrix will be calculated for all valid points and contain the `null` value for all entries that could not be calculated. The `hint` field of the response will also contain additional information about what went wrong (see its documentation).
	FailFast *bool `json:"fail_fast,omitempty"`
	// Specifies if turn restrictions should be considered. Enabling this option increases the matrix computation time. Only supported for motor vehicles and OpenStreetMap.
	TurnCosts *bool `json:"turn_costs,omitempty"`
}

// NewMatrixRequest instantiates a new MatrixRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMatrixRequest() *MatrixRequest {
	this := MatrixRequest{}
	var failFast bool = true
	this.FailFast = &failFast
	var turnCosts bool = false
	this.TurnCosts = &turnCosts
	return &this
}

// NewMatrixRequestWithDefaults instantiates a new MatrixRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMatrixRequestWithDefaults() *MatrixRequest {
	this := MatrixRequest{}
	var failFast bool = true
	this.FailFast = &failFast
	var turnCosts bool = false
	this.TurnCosts = &turnCosts
	return &this
}

// GetFromPoints returns the FromPoints field value if set, zero value otherwise.
func (o *MatrixRequest) GetFromPoints() [][]float64 {
	if o == nil || o.FromPoints == nil {
		var ret [][]float64
		return ret
	}
	return *o.FromPoints
}

// GetFromPointsOk returns a tuple with the FromPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatrixRequest) GetFromPointsOk() (*[][]float64, bool) {
	if o == nil || o.FromPoints == nil {
		return nil, false
	}
	return o.FromPoints, true
}

// HasFromPoints returns a boolean if a field has been set.
func (o *MatrixRequest) HasFromPoints() bool {
	if o != nil && o.FromPoints != nil {
		return true
	}

	return false
}

// SetFromPoints gets a reference to the given [][]float64 and assigns it to the FromPoints field.
func (o *MatrixRequest) SetFromPoints(v [][]float64) {
	o.FromPoints = &v
}

// GetToPoints returns the ToPoints field value if set, zero value otherwise.
func (o *MatrixRequest) GetToPoints() [][]float64 {
	if o == nil || o.ToPoints == nil {
		var ret [][]float64
		return ret
	}
	return *o.ToPoints
}

// GetToPointsOk returns a tuple with the ToPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatrixRequest) GetToPointsOk() (*[][]float64, bool) {
	if o == nil || o.ToPoints == nil {
		return nil, false
	}
	return o.ToPoints, true
}

// HasToPoints returns a boolean if a field has been set.
func (o *MatrixRequest) HasToPoints() bool {
	if o != nil && o.ToPoints != nil {
		return true
	}

	return false
}

// SetToPoints gets a reference to the given [][]float64 and assigns it to the ToPoints field.
func (o *MatrixRequest) SetToPoints(v [][]float64) {
	o.ToPoints = &v
}

// GetFromPointHints returns the FromPointHints field value if set, zero value otherwise.
func (o *MatrixRequest) GetFromPointHints() []string {
	if o == nil || o.FromPointHints == nil {
		var ret []string
		return ret
	}
	return *o.FromPointHints
}

// GetFromPointHintsOk returns a tuple with the FromPointHints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatrixRequest) GetFromPointHintsOk() (*[]string, bool) {
	if o == nil || o.FromPointHints == nil {
		return nil, false
	}
	return o.FromPointHints, true
}

// HasFromPointHints returns a boolean if a field has been set.
func (o *MatrixRequest) HasFromPointHints() bool {
	if o != nil && o.FromPointHints != nil {
		return true
	}

	return false
}

// SetFromPointHints gets a reference to the given []string and assigns it to the FromPointHints field.
func (o *MatrixRequest) SetFromPointHints(v []string) {
	o.FromPointHints = &v
}

// GetToPointHints returns the ToPointHints field value if set, zero value otherwise.
func (o *MatrixRequest) GetToPointHints() []string {
	if o == nil || o.ToPointHints == nil {
		var ret []string
		return ret
	}
	return *o.ToPointHints
}

// GetToPointHintsOk returns a tuple with the ToPointHints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatrixRequest) GetToPointHintsOk() (*[]string, bool) {
	if o == nil || o.ToPointHints == nil {
		return nil, false
	}
	return o.ToPointHints, true
}

// HasToPointHints returns a boolean if a field has been set.
func (o *MatrixRequest) HasToPointHints() bool {
	if o != nil && o.ToPointHints != nil {
		return true
	}

	return false
}

// SetToPointHints gets a reference to the given []string and assigns it to the ToPointHints field.
func (o *MatrixRequest) SetToPointHints(v []string) {
	o.ToPointHints = &v
}

// GetSnapPreventions returns the SnapPreventions field value if set, zero value otherwise.
func (o *MatrixRequest) GetSnapPreventions() []string {
	if o == nil || o.SnapPreventions == nil {
		var ret []string
		return ret
	}
	return *o.SnapPreventions
}

// GetSnapPreventionsOk returns a tuple with the SnapPreventions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatrixRequest) GetSnapPreventionsOk() (*[]string, bool) {
	if o == nil || o.SnapPreventions == nil {
		return nil, false
	}
	return o.SnapPreventions, true
}

// HasSnapPreventions returns a boolean if a field has been set.
func (o *MatrixRequest) HasSnapPreventions() bool {
	if o != nil && o.SnapPreventions != nil {
		return true
	}

	return false
}

// SetSnapPreventions gets a reference to the given []string and assigns it to the SnapPreventions field.
func (o *MatrixRequest) SetSnapPreventions(v []string) {
	o.SnapPreventions = &v
}

// GetFromCurbsides returns the FromCurbsides field value if set, zero value otherwise.
func (o *MatrixRequest) GetFromCurbsides() []string {
	if o == nil || o.FromCurbsides == nil {
		var ret []string
		return ret
	}
	return *o.FromCurbsides
}

// GetFromCurbsidesOk returns a tuple with the FromCurbsides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatrixRequest) GetFromCurbsidesOk() (*[]string, bool) {
	if o == nil || o.FromCurbsides == nil {
		return nil, false
	}
	return o.FromCurbsides, true
}

// HasFromCurbsides returns a boolean if a field has been set.
func (o *MatrixRequest) HasFromCurbsides() bool {
	if o != nil && o.FromCurbsides != nil {
		return true
	}

	return false
}

// SetFromCurbsides gets a reference to the given []string and assigns it to the FromCurbsides field.
func (o *MatrixRequest) SetFromCurbsides(v []string) {
	o.FromCurbsides = &v
}

// GetToCurbsides returns the ToCurbsides field value if set, zero value otherwise.
func (o *MatrixRequest) GetToCurbsides() []string {
	if o == nil || o.ToCurbsides == nil {
		var ret []string
		return ret
	}
	return *o.ToCurbsides
}

// GetToCurbsidesOk returns a tuple with the ToCurbsides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatrixRequest) GetToCurbsidesOk() (*[]string, bool) {
	if o == nil || o.ToCurbsides == nil {
		return nil, false
	}
	return o.ToCurbsides, true
}

// HasToCurbsides returns a boolean if a field has been set.
func (o *MatrixRequest) HasToCurbsides() bool {
	if o != nil && o.ToCurbsides != nil {
		return true
	}

	return false
}

// SetToCurbsides gets a reference to the given []string and assigns it to the ToCurbsides field.
func (o *MatrixRequest) SetToCurbsides(v []string) {
	o.ToCurbsides = &v
}

// GetOutArrays returns the OutArrays field value if set, zero value otherwise.
func (o *MatrixRequest) GetOutArrays() []string {
	if o == nil || o.OutArrays == nil {
		var ret []string
		return ret
	}
	return *o.OutArrays
}

// GetOutArraysOk returns a tuple with the OutArrays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatrixRequest) GetOutArraysOk() (*[]string, bool) {
	if o == nil || o.OutArrays == nil {
		return nil, false
	}
	return o.OutArrays, true
}

// HasOutArrays returns a boolean if a field has been set.
func (o *MatrixRequest) HasOutArrays() bool {
	if o != nil && o.OutArrays != nil {
		return true
	}

	return false
}

// SetOutArrays gets a reference to the given []string and assigns it to the OutArrays field.
func (o *MatrixRequest) SetOutArrays(v []string) {
	o.OutArrays = &v
}

// GetVehicle returns the Vehicle field value if set, zero value otherwise.
func (o *MatrixRequest) GetVehicle() VehicleProfileId {
	if o == nil || o.Vehicle == nil {
		var ret VehicleProfileId
		return ret
	}
	return *o.Vehicle
}

// GetVehicleOk returns a tuple with the Vehicle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatrixRequest) GetVehicleOk() (*VehicleProfileId, bool) {
	if o == nil || o.Vehicle == nil {
		return nil, false
	}
	return o.Vehicle, true
}

// HasVehicle returns a boolean if a field has been set.
func (o *MatrixRequest) HasVehicle() bool {
	if o != nil && o.Vehicle != nil {
		return true
	}

	return false
}

// SetVehicle gets a reference to the given VehicleProfileId and assigns it to the Vehicle field.
func (o *MatrixRequest) SetVehicle(v VehicleProfileId) {
	o.Vehicle = &v
}

// GetFailFast returns the FailFast field value if set, zero value otherwise.
func (o *MatrixRequest) GetFailFast() bool {
	if o == nil || o.FailFast == nil {
		var ret bool
		return ret
	}
	return *o.FailFast
}

// GetFailFastOk returns a tuple with the FailFast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatrixRequest) GetFailFastOk() (*bool, bool) {
	if o == nil || o.FailFast == nil {
		return nil, false
	}
	return o.FailFast, true
}

// HasFailFast returns a boolean if a field has been set.
func (o *MatrixRequest) HasFailFast() bool {
	if o != nil && o.FailFast != nil {
		return true
	}

	return false
}

// SetFailFast gets a reference to the given bool and assigns it to the FailFast field.
func (o *MatrixRequest) SetFailFast(v bool) {
	o.FailFast = &v
}

// GetTurnCosts returns the TurnCosts field value if set, zero value otherwise.
func (o *MatrixRequest) GetTurnCosts() bool {
	if o == nil || o.TurnCosts == nil {
		var ret bool
		return ret
	}
	return *o.TurnCosts
}

// GetTurnCostsOk returns a tuple with the TurnCosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatrixRequest) GetTurnCostsOk() (*bool, bool) {
	if o == nil || o.TurnCosts == nil {
		return nil, false
	}
	return o.TurnCosts, true
}

// HasTurnCosts returns a boolean if a field has been set.
func (o *MatrixRequest) HasTurnCosts() bool {
	if o != nil && o.TurnCosts != nil {
		return true
	}

	return false
}

// SetTurnCosts gets a reference to the given bool and assigns it to the TurnCosts field.
func (o *MatrixRequest) SetTurnCosts(v bool) {
	o.TurnCosts = &v
}

func (o MatrixRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FromPoints != nil {
		toSerialize["from_points"] = o.FromPoints
	}
	if o.ToPoints != nil {
		toSerialize["to_points"] = o.ToPoints
	}
	if o.FromPointHints != nil {
		toSerialize["from_point_hints"] = o.FromPointHints
	}
	if o.ToPointHints != nil {
		toSerialize["to_point_hints"] = o.ToPointHints
	}
	if o.SnapPreventions != nil {
		toSerialize["snap_preventions"] = o.SnapPreventions
	}
	if o.FromCurbsides != nil {
		toSerialize["from_curbsides"] = o.FromCurbsides
	}
	if o.ToCurbsides != nil {
		toSerialize["to_curbsides"] = o.ToCurbsides
	}
	if o.OutArrays != nil {
		toSerialize["out_arrays"] = o.OutArrays
	}
	if o.Vehicle != nil {
		toSerialize["vehicle"] = o.Vehicle
	}
	if o.FailFast != nil {
		toSerialize["fail_fast"] = o.FailFast
	}
	if o.TurnCosts != nil {
		toSerialize["turn_costs"] = o.TurnCosts
	}
	return json.Marshal(toSerialize)
}

type NullableMatrixRequest struct {
	value *MatrixRequest
	isSet bool
}

func (v NullableMatrixRequest) Get() *MatrixRequest {
	return v.value
}

func (v *NullableMatrixRequest) Set(val *MatrixRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMatrixRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMatrixRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMatrixRequest(val *MatrixRequest) *NullableMatrixRequest {
	return &NullableMatrixRequest{value: val, isSet: true}
}

func (v NullableMatrixRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMatrixRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


