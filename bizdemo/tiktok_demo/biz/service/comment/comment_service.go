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

package service

import (
	"context"
	"log"
	"offer_tiktok/biz/dal/db"
	"offer_tiktok/biz/model/interact/comment"
	"offer_tiktok/pkg/errno"

	user_service "offer_tiktok/biz/service/user"

	"github.com/cloudwego/hertz/pkg/app"
)

type CommentService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewCommentService(ctx context.Context, c *app.RequestContext) *CommentService {
	return &CommentService{ctx: ctx, c: c}
}

func (c *CommentService) AddNewComment(req *comment.DouyinCommentActionRequest) (*comment.Comment, error) {
	current_user_id, _ := c.c.Get("current_user_id")
	video_id := req.VideoId
	action_type := req.ActionType
	comment_text := req.CommentText
	comment_id := req.CommentId
	comment := &comment.Comment{}
	// 发表评论
	if action_type == 1 {
		db_comment := &db.Comment{
			UserId:      current_user_id.(int64),
			VideoId:     video_id,
			CommentText: comment_text,
		}
		err := db.AddNewComment(db_comment)
		if err != nil {
			return comment, err
		}
		comment.Id = db_comment.ID
		comment.CreateDate = db_comment.CreatedAt.Format("01-02")
		comment.Content = db_comment.CommentText
		comment.User, err = c.getUserInfoById(current_user_id.(int64), current_user_id.(int64))
		if err != nil {
			return comment, err
		}
		return comment, nil
	} else {
		err := db.DeleteCommentById(comment_id)
		if err != nil {
			return comment, err
		}
		return comment, nil
	}
}

func (c *CommentService) getUserInfoById(current_user_id, user_id int64) (*comment.User, error) {
	u, err := user_service.NewUserService(c.ctx, c.c).GetUserInfo(user_id, current_user_id)
	var comment_user *comment.User
	if err != nil {
		return comment_user, err
	}
	comment_user = &comment.User{
		Id:              u.Id,
		Name:            u.Name,
		FollowCount:     u.FollowCount,
		FollowerCount:   u.FollowerCount,
		IsFollow:        u.IsFollow,
		Avatar:          u.Avatar,
		BackgroundImage: u.BackgroundImage,
		Signature:       u.Signature,
		TotalFavorited:  u.TotalFavorited,
		WorkCount:       u.WorkCount,
		FavoriteCount:   u.FavoriteCount,
	}
	return comment_user, nil
}

func (c *CommentService) CommentList(req *comment.DouyinCommentListRequest) (*comment.DouyinCommentListResponse, error) {
	resp := &comment.DouyinCommentListResponse{}
	video_id := req.VideoId

	// 获取current_user_id
	current_user_id, _ := c.c.Get("current_user_id")

	dbcomments, err := db.GetCommentListByVideoID(video_id)
	if err != nil {
		return resp, err
	}
	var comments []*comment.Comment
	err = c.copyComment(&comments, &dbcomments, current_user_id.(int64))
	if err != nil {
		return resp, err
	}
	resp.CommentList = comments
	resp.StatusMsg = errno.SuccessMsg
	resp.StatusCode = errno.SuccessCode
	return resp, nil
}

func (c *CommentService) copyComment(result *[]*comment.Comment, data *[]*db.Comment, current_user_id int64) error {
	for _, item := range *data {
		comment := c.createComment(item, current_user_id)
		*result = append(*result, comment)
	}
	return nil
}

/**
 * @description: 将 db.Comment 拼接成 comment.Comment
 * @param {*db.comment} data
 * @param {int64} userId
 * @return {*}
 */
func (c *CommentService) createComment(data *db.Comment, userId int64) *comment.Comment {
	comment := &comment.Comment{
		Id:         data.ID,
		Content:    data.CommentText,
		CreateDate: data.CreatedAt.Format("01-02"),
	}

	user_info, err := c.getUserInfoById(userId, data.UserId)
	if err != nil {
		log.Printf("func error")
	}
	comment.User = user_info
	return comment
}
