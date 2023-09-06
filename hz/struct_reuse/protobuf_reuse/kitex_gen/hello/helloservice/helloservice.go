// Code generated by Kitex v0.7.1. DO NOT EDIT.

package helloservice

import (
	hello "a/b/c/kitex_gen/hello"
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return helloServiceServiceInfo
}

var helloServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "HelloService"
	handlerType := (*hello.HelloService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Method1": kitex.NewMethodInfo(method1Handler, newMethod1Args, newMethod1Result, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "hello",
		"ServiceFilePath": "",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.7.1",
		Extra:           extra,
	}
	return svcInfo
}

func method1Handler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(hello.HelloReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(hello.HelloService).Method1(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *Method1Args:
		success, err := handler.(hello.HelloService).Method1(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*Method1Result)
		realResult.Success = success
	}
	return nil
}
func newMethod1Args() interface{} {
	return &Method1Args{}
}

func newMethod1Result() interface{} {
	return &Method1Result{}
}

type Method1Args struct {
	Req *hello.HelloReq
}

func (p *Method1Args) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(hello.HelloReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *Method1Args) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *Method1Args) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *Method1Args) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *Method1Args) Unmarshal(in []byte) error {
	msg := new(hello.HelloReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var Method1Args_Req_DEFAULT *hello.HelloReq

func (p *Method1Args) GetReq() *hello.HelloReq {
	if !p.IsSetReq() {
		return Method1Args_Req_DEFAULT
	}
	return p.Req
}

func (p *Method1Args) IsSetReq() bool {
	return p.Req != nil
}

func (p *Method1Args) GetFirstArgument() interface{} {
	return p.Req
}

type Method1Result struct {
	Success *hello.HelloResp
}

var Method1Result_Success_DEFAULT *hello.HelloResp

func (p *Method1Result) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(hello.HelloResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *Method1Result) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *Method1Result) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *Method1Result) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *Method1Result) Unmarshal(in []byte) error {
	msg := new(hello.HelloResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *Method1Result) GetSuccess() *hello.HelloResp {
	if !p.IsSetSuccess() {
		return Method1Result_Success_DEFAULT
	}
	return p.Success
}

func (p *Method1Result) SetSuccess(x interface{}) {
	p.Success = x.(*hello.HelloResp)
}

func (p *Method1Result) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *Method1Result) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Method1(ctx context.Context, Req *hello.HelloReq) (r *hello.HelloResp, err error) {
	var _args Method1Args
	_args.Req = Req
	var _result Method1Result
	if err = p.c.Call(ctx, "Method1", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
