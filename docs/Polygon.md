# Polygon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Coordinates** | Pointer to [**[][][]float32**]([][]float32.md) |  | [optional] 

## Methods

### NewPolygon

`func NewPolygon() *Polygon`

NewPolygon instantiates a new Polygon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolygonWithDefaults

`func NewPolygonWithDefaults() *Polygon`

NewPolygonWithDefaults instantiates a new Polygon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Polygon) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Polygon) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Polygon) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Polygon) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCoordinates

`func (o *Polygon) GetCoordinates() [][][]float32`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *Polygon) GetCoordinatesOk() (*[][][]float32, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *Polygon) SetCoordinates(v [][][]float32)`

SetCoordinates sets Coordinates field to given value.

### HasCoordinates

`func (o *Polygon) HasCoordinates() bool`

HasCoordinates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


