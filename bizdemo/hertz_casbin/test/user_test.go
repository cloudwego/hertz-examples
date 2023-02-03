package test

import (
	"encoding/json"
	"fmt"
	"github.com/darrenli6/hertz-examples/bizdemo/hertz_casbin/pkg/consts"
	"github.com/darrenli6/hertz-examples/bizdemo/hertz_casbin/pkg/utils"
	"testing"
)

func TestUserLogin(t *testing.T) {
	m := consts.M{
		"username": "admin",
		"password": "123",
	}

	data, _ := json.Marshal(m)
	rep, err := utils.HttpPost(userServiceAddr+"/v1/login", data)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(rep))
}

func TestRoleUserLogin(t *testing.T) {
	m := consts.M{
		"username": "role_user",
		"password": "123",
	}

	data, _ := json.Marshal(m)
	rep, err := utils.HttpPost(userServiceAddr+"/v1/login", data)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(rep))
}

func TestPermissionUserLogin(t *testing.T) {
	m := consts.M{
		"username": "permission_user",
		"password": "123",
	}

	data, _ := json.Marshal(m)
	rep, err := utils.HttpPost(userServiceAddr+"/v1/login", data)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(rep))
}
