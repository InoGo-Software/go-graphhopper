# GroupRelation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Specifies the type of relation. It must be either of type &#x60;in_sequence&#x60; or &#x60;in_direct_sequence&#x60;.  | 
**Groups** | **[]string** | An array of groups that should be related | 

## Methods

### NewGroupRelation

`func NewGroupRelation(type_ string, groups []string, ) *GroupRelation`

NewGroupRelation instantiates a new GroupRelation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupRelationWithDefaults

`func NewGroupRelationWithDefaults() *GroupRelation`

NewGroupRelationWithDefaults instantiates a new GroupRelation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GroupRelation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GroupRelation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GroupRelation) SetType(v string)`

SetType sets Type field to given value.


### GetGroups

`func (o *GroupRelation) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *GroupRelation) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *GroupRelation) SetGroups(v []string)`

SetGroups sets Groups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


