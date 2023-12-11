// Code generated by hertz generator.

package example

import (
	"context"

	example "a/b/c/kitex_gen/hello/example"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// HelloMethod .
// @router /hello [GET]
func HelloMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req example.HelloReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(example.HelloResp)

	c.JSON(consts.StatusOK, resp)
}
