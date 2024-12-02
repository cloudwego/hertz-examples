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
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/stretchr/testify/assert"
)

func TestRateLimitMiddleware(t *testing.T) {
	h := server.Default(server.WithHostPorts("127.0.0.1:8888"))
	h.Use(rateLimitMiddleware(100)) // set rate limit to 100KB/s

	h.GET("/rateLimit", func(ctx context.Context, c *app.RequestContext) {
		// Simulate a large response
		data := make([]byte, 200*1024) // 200KB of data
		c.Data(consts.StatusOK, "application/octet-stream", data)
	})

	go h.Spin()

	time.Sleep(1 * time.Second) // Wait for the server to start

	startTime := time.Now()
	resp, err := http.Get("http://127.0.0.1:8888/rateLimit")
	assert.NoError(t, err)
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)

	elapsedTime := time.Since(startTime)
	expectedTime := 2 * time.Second // 200KB at 100KB/s should take 2 seconds

	assert.GreaterOrEqual(t, elapsedTime, expectedTime, "Response time should be at least 2 seconds")
}
