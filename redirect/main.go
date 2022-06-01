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
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
	h := server.Default(server.WithHostPorts("127.0.0.1:8080"))

	h.GET("/externalRedirect", func(ctx context.Context, c *app.RequestContext) {
		c.Redirect(consts.StatusMovedPermanently, []byte("http://www.google.com/"))
	})

	h.GET("/internalRedirect", func(ctx context.Context, c *app.RequestContext) {
		c.Redirect(consts.StatusFound, []byte("/foo"))
	})

	h.GET("/foo", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "hello, world")
	})

	h.Spin()
}
