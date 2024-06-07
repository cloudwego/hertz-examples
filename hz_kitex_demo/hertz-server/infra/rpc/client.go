package rpc

import (
	"hertz-examples/hz_demo/hertz-server/kitex_gen/student/management/studentmanagement"

	client "github.com/cloudwego/kitex/client"
)

var Client studentmanagement.Client

func InitClient() {
	Client, _ = studentmanagement.NewClient("student", client.WithHostPorts("127.0.0.1:8888"))
}
