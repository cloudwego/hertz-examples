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
	"fmt"
	"log"
	"offer_tiktok/biz/dal/db"
	"offer_tiktok/pkg/constants"
	"offer_tiktok/pkg/utils"
	"sync"
	"time"

	feed "offer_tiktok/biz/model/basic/feed"
	user_service "offer_tiktok/biz/service/user"

	"github.com/cloudwego/hertz/pkg/app"
)

type FeedService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewFeedService(ctx context.Context, c *app.RequestContext) *FeedService {
	return &FeedService{ctx: ctx, c: c}
}

func (s *FeedService) Feed(req *feed.DouyinFeedRequest) (*feed.DouyinFeedResponse, error) {
	resp := &feed.DouyinFeedResponse{}
	var lastTime time.Time
	if req.LatestTime == 0 {
		lastTime = time.Now()
	} else {
		lastTime = time.Unix(req.LatestTime/1000, 0)
	}
	fmt.Printf("LastTime: %v\n", lastTime)
	current_user_id, exists := s.c.Get("current_user_id")
	// 如果当前用户没有登陆，则将 current_user_id 赋值为 0
	if !exists {
		current_user_id = int64(0)
	}

	dbVideos, err := db.GetVideosByLastTime(lastTime)
	if err != nil {
		return resp, err
	}

	videos := make([]*feed.Video, 0, constants.VideoFeedCount)
	err = s.CopyVideos(&videos, &dbVideos, current_user_id.(int64))
	if err != nil {
		return resp, nil
	}
	resp.VideoList = videos
	if len(dbVideos) != 0 {
		resp.NextTime = dbVideos[len(dbVideos)-1].PublishTime.Unix()
	}
	return resp, nil
}

func (s *FeedService) CopyVideos(result *[]*feed.Video, data *[]*db.Video, userId int64) error {
	for _, item := range *data {
		video := s.createVideo(item, userId)
		*result = append(*result, video)
	}
	return nil
}

/**
 * @description: 将 db.Video 拼接成 feed.Video
 * @param {*db.Video} data
 * @param {int64} userId
 * @return {*}
 */
func (s *FeedService) createVideo(data *db.Video, userId int64) *feed.Video {
	video := &feed.Video{
		Id:       data.ID,
		PlayUrl:  utils.URLconvert(s.ctx, s.c, data.PlayURL),
		CoverUrl: utils.URLconvert(s.ctx, s.c, data.CoverURL),
		Title:    data.Title,
	}

	var wg sync.WaitGroup
	wg.Add(4)

	// 获取作者信息
	go func() {
		author, err := user_service.NewUserService(s.ctx, s.c).GetUserInfo(data.AuthorID, userId)
		if err != nil {
			log.Printf("GetUserInfo func error:" + err.Error())
		}
		video.Author = &feed.User{
			Id:              author.Id,
			Name:            author.Name,
			FollowCount:     author.FollowCount,
			FollowerCount:   author.FollowerCount,
			IsFollow:        author.IsFollow,
			Avatar:          author.Avatar,
			BackgroundImage: author.BackgroundImage,
			Signature:       author.Signature,
			TotalFavorited:  author.TotalFavorited,
			WorkCount:       author.WorkCount,
			FavoriteCount:   author.FavoriteCount,
		}

		wg.Done()
	}()

	// 获取视频点赞数量
	go func() {
		err := *new(error)
		video.FavoriteCount, err = db.GetFavoriteCount(data.ID)
		if err != nil {
			log.Printf("GetFavoriteCount func error:" + err.Error())
		}
		wg.Done()
	}()

	go func() {
		err := *new(error)
		video.CommentCount, err = db.GetCommentCountByVideoID(data.ID)
		if err != nil {
			log.Printf("GetCommentCountByVideoID func error:" + err.Error())
		}
		wg.Done()
	}()

	go func() {
		err := *new(error)
		video.IsFavorite, err = db.QueryFavoriteExist(data.ID, userId)
		if err != nil {
			log.Printf("QueryFavoriteExist func error:" + err.Error())
		}
		wg.Done()
	}()

	wg.Wait()
	return video
}
