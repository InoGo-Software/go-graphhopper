# IsochroneResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Polygons** | Pointer to [**[]IsochroneResponsePolygon**](IsochroneResponsePolygon.md) | The list of polygons in GeoJson format. It can be used e.g. in the Leaflet framework:  &#x60;&#x60;&#x60; L.geoJson(json.polygons).addTo(map) &#x60;&#x60;&#x60;  The number of polygon is identical to the specified buckets in the query. Every polygon contains the bucket number in the properties section of the GeoJson.  | [optional] 
**Copyrights** | Pointer to **[]string** |  | [optional] 

## Methods

### NewIsochroneResponse

`func NewIsochroneResponse() *IsochroneResponse`

NewIsochroneResponse instantiates a new IsochroneResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIsochroneResponseWithDefaults

`func NewIsochroneResponseWithDefaults() *IsochroneResponse`

NewIsochroneResponseWithDefaults instantiates a new IsochroneResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolygons

`func (o *IsochroneResponse) GetPolygons() []IsochroneResponsePolygon`

GetPolygons returns the Polygons field if non-nil, zero value otherwise.

### GetPolygonsOk

`func (o *IsochroneResponse) GetPolygonsOk() (*[]IsochroneResponsePolygon, bool)`

GetPolygonsOk returns a tuple with the Polygons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolygons

`func (o *IsochroneResponse) SetPolygons(v []IsochroneResponsePolygon)`

SetPolygons sets Polygons field to given value.

### HasPolygons

`func (o *IsochroneResponse) HasPolygons() bool`

HasPolygons returns a boolean if a field has been set.

### GetCopyrights

`func (o *IsochroneResponse) GetCopyrights() []string`

GetCopyrights returns the Copyrights field if non-nil, zero value otherwise.

### GetCopyrightsOk

`func (o *IsochroneResponse) GetCopyrightsOk() (*[]string, bool)`

GetCopyrightsOk returns a tuple with the Copyrights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyrights

`func (o *IsochroneResponse) SetCopyrights(v []string)`

SetCopyrights sets Copyrights field to given value.

### HasCopyrights

`func (o *IsochroneResponse) HasCopyrights() bool`

HasCopyrights returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


