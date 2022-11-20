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
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
	h := server.Default(server.WithHostPorts("127.0.0.1:8080"))

	// register route
	RegisterRoute(h)

	// register route with handle
	RegisterRouteWithHandle(h)

	// register group route
	RegisterGroupRoute(h)

	// register parameter route
	RegisterParaRoute(h)

	h.Spin()
}

func RegisterRoute(h *server.Hertz) {
	h.GET("/get", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "get")
	})
	h.POST("/post", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "post")
	})
	h.PUT("/put", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "put")
	})
	h.DELETE("/delete", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "delete")
	})
	h.PATCH("/patch", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "patch")
	})
	h.HEAD("/head", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "head")
	})
	h.OPTIONS("/options", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "options")
	})
}

func RegisterRouteWithHandle(h *server.Hertz) {
	h.Handle(consts.MethodGet, "/hget", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "hget")
	})
	h.Handle(consts.MethodPost, "/hpost", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "hpost")
	})
	h.Handle(consts.MethodPut, "/hput", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "hput")
	})
	h.Handle(consts.MethodDelete, "/hdelete", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "hdelete")
	})
	h.Handle(consts.MethodPatch, "/hpatch", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "hpatch")
	})
	h.Handle(consts.MethodHead, "/hhead", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "hhead")
	})
	h.Handle(consts.MethodOptions, "/hoptions", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "hoptions")
	})
}

func loginEndpoint(ctx context.Context, c *app.RequestContext) {}

func submitEndpoint(ctx context.Context, c *app.RequestContext) {}

func readEndpoint(ctx context.Context, c *app.RequestContext) {}

func RegisterGroupRoute(h *server.Hertz) {
	// Simple group: v1
	v1 := h.Group("/v1")
	{
		// loginEndpoint is a handler func
		v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/streaming_read", readEndpoint)
	}

	// Simple group: v2
	v2 := h.Group("/v2")
	{
		v2.POST("/login", loginEndpoint)
		v2.POST("/submit", submitEndpoint)
		v2.POST("/streaming_read", readEndpoint)
	}
}

func RegisterParaRoute(h *server.Hertz) {
	// This handler will match: "/hertz/version", but will not match : "/hertz/" or "/hertz"
	h.GET("/hertz/:version", func(ctx context.Context, c *app.RequestContext) {
		version := c.Param("version")
		c.String(consts.StatusOK, "Hello %s", version)
	})

	// However, this one will match "/hertz/v1/" and "/hertz/v2/send"
	h.GET("/hertz/:version/*action", func(ctx context.Context, c *app.RequestContext) {
		version := c.Param("version")
		action := c.Param("action")
		message := version + " is " + action
		c.String(consts.StatusOK, message)
	})

	// For each matched request Context will hold the route definition
	h.POST("/hertz/:version/*action", func(ctx context.Context, c *app.RequestContext) {
		// c.FullPath() == "/hertz/:version/*action" // true
		c.String(consts.StatusOK, c.FullPath())
	})
}
