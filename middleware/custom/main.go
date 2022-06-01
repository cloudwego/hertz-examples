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
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func MyMiddleware1() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// pre-handle
		fmt.Println("pre-handle")
	}
}

func MyMiddleware2() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// pre-handle
		fmt.Println("pre-handle")

		c.Next(ctx) // call the next middleware(handler)
		// post-handle
		fmt.Println("post-handle")
	}
}

func main() {
	h := server.Default(server.WithHostPorts("127.0.0.1:8080"))

	h.Use(MyMiddleware1())
	h.Use(MyMiddleware2())

	h.GET("/middleware", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "Hello hertz!")
	})

	h.Spin()
}
