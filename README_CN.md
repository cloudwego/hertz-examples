# Hertz Examples

[English](README.md) | 中文

## 如何运行

请进入对应示例目录查看运行说明。除非示例另有说明，所有命令都应在仓库根目录执行。

## 业务示例

- [bizdemo/hertz_gorm](bizdemo/hertz_gorm)：使用 Thrift IDL、`hz`、Hertz 参数绑定与校验、GORM 和 MySQL 构建 Hertz 服务。
- [bizdemo/hertz_jwt](bizdemo/hertz_jwt)：在 Hertz 服务中使用 JWT 认证、GORM 和 MySQL。
- [bizdemo/hertz_session](bizdemo/hertz_session)：在 Hertz 服务中使用基于 Redis 的分布式 Session 和 CSRF 防护。
- [bizdemo/hertz_swagger_gen](bizdemo/hertz_swagger_gen)：使用 `thrift-gen-http-swagger` 从 Thrift IDL 生成 Swagger 文档和 Swagger UI 服务。
- [bizdemo/tiktok_demo](bizdemo/tiktok_demo)：包含用户、视频、社交、Feed、点赞、评论和消息服务的简易 Hertz 后端。
- [formulago](https://github.com/chenghonour/formulago)：使用 Hertz 和 Ent 构建的企业级后台管理框架。
- [gpress](https://github.com/springrain/gpress)：使用 Hertz 和 Go 模板构建的 Web3 内容平台，支持 FTS5 全文搜索和 Hugo 生态。

## 服务端

- [hello](hello)：基础的 Hertz “hello world” 服务。
- [config](config)：配置 Hertz 服务端。
- [protocol](protocol)：在 Hertz 中使用 HTTP/1.1 和 TLS，并链接到其他协议示例。
  - [HTTP1](https://github.com/cloudwego/hertz-examples/tree/main/protocol/http1)：Hertz HTTP/1.1 服务示例。
  - [TLS](https://github.com/cloudwego/hertz-examples/tree/main/protocol/tls)：TLS 服务端，以及使用 Hertz 客户端发送 TLS 请求的示例。
  - [HTTP2](https://github.com/hertz-contrib/http2/tree/main/examples)：外部 Hertz HTTP/2 示例。
  - [HTTP3](https://github.com/hertz-contrib/http3/tree/main/examples/quic-go)：基于 QUIC 的外部 Hertz HTTP/3 示例。
  - [WebSocket](https://github.com/hertz-contrib/websocket/tree/main/examples)：外部 Hertz WebSocket 示例。
  - [SSE](sse)：使用 Hertz 内置 SSE 包实现的服务端和客户端示例。
- [middleware](middleware)：服务端中间件示例，包括 BasicAuth、CORS、CSRF、自定义中间件、pprof、RequestID、Gzip 和负载均衡。
  - [basicauth](middleware/basicauth)：使用 BasicAuth 中间件。
  - [CORS](middleware/CORS)：使用 CORS 中间件。
  - [csrf](middleware/csrf)：配置 CSRF 中间件，包括自定义 token 提取和跳过规则。
  - [custom](middleware/custom)：编写并注册自定义中间件。
  - [pprof](middleware/pprof)：配置 pprof 中间件。
  - [requestid](middleware/requestid)：配置 RequestID 中间件和自定义 ID 处理逻辑。
  - [gzip](middleware/gzip)：配置 Gzip 压缩以及路由/路径排除规则。
  - [loadbalance](middleware/loadbalance)：配置非默认的负载均衡算法。
  - [Recovery](https://www.cloudwego.io/zh/docs/hertz/tutorials/basic-feature/middleware/recovery/)：使用内置 Recovery 中间件。
  - [JWT](https://github.com/hertz-contrib/jwt/blob/main/example/basic/main.go)：使用 JWT 认证中间件。
  - [i18n](https://github.com/hertz-contrib/i18n/blob/main/example/main.go)：使用国际化中间件。
  - [session](https://github.com/hertz-contrib/sessions)：使用支持 Cookie 或 Redis 后端的 Session 中间件。
  - [KeyAuth](https://github.com/hertz-contrib/keyauth)：使用可配置的 key 查找方式认证请求。
  - [Swagger](https://github.com/hertz-contrib/swagger/blob/main/example/basic/main.go)：在 Hertz 中提供 Swagger 文档。
  - [access log](https://github.com/hertz-contrib/logger/blob/main/accesslog/example/main.go)：记录 HTTP 访问详情。
  - [Secure](https://github.com/hertz-contrib/secure/blob/main/example/custom/main.go)：配置安全相关的 HTTP Header 和重定向。
  - [Sentry](https://github.com/hertz-contrib/hertzsentry)：将 Hertz 请求错误上报到 Sentry。
  - [Casbin](https://github.com/hertz-contrib/casbin/blob/main/example/main.go)：使用 Casbin 授权中间件。
  - [ETag](https://github.com/hertz-contrib/etag)：添加 ETag 响应 Header 并自定义 ETag 行为。
  - [Cache](https://github.com/hertz-contrib/cache)：使用内存或 Redis 后端缓存 Hertz 响应。
  - [Paseto](https://github.com/hertz-contrib/paseto)：使用 PASETO token 保护路由。
- [binding](binding)：绑定并校验请求参数。
- [parameters](parameter)：读取 query、form 和 cookie 参数。
- [file](file)：上传、下载文件，提供静态文件服务，并从文件或模板渲染 HTML。
- [render](render)：渲染 JSON、HTML、Protobuf、文本、XML 和自定义 YAML 响应。
- [redirect](redirect)：将请求重定向到内部或外部 URI。
- [route](route)：注册静态路由、路由组、参数路由和带中间件的路由，并查看路由信息。
- [streaming](streaming)：在 Hertz 服务端流式读取请求体和写入响应体。
- [graceful_shutdown](graceful_shutdown)：优雅退出 Hertz 服务。
- [unit_test](unit_test)：使用 `ResponseRecord` 和 `PerformRequest` 在没有网络传输的情况下测试 Hertz Handler。
- [monitoring](monitoring)：将 Hertz 和 Go 运行时指标导出到 Prometheus，并使用 Grafana 展示。
- [opentelemetry](opentelemetry)：对 Hertz 到 Kitex 的请求进行链路追踪，通过 OpenTelemetry Collector 导出数据，并在 Jaeger 与 Grafana/VictoriaMetrics 中查看链路和指标。
- [multiple_service](multiple_service)：在不同端口运行多个 Hertz 服务。
- [adaptor](adaptor)：将标准库 `http.Handler` 实现适配到 Hertz 中，并包含 [Jade](https://github.com/Joker/jade) 模板示例。
- [sentinel](sentinel)：在 Hertz 中使用 Sentinel 进行流量控制。
- [reverseproxy](reverseproxy)：使用 Hertz 反向代理扩展，包含标准代理、TLS、服务发现、响应修改、SSE、中间件和 WebSocket 示例。
- [hlog](hlog)：使用 Hertz 日志及其日志扩展。
- [trailer](trailer)：使用 Hertz 读取和写入 HTTP trailer。
- [graphql-go](graphql-go)：在 Hertz 服务端处理 GraphQL 请求。

## 客户端

- [client/send_request](client/send_request)：使用 Hertz 客户端发送请求，包括 deadline、重定向和超时。
- [client/config](client/config)：配置 Hertz 客户端。
- [protocol/tls](protocol/tls)：使用 Hertz 客户端向示例 TLS 服务端发送 TLS 请求。
- [client/add_parameters](client/add_parameters)：为 Hertz 客户端请求添加 query、form、path 和 body 参数。
- [client/upload_file](client/upload_file)：使用 Hertz 客户端上传文件。
- [client/middleware](client/middleware)：使用客户端中间件。
- [client/streaming_read](client/streaming_read)：使用 Hertz 客户端读取流式响应。
- [client/forward_proxy](client/forward_proxy)：为 Hertz 客户端配置正向代理。
- [trailer](trailer)：发送请求并使用 Hertz 客户端读取 HTTP trailer。

## Hz

- [hz/thrift](hz/thrift)：使用 `hz` 和 Thrift IDL 生成服务端代码。
- [hz/protobuf](hz/protobuf)：使用 `hz` 和 Protobuf IDL 生成服务端代码。
- [hz/hz_client](hz/hz_client)：使用 `hz` 生成客户端代码。
- [hz/template](hz/template)：使用自定义 `hz` 模板生成服务端项目。
- [hz/plugin](hz/plugin)：在 Thrift 和 Protobuf 项目中使用第三方 `hz` 代码生成插件。
- [hz/struct_reuse](hz/struct_reuse)：复用 Kitex 生成的 Thrift 结构体作为 Hertz model。

## 注意

除非示例自己的 README 另有说明，执行示例命令时请以 `hertz-examples` 仓库根目录为当前目录。
