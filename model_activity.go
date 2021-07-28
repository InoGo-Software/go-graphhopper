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
	"time"
)

// Activity struct for Activity
type Activity struct {
	// type of activity
	Type *string `json:"type,omitempty"`
	// Id referring to the underlying service or shipment, i.e. the shipment or service this activity belongs to
	Id *string `json:"id,omitempty"`
	// Id that refers to address
	LocationId *string `json:"location_id,omitempty"`
	Address *ResponseAddress `json:"address,omitempty"`
	// Arrival time at this activity in seconds. If type is `start`, this is not available (since it makes no sense to have `arr_time` at start). However, `end_time` is available and actually means \\\"departure time\\\" at start location. It is important to note that `arr_time` does not necessarily mean \\\"start of underlying activity\\\", it solely means arrival time at activity location. If this activity has no time windows and if there are no further preparation times, `arr_time` is equal to activity start time.
	ArrTime *int64 `json:"arr_time,omitempty"`
	// End time of and thus departure time at this activity. If type is `end`, this is not available (since it makes no sense to have an `end_time` at end) `end_time` at each activity is equal to the departure time at the activity location.
	EndTime *int64 `json:"end_time,omitempty"`
	// End date time with offset like this 1970-01-01T01:00+01:00. If you do not use time-dependent optimization, this is `null`.
	EndDateTime *time.Time `json:"end_date_time,omitempty"`
	// Arrival date time with offset like this 1970-01-01T01:00+01:00. If you do not use time-dependent optimization, this is `null`.
	ArrDateTime *time.Time `json:"arr_date_time,omitempty"`
	// Waiting time at this activity in seconds. A waiting time can occur if the activity has at least one time window. If `arr_time` < `time_window.earliest` a waiting time of `time_window_earliest` - `arr_time` occurs.
	WaitingTime *int64 `json:"waiting_time,omitempty"`
	// preparation time at this activity in seconds
	PreparationTime *int64 `json:"preparation_time,omitempty"`
	// cumulated distance from start to this activity in m
	Distance *int64 `json:"distance,omitempty"`
	// cumulated driving time from start to this driver activity in seconds
	DrivingTime *int64 `json:"driving_time,omitempty"`
	// Array with size/capacity dimensions before this activity
	LoadBefore *[]int32 `json:"load_before,omitempty"`
	// Array with size/capacity dimensions after this activity
	LoadAfter *[]int32 `json:"load_after,omitempty"`
}

// NewActivity instantiates a new Activity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivity() *Activity {
	this := Activity{}
	return &this
}

// NewActivityWithDefaults instantiates a new Activity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivityWithDefaults() *Activity {
	this := Activity{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Activity) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Activity) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Activity) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Activity) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Activity) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Activity) SetId(v string) {
	o.Id = &v
}

// GetLocationId returns the LocationId field value if set, zero value otherwise.
func (o *Activity) GetLocationId() string {
	if o == nil || o.LocationId == nil {
		var ret string
		return ret
	}
	return *o.LocationId
}

// GetLocationIdOk returns a tuple with the LocationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetLocationIdOk() (*string, bool) {
	if o == nil || o.LocationId == nil {
		return nil, false
	}
	return o.LocationId, true
}

// HasLocationId returns a boolean if a field has been set.
func (o *Activity) HasLocationId() bool {
	if o != nil && o.LocationId != nil {
		return true
	}

	return false
}

// SetLocationId gets a reference to the given string and assigns it to the LocationId field.
func (o *Activity) SetLocationId(v string) {
	o.LocationId = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *Activity) GetAddress() ResponseAddress {
	if o == nil || o.Address == nil {
		var ret ResponseAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetAddressOk() (*ResponseAddress, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *Activity) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given ResponseAddress and assigns it to the Address field.
func (o *Activity) SetAddress(v ResponseAddress) {
	o.Address = &v
}

// GetArrTime returns the ArrTime field value if set, zero value otherwise.
func (o *Activity) GetArrTime() int64 {
	if o == nil || o.ArrTime == nil {
		var ret int64
		return ret
	}
	return *o.ArrTime
}

// GetArrTimeOk returns a tuple with the ArrTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetArrTimeOk() (*int64, bool) {
	if o == nil || o.ArrTime == nil {
		return nil, false
	}
	return o.ArrTime, true
}

// HasArrTime returns a boolean if a field has been set.
func (o *Activity) HasArrTime() bool {
	if o != nil && o.ArrTime != nil {
		return true
	}

	return false
}

// SetArrTime gets a reference to the given int64 and assigns it to the ArrTime field.
func (o *Activity) SetArrTime(v int64) {
	o.ArrTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *Activity) GetEndTime() int64 {
	if o == nil || o.EndTime == nil {
		var ret int64
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetEndTimeOk() (*int64, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *Activity) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given int64 and assigns it to the EndTime field.
func (o *Activity) SetEndTime(v int64) {
	o.EndTime = &v
}

// GetEndDateTime returns the EndDateTime field value if set, zero value otherwise.
func (o *Activity) GetEndDateTime() time.Time {
	if o == nil || o.EndDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDateTime
}

// GetEndDateTimeOk returns a tuple with the EndDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetEndDateTimeOk() (*time.Time, bool) {
	if o == nil || o.EndDateTime == nil {
		return nil, false
	}
	return o.EndDateTime, true
}

// HasEndDateTime returns a boolean if a field has been set.
func (o *Activity) HasEndDateTime() bool {
	if o != nil && o.EndDateTime != nil {
		return true
	}

	return false
}

// SetEndDateTime gets a reference to the given time.Time and assigns it to the EndDateTime field.
func (o *Activity) SetEndDateTime(v time.Time) {
	o.EndDateTime = &v
}

// GetArrDateTime returns the ArrDateTime field value if set, zero value otherwise.
func (o *Activity) GetArrDateTime() time.Time {
	if o == nil || o.ArrDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ArrDateTime
}

// GetArrDateTimeOk returns a tuple with the ArrDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetArrDateTimeOk() (*time.Time, bool) {
	if o == nil || o.ArrDateTime == nil {
		return nil, false
	}
	return o.ArrDateTime, true
}

// HasArrDateTime returns a boolean if a field has been set.
func (o *Activity) HasArrDateTime() bool {
	if o != nil && o.ArrDateTime != nil {
		return true
	}

	return false
}

// SetArrDateTime gets a reference to the given time.Time and assigns it to the ArrDateTime field.
func (o *Activity) SetArrDateTime(v time.Time) {
	o.ArrDateTime = &v
}

// GetWaitingTime returns the WaitingTime field value if set, zero value otherwise.
func (o *Activity) GetWaitingTime() int64 {
	if o == nil || o.WaitingTime == nil {
		var ret int64
		return ret
	}
	return *o.WaitingTime
}

// GetWaitingTimeOk returns a tuple with the WaitingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetWaitingTimeOk() (*int64, bool) {
	if o == nil || o.WaitingTime == nil {
		return nil, false
	}
	return o.WaitingTime, true
}

// HasWaitingTime returns a boolean if a field has been set.
func (o *Activity) HasWaitingTime() bool {
	if o != nil && o.WaitingTime != nil {
		return true
	}

	return false
}

// SetWaitingTime gets a reference to the given int64 and assigns it to the WaitingTime field.
func (o *Activity) SetWaitingTime(v int64) {
	o.WaitingTime = &v
}

// GetPreparationTime returns the PreparationTime field value if set, zero value otherwise.
func (o *Activity) GetPreparationTime() int64 {
	if o == nil || o.PreparationTime == nil {
		var ret int64
		return ret
	}
	return *o.PreparationTime
}

// GetPreparationTimeOk returns a tuple with the PreparationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetPreparationTimeOk() (*int64, bool) {
	if o == nil || o.PreparationTime == nil {
		return nil, false
	}
	return o.PreparationTime, true
}

// HasPreparationTime returns a boolean if a field has been set.
func (o *Activity) HasPreparationTime() bool {
	if o != nil && o.PreparationTime != nil {
		return true
	}

	return false
}

// SetPreparationTime gets a reference to the given int64 and assigns it to the PreparationTime field.
func (o *Activity) SetPreparationTime(v int64) {
	o.PreparationTime = &v
}

// GetDistance returns the Distance field value if set, zero value otherwise.
func (o *Activity) GetDistance() int64 {
	if o == nil || o.Distance == nil {
		var ret int64
		return ret
	}
	return *o.Distance
}

// GetDistanceOk returns a tuple with the Distance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetDistanceOk() (*int64, bool) {
	if o == nil || o.Distance == nil {
		return nil, false
	}
	return o.Distance, true
}

// HasDistance returns a boolean if a field has been set.
func (o *Activity) HasDistance() bool {
	if o != nil && o.Distance != nil {
		return true
	}

	return false
}

// SetDistance gets a reference to the given int64 and assigns it to the Distance field.
func (o *Activity) SetDistance(v int64) {
	o.Distance = &v
}

// GetDrivingTime returns the DrivingTime field value if set, zero value otherwise.
func (o *Activity) GetDrivingTime() int64 {
	if o == nil || o.DrivingTime == nil {
		var ret int64
		return ret
	}
	return *o.DrivingTime
}

// GetDrivingTimeOk returns a tuple with the DrivingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetDrivingTimeOk() (*int64, bool) {
	if o == nil || o.DrivingTime == nil {
		return nil, false
	}
	return o.DrivingTime, true
}

// HasDrivingTime returns a boolean if a field has been set.
func (o *Activity) HasDrivingTime() bool {
	if o != nil && o.DrivingTime != nil {
		return true
	}

	return false
}

// SetDrivingTime gets a reference to the given int64 and assigns it to the DrivingTime field.
func (o *Activity) SetDrivingTime(v int64) {
	o.DrivingTime = &v
}

// GetLoadBefore returns the LoadBefore field value if set, zero value otherwise.
func (o *Activity) GetLoadBefore() []int32 {
	if o == nil || o.LoadBefore == nil {
		var ret []int32
		return ret
	}
	return *o.LoadBefore
}

// GetLoadBeforeOk returns a tuple with the LoadBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetLoadBeforeOk() (*[]int32, bool) {
	if o == nil || o.LoadBefore == nil {
		return nil, false
	}
	return o.LoadBefore, true
}

// HasLoadBefore returns a boolean if a field has been set.
func (o *Activity) HasLoadBefore() bool {
	if o != nil && o.LoadBefore != nil {
		return true
	}

	return false
}

// SetLoadBefore gets a reference to the given []int32 and assigns it to the LoadBefore field.
func (o *Activity) SetLoadBefore(v []int32) {
	o.LoadBefore = &v
}

// GetLoadAfter returns the LoadAfter field value if set, zero value otherwise.
func (o *Activity) GetLoadAfter() []int32 {
	if o == nil || o.LoadAfter == nil {
		var ret []int32
		return ret
	}
	return *o.LoadAfter
}

// GetLoadAfterOk returns a tuple with the LoadAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetLoadAfterOk() (*[]int32, bool) {
	if o == nil || o.LoadAfter == nil {
		return nil, false
	}
	return o.LoadAfter, true
}

// HasLoadAfter returns a boolean if a field has been set.
func (o *Activity) HasLoadAfter() bool {
	if o != nil && o.LoadAfter != nil {
		return true
	}

	return false
}

// SetLoadAfter gets a reference to the given []int32 and assigns it to the LoadAfter field.
func (o *Activity) SetLoadAfter(v []int32) {
	o.LoadAfter = &v
}

func (o Activity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LocationId != nil {
		toSerialize["location_id"] = o.LocationId
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.ArrTime != nil {
		toSerialize["arr_time"] = o.ArrTime
	}
	if o.EndTime != nil {
		toSerialize["end_time"] = o.EndTime
	}
	if o.EndDateTime != nil {
		toSerialize["end_date_time"] = o.EndDateTime
	}
	if o.ArrDateTime != nil {
		toSerialize["arr_date_time"] = o.ArrDateTime
	}
	if o.WaitingTime != nil {
		toSerialize["waiting_time"] = o.WaitingTime
	}
	if o.PreparationTime != nil {
		toSerialize["preparation_time"] = o.PreparationTime
	}
	if o.Distance != nil {
		toSerialize["distance"] = o.Distance
	}
	if o.DrivingTime != nil {
		toSerialize["driving_time"] = o.DrivingTime
	}
	if o.LoadBefore != nil {
		toSerialize["load_before"] = o.LoadBefore
	}
	if o.LoadAfter != nil {
		toSerialize["load_after"] = o.LoadAfter
	}
	return json.Marshal(toSerialize)
}

type NullableActivity struct {
	value *Activity
	isSet bool
}

func (v NullableActivity) Get() *Activity {
	return v.value
}

func (v *NullableActivity) Set(val *Activity) {
	v.value = val
	v.isSet = true
}

func (v NullableActivity) IsSet() bool {
	return v.isSet
}

func (v *NullableActivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivity(val *Activity) *NullableActivity {
	return &NullableActivity{value: val, isSet: true}
}

func (v NullableActivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


