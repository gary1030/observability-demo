package main

import (
	"github.com/gary1030/learning-o11y/internal/handler"
	"github.com/gary1030/learning-o11y/pkg/log"
	"github.com/gary1030/learning-o11y/pkg/prom"
	"github.com/gin-gonic/gin"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	prom.Init()
	log.InitLogFormatter()

	appRouter := gin.New()
	appRouter.Use(prom.GinPromMiddleware)
	appRouter.Use(gin.LoggerWithFormatter(log.GinLogFormatter))
	appRouter.Use(gin.Recovery())

	app := appRouter.Group("/")
	app.GET("/hello", handler.HelloHandler)
	app.GET("/ping", handler.PingHandler)

	go func() {
		if err := appRouter.Run(":8080"); err != nil {
			log.Fatalf("Failed to start application server: %v", err)
		}
	}()

	metricsRouter := gin.New()
	metricsRouter.Use(gin.LoggerWithFormatter(log.GinLogFormatter))
	metricsRouter.Use(gin.Recovery())
	metricsRouter.GET("/metrics", prom.PromHandler(promhttp.Handler()))

	if err := metricsRouter.Run(":9090"); err != nil {
		log.Fatalf("Failed to start metrics server: %v", err)
	}
}
