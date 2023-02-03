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

import "github.com/cloudwego/hertz-examples/bizdemo/hertz_casbin/biz/model/casbin"

func CreatePermission(permission *casbin.Permission) error {
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
