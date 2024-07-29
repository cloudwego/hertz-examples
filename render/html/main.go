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
	"html/template"
	"net/http"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

func main() {
	h := server.Default(server.WithHostPorts(":8080"))

	h.Delims("{[{", "}]}")
	h.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	h.LoadHTMLGlob("render/html/*")

	h.GET("/index", func(ctx context.Context, c *app.RequestContext) {
		ctx.HTML(http.StatusOK, "index.tmpl", utils.H{
			"title": "Main website",
		})
	})

	h.GET("/raw", func(ctx context.Context, c *app.RequestContext) {
		ctx.HTML(http.StatusOK, "template1.html", map[string]interface{}{
			"now": time.Date(2017, 0o7, 0o1, 0, 0, 0, 0, time.UTC),
		})
	})
	h.Spin()
}
