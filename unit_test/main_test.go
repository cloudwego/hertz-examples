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
	"bytes"
	"context"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/ut"
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/stretchr/testify/assert"
)

func TestPerformRequest(t *testing.T) {
	opt := config.NewOptions([]config.Option{})
	router := route.NewEngine(opt)
	// context uri
	router.PUT("/hey/:version", func(ctx context.Context, c *app.RequestContext) {
		version := c.Param("version")
		if string(c.Request.Body()) == "1" {
			assert.Equal(t, "close", c.Request.Header.Get("Connection"))
			c.Response.SetConnectionClose()
			c.JSON(201, map[string]string{"hi": version})
		} else if string(c.Request.Body()) == "" {
			c.AbortWithMsg("unauthorized", 401)
		} else {
			assert.Equal(t, "PUT /hey/dy HTTP/1.1\r\nContent-Type: application/x-www-form-urlencoded\r\nTransfer-Encoding: chunked\r\n\r\n", string(c.Request.Header.Header()))
			c.String(202, "body:%v", string(c.Request.Body()))
		}
	})

	router.GET("/hey/header", func(ctx context.Context, c *app.RequestContext) {
		assert.Equal(t, "application/json", string(c.GetHeader("Content-Type")))
		assert.Equal(t, 1, c.Request.Header.ContentLength())
		assert.Equal(t, "a", c.Request.Header.Get("dummy"))
	})

	type testReq struct {
		Version string `json:"version"`
	}
	router.POST("/hey/json", func(ctx context.Context, c *app.RequestContext) {
		assert.Equal(t, "application/json", string(c.GetHeader("Content-Type")))
		var req testReq
		if err := c.BindAndValidate(&req); err != nil {
			panic(err)
		}
		assert.Equal(t, "v0.0.1", req.Version)
	})

	w := ut.PerformRequest(router, "PUT", "/hey/version1", &ut.Body{Body: bytes.NewBufferString("1"), Len: 1},
		ut.Header{Key: "Connection", Value: "close"})
	resp := w.Result()
	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, "{\"hi\":\"version1\"}", string(resp.Body()))
	assert.Equal(t, "application/json; charset=utf-8", string(resp.Header.ContentType()))
	assert.Equal(t, true, resp.Header.ConnectionClose())

	// unauthorized user
	w = ut.PerformRequest(router, "PUT", "/hey/version2", nil)
	_ = w.Result()
	resp = w.Result()
	assert.Equal(t, 401, resp.StatusCode())
	assert.Equal(t, "unauthorized", string(resp.Body()))
	assert.Equal(t, "text/plain; charset=utf-8", string(resp.Header.ContentType()))
	assert.Equal(t, 12, resp.Header.ContentLength())

	// special header
	ut.PerformRequest(router, "GET", "/hey/header", nil,
		ut.Header{Key: "content-type", Value: "application/json"},
		ut.Header{Key: "content-length", Value: "1"},
		ut.Header{Key: "dummy", Value: "a"},
		ut.Header{Key: "dummy", Value: "b"},
	)

	// not found
	w = ut.PerformRequest(router, "GET", "/hey", nil)
	resp = w.Result()
	assert.Equal(t, 404, resp.StatusCode())

	// fake body
	w = ut.PerformRequest(router, "GET", "/hey", nil)
	_, err := w.WriteString(", faker")
	resp = w.Result()
	assert.Nil(t, err)
	assert.Equal(t, 404, resp.StatusCode())
	assert.Equal(t, "Not Found, faker", string(resp.Body()))

	// json bind
	json := `{"version":"v0.0.1"}`
	w = ut.PerformRequest(router, "POST", "/hey/json",
		&ut.Body{Body: bytes.NewBufferString(json), Len: len(json)},
		ut.Header{Key: "Content-Type", Value: "application/json"},
	)
	resp = w.Result()
	assert.Equal(t, 200, resp.StatusCode())
}
