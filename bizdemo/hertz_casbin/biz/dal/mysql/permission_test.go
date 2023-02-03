package mysql

import (
	"fmt"
	"github.com/darrenli6/hertz-examples/bizdemo/hertz_casbin/biz/model/casbin"
	"testing"
)

func init() {
	Init()
}

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
	err = CreatePermisson(&permission)
	if err != nil {
		t.Fatal(err)
	}

}

func TestQueryPermission(t *testing.T) {
	qRole, err := QueryPermissionByV("/v1/role/create", "POST")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(qRole)

}

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
