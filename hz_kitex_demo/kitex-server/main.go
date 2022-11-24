package main

import (
	"log"

	management "hertz-examples/hz_demo/kitex-server/kitex_gen/student/management/studentmanagement"
)

func main() {
	svr := management.NewServer(new(StudentManagementImpl))

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
