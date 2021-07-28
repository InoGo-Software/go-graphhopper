# CostMatrix

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | type of cost matrix, currently default or google are supported | [optional] 
**LocationIds** | Pointer to **[]string** |  | [optional] 
**Data** | Pointer to [**CostMatrixData**](CostMatrixData.md) |  | [optional] 
**Profile** | Pointer to **string** | vehicle profile or empty if catch all fallback | [optional] 

## Methods

### NewCostMatrix

`func NewCostMatrix() *CostMatrix`

NewCostMatrix instantiates a new CostMatrix object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostMatrixWithDefaults

`func NewCostMatrixWithDefaults() *CostMatrix`

NewCostMatrixWithDefaults instantiates a new CostMatrix object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CostMatrix) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CostMatrix) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CostMatrix) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CostMatrix) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLocationIds

`func (o *CostMatrix) GetLocationIds() []string`

GetLocationIds returns the LocationIds field if non-nil, zero value otherwise.

### GetLocationIdsOk

`func (o *CostMatrix) GetLocationIdsOk() (*[]string, bool)`

GetLocationIdsOk returns a tuple with the LocationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationIds

`func (o *CostMatrix) SetLocationIds(v []string)`

SetLocationIds sets LocationIds field to given value.

### HasLocationIds

`func (o *CostMatrix) HasLocationIds() bool`

HasLocationIds returns a boolean if a field has been set.

### GetData

`func (o *CostMatrix) GetData() CostMatrixData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CostMatrix) GetDataOk() (*CostMatrixData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CostMatrix) SetData(v CostMatrixData)`

SetData sets Data field to given value.

### HasData

`func (o *CostMatrix) HasData() bool`

HasData returns a boolean if a field has been set.

### GetProfile

`func (o *CostMatrix) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *CostMatrix) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *CostMatrix) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *CostMatrix) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


