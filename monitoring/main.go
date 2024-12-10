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
	prometheus "github.com/hertz-contrib/monitor-prometheus"
)

func main() {
	h := server.Default(
		server.WithHostPorts("127.0.0.1:8080"),
		server.WithTracer(
			prometheus.NewServerTracer(":9091", "/hertz",
				prometheus.WithEnableGoCollector(true), // enable go runtime metric collector
			),
		),
	)

	h.GET("/metricGet", func(ctx context.Context, c *app.RequestContext) {
		c.String(200, "hello get")
	})

	h.POST("/metricPost", func(ctx context.Context, c *app.RequestContext) {
		time.Sleep(100 * time.Millisecond)
		c.String(200, "hello post")
	})

	h.Spin()
}
