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

	"github.com/cloudwego/hertz-examples/bizdemo/hertz_gen/biz/model/api/user"
	"github.com/cloudwego/hertz-examples/bizdemo/hertz_gen/biz/model/orm_gen"
	"github.com/cloudwego/hertz-examples/bizdemo/hertz_gen/biz/model/pack"
	"github.com/cloudwego/hertz-examples/bizdemo/hertz_gen/biz/model/query"
	"github.com/cloudwego/hertz/pkg/app"
)

// CreateUserResponse .
// @router /v1/user/create [POST]
func CreateUserResponse(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.CreateUserReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := new(user.CreateUserResp)
	err = query.User.WithContext(ctx).Create(&orm_gen.User{
		Name:      req.Name,
		Gender:    int32(req.Gender),
		Age:       int32(req.Age),
		Introduce: req.Introduce,
	})
	if err != nil {
		resp.Code = user.Code_ParamInvalid
		resp.Msg = err.Error()
		c.JSON(200, resp)
		return
	}

	resp.Code = user.Code_Success
	resp.Msg = "创建记录成功"
	c.JSON(200, resp)
}

// QueryUserResponse .
// @router /v1/user/query [POST]
func QueryUserResponse(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.QueryUserReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	resp := new(user.QueryUserResp)
	u, m := query.User, query.User.WithContext(ctx)
	if req.Keyword != "" {
		m = m.Where(u.Introduce.Like("%" + req.Keyword + "%"))
	}

	var total int64
	total, err = m.Count()
	if err != nil {
		resp.Code = user.Code_DBErr
		resp.Msg = err.Error()
		c.JSON(200, resp)
		return
	}

	var users []*orm_gen.User
	if total > 0 {
		users, err = m.Limit(int(req.PageSize)).Offset(int(req.PageSize * (req.Page - 1))).Find()
		if err != nil {
			resp.Code = user.Code_DBErr
			resp.Msg = err.Error()
			c.JSON(200, resp)
			return
		}
	}

	resp.Code = user.Code_Success
	resp.Total = total
	resp.User = pack.Users(users)
	c.JSON(200, resp)
}

// UpdateUserResponse .
// @router /v1/user/update/:user_id [POST]
func UpdateUserResponse(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UpdateUserReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := new(user.UpdateUserResp)
	u := &orm_gen.User{}
	u.ID = req.UserID
	u.Name = req.Name
	u.Gender = int32(req.Gender)
	u.Age = int32(req.Age)
	u.Introduce = req.Introduce
	_, err = query.User.WithContext(ctx).Updates(u)
	if err != nil {
		resp.Code = user.Code_DBErr
		resp.Msg = err.Error()
		c.JSON(200, resp)
		return
	}

	resp.Code = user.Code_Success
	resp.Msg = "更新记录成功"
	c.JSON(200, resp)
}

// DeleteUserResponse .
// @router /v1/user/delete/:user_id [POST]
func DeleteUserResponse(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.DeleteUserReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := new(user.DeleteUserResp)
	_, err = query.User.WithContext(ctx).Where(query.User.ID.Eq(req.UserID)).Delete()
	if err != nil {
		resp.Code = user.Code_DBErr
		resp.Msg = err.Error()
		c.JSON(200, resp)
		return
	}

	resp.Code = user.Code_Success
	resp.Msg = "删除记录成功"
	c.JSON(200, resp)
}
