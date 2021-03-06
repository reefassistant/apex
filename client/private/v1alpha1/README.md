# Go API client for v1alpha1

Private (authenticated) API to integrate with your local Apex aquarium controller (AOS 5 compatible).


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0-alpha.1
- Package version: 1.0.0-alpha.1
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./v1alpha1"
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
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
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

All URIs are relative to *http://apex.local*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**GetConfig**](docs/DefaultApi.md#getconfig) | **Get** /rest/config | Fetch configuration.
*DefaultApi* | [**GetStatus**](docs/DefaultApi.md#getstatus) | **Get** /rest/status | Fetch status.
*DefaultApi* | [**GetUser**](docs/DefaultApi.md#getuser) | **Get** /rest/user | Fetch user information.
*DefaultApi* | [**LoginUser**](docs/DefaultApi.md#loginuser) | **Post** /rest/login | User login and return authentication cookie.


## Documentation For Models

 - [CalibrationConfig](docs/CalibrationConfig.md)
 - [ClockConfig](docs/ClockConfig.md)
 - [DisplayConfig](docs/DisplayConfig.md)
 - [GetConfigResponse](docs/GetConfigResponse.md)
 - [GetStatusResponse](docs/GetStatusResponse.md)
 - [GetUserResponse](docs/GetUserResponse.md)
 - [InputConfig](docs/InputConfig.md)
 - [LoginRequest](docs/LoginRequest.md)
 - [LoginResponse](docs/LoginResponse.md)
 - [MiscellaneousConfig](docs/MiscellaneousConfig.md)
 - [ModuleConfig](docs/ModuleConfig.md)
 - [NetworkConfig](docs/NetworkConfig.md)
 - [OutputConfig](docs/OutputConfig.md)
 - [ProfileConfig](docs/ProfileConfig.md)
 - [SeasonConfig](docs/SeasonConfig.md)
 - [SystemStatus](docs/SystemStatus.md)


## Documentation For Authorization



### cookieAuth

- **Type**: API key
- **API key parameter name**: connect.sid
- **Location**: 

Note, each API key must be added to a map of `map[string]APIKey` where the key is: connect.sid and passed in as the auth context for each request.


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



