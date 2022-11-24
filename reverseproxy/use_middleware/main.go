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

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/reverseproxy"
)

func main() {
	r := server.Default(server.WithHostPorts("127.0.0.1:9998"))

	r2 := server.Default(server.WithHostPorts("127.0.0.1:9997"))

	proxy, err := reverseproxy.NewSingleHostReverseProxy("http://127.0.0.1:9997")
	if err != nil {
		panic(err)
	}

	r.Use(func(c context.Context, ctx *app.RequestContext) {
		if ctx.Query("country") == "cn" {
			proxy.ServeHTTP(c, ctx)
			ctx.Response.Header.Set("key", "value")
			ctx.Abort()
		} else {
			ctx.Next(c)
		}
	})

	r.GET("/backend", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(200, utils.H{
			"message": "pong1",
		})
	})

	r2.GET("/backend", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(200, utils.H{
			"message": "pong2",
		})
	})

	go r.Spin()
	r2.Spin()
}
