# Stop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**Address**](Address.md) |  | [optional] 
**Duration** | Pointer to **int64** | Specifies the duration of the pickup or delivery in seconds, e.g. how long it takes unload items at the customer site. | [optional] [default to 0]
**PreparationTime** | Pointer to **int64** | Specifies the preparation time in seconds. It can be used to model parking lot search time since if you have 3 identical locations in a row, it only falls due once. | [optional] [default to 0]
**TimeWindows** | Pointer to [**[]TimeWindow**](TimeWindow.md) | Specifies an array of time window objects (see time window object below). For example, if an item needs to be delivered between 7am and 10am then specify the array as follows: [ { \&quot;earliest\&quot;: 25200, \&quot;latest\&quot; : 32400 } ] (starting the day from 0 in seconds). | [optional] 
**Group** | Pointer to **string** | Group this stop belongs to. See the group relation and [this post](https://discuss.graphhopper.com/t/4040) on how to utilize this. | [optional] 

## Methods

### NewStop

`func NewStop() *Stop`

NewStop instantiates a new Stop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStopWithDefaults

`func NewStopWithDefaults() *Stop`

NewStopWithDefaults instantiates a new Stop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *Stop) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Stop) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Stop) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Stop) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetDuration

`func (o *Stop) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Stop) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Stop) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *Stop) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetPreparationTime

`func (o *Stop) GetPreparationTime() int64`

GetPreparationTime returns the PreparationTime field if non-nil, zero value otherwise.

### GetPreparationTimeOk

`func (o *Stop) GetPreparationTimeOk() (*int64, bool)`

GetPreparationTimeOk returns a tuple with the PreparationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparationTime

`func (o *Stop) SetPreparationTime(v int64)`

SetPreparationTime sets PreparationTime field to given value.

### HasPreparationTime

`func (o *Stop) HasPreparationTime() bool`

HasPreparationTime returns a boolean if a field has been set.

### GetTimeWindows

`func (o *Stop) GetTimeWindows() []TimeWindow`

GetTimeWindows returns the TimeWindows field if non-nil, zero value otherwise.

### GetTimeWindowsOk

`func (o *Stop) GetTimeWindowsOk() (*[]TimeWindow, bool)`

GetTimeWindowsOk returns a tuple with the TimeWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindows

`func (o *Stop) SetTimeWindows(v []TimeWindow)`

SetTimeWindows sets TimeWindows field to given value.

### HasTimeWindows

`func (o *Stop) HasTimeWindows() bool`

HasTimeWindows returns a boolean if a field has been set.

### GetGroup

`func (o *Stop) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *Stop) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *Stop) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *Stop) HasGroup() bool`

HasGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


