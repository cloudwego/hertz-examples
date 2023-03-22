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

package redis

import (
	"fmt"
	"strconv"
	"testing"
)

func TestQueryCount(t *testing.T) {
	InitRedis()
	user_id := 1003
	if cnt, err := RdbFollowing.SCard(strconv.Itoa(user_id)).Result(); cnt > 0 {
		// 更新过期时间。
		RdbFollowing.Expire(strconv.Itoa(int(user_id)), ExpireTime)
		fmt.Println(cnt, err)
	}
}

func TestAddFollow(t *testing.T) {
	InitRedis()
	user_id := 1003
	RdbFollowing.SAdd(strconv.Itoa(user_id), 1005)
	RdbFollowing.SAdd(strconv.Itoa(user_id), 1006)
}
