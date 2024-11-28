# hertz_swagger_gen

## 介绍

一个使用 `Hertz` 和 `thrift-gen-http-swagger` 的示例。

- 使用 `thrift` IDL 定义 `HTTP` 接口
- 使用 `hz` 生成代码
- 使用 `Gorm` and `MySQL`
- 使用 `thrift-gen-http-swagger` 插件生成 `swagger` 文件和 `swagger ui` 服务

- `/swagger` 提供 `swagger` 文件和 `swagger ui` 服务器
- `/handler` 包含更新用户、添加用户、删除用户、查询用户的基础业务逻辑

## IDL

该示例使用 `thrift` IDL 来定义 `HTTP` 接口。具体的接口定义在 [user.thrift](idl/user.thrift) 中。

## 代码生成工具

该示例使用 `hz` 来生成代码。`hz` 的使用可以参考 [hz](https://www.cloudwego.io/docs/hertz/tutorials/toolkit/toolkit/)。

使用的 `hz` 命令可以在 [Makefile](Makefile) 中找到。

## 绑定和验证

绑定和验证的使用可以参考[Binding and Validate](https://www.cloudwego.io/docs/hertz/tutorials/basic-feature/binding-and-validate/)

## 插件

`thrift-gen-http-swagger` 通过代码生成的`swagger`文档和`swagger ui`服务。

插件的使用可参考 [swagger-generate](https://github.com/hertz-contrib/swagger-generate)。

## 如何运行

### 运行 MySQL docker

```bash
cd bizdemo/hertz_swagger_gen && docker-compose up
```

### 运行示例

```bash
cd bizdemo/hertz_swagger_gen
go run .
```