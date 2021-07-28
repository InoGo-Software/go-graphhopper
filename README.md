# Go API client for openapi


With the [GraphHopper Directions API](https://www.graphhopper.com/products/) you can integrate A-to-B route planning, turn-by-turn navigation,
route optimization, isochrone calculations and other tools in your application.

The GraphHopper Directions API consists of the following RESTful web services:

 * [Routing API](#tag/Routing-API),
 * [Route Optimization API](#tag/Route-Optimization-API),
 * [Isochrone API](#tag/Isochrone-API),
 * [Map Matching API](#tag/Map-Matching-API),
 * [Matrix API](#tag/Matrix-API),
 * [Geocoding API](#tag/Geocoding-API) and
 * [Cluster API](#tag/Cluster-API).

# Explore our APIs

## Get started

1. [Sign up for GraphHopper](https://support.graphhopper.com/a/solutions/articles/44001976025)
2. [Create an API key](https://support.graphhopper.com/a/solutions/articles/44001976027)

Each API part has its own documentation. Jump to the desired API part and learn about the API through the given examples and tutorials.

In addition, for each API there are specific sample requests that you can send via Insomnia or Postman to see what the requests and responses look like.

## Insomnia

To explore our APIs with [Insomnia](https://insomnia.rest/), follow these steps:

1. Open Insomnia and Import [our workspace](https://raw.githubusercontent.com/graphhopper/directions-api-doc/master/web/restclients/GraphHopper-Direction-API-Insomnia.json).
2. Specify [your API key](https://graphhopper.com/dashboard/#/register) in your workspace: Manage Environments -> Base Environment -> `\"api_key\": your API key`
3. Start exploring

![Insomnia](./img/insomnia.png)

## Postman

To explore our APIs with [Postman](https://www.getpostman.com/), follow these steps:

1. Import our [request collections](https://raw.githubusercontent.com/graphhopper/directions-api-doc/master/web/restclients/graphhopper_directions_api.postman_collection.json) as well as our [environment file](https://raw.githubusercontent.com/graphhopper/directions-api-doc/master/web/restclients/graphhopper_directions_api.postman_environment.json).
2. Specify [your API key](https://graphhopper.com/dashboard/#/register) in your environment: `\"api_key\": your API key`
3. Start exploring

![Postman](./img/postman.png)

## API Client Libraries

To speed up development and make coding easier, we offer the following client libraries:

 * [JavaScript client](https://github.com/graphhopper/directions-api-js-client) - try the [live examples](https://graphhopper.com/api/1/examples/)
 * [Others](https://github.com/graphhopper/directions-api-clients) like C#, Ruby, PHP, Python, ... automatically created for the Route Optimization API

### Bandwidth reduction

If you create your own client, make sure it supports http/2 and gzipped responses for best speed.

If you use the Matrix, the Route Optimization API or the Cluster API and want to solve large problems, we recommend you to reduce bandwidth
by [compressing your POST request](https://gist.github.com/karussell/82851e303ea7b3459b2dea01f18949f4)
and specifying the header as follows: `Content-Encoding: gzip`. This will also avoid the HTTP 413 error \"Request Entity Too Large\".

## Contact Us

If you have problems or questions, please read the following information:

- [FAQ](https://graphhopper.com/api/1/docs/FAQ/)
- [Public forum](https://discuss.graphhopper.com/c/directions-api)
- [Contact us](https://www.graphhopper.com/contact-form/)
- [GraphHopper Status Page](https://status.graphhopper.com/)

To stay informed about the latest developments, you can

- follow us on [twitter](https://twitter.com/graphhopper/),
- read [our blog](https://graphhopper.com/blog/),
- watch [our documentation repository](https://github.com/graphhopper/directions-api-doc),
- sign up for our newsletter or
- [our forum](https://discuss.graphhopper.com/c/directions-api).

Select the channel you like the most.


# Map Data and Routing Profiles

Currently, our main data source is [OpenStreetMap](https://www.openstreetmap.org). We also integrated other network data providers.
This chapter gives an overview about the options you have.

## OpenStreetMap

#### Geographical Coverage

[OpenStreetMap](https://www.openstreetmap.org) covers the whole world. If you want to see for yourself if we can provide data suitable for your region,
please visit [GraphHopper Maps](https://graphhopper.com/maps/).
You can edit and modify OpenStreetMap data if you find that important information is missing, e.g. a weight limit for a bridge.
[Here](https://wiki.openstreetmap.org/wiki/Beginners%27_guide) is a beginner's guide that shows how to add data. If you have edited data, we usually consider your data after 1 week at the latest.

#### Supported Vehicle Profiles

The Routing, Matrix and Route Optimization APIs support the following vehicle profiles:

Name       | Description           | Restrictions              | Icon
-----------|:----------------------|:--------------------------|:---------------------------------------------------------
car        | Car mode              | car access                | ![car image](https://graphhopper.com/maps/img/car.png)
small_truck| Small truck like a Mercedes Sprinter, Ford Transit or Iveco Daily | height=2.7m, width=2+0.4m, length=5.5m, weight=2080+1400 kg | ![small truck image](https://graphhopper.com/maps/img/small_truck.png)
truck      | Truck like a MAN or Mercedes-Benz Actros | height=3.7m, width=2.6+0.5m, length=12m, weight=13000 + 13000 kg, hgv=yes, 3 Axes | ![truck image](https://graphhopper.com/maps/img/truck.png)
scooter    | Moped mode | Fast inner city, often used for food delivery, is able to ignore certain bollards, maximum speed of roughly 50km/h | ![scooter image](https://graphhopper.com/maps/img/scooter.png)
foot       | Pedestrian or walking without dangerous [SAC-scales](https://wiki.openstreetmap.org/wiki/Key:sac_scale) | foot access         | ![foot image](https://graphhopper.com/maps/img/foot.png)
hike       | Pedestrian or walking with priority for more beautiful hiking tours and potentially a bit longer than `foot`. Walking duration is influenced by elevation differences.  | foot access         | ![hike image](https://graphhopper.com/maps/img/hike.png)
bike       | Trekking bike avoiding hills | bike access  | ![bike image](https://graphhopper.com/maps/img/bike.png)
mtb        | Mountainbike          | bike access         | ![Mountainbike image](https://graphhopper.com/maps/img/mtb.png)
racingbike| Bike preferring roads | bike access         | ![racingbike image](https://graphhopper.com/maps/img/racingbike.png)

Please note:

 * all motor vehicles (`car`, `small_truck`, `truck` and `scooter`) support turn restrictions via `turn_costs=true`
 * the free package supports only the vehicle profiles `car`, `bike` or `foot`
 * up to 2 different vehicle profiles can be used in a single optimization request. The number of vehicles is unaffected and depends on your subscription.
 * we offer custom vehicle profiles with different properties, different speed profiles or different access options. To find out more about custom profiles, please [contact us](https://www.graphhopper.com/contact-form/).
 * a sophisticated `motorcycle` profile is available up on request. It is powered by the [Kurviger](https://kurviger.de/en) Routing API and favors curves and slopes while avoiding cities and highways.
 
## TomTom

If you want to include traffic, you can purchase the TomTom Add-on.
This Add-on only uses TomTom's road network and historical traffic information.
Live traffic is not yet considered. If you are interested to learn how we consider traffic information, we recommend that you read [this article](https://www.graphhopper.com/blog/2017/11/06/time-dependent-optimization/).

Please note the following:

 * Currently we only offer this for our [Route Optimization API](#tag/Route-Optimization-API).
 * In addition to our terms, you need to accept TomTom's [End User License Aggreement](https://www.graphhopper.com/tomtom-end-user-license-agreement/).
 * We do *not* use TomTom's web services. We only use their data with our software.
 
[Contact us](https://www.graphhopper.com/contact-form/) for more details.

#### Geographical Coverage

We offer

- Europe including Russia
- North, Central and South America
- Saudi Arabia
- United Arab Emirates
- South Africa
- Australia

#### Supported Vehicle Profiles

Name       | Description           | Restrictions              | Icon
-----------|:----------------------|:--------------------------|:---------------------------------------------------------
car        | Car mode              | car access                | ![car image](https://graphhopper.com/maps/img/car.png)
small_truck| Small truck like a Mercedes Sprinter, Ford Transit or Iveco Daily | height=2.7m, width=2+0.4m, length=5.5m, weight=2080+1400 kg | ![small truck image](https://graphhopper.com/maps/img/small_truck.png)


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://www.graphhopper.com/](https://www.graphhopper.com/)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./openapi"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://graphhopper.com/api/1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ClusterAPIApi* | [**AsyncClusteringProblem**](docs/ClusterAPIApi.md#asyncclusteringproblem) | **Post** /cluster/calculate | Batch Cluster Endpoint
*ClusterAPIApi* | [**GetClusterSolution**](docs/ClusterAPIApi.md#getclustersolution) | **Get** /cluster/solution/{jobId} | GET Batch Solution Endpoint
*ClusterAPIApi* | [**SolveClusteringProblem**](docs/ClusterAPIApi.md#solveclusteringproblem) | **Post** /cluster | POST Cluster Endpoint
*GeocodingAPIApi* | [**GetGeocode**](docs/GeocodingAPIApi.md#getgeocode) | **Get** /geocode | Geocoding Endpoint
*IsochroneAPIApi* | [**GetIsochrone**](docs/IsochroneAPIApi.md#getisochrone) | **Get** /isochrone | Isochrone Endpoint
*MapMatchingAPIApi* | [**PostGPX**](docs/MapMatchingAPIApi.md#postgpx) | **Post** /match | Map-match a GPX file
*MatrixAPIApi* | [**CalculateMatrix**](docs/MatrixAPIApi.md#calculatematrix) | **Post** /matrix/calculate | Batch Matrix Endpoint
*MatrixAPIApi* | [**GetMatrix**](docs/MatrixAPIApi.md#getmatrix) | **Get** /matrix | GET Matrix Endpoint
*MatrixAPIApi* | [**GetMatrixSolution**](docs/MatrixAPIApi.md#getmatrixsolution) | **Get** /matrix/solution/{jobId} | GET Batch Matrix Endpoint
*MatrixAPIApi* | [**PostMatrix**](docs/MatrixAPIApi.md#postmatrix) | **Post** /matrix | POST Matrix Endpoint
*RouteOptimizationAPIApi* | [**AsyncVRP**](docs/RouteOptimizationAPIApi.md#asyncvrp) | **Post** /vrp/optimize | POST route optimization problem (batch mode)
*RouteOptimizationAPIApi* | [**GetSolution**](docs/RouteOptimizationAPIApi.md#getsolution) | **Get** /vrp/solution/{jobId} | GET the solution (batch mode)
*RouteOptimizationAPIApi* | [**SolveVRP**](docs/RouteOptimizationAPIApi.md#solvevrp) | **Post** /vrp | POST route optimization problem
*RoutingAPIApi* | [**GetRoute**](docs/RoutingAPIApi.md#getroute) | **Get** /route | GET Route Endpoint
*RoutingAPIApi* | [**PostRoute**](docs/RoutingAPIApi.md#postroute) | **Post** /route | POST Route Endpoint
*RoutingAPIApi* | [**RouteInfoGet**](docs/RoutingAPIApi.md#routeinfoget) | **Get** /route/info | Coverage information


## Documentation For Models

 - [Activity](docs/Activity.md)
 - [Address](docs/Address.md)
 - [Algorithm](docs/Algorithm.md)
 - [BadRequest](docs/BadRequest.md)
 - [Cluster](docs/Cluster.md)
 - [ClusterConfiguration](docs/ClusterConfiguration.md)
 - [ClusterConfigurationClustering](docs/ClusterConfigurationClustering.md)
 - [ClusterConfigurationRouting](docs/ClusterConfigurationRouting.md)
 - [ClusterCustomer](docs/ClusterCustomer.md)
 - [ClusterCustomerAddress](docs/ClusterCustomerAddress.md)
 - [ClusterRequest](docs/ClusterRequest.md)
 - [ClusterResponse](docs/ClusterResponse.md)
 - [Configuration](docs/Configuration.md)
 - [CostMatrix](docs/CostMatrix.md)
 - [CostMatrixData](docs/CostMatrixData.md)
 - [CostMatrixDataInfo](docs/CostMatrixDataInfo.md)
 - [Detail](docs/Detail.md)
 - [DriveTimeBreak](docs/DriveTimeBreak.md)
 - [ErrorMessage](docs/ErrorMessage.md)
 - [GHError](docs/GHError.md)
 - [GHErrorHints](docs/GHErrorHints.md)
 - [GeocodingLocation](docs/GeocodingLocation.md)
 - [GeocodingPoint](docs/GeocodingPoint.md)
 - [GeocodingResponse](docs/GeocodingResponse.md)
 - [GroupRelation](docs/GroupRelation.md)
 - [InfoResponse](docs/InfoResponse.md)
 - [InlineResponse404](docs/InlineResponse404.md)
 - [InternalErrorMessage](docs/InternalErrorMessage.md)
 - [IsochroneResponse](docs/IsochroneResponse.md)
 - [IsochroneResponsePolygon](docs/IsochroneResponsePolygon.md)
 - [IsochroneResponsePolygonProperties](docs/IsochroneResponsePolygonProperties.md)
 - [JobId](docs/JobId.md)
 - [JobRelation](docs/JobRelation.md)
 - [LineString](docs/LineString.md)
 - [MatrixRequest](docs/MatrixRequest.md)
 - [MatrixResponse](docs/MatrixResponse.md)
 - [MatrixResponseHints](docs/MatrixResponseHints.md)
 - [Objective](docs/Objective.md)
 - [Polygon](docs/Polygon.md)
 - [Request](docs/Request.md)
 - [Response](docs/Response.md)
 - [ResponseAddress](docs/ResponseAddress.md)
 - [ResponseInfo](docs/ResponseInfo.md)
 - [Route](docs/Route.md)
 - [RoutePoint](docs/RoutePoint.md)
 - [RouteRequest](docs/RouteRequest.md)
 - [RouteResponse](docs/RouteResponse.md)
 - [RouteResponsePath](docs/RouteResponsePath.md)
 - [RouteResponsePathInstructions](docs/RouteResponsePathInstructions.md)
 - [Routing](docs/Routing.md)
 - [Service](docs/Service.md)
 - [Shipment](docs/Shipment.md)
 - [SnappedWaypoint](docs/SnappedWaypoint.md)
 - [Solution](docs/Solution.md)
 - [SolutionUnassigned](docs/SolutionUnassigned.md)
 - [Stop](docs/Stop.md)
 - [SymmetricalMatrixRequest](docs/SymmetricalMatrixRequest.md)
 - [TimeWindow](docs/TimeWindow.md)
 - [TimeWindowBreak](docs/TimeWindowBreak.md)
 - [Vehicle](docs/Vehicle.md)
 - [VehicleProfileId](docs/VehicleProfileId.md)
 - [VehicleType](docs/VehicleType.md)


## Documentation For Authorization



### api_key

- **Type**: API key
- **API key parameter name**: key
- **Location**: URL query string

Note, each API key must be added to a map of `map[string]APIKey` where the key is: key and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@graphhopper.com

