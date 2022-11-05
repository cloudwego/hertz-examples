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

package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	hertzlogrus "github.com/hertz-contrib/logger/logrus"
	"github.com/hertz-contrib/requestid"
	"github.com/sirupsen/logrus"
)

type RequestIdHook struct{}

func (h *RequestIdHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *RequestIdHook) Fire(e *logrus.Entry) error {
	ctx := e.Context
	if ctx == nil {
		return nil
	}
	value := ctx.Value("X-Request-ID")
	if value != nil {
		e.Data["log_id"] = value
	}
	return nil
}

func main() {
	h := server.Default()
	logger := hertzlogrus.NewLogger(hertzlogrus.WithHook(&RequestIdHook{}))
	hlog.SetLogger(logger)

	h.Use(requestid.New())

	// Example ping request.
	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		hlog.CtxInfof(ctx, "test log")
		c.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	})
	h.Spin()
}
