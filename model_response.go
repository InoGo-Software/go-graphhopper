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

// Response struct for Response
type Response struct {
	Copyrights *[]string `json:"copyrights,omitempty"`
	// Indicates the current status of the job
	Status *string `json:"status,omitempty"`
	// Waiting time in ms
	WaitingTimeInQueue *int64 `json:"waiting_time_in_queue,omitempty"`
	// Processing time in ms. If job is still waiting in queue, processing_time is 0
	ProcessingTime *int64 `json:"processing_time,omitempty"`
	Solution *Solution `json:"solution,omitempty"`
}

// NewResponse instantiates a new Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponse() *Response {
	this := Response{}
	return &this
}

// NewResponseWithDefaults instantiates a new Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseWithDefaults() *Response {
	this := Response{}
	return &this
}

// GetCopyrights returns the Copyrights field value if set, zero value otherwise.
func (o *Response) GetCopyrights() []string {
	if o == nil || o.Copyrights == nil {
		var ret []string
		return ret
	}
	return *o.Copyrights
}

// GetCopyrightsOk returns a tuple with the Copyrights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Response) GetCopyrightsOk() (*[]string, bool) {
	if o == nil || o.Copyrights == nil {
		return nil, false
	}
	return o.Copyrights, true
}

// HasCopyrights returns a boolean if a field has been set.
func (o *Response) HasCopyrights() bool {
	if o != nil && o.Copyrights != nil {
		return true
	}

	return false
}

// SetCopyrights gets a reference to the given []string and assigns it to the Copyrights field.
func (o *Response) SetCopyrights(v []string) {
	o.Copyrights = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Response) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Response) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Response) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Response) SetStatus(v string) {
	o.Status = &v
}

// GetWaitingTimeInQueue returns the WaitingTimeInQueue field value if set, zero value otherwise.
func (o *Response) GetWaitingTimeInQueue() int64 {
	if o == nil || o.WaitingTimeInQueue == nil {
		var ret int64
		return ret
	}
	return *o.WaitingTimeInQueue
}

// GetWaitingTimeInQueueOk returns a tuple with the WaitingTimeInQueue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Response) GetWaitingTimeInQueueOk() (*int64, bool) {
	if o == nil || o.WaitingTimeInQueue == nil {
		return nil, false
	}
	return o.WaitingTimeInQueue, true
}

// HasWaitingTimeInQueue returns a boolean if a field has been set.
func (o *Response) HasWaitingTimeInQueue() bool {
	if o != nil && o.WaitingTimeInQueue != nil {
		return true
	}

	return false
}

// SetWaitingTimeInQueue gets a reference to the given int64 and assigns it to the WaitingTimeInQueue field.
func (o *Response) SetWaitingTimeInQueue(v int64) {
	o.WaitingTimeInQueue = &v
}

// GetProcessingTime returns the ProcessingTime field value if set, zero value otherwise.
func (o *Response) GetProcessingTime() int64 {
	if o == nil || o.ProcessingTime == nil {
		var ret int64
		return ret
	}
	return *o.ProcessingTime
}

// GetProcessingTimeOk returns a tuple with the ProcessingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Response) GetProcessingTimeOk() (*int64, bool) {
	if o == nil || o.ProcessingTime == nil {
		return nil, false
	}
	return o.ProcessingTime, true
}

// HasProcessingTime returns a boolean if a field has been set.
func (o *Response) HasProcessingTime() bool {
	if o != nil && o.ProcessingTime != nil {
		return true
	}

	return false
}

// SetProcessingTime gets a reference to the given int64 and assigns it to the ProcessingTime field.
func (o *Response) SetProcessingTime(v int64) {
	o.ProcessingTime = &v
}

// GetSolution returns the Solution field value if set, zero value otherwise.
func (o *Response) GetSolution() Solution {
	if o == nil || o.Solution == nil {
		var ret Solution
		return ret
	}
	return *o.Solution
}

// GetSolutionOk returns a tuple with the Solution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Response) GetSolutionOk() (*Solution, bool) {
	if o == nil || o.Solution == nil {
		return nil, false
	}
	return o.Solution, true
}

// HasSolution returns a boolean if a field has been set.
func (o *Response) HasSolution() bool {
	if o != nil && o.Solution != nil {
		return true
	}

	return false
}

// SetSolution gets a reference to the given Solution and assigns it to the Solution field.
func (o *Response) SetSolution(v Solution) {
	o.Solution = &v
}

func (o Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Copyrights != nil {
		toSerialize["copyrights"] = o.Copyrights
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.WaitingTimeInQueue != nil {
		toSerialize["waiting_time_in_queue"] = o.WaitingTimeInQueue
	}
	if o.ProcessingTime != nil {
		toSerialize["processing_time"] = o.ProcessingTime
	}
	if o.Solution != nil {
		toSerialize["solution"] = o.Solution
	}
	return json.Marshal(toSerialize)
}

type NullableResponse struct {
	value *Response
	isSet bool
}

func (v NullableResponse) Get() *Response {
	return v.value
}

func (v *NullableResponse) Set(val *Response) {
	v.value = val
	v.isSet = true
}

func (v NullableResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponse(val *Response) *NullableResponse {
	return &NullableResponse{value: val, isSet: true}
}

func (v NullableResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


