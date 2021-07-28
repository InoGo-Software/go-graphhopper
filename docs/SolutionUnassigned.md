# SolutionUnassigned

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Services** | Pointer to **[]string** | An array of ids of unassigned services | [optional] 
**Shipments** | Pointer to **[]string** | An array of ids of unassigned shipments | [optional] 
**Breaks** | Pointer to **[]string** | An array of ids of unassigned breaks | [optional] 
**Details** | Pointer to [**[]Detail**](Detail.md) | An array of details, i.e. reason for unassigned services or shipments | [optional] 

## Methods

### NewSolutionUnassigned

`func NewSolutionUnassigned() *SolutionUnassigned`

NewSolutionUnassigned instantiates a new SolutionUnassigned object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSolutionUnassignedWithDefaults

`func NewSolutionUnassignedWithDefaults() *SolutionUnassigned`

NewSolutionUnassignedWithDefaults instantiates a new SolutionUnassigned object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServices

`func (o *SolutionUnassigned) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *SolutionUnassigned) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *SolutionUnassigned) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *SolutionUnassigned) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetShipments

`func (o *SolutionUnassigned) GetShipments() []string`

GetShipments returns the Shipments field if non-nil, zero value otherwise.

### GetShipmentsOk

`func (o *SolutionUnassigned) GetShipmentsOk() (*[]string, bool)`

GetShipmentsOk returns a tuple with the Shipments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipments

`func (o *SolutionUnassigned) SetShipments(v []string)`

SetShipments sets Shipments field to given value.

### HasShipments

`func (o *SolutionUnassigned) HasShipments() bool`

HasShipments returns a boolean if a field has been set.

### GetBreaks

`func (o *SolutionUnassigned) GetBreaks() []string`

GetBreaks returns the Breaks field if non-nil, zero value otherwise.

### GetBreaksOk

`func (o *SolutionUnassigned) GetBreaksOk() (*[]string, bool)`

GetBreaksOk returns a tuple with the Breaks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreaks

`func (o *SolutionUnassigned) SetBreaks(v []string)`

SetBreaks sets Breaks field to given value.

### HasBreaks

`func (o *SolutionUnassigned) HasBreaks() bool`

HasBreaks returns a boolean if a field has been set.

### GetDetails

`func (o *SolutionUnassigned) GetDetails() []Detail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *SolutionUnassigned) GetDetailsOk() (*[]Detail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *SolutionUnassigned) SetDetails(v []Detail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *SolutionUnassigned) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


