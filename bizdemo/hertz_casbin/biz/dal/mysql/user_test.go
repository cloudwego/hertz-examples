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

package mysql

import (
	"fmt"
	"testing"

	"github.com/cloudwego/hertz-examples/bizdemo/hertz_casbin/biz/model/casbin"
	"github.com/cloudwego/hertz-examples/bizdemo/hertz_casbin/pkg/utils"
)

func init() {
	Init()
}

// create the user data
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

// query user by the username of user
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
