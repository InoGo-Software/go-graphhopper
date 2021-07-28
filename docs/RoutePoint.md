# RoutePoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Coordinates** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewRoutePoint

`func NewRoutePoint() *RoutePoint`

NewRoutePoint instantiates a new RoutePoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutePointWithDefaults

`func NewRoutePointWithDefaults() *RoutePoint`

NewRoutePointWithDefaults instantiates a new RoutePoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RoutePoint) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoutePoint) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoutePoint) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RoutePoint) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCoordinates

`func (o *RoutePoint) GetCoordinates() []map[string]interface{}`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *RoutePoint) GetCoordinatesOk() (*[]map[string]interface{}, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *RoutePoint) SetCoordinates(v []map[string]interface{})`

SetCoordinates sets Coordinates field to given value.

### HasCoordinates

`func (o *RoutePoint) HasCoordinates() bool`

HasCoordinates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


