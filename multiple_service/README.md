# multiple service

This is "multiple service" example.

multiple service is a multi-port service.

Users can request different results from the same path, different ports. 

eg: `curl --location --request GET '127.0.0.1:8080/ping'`or `curl --location --request GET '127.0.0.1:8081/ping'`
- The response result for 8080 is `pong1` and for 8081 is `pong2`.

## how to run
* `go run multiple_service/main.go`

* `curl --location --request GET '127.0.0.1:8080/ping'`or `curl --location --request GET '127.0.0.1:8081/ping'`
