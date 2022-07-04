package main

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzlogrus "github.com/hertz-contrib/obs-opentelemetry/logging/logrus"
	"github.com/hertz-contrib/obs-opentelemetry/provider"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

func main() {
	hlog.SetLogger(hertzlogrus.NewLogger())
	hlog.SetLevel(hlog.LevelDebug)

	serviceName := "demo-hertz-client"

	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName),
		// Support setting ExportEndpoint via environment variables: OTEL_EXPORTER_OTLP_ENDPOINT
		provider.WithExportEndpoint("host.docker.internal:4317"),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	c, _ := client.NewClient()
	c.Use(hertztracing.ClientMiddleware())

	for {
		ctx, span := otel.Tracer("github.com/hertz-contrib/obs-opentelemetry").
			Start(context.Background(), "loop")
		defer span.End()

		_, b, err := c.Get(ctx, nil, "http://0.0.0.0:8888/ping?foo=bar")
		if err != nil {
			hlog.CtxErrorf(ctx, err.Error())
		}

		span.SetAttributes(attribute.String("msg", string(b)))

		hlog.CtxInfof(ctx, "hertz client %s", string(b))

		<-time.After(time.Second)
	}

}
