package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	example "github.com/hertz/hello/biz/model/hello/example"
)

type HelloMethodService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHelloMethodService(Context context.Context, RequestContext *app.RequestContext) *HelloMethodService {
	return &HelloMethodService{RequestContext: RequestContext, Context: Context}
}

func (h *HelloMethodService) Run(req *example.HelloReq) (resp *example.HelloResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
