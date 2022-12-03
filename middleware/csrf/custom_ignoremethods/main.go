package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/csrf"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/cookie"
)

func main() {
	h := server.Default()

	store := cookie.NewStore([]byte("secret"))
	h.Use(sessions.Sessions("session", store))
	h.Use(csrf.New(csrf.WithIgnoredMethods([]string{"GET", "HEAD", "TRACE"})))

	h.GET("/protected", func(c context.Context, ctx *app.RequestContext) {
		ctx.String(200, csrf.GetToken(ctx))
	})

	h.OPTIONS("/protected", func(c context.Context, ctx *app.RequestContext) {
		ctx.String(200, "success")
	})
	h.Spin()
}
