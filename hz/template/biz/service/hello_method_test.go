package service

import (
	"context"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	example "github.com/hertz/hello/biz/model/hello/example"
)

func TestHelloMethodService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewHelloMethodService(ctx, c)
	// init req and assert value
	req := &example.HelloReq{}
	resp, err := s.Run(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
