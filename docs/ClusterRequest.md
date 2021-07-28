# ClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to [**ClusterConfiguration**](ClusterConfiguration.md) |  | [optional] 
**Customers** | Pointer to [**[]ClusterCustomer**](ClusterCustomer.md) |  | [optional] 

## Methods

### NewClusterRequest

`func NewClusterRequest() *ClusterRequest`

NewClusterRequest instantiates a new ClusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterRequestWithDefaults

`func NewClusterRequestWithDefaults() *ClusterRequest`

NewClusterRequestWithDefaults instantiates a new ClusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *ClusterRequest) GetConfiguration() ClusterConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ClusterRequest) GetConfigurationOk() (*ClusterConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ClusterRequest) SetConfiguration(v ClusterConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *ClusterRequest) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetCustomers

`func (o *ClusterRequest) GetCustomers() []ClusterCustomer`

GetCustomers returns the Customers field if non-nil, zero value otherwise.

### GetCustomersOk

`func (o *ClusterRequest) GetCustomersOk() (*[]ClusterCustomer, bool)`

GetCustomersOk returns a tuple with the Customers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomers

`func (o *ClusterRequest) SetCustomers(v []ClusterCustomer)`

SetCustomers sets Customers field to given value.

### HasCustomers

`func (o *ClusterRequest) HasCustomers() bool`

HasCustomers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


