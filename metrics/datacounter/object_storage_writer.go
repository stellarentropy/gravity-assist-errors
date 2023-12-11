package datacounter

import (
	"context"
	"sync/atomic"

	"github.com/stellarentropy/gravity-assist-common/metrics/tracer"

	"cloud.google.com/go/storage"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

// ObjectStorageWriterCounter aggregates and reports the amount of data sent to
// an object storage service, while also monitoring the performance of these
// write operations. It encapsulates the functionality to handle concurrent
// writes effectively, ensuring that the byte count is precise and consistent.
// Additionally, it offers mechanisms to conclude the writing session and
// finalize the collection of related metrics. Use this type when you need a
// reliable way to track both the volume of data written to object storage and
// the associated latency of these operations.
type ObjectStorageWriterCounter struct {
	ctx       context.Context
	count     uint64
	component string
	client    *storage.Client
	Writer    *storage.Writer
}

// NewObjectStorageWriterCounter returns a new instance of
// [ObjectStorageWriterCounter], which is responsible for monitoring and
// aggregating the volume of data written to an object storage service, as well
// as tracking the latency of write operations. It utilizes a provided
// [storage.Writer] and [storage.Client] to interface with the object storage
// system.
func NewObjectStorageWriterCounter(ctx context.Context, component string, w *storage.Writer, client *storage.Client) *ObjectStorageWriterCounter {
	return &ObjectStorageWriterCounter{
		Writer:    w,
		ctx:       ctx,
		component: component,
		client:    client,
	}
}

// Write writes a slice of bytes to the object storage, updates the count of
// successfully written bytes, and records the corresponding metrics. It returns
// the number of bytes written and any error that may have occurred during the
// write operation. Only positive byte counts are considered valid and included
// in the total count.
func (counter *ObjectStorageWriterCounter) Write(buf []byte) (int, error) {
	n, err := counter.Writer.Write(buf)

	// Write() should always return a non-negative `n`.
	// But since `n` is a signed integer, some custom
	// implementation of an io.Writer may return negative
	// values.
	//
	// Excluding such invalid values from counting,
	// thus `if n >= 0`:
	if n >= 0 {
		atomic.AddUint64(&counter.count, uint64(n))

		err = tracer.AddInt64(counter.ctx, counter.component, "object_storage.bytes.written", int64(n),
			metric.AddOption(metric.WithAttributes(
				attribute.KeyValue{
					Key:   "gcs.bucket.name",
					Value: attribute.StringValue(counter.Writer.Bucket),
				},
			)),
		)
	}

	return n, err
}

// Count retrieves the total number of bytes that have been successfully written
// to object storage using the [*ObjectStorageWriterCounter]. It ensures
// thread-safe access to the byte count and is suitable for use in concurrent
// operations. The method returns the cumulative byte count as a [uint64].
func (counter *ObjectStorageWriterCounter) Count() uint64 {
	return atomic.LoadUint64(&counter.count)
}

// Close finalizes the operation of the [*ObjectStorageWriterCounter],
// completing the metric tracking and closing the underlying writer if
// applicable. It returns any error that occurs during the closure of the
// underlying writer.
func (counter *ObjectStorageWriterCounter) Close() error {
	return counter.Writer.Close()
}
