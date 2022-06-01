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
)

type Args struct {
	Query      string   `query:"query"`
	QuerySlice []string `query:"q"`
	Path       string   `path:"path"`
	Header     string   `header:"header"`
	Form       string   `form:"form"`
	Json       string   `json:"json"`
	Vd         int      `query:"vd" vd:"$==0||$==1"`
}

func main() {
	h := server.Default(server.WithHostPorts("127.0.0.1:8080"))

	h.POST("v:path/bind", func(ctx context.Context, c *app.RequestContext) {
		var arg Args
		err := c.BindAndValidate(&arg)
		if err != nil {
			panic(err)
		}
		fmt.Println(arg)
	})

	h.Spin()
}
