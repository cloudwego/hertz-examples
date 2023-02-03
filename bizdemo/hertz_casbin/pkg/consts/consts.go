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

package consts

import (
	"github.com/cloudwego/hertz-examples/bizdemo/hertz_casbin/biz/model/casbin"
	"github.com/golang-jwt/jwt/v4"
)

var (
	MysqlDSN   = "root:casbin@tcp(localhost:9912)/casbin?charset=utf8&parseTime=True&loc=Local"
	EmqxKey    = "1f9c5b734fe27865"
	EmqxSecret = "lV9C2iefOp9Cr9BeiB5rr3N9CBolJjKk3HruhqEpHQxsuVD"
)

type M map[string]interface{}

type UserClaim struct {
	Id       uint          `json:"id"`
	Username string        `json:"username"`
	Roles    []casbin.Role `json:"rids"`
	jwt.RegisteredClaims
}

var (
	JwtKey = "darren"
)
