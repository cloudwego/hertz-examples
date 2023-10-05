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

package db

import (
	"fmt"
	"testing"
)

func TestCreateUser(t *testing.T) {
	Init()
	u := &User{
		UserName: "test",
		Password: "123456",
	}
	user_id, err := CreateUser(u)
	if err != nil {
		fmt.Printf("%v", false)
		return
	}
	fmt.Printf("%v", user_id)
}

func TestQueryUser(t *testing.T) {
	Init()
	user, err := QueryUser("test")
	if err != nil {
		fmt.Println(false)
		return
	}

	fmt.Printf("%v", user)
}

func TestQueryUser2(t *testing.T) {
	Init()
	user, err := QueryUser("ttttttt")
	if err != nil {
	}
	if *user == (User{}) {
		fmt.Println(true)
		return
	}
	fmt.Println(false)
}

func TestVerifyUser(t *testing.T) {
	Init()
	user_id, err := VerifyUser("test", "123456")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user_id)
}

func TestVerifyUser2(t *testing.T) {
	Init()
	user_id, err := VerifyUser("test", "1234523426")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user_id)
}

func TestQueryUserById(t *testing.T) {
	Init()
	u, err := QueryUserById(int64(1001))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u)
}
