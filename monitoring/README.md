# Prometheus monitoring for Hertz

## Usage Example

### Server

```
import (
   "context"
	"time"

	"github.com/cloudwego/monitor-prometheus"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
...
	h := server.Default(server.WithHostPorts("127.0.0.1:8080"), server.WithTracer(prometheus.NewServerTracer(":9091", "/hertz")))

	h.GET("/metricGet", func(ctx context.Context, c *app.RequestContext) {
		ctx.String(200, "hello get")
	})

	h.POST("/metricPost", func(ctx context.Context, c *app.RequestContext) {
		time.Sleep(100 * time.Millisecond)
		ctx.String(200, "hello post")
	})

	h.Spin()
```
## HOW-TO-RUN
1. install docker and start docker
2. change $inetIP to local ip in line 30 of prometheus.yml
3. run Prometheus and Grafana  
   `docker-compose up`
4. run Hertz server   
   `go run main.go`
5. run Hertz client  
   `go run client/main.go`
6. visit `http://localhost:3000`, the account password is `admin` by default
7. configure Prometheus data sources
   1. `Configuration`
   2. `Data Source`
   3. `Add data source`
   4. Select `Prometheus` and fill the URL with `http://prometheus:9090` 
   5. click `Save & Test` after configuration to test if it works
8. add dashboard `Create` -> `dashboard`, add monitoring metrics such as throughput and pct99 according to your needs, for example:  
   
   - server throughput of succeed requests

   `sum(rate(hertz_server_throughput{statusCode="200"}[1m])) by (method)`
   
   - server latency pct99 of succeed requests

   `histogram_quantile(0.9,sum(rate(hertz_server_latency_us_bucket{statusCode="200"}[1m]))by(le))`

For more information about hertz monitoring, please click [monitoring](https://www.cloudwego.io/zh/docs/hertz/tutorials/framework-exten/monitor/)
