package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/csrf"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/cookie"
)

func isPostMethod(_ context.Context, ctx *app.RequestContext) bool {
	if string(ctx.Method()) == "POST" {
		return true
	} else {
		return false
	}
}

func main() {
	h := server.Default()

	store := cookie.NewStore([]byte("store"))
	h.Use(sessions.Sessions("csrf-session", store))

	//  skip csrf middleware when request method is post
	h.Use(csrf.New(csrf.WithNext(isPostMethod)))

	h.POST("/protected", func(c context.Context, ctx *app.RequestContext) {
		ctx.String(200, "success even no csrf-token in header")
	})
	h.Spin()
}
