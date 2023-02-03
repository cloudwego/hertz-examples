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
	"github.com/darrenli6/hertz-examples/bizdemo/hertz_casbin/pkg/consts"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(consts.MysqlDSN), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	m := DB.Migrator()
	if m.HasTable(&casbin.User{}) {
		return
	}
	if err = m.CreateTable(&casbin.User{}, &casbin.Role{}, &casbin.Permission{}, &casbin.UserRole{}, &casbin.PermissionRole{}); err != nil {
		panic(err)
	}
}
