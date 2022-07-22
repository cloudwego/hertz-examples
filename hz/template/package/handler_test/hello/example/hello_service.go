// this is my custom handler.

package example

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	example "github.com/hertz/hello/biz/model/hello/example"
)

// HelloMethod .
// @router /hello [GET]
func HelloMethod(ctx context.Context, c *app.RequestContext) {
	//  you can code something
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

// HelloMethod2 .
// @router /hello2 [GET]
func HelloMethod2(ctx context.Context, c *app.RequestContext) {
	// this my demo
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
