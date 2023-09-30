package main

import (
	"fmt"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/hertz-examples/graphql/graph"
	"github.com/hertz-contrib/hertz-examples/graphql/graph/model"
	"github.com/hertz-contrib/pprof/adaptor"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const defaultPort = "8080"

var db *gorm.DB

func initDB() {
	var err error
	dataSourceName := "gorm:gorm@tcp(localhost:3306)/test_db?parseTime=True"
	db, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	// Migration to create tables for Order and Item schema
	db.AutoMigrate(&model.Order{}, &model.Item{})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	initDB()
	h := server.Default(server.WithHostPorts(":" + port))

	h.POST("/query", graphHandler())
	h.GET("/", playgroundHandler())
	h.Spin()
}

func graphHandler() app.HandlerFunc {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{DB: db}}))
	return adaptor.NewHertzHTTPHandlerFunc(srv.ServeHTTP)
}

func playgroundHandler() app.HandlerFunc {
	h := playground.Handler("GraphQL playground", "/query")
	return adaptor.NewHertzHTTPHandlerFunc(h)
}
