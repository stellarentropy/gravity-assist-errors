package tracer

import (
	"context"

	config "github.com/stellarentropy/gravity-assist-common/config/common"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func NewSpan(ctx context.Context, name string, attributes ...attribute.KeyValue) (context.Context, trace.Span) {
	if !config.Common.EnableTraceCollection {
		return ctx, &MockTracer{}
	}

	ctx, span := otel.GetTracerProvider().Tracer("gravity-assist").Start(ctx, name)

	span.SetAttributes(attributes...)

	return ctx, span
}

func RecordError(span trace.Span, description string, err error) {
	if config.Common.EnableTraceCollection {
		span.RecordError(err)
		span.SetStatus(codes.Error, description)
	}
}
