# ResponseInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Copyrights** | Pointer to **[]string** | Attribution according to our documentation is necessary if no white-label option included. | [optional] 
**Took** | Pointer to **float64** |  | [optional] 

## Methods

### NewResponseInfo

`func NewResponseInfo() *ResponseInfo`

NewResponseInfo instantiates a new ResponseInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseInfoWithDefaults

`func NewResponseInfoWithDefaults() *ResponseInfo`

NewResponseInfoWithDefaults instantiates a new ResponseInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCopyrights

`func (o *ResponseInfo) GetCopyrights() []string`

GetCopyrights returns the Copyrights field if non-nil, zero value otherwise.

### GetCopyrightsOk

`func (o *ResponseInfo) GetCopyrightsOk() (*[]string, bool)`

GetCopyrightsOk returns a tuple with the Copyrights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyrights

`func (o *ResponseInfo) SetCopyrights(v []string)`

SetCopyrights sets Copyrights field to given value.

### HasCopyrights

`func (o *ResponseInfo) HasCopyrights() bool`

HasCopyrights returns a boolean if a field has been set.

### GetTook

`func (o *ResponseInfo) GetTook() float64`

GetTook returns the Took field if non-nil, zero value otherwise.

### GetTookOk

`func (o *ResponseInfo) GetTookOk() (*float64, bool)`

GetTookOk returns a tuple with the Took field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTook

`func (o *ResponseInfo) SetTook(v float64)`

SetTook sets Took field to given value.

### HasTook

`func (o *ResponseInfo) HasTook() bool`

HasTook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


