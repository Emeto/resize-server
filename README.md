# resize-server
This is a HTTP-accessible server for scaling images on request. Think of it as [placeholder.com](https://placeholder.com) but for any image hosted on the internet.

## Table of contents

1. [How to use it?](#how-to-use-it)
2. [Configuration](#configuration)
3. [Misc. endpoints](#misc-endpoints)
4. [Build](#build)

## How to use it?
To use this server, you only need to send a GET request to the `/v1/scale/do` endpoint and specify these two query parameters :

| Key  | Value                                                                                                                                      |
|------|--------------------------------------------------------------------------------------------------------------------------------------------|
| path | The URL to the image                                                                                                                       |
| size | Output size in pixels with the following format `600x300`. Set the height portion to 0 to keep the aspect ratio of the image i.e.: `200x0` |                                                                        
The scaled image will be directly output to the response. The response `Content-Type` header will be set the same as the source MIME type.

## Configuration
Configuration is done using environment variables, configuration is pretty basic, **if the defaults are good enough for your use case, setting them is not required**.

| Key       | Value                                                                                                          | Default |
|-----------|----------------------------------------------------------------------------------------------------------------|---------|
| HTTP_PORT | Port of the built-in HTTP server                                                                               | 80      |
| LOG_LEVEL | Minimum level of logs that output to the console, accept the following values : `info` `warn` `error` `debug`  | debug   |

Definition of these variables depends on how the server is deployed :

### Pre-built binaries
Create an `.env` file in the same directory as the binary and paste this in :
```dotenv
HTTP_PORT=80
LOG_LEVEL=debug
```

### Docker
Specify the variables when running `docker run` or set them in the `docker-compose.yml` file from the source

## Misc. endpoints

| Endpoint | Method | Description                                     |
|----------|--------|-------------------------------------------------|
| /healthz | GET    | Heartbeat function, just to tell you it's alive |
| /metrics | GET    | Metrics endpoint for Prometheus                 |

## Build
Building from source **requires Go 1.19 or later**

```commandline
go build -o ./app ./cmd/app
```
