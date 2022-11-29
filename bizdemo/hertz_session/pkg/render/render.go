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

package render

import (
	"context"
	"html/template"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	hutils "github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/sessions"
	"hertz-examples/bizdemo/hertz_session/pkg/consts"
	"hertz-examples/bizdemo/hertz_session/pkg/utils"
)

func InitHTML(h *server.Hertz) {
	h.Delims("{[{", "}]}")
	h.SetFuncMap(template.FuncMap{
		"BuildMsg": utils.BuildMsg,
	})
	h.LoadHTMLGlob("html/*")
	// register.html
	h.GET("/register", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(http.StatusOK, "register.html", hutils.H{
			"message": utils.BuildMsg("Please Sign Up"),
		})
	})
	// login.html
	h.GET("/login", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(http.StatusOK, "login.html", hutils.H{
			"message": utils.BuildMsg("Please Login"),
		})
	})
	// page.html
	h.GET("/page", func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		username := session.Get(consts.Username)
		if username == nil {
			c.HTML(http.StatusOK, "page.html", hutils.H{
				"message": utils.BuildMsg(consts.PageErr),
			})
			c.Redirect(http.StatusMovedPermanently, []byte("/login"))
			return
		}
		c.HTML(http.StatusOK, "page.html", hutils.H{
			"message": utils.BuildMsg(username.(string)),
		})
	})
}