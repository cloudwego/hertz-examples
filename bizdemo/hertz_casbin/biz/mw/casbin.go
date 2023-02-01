package mw

import (
	"github.com/casbin/casbin"
	"log"
)

var authEnforcer *casbin.Enforcer

func InitCasbin() {
	var err error
	authEnforcer, err = casbin.NewEnforcerSafe("./conf/auth_model.conf", "./conf/policy.csv")
	if err != nil {
		log.Fatal(err)
	}
}

func Authorize(rvals ...interface{}) (result bool, err error) {
	// casbin enforce
	res, err1 := authEnforcer.EnforceSafe(rvals[0], rvals[1], rvals[2])
	return res, err1
}
