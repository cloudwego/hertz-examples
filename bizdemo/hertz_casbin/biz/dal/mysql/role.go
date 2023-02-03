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
	"github.com/darrenli6/hertz-examples/bizdemo/hertz_casbin/biz/model/casbin"
)

func CreateRole(role *casbin.Role) error {
	return DB.Create(role).Error
}

func BindRole(userRole *casbin.UserRole) error {
	return DB.Create(userRole).Error
}

func QueryRoleById(id int) (*casbin.Role, error) {
	var role casbin.Role
	DB.First(&role, id)
	return &role, nil
}

func QueryRoleByName(name string) (*casbin.Role, error) {
	var role casbin.Role
	DB.Where("name= ?", name).First(&role)
	return &role, nil
}

func QueryUserRoleByIds(uid, rid int) []casbin.UserRole {

	var userRole []casbin.UserRole
	tx := DB.Model(new(casbin.UserRole))
	if uid != 0 {
		tx.Where("uid= ?", uid)
	}
	if rid != 0 {
		tx.Where("rid= ?", rid)
	}

	tx.Select("rid,uid,id").Find(&userRole)

	return userRole
}

func QueryRolesByUid(uid int) []casbin.Role {

	var userRole []casbin.Role
	_ = DB.Model(new(casbin.UserRole)).
		Joins("LEFT JOIN roles on roles.id=user_roles.rid  ").
		Select("roles.id,roles.name").
		Where("user_roles.uid=?", uid).
		Scan(&userRole)

	return userRole
}
