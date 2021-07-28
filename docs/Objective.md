# Objective

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of objective function, i.e. &#x60;min&#x60; or &#x60;min-max&#x60;.   * &#x60;min&#x60;: Minimizes the objective value.  * &#x60;min-max&#x60;: Minimizes the maximum objective value.  For instance, &#x60;min&#x60; -&gt; &#x60;vehicles&#x60; minimizes the number of employed vehicles. &#x60;min&#x60; -&gt; &#x60;completion_time&#x60; minimizes the sum of your vehicle routes&#39; completion time.  If you use, for example, &#x60;min-max&#x60; -&gt; &#x60;completion_time&#x60;, it minimizes the maximum of your vehicle routes&#39; completion time, i.e. it minimizes the overall makespan. This only makes sense if you have more than one vehicle. In case of one vehicle, switching from &#x60;min&#x60; to &#x60;min-max&#x60; should not have any impact. If you have more than one vehicle, then the algorithm tries to constantly move stops from one vehicle to another such that the completion time of longest vehicle route can be further reduced. For example, if you have one vehicle that takes 8 hours to serve all customers, adding another vehicle (and using &#x60;min-max&#x60;) might halve the time to serve all customers to 4 hours. However, this usually comes with higher transport costs.  If you want to minimize &#x60;vehicles&#x60; first and, second, &#x60;completion_time&#x60;, you can also combine different objectives like this:  &#x60;&#x60;&#x60;json \&quot;objectives\&quot; : [    {       \&quot;type\&quot;: \&quot;min\&quot;,       \&quot;value\&quot;: \&quot;vehicles\&quot;    },    {       \&quot;type\&quot;: \&quot;min\&quot;,       \&quot;value\&quot;: \&quot;completion_time\&quot;    } ] &#x60;&#x60;&#x60;  If you want to balance activities or the number of stops among all employed drivers, you need to specify it as follows:  &#x60;&#x60;&#x60;json \&quot;objectives\&quot; : [    {       \&quot;type\&quot;: \&quot;min-max\&quot;,       \&quot;value\&quot;: \&quot;completion_time\&quot;    },    {       \&quot;type\&quot;: \&quot;min-max\&quot;,       \&quot;value\&quot;: \&quot;activities\&quot;    } ] &#x60;&#x60;&#x60;  | [default to "min"]
**Value** | **string** | The value of the objective function. The objective value &#x60;transport_time&#x60; solely considers the time your drivers spend on the road, i.e. transport time. In contrary to &#x60;transport_time&#x60;, &#x60;completion_time&#x60; also takes waiting times at customer sites into account. The &#x60;completion_time&#x60; of a route is defined as the time from starting to ending the route, i.e. the route&#39;s transport time, the sum of waiting times plus the sum of activity durations. Note that choosing &#x60;transport_time&#x60; or &#x60;completion_time&#x60; only makes a difference if you specified time windows for your services/shipments since only in scenarios with time windows waiting times can occur. The objective value &#x60;vehicles&#x60; can only be used along with &#x60;min&#x60; and minimizes vehicles.  | [default to "transport_time"]

## Methods

### NewObjective

`func NewObjective(type_ string, value string, ) *Objective`

NewObjective instantiates a new Objective object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectiveWithDefaults

`func NewObjectiveWithDefaults() *Objective`

NewObjectiveWithDefaults instantiates a new Objective object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Objective) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Objective) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Objective) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *Objective) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Objective) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Objective) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


