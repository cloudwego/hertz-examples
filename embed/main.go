package main

import (
	"embed"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/pprof/adaptor"
)

//go:embed index.html
var index embed.FS

//go:embed test
var test embed.FS

func main() {
	srv := server.New()

	srv.GET("/", adaptor.NewHertzHTTPHandler(http.FileServer(http.FS(index))))
	srv.GET("/test/*filepath", adaptor.NewHertzHTTPHandler(http.FileServerFS(test)))

	srv.Run()
}
