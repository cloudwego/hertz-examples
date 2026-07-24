# Hertz Examples

English | [中文](README_CN.md)

## How to run

Enter an example directory for its specific run instructions. Unless an example says otherwise, run commands from the repository root.

## Bizdemo

- [bizdemo/hertz_gorm](bizdemo/hertz_gorm): Use Thrift IDL, `hz`, Hertz binding and validation, GORM, and MySQL in a Hertz service.
- [bizdemo/hertz_jwt](bizdemo/hertz_jwt): Build a Hertz service with JWT authentication, GORM, and MySQL.
- [bizdemo/hertz_session](bizdemo/hertz_session): Use Redis-backed distributed sessions and CSRF protection in a Hertz service.
- [bizdemo/hertz_swagger_gen](bizdemo/hertz_swagger_gen): Generate Swagger documentation and a Swagger UI service from Thrift IDL with `thrift-gen-http-swagger`.
- [bizdemo/tiktok_demo](bizdemo/tiktok_demo): A simple Hertz backend for user, video, social, feed, favorite, comment, and message services.
- [formulago](https://github.com/chenghonour/formulago): An enterprise administration framework built with Hertz and Ent.
- [gpress](https://github.com/springrain/gpress): A Web3 content platform built with Hertz and Go templates, with FTS5 full-text search and Hugo compatibility.

## Server

- [hello](hello): A basic Hertz “hello world” server.
- [config](config): Configure a Hertz server.
- [protocol](protocol): Use HTTP/1.1 and TLS with Hertz. It also links to external protocol examples.
  - [HTTP1](https://github.com/cloudwego/hertz-examples/tree/main/protocol/http1): A Hertz HTTP/1.1 server.
  - [TLS](https://github.com/cloudwego/hertz-examples/tree/main/protocol/tls): A TLS server plus a Hertz client that sends a TLS request.
  - [HTTP2](https://github.com/hertz-contrib/http2/tree/main/examples): The external Hertz HTTP/2 examples.
  - [HTTP3](https://github.com/hertz-contrib/http3/tree/main/examples/quic-go): The external Hertz HTTP/3 example based on QUIC.
  - [WebSocket](https://github.com/hertz-contrib/websocket/tree/main/examples): The external Hertz WebSocket examples.
  - [SSE](sse): A Hertz Server-Sent Events server and client using Hertz’s built-in SSE package.
- [middleware](middleware): Server middleware examples, including BasicAuth, CORS, CSRF, custom middleware, pprof, RequestID, Gzip, and load balancing.
  - [basicauth](middleware/basicauth): Use BasicAuth middleware.
  - [CORS](middleware/CORS): Use CORS middleware.
  - [csrf](middleware/csrf): Configure CSRF middleware, including custom token extraction and skip rules.
  - [custom](middleware/custom): Write and register custom middleware.
  - [pprof](middleware/pprof): Configure pprof middleware.
  - [requestid](middleware/requestid): Configure RequestID middleware and custom ID handling.
  - [gzip](middleware/gzip): Configure Gzip compression and route/path exclusions.
  - [loadbalance](middleware/loadbalance): Configure a non-default load-balancing algorithm.
  - [Recovery](https://www.cloudwego.io/docs/hertz/tutorials/basic-feature/middleware/recovery/): Use the built-in Recovery middleware.
  - [JWT](https://github.com/hertz-contrib/jwt/blob/main/example/basic/main.go): Use JWT authentication middleware.
  - [i18n](https://github.com/hertz-contrib/i18n/blob/main/example/main.go): Use internationalization middleware.
  - [session](https://github.com/hertz-contrib/sessions): Use session middleware with cookie or Redis backends.
  - [KeyAuth](https://github.com/hertz-contrib/keyauth): Authenticate requests with a configurable key lookup.
  - [Swagger](https://github.com/hertz-contrib/swagger/blob/main/example/basic/main.go): Serve Swagger documentation with Hertz.
  - [access log](https://github.com/hertz-contrib/logger/blob/main/accesslog/example/main.go): Log HTTP access details.
  - [Secure](https://github.com/hertz-contrib/secure/blob/main/example/custom/main.go): Configure security-related HTTP headers and redirects.
  - [Sentry](https://github.com/hertz-contrib/hertzsentry): Report Hertz request errors to Sentry.
  - [Casbin](https://github.com/hertz-contrib/casbin/blob/main/example/main.go): Apply Casbin authorization middleware.
  - [ETag](https://github.com/hertz-contrib/etag): Add ETag response headers and customize ETag behavior.
  - [Cache](https://github.com/hertz-contrib/cache): Cache Hertz responses with memory or Redis backends.
  - [Paseto](https://github.com/hertz-contrib/paseto): Protect routes with PASETO tokens.
- [binding](binding): Bind request parameters and validate them.
- [parameters](parameter): Read query, form, and cookie parameters.
- [file](file): Upload and download files, serve static files, and render HTML from files or templates.
- [render](render): Render JSON, HTML, Protobuf, text, XML, and customized YAML responses.
- [redirect](redirect): Redirect requests to internal or external URIs.
- [route](route): Register static, grouped, parameterized, and middleware-protected routes, and inspect route metadata.
- [streaming](streaming): Stream request and response bodies from a Hertz server.
- [graceful_shutdown](graceful_shutdown): Gracefully shut down a Hertz server.
- [unit_test](unit_test): Test Hertz handlers with `ResponseRecord` and `PerformRequest` without network transmission.
- [monitoring](monitoring): Export Hertz and Go runtime metrics to Prometheus and visualize them in Grafana.
- [opentelemetry](opentelemetry): Trace a Hertz-to-Kitex request, export telemetry through OpenTelemetry Collector, and view traces and metrics in Jaeger and Grafana/VictoriaMetrics.
- [multiple_service](multiple_service): Run multiple Hertz services on different ports.
- [adaptor](adaptor): Adapt standard-library `http.Handler` implementations for use in Hertz, including a [Jade](https://github.com/Joker/jade) template example.
- [sentinel](sentinel): Use Sentinel for traffic control with Hertz.
- [reverseproxy](reverseproxy): Use the Hertz reverse proxy extension, with examples for standard proxying, TLS, discovery, response modification, SSE, middleware, and WebSocket.
- [hlog](hlog): Use Hertz logging and its logger extensions.
- [trailer](trailer): Read and write HTTP trailers with Hertz.
- [graphql-go](graphql-go): Serve GraphQL requests from a Hertz server.

## Client

- [client/send_request](client/send_request): Send requests with the Hertz client, including deadlines, redirects, and timeouts.
- [client/config](client/config): Configure a Hertz client.
- [protocol/tls](protocol/tls): Send a TLS request with the Hertz client to the example TLS server.
- [client/add_parameters](client/add_parameters): Add query, form, path, and body parameters to Hertz client requests.
- [client/upload_file](client/upload_file): Upload files with the Hertz client.
- [client/middleware](client/middleware): Use client middleware.
- [client/streaming_read](client/streaming_read): Read a streaming response with the Hertz client.
- [client/forward_proxy](client/forward_proxy): Configure a forward proxy for the Hertz client.
- [trailer](trailer): Send a request and read its HTTP trailer with the Hertz client.

## Hz

- [hz/thrift](hz/thrift): Use `hz` with Thrift IDL to generate server code.
- [hz/protobuf](hz/protobuf): Use `hz` with Protobuf IDL to generate server code.
- [hz/hz_client](hz/hz_client): Use `hz` to generate client code.
- [hz/template](hz/template): Use a custom `hz` template to generate a server project.
- [hz/plugin](hz/plugin): Use third-party `hz` code-generation plugins with Thrift and Protobuf projects.
- [hz/struct_reuse](hz/struct_reuse): Reuse Kitex-generated Thrift structs as Hertz models.

## Note

Commands for an example should be run from the `hertz-examples` repository root unless that example’s own README says otherwise.
