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
	"time"

	"gorm.io/gorm"

	redis "github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/biz/mw/redis"
	"github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/pkg/constants"
)

// Follows follower is fan of user
type Follows struct {
	ID         int64          `json:"id"`
	UserId     int64          `json:"user_id"`
	FollowerId int64          `json:"follower_id"`
	CreatedAt  time.Time      `json:"create_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"delete_at"`
}

// register redis operate strategy
var rdFollows redis.Follows

// TableName set table name to make gorm can correctly identify
func (Follows) TableName() string {
	return constants.FollowsTableName
}

func AddNewFollow(follow *Follows) (bool, error) {
	err := DB.Create(follow).Error
	if err != nil {
		return false, err
	}
	// add data to redis
	if rdFollows.CheckFollow(follow.FollowerId) {
		rdFollows.AddFollow(follow.UserId, follow.FollowerId)
	}
	if rdFollows.CheckFollower(follow.UserId) {
		rdFollows.AddFollower(follow.UserId, follow.FollowerId)
	}

	return true, nil
}

// DeleteFollow delete follow relation in db and update redis
func DeleteFollow(follow *Follows) (bool, error) {
	err := DB.Where("user_id = ? AND follower_id = ?", follow.UserId, follow.FollowerId).Delete(follow).Error
	if err != nil {
		return false, err
	}
	// if redis hit del
	if rdFollows.CheckFollow(follow.FollowerId) {
		rdFollows.DelFollow(follow.UserId, follow.FollowerId)
	}
	if rdFollows.CheckFollower(follow.UserId) {
		rdFollows.DelFollower(follow.UserId, follow.FollowerId)
	}
	return true, nil
}

// QueryFollowExist check the relation of user and follower
func QueryFollowExist(user_id, follower_id int64) (bool, error) {
	if rdFollows.CheckFollow(follower_id) {
		return rdFollows.ExistFollow(user_id, follower_id), nil
	}
	if rdFollows.CheckFollower(user_id) {
		return rdFollows.ExistFollower(user_id, follower_id), nil
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
	return true, nil
}

// GetFollowCount query the number of users following
func GetFollowCount(follower_id int64) (int64, error) {
	if rdFollows.CheckFollow(follower_id) {
		return rdFollows.CountFollow(follower_id)
	}

	// Not in the cache, go to the database to find and update the cache
	followings, err := getFollowIdList(follower_id)
	if err != nil {
		return 0, err
	}
	// update redis asynchronously
	go addFollowRelationToRedis(follower_id, followings)
	return int64(len(followings)), nil
}

// addFollowRelationToRedis update redis.RdbFollowing
func addFollowRelationToRedis(follower_id int64, followings []int64) {
	for _, following := range followings {
		rdFollows.AddFollow(following, follower_id)
	}
}

// GetFollowerCount query the number of followers of a user
func GetFollowerCount(user_id int64) (int64, error) {
	if rdFollows.CheckFollower(user_id) {
		return rdFollows.CountFollower(user_id)
	}
	// Not in the cache, go to the database to find and update the cache
	followers, err := getFollowerIdList(user_id)
	if err != nil {
		return 0, err
	}
	// update redis asynchronously
	go addFollowerRelationToRedis(user_id, followers)
	return int64(len(followers)), nil
}

// addFollowerRelationToRedis update redis.RdbFollower
func addFollowerRelationToRedis(user_id int64, followers []int64) {
	for _, follower := range followers {
		rdFollows.AddFollower(user_id, follower)
	}
}

// getFollowIdList find user_id follow id list in db
func getFollowIdList(follower_id int64) ([]int64, error) {
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

// GetFollowIdList find user_id follow id list in db or rdb
func GetFollowIdList(follower_id int64) ([]int64, error) {
	if rdFollows.CheckFollow(follower_id) {
		return rdFollows.GetFollow(follower_id), nil
	}
	return getFollowIdList(follower_id)
}

// getFollowerIdList get follower id list in db
func getFollowerIdList(user_id int64) ([]int64, error) {
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

// GetFollowerIdList get follower id list in db or rdb
func GetFollowerIdList(user_id int64) ([]int64, error) {
	if rdFollows.CheckFollower(user_id) {
		return rdFollows.GetFollower(user_id), nil
	}
	return getFollowerIdList(user_id)
}

func GetFriendIdList(user_id int64) ([]int64, error) {
	if !rdFollows.CheckFollow(user_id) {
		following, err := getFollowIdList(user_id)
		if err != nil {
			return *new([]int64), err
		}
		addFollowRelationToRedis(user_id, following)
	}
	if !rdFollows.CheckFollower(user_id) {
		followers, err := getFollowerIdList(user_id)
		if err != nil {
			return *new([]int64), err
		}
		addFollowerRelationToRedis(user_id, followers)
	}
	return rdFollows.GetFriend(user_id), nil
}
