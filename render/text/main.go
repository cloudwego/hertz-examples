package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
	h := server.Default(server.WithHostPorts(":8080"))

	h.GET("someText", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "message", "hello,world")
	})

	h.Spin()
}
