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

// ClusterConfiguration struct for ClusterConfiguration
type ClusterConfiguration struct {
	// Specifies the response format. You can either choose `geojson` or `json`.
	ResponseType *string `json:"response_type,omitempty"`
	Routing *ClusterConfigurationRouting `json:"routing,omitempty"`
	Clustering *ClusterConfigurationClustering `json:"clustering,omitempty"`
}

// NewClusterConfiguration instantiates a new ClusterConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterConfiguration() *ClusterConfiguration {
	this := ClusterConfiguration{}
	var responseType string = "json"
	this.ResponseType = &responseType
	return &this
}

// NewClusterConfigurationWithDefaults instantiates a new ClusterConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterConfigurationWithDefaults() *ClusterConfiguration {
	this := ClusterConfiguration{}
	var responseType string = "json"
	this.ResponseType = &responseType
	return &this
}

// GetResponseType returns the ResponseType field value if set, zero value otherwise.
func (o *ClusterConfiguration) GetResponseType() string {
	if o == nil || o.ResponseType == nil {
		var ret string
		return ret
	}
	return *o.ResponseType
}

// GetResponseTypeOk returns a tuple with the ResponseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterConfiguration) GetResponseTypeOk() (*string, bool) {
	if o == nil || o.ResponseType == nil {
		return nil, false
	}
	return o.ResponseType, true
}

// HasResponseType returns a boolean if a field has been set.
func (o *ClusterConfiguration) HasResponseType() bool {
	if o != nil && o.ResponseType != nil {
		return true
	}

	return false
}

// SetResponseType gets a reference to the given string and assigns it to the ResponseType field.
func (o *ClusterConfiguration) SetResponseType(v string) {
	o.ResponseType = &v
}

// GetRouting returns the Routing field value if set, zero value otherwise.
func (o *ClusterConfiguration) GetRouting() ClusterConfigurationRouting {
	if o == nil || o.Routing == nil {
		var ret ClusterConfigurationRouting
		return ret
	}
	return *o.Routing
}

// GetRoutingOk returns a tuple with the Routing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterConfiguration) GetRoutingOk() (*ClusterConfigurationRouting, bool) {
	if o == nil || o.Routing == nil {
		return nil, false
	}
	return o.Routing, true
}

// HasRouting returns a boolean if a field has been set.
func (o *ClusterConfiguration) HasRouting() bool {
	if o != nil && o.Routing != nil {
		return true
	}

	return false
}

// SetRouting gets a reference to the given ClusterConfigurationRouting and assigns it to the Routing field.
func (o *ClusterConfiguration) SetRouting(v ClusterConfigurationRouting) {
	o.Routing = &v
}

// GetClustering returns the Clustering field value if set, zero value otherwise.
func (o *ClusterConfiguration) GetClustering() ClusterConfigurationClustering {
	if o == nil || o.Clustering == nil {
		var ret ClusterConfigurationClustering
		return ret
	}
	return *o.Clustering
}

// GetClusteringOk returns a tuple with the Clustering field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterConfiguration) GetClusteringOk() (*ClusterConfigurationClustering, bool) {
	if o == nil || o.Clustering == nil {
		return nil, false
	}
	return o.Clustering, true
}

// HasClustering returns a boolean if a field has been set.
func (o *ClusterConfiguration) HasClustering() bool {
	if o != nil && o.Clustering != nil {
		return true
	}

	return false
}

// SetClustering gets a reference to the given ClusterConfigurationClustering and assigns it to the Clustering field.
func (o *ClusterConfiguration) SetClustering(v ClusterConfigurationClustering) {
	o.Clustering = &v
}

func (o ClusterConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseType != nil {
		toSerialize["response_type"] = o.ResponseType
	}
	if o.Routing != nil {
		toSerialize["routing"] = o.Routing
	}
	if o.Clustering != nil {
		toSerialize["clustering"] = o.Clustering
	}
	return json.Marshal(toSerialize)
}

type NullableClusterConfiguration struct {
	value *ClusterConfiguration
	isSet bool
}

func (v NullableClusterConfiguration) Get() *ClusterConfiguration {
	return v.value
}

func (v *NullableClusterConfiguration) Set(val *ClusterConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterConfiguration(val *ClusterConfiguration) *NullableClusterConfiguration {
	return &NullableClusterConfiguration{value: val, isSet: true}
}

func (v NullableClusterConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


