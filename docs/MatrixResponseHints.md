# MatrixResponseHints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Short description of this hint | [optional] 
**Details** | Pointer to **string** | Details of this hint | [optional] 
**InvalidFromPoints** | Pointer to **[]float32** | Optional. An array of from_point indices of points that could not be found. Will only be added if &#x60;fail_fast&#x3D;false&#x60; and some &#x60;from_point&#x60;s were not found.&#x60; | [optional] 
**InvalidToPoints** | Pointer to **[]float32** | Optional. An array of to_point indices of points that could not be found. Will only be added if &#x60;fail_fast&#x3D;false&#x60; and some &#x60;to_point&#x60;s were not found.&#x60; | [optional] 
**PointPairs** | Pointer to **[][]float32** | Optional. An array of two-element arrays representing the from/to_point indices of points for which no connection could be found. Will only be added if &#x60;fail_fast&#x3D;false&#x60; and some connections were not found. | [optional] 

## Methods

### NewMatrixResponseHints

`func NewMatrixResponseHints() *MatrixResponseHints`

NewMatrixResponseHints instantiates a new MatrixResponseHints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatrixResponseHintsWithDefaults

`func NewMatrixResponseHintsWithDefaults() *MatrixResponseHints`

NewMatrixResponseHintsWithDefaults instantiates a new MatrixResponseHints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *MatrixResponseHints) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MatrixResponseHints) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MatrixResponseHints) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MatrixResponseHints) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDetails

`func (o *MatrixResponseHints) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *MatrixResponseHints) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *MatrixResponseHints) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *MatrixResponseHints) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetInvalidFromPoints

`func (o *MatrixResponseHints) GetInvalidFromPoints() []float32`

GetInvalidFromPoints returns the InvalidFromPoints field if non-nil, zero value otherwise.

### GetInvalidFromPointsOk

`func (o *MatrixResponseHints) GetInvalidFromPointsOk() (*[]float32, bool)`

GetInvalidFromPointsOk returns a tuple with the InvalidFromPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidFromPoints

`func (o *MatrixResponseHints) SetInvalidFromPoints(v []float32)`

SetInvalidFromPoints sets InvalidFromPoints field to given value.

### HasInvalidFromPoints

`func (o *MatrixResponseHints) HasInvalidFromPoints() bool`

HasInvalidFromPoints returns a boolean if a field has been set.

### GetInvalidToPoints

`func (o *MatrixResponseHints) GetInvalidToPoints() []float32`

GetInvalidToPoints returns the InvalidToPoints field if non-nil, zero value otherwise.

### GetInvalidToPointsOk

`func (o *MatrixResponseHints) GetInvalidToPointsOk() (*[]float32, bool)`

GetInvalidToPointsOk returns a tuple with the InvalidToPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidToPoints

`func (o *MatrixResponseHints) SetInvalidToPoints(v []float32)`

SetInvalidToPoints sets InvalidToPoints field to given value.

### HasInvalidToPoints

`func (o *MatrixResponseHints) HasInvalidToPoints() bool`

HasInvalidToPoints returns a boolean if a field has been set.

### GetPointPairs

`func (o *MatrixResponseHints) GetPointPairs() [][]float32`

GetPointPairs returns the PointPairs field if non-nil, zero value otherwise.

### GetPointPairsOk

`func (o *MatrixResponseHints) GetPointPairsOk() (*[][]float32, bool)`

GetPointPairsOk returns a tuple with the PointPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointPairs

`func (o *MatrixResponseHints) SetPointPairs(v [][]float32)`

SetPointPairs sets PointPairs field to given value.

### HasPointPairs

`func (o *MatrixResponseHints) HasPointPairs() bool`

HasPointPairs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


