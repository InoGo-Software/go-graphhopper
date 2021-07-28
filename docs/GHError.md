# GHError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Hints** | Pointer to [**[]GHErrorHints**](GHErrorHints.md) | Optional error information. | [optional] 

## Methods

### NewGHError

`func NewGHError() *GHError`

NewGHError instantiates a new GHError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGHErrorWithDefaults

`func NewGHErrorWithDefaults() *GHError`

NewGHErrorWithDefaults instantiates a new GHError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *GHError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GHError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GHError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GHError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetHints

`func (o *GHError) GetHints() []GHErrorHints`

GetHints returns the Hints field if non-nil, zero value otherwise.

### GetHintsOk

`func (o *GHError) GetHintsOk() (*[]GHErrorHints, bool)`

GetHintsOk returns a tuple with the Hints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHints

`func (o *GHError) SetHints(v []GHErrorHints)`

SetHints sets Hints field to given value.

### HasHints

`func (o *GHError) HasHints() bool`

HasHints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


