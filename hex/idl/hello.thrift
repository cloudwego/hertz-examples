// idl/hello.thrift
namespace go hello.example

struct HelloReq {
    1: string Name (api.query="name"); // 对应HTTP query参数
}

struct HelloResp {
    1: string RespBody;
}


service HelloService {
    HelloResp HelloMethod(1: HelloReq request) (api.get="/hello"); // 对应HTTP路由url
}