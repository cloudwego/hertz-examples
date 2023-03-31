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

	"github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/biz/mw/redis"

	"gorm.io/gorm"

	"github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/pkg/constants"
)

// register redis operate strategy
var rdFavorite redis.Favorite

type Favorites struct {
	ID        int64          `json:"id"`
	UserId    int64          `json:"user_id"`
	VideoId   int64          `json:"video_id"`
	CreatedAt time.Time      `json:"create_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"delete_at"`
}

// TableName set table name to make gorm can correctly identify
func (Favorites) TableName() string {
	return constants.FavoritesTableName
}

// AddNewFavorite add favorite relation
func AddNewFavorite(favorite *Favorites) (bool, error) {
	err := DB.Create(favorite).Error
	if err != nil {
		return false, err
	}
	// add data to redis
	if rdFavorite.CheckLiked(favorite.VideoId) {
		rdFavorite.AddLiked(favorite.UserId, favorite.VideoId)
	}
	if rdFavorite.CheckLike(favorite.UserId) {
		rdFavorite.AddLike(favorite.UserId, favorite.VideoId)
	}

	return true, nil
}

// DeleteFavorite delete favorite relation
func DeleteFavorite(favorite *Favorites) (bool, error) {
	err := DB.Where("video_id = ? AND user_id = ?", favorite.VideoId, favorite.UserId).Delete(favorite).Error
	if err != nil {
		return false, err
	}
	// del data if hit
	if rdFavorite.CheckLiked(favorite.VideoId) {
		rdFavorite.DelLiked(favorite.UserId, favorite.VideoId)
	}
	if rdFavorite.CheckLike(favorite.UserId) {
		rdFavorite.DelLike(favorite.UserId, favorite.VideoId)
	}
	return true, nil
}

// QueryFavoriteExist query the like record by video_id and user_id
func QueryFavoriteExist(user_id, video_id int64) (bool, error) {
	if rdFavorite.CheckLiked(video_id) {
		return rdFavorite.ExistLiked(user_id, video_id), nil
	}
	if rdFavorite.CheckLike(user_id) {
		return rdFavorite.ExistLike(user_id, video_id), nil
	}
	var sum int64
	err := DB.Model(&Favorites{}).Where("video_id = ? AND user_id = ?", video_id, user_id).Count(&sum).Error
	if err != nil {
		return false, err
	}
	if sum == 0 {
		return false, nil
	}
	return true, nil
}

// QueryTotalFavoritedByAuthorID query the like num of all the video published by  the user
func QueryTotalFavoritedByAuthorID(author_id int64) (int64, error) {
	var sum int64
	err := DB.Table(constants.FavoritesTableName).Joins("JOIN videos ON likes.video_id = videos.id").
		Where("videos.author_id = ?", author_id).Count(&sum).Error
	if err != nil {
		return 0, err
	}
	return sum, nil
}

// getFavoriteIdList get the id list of video liked by the user in db
func getFavoriteIdList(user_id int64) ([]int64, error) {
	var favorite_actions []Favorites
	err := DB.Where("user_id = ?", user_id).Find(&favorite_actions).Error
	if err != nil {
		return nil, err
	}
	var result []int64
	for _, v := range favorite_actions {
		result = append(result, v.VideoId)
	}
	return result, nil
}

// GetFavoriteIdList get the id list of video liked by the user
func GetFavoriteIdList(user_id int64) ([]int64, error) {
	if rdFavorite.CheckLike(user_id) {
		return rdFavorite.GetLike(user_id), nil
	}
	return getFavoriteIdList(user_id)
}

// GetFavoriteCountByUserID get the num of the video liked by user
func GetFavoriteCountByUserID(user_id int64) (int64, error) {
	if rdFavorite.CheckLike(user_id) {
		return rdFavorite.CountLike(user_id)
	}
	// Not in the cache, go to the database to find and update the cache
	videos, err := getFavoriteIdList(user_id)
	if err != nil {
		return 0, err
	}

	// update redis asynchronously
	go func(user int64, videos []int64) {
		for _, video := range videos {
			rdFavorite.AddLiked(user, video)
		}
	}(user_id, videos)

	return int64(len(videos)), nil
}

// getFavoriterIdList get the id list of liker of video in db
func getFavoriterIdList(video_id int64) ([]int64, error) {
	var favorite_actions []Favorites
	err := DB.Where("video_id = ?", video_id).Find(&favorite_actions).Error
	if err != nil {
		return nil, err
	}
	var result []int64
	for _, v := range favorite_actions {
		result = append(result, v.UserId)
	}
	return result, nil
}

// GetFavoriterIdList get the id list of liker of  video
func GetFavoriterIdList(video_id int64) ([]int64, error) {
	if rdFavorite.CheckLiked(video_id) {
		return rdFavorite.GetLiked(video_id), nil
	}
	return getFavoriterIdList(video_id)
}

// GetFavoriteCount count the favorite of video
func GetFavoriteCount(video_id int64) (int64, error) {
	if rdFavorite.CheckLiked(video_id) {
		return rdFavorite.CountLiked(video_id)
	}
	// Not in the cache, go to the database to find and update the cache
	likes, err := getFavoriterIdList(video_id)
	if err != nil {
		return 0, err
	}

	// update redis asynchronously
	go func(users []int64, video int64) {
		for _, user := range users {
			rdFavorite.AddLiked(user, video)
		}
	}(likes, video_id)
	return int64(len(likes)), nil
}
