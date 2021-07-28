# ClusterCustomer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | id of customer | [optional] 
**Address** | Pointer to [**ClusterCustomerAddress**](ClusterCustomerAddress.md) |  | [optional] 
**Quantity** | Pointer to **float32** | demand of customer | [optional] 

## Methods

### NewClusterCustomer

`func NewClusterCustomer() *ClusterCustomer`

NewClusterCustomer instantiates a new ClusterCustomer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterCustomerWithDefaults

`func NewClusterCustomerWithDefaults() *ClusterCustomer`

NewClusterCustomerWithDefaults instantiates a new ClusterCustomer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterCustomer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterCustomer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterCustomer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterCustomer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAddress

`func (o *ClusterCustomer) GetAddress() ClusterCustomerAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ClusterCustomer) GetAddressOk() (*ClusterCustomerAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ClusterCustomer) SetAddress(v ClusterCustomerAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ClusterCustomer) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetQuantity

`func (o *ClusterCustomer) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ClusterCustomer) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ClusterCustomer) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ClusterCustomer) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


