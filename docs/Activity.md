# Activity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | type of activity | [optional] 
**Id** | Pointer to **string** | Id referring to the underlying service or shipment, i.e. the shipment or service this activity belongs to | [optional] 
**LocationId** | Pointer to **string** | Id that refers to address | [optional] 
**Address** | Pointer to [**ResponseAddress**](ResponseAddress.md) |  | [optional] 
**ArrTime** | Pointer to **int64** | Arrival time at this activity in seconds. If type is &#x60;start&#x60;, this is not available (since it makes no sense to have &#x60;arr_time&#x60; at start). However, &#x60;end_time&#x60; is available and actually means \\\&quot;departure time\\\&quot; at start location. It is important to note that &#x60;arr_time&#x60; does not necessarily mean \\\&quot;start of underlying activity\\\&quot;, it solely means arrival time at activity location. If this activity has no time windows and if there are no further preparation times, &#x60;arr_time&#x60; is equal to activity start time. | [optional] 
**EndTime** | Pointer to **int64** | End time of and thus departure time at this activity. If type is &#x60;end&#x60;, this is not available (since it makes no sense to have an &#x60;end_time&#x60; at end) &#x60;end_time&#x60; at each activity is equal to the departure time at the activity location. | [optional] 
**EndDateTime** | Pointer to **time.Time** | End date time with offset like this 1970-01-01T01:00+01:00. If you do not use time-dependent optimization, this is &#x60;null&#x60;. | [optional] 
**ArrDateTime** | Pointer to **time.Time** | Arrival date time with offset like this 1970-01-01T01:00+01:00. If you do not use time-dependent optimization, this is &#x60;null&#x60;. | [optional] 
**WaitingTime** | Pointer to **int64** | Waiting time at this activity in seconds. A waiting time can occur if the activity has at least one time window. If &#x60;arr_time&#x60; &lt; &#x60;time_window.earliest&#x60; a waiting time of &#x60;time_window_earliest&#x60; - &#x60;arr_time&#x60; occurs. | [optional] 
**PreparationTime** | Pointer to **int64** | preparation time at this activity in seconds | [optional] 
**Distance** | Pointer to **int64** | cumulated distance from start to this activity in m | [optional] 
**DrivingTime** | Pointer to **int64** | cumulated driving time from start to this driver activity in seconds | [optional] 
**LoadBefore** | Pointer to **[]int32** | Array with size/capacity dimensions before this activity | [optional] 
**LoadAfter** | Pointer to **[]int32** | Array with size/capacity dimensions after this activity | [optional] 

## Methods

### NewActivity

`func NewActivity() *Activity`

NewActivity instantiates a new Activity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityWithDefaults

`func NewActivityWithDefaults() *Activity`

NewActivityWithDefaults instantiates a new Activity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Activity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Activity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Activity) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Activity) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *Activity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Activity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Activity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Activity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocationId

`func (o *Activity) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *Activity) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *Activity) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *Activity) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### GetAddress

`func (o *Activity) GetAddress() ResponseAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Activity) GetAddressOk() (*ResponseAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Activity) SetAddress(v ResponseAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Activity) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetArrTime

`func (o *Activity) GetArrTime() int64`

GetArrTime returns the ArrTime field if non-nil, zero value otherwise.

### GetArrTimeOk

`func (o *Activity) GetArrTimeOk() (*int64, bool)`

GetArrTimeOk returns a tuple with the ArrTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrTime

`func (o *Activity) SetArrTime(v int64)`

SetArrTime sets ArrTime field to given value.

### HasArrTime

`func (o *Activity) HasArrTime() bool`

HasArrTime returns a boolean if a field has been set.

### GetEndTime

`func (o *Activity) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *Activity) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *Activity) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *Activity) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetEndDateTime

`func (o *Activity) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *Activity) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *Activity) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *Activity) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### GetArrDateTime

`func (o *Activity) GetArrDateTime() time.Time`

GetArrDateTime returns the ArrDateTime field if non-nil, zero value otherwise.

### GetArrDateTimeOk

`func (o *Activity) GetArrDateTimeOk() (*time.Time, bool)`

GetArrDateTimeOk returns a tuple with the ArrDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrDateTime

`func (o *Activity) SetArrDateTime(v time.Time)`

SetArrDateTime sets ArrDateTime field to given value.

### HasArrDateTime

`func (o *Activity) HasArrDateTime() bool`

HasArrDateTime returns a boolean if a field has been set.

### GetWaitingTime

`func (o *Activity) GetWaitingTime() int64`

GetWaitingTime returns the WaitingTime field if non-nil, zero value otherwise.

### GetWaitingTimeOk

`func (o *Activity) GetWaitingTimeOk() (*int64, bool)`

GetWaitingTimeOk returns a tuple with the WaitingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitingTime

`func (o *Activity) SetWaitingTime(v int64)`

SetWaitingTime sets WaitingTime field to given value.

### HasWaitingTime

`func (o *Activity) HasWaitingTime() bool`

HasWaitingTime returns a boolean if a field has been set.

### GetPreparationTime

`func (o *Activity) GetPreparationTime() int64`

GetPreparationTime returns the PreparationTime field if non-nil, zero value otherwise.

### GetPreparationTimeOk

`func (o *Activity) GetPreparationTimeOk() (*int64, bool)`

GetPreparationTimeOk returns a tuple with the PreparationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparationTime

`func (o *Activity) SetPreparationTime(v int64)`

SetPreparationTime sets PreparationTime field to given value.

### HasPreparationTime

`func (o *Activity) HasPreparationTime() bool`

HasPreparationTime returns a boolean if a field has been set.

### GetDistance

`func (o *Activity) GetDistance() int64`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *Activity) GetDistanceOk() (*int64, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *Activity) SetDistance(v int64)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *Activity) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### GetDrivingTime

`func (o *Activity) GetDrivingTime() int64`

GetDrivingTime returns the DrivingTime field if non-nil, zero value otherwise.

### GetDrivingTimeOk

`func (o *Activity) GetDrivingTimeOk() (*int64, bool)`

GetDrivingTimeOk returns a tuple with the DrivingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivingTime

`func (o *Activity) SetDrivingTime(v int64)`

SetDrivingTime sets DrivingTime field to given value.

### HasDrivingTime

`func (o *Activity) HasDrivingTime() bool`

HasDrivingTime returns a boolean if a field has been set.

### GetLoadBefore

`func (o *Activity) GetLoadBefore() []int32`

GetLoadBefore returns the LoadBefore field if non-nil, zero value otherwise.

### GetLoadBeforeOk

`func (o *Activity) GetLoadBeforeOk() (*[]int32, bool)`

GetLoadBeforeOk returns a tuple with the LoadBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBefore

`func (o *Activity) SetLoadBefore(v []int32)`

SetLoadBefore sets LoadBefore field to given value.

### HasLoadBefore

`func (o *Activity) HasLoadBefore() bool`

HasLoadBefore returns a boolean if a field has been set.

### GetLoadAfter

`func (o *Activity) GetLoadAfter() []int32`

GetLoadAfter returns the LoadAfter field if non-nil, zero value otherwise.

### GetLoadAfterOk

`func (o *Activity) GetLoadAfterOk() (*[]int32, bool)`

GetLoadAfterOk returns a tuple with the LoadAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadAfter

`func (o *Activity) SetLoadAfter(v []int32)`

SetLoadAfter sets LoadAfter field to given value.

### HasLoadAfter

`func (o *Activity) HasLoadAfter() bool`

HasLoadAfter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


