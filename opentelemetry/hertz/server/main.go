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
	"os"

	"github.com/cloudwego/hertz-examples/opentelemetry/kitex/kitex_gen/api"
	"github.com/cloudwego/hertz-examples/opentelemetry/kitex/kitex_gen/api/echo"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	kclient "github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	hertzlogrus "github.com/hertz-contrib/obs-opentelemetry/logging/logrus"
	"github.com/hertz-contrib/obs-opentelemetry/provider"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	kitextracing "github.com/kitex-contrib/obs-opentelemetry/tracing"
)

func main() {
	hlog.SetLogger(hertzlogrus.NewLogger())
	hlog.SetLevel(hlog.LevelDebug)

	serviceName := "demo-hertz-server"

	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName),
		// Support setting ExportEndpoint via environment variables: OTEL_EXPORTER_OTLP_ENDPOINT
		provider.WithExportEndpoint("localhost:4317"),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	tracer, cfg := hertztracing.NewServerTracer()
	h := server.Default(tracer)
	h.Use(hertztracing.ServerMiddleware(cfg))

	demoKitexServerAddr, ok := os.LookupEnv("DEMO_KITEX_SERVER_ENDPOINT")
	if !ok {
		demoKitexServerAddr = "0.0.0.0:8181"
	}
	client, err := echo.NewClient(serviceName,
		kclient.WithHostPorts(demoKitexServerAddr),
		kclient.WithSuite(kitextracing.NewClientSuite()),
		kclient.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),
	)
	if err != nil {
		hlog.Fatal(err)
	}

	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		req := &api.Request{Message: "my request"}
		resp, err := client.Echo(c, req)
		if err != nil {
			hlog.Errorf(err.Error())
		}
		hlog.CtxDebugf(c, "message received successfully: %s", req.Message)
		ctx.JSON(consts.StatusOK, resp)
	})

	h.Spin()
}
