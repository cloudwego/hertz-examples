package mysql

import (
	"fmt"
	"github.com/darrenli6/hertz-examples/bizdemo/hertz_casbin/biz/model/casbin"
	"github.com/darrenli6/hertz-examples/bizdemo/hertz_casbin/pkg/utils"
	"testing"
)

func init() {
	Init()
}

func TestCreateUser(t *testing.T) {

	user := casbin.User{
		Username: "admin",
		Password: utils.Md5("123"),
	}

	qUser, err := QueryUserByUsername(user.Username)
	if err != nil {
		t.Fatal(err)
	}

	if qUser.ID != 0 {
		fmt.Println("User already exists")
		return
	}
	err = CreateUser(&user)
	if err != nil {
		t.Fatal(err)
	}

	rUser, err := QueryUser(user.Username, user.Password)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(rUser)

}

func TestQueryUser(t *testing.T) {
	qUser, err := QueryUserByUsername("admin")
	if err != nil {
		t.Fatal(err)
	}

	if qUser.ID != 0 {
		fmt.Println("User already exists")
	}

	fmt.Println(qUser)

}
