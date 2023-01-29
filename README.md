# Hertz Examples

English | [中文](README_CN.md)

## How to run

You can enter the example for information about  "How to run"

## Bizdemo

- [bizdemo/hertz_gorm:](bizdemo/hertz_gorm) Example of using gorm in hertz server
- [bizdemo/hertz_gorm_gen:](bizdemo/hertz_gorm_gen) Example of using gorm/gen & proto IDL in hertz server
- [bizdemo/hertz_jwt:](bizdemo/hertz_jwt) Example of using jwt in hertz server
- [bizdemo/hertz_session:](bizdemo/hertz_session) Example of using distributed session and csrf in hertz server

## Server

- [hello:](hello) Example of launching a hertz "hello world" application
- [config:](config) Example of configuring hertz server
- [protocol:](protocol) Example of using http1, tls and other protocols of hertz
- [middleware:](middleware) Example of using middleware of hertz
  - [basicauth:](middleware/basicauth) Example of using BasicAuth middleware
  - [cors:](middleware/CORS) Example of using CORS middleware
  - [csrf:](middleware/csrf) Example of using csrf middleware
  - [custom:](middleware/custom) Example of using custom middleware
  - [pprof:](middleware/pprof) Example of using pprof middleware
  - [requestid:](middleware/requestid) Example of using RequestID middleware
  - [gzip:](middleware/gzip) Example of using Gzip middleware
- [binding:](binding) Example of parameter binding and validation
- [parameters:](parameter) Example of getting query, form, cookie
- [file:](file) Examples of file upload, file download, and static file services
- [render:](render) Example of render body as json, html, protobuf
- [redirect:](redirect)  Examples of redirects to internal/external URI
- [streaming:](streaming) Example of streaming read/write for hertz server
- [graceful_shutdown:](graceful_shutdown) Example of a graceful shutdown for hertz server
- [unit_test:](unit_test) Example of writing unit tests using the interface provided by hertz without network transmission
- [tracer:](tracer) Example of using Jaeger for link tracing
- [monitoring:](monitoring) Example of using Prometheus for metrics monitoring
- [multiple_service:](multiple_service) Example of using hertz with multiple services
- [adaptor:](adaptor) Example of using adaptor to integrate hertz with package built for `http.Handler` interface , including a demonstration on using [jade](https://github.com/Joker/jade)
  as template engine.
- [sentinel:](sentinel) Example of using sentinel-golang in hertz
- [reverseproxy:](reverseproxy/standard) Example of using reverseproxy in hertz server
- [gzip:](middleware/gzip) Example of using gzip middleware in hertz server
- [hlog:](hlog) Example of using hlog and its log extension

## Client

- [client/send_request:](client/send_request) Example of sending http requests using hertz client
- [client/config:](client/config) Example of configuring hertz client
- [protocol/tls:](protocol/tls) Example of sending a tls request using hertz client
- [client/add_parameters:](client/add_parameters) Example of adding request parameters using the hertz client
- [client/upload_file:](client/upload_file) Example of uploading a file using hertz client
- [client/middleware:](client/middleware) Example of using hertz client middleware
- [client/streaming_read:](client/streaming_read) Example of streaming read response using hertz client
- [client/forward_proxy:](client/forward_proxy) Example of configuring a forward proxy using hertz client

## Hz

- [hz/thrift:](hz/thrift) Example of using hz with thrift to generate server code
- [hz/protobuf:](hz/protobuf) Example of using hz with protobuf to generate server code
- [hz/hz_client:](hz/hz_client) Example of using hz to generate client code
- [hz/template:](hz/template) Example of using hz custom templates to generate server code
- [hz/plugin:](hz/plugin) Example of using hz to access third-party plugins

## Note

All commands to execute the example should be executed under "hertz-example".
