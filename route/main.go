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
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/basic_auth"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
	h := server.Default(server.WithHostPorts("127.0.0.1:8080"))

	// register static route
	RegisterRoute(h)

	// register route group
	RegisterGroupRoute(h)

	// register use middleware with route group
	RegisterGroupRouteWithMiddleware(h)

	// register parameter route
	RegisterParaRoute(h)

	// register use anonymous function or decorator to register routes
	RegisterAnonFunOrDecRoute(h)

	// register Get route info
	RegisterGetRoutesInfo(h)

	h.Spin()
}

// RegisterRoute static route
func RegisterRoute(h *server.Hertz) {
	h.StaticFS("/", &app.FS{Root: "./", GenerateIndexPages: true})

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
	h.Any("/ping_any", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "any")
	})
	h.Handle("LOAD", "/load", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "load")
	})
}

// RegisterGroupRoute group route
func RegisterGroupRoute(h *server.Hertz) {
	// Simple group: v1
	v1 := h.Group("/v1")
	{
		// loginEndpoint is a handler func
		v1.GET("/get", func(ctx context.Context, c *app.RequestContext) {
			c.String(consts.StatusOK, "get")
		})
		v1.POST("/post", func(ctx context.Context, c *app.RequestContext) {
			c.String(consts.StatusOK, "post")
		})
	}

	// Simple group: v2
	v2 := h.Group("/v2")
	{
		v2.PUT("/put", func(ctx context.Context, c *app.RequestContext) {
			c.String(consts.StatusOK, "put")
		})
		v2.DELETE("/delete", func(ctx context.Context, c *app.RequestContext) {
			c.String(consts.StatusOK, "delete")
		})
	}
}

// RegisterGroupRouteWithMiddleware route groups that incorporate middleware
func RegisterGroupRouteWithMiddleware(h *server.Hertz) {
	// The following example uses the BasicAuth middleware in a route group.

	// Sample Code 1:
	//
	// Bind the middleware directly to the routing group
	example1 := h.Group("/example1", basic_auth.BasicAuth(map[string]string{"test": "test"}))
	example1.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		ctx.String(consts.StatusOK, "ping")
	})

	// Sample Code 2:
	//
	// use `Use` method
	example2 := h.Group("/example2")
	example2.Use(basic_auth.BasicAuth(map[string]string{"test": "test"}))
	example2.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		ctx.String(consts.StatusOK, "ping")
	})
}

// RegisterParaRoute parameter route
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

// RegisterAnonFunOrDecRoute Use anonymous function or decorator to register routes
func RegisterAnonFunOrDecRoute(h *server.Hertz) {
	h.AnyEX("/ping", func(ctx context.Context, c *app.RequestContext) {
		ctx.String(consts.StatusOK, app.GetHandlerName(ctx.Handler()))
	}, "ping_handler")
}

// RegisterGetRoutesInfo Get route info
func RegisterGetRoutesInfo(h *server.Hertz) {
	h.GET("/getRoutesInfo", func(ctx context.Context, c *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	})
	routeInfo := h.Routes()
	hlog.Info(routeInfo)
}
