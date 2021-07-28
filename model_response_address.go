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

// ResponseAddress Address of activity
type ResponseAddress struct {
	// Specifies the id of the location.
	LocationId *string `json:"location_id,omitempty"`
	// Name of location.
	Name *string `json:"name,omitempty"`
	// Longitude of location.
	Lon *float64 `json:"lon,omitempty"`
	// Latitude of location.
	Lat *float64 `json:"lat,omitempty"`
	// Optional parameter. Specifies a hint for each address to better snap the coordinates (lon,lat) to road network. E.g. if there is an address or house with two or more neighboring streets you can control for which street the closest location is looked up.
	StreetHint *string `json:"street_hint,omitempty"`
	SnappedWaypoint *SnappedWaypoint `json:"snapped_waypoint,omitempty"`
}

// NewResponseAddress instantiates a new ResponseAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseAddress() *ResponseAddress {
	this := ResponseAddress{}
	return &this
}

// NewResponseAddressWithDefaults instantiates a new ResponseAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseAddressWithDefaults() *ResponseAddress {
	this := ResponseAddress{}
	return &this
}

// GetLocationId returns the LocationId field value if set, zero value otherwise.
func (o *ResponseAddress) GetLocationId() string {
	if o == nil || o.LocationId == nil {
		var ret string
		return ret
	}
	return *o.LocationId
}

// GetLocationIdOk returns a tuple with the LocationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAddress) GetLocationIdOk() (*string, bool) {
	if o == nil || o.LocationId == nil {
		return nil, false
	}
	return o.LocationId, true
}

// HasLocationId returns a boolean if a field has been set.
func (o *ResponseAddress) HasLocationId() bool {
	if o != nil && o.LocationId != nil {
		return true
	}

	return false
}

// SetLocationId gets a reference to the given string and assigns it to the LocationId field.
func (o *ResponseAddress) SetLocationId(v string) {
	o.LocationId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ResponseAddress) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAddress) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ResponseAddress) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ResponseAddress) SetName(v string) {
	o.Name = &v
}

// GetLon returns the Lon field value if set, zero value otherwise.
func (o *ResponseAddress) GetLon() float64 {
	if o == nil || o.Lon == nil {
		var ret float64
		return ret
	}
	return *o.Lon
}

// GetLonOk returns a tuple with the Lon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAddress) GetLonOk() (*float64, bool) {
	if o == nil || o.Lon == nil {
		return nil, false
	}
	return o.Lon, true
}

// HasLon returns a boolean if a field has been set.
func (o *ResponseAddress) HasLon() bool {
	if o != nil && o.Lon != nil {
		return true
	}

	return false
}

// SetLon gets a reference to the given float64 and assigns it to the Lon field.
func (o *ResponseAddress) SetLon(v float64) {
	o.Lon = &v
}

// GetLat returns the Lat field value if set, zero value otherwise.
func (o *ResponseAddress) GetLat() float64 {
	if o == nil || o.Lat == nil {
		var ret float64
		return ret
	}
	return *o.Lat
}

// GetLatOk returns a tuple with the Lat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAddress) GetLatOk() (*float64, bool) {
	if o == nil || o.Lat == nil {
		return nil, false
	}
	return o.Lat, true
}

// HasLat returns a boolean if a field has been set.
func (o *ResponseAddress) HasLat() bool {
	if o != nil && o.Lat != nil {
		return true
	}

	return false
}

// SetLat gets a reference to the given float64 and assigns it to the Lat field.
func (o *ResponseAddress) SetLat(v float64) {
	o.Lat = &v
}

// GetStreetHint returns the StreetHint field value if set, zero value otherwise.
func (o *ResponseAddress) GetStreetHint() string {
	if o == nil || o.StreetHint == nil {
		var ret string
		return ret
	}
	return *o.StreetHint
}

// GetStreetHintOk returns a tuple with the StreetHint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAddress) GetStreetHintOk() (*string, bool) {
	if o == nil || o.StreetHint == nil {
		return nil, false
	}
	return o.StreetHint, true
}

// HasStreetHint returns a boolean if a field has been set.
func (o *ResponseAddress) HasStreetHint() bool {
	if o != nil && o.StreetHint != nil {
		return true
	}

	return false
}

// SetStreetHint gets a reference to the given string and assigns it to the StreetHint field.
func (o *ResponseAddress) SetStreetHint(v string) {
	o.StreetHint = &v
}

// GetSnappedWaypoint returns the SnappedWaypoint field value if set, zero value otherwise.
func (o *ResponseAddress) GetSnappedWaypoint() SnappedWaypoint {
	if o == nil || o.SnappedWaypoint == nil {
		var ret SnappedWaypoint
		return ret
	}
	return *o.SnappedWaypoint
}

// GetSnappedWaypointOk returns a tuple with the SnappedWaypoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAddress) GetSnappedWaypointOk() (*SnappedWaypoint, bool) {
	if o == nil || o.SnappedWaypoint == nil {
		return nil, false
	}
	return o.SnappedWaypoint, true
}

// HasSnappedWaypoint returns a boolean if a field has been set.
func (o *ResponseAddress) HasSnappedWaypoint() bool {
	if o != nil && o.SnappedWaypoint != nil {
		return true
	}

	return false
}

// SetSnappedWaypoint gets a reference to the given SnappedWaypoint and assigns it to the SnappedWaypoint field.
func (o *ResponseAddress) SetSnappedWaypoint(v SnappedWaypoint) {
	o.SnappedWaypoint = &v
}

func (o ResponseAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LocationId != nil {
		toSerialize["location_id"] = o.LocationId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Lon != nil {
		toSerialize["lon"] = o.Lon
	}
	if o.Lat != nil {
		toSerialize["lat"] = o.Lat
	}
	if o.StreetHint != nil {
		toSerialize["street_hint"] = o.StreetHint
	}
	if o.SnappedWaypoint != nil {
		toSerialize["snapped_waypoint"] = o.SnappedWaypoint
	}
	return json.Marshal(toSerialize)
}

type NullableResponseAddress struct {
	value *ResponseAddress
	isSet bool
}

func (v NullableResponseAddress) Get() *ResponseAddress {
	return v.value
}

func (v *NullableResponseAddress) Set(val *ResponseAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseAddress(val *ResponseAddress) *NullableResponseAddress {
	return &NullableResponseAddress{value: val, isSet: true}
}

func (v NullableResponseAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


