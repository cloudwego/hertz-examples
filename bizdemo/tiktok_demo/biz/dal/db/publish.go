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
	"offer_tiktok/pkg/constants"
	"time"
)

type Video struct {
	ID          int64
	AuthorID    int64
	PlayURL     string
	CoverURL    string
	PublishTime time.Time
	Title       string
}

func (Video) TableName() string {
	return constants.VideosTableName
}

func CreateVideo(video *Video) (Video_id int64, err error) {
	err = DB.Create(video).Error
	if err != nil {
		return 0, err
	}
	return video.ID, err
}

func GetVideosByLastTime(lastTime time.Time) ([]*Video, error) {
	videos := make([]*Video, constants.VideoFeedCount)
	err := DB.Where("publish_time < ?", lastTime).Order("publish_time desc").Limit(constants.VideoFeedCount).Find(&videos).Error
	if err != nil {
		return videos, err
	}
	return videos, nil
}

func GetVideoByUserID(user_id int64) ([]*Video, error) {
	var videos []*Video
	err := DB.Where("author_id = ?", user_id).Find(&videos).Error
	if err != nil {
		return videos, err
	}
	return videos, err
}

func GetVideoListByVideoIDList(video_id_list []int64) ([]*Video, error) {
	var video_list []*Video
	var err error
	for _, item := range video_id_list {
		var video *Video
		err = DB.Where("id = ?", item).Find(&video).Error
		if err != nil {
			return video_list, err
		}
		video_list = append(video_list, video)
	}

	return video_list, err
}

func GetWorkCount(user_id int64) (int64, error) {
	var count int64
	err := DB.Model(&Video{}).Where("author_id = ?", user_id).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func CheckVideoExistById(video_id int64) (bool, error) {
	var video Video
	if err := DB.Where("id = ?", video_id).Find(&video).Error; err != nil {
		return false, err
	}
	if video == (Video{}) {
		return false, nil
	}
	return true, nil
}
