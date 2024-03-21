package main

import (
	"log"

	"github.com/gary1030/learning-o11y/internal/handler"
	"github.com/gary1030/learning-o11y/pkg/prom"
	"github.com/gin-gonic/gin"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	prom.Init()
	r := gin.Default()
	r.GET("/metrics", prom.PromHandler(promhttp.Handler()))

	app := r.Group("/")
	app.Use(prom.GinPromMiddleware)
	app.GET("/hello", handler.HelloHandler)
	app.GET("/ping", handler.PingHandler)

	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("impossible to start server: %s", err)
	}
}
