# Hertz Examples

[English](README.md) | 中文

## 如何运行

您可以进入相关示例以获取有关“如何运行”的信息

## Bizdemo

- [hertz_gorm:](bizdemo/hertz_gorm) 在 hertz server 中使用 gorm 的示例
- [hertz_gorm_gen:](bizdemo/hertz_gorm_gen) 在 hertz server 中使用 gorm/gen & proto IDL 的示例
- [hertz_jwt:](bizdemo/hertz_jwt) 在 hertz server 中使用 jwt 的示例
- [hertz_session:](bizdemo/hertz_session) 在 hertz server 中使用分布式 session 和 csrf 的示例
- [hertz_swagger_gen:](bizdemo/hertz_swagger_gen) 在 hertz server 中使用插件生成 swagger 服务的示例
- [tiktok_demo:](bizdemo/tiktok_demo) 在 hertz server 中实现极简版抖音服务端的示例
- [formulago:](https://github.com/chenghonour/formulago) 使用 hertz 与 ent 实现的生产级别后台管理系统框架
- [gpress:](https://github.com/springrain/gpress) 使用 hertz 与 zorm 实现的生产级别云原生高性能内容平台

## Server

- [hello:](hello) 启动对于 hertz 来说相当于 "hello world" 的示例
- [config:](config) 配置 Hertz server 的示例
- [protocol:](protocol) 使用 http1, TLS 以及其他协议的示例
  - [HTTP1](https://github.com/cloudwego/hertz-examples/tree/main/protocol/http1) hertz 使用 HTTP1 协议的示例
  - [TLS](https://github.com/cloudwego/hertz-examples/tree/main/protocol/tls) hertz 使用 TLS 协议的示例
  - [HTTP2](https://github.com/hertz-contrib/http2/tree/main/examples) hertz 使用 HTTP2 协议的示例
  - [HTTP3](https://github.com/hertz-contrib/http3/tree/main/examples/quic-go) hertz 使用 HTTP3 协议的示例
  - [Websocket](https://github.com/hertz-contrib/websocket/tree/main/examples) hertz 使用 Websocket 协议的示例
  - [SSE](https://github.com/hertz-contrib/sse/tree/main/examples) hertz 使用 SSE 协议的示例
- [middleware:](middleware) 使用 hertz 中间件的示例
  - [basicauth:](middleware/basicauth) 使用 BasicAuth 中间件的示例
  - [cors:](middleware/CORS) 使用 CORS 中间件的示例
  - [csrf:](middleware/csrf) 使用 csrf 中间件示例
  - [custom:](middleware/custom) 自定义 middleware 的示例
  - [pprof:](middleware/pprof) 使用 pprof 中间件的示例
  - [requestid:](middleware/requestid) 使用 RequestID 中间件的示例
  - [gzip:](middleware/gzip) 使用 Gzip 中间件的示例
  - [loadbalance:](middleware/loadbalance) 使用 Loadbalance 中间件的示例
  - [Recovery](https://www.cloudwego.io/zh/docs/hertz/tutorials/basic-feature/middleware/recovery/) 使用 Recovery 中间件的示例
  - [jwt](https://github.com/hertz-contrib/jwt/tree/main/example/basic) 使用 jwt 中间件的示例
  - [i18n](https://github.com/hertz-contrib/i18n/tree/main/example) 使用 i18n 中间件的示例
  - [session](https://github.com/hertz-contrib/sessions/tree/main/_example) 使用 session 中间件的示例
  - [KeyAuth](https://github.com/hertz-contrib/keyauth/tree/main/example) 使用 KeyAuth 中间件的示例
  - [Swagger](https://github.com/hertz-contrib/swagger/tree/main/example/basic) 使用 Swagger 中间件的示例
  - [access log](https://github.com/hertz-contrib/logger/tree/main/accesslog/example) 使用 access log 中间件的示例
  - [Secure](https://github.com/hertz-contrib/secure/tree/main/example/custom) 使用 Secure 中间件的示例
  - [Sentry](https://github.com/hertz-contrib/hertzsentry) 使用 Sentry 中间件的示例
  - [Casbin](https://github.com/hertz-contrib/casbin/tree/main/example) 使用 Casbin 中间件的示例
  - [ETag](https://github.com/hertz-contrib/etag/tree/main/example) 使用 ETag 中间件的示例
  - [Cache](https://github.com/hertz-contrib/cache/tree/main/example) 使用 Cache 中间件的示例
  - [Paseto](https://github.com/hertz-contrib/paseto/tree/main/example) 使用 Paseto 中间件的示例
- [binding:](binding) 绑定参数和验证参数的示例
- [parameters:](parameter) 获取 query, form, cookie 等类型参数的示例
- [file:](file) 关于如何上传，下载文件和搭建静态文件服务的示例
- [render:](render) 渲染 json, html, protobuf 的示例
- [redirect:](redirect) 重定向到内部/外部URI的示例
- [streaming:](streaming) hertz server 的流读/写示例
- [graceful_shutdown:](graceful_shutdown) hertz server 如何优雅退出的示例
- [unit_test:](unit_test) 使用 hertz 提供的没有网络传输的接口编写单元测试的示例
- [tracer:](tracer) 使用 Jaeger 进行链接追踪的示例
- [monitoring:](monitoring) 使用 Prometheus 进行服务监控的示例
- [multiple_service:](multiple_service) 使用 Hertz 启动多端口服务的示例
- [adaptor:](adaptor) 使用 adaptor 集成基于`http.Handler`接口开发的工具, 包含使用 [jade](https://github.com/Joker/jade) 作为模版引擎的示例
- [sentinel:](sentinel) sentinel-golang 结合 hertz 使用的示例
- [reverseproxy:](reverseproxy) 在 hertz server 中使用反向代理的示例
- [hlog:](hlog) 使用 hlog 以及其日志拓展的示例
- [graphql-go:](graphql-go) 在 hertz server 中使用 graphql 的示例


## Client

- [client/send_request:](client/send_request) 使用 hertz 客户端发送http请求的示例
- [client/config:](client/config) 配置 hertz 客户端的示例
- [protocol/tls:](protocol/tls) 使用 hertz 客户端发送 TLS 请求的示例
- [client/add_parameters:](client/add_parameters) 使用 hertz 客户端添加请求参数的示例
- [client/upload_file:](client/upload_file) 使用 hertz 客户端上载文件的示例
- [client/middleware:](client/middleware) 使用 hertz 客户端中间件的示例
- [client/streaming_read:](client/streaming_read) 使用 hertz 客户端的流式读取响应示例
- [client/forward_proxy:](client/forward_proxy) 使用 hertz 客户端配置转发代理的示例

## Hz

- [hz/thrift:](hz/thrift) 使用 hz 与 thrift 生成服务端代码的示例
- [hz/protobuf:](hz/protobuf) 使用 hz 与 protobuf 生成服务端代码的示例
- [hz/hz_client:](hz/hz_client) 使用 hz 生成客户端代码的示例
- [hz/template:](hz/template) 使用 hz 自定义模版生成服务端代码的示例
- [hz/plugin:](hz/plugin) 使用 hz 接入第三方插件的示例
- [hz/struct_reuse:](hz/struct_reuse) 使用 hz 与 kitex_gen 作为 hertz model 的代码示例

## Note

执行示例的所有命令都应在 hertz-examples 下执行。
