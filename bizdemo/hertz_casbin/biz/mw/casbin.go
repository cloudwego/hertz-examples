package mw

import (
	"github.com/casbin/casbin"
	xormadapter "github.com/casbin/xorm-adapter"
	"github.com/darrenli6/hertz-examples/bizdemo/hertz_casbin/pkg/consts"
)

var AuthEnforcer *casbin.Enforcer

func InitCasbin() {
	adapter := xormadapter.NewAdapter("mysql", consts.MysqlDSN, true)

	enforcer := casbin.NewEnforcer("conf/auth_model.conf", adapter)

	AuthEnforcer = enforcer
}

func Authorize(rvals ...interface{}) (result bool, err error) {
	// casbin enforce
	res, err1 := AuthEnforcer.EnforceSafe(rvals[0], rvals[1], rvals[2])
	return res, err1
}
