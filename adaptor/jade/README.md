# Using Adaptor to Work with [jade](https://github.com/Joker/jade)

This example demonstrates how to use adaptor to integrate [jade](https://github.com/Joker/jade) with Hertz as template engine.
### Server

```
package main

//go:generate jade -pkg=main -writer hello.jade

import (
	"fmt"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/adaptor"
)

func main() {
	h := server.Default()

	h.GET("/hello", adaptor.HertzHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// You may build more logic on req
		fmt.Println(r.URL.String())

		Jade_hello("Hertz", w)
	})))

	h.Spin()
}

```
## How to run
1. install jade by running `go install github.com/Joker/jade/cmd/jade@latest`. Please refer to [jade](https://github.com/Joker/jade) documentation should you require any help.
2. run `jade -writer -pkg=main adaptor/jade/hello.jade` to generate go code from template file.
3. run `go run hello/main.go`. This will spin up hertz listening on 8080.
4. run `curl --location --request GET '127.0.0.1:8080/hello'`