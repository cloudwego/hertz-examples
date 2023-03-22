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
	"fmt"
	"testing"
)

func TestAddNewFavorite(t *testing.T) {
	Init()
	_, err := AddNewFavorite(&Favorites{
		UserId:  1000,
		VideoId: 115,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("succ")
}

func TestQueryTotalFavoritedByAuthorID(t *testing.T) {
	Init()
	author_id := 1000
	sum, err := QueryTotalFavoritedByAuthorID(int64(author_id))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sum)
}
