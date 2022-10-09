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

package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `json:"user_name" column:"user_name"`
	Email    string `json:"email" column:"email"`
	Password string `json:"password" column:"password"`
}

func (u *User) TableName() string {
	return "users"
}
