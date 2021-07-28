/*
 * GraphHopper Directions API
 *
 *  With the [GraphHopper Directions API](https://www.graphhopper.com/products/) you can integrate A-to-B route planning, turn-by-turn navigation, route optimization, isochrone calculations and other tools in your application.  The GraphHopper Directions API consists of the following RESTful web services:   * [Routing API](#tag/Routing-API),  * [Route Optimization API](#tag/Route-Optimization-API),  * [Isochrone API](#tag/Isochrone-API),  * [Map Matching API](#tag/Map-Matching-API),  * [Matrix API](#tag/Matrix-API),  * [Geocoding API](#tag/Geocoding-API) and  * [Cluster API](#tag/Cluster-API).  # Explore our APIs  ## Get started  1. [Sign up for GraphHopper](https://support.graphhopper.com/a/solutions/articles/44001976025) 2. [Create an API key](https://support.graphhopper.com/a/solutions/articles/44001976027)  Each API part has its own documentation. Jump to the desired API part and learn about the API through the given examples and tutorials.  In addition, for each API there are specific sample requests that you can send via Insomnia or Postman to see what the requests and responses look like.  ## Insomnia  To explore our APIs with [Insomnia](https://insomnia.rest/), follow these steps:  1. Open Insomnia and Import [our workspace](https://raw.githubusercontent.com/graphhopper/directions-api-doc/master/web/restclients/GraphHopper-Direction-API-Insomnia.json). 2. Specify [your API key](https://graphhopper.com/dashboard/#/register) in your workspace: Manage Environments -> Base Environment -> `\"api_key\": your API key` 3. Start exploring  ![Insomnia](./img/insomnia.png)  ## Postman  To explore our APIs with [Postman](https://www.getpostman.com/), follow these steps:  1. Import our [request collections](https://raw.githubusercontent.com/graphhopper/directions-api-doc/master/web/restclients/graphhopper_directions_api.postman_collection.json) as well as our [environment file](https://raw.githubusercontent.com/graphhopper/directions-api-doc/master/web/restclients/graphhopper_directions_api.postman_environment.json). 2. Specify [your API key](https://graphhopper.com/dashboard/#/register) in your environment: `\"api_key\": your API key` 3. Start exploring  ![Postman](./img/postman.png)  ## API Client Libraries  To speed up development and make coding easier, we offer the following client libraries:   * [JavaScript client](https://github.com/graphhopper/directions-api-js-client) - try the [live examples](https://graphhopper.com/api/1/examples/)  * [Others](https://github.com/graphhopper/directions-api-clients) like C#, Ruby, PHP, Python, ... automatically created for the Route Optimization API  ### Bandwidth reduction  If you create your own client, make sure it supports http/2 and gzipped responses for best speed.  If you use the Matrix, the Route Optimization API or the Cluster API and want to solve large problems, we recommend you to reduce bandwidth by [compressing your POST request](https://gist.github.com/karussell/82851e303ea7b3459b2dea01f18949f4) and specifying the header as follows: `Content-Encoding: gzip`. This will also avoid the HTTP 413 error \"Request Entity Too Large\".  ## Contact Us  If you have problems or questions, please read the following information:  - [FAQ](https://graphhopper.com/api/1/docs/FAQ/) - [Public forum](https://discuss.graphhopper.com/c/directions-api) - [Contact us](https://www.graphhopper.com/contact-form/) - [GraphHopper Status Page](https://status.graphhopper.com/)  To stay informed about the latest developments, you can  - follow us on [twitter](https://twitter.com/graphhopper/), - read [our blog](https://graphhopper.com/blog/), - watch [our documentation repository](https://github.com/graphhopper/directions-api-doc), - sign up for our newsletter or - [our forum](https://discuss.graphhopper.com/c/directions-api).  Select the channel you like the most.   # Map Data and Routing Profiles  Currently, our main data source is [OpenStreetMap](https://www.openstreetmap.org). We also integrated other network data providers. This chapter gives an overview about the options you have.  ## OpenStreetMap  #### Geographical Coverage  [OpenStreetMap](https://www.openstreetmap.org) covers the whole world. If you want to see for yourself if we can provide data suitable for your region, please visit [GraphHopper Maps](https://graphhopper.com/maps/). You can edit and modify OpenStreetMap data if you find that important information is missing, e.g. a weight limit for a bridge. [Here](https://wiki.openstreetmap.org/wiki/Beginners%27_guide) is a beginner's guide that shows how to add data. If you have edited data, we usually consider your data after 1 week at the latest.  #### Supported Vehicle Profiles  The Routing, Matrix and Route Optimization APIs support the following vehicle profiles:  Name       | Description           | Restrictions              | Icon -----------|:----------------------|:--------------------------|:--------------------------------------------------------- car        | Car mode              | car access                | ![car image](https://graphhopper.com/maps/img/car.png) small_truck| Small truck like a Mercedes Sprinter, Ford Transit or Iveco Daily | height=2.7m, width=2+0.4m, length=5.5m, weight=2080+1400 kg | ![small truck image](https://graphhopper.com/maps/img/small_truck.png) truck      | Truck like a MAN or Mercedes-Benz Actros | height=3.7m, width=2.6+0.5m, length=12m, weight=13000 + 13000 kg, hgv=yes, 3 Axes | ![truck image](https://graphhopper.com/maps/img/truck.png) scooter    | Moped mode | Fast inner city, often used for food delivery, is able to ignore certain bollards, maximum speed of roughly 50km/h | ![scooter image](https://graphhopper.com/maps/img/scooter.png) foot       | Pedestrian or walking without dangerous [SAC-scales](https://wiki.openstreetmap.org/wiki/Key:sac_scale) | foot access         | ![foot image](https://graphhopper.com/maps/img/foot.png) hike       | Pedestrian or walking with priority for more beautiful hiking tours and potentially a bit longer than `foot`. Walking duration is influenced by elevation differences.  | foot access         | ![hike image](https://graphhopper.com/maps/img/hike.png) bike       | Trekking bike avoiding hills | bike access  | ![bike image](https://graphhopper.com/maps/img/bike.png) mtb        | Mountainbike          | bike access         | ![Mountainbike image](https://graphhopper.com/maps/img/mtb.png) racingbike| Bike preferring roads | bike access         | ![racingbike image](https://graphhopper.com/maps/img/racingbike.png)  Please note:   * all motor vehicles (`car`, `small_truck`, `truck` and `scooter`) support turn restrictions via `turn_costs=true`  * the free package supports only the vehicle profiles `car`, `bike` or `foot`  * up to 2 different vehicle profiles can be used in a single optimization request. The number of vehicles is unaffected and depends on your subscription.  * we offer custom vehicle profiles with different properties, different speed profiles or different access options. To find out more about custom profiles, please [contact us](https://www.graphhopper.com/contact-form/).  * a sophisticated `motorcycle` profile is available up on request. It is powered by the [Kurviger](https://kurviger.de/en) Routing API and favors curves and slopes while avoiding cities and highways.   ## TomTom  If you want to include traffic, you can purchase the TomTom Add-on. This Add-on only uses TomTom's road network and historical traffic information. Live traffic is not yet considered. If you are interested to learn how we consider traffic information, we recommend that you read [this article](https://www.graphhopper.com/blog/2017/11/06/time-dependent-optimization/).  Please note the following:   * Currently we only offer this for our [Route Optimization API](#tag/Route-Optimization-API).  * In addition to our terms, you need to accept TomTom's [End User License Aggreement](https://www.graphhopper.com/tomtom-end-user-license-agreement/).  * We do *not* use TomTom's web services. We only use their data with our software.   [Contact us](https://www.graphhopper.com/contact-form/) for more details.  #### Geographical Coverage  We offer  - Europe including Russia - North, Central and South America - Saudi Arabia - United Arab Emirates - South Africa - Australia  #### Supported Vehicle Profiles  Name       | Description           | Restrictions              | Icon -----------|:----------------------|:--------------------------|:--------------------------------------------------------- car        | Car mode              | car access                | ![car image](https://graphhopper.com/maps/img/car.png) small_truck| Small truck like a Mercedes Sprinter, Ford Transit or Iveco Daily | height=2.7m, width=2+0.4m, length=5.5m, weight=2080+1400 kg | ![small truck image](https://graphhopper.com/maps/img/small_truck.png) 
 *
 * API version: 1.0.0
 * Contact: support@graphhopper.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"reflect"
)

// Linger please
var (
	_ _context.Context
)

// RoutingAPIApiService RoutingAPIApi service
type RoutingAPIApiService service

type ApiGetRouteRequest struct {
	ctx _context.Context
	ApiService *RoutingAPIApiService
	point *[]string
	pointHint *[]string
	snapPrevention *[]string
	vehicle *VehicleProfileId
	curbside *[]string
	turnCosts *bool
	locale *string
	elevation *bool
	details *[]string
	optimize *string
	instructions *bool
	calcPoints *bool
	debug *bool
	pointsEncoded *bool
	chDisable *bool
	weighting *string
	heading *[]int32
	headingPenalty *int32
	passThrough *bool
	blockArea *string
	avoid *string
	algorithm *string
	roundTripDistance *int32
	roundTripSeed *int64
	alternativeRouteMaxPaths *int32
	alternativeRouteMaxWeightFactor *float32
	alternativeRouteMaxShareFactor *float32
}

func (r ApiGetRouteRequest) Point(point []string) ApiGetRouteRequest {
	r.point = &point
	return r
}
func (r ApiGetRouteRequest) PointHint(pointHint []string) ApiGetRouteRequest {
	r.pointHint = &pointHint
	return r
}
func (r ApiGetRouteRequest) SnapPrevention(snapPrevention []string) ApiGetRouteRequest {
	r.snapPrevention = &snapPrevention
	return r
}
func (r ApiGetRouteRequest) Vehicle(vehicle VehicleProfileId) ApiGetRouteRequest {
	r.vehicle = &vehicle
	return r
}
func (r ApiGetRouteRequest) Curbside(curbside []string) ApiGetRouteRequest {
	r.curbside = &curbside
	return r
}
func (r ApiGetRouteRequest) TurnCosts(turnCosts bool) ApiGetRouteRequest {
	r.turnCosts = &turnCosts
	return r
}
func (r ApiGetRouteRequest) Locale(locale string) ApiGetRouteRequest {
	r.locale = &locale
	return r
}
func (r ApiGetRouteRequest) Elevation(elevation bool) ApiGetRouteRequest {
	r.elevation = &elevation
	return r
}
func (r ApiGetRouteRequest) Details(details []string) ApiGetRouteRequest {
	r.details = &details
	return r
}
func (r ApiGetRouteRequest) Optimize(optimize string) ApiGetRouteRequest {
	r.optimize = &optimize
	return r
}
func (r ApiGetRouteRequest) Instructions(instructions bool) ApiGetRouteRequest {
	r.instructions = &instructions
	return r
}
func (r ApiGetRouteRequest) CalcPoints(calcPoints bool) ApiGetRouteRequest {
	r.calcPoints = &calcPoints
	return r
}
func (r ApiGetRouteRequest) Debug(debug bool) ApiGetRouteRequest {
	r.debug = &debug
	return r
}
func (r ApiGetRouteRequest) PointsEncoded(pointsEncoded bool) ApiGetRouteRequest {
	r.pointsEncoded = &pointsEncoded
	return r
}
func (r ApiGetRouteRequest) ChDisable(chDisable bool) ApiGetRouteRequest {
	r.chDisable = &chDisable
	return r
}
func (r ApiGetRouteRequest) Weighting(weighting string) ApiGetRouteRequest {
	r.weighting = &weighting
	return r
}
func (r ApiGetRouteRequest) Heading(heading []int32) ApiGetRouteRequest {
	r.heading = &heading
	return r
}
func (r ApiGetRouteRequest) HeadingPenalty(headingPenalty int32) ApiGetRouteRequest {
	r.headingPenalty = &headingPenalty
	return r
}
func (r ApiGetRouteRequest) PassThrough(passThrough bool) ApiGetRouteRequest {
	r.passThrough = &passThrough
	return r
}
func (r ApiGetRouteRequest) BlockArea(blockArea string) ApiGetRouteRequest {
	r.blockArea = &blockArea
	return r
}
func (r ApiGetRouteRequest) Avoid(avoid string) ApiGetRouteRequest {
	r.avoid = &avoid
	return r
}
func (r ApiGetRouteRequest) Algorithm(algorithm string) ApiGetRouteRequest {
	r.algorithm = &algorithm
	return r
}
func (r ApiGetRouteRequest) RoundTripDistance(roundTripDistance int32) ApiGetRouteRequest {
	r.roundTripDistance = &roundTripDistance
	return r
}
func (r ApiGetRouteRequest) RoundTripSeed(roundTripSeed int64) ApiGetRouteRequest {
	r.roundTripSeed = &roundTripSeed
	return r
}
func (r ApiGetRouteRequest) AlternativeRouteMaxPaths(alternativeRouteMaxPaths int32) ApiGetRouteRequest {
	r.alternativeRouteMaxPaths = &alternativeRouteMaxPaths
	return r
}
func (r ApiGetRouteRequest) AlternativeRouteMaxWeightFactor(alternativeRouteMaxWeightFactor float32) ApiGetRouteRequest {
	r.alternativeRouteMaxWeightFactor = &alternativeRouteMaxWeightFactor
	return r
}
func (r ApiGetRouteRequest) AlternativeRouteMaxShareFactor(alternativeRouteMaxShareFactor float32) ApiGetRouteRequest {
	r.alternativeRouteMaxShareFactor = &alternativeRouteMaxShareFactor
	return r
}

func (r ApiGetRouteRequest) Execute() (RouteResponse, *_nethttp.Response, error) {
	return r.ApiService.GetRouteExecute(r)
}

/*
 * GetRoute GET Route Endpoint
 * The GET request is the most simple one: just specify the parameter in the URL and you are done.
Can be tried directly in every browser.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetRouteRequest
 */
func (a *RoutingAPIApiService) GetRoute(ctx _context.Context) ApiGetRouteRequest {
	return ApiGetRouteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return RouteResponse
 */
func (a *RoutingAPIApiService) GetRouteExecute(r ApiGetRouteRequest) (RouteResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  RouteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoutingAPIApiService.GetRoute")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/route"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.point == nil {
		return localVarReturnValue, nil, reportError("point is required and must be specified")
	}

	{
		t := *r.point
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("point", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("point", parameterToString(t, "multi"))
		}
	}
	if r.pointHint != nil {
		t := *r.pointHint
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("point_hint", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("point_hint", parameterToString(t, "multi"))
		}
	}
	if r.snapPrevention != nil {
		t := *r.snapPrevention
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("snap_prevention", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("snap_prevention", parameterToString(t, "multi"))
		}
	}
	if r.vehicle != nil {
		localVarQueryParams.Add("vehicle", parameterToString(*r.vehicle, ""))
	}
	if r.curbside != nil {
		t := *r.curbside
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("curbside", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("curbside", parameterToString(t, "multi"))
		}
	}
	if r.turnCosts != nil {
		localVarQueryParams.Add("turn_costs", parameterToString(*r.turnCosts, ""))
	}
	if r.locale != nil {
		localVarQueryParams.Add("locale", parameterToString(*r.locale, ""))
	}
	if r.elevation != nil {
		localVarQueryParams.Add("elevation", parameterToString(*r.elevation, ""))
	}
	if r.details != nil {
		t := *r.details
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("details", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("details", parameterToString(t, "multi"))
		}
	}
	if r.optimize != nil {
		localVarQueryParams.Add("optimize", parameterToString(*r.optimize, ""))
	}
	if r.instructions != nil {
		localVarQueryParams.Add("instructions", parameterToString(*r.instructions, ""))
	}
	if r.calcPoints != nil {
		localVarQueryParams.Add("calc_points", parameterToString(*r.calcPoints, ""))
	}
	if r.debug != nil {
		localVarQueryParams.Add("debug", parameterToString(*r.debug, ""))
	}
	if r.pointsEncoded != nil {
		localVarQueryParams.Add("points_encoded", parameterToString(*r.pointsEncoded, ""))
	}
	if r.chDisable != nil {
		localVarQueryParams.Add("ch.disable", parameterToString(*r.chDisable, ""))
	}
	if r.weighting != nil {
		localVarQueryParams.Add("weighting", parameterToString(*r.weighting, ""))
	}
	if r.heading != nil {
		t := *r.heading
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("heading", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("heading", parameterToString(t, "multi"))
		}
	}
	if r.headingPenalty != nil {
		localVarQueryParams.Add("heading_penalty", parameterToString(*r.headingPenalty, ""))
	}
	if r.passThrough != nil {
		localVarQueryParams.Add("pass_through", parameterToString(*r.passThrough, ""))
	}
	if r.blockArea != nil {
		localVarQueryParams.Add("block_area", parameterToString(*r.blockArea, ""))
	}
	if r.avoid != nil {
		localVarQueryParams.Add("avoid", parameterToString(*r.avoid, ""))
	}
	if r.algorithm != nil {
		localVarQueryParams.Add("algorithm", parameterToString(*r.algorithm, ""))
	}
	if r.roundTripDistance != nil {
		localVarQueryParams.Add("round_trip.distance", parameterToString(*r.roundTripDistance, ""))
	}
	if r.roundTripSeed != nil {
		localVarQueryParams.Add("round_trip.seed", parameterToString(*r.roundTripSeed, ""))
	}
	if r.alternativeRouteMaxPaths != nil {
		localVarQueryParams.Add("alternative_route.max_paths", parameterToString(*r.alternativeRouteMaxPaths, ""))
	}
	if r.alternativeRouteMaxWeightFactor != nil {
		localVarQueryParams.Add("alternative_route.max_weight_factor", parameterToString(*r.alternativeRouteMaxWeightFactor, ""))
	}
	if r.alternativeRouteMaxShareFactor != nil {
		localVarQueryParams.Add("alternative_route.max_share_factor", parameterToString(*r.alternativeRouteMaxShareFactor, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("key", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v GHError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v GHError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v GHError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v GHError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 501 {
			var v GHError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPostRouteRequest struct {
	ctx _context.Context
	ApiService *RoutingAPIApiService
	routeRequest *RouteRequest
}

func (r ApiPostRouteRequest) RouteRequest(routeRequest RouteRequest) ApiPostRouteRequest {
	r.routeRequest = &routeRequest
	return r
}

func (r ApiPostRouteRequest) Execute() (RouteResponse, *_nethttp.Response, error) {
	return r.ApiService.PostRouteExecute(r)
}

/*
 * PostRoute POST Route Endpoint
 * Please see the [GET endpoint](#operation/getRoute) for a simpler method on how to get started.
If you are familiar with POST requests and JSON then do not hesitate to continue here.

Especially when you use many locations you should get familiar with this POST endpoint as the GET endpoint
has an URL length limitation. Additionally the request of this POST endpoint can be compressed and can slightly
speed up the request.

To do a request you send JSON data. Both request scenarios GET and POST are identical except that all singular parameter names are named as their plural for a POST request.
The effected parameters are: `points`, `point_hints` and `snap_preventions`.

**Please note that in opposite to the GET endpoint, points are specified in the order of `longitude, latitude`**.

For example `point=10,11&point=20,22` will be converted to the `points` array (plural):
```json
{ "points": [[11,10], [22,20]] }
```
Note again that also the order changes from `[latitude,longitude]` to `[longitude,latitude]`
similar to [GeoJson](http://geojson.org/geojson-spec.html#examples).

Example:
```bash
curl -X POST -H "Content-Type: application/json" "https://graphhopper.com/api/1/route?key=[YOUR_KEY]" -d '{"elevation":false,"points":[[-0.087891,51.534377],[-0.090637,51.467697]],"vehicle":"car"}'
```

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiPostRouteRequest
 */
func (a *RoutingAPIApiService) PostRoute(ctx _context.Context) ApiPostRouteRequest {
	return ApiPostRouteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return RouteResponse
 */
func (a *RoutingAPIApiService) PostRouteExecute(r ApiPostRouteRequest) (RouteResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  RouteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoutingAPIApiService.PostRoute")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/route"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.routeRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("key", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v GHError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v GHError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v GHError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v GHError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 501 {
			var v GHError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiRouteInfoGetRequest struct {
	ctx _context.Context
	ApiService *RoutingAPIApiService
}


func (r ApiRouteInfoGetRequest) Execute() (InfoResponse, *_nethttp.Response, error) {
	return r.ApiService.RouteInfoGetExecute(r)
}

/*
 * RouteInfoGet Coverage information
 * Use this to find out details about the supported vehicle profiles and features, or if you just need to ping the server.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiRouteInfoGetRequest
 */
func (a *RoutingAPIApiService) RouteInfoGet(ctx _context.Context) ApiRouteInfoGetRequest {
	return ApiRouteInfoGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return InfoResponse
 */
func (a *RoutingAPIApiService) RouteInfoGetExecute(r ApiRouteInfoGetRequest) (InfoResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InfoResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoutingAPIApiService.RouteInfoGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/route/info"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("key", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
