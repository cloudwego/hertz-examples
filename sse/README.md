# Hertz SSE Example

This example demonstrates how to use the Hertz framework to provide Server-Sent Events (SSE) service and how to use the Hertz client to receive SSE events.

## Project Structure

```
/sse
  /server - Hertz SSE server example
  /client - Hertz SSE client example
```

## Server Side (Hertz)

The server side implements a simple SSE service using the Hertz framework's SSE package. It sends one event per second, for a total of 10 events.

Main features:

- Use `sse.NewWriter` to create SSE writer
- Use `writer.WriteEvent` to send events with ID, type, and data
- Support `Last-Event-ID` header for connection recovery

## Client Side (Hertz)

The client uses Hertz framework's client API and SSE package to connect to the SSE server and receive events.

Main features:

- Use `client.NewClient()` to create Hertz client
- Use `sse.AddAcceptMIME()` to add SSE-related request headers
- Use `sse.NewReader()` to create SSE reader
- Use `reader.ForEach()` to iterate over SSE events
- Support graceful exit through Ctrl+C

## Running the Example

1. Start the server:

```bash
cd server
go run main.go
```

2. Start the client in another terminal:

```bash
cd client
go run main.go
```

## SSE Protocol Introduction

Server-Sent Events (SSE) is a server push technology that allows servers to push real-time updates to clients over HTTP connections. Unlike WebSocket, SSE is unidirectional (server to client only) and based on standard HTTP protocol.

SSE event format:

```
id: event-id
event: event-type
data: event-data

```

Each event consists of one or more fields, separated by newlines, with events separated by blank lines.

## References

- [Hertz Framework SSE Package Documentation](https://pkg.go.dev/github.com/cloudwego/hertz/pkg/protocol/sse)
- [MDN Server-Sent Events Guide](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events)