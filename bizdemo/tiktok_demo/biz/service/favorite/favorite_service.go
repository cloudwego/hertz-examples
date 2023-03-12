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
	"offer_tiktok/biz/dal/db"
	"offer_tiktok/biz/model/basic/feed"
	"offer_tiktok/pkg/constants"
	"offer_tiktok/pkg/errno"

	favorite "offer_tiktok/biz/model/interact/favorite"
	feed_service "offer_tiktok/biz/service/feed"

	"github.com/cloudwego/hertz/pkg/app"
)

type FavoriteService struct {
	ctx context.Context
	c   *app.RequestContext
}

// new FavoriteService
func NewFavoriteService(ctx context.Context, c *app.RequestContext) *FavoriteService {
	return &FavoriteService{ctx: ctx, c: c}
}

// like action, include like and unlike
// request parameters:
// string token = 1;       // 用户鉴权token
// int64 to_user_id = 2;   // 对方用户id
// int32 action_type = 3;  // 1-点赞，2-取消点赞
func (r *FavoriteService) FavoriteAction(req *favorite.DouyinFavoriteActionRequest) (flag bool, err error) {
	// 颁发和验证token的行为均交给jwt处理，当发送到handler层时，默认已通过验证
	// 只需要检查参数VideoID的合法性

	_, err = db.CheckVideoExistById(req.VideoId) // zheli
	if err != nil {
		return false, err
	}
	if req.ActionType != constants.FavoriteActionType && req.ActionType != constants.UnFavoriteActionType {
		return false, errno.ParamErr
	}
	// 获取current_user_id
	current_user_id, _ := r.c.Get("current_user_id")
	// // 不准自己关注自己
	// if req.ToUserId == current_user_id.(int64) {
	// 	return false, errno.ParamErr
	// }
	new_favorite_relation := &db.Favorites{
		UserId:  current_user_id.(int64),
		VideoId: req.VideoId,
	}
	// 请求参数校验完毕，检查favorite表中是否已经存在这两者的关系
	favorite_exist, _ := db.CheckFavoriteRelationExist(new_favorite_relation)
	if req.ActionType == constants.FavoriteActionType {
		if favorite_exist {
			return false, errno.FavoriteRelationAlreadyExistErr
		}
		flag, err = db.AddNewFavorite(new_favorite_relation)
	} else {
		if !favorite_exist {
			return false, errno.FavoriteRelationNotExistErr
		}
		flag, err = db.DeleteFavorite(new_favorite_relation)
	}
	return flag, err
}

// 获取用户点赞的所有视频列表，需要注意的是这里的token是客户端当前用户，而user_id可以是任意用户//zheli
// request parameters:
// string token;       // 用户鉴权token
// int64  user_id;     // 用户id
func (r *FavoriteService) GetFavoriteList(req *favorite.DouyinFavoriteListRequest) (favoritelist []*favorite.Video, err error) {
	query_user_id := req.UserId
	_, err = db.CheckUserExistById(query_user_id)

	if err != nil {
		return nil, err
	}
	// 获取current_user_id
	current_user_id, _ := r.c.Get("current_user_id")

	video_id_list, err := db.GetFavoriteIdList(query_user_id)

	dbVideos, err := db.GetVideoListByVideoIDList(video_id_list)
	var videos []*feed.Video
	f := feed_service.NewFeedService(r.ctx, r.c)
	err = f.CopyVideos(&videos, &dbVideos, current_user_id.(int64))
	for _, item := range videos {
		video := &favorite.Video{
			Id: item.Id,
			Author: favorite.User{
				Id:              item.Author.Id,
				Name:            item.Author.Name,
				FollowCount:     item.Author.FollowCount,
				FollowerCount:   item.Author.FollowerCount,
				Avatar:          item.Author.Avatar,
				BackgroundImage: item.Author.BackgroundImage,
				Signature:       item.Author.Signature,
				TotalFavorited:  item.Author.TotalFavorited,
				WorkCount:       item.Author.WorkCount,
			},
			PlayUrl:       item.PlayUrl,
			CoverUrl:      item.CoverUrl,
			FavoriteCount: item.FavoriteCount,
			CommentCount:  item.CommentCount,
			IsFavorite:    item.IsFavorite,
			Title:         item.Title,
		}
		favoritelist = append(favoritelist, video)
	}
	return favoritelist, err
}
