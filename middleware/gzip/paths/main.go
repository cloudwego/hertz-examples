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
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/gzip"
)

func main() {
	h := server.Default(server.WithHostPorts(":8080"))
	h.Use(
		gzip.Gzip(
			gzip.DefaultCompression,
			// This WithExcludedPaths takes as its parameter the file path
			gzip.WithExcludedPaths([]string{"/api/"}),
		),
	)
	// This is before compression
	h.GET("/api/book", func(ctx context.Context, c *app.RequestContext) {
		c.String(http.StatusOK, "This is a book!")
	})
	// This is the compressed
	h.GET("/book", func(ctx context.Context, c *app.RequestContext) {
		c.String(http.StatusOK, "This is a book!")
	})
	h.Spin()
}
