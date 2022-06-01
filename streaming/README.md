# Streaming
You can learn about how to use streaming read/write using hertz server：
* server-read: how to stream read request's body using hertz
* server-write: how to stream write request's body using hertz, and providing a demo about how to use chunk for streaming write body


For more information about streaming, please click [streaming](https://www.cloudwego.io/zh/docs/hertz/tutorials/basic-feature/stream/)
## How to run
* start streaming write server  
  `go run streaming/streaming_write/main.go`
* send client streaming read request  
  `go run client/streaming_read/main.go`

