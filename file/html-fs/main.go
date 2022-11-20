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
	"strings"
)

func main() {
	h := server.Default(server.WithHostPorts("127.0.0.1:8080"))
	h.LoadHTMLGlob("views/*")

	prefix := "/public"
	root := "./assets"
	fs := &app.FS{Root: root, PathRewrite: getPathRewriter(prefix, root)}
	h.StaticFS(prefix, fs)

	h.GET("/", func(c context.Context, ctx *app.RequestContext) {
		ctx.HTML(200, "index.html", nil)
	})
	h.Spin()
}

func getPathRewriter(prefix, root string) app.PathRewriteFunc {
	// For security, we want to restrict to the current work directory.
	if root == "" {
		root = "."
	}
	// Cannot have an empty prefix
	if prefix == "" {
		prefix = "/"
	}
	// Prefix always start with a '/' or '*'
	if prefix[0] != '/' {
		prefix = "/" + prefix
	}

	// Strip trailing slashes from the root path
	if len(root) > 0 && root[len(root)-1] == '/' {
		root = root[:len(root)-1]
	}
	// Is prefix a direct wildcard?
	isStar := prefix == "/*"
	// Is prefix a partial wildcard?
	if strings.Contains(prefix, "*") {
		isStar = true
		prefix = strings.Split(prefix, "*")[0]
		// Fix this later
	}
	prefixLen := len(prefix)
	if prefixLen > 1 && prefix[prefixLen-1:] == "/" {
		// /john/ -> /john
		prefixLen--
		prefix = prefix[:prefixLen]
	}
	return func(ctx *app.RequestContext) []byte {
		path := ctx.Path()
		if len(path) >= prefixLen {
			if isStar && string(path[0:prefixLen]) == prefix {
				path = append(path[0:0], '/')
			} else {
				path = path[prefixLen:]
				if len(path) == 0 || path[len(path)-1] != '/' {
					path = append(path, '/')
				}
			}
		}
		if len(path) > 0 && path[0] != '/' {
			path = append([]byte("/"), path...)
		}
		return path
	}
}
