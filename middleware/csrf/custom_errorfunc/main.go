/*
 * Copyright 2022 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
	if err == nil {
		return
	}

	switch err.Err.(type) {
	case error:
		switch {
		case errors.Is(err, errMissingForm), errors.Is(err, errMissingParam), errors.Is(err, errMissingHeader), errors.Is(err, errMissingQuery):
			ctx.String(http.StatusBadRequest, err.Error()) // extract csrf-token failed
		case errors.Is(err, errMissingSalt):
			fmt.Println(err.Error())
			ctx.String(http.StatusInternalServerError, err.Error()) // get salt failed, which is unexpected
		case errors.Is(err, errInvalidToken):
			ctx.String(http.StatusBadRequest, err.Error()) // csrf-token is invalid
		default:
			ctx.String(http.StatusInternalServerError, "Unknown error") // handle unknown errors
		}
	}

	ctx.Abort()
}

func main() {
	h := server.Default()

	store := cookie.NewStore([]byte("store"))
	h.Use(sessions.New("csrf-session", store))
	h.Use(csrf.New(csrf.WithErrorFunc(myErrFunc)))

	h.GET("/protected", func(c context.Context, ctx *app.RequestContext) {
		ctx.String(200, csrf.GetToken(ctx))
	})
	h.POST("/protected", func(c context.Context, ctx *app.RequestContext) {
		ctx.String(200, "CSRF token is valid")
	})

	h.Spin()
}
