package otel

import (
	"context"
	"runtime"

	"github.com/gary1030/learning-o11y/server/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func StartNewSpan(ctx context.Context) (context.Context, trace.Span) {
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		ctx, span := otel.Tracer(config.ServiceName).Start(ctx, details.Name())
		return ctx, span
	}

	ctx, span := otel.Tracer("").Start(ctx, "")
	return ctx, span
}
