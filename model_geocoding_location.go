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

// GeocodingLocation struct for GeocodingLocation
type GeocodingLocation struct {
	Point *GeocodingPoint `json:"point,omitempty"`
	// The OSM ID of the entity
	OsmId *string `json:"osm_id,omitempty"`
	// N = node, R = relation, W = way
	OsmType *string `json:"osm_type,omitempty"`
	// The OSM key of the entity
	OsmKey *string `json:"osm_key,omitempty"`
	// The name of the entity. Can be a boundary, POI, address, etc
	Name *string `json:"name,omitempty"`
	// The country of the address
	Country *string `json:"country,omitempty"`
	// The city of the address
	City *string `json:"city,omitempty"`
	// The state of the address
	State *string `json:"state,omitempty"`
	// The street of the address
	Street *string `json:"street,omitempty"`
	// The housenumber of the address
	Housenumber *string `json:"housenumber,omitempty"`
	// The postcode of the address
	Postcode *string `json:"postcode,omitempty"`
}

// NewGeocodingLocation instantiates a new GeocodingLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeocodingLocation() *GeocodingLocation {
	this := GeocodingLocation{}
	return &this
}

// NewGeocodingLocationWithDefaults instantiates a new GeocodingLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeocodingLocationWithDefaults() *GeocodingLocation {
	this := GeocodingLocation{}
	return &this
}

// GetPoint returns the Point field value if set, zero value otherwise.
func (o *GeocodingLocation) GetPoint() GeocodingPoint {
	if o == nil || o.Point == nil {
		var ret GeocodingPoint
		return ret
	}
	return *o.Point
}

// GetPointOk returns a tuple with the Point field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeocodingLocation) GetPointOk() (*GeocodingPoint, bool) {
	if o == nil || o.Point == nil {
		return nil, false
	}
	return o.Point, true
}

// HasPoint returns a boolean if a field has been set.
func (o *GeocodingLocation) HasPoint() bool {
	if o != nil && o.Point != nil {
		return true
	}

	return false
}

// SetPoint gets a reference to the given GeocodingPoint and assigns it to the Point field.
func (o *GeocodingLocation) SetPoint(v GeocodingPoint) {
	o.Point = &v
}

// GetOsmId returns the OsmId field value if set, zero value otherwise.
func (o *GeocodingLocation) GetOsmId() string {
	if o == nil || o.OsmId == nil {
		var ret string
		return ret
	}
	return *o.OsmId
}

// GetOsmIdOk returns a tuple with the OsmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeocodingLocation) GetOsmIdOk() (*string, bool) {
	if o == nil || o.OsmId == nil {
		return nil, false
	}
	return o.OsmId, true
}

// HasOsmId returns a boolean if a field has been set.
func (o *GeocodingLocation) HasOsmId() bool {
	if o != nil && o.OsmId != nil {
		return true
	}

	return false
}

// SetOsmId gets a reference to the given string and assigns it to the OsmId field.
func (o *GeocodingLocation) SetOsmId(v string) {
	o.OsmId = &v
}

// GetOsmType returns the OsmType field value if set, zero value otherwise.
func (o *GeocodingLocation) GetOsmType() string {
	if o == nil || o.OsmType == nil {
		var ret string
		return ret
	}
	return *o.OsmType
}

// GetOsmTypeOk returns a tuple with the OsmType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeocodingLocation) GetOsmTypeOk() (*string, bool) {
	if o == nil || o.OsmType == nil {
		return nil, false
	}
	return o.OsmType, true
}

// HasOsmType returns a boolean if a field has been set.
func (o *GeocodingLocation) HasOsmType() bool {
	if o != nil && o.OsmType != nil {
		return true
	}

	return false
}

// SetOsmType gets a reference to the given string and assigns it to the OsmType field.
func (o *GeocodingLocation) SetOsmType(v string) {
	o.OsmType = &v
}

// GetOsmKey returns the OsmKey field value if set, zero value otherwise.
func (o *GeocodingLocation) GetOsmKey() string {
	if o == nil || o.OsmKey == nil {
		var ret string
		return ret
	}
	return *o.OsmKey
}

// GetOsmKeyOk returns a tuple with the OsmKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeocodingLocation) GetOsmKeyOk() (*string, bool) {
	if o == nil || o.OsmKey == nil {
		return nil, false
	}
	return o.OsmKey, true
}

// HasOsmKey returns a boolean if a field has been set.
func (o *GeocodingLocation) HasOsmKey() bool {
	if o != nil && o.OsmKey != nil {
		return true
	}

	return false
}

// SetOsmKey gets a reference to the given string and assigns it to the OsmKey field.
func (o *GeocodingLocation) SetOsmKey(v string) {
	o.OsmKey = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GeocodingLocation) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeocodingLocation) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GeocodingLocation) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GeocodingLocation) SetName(v string) {
	o.Name = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *GeocodingLocation) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeocodingLocation) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *GeocodingLocation) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *GeocodingLocation) SetCountry(v string) {
	o.Country = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *GeocodingLocation) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeocodingLocation) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *GeocodingLocation) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *GeocodingLocation) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *GeocodingLocation) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeocodingLocation) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *GeocodingLocation) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *GeocodingLocation) SetState(v string) {
	o.State = &v
}

// GetStreet returns the Street field value if set, zero value otherwise.
func (o *GeocodingLocation) GetStreet() string {
	if o == nil || o.Street == nil {
		var ret string
		return ret
	}
	return *o.Street
}

// GetStreetOk returns a tuple with the Street field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeocodingLocation) GetStreetOk() (*string, bool) {
	if o == nil || o.Street == nil {
		return nil, false
	}
	return o.Street, true
}

// HasStreet returns a boolean if a field has been set.
func (o *GeocodingLocation) HasStreet() bool {
	if o != nil && o.Street != nil {
		return true
	}

	return false
}

// SetStreet gets a reference to the given string and assigns it to the Street field.
func (o *GeocodingLocation) SetStreet(v string) {
	o.Street = &v
}

// GetHousenumber returns the Housenumber field value if set, zero value otherwise.
func (o *GeocodingLocation) GetHousenumber() string {
	if o == nil || o.Housenumber == nil {
		var ret string
		return ret
	}
	return *o.Housenumber
}

// GetHousenumberOk returns a tuple with the Housenumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeocodingLocation) GetHousenumberOk() (*string, bool) {
	if o == nil || o.Housenumber == nil {
		return nil, false
	}
	return o.Housenumber, true
}

// HasHousenumber returns a boolean if a field has been set.
func (o *GeocodingLocation) HasHousenumber() bool {
	if o != nil && o.Housenumber != nil {
		return true
	}

	return false
}

// SetHousenumber gets a reference to the given string and assigns it to the Housenumber field.
func (o *GeocodingLocation) SetHousenumber(v string) {
	o.Housenumber = &v
}

// GetPostcode returns the Postcode field value if set, zero value otherwise.
func (o *GeocodingLocation) GetPostcode() string {
	if o == nil || o.Postcode == nil {
		var ret string
		return ret
	}
	return *o.Postcode
}

// GetPostcodeOk returns a tuple with the Postcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeocodingLocation) GetPostcodeOk() (*string, bool) {
	if o == nil || o.Postcode == nil {
		return nil, false
	}
	return o.Postcode, true
}

// HasPostcode returns a boolean if a field has been set.
func (o *GeocodingLocation) HasPostcode() bool {
	if o != nil && o.Postcode != nil {
		return true
	}

	return false
}

// SetPostcode gets a reference to the given string and assigns it to the Postcode field.
func (o *GeocodingLocation) SetPostcode(v string) {
	o.Postcode = &v
}

func (o GeocodingLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Point != nil {
		toSerialize["point"] = o.Point
	}
	if o.OsmId != nil {
		toSerialize["osm_id"] = o.OsmId
	}
	if o.OsmType != nil {
		toSerialize["osm_type"] = o.OsmType
	}
	if o.OsmKey != nil {
		toSerialize["osm_key"] = o.OsmKey
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Street != nil {
		toSerialize["street"] = o.Street
	}
	if o.Housenumber != nil {
		toSerialize["housenumber"] = o.Housenumber
	}
	if o.Postcode != nil {
		toSerialize["postcode"] = o.Postcode
	}
	return json.Marshal(toSerialize)
}

type NullableGeocodingLocation struct {
	value *GeocodingLocation
	isSet bool
}

func (v NullableGeocodingLocation) Get() *GeocodingLocation {
	return v.value
}

func (v *NullableGeocodingLocation) Set(val *GeocodingLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableGeocodingLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableGeocodingLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeocodingLocation(val *GeocodingLocation) *NullableGeocodingLocation {
	return &NullableGeocodingLocation{value: val, isSet: true}
}

func (v NullableGeocodingLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeocodingLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


