package main

import (
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/csrf"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/cookie"
)

func myExtractor(_ context.Context, ctx *app.RequestContext) (string, error) {
	token := ctx.FormValue("csrf-token")
	if token == nil {
		return "", errors.New("missing token in form-data")
	}
	return string(token), nil
}

func main() {
	h := server.Default()

	store := cookie.NewStore([]byte("secret"))
	h.Use(sessions.Sessions("csrf-session", store))
	h.Use(csrf.New(csrf.WithExtractor(myExtractor)))

	h.GET("/protected", func(c context.Context, ctx *app.RequestContext) {
		ctx.String(200, csrf.GetToken(ctx))
	})
	h.POST("/protected", func(c context.Context, ctx *app.RequestContext) {
		ctx.String(200, "CSRF token is valid")
	})

	h.Spin()
}
