# resize-server
This is a HTTP-accessible server for scaling images on request. Think of it as [placeholder.com](https://placeholder.com) but for any image hosted on the internet.

## Table of contents

1. [How to use it?](#how-to-use-it)
2. [Configuration](#configuration)
3. [Misc. endpoints](#misc-endpoints)
4. [Build](#build)

## How to use it?
To use this server, you only need to send a GET request to the `/v1/scale/do` endpoint and specify these two query parameters :

| Key  | Value                                                                                                                                                       |
|------|-------------------------------------------------------------------------------------------------------------------------------------------------------------|
| path | The URL to the image                                                                                                                                        |
| size | Output size in pixels with the following format `600x300` OR a preset name. Set the height portion to 0 to keep the aspect ratio of the image i.e.: `200x0` |                                                                        

The scaled image will be directly output to the response. The response `Content-Type` header will be set the same as the source MIME type.

## Configuration
Configuration is done in the `./config/config.json` file, configuration is pretty basic.

| Key             | Value                                                                                                          | Default |
|-----------------|----------------------------------------------------------------------------------------------------------------|---------|
| http.port       | Port of the built-in HTTP server                                                                               | 80      |
| logger.logLevel | Minimum level of logs that output to the console, accept the following values : `info` `warn` `error` `debug`  | debug   |

### Presets
You can add presets in your config file. They can be used for the `size` query parameter when making a request.

| Key       | Type   | Description                                                                                                                   |
|-----------|--------|-------------------------------------------------------------------------------------------------------------------------------|
| name      | string | The name of your preset, this will be used in the `size` query parameter                                                      |
| width     | number | The output width in pixels                                                                                                    |
| height    | number | The output height in pixels                                                                                                   |
| interFunc | string | The interpolation function name. One of `NearestNeighbor`, `Bilinear`, `Bicubic`, `MitchellNetravali`, `Lanczos2`, `Lanczos3` |



## Misc. Endpoints

| Endpoint | Method | Description                                     |
|----------|--------|-------------------------------------------------|
| /healthz | GET    | Heartbeat function, just to tell you it's alive |
| /metrics | GET    | Metrics endpoint for Prometheus                 |

## Build
Building from source **requires Go 1.19 or later**

```commandline
go build -o ./app ./cmd/app
```
