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

// Vehicle struct for Vehicle
type Vehicle struct {
	// Specifies the ID of the vehicle. Ids must be unique, i.e. if there are two vehicles with the same ID, an error is returned.
	VehicleId string `json:"vehicle_id"`
	// The type ID assigns a vehicle type to this vehicle. You can specify types in the array of vehicle types. If you omit the type ID, the default type is used. The default type is a `car` with a capacity of 0.
	TypeId *string `json:"type_id,omitempty"`
	StartAddress Address `json:"start_address"`
	EndAddress *Address `json:"end_address,omitempty"`
	Break *AnyOfTimeWindowBreakDriveTimeBreak `json:"break,omitempty"`
	// If it is false, the algorithm decides where to end the vehicle route. It ends in one of your customers' locations. The end is chosen such that it contributes to the overall objective function, e.g. min transport_time. If it is true, you can either specify a specific end location (which is then regarded as end depot) or you can leave it and the driver returns to its start location.
	ReturnToDepot *bool `json:"return_to_depot,omitempty"`
	// Earliest start of vehicle in seconds. It is recommended to use the unix timestamp.
	EarliestStart *int64 `json:"earliest_start,omitempty"`
	// Latest end of vehicle in seconds, i.e. the time the vehicle needs to be at its end location at latest.
	LatestEnd *int64 `json:"latest_end,omitempty"`
	// Array of skills, i.e. array of string (not case sensitive).
	Skills *[]string `json:"skills,omitempty"`
	// Specifies the maximum distance (in meters) a vehicle can go.
	MaxDistance *int64 `json:"max_distance,omitempty"`
	// Specifies the maximum drive time (in seconds) a vehicle/driver can go, i.e. the maximum time on the road (service and waiting times are not included here)
	MaxDrivingTime *int64 `json:"max_driving_time,omitempty"`
	// Specifies the maximum number of jobs a vehicle can load.
	MaxJobs *int32 `json:"max_jobs,omitempty"`
	// Specifies the minimum number of jobs a vehicle should load. This is a soft constraint, i.e. if it is not possible to fulfill “min_jobs”, we will still try to get as close as possible to this constraint.
	MinJobs *int32 `json:"min_jobs,omitempty"`
	// Specifies the maximum number of activities a vehicle can conduct.
	MaxActivities *int32 `json:"max_activities,omitempty"`
	// Indicates whether a vehicle should be moved even though it has not been assigned any jobs.
	MoveToEndAddress *bool `json:"move_to_end_address,omitempty"`
}

// NewVehicle instantiates a new Vehicle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVehicle(vehicleId string, startAddress Address) *Vehicle {
	this := Vehicle{}
	this.VehicleId = vehicleId
	var typeId string = "default-type"
	this.TypeId = &typeId
	this.StartAddress = startAddress
	var returnToDepot bool = true
	this.ReturnToDepot = &returnToDepot
	var earliestStart int64 = 0
	this.EarliestStart = &earliestStart
	return &this
}

// NewVehicleWithDefaults instantiates a new Vehicle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVehicleWithDefaults() *Vehicle {
	this := Vehicle{}
	var typeId string = "default-type"
	this.TypeId = &typeId
	var returnToDepot bool = true
	this.ReturnToDepot = &returnToDepot
	var earliestStart int64 = 0
	this.EarliestStart = &earliestStart
	return &this
}

// GetVehicleId returns the VehicleId field value
func (o *Vehicle) GetVehicleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VehicleId
}

// GetVehicleIdOk returns a tuple with the VehicleId field value
// and a boolean to check if the value has been set.
func (o *Vehicle) GetVehicleIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VehicleId, true
}

// SetVehicleId sets field value
func (o *Vehicle) SetVehicleId(v string) {
	o.VehicleId = v
}

// GetTypeId returns the TypeId field value if set, zero value otherwise.
func (o *Vehicle) GetTypeId() string {
	if o == nil || o.TypeId == nil {
		var ret string
		return ret
	}
	return *o.TypeId
}

// GetTypeIdOk returns a tuple with the TypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetTypeIdOk() (*string, bool) {
	if o == nil || o.TypeId == nil {
		return nil, false
	}
	return o.TypeId, true
}

// HasTypeId returns a boolean if a field has been set.
func (o *Vehicle) HasTypeId() bool {
	if o != nil && o.TypeId != nil {
		return true
	}

	return false
}

// SetTypeId gets a reference to the given string and assigns it to the TypeId field.
func (o *Vehicle) SetTypeId(v string) {
	o.TypeId = &v
}

// GetStartAddress returns the StartAddress field value
func (o *Vehicle) GetStartAddress() Address {
	if o == nil {
		var ret Address
		return ret
	}

	return o.StartAddress
}

// GetStartAddressOk returns a tuple with the StartAddress field value
// and a boolean to check if the value has been set.
func (o *Vehicle) GetStartAddressOk() (*Address, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StartAddress, true
}

// SetStartAddress sets field value
func (o *Vehicle) SetStartAddress(v Address) {
	o.StartAddress = v
}

// GetEndAddress returns the EndAddress field value if set, zero value otherwise.
func (o *Vehicle) GetEndAddress() Address {
	if o == nil || o.EndAddress == nil {
		var ret Address
		return ret
	}
	return *o.EndAddress
}

// GetEndAddressOk returns a tuple with the EndAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetEndAddressOk() (*Address, bool) {
	if o == nil || o.EndAddress == nil {
		return nil, false
	}
	return o.EndAddress, true
}

// HasEndAddress returns a boolean if a field has been set.
func (o *Vehicle) HasEndAddress() bool {
	if o != nil && o.EndAddress != nil {
		return true
	}

	return false
}

// SetEndAddress gets a reference to the given Address and assigns it to the EndAddress field.
func (o *Vehicle) SetEndAddress(v Address) {
	o.EndAddress = &v
}

// GetBreak returns the Break field value if set, zero value otherwise.
func (o *Vehicle) GetBreak() AnyOfTimeWindowBreakDriveTimeBreak {
	if o == nil || o.Break == nil {
		var ret AnyOfTimeWindowBreakDriveTimeBreak
		return ret
	}
	return *o.Break
}

// GetBreakOk returns a tuple with the Break field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetBreakOk() (*AnyOfTimeWindowBreakDriveTimeBreak, bool) {
	if o == nil || o.Break == nil {
		return nil, false
	}
	return o.Break, true
}

// HasBreak returns a boolean if a field has been set.
func (o *Vehicle) HasBreak() bool {
	if o != nil && o.Break != nil {
		return true
	}

	return false
}

// SetBreak gets a reference to the given AnyOfTimeWindowBreakDriveTimeBreak and assigns it to the Break field.
func (o *Vehicle) SetBreak(v AnyOfTimeWindowBreakDriveTimeBreak) {
	o.Break = &v
}

// GetReturnToDepot returns the ReturnToDepot field value if set, zero value otherwise.
func (o *Vehicle) GetReturnToDepot() bool {
	if o == nil || o.ReturnToDepot == nil {
		var ret bool
		return ret
	}
	return *o.ReturnToDepot
}

// GetReturnToDepotOk returns a tuple with the ReturnToDepot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetReturnToDepotOk() (*bool, bool) {
	if o == nil || o.ReturnToDepot == nil {
		return nil, false
	}
	return o.ReturnToDepot, true
}

// HasReturnToDepot returns a boolean if a field has been set.
func (o *Vehicle) HasReturnToDepot() bool {
	if o != nil && o.ReturnToDepot != nil {
		return true
	}

	return false
}

// SetReturnToDepot gets a reference to the given bool and assigns it to the ReturnToDepot field.
func (o *Vehicle) SetReturnToDepot(v bool) {
	o.ReturnToDepot = &v
}

// GetEarliestStart returns the EarliestStart field value if set, zero value otherwise.
func (o *Vehicle) GetEarliestStart() int64 {
	if o == nil || o.EarliestStart == nil {
		var ret int64
		return ret
	}
	return *o.EarliestStart
}

// GetEarliestStartOk returns a tuple with the EarliestStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetEarliestStartOk() (*int64, bool) {
	if o == nil || o.EarliestStart == nil {
		return nil, false
	}
	return o.EarliestStart, true
}

// HasEarliestStart returns a boolean if a field has been set.
func (o *Vehicle) HasEarliestStart() bool {
	if o != nil && o.EarliestStart != nil {
		return true
	}

	return false
}

// SetEarliestStart gets a reference to the given int64 and assigns it to the EarliestStart field.
func (o *Vehicle) SetEarliestStart(v int64) {
	o.EarliestStart = &v
}

// GetLatestEnd returns the LatestEnd field value if set, zero value otherwise.
func (o *Vehicle) GetLatestEnd() int64 {
	if o == nil || o.LatestEnd == nil {
		var ret int64
		return ret
	}
	return *o.LatestEnd
}

// GetLatestEndOk returns a tuple with the LatestEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetLatestEndOk() (*int64, bool) {
	if o == nil || o.LatestEnd == nil {
		return nil, false
	}
	return o.LatestEnd, true
}

// HasLatestEnd returns a boolean if a field has been set.
func (o *Vehicle) HasLatestEnd() bool {
	if o != nil && o.LatestEnd != nil {
		return true
	}

	return false
}

// SetLatestEnd gets a reference to the given int64 and assigns it to the LatestEnd field.
func (o *Vehicle) SetLatestEnd(v int64) {
	o.LatestEnd = &v
}

// GetSkills returns the Skills field value if set, zero value otherwise.
func (o *Vehicle) GetSkills() []string {
	if o == nil || o.Skills == nil {
		var ret []string
		return ret
	}
	return *o.Skills
}

// GetSkillsOk returns a tuple with the Skills field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetSkillsOk() (*[]string, bool) {
	if o == nil || o.Skills == nil {
		return nil, false
	}
	return o.Skills, true
}

// HasSkills returns a boolean if a field has been set.
func (o *Vehicle) HasSkills() bool {
	if o != nil && o.Skills != nil {
		return true
	}

	return false
}

// SetSkills gets a reference to the given []string and assigns it to the Skills field.
func (o *Vehicle) SetSkills(v []string) {
	o.Skills = &v
}

// GetMaxDistance returns the MaxDistance field value if set, zero value otherwise.
func (o *Vehicle) GetMaxDistance() int64 {
	if o == nil || o.MaxDistance == nil {
		var ret int64
		return ret
	}
	return *o.MaxDistance
}

// GetMaxDistanceOk returns a tuple with the MaxDistance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetMaxDistanceOk() (*int64, bool) {
	if o == nil || o.MaxDistance == nil {
		return nil, false
	}
	return o.MaxDistance, true
}

// HasMaxDistance returns a boolean if a field has been set.
func (o *Vehicle) HasMaxDistance() bool {
	if o != nil && o.MaxDistance != nil {
		return true
	}

	return false
}

// SetMaxDistance gets a reference to the given int64 and assigns it to the MaxDistance field.
func (o *Vehicle) SetMaxDistance(v int64) {
	o.MaxDistance = &v
}

// GetMaxDrivingTime returns the MaxDrivingTime field value if set, zero value otherwise.
func (o *Vehicle) GetMaxDrivingTime() int64 {
	if o == nil || o.MaxDrivingTime == nil {
		var ret int64
		return ret
	}
	return *o.MaxDrivingTime
}

// GetMaxDrivingTimeOk returns a tuple with the MaxDrivingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetMaxDrivingTimeOk() (*int64, bool) {
	if o == nil || o.MaxDrivingTime == nil {
		return nil, false
	}
	return o.MaxDrivingTime, true
}

// HasMaxDrivingTime returns a boolean if a field has been set.
func (o *Vehicle) HasMaxDrivingTime() bool {
	if o != nil && o.MaxDrivingTime != nil {
		return true
	}

	return false
}

// SetMaxDrivingTime gets a reference to the given int64 and assigns it to the MaxDrivingTime field.
func (o *Vehicle) SetMaxDrivingTime(v int64) {
	o.MaxDrivingTime = &v
}

// GetMaxJobs returns the MaxJobs field value if set, zero value otherwise.
func (o *Vehicle) GetMaxJobs() int32 {
	if o == nil || o.MaxJobs == nil {
		var ret int32
		return ret
	}
	return *o.MaxJobs
}

// GetMaxJobsOk returns a tuple with the MaxJobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetMaxJobsOk() (*int32, bool) {
	if o == nil || o.MaxJobs == nil {
		return nil, false
	}
	return o.MaxJobs, true
}

// HasMaxJobs returns a boolean if a field has been set.
func (o *Vehicle) HasMaxJobs() bool {
	if o != nil && o.MaxJobs != nil {
		return true
	}

	return false
}

// SetMaxJobs gets a reference to the given int32 and assigns it to the MaxJobs field.
func (o *Vehicle) SetMaxJobs(v int32) {
	o.MaxJobs = &v
}

// GetMinJobs returns the MinJobs field value if set, zero value otherwise.
func (o *Vehicle) GetMinJobs() int32 {
	if o == nil || o.MinJobs == nil {
		var ret int32
		return ret
	}
	return *o.MinJobs
}

// GetMinJobsOk returns a tuple with the MinJobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetMinJobsOk() (*int32, bool) {
	if o == nil || o.MinJobs == nil {
		return nil, false
	}
	return o.MinJobs, true
}

// HasMinJobs returns a boolean if a field has been set.
func (o *Vehicle) HasMinJobs() bool {
	if o != nil && o.MinJobs != nil {
		return true
	}

	return false
}

// SetMinJobs gets a reference to the given int32 and assigns it to the MinJobs field.
func (o *Vehicle) SetMinJobs(v int32) {
	o.MinJobs = &v
}

// GetMaxActivities returns the MaxActivities field value if set, zero value otherwise.
func (o *Vehicle) GetMaxActivities() int32 {
	if o == nil || o.MaxActivities == nil {
		var ret int32
		return ret
	}
	return *o.MaxActivities
}

// GetMaxActivitiesOk returns a tuple with the MaxActivities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetMaxActivitiesOk() (*int32, bool) {
	if o == nil || o.MaxActivities == nil {
		return nil, false
	}
	return o.MaxActivities, true
}

// HasMaxActivities returns a boolean if a field has been set.
func (o *Vehicle) HasMaxActivities() bool {
	if o != nil && o.MaxActivities != nil {
		return true
	}

	return false
}

// SetMaxActivities gets a reference to the given int32 and assigns it to the MaxActivities field.
func (o *Vehicle) SetMaxActivities(v int32) {
	o.MaxActivities = &v
}

// GetMoveToEndAddress returns the MoveToEndAddress field value if set, zero value otherwise.
func (o *Vehicle) GetMoveToEndAddress() bool {
	if o == nil || o.MoveToEndAddress == nil {
		var ret bool
		return ret
	}
	return *o.MoveToEndAddress
}

// GetMoveToEndAddressOk returns a tuple with the MoveToEndAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetMoveToEndAddressOk() (*bool, bool) {
	if o == nil || o.MoveToEndAddress == nil {
		return nil, false
	}
	return o.MoveToEndAddress, true
}

// HasMoveToEndAddress returns a boolean if a field has been set.
func (o *Vehicle) HasMoveToEndAddress() bool {
	if o != nil && o.MoveToEndAddress != nil {
		return true
	}

	return false
}

// SetMoveToEndAddress gets a reference to the given bool and assigns it to the MoveToEndAddress field.
func (o *Vehicle) SetMoveToEndAddress(v bool) {
	o.MoveToEndAddress = &v
}

func (o Vehicle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["vehicle_id"] = o.VehicleId
	}
	if o.TypeId != nil {
		toSerialize["type_id"] = o.TypeId
	}
	if true {
		toSerialize["start_address"] = o.StartAddress
	}
	if o.EndAddress != nil {
		toSerialize["end_address"] = o.EndAddress
	}
	if o.Break != nil {
		toSerialize["break"] = o.Break
	}
	if o.ReturnToDepot != nil {
		toSerialize["return_to_depot"] = o.ReturnToDepot
	}
	if o.EarliestStart != nil {
		toSerialize["earliest_start"] = o.EarliestStart
	}
	if o.LatestEnd != nil {
		toSerialize["latest_end"] = o.LatestEnd
	}
	if o.Skills != nil {
		toSerialize["skills"] = o.Skills
	}
	if o.MaxDistance != nil {
		toSerialize["max_distance"] = o.MaxDistance
	}
	if o.MaxDrivingTime != nil {
		toSerialize["max_driving_time"] = o.MaxDrivingTime
	}
	if o.MaxJobs != nil {
		toSerialize["max_jobs"] = o.MaxJobs
	}
	if o.MinJobs != nil {
		toSerialize["min_jobs"] = o.MinJobs
	}
	if o.MaxActivities != nil {
		toSerialize["max_activities"] = o.MaxActivities
	}
	if o.MoveToEndAddress != nil {
		toSerialize["move_to_end_address"] = o.MoveToEndAddress
	}
	return json.Marshal(toSerialize)
}

type NullableVehicle struct {
	value *Vehicle
	isSet bool
}

func (v NullableVehicle) Get() *Vehicle {
	return v.value
}

func (v *NullableVehicle) Set(val *Vehicle) {
	v.value = val
	v.isSet = true
}

func (v NullableVehicle) IsSet() bool {
	return v.isSet
}

func (v *NullableVehicle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVehicle(val *Vehicle) *NullableVehicle {
	return &NullableVehicle{value: val, isSet: true}
}

func (v NullableVehicle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVehicle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


