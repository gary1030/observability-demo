package log

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GinLogFormatter(param gin.LogFormatterParams) string {
	log.WithFields(log.Fields{
		"status_code":   param.StatusCode,
		"latency":       param.Latency,
		"client_ip":     param.ClientIP,
		"method":        param.Method,
		"path":          param.Path,
		"request_proto": param.Request.Proto,
		"user_agent":    param.Request.UserAgent(),
		"error_message": param.ErrorMessage,
	}).Info("[GIN INFO]")

	return ""
}
