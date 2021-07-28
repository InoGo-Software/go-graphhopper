# DriveTimeBreak

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **int64** | Specifies the duration of the break in seconds. | 
**MaxDrivingTime** | **int64** | Specifies the max driving time (in a row) without break in seconds. | 
**InitialDrivingTime** | Pointer to **int64** | Specifies the initial (current) driving time of a driver to allow dynamic adaptations in seconds. | [optional] 
**PossibleSplit** | Pointer to **[]int64** | Array specifying how a break duration (in seconds) can be split into several smaller breaks | [optional] 

## Methods

### NewDriveTimeBreak

`func NewDriveTimeBreak(duration int64, maxDrivingTime int64, ) *DriveTimeBreak`

NewDriveTimeBreak instantiates a new DriveTimeBreak object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriveTimeBreakWithDefaults

`func NewDriveTimeBreakWithDefaults() *DriveTimeBreak`

NewDriveTimeBreakWithDefaults instantiates a new DriveTimeBreak object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *DriveTimeBreak) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *DriveTimeBreak) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *DriveTimeBreak) SetDuration(v int64)`

SetDuration sets Duration field to given value.


### GetMaxDrivingTime

`func (o *DriveTimeBreak) GetMaxDrivingTime() int64`

GetMaxDrivingTime returns the MaxDrivingTime field if non-nil, zero value otherwise.

### GetMaxDrivingTimeOk

`func (o *DriveTimeBreak) GetMaxDrivingTimeOk() (*int64, bool)`

GetMaxDrivingTimeOk returns a tuple with the MaxDrivingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDrivingTime

`func (o *DriveTimeBreak) SetMaxDrivingTime(v int64)`

SetMaxDrivingTime sets MaxDrivingTime field to given value.


### GetInitialDrivingTime

`func (o *DriveTimeBreak) GetInitialDrivingTime() int64`

GetInitialDrivingTime returns the InitialDrivingTime field if non-nil, zero value otherwise.

### GetInitialDrivingTimeOk

`func (o *DriveTimeBreak) GetInitialDrivingTimeOk() (*int64, bool)`

GetInitialDrivingTimeOk returns a tuple with the InitialDrivingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialDrivingTime

`func (o *DriveTimeBreak) SetInitialDrivingTime(v int64)`

SetInitialDrivingTime sets InitialDrivingTime field to given value.

### HasInitialDrivingTime

`func (o *DriveTimeBreak) HasInitialDrivingTime() bool`

HasInitialDrivingTime returns a boolean if a field has been set.

### GetPossibleSplit

`func (o *DriveTimeBreak) GetPossibleSplit() []int64`

GetPossibleSplit returns the PossibleSplit field if non-nil, zero value otherwise.

### GetPossibleSplitOk

`func (o *DriveTimeBreak) GetPossibleSplitOk() (*[]int64, bool)`

GetPossibleSplitOk returns a tuple with the PossibleSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPossibleSplit

`func (o *DriveTimeBreak) SetPossibleSplit(v []int64)`

SetPossibleSplit sets PossibleSplit field to given value.

### HasPossibleSplit

`func (o *DriveTimeBreak) HasPossibleSplit() bool`

HasPossibleSplit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


