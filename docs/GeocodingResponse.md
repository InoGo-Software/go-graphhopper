# GeocodingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hits** | Pointer to [**[]GeocodingLocation**](GeocodingLocation.md) |  | [optional] 
**Took** | Pointer to **float32** | in ms | [optional] 

## Methods

### NewGeocodingResponse

`func NewGeocodingResponse() *GeocodingResponse`

NewGeocodingResponse instantiates a new GeocodingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeocodingResponseWithDefaults

`func NewGeocodingResponseWithDefaults() *GeocodingResponse`

NewGeocodingResponseWithDefaults instantiates a new GeocodingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHits

`func (o *GeocodingResponse) GetHits() []GeocodingLocation`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *GeocodingResponse) GetHitsOk() (*[]GeocodingLocation, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *GeocodingResponse) SetHits(v []GeocodingLocation)`

SetHits sets Hits field to given value.

### HasHits

`func (o *GeocodingResponse) HasHits() bool`

HasHits returns a boolean if a field has been set.

### GetTook

`func (o *GeocodingResponse) GetTook() float32`

GetTook returns the Took field if non-nil, zero value otherwise.

### GetTookOk

`func (o *GeocodingResponse) GetTookOk() (*float32, bool)`

GetTookOk returns a tuple with the Took field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTook

`func (o *GeocodingResponse) SetTook(v float32)`

SetTook sets Took field to given value.

### HasTook

`func (o *GeocodingResponse) HasTook() bool`

HasTook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


