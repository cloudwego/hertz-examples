package test

import (
	"encoding/json"
	"fmt"
	"github.com/darrenli6/hertz-examples/bizdemo/hertz_casbin/pkg/consts"
	"github.com/darrenli6/hertz-examples/bizdemo/hertz_casbin/pkg/utils"
	"testing"
)

func TestRoleAdd(t *testing.T) {
 
	header, _ = json.Marshal(rolem1)

	m := consts.M{
		"name": "admin",
	}

	data, _ := json.Marshal(m)
	rep, err := utils.HttpPost(userServiceAddr+"/v1/role/create", data, header...)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(rep))
}
