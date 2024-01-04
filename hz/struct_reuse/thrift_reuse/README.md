## 使用 struct_reuse

### 准备工作

安装代码生成工具

```
go install github.com/cloudwego/hertz/cmd/hz@latest

go install github.com/cloudwego/kitex/tool/cmd/kitex@latest
```

初始化 go.mod

```
go mod init a/b/c
```

### 使用 kitex 生成 model
```
kitex --module=a/b/c ../thrift/hello.thrift
```

### 使用 hz 更新 model
```
hz model --mod=a/b/c --model_dir=kitex_gen -t=ignore_initialisms -t=gen_setter -t=gen_deep_equal -t=compatible_names -t=frugal_tag --idl=../thrift/hello.thrift
```
hz 需要为生成的结构体添加 tag，所以让 hz 去覆盖公共的 idl 文件(hello.thrift)，从而增加 tag

--model_dir 指定生成路径，-t 指定 thriftgo 参数，--idl 指定 thrift 文件路径