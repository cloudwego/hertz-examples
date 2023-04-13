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
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
	h := server.Default(server.WithHostPorts("127.0.0.1:8080"))

	h.Static("/static", "./")

	h.StaticFile("/main", "./main.go")

	indexNames := []string{"1.txt", "2.txt"}
	h.StaticFS("/static1", &app.FS{
		Root:        "./",
		PathRewrite: app.NewPathSlashesStripper(1),
		PathNotFound: func(_ context.Context, ctx *app.RequestContext) {
			ctx.JSON(consts.StatusNotFound, "The requested resource does not exist")
		},

		// set the time interval for automatically closing inactive file handlers
		CacheDuration: time.Second * 5,

		// set indexNames, when you access the directory, you will get one of the file in the slice
		IndexNames: indexNames,

		// set the compress true, server adds a `CompressedFileSuffix` suffix to the original file name
		// server attempts to save the resulting compressed file under the new file name
		Compress:             true,
		CompressedFileSuffix: "hertz",

		// enables clients to request a specific range of bytes from a file on the server
		AcceptByteRange: true,
	})

	h.StaticFS("/static2", &app.FS{
		PathRewrite: app.NewPathSlashesStripper(1),

		// set GenerateIndexPages true, when you access the directory, you will get the index(without `IndexNames`)
		GenerateIndexPages: true,
	})

	h.Spin()
}
