// Code generated by hertz generator. DO NOT EDIT.

package HzClient

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	get "hertz-examples/hz_client/server/biz/handler/get"
	post "hertz-examples/hz_client/server/biz/handler/post"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	root.POST("/body", append(_bodymethodMw(), post.BodyMethod)...)
	root.POST("/form", append(_formmethodMw(), post.FormMethod)...)
	root.POST("/path:path1", append(_pathmethodMw(), post.PathMethod)...)
	root.GET("/query", append(_querymethodMw(), get.QueryMethod)...)
}
