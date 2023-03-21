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

	"github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/biz/model/common"

	"github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/biz/dal/db"
	"github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/pkg/errno"

	user_service "github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/biz/service/user"

	"github.com/cloudwego/hertz/pkg/app"

	relation "github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/biz/model/social/relation"
)

type FriendListService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewFriendListService(ctx context.Context, c *app.RequestContext) *FriendListService {
	return &FriendListService{ctx: ctx, c: c}
}

func (s *FriendListService) GetFriendList(req *relation.DouyinRelationFriendListRequest) ([]*relation.FriendUser, error) {
	user_id := req.UserId
	current_user_id, _ := s.c.Get("current_user_id")

	if current_user_id.(int64) != user_id {
		return nil, errno.FriendListNoPermissionErr
	}

	var friendList []*relation.FriendUser

	followerIdList, err := db.GetFollowerIdList(user_id)
	if err != nil {
		return friendList, err
	}

	for _, id := range followerIdList {
		isFriend, err := db.QueryFollowExist(&db.Follows{UserId: user_id, FollowerId: id})
		if err != nil {
			return friendList, err
		}
		if isFriend {
			user_info, err := user_service.NewUserService(s.ctx, s.c).GetUserInfo(id, user_id)
			if err != nil {
				log.Printf("func error: GetFriendList -> GetUserInfo")
			}
			message, err := db.GetLatestMessageByIdPair(user_id, id)
			if err != nil {
				log.Printf("func error: GetFriendList -> GetLatestMessageByIdPair")
			}
			var msgType int64
			if message == nil { // No chat history
				msgType = 2
				message = &db.Messages{}
			} else if user_id == message.FromUserId {
				msgType = 1
			} else {
				msgType = 0
			}
			friendList = append(friendList, &relation.FriendUser{
				User: common.User{
					Id:              user_info.Id,
					Name:            user_info.Name,
					FollowCount:     user_info.FollowCount,
					FollowerCount:   user_info.FollowerCount,
					IsFollow:        user_info.IsFollow,
					Avatar:          user_info.Avatar,
					BackgroundImage: user_info.BackgroundImage,
					Signature:       user_info.Signature,
					TotalFavorited:  user_info.TotalFavorited,
					WorkCount:       user_info.WorkCount,
					FavoriteCount:   user_info.FavoriteCount,
				},
				Message: message.Content,
				MsgType: msgType,
			})
		}
	}

	return friendList, nil
}
