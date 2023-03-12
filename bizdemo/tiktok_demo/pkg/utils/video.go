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

package utils

import (
	"context"
	"fmt"
	"offer_tiktok/biz/mw/minio"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func NewFileName(user_id, time int64) string {
	return fmt.Sprintf("%d.%d", user_id, time)
}

// URLconvert
/**
 * @description: 将数据库中存放的url转换为前端可访问的完整url
 * @param {context.Context} ctx
 * @param {*app.RequestContext} c
 * @param {string} path
 * @return {string} fullURL
 */
func URLconvert(ctx context.Context, c *app.RequestContext, path string) (fullURL string) {
	if len(path) == 0 {
		return ""
	}
	arr := strings.Split(path, "/")
	u, err := minio.GetObjURL(ctx, arr[0], arr[1])
	if err != nil {
		hlog.CtxInfof(ctx, err.Error())
		return ""
	}
	u.Scheme = string(c.URI().Scheme())
	u.Host = string(c.URI().Host())
	u.Path = "/src" + u.Path
	return u.String()
}
