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

// RouteResponsePathInstructions struct for RouteResponsePathInstructions
type RouteResponsePathInstructions struct {
	// A description what the user has to do in order to follow the route. The language depends on the locale parameter. 
	Text *string `json:"text,omitempty"`
	// The name of the street to turn onto in order to follow the route. 
	StreetName *string `json:"street_name,omitempty"`
	// The distance for this instruction, in meters. 
	Distance *float64 `json:"distance,omitempty"`
	// The duration for this instruction, in milliseconds. 
	Time *int32 `json:"time,omitempty"`
	// Two indices into `points`, referring to the beginning and the end of the segment of the route this instruction refers to. 
	Interval *[]int32 `json:"interval,omitempty"`
	// A number which specifies the sign to show:  | sign | description  | |---|---| |-98| an U-turn without the knowledge if it is a right or left U-turn | | -8| a left U-turn | | -7| keep left | | -6| **not yet used**: leave roundabout | | -3| turn sharp left | | -2| turn left | | -1| turn slight left | |  0| continue on street | |  1| turn slight right | |  2| turn right | |  3| turn sharp right | |  4| the finish instruction before the last point | |  5| the instruction before a via point | |  6| the instruction before entering a roundabout | |  7| keep right | |  8| a right U-turn | |  *| **For future compatibility** it is important that all clients are able to handle also unknown instruction sign numbers 
	Sign *int32 `json:"sign,omitempty"`
	// Only available for roundabout instructions (sign is 6). The count of exits at which the route leaves the roundabout. 
	ExitNumber *int32 `json:"exit_number,omitempty"`
	// Only available for roundabout instructions (sign is 6). The radian of the route within the roundabout `0 < r < 2*PI` for clockwise and `-2*PI < r < 0` for counterclockwise turns. 
	TurnAngle *float64 `json:"turn_angle,omitempty"`
}

// NewRouteResponsePathInstructions instantiates a new RouteResponsePathInstructions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRouteResponsePathInstructions() *RouteResponsePathInstructions {
	this := RouteResponsePathInstructions{}
	return &this
}

// NewRouteResponsePathInstructionsWithDefaults instantiates a new RouteResponsePathInstructions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteResponsePathInstructionsWithDefaults() *RouteResponsePathInstructions {
	this := RouteResponsePathInstructions{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *RouteResponsePathInstructions) GetText() string {
	if o == nil || o.Text == nil {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteResponsePathInstructions) GetTextOk() (*string, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *RouteResponsePathInstructions) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *RouteResponsePathInstructions) SetText(v string) {
	o.Text = &v
}

// GetStreetName returns the StreetName field value if set, zero value otherwise.
func (o *RouteResponsePathInstructions) GetStreetName() string {
	if o == nil || o.StreetName == nil {
		var ret string
		return ret
	}
	return *o.StreetName
}

// GetStreetNameOk returns a tuple with the StreetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteResponsePathInstructions) GetStreetNameOk() (*string, bool) {
	if o == nil || o.StreetName == nil {
		return nil, false
	}
	return o.StreetName, true
}

// HasStreetName returns a boolean if a field has been set.
func (o *RouteResponsePathInstructions) HasStreetName() bool {
	if o != nil && o.StreetName != nil {
		return true
	}

	return false
}

// SetStreetName gets a reference to the given string and assigns it to the StreetName field.
func (o *RouteResponsePathInstructions) SetStreetName(v string) {
	o.StreetName = &v
}

// GetDistance returns the Distance field value if set, zero value otherwise.
func (o *RouteResponsePathInstructions) GetDistance() float64 {
	if o == nil || o.Distance == nil {
		var ret float64
		return ret
	}
	return *o.Distance
}

// GetDistanceOk returns a tuple with the Distance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteResponsePathInstructions) GetDistanceOk() (*float64, bool) {
	if o == nil || o.Distance == nil {
		return nil, false
	}
	return o.Distance, true
}

// HasDistance returns a boolean if a field has been set.
func (o *RouteResponsePathInstructions) HasDistance() bool {
	if o != nil && o.Distance != nil {
		return true
	}

	return false
}

// SetDistance gets a reference to the given float64 and assigns it to the Distance field.
func (o *RouteResponsePathInstructions) SetDistance(v float64) {
	o.Distance = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *RouteResponsePathInstructions) GetTime() int32 {
	if o == nil || o.Time == nil {
		var ret int32
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteResponsePathInstructions) GetTimeOk() (*int32, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *RouteResponsePathInstructions) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given int32 and assigns it to the Time field.
func (o *RouteResponsePathInstructions) SetTime(v int32) {
	o.Time = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *RouteResponsePathInstructions) GetInterval() []int32 {
	if o == nil || o.Interval == nil {
		var ret []int32
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteResponsePathInstructions) GetIntervalOk() (*[]int32, bool) {
	if o == nil || o.Interval == nil {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *RouteResponsePathInstructions) HasInterval() bool {
	if o != nil && o.Interval != nil {
		return true
	}

	return false
}

// SetInterval gets a reference to the given []int32 and assigns it to the Interval field.
func (o *RouteResponsePathInstructions) SetInterval(v []int32) {
	o.Interval = &v
}

// GetSign returns the Sign field value if set, zero value otherwise.
func (o *RouteResponsePathInstructions) GetSign() int32 {
	if o == nil || o.Sign == nil {
		var ret int32
		return ret
	}
	return *o.Sign
}

// GetSignOk returns a tuple with the Sign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteResponsePathInstructions) GetSignOk() (*int32, bool) {
	if o == nil || o.Sign == nil {
		return nil, false
	}
	return o.Sign, true
}

// HasSign returns a boolean if a field has been set.
func (o *RouteResponsePathInstructions) HasSign() bool {
	if o != nil && o.Sign != nil {
		return true
	}

	return false
}

// SetSign gets a reference to the given int32 and assigns it to the Sign field.
func (o *RouteResponsePathInstructions) SetSign(v int32) {
	o.Sign = &v
}

// GetExitNumber returns the ExitNumber field value if set, zero value otherwise.
func (o *RouteResponsePathInstructions) GetExitNumber() int32 {
	if o == nil || o.ExitNumber == nil {
		var ret int32
		return ret
	}
	return *o.ExitNumber
}

// GetExitNumberOk returns a tuple with the ExitNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteResponsePathInstructions) GetExitNumberOk() (*int32, bool) {
	if o == nil || o.ExitNumber == nil {
		return nil, false
	}
	return o.ExitNumber, true
}

// HasExitNumber returns a boolean if a field has been set.
func (o *RouteResponsePathInstructions) HasExitNumber() bool {
	if o != nil && o.ExitNumber != nil {
		return true
	}

	return false
}

// SetExitNumber gets a reference to the given int32 and assigns it to the ExitNumber field.
func (o *RouteResponsePathInstructions) SetExitNumber(v int32) {
	o.ExitNumber = &v
}

// GetTurnAngle returns the TurnAngle field value if set, zero value otherwise.
func (o *RouteResponsePathInstructions) GetTurnAngle() float64 {
	if o == nil || o.TurnAngle == nil {
		var ret float64
		return ret
	}
	return *o.TurnAngle
}

// GetTurnAngleOk returns a tuple with the TurnAngle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteResponsePathInstructions) GetTurnAngleOk() (*float64, bool) {
	if o == nil || o.TurnAngle == nil {
		return nil, false
	}
	return o.TurnAngle, true
}

// HasTurnAngle returns a boolean if a field has been set.
func (o *RouteResponsePathInstructions) HasTurnAngle() bool {
	if o != nil && o.TurnAngle != nil {
		return true
	}

	return false
}

// SetTurnAngle gets a reference to the given float64 and assigns it to the TurnAngle field.
func (o *RouteResponsePathInstructions) SetTurnAngle(v float64) {
	o.TurnAngle = &v
}

func (o RouteResponsePathInstructions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	if o.StreetName != nil {
		toSerialize["street_name"] = o.StreetName
	}
	if o.Distance != nil {
		toSerialize["distance"] = o.Distance
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	if o.Interval != nil {
		toSerialize["interval"] = o.Interval
	}
	if o.Sign != nil {
		toSerialize["sign"] = o.Sign
	}
	if o.ExitNumber != nil {
		toSerialize["exit_number"] = o.ExitNumber
	}
	if o.TurnAngle != nil {
		toSerialize["turn_angle"] = o.TurnAngle
	}
	return json.Marshal(toSerialize)
}

type NullableRouteResponsePathInstructions struct {
	value *RouteResponsePathInstructions
	isSet bool
}

func (v NullableRouteResponsePathInstructions) Get() *RouteResponsePathInstructions {
	return v.value
}

func (v *NullableRouteResponsePathInstructions) Set(val *RouteResponsePathInstructions) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteResponsePathInstructions) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteResponsePathInstructions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteResponsePathInstructions(val *RouteResponsePathInstructions) *NullableRouteResponsePathInstructions {
	return &NullableRouteResponsePathInstructions{value: val, isSet: true}
}

func (v NullableRouteResponsePathInstructions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteResponsePathInstructions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


