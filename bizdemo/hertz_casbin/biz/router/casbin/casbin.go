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


// Code generated by hertz generator. DO NOT EDIT.

package Casbin

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	casbin "github.com/cloudwego/hertz-examples/bizdemo/hertz_casbin/biz/handler/casbin"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_v1 := root.Group("/v1", _v1Mw()...)
		_v1.POST("/login", append(_loginMw(), casbin.Login)...)
		{
			_permission := _v1.Group("/permission", _permissionMw()...)
			{
				_create := _permission.Group("/create", _createMw()...)
				_create.POST("/", append(_createpermissionMw(), casbin.CreatePermission)...)
			}
		}
		{
			_permissionrole := _v1.Group("/permissionrole", _permissionroleMw()...)
			{
				_bind := _permissionrole.Group("/bind", _bindMw()...)
				_bind.POST("/", append(_bindpermissionroleMw(), casbin.BindPermissionRole)...)
			}
		}
		{
			_role := _v1.Group("/role", _roleMw()...)
			{
				_bind0 := _role.Group("/bind", _bind0Mw()...)
				_bind0.POST("/", append(_bindroleMw(), casbin.BindRole)...)
			}
			{
				_create0 := _role.Group("/create", _create0Mw()...)
				_create0.POST("/", append(_createroleMw(), casbin.CreateRole)...)
			}
		}
	}
}