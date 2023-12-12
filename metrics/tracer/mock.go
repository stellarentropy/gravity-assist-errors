package tracer

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type MockTracer struct {
	trace.Span
}

func (t *MockTracer) span() {}

func (t *MockTracer) End(...trace.SpanEndOption) {}

func (t *MockTracer) AddEvent(string, ...trace.EventOption) {}

func (t *MockTracer) IsRecording() bool { return false }

func (t *MockTracer) RecordError(error, ...trace.EventOption) {}

func (t *MockTracer) SpanContext() trace.SpanContext {
	return trace.NewSpanContext(trace.SpanContextConfig{})
}

func (t *MockTracer) SetStatus(codes.Code, string) {}

func (t *MockTracer) SetName(string) {}

func (t *MockTracer) SetAttributes(...attribute.KeyValue) {}

func (t *MockTracer) TracerProvider() trace.TracerProvider {
	return otel.GetTracerProvider()
}
