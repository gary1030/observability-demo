package otel

import (
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/contrib/propagators/jaeger"
	"go.opentelemetry.io/otel/propagation"
)

func PropagationExtractOption() otelgin.Option {
	tc := propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
		jaeger.Jaeger{},
	)
	return otelgin.WithPropagators(tc)
}
