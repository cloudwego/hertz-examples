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
	"net/url"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.Default(server.WithHostPorts("127.0.0.1:8080"))

	// FileAttachment() sets the "content-disposition" header and returns the file as an "attachment".
	h.GET("/fileAttachment", func(ctx context.Context, c *app.RequestContext) {
		// If you use Chinese, need to encode
		fileName := url.QueryEscape("hertz")
		c.FileAttachment("./file/download/file.txt", fileName)
	})

	// File() will return the contents of the file directly
	h.GET("/file", func(ctx context.Context, c *app.RequestContext) {
		c.File("./file/download/file.txt")
	})

	h.Spin()
}
