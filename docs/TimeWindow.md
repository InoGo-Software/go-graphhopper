# TimeWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Earliest** | Pointer to **int64** | Specifies the opening time of the time window in seconds, i.e. the earliest time the service can start. | [optional] [default to 0]
**Latest** | Pointer to **int64** | Specifies the closing time of the time window in seconds, i.e. the latest time the service can start. | [optional] 

## Methods

### NewTimeWindow

`func NewTimeWindow() *TimeWindow`

NewTimeWindow instantiates a new TimeWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeWindowWithDefaults

`func NewTimeWindowWithDefaults() *TimeWindow`

NewTimeWindowWithDefaults instantiates a new TimeWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEarliest

`func (o *TimeWindow) GetEarliest() int64`

GetEarliest returns the Earliest field if non-nil, zero value otherwise.

### GetEarliestOk

`func (o *TimeWindow) GetEarliestOk() (*int64, bool)`

GetEarliestOk returns a tuple with the Earliest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliest

`func (o *TimeWindow) SetEarliest(v int64)`

SetEarliest sets Earliest field to given value.

### HasEarliest

`func (o *TimeWindow) HasEarliest() bool`

HasEarliest returns a boolean if a field has been set.

### GetLatest

`func (o *TimeWindow) GetLatest() int64`

GetLatest returns the Latest field if non-nil, zero value otherwise.

### GetLatestOk

`func (o *TimeWindow) GetLatestOk() (*int64, bool)`

GetLatestOk returns a tuple with the Latest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatest

`func (o *TimeWindow) SetLatest(v int64)`

SetLatest sets Latest field to given value.

### HasLatest

`func (o *TimeWindow) HasLatest() bool`

HasLatest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


