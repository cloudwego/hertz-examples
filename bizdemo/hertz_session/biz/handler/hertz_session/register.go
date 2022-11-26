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
	"hertz-examples/bizdemo/hertz_session/biz/consts"
	"hertz-examples/bizdemo/hertz_session/biz/dal/mysql"
	"hertz-examples/bizdemo/hertz_session/biz/model"
	"hertz-examples/bizdemo/hertz_session/biz/utils"
)

type registerStruct struct {
	Username string `form:"username" json:"username" vd:"(len($) > 0 && len($) < 128); msg:'Illegal format'"`
	Email    string `form:"email" json:"email" vd:"(len($) > 0 && len($) < 128) && email($); msg:'Illegal format'"`
	Password string `form:"password" json:"password" vd:"(len($) > 0 && len($) < 128); msg:'Illegal format'"`
}

// Register user register handler
func Register(_ context.Context, c *app.RequestContext) {
	var err error
	var req registerStruct
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(http.StatusOK, hutils.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	users, err := mysql.FindUserByNameOrEmail(req.Username, req.Email)
	if err != nil {
		c.JSON(http.StatusOK, hutils.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	if len(users) != 0 {
		c.JSON(http.StatusOK, hutils.H{
			"code":    http.StatusBadRequest,
			"message": consts.RegisterErr,
		})
		return
	}
	if err = mysql.CreateUsers([]*model.User{
		{
			Username: req.Username,
			Password: utils.MD5(req.Password),
			Email:    req.Email,
		},
	}); err != nil {
		c.JSON(http.StatusOK, hutils.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, hutils.H{
		"code":    http.StatusOK,
		"message": consts.Success,
	})
}
