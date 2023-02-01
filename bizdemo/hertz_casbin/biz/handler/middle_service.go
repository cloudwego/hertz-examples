package handler

import (
	"context"
	"github.com/cloudwego/hertz-examples/bizdemo/hertz_session/biz/dal/mysql"
	"github.com/cloudwego/hertz-examples/bizdemo/hertz_session/biz/mw"
	"github.com/cloudwego/hertz-examples/bizdemo/hertz_session/pkg/consts"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/sessions"
)

func Auth(_ context.Context, c *app.RequestContext) map[string]interface{} {

	session := sessions.Default(c)

	role := session.Get(consts.Role)
	if role == "" || role == nil {
		role = "anonymous"
	}

	if c.FullPath() == "/login" || c.FullPath() == "/register" {
		role = "anonymous"
	}

	// if it's a member, check if the user still exists
	if role == "member" {
		username := session.Get(consts.Username)

		_, err := mysql.CheckUserExists(username.(string))

		if err != nil {

			return utils.H{"code": 1, "message": "FORBIDDEN"}
		}
	}

	// casbin enforce
	res, err := mw.Authorize(role, c.FullPath(), string(c.Request.Header.Method()))
	if err != nil {
		return utils.H{"code": 1, "message": "ERROR"}
	}

	if res {
		return utils.H{"code": 0, "message": "ok"}
	} else {

		return utils.H{"code": 1, "message": "FORBIDDEN"}
	}
}
