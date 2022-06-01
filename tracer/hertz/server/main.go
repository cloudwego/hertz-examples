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
	"fmt"
	"io"

	"github.com/cloudwego/hertz-examples/tracer/kitex/kitex_gen/api"
	"github.com/cloudwego/hertz-examples/tracer/kitex/kitex_gen/api/echo"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	kclient "github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/hertz-contrib/tracer/hertz"
	kopentracing "github.com/kitex-contrib/tracer-opentracing"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

/*
export JAEGER_DISABLED=false
export JAEGER_SAMPLER_TYPE="const"
export JAEGER_SAMPLER_PARAM=1
export JAEGER_REPORTER_LOG_SPANS=true
export JAEGER_AGENT_HOST="127.0.0.1"
export JAEGER_AGENT_PORT=6831
*/

func InitTracer(serviceName string) (opentracing.Tracer, io.Closer) {
	cfg, _ := jaegercfg.FromEnv()
	cfg.ServiceName = serviceName
	tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	// opentracing.InitGlobalTracer(tracer)
	return tracer, closer
}

func main() {
	ht, hc := InitTracer("hertz-server")
	kt, kc := InitTracer("kitex-client")
	defer hc.Close()
	defer kc.Close()

	client, err := echo.NewClient("echo",
		kclient.WithHostPorts("0.0.0.0:5555"),
		kclient.WithSuite(kopentracing.NewClientSuite(kt, func(c context.Context) string {
			endpoint := rpcinfo.GetRPCInfo(c).From()
			return endpoint.ServiceName() + "::" + endpoint.Method()
		})))
	if err != nil {
		panic(err)
	}

	h := server.Default(server.WithTracer(hertz.NewTracer(ht, func(c *app.RequestContext) string {
		return "test.hertz.server" + "::" + c.FullPath()
	})))
	h.Use(hertz.ServerCtx())

	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		req := &api.Request{Message: "my request"}
		resp, err := client.Echo(c, req)
		if err != nil {
			hlog.Errorf(err.Error())
		}
		ctx.JSON(consts.StatusOK, resp)
	})

	h.Spin()
}
