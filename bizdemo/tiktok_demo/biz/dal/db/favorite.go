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

	"github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/pkg/constants"
	"github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/pkg/errno"
)

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
	return true, nil
}

// DeleteFavorite delete favorite relation
func DeleteFavorite(favorite *Favorites) (bool, error) {
	err := DB.Where("video_id = ? AND user_id = ?", favorite.VideoId, favorite.UserId).Delete(favorite).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

// QueryFavoriteExist query the publishing record by video_id and user_id
func QueryFavoriteExist(video_id, user_id int64) (bool, error) {
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

// GetFavoriteCount count the favorite of video
func GetFavoriteCount(video_id int64) (int64, error) {
	var count int64
	err := DB.Model(&Favorites{}).Where("video_id = ?", video_id).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// GetFavoriteIdList get the id list of video liked by the user
func GetFavoriteIdList(user_id int64) ([]int64, error) {
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

// GetFavoriteCountByUserID get the num of the video liked by user
func GetFavoriteCountByUserID(user_id int64) (int64, error) {
	var count int64
	err := DB.Model(&Favorites{}).Where("user_id = ?", user_id).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// GetFavoriterIdList get the id list of liker of  video
func GetFavoriterIdList(video_id int64) ([]int64, error) {
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

// CheckFavoriteRelationExist check the favorite relation if exist
func CheckFavoriteRelationExist(favorite *Favorites) (bool, error) {
	err := DB.Where("video_id = ? AND user_id = ?", favorite.VideoId, favorite.UserId).Find(&favorite).Error
	if err != nil {
		return false, err
	}
	// If no matching data is found, an empty structure will be returned
	if favorite.ID == 0 {
		err := errno.FavoriteRelationNotExistErr
		return false, err
	}
	return true, nil
}
