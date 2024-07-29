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
	"log"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	hertzSentinel "github.com/hertz-contrib/opensergo/sentinel/adapter"
)

func initSentinel() {
	err := sentinel.InitDefault()
	if err != nil {
		log.Fatalf("Unexpected error: %+v", err)
	}
	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource:               "server_test",
			Threshold:              0.0,
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject,
			StatIntervalInMs:       1000,
		},
	})
	if err != nil {
		log.Fatalf("Unexpected error: %+v", err)
		return
	}
}

func main() {
	initSentinel()
	h := server.Default(server.WithHostPorts(":8081"))
	h.Use(hertzSentinel.SentinelServerMiddleware(
		// customize resource extractor if required
		// method_path by default
		hertzSentinel.WithServerResourceExtractor(func(ctx context.Context, c *app.RequestContext) string {
			return "server_test"
		}),
		// customize block fallback if required
		// abort with status 429 by default
		hertzSentinel.WithServerBlockFallback(func(ctx context.Context, c *app.RequestContext) {
			c.AbortWithStatusJSON(400, utils.H{
				"err":  "too many request; the quota used up",
				"code": 10222,
			})
		}),
	))
	h.GET("/server_test", func(ctx context.Context, c *app.RequestContext) {})
	h.Spin()
}
