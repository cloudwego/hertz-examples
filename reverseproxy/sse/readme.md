## sse

### Introduction

This example shows how to use sse with reverseproxy.

### Usage

* Start server

```bash
go run main.go
```

* Test
```bash
curl -N --location http://localhost:8080/proxy/sse
```

Result:
```bash
> curl -N --location http://localhost:8080/proxy/sse
event:timestamp
data:2023-10-09T22:32:14+08:00

event:timestamp
data:2023-10-09T22:32:14+08:00

event:timestamp
data:2023-10-09T22:32:14+08:00

event:timestamp
data:2023-10-09T22:32:14+08:00

event:timestamp
data:2023-10-09T22:32:14+08:00

event:timestamp
data:2023-10-09T22:32:15+08:00

event:timestamp
data:2023-10-09T22:32:15+08:00

event:timestamp
data:2023-10-09T22:32:15+08:00

event:timestamp
data:2023-10-09T22:32:15+08:00

event:timestamp
data:2023-10-09T22:32:15+08:00
```
