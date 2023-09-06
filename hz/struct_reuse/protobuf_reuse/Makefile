hz_install:
	go install github.com/cloudwego/hertz/cmd/hz@latest

kitex_install:
	go install github.com/cloudwego/kitex/tool/cmd/kitex@latest

mod_init:
	go mod init a/b/c

kitex_model:
	kitex --module=a/b/c -I=../protobuf ../protobuf/hello/hello.proto

hz_model:
	hz model --mod=a/b/c --model_dir=kitex_gen -I=../protobuf --idl=../protobuf/hello/hello.proto

# 使用 hz_model 去覆盖 kitex_model，
# 注意：
# 1.所有 proto 中 go_package 的定义必须是这种形式 "{GOMOD}/kitex_gen/***" 其中gomod为项目的 go module 名
# 2. hz 的命令必须加 --model_dir=kitex_gen，才能正确覆盖 kitex_modle
# 3. 目前 hz无法生成 kitex 需要的接口信息，需要进一步对齐
