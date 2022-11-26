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

package hertz_session

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	hutils "github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/sessions"
	"hertz-examples/bizdemo/hertz_session/biz/consts"
)

// Page user page handler
func Page(_ context.Context, c *app.RequestContext) {
	session := sessions.Default(c)
	username := session.Get(consts.Username)
	if username == nil {
		c.JSON(http.StatusOK, hutils.H{
			"code":    http.StatusBadRequest,
			"message": consts.PageErr,
		})
		return
	}
	c.JSON(http.StatusOK, hutils.H{
		"code":    http.StatusOK,
		"message": fmt.Sprintf("Welcome, %v!", username.(string)),
	})
}
