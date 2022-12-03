package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/csrf"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/cookie"
)

func myErrFunc(_ context.Context, ctx *app.RequestContext) {
	if ctx.Errors.Last() == nil {
		fmt.Errorf("myErrFunc called when no error occurs")
	}
	ctx.AbortWithMsg(ctx.Errors.Last().Error(), http.StatusBadRequest)
}

func main() {
	h := server.Default()

	store := cookie.NewStore([]byte("store"))
	h.Use(sessions.Sessions("csrf-session", store))
	h.Use(csrf.New(csrf.WithErrorFunc(myErrFunc)))

	h.GET("/protected", func(c context.Context, ctx *app.RequestContext) {
		ctx.String(200, csrf.GetToken(ctx))
	})
	h.POST("/protected", func(c context.Context, ctx *app.RequestContext) {
		ctx.String(200, "CSRF token is valid")
	})

	h.Spin()
}
