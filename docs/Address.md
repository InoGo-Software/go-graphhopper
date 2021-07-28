# Address

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationId** | **string** | Specifies the id of the location. | 
**Name** | Pointer to **string** | Name of location. | [optional] 
**Lon** | **float64** | Longitude of location. | 
**Lat** | **float64** | Latitude of location. | 
**StreetHint** | Pointer to **string** | Optional parameter. Specifies a hint for each address to better snap the coordinates (lon,lat) to road network. E.g. if there is an address or house with two or more neighboring streets you can control for which street the closest location is looked up. | [optional] 
**Curbside** | Pointer to **string** | Optional parameter. Specifies on which side a point should be relative to the driver when she leaves/arrives at a start/target/via point. Only supported for motor vehicles and OpenStreetMap. | [optional] [default to "any"]

## Methods

### NewAddress

`func NewAddress(locationId string, lon float64, lat float64, ) *Address`

NewAddress instantiates a new Address object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressWithDefaults

`func NewAddressWithDefaults() *Address`

NewAddressWithDefaults instantiates a new Address object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationId

`func (o *Address) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *Address) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *Address) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.


### GetName

`func (o *Address) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Address) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Address) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Address) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLon

`func (o *Address) GetLon() float64`

GetLon returns the Lon field if non-nil, zero value otherwise.

### GetLonOk

`func (o *Address) GetLonOk() (*float64, bool)`

GetLonOk returns a tuple with the Lon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLon

`func (o *Address) SetLon(v float64)`

SetLon sets Lon field to given value.


### GetLat

`func (o *Address) GetLat() float64`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *Address) GetLatOk() (*float64, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *Address) SetLat(v float64)`

SetLat sets Lat field to given value.


### GetStreetHint

`func (o *Address) GetStreetHint() string`

GetStreetHint returns the StreetHint field if non-nil, zero value otherwise.

### GetStreetHintOk

`func (o *Address) GetStreetHintOk() (*string, bool)`

GetStreetHintOk returns a tuple with the StreetHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetHint

`func (o *Address) SetStreetHint(v string)`

SetStreetHint sets StreetHint field to given value.

### HasStreetHint

`func (o *Address) HasStreetHint() bool`

HasStreetHint returns a boolean if a field has been set.

### GetCurbside

`func (o *Address) GetCurbside() string`

GetCurbside returns the Curbside field if non-nil, zero value otherwise.

### GetCurbsideOk

`func (o *Address) GetCurbsideOk() (*string, bool)`

GetCurbsideOk returns a tuple with the Curbside field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurbside

`func (o *Address) SetCurbside(v string)`

SetCurbside sets Curbside field to given value.

### HasCurbside

`func (o *Address) HasCurbside() bool`

HasCurbside returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


