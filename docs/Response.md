# Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Copyrights** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to **string** | Indicates the current status of the job | [optional] 
**WaitingTimeInQueue** | Pointer to **int64** | Waiting time in ms | [optional] 
**ProcessingTime** | Pointer to **int64** | Processing time in ms. If job is still waiting in queue, processing_time is 0 | [optional] 
**Solution** | Pointer to [**Solution**](Solution.md) |  | [optional] 

## Methods

### NewResponse

`func NewResponse() *Response`

NewResponse instantiates a new Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseWithDefaults

`func NewResponseWithDefaults() *Response`

NewResponseWithDefaults instantiates a new Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCopyrights

`func (o *Response) GetCopyrights() []string`

GetCopyrights returns the Copyrights field if non-nil, zero value otherwise.

### GetCopyrightsOk

`func (o *Response) GetCopyrightsOk() (*[]string, bool)`

GetCopyrightsOk returns a tuple with the Copyrights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyrights

`func (o *Response) SetCopyrights(v []string)`

SetCopyrights sets Copyrights field to given value.

### HasCopyrights

`func (o *Response) HasCopyrights() bool`

HasCopyrights returns a boolean if a field has been set.

### GetStatus

`func (o *Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWaitingTimeInQueue

`func (o *Response) GetWaitingTimeInQueue() int64`

GetWaitingTimeInQueue returns the WaitingTimeInQueue field if non-nil, zero value otherwise.

### GetWaitingTimeInQueueOk

`func (o *Response) GetWaitingTimeInQueueOk() (*int64, bool)`

GetWaitingTimeInQueueOk returns a tuple with the WaitingTimeInQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitingTimeInQueue

`func (o *Response) SetWaitingTimeInQueue(v int64)`

SetWaitingTimeInQueue sets WaitingTimeInQueue field to given value.

### HasWaitingTimeInQueue

`func (o *Response) HasWaitingTimeInQueue() bool`

HasWaitingTimeInQueue returns a boolean if a field has been set.

### GetProcessingTime

`func (o *Response) GetProcessingTime() int64`

GetProcessingTime returns the ProcessingTime field if non-nil, zero value otherwise.

### GetProcessingTimeOk

`func (o *Response) GetProcessingTimeOk() (*int64, bool)`

GetProcessingTimeOk returns a tuple with the ProcessingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingTime

`func (o *Response) SetProcessingTime(v int64)`

SetProcessingTime sets ProcessingTime field to given value.

### HasProcessingTime

`func (o *Response) HasProcessingTime() bool`

HasProcessingTime returns a boolean if a field has been set.

### GetSolution

`func (o *Response) GetSolution() Solution`

GetSolution returns the Solution field if non-nil, zero value otherwise.

### GetSolutionOk

`func (o *Response) GetSolutionOk() (*Solution, bool)`

GetSolutionOk returns a tuple with the Solution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolution

`func (o *Response) SetSolution(v Solution)`

SetSolution sets Solution field to given value.

### HasSolution

`func (o *Response) HasSolution() bool`

HasSolution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


