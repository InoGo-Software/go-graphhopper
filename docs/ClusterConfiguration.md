# ClusterConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseType** | Pointer to **string** | Specifies the response format. You can either choose &#x60;geojson&#x60; or &#x60;json&#x60;. | [optional] [default to "json"]
**Routing** | Pointer to [**ClusterConfigurationRouting**](ClusterConfigurationRouting.md) |  | [optional] 
**Clustering** | Pointer to [**ClusterConfigurationClustering**](ClusterConfigurationClustering.md) |  | [optional] 

## Methods

### NewClusterConfiguration

`func NewClusterConfiguration() *ClusterConfiguration`

NewClusterConfiguration instantiates a new ClusterConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterConfigurationWithDefaults

`func NewClusterConfigurationWithDefaults() *ClusterConfiguration`

NewClusterConfigurationWithDefaults instantiates a new ClusterConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseType

`func (o *ClusterConfiguration) GetResponseType() string`

GetResponseType returns the ResponseType field if non-nil, zero value otherwise.

### GetResponseTypeOk

`func (o *ClusterConfiguration) GetResponseTypeOk() (*string, bool)`

GetResponseTypeOk returns a tuple with the ResponseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseType

`func (o *ClusterConfiguration) SetResponseType(v string)`

SetResponseType sets ResponseType field to given value.

### HasResponseType

`func (o *ClusterConfiguration) HasResponseType() bool`

HasResponseType returns a boolean if a field has been set.

### GetRouting

`func (o *ClusterConfiguration) GetRouting() ClusterConfigurationRouting`

GetRouting returns the Routing field if non-nil, zero value otherwise.

### GetRoutingOk

`func (o *ClusterConfiguration) GetRoutingOk() (*ClusterConfigurationRouting, bool)`

GetRoutingOk returns a tuple with the Routing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouting

`func (o *ClusterConfiguration) SetRouting(v ClusterConfigurationRouting)`

SetRouting sets Routing field to given value.

### HasRouting

`func (o *ClusterConfiguration) HasRouting() bool`

HasRouting returns a boolean if a field has been set.

### GetClustering

`func (o *ClusterConfiguration) GetClustering() ClusterConfigurationClustering`

GetClustering returns the Clustering field if non-nil, zero value otherwise.

### GetClusteringOk

`func (o *ClusterConfiguration) GetClusteringOk() (*ClusterConfigurationClustering, bool)`

GetClusteringOk returns a tuple with the Clustering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClustering

`func (o *ClusterConfiguration) SetClustering(v ClusterConfigurationClustering)`

SetClustering sets Clustering field to given value.

### HasClustering

`func (o *ClusterConfiguration) HasClustering() bool`

HasClustering returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


