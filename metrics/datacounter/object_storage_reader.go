package datacounter

import (
	"bytes"
	"context"
	"io"
	"sync/atomic"

	"github.com/stellarentropy/gravity-assist-common/metrics/tracer"

	"cloud.google.com/go/storage"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

// ObjectStorageReaderCounter tracks the amount of data read from an object
// storage service, ensuring thread-safe tallying of bytes transferred during
// read operations. It encapsulates a reader for monitoring and recording the
// volume of data accessed, with optional support for random-access reading
// capabilities. The counter integrates with a metrics system to log read
// latencies and byte counts, providing valuable insights for monitoring object
// storage interactions.
type ObjectStorageReaderCounter struct {
	ctx        context.Context
	count      uint64
	component  string
	client     *storage.Client
	objHandler *storage.ObjectHandle
	seeker     bool
	seekReader *bytes.Reader
	Reader     *storage.Reader
}

// NewObjectStorageReaderCounter creates a new instance of
// [ObjectStorageReaderCounter] that wraps an existing [storage.Reader] to
// monitor and count the bytes read from an object in cloud storage. It also
// enables optional random-access read capabilities and integrates with a
// metrics system for monitoring read operations. The function accepts a
// [storage.Reader], a [storage.Client], a [storage.ObjectHandle], and a boolean
// indicating whether random-access reads should be enabled. It returns the
// newly created [ObjectStorageReaderCounter].
func NewObjectStorageReaderCounter(ctx context.Context, component string, r *storage.Reader, client *storage.Client, objHandler *storage.ObjectHandle, seeker bool) *ObjectStorageReaderCounter {
	reader := &ObjectStorageReaderCounter{
		Reader:     r,
		ctx:        ctx,
		component:  component,
		objHandler: objHandler,
		seeker:     seeker,
		client:     client,
	}

	if reader.seeker {
		buf := bytes.NewBuffer(nil)
		_, _ = io.Copy(buf, reader.Reader)

		reader.seekReader = bytes.NewReader(buf.Bytes())
	}

	return reader
}

// Read retrieves data from the underlying [io.Reader] into the provided buffer
// and updates the read byte count. It returns the number of bytes read along
// with any error encountered, ensuring that only non-negative byte counts are
// considered for updating the metrics. The function is concurrency-safe and
// integrates with metric tracking for object storage reads.
func (counter *ObjectStorageReaderCounter) Read(buf []byte) (int, error) {
	n, err := counter.Reader.Read(buf)

	// Read() should always return a non-negative `n`.
	// But since `n` is a signed integer, some custom
	// implementation of an io.Reader may return negative
	// values.
	//
	// Excluding such invalid values from counting,
	// thus `if n >= 0`:
	if n >= 0 {
		atomic.AddUint64(&counter.count, uint64(n))

		tracer.MustAddInt64(counter.ctx, counter.component, "object_storage.bytes.read", int64(n),
			metric.AddOption(metric.WithAttributes(
				attribute.KeyValue{
					Key:   "gcs.bucket.name",
					Value: attribute.StringValue(counter.objHandler.BucketName()),
				},
			)),
		)
	}

	return n, err
}

// Count returns the cumulative number of bytes read from the object storage by
// this instance. It ensures thread-safety, allowing for accurate byte count
// retrieval at any point during the object's read operations.
func (counter *ObjectStorageReaderCounter) Count() uint64 {
	return atomic.LoadUint64(&counter.count)
}

// Close terminates the underlying [io.Reader] and finalizes the tracking of
// read latency metrics. If the [io.Reader] implements the [io.Closer] interface
// and an error occurs during its closure, that error is returned. If the
// underlying [io.Reader] does not implement the [io.Closer] interface, Close
// performs no operation and returns nil.
func (counter *ObjectStorageReaderCounter) Close() error {
	return counter.Reader.Close()
}

// Size retrieves the size of the underlying object in the storage, expressed in
// bytes. This information is accessed through the attributes of the underlying
// [storage.Reader].
func (counter *ObjectStorageReaderCounter) Size() int64 {
	return counter.Reader.Attrs.Size
}

// ReadAt reads data from the storage object starting at a specified offset into
// the provided buffer. It updates the byte count and integrates with metric
// tracking, returning the number of bytes read and any error encountered. This
// method is concurrency-safe and supports concurrent access by multiple
// goroutines. If the instance is configured for random-access reads, it
// performs the operation using an internal reader; otherwise, it creates a new
// range reader for each call.
func (counter *ObjectStorageReaderCounter) ReadAt(buf []byte, off int64) (int, error) {
	var n int
	var err error

	if counter.seeker {
		n, err = counter.seekReader.ReadAt(buf, off)
	} else {
		r, err2 := counter.objHandler.NewRangeReader(context.Background(), off, int64(len(buf)))

		if err2 != nil {
			return 0, err2
		}

		n, err = r.Read(buf)
	}

	if n >= 0 {
		atomic.AddUint64(&counter.count, uint64(n))

		err = tracer.AddInt64(counter.ctx, counter.component, "object_storage.bytes.read", int64(n),
			metric.AddOption(metric.WithAttributes(
				attribute.KeyValue{
					Key:   "bucket",
					Value: attribute.StringValue(counter.objHandler.BucketName()),
				},
			)),
		)
	}

	return n, err
}
