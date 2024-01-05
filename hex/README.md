# CWGO Hex Usage

## Introduce
The main power of `cwgo hex` is to allow hertz and kitex to listen on the same port and use protocol sniffing to distribute requests to kitex and hertz for processing
## Install
```
# Go 1.15 and earlier version
GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get github.com/cloudwego/cwgo@latest

# Go 1.16 and later version
GOPROXY=https://goproxy.cn/,direct go install github.com/cloudwego/cwgo@latest
```
## Usage
- init go.mod `go mod init cwgo/example/hex`
- generate code `cwgo  server --type RPC  --idl idl/hello.thrift --service p.s.m --hex`
- mod tidy `go mod tidy`

## Test
- `cd /cwgo/example/hex`
- start server: `go run .`
- test rpc: `go run client/main.go`
  - `HelloResp({RespBody:[KITEX] hello, hex})`
- test http: `curl 127.0.0.1:8888/hello?name=hex`
  - `{"RespBody":"[HERTZ] hello, hex"}`