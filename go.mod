module github.com/cloudwego/hertz-examples

go 1.16

require (
	github.com/alibaba/sentinel-golang v1.0.4
	github.com/apache/thrift v0.16.0
	github.com/cloudwego/hertz v0.6.8
	github.com/cloudwego/kitex v0.6.1
	github.com/hertz-contrib/cors v0.0.0-20230423034624-2bc83a8400f0
	github.com/hertz-contrib/csrf v0.1.1
	github.com/hertz-contrib/gzip v0.0.1
	github.com/hertz-contrib/loadbalance v0.1.0
	github.com/hertz-contrib/logger/logrus v0.0.0-20221104075115-aecbfb39bbfe
	github.com/hertz-contrib/logger/zap v0.0.0-20221104075115-aecbfb39bbfe
	github.com/hertz-contrib/logger/zerolog v0.0.0-20221111024215-1ee59ae719d7
	github.com/hertz-contrib/monitor-prometheus v0.0.0-20220908085834-f3fe5f5e72ed
	github.com/hertz-contrib/obs-opentelemetry/logging/logrus v0.1.1
	github.com/hertz-contrib/obs-opentelemetry/provider v0.2.2
	github.com/hertz-contrib/obs-opentelemetry/tracing v0.2.2
	github.com/hertz-contrib/opensergo v0.0.1
	github.com/hertz-contrib/pprof v0.1.0
	github.com/hertz-contrib/registry/nacos v0.0.0-20221111034347-1885e5d5c1c9
	github.com/hertz-contrib/requestid v1.1.0
	github.com/hertz-contrib/reverseproxy v0.0.0-20220907134658-6a05798e1cc5
	github.com/hertz-contrib/sessions v1.0.2
	github.com/hertz-contrib/tracer v0.0.0-20220601062646-788b1565bdab
	github.com/kitex-contrib/obs-opentelemetry v0.2.3
	github.com/kitex-contrib/obs-opentelemetry/logging/logrus v0.0.0-20230530060140-c76e27f58391
	github.com/kitex-contrib/tracer-opentracing v0.0.3
	github.com/opentracing/opentracing-go v1.2.0
	github.com/sirupsen/logrus v1.9.2
	github.com/stretchr/testify v1.8.4
	github.com/uber/jaeger-client-go v2.30.0+incompatible
	go.opentelemetry.io/otel v1.16.0
	google.golang.org/protobuf v1.30.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.1764 // indirect
	github.com/benbjohnson/clock v1.3.0 // indirect
	github.com/bytedance/go-tagexpr/v2 v2.9.7 // indirect
	github.com/go-errors/errors v1.4.2 // indirect
	github.com/hertz-contrib/logger/slog v1.0.0
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/prometheus/client_golang v1.13.0 // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	golang.org/x/arch v0.3.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1
)

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0
