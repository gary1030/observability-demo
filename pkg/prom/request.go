package prom

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	HttpReqTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "status"},
	)

	HttpReqInFlight = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "http_requests_in_flight",
			Help: "In flight number of HTTP requests",
		},
	)

	ReqDuration = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "request_duration_seconds",
			Help:    "HTTP request duration distribution in milliseconds",
			Buckets: []float64{0.1, 0.2, 0.5, 1, 2, 5, 10, 100, 1000},
		},
	)
)
