/*
 * Copyright 2023 CloudWeGo Authors
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

/*
 * @Description:
 * @Author: kpole
 * @Date: 2023-02-20 21:03:18
 * @LastEditors: kpole
 */
package db

import (
	"offer_tiktok/pkg/constants"
	"offer_tiktok/pkg/errno"
)

type User struct {
	ID              int64  `json:"id"`
	UserName        string `json:"user_name"`
	Password        string `json:"password"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
}

func (User) TableName() string {
	return constants.UserTableName
}

// CreateUser create user info
func CreateUser(user *User) (int64, error) {
	err := DB.Create(user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, err
}

// 根据 user_name 查询 User
func QueryUser(userName string) (*User, error) {
	var user User
	if err := DB.Where("user_name = ?", userName).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func QueryUserById(user_id int64) (*User, error) {
	var user User
	if err := DB.Where("id = ?", user_id).Find(&user).Error; err != nil {
		return nil, err
	}
	if user == (User{}) {
		err := errno.UserIsNotExistErr
		return nil, err
	}
	return &user, nil
}

func VerifyUser(userName, password string) (int64, error) {
	var user User
	if err := DB.Where("user_name = ? AND password = ?", userName, password).Find(&user).Error; err != nil {
		return 0, err
	}
	if user.ID == 0 {
		err := errno.PasswordIsNotVerified
		return user.ID, err
	}
	return user.ID, nil
}

func CheckUserExistById(user_id int64) (bool, error) {
	var user User
	if err := DB.Where("id = ?", user_id).Find(&user).Error; err != nil {
		return false, err
	}
	if user == (User{}) {
		return false, nil
	}
	return true, nil
}
