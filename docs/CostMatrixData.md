# CostMatrixData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Times** | Pointer to **[][]int64** |  | [optional] 
**Distances** | Pointer to **[][]float64** |  | [optional] 
**Info** | Pointer to [**CostMatrixDataInfo**](CostMatrixDataInfo.md) |  | [optional] 

## Methods

### NewCostMatrixData

`func NewCostMatrixData() *CostMatrixData`

NewCostMatrixData instantiates a new CostMatrixData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostMatrixDataWithDefaults

`func NewCostMatrixDataWithDefaults() *CostMatrixData`

NewCostMatrixDataWithDefaults instantiates a new CostMatrixData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimes

`func (o *CostMatrixData) GetTimes() [][]int64`

GetTimes returns the Times field if non-nil, zero value otherwise.

### GetTimesOk

`func (o *CostMatrixData) GetTimesOk() (*[][]int64, bool)`

GetTimesOk returns a tuple with the Times field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimes

`func (o *CostMatrixData) SetTimes(v [][]int64)`

SetTimes sets Times field to given value.

### HasTimes

`func (o *CostMatrixData) HasTimes() bool`

HasTimes returns a boolean if a field has been set.

### GetDistances

`func (o *CostMatrixData) GetDistances() [][]float64`

GetDistances returns the Distances field if non-nil, zero value otherwise.

### GetDistancesOk

`func (o *CostMatrixData) GetDistancesOk() (*[][]float64, bool)`

GetDistancesOk returns a tuple with the Distances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistances

`func (o *CostMatrixData) SetDistances(v [][]float64)`

SetDistances sets Distances field to given value.

### HasDistances

`func (o *CostMatrixData) HasDistances() bool`

HasDistances returns a boolean if a field has been set.

### GetInfo

`func (o *CostMatrixData) GetInfo() CostMatrixDataInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CostMatrixData) GetInfoOk() (*CostMatrixDataInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CostMatrixData) SetInfo(v CostMatrixDataInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *CostMatrixData) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


