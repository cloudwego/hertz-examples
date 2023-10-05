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
	"time"
)

func TestGetVideoByLastTime(t *testing.T) {
	Init()
	lastTime := time.Now()
	videos, err := GetVideosByLastTime(lastTime)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	for _, video := range videos {
		fmt.Printf("%#v\n", video)
	}
}

func TestGetVideoByUserID(t *testing.T) {
	Init()
	user_id := int64(1000)
	videos, err := GetVideoByUserID(user_id)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	for _, video := range videos {
		fmt.Printf("%#v\n", video)
	}
}
