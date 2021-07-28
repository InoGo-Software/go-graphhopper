# MatrixResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Distances** | Pointer to **[][]float32** | The distance matrix for the specified points in the same order as the time matrix. The distances are in meters. If &#x60;fail_fast&#x3D;false&#x60; the matrix will contain &#x60;null&#x60; for connections that could not be found. | [optional] 
**Times** | Pointer to **[][]float32** | The time matrix for the specified points in the order [[from1-&gt;to1, from1-&gt;to2, ...], [from2-&gt;to1, from2-&gt;to2, ...], ...]. The times are in seconds. If &#x60;fail_fast&#x3D;false&#x60; the matrix will contain &#x60;null&#x60; for connections that could not be found. | [optional] 
**Weights** | Pointer to **[][]float64** | The weight matrix for the specified points in the same order as the time matrix. The weights for different vehicles can have a different unit but the weights array is perfectly suited as input for Vehicle Routing Problems as it is currently faster to calculate. If &#x60;fail_fast&#x3D;false&#x60; the matrix will contain &#x60;null&#x60; for connections that could not be found. | [optional] 
**Info** | Pointer to [**ResponseInfo**](ResponseInfo.md) |  | [optional] 
**Hints** | Pointer to [**[]MatrixResponseHints**](MatrixResponseHints.md) | Optional. Additional response data. | [optional] 

## Methods

### NewMatrixResponse

`func NewMatrixResponse() *MatrixResponse`

NewMatrixResponse instantiates a new MatrixResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatrixResponseWithDefaults

`func NewMatrixResponseWithDefaults() *MatrixResponse`

NewMatrixResponseWithDefaults instantiates a new MatrixResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistances

`func (o *MatrixResponse) GetDistances() [][]float32`

GetDistances returns the Distances field if non-nil, zero value otherwise.

### GetDistancesOk

`func (o *MatrixResponse) GetDistancesOk() (*[][]float32, bool)`

GetDistancesOk returns a tuple with the Distances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistances

`func (o *MatrixResponse) SetDistances(v [][]float32)`

SetDistances sets Distances field to given value.

### HasDistances

`func (o *MatrixResponse) HasDistances() bool`

HasDistances returns a boolean if a field has been set.

### GetTimes

`func (o *MatrixResponse) GetTimes() [][]float32`

GetTimes returns the Times field if non-nil, zero value otherwise.

### GetTimesOk

`func (o *MatrixResponse) GetTimesOk() (*[][]float32, bool)`

GetTimesOk returns a tuple with the Times field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimes

`func (o *MatrixResponse) SetTimes(v [][]float32)`

SetTimes sets Times field to given value.

### HasTimes

`func (o *MatrixResponse) HasTimes() bool`

HasTimes returns a boolean if a field has been set.

### GetWeights

`func (o *MatrixResponse) GetWeights() [][]float64`

GetWeights returns the Weights field if non-nil, zero value otherwise.

### GetWeightsOk

`func (o *MatrixResponse) GetWeightsOk() (*[][]float64, bool)`

GetWeightsOk returns a tuple with the Weights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeights

`func (o *MatrixResponse) SetWeights(v [][]float64)`

SetWeights sets Weights field to given value.

### HasWeights

`func (o *MatrixResponse) HasWeights() bool`

HasWeights returns a boolean if a field has been set.

### GetInfo

`func (o *MatrixResponse) GetInfo() ResponseInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *MatrixResponse) GetInfoOk() (*ResponseInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *MatrixResponse) SetInfo(v ResponseInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *MatrixResponse) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetHints

`func (o *MatrixResponse) GetHints() []MatrixResponseHints`

GetHints returns the Hints field if non-nil, zero value otherwise.

### GetHintsOk

`func (o *MatrixResponse) GetHintsOk() (*[]MatrixResponseHints, bool)`

GetHintsOk returns a tuple with the Hints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHints

`func (o *MatrixResponse) SetHints(v []MatrixResponseHints)`

SetHints sets Hints field to given value.

### HasHints

`func (o *MatrixResponse) HasHints() bool`

HasHints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


