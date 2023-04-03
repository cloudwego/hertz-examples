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
	"sync"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/biz/dal/db"
	user "github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/biz/model/basic/user"
	"github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/biz/model/common"
	"github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/pkg/constants"
	"github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/pkg/errno"
	"github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/pkg/utils"
)

type UserService struct {
	ctx context.Context
	c   *app.RequestContext
}

// NewUserService create user service
func NewUserService(ctx context.Context, c *app.RequestContext) *UserService {
	return &UserService{ctx: ctx, c: c}
}

// UserRegister register user return user id.
func (s *UserService) UserRegister(req *user.DouyinUserRegisterRequest) (user_id int64, err error) {
	user, err := db.QueryUser(req.Username)
	if err != nil {
		return 0, err
	}
	if *user != (db.User{}) {
		return 0, errno.UserAlreadyExistErr
	}

	passWord, err := utils.Crypt(req.Password)
	user_id, err = db.CreateUser(&db.User{
		UserName:        req.Username,
		Password:        passWord,
		Avatar:          constants.TestAva,
		BackgroundImage: constants.TestBackground,
		Signature:       constants.TestSign,
	})
	return user_id, nil
}

// UserInfo the function of user api
func (s *UserService) UserInfo(req *user.DouyinUserRequest) (*common.User, error) {
	query_user_id := req.UserId
	current_user_id, exists := s.c.Get("current_user_id")
	if !exists {
		current_user_id = 0
	}
	return s.GetUserInfo(query_user_id, current_user_id.(int64))
}

// GetUserInfo
//
//	@Description: Query the information of query_user_id according to the current user user_id
//	@receiver *UserService
//	@param query_user_id int64
//	@param user_id int64  "Currently logged-in user id, may be 0"
//	@return *user.User
//	@return error
func (s *UserService) GetUserInfo(query_user_id, user_id int64) (*common.User, error) {
	u := &common.User{}
	errChan := make(chan error, 7)
	defer close(errChan)
	var wg sync.WaitGroup
	wg.Add(7)
	go func() {
		dbUser, err := db.QueryUserById(query_user_id)
		if err != nil {
			errChan <- err
		} else {
			u.Name = dbUser.UserName
			u.Avatar = utils.URLconvert(s.ctx, s.c, dbUser.Avatar)
			u.BackgroundImage = utils.URLconvert(s.ctx, s.c, dbUser.BackgroundImage)
			u.Signature = dbUser.Signature
		}
		wg.Done()
	}()

	go func() {
		WorkCount, err := db.GetWorkCount(query_user_id)
		if err != nil {
			errChan <- err
		} else {
			u.WorkCount = WorkCount
		}
		wg.Done()
	}()

	go func() {
		FollowCount, err := db.GetFollowCount(query_user_id)
		if err != nil {
			errChan <- err
			return
		} else {
			u.FollowCount = FollowCount
		}
		wg.Done()
	}()

	go func() {
		FollowerCount, err := db.GetFollowerCount(query_user_id)
		if err != nil {
			errChan <- err
		} else {
			u.FollowerCount = FollowerCount
		}
		wg.Done()
	}()

	go func() {
		if user_id != 0 {
			IsFollow, err := db.QueryFollowExist(user_id, query_user_id)
			if err != nil {
				errChan <- err
			} else {
				u.IsFollow = IsFollow
			}
		} else {
			u.IsFollow = false
		}
		wg.Done()
	}()

	go func() {
		FavoriteCount, err := db.GetFavoriteCountByUserID(query_user_id)
		if err != nil {
			errChan <- err
		} else {
			u.FavoriteCount = FavoriteCount
		}
		wg.Done()
	}()

	go func() {
		TotalFavorited, err := db.QueryTotalFavoritedByAuthorID(query_user_id)
		if err != nil {
			errChan <- err
		} else {
			u.TotalFavorited = TotalFavorited
		}
		wg.Done()
	}()

	wg.Wait()
	select {
	case result := <-errChan:
		return &common.User{}, result
	default:
	}
	u.Id = query_user_id
	return u, nil
}
