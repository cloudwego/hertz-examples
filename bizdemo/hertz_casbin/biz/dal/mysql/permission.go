package mysql

import "github.com/darrenli6/hertz-examples/bizdemo/hertz_casbin/biz/model/casbin"

func CreatePermisson(permission *casbin.Permission) error {
	return DB.Create(permission).Error
}

func BindPermissionRole(permissionRole *casbin.PermissionRole) error {
	return DB.Create(permissionRole).Error
}

func QueryPermissionById(id int) (*casbin.Permission, error) {
	var permission casbin.Permission
	DB.First(&permission, id)
	return &permission, nil
}

func QueryPermissionByV(v1 string, v2 string) (*casbin.Permission, error) {
	var permission casbin.Permission
	DB.Where("v1= ? AND v2 =?", v1, v2).First(&permission)
	return &permission, nil
}

func QuerypermissionRoleByIds(pid, rid int) []casbin.PermissionRole {

	var permissionRole []casbin.PermissionRole
	tx := DB.Model(new(casbin.PermissionRole))
	if pid != 0 {
		tx.Where("pid= ?", pid)
	}
	if rid != 0 {
		tx.Where("rid= ?", rid)
	}

	tx.Select("pid,rid,id").Find(&permissionRole)

	return permissionRole
}
