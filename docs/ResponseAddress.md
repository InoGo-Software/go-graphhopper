# ResponseAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationId** | Pointer to **string** | Specifies the id of the location. | [optional] 
**Name** | Pointer to **string** | Name of location. | [optional] 
**Lon** | Pointer to **float64** | Longitude of location. | [optional] 
**Lat** | Pointer to **float64** | Latitude of location. | [optional] 
**StreetHint** | Pointer to **string** | Optional parameter. Specifies a hint for each address to better snap the coordinates (lon,lat) to road network. E.g. if there is an address or house with two or more neighboring streets you can control for which street the closest location is looked up. | [optional] 
**SnappedWaypoint** | Pointer to [**SnappedWaypoint**](SnappedWaypoint.md) |  | [optional] 

## Methods

### NewResponseAddress

`func NewResponseAddress() *ResponseAddress`

NewResponseAddress instantiates a new ResponseAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseAddressWithDefaults

`func NewResponseAddressWithDefaults() *ResponseAddress`

NewResponseAddressWithDefaults instantiates a new ResponseAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationId

`func (o *ResponseAddress) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *ResponseAddress) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *ResponseAddress) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *ResponseAddress) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### GetName

`func (o *ResponseAddress) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseAddress) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseAddress) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResponseAddress) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLon

`func (o *ResponseAddress) GetLon() float64`

GetLon returns the Lon field if non-nil, zero value otherwise.

### GetLonOk

`func (o *ResponseAddress) GetLonOk() (*float64, bool)`

GetLonOk returns a tuple with the Lon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLon

`func (o *ResponseAddress) SetLon(v float64)`

SetLon sets Lon field to given value.

### HasLon

`func (o *ResponseAddress) HasLon() bool`

HasLon returns a boolean if a field has been set.

### GetLat

`func (o *ResponseAddress) GetLat() float64`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *ResponseAddress) GetLatOk() (*float64, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *ResponseAddress) SetLat(v float64)`

SetLat sets Lat field to given value.

### HasLat

`func (o *ResponseAddress) HasLat() bool`

HasLat returns a boolean if a field has been set.

### GetStreetHint

`func (o *ResponseAddress) GetStreetHint() string`

GetStreetHint returns the StreetHint field if non-nil, zero value otherwise.

### GetStreetHintOk

`func (o *ResponseAddress) GetStreetHintOk() (*string, bool)`

GetStreetHintOk returns a tuple with the StreetHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetHint

`func (o *ResponseAddress) SetStreetHint(v string)`

SetStreetHint sets StreetHint field to given value.

### HasStreetHint

`func (o *ResponseAddress) HasStreetHint() bool`

HasStreetHint returns a boolean if a field has been set.

### GetSnappedWaypoint

`func (o *ResponseAddress) GetSnappedWaypoint() SnappedWaypoint`

GetSnappedWaypoint returns the SnappedWaypoint field if non-nil, zero value otherwise.

### GetSnappedWaypointOk

`func (o *ResponseAddress) GetSnappedWaypointOk() (*SnappedWaypoint, bool)`

GetSnappedWaypointOk returns a tuple with the SnappedWaypoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnappedWaypoint

`func (o *ResponseAddress) SetSnappedWaypoint(v SnappedWaypoint)`

SetSnappedWaypoint sets SnappedWaypoint field to given value.

### HasSnappedWaypoint

`func (o *ResponseAddress) HasSnappedWaypoint() bool`

HasSnappedWaypoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


