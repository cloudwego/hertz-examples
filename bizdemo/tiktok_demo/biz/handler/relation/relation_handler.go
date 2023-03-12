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

// Code generated by hertz generator.

package relation

import (
	"context"
	"offer_tiktok/biz/pack"
	"offer_tiktok/pkg/errno"

	relation "offer_tiktok/biz/model/social/relation"

	follow_service "offer_tiktok/biz/service/relation/follow"
	followerList_service "offer_tiktok/biz/service/relation/follower"
	friendList_service "offer_tiktok/biz/service/relation/friend"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// RelationAction .
// @router /douyin/relation/action/ [POST]
func RelationAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req relation.DouyinRelationActionRequest
	err = c.BindAndValidate(&req)
	// hlog.CtxInfof(ctx, "RelationAction: usr_token: %s follower_id: %d action_type: %d", req.Token, req.ToUserId, req.ActionType)
	if err != nil {
		resp := pack.BuildBaseResp(err)
		c.JSON(consts.StatusOK, relation.DouyinRelationActionResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	_, err = follow_service.NewRelationService(ctx, c).FollowAction(&req)
	if err != nil {
		resp := pack.BuildBaseResp(err)
		c.JSON(consts.StatusOK, relation.DouyinRelationActionResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	c.JSON(consts.StatusOK, relation.DouyinRelationActionResponse{
		StatusCode: errno.SuccessCode,
		StatusMsg:  errno.SuccessMsg,
	})
}

// RelationFollowList .
// @router /douyin/relation/follow/list/ [GET]
func RelationFollowList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req relation.DouyinRelationFollowListRequest
	err = c.BindAndValidate(&req)

	// hlog.CtxInfof(ctx, "RelationGetFollowList: usr_id: %d user_token: %s", req.UserId, req.Token)
	if err != nil {
		resp := pack.BuildBaseResp(err)
		c.JSON(consts.StatusOK, relation.DouyinRelationFollowListResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
			UserList:   nil,
		})
		return
	}

	FollowInfo, err := follow_service.NewRelationService(ctx, c).GetFollowList(&req)
	if err != nil {
		resp := pack.BuildBaseResp(err)
		c.JSON(consts.StatusOK, relation.DouyinRelationFollowListResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
			UserList:   nil,
		})
		return
	}

	c.JSON(consts.StatusOK, relation.DouyinRelationFollowListResponse{
		StatusCode: errno.SuccessCode,
		StatusMsg:  errno.SuccessMsg,
		UserList:   FollowInfo,
	})
}

// RelationFollowerList .
// @router /douyin/relation/follower/list/ [GET]
func RelationFollowerList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req relation.DouyinRelationFollowerListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := pack.BuildBaseResp(err)
		// c.String(consts.StatusBadRequest, err.Error())
		c.JSON(consts.StatusOK, relation.DouyinRelationFollowerListResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
			UserList:   nil,
		})
		return
	}

	followerList, err := followerList_service.NewFollowerListService(ctx, c).GetFollowerList(&req)
	if err != nil {
		resp := pack.BuildBaseResp(err)
		c.JSON(consts.StatusOK, relation.DouyinRelationFollowerListResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
			UserList:   nil,
		})
	} else {
		c.JSON(consts.StatusOK, relation.DouyinRelationFollowerListResponse{
			StatusCode: errno.SuccessCode,
			StatusMsg:  errno.SuccessMsg,
			UserList:   followerList,
		})
	}
}

// RelationFriendList .
// @router /douyin/relation/friend/list/ [GET]
func RelationFriendList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req relation.DouyinRelationFriendListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := pack.BuildBaseResp(err)
		// c.String(consts.StatusBadRequest, err.Error())
		c.JSON(consts.StatusOK, relation.DouyinRelationFriendListResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
			UserList:   nil,
		})
		return
	}

	friendList, err := friendList_service.NewFriendListService(ctx, c).GetFriendList(&req)
	if err != nil {
		resp := pack.BuildBaseResp(err)
		c.JSON(consts.StatusOK, relation.DouyinRelationFriendListResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
			UserList:   nil,
		})
	} else {
		c.JSON(consts.StatusOK, relation.DouyinRelationFriendListResponse{
			StatusCode: errno.SuccessCode,
			StatusMsg:  errno.SuccessMsg,
			UserList:   friendList,
		})
	}
}
