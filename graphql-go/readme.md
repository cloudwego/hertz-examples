## graphql-go

### Introduction

This example shows how to use graphql-go in hertz server.

### Usage

* Start server
```shell
go run main.go
```

* Test
- Use curl to send a request

```shell
curl -X POST 'http://localhost:9090/' --header 'content-type: application/json' --data-raw '{"query":"query{hello}"}'
```
Result:
```shell
> curl -X POST 'http://localhost:9090/' --header 'content-type: application/json' --data-raw '{"query":"query{hello}"}'
{"data":{"hello":"world"}}
``` 

```shell
curl 'http://localhost:9090/?query=query%7Bhello%7D'
```

Result:
```shell
> curl 'http://localhost:9090/?query=query%7Bhello%7D'
{"data":{"hello":"world"}}
```
- Use Graphql Playground to send a request to http://localhost:9090/

```graphql
query {
  hello
}
```

```json
{
  "data": {
    "hello": "world"
  }
}
```

### Warning

We do not recommend using [gqlgen](https://github.com/99designs/gqlgen) in hertz. Because adapting the Go standard library `http.Handler` and `http.HandlerFunc` methods results in a performance hit.
