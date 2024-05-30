package otelresty

import (
	"github.com/gary1030/learning-o11y/config"
	"github.com/go-resty/resty/v2"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

func NewClient() *resty.Client {
	client := resty.New()
	otelMiddleware(client)
	return client
}

func otelMiddleware(client *resty.Client) {
	propagator := otel.GetTextMapPropagator()
	tracer := otel.Tracer(config.ServiceName)

	client.OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
		ctx := req.Context()
		spanCtx, span := tracer.Start(ctx, "HTTP "+req.Method, trace.WithSpanKind(trace.SpanKindClient))
		span.SetAttributes(attribute.String("endpoint", req.URL))
		propagator.Inject(spanCtx, propagation.HeaderCarrier(req.Header))
		req.SetContext(spanCtx)
		return nil
	})

	client.OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
		span := trace.SpanFromContext(resp.Request.Context())
		span.SetAttributes(attribute.Int("status_code", resp.StatusCode()))
		span.SetAttributes(attribute.Int64("response_size", resp.Size()))
		span.End()
		return nil
	})
}
