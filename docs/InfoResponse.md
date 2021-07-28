# InfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | The version of the GraphHopper server that provided this response. This is not related to the API version.  | [optional] 
**Bbox** | Pointer to **string** | The bounding box of the geographical area covered by this GraphHopper instance. Format: &#x60;\&quot;minLon,minLat,maxLon,maxLat\&quot;  | [optional] 
**Features** | Pointer to **map[string]interface{}** | The supported features, such as elevation, per vehicle profile.  | [optional] 

## Methods

### NewInfoResponse

`func NewInfoResponse() *InfoResponse`

NewInfoResponse instantiates a new InfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfoResponseWithDefaults

`func NewInfoResponseWithDefaults() *InfoResponse`

NewInfoResponseWithDefaults instantiates a new InfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *InfoResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InfoResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InfoResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *InfoResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBbox

`func (o *InfoResponse) GetBbox() string`

GetBbox returns the Bbox field if non-nil, zero value otherwise.

### GetBboxOk

`func (o *InfoResponse) GetBboxOk() (*string, bool)`

GetBboxOk returns a tuple with the Bbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBbox

`func (o *InfoResponse) SetBbox(v string)`

SetBbox sets Bbox field to given value.

### HasBbox

`func (o *InfoResponse) HasBbox() bool`

HasBbox returns a boolean if a field has been set.

### GetFeatures

`func (o *InfoResponse) GetFeatures() map[string]interface{}`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *InfoResponse) GetFeaturesOk() (*map[string]interface{}, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *InfoResponse) SetFeatures(v map[string]interface{})`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *InfoResponse) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


