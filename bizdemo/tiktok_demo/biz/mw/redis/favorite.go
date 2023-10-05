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

package redis

import (
	"strconv"
)

const (
	likeSuffix  = ":like"
	likedSuffix = ":liked"
)

type (
	Favorite struct{}
)

func (f Favorite) AddLike(user_id, video_id int64) {
	add(rdbFavorite, strconv.FormatInt(user_id, 10)+likeSuffix, video_id)
}

func (f Favorite) AddLiked(user_id, video_id int64) {
	add(rdbFavorite, strconv.FormatInt(video_id, 10)+likedSuffix, user_id)
}

func (f Favorite) DelLike(user_id, video_id int64) {
	del(rdbFavorite, strconv.FormatInt(user_id, 10)+likeSuffix, video_id)
}

func (f Favorite) DelLiked(user_id, video_id int64) {
	del(rdbFavorite, strconv.FormatInt(video_id, 10)+likedSuffix, user_id)
}

func (f Favorite) CheckLike(user_id int64) bool {
	return check(rdbFavorite, strconv.FormatInt(user_id, 10)+likeSuffix)
}

func (f Favorite) CheckLiked(video_id int64) bool {
	return check(rdbFavorite, strconv.FormatInt(video_id, 10)+likedSuffix)
}

func (f Favorite) ExistLike(user_id, video_id int64) bool {
	return exist(rdbFavorite, strconv.FormatInt(user_id, 10)+likeSuffix, video_id)
}

func (f Favorite) ExistLiked(user_id, video_id int64) bool {
	return exist(rdbFavorite, strconv.FormatInt(video_id, 10)+likedSuffix, user_id)
}

func (f Favorite) CountLike(user_id int64) (int64, error) {
	return count(rdbFavorite, strconv.FormatInt(user_id, 10)+likeSuffix)
}

func (f Favorite) CountLiked(video_id int64) (int64, error) {
	return count(rdbFavorite, strconv.FormatInt(video_id, 10)+likedSuffix)
}

func (f Favorite) GetLike(user_id int64) []int64 {
	return get(rdbFavorite, strconv.FormatInt(user_id, 10)+likeSuffix)
}

func (f Favorite) GetLiked(video_id int64) []int64 {
	return get(rdbFavorite, strconv.FormatInt(video_id, 10)+likedSuffix)
}
