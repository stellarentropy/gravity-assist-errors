package tracer

import (
	"context"
	"sync"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

var metrics = make(map[string]metric.Meter)

var counters = make(map[string]metric.Int64Counter)

var lock = sync.Mutex{}

func NewMetric(ctx context.Context, component string, opts ...metric.MeterOption) metric.Meter {
	lock.Lock()
	defer lock.Unlock()

	if m, ok := metrics[component]; ok {
		return m
	}

	m := otel.GetMeterProvider().Meter(component, opts...)

	metrics[component] = m

	return m
}

func AddInt64(ctx context.Context, component string, name string, value int64, opts ...metric.AddOption) error {
	lock.Lock()
	defer lock.Unlock()

	m := NewMetric(ctx, component)

	c, ok := counters[name]
	if !ok {
		counter, err := m.Int64Counter(name)
		if err != nil {
			return err
		}
		counters[name] = counter
		c = counter
	}

	c.Add(ctx, value, opts...)

	return nil
}
