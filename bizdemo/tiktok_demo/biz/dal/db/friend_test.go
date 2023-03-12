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

func TestGetFriendList(t *testing.T) {
	Init()
	followerList, err := GetFollowerIdList(1001)
	if err != nil {
		fmt.Println("false")
		return
	}
	for _, followerId := range followerList {
		isFriend, err := QueryFollowExist(&Follows{UserId: 1001, FollowerId: followerId})
		if err != nil {
			fmt.Println("false")
			return
		}
		if isFriend {
			fmt.Println(followerId)
		}
	}
}
