# Using Adaptor

This example demonstrates how to use get standard `http.ResponseWriter` and `http.Request` from `app.RequestContext` . This is useful if you want to integrate packages build for `net/http` with Hertz, especially these use `Handler` to register handler. Be mindful that this compatibility comes at the cost performance loss. 

```
package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
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

	h.GET("/hello", func(ctx context.Context, c *app.RequestContext) {
		req, err := adaptor.GetCompatRequest(&c.Request)
		if err != nil {
			fmt.Println(err)
			return
		}
		// You may build more logic on req
		fmt.Println(req.URL.String())

		// caution: don't pass in c.GetResponse() as it return a copy of response
		rw := adaptor.GetCompatResponseWriter(&c.Response)

		handler(rw, req)
	})

	h.Spin()
}

```
## How to run
1. run `go run hello/main.go`. This will spin up hertz listening on 8080.
2. run `curl --location --request GET '127.0.0.1:8080/hello'`