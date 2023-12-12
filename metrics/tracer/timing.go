package tracer

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func NewSpan(ctx context.Context, name string, attributes ...attribute.KeyValue) (context.Context, trace.Span) {
	ctx, span := otel.GetTracerProvider().Tracer("gravity-assist").Start(ctx, name)

	span.SetAttributes(attributes...)

	return ctx, span
}

func RecordError(span trace.Span, description string, err error) {
	span.RecordError(err)
	span.SetStatus(codes.Error, description)
}
