# discovery

1. start nacos
```shell
make prepare
```
2. start registry server
```shell
go run registry/main.go
```
3. start discovery server
```shell
go run registry/main.go
```
4. vist discovery server
```shell
curl --location --request GET 'http://127.0.0.1:8741/backend'
```
