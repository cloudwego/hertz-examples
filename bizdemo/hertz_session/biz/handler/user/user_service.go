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

// Code generated by hertz generator.

package user

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz-examples/bizdemo/hertz_session/biz/dal/mysql"
	"github.com/cloudwego/hertz-examples/bizdemo/hertz_session/biz/model/user"
	"github.com/cloudwego/hertz-examples/bizdemo/hertz_session/pkg/consts"
	"github.com/cloudwego/hertz-examples/bizdemo/hertz_session/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"
	hutils "github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/sessions"
)

// Register .
// @router /register [POST]
func Register(_ context.Context, c *app.RequestContext) {
	var err error
	var req user.RegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.HTML(http.StatusOK, "register.html", hutils.H{
			"message": utils.BuildMsg(err.Error()),
		})
		return
	}
	users, err := mysql.FindUserByNameOrEmail(req.Username, req.Email)
	if err != nil {
		c.HTML(http.StatusOK, "register.html", hutils.H{
			"message": utils.BuildMsg(err.Error()),
		})
		return
	}
	if len(users) != 0 {
		c.HTML(http.StatusOK, "register.html", hutils.H{
			"message": utils.BuildMsg(consts.RegisterErr),
		})
		return
	}
	if err = mysql.CreateUsers([]*mysql.User{
		{
			Username: req.Username,
			Password: utils.MD5(req.Password),
			Email:    req.Email,
		},
	}); err != nil {
		c.HTML(http.StatusOK, "register.html", hutils.H{
			"message": utils.BuildMsg(consts.RegisterErr),
		})
		return
	}
	c.HTML(http.StatusOK, "register.html", hutils.H{
		"message": consts.Success,
	})
}

// Login .
// @router /login [POST]
func Login(_ context.Context, c *app.RequestContext) {
	var err error
	var req user.LoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.HTML(http.StatusOK, "login.html", hutils.H{
			"message": utils.BuildMsg(err.Error()),
		})
		return
	}
	users, err := mysql.CheckUser(req.Username, utils.MD5(req.Password))
	if err != nil {
		c.HTML(http.StatusOK, "login.html", hutils.H{
			"message": utils.BuildMsg(err.Error()),
		})
		return
	}
	if len(users) == 0 {
		c.HTML(http.StatusOK, "login.html", hutils.H{
			"message": utils.BuildMsg(consts.LoginErr),
		})
		return
	}
	session := sessions.Default(c)
	session.Set(consts.Username, req.Username)
	_ = session.Save()
	c.Redirect(http.StatusMovedPermanently, []byte("/page"))
}

// Logout .
// @router /logout [GET]
func Logout(_ context.Context, c *app.RequestContext) {
	session := sessions.Default(c)
	session.Delete(consts.Username)
	_ = session.Save()
	c.Redirect(http.StatusMovedPermanently, []byte("/login.html"))
}
