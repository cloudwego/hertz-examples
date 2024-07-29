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

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/csrf"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/cookie"
)

func myExtractor(_ context.Context, ctx *app.RequestContext) (string, error) {
	token := ctx.FormValue("csrf-token")
	if token == nil {
		return "", errors.New("missing token in form-data") // get csrf-token from form-data failed
	}
	return string(token), nil
}

func main() {
	h := server.Default()

	store := cookie.NewStore([]byte("secret"))
	h.Use(sessions.New("csrf-session", store))
	h.Use(csrf.New(csrf.WithExtractor(myExtractor)))

	h.GET("/protected", func(ctx context.Context, c *app.RequestContext) {
		ctx.String(200, csrf.GetToken(ctx))
	})
	h.POST("/protected", func(ctx context.Context, c *app.RequestContext) {
		ctx.String(200, "CSRF token is valid")
	})

	h.Spin()
}
