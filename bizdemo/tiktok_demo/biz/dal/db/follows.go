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
	"strconv"
	"time"

	"gorm.io/gorm"

	redis "github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/biz/mw/redis"
	"github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/pkg/constants"
)

// Follows user_id follow follower_id
type Follows struct {
	ID         int64          `json:"id"`
	UserId     int64          `json:"user_id"`
	FollowerId int64          `json:"follower_id"`
	CreatedAt  time.Time      `json:"create_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"delete_at"`
}

// TableName set table name to make gorm can correctly identify
func (Follows) TableName() string {
	return constants.FollowsTableName
}

func AddNewFollow(follow *Follows) (bool, error) {
	err := DB.Create(follow).Error
	if err != nil {
		return false, err
	}
	// if redis hit update
	if cnt, _ := redis.RdbFollowing.SCard(strconv.Itoa(int(follow.UserId))).Result(); cnt > 0 {
		redis.RdbFollowing.SAdd(strconv.Itoa(int(follow.UserId)), follow.FollowerId)
		redis.RdbFollowing.Expire(strconv.Itoa(int(follow.UserId)), redis.ExpireTime)
	}
	if cnt, _ := redis.RdbFollower.SCard(strconv.Itoa(int(follow.FollowerId))).Result(); cnt > 0 {
		redis.RdbFollower.SAdd(strconv.Itoa(int(follow.FollowerId)), follow.UserId)
		redis.RdbFollower.Expire(strconv.Itoa(int(follow.FollowerId)), redis.ExpireTime)
	}
	return true, nil
}

// DeleteFollow delete follow relation in db and update redis
func DeleteFollow(follow *Follows) (bool, error) {
	err := DB.Where("user_id = ? AND follower_id = ?", follow.UserId, follow.FollowerId).Delete(follow).Error
	if err != nil {
		return false, err
	}
	// if redis hit update
	if cnt, _ := redis.RdbFollowing.SCard(strconv.Itoa(int(follow.UserId))).Result(); cnt > 0 {
		redis.RdbFollowing.SRem(strconv.Itoa(int(follow.UserId)), follow.FollowerId)
		redis.RdbFollowing.Expire(strconv.Itoa(int(follow.UserId)), redis.ExpireTime)
	}
	if cnt, _ := redis.RdbFollower.SCard(strconv.Itoa(int(follow.FollowerId))).Result(); cnt > 0 {
		redis.RdbFollower.SRem(strconv.Itoa(int(follow.FollowerId)), follow.UserId)
		redis.RdbFollower.Expire(strconv.Itoa(int(follow.FollowerId)), redis.ExpireTime)
	}
	return true, nil
}

// QueryFollowExist check the relation of user and follower
func QueryFollowExist(user_id, follower_id int64) (bool, error) {
	if exist, err := redis.RdbFollowing.SIsMember(strconv.Itoa(int(user_id)), follower_id).Result(); exist {
		redis.RdbFollowing.Expire(strconv.Itoa(int(user_id)), redis.ExpireTime)
		return true, err
	}
	follow := Follows{
		UserId:     user_id,
		FollowerId: follower_id,
	}
	err := DB.Where("user_id = ? AND follower_id = ?", user_id, follower_id).Find(&follow).Error
	if err != nil {
		return false, err
	}
	if follow.ID == 0 {
		return false, nil
	}
	// Update redis asynchronously
	go addRelationToRedis(int(follow.UserId), int(follow.FollowerId))
	return true, nil
}

// GetFollowCount query the number of users following
func GetFollowCount(user_id int64) (int64, error) {
	if count, err := redis.RdbFollowing.SCard(strconv.Itoa(int(user_id))).Result(); count > 0 {
		// update expiration time
		redis.RdbFollowing.Expire(strconv.Itoa(int(user_id)), redis.ExpireTime)
		return count, err
	}
	// Not in the cache, go to the database to find and update the cache
	followings, err := GetFollowIdList(user_id)
	if err != nil {
		return 0, err
	}
	// update cache
	go AddNewFollowRelationToRedis(user_id, followings)
	return int64(len(followings)), nil
}

// AddNewFollowRelationToRedis update redis.RdbFollowing
func AddNewFollowRelationToRedis(user_id int64, followings []int64) {
	for _, following := range followings {
		redis.RdbFollowing.SAdd(strconv.Itoa(int(user_id)), following)
	}
	// update expiration time to keep data hot
	redis.RdbFollowing.Expire(strconv.Itoa(int(user_id)), redis.ExpireTime)
}

// GetFolloweeCount query the number of followers of a user
func GetFolloweeCount(follower_id int64) (int64, error) {
	if count, err := redis.RdbFollower.SCard(strconv.Itoa(int(follower_id))).Result(); count > 0 {
		// update expire time
		redis.RdbFollower.Expire(strconv.Itoa(int(follower_id)), redis.ExpireTime)
		return count, err
	}
	// Not in the cache, go to the database to find and update the cache
	followers, err := GetFollowerIdList(follower_id)
	if err != nil {
		return 0, err
	}
	go AddNewFollowerRelationToRedis(follower_id, followers)
	return int64(len(followers)), nil
}

// AddNewFollowerRelationToRedis update redis.RdbFollower
func AddNewFollowerRelationToRedis(user_id int64, followers []int64) {
	for _, follower := range followers {
		redis.RdbFollower.SAdd(strconv.Itoa(int(user_id)), follower)
	}
	redis.RdbFollower.Expire(strconv.Itoa(int(user_id)), redis.ExpireTime)
}

// GetFollowIdList find user_id follow id list
func GetFollowIdList(user_id int64) ([]int64, error) {
	var follow_actions []Follows
	err := DB.Where("user_id = ?", user_id).Find(&follow_actions).Error
	if err != nil {
		return nil, err
	}
	var result []int64
	for _, v := range follow_actions {
		result = append(result, v.FollowerId)
	}
	return result, nil
}

// GetFollowerIdList get follower id list
func GetFollowerIdList(follower_id int64) ([]int64, error) {
	var follow_actions []Follows
	err := DB.Where("follower_id = ?", follower_id).Find(&follow_actions).Error
	if err != nil {
		return nil, err
	}
	var result []int64
	for _, v := range follow_actions {
		result = append(result, v.UserId)
	}
	return result, nil
}

// CheckFollowRelationExist use user id and follower id to check the relation
func CheckFollowRelationExist(follow *Follows) (bool, error) {
	if exist, err := redis.RdbFollowing.SIsMember(strconv.Itoa(int(follow.UserId)), follow.FollowerId).Result(); exist {
		redis.RdbFollowing.Expire(strconv.Itoa(int(follow.UserId)), redis.ExpireTime)
		return true, err
	}
	err := DB.Where("user_id = ? AND follower_id = ?", follow.UserId, follow.FollowerId).Find(&follow).Error
	if err != nil {
		return false, err
	}
	if follow.ID == 0 {
		return false, nil
	}
	go addRelationToRedis(int(follow.UserId), int(follow.FollowerId))
	return true, nil
}

// addRelationToRedis add relation and extend expiration time
func addRelationToRedis(user_id, follow_id int) {
	redis.RdbFollowing.SAdd(strconv.Itoa(user_id), follow_id)
	redis.RdbFollowing.Expire(strconv.Itoa(user_id), redis.ExpireTime)
}
