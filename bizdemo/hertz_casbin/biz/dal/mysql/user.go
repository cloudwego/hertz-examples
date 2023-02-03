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

func CreateUser(user *casbin.User) error {
	return DB.Create(user).Error

}

func QueryUser(username, password string) (*casbin.User, error) {
	var user casbin.User
	DB.Where("username=? AND password =? ", username, password).First(&user)
	return &user, nil
}

func QueryUserByUsername(username string) (*casbin.User, error) {
	var user casbin.User
	DB.Where("username = ?", username).First(&user)
	return &user, nil
}

func QueryUserById(id int) (*casbin.User, error) {
	var user casbin.User
	DB.First(&user, id)
	return &user, nil
}
