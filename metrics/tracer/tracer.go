package tracer

import (
	"context"
	"sync"
	"time"

	config "github.com/stellarentropy/gravity-assist-common/config/common"

	mexporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/metric"
	texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	"go.opentelemetry.io/contrib/detectors/gcp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

func StartTracer(ctx context.Context) (*sdktrace.TracerProvider, error) {
	gcpTraceExporter, err := texporter.New(texporter.WithProjectID(config.Common.GoogleProjectId))
	if err != nil {
		return nil, err
	}

	gcpMetricExporter, err := mexporter.New(mexporter.WithProjectID(config.Common.GoogleProjectId))
	if err != nil {
		return nil, err
	}

	res, err := resource.New(ctx,
		resource.WithDetectors(gcp.NewDetector()),
		resource.WithTelemetrySDK(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String(config.Common.ServiceName),
		),
	)
	if err != nil {
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(gcpTraceExporter),
		sdktrace.WithResource(res),
	)

	metricReader := sdkmetric.NewPeriodicReader(gcpMetricExporter, sdkmetric.WithInterval(10*time.Second))

	mp := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(metricReader),
		sdkmetric.WithResource(res),
	)

	otel.SetTracerProvider(tp)
	otel.SetMeterProvider(mp)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	return tp, nil
}

func Start(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	tp, err := StartTracer(ctx)
	if err != nil {
		logger.Error().Err(err).Msg("error starting tracer")
		panic(err)
	}

	defer func() {
		if err := tp.Shutdown(ctx); err != nil {
			logger.Error().Err(err).Msg("error shutting down tracer")
			panic(err)
		}
	}()

	<-ctx.Done()
}
