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

package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/cloudwego/hertz-examples/bizdemo/hertz_casbin/pkg/consts"
	"github.com/cloudwego/hertz-examples/bizdemo/hertz_casbin/pkg/utils"
)

// Test add permission
func TestPermissionAdd(t *testing.T) {

	header, _ = json.Marshal(rolem1)
	m := consts.M{
		"v1": "/v1/permission/create/",
		"v2": "POST",
	}

	data, _ := json.Marshal(m)
	rep, err := utils.HttpPost(userServiceAddr+"/v1/permission/create", data, header...)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(rep))
}

// Test bind permission
func TestPermissionBind(t *testing.T) {
	m := consts.M{
		"pid": "1",
		"rid": "1",
	}

	data, _ := json.Marshal(m)
	rep, err := utils.HttpPost(userServiceAddr+"/v1/permissionrole/bind/", data, header...)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(rep))
}
