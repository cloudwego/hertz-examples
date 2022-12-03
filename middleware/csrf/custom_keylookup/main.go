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

	store := cookie.NewStore([]byte("store"))
	h.Use(sessions.Sessions("csrf-session", store))
	h.Use(csrf.New(csrf.WithKeyLookUp("form:csrf")))

	h.GET("/protected", func(c context.Context, ctx *app.RequestContext) {
		ctx.String(200, csrf.GetToken(ctx))
	})
	h.POST("/protected", func(c context.Context, ctx *app.RequestContext) {
		ctx.String(200, "CSRF token is valid")
	})

	h.Spin()
}
