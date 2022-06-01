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
	h := server.Default(server.WithHostPorts("127.0.0.1:8080"))

	// content-type : application/x-www-form-urlencoded
	h.POST("/urlencoded", func(ctx context.Context, c *app.RequestContext) {
		name := c.PostForm("name")
		message := c.PostForm("message")

		c.PostArgs().VisitAll(func(key, value []byte) {
			if string(key) == "name" {
				fmt.Printf("This is %s!", string(value))
			}
		})

		c.String(consts.StatusOK, "name: %s; message: %s", name, message)
	})

	// content-type : multipart/form-data
	h.POST("/formdata", func(ctx context.Context, c *app.RequestContext) {
		id := c.FormValue("id")
		name := c.FormValue("name")
		message := c.FormValue("message")

		c.String(consts.StatusOK, "id: %s; name: %s; message: %s\n", id, name, message)
	})

	h.Spin()
}
