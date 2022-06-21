## 基于 thrift IDL 创建项目

### new: 创建一个新项目

1.  在当前目录下创建 thrift idl 文件

```
// idl/hello.thrift
namespace go hello.example

struct HelloReq {
    1: string Name (api.query="name");
}

struct HelloResp {
    1: string RespBody;
}


service HelloService {
    HelloResp HelloMethod(1: HelloReq request) (api.get="/hello");
}
```

2.  创建新项目

```
// GOPATH 下执行
hz new -idl idl/hello.thrift

// 整理 & 拉取依赖
go mod tidy
```

3.  修改handler，添加自己的逻辑

```
// handler path: biz/handler/hello/example/hello_service.go
// 其中 "hello/example" 是 thrift idl 的 namespace
// "hello_service.go" 是 thrift idl 中 service 的名字，所有 service 定义的方法都会生成在这个文件中

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

1.  如果你的 thrift idl 有更新，例如：

```
// idl/hello.thrift
namespace go hello.example

struct HelloReq {
    1: string Name (api.query="name");
}

struct HelloResp {
    1: string RespBody;
}

struct OtherReq {
    1: string Other (api.body="other");
}

struct OtherResp {
    1: string Resp;
}


service HelloService {
    HelloResp HelloMethod(1: HelloReq request) (api.get="/hello");
    OtherResp OtherMethod(1: OtherReq request) (api.post="/other");
}

service NewService {
    HelloResp NewMethod(1: HelloReq request) (api.get="/new");
}
```

2.  切换到执行 new 命令的目录，更新修改后的 thrift idl

```
hz update -idl idl/hello.thrift
```

3. 可以看到

   在 "biz/handler/hello/example/hello_service.go" 下新增了新的方法<br>
   在 "biz/handler/hello/example" 下新增了文件 "new_service.go" 以及对应的 "NewMethod" 方法。

下面我们来开发 "OtherMethod" 接口

```
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

   // 你可以修改整个函数的逻辑，而不仅仅局限于当前模板
   resp.RespBody = "hello," + req.Name // 添加的逻辑

   c.JSON(200, resp)
}

// OtherMethod .
// @router /other [POST]
func OtherMethod(ctx context.Context, c *app.RequestContext) {
   var err error
   // example.OtherReq 对应的model文件也会重新生成
   var req example.OtherReq
   err = c.BindAndValidate(&req)
   if err != nil {
      c.String(400, err.Error())
      return
   }

   resp := new(example.OtherResp)

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