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
)

func init() {
	Init()
}

// create permission data
func TestCreatePermission(t *testing.T) {

	permission := casbin.Permission{
		V1: "/1",
		V2: "/2",
	}

	qRole, err := QueryPermissionByV(permission.V1, permission.V2)
	if err != nil {
		t.Fatal(err)
	}

	if qRole.ID != 0 {
		fmt.Println("Permission already exists")
		return
	}
	err = CreatePermission(&permission)
	if err != nil {
		t.Fatal(err)
	}

}

// query permission data
func TestQueryPermission(t *testing.T) {
	qRole, err := QueryPermissionByV("/v1/role/create", "POST")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(qRole)

}

// bind permission and role
func TestBindPermission(t *testing.T) {

	permissionRole := casbin.PermissionRole{
		Pid: 1,
		Rid: 1,
	}

	rpermissionRole := QuerypermissionRoleByIds(int(permissionRole.Pid), int(permissionRole.Rid))

	if len(rpermissionRole) > 0 {
		t.Fatal("Data already exists")
	}

	// 检查用户
	role, err := QueryRoleById(int(permissionRole.Rid))
	if err != nil {
		t.Fatal(err)
	}

	if role.ID == 0 {
		t.Fatal("role data does not exist ")
	}

	rPermission, err := QueryPermissionById(int(permissionRole.Pid))
	if err != nil {
		t.Fatal(err)
	}

	if rPermission.ID == 0 {
		t.Fatal("Permission data does not exist")
	}

	err = BindPermissionRole(&permissionRole)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(permissionRole)

}
