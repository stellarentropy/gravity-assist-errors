package tracer

import (
	"context"
	"sync"

	config "github.com/stellarentropy/gravity-assist-common/config/common"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

var metrics = make(map[string]metric.Meter)

var counters = make(map[string]metric.Int64Counter)

var metricsLock = sync.Mutex{}
var countersLock = sync.Mutex{}

func NewMetric(ctx context.Context, component string, opts ...metric.MeterOption) metric.Meter {
	metricsLock.Lock()
	defer metricsLock.Unlock()

	if m, ok := metrics[component]; ok {
		return m
	}

	m := otel.GetMeterProvider().Meter(component, opts...)

	metrics[component] = m

	return m
}

func AddInt64(ctx context.Context, component string, name string, value int64, opts ...metric.AddOption) error {
	if !config.Common.EnableMetricCollection {
		return nil
	}

	m := NewMetric(ctx, component)

	countersLock.Lock()

	c, ok := counters[name]
	if !ok {
		counter, err := m.Int64Counter(name)
		if err != nil {
			countersLock.Unlock()
			return err
		}
		counters[name] = counter
		c = counter
	}
	countersLock.Unlock()

	c.Add(ctx, value, opts...)

	return nil
}

func MustAddInt64(ctx context.Context, component string, name string, value int64, opts ...metric.AddOption) {
	_ = AddInt64(ctx, component, name, value, opts...)
}
