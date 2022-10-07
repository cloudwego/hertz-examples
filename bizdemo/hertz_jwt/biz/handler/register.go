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

package handler

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz-examples/bizdemo/hertz_jwt/biz/dal/mysql"
	"github.com/cloudwego/hertz-examples/bizdemo/hertz_jwt/biz/model"
	utils2 "github.com/cloudwego/hertz-examples/bizdemo/hertz_jwt/biz/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

// Register user register handler
func Register(ctx context.Context, c *app.RequestContext) {
	var registerStruct struct {
		Username string `form:"username" json:"username" query:"username" vd:"(len($) > 0 && len($) < 128); msg:'Illegal format'"`
		Email    string `form:"email" json:"email" query:"email" vd:"(len($) > 0 && len($) < 128) && email($); msg:'Illegal format'"`
		Password string `form:"password" json:"password" query:"password" vd:"(len($) > 0 && len($) < 128); msg:'Illegal format'"`
	}

	if err := c.BindAndValidate(&registerStruct); err != nil {
		c.JSON(http.StatusOK, utils.H{
			"message": err.Error(),
			"code":    http.StatusBadRequest,
		})
		return
	}
	users, err := mysql.FindUserByNameOrEmail(registerStruct.Username, registerStruct.Email)
	if err != nil {
		c.JSON(http.StatusOK, utils.H{
			"message": err.Error(),
			"code":    http.StatusBadRequest,
		})
		return
	}

	if len(users) != 0 {
		c.JSON(http.StatusOK, utils.H{
			"message": "user already exists",
			"code":    http.StatusBadRequest,
		})
		return
	}

	if err = mysql.CreateUsers([]*model.User{
		{
			UserName: registerStruct.Username,
			Email:    registerStruct.Email,
			Password: utils2.MD5(registerStruct.Password),
		},
	}); err != nil {
		c.JSON(http.StatusOK, utils.H{
			"message": err.Error(),
			"code":    http.StatusBadRequest,
		})
		return
	}

	c.JSON(http.StatusOK, utils.H{
		"message": "success",
		"code":    http.StatusOK,
	})
}
