# SymmetricalMatrixRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Points** | Pointer to **[][]float64** | Specify multiple points for which the weight-, route-, time- or distance-matrix should be calculated as follows: &#x60;[longitude,latitude]&#x60;. In this case the origins are identical to the destinations. Thus, if there are N points, NxN entries are calculated. The order of the point parameter is important. Specify at least three points. Cannot be used together with &#x60;from_point&#x60; or &#x60;to_point.&#x60;. | [optional] 
**PointHints** | Pointer to **[]string** | Optional parameter. Specifies a hint for each point in the &#x60;points&#x60; array to prefer a certain street for the closest location lookup. E.g. if there is an address or house with two or more neighboring streets you can control for which street the closest location is looked up. | [optional] 
**SnapPreventions** | Pointer to **[]string** | Optional parameter to avoid snapping to a certain road class or road environment. Current supported values &#x60;motorway&#x60;, &#x60;trunk&#x60;, &#x60;ferry&#x60;, &#x60;tunnel&#x60;, &#x60;bridge&#x60; and &#x60;ford&#x60; | [optional] 
**Curbsides** | Pointer to **[]string** | Optional parameter. It specifies on which side a point should be relative to the driver when she leaves/arrives at a start/target/via point. You need to specify this parameter for either none or all points. Only supported for motor vehicles and OpenStreetMap. | [optional] 
**OutArrays** | Pointer to **[]string** | Specifies which matrices should be included in the response. Specify one or more of the following options &#x60;weights&#x60;, &#x60;times&#x60;, &#x60;distances&#x60;. The units of the entries of &#x60;distances&#x60; are meters, of &#x60;times&#x60; are seconds and of &#x60;weights&#x60; is arbitrary and it can differ for different vehicles or versions of this API. | [optional] 
**Vehicle** | Pointer to [**VehicleProfileId**](VehicleProfileId.md) |  | [optional] 
**FailFast** | Pointer to **bool** | Specifies whether or not the matrix calculation should return with an error as soon as possible in case some points cannot be found or some points are not connected. If set to &#x60;false&#x60; the time/weight/distance matrix will be calculated for all valid points and contain the &#x60;null&#x60; value for all entries that could not be calculated. The &#x60;hint&#x60; field of the response will also contain additional information about what went wrong (see its documentation). | [optional] [default to true]
**TurnCosts** | Pointer to **bool** | Specifies if turn restrictions should be considered. Enabling this option increases the matrix computation time. Only supported for motor vehicles and OpenStreetMap. | [optional] [default to false]

## Methods

### NewSymmetricalMatrixRequest

`func NewSymmetricalMatrixRequest() *SymmetricalMatrixRequest`

NewSymmetricalMatrixRequest instantiates a new SymmetricalMatrixRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSymmetricalMatrixRequestWithDefaults

`func NewSymmetricalMatrixRequestWithDefaults() *SymmetricalMatrixRequest`

NewSymmetricalMatrixRequestWithDefaults instantiates a new SymmetricalMatrixRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoints

`func (o *SymmetricalMatrixRequest) GetPoints() [][]float64`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *SymmetricalMatrixRequest) GetPointsOk() (*[][]float64, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *SymmetricalMatrixRequest) SetPoints(v [][]float64)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *SymmetricalMatrixRequest) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### GetPointHints

`func (o *SymmetricalMatrixRequest) GetPointHints() []string`

GetPointHints returns the PointHints field if non-nil, zero value otherwise.

### GetPointHintsOk

`func (o *SymmetricalMatrixRequest) GetPointHintsOk() (*[]string, bool)`

GetPointHintsOk returns a tuple with the PointHints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointHints

`func (o *SymmetricalMatrixRequest) SetPointHints(v []string)`

SetPointHints sets PointHints field to given value.

### HasPointHints

`func (o *SymmetricalMatrixRequest) HasPointHints() bool`

HasPointHints returns a boolean if a field has been set.

### GetSnapPreventions

`func (o *SymmetricalMatrixRequest) GetSnapPreventions() []string`

GetSnapPreventions returns the SnapPreventions field if non-nil, zero value otherwise.

### GetSnapPreventionsOk

`func (o *SymmetricalMatrixRequest) GetSnapPreventionsOk() (*[]string, bool)`

GetSnapPreventionsOk returns a tuple with the SnapPreventions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapPreventions

`func (o *SymmetricalMatrixRequest) SetSnapPreventions(v []string)`

SetSnapPreventions sets SnapPreventions field to given value.

### HasSnapPreventions

`func (o *SymmetricalMatrixRequest) HasSnapPreventions() bool`

HasSnapPreventions returns a boolean if a field has been set.

### GetCurbsides

`func (o *SymmetricalMatrixRequest) GetCurbsides() []string`

GetCurbsides returns the Curbsides field if non-nil, zero value otherwise.

### GetCurbsidesOk

`func (o *SymmetricalMatrixRequest) GetCurbsidesOk() (*[]string, bool)`

GetCurbsidesOk returns a tuple with the Curbsides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurbsides

`func (o *SymmetricalMatrixRequest) SetCurbsides(v []string)`

SetCurbsides sets Curbsides field to given value.

### HasCurbsides

`func (o *SymmetricalMatrixRequest) HasCurbsides() bool`

HasCurbsides returns a boolean if a field has been set.

### GetOutArrays

`func (o *SymmetricalMatrixRequest) GetOutArrays() []string`

GetOutArrays returns the OutArrays field if non-nil, zero value otherwise.

### GetOutArraysOk

`func (o *SymmetricalMatrixRequest) GetOutArraysOk() (*[]string, bool)`

GetOutArraysOk returns a tuple with the OutArrays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutArrays

`func (o *SymmetricalMatrixRequest) SetOutArrays(v []string)`

SetOutArrays sets OutArrays field to given value.

### HasOutArrays

`func (o *SymmetricalMatrixRequest) HasOutArrays() bool`

HasOutArrays returns a boolean if a field has been set.

### GetVehicle

`func (o *SymmetricalMatrixRequest) GetVehicle() VehicleProfileId`

GetVehicle returns the Vehicle field if non-nil, zero value otherwise.

### GetVehicleOk

`func (o *SymmetricalMatrixRequest) GetVehicleOk() (*VehicleProfileId, bool)`

GetVehicleOk returns a tuple with the Vehicle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicle

`func (o *SymmetricalMatrixRequest) SetVehicle(v VehicleProfileId)`

SetVehicle sets Vehicle field to given value.

### HasVehicle

`func (o *SymmetricalMatrixRequest) HasVehicle() bool`

HasVehicle returns a boolean if a field has been set.

### GetFailFast

`func (o *SymmetricalMatrixRequest) GetFailFast() bool`

GetFailFast returns the FailFast field if non-nil, zero value otherwise.

### GetFailFastOk

`func (o *SymmetricalMatrixRequest) GetFailFastOk() (*bool, bool)`

GetFailFastOk returns a tuple with the FailFast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailFast

`func (o *SymmetricalMatrixRequest) SetFailFast(v bool)`

SetFailFast sets FailFast field to given value.

### HasFailFast

`func (o *SymmetricalMatrixRequest) HasFailFast() bool`

HasFailFast returns a boolean if a field has been set.

### GetTurnCosts

`func (o *SymmetricalMatrixRequest) GetTurnCosts() bool`

GetTurnCosts returns the TurnCosts field if non-nil, zero value otherwise.

### GetTurnCostsOk

`func (o *SymmetricalMatrixRequest) GetTurnCostsOk() (*bool, bool)`

GetTurnCostsOk returns a tuple with the TurnCosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurnCosts

`func (o *SymmetricalMatrixRequest) SetTurnCosts(v bool)`

SetTurnCosts sets TurnCosts field to given value.

### HasTurnCosts

`func (o *SymmetricalMatrixRequest) HasTurnCosts() bool`

HasTurnCosts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


