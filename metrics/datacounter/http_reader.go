package datacounter

import (
	"context"
	"io"
	"sync/atomic"

	"github.com/stellarentropy/gravity-assist-common/metrics/tracer"
)

// HTTPReaderCounter monitors the amount of data read from an [io.Reader] and
// accumulates it safely across multiple goroutines. It wraps an
// [io.ReadCloser], providing additional functionality to track and update
// metrics related to read operations. The type also offers a method to retrieve
// the total byte count in a thread-safe manner, ensuring data integrity when
// accessed concurrently.
type HTTPReaderCounter struct {
	ctx       context.Context
	count     uint64
	component string
	Reader    io.ReadCloser
}

// NewHTTPReaderCounter wraps an [io.ReadCloser] to monitor the amount of data
// being read, updating metrics for each read operation and providing a
// concurrent-safe total byte count. It returns a new instance of
// [HTTPReaderCounter].
func NewHTTPReaderCounter(ctx context.Context, component string, r io.ReadCloser) *HTTPReaderCounter {
	return &HTTPReaderCounter{
		ctx:       ctx,
		Reader:    r,
		component: component,
	}
}

// Read fetches data from the underlying [io.Reader] into a provided buffer and
// updates the count of bytes read. It ensures that only non-negative byte
// counts are recorded in the common byte counter and updates associated metrics
// for HTTPBytesRead. The method returns the number of bytes read and any error
// that may have occurred during the reading operation.
func (counter *HTTPReaderCounter) Read(buf []byte) (int, error) {
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
		err = tracer.AddInt64(counter.ctx, counter.component, "http.server.bytes.read", int64(n))
	}

	return n, err
}

// Count returns the total number of bytes read from the underlying [io.Reader]
// in a thread-safe manner.
func (counter *HTTPReaderCounter) Count() uint64 {
	return atomic.LoadUint64(&counter.count)
}

// Close terminates the underlying [io.Reader] of the [*HTTPReaderCounter], ends
// the tracking of metrics, and returns any error that occurred during the
// closing operation. If the underlying [io.Reader] does not implement the
// [io.Closer] interface, Close performs no operation and returns nil.
func (counter *HTTPReaderCounter) Close() error {
	return counter.Reader.Close()
}
