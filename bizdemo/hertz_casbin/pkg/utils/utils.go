/*
 * Copyright 2022 CloudWeGo Authors
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
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/cloudwego/hertz-examples/bizdemo/hertz_session/pkg/consts"
	"github.com/cloudwego/hertz/pkg/app"
)

// MD5 use md5 to encrypt strings
func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// BuildMsg render message for the html page
func BuildMsg(msg string) string {
	return fmt.Sprintf("%v", msg)
}

// IsLogout if user already login then return false
func IsLogout(_ context.Context, c *app.RequestContext) bool {
	if string(c.Cookie(consts.HertzSession)) == "" {
		return true
	}
	return false
}
