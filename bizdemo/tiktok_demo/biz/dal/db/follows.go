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

// Follows user_id 关注了 follower_id
type Follows struct {
	ID         int64          `json:"id"`
	UserId     int64          `json:"user_id"`
	FollowerId int64          `json:"follower_id"`
	CreatedAt  time.Time      `json:"create_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"delete_at"`
}

func (Follows) TableName() string {
	return constants.FollowsTableName
}

func AddNewFollow(follow *Follows) (bool, error) {
	err := DB.Create(follow).Error
	if err != nil {
		return false, err
	}
	// 如果缓存中有此关系，更新
	if cnt, _ := redis.RdbFollowing.SCard(strconv.Itoa(int(follow.UserId))).Result(); cnt > 0 {
		redis.RdbFollowing.SAdd(strconv.Itoa(int(follow.UserId)), follow.FollowerId)
		redis.RdbFollower.Expire(strconv.Itoa(int(follow.UserId)), redis.ExpireTime)
	}
	if cnt, _ := redis.RdbFollower.SCard(strconv.Itoa(int(follow.FollowerId))).Result(); cnt > 0 {
		redis.RdbFollower.SAdd(strconv.Itoa(int(follow.FollowerId)), follow.UserId)
		redis.RdbFollower.Expire(strconv.Itoa(int(follow.FollowerId)), redis.ExpireTime)
	}
	return true, nil
}

func DeleteFollow(follow *Follows) (bool, error) {
	err := DB.Where("user_id = ? AND follower_id = ?", follow.UserId, follow.FollowerId).Delete(follow).Error
	if err != nil {
		return false, err
	}
	// 如果缓存中有此关系，更新
	if cnt, _ := redis.RdbFollowing.SCard(strconv.Itoa(int(follow.UserId))).Result(); cnt > 0 {
		redis.RdbFollowing.SRem(strconv.Itoa(int(follow.UserId)), follow.FollowerId)
		redis.RdbFollower.Expire(strconv.Itoa(int(follow.UserId)), redis.ExpireTime)
	}
	if cnt, _ := redis.RdbFollower.SCard(strconv.Itoa(int(follow.FollowerId))).Result(); cnt > 0 {
		redis.RdbFollower.SRem(strconv.Itoa(int(follow.FollowerId)), follow.UserId)
		redis.RdbFollower.Expire(strconv.Itoa(int(follow.FollowerId)), redis.ExpireTime)
	}
	return true, nil
}

func QueryFollowExist(follow *Follows) (bool, error) {
	if exist, err := redis.RdbFollowing.SIsMember(strconv.Itoa(int(follow.UserId)), follow.FollowerId).Result(); exist {
		// fmt.Printf("在redis中获取到user_id:%d关注了user:%d的关系\n", follow.UserId, follow.FollowerId)
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
	// 异步更新redis
	go addRelationToRedis(int(follow.UserId), int(follow.FollowerId))
	return true, nil
}

// GetFollowCount 查询用户的关注数量
func GetFollowCount(user_id int64) (int64, error) {
	if count, err := redis.RdbFollowing.SCard(strconv.Itoa(int(user_id))).Result(); count > 0 {
		// 更新过期时间。
		redis.RdbFollowing.Expire(strconv.Itoa(int(user_id)), redis.ExpireTime)
		// fmt.Printf("在redis中获取到user_id:%d的关注数%d\n", user_id, count)
		return count, err
	}
	// 缓存中没有，去数据库查找并更新缓存
	followings, err := GetFollowIdList(user_id)
	if err != nil {
		return 0, err
	}
	// 更新Redis
	go AddNewFollowRelationToRedis(user_id, followings)
	return int64(len(followings)), nil
}

// AddNewFollowRelationToRedis 更新redis.RdbFollowing

func AddNewFollowRelationToRedis(user_id int64, followings []int64) {
	for _, following := range followings {
		// redis.RdbFollowing.SAdd(strconv.Itoa(int(user_id)), -1)
		redis.RdbFollowing.SAdd(strconv.Itoa(int(user_id)), following)
	}
	// 更新过期时间，保持数据热度
	redis.RdbFollowing.Expire(strconv.Itoa(int(user_id)), redis.ExpireTime)
}

// GetFolloweeCount 查询用户的粉丝数量
func GetFolloweeCount(follower_id int64) (int64, error) {
	if count, err := redis.RdbFollower.SCard(strconv.Itoa(int(follower_id))).Result(); count > 0 {
		// 更新过期时间。
		redis.RdbFollower.Expire(strconv.Itoa(int(follower_id)), redis.ExpireTime)
		// fmt.Printf("在redis中获取到user_id:%d的粉丝数%d\n", follower_id, count)
		return count, err
	}
	// 缓存中没有，去数据库查找并更新缓存
	followers, err := GetFollowerIdList(follower_id)
	if err != nil {
		return 0, err
	}
	// 更新Redis
	go AddNewFollowerRelationToRedis(follower_id, followers)
	return int64(len(followers)), nil
}

// AddNewFollowerRelationToRedis 更新redis.RdbFollower
func AddNewFollowerRelationToRedis(user_id int64, followers []int64) {
	for _, follower := range followers {
		redis.RdbFollower.SAdd(strconv.Itoa(int(user_id)), follower)
	}
	// 更新过期时间，保持数据热度
	redis.RdbFollower.Expire(strconv.Itoa(int(user_id)), redis.ExpireTime)
}

// GetFollowIdList 获得 user_id 关注的人的 id
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

// GetFollowerIdList 获得 follower_id 所有粉丝的 id
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

func CheckFollowRelationExist(follow *Follows) (bool, error) {
	if exist, err := redis.RdbFollowing.SIsMember(strconv.Itoa(int(follow.UserId)), follow.FollowerId).Result(); exist {
		// fmt.Printf("在redis中获取到user_id:%d关注了user:%d的关系\n", follow.UserId, follow.FollowerId)
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
	// 异步更新redis
	go addRelationToRedis(int(follow.UserId), int(follow.FollowerId))
	return true, nil
}

func addRelationToRedis(user_id, follow_id int) {
	redis.RdbFollowing.SAdd(strconv.Itoa(user_id), follow_id)
	redis.RdbFollowing.Expire(strconv.Itoa(user_id), redis.ExpireTime)
}
