# Using Adaptor

This example demonstrates how to wrap a standard http.Handler so that it can run inside Hertz. This is useful if you want to integrate packages built for `net/http` with Hertz, especially for libraries that expose handlers via the `http.Handler` interface. Be mindful that this compatability comes at a cost of some performance loss.

```
package main

import (
	"fmt"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/adaptor"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)

	_, err := w.Write([]byte("Hello World"))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	h := server.Default()

	// Wrap the standard http.HandlerFunc using HertzHandler
	h.GET("/hello", adaptor.HertzHandler(http.HandlerFunc(handler)))

	h.Spin()
}

```
## How to run
1. run `go run hello/main.go`. This will spin up hertz listening on 8080.
2. run `curl --location --request GET '127.0.0.1:8080/hello'`