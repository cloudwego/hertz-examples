# Websocket Reverse Proxy

This is an echo server example with websocket reverse proxy.

## Quick Start

```shell
go run main.go
```

## Example Usage

```text
2023/12/02 16:09:16.219292 engine.go:668: [Debug] HERTZ: Method=GET    absolutePath=/ws                       --> handlerName=github.com/hertz-contrib/reverseproxy.(*WSReverseProxy
).ServeHTTP-fm (num=2 handlers)
2023/12/02 16:09:16.227612 engine.go:396: [Info] HERTZ: Using network library=standard
2023/12/02 16:09:16.229601 transport.go:65: [Info] HERTZ: HTTP server listening on address=127.0.0.1:8080
2023/12/02 16:09:17.240885 engine.go:668: [Debug] HERTZ: Method=GET    absolutePath=/backend                  --> handlerName=main.main.func1.1 (num=2 handlers)
2023/12/02 16:09:17.240885 engine.go:396: [Info] HERTZ: Using network library=standard
2023/12/02 16:09:17.241777 transport.go:65: [Info] HERTZ: HTTP server listening on address=127.0.0.1:9090
2023/12/02 16:09:18.255589 ws_reverse_proxy.go:105: [Debug] upgrade handler working...
send: hello
receive: hello
send: world
receive: world
send: 
```
