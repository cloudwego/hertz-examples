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
	"github.com/cloudwego/hertz-examples/bizdemo/hertz_session/pkg/consts"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

func (u *User) TableName() string {
	return consts.UserTableName
}

func CreateUsers(users []*User) error {
	return DB.Create(users).Error
}

func FindUserByNameOrEmail(username, email string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.Where("username = ?", username).Or("email = ?", email).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func CheckUser(username, password string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.Where("username = ? AND password = ?", username, password).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func CheckUserExists(username string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.Where("username = ?   ", username).Limit(1).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func RoleList(role string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.Where("role = ?   ", role).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
