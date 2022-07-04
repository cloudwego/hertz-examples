module github.com/cloudwego/hertz-examples

go 1.16

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/apache/thrift v0.13.0
	github.com/cloudwego/hertz v0.1.0
	github.com/cloudwego/kitex v0.3.1
	github.com/hertz-contrib/cors v0.0.0-20220601061225-50f4e582beaf
	github.com/hertz-contrib/monitor-prometheus v0.0.0-20220601062737-825a4fc4595f
	github.com/hertz-contrib/obs-opentelemetry/logging/logrus v0.0.0-00010101000000-000000000000
	github.com/hertz-contrib/obs-opentelemetry/provider v0.0.0-00010101000000-000000000000
	github.com/hertz-contrib/obs-opentelemetry/tracing v0.0.0-00010101000000-000000000000
	github.com/hertz-contrib/tracer v0.0.0-20220601062646-788b1565bdab
	github.com/kitex-contrib/obs-opentelemetry v0.0.0-20220616115444-37518030dbb3
	github.com/kitex-contrib/obs-opentelemetry/logging/logrus v0.0.0-20220616115444-37518030dbb3
	github.com/kitex-contrib/tracer-opentracing v0.0.3
	github.com/opentracing/opentracing-go v1.2.0
	github.com/stretchr/testify v1.8.0
	github.com/uber/jaeger-client-go v2.30.0+incompatible
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	go.opentelemetry.io/otel v1.7.0
	go.uber.org/atomic v1.9.0 // indirect
	google.golang.org/protobuf v1.28.0
)

replace (
	github.com/hertz-contrib/obs-opentelemetry/logging/logrus => github.com/CoderPoet/obs-opentelemetry-hertz/logging/logrus v0.0.0-20220704030955-4c6354e7f6a1
	github.com/hertz-contrib/obs-opentelemetry/provider => github.com/CoderPoet/obs-opentelemetry-hertz/provider v0.0.0-20220704030955-4c6354e7f6a1
	github.com/hertz-contrib/obs-opentelemetry/tracing => github.com/CoderPoet/obs-opentelemetry-hertz/tracing v0.0.0-20220704030955-4c6354e7f6a1
)
