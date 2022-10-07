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
	"github.com/cloudwego/hertz-examples/bizdemo/hertz_jwt/biz/model"
)

func CreateUsers(users []*model.User) error {
	return DB.Create(users).Error
}

func FindUserByNameOrEmail(userName, email string) ([]*model.User, error) {
	res := make([]*model.User, 0)
	if err := DB.Where(DB.Or("user_name = ?", userName).
		Or("email = ?", email)).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func CheckUser(account, password string) ([]*model.User, error) {
	res := make([]*model.User, 0)
	if err := DB.Where(DB.Or("user_name = ?", account).
		Or("email = ?", account)).Where("password = ?", password).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
