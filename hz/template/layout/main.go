// This is my customized hertz layout.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz/hello/biz/router"
)

func main() {
	h := server.Default()

	router.GeneratedRegister(h)

	// do what you wanted
	// add some render data: this is customized render data
	h.Spin()
}
