// Code generated by hertz generator.

package example

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	example "hertz-examples/hz/thrift/biz/model/hello/example"
)

// HelloMethod .
// @router /hello [GET]
func HelloMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req example.HelloReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := new(example.HelloResp)

	c.JSON(200, resp)
}
