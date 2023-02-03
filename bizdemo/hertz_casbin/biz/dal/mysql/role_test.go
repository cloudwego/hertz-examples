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
	"github.com/darrenli6/hertz-examples/bizdemo/hertz_casbin/biz/model/casbin"
	"testing"
)

func init() {
	Init()
}

func TestCreateRole(t *testing.T) {

	role := casbin.Role{
		Name: "admin",
	}

	qRole, err := QueryRoleByName(role.Name)
	if err != nil {
		t.Fatal(err)
	}

	if qRole.ID != 0 {
		fmt.Println("User already exists")
		return
	}
	err = CreateRole(&role)
	if err != nil {
		t.Fatal(err)
	}

}

func TestQueryRole(t *testing.T) {
	qRole, err := QueryRoleByName("admin")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(qRole)

}

func TestQueryRolesByUid(t *testing.T) {
	qRoles := QueryRolesByUid(int(5))

	fmt.Println(qRoles)

}
func TestBindUserRole(t *testing.T) {

	userRole := casbin.UserRole{
		UID: 5,
		Rid: 1,
	}

	userRoles1 := QueryUserRoleByIds(int(userRole.UID), int(userRole.Rid))

	if len(userRoles1) > 0 {
		t.Fatal("Data already exists")
	}

	// check user
	user, err := QueryUserById(int(userRole.UID))
	if err != nil {
		t.Fatal(err)
	}

	if user.ID == 0 {
		t.Fatal("User data does not exist")
	}

	crrole, err := QueryRoleById(int(userRole.Rid))
	if err != nil {
		t.Fatal(err)
	}

	if crrole.ID == 0 {
		t.Fatal("Role data does not exist")
	}

	err = BindRole(&userRole)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(userRole)

}
