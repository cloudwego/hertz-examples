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

package relation

type DouyinRelationActionRequest struct {
	Token      string `protobuf:"bytes,1,req,name=token" json:"token,required" form:"token,required" query:"token,required"`                                          // 用户鉴权token
	ToUserId   int64  `protobuf:"varint,2,req,name=to_user_id,json=toUserId" json:"to_user_id,required" form:"to_user_id,required" query:"to_user_id,required"`       // 对方用户id
	ActionType int32  `protobuf:"varint,3,req,name=action_type,json=actionType" json:"action_type,required" form:"action_type,required" query:"action_type,required"` // 1-关注，2-取消关注
}

type DouyinRelationActionResponse struct {
	StatusCode int32  `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,required" form:"status_code,required" query:"status_code,required"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty" form:"status_msg" query:"status_msg"`                        // 返回状态描述
}

type DouyinRelationFollowListRequest struct {
	UserId int64  `protobuf:"varint,1,req,name=user_id,json=userId" json:"user_id,required" form:"user_id,required" query:"user_id,required"` // 用户id
	Token  string `protobuf:"bytes,2,req,name=token" json:"token,required" form:"token,required" query:"token,required"`                      // 用户鉴权token
}

type DouyinRelationFollowListResponse struct {
	StatusCode int32  `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,required" form:"status_code,required" query:"status_code,required"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty" form:"status_msg" query:"status_msg"`                        // 返回状态描述
	UserList   []User `protobuf:"bytes,3,rep,name=user_list,json=userList" json:"user_list" form:"user_list" query:"user_list"`                                       // 用户信息列表
}

type User struct {
	Id              int64  `protobuf:"varint,1,req,name=id" json:"id,required" form:"id,required" query:"id,required"`                                           // 用户id
	Name            string `protobuf:"bytes,2,req,name=name" json:"name,required" form:"name,required" query:"name,required"`                                    // 用户名称
	FollowCount     int64  `protobuf:"varint,3,opt,name=follow_count,json=followCount" json:"follow_count" form:"follow_count" query:"follow_count"`             // 关注总数
	FollowerCount   int64  `protobuf:"varint,4,opt,name=follower_count,json=followerCount" json:"follower_count" form:"follower_count" query:"follower_count"`   // 粉丝总数
	IsFollow        bool   `protobuf:"varint,5,req,name=is_follow,json=isFollow" json:"is_follow,required" form:"is_follow,required" query:"is_follow,required"` // true-已关注，false-未关注
	Avatar          string `json:"avatar" form:"avatar" query:"avatar"`
	BackgroundImage string `json:"background_image" form:"background_image" query:"background_image"`
	Signature       string `json:"signature" form:"signature" query:"signature"`
	TotalFavorited  int64  `json:"total_favorited" form:"total_favorited" query:"total_favorited"`
	WorkCount       int64  `json:"work_count" form:"work_count" query:"work_count"`
	FavoriteCount   int64  `json:"favorite_count" form:"favorite_count" query:"favorite_count"`
}

type DouyinRelationFollowerListRequest struct {
	UserId int64  `protobuf:"varint,1,req,name=user_id,json=userId" json:"user_id,required" form:"user_id,required" query:"user_id,required"` // 用户id
	Token  string `protobuf:"bytes,2,req,name=token" json:"token,required" form:"token,required" query:"token,required"`                      // 用户鉴权token
}

type DouyinRelationFollowerListResponse struct {
	StatusCode int32   `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,required" form:"status_code,required" query:"status_code,required"` // 状态码，0-成功，其他值-失败
	StatusMsg  string  `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty" form:"status_msg" query:"status_msg"`                        // 返回状态描述
	UserList   []*User `protobuf:"bytes,3,rep,name=user_list,json=userList" json:"user_list" form:"user_list" query:"user_list"`                                       // 用户列表
}

type DouyinRelationFriendListRequest struct {
	UserId int64  `protobuf:"varint,1,req,name=user_id,json=userId" json:"user_id,required" form:"user_id,required" query:"user_id,required"` // 用户id
	Token  string `protobuf:"bytes,2,req,name=token" json:"token,required" form:"token,required" query:"token,required"`                      // 用户鉴权token
}

type DouyinRelationFriendListResponse struct {
	StatusCode int32         `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,required" form:"status_code,required" query:"status_code,required"` // 状态码，0-成功，其他值-失败
	StatusMsg  string        `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty" form:"status_msg" query:"status_msg"`                        // 返回状态描述
	UserList   []*FriendUser `protobuf:"bytes,3,rep,name=user_list,json=userList" json:"user_list" form:"user_list" query:"user_list"`                                       // 用户列表
}

type FriendUser struct {
	Id              int64  `protobuf:"varint,1,req,name=id" json:"id,required" form:"id,required" query:"id,required"`                                           // 用户id
	Name            string `protobuf:"bytes,2,req,name=name" json:"name,required" form:"name,required" query:"name,required"`                                    // 用户名称
	FollowCount     int64  `protobuf:"varint,3,opt,name=follow_count,json=followCount" json:"follow_count" form:"follow_count" query:"follow_count"`             // 关注总数
	FollowerCount   int64  `protobuf:"varint,4,opt,name=follower_count,json=followerCount" json:"follower_count" form:"follower_count" query:"follower_count"`   // 粉丝总数
	IsFollow        bool   `protobuf:"varint,5,req,name=is_follow,json=isFollow" json:"is_follow,required" form:"is_follow,required" query:"is_follow,required"` // true-已关注，false-未关注
	Avatar          string `json:"avatar" form:"avatar" query:"avatar"`
	BackgroundImage string `json:"background_image" form:"background_image" query:"background_image"`
	Signature       string `json:"signature" form:"signature" query:"signature"`
	TotalFavorited  int64  `json:"total_favorited" form:"total_favorited" query:"total_favorited"`
	WorkCount       int64  `json:"work_count" form:"work_count" query:"work_count"`
	FavoriteCount   int64  `json:"favorite_count" form:"favorite_count" query:"favorite_count"`
	Message         string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty" form:"message" query:"message"`                   // 和该好友的最新聊天消息
	MsgType         int64  `protobuf:"varint,3,req,name=msgType" json:"msgType,required" form:"msgType,required" query:"msgType,required"` // message消息的类型，0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
}
