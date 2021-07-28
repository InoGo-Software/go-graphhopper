# BadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Short error message | [optional] 
**Hints** | Pointer to [**[]ErrorMessage**](ErrorMessage.md) | Optional error information. | [optional] 
**Status** | Pointer to **string** | status | [optional] [default to "finished"]

## Methods

### NewBadRequest

`func NewBadRequest() *BadRequest`

NewBadRequest instantiates a new BadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBadRequestWithDefaults

`func NewBadRequestWithDefaults() *BadRequest`

NewBadRequestWithDefaults instantiates a new BadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *BadRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BadRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BadRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BadRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetHints

`func (o *BadRequest) GetHints() []ErrorMessage`

GetHints returns the Hints field if non-nil, zero value otherwise.

### GetHintsOk

`func (o *BadRequest) GetHintsOk() (*[]ErrorMessage, bool)`

GetHintsOk returns a tuple with the Hints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHints

`func (o *BadRequest) SetHints(v []ErrorMessage)`

SetHints sets Hints field to given value.

### HasHints

`func (o *BadRequest) HasHints() bool`

HasHints returns a boolean if a field has been set.

### GetStatus

`func (o *BadRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BadRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BadRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BadRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


