# ClusterResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Copyrights** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to **string** | Indicates the current status of the job | [optional] 
**WaitingTimeInQueue** | Pointer to **float64** |  | [optional] 
**ProcessingTime** | Pointer to **float64** |  | [optional] 
**Clusters** | Pointer to [**[]Cluster**](Cluster.md) |  | [optional] 

## Methods

### NewClusterResponse

`func NewClusterResponse() *ClusterResponse`

NewClusterResponse instantiates a new ClusterResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterResponseWithDefaults

`func NewClusterResponseWithDefaults() *ClusterResponse`

NewClusterResponseWithDefaults instantiates a new ClusterResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCopyrights

`func (o *ClusterResponse) GetCopyrights() []string`

GetCopyrights returns the Copyrights field if non-nil, zero value otherwise.

### GetCopyrightsOk

`func (o *ClusterResponse) GetCopyrightsOk() (*[]string, bool)`

GetCopyrightsOk returns a tuple with the Copyrights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyrights

`func (o *ClusterResponse) SetCopyrights(v []string)`

SetCopyrights sets Copyrights field to given value.

### HasCopyrights

`func (o *ClusterResponse) HasCopyrights() bool`

HasCopyrights returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWaitingTimeInQueue

`func (o *ClusterResponse) GetWaitingTimeInQueue() float64`

GetWaitingTimeInQueue returns the WaitingTimeInQueue field if non-nil, zero value otherwise.

### GetWaitingTimeInQueueOk

`func (o *ClusterResponse) GetWaitingTimeInQueueOk() (*float64, bool)`

GetWaitingTimeInQueueOk returns a tuple with the WaitingTimeInQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitingTimeInQueue

`func (o *ClusterResponse) SetWaitingTimeInQueue(v float64)`

SetWaitingTimeInQueue sets WaitingTimeInQueue field to given value.

### HasWaitingTimeInQueue

`func (o *ClusterResponse) HasWaitingTimeInQueue() bool`

HasWaitingTimeInQueue returns a boolean if a field has been set.

### GetProcessingTime

`func (o *ClusterResponse) GetProcessingTime() float64`

GetProcessingTime returns the ProcessingTime field if non-nil, zero value otherwise.

### GetProcessingTimeOk

`func (o *ClusterResponse) GetProcessingTimeOk() (*float64, bool)`

GetProcessingTimeOk returns a tuple with the ProcessingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingTime

`func (o *ClusterResponse) SetProcessingTime(v float64)`

SetProcessingTime sets ProcessingTime field to given value.

### HasProcessingTime

`func (o *ClusterResponse) HasProcessingTime() bool`

HasProcessingTime returns a boolean if a field has been set.

### GetClusters

`func (o *ClusterResponse) GetClusters() []Cluster`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *ClusterResponse) GetClustersOk() (*[]Cluster, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *ClusterResponse) SetClusters(v []Cluster)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *ClusterResponse) HasClusters() bool`

HasClusters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


