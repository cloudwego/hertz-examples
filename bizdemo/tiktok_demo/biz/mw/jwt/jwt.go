/*
 * Copyright 2023 CloudWeGo Authors
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

package jwt

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/jwt"

	db "github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/biz/dal/db"
	"github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/biz/model/basic/user"
	"github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/pkg/errno"
	_ "github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/pkg/errno"
	"github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/pkg/utils"
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
	identity      = "user_id"
)

func Init() {
	JwtMiddleware, _ = jwt.New(&jwt.HertzJWTMiddleware{
		Key:         []byte("tiktok secret key"),
		TokenLookup: "query:token,form:token",
		Timeout:     24 * time.Hour,
		IdentityKey: identity,
		// Verify password at login
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginRequest user.DouyinUserLoginRequest
			if err := c.BindAndValidate(&loginRequest); err != nil {
				return nil, err
			}
			user, err := db.QueryUser(loginRequest.Username)
			if ok := utils.VerifyPassword(loginRequest.Password, user.Password); !ok {
				err = errno.PasswordIsNotVerified
				return nil, err
			}
			if err != nil {
				return nil, err
			}
			c.Set("user_id", user.ID)
			return user.ID, nil
		},
		// Set the payload in the token
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					identity: v,
				}
			}
			return jwt.MapClaims{}
		},
		// build login response if verify password successfully
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			hlog.CtxInfof(ctx, "Login success ，token is issued clientIP: "+c.ClientIP())
			c.Set("token", token)
		},
		// Verify token and get the id of logged-in user
		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			if v, ok := data.(float64); ok {
				current_user_id := int64(v)
				c.Set("current_user_id", current_user_id)
				hlog.CtxInfof(ctx, "Token is verified clientIP: "+c.ClientIP())
				return true
			}
			return false
		},
		// Validation failed, build the message
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(consts.StatusOK, user.DouyinUserLoginResponse{
				StatusCode: errno.AuthorizationFailedErrCode,
				StatusMsg:  message,
			})
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			resp := utils.BuildBaseResp(e)
			return resp.StatusMsg
		},
	})
}
