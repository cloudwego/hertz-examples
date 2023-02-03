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

package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/cloudwego/hertz-examples/bizdemo/hertz_casbin/pkg/consts"
	"github.com/cloudwego/hertz-examples/bizdemo/hertz_casbin/pkg/utils"
)

// Test the login of the admin user to get the token
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

// Test the login of the role user to get the token
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

// Test the login of the permission user to get the token
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
