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

	"github.com/hertz-contrib/csrf"

	"github.com/cloudwego/hertz-examples/bizdemo/hertz_session/pkg/consts"
	"github.com/cloudwego/hertz-examples/bizdemo/hertz_session/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	hutils "github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/sessions"
)

// InitHTML render HTML page
func InitHTML(h *server.Hertz) {
	h.Delims("{[{", "}]}")
	h.SetFuncMap(template.FuncMap{
		"BuildMsg": utils.BuildMsg,
	})
	// load templates
	h.LoadHTMLGlob("static/html/*")
	h.Static("/", "./static")
	token := ""
	// register.html
	h.GET("/register.html", func(ctx context.Context, c *app.RequestContext) {
		if !utils.IsLogout(ctx, c) {
			token = csrf.GetToken(c)
		}
		c.HTML(http.StatusOK, "register.html", hutils.H{
			"message": utils.BuildMsg("Register a new membership"),
			"token":   utils.BuildMsg(token),
		})
	})
	// login.html
	h.GET("/login.html", func(ctx context.Context, c *app.RequestContext) {
		if !utils.IsLogout(ctx, c) {
			token = csrf.GetToken(c)
		}
		c.HTML(http.StatusOK, "login.html", hutils.H{
			"message": utils.BuildMsg("Sign in to start your session"),
			"token":   utils.BuildMsg(token),
		})
	})
	// index.html
	h.GET("/index.html", func(ctx context.Context, c *app.RequestContext) {
		if !utils.IsLogout(ctx, c) {
			token = csrf.GetToken(c)
		}
		session := sessions.Default(c)
		username := session.Get(consts.Username)
		if username == nil {
			c.HTML(http.StatusOK, "index.html", hutils.H{
				"message": utils.BuildMsg(consts.PageErr),
				"token":   utils.BuildMsg(token),
			})
			c.Redirect(http.StatusMovedPermanently, []byte("/login.html"))
			return
		}
		c.HTML(http.StatusOK, "index.html", hutils.H{
			"message": utils.BuildMsg(username.(string)),
			"token":   utils.BuildMsg(token),
		})
	})
}
