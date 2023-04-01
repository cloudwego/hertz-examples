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
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.Default(server.WithHostPorts("127.0.0.1:8080"))

	h.Static("/static", "./")

	h.StaticFile("/main", "./main.go")

	h.StaticFS("/static1", &app.FS{
		PathRewrite: app.NewPathSlashesStripper(1),
		PathNotFound: func(_ context.Context, ctx *app.RequestContext) {
			ctx.JSON(consts.StatusNotFound, "The requested resource does not exist")
		},
	})

	h.Spin()
}
