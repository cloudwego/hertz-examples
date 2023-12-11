hz_install:
	go install github.com/cloudwego/hertz/cmd/hz@latest

kitex_install:
	go install github.com/cloudwego/kitex/tool/cmd/kitex@latest

mod_init:
	go mod init a/b/c

kitex_model:
	kitex --module=a/b/c ../thrift/hello.thrift
hz_model:
	hz model --mod=a/b/c --model_dir=kitex_gen -t=ignore_initialisms -t=gen_setter -t=gen_deep_equal -t=compatible_names -t=frugal_tag --idl=../thrift/hello.thrift

# 以上使用 hz_model 覆盖 kitex_model, 从而可以生成相同的结构体，也可以进而使用 hz 的 new/update 命令和 kitex 的 service 选项，创建服务代码。