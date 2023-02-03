package test

import (
	"encoding/json"
	"fmt"
	"github.com/darrenli6/hertz-examples/bizdemo/hertz_casbin/pkg/consts"
	"github.com/darrenli6/hertz-examples/bizdemo/hertz_casbin/pkg/utils"
	"testing"
)

func TestPermissionAdd(t *testing.T) {

	header, _ = json.Marshal(rolem1)
	m := consts.M{
		"v1": "/v1/permission/create/",
		"v2": "POST",
	}

	data, _ := json.Marshal(m)
	rep, err := utils.HttpPost(userServiceAddr+"/v1/permission/create", data, header...)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(rep))
}

func TestPermissionBind(t *testing.T) {
	m := consts.M{
		"pid": "1",
		"rid": "1",
	}

	data, _ := json.Marshal(m)
	rep, err := utils.HttpPost(userServiceAddr+"/v1/permissionrole/bind/", data, header...)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(rep))
}
