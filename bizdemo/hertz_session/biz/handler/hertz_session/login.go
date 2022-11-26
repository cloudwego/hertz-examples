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

package hertz_session

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	hutils "github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/sessions"
	"hertz-examples/bizdemo/hertz_session/biz/consts"
	"hertz-examples/bizdemo/hertz_session/biz/dal/mysql"
	"hertz-examples/bizdemo/hertz_session/biz/utils"
)

type loginStruct struct {
	Username string `form:"username" json:"username" vd:"(len($) > 0 && len($) < 30); msg:'Illegal format'"`
	Password string `form:"password" json:"password" vd:"(len($) > 0 && len($) < 30); msg:'Illegal format'"`
}

// Login user login handler
func Login(_ context.Context, c *app.RequestContext) {
	var err error
	var req loginStruct
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(http.StatusOK, hutils.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	users, err := mysql.CheckUser(req.Username, utils.MD5(req.Password))
	if err != nil {
		c.JSON(http.StatusOK, hutils.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	if len(users) == 0 {
		c.JSON(http.StatusOK, hutils.H{
			"code":    http.StatusBadRequest,
			"message": consts.LoginErr,
		})
		return
	}
	session := sessions.Default(c)
	session.Set(consts.Username, req.Username)
	_ = session.Save()
	c.Redirect(http.StatusMovedPermanently, []byte("/page"))
}
