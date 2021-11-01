# Go API client for v1

Public (unauthenticated) API to integrate with your local Apex aquarium controller (AOS 5 compatible).


## Documentation for API Endpoints

All URIs are relative to *http://apex.local*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**GetDatalog**](docs/DefaultApi.md#getdatalog) | **Get** /cgi-bin/datalog.json | Fetch input log records.
*DefaultApi* | [**GetOutlog**](docs/DefaultApi.md#getoutlog) | **Get** /cgi-bin/outlog.json | Fetch output log records.
*DefaultApi* | [**GetStatus**](docs/DefaultApi.md#getstatus) | **Get** /cgi-bin/status.json | Fetch system info and current input/output states.


## Documentation For Models

 - [DatalogContainer](docs/DatalogContainer.md)
 - [DatalogData](docs/DatalogData.md)
 - [DatalogRecord](docs/DatalogRecord.md)
 - [DatalogResponse](docs/DatalogResponse.md)
 - [OutlogContainer](docs/OutlogContainer.md)
 - [OutlogRecord](docs/OutlogRecord.md)
 - [OutlogResponse](docs/OutlogResponse.md)
 - [StatusContainer](docs/StatusContainer.md)
 - [StatusFeed](docs/StatusFeed.md)
 - [StatusInput](docs/StatusInput.md)
 - [StatusLink](docs/StatusLink.md)
 - [StatusOutput](docs/StatusOutput.md)
 - [StatusPower](docs/StatusPower.md)
 - [StatusResponse](docs/StatusResponse.md)


## Documentation For Authorization

 Endpoints do not require authorization.


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
