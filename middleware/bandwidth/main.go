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
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func rateLimitMiddleware(rateLimitKBps int) app.HandlerFunc {
	bytesPerSecond := rateLimitKBps * 1024
	return func(ctx context.Context, c *app.RequestContext) {
		startTime := time.Now()
		c.Next(ctx)
		elapsedTime := time.Since(startTime)
		expectedTime := time.Duration(len(c.Response.Body()) * int(time.Second) / bytesPerSecond)
		if elapsedTime < expectedTime {
			time.Sleep(expectedTime - elapsedTime)
		}
	}
}

func main() {
	h := server.Default(server.WithHostPorts("127.0.0.1:8888"))
	h.Use(rateLimitMiddleware(100)) // set rate limit to 100KB/s

	h.GET("/rateLimit", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "hello hertz")
	})

	h.Spin()
}
