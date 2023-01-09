package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/csrf"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/cookie"
)

var (
	errMissingHeader = errors.New("[CSRF] missing csrf token in header")
	errMissingQuery  = errors.New("[CSRF] missing csrf token in query")
	errMissingParam  = errors.New("[CSRF] missing csrf token in param")
	errMissingForm   = errors.New("[CSRF] missing csrf token in form")
	errMissingSalt   = errors.New("[CSRF] missing salt")
	errInvalidToken  = errors.New("[CSRF] invalid token")
)

// myErrFunc is executed when an error occurs in csrf middleware.
func myErrFunc(_ context.Context, ctx *app.RequestContext) {
	err := ctx.Errors.Last()
	switch err {
	case errMissingForm, errMissingParam, errMissingHeader, errMissingQuery:
		ctx.String(http.StatusBadRequest, err.Error()) // extract csrf-token failed
	case errMissingSalt:
		fmt.Println(err.Error())
		ctx.String(http.StatusInternalServerError, err.Error()) // get salt failed,which is unexpected
	case errInvalidToken:
		ctx.String(http.StatusBadRequest, err.Error()) //csrf-token is invalid
	}
	ctx.Abort()
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
