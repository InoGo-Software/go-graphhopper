# RouteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paths** | Pointer to [**[]RouteResponsePath**](RouteResponsePath.md) |  | [optional] 
**Info** | Pointer to [**ResponseInfo**](ResponseInfo.md) |  | [optional] 

## Methods

### NewRouteResponse

`func NewRouteResponse() *RouteResponse`

NewRouteResponse instantiates a new RouteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteResponseWithDefaults

`func NewRouteResponseWithDefaults() *RouteResponse`

NewRouteResponseWithDefaults instantiates a new RouteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaths

`func (o *RouteResponse) GetPaths() []RouteResponsePath`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *RouteResponse) GetPathsOk() (*[]RouteResponsePath, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *RouteResponse) SetPaths(v []RouteResponsePath)`

SetPaths sets Paths field to given value.

### HasPaths

`func (o *RouteResponse) HasPaths() bool`

HasPaths returns a boolean if a field has been set.

### GetInfo

`func (o *RouteResponse) GetInfo() ResponseInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *RouteResponse) GetInfoOk() (*ResponseInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *RouteResponse) SetInfo(v ResponseInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *RouteResponse) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


