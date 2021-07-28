# ClusterCustomerAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lon** | Pointer to **float64** | Longitude | [optional] 
**Lat** | Pointer to **float64** | Latitude | [optional] 
**StreetHint** | Pointer to **string** | Optional parameter. Specifies a hint for each address to better snap the coordinates (lon,lat) to road network. E.g. if there is an address or house with two or more neighboring streets you can control for which street the closest location is looked up. | [optional] 

## Methods

### NewClusterCustomerAddress

`func NewClusterCustomerAddress() *ClusterCustomerAddress`

NewClusterCustomerAddress instantiates a new ClusterCustomerAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterCustomerAddressWithDefaults

`func NewClusterCustomerAddressWithDefaults() *ClusterCustomerAddress`

NewClusterCustomerAddressWithDefaults instantiates a new ClusterCustomerAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLon

`func (o *ClusterCustomerAddress) GetLon() float64`

GetLon returns the Lon field if non-nil, zero value otherwise.

### GetLonOk

`func (o *ClusterCustomerAddress) GetLonOk() (*float64, bool)`

GetLonOk returns a tuple with the Lon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLon

`func (o *ClusterCustomerAddress) SetLon(v float64)`

SetLon sets Lon field to given value.

### HasLon

`func (o *ClusterCustomerAddress) HasLon() bool`

HasLon returns a boolean if a field has been set.

### GetLat

`func (o *ClusterCustomerAddress) GetLat() float64`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *ClusterCustomerAddress) GetLatOk() (*float64, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *ClusterCustomerAddress) SetLat(v float64)`

SetLat sets Lat field to given value.

### HasLat

`func (o *ClusterCustomerAddress) HasLat() bool`

HasLat returns a boolean if a field has been set.

### GetStreetHint

`func (o *ClusterCustomerAddress) GetStreetHint() string`

GetStreetHint returns the StreetHint field if non-nil, zero value otherwise.

### GetStreetHintOk

`func (o *ClusterCustomerAddress) GetStreetHintOk() (*string, bool)`

GetStreetHintOk returns a tuple with the StreetHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetHint

`func (o *ClusterCustomerAddress) SetStreetHint(v string)`

SetStreetHint sets StreetHint field to given value.

### HasStreetHint

`func (o *ClusterCustomerAddress) HasStreetHint() bool`

HasStreetHint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


