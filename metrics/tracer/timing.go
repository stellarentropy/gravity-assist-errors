package tracer

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func NewSpan(ctx context.Context, name string, attributes ...attribute.KeyValue) trace.Span {
	span := trace.SpanFromContext(ctx)
	span.SetName(name)

	span.SetAttributes(attributes...)

	return span
}
