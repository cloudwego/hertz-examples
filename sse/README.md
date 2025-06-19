# Hertz SSE 示例

这个示例展示了如何使用 Hertz 框架提供 Server-Sent Events (SSE) 服务，以及如何使用 Hertz 客户端接收 SSE 事件。

## 项目结构

```
/sse
  /server - Hertz SSE 服务器示例
  /client - Hertz SSE 客户端示例
```

## 服务器端 (Hertz)

服务器端使用 Hertz 框架的 SSE 包实现了一个简单的 SSE 服务。它会每秒发送一个事件，总共发送 10 个事件。

主要功能：

- 使用 `sse.NewWriter` 创建 SSE 写入器
- 使用 `writer.WriteEvent` 发送带有 ID、类型和数据的事件
- 支持 `Last-Event-ID` 头部，可用于恢复连接

## 客户端 (Hertz)

客户端使用 Hertz 框架的客户端 API 和 SSE 包连接到 SSE 服务器并接收事件。

主要功能：

- 使用 `client.NewClient()` 创建 Hertz 客户端
- 使用 `sse.AddAcceptMIME()` 添加 SSE 相关的请求头
- 使用 `sse.NewReader()` 创建 SSE 读取器
- 使用 `reader.ForEach()` 迭代处理 SSE 事件
- 支持通过 Ctrl+C 优雅退出

## 运行示例

1. 启动服务器：

```bash
cd server
go run main.go
```

2. 在另一个终端启动客户端：

```bash
cd client
go run main.go
```

## SSE 协议简介

Server-Sent Events (SSE) 是一种服务器推送技术，允许服务器通过 HTTP 连接向客户端推送实时更新。与 WebSocket 不同，SSE 是单向的（只能从服务器到客户端），并且基于标准 HTTP 协议。

SSE 事件格式：

```
id: event-id
event: event-type
data: event-data

```

每个事件由一个或多个字段组成，字段之间用换行符分隔，事件之间用空行分隔。

## 参考资料

- [Hertz 框架 SSE 包文档](https://pkg.go.dev/github.com/cloudwego/hertz/pkg/protocol/sse)
- [MDN Server-Sent Events 指南](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events)