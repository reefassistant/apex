# Neptune Apex Go API Client

This package provides a Go API client to interface with your Neptune Apex aquarium controller over your local network.

> :warning: **This API client is not designed nor to be used against Apex Fusion !**

## Usage

Below is a conceptual example how to instantiate a client object and perform requests.

```go
client, _ := apex.NewDefaultClient("http://your-apex-hostname.local")
resp, _, _ := client.PublicV1().DefaultApi.GetStatus(ctx).Execute()
fmt.Printf("software: %q\n", resp.Istat.Software)
```

Check the [example](../example/main.go) for more details.

To execute the above example against your local Apex controller, you can use below command. Replace `http://your-apex-hostname.local`  with the URL of your own Apex.

```
$ APEX_URL=http://your-apex-hostname.local \
    APEX_USERNAME=your_username \
    APEX_PASSWORD=your_password \
    go run ./example

software: "5.06_8C21"
serial: "AC5:12345"
inputs: 82
outputs: 89
session id: "AwmZi67ax3H8xXWUG9RCBZ95"
confirmed login as: "johndoe"
```

## API documentation

The available APIs are defined using [OAS 3.0](https://swagger.io/resources/open-api/) (OpenAPI Specification 3.0):

* [apis/public/v1](../apis/public/v1/openapi.yaml)
* [apis/private/v1alpha1](../apis/private/v1alpha1/openapi.yaml)

The lower level Go API clients which are wrapped by `go.reefassistant.com/apex.Client` are generated from their OAS 3.0 spec using the [OpenAPI Generator](https://openapi-generator.tech) project. Documentation for these lower level APIs can be found here:

* [client/public/v1](../client/public/v1/README.md)
* [client/private/v1alpha1](../client/private/v1alpha1/README.md)
