# Hertz Examples
[English](README.md) | 中文
## 如何运行
您可以进入相关示例以获取有关“如何运行”的信息
## Server
- hello: 启动对于 hertz 来说相当于 "hello world" 的示例
- config: 配置 Hertz server 的示例
- protocol: 使用 http1, TLS 以及其他协议的示例
- middleware: 使用 hertz 中间件的示例
- binding: 绑定参数和验证参数的示例
- parameters: 获取 query, form, cookie 等类型参数的示例
- file: 关于如何上传，下载文件和搭建静态文件服务的示例
- render: 渲染 json, html, protobuf 的示例
- redirect: 重定向到内部/外部URI的示例
- streaming: hertz server 的流读/写示例
- graceful_shutdown: hertz server 如何优雅退出的示例
- unit_test: 使用 hertz 提供的没有网络传输的接口编写单元测试的示例
- tracer: 使用 Jaeger 进行链接追踪的示例
- monitoring: 使用 Prometheus 进行服务监控的示例

## Client
- client/send_request: 使用 hertz 客户端发送http请求的示例
- client/config: 配置 hertz 客户端的示例
- protocol/tls: 使用 hertz 客户端发送 TLS 请求的示例
- client/add_parameters: 使用 hertz 客户端添加请求参数的示例
- client/upload_file: 使用 hertz 客户端上载文件的示例
- client/middleware: 使用 hertz 客户端中间件的示例
- client/streaming_read: 使用 hertz 客户端的流式读取响应示例
- client/forward_proxy: 使用 hertz 客户端配置转发代理的示例

## Note
执行示例的所有命令都应在 hertz-examples 下执行。