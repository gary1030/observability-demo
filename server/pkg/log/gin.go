package log

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"
)

func GinLogFormatter(param gin.LogFormatterParams) string {
	ctx := param.Request.Context()
	span := trace.SpanFromContext(ctx)

	traceID := span.SpanContext().TraceID().String()

	log.WithFields(log.Fields{
		"status_code":   param.StatusCode,
		"latency":       param.Latency,
		"client_ip":     param.ClientIP,
		"method":        param.Method,
		"path":          param.Path,
		"request_proto": param.Request.Proto,
		"user_agent":    param.Request.UserAgent(),
		"error_message": param.ErrorMessage,
		"trace_id":      traceID,
	}).Info("[GIN INFO]")

	return ""
}
