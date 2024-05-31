package prom

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	ResSizeBytes = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "response_size_bytes",
			Help:    "HTTP response size bytes distribution",
			Buckets: []float64{0, 1, 5, 10, 50, 100, 1000},
		},
	)
)
