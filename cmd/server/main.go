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

	appRouter := gin.Default()
	appRouter.Use(prom.GinPromMiddleware)

	app := appRouter.Group("/")
	app.GET("/hello", handler.HelloHandler)
	app.GET("/ping", handler.PingHandler)

	go func() {
		if err := appRouter.Run(":8080"); err != nil {
			log.Fatalf("Failed to start application server: %v", err)
		}
	}()

	metricsRouter := gin.Default()
	metricsRouter.GET("/metrics", prom.PromHandler(promhttp.Handler()))

	if err := metricsRouter.Run(":9090"); err != nil {
		log.Fatalf("Failed to start metrics server: %v", err)
	}
}
