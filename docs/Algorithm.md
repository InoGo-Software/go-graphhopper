# Algorithm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProblemType** | Pointer to **string** |  | [optional] 
**Objective** | Pointer to **string** |  | [optional] 

## Methods

### NewAlgorithm

`func NewAlgorithm() *Algorithm`

NewAlgorithm instantiates a new Algorithm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlgorithmWithDefaults

`func NewAlgorithmWithDefaults() *Algorithm`

NewAlgorithmWithDefaults instantiates a new Algorithm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProblemType

`func (o *Algorithm) GetProblemType() string`

GetProblemType returns the ProblemType field if non-nil, zero value otherwise.

### GetProblemTypeOk

`func (o *Algorithm) GetProblemTypeOk() (*string, bool)`

GetProblemTypeOk returns a tuple with the ProblemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblemType

`func (o *Algorithm) SetProblemType(v string)`

SetProblemType sets ProblemType field to given value.

### HasProblemType

`func (o *Algorithm) HasProblemType() bool`

HasProblemType returns a boolean if a field has been set.

### GetObjective

`func (o *Algorithm) GetObjective() string`

GetObjective returns the Objective field if non-nil, zero value otherwise.

### GetObjectiveOk

`func (o *Algorithm) GetObjectiveOk() (*string, bool)`

GetObjectiveOk returns a tuple with the Objective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjective

`func (o *Algorithm) SetObjective(v string)`

SetObjective sets Objective field to given value.

### HasObjective

`func (o *Algorithm) HasObjective() bool`

HasObjective returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


