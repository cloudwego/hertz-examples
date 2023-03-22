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
	"time"
)

func TestAddNewMessage(t *testing.T) {
	Init()
	// one_min_before, _ := time.ParseDuration("-1m")
	message := &Messages{
		ToUserId:   1001,
		FromUserId: 1004,
		Content:    "test: 1004 -> 1001, this is a message",
		// CreatedAt:  time.Now().Add(one_min_before),
	}
	is_succ, err := AddNewMessage(message)
	if err != nil {
		fmt.Println("err 2")
	}
	if !is_succ {
		fmt.Println("failed 1")
	}
	time.Sleep(time.Second)
	message = &Messages{
		ToUserId:   1004,
		FromUserId: 1001,
		Content:    "test: 1001 -> 1004, this is latest message",
		// CreatedAt:  time.Now(),
	}
	is_succ, err = AddNewMessage(message)
	if err != nil {
		fmt.Println("err 2")
	}
	if !is_succ {
		fmt.Println("failed 2")
	}
}

func TestGetMessageByIdPair(t *testing.T) {
	Init()
	user_id1, user_id2 := 1004, 1001
	// 假设过来的是毫秒
	pre_msg_timestamp := int64(1676819765000)
	pre_msg_time := time.Unix(pre_msg_timestamp/1000, pre_msg_timestamp%1000*1000000)
	fmt.Println(pre_msg_time)

	messages, err := GetMessageByIdPair(int64(user_id1), int64(user_id2), pre_msg_time)
	if err != nil {
		fmt.Println("get error")
	}
	for _, message := range messages {
		fmt.Printf("%v\n", message)
	}
	fmt.Println("OK")
}

// 查看好友列表时需要返回最新一条的聊天消息，故在此测试
func TestGetLatestMessage(t *testing.T) {
	Init()
	var id1, id2 int64 = 1001, 1005
	message, err := GetLatestMessageByIdPair(id1, id2)
	if err != nil {
		fmt.Println("false")
	} else if message == nil {
		fmt.Println("1001 与 1005 没有消息")
	} else {
		fmt.Printf("%v\n", message)
	}
}
