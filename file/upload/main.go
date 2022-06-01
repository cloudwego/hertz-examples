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

func main() {
	// WithMaxRequestBodySize can set the size of the body
	h := server.Default(server.WithHostPorts("127.0.0.1:8080"), server.WithMaxRequestBodySize(20<<20))

	h.POST("/singleFile", func(ctx context.Context, c *app.RequestContext) {
		// single file
		file, _ := c.FormFile("file")
		fmt.Println(file.Filename)

		// Upload the file to specific dst
		c.SaveUploadedFile(file, fmt.Sprintf("./file/upload/%s", file.Filename))

		c.String(consts.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	h.POST("/multiFile", func(ctx context.Context, c *app.RequestContext) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["file"]

		for _, file := range files {
			fmt.Println(file.Filename)

			// Upload the file to specific dst.
			c.SaveUploadedFile(file, fmt.Sprintf("./file/upload/%s", file.Filename))
		}
		c.String(consts.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})

	h.Spin()
}
