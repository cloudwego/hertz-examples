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
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/reverseproxy"
	"github.com/hertz-contrib/sse"
	"net/http"
	"time"
)

var proxy *reverseproxy.ReverseProxy

func main() {
	proxy, _ = reverseproxy.NewSingleHostReverseProxy("http://127.0.0.1:8001",
		client.WithResponseBodyStream(true),
	)

	h := server.Default(server.WithHostPorts("127.0.0.1:8080"))
	h.GET("/proxy", proxyToServer2)
	h.GET("/proxy/:path", proxyToServer2)

	h2 := server.Default(server.WithHostPorts("127.0.0.1:8001"))
	h2.GET("/sse", sseHandler)

	go h.Spin()
	h2.Spin()
}

func proxyToServer2(ctx context.Context, c *app.RequestContext) {
	path := c.Param("path")
	c.Request.Header.Set("Secret", "123456")
	c.Request.URI().SetPath(path)
	proxy.ServeHTTP(ctx, c)
	c.Abort()
}

func sseHandler(ctx context.Context, c *app.RequestContext) {
	// you must set status code and response headers before first render call
	c.Response.Header.Set("X-Accel-Buffering", "no")
	c.SetStatusCode(http.StatusOK)
	s := sse.NewStream(c)
	for i := 0; i < 10; i++ {
		event := &sse.Event{
			Event: "timestamp",
			Data:  []byte(time.Now().Format(time.RFC3339)),
		}
		err := s.Publish(event)
		if err != nil {
			return
		}
		time.Sleep(200 * time.Millisecond)
	}
}
