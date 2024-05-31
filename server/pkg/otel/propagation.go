package otel

import (
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel/propagation"
)

// extract traceparent and tracestate from upstream
func PropagationExtractOption() otelgin.Option {
	tc := propagation.TraceContext{}
	return otelgin.WithPropagators(tc)
}
