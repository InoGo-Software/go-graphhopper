# Detail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Id of unassigned service/shipment | [optional] 
**Code** | Pointer to **int32** | Reason code  Code   |  Reason :------|:--------- 1 | cannot serve required skill 2 | cannot be visited within time window 3 | does not fit into any vehicle due to capacity 4 | cannot be assigned due to max distance constraint of vehicles 21 | could not be assigned due to relation constraint 22 | could not be assigned due to allowed vehicle constraint 23 | could not be assigned due to max-time-in-vehicle constraint 24 | driver does not need a break 25 | could not be assigned due to disallowed vehicle constraint 26 | could not be assigned due to max drive time constraint 27 | could not be assigned due to max job constraint 28 | could not be assigned due to max activity constraint 50 | underlying location cannot be accessed over road network by at least one vehicle  | [optional] 
**Reason** | Pointer to **string** | Human readable reason as listed above | [optional] 

## Methods

### NewDetail

`func NewDetail() *Detail`

NewDetail instantiates a new Detail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetailWithDefaults

`func NewDetailWithDefaults() *Detail`

NewDetailWithDefaults instantiates a new Detail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Detail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Detail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Detail) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Detail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *Detail) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Detail) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Detail) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *Detail) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetReason

`func (o *Detail) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Detail) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Detail) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *Detail) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


