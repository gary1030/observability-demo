package main

import (
	"context"

	"github.com/gary1030/learning-o11y/server/config"
	"github.com/gary1030/learning-o11y/server/internal/handler"
	"github.com/gary1030/learning-o11y/server/internal/routes"
	"github.com/gary1030/learning-o11y/server/pkg/log"
	"github.com/gary1030/learning-o11y/server/pkg/otel"
	"github.com/gary1030/learning-o11y/server/pkg/prom"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	prom.Init()
	log.InitLogFormatter()
	config.Init()

	tp, err := otel.InitTracer()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Fatalf("Error shutting down tracer provider: %v", err)
		}
	}()

	appRouter := gin.New()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "traceparent", "uber-trace-id")
	appRouter.Use(cors.New(corsConfig))

	appRouter.Use(prom.GinPromMiddleware)
	appRouter.Use(gin.Recovery())
	appRouter.Use(otel.Register())
	appRouter.Use(gin.LoggerWithFormatter(log.GinLogFormatter))

	routes.SetTaskRoute(appRouter)
	routes.SetJokeRoute(appRouter)

	app := appRouter.Group("/")
	app.GET("/hello", handler.HelloHandler)
	app.GET("/ping", handler.PingHandler)

	go func() {
		if err := appRouter.Run(":" + config.ServicePort); err != nil {
			log.Fatalf("Failed to start application server: %v", err)
		}
	}()

	metricsRouter := gin.New()
	metricsRouter.Use(gin.Recovery())
	metricsRouter.GET("/metrics", prom.PromHandler(promhttp.Handler()))

	if err := metricsRouter.Run(":9999"); err != nil {
		log.Fatalf("Failed to start metrics server: %v", err)
	}
}
