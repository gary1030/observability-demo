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
			Help:    "HTTP request duration distribution in seconds",
			Buckets: []float64{1, 2, 5, 10, 100, 500, 1000},
		},
	)
)
