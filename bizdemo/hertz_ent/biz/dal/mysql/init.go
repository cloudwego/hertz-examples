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
	"context"

	"github.com/cloudwego/hertz-examples/bizdemo/hertz_ent/biz/model/ent"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dsn    = "ent:ent@tcp(localhost:3306)/ent?charset=utf8&parseTime=True&loc=Local"
	Client *ent.Client
)

func Init() {
	var err error
	Client, err = ent.Open("mysql", dsn)
	if err != nil {
		hlog.Fatalf("failed opening connection to mysql: %v", err)
	}

	// Run the auto migration tool.
	if err := Client.Schema.Create(context.Background()); err != nil {
		hlog.Fatalf("failed creating schema resources: %v", err)
	}
}
