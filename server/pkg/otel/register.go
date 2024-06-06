package otel

import (
	"github.com/gary1030/learning-o11y/server/config"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func Register() gin.HandlerFunc {
	return otelgin.Middleware(config.ServiceName)
}
