# GeocodingLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Point** | Pointer to [**GeocodingPoint**](GeocodingPoint.md) |  | [optional] 
**OsmId** | Pointer to **string** | The OSM ID of the entity | [optional] 
**OsmType** | Pointer to **string** | N &#x3D; node, R &#x3D; relation, W &#x3D; way | [optional] 
**OsmKey** | Pointer to **string** | The OSM key of the entity | [optional] 
**Name** | Pointer to **string** | The name of the entity. Can be a boundary, POI, address, etc | [optional] 
**Country** | Pointer to **string** | The country of the address | [optional] 
**City** | Pointer to **string** | The city of the address | [optional] 
**State** | Pointer to **string** | The state of the address | [optional] 
**Street** | Pointer to **string** | The street of the address | [optional] 
**Housenumber** | Pointer to **string** | The housenumber of the address | [optional] 
**Postcode** | Pointer to **string** | The postcode of the address | [optional] 

## Methods

### NewGeocodingLocation

`func NewGeocodingLocation() *GeocodingLocation`

NewGeocodingLocation instantiates a new GeocodingLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeocodingLocationWithDefaults

`func NewGeocodingLocationWithDefaults() *GeocodingLocation`

NewGeocodingLocationWithDefaults instantiates a new GeocodingLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoint

`func (o *GeocodingLocation) GetPoint() GeocodingPoint`

GetPoint returns the Point field if non-nil, zero value otherwise.

### GetPointOk

`func (o *GeocodingLocation) GetPointOk() (*GeocodingPoint, bool)`

GetPointOk returns a tuple with the Point field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoint

`func (o *GeocodingLocation) SetPoint(v GeocodingPoint)`

SetPoint sets Point field to given value.

### HasPoint

`func (o *GeocodingLocation) HasPoint() bool`

HasPoint returns a boolean if a field has been set.

### GetOsmId

`func (o *GeocodingLocation) GetOsmId() string`

GetOsmId returns the OsmId field if non-nil, zero value otherwise.

### GetOsmIdOk

`func (o *GeocodingLocation) GetOsmIdOk() (*string, bool)`

GetOsmIdOk returns a tuple with the OsmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsmId

`func (o *GeocodingLocation) SetOsmId(v string)`

SetOsmId sets OsmId field to given value.

### HasOsmId

`func (o *GeocodingLocation) HasOsmId() bool`

HasOsmId returns a boolean if a field has been set.

### GetOsmType

`func (o *GeocodingLocation) GetOsmType() string`

GetOsmType returns the OsmType field if non-nil, zero value otherwise.

### GetOsmTypeOk

`func (o *GeocodingLocation) GetOsmTypeOk() (*string, bool)`

GetOsmTypeOk returns a tuple with the OsmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsmType

`func (o *GeocodingLocation) SetOsmType(v string)`

SetOsmType sets OsmType field to given value.

### HasOsmType

`func (o *GeocodingLocation) HasOsmType() bool`

HasOsmType returns a boolean if a field has been set.

### GetOsmKey

`func (o *GeocodingLocation) GetOsmKey() string`

GetOsmKey returns the OsmKey field if non-nil, zero value otherwise.

### GetOsmKeyOk

`func (o *GeocodingLocation) GetOsmKeyOk() (*string, bool)`

GetOsmKeyOk returns a tuple with the OsmKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsmKey

`func (o *GeocodingLocation) SetOsmKey(v string)`

SetOsmKey sets OsmKey field to given value.

### HasOsmKey

`func (o *GeocodingLocation) HasOsmKey() bool`

HasOsmKey returns a boolean if a field has been set.

### GetName

`func (o *GeocodingLocation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GeocodingLocation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GeocodingLocation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GeocodingLocation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCountry

`func (o *GeocodingLocation) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *GeocodingLocation) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *GeocodingLocation) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *GeocodingLocation) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCity

`func (o *GeocodingLocation) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *GeocodingLocation) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *GeocodingLocation) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *GeocodingLocation) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *GeocodingLocation) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GeocodingLocation) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GeocodingLocation) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GeocodingLocation) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStreet

`func (o *GeocodingLocation) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *GeocodingLocation) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *GeocodingLocation) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *GeocodingLocation) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### GetHousenumber

`func (o *GeocodingLocation) GetHousenumber() string`

GetHousenumber returns the Housenumber field if non-nil, zero value otherwise.

### GetHousenumberOk

`func (o *GeocodingLocation) GetHousenumberOk() (*string, bool)`

GetHousenumberOk returns a tuple with the Housenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHousenumber

`func (o *GeocodingLocation) SetHousenumber(v string)`

SetHousenumber sets Housenumber field to given value.

### HasHousenumber

`func (o *GeocodingLocation) HasHousenumber() bool`

HasHousenumber returns a boolean if a field has been set.

### GetPostcode

`func (o *GeocodingLocation) GetPostcode() string`

GetPostcode returns the Postcode field if non-nil, zero value otherwise.

### GetPostcodeOk

`func (o *GeocodingLocation) GetPostcodeOk() (*string, bool)`

GetPostcodeOk returns a tuple with the Postcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostcode

`func (o *GeocodingLocation) SetPostcode(v string)`

SetPostcode sets Postcode field to given value.

### HasPostcode

`func (o *GeocodingLocation) HasPostcode() bool`

HasPostcode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


