# TimeWindowBreak

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Earliest** | **int64** | Specifies the earliest start time of the break in seconds. | 
**Latest** | **int64** | Specifies the latest start time of break in seconds. | 
**Duration** | **int64** | Specifies the duration of the break in seconds. | 

## Methods

### NewTimeWindowBreak

`func NewTimeWindowBreak(earliest int64, latest int64, duration int64, ) *TimeWindowBreak`

NewTimeWindowBreak instantiates a new TimeWindowBreak object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeWindowBreakWithDefaults

`func NewTimeWindowBreakWithDefaults() *TimeWindowBreak`

NewTimeWindowBreakWithDefaults instantiates a new TimeWindowBreak object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEarliest

`func (o *TimeWindowBreak) GetEarliest() int64`

GetEarliest returns the Earliest field if non-nil, zero value otherwise.

### GetEarliestOk

`func (o *TimeWindowBreak) GetEarliestOk() (*int64, bool)`

GetEarliestOk returns a tuple with the Earliest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliest

`func (o *TimeWindowBreak) SetEarliest(v int64)`

SetEarliest sets Earliest field to given value.


### GetLatest

`func (o *TimeWindowBreak) GetLatest() int64`

GetLatest returns the Latest field if non-nil, zero value otherwise.

### GetLatestOk

`func (o *TimeWindowBreak) GetLatestOk() (*int64, bool)`

GetLatestOk returns a tuple with the Latest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatest

`func (o *TimeWindowBreak) SetLatest(v int64)`

SetLatest sets Latest field to given value.


### GetDuration

`func (o *TimeWindowBreak) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *TimeWindowBreak) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *TimeWindowBreak) SetDuration(v int64)`

SetDuration sets Duration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


