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
	"strconv"
)

const (
	followerSuffix = ":follower"
	followSuffix   = ":follow"
)

type (
	Follows struct{}
)

func (f Follows) AddFollow(user_id, follower_id int64) {
	add(rdbFollows, strconv.FormatInt(follower_id, 10)+followSuffix, user_id)
}

func (f Follows) AddFollower(user_id, follower_id int64) {
	add(rdbFollows, strconv.FormatInt(user_id, 10)+followerSuffix, follower_id)
}

func (f Follows) DelFollow(user_id, follower_id int64) {
	del(rdbFollows, strconv.FormatInt(follower_id, 10)+followSuffix, user_id)
}

func (f Follows) DelFollower(user_id, follower_id int64) {
	del(rdbFollows, strconv.FormatInt(user_id, 10)+followerSuffix, follower_id)
}

func (f Follows) CheckFollow(follower_id int64) bool {
	return check(rdbFollows, strconv.FormatInt(follower_id, 10)+followSuffix)
}

func (f Follows) CheckFollower(user_id int64) bool {
	return check(rdbFollows, strconv.FormatInt(user_id, 10)+followerSuffix)
}

func (f Follows) ExistFollow(user_id, follower_id int64) bool {
	return exist(rdbFollows, strconv.FormatInt(follower_id, 10)+followSuffix, user_id)
}

func (f Follows) ExistFollower(user_id, follower_id int64) bool {
	return exist(rdbFollows, strconv.FormatInt(user_id, 10)+followerSuffix, follower_id)
}

func (f Follows) CountFollow(follower_id int64) (int64, error) {
	return count(rdbFollows, strconv.FormatInt(follower_id, 10)+followSuffix)
}

func (f Follows) CountFollower(user_id int64) (int64, error) {
	return count(rdbFollows, strconv.FormatInt(user_id, 10)+followerSuffix)
}

func (f Follows) GetFollow(follower_id int64) []int64 {
	return get(rdbFollows, strconv.FormatInt(follower_id, 10)+followSuffix)
}

func (f Follows) GetFollower(user_id int64) []int64 {
	return get(rdbFollows, strconv.FormatInt(user_id, 10)+followerSuffix)
}

// GetFriend get the friend of the id via intersection
func (f Follows) GetFriend(id int64) (friends []int64) {
	ks1 := strconv.FormatInt(id, 10) + followSuffix
	ks2 := strconv.FormatInt(id, 10) + followerSuffix
	v, _ := rdbFollows.SInter(ks1, ks2).Result()
	for _, vs := range v {
		v_i64, _ := strconv.ParseInt(vs, 10, 64)
		friends = append(friends, v_i64)
	}
	return friends
}
