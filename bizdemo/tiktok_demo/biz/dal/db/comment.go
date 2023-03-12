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
	"offer_tiktok/pkg/errno"
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID          int64          `json:"id"`
	UserId      int64          `json:"user_id"`
	VideoId     int64          `json:"video_id"`
	CommentText string         `json:"comment_text"`
	CreatedAt   time.Time      `json:"created_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (Comment) TableName() string {
	return constants.CommentTableName
}

func AddNewComment(comment *Comment) error {
	if ok, _ := CheckUserExistById(comment.UserId); !ok {
		return errno.UserIsNotExistErr
	}
	if ok, _ := CheckVideoExistById(comment.VideoId); !ok {
		return errno.VideoIsNotExistErr
	}
	err := DB.Create(comment).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteCommentById(comment_id int64) error {
	if ok, _ := CheckCommentExist(comment_id); !ok {
		return errno.CommentIsNotExistErr
	}
	comment := &Comment{}
	err := DB.Where("id = ?", comment_id).Delete(comment).Error
	if err != nil {
		return err
	}
	return nil
}

func CheckCommentExist(comment_id int64) (bool, error) {
	comment := &Comment{}
	err := DB.Where("id = ?", comment_id).Find(comment).Error
	if err != nil {
		return false, err
	}
	if comment.ID == 0 {
		return false, nil
	}
	return true, nil
}

func GetCommentListByVideoID(video_id int64) ([]*Comment, error) {
	var comment_list []*Comment
	if ok, _ := CheckVideoExistById(video_id); !ok {
		return comment_list, errno.VideoIsNotExistErr
	}
	err := DB.Table(constants.CommentTableName).Where("video_id = ?", video_id).Find(&comment_list).Error
	if err != nil {
		return comment_list, err
	}
	return comment_list, nil
}

func GetCommentCountByVideoID(video_id int64) (int64, error) {
	var sum int64
	err := DB.Model(&Comment{}).Where("video_id = ?", video_id).Count(&sum).Error
	if err != nil {
		return sum, err
	}
	return sum, nil
}
