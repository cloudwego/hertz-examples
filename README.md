# Hertz Examples

English | [中文](README_CN.md)

## How to run

You can enter the example for information about  "How to run"

## Bizdemo

- [bizdemo/hertz_gorm:](bizdemo/hertz_gorm) Example of using gorm in hertz server
- [bizdemo/hertz_jwt:](bizdemo/hertz_jwt) Example of using jwt in hertz server
- [bizdemo/hertz_session:](bizdemo/hertz_session) Example of using distributed session and csrf in hertz server
- [bizdemo/hertz_swagger_gen:](bizdemo/hertz_swagger_gen) Example of using plugin to generate swagger service in hertz server
- [bizdemo/tiktok_demo:](bizdemo/tiktok_demo) Example of simple tiktok in hertz server
- [formulago:](https://github.com/chenghonour/formulago) Production-level backend management system framework implemented using hertz and ent
- [gpress:](https://github.com/springrain/gpress) Production-grade cloud-native high-performance content platform using hertz and zorm

## Server

- [hello:](hello) Example of launching a hertz "hello world" application
- [config:](config) Example of configuring hertz server
- [protocol:](protocol) Example of using http1, tls and other protocols of hertz
  - [HTTP1](https://github.com/cloudwego/hertz-examples/tree/main/protocol/http1) Example of hertz using HTTP1 protocol
  - [TLS](https://github.com/cloudwego/hertz-examples/tree/main/protocol/tls) Example of hertz using TLS protocol
  - [HTTP2](https://github.com/hertz-contrib/http2/tree/main/examples) Example of hertz using HTTP2 protocol
  - [HTTP3](https://github.com/hertz-contrib/http3/tree/main/examples/quic-go) Example of hertz using HTTP3 protocol
  - [Websocket](https://github.com/hertz-contrib/websocket/tree/main/examples) Example of hertz using Websocket protocol
  - [SSE](https://github.com/hertz-contrib/sse/tree/main/examples) Example of hertz using SSE protocol
- [middleware:](middleware) Example of using middleware of hertz
  - [basicauth:](middleware/basicauth) Example of using BasicAuth middleware
  - [cors:](middleware/CORS) Example of using CORS middleware
  - [csrf:](middleware/csrf) Example of using csrf middleware
  - [custom:](middleware/custom) Example of using custom middleware
  - [pprof:](middleware/pprof) Example of using pprof middleware
  - [requestid:](middleware/requestid) Example of using RequestID middleware
  - [gzip:](middleware/gzip) Example of using Gzip middleware
  - [loadbalance:](middleware/loadbalance) Example of using Loadbalance middleware
  - [Recovery](https://www.cloudwego.io/docs/hertz/tutorials/basic-feature/middleware/recovery/) Example of using Recovery middleware
  - [jwt](https://github.com/hertz-contrib/jwt/tree/main/example/basic) Example of using jwt middleware
  - [i18n](https://github.com/hertz-contrib/i18n/tree/main/example) Example of using i18n middleware
  - [session](https://github.com/hertz-contrib/sessions/tree/main/_example) Example of using session middleware
  - [KeyAuth](https://github.com/hertz-contrib/keyauth/tree/main/example) Example of using KeyAuth middleware
  - [Swagger](https://github.com/hertz-contrib/swagger/tree/main/example/basic) Example of using Swagger middleware
  - [access log](https://github.com/hertz-contrib/logger/tree/main/accesslog/example) Example of using access log middleware
  - [Secure](https://github.com/hertz-contrib/secure/tree/main/example/custom) Example of using Secure middleware
  - [Sentry](https://github.com/hertz-contrib/hertzsentry) Example of using Sentry middleware
  - [Casbin](https://github.com/hertz-contrib/casbin/tree/main/example) Example of using Casbin middleware
  - [ETag](https://github.com/hertz-contrib/etag/tree/main/example) Example of using ETag middleware
  - [Cache](https://github.com/hertz-contrib/cache/tree/main/example) Example of using Cache middleware
  - [Paseto](https://github.com/hertz-contrib/paseto/tree/main/example) Example of using Paseto middleware
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
- [trailer:](trailer) Example of read/write trailers for hertz server
- [graphql-go:](graphql-go) Example of using graphql in hertz server


## Client

- [client/send_request:](client/send_request) Example of sending http requests using hertz client
- [client/config:](client/config) Example of configuring hertz client
- [protocol/tls:](protocol/tls) Example of sending a tls request using hertz client
- [client/add_parameters:](client/add_parameters) Example of adding request parameters using the hertz client
- [client/upload_file:](client/upload_file) Example of uploading a file using hertz client
- [client/middleware:](client/middleware) Example of using hertz client middleware
- [client/streaming_read:](client/streaming_read) Example of streaming read response using hertz client
- [client/forward_proxy:](client/forward_proxy) Example of configuring a forward proxy using hertz client
- [trailer:](trailer) Example of sending a request with trailer using hertz client

## Hz

- [hz/thrift:](hz/thrift) Example of using hz with thrift to generate server code
- [hz/protobuf:](hz/protobuf) Example of using hz with protobuf to generate server code
- [hz/hz_client:](hz/hz_client) Example of using hz to generate client code
- [hz/template:](hz/template) Example of using hz custom templates to generate server code
- [hz/plugin:](hz/plugin) Example of using hz to access third-party plugins
- [hz/struct_reuse:](hz/struct_reuse) Example of using hz to use kitex_gen as hertz model

## Note

All commands to execute the example should be executed under "hertz-example".
