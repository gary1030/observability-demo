package prom

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

func PromHandler(handler http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}

func Init() {
	prometheus.MustRegister(HttpReqTotal)
	prometheus.MustRegister(HttpReqInFlight)
	prometheus.MustRegister(ReqDuration)
	prometheus.MustRegister(ResSizeBytes)
}

func GinPromMiddleware(c *gin.Context) {
	start := time.Now()

	HttpReqInFlight.Inc()

	c.Next()

	// after request
	HttpReqTotal.With(prometheus.Labels{"method": c.Request.Method, "status": strconv.Itoa(c.Writer.Status())}).Inc()
	HttpReqInFlight.Dec()

	if responseSize := c.Writer.Size(); responseSize != -1 {
		ResSizeBytes.Observe(float64(responseSize))
	} else {
		ResSizeBytes.Observe(0)
	}

	elapsedTime := float64(time.Since(start)) / float64(time.Millisecond)
	ReqDuration.Observe(elapsedTime)
}
