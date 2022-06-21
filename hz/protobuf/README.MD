## 基于 protobuf IDL 创建项目

### new: 创建一个新项目

1. 在当前目录下创建 protobuf idl 文件

   *注*：为在 protobuf 中支持 api 注解，请在使用了注解的 proto 文件中，import 下面的文件

```
syntax = "proto2";

package api;

import "google/protobuf/descriptor.proto";

option go_package = "/api";

extend google.protobuf.FieldOptions {
  optional string raw_body = 50101;
  optional string query = 50102;
  optional string header = 50103;
  optional string cookie = 50104;
  optional string body = 50105;
  optional string path = 50106;
  optional string vd = 50107;
  optional string form = 50108;
  optional string go_tag = 51001;
  optional string js_conv = 50109;
}

extend google.protobuf.MethodOptions {
  optional string get = 50201;
  optional string post = 50202;
  optional string put = 50203;
  optional string delete = 50204;
  optional string patch = 50205;
  optional string options = 50206;
  optional string head = 50207;
  optional string any = 50208;
  optional string gen_path = 50301;
  optional string api_version = 50302;
  optional string tag = 50303;
  optional string name = 50304;
  optional string api_level = 50305;
  optional string serializer = 50306;
  optional string param = 50307;
  optional string baseurl = 50308;
}

extend google.protobuf.EnumValueOptions {
  optional int32 http_code = 50401;
}
```

主 idl 定义：

```
// idl/hello/hello.proto
syntax = "proto3";

package hello;

option go_package = "hertz/hello";

import "api.proto";

message HelloReq {
  string Name = 1[(api.query)="name"];
}

message HelloResp {
  string RespBody = 1;
}

service HelloService {
  rpc Method1(HelloReq) returns(HelloResp) {
    option (api.get) = "/hello";
  }
}
```

2.  创建新项目

```
// GOPATH 下执行, 如果主IDL的依赖和主IDL不在同一路径下，需要加入 -I 选项，其含义为IDL搜索路径，等同于 protoc 的 -I 命令
hz new -I idl -idl idl/hello/hello.proto

// 整理 & 拉取依赖
go mod tidy
```

3.  修改handler，添加自己的逻辑

```
// handler path: biz/handler/hello/hello_service.go
// 其中 "/hello" 是 protobuf idl 中 go_package 的最后一级
// "hello_service.go" 是 protobuf idl 中 service 的名字，所有 service 定义的方法都会生成在这个文件中

// Method1 .
// @router /hello [GET]
func Method1(ctx context.Context, c *app.RequestContext) {
   var err error
   var req hello.HelloReq
   err = c.BindAndValidate(&req)
   if err != nil {
      c.String(400, err.Error())
      return
   }

   resp := new(hello.HelloResp)

   // 你可以修改整个函数的逻辑，而不仅仅局限于当前模板
   resp.RespBody = "hello," + req.Name // 添加的逻辑

   c.JSON(200, resp)
}
```

4.  编译项目

```
go build
```

5.  运行项目并测试

运行项目：

```
./{{your binary}}
```

测试：

```
curl --location --request GET 'http://127.0.0.1:8888/hello?name=hertz'
```

如果返回`{"RespBody":"hello,hertz"}`，说明接口调通。

### update: 更新一个已有的项目

1.  如果你的 protobuf idl 有更新，例如：

```
// idl/hello/hello.proto
syntax = "proto3";

package hello;

option go_package = "hertz/hello";

import "api.proto";

message HelloReq {
  string Name = 1[(api.query)="name"];
}

message HelloResp {
  string RespBody = 1;
}

message OtherReq {
  string Other = 1[(api.body)="other"];
}

message OtherResp {
  string Resp = 1;
}

service HelloService {
  rpc Method1(HelloReq) returns(HelloResp) {
    option (api.get) = "/hello";
  }
  rpc Method2(OtherReq) returns(OtherResp) {
    option (api.post) = "/other";
  }
}

service NewService {
  rpc Method3(OtherReq) returns(OtherResp) {
    option (api.get) = "/new";
  }
}
```

2.  切换到执行 new 命令的目录，更新修改后的 protobuf idl

```
hz update -I idl -idl idl/hello/hello.proto
```

3.  可以看到 "biz/handler/hello/hello_service.go" 下新增了新的方法，在 "biz/handler/hello" 下新增了文件 "new_service.go" 以及对应的 "Method3" 方法。

下面我们来开发 "Method2" 接口

```
// Method1 .
// @router /hello [GET]
func Method1(ctx context.Context, c *app.RequestContext) {
   var err error
   var req hello.HelloReq
   err = c.BindAndValidate(&req)
   if err != nil {
      c.String(400, err.Error())
      return
   }

   resp := new(hello.HelloResp)

   // 你可以修改整个函数的逻辑，而不仅仅局限于当前模板
   resp.RespBody = "hello," + req.Name // 添加的逻辑

   c.JSON(200, resp)
}

// Method2 .
// @router /other [POST]
func Method2(ctx context.Context, c *app.RequestContext) {
   var err error
   var req hello.OtherReq
   err = c.BindAndValidate(&req)
   if err != nil {
      c.String(400, err.Error())
      return
   }

   resp := new(hello.OtherResp)

   // 增加的逻辑
   resp.Resp = "Other method: " + req.Other

   c.JSON(200, resp)
}
```

4.  编译项目

```
go build
```

5.  运行项目并测试

运行项目：

```
./{{your binary}}
```

测试：

```
curl --location --request POST 'http://127.0.0.1:8888/other' \
--header 'Content-Type: application/json' \
--data-raw '{
    "Other": "other method"
}'
```

如果返回`{"Resp":"Other method: other method"}`，说明接口调通。