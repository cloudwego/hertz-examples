package main

import (
	management "hertz-examples/hz_demo/kitex-server/kitex_gen/student/management/studentmanagement"
	"log"
)

func main() {
	svr := management.NewServer(new(StudentManagementImpl))

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
