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

// MatrixResponse struct for MatrixResponse
type MatrixResponse struct {
	// The distance matrix for the specified points in the same order as the time matrix. The distances are in meters. If `fail_fast=false` the matrix will contain `null` for connections that could not be found.
	Distances *[][]float32 `json:"distances,omitempty"`
	// The time matrix for the specified points in the order [[from1->to1, from1->to2, ...], [from2->to1, from2->to2, ...], ...]. The times are in seconds. If `fail_fast=false` the matrix will contain `null` for connections that could not be found.
	Times *[][]float32 `json:"times,omitempty"`
	// The weight matrix for the specified points in the same order as the time matrix. The weights for different vehicles can have a different unit but the weights array is perfectly suited as input for Vehicle Routing Problems as it is currently faster to calculate. If `fail_fast=false` the matrix will contain `null` for connections that could not be found.
	Weights *[][]float64 `json:"weights,omitempty"`
	Info *ResponseInfo `json:"info,omitempty"`
	// Optional. Additional response data.
	Hints *[]MatrixResponseHints `json:"hints,omitempty"`
}

// NewMatrixResponse instantiates a new MatrixResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMatrixResponse() *MatrixResponse {
	this := MatrixResponse{}
	return &this
}

// NewMatrixResponseWithDefaults instantiates a new MatrixResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMatrixResponseWithDefaults() *MatrixResponse {
	this := MatrixResponse{}
	return &this
}

// GetDistances returns the Distances field value if set, zero value otherwise.
func (o *MatrixResponse) GetDistances() [][]float32 {
	if o == nil || o.Distances == nil {
		var ret [][]float32
		return ret
	}
	return *o.Distances
}

// GetDistancesOk returns a tuple with the Distances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatrixResponse) GetDistancesOk() (*[][]float32, bool) {
	if o == nil || o.Distances == nil {
		return nil, false
	}
	return o.Distances, true
}

// HasDistances returns a boolean if a field has been set.
func (o *MatrixResponse) HasDistances() bool {
	if o != nil && o.Distances != nil {
		return true
	}

	return false
}

// SetDistances gets a reference to the given [][]float32 and assigns it to the Distances field.
func (o *MatrixResponse) SetDistances(v [][]float32) {
	o.Distances = &v
}

// GetTimes returns the Times field value if set, zero value otherwise.
func (o *MatrixResponse) GetTimes() [][]float32 {
	if o == nil || o.Times == nil {
		var ret [][]float32
		return ret
	}
	return *o.Times
}

// GetTimesOk returns a tuple with the Times field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatrixResponse) GetTimesOk() (*[][]float32, bool) {
	if o == nil || o.Times == nil {
		return nil, false
	}
	return o.Times, true
}

// HasTimes returns a boolean if a field has been set.
func (o *MatrixResponse) HasTimes() bool {
	if o != nil && o.Times != nil {
		return true
	}

	return false
}

// SetTimes gets a reference to the given [][]float32 and assigns it to the Times field.
func (o *MatrixResponse) SetTimes(v [][]float32) {
	o.Times = &v
}

// GetWeights returns the Weights field value if set, zero value otherwise.
func (o *MatrixResponse) GetWeights() [][]float64 {
	if o == nil || o.Weights == nil {
		var ret [][]float64
		return ret
	}
	return *o.Weights
}

// GetWeightsOk returns a tuple with the Weights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatrixResponse) GetWeightsOk() (*[][]float64, bool) {
	if o == nil || o.Weights == nil {
		return nil, false
	}
	return o.Weights, true
}

// HasWeights returns a boolean if a field has been set.
func (o *MatrixResponse) HasWeights() bool {
	if o != nil && o.Weights != nil {
		return true
	}

	return false
}

// SetWeights gets a reference to the given [][]float64 and assigns it to the Weights field.
func (o *MatrixResponse) SetWeights(v [][]float64) {
	o.Weights = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *MatrixResponse) GetInfo() ResponseInfo {
	if o == nil || o.Info == nil {
		var ret ResponseInfo
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatrixResponse) GetInfoOk() (*ResponseInfo, bool) {
	if o == nil || o.Info == nil {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *MatrixResponse) HasInfo() bool {
	if o != nil && o.Info != nil {
		return true
	}

	return false
}

// SetInfo gets a reference to the given ResponseInfo and assigns it to the Info field.
func (o *MatrixResponse) SetInfo(v ResponseInfo) {
	o.Info = &v
}

// GetHints returns the Hints field value if set, zero value otherwise.
func (o *MatrixResponse) GetHints() []MatrixResponseHints {
	if o == nil || o.Hints == nil {
		var ret []MatrixResponseHints
		return ret
	}
	return *o.Hints
}

// GetHintsOk returns a tuple with the Hints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatrixResponse) GetHintsOk() (*[]MatrixResponseHints, bool) {
	if o == nil || o.Hints == nil {
		return nil, false
	}
	return o.Hints, true
}

// HasHints returns a boolean if a field has been set.
func (o *MatrixResponse) HasHints() bool {
	if o != nil && o.Hints != nil {
		return true
	}

	return false
}

// SetHints gets a reference to the given []MatrixResponseHints and assigns it to the Hints field.
func (o *MatrixResponse) SetHints(v []MatrixResponseHints) {
	o.Hints = &v
}

func (o MatrixResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Distances != nil {
		toSerialize["distances"] = o.Distances
	}
	if o.Times != nil {
		toSerialize["times"] = o.Times
	}
	if o.Weights != nil {
		toSerialize["weights"] = o.Weights
	}
	if o.Info != nil {
		toSerialize["info"] = o.Info
	}
	if o.Hints != nil {
		toSerialize["hints"] = o.Hints
	}
	return json.Marshal(toSerialize)
}

type NullableMatrixResponse struct {
	value *MatrixResponse
	isSet bool
}

func (v NullableMatrixResponse) Get() *MatrixResponse {
	return v.value
}

func (v *NullableMatrixResponse) Set(val *MatrixResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMatrixResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMatrixResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMatrixResponse(val *MatrixResponse) *NullableMatrixResponse {
	return &NullableMatrixResponse{value: val, isSet: true}
}

func (v NullableMatrixResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMatrixResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


