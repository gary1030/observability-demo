package main

import (
	"context"

	"github.com/gary1030/learning-o11y/config"
	"github.com/gary1030/learning-o11y/internal/handler"
	"github.com/gary1030/learning-o11y/internal/routes"
	"github.com/gary1030/learning-o11y/pkg/log"
	"github.com/gary1030/learning-o11y/pkg/otel"
	"github.com/gary1030/learning-o11y/pkg/prom"
	"github.com/gin-gonic/gin"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
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
	appRouter.Use(prom.GinPromMiddleware)
	appRouter.Use(gin.Recovery())
	appRouter.Use(otelgin.Middleware(config.ServiceName))
	appRouter.Use(gin.LoggerWithFormatter(log.GinLogFormatter))

	routes.SetTaskRoute(appRouter)

	app := appRouter.Group("/")
	app.GET("/hello", handler.HelloHandler)
	app.GET("/ping", handler.PingHandler)

	go func() {
		if err := appRouter.Run(":8081"); err != nil {
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
